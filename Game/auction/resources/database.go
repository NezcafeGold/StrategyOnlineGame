package resources

import (
	"../../player/auction_inventory"
	"../../utils/config"
	"../../utils/logger"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
)

// Database - база данных.
type Database struct {
	Session  *mgo.Session  // Сессия бд
	Database *mgo.Database // Бд
}

// NewDatabase - создание базы данных.
func NewDatabase() *Database {
	sessionDB, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	db := sessionDB.DB(config.Database)
	if db == nil {
		logger.Logger.Fatal("Ошибка подключения к \"" + config.Database + "\".")
	}

	return &Database{
		Session:  sessionDB,
		Database: db,
	}
}

// GetUserAuctionInventoryByID - получить инвентарь аукциона пользователя из базы данных по ID.
func (database *Database) GetUserAuctionInventoryByID(id string) (*auction_inventory.AuctionInventory, error) {
	ai := auction_inventory.New(id)

	// Получение инвентаря аукциона
	{
		collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
		query := bson.M{
			"_id" : ai.ID,
		}
		err := collectionAuctionInventory.Find(query).One(ai)
		if err != nil{
			return nil, err
		}

		ai.RwMutex = new(sync.RWMutex)
	}

	return ai, nil
}

// UpdateUserAuctionInventoryByID - обновить инвентарь аукциона пользователя.
func (database *Database) UpdateUserAuctionInventoryByID(auctionInventory *auction_inventory.AuctionInventory) error {
	// Сохранение инвентаря аукциона
	{
		collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
		query := bson.M{
			"_id" : auctionInventory.ID,
		}
		err := collectionAuctionInventory.Update(query, auctionInventory)
		if err != nil{
			return err
		}
	}

	return nil
}

// SyncingAuctionResourcesBid - синхронизировать Bid аукциона ресурсов.
func (database *Database) SyncingAuctionResourcesBid(bids *Bids) {
	collectionBidAuction := database.Database.C(config.CollectionBidAuctionResources)

	wg := new(sync.WaitGroup)

	wg.Add(4)
	
	// Сохранение еды.
	go func(){
		defer wg.Done()
		
		bid := &BidModel{
			Type:         "food",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  bids.food.ValueOrders,
			AveragePrice: bids.food.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range bids.GetAllFoodOrders() {
			bid.rwMutex.Lock()
			bid.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			bid.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": bid.Type,
		}

		err := collectionBidAuction.Update(query, bid)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение дерева.
	go func(){
		defer wg.Done()
		
		bid := &BidModel{
			Type:         "wood",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  bids.wood.ValueOrders,
			AveragePrice: bids.wood.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range bids.GetAllWoodOrders() {
			bid.rwMutex.Lock()
			bid.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			bid.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": bid.Type,
		}

		err := collectionBidAuction.Update(query, bid)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение камня.
	go func(){
		defer wg.Done()

		bid := &BidModel{
			Type:         "stone",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  bids.stone.ValueOrders,
			AveragePrice: bids.stone.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range bids.GetAllStoneOrders() {
			bid.rwMutex.Lock()
			bid.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			bid.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": bid.Type,
		}

		err := collectionBidAuction.Update(query, bid)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение железа.
	go func(){
		defer wg.Done()

		bid := &BidModel{
			Type:         "iron",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  bids.iron.ValueOrders,
			AveragePrice: bids.iron.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range bids.GetAllIronOrders() {
			bid.rwMutex.Lock()
			bid.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			bid.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": bid.Type,
		}

		err := collectionBidAuction.Update(query, bid)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	wg.Wait()
}

// GetAuctionResourcesBid - получение Bid аукциона ресурсов.
func (database *Database) GetAuctionResourcesBid(bids *Bids) {
	collectionBidAuction := database.Database.C(config.CollectionBidAuctionResources)
	
	wg := new(sync.WaitGroup)
	
	wg.Add(4)
	
	// Получение еды.
	go func(){
		defer wg.Done()
		
		modelFood := new(BidModel)
		query := bson.M{
			"_id" : "food",
		}

		err := collectionBidAuction.Find(query).One(modelFood)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		bids.food.AveragePrice = modelFood.AveragePrice
		bids.food.ValueOrders = modelFood.ValueOrders
		bids.food.rwMutex = new(sync.RWMutex)
		bids.food.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelFood.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			bids.iron.rwMutex.Lock()
			bids.iron.Orders[float32(price)] = ordersList
			bids.iron.rwMutex.Unlock()
		}
	}()

	// Получение дерева.
	go func(){
		defer wg.Done()
		
		modelWood := new(BidModel)
		query := bson.M{
			"_id" : "wood",
		}

		err := collectionBidAuction.Find(query).One(modelWood)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		bids.wood.AveragePrice = modelWood.AveragePrice
		bids.wood.ValueOrders = modelWood.ValueOrders
		bids.wood.rwMutex = new(sync.RWMutex)
		bids.wood.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelWood.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			bids.iron.rwMutex.Lock()
			bids.iron.Orders[float32(price)] = ordersList
			bids.iron.rwMutex.Unlock()
		}
	}()

	// Получение камня.
	go func(){
		defer wg.Done()
		
		modelStone := new(BidModel)
		query := bson.M{
			"_id" : "stone",
		}

		err := collectionBidAuction.Find(query).One(modelStone)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		bids.stone.AveragePrice = modelStone.AveragePrice
		bids.stone.ValueOrders = modelStone.ValueOrders
		bids.stone.rwMutex = new(sync.RWMutex)
		bids.stone.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelStone.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			bids.iron.rwMutex.Lock()
			bids.iron.Orders[float32(price)] = ordersList
			bids.iron.rwMutex.Unlock()
		}
	}()

	// Получение железа.
	go func(){
		defer wg.Done()
		
		modelIron := new(BidModel)
		query := bson.M{
			"_id" : "iron",
		}

		err := collectionBidAuction.Find(query).One(modelIron)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		bids.iron.AveragePrice = modelIron.AveragePrice
		bids.iron.ValueOrders = modelIron.ValueOrders
		bids.iron.rwMutex = new(sync.RWMutex)
		bids.iron.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelIron.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			bids.iron.rwMutex.Lock()
			bids.iron.Orders[float32(price)] = ordersList
			bids.iron.rwMutex.Unlock()
		}
	}()

	wg.Wait()
}

// SyncingAuctionResourcesAsk - синхронизировать Ask аукциона ресурсов.
func (database *Database) SyncingAuctionResourcesAsk(asks *Asks) {
	collectionAskAuction := database.Database.C(config.CollectionAskAuctionResources)
	wg := new(sync.WaitGroup)

	wg.Add(4)

	// Сохранение еды.
	go func(){
		defer wg.Done()

		ask := &AskModel{
			Type:         "food",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  asks.food.ValueOrders,
			AveragePrice: asks.food.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range asks.GetAllFoodOrders() {
			ask.rwMutex.Lock()
			ask.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			ask.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": ask.Type,
		}

		err := collectionAskAuction.Update(query, ask)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение дерева.
	go func(){
		defer wg.Done()

		ask := &AskModel{
			Type:         "wood",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  asks.wood.ValueOrders,
			AveragePrice: asks.wood.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range asks.GetAllWoodOrders() {
			ask.rwMutex.Lock()
			ask.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			ask.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": ask.Type,
		}

		err := collectionAskAuction.Update(query, ask)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение камня.
	go func(){
		defer wg.Done()

		ask := &AskModel{
			Type:         "stone",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  asks.stone.ValueOrders,
			AveragePrice: asks.stone.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range asks.GetAllStoneOrders() {
			ask.rwMutex.Lock()
			ask.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			ask.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": ask.Type,
		}

		err := collectionAskAuction.Update(query, ask)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	// Сохранение железа.
	go func(){
		defer wg.Done()

		ask := &AskModel{
			Type:         "iron",
			Orders:       make(map[string]*OrdersModel),
			ValueOrders:  asks.iron.ValueOrders,
			AveragePrice: asks.iron.AveragePrice,

			rwMutex: new(sync.RWMutex),
		}

		for price, orders := range asks.GetAllIronOrders() {
			ask.rwMutex.Lock()
			ask.Orders[fmt.Sprintf("%.1f", price)] = &OrdersModel{
				Orders:            orders.GetAll(),
				ValueAllResources: orders.ValueAllResources,
			}
			ask.rwMutex.Unlock()
		}

		query := bson.M{
			"_id": ask.Type,
		}

		err := collectionAskAuction.Update(query, ask)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}
	}()

	wg.Wait()
}

// GetAuctionResourcesAsk - получение Ask аукциона ресурсов.
func (database *Database) GetAuctionResourcesAsk(asks *Asks) {
	collectionAskAuction := database.Database.C(config.CollectionAskAuctionResources)

	wg := new(sync.WaitGroup)

	wg.Add(4)

	// Получение еды.
	go func(){
		defer wg.Done()

		modelFood := new(AskModel)
		query := bson.M{
			"_id" : "food",
		}

		err := collectionAskAuction.Find(query).One(modelFood)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		asks.food.AveragePrice = modelFood.AveragePrice
		asks.food.ValueOrders = modelFood.ValueOrders
		asks.food.rwMutex = new(sync.RWMutex)
		asks.food.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelFood.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			asks.iron.rwMutex.Lock()
			asks.iron.Orders[float32(price)] = ordersList
			asks.iron.rwMutex.Unlock()
		}
	}()

	// Получение дерева.
	go func(){
		defer wg.Done()

		modelWood := new(AskModel)
		query := bson.M{
			"_id" : "wood",
		}

		err := collectionAskAuction.Find(query).One(modelWood)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		asks.wood.AveragePrice = modelWood.AveragePrice
		asks.wood.ValueOrders = modelWood.ValueOrders
		asks.wood.rwMutex = new(sync.RWMutex)
		asks.wood.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelWood.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			asks.iron.rwMutex.Lock()
			asks.iron.Orders[float32(price)] = ordersList
			asks.iron.rwMutex.Unlock()
		}
	}()

	// Получение камня.
	go func(){
		defer wg.Done()

		modelStone := new(AskModel)
		query := bson.M{
			"_id" : "stone",
		}

		err := collectionAskAuction.Find(query).One(modelStone)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		asks.stone.AveragePrice = modelStone.AveragePrice
		asks.stone.ValueOrders = modelStone.ValueOrders
		asks.stone.rwMutex = new(sync.RWMutex)
		asks.stone.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelStone.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			asks.iron.rwMutex.Lock()
			asks.iron.Orders[float32(price)] = ordersList
			asks.iron.rwMutex.Unlock()
		}
	}()

	// Получение железа.
	go func(){
		defer wg.Done()

		modelIron := new(AskModel)
		query := bson.M{
			"_id" : "iron",
		}

		err := collectionAskAuction.Find(query).One(modelIron)
		if err != nil{
			logger.Logger.Fatal(err.Error())
		}

		asks.iron.AveragePrice = modelIron.AveragePrice
		asks.iron.ValueOrders = modelIron.ValueOrders
		asks.iron.rwMutex = new(sync.RWMutex)
		asks.iron.Orders = make(map[float32]*Orders)

		for priceStr, orders := range modelIron.Orders {
			ordersList := newOrders()

			for _, order := range orders.Orders {
				ordersList.Add(order)
			}
			ordersList.ValueAllResources = orders.ValueAllResources

			price, err := strconv.ParseFloat(priceStr, 32)
			if err != nil {
				logger.Logger.Fatal(err.Error())
			}

			asks.iron.rwMutex.Lock()
			asks.iron.Orders[float32(price)] = ordersList
			asks.iron.rwMutex.Unlock()
		}
	}()

	wg.Wait()
}