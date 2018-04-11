package main

// Reads a Charles Schwab detailed CSV file and outputs a TXF file.
//
/*
	0		"Symbol"
	1		"Name"
	2		"Closed Date"
	3		"Opened Date"
	4		"Quantity"
	5		"Proceeds Per Share"
	6		"Cost Per Share"
	7		"Proceeds"
	8		"Cost Basis"
	9		"Gain/Loss ($)"
	10		"Gain/Loss (%)"
	11		"Long Term Gain/Loss"
	12		"Short Term Gain/Loss"
	13		"Term"
	14		"Unadjusted Cost Basis"
	15		"Wash Sale?"
	16		"Disallowed Loss"
			"Transaction Closed Date"
			"Transaction Cost Basis"
			"Total Transaction Gain/Loss ($)"
			"Total Transaction Gain/Loss (%)"
			"LT Transaction Gain/Loss ($)"
			"LT Transaction Gain/Loss (%)"
			"ST Transaction Gain/Loss ($)"
			"ST Transaction Gain/Loss (%)"
*/

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"./txf"
)

func main() {
	csvFile, _ := os.Open(os.Args[1])
	// Discard the first line...
	bufioreader := bufio.NewReader(csvFile)
	bufioreader.ReadLine()
	reader := csv.NewReader(bufioreader)

	out := txf.New()

	// Read the header line
	_, err := reader.Read()
	if err == io.EOF {
		return
	}

	// Read the data lines.
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line[15] != "No" {
			log.Fatal(line[15], " Wash sale not handled.")
		}

		symbol := line[0]
		quantity := line[4]
		dateAcquired := line[3]
		dateSold := line[2]
		proceeds := line[7]
		costBasis := line[8]
		ltGain := line[11]
		stGain := line[12]

		// What refNum do we want? If we're dealing with charles schwab, then they did not report
		// the transaction, so we probably want...
		/*
			Sales get reported on Copy A, B or C of the new Form 8949, and then flow to the Sch D.
			These new taxrefs will allow you to indicate which copy of Form 8949 a sale belongs on:
			Form 8949 Copy A : (you repored cost basis for this sale to the IRS using Form 1099B Box 3)
				321 (Short term holding period); 323 (long term holding period); 673 (you don't know the holding period); 682 (wash);

			Form 8949 Copy B : (you provide cost basis to customer, but you do NOT report it to the IRS using Form 1099B Box 3)
				711 (short term, Copy B); 713 (long term, Copy B); 715 (unknown holding period, Copy B); 718 (wash, Copy B)

			Form 8949 Copy C : (no 1099B issued customer or IRS)
				712 (short term, Copy C); 714 (long term, Copy C); 716 (unknown holding period, Copy C)
		*/
		refNum := txf.D_Short_Long_8949_Copy_B // Unknown holding period
		if ltGain != "" {
			refNum = txf.D_LT_gain_loss_8949_Copy_B
		}
		if stGain != "" {
			refNum = txf.D_ST_gain_loss_8949_Copy_B
		}

		r := txf.NewRecord5(
			refNum,
			fmt.Sprintf("%s %s", quantity, symbol),
			txf.ParseDate(dateAcquired),
			txf.ParseDate(dateSold),
			txf.ParseCurrency(costBasis),
			txf.ParseCurrency(proceeds),
			0)

		/*
			if txf.ParseCurrency(costBasis) > txf.ParseCurrency(proceeds) {
				r.X = []string{fmt.Sprintf("LOSS %s%s", ltGain, stGain)}
			}
		*/
		out.AddRecord(r)
	}

	out.Write(os.Stdout)
}
