package statutory_holiday

import (
	"log"
	"testing"
	"time"
)

func TestStatutoryHoliday_IsHoliday(t *testing.T) {

	isHoliday, err:= NewStatutoryHoliday(time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local)).IsHoliday()
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println(isHoliday)
}
