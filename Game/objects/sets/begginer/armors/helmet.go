package armors

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualHelmet = []generate_items.GenerateObjFloat32 {
		{5.0, 50},
		{6.0, 30},
		{7.0, 10},
		{8.0, 5},
		{9.0, 1},
	}
	generateUsualHelmetAllWeight = generate_items.WeightSumObjFloat32(generateUsualHelmet)

	generateUnusualHelmet = []generate_items.GenerateObjFloat32 {
		{10.0, 50},
		{11.0, 30},
		{12.0, 10},
		{13.0, 5},
		{14.0, 1},
	}
	generateUnusualHelmetAllWeight = generate_items.WeightSumObjFloat32(generateUnusualHelmet)

	generateRareHelmet = []generate_items.GenerateObjFloat32 {
		{15.0, 50},
		{16.0, 30},
		{17.0, 10},
		{18.0, 5},
		{19.0, 1},
	}
	generateRareHelmetAllWeight = generate_items.WeightSumObjFloat32(generateRareHelmet)

	generateEpicHelmet = []generate_items.GenerateObjFloat32 {
		{20.0, 50},
		{21.0, 30},
		{22.0, 10},
		{23.0, 5},
		{24.0, 1},
	}
	generateEpicHelmetAllWeight = generate_items.WeightSumObjFloat32(generateEpicHelmet)

	generateLegendaryHelmet = []generate_items.GenerateObjFloat32 {
		{25.0, 50},
		{26.0, 30},
		{27.0, 10},
		{28.0, 5},
		{29.0, 1},
	}
	generateLegendaryHelmetAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryHelmet)
)

// GenerationUsualBeginnerHelmet - генерация обычного шлема.
func GenerationUsualBeginnerHelmet() *equipments.Armor {
	helmet := equipments.Armor {
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Шлем новичка",
		TypeArmor:     equipments.Helmet,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Defense:       0,
		Weight:        0.3,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	helmet.Defense = generate_items.GenerateValueObjFloat32(generateUsualHelmetAllWeight, generateUsualHelmet)

	return &helmet
}

// GenerationUnusualBeginnerHelmet - генерация необычного шлема.
func GenerationUnusualBeginnerHelmet() *equipments.Armor {
	helmet := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Шлем новичка",
		TypeArmor:     equipments.Helmet,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Defense:       0,
		Weight:        0.3,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	helmet.Defense = generate_items.GenerateValueObjFloat32(generateUnusualHelmetAllWeight, generateUnusualHelmet)

	return &helmet
}

// GenerationRareBeginnerHelmet - генерация редкого шлема.
func GenerationRareBeginnerHelmet() *equipments.Armor {
	helmet := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Шлем новичка",
		TypeArmor:     equipments.Helmet,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Defense:       0,
		Weight:        0.3,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	helmet.Defense = generate_items.GenerateValueObjFloat32(generateRareHelmetAllWeight, generateRareHelmet)

	return &helmet
}

// GenerationEpicBeginnerHelmet - генерация эпического шлема.
func GenerationEpicBeginnerHelmet() *equipments.Armor {
	helmet := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Шлем новичка",
		TypeArmor:     equipments.Helmet,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Defense:       0,
		Weight:        0.3,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	helmet.Defense = generate_items.GenerateValueObjFloat32(generateEpicHelmetAllWeight, generateEpicHelmet)

	return &helmet
}

// GenerationLegendaryBeginnerHelmet - генерация легендарного шлема.
func GenerationLegendaryBeginnerHelmet() *equipments.Armor {
	helmet := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Шлем новичка",
		TypeArmor:     equipments.Helmet,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Defense:       0,
		Weight:        0.3,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	helmet.Defense = generate_items.GenerateValueObjFloat32(generateLegendaryHelmetAllWeight, generateLegendaryHelmet)

	return &helmet
}