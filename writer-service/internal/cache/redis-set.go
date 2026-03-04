package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"writer/internal/model"
)


func SetJSON(
	ctx context.Context,
	cacheKey string, 
	value any, 
	expr int,
) error {
	jsonData, err := json.Marshal(value);
	if err != nil {
		fmt.Println("Failed to set cache:" + err.Error())
		return err
	}

	Client.Set(
		ctx, 
		cacheKey, 
		jsonData, 
		time.Duration(expr)*time.Minute,
	)

	return nil
}

func SetWriterJSONs(
	ctx context.Context,
	cacheKey string, 
	values []*model.Writer,
	expr int64,
) error {
	pipe := Client.Pipeline()

	for _, value := range values {
		jsonData, _ := json.Marshal(value)
		pipe.Set(
			ctx, fmt.Sprintf(cacheKey, value.Id), 
			jsonData, time.Duration(expr)*time.Minute,
		)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

