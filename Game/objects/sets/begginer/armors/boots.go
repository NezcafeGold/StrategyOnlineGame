package armors

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualBoots = []generate_items.GenerateObjFloat32 {
		{4.0, 50},
		{5.0, 30},
		{6.0, 10},
		{7.0, 5},
		{8.0, 1},
	}
	generateUsualBootsAllWeight = generate_items.WeightSumObjFloat32(generateUsualBoots)

	generateUnusualBoots = []generate_items.GenerateObjFloat32 {
		{9.0, 50},
		{10.0, 30},
		{11.0, 10},
		{12.0, 5},
		{13.0, 1},
	}
	generateUnusualBootsAllWeight = generate_items.WeightSumObjFloat32(generateUnusualBoots)

	generateRareBoots = []generate_items.GenerateObjFloat32 {
		{14.0, 50},
		{15.0, 30},
		{16.0, 10},
		{17.0, 5},
		{18.0, 1},
	}
	generateRareBootsAllWeight = generate_items.WeightSumObjFloat32(generateRareBoots)

	generateEpicBoots = []generate_items.GenerateObjFloat32 {
		{19.0, 50},
		{20.0, 30},
		{21.0, 10},
		{22.0, 5},
		{23.0, 1},
	}
	generateEpicBootsAllWeight = generate_items.WeightSumObjFloat32(generateEpicBoots)

	generateLegendaryBoots = []generate_items.GenerateObjFloat32 {
		{24.0, 50},
		{25.0, 30},
		{26.0, 10},
		{27.0, 5},
		{28.0, 1},
	}
	generateLegendaryBootsAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryBoots)
)

// GenerationUsualBeginnerBoots - генерация обычных ботинков.
func GenerationUsualBeginnerBoots() *equipments.Armor {
	boots := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Ботинки новичка",
		TypeArmor:     equipments.Boots,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Defense:       0,
		Weight:        0.4,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	boots.Defense = generate_items.GenerateValueObjFloat32(generateUsualBootsAllWeight, generateUsualBoots)

	return &boots
}

// GenerationUnusualBeginnerBoots - генерация необычных ботинков.
func GenerationUnusualBeginnerBoots() *equipments.Armor {
	boots := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Ботинки новичка",
		TypeArmor:     equipments.Boots,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Defense:       0,
		Weight:        0.4,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	boots.Defense = generate_items.GenerateValueObjFloat32(generateUnusualBootsAllWeight, generateUnusualBoots)

	return &boots
}

// GenerationRareBeginnerBoots - генерация редких ботинков.
func GenerationRareBeginnerBoots() *equipments.Armor {
	boots := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Ботинки новичка",
		TypeArmor:     equipments.Boots,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Defense:       0,
		Weight:        0.4,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	boots.Defense = generate_items.GenerateValueObjFloat32(generateRareBootsAllWeight, generateRareBoots)

	return &boots
}

// GenerationEpicBeginnerBoots - генерация эпических ботинков.
func GenerationEpicBeginnerBoots() *equipments.Armor {
	boots := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Ботинки новичка",
		TypeArmor:     equipments.Boots,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Defense:       0,
		Weight:        0.4,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	boots.Defense = generate_items.GenerateValueObjFloat32(generateEpicBootsAllWeight, generateEpicBoots)

	return &boots
}

// GenerationLegendaryBeginnerBoots - генерация легендарных ботинков.
func GenerationLegendaryBeginnerBoots() *equipments.Armor {
	boots := equipments.Armor{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Ботинки новичка",
		TypeArmor:     equipments.Boots,
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Defense:       0,
		Weight:        0.4,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	boots.Defense = generate_items.GenerateValueObjFloat32(generateLegendaryBootsAllWeight, generateLegendaryBoots)

	return &boots
}