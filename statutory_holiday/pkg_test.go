package statutory_holiday

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHoliday_IsHoliday(t *testing.T) {
	isHoliday, err := NewStatutoryHoliday(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local)).IsHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}
	assert.Equal(t, true, isHoliday)

	isHoliday, err = NewStatutoryHoliday(time.Date(2021, 1, 5, 0, 0, 0, 0, time.Local)).IsHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}
	assert.Equal(t, false, isHoliday)
}

func TestHoliday_GetCurrentDateTheDayOfHoliday(t *testing.T) {
	IsHoliday, result, err := NewStatutoryHoliday(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)).GetCurrentDateTheDayOfHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}

	assert.Equal(t, true, IsHoliday)
	assert.Equal(t, 1, result.HolidayFirstFewDay)

	IsHoliday, result, err = NewStatutoryHoliday(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local)).GetCurrentDateTheDayOfHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}

	assert.Equal(t, true, IsHoliday)
	assert.Equal(t, 2, result.HolidayFirstFewDay)

	IsHoliday, result, err = NewStatutoryHoliday(time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local)).GetCurrentDateTheDayOfHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}

	assert.Equal(t, false, IsHoliday)
	assert.Equal(t, 0, result.HolidayFirstFewDay)
}

func TestHoliday_GetHolidayDetail(t *testing.T) {
	isHoliday, holidayDetail, err := NewStatutoryHoliday(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)).GetHolidayDetail()
	if err != nil {
		log.Fatalf("%v", err)
	}

	assert.Equal(t, true, isHoliday)
	assert.Equal(t, NewYear.Label(), holidayDetail.HolidayName)
	assert.Equal(t, []time.Time{}, holidayDetail.HolidayRest)
}