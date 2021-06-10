package tools

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualAxe = []generate_items.GenerateObjFloat32 {
		{60.0, 50},
		{70.0, 30},
		{80.0, 10},
		{90.0, 5},
		{100.0, 1},
	}
	generateUsualAxeAllWeight = generate_items.WeightSumObjFloat32(generateUsualAxe)

	generateUnusualAxe = []generate_items.GenerateObjFloat32 {
		{110.0, 50},
		{120.0, 30},
		{130.0, 10},
		{140.0, 5},
		{150.0, 1},
	}
	generateUnusualAxeAllWeight = generate_items.WeightSumObjFloat32(generateUnusualAxe)

	generateRareAxe = []generate_items.GenerateObjFloat32 {
		{160.0, 50},
		{170.0, 30},
		{180.0, 10},
		{190.0, 5},
		{200.0, 1},
	}
	generateRareAxeAllWeight = generate_items.WeightSumObjFloat32(generateRareAxe)

	generateEpicAxe = []generate_items.GenerateObjFloat32 {
		{210.0, 50},
		{220.0, 30},
		{230.0, 10},
		{240.0, 5},
		{250.0, 1},
	}
	generateEpicAxeAllWeight = generate_items.WeightSumObjFloat32(generateEpicAxe)

	generateLegendaryAxe = []generate_items.GenerateObjFloat32 {
		{260.0, 50},
		{270.0, 30},
		{280.0, 10},
		{290.0, 5},
		{300.0, 1},
	}
	generateLegendaryAxeAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryAxe)
)

// GenerationUsualBeginnerAxe - генерация обычного топора.
func GenerationUsualBeginnerAxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Топор новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Usual,
		SpeedWork:           0,
		ChanceDroppingItems: 5,
		TypeTool:            equipments.Axe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUsualAxeAllWeight, generateUsualAxe)

	return &tool
}

// GenerationUnusualBeginnerAxe - генерация необычного топора.
func GenerationUnusualBeginnerAxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Топор новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Unusual,
		SpeedWork:           0,
		ChanceDroppingItems: 10,
		TypeTool:            equipments.Axe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUnusualAxeAllWeight, generateUnusualAxe)

	return &tool
}

// GenerationRareBeginnerAxe - генерация редкого топора.
func GenerationRareBeginnerAxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Топор новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Rare,
		SpeedWork:           0,
		ChanceDroppingItems: 15,
		TypeTool:            equipments.Axe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateRareAxeAllWeight, generateRareAxe)

	return &tool
}

// GenerationEpicBeginnerAxe - генерация эпического топора.
func GenerationEpicBeginnerAxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Топор новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Epic,
		SpeedWork:           0,
		ChanceDroppingItems: 20,
		TypeTool:            equipments.Axe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateEpicAxeAllWeight, generateEpicAxe)

	return &tool
}

// GenerationLegendaryBeginnerAxe - генерация легендарного топора.
func GenerationLegendaryBeginnerAxe() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Топор новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Legendary,
		SpeedWork:           0,
		ChanceDroppingItems: 25,
		TypeTool:            equipments.Axe,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateLegendaryAxeAllWeight, generateLegendaryAxe)

	return &tool
}