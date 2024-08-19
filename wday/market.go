package wday

import "time"

func IsMarketOpenNow() bool {
	return IsMarketOpen(time.Now())
}

func IsMarketOpen(currentTime time.Time) bool {
	c := Config()

	if c.MarketOpen == c.MarketClose {
		return true
	}

	currentInt := currentTime.Hour()*100 + currentTime.Minute()
	return currentInt >= c.MarketOpen && currentInt < c.MarketClose
}
