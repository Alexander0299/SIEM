package service

import (
	"fmt"
	"sync"
	"time"
)

var (
	prevLogCount   int
	prevUserCount  int
	prevAlertCount int
)

func LogChanges(logMutex, userMutex, alertMutex *sync.Mutex) {
	for {
		time.Sleep(200 * time.Millisecond)

		logMutex.Lock()
		if len(logs) > prevLogCount {
			for i := prevLogCount; i < len(logs); i++ {
				fmt.Println("Новый лог:", logs[i])
			}
			prevLogCount = len(logs)
		}
		logMutex.Unlock()

		userMutex.Lock()
		if len(users) > prevUserCount {
			for i := prevUserCount; i < len(users); i++ {
				fmt.Println("Новый пользователь:", users[i])
			}
			prevUserCount = len(users)
		}
		userMutex.Unlock()

		alertMutex.Lock()
		if len(alerts) > prevAlertCount {
			for i := prevAlertCount; i < len(alerts); i++ {
				fmt.Println("Новое предупреждение:", alerts[i])
			}
			prevAlertCount = len(alerts)
		}
		alertMutex.Unlock()
	}
}
