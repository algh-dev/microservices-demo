package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/algh-dev/microservices-demo/cache"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type MovieRepository struct {
	client *redis.Client
	ctx context.Context
}

func NewMovieRepository() MovieRepository {
	redisClient := cache.GetRedisClient()

	return MovieRepository{
		client: redisClient,
		ctx: context.Background(),
	}
}

func (repository MovieRepository) CreateMovie(movie *cache.Movie) (*cache.Movie, error) {
	//cache, ctx := GetRedisClient()

	movie.Id = uuid.New().String()

	err := repository.client.HSet(repository.ctx, "movies", movie.Id, *movie).Err()

	if (err != nil) {
		return nil, err
	}
	log.Println("Movie saved in Redis with id ", movie.Id)
	return movie, nil

}

func (repository MovieRepository) GetMovie(id string) (*cache.Movie, error) {
	//cache, ctx := GetRedisClient()

	val, err := repository.client.HGet(repository.ctx, "movies", id).Result()

	if err != nil && err.Error() == "redis: nil" {
		return nil, fmt.Errorf("Movie not found, id=%s", id)
	}

	movie := &cache.Movie{}

	if err := movie.UnmarshalBinary([]byte(val)); err != nil {
		return nil, err
	} 

	return movie, nil
}