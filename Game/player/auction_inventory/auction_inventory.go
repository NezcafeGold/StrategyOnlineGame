package auction_inventory

import (
	"../../objects/chests"
	"../../objects/equipments"
	"../../objects/materials"
	"../../objects/resources"
	"../../objects/summoning_scrolls"
	"sync"
)

// AuctionInventory - инвентарь аукциона.
type AuctionInventory struct {
	ID               string                                         `json:"id"                bson:"_id"`               // ID.
	Weapons          map[string]*equipments.Weapon                  `json:"weapons"           bson:"weapons"`           // Список оружия.
	Tools            map[string]*equipments.Tool                    `json:"tools"             bson:"tools"`             // Список инструментов.
	Chests           map[string]*chests.Chest                       `json:"chests"            bson:"chests"`            // Список сундуков.
	SummoningScrolls map[string]*summoning_scrolls.SummoningScrolls `json:"summoning_scrolls" bson:"summoning_scrolls"` // Список свитков призыва.
	Armors           map[string]*equipments.Armor                   `json:"armors"            bson:"armors"`            // Список брони.
	Materials        map[string]*materials.Material                 `json:"materials"         bson:"materials"`         // Список материалов
	Resources        *resources.Resources                           `json:"resources"         bson:"resources"`         // Ресурсы.
	RwMutex          *sync.RWMutex                                  // sync.RWMutex
}

// New - создать инвентарь аукциона.
func New(id string) *AuctionInventory {
	return &AuctionInventory{
		ID:               id,
		Weapons:          make(map[string]*equipments.Weapon),
		Tools:            make(map[string]*equipments.Tool),
		Chests:           make(map[string]*chests.Chest),
		Armors:           make(map[string]*equipments.Armor),
		SummoningScrolls: make(map[string]*summoning_scrolls.SummoningScrolls),
		Materials:        make(map[string]*materials.Material),
		Resources:        resources.New(),
		RwMutex:          new(sync.RWMutex),
	}
}
