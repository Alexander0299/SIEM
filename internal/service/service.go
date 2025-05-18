package service

import (
	"context"
	"log"
	"siem-sistem/internal/model"
	"siem-sistem/internal/storage"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	idMutex sync.Mutex

	mongoClient *mongo.Client
	redisClient *redis.Client

	mongoUsers  *mongo.Collection
	mongoAlerts *mongo.Collection
	mongoLogs   *mongo.Collection
)

func InitStorage() {
	mongoClient = storage.InitMongo()
	redisClient = storage.InitRedis()

	db := mongoClient.Database("siem")

	mongoUsers = db.Collection("users")
	mongoAlerts = db.Collection("alerts")
	mongoLogs = db.Collection("logs")
}

func GetNextID(entity string) int {
	idMutex.Lock()
	defer idMutex.Unlock()

	ctx := context.Background()
	val, err := redisClient.Incr(ctx, "id:"+entity).Result()
	if err != nil {
		log.Printf("Ошибка Redis при генерации ID: %v", err)
		return 0
	}
	return int(val)
}

func AddUsers(ctx context.Context, usersChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			usersChan <- model.User{
				ID:    GetNextID("user"),
				Login: "Alex",
			}
		}
	}
}

func AddAlerts(ctx context.Context, alertsChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			alertsChan <- model.Alert{
				ID:      GetNextID("alert"),
				Massage: "Попытка взлома",
			}
		}
	}
}

func AddLogs(ctx context.Context, logsChan chan model.Inter) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			logsChan <- model.Log{
				ID:   GetNextID("log"),
				Area: "Антивирус Касперского",
			}
		}
	}
}

func Logger(usersChan, alertsChan, logsChan chan model.Inter) {
	users := []model.User{}
	alerts := []model.Alert{}
	logs := []model.Log{}

	var totalUsers, totalAlerts, totalLogs int

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case item, ok := <-usersChan:
			if !ok {
				usersChan = nil
				continue
			}
			if user, ok := item.(model.User); ok {
				users = append(users, user)
				totalUsers++
			}

		case item, ok := <-alertsChan:
			if !ok {
				alertsChan = nil
				continue
			}
			if alert, ok := item.(model.Alert); ok {
				alerts = append(alerts, alert)
				totalAlerts++
			}

		case item, ok := <-logsChan:
			if !ok {
				logsChan = nil
				continue
			}
			if logItem, ok := item.(model.Log); ok {
				logs = append(logs, logItem)
				totalLogs++
			}

		case <-ticker.C:
			log.Printf("Количество пользователей=%d, Количество уведомлений=%d, Количество логов=%d",
				totalUsers, totalAlerts, totalLogs)

			if err := RewriteUsers(users, ""); err != nil {
				log.Printf("Ошибка сохранения пользователей: %v", err)
			}
			if err := RewriteAlerts(alerts, ""); err != nil {
				log.Printf("Ошибка сохранения уведомлений: %v", err)
			}
			if err := RewriteLogs(logs, ""); err != nil {
				log.Printf("Ошибка сохранения логов: %v", err)
			}

			users = []model.User{}
			alerts = []model.Alert{}
			logs = []model.Log{}
		}

		if usersChan == nil && alertsChan == nil && logsChan == nil {
			return
		}
	}
}

func RewriteUsers(users []model.User, _ string) error {
	collection := mongoClient.Database("siem").Collection("users")
	ctx := context.Background()

	docs := make([]interface{}, len(users))
	for i, u := range users {
		docs[i] = u
	}

	if _, err := collection.InsertMany(ctx, docs); err != nil {
		return err
	}
	return nil
}

func RewriteAlerts(alerts []model.Alert, _ string) error {
	collection := mongoClient.Database("siem").Collection("alerts")
	ctx := context.Background()

	docs := make([]interface{}, len(alerts))
	for i, a := range alerts {
		docs[i] = a
	}

	if _, err := collection.InsertMany(ctx, docs); err != nil {
		return err
	}
	return nil
}

func RewriteLogs(logs []model.Log, _ string) error {
	collection := mongoClient.Database("siem").Collection("logs")
	ctx := context.Background()

	docs := make([]interface{}, len(logs))
	for i, l := range logs {
		docs[i] = l
	}

	if _, err := collection.InsertMany(ctx, docs); err != nil {
		return err
	}
	return nil
}
func LoadUsersFrom(_ string) []model.User {
	ctx := context.Background()
	cursor, err := mongoUsers.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Ошибка чтения пользователей из MongoDB: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var users []model.User
	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Ошибка декодирования пользователя: %v", err)
			continue
		}
		users = append(users, user)
	}
	return users
}

func LoadAlertsFrom(_ string) []model.Alert {
	ctx := context.Background()
	cursor, err := mongoAlerts.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Ошибка чтения уведомлений из MongoDB: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var alerts []model.Alert
	for cursor.Next(ctx) {
		var alert model.Alert
		if err := cursor.Decode(&alert); err != nil {
			log.Printf("Ошибка декодирования уведомления: %v", err)
			continue
		}
		alerts = append(alerts, alert)
	}
	return alerts
}

func LoadLogsFrom(_ string) []model.Log {
	ctx := context.Background()
	cursor, err := mongoLogs.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Ошибка чтения логов из MongoDB: %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var logs []model.Log
	for cursor.Next(ctx) {
		var logEntry model.Log
		if err := cursor.Decode(&logEntry); err != nil {
			log.Printf("Ошибка декодирования лога: %v", err)
			continue
		}
		logs = append(logs, logEntry)
	}
	return logs
}
