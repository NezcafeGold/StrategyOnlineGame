package user

import (
	"../objects/chests"
	"../objects/equipments"
	"../objects/materials"
	"../objects/summoning_scrolls"
	"../objects/units"
	"../player/auction_inventory"
	"../utils/config"
	"../utils/logger"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// AddUser - добавить пользователя в базу данных.
func (database *Database) AddUser(u User) error {
	u_ := new(User)

	collectionUsers := database.Database.C(config.CollectionUsers)
	queryNickname := bson.M{
		"nickname" : u.Nickname,
	}
	queryEmail := bson.M{
		"email" : u.Email,
	}

	errNickname := collectionUsers.Find(queryNickname).One(u_)
	errEmail := collectionUsers.Find(queryEmail).One(u_)
	if errNickname != nil && errEmail != nil {
		if u_.Email == "" {
			// Сохранение инвентаря
			{
				collectionInventory := database.Database.C(config.CollectionInventory)
				err := collectionInventory.Insert(u.Player.Inventory)
				if err != nil {
					return err
				}
			}

			// Сохранение юнитов
			{
				collectionUnits := database.Database.C(config.CollectionUnits)
				err := collectionUnits.Insert(u.Player.Units)
				if err != nil {
					return err
				}
			}

			// Сохранение инвентаря аукциона
			{
				collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
				err := collectionAuctionInventory.Insert(u.Player.AuctionInventory)
				if err != nil{
					return err
				}
			}

			// Сохронение пользователя
			{
				resetToZeroUserMap(&u)

				err := collectionUsers.Insert(u)
				if err != nil {
					return err
				}
			}

			return nil
		}
	}

	return errors.New("Пользователь уже существует. ")
}

// GetUserByID - получить пользователя из базы данных по ID.
func (database *Database) GetUserByID(id string) (*User, error) {
	u := new(User)

	// Получение пользователя
	{
		collectionUsers := database.Database.C(config.CollectionUsers)
		query := bson.M{
			"_id": id,
		}
		err := collectionUsers.Find(query).One(u)
		if err != nil {
			return nil, err
		}
	}

	// Получение инвентаря
	{
		collectionInventory := database.Database.C(config.CollectionInventory)
		query := bson.M{
			"_id" : u.Player.Inventory.ID,
		}
		err := collectionInventory.Find(query).One(u.Player.Inventory)
		if err != nil{
			return nil, err
		}

		u.Player.Inventory.RwMutex = new(sync.RWMutex)
	}

	// Получение юнитов
	{
		collectionUnits := database.Database.C(config.CollectionUnits)
		query := bson.M{
			"_id" : u.Player.Units.ID,
		}
		err := collectionUnits.Find(query).One(u.Player.Units)
		if err != nil{
			return nil, err
		}

		u.Player.Units.RwMutex = new(sync.RWMutex)
	}

	// Получение инвентаря аукциона
	{
		collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
		query := bson.M{
			"_id" : u.ID,
		}
		err := collectionAuctionInventory.Find(query).One(u.Player.AuctionInventory)
		if err != nil{
			return nil, err
		}

		u.Player.AuctionInventory.RwMutex = new(sync.RWMutex)
	}

	return u, nil
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

// GetUser - получить пользователя из базы данных.
func (database *Database) GetUser(email string) (*User, error) {
	u := new(User)

	// Получение пользователя
	{
		collectionUsers := database.Database.C(config.CollectionUsers)
		query := bson.M{
			"email": email,
		}
		err := collectionUsers.Find(query).One(u)
		if err != nil {
			return nil, err
		}
	}

	// Получение инвентаря
	{
		collectionInventory := database.Database.C(config.CollectionInventory)
		query := bson.M{
			"_id" : u.Player.Inventory.ID,
		}
		err := collectionInventory.Find(query).One(u.Player.Inventory)
		if err != nil{
			return nil, err
		}

		u.Player.Inventory.RwMutex = new(sync.RWMutex)
	}

	// Получение юнитов
	{
		collectionUnits := database.Database.C(config.CollectionUnits)
		query := bson.M{
			"_id" : u.Player.Units.ID,
		}
		err := collectionUnits.Find(query).One(u.Player.Units)
		if err != nil{
			return nil, err
		}

		u.Player.Units.RwMutex = new(sync.RWMutex)
	}

	// Получение инвентаря аукциона
	{
		collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
		query := bson.M{
			"_id" : u.Player.AuctionInventory.ID,
		}
		err := collectionAuctionInventory.Find(query).One(u.Player.AuctionInventory)
		if err != nil{
			return nil, err
		}

		u.Player.AuctionInventory.RwMutex = new(sync.RWMutex)
	}

	return u, nil
}

// SyncingUser - синхронизировать данные пользователя.
func (database *Database) SyncingUser(u User) error {
	u.Syncing()

	// Сохранение инвентаря
	{
		collectionInventory := database.Database.C(config.CollectionInventory)
		query := bson.M{
			"_id" : u.Player.Inventory.ID,
		}
		err := collectionInventory.Update(query, u.Player.Inventory)
		if err != nil{
			return err
		}
	}

	// Сохранение юнитов
	{
		collectionUnits := database.Database.C(config.CollectionUnits)
		query := bson.M{
			"_id" : u.Player.Units.ID,
		}
		err := collectionUnits.Update(query, u.Player.Units)
		if err != nil{
			return err
		}
	}

	// Сохранение инвентаря аукциона
	{
		collectionAuctionInventory := database.Database.C(config.CollectionAuctionInventory)
		query := bson.M{
			"_id" : u.Player.AuctionInventory.ID,
		}
		err := collectionAuctionInventory.Update(query, u.Player.AuctionInventory)
		if err != nil{
			return err
		}
	}

	// Сохранение пользователя
	{
		resetToZeroUserMap(&u)

		collectionUsers := database.Database.C(config.CollectionUsers)
		query := bson.M{
			"_id" : u.ID,
		}
		err := collectionUsers.Update(query, u)
		if err != nil{
			return err
		}
	}

	return nil
}

// resetToZeroUserMap - делает мапы в структуре User пустыми.
func resetToZeroUserMap(u *User) {
	u.Player.Inventory.Weapons = make(map[string]*equipments.Weapon)
	u.Player.Inventory.Tools = make(map[string]*equipments.Tool)
	u.Player.Inventory.Chests = make(map[string]*chests.Chest)
	u.Player.Inventory.SummoningScrolls = make(map[string]*summoning_scrolls.SummoningScrolls)
	u.Player.Inventory.Armors = make(map[string]*equipments.Armor)
	u.Player.Inventory.Materials = make(map[string]*materials.Material)

	u.Player.AuctionInventory.Weapons = make(map[string]*equipments.Weapon)
	u.Player.AuctionInventory.Tools = make(map[string]*equipments.Tool)
	u.Player.AuctionInventory.Chests = make(map[string]*chests.Chest)
	u.Player.AuctionInventory.SummoningScrolls = make(map[string]*summoning_scrolls.SummoningScrolls)
	u.Player.AuctionInventory.Armors = make(map[string]*equipments.Armor)
	u.Player.AuctionInventory.Materials = make(map[string]*materials.Material)

	u.Player.Units.Units = make(map[string]*units.Unit)
}