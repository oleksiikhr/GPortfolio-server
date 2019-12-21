package redis

import (
	"encoding/json"
	"time"

	"github.com/GPortfolio/server/services/helpers"
	"github.com/go-redis/redis/v7"
)

// Redis client
type Redis struct {
	Client *redis.Client
}

// Structure stores in Redis
type Structure struct {
	Value interface{}
	Pass  string
}

// KeyPass
type KeyPass struct {
	Key  string
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

// SecSet generate random key and value and store value
// to Redis
func (r *Redis) SecSet(value interface{}, t time.Duration) (*KeyPass, error) {
	keyPass := generateKeyPass()
	encoding, err := json.Marshal(Structure{
		Value: value,
		Pass:  keyPass.Pass,
	})

	if err != nil {
		return keyPass, err
	}

	err = r.Client.Set("sec:"+keyPass.Key, encoding, t).Err()
	if err != nil {
		return keyPass, err
	}

	return keyPass, nil
}

// SecGetHard try get data by key and password
// if key equals but password not - destroy data
func (r *Redis) SecGetHard(key, password string) interface{} {
	value, err := r.Client.Get("sec:" + key).Bytes()
	if err != nil {
		return nil
	}

	var structure *Structure
	err = json.Unmarshal(value, &structure)
	if err != nil {
		return nil
	}

	// Destroy data if the attempt is unsuccessful
	if structure.Pass != password {
		r.Client.Del("sec:" + key)
		return nil
	}

	return structure.Value
}

// generateKeyPass for store value in database*
func generateKeyPass() *KeyPass {
	return &KeyPass{
		Key:  helpers.Rnd(70),
		Pass: helpers.Rnd(30),
	}
}
