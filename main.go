package main

import (
	"database/sql"
	"log"
	"main/notification"
	"main/order"
	"main/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// БД
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepo(db)
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	// EmailSender
	emailNotifier := notification.NewEmailSender()
	orderService := order.NewOrderService(repo, emailNotifier)

	err = orderService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// SMSSender
	smsNotifier := notification.NewSMSSender("+79991234567")
	orderService2 := order.NewOrderService(repo, smsNotifier)

	err = orderService2.CreateOrder("Петр", []string{"orange", "grape"}, 15.75)
	if err != nil {
		log.Fatal(err)
	}
}
