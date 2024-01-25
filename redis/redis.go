package redis

import (
	"log"

	"github.com/go-redis/redis"
)

func RedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

// ! StoreUserDetailInRedis function will store the user details in key,value pair in redis
func StoreUserDetailInRedis(userDetails map[string]interface{}) error {

	client := RedisConnection()
	err := client.HMSet("details", userDetails).Err()
	if err != nil {
		log.Println("Error in Storing the user details to Redis ", err)
		return err
	}
	return nil

}

// !GetUserDetailsFromRedis function will return the user details from redis
func GetUserDetailsFromRedis() (map[string]string, error) {
	client := RedisConnection()
	hashKey := "details"
	result, err := client.HGetAll(hashKey).Result()
	if err != nil {
		log.Println("Error in retrieving the user details from Redis ", err)
		return nil, err
	}

	userDetails := map[string]string{}
	for key, value := range result {
		userDetails[key] = value
	}
	return userDetails, nil

}
