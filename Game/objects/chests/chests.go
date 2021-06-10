package chests

import (
	"../../utils/generate_items"
	"../equipments"
)

type TypeChest int

// Chest - сундук с предметами.
type Chest struct {
	ID        string    `json:"id"          bson:"_id"`         // ID.
	Title     string    `json:"title"       bson:"title"`       // Название.
	TypeChest TypeChest `json:"type_chest"  bson:"type_chest"`  // Тип сундука.
	IsOpen    bool      `json:"is_open"     bson:"is_open"`     // Статус.
	CanBeSold bool      `json:"can_be_sold" bson:"can_be_sold"` // Можно ли продать.
}

// Качество
var (
	generateQuality = []generate_items.GenerateObjInt {
		{int(equipments.Legendary), 1},
		{int(equipments.Epic), 5},
		{int(equipments.Rare), 15},
		{int(equipments.Unusual), 50},
		{int(equipments.Usual), 70},
	}
	generateQualityAllWeight = generate_items.WeightSumObjInt(generateQuality)
)

const (
	Beginner TypeChest = iota          // Сундук новичка.
	BeginnerArmors TypeChest = iota    // Сундук брони новичка.
	BeginnerTools TypeChest = iota     // Набор инструментов новичка.
	BeginnerJewelries TypeChest = iota // Набор бижутерии новичка.
	BeginnerWeapons TypeChest = iota   // Набор оружия новичка.
)

// Open - открыть сундук.
func (chest *Chest) Open() (interface{}, bool) {
	if !chest.IsOpen {

		chest.IsOpen = true

		switch chest.TypeChest {
		case Beginner:
			return OpenChestSetBeginner()
		case BeginnerWeapons:
			return OpenChestSetBeginnerWeapons()
		case BeginnerArmors:
			return OpenChestSetBeginnerArmors()
		case BeginnerTools:
			return OpenChestSetBeginnerTools()
		}
	}

	return nil, false
}