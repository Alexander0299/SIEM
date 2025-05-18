package storage

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/redis/go-redis/v9"
)

var MongoClient *mongo.Client
var RedisClient *redis.Client

func InitMongo() *mongo.Client {
	// Читаем адрес из переменной окружения
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		// По умолчанию для локального запуска
		uri = "mongodb://localhost:27017"
	}

	// Создаём контекст с таймаутом на подключение
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Пытаемся подключиться
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB (%s): %v", uri, err)
	}

	// Проверяем связь ping’ом
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Не удалось ping к MongoDB (%s): %v", uri, err)
	}

	MongoClient = client
	log.Printf("Успешно подключились к MongoDB: %s", uri)
	return client
}
func InitRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379" // значение по умолчанию для локального запуска
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: addr, // <— используем переменную, а не хардкод
		DB:   0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Ошибка подключения к Redis (%s): %v", addr, err)
	}
	RedisClient = rdb
	return rdb
}
