package main

import (
	"fmt"
	"html/template"
	"main/internal/cache"
	"main/internal/configs"
	"main/internal/db"
	"main/internal/kafka/consumer"
	"main/internal/logger"
	"main/internal/types"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	logger.InitLogger()

	configs.InitConfig()
	c := cache.New()
	d := db.New()

	conf := configs.GetConfig()

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Dbname)

	d.Connect("postgres", dbInfo)
	defer d.Close()

	c.RestoreDataFromDB(d)
	go consumer.StartConsumer(d, c)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Err(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, types.Order{})
	})
	http.HandleFunc("/order", func(w http.ResponseWriter, req *http.Request) {
		uid := req.URL.Query().Get("uid")
		order, err := c.GetOrderById(uid)

		log.Debug().Msg(fmt.Sprintf("ORDER: %#v\n", order))

		if err != nil {
			log.Info().Msg("no order with that uid")
		}
		tmpl.Execute(w, order)

	})
	log.Err(http.ListenAndServe(":8000", nil))
}
