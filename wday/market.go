package wday

import "time"

func IsMarketOpen(currentTime time.Time) bool {
	c := Config()

	if c.MarketOpen == c.MarketClose {
		return true
	}

	currentInt := currentTime.Hour()*100 + currentTime.Minute()
	return c.MarketOpen < currentInt && currentInt > c.MarketClose
}
