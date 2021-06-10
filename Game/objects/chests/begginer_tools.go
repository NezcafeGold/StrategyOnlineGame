package chests

import (
	"../../utils/generate_items"
	"../equipments"
	"../sets/begginer/tools"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateChestSetBeginnerTools = []generate_items.GenerateObjFloat32 {
		{19.0, 1},
		{20.0, 1},
		{21.0, 1},
		{22.0, 1},
	}
	generateChestSetBeginnerToolsAllWeight = generate_items.WeightSumObjFloat32(generateChestSetBeginnerTools)
)

// NewChestSetBeginnerTools - создать сундук с инструментами новичка.
func NewChestSetBeginnerTools() *Chest {
	return &Chest{
		ID:        bson.NewObjectId().Hex(),
		Title:     "Сундук c инструментами новичка",
		TypeChest: BeginnerTools,
		IsOpen:    false,
		CanBeSold: true,
	}
}

// OpenChestSetBeginnerTools - открыть сундук с инструментами новичка.
func OpenChestSetBeginnerTools() (interface{}, bool) {
	typeEquipment := generate_items.GenerateValueObjFloat32(generateChestSetBeginnerToolsAllWeight, generateChestSetBeginnerTools)

	switch typeEquipment {
	case 19.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return tools.GenerationUsualBeginnerAxe(), true
		case equipments.Unusual:
			return tools.GenerationUnusualBeginnerAxe(), true
		case equipments.Rare:
			return tools.GenerationRareBeginnerAxe(), true
		case equipments.Epic:
			return tools.GenerationEpicBeginnerAxe(), true
		case equipments.Legendary:
			return tools.GenerationLegendaryBeginnerAxe(), true
		}
	}
	case 20.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return tools.GenerationUsualBeginnerPickaxe(), true
		case equipments.Unusual:
			return tools.GenerationUnusualBeginnerPickaxe(), true
		case equipments.Rare:
			return tools.GenerationRareBeginnerPickaxe(), true
		case equipments.Epic:
			return tools.GenerationEpicBeginnerPickaxe(), true
		case equipments.Legendary:
			return tools.GenerationLegendaryBeginnerPickaxe(), true
		}
	}
	case 21.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return tools.GenerationUsualBeginnerSickle(), true
		case equipments.Unusual:
			return tools.GenerationUnusualBeginnerSickle(), true
		case equipments.Rare:
			return tools.GenerationRareBeginnerSickle(), true
		case equipments.Epic:
			return tools.GenerationEpicBeginnerSickle(), true
		case equipments.Legendary:
			return tools.GenerationLegendaryBeginnerSickle(), true
		}
	}
	case 22.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return tools.GenerationUsualBeginnerHammer(), true
		case equipments.Unusual:
			return tools.GenerationUnusualBeginnerHammer(), true
		case equipments.Rare:
			return tools.GenerationRareBeginnerHammer(), true
		case equipments.Epic:
			return tools.GenerationEpicBeginnerHammer(), true
		case equipments.Legendary:
			return tools.GenerationLegendaryBeginnerHammer(), true
		}
	}
	}

	return nil, false
}
