package inventory

import (
	"../../objects/chests"
	"../../objects/equipments"
	"../../objects/materials"
	"../../objects/summoning_scrolls"
	"../../utils/config"
	"sync"
)

// Inventory - инвентарь.
type Inventory struct {
	ID               string                                         `json:"id"                bson:"_id"`               // ID.
	MaxSize          uint16                                         `json:"max_size"          bson:"max_size"`          // Максимальный размер.
	Size             uint16                                         `json:"size"              bson:"size"`              // Текущий размер.
	Weapons          map[string]*equipments.Weapon                  `json:"weapons"           bson:"weapons"`           // Список оружия.
	Tools            map[string]*equipments.Tool                    `json:"tools"             bson:"tools"`             // Список инструментов.
	Chests           map[string]*chests.Chest                       `json:"chests"            bson:"chests"`            // Список сундуков.
	SummoningScrolls map[string]*summoning_scrolls.SummoningScrolls `json:"summoning_scrolls" bson:"summoning_scrolls"` // Список свитков призыва.
	Armors           map[string]*equipments.Armor                   `json:"armors"            bson:"armors"`            // Список брони.
	Materials        map[string]*materials.Material                 `json:"materials"         bson:"materials"`         // Список материалов
	RwMutex          *sync.RWMutex                                                                                      // sync.RWMutex
}

// Item - предмет.
type Item interface {}

// New - создание нового инвентаря.
func New(id string) *Inventory {
	return &Inventory{
		ID:               id,
		MaxSize:          config.MaxSizeInventory,
		Size:             0,
		Weapons:          make(map[string]*equipments.Weapon),
		Tools:            make(map[string]*equipments.Tool),
		Chests:           make(map[string]*chests.Chest),
		Armors:           make(map[string]*equipments.Armor),
		SummoningScrolls: make(map[string]*summoning_scrolls.SummoningScrolls),
		Materials:        make(map[string]*materials.Material),
		RwMutex:          new(sync.RWMutex),
	}
}

// Add - добавить предмет в инвентарь.
func (inventory *Inventory) Add(item Item) {
	if inventory.Size < inventory.MaxSize {
		switch object := item.(type) {
		case *equipments.Weapon: {
			inventory.RwMutex.Lock()
			inventory.Weapons[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case equipments.Weapon: {
			inventory.RwMutex.Lock()
			inventory.Weapons[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		case *equipments.Tool: {
			inventory.RwMutex.Lock()
			inventory.Tools[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case equipments.Tool: {
			inventory.RwMutex.Lock()
			inventory.Tools[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		case *chests.Chest: {
			inventory.RwMutex.Lock()
			inventory.Chests[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case chests.Chest: {
			inventory.RwMutex.Lock()
			inventory.Chests[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		case *summoning_scrolls.SummoningScrolls: {
			inventory.RwMutex.Lock()
			inventory.SummoningScrolls[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case summoning_scrolls.SummoningScrolls: {
			inventory.RwMutex.Lock()
			inventory.SummoningScrolls[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		case *materials.Material: {
			inventory.RwMutex.Lock()
			inventory.Materials[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case materials.Material: {
			inventory.RwMutex.Lock()
			inventory.Materials[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		case *equipments.Armor: {
			inventory.RwMutex.Lock()
			inventory.Armors[object.ID] = object
			inventory.RwMutex.Unlock()
		}
		case equipments.Armor: {
			inventory.RwMutex.Lock()
			inventory.Armors[object.ID] = &object
			inventory.RwMutex.Unlock()
		}
		}
	}

	inventory.UpdateSize()
}

// UpdateSize - обновить размер инвентаря.
func (inventory *Inventory) UpdateSize() {
	inventory.RwMutex.Lock()
	inventory.Size = uint16(len(inventory.Weapons) + len(inventory.Tools) + len(inventory.Chests) + len(inventory.Armors) + len(inventory.SummoningScrolls) + len(inventory.Materials))
	inventory.RwMutex.Unlock()
}