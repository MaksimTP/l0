package main

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/cache"
	"main/internal/configs"
	"main/internal/db"
	"main/internal/kafka/consumer"
	"main/internal/types"
	"net/http"
)

func main() {
	configs.InitConfig()
	c := cache.NewCache()
	d := db.DataBase{}

	conf := configs.GetConfig()

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Dbname)

	d.Connect("postgres", dbInfo)
	defer d.Close()

	c.RestoreDataFromDB(d)
	go consumer.StartConsumer(d, c)

	tmpl, err := template.ParseFiles("../static/index.html")
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, types.Order{})
	})
	http.HandleFunc("/order", func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		order, err := c.GetOrderById(uid)
		log.Printf("ORDER: %#v\n", order)
		if err != nil {
			log.Println("no order with that uid")
		}
		tmpl.Execute(w, order)

	})
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
