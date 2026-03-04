package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)


func GetJSON[T any](
	ctx context.Context,
	cacheKey string,
) (T, error) {
	var result T

	cached, err := get(ctx, cacheKey)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(cached), &result)
	if err != nil {
		return result, err
	}

	fmt.Println("Cache hit => " + cacheKey)
	return result, nil
}

func GetInt(
	ctx context.Context,
	cacheKey string,
) (int64, error) {
	cached, err := get(ctx, cacheKey)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseInt(cached, 10, 64)
	if err != nil {
		return 0, err
	}

	fmt.Println("Cache hit => " + cacheKey)
	return  result, nil
}



func get(
	ctx context.Context, 
	key string,
) (string, error) {
	return Client.Get(ctx, key).Result()
}
