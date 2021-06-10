package game

import (
	"./player/auction_inventory"
	"./user"
)

// userPage - страница с методами пользователя.
type userPage struct {
	ListOnlineUsers *user.ListOnlineUsers // Онлайн пользователи
	database        *user.Database        // Бд пользователей
}

// newUserPage - создать страницу с методами пользователя.
func newUserPage() *userPage {
	return &userPage{
		ListOnlineUsers: user.NewListOnlineUsers(),
		database:        user.NewDatabase(),
	}
}

// New - создание нового аккаунта пользователя.
func (page *userPage) New(email, nickname string) *user.User {
	return user.New(email, nickname)
}

// Save - сохроняет аккаунт пользователя в базу данных.
func (page *userPage) Save(u user.User) error {
	return page.database.AddUser(u)
}

// Get - получить аккаунт пользователя из базы данных.
func (page *userPage) Get(email string) (*user.User, error) {
	u, err := page.database.GetUser(email)
	if err == nil && u != nil {
		err = page.ListOnlineUsers.Add(u)
		if err != nil {
			return nil, err
		}
	}

	return u, err
}

// GetByID - получить аккаунт пользователя из базы данных по ID.
func (page *userPage) GetByID(id string) (*user.User, error) {
	return page.database.GetUserByID(id)
}

// GetAuctionInventoryByID - получить инвентарь аукциона пользователя из базы данных по ID.
func (page *userPage) GetAuctionInventoryByID(id string) (*auction_inventory.AuctionInventory, error) {
	return page.database.GetUserAuctionInventoryByID(id)
}

// UpdateAuctionInventoryByID - обновить инвентарь аукциона пользователя.
func (page *userPage) UpdateAuctionInventoryByID(auctionInventory *auction_inventory.AuctionInventory) error {
	return page.database.UpdateUserAuctionInventoryByID(auctionInventory)
}

// Syncing - синхронизировать данные пользователя.
func (page *userPage) Syncing(u user.User) error {
	return page.database.SyncingUser(u)
}

// Disconnect - отключить пользователя.
func (page *userPage) Disconnect(u *user.User) error {
	err := page.ListOnlineUsers.Delete(u.ID)
	if err != nil {
		return err
	}

	return page.Syncing(*u)
}