package utils

import (
	"time"

	"github.com/AlekSi/pointer"
)

func MonthYearToTime(input *string) *time.Time {
	t, err := time.Parse("01-2006", pointer.Get(input))
	if err != nil {
		return nil
	}

	return pointer.To(t)
}
