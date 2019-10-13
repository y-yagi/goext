package timeext_test

import (
	"testing"
	"time"

	"github.com/y-yagi/goext/timeext"
)

func TestBeginningOfMonth(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2016-01-10")
	firstDate := timeext.BeginningOfMonth(date)

	if (firstDate.Year() != 2016) || (firstDate.Month() != 1) || (firstDate.Day() != 1) {
		t.Errorf("Expect beginningOfmonth is 2016-01-01, but %s", firstDate)
	}
}

func TestEndOfMonth(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2016-01-10")
	endDate := timeext.EndOfMonth(date)

	if (endDate.Year() != 2016) || (endDate.Month() != 1) || (endDate.Day() != 31) {
		t.Errorf("Expect endOfMonth is 2016-01-31, but %s", endDate)
	}

	date, _ = time.Parse("2006-01-02", "2016-02-02")
	endDate = timeext.EndOfMonth(date)

	if (endDate.Year() != 2016) || (endDate.Month() != 2) || (endDate.Day() != 29) {
		t.Errorf("Expect endOfMonth is 2016-02-29, but %s", endDate)
	}
}
