package resources

import "sync"

// BidModel - модель заявок покупателей.
type BidModel struct {
	Type         string                    `json:"type"            bson:"_id"`           // Тип
	Orders       map[string]*OrdersModel   `json:"orders"          bson:"orders"`        // Список ордеров
	ValueOrders  uint32                    `json:"value_orders"    bson:"value_orders"`  // Кол-во ордеров
	AveragePrice float32                   `json:"average_price"   bson:"average_price"` // Средняя стоимость

	rwMutex *sync.RWMutex // RWMutex
}

// AskModel - модель заявок продавцов.
type AskModel struct {
	Type         string                    `json:"type"            bson:"_id"`           // Тип
	Orders       map[string]*OrdersModel   `json:"orders"          bson:"orders"`        // Список ордеров
	ValueOrders  uint32                    `json:"value_orders"    bson:"value_orders"`  // Кол-во ордеров
	AveragePrice float32                   `json:"average_price"   bson:"average_price"` // Средняя стоимость

	rwMutex *sync.RWMutex // RWMutex
}

// OrdersModel - модель ордеров.
type OrdersModel struct {
	Orders            []interface{} `json:"orders"              bson:"orders"`              // Список ордеров
	ValueAllResources uint64        `json:"value_all_resources" bson:"value_all_resources"` // Общее кол-во ресурсов на цене
}
