package main

import (
	"database/sql"
	"log"

	"solid/notification"
	"solid/order"
	"solid/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepo(db)
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	log.Println("EMAIL УВЕДОМЛЕНИЯ:")
	emailNotifier := notification.NewEmailSender()
	emailService := order.NewOrderService(repo, emailNotifier)

	if err := emailService.CreateOrder("Иван Петров", []string{"Apple", "Banana", "Orange"}, 150.50); err != nil {
		log.Printf("Ошибка: %v", err)
	}

	log.Println("\nSMS УВЕДОМЛЕНИЯ:")
	smsNotifier := notification.NewSMSSender("+7 123 456-78-90")
	smsService := order.NewOrderService(repo, smsNotifier)

	if err := smsService.CreateOrder("Мария Смирнова", []string{"Milk", "Bread", "Butter"}, 85.75); err != nil {
		log.Printf("Ошибка: %v", err)
	}

	log.Println("\nВАЛИДАЦИЯ:")
	invalidService := order.NewOrderService(repo, emailNotifier)

	log.Println("Создание заказа с пустым именем:")
	if err := invalidService.CreateOrder("", []string{"item"}, 100.0); err != nil {
		log.Printf("Ошибка валидации: %v", err)
	}

	log.Println("\nсоздание заказа с пустым списком товаров:")
	if err := invalidService.CreateOrder("Клиент", []string{}, 100.0); err != nil {
		log.Printf("Ошибка валидации: %v", err)
	}

	log.Println("\nСоздание заказа с отрицательной суммой:")
	if err := invalidService.CreateOrder("Клиент", []string{"item"}, -50.0); err != nil {
		log.Printf("Ошибка валидации: %v", err)
	}
}
