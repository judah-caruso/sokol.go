package time_test

import (
	"testing"

	"github.com/judahcaruso/sokol.go/time"
)

func TestTime(t *testing.T) {
	time.Setup()

	start := time.Now()
	if start == 0 {
		t.FailNow()
	}
}
