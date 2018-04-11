package txf

// A quick-and-dirty TXF file format reader / writer.
// https://turbotax.intuit.com/txf/TXF042.jsp

import (
	"fmt"
	"io"
	"time"
)

/*
All Tax Exchange Files should have a header which defines what version
of this spec was used as well as what program was used to create
the import file.

The fields for the Header of the file are:

  V version
  A accounting program name/version
  D export date

For each exchange record there are a set of common fields. This is the
recommended order:

  T type
  N refnum
  C copy
  L line
  X detail

There are 7 record formats, labeled 0 through 6 with various
additional fields.

The fields for these record formats are:

Record Format 0
  Note: used for boolean records such as "spouse" indicator

Record Format 1 (placed before the 'X' detail line in the record)
  $ amount

Record Format 2
  P Text value

  Note: String value used for "description" fields such as Sched C Line A

Record Format 3
  $ amount
  P description

  Note: used for fields such as interest where bank name/account#
        and $ amount required.

Record Format 4
  P security
  D date acquired
  D date sold
  $ cost basis
  $ sales net (net of commission)

Record Format 5
  P security
  D date acquired
  D date sold
  $ cost basis
  $ sales net
  $ Depreciation Allowed OR Disallowed wash sale amount (added for TY11)

  Note: expanded Format 5 to use for new TY11 cost basis reporting rules. Format 5 will
   be used with Taxrefs 321,323,673,682,711,712,713,714,715,716 when reporting a disallowed wash sale amount.

Record Format 6
  D date paid
  $ amount paid
  P state initials

  Note: used for quarterly federal and state estimated tax payments; state
        initials ignored for federal estimates.


B. Codes and Definition of Fields

"V" Version
  Indicates which version of the tax line item maintenance file was used to
  export this file. The current number can be found at the top of this file.

"A" Accounting Program Name/Version
  Indicates which program (including version) exported this file.

"D" Export Date
  Date on which this file was exported.

"T" Type
  Either S for summary or D for detail.  Default value is S.

"N" Refnum
  The tax category refnum for this item.  There is not default value (it
  must be supplied).  This refnum should come from the tax line item
  maintenance file.

"C" Copy
  Integer between 1 and 255.  Default value is 1.  This is the copy number
  for multi-copy forms (such as Schedule C.)  If there is only one copy, the
  value should be 1.

"L" Line
  A positive integer.  The default value is 1.  This is the line number for
  multi-line items, such as interest income on Schedule B.  If there is only
  one line item, this number should be 1.

"X" Detail
  Extra text that can be used by the tax software, if desired.  The initial
  intent of this field is for use in a supporting statement.

"$" Amount
  This is the dollar amount for the item.  Income, gains, and money received
  are positive numbers.  Expenses, losses, and money spent (including tax
  payments) are negative numbers.

"P" Description
  A string describing this particular line.  This is the value that is
  different and describes the line and therefore should appear on the line
  on the tax form.

"P" Security
  Name of security for schedule D.

"D" Date Acquired
  Acquisition date.  It should be in the form MM/DD/YYYY.

"D" Date Sold
  Sale date.  It should be in the form MM/DD/YYYY.

"D" Date Paid
  Payment date.  It should be in the form MM/DD/YYYY.


Default means that a line does not need to be supplied not that if the
field is supplied with no value the "default" will be used.  So if
NO "C" field is supplied then the default of Copy 1 is to be used.

*/
type Record struct {
	T bool     // type true == detail
	N RefNum   // RefNum 0-999
	C int      // copy
	L int      // line
	X []string // detail
	// 1,2,3,6
	DatePaid time.Time // D
	Amount   Dollars   // $
	Desc     string    // P  2...

	// 4,5
	DateAcquired time.Time // D
	DateSold     time.Time // D
	CostBasis    Dollars   //
	SalesNet     Dollars   //
	WashSale     Dollars   // .. or depreciation
}

func (r *Record) Write(w io.Writer) {
	if r.T {
		w.Write([]byte("TD\n"))
	} else {
		w.Write([]byte("TS\n"))
	}
	r.N.Write(w)
	if r.C > 0 {
		fmt.Fprintf(w, "C%d\n", r.C)
	}
	if r.L > 0 {
		fmt.Fprintf(w, "L%d\n", r.C)
	}

	// 1, 2, 3, 6
	if !r.DatePaid.IsZero() {
		w.Write([]byte(r.DatePaid.Format(outDateFormat)))
	}
	if r.Amount > 0 {
		r.Amount.Write(w)
	}
	if r.Desc != "" {
		fmt.Fprintf(w, "P%s\n", r.Desc)
	}

	// 4,5
	if !r.DateAcquired.IsZero() {
		w.Write([]byte(r.DateAcquired.Format(outDateFormat)))
	}
	if !r.DateSold.IsZero() {
		w.Write([]byte(r.DateSold.Format(outDateFormat)))
	}
	if r.CostBasis > 0 {
		r.CostBasis.Write(w)
	}
	if r.SalesNet > 0 {
		r.SalesNet.Write(w)
	}
	if r.WashSale > 0 {
		r.WashSale.Write(w)
	}
	for i := range r.X {
		fmt.Fprintf(w, "X%s\n", r.X[i])
	}
	w.Write([]byte("^\n"))
}

func NewRecord0(RefNum RefNum) *Record {
	return &Record{T: false, N: RefNum, C: 1, L: 1}
}

func NewRecord1(RefNum RefNum, Amount Dollars) *Record {
	return &Record{T: false, N: RefNum, C: 1, L: 1, Amount: Amount}
}

func NewRecord2(RefNum RefNum, Desc string) *Record {
	return &Record{T: false, N: RefNum, C: 1, L: 1, Desc: Desc}
}

func NewRecord3(RefNum RefNum, Desc string, Amount Dollars) *Record {
	return &Record{T: false, N: RefNum, C: 1, L: 1, Amount: Amount, Desc: Desc}
}

func NewRecord4(
	RefNum RefNum,
	Security string,
	DateAcquired time.Time,
	DateSold time.Time,
	CostBasis Dollars,
	SalesNet Dollars) *Record {
	return &Record{T: true, N: RefNum, C: 1, L: 1, Desc: Security,
		DateAcquired: DateAcquired,
		DateSold:     DateSold,
		CostBasis:    CostBasis,
		SalesNet:     SalesNet,
	}
}

func NewRecord5(
	RefNum RefNum,
	Security string,
	DateAcquired time.Time,
	DateSold time.Time,
	CostBasis Dollars,
	SalesNet Dollars,
	WashSale Dollars) *Record {
	return &Record{T: true, N: RefNum, C: 1, L: 1, Desc: Security,
		DateAcquired: DateAcquired,
		DateSold:     DateSold,
		CostBasis:    CostBasis,
		SalesNet:     SalesNet,
		WashSale:     WashSale,
	}
}

func NewRecord6(
	RefNum RefNum,
	DatePaid time.Time,
	Amount Dollars,
	State string) *Record {
	return &Record{T: false, N: RefNum, C: 1, L: 1, Amount: Amount, Desc: State, DatePaid: DatePaid}
}

// TXF holds the individual records in a TXF file.
type TXF struct {
	V    int
	A    string
	D    time.Time
	data []*Record
}

// Returns a new TXF object
func New() *TXF {
	return &TXF{V: 42, A: "GoTXF", D: time.Now()}
}

func (txf *TXF) writeHeader(w io.Writer) {
	fmt.Fprintf(w, "V%03d\nA%s\n%s^\n", txf.V, txf.A, txf.D.Format(outDateFormat))
}

func (txf *TXF) Write(w io.Writer) {
	txf.writeHeader(w)
	for i := range txf.data {
		txf.data[i].Write(w)
	}
}

// Adds a txf record to the file.
func (txf *TXF) AddRecord(r *Record) {
	txf.data = append(txf.data, r)
}
