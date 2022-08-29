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

}
var mainCtx = context.Background() 
func RegisterNode()int{
	resp := nodeSpace.Incr(mainCtx, "cluster_size")
	if resp.Err()!=nil{
		log.Fatalf("couldn't get cluster size from redis %v:", resp.Err())
	}
	return int(resp.Val())
}
