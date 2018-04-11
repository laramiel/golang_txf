# golang_txf

This is a quick and dirty hack to convert Charles Schwab detailed
transaction CSV files to TXF files.

In the spirit of 

* https://github.com/jacoblyw/csv-to-txf-converter
* https://github.com/ntulpule/Schwab-1099B-Parser

To use it, clone the repo and...

`go run main.go  Detailed.CSV`


## HOWTO

#### Step 1. Get detailed CSV from Charles Schwab

1. Go to the Account *History* Page
2. Select the *Realized Gain / Loss* tab
3. Choose the date range. Typically 1 year. 01/01/2016 - 12/31/2016
4. Select the *Export* link in the upper right.
5. Download the detailed report.

#### Step 2. Run the program

`go run main.go Detailed.CSV`


