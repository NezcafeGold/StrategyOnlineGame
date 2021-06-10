package chests

import (
	"../../utils/generate_items"
	"../equipments"
	"../sets/begginer/armors"
	"gopkg.in/mgo.v2/bson"
)

var (
	GenerateChestSetBeginnerArmors = []generate_items.GenerateObjFloat32 {
		{1.0, 1},
		{2.0, 1},
		{3.0, 1},
		{4.0, 1},
	}
	GenerateChestSetBeginnerArmorsAllWeight = generate_items.WeightSumObjFloat32(GenerateChestSetBeginnerArmors)
)

// NewChestSetBeginnerArmors - создать сундук с броней новичка.
func NewChestSetBeginnerArmors() *Chest {
	return &Chest{
		ID:        bson.NewObjectId().Hex(),
		Title:     "Сундук с броней новичка",
		TypeChest: BeginnerArmors,
		IsOpen:    false,
		CanBeSold: true,
	}
}

// OpenChestSetBeginnerArmors - открыть сундук с броней новичка.
func OpenChestSetBeginnerArmors() (interface{}, bool) {
	typeEquipment := generate_items.GenerateValueObjFloat32(GenerateChestSetBeginnerArmorsAllWeight, GenerateChestSetBeginnerArmors)

	switch typeEquipment {
	case 1.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return armors.GenerationUsualBeginnerHelmet(), true
		case equipments.Unusual:
			return armors.GenerationUnusualBeginnerHelmet(), true
		case equipments.Rare:
			return armors.GenerationRareBeginnerHelmet(), true
		case equipments.Epic:
			return armors.GenerationEpicBeginnerHelmet(), true
		case equipments.Legendary:
			return armors.GenerationLegendaryBeginnerHelmet(), true
		}
	}
	case 2.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return armors.GenerationUsualBeginnerChestplate(), true
		case equipments.Unusual:
			return armors.GenerationUnusualBeginnerChestplate(), true
		case equipments.Rare:
			return armors.GenerationRareBeginnerChestplate(), true
		case equipments.Epic:
			return armors.GenerationEpicBeginnerChestplate(), true
		case equipments.Legendary:
			return armors.GenerationLegendaryBeginnerChestplate(), true
		}
	}
	case 3.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return armors.GenerationUsualBeginnerLeggings(), true
		case equipments.Unusual:
			return armors.GenerationUnusualBeginnerLeggings(), true
		case equipments.Rare:
			return armors.GenerationRareBeginnerLeggings(), true
		case equipments.Epic:
			return armors.GenerationEpicBeginnerLeggings(), true
		case equipments.Legendary:
			return armors.GenerationLegendaryBeginnerLeggings(), true
		}
	}
	case 4.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return armors.GenerationUsualBeginnerBoots(), true
		case equipments.Unusual:
			return armors.GenerationUnusualBeginnerBoots(), true
		case equipments.Rare:
			return armors.GenerationRareBeginnerBoots(), true
		case equipments.Epic:
			return armors.GenerationEpicBeginnerBoots(), true
		case equipments.Legendary:
			return armors.GenerationLegendaryBeginnerBoots(), true
		}
	}
	}

	return nil, false
}

