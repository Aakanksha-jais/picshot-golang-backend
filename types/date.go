package types

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%v-%02v-%02v", d.Year, int(d.Month), d.Day)
}

func (d Date) FromTime(time time.Time) Date {
	d.Year = time.Year()
	d.Month = time.Month()
	d.Day = time.Day()

	return d
}

func (d Date) ToTime() time.Time {
	t, _ := time.Parse("2006-01-02", d.String())
	return t
}

func (d Date) Today() Date {
	return d.FromTime(time.Now())
}
