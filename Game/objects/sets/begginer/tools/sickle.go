package tools

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualSickle = []generate_items.GenerateObjFloat32 {
		{110.0, 50},
		{110.0, 30},
		{120.0, 10},
		{130.0, 5},
		{140.0, 1},
	}
	generateUsualSickleAllWeight = generate_items.WeightSumObjFloat32(generateUsualSickle)

	generateUnusualSickle = []generate_items.GenerateObjFloat32 {
		{150.0, 50},
		{160.0, 30},
		{170.0, 10},
		{180.0, 5},
		{190.0, 1},
	}
	generateUnusualSickleAllWeight = generate_items.WeightSumObjFloat32(generateUnusualSickle)

	generateRareSickle = []generate_items.GenerateObjFloat32 {
		{200.0, 50},
		{210.0, 30},
		{220.0, 10},
		{230.0, 5},
		{240.0, 1},
	}
	generateRareSickleAllWeight = generate_items.WeightSumObjFloat32(generateRareSickle)

	generateEpicSickle = []generate_items.GenerateObjFloat32 {
		{250.0, 50},
		{260.0, 30},
		{270.0, 10},
		{280.0, 5},
		{290.0, 1},
	}
	generateEpicSickleAllWeight = generate_items.WeightSumObjFloat32(generateEpicSickle)

	generateLegendarySickle = []generate_items.GenerateObjFloat32 {
		{300.0, 50},
		{310.0, 30},
		{320.0, 10},
		{330.0, 5},
		{340.0, 1},
	}
	generateLegendarySickleAllWeight = generate_items.WeightSumObjFloat32(generateLegendarySickle)
)

// GenerationUsualBeginnerSickle - генерация обычного серпа.
func GenerationUsualBeginnerSickle() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Серп новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Usual,
		SpeedWork:           0,
		ChanceDroppingItems: 5,
		TypeTool:            equipments.Sickle,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUsualSickleAllWeight, generateUsualSickle)

	return &tool
}

// GenerationUnusualBeginnerSickle - генерация необычного серпа.
func GenerationUnusualBeginnerSickle() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Серп новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Unusual,
		SpeedWork:           0,
		ChanceDroppingItems: 10,
		TypeTool:            equipments.Sickle,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUnusualSickleAllWeight, generateUnusualSickle)

	return &tool
}

// GenerationRareBeginnerSickle - генерация редкого серпа.
func GenerationRareBeginnerSickle() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Серп новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Rare,
		SpeedWork:           0,
		ChanceDroppingItems: 15,
		TypeTool:            equipments.Sickle,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateRareSickleAllWeight, generateRareSickle)

	return &tool
}

// GenerationEpicBeginnerSickle - генерация эпического серпа.
func GenerationEpicBeginnerSickle() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Серп новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Epic,
		SpeedWork:           0,
		ChanceDroppingItems: 20,
		TypeTool:            equipments.Sickle,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateEpicSickleAllWeight, generateEpicSickle)

	return &tool
}

// GenerationLegendaryBeginnerSickle - генерация легендарного серпа.
func GenerationLegendaryBeginnerSickle() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Серп новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Legendary,
		SpeedWork:           0,
		ChanceDroppingItems: 25,
		TypeTool:            equipments.Sickle,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateLegendarySickleAllWeight, generateLegendarySickle)
	
	return &tool
}
