package types

import (
	"encoding/json"
)

type Order struct {
	OrderUid          string   `json:"order_uid" fake:"{regex:[a-zA-Z0-9]{14}}test"`
	TrackNumber       string   `json:"track_number" fake:"{regex:[A-Z]{14}}TEST"`
	Entry             string   `json:"entry" fake:"{regex:[A-Z]{4}}"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Item   `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	Shardkey          string   `json:"shardkey"`
	SmID              int64    `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name" fake:"{firstname}"`
	Phone   string `json:"phone" fake:"{phone}"`
	Zip     string `json:"zip" fake:"{zip}"`
	City    string `json:"city" fake:"{city}"`
	Address string `json:"address" fake:"{address}"`
	Region  string `json:"region" fake:"{region}"`
	Email   string `json:"email" fake:"{email}"`
}

type Item struct {
	ChrtID      int64  `json:"chrt_id" fake:"{regex:[0-9]{7}}"`
	TrackNumber string `json:"track_number" fake:"{regex:[A-Z]{14}}TEST"`
	Price       int64  `json:"price" fake:{int}`
	Rid         string `json:"rid" fake:"{regex:[a-zA-Z0-9]{14}}test"`
	Name        string `json:"name" fake:"{randomstring:[pickle, lion's head, shoes]}"`
	Sale        int64  `json:"sale"`
	Size        string `json:"size" fake:"{randomstring:[0,1,2,3]}"`
	TotalPrice  int64  `json:"total_price"`
	NmID        int64  `json:"nm_id"`
	Brand       string `json:"brand" fake:"{company}"`
	Status      int64  `json:"status" fake:"{randomstring:[202,200,400,201]}"`
}

type Payment struct {
	Transaction  string `json:"transaction" fake:"{regex:[a-zA-Z0-9]{14}}test"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" fake:"{currency}"`
	Provider     string `json:"provider" fake:"{company}"`
	Amount       int64  `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank" fake:"{randomstring:[sber,alpha,tbank]}"`
	DeliveryCost int64  `json:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total"`
	CustomFee    int64  `json:"custom_fee"`
}

func ReadJSON(data []byte) (Order, error) {
	res := Order{}
	err := json.Unmarshal(data, &res)
	return res, err
}
