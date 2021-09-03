package solar_terms

import (
	"fmt"
	"testing"
	"time"
)

func TestSolar_IsSolarTerm(t *testing.T) {
	solar := &Solar{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)}
	terms, err := solar.getCurrentYearSolarTerms()
	fmt.Println(terms)
	fmt.Println(err)
}
