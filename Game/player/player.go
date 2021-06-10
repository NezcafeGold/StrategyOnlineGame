package player

import (
	"../objects/resources"
	"./auction_inventory"
	"./inventory"
	"./units"
	"errors"
)

// Player - игрок.
type Player struct {
	Level      uint8     `json:"level"      bson:"level"`      // Уровень.
	Experience uint64    `json:"experience" bson:"experience"` // Опыт.

	AuctionInventory *auction_inventory.AuctionInventory `json:"auction_inventory" bson:"auction_inventory"` // Инвентарь аукциона.
	Resources        *resources.Resources                `json:"resources"         bson:"resources"`         // Ресурсы.
	Inventory        *inventory.Inventory                `json:"inventory"         bson:"inventory"`         // Инвентарь.
	Units            *units.Storage                      `json:"units"             bson:"units"`             // Юниты.
}

// New - создание нового игрока.
func New(id string) *Player{
	player :=  &Player{
		Level:      0,
		Experience: 0,

		Resources:        resources.New(),
		Inventory:        inventory.New(id),
		Units:            units.New(id),
	}

	player.AuctionInventory = auction_inventory.New(id)

	return player
}

// PickUpWoodFromAuction - забрать дерево с инвентаря аукциона.
func (player *Player) PickUpWoodFromAuction() error {
	if player.AuctionInventory.Resources.Wood== 0 {
		return errors.New("Нет дерева которое можно забрать. ")
	}

	player.Resources.Wood = player.AuctionInventory.Resources.Wood
	player.AuctionInventory.Resources.Wood = 0

	return nil
}

// PickUpStoneFromAuction - забрать камня с инвентаря аукциона.
func (player *Player) PickUpStoneFromAuction() error {
	if player.AuctionInventory.Resources.Stone == 0 {
		return errors.New("Нет камня который можно забрать. ")
	}

	player.Resources.Stone += player.AuctionInventory.Resources.Stone
	player.AuctionInventory.Resources.Stone = 0

	return nil
}

// PickUpFoodFromAuction - забрать еду с инвентаря аукциона.
func (player *Player) PickUpFoodFromAuction() error {
	if player.AuctionInventory.Resources.Food == 0 {
		return errors.New("Нет еды которую можно забрать. ")
	}

	player.Resources.Food += player.AuctionInventory.Resources.Food
	player.AuctionInventory.Resources.Food = 0

	return nil
}

// PickUpIronFromAuction - забрать железо с инвентаря аукциона.
func (player *Player) PickUpIronFromAuction() error {
	if player.AuctionInventory.Resources.Iron == 0 {
		return errors.New("Нет железа которое можно забрать. ")
	}

	player.Resources.Iron += player.AuctionInventory.Resources.Iron
	player.AuctionInventory.Resources.Iron = 0

	return nil
}

// PickUpGoldFromAuction - забрать золото с инвентаря аукциона.
func (player *Player) PickUpGoldFromAuction() error {
	if player.AuctionInventory.Resources.Gold == 0 {
		return errors.New("Нет золота которое можно забрать. ")
	}

	player.Resources.Gold += player.AuctionInventory.Resources.Gold
	player.AuctionInventory.Resources.Gold = 0

	return nil
}

// PickUpAllResourcesFromAuction - забрать все ресурсы с инвентаря аукциона.
func (player *Player) PickUpAllResourcesFromAuction() error {
	err := player.PickUpWoodFromAuction()
	if err != nil {
		return err
	}

	err = player.PickUpFoodFromAuction()
	if err != nil {
		return err
	}

	err = player.PickUpStoneFromAuction()
	if err != nil {
		return err
	}

	err = player.PickUpIronFromAuction()
	if err != nil {
		return err
	}

	err = player.PickUpGoldFromAuction()
	if err != nil {
		return err
	}

	return nil
}