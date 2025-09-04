package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_basic() {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(2)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11433.crce179.ap-south-1-1.ec2.redns.redis-cloud.com:11433",
		Username: "default",
		Password: "wNUf8rJ9V7AA8oSFucSxI6uOZ96YYbNL",
		DB:       0,
	})
	defer rdb.Close()
	rdb.LPush(ctx, "Persons", "123")
	rdb.LPush(ctx, "Persons", "345")
	rdb.LPush(ctx, "Persons", "567")
	rdb.Set(ctx, "foo", "bar", 0)
	rdb.Set(ctx, "boo", "bar", 0)
	rdb.Set(ctx, "coo", "bar", 0)
	rdb.Set(ctx, "loo", "bar", 0)
	rdb.SAdd(ctx, "Players", "Harsh")
	rdb.SAdd(ctx, "Players", "Hemant", "Aashish")
	fmt.Println(rdb.SMembers(ctx, "Players").Result())
	rdb.SRem(ctx, "Players", "Hemant")
	fmt.Println(rdb.SMembers(ctx, "Players").Result())
	rdb.SAdd(ctx, "Players", "Aashish")
	fmt.Println(rdb.SMembers(ctx, "Players").Result())
	fmt.Println(rdb.DBSize(ctx).Result())
	time.Sleep(1 * time.Second)
	go func() {
		defer wg.Done()
		le, _ := rdb.LLen(ctx, "Persons").Result()
		for i := 0; i < int(le); i++ {
			result, err := rdb.LPop(ctx, "Persons").Result()

			if err != nil {
				panic(err)
			}

			fmt.Println(result) // >>> bar
		}
	}()
	go func() {
		defer wg.Done()
		result, err := rdb.Get(ctx, "foo").Result()

		if err != nil {
			panic(err)
		}

		fmt.Println(result) // >>> bar
	}()
	wg.Wait()
}

func main() {
	ExampleClient_connect_basic()
}
