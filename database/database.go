package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnSQL() (*gorm.DB, error) {
	// mendapatkan data DB dari ENV
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	if DBHost == "" || DBUser == "" || DBName == "" || DBPass == "" || DBPort == "" {
		return nil, fmt.Errorf("Failed Connect DB : Invalid Data DB")
	}

	// membuat koneksi ke DB SQL
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed Connect DB : " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed Connect DB : " + err.Error())
	}

	// menset data max idle connection, max open connection, dan max lifetime connection
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Success")
	return db, nil
}

func NewDatabaseRedis() (map[string]*redis.Client, error) {
	// mendapatkan data redis dari env
	fmt.Println("starting redis....")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PAS")
	redisPool, _ := strconv.Atoi(os.Getenv("REDIS_POOL"))
	redisAge, _ := strconv.Atoi(os.Getenv("REDIS_AGE"))
	redisType := strings.Split(os.Getenv("REDIS_TYPE"), ",")
	if redisPort == "" {
		redisPort = "6379"
	}

	// membuat koneksi ke redis sebanyak data yang di dapat dari env
	redisDB := map[string]*redis.Client{}
	for i, v := range redisType {
		rdb := redis.NewClient(&redis.Options{
			Addr:       fmt.Sprintf("%s:%s", redisHost, redisPort),
			Password:   redisPass, // no password set
			DB:         i,         // use default DB
			PoolSize:   redisPool,
			MaxConnAge: time.Duration(redisAge) * time.Second,
		})

		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			return nil, fmt.Errorf("Failed to connect redis with DB "+strconv.Itoa(i)+" and error : ", err)
		}
		redisDB[v] = rdb
	}

	fmt.Println("Successfully connect to Redis")
	return redisDB, nil
}
