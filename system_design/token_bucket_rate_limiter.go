package system_design

import (
	"fmt"
	"math"
	"time"
)

const (
	MAX_BUCKET_SIZE float64 = 3
	REFILL_RATE     int     = 1
)

type TokenBucket struct {
	currentBucketSize   float64
	lastRefillTimestamp int
}

func (tb *TokenBucket) allowRequest(tokens float64) bool {
	tb.refill()
	if tb.currentBucketSize > 0 {
		tb.currentBucketSize -= 1
		return true
	}
	return false
}

func (tb *TokenBucket) refill() {
	now := time.Now().Nanosecond()
	elapsedTime := now - tb.lastRefillTimestamp
	tokensToAdd := (elapsedTime/1e9) * REFILL_RATE

	tb.currentBucketSize = math.Min(tb.currentBucketSize+float64(tokensToAdd), MAX_BUCKET_SIZE);
	tb.lastRefillTimestamp = now
}

func main(){
	obj := TokenBucket{
		currentBucketSize: 3,
		lastRefillTimestamp: 0,
	}

	fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
	fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
	fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //true
	fmt.Printf("Request processed: %v\n", obj.allowRequest(1)) //false, request dropped
}