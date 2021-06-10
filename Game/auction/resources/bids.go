package resources

import (
	"../../objects/resources"
	"../../user"
	"../../utils/config"
	"errors"
)

// Bids - заявки покупателей на ресурсы.
type Bids struct {
	Database        *Database
	ListOnlineUsers *user.ListOnlineUsers

	food  *Bid
	wood  *Bid
	stone *Bid
	iron  *Bid
}

// newBids - создание заявок покупателей  на ресурсы.
func newBids(listOnlineUsers *user.ListOnlineUsers) *Bids {
	bids := &Bids{
		Database:        NewDatabase(),
		ListOnlineUsers: listOnlineUsers,

		food:  newBid(),
		wood:  newBid(),
		stone: newBid(),
		iron:  newBid(),
	}

	bids.Database.GetAuctionResourcesBid(bids)

	return bids
}

// Syncing - синхронизировать заявки покупателей.
func (bids *Bids) Syncing() {
	bids.Database.SyncingAuctionResourcesBid(bids)
}

// SellFoodOrder - продажа еды в ордер.
func (bids *Bids) SellFoodOrder(price float32, value uint32, u *user.User) error {
	_, err := bids.GetFoodOrderByID(u.ID, price)
	if err == nil {
		return errors.New("Нельзя продать еду за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Food(price * float32(value)) > u.Player.Resources.Food {
		return errors.New("У игрока недостаточно еды для продажи. ")
	}

	err, users := bids.food.SellOrder(price, value)
	if err != nil {
		return err
	}

	for id, value := range users {
		listOnlineUsers, err := bids.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Food += resources.Food(value)

			continue
		}

		auctionInventory, err := bids.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			return err
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Food += resources.Food(value)

			err = bids.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				return err
			}
		}
	}

	u.Player.Resources.Food -= resources.Food(value)
	u.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

	return nil
}

// SellWoodOrder - продажа дерево в ордер.
func (bids *Bids) SellWoodOrder(price float32, value uint32, u *user.User) error {
	_, err := bids.GetWoodOrderByID(u.ID, price)
	if err == nil {
		return errors.New("Нельзя продать дерево за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Wood(price * float32(value)) > u.Player.Resources.Wood {
		return errors.New("У игрока недостаточно дерева для продажи. ")
	}

	err, users := bids.wood.SellOrder(price, value)
	if err != nil {
		return err
	}

	for id, value := range users {
		listOnlineUsers, err := bids.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Wood += resources.Wood(value)

			continue
		}

		auctionInventory, err := bids.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			return err
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Wood += resources.Wood(value)

			err = bids.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				return err
			}
		}
	}

	u.Player.Resources.Wood -= resources.Wood(value)
	u.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

	return nil
}

// SellStoneOrder - продажа камня в ордере.
func (bids *Bids) SellStoneOrder(price float32, value uint32, u *user.User) error {
	_, err := bids.GetStoneOrderByID(u.ID, price)
	if err == nil {
		return errors.New("Нельзя продать камень за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Stone(price * float32(value)) > u.Player.Resources.Stone {
		return errors.New("У игрока недостаточно камня для продажи. ")
	}

	err, users := bids.stone.SellOrder(price, value)
	if err != nil {
		return err
	}

	for id, value := range users {
		listOnlineUsers, err := bids.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Stone += resources.Stone(value)

			continue
		}

		auctionInventory, err := bids.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			return err
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Stone += resources.Stone(value)

			err = bids.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				return err
			}
		}
	}

	u.Player.Resources.Stone -= resources.Stone(value)
	u.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

	return nil
}

// SellIronOrder - продажа железо в ордере.
func (bids *Bids) SellIronOrder(price float32, value uint32, u *user.User) error {
	_, err := bids.GetIronOrderByID(u.ID, price)
	if err == nil  {
		return errors.New("Нельзя продать желео за эту стоимость пока у вас стоит ордер. ")
	}

	if resources.Iron(price * float32(value)) > u.Player.Resources.Iron {
		return errors.New("У игрока недостаточно железа для продажи. ")
	}

	err, users := bids.iron.SellOrder(price, value)
	if err != nil {
		return err
	}

	for id, value := range users {
		listOnlineUsers, err := bids.ListOnlineUsers.Get(id)
		if err != nil {
			return err
		}

		if listOnlineUsers != nil {
			listOnlineUsers.Player.AuctionInventory.Resources.Iron += resources.Iron(value)

			continue
		}

		auctionInventory, err := bids.Database.GetUserAuctionInventoryByID(id)
		if err != nil {
			return err
		}

		if auctionInventory != nil {
			auctionInventory.Resources.Iron += resources.Iron(value)

			err = bids.Database.UpdateUserAuctionInventoryByID(auctionInventory)
			if err != nil {
				return err
			}
		}
	}

	u.Player.Resources.Iron -= resources.Iron(value)
	u.Player.AuctionInventory.Resources.Gold += resources.Gold(price * float32(value)) - (resources.Gold(price * float32(value) / 100 * config.TaxOnAuctionItems))

	return nil
}

// AddFoodOrder - выставить ордер на еду.
func (bids *Bids) AddFoodOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Gold(float32(value) * costPerPiece) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для ордера. ")
	}

	u.Player.Resources.Gold -= resources.Gold(float32(value) * costPerPiece)

	or, err := bids.food.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return bids.food.AddOrder(newOrder)
	}

	or.Value += value
	bids.food.rwMutex.Lock()
	bids.food.Orders[costPerPiece].ValueAllResources += uint64(value)
	bids.food.rwMutex.Unlock()

	return nil
}

// AddWoodOrder - выставить ордер на дерево.
func (bids *Bids) AddWoodOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Gold(float32(value) * costPerPiece) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для ордера. ")
	}

	u.Player.Resources.Gold -= resources.Gold(float32(value) * costPerPiece)

	or, err := bids.wood.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return bids.wood.AddOrder(newOrder)
	}

	or.Value += value
	bids.wood.rwMutex.Lock()
	bids.wood.Orders[costPerPiece].ValueAllResources += uint64(value)
	bids.wood.rwMutex.Unlock()

	return nil
}

// AddStoneOrder - выставить ордер на камень.
func (bids *Bids) AddStoneOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Gold(float32(value) * costPerPiece) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для ордера. ")
	}

	u.Player.Resources.Gold -= resources.Gold(float32(value) * costPerPiece)

	or, err := bids.stone.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return bids.stone.AddOrder(newOrder)
	}

	or.Value += value
	bids.stone.rwMutex.Lock()
	bids.stone.Orders[costPerPiece].ValueAllResources += uint64(value)
	bids.stone.rwMutex.Unlock()

	return nil
}

// AddIronOrder - выставить ордер на железо.
func (bids *Bids) AddIronOrder(costPerPiece float32, value uint32, u *user.User) error {
	if resources.Gold(float32(value) * costPerPiece) > u.Player.Resources.Gold {
		return errors.New("У игрока недостаточно золота для ордера. ")
	}

	u.Player.Resources.Gold -= resources.Gold(float32(value) * costPerPiece)

	or, err := bids.iron.GetOrderByID(u.ID, costPerPiece)
	if err != nil {
		newOrder := newOrder(costPerPiece, value, Owner{ID: u.ID, NickName: u.Nickname})

		return bids.iron.AddOrder(newOrder)
	}

	or.Value += value
	bids.iron.rwMutex.Lock()
	bids.iron.Orders[costPerPiece].ValueAllResources += uint64(value)
	bids.iron.rwMutex.Unlock()

	return nil
}

// GetFoodOrderByID - получение ордера на еду пользователя.
func (bids *Bids) GetFoodOrderByID(id string, price float32) (*Order, error) {
	return bids.food.GetOrderByID(id, price)
}

// GetWoodOrderByID - получение ордера на дерево пользователя.
func (bids *Bids) GetWoodOrderByID(id string, price float32) (*Order, error) {
	return bids.wood.GetOrderByID(id, price)
}

// GetStoneOrderByID - получение ордера на камень пользователя.
func (bids *Bids) GetStoneOrderByID(id string, price float32) (*Order, error) {
	return bids.stone.GetOrderByID(id, price)
}

// GetIronOrderByID - получение ордера на железо пользователя.
func (bids *Bids) GetIronOrderByID(id string, price float32) (*Order, error) {
	return bids.iron.GetOrderByID(id, price)
}

// GetFoodOrders - получение ордеров на еду.
func (bids *Bids) GetFoodOrders(price float32) (*Orders, error) {
	return bids.food.GetOrders(price)
}

// GetWoodOrders - получение ордеров на дерево.
func (bids *Bids) GetWoodOrders(price float32) (*Orders, error) {
	return bids.wood.GetOrders(price)
}

// GetStoneOrders - получение ордеров на камень.
func (bids *Bids) GetStoneOrders(price float32) (*Orders, error) {
	return bids.stone.GetOrders(price)
}

// GetIronOrders - получение ордеров на железо.
func (bids *Bids) GetIronOrders(price float32) (*Orders, error) {
	return bids.iron.GetOrders(price)
}

// GetAllFoodOrders - получение всех ордеров на еду.
func (bids *Bids) GetAllFoodOrders() map[float32]*Orders {
	return bids.food.GetAllOrders()
}

// GetAllWoodOrders - получение всех ордеров на дерево.
func (bids *Bids) GetAllWoodOrders() map[float32]*Orders {
	return bids.wood.GetAllOrders()
}

// GetAllStoneOrders - получение всех ордеров на камень.
func (bids *Bids) GetAllStoneOrders() map[float32]*Orders {
	return bids.stone.GetAllOrders()
}

// GetAllIronOrders - получение всех ордеров на железо.
func (bids *Bids) GetAllIronOrders() map[float32]*Orders {
	return bids.iron.GetAllOrders()
}

// GetAveragePriceFoodOrders - получение средней стоимости еды.
func (bids *Bids) GetAveragePriceFoodOrders() float32 {
	return bids.food.AveragePrice
}

// GetAveragePriceWoodOrders - получение средней стоимости дерева.
func (bids *Bids) GetAveragePriceWoodOrders() float32 {
	return bids.wood.AveragePrice
}

// GetAveragePriceStoneOrders - получение средней стоимости камня.
func (bids *Bids) GetAveragePriceStoneOrders() float32 {
	return bids.stone.AveragePrice
}

// GetAveragePriceIronOrders - получение средней стоимости железа.
func (bids *Bids) GetAveragePriceIronOrders() float32 {
	return bids.iron.AveragePrice
}

// GetValueOrdersFoodOrders - получение кол-во ордеров на еды.
func (bids *Bids) GetValueOrdersFoodOrders() uint32 {
	return bids.food.ValueOrders
}

// GetValueOrdersWoodOrders - получение кол-во ордеров на дерево.
func (bids *Bids) GetValueOrdersWoodOrders() uint32 {
	return bids.wood.ValueOrders
}

// GetValueOrdersStoneOrders - получение кол-во ордеров на камень.
func (bids *Bids) GetValueOrdersStoneOrders() uint32 {
	return bids.stone.ValueOrders
}

// GetValueOrdersIronOrders - получение кол-во ордеров на железо.
func (bids *Bids) GetValueOrdersIronOrders() uint32 {
	return bids.iron.ValueOrders
}