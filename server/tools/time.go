package tools

import (
	"time"

	"github.com/VinceZCL/FinalYearProject/app/config"
)

func GetTimes(dateStr string) (start, end time.Time, err error) {

	loc, err := time.LoadLocation(config.Get().Database.Location)
	if err != nil {
		return
	}

	day := time.Now().In(loc)
	if dateStr != "" {
		day, err = time.ParseInLocation(time.DateOnly, dateStr, loc)
		if err != nil {
			return
		}
	}
	start = time.Date(
		day.Year(),
		day.Month(),
		day.Day(),
		0, 0, 0, 0,
		loc,
	).UTC()
	end = start.Add(24 * time.Hour)
	return
}
