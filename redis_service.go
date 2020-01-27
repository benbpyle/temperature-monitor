package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

// RedisService service
type RedisService struct {
	pool *redis.Pool
	conn redis.Conn
}

// New return new service
func New(url string) *RedisService {
	if &url == nil {
		log.Fatal("input is required")
	}
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", url)
		},
	}

	// Get a connection
	conn := redispool.Get()
	defer conn.Close()
	// Test the connection
	p, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("can't connect to the redis database, got error:\n%v", err)
	}

	log.Println("Connected: ", p)

	return &RedisService{
		pool: redispool,
		conn: conn,
	}
}

// Publish publish key value
func (s *RedisService) Publish(p *Peripheral) error {
	json, err := json2.Marshal(p)
	fmt.Println("Publishing: ", p)
	conn := s.pool.Get()
	_, _ = conn.Do("PUBLISH", "c1", json)

	if err != nil {
		return fmt.Errorf("error publishing: %s", err)
	}

	return nil
}
