package resources

import (
	"../../utils/config"
	"errors"
	"sync"
)

// Bid - заявки покупателей.
type Bid struct {
	Orders       map[float32]*Orders `json:"orders"          bson:"orders"`        // Список ордеров
	ValueOrders  uint32              `json:"value_orders"    bson:"value_orders"`  // Кол-во ордеров
	AveragePrice float32             `json:"average_price"   bson:"average_price"` // Средняя стоимость

	rwMutex *sync.RWMutex // RWMutex
}

// newBid - создание заявок продовцов.
func newBid() *Bid {
	return &Bid{
		Orders:       make(map[float32]*Orders),
		AveragePrice: 0,
		ValueOrders:  0,
		rwMutex:      new(sync.RWMutex),
	}
}

// AddOrder - добавить ордер.
func (bid *Bid) AddOrder(order *Order) error {
	if order.CostPerPiece > config.MaxPriceItem {
		return errors.New("Стоимость предмета слишком высока. ")
	}

	if order.Value > config.MaxValueItem {
		return errors.New("Слишком много ресурсов для продажи. ")
	}

	if order.CostPerPiece > order.CostPerPiece * config.PriceStepFromAverageCost {
		return errors.New("Стоимость предмета слишком высока. ")
	}

	if order.CostPerPiece < order.CostPerPiece / config.PriceStepFromAverageCost  {
		return errors.New("Стоимость предмета слишком низкая. ")
	}

	bid.rwMutex.RLock()
	list := bid.Orders[order.CostPerPiece]
	bid.rwMutex.RUnlock()

	if list == nil {
		bid.rwMutex.Lock()
		bid.Orders[order.CostPerPiece] = newOrders()
		bid.rwMutex.Unlock()

		bid.rwMutex.RLock()
		list = bid.Orders[order.CostPerPiece]
		bid.rwMutex.RUnlock()
	}

	bid.rwMutex.RLock()
	if bid.Orders[order.CostPerPiece].Len() > config.MaxNumbersOfOrdersPerPrice {
		bid.rwMutex.RUnlock()

		return errors.New("Слишком много товаров на эту стоимость. ")
	}
	bid.rwMutex.RUnlock()

	bid.rwMutex.Lock()
	bid.Orders[order.CostPerPiece].Add(order)
	bid.rwMutex.Unlock()

	bid.ValueOrders++

	var totalPrice float32
	for price, _ := range bid.Orders {
		totalPrice += float32(price)
	}

	bid.AveragePrice = totalPrice / float32(len(bid.Orders))

	bid.rwMutex.Lock()
	bid.Orders[order.CostPerPiece].ValueAllResources += uint64(order.Value)
	bid.rwMutex.Unlock()

	return nil
}

// GetOrders - получение ордеров.
func (bid *Bid) GetOrders(price float32) (*Orders, error) {
	bid.rwMutex.RLock()
	or, ok := bid.Orders[price]
	bid.rwMutex.RUnlock()

	if ok {
		return or, nil
	}

	return nil, errors.New("Ордера не найдены. ")
}

// GetOrderByID - получение ордера по ID продовца.
func (bid *Bid) GetOrderByID(id string, price float32) (*Order, error) {
	bid.rwMutex.RLock()
	ors, ok := bid.Orders[price]
	bid.rwMutex.RUnlock()

	if ok {
		for _, or := range ors.GetAll() {
			if or.(*Order).ID == id {
				return or.(*Order), nil
			}
		}
	}

	return nil, errors.New("Ордер не найден. ")
}

// SellOrder - продажа ресурса в ордере.
func (bid *Bid) SellOrder(price float32, value uint32) (error, map[string]uint32) {
	users := make(map[string]uint32)

	if price == 0.0 {
		return errors.New("Стоимость не может быть равна 0. "), users
	} else if value == 0 {
		return errors.New("Количество не может быть равно 0. "), users
	}

	bid.rwMutex.RLock()
	orders := bid.Orders[price]
	if orders == nil {
		return errors.New("Ордеров не найдено. "), users
	}
	bid.rwMutex.RUnlock()

	element := orders.Get()
	if element == nil {
		return errors.New("Нет ордеров. "), users
	}

	if uint64(value) > orders.ValueAllResources {
		return errors.New("Слишком большое количество требуемы ресурсов для продажи. "), users
	}

	resources := value

	for resources > 0 {
		element := orders.Get()
		or := element.Value.(*Order)

		if resources > or.Value {
			users[or.Owner.ID] = or.Value

			resources -= or.Value
			orders.Remove(element)

			continue
		}

		or.Value -= resources
		users[or.Owner.ID] = resources

		break
	}
	orders.ValueAllResources -= uint64(value)

	return nil, users
}

// GetAllOrders - получение всех ордеров.
func (bid *Bid) GetAllOrders() map[float32]*Orders {
	return bid.Orders
}
