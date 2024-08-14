package wday

import (
	"github.com/alphadose/haxmap"
)

func getHolidays(country string) *haxmap.Map[uint16, Empty] {
	switch country {
	case "CO":
		return holidaysCO
	case "PE":
		return holidaysPE
	case "CL":
		return holidaysCL
	}
	return nil
}

func init() {
	countries := Config().Countries
	if countries == nil {
		go initCO()
		go initCL()
		go initPE()
		return
	}

	for _, country := range countries {
		switch country {
		case "CO":
			go initCO()
		case "PE":
			go initPE()
		case "CL":
			go initCL()
		}
	}

}
