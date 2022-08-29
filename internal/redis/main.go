package redisTools

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)
const (
	Address  = "localhost:6379"
	Password = ""
	PoolSize    = 1000
	PoolTimeout = time.Second * 30
)

var nodeSpace *redis.Client

func init() {
	nodeSpace = redis.NewClient(&redis.Options{
		Addr:        Address,
		Password:    Password,
		DB:          2,
		PoolTimeout: PoolTimeout,
		PoolSize:    PoolSize,
	})
	
	ctx := context.Background()
	err := nodeSpace.Ping(ctx).Err()
	ticker := time.NewTicker(time.Minute)

	for err!=nil{
		select{
		case <-ticker.C:
			log.Fatalf("Coudn't connect to redis db")
		case <-time.After(time.Second):
			err = nodeSpace.Ping(ctx).Err()
			log.Println("waiting for redis")
		}
	}

}
var mainCtx = context.Background() 
func RegisterNode()int{
	resp := nodeSpace.Incr(mainCtx, "cluster_size")
	if resp.Err()!=nil{
		log.Fatalf("couldn't get cluster size from redis %v:", resp.Err())
	}
	return int(resp.Val())
}
