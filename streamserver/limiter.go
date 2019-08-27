package main

import "log"

//流控方案，server多请求会导致带宽打满，server会crash
//bucket：token 1～token n
//一个request来了，发一个token，一个response来了，还回一个token

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}

	cl.bucket <- 1
	return true
}
