package db

import (
	"database/sql"
	"fmt"
	"log"
	"main/internal/types"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty"
	dbname   = "wb_lvl0"
)

type DataBase struct {
	db *sql.DB
}

var DbInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func (d *DataBase) Connect(driverName string, dbInfo string) {
	db, err := sql.Open(driverName, dbInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	d.db = db
	log.Println("Connected to postgresql server with params:", dbInfo)
}

func (d *DataBase) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

const (
	insertStatementPayment = `
INSERT INTO payment (id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	insertStatementDelivery = `
INSERT INTO delivery (id, name, phone, zip, city, address, region, email)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	insertStatementItem = `
INSERT INTO item (id, order_uid, chrt_id, track_number, price, rid, sale, size, total_price, nm_id, brand, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	insertStatementOrder = `
INSERT INTO "order" (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
)

func (d *DataBase) GetNextIdToInsert(tableName string) int {
	query := "SELECT COUNT(*) FROM " + tableName
	rows, _ := d.db.Query(query)
	var id int
	rows.Next()
	err := rows.Scan(&id)
	if err != nil {
		log.Println(err.Error())
	}
	id++
	return id
}

func (d *DataBase) InsertData(data types.Order) {

	deliveryID := d.GetNextIdToInsert("delivery")
	_, err := d.db.Exec(insertStatementDelivery, deliveryID, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region, data.Delivery.Email)
	if err != nil {
		log.Println(err.Error())
	}

	paymentId := d.GetNextIdToInsert("payment")
	_, err = d.db.Exec(insertStatementPayment, paymentId, data.Payment.Transaction, data.Payment.RequestID, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.PaymentDt, data.Payment.Bank, data.Payment.DeliveryCost, data.Payment.GoodsTotal, data.Payment.CustomFee)
	if err != nil {
		log.Println(err.Error())
	}
	itemId := d.GetNextIdToInsert("item")

	for _, v := range data.Items {
		_, err = d.db.Exec(insertStatementItem, itemId, data.OrderUid, v.ChrtID, v.TrackNumber, v.Price, v.Rid, v.Sale, v.Size, v.TotalPrice, v.NmID, v.Brand, v.Status)
		if err != nil {
			log.Println(err.Error())
		}
		itemId++
	}

	_, err = d.db.Exec(insertStatementOrder, data.OrderUid, data.TrackNumber, data.Entry, deliveryID, paymentId, data.Locale, data.InternalSignature, data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID, data.DateCreated, data.OofShard)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Inserted data with order uid", data.OrderUid)
	}
}

func (d *DataBase) GetAllData() []types.Order {
	orders := make([]types.Order, 0)
	rows, err := d.db.Query(`SELECT * FROM "order" as o
	JOIN delivery as d on o.delivery_id = d.id
	JOIN payment as p on o.payment_id = p.id
	JOIN item as i on i.order_uid = o.order_uid`)

	if err != nil {
		log.Println(err)
	}
	// 43
	var order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, date_created, oof_shard, id, name, phone, zip, city, address, region, email, id1, transaction, request_id, currency, provider, bank, id2, order_uid2, track_number1, rid, size, brand string
	var amount, payment_dt, delivery_cost, goods_total, custom_fee, chrt_id, price, sale, total_price, nm_id, status, sm_id int64
	is_new_order := true
	prev_order_uid := ""
	for rows.Next() {
		rows.Scan(&order_uid, &track_number, &entry, &delivery_id, &payment_id, &locale, &internal_signature, &customer_id, &delivery_service, &shardkey, &sm_id, &date_created, &oof_shard, &id, &name, &phone, &zip, &city, &address, &region, &email, &id1, &transaction, &request_id, &currency, &provider, &amount, &payment_dt, &bank, &delivery_cost, &goods_total, &custom_fee, &id2, &order_uid2, &chrt_id, &track_number1, &price, &rid, &sale, &size, &total_price, &nm_id, &brand, &status)
		if len(orders) != 0 {
			is_new_order = prev_order_uid == order_uid
		}

		if is_new_order {
			orders = append(orders, types.Order{OrderUid: order_uid, TrackNumber: track_number, Entry: entry, Delivery: types.Delivery{Name: name, Phone: phone, Zip: zip, City: city, Address: address, Region: region, Email: email}, Payment: types.Payment{Transaction: transaction, RequestID: request_id, Currency: currency, Provider: provider, Amount: amount, PaymentDt: payment_dt, Bank: bank, DeliveryCost: delivery_cost, GoodsTotal: goods_total, CustomFee: custom_fee}, Items: []types.Item{{ChrtID: chrt_id, TrackNumber: track_number1, Price: price, Rid: rid, Name: name, Sale: sale, Size: size, TotalPrice: total_price, NmID: nm_id, Brand: brand, Status: status}}, Locale: locale, InternalSignature: internal_signature, CustomerID: customer_id, DeliveryService: delivery_service, Shardkey: shardkey, SmID: sm_id, DateCreated: date_created, OofShard: oof_shard})
		} else {
			for k, order := range orders {
				if order.OrderUid == order_uid {
					orders[k].Items = append(orders[k].Items, types.Item{ChrtID: chrt_id, TrackNumber: track_number1, Price: price, Rid: rid, Name: name, Sale: sale, Size: size, TotalPrice: total_price, NmID: nm_id, Brand: brand, Status: status})
				}
			}
		}
		prev_order_uid = order_uid
	}
	return orders
}
