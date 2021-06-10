package auction

import (
	"../user"
	"./resources"
)

// Auction - аукцион ресурсов.
type Auction struct {
	Resources  *resources.Resources     // Аукцион ресурсов
}

// New - создание аукциона.
func New(listOnlineUsers *user.ListOnlineUsers) *Auction {
	auction := new(Auction)

	auction.Resources = resources.New(listOnlineUsers)

	return auction
}

// Syncing - синхронизация аукционов.
func (auction *Auction) Syncing() {
	auction.Resources.Syncing()
}
