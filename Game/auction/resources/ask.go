package resources

import (
	"../../utils/config"
	"errors"
	"sync"
)

// Ask - заявки продовцов.
type Ask struct {
	Orders       map[float32]*Orders `json:"orders"          bson:"orders"`        // Список ордеров.
	ValueOrders  uint32              `json:"value_orders"    bson:"value_orders"`  // Кол-во ордеров.
	AveragePrice float32             `json:"average_price"   bson:"average_price"` // Средняя стоимость.

	rwMutex *sync.RWMutex // RWMutex
}

// newAsk - создание заявок продовцов.
func newAsk() *Ask {
	return &Ask{
		Orders:       make(map[float32]*Orders),
		AveragePrice: 0,
		ValueOrders:  0,
		rwMutex:      new(sync.RWMutex),
	}
}

// AddOrder - добавить ордер.
func (ask *Ask) AddOrder(order *Order) error {
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

	ask.rwMutex.RLock()
	list := ask.Orders[order.CostPerPiece]
	ask.rwMutex.RUnlock()

	if list == nil {
		ask.rwMutex.Lock()
		ask.Orders[order.CostPerPiece] = newOrders()
		ask.rwMutex.Unlock()

		ask.rwMutex.RLock()
		list = ask.Orders[order.CostPerPiece]
		ask.rwMutex.RUnlock()
	}

	ask.rwMutex.RLock()
	if ask.Orders[order.CostPerPiece].Len() > config.MaxNumbersOfOrdersPerPrice {
		ask.rwMutex.RUnlock()

		return errors.New("Слишком много товаров на эту стоимость. ")
	}
	ask.rwMutex.RUnlock()

	ask.rwMutex.Lock()
	ask.Orders[order.CostPerPiece].Add(order)
	ask.rwMutex.Unlock()

	ask.ValueOrders++

	var totalPrice float32
	for price, _ := range ask.Orders {
		totalPrice += price
	}

	ask.AveragePrice = totalPrice / float32(len(ask.Orders))

	ask.rwMutex.Lock()
	ask.Orders[order.CostPerPiece].ValueAllResources += uint64(order.Value)
	ask.rwMutex.Unlock()

	return nil
}

// GetOrders - получение ордеров.
func (ask *Ask) GetOrders(price float32) (*Orders, error) {
	ask.rwMutex.RLock()
	or, ok := ask.Orders[price]
	ask.rwMutex.RUnlock()

	if ok {
		return or, nil
	}

	return nil, errors.New("Ордера не найдены. ")
}

// GetOrderByID - получение ордера по ID продовца.
func (ask *Ask) GetOrderByID(id string, price float32) (*Order, error) {
	ask.rwMutex.RLock()
	ors, ok := ask.Orders[price]
	ask.rwMutex.RUnlock()

	if ok {
		for _, or := range ors.GetAll() {
			if or.(Order).ID == id {
				return or.(*Order), nil
			}
		}
	}

	return nil, errors.New("Ордер не найден. ")
}

// BuyOrder - покупка ордера.
func (ask *Ask) BuyOrder(price float32, value uint32) (error, map[string]uint32) {
	users := make(map[string]uint32)

	if price == 0.0 {
		return errors.New("Стоимость не может быть равна 0. "), users
	} else if value == 0 {
		return errors.New("Количество не может быть равно 0. "), users
	}

	ask.rwMutex.RLock()
	orders := ask.Orders[price]
	if orders == nil {
		return errors.New("Ордеров не найдено. "), users
	}
	ask.rwMutex.RUnlock()

	element := orders.Get()
	if element == nil {
		return errors.New("Нет ордеров. "), users
	}

	if uint64(value) > orders.ValueAllResources {
		return errors.New("Слишком большое количество требуемы ресурсов для покупки. "), users
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
func (ask *Ask) GetAllOrders() map[float32]*Orders {
	return ask.Orders
}
