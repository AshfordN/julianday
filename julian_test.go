package julian

import (
	"testing"
	"time"
)

const TestJD = 2458520

func TestGregorianConversion(t *testing.T) {
	t.Run("from gregorian", func(t *testing.T) {
		jd := FromGregorian(2019, 2, 5)
		if jd != TestJD {
			t.Fail()
		}
	})

	t.Run("from time", func(t *testing.T) {
		tmp := time.Date(2019, 2, 5, 12, 0, 0, 0, time.UTC)
		jd := FromTime(tmp)
		if jd != TestJD {
			t.Fail()
		}
	})
}

func TestJulianConversion(t *testing.T) {
	jd := Day(TestJD)

	t.Run("to gregorian", func(t *testing.T) {
		y, m, d := jd.ToGregorian()
		if y != 2019 || m != 2 || d != 5 {
			t.Fail()
		}
	})

	t.Run("to time", func(t *testing.T) {
		tmp := jd.ToTime(time.UTC)
		if tmp.Year() != 2019 || tmp.Month() != time.February || tmp.Day() != 5 || tmp.Hour() != 12 {
			t.Fail()
		}
	})
}

func TestWeekDay(t *testing.T) {
	jd := Day(TestJD)
	if jd.WeekDay() != 2 { //tuesday
		t.Fail()
	}
}

func TestArithmetic(t *testing.T) {
	jd := Day(TestJD)

	t.Run("retreat", func(t *testing.T) {
		if jd.Sub(5).Int() != TestJD-5 {
			t.Fail()
		}
	})

	t.Run("advance", func(t *testing.T) {
		if jd.Add(5).Int() != TestJD+5 {
			t.Fail()
		}
	})
}
