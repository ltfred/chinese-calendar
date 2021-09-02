package statutory_holiday

import (
	"time"

	"github.com/pkg/errors"
)

type (
	StatutoryHoliday interface {
		IsHoliday() (bool, error)
		GetHolidayDetail() (bool, HolidayDetailResp, error)
		GetCurrentDateTheDayOfHoliday() (bool, CurrentDateTheDayOfHolidayResp, error)
	}

	Holiday struct {
		Day time.Time
	}

	HolidayDetailResp struct {
		HolidayName string
		HolidayTime []time.Time
		HolidayRest []time.Time
	}

	CurrentDateTheDayOfHolidayResp struct {
		HolidayName        string
		HolidayFirstFewDay int
	}
)

func NewStatutoryHoliday(t time.Time) StatutoryHoliday {
	return &Holiday{
		Day: t,
	}
}

func (s *Holiday) IsHoliday() (bool, error) {
	holidays, newDate, err := s.getHolidays()
	if err != nil {
		return false, err
	}

	for _, v := range holidays {
		for _, vv := range v.HolidayTime {
			if newDate == vv {
				return true, nil
			}
		}
	}

	return false, nil
}

func (s *Holiday) GetHolidayDetail() (bool, HolidayDetailResp, error) {
	holidays, newDate, err := s.getHolidays()
	if err != nil {
		return false, HolidayDetailResp{}, err
	}

	for _, v := range holidays {
		for _, vv := range v.HolidayTime {
			if newDate == vv {
				return true, HolidayDetailResp{
					HolidayName: v.HolidayName.Label(),
					HolidayTime: v.HolidayTime,
					HolidayRest: v.HolidayReset,
				}, nil
			}
		}
	}

	return false, HolidayDetailResp{}, nil
}

func (s *Holiday) GetCurrentDateTheDayOfHoliday() (bool, CurrentDateTheDayOfHolidayResp, error) {
	holidays, newDate, err := s.getHolidays()
	if err != nil {
		return false, CurrentDateTheDayOfHolidayResp{}, err
	}

	for _, v := range holidays {
		for i, vv := range v.HolidayTime {
			if newDate == vv {
				return true, CurrentDateTheDayOfHolidayResp{
					HolidayName:        v.HolidayName.Label(),
					HolidayFirstFewDay: i + 1,
				}, nil
			}
		}
	}

	return false, CurrentDateTheDayOfHolidayResp{}, nil
}

func (s *Holiday) getHolidays() ([]StatutoryHolidayDefine, time.Time, error) {
	year, mon, day := s.Day.Date()
	newDate := time.Date(year, mon, day, 0, 0, 0, 0, time.Local)
	holidays, ok := Holidays[year]
	if !ok {
		return nil, newDate, errors.New("not support the holidays of the year")
	}

	return holidays, newDate, nil
}
