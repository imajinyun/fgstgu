package c01

import (
	"testing"
	"time"

	"github.com/uniplaces/carbon"
)

func TestCarbon(t *testing.T) {
	t.Logf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("America/Vancouver")
	t.Logf("Right now in Vancouver is %s\n", today)
	t.Logf("Tomorrow is %s\n", carbon.Now().AddDay())
	t.Logf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlympics, _ := carbon.CreateFromDate(2020, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	t.Logf("Next olympics are in %d\n", nextOlympics.Year())
	if carbon.Now().IsWeekday() {
		t.Log("🐶 Working in Company!")
	}
}
