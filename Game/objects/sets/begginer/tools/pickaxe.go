package tools

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualPickaxe = []generate_items.GenerateObjFloat32 {
		{60.0, 50},
		{70.0, 30},
		{80.0, 10},
		{90.0, 5},
		{100.0, 1},
	}
	generateUsualPickaxeAllWeight = generate_items.WeightSumObjFloat32(generateUsualPickaxe)

	generateUnusualPickaxe = []generate_items.GenerateObjFloat32 {
		{110.0, 50},
		{120.0, 30},
		{130.0, 10},
		{140.0, 5},
		{150.0, 1},
	}
	generateUnusualPickaxeAllWeight = generate_items.WeightSumObjFloat32(generateUsualPickaxe)

	generateRarePickaxe = []generate_items.GenerateObjFloat32 {
		{160.0, 50},
		{170.0, 30},
		{180.0, 10},
		{190.0, 5},
		{200.0, 1},
	}
	generateRarePickaxeAllWeight = generate_items.WeightSumObjFloat32(generateRarePickaxe)

	generateEpicPickaxe = []generate_items.GenerateObjFloat32 {
		{210.0, 50},
		{220.0, 30},
		{230.0, 10},
		{240.0, 5},
		{250.0, 1},
	}
	generateEpicPickaxeAllWeight = generate_items.WeightSumObjFloat32(generateEpicPickaxe)

	generateLegendaryPickaxe = []generate_items.GenerateObjFloat32 {
		{260.0, 50},
		{270.0, 30},
		{280.0, 10},
		{290.0, 5},
		{300.0, 1},
	}
	generateLegendaryPickaxeAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryPickaxe)
)

// GenerationUsualBeginnerPickaxe - генерация обычной кирки.
func GenerationUsualBeginnerPickaxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Кирка новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Usual,
		SpeedWork:           0,
		ChanceDroppingItems: 5,
		TypeTool:            equipments.Pickaxe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUsualPickaxeAllWeight, generateUsualPickaxe)

	return &tool
}

// GenerationUnusualBeginnerPickaxe - генерация необычной кирки.
func GenerationUnusualBeginnerPickaxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Кирка новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Unusual,
		SpeedWork:           0,
		ChanceDroppingItems: 10,
		TypeTool:            equipments.Pickaxe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUnusualPickaxeAllWeight, generateUnusualPickaxe)

	return &tool
}

// GenerationRareBeginnerPickaxe - генерация редкой кирки.
func GenerationRareBeginnerPickaxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Кирка новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Rare,
		SpeedWork:           0,
		ChanceDroppingItems: 15,
		TypeTool:            equipments.Pickaxe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateRarePickaxeAllWeight, generateRarePickaxe)

	return &tool
}

// GenerationEpicBeginnerPickaxe - генерация эпической кирки.
func GenerationEpicBeginnerPickaxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Кирка новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Epic,
		SpeedWork:           0,
		ChanceDroppingItems: 20,
		TypeTool:            equipments.Pickaxe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateEpicPickaxeAllWeight, generateEpicPickaxe)

	return &tool
}

// GenerationLegendaryBeginnerPickaxe - генерация легендарной кирки.
func GenerationLegendaryBeginnerPickaxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Кирка новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Legendary,
		SpeedWork:           0,
		ChanceDroppingItems: 25,
		TypeTool:            equipments.Pickaxe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateLegendaryPickaxeAllWeight, generateLegendaryPickaxe)

	return &tool
}

