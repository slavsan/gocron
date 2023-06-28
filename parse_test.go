package gocron

import (
	"reflect"
	"testing"
	"time"
)

const extendedRFC3339 = "2006-01-02T15:04:05.000Z07:00"

func TestParse(t *testing.T) {
	testCases := []struct {
		currentTime    string
		cron           string
		tillInitialRun time.Duration
		shouldStartAt  string
		interval       time.Duration
		secondRunAt    string
	}{
		// with invalid cron
		// TODO: add test cases
		// every minute
		{
			"2022-07-14T17:43:28.371Z",
			"* * * * *",
			(31 * time.Second) + (629 * time.Millisecond),
			"2022-07-14T17:44:00.000Z",
			1 * time.Minute,
			"2022-07-14T17:45:00.000Z",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			clock := newMockClock(tc.currentTime)
			tillInitialRun, interval := parse(tc.cron, clock)
			if !reflect.DeepEqual(tc.tillInitialRun, tillInitialRun) {
				t.Errorf(
					"duration till initial run does not match\n\texpected: %v\n\t  actual: %v\n",
					tc.tillInitialRun, tillInitialRun,
				)
			}
			now := clock.Now().Add(tillInitialRun).UTC()
			shouldStartAt := now.Format(extendedRFC3339)
			if !reflect.DeepEqual(tc.shouldStartAt, shouldStartAt) {
				t.Errorf(
					"formatted initial run time does not match\n\texpected: %v\n\t  actual: %v\n",
					tc.shouldStartAt, shouldStartAt,
				)
			}
			if !reflect.DeepEqual(tc.interval, interval) {
				t.Errorf(
					"interval duration does not match\n\texpected: %v\n\t  actual: %v\n",
					tc.interval, interval,
				)
			}
			now = now.Add(interval)
			secondRunAt := now.Format(extendedRFC3339)
			if !reflect.DeepEqual(tc.secondRunAt, secondRunAt) {
				t.Errorf(
					"secondRunAt time does not match\n\texpected: %v\n\t  actual: %v\n",
					tc.secondRunAt, secondRunAt,
				)
			}
		})
	}
}

type clockMock struct {
	now time.Time
}

func newMockClock(date string) Clock {
	now, err := time.Parse(time.RFC3339, date)
	if err != nil {
		panic(err)
	}
	return &clockMock{now: now}
}

func (m *clockMock) Now() time.Time {
	return m.now
}
