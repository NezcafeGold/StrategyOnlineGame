package weapons

import (
	"../../../../utils/generate_items"
	"../../../equipments"
	"../../../sets"
	"gopkg.in/mgo.v2/bson"
)

var (
	generateUsualShortBow = []generate_items.GenerateObjFloat32 {
		{7.0, 50},
		{8.0, 30},
		{9.0, 10},
		{10.0, 5},
		{11.0, 1},
	}
	generateUsualShortBowAllWeight = generate_items.WeightSumObjFloat32(generateUsualShortBow)

	generateUnusualShortBow = []generate_items.GenerateObjFloat32 {
		{12.0, 50},
		{13.0, 30},
		{14.0, 10},
		{15.0, 5},
		{16.0, 1},
	}
	generateUnusualShortBowAllWeight = generate_items.WeightSumObjFloat32(generateUnusualShortBow)

	generateRareShortBow = []generate_items.GenerateObjFloat32 {
		{17.0, 50},
		{18.0, 30},
		{19.0, 10},
		{20.0, 5},
		{21.0, 1},
	}
	generateRareShortBowAllWeight = generate_items.WeightSumObjFloat32(generateRareShortBow)

	generateEpicShortBow = []generate_items.GenerateObjFloat32 {
		{22.0, 50},
		{23.0, 30},
		{24.0, 10},
		{25.0, 5},
		{26.0, 1},
	}
	generateEpicShortBowAllWeight = generate_items.WeightSumObjFloat32(generateEpicShortBow)

	generateLegendaryShortBow = []generate_items.GenerateObjFloat32 {
		{27.0, 50},
		{28.0, 30},
		{29.0, 10},
		{30.0, 5},
		{31.0, 1},
	}
	generateLegendaryShortBowAllWeight = generate_items.WeightSumObjFloat32(generateLegendaryShortBow)
)

// GenerationUsualBeginnerShortBow - генерация обычного короткого лука.
func GenerationUsualBeginnerShortBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Короткий лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Usual,
		Damage:        0,
		TypeWeapon:    equipments.ShortBow,
		SpeedAttack:   2,
		AttackRadius:  15,
		Weight:        0.7,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateUsualShortBowAllWeight, generateUsualShortBow)

	return &weapon
}

// GenerationUnusualBeginnerShortBow - генерация необычного короткого лука.
func GenerationUnusualBeginnerShortBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Короткий лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Unusual,
		Damage:        0,
		TypeWeapon:    equipments.ShortBow,
		SpeedAttack:   2,
		AttackRadius:  15,
		Weight:        0.7,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateUnusualShortBowAllWeight, generateUnusualShortBow)

	return &weapon
}

// GenerationRareBeginnerShortBow - генерация редкого короткого лука.
func GenerationRareBeginnerShortBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Короткий лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Rare,
		Damage:        0,
		TypeWeapon:    equipments.ShortBow,
		SpeedAttack:   2,
		AttackRadius:  15,
		Weight:        0.7,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateRareShortBowAllWeight, generateRareShortBow)

	return &weapon
}

// GenerationEpicBeginnerShortBow - генерация эпического короткого лука.
func GenerationEpicBeginnerShortBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Короткий лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Epic,
		Damage:        0,
		TypeWeapon:    equipments.ShortBow,
		SpeedAttack:   2,
		AttackRadius:  15,
		Weight:        0.7,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateEpicShortBowAllWeight, generateEpicShortBow)

	return &weapon
}

// GenerationLegendaryBeginnerShortBow - генерация легендарного короткого лука.
func GenerationLegendaryBeginnerShortBow() *equipments.Weapon {
	weapon := equipments.Weapon{
		ID:            bson.NewObjectId().Hex(),
		IDSet:         sets.Begginer,
		Title:         "Короткий лук новичка",
		Durability:    100,
		MaxDurability: 100,
		Quality:       equipments.Legendary,
		Damage:        0,
		TypeWeapon:    equipments.ShortBow,
		SpeedAttack:   2,
		AttackRadius:  15,
		Weight:        0.7,
		RequiredLevel: 1,
		CanBeSold:     true,
	}

	weapon.Damage = generate_items.GenerateValueObjFloat32(generateLegendaryShortBowAllWeight, generateLegendaryShortBow)

	return &weapon
}
