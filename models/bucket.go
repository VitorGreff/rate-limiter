package models

import "errors"

type Bucket struct {
	Capacity int
	IpAddr   string
}

func BuildBucket(ipaddr string) Bucket {
	return Bucket{Capacity: 10, IpAddr: ipaddr}
}

func IsAddrNew(buckets []Bucket, ipaddr string) bool {
	for _, b := range buckets {
		if b.IpAddr == ipaddr {
			return false
		}
	}
	return true
}

func (b *Bucket) AddToken() error {
	if b.Capacity < 10{
		b.Capacity++
		return nil
	}
	return errors.New("bucket reached its full capacity!")
}

func (b *Bucket) ConsumeToken() error {
	if b.Capacity > 0 {
		b.Capacity--
		return nil
	}
	return errors.New("empty bucket!")
}
