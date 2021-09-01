package statutory_holiday

import "time"

type StatutoryHolidayEnum int

const (
	Unknown StatutoryHolidayEnum = iota
	NewYear
	SpringFestival
	TombSweepingDay
	LabourDay
	DragonBoatFestival
	NationalDay
	MidAutumnFestival
)

var HolidayNameMap = map[StatutoryHolidayEnum]string{
	NewYear:            "元旦",
	SpringFestival:     "春节",
	TombSweepingDay:    "清明节",
	LabourDay:          "劳动节",
	DragonBoatFestival: "端午节",
	NationalDay:        "国庆节",
	MidAutumnFestival:  "中秋节",
}

type StatutoryHolidayDefine struct {
	HolidayTime  []time.Time
	HolidayName  StatutoryHolidayEnum
	HolidayReset []time.Time
}

var Holidays = map[int][]StatutoryHolidayDefine{
	2021: {
		{
			HolidayTime: []time.Time{
				time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
				time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
				time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			},
			HolidayName:  NewYear,
			HolidayReset: nil,
		},
	},
}
