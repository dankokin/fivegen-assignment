package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/dankokin/fivegen-assignment/models"
)

type RedisDataStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisDataStore(addr string, pass string, db int, ctx context.Context) *RedisDataStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return &RedisDataStore{
		client: redisClient,
		ctx:    ctx,
	}
}

func (r *RedisDataStore) UploadFileName(file *models.File, errChan chan error) {
	jFile, err := json.Marshal(file)
	if err != nil {
		errChan <- err
		return
	}

	err = r.client.Set(r.ctx, file.ShortUrl, jFile, 0).Err()
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}

func (r *RedisDataStore) DownloadFileName(url string) *models.File {
	jFile, err := r.client.Get(r.ctx, url).Result()
	if err == redis.Nil {
		return nil
	}

	var file models.File
	if err = json.Unmarshal([]byte(jFile), &file); err != nil {
		return nil
	}

	return &file
}

func (r *RedisDataStore) IsExists(key string, fileDataHash string) bool {
	value, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return false
	}

	var file models.File
	if err = json.Unmarshal([]byte(value), &file); err != nil {
		return true
	}

	if file.HashedName == fileDataHash {
		return false
	} else {
		return true
	}
}

func (r *RedisDataStore) AllFilesRecords() ([]string, error) {
	return nil, nil
}
