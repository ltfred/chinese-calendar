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

func (v StatutoryHolidayEnum) Label() string {
	switch v {
	case Unknown:
		return ""
	case NewYear:
		return "元旦"
	case SpringFestival:
		return "春节"
	case TombSweepingDay:
		return "劳动节"
	case LabourDay:
		return "清明节"
	case DragonBoatFestival:
		return "端午节"
	case NationalDay:
		return "国庆节"
	case MidAutumnFestival:
		return "中秋节"
	default:
		return ""
	}
}

type StatutoryHolidayDefine struct {
	HolidayName  StatutoryHolidayEnum
	HolidayTime  []time.Time
	HolidayReset []time.Time
}

var Holidays = map[int][]StatutoryHolidayDefine{
	2021: {
		{
			HolidayName: NewYear,
			HolidayTime: []time.Time{
				time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
				time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
				time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			},
			HolidayReset: []time.Time{},
		},
		{
			HolidayName: SpringFestival,
			HolidayTime: []time.Time{
				time.Date(2021, 2, 11, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 13, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 14, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 15, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 16, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 17, 0, 0, 0, 0, time.Local),
			},
			HolidayReset: []time.Time{
				time.Date(2021, 2, 7, 0, 0, 0, 0, time.Local),
				time.Date(2021, 2, 20, 0, 0, 0, 0, time.Local),
			},
		},
	},
}
