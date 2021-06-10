package armors

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualLeggings = []generate_items.GenerateObjFloat32 {
		{6.0, 50},
		{7.0, 30},
		{8.0, 10},
		{9.0, 5},
		{10.0, 1},
	}
	generateUsualLeggingsAllWeight = generate_items.WeightSumObjFloat32(generateUsualLeggings)

	generateUnusualLeggings = []generate_items.GenerateObjFloat32 {
		{11.0, 50},
		{12.0, 30},
		{13.0, 10},
		{14.0, 5},
		{15.0, 1},
	}
	generateUnusualLeggingsAllWeight = generate_items.WeightSumObjFloat32(generateUnusualLeggings)

	generateRareLeggings = []generate_items.GenerateObjFloat32 {
		{16.0, 50},
		{17.0, 30},
		{18.0, 10},
		{19.0, 5},
		{20.0, 1},
	}
	generateRareLeggingsAllWeight = generate_items.WeightSumObjFloat32(generateRareLeggings)

	generateEpicLeggings = []generate_items.GenerateObjFloat32 {
		{21.0, 50},
		{22.0, 30},
		{23.0, 10},
		{24.0, 5},
		{25.0, 1},
	}
	generateEpicLeggingsAllWeight = generate_items.WeightSumObjFloat32(generateEpicLeggings)

	generateLegendaryLeggings = []generate_items.GenerateObjFloat32 {
		{26.0, 50},
		{27.0, 30},
		{28.0, 10},
		{29.0, 5},
		{30.0, 1},
	}
	generateLegendaryLeggingsAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryLeggings)
)

// GenerationUsualBeginnerLeggings - генерация обычных штанов.
func GenerationUsualBeginnerLeggings() *equipments.Armor {
	leggings := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Штаны новичка",
		TypeArmor:     equipments.Leggings,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Defense:       0,
		Weight:        0.5,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	leggings.Defense = generate_items.GenerateValueObjFloat32(generateUsualLeggingsAllWeight, generateUsualLeggings)

	return &leggings
}

// GenerationUnusualBeginnerLeggings - генерация необычных штанов.
func GenerationUnusualBeginnerLeggings() *equipments.Armor {
	leggings := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Штаны новичка",
		TypeArmor:     equipments.Leggings,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Defense:       0,
		Weight:        0.5,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	leggings.Defense = generate_items.GenerateValueObjFloat32(generateUnusualLeggingsAllWeight, generateUnusualLeggings)

	return &leggings
}

// GenerationRareBeginnerLeggings - генерация редких штанов.
func GenerationRareBeginnerLeggings() *equipments.Armor {
	leggings := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Штаны новичка",
		TypeArmor:     equipments.Leggings,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Defense:       0,
		Weight:        0.5,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	leggings.Defense = generate_items.GenerateValueObjFloat32(generateRareLeggingsAllWeight, generateRareLeggings)

	return &leggings
}

// GenerationEpicBeginnerLeggings - генерация эпических штанов.
func GenerationEpicBeginnerLeggings() *equipments.Armor {
	leggings := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Штаны новичка",
		TypeArmor:     equipments.Leggings,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Defense:       0,
		Weight:        0.5,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	leggings.Defense = generate_items.GenerateValueObjFloat32(generateEpicLeggingsAllWeight, generateEpicLeggings)

	return &leggings
}

// GenerationLegendaryBeginnerLeggings - генерация легендарных штанов.
func GenerationLegendaryBeginnerLeggings() *equipments.Armor {
	leggings := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Штаны новичка",
		TypeArmor:     equipments.Leggings,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Defense:       0,
		Weight:        0.5,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	leggings.Defense = generate_items.GenerateValueObjFloat32(generateLegendaryLeggingsAllWeight, generateLegendaryLeggings)

	return &leggings
}