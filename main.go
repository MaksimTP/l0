package main

import (
	"html/template"
	"log"
	"main/cache"
	"main/db"
	"main/kafka/consumer"
	"main/model"
	"net/http"
)

func main() {
	c := cache.NewCache()
	d := db.DataBase{}
	d.Connect("postgres", db.DbInfo)
	c.RestoreDataFromDB(d)
	go consumer.StartConsumer(d, c)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, model.Order{})
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
