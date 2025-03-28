package main

import (
	"fmt"
	"siem-system/internal/repository"
)

func main() {
	repo := repository.NewRepository("logs.csv", "users.csv", "alerts.csv")

	if err := repo.Load(); err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		return
	}

	fmt.Println("Данные восстановлены.")

}
