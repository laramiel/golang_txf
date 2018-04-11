package txf

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	// DateFormat formats a date as required by TXF.
	DateFormat    = "01/02/2006"
	outDateFormat = "D01/02/2006\n"
)

// ParseDate prases a TXF date
func ParseDate(in string) time.Time {
	t, err := time.Parse(DateFormat, in)
	if err != nil {
		panic(err)
	}
	return t
}

// Dollars is used to represent money, in cents.
type Dollars int

func (d Dollars) Write(w io.Writer) {
	a := int64(d) / 100
	b := int64(d) % 100
	fmt.Fprintf(w, "$%d.%02d\n", a, b)
}

func parseCurrencyImpl(in string) Dollars {
	i, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		panic(err)
	}
	return Dollars(i)
}

// Parse a currency string
func ParseCurrency(in string) Dollars {
	if in == "" || in == "0" {
		return 0
	}
	if in[0] == '-' {
		return -ParseCurrency(in[1:])
	} else if in[0] == '(' && in[len(in)-1] == ')' {
		return -ParseCurrency(in[1 : len(in)-1])
	}
	if in[0] == '$' {
		return ParseCurrency(in[1:])
	}

	j := strings.LastIndexByte(in, '.')
	if j == -1 {
		return parseCurrencyImpl(in + "00")
	}
	return parseCurrencyImpl(in[0:j] + in[j+1:len(in)])
}
