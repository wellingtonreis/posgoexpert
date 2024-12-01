package limiter

import "errors"

type RateLimiter struct {
	Storage       Storage
	MaxRequests   int
	BlockDuration int64
}

func (rl *RateLimiter) Allow(key string) (bool, error) {
	count, err := rl.Storage.Increment(key)
	if err != nil {
		return false, err
	}

	if count > rl.MaxRequests {
		err := rl.Storage.Block(key, rl.BlockDuration)
		if err != nil {
			return false, err
		}
		return false, errors.New("rate limit exceeded")
	}

	return true, nil
}
