package models

import (
	"errors"
	"time"
)

type Bucket struct {
	Capacity           int
	CurrentTokenNumber int
	IpAddr             string
	RefillRate         float64
	LastRefill         time.Time
}

func BuildBucket(ipaddr string, refillRate float64) *Bucket {
	return &Bucket{Capacity: 10, CurrentTokenNumber: 10, IpAddr: ipaddr, RefillRate: refillRate, LastRefill: time.Now()}
}

func (b *Bucket) TakeToken() error {
	now := time.Now()
	// time passed since last refill in seconds
	elapsedSeconds := now.Sub(b.LastRefill).Seconds()

  // if enough time has passed to add tokens to the bucket
	if elapsedSeconds >= b.RefillRate {
		// 20 -> 2
		// 60 -> 1
		b.CurrentTokenNumber += int(elapsedSeconds / b.RefillRate)

		if b.CurrentTokenNumber >= b.Capacity {
			b.CurrentTokenNumber = b.Capacity
		}

		b.LastRefill = now
	}

	if b.CurrentTokenNumber > 0 {
		b.CurrentTokenNumber--
		return nil
	}

	return errors.New("empty bucket encountered")
}
