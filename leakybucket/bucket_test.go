package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetBucket(t *testing.T) {
	cases := []struct {
		qps      int
		count    int
		interval time.Duration
	}{
		{1, 1, time.Second},
		{2, 1, 500 * time.Millisecond},
		{100, 1, 10 * time.Millisecond},
		{101, 101, time.Second},
		{110, 11, 100 * time.Millisecond},
		{150, 3, 20 * time.Millisecond},
		{1000, 10, 10 * time.Millisecond},
		{1010, 101, 100 * time.Millisecond},
		{1500, 15, 10 * time.Millisecond},
		{1511, 1511, time.Second},
	}
	l := newBucket()
	for _, want := range cases {
		if l.set(want.qps, LEAKY_BUCKET, "test"); l.count != want.count || l.interval != want.interval {
			fmt.Println("qps %v want count %v interval %v, got count %v interval %v\n", want.qps, want.count, want.interval, l.count, l.interval)
			//fmt.Println()
		}
	}

}
