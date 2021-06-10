package resources

import (
	"../../objects/resources"
	"../../user"
	"../../utils/config"
	"../../utils/logger"
	"errors"
)

// Asks - заявки продовцов на ресурсы.
type Asks struct {
	Database        *Database
	ListOnlineUsers *user.ListOnlineUsers

	food  *Ask
	wood  *Ask
	stone *Ask
	iron  *Ask
}

// newAsks - создание заявки продовцов на ресурсы.
func newAsks(listOnlineUsers *user.ListOnlineUsers) *Asks {
	asks := &Asks{
		Database:        NewDatabase(),
		ListOnlineUsers: listOnlineUsers,

		food:  newAsk(),
		wood:  newAsk(),
		stone: newAsk(),
		iron:  newAsk(),
	}

	asks.Database.GetAuctionResourcesAsk(asks)

	return asks
}

func (asks *Asks) Syncing() {
	asks.Database.SyncingAuctionResourcesAsk(asks)
}

// BuyFoodOrder - покупка ордера на еду.
func (asks *Asks) BuyFoodOrder(price float32, value uint32, u *user.User) error {
	_, err := asks.GetFoodOrderByID(u.ID, price)
	if err != nil {
		return errors.New("Нельзя купить еды за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Gold(price * float32(value)) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для покупки. ")
	}

	err, users := asks.food.BuyOrder(price, value)
	if err != nil {
		return err
	}

	u.Player.Resources.Gold -= resources.Gold(price * float32(value))
	u.Player.AuctionInventory.Resources.Food += resources.Food(value)

	for id, value := range users {
		if id == u.ID {
			continue
		}

		listOnlineUsers, err := asks.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			continue
		}

		auctionInventory, err := asks.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			err = asks.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				logger.Logger.Error(err.Error())
			}
		}
	}

	return nil
}

// BuyWoodOrder - покупка ордера на дерево.
func (asks *Asks) BuyWoodOrder(price float32, value uint32, u *user.User) error {
	_, err := asks.GetWoodOrderByID(u.ID, price)
	if err != nil {
		return errors.New("Нельзя купить дерево за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Gold(price * float32(value)) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для покупки. ")
	}

	err, users := asks.wood.BuyOrder(price, value)
	if err != nil {
		return err
	}

	u.Player.Resources.Gold -= resources.Gold(price * float32(value))
	u.Player.AuctionInventory.Resources.Wood += resources.Wood(value)

	for id, value := range users {
		listOnlineUsers, err := asks.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			continue
		}

		auctionInventory, err := asks.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			err = asks.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				logger.Logger.Error(err.Error())
			}
		}
	}

	return nil
}

// BuyStoneOrder - покупка ордера на камень.
func (asks *Asks) BuyStoneOrder(price float32, value uint32, u *user.User) error {
	_, err := asks.GetStoneOrderByID(u.ID, price)
	if err != nil {
		return errors.New("Нельзя купить камень за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Gold(price * float32(value)) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для покупки. ")
	}

	err, users := asks.stone.BuyOrder(price, value)
	if err != nil {
		return err
	}

	u.Player.Resources.Gold -= resources.Gold(price * float32(value))
	u.Player.AuctionInventory.Resources.Stone += resources.Stone(value)

	for id, value := range users {
		if id == u.ID {
			continue
		}

		listOnlineUsers, err := asks.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			continue
		}

		auctionInventory, err := asks.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			err = asks.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				logger.Logger.Error(err.Error())
			}
		}
	}

	return nil
}

// BuyIronOrder - покупка ордера на железо.
func (asks *Asks) BuyIronOrder(price float32, value uint32, u *user.User) error {
	_, err := asks.GetIronOrderByID(u.ID, price)
	if err != nil {
		return errors.New("Нельзя купить железо за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Gold(price * float32(value)) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для покупки. ")
	}

	err, users := asks.iron.BuyOrder(price, value)
	if err != nil {
		return err
	}

	u.Player.Resources.Gold -= resources.Gold(price * float32(value))
	u.Player.AuctionInventory.Resources.Iron += resources.Iron(value)

	for id, value := range users {
		if id == u.ID {
			continue
		}

		listOnlineUsers, err := asks.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			continue
		}

		auctionInventory, err := asks.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

			err = asks.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				logger.Logger.Error(err.Error())
			}
		}
	}

	return nil
}

// AddFoodOrder - выставить ордер на еду.
func (asks *Asks) AddFoodOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Food(value) > u.Player.Resources.Food {
		return errors.New("У игрока недостаточно еды для ордера. ")
	}

	u.Player.Resources.Food -= resources.Food(value)

	or, err := asks.food.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return asks.food.AddOrder(newOrder)
	}

	or.Value += value
	asks.food.rwMutex.Lock()
	asks.food.Orders[costPerPiece].ValueAllResources += uint64(value)
	asks.food.rwMutex.Unlock()

	return nil
}

// AddWoodOrder - выставить ордер на дерево.
func (asks *Asks) AddWoodOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Wood(value) > u.Player.Resources.Wood {
		return errors.New("У игрока недостаточно дерева для ордера. ")
	}

	u.Player.Resources.Wood -= resources.Wood(value)

	or, err := asks.wood.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return asks.wood.AddOrder(newOrder)
	}

	or.Value += value
	asks.wood.rwMutex.Lock()
	asks.wood.Orders[costPerPiece].ValueAllResources += uint64(value)
	asks.wood.rwMutex.Unlock()

	return nil
}

// AddStoneOrder - выставить ордер на камень.
func (asks *Asks) AddStoneOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Stone(value) > u.Player.Resources.Stone {
		return errors.New("У игрока недостаточно камня для ордера. ")
	}

	u.Player.Resources.Stone -= resources.Stone(value)

	or, err := asks.stone.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return asks.stone.AddOrder(newOrder)
	}

	or.Value += value
	asks.stone.rwMutex.Lock()
	asks.stone.Orders[costPerPiece].ValueAllResources += uint64(value)
	asks.stone.rwMutex.Unlock()

	return nil
}

// AddIronOrder - выставить ордер на железо.
func (asks *Asks) AddIronOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Iron(value) > u.Player.Resources.Iron {
		return errors.New("У игрока недостаточно железа для ордера. ")
	}

	u.Player.Resources.Iron -= resources.Iron(value)

	or, err := asks.iron.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return asks.iron.AddOrder(newOrder)
	}

	or.Value += value
	asks.iron.rwMutex.Lock()
	asks.iron.Orders[costPerPiece].ValueAllResources += uint64(value)
	asks.iron.rwMutex.Unlock()

	return nil
}

// GetFoodOrderByID - получение ордера на еду пользователя.
func (asks *Asks) GetFoodOrderByID(id string, price float32) (*Order, error) {
	return asks.food.GetOrderByID(id, price)
}

// GetWoodOrderByID - получение ордера на дерево пользователя.
func (asks *Asks) GetWoodOrderByID(id string, price float32) (*Order, error) {
	return asks.wood.GetOrderByID(id, price)
}

// GetStoneOrderByID - получение ордера на камень пользователя.
func (asks *Asks) GetStoneOrderByID(id string, price float32) (*Order, error) {
	return asks.stone.GetOrderByID(id, price)
}

// GetIronOrderByID - получение ордера на железо пользователя.
func (asks *Asks) GetIronOrderByID(id string, price float32) (*Order, error) {
	return asks.iron.GetOrderByID(id, price)
}

// GetFoodOrders - получение ордеров на еду.
func (asks *Asks) GetFoodOrders(price float32) (*Orders, error) {
	return asks.food.GetOrders(price)
}

// GetWoodOrders - получение ордеров на дерево.
func (asks *Asks) GetWoodOrders(price float32) (*Orders, error) {
	return asks.wood.GetOrders(price)
}

// GetStoneOrders - получение ордеров на камень.
func (asks *Asks) GetStoneOrders(price float32) (*Orders, error) {
	return asks.stone.GetOrders(price)
}

// GetIronOrders - получение ордеров на железо.
func (asks *Asks) GetIronOrders(price float32) (*Orders, error) {
	return asks.iron.GetOrders(price)
}

// GetAllFoodOrders - получение всех ордеров на еду.
func (asks *Asks) GetAllFoodOrders() map[float32]*Orders {
	return asks.food.GetAllOrders()
}

// GetAllWoodOrders - получение всех ордеров на дерево.
func (asks *Asks) GetAllWoodOrders() map[float32]*Orders {
	return asks.wood.GetAllOrders()
}

// GetAllStoneOrders - получение всех ордеров на камень.
func (asks *Asks) GetAllStoneOrders() map[float32]*Orders {
	return asks.stone.GetAllOrders()
}

// GetAllIronOrders - получение всех ордеров на железо.
func (asks *Asks) GetAllIronOrders() map[float32]*Orders {
	return asks.iron.GetAllOrders()
}

// GetAveragePriceFoodOrders - получение средней стоимости еды.
func (asks *Asks) GetAveragePriceFoodOrders() float32 {
	return asks.food.AveragePrice
}

// GetAveragePriceWoodOrders - получение средней стоимости дерева.
func (asks *Asks) GetAveragePriceWoodOrders() float32 {
	return asks.wood.AveragePrice
}

// GetAveragePriceStoneOrders - получение средней стоимости камня.
func (asks *Asks) GetAveragePriceStoneOrders() float32 {
	return asks.stone.AveragePrice
}

// GetAveragePriceIronOrders - получение средней стоимости железа.
func (asks *Asks) GetAveragePriceIronOrders() float32 {
	return asks.iron.AveragePrice
}

// GetValueOrdersFoodOrders - получение кол-во ордеров на еды.
func (asks *Asks) GetValueOrdersFoodOrders() uint32 {
	return asks.food.ValueOrders
}

// GetValueOrdersWoodOrders - получение кол-во ордеров на дерево.
func (asks *Asks) GetValueOrdersWoodOrders() uint32 {
	return asks.wood.ValueOrders
}

// GetValueOrdersStoneOrders - получение кол-во ордеров на камень.
func (asks *Asks) GetValueOrdersStoneOrders() uint32 {
	return asks.stone.ValueOrders
}

// GetValueOrdersIronOrders - получение кол-во ордеров на железо.
func (asks *Asks) GetValueOrdersIronOrders() uint32 {
	return asks.iron.ValueOrders
}