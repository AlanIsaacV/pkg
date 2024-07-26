package wday

import (
	"time"

	"github.com/rs/zerolog/log"
)

func IsWorkDayToday(country string) bool {
	return IsWorkDay(country, time.Now())
}

func IsWorkDay(country string, day time.Time) bool {
	_, m, d := day.Date()
	return isWorkDay(country, m, d, day.Weekday())
}

func isWorkDay(country string, month time.Month, day int, weekday time.Weekday) bool {
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	holidays := getHolidays(country)
	if holidays == nil {
		log.Warn().Msgf("Country %s not found", country)
		return true
	}

	holiday := uint16(int(month)*100 + day)
	_, exists := holidays.Get(holiday)
	return !exists
}

func NextWorkDay(country string) time.Time {
	return GetWorkDay(country, time.Now(), 1, true)
}

func PrevWorkDay(country string) time.Time {
	return GetWorkDay(country, time.Now(), 1, false)
}

func NextWorkDayOffset(country string, offset uint8) time.Time {
	return GetWorkDay(country, time.Now(), offset, true)
}

func PrevWorkDayOffset(country string, offset uint8) time.Time {
	return GetWorkDay(country, time.Now(), offset, false)
}

func GetWorkDay(country string, day time.Time, offset uint8, nextDay bool) time.Time {
	var delta int
	if nextDay {
		delta = 1
	} else {
		delta = -1
	}

	if !IsWorkDay(country, day) {
		offset += 1
	}
	for offset > 0 {
		day = day.AddDate(0, 0, delta)
		if IsWorkDay(country, day) {
			offset -= 1
		}
	}

	return day
}
