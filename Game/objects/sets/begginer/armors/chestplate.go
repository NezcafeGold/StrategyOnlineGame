package armors

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualChestplate = []generate_items.GenerateObjFloat32 {
		{7.0, 50},
		{8.0, 30},
		{9.0, 10},
		{10.0, 5},
		{11.0, 1},
	}
	generateUsualChestplateAllWeight = generate_items.WeightSumObjFloat32(generateUsualChestplate)

	generateUnusualChestplate = []generate_items.GenerateObjFloat32 {
		{12.0, 50},
		{13.0, 30},
		{14.0, 10},
		{15.0, 5},
		{16.0, 1},
	}
	generateUnusualChestplateAllWeight = generate_items.WeightSumObjFloat32(generateUnusualChestplate)

	generateRareChestplate = []generate_items.GenerateObjFloat32 {
		{17.0, 50},
		{18.0, 30},
		{19.0, 10},
		{20.0, 5},
		{21.0, 1},
	}
	generateRareChestplateAllWeight = generate_items.WeightSumObjFloat32(generateRareChestplate)

	generateEpicChestplate = []generate_items.GenerateObjFloat32 {
		{22.0, 50},
		{23.0, 30},
		{24.0, 10},
		{25.0, 5},
		{26.0, 1},
	}
	generateEpicChestplateAllWeight = generate_items.WeightSumObjFloat32(generateEpicChestplate)

	generateLegendaryChestplate = []generate_items.GenerateObjFloat32 {
		{27.0, 50},
		{28.0, 30},
		{29.0, 10},
		{30.0, 5},
		{31.0, 1},
	}
	generateLegendaryChestplateAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryChestplate)
)

// GenerationUsualBeginnerChestplate - генерация обычного нагрудника.
func GenerationUsualBeginnerChestplate() *equipments.Armor {
	chestplate := equipments.Armor {
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Нагрудник новичка",
		TypeArmor:     equipments.Chestplate,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Defense:       0,
		Weight:        0.6,
		RequiredLevel: 1,
		CanBeSold:   true,
	}

	chestplate.Defense = generate_items.GenerateValueObjFloat32(generateUsualChestplateAllWeight, generateUsualChestplate)

	return &chestplate
}

// GenerationUnusualBeginnerChestplate - генерация необычного нагрудника.
func GenerationUnusualBeginnerChestplate() *equipments.Armor {
	chestplate := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Нагрудник новичка",
		TypeArmor:     equipments.Chestplate,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Defense:       0,
		Weight:        0.6,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	chestplate.Defense = generate_items.GenerateValueObjFloat32(generateUnusualChestplateAllWeight, generateUnusualChestplate)

	return &chestplate
}

// GenerationRareBeginnerChestplate - генерация редкого нагрудника.
func GenerationRareBeginnerChestplate() *equipments.Armor {
	chestplate := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Нагрудник новичка",
		TypeArmor:     equipments.Chestplate,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Defense:       0,
		Weight:        0.6,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	chestplate.Defense = generate_items.GenerateValueObjFloat32(generateRareChestplateAllWeight, generateRareChestplate)

	return &chestplate
}

// GenerationEpicBeginnerChestplate - генерация эпического нагрудника.
func GenerationEpicBeginnerChestplate() *equipments.Armor {
	chestplate := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Нагрудник новичка",
		TypeArmor:     equipments.Chestplate,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Defense:       0,
		Weight:        0.6,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	chestplate.Defense = generate_items.GenerateValueObjFloat32(generateEpicChestplateAllWeight, generateEpicChestplate)

	return &chestplate
}

// GenerationLegendaryBeginnerChestplate - генерация легендарного нагрудника.
func GenerationLegendaryBeginnerChestplate() *equipments.Armor {
	chestplate := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Нагрудник новичка",
		TypeArmor:     equipments.Chestplate,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Defense:       0,
		Weight:        0.6,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	chestplate.Defense = generate_items.GenerateValueObjFloat32(generateLegendaryChestplateAllWeight, generateLegendaryChestplate)

	return &chestplate
}
