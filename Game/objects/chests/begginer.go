package chests

import (
	"../../utils/generate_items"
	"../equipments"
	"../sets/begginer/armors"
	"../sets/begginer/tools"
	"../sets/begginer/weapons"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateChestSetBeginner = []generate_items.GenerateObjFloat32 {
		{1.0, 1},	// Шлем
		{2.0, 1},	// Нагрудник
		{3.0, 1},	// Штаны
		{4.0, 1},	// Ботинки

		{5.0, 1},	// Лук
		{6.0, 1},	// Короткий лук

		{19.0, 1},	// Топор
		{20.0, 1},	// Кирка
		{21.0, 1},	// Серп
		{22.0, 1},	// Молоток

		{23.0, 1},	// Кольцо
		{24.0, 1},	// Серьга
		{25.0, 1},	// Пояс
		{26.0, 1},	// Ожерелье
	}
	generateChestSetBeginnerAllWeight = generate_items.WeightSumObjFloat32(generateChestSetBeginner)
)

// NewChestSetBeginner - создать сундук новичка.
func NewChestSetBeginner() *Chest {
	return &Chest{
		ID:        bson.NewObjectId().Hex(),
		Title:     "Сундук новичка",
		TypeChest: Beginner,
		IsOpen:    false,
		CanBeSold: true,
	}
}

// OpenChestSetBeginner - открыть сундук новичка.
func OpenChestSetBeginner() (interface{}, bool) {
	typeEquipment := generate_items.GenerateValueObjFloat32(generateChestSetBeginnerAllWeight, generateChestSetBeginner)

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

	case 5.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return weapons.GenerationUsualBeginnerBow(), true
		case equipments.Unusual:
			return weapons.GenerationUnusualBeginnerBow(), true
		case equipments.Rare:
			return weapons.GenerationRareBeginnerBow(), true
		case equipments.Epic:
			return weapons.GenerationEpicBeginnerBow(), true
		case equipments.Legendary:
			return weapons.GenerationLegendaryBeginnerBow(), true
		}
	}
	case 6.0: {
		quality := generate_items.GenerateValueObjInt(generateQualityAllWeight, generateQuality)

		switch equipments.Quality(quality) {
		case equipments.Usual:
			return weapons.GenerationUsualBeginnerShortBow(), true
		case equipments.Unusual:
			return weapons.GenerationUnusualBeginnerShortBow(), true
		case equipments.Rare:
			return weapons.GenerationRareBeginnerShortBow(), true
		case equipments.Epic:
			return weapons.GenerationEpicBeginnerShortBow(), true
		case equipments.Legendary:
			return weapons.GenerationLegendaryBeginnerShortBow(), true
		}
	}

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