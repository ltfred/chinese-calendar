package solar_terms

import (
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type SolarTerm interface {
	IsSolarTerm() (bool, string, error)
}

type Solar struct {
	t time.Time
}

func NewSolarTerm(t time.Time) SolarTerm {
	return &Solar{t: t}
}

func (s *Solar) IsSolarTerm() (bool, string, error) {

	date := time.Date(s.t.Year(), s.t.Month(), s.t.Day(), 0,0,0,0, time.Local)

	terms, err := s.getCurrentYearSolarTerms()
	if err != nil {
		return false, "", err
	}

	for k, v := range terms {
		if date == v {
			return true, k.Label(), nil
		}
	}

	return false, "", nil

}

func (s *Solar) getCurrentYearSolarTerms() (map[SolarTermsEnum]time.Time, error) {
	yearStr := strconv.Itoa(s.t.Year())
	s2 := string([]byte(yearStr)[:2])
	pre, err := strconv.Atoi(s2)
	if err != nil {
		return nil, err
	}

	century := pre + 1

	if century > 21 || century < 20 {
		return nil, errors.New("no support the year")
	}

	solarTerms := make(map[SolarTermsEnum]time.Time, 0)

	for k, v := range solarTermCValueMap {
		n := float64(s.t.Year()-(pre*100))*0.2422 + v[century].SolarTermsCValue
		var l int
		switch k {
		case SlightCold, GreatCold, BeginningOfSpring, RainWater:
			l = ((s.t.Year() - (pre * 100)) - 1) / 4
		default:
			l = (s.t.Year() - (pre * 100)) / 4
		}

		day := int(n) - l

		for _, y := range v[century].SolarTermsSpecialYear {
			if s.t.Year() == y {
				day = day + v[century].SolarTermsSpecial
			}
		}

		solarTerms[k] = time.Date(s.t.Year(), v[century].SolarTermsMonth, day, 0, 0, 0, 0, time.Local)

	}

	return solarTerms, nil
}
