package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"
)

//漏桶算法

type Strategy int

const (
	LEAKY_BUCKET Strategy = iota
	TOKEN_BUCKET
)

type bucket struct {
	interval      time.Duration
	ticker        *time.Ticker
	stopTicker    chan struct{}
	count         int
	volume        int
	current       int
	qps           int
	refreshTicker bool
	sync.RWMutex
}

func newBucket() *bucket {
	return &bucket{
		stopTicker:    make(chan struct{}),
		refreshTicker: true,
		qps:           -1,
	}
}

func (b *bucket) set(qps int, strategy Strategy, name string) *bucket {
	qpsStr := strconv.Itoa(qps)
	if qps <= 0 {
		qpsStr = "no limit"
		qps = 0
	}

	if b.qps == qps {
		b.refreshTicker = false
		return b
	}

	b.qps = qps
	b.refreshTicker = true
	fmt.Printf("set %s qps to %s\n", name, qpsStr)

	if qps <= 0 {
		b.Lock()
		b.interval = 0
		b.count = 0
		b.volume = 0
		b.current = 0
		b.Unlock()
		return b
	}

	count := 1
	interval := time.Second / time.Duration(qps)
	if qps > 100 {
		gcd := big.NewInt(0).GCD(nil, nil, big.NewInt(int64(qps)), big.NewInt(100)).Int64()
		count = qps / int(gcd)
		interval = 10 * time.Millisecond * time.Duration(100/gcd)
	}

	volume := 0
	if strategy == TOKEN_BUCKET {
		volume = 2 * count
	} else {
		volume = count
	}

	b.Lock()
	b.count = count
	b.volume = volume
	if b.current > volume {
		b.current = volume
	}
	b.interval = interval
	b.Unlock()
	return b
}

func (b *bucket) startTicker() *bucket {
	if !b.refreshTicker {
		return b
	}

	b.refreshTicker = false
	b.Lock()
	defer b.Unlock()

	if b.ticker != nil {
		b.ticker.Stop()
		b.stopTicker <- struct{}{}
		<-b.stopTicker
		b.ticker = nil
	}

	if b.interval != 0 {
		b.ticker = time.NewTicker(b.interval)
	}

	if b.ticker != nil {
		go func() {
			for {
				select {
				case <-b.ticker.C:
					b.Lock()
					b.current -= b.count
					if b.current < 0 {
						b.current = 0
					}
					b.Unlock()
				case <-b.stopTicker:
					b.stopTicker <- struct{}{}
					return
				}
			}
		}()
	}
	return b
}

func (b *bucket) add() bool {
	b.Lock()
	if b.interval == 0 {
		b.Unlock()
		return true
	}
	success := true
	if b.current+1 > b.volume {
		success = false
	} else {
		b.current++
	}
	b.Unlock()
	return success
}

func (b *bucket) getQPS() int {
	b.RLock()
	qps := b.qps
	b.RUnlock()
	return qps
}
