package weapons

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualBow = []generate_items.GenerateObjFloat32 {
		{10.0, 50},
		{11.0, 30},
		{12.0, 10},
		{13.0, 5},
		{14.0, 1},
	}
	generateUsualBowAllWeight = generate_items.WeightSumObjFloat32(generateUsualBow)

	generateUnusualBow = []generate_items.GenerateObjFloat32 {
		{15.0, 50},
		{16.0, 30},
		{17.0, 10},
		{18.0, 5},
		{19.0, 1},
	}
	generateUnusualBowAllWeight = generate_items.WeightSumObjFloat32(generateUnusualBow)

	generateRareBow = []generate_items.GenerateObjFloat32 {
		{20.0, 50},
		{21.0, 30},
		{22.0, 10},
		{23.0, 5},
		{24.0, 1},
	}
	generateRareBowAllWeight = generate_items.WeightSumObjFloat32(generateRareBow)

	generateEpicBow = []generate_items.GenerateObjFloat32 {
		{25.0, 50},
		{26.0, 30},
		{27.0, 10},
		{28.0, 5},
		{29.0, 1},
	}
	generateEpicBowAllWeight = generate_items.WeightSumObjFloat32(generateEpicBow)

	generateLegendaryBow = []generate_items.GenerateObjFloat32 {
		{30.0, 50},
		{31.0, 30},
		{32.0, 10},
		{33.0, 5},
		{34.0, 1},
	}
	generateLegendaryBowAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryBow)
)

// GenerationUsualBeginnerBow - генерация обычного лука
func GenerationUsualBeginnerBow() *equipments.Weapon {
	weapon := equipments.Weapon {
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Damage:        0,
		TypeWeapon:    equipments.Bow,
		SpeedAttack:   1,
		AttackRadius:  20,
		Weight:        1.1,
		RequiredLevel: 1,
		CanBeSold:   true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateUsualBowAllWeight, generateUsualBow)

	return &weapon
}

// GenerationUnusualBeginnerBow - генерация необычного лука.
func GenerationUnusualBeginnerBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Damage:        0,
		TypeWeapon:    equipments.Bow,
		SpeedAttack:   1,
		AttackRadius:  20,
		Weight:        1.1,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateUnusualBowAllWeight, generateUnusualBow)

	return &weapon
}

// GenerationRareBeginnerBow - генерация редкого лука.
func GenerationRareBeginnerBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Damage:        0,
		TypeWeapon:    equipments.Bow,
		SpeedAttack:   1,
		AttackRadius:  20,
		Weight:        1.1,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateRareBowAllWeight, generateRareBow)

	return &weapon
}

// GenerationEpicBeginnerBow - генерация эпического лука.
func GenerationEpicBeginnerBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Damage:        0,
		TypeWeapon:    equipments.Bow,
		SpeedAttack:   1,
		AttackRadius:  20,
		Weight:        1.1,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateEpicBowAllWeight, generateEpicBow)

	return &weapon
}

// GenerationLegendaryBeginnerBow - генерация легендарного лука.
func GenerationLegendaryBeginnerBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Damage:        0,
		TypeWeapon:    equipments.Bow,
		SpeedAttack:   1,
		AttackRadius:  20,
		Weight:        1.1,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateLegendaryBowAllWeight, generateLegendaryBow)

	return &weapon
}
