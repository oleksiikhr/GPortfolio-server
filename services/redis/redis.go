package redis

import (
	"encoding/json"
	"time"

	"github.com/GPortfolio/server/services/helpers"
	"github.com/go-redis/redis/v7"
)

// TODO
type Redis struct {
	Client *redis.Client
}

// TODO
type Structure struct {
	Value interface{}
	Pass  string
}

// TODO
type KeyPass struct {
	Key string
	Pass string
}

// NewRedis create a new connection to Redis client
func NewRedis(addr string) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Check connection
	_, err := client.Ping().Result()

	return &Redis{Client: client}, err
}

// TODO
func (r *Redis) RndSet(value interface{}, t time.Duration) (*KeyPass, error) {
	keyPass := generateKeyPass()
	encoding, err := json.Marshal(Structure{
		Value: value,
		Pass:  keyPass.Pass,
	})

	if err != nil {
		return keyPass, err
	}

	err = r.Client.Set(keyPass.Key, encoding, t).Err()
	if err != nil {
		return keyPass, err
	}

	return keyPass, nil
}

// TODO
func generateKeyPass() *KeyPass {
	return &KeyPass{
		Key: helpers.Rnd(40),
		Pass: helpers.Rnd(60),
	}
}
