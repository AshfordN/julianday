package julian

import (
	"fmt"
	"time"
)

type Julian int

//Int returns the integer representation of the julian day
func (jd *Julian) Int() int {
	return int(*jd)
}

//Format returns the string representation of the julian day
func (jd *Julian) Format() string {
	y, m, d := jd.ToGregorian()
	return fmt.Sprintf("%d-%d-%d", d, m, y)
}

//FromTime returns the julian day equivalent of the time object
func FromTime(t time.Time) Julian {
	y, month, d := t.Date()
	m := int(month) //convert month to int
	return FromGregorian(y, m, d)
}

//FromGregorian returns the julian day equivalent from the gregorian day representation
//International Amateur-Professional Photoelectric Photometry Communication, No. 13, p.16
//http://adsabs.harvard.edu/full/1983IAPPP..13...16F
func FromGregorian(y, m, d int) Julian {
	return Julian(367*y - 7*(y+(m+9)/12)/4 - 3*((y+(m-9)/7)/100+1)/4 + (275*m)/9 + d + 1721029)
}

//WeekDay returns the week day of the julian day
//where 0 = Sunday and 6 = Saturday
func (jd *Julian) WeekDay() int {
	return (jd.Int() + 1) % 7
}

//ToGregorian returns the year, month and day corresponding to the julian date
func (jd *Julian) ToGregorian() (int, int, int) {
	l := jd.Int() + 68569
	n := 4 * l / 146097
	l = l - (146097*n+3)/4
	y := 4000 * (l + 1) / 1461001
	l = l - 1461*y/4 + 31
	m := 80 * l / 2447
	d := l - 2447*m/80
	l = m / 11
	m = m + 2 - 12*l
	y = 100*(n-49) + y + l
	return y, m, d
}

//ToTime returns the time object equivalent of the julian day
func (jd *Julian) ToTime(loc *time.Location) time.Time {
	y, m, d := jd.ToGregorian()
	return time.Date(y, time.Month(m), d, 12, 0, 0, 0, loc)
}
