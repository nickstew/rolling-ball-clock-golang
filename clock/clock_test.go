package clock

import (
	"testing"
)

func TestClock30(t *testing.T) {
	clock := New(30)
	result := clock.FindCycleDays()
	if (result != 15) {
		t.Fatal(result)
	}
}

func TestClock45(t *testing.T) {
	result := New(45).FindCycleDays()
	if (result != 378) {
		t.Fatal(result)
	}
}

func BenchmarkClock123(b *testing.B) {
	result := New(123).FindCycleDays()
	b.StopTimer()
	if (result != 108855) {
		b.Fatal(result)
	}
}