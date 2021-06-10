package tools

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualHammer = []generate_items.GenerateObjFloat32 {
		{1.0, 30},
		{2.0, 10},
		{3.0, 1},
	}
	generateUsualHammerAllWeight = generate_items.WeightSumObjFloat32(generateUsualHammer)

	generateUnusualHammer = []generate_items.GenerateObjFloat32 {
		{4.0, 30},
		{5.0, 10},
		{6.0, 1},
	}
	generateUnusualHammerAllWeight = generate_items.WeightSumObjFloat32(generateUnusualHammer)

	generateRareHammer = []generate_items.GenerateObjFloat32 {
		{7.0, 30},
		{8.0, 10},
		{9.0, 1},
	}
	generateRareHammerAllWeight = generate_items.WeightSumObjFloat32(generateRareHammer)

	generateEpicHammer = []generate_items.GenerateObjFloat32 {
		{10.0, 30},
		{11.0, 10},
		{12.0, 1},
	}
	generateEpicHammerAllWeight = generate_items.WeightSumObjFloat32(generateEpicHammer)

	generateLegendaryHammer = []generate_items.GenerateObjFloat32 {
		{13.0, 30},
		{14.0, 10},
		{15.0, 1},
	}
	generateLegendaryHammerAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryHammer)
)

// GenerationUsualBeginnerHammer - генерация обычного молотка.
func GenerationUsualBeginnerHammer() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Молоток новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Usual,
		SpeedWork:           0,
		ChanceDroppingItems: 0,
		TypeTool:            equipments.Hammer,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUsualHammerAllWeight, generateUsualHammer)

	return &tool
}

// GenerationUnusualBeginnerHammer - генерация необычного молотка.
func GenerationUnusualBeginnerHammer() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Молоток новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Unusual,
		SpeedWork:           0,
		ChanceDroppingItems: 0,
		TypeTool:            equipments.Hammer,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateUnusualHammerAllWeight, generateUnusualHammer)


	return &tool
}

// GenerationRareBeginnerHammer - генерация редкого молотка.
func GenerationRareBeginnerHammer() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Молоток новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Rare,
		SpeedWork:           0,
		ChanceDroppingItems: 0,
		TypeTool:            equipments.Hammer,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateRareHammerAllWeight, generateRareHammer)

	return &tool
}

// GenerationEpicBeginnerHammer - генерация эпического молотка.
func GenerationEpicBeginnerHammer() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Молоток новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Epic,
		SpeedWork:           0,
		ChanceDroppingItems: 0,
		TypeTool:            equipments.Hammer,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateEpicHammerAllWeight, generateEpicHammer)

	return &tool
}

// GenerationLegendaryBeginnerHammer - генерация легендарного молотка.
func GenerationLegendaryBeginnerHammer() *equipments.Tool {
	tool := equipments.Tool{
		ID:                  bson.NewObjectId().Hex(),
		IDSet:               sets.Begginer,
		Title:               "Молоток новичка",
		Durability:          100,
		MaxDurability:       100,
		Quality:             equipments.Legendary,
		SpeedWork:           0,
		ChanceDroppingItems: 0,
		TypeTool:            equipments.Hammer,
		AttackRadius:        1,
		Weight:              1,
		RequiredLevel:       1,
		CanBeSold:           true,
	}

	tool.SpeedWork = generate_items.GenerateValueObjFloat32(generateLegendaryHammerAllWeight, generateLegendaryHammer)

	return &tool
}