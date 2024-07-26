package wday

import "github.com/cornelk/hashmap"

type Empty struct{}

var (
	void       = Empty{}
	holidaysCO = hashmap.NewSized[uint16, Empty](18)
	holidaysCL = hashmap.NewSized[uint16, Empty](13)
	holidaysPE = hashmap.NewSized[uint16, Empty](13)
)

func initCO() {
	holidaysCO.Set(101, void)
	holidaysCO.Set(108, void)
	holidaysCO.Set(325, void)
	holidaysCO.Set(328, void)
	holidaysCO.Set(329, void)
	holidaysCO.Set(501, void)
	holidaysCO.Set(513, void)
	holidaysCO.Set(603, void)
	holidaysCO.Set(610, void)
	holidaysCO.Set(701, void)
	holidaysCO.Set(719, void)
	holidaysCO.Set(807, void)
	holidaysCO.Set(819, void)
	holidaysCO.Set(1014, void)
	holidaysCO.Set(1104, void)
	holidaysCO.Set(1111, void)
	holidaysCO.Set(1208, void)
	holidaysCO.Set(1225, void)
}

func initCL() {
	holidaysCL.Set(101, void)
	holidaysCL.Set(329, void)
	holidaysCL.Set(501, void)
	holidaysCL.Set(521, void)
	holidaysCL.Set(620, void)
	holidaysCL.Set(716, void)
	holidaysCL.Set(815, void)
	holidaysCL.Set(918, void)
	holidaysCL.Set(919, void)
	holidaysCL.Set(920, void)
	holidaysCL.Set(1031, void)
	holidaysCL.Set(1101, void)
	holidaysCL.Set(1225, void)
}

func initPE() {
	holidaysPE.Set(101, void)
	holidaysPE.Set(328, void)
	holidaysPE.Set(329, void)
	holidaysPE.Set(501, void)
	holidaysPE.Set(607, void)
	holidaysPE.Set(723, void)
	holidaysPE.Set(729, void)
	holidaysPE.Set(806, void)
	holidaysPE.Set(830, void)
	holidaysPE.Set(1008, void)
	holidaysPE.Set(1101, void)
	holidaysPE.Set(1209, void)
	holidaysPE.Set(1225, void)
}
