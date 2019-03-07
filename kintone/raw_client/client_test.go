package raw_client

import (
	"testing"
	"time"
)

func TestGetWaitDuration(t *testing.T) {
	patterns := []struct {
		retry             int
		expectedMinSecond int
		expectedMaxSecond int
	}{
		{0, 1, 2},
		{1, 2, 3},
		{2, 4, 5},
		{3, 8, 9},
		{4, 16, 17},
		{5, 32, 33},
		{6, 64, 64},
		{7, 64, 64},
	}

	for _, pattern := range patterns {
		actual := getWaitDuration(pattern.retry)
		expectedMinDuration := time.Duration(pattern.expectedMinSecond) * time.Second
		expectedMaxDuration := time.Duration(pattern.expectedMaxSecond) * time.Second
		if actual < expectedMinDuration || actual > expectedMaxDuration {
			t.Errorf("getWaitDuration( retry: %d ): expected( %d <= x <= %d ), actual %d",
				pattern.retry, expectedMinDuration, expectedMaxDuration, actual)
		}
	}
}
