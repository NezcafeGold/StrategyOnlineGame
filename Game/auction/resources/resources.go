package resources

import (
	"../../user"
	"sync"
)

// Resources - аукцион ресурсов.
type Resources struct {
	Asks *Asks  // Заявки продовцов.
	Bids *Bids  // Заявки покупателей.
}

// New - создание аукциона ресурсов
func New(listOnlineUsers *user.ListOnlineUsers) *Resources {
	return &Resources{
		Asks: newAsks(listOnlineUsers),
		Bids: newBids(listOnlineUsers),
	}
}

// Syncing - синхронизация аукциона ресурсов.
func (auction *Resources) Syncing() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		defer wg.Done()

		auction.Asks.Syncing()
	}()

	go func() {
		defer wg.Done()

		auction.Bids.Syncing()
	}()

	wg.Wait()
}
