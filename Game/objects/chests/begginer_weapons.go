package chests

import (
	"../../utils/generate_items"
	"../equipments"
	"../sets/begginer/weapons"
	"gopkg.in/mgo.v2/bson"
)

var (
	GenerateChestSetBeginnerWeapons = []generate_items.GenerateObjFloat32 {
		{5.0, 1},	// Лук
		{6.0, 1},	// Короткий лук
		//tools.GenerateObjFloat32{7.0, 1},	// Длинный лук
		//
		//tools.GenerateObjFloat32{8.0, 1},	// Меч
		//tools.GenerateObjFloat32{9.0, 1},	// Двуручный меч
		//tools.GenerateObjFloat32{10.0, 1},	// Два меча
		//
		//tools.GenerateObjFloat32{11.0, 1},	// Копье
		//tools.GenerateObjFloat32{12.0, 1},	// Короткое копье
		//tools.GenerateObjFloat32{13.0, 1},	// Длинное копье
		//
		//tools.GenerateObjFloat32{14.0, 1},	// Алебарда
		//tools.GenerateObjFloat32{15.0, 1},	// Кинжалы
		//
		//tools.GenerateObjFloat32{16.0, 1},	// Посох
		//tools.GenerateObjFloat32{17.0, 1},	// Маг книга
		//tools.GenerateObjFloat32{18.0, 1},	// Кристалл
	}

	GenerateChestSetBeginnerWeaponsAllWeight = generate_items.WeightSumObjFloat32(GenerateChestSetBeginnerWeapons)
)

// NewChestSetBeginnerWeapons - создать сундук с оружием новичка.
func NewChestSetBeginnerWeapons() *Chest {
	return &Chest{
		ID:        bson.NewObjectId().Hex(),
		Title:     "Сундук c оружием новичка",
		TypeChest: BeginnerWeapons,
		IsOpen:    false,
		CanBeSold: true,
	}
}

// OpenChestSetBeginnerWeapons - открыть сундук с оружием новичка.
func OpenChestSetBeginnerWeapons() (interface{}, bool) {
	typeEquipment := generate_items.GenerateValueObjFloat32(GenerateChestSetBeginnerWeaponsAllWeight, GenerateChestSetBeginnerWeapons)

	switch typeEquipment {
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
	//case 7.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerLongBow(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerLongBow(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerLongBow(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerLongBow(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerLongBow(), true
	//	}
	//}
	//
	//case 8.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerSword(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerSword(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerSword(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerSword(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerSword(), true
	//	}
	//}
	//case 9.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerTwoHandedSword(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerTwoHandedSword(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerTwoHandedSword(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerTwoHandedSword(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerTwoHandedSword(), true
	//	}
	//}
	//case 10.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerDoubleSwords(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerDoubleSwords(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerDoubleSwords(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerDoubleSwords(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerDoubleSwords(), true
	//	}
	//}
	//
	//case 11.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerSpear(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerSpear(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerSpear(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerSpear(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerSpear(), true
	//	}
	//}
	//case 12.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerShortSpear(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerShortSpear(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerShortSpear(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerShortSpear(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerShortSpear(), true
	//	}
	//}
	//case 13.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerLongSpear(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerLongSpear(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerLongSpear(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerLongSpear(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerLongSpear(), true
	//	}
	//}
	//
	//case 14.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerHalberd(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerHalberd(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerHalberd(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerHalberd(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerHalberd(), true
	//	}
	//}
	//case 15.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerDaggers(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerDaggers(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerDaggers(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerDaggers(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerDaggers(), true
	//	}
	//}
	//
	//case 16.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerStaff(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerStaff(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerStaff(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerStaff(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerStaff(), true
	//	}
	//}
	//case 17.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerMagicBook(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerMagicBook(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerMagicBook(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerMagicBook(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerMagicBook(), true
	//	}
	//}
	//case 18.0: {
	//	var quality equipments.Quality
	//
	//	var count int
	//	rnd := crnd.Intn(GenerateQualityAllWeight) + 1
	//
	//	for _, GenerateObjQuality := range GenerateQuality {
	//		count += GenerateObjQuality.Weight
	//
	//		if count >= rnd {
	//			quality = GenerateObjQuality.Value
	//			break
	//		}
	//	}
	//
	//	switch quality {
	//	case equipments.Usual:
	//		return weapons.GenerationUsualBeginnerCrystal(), true
	//	case equipments.Unusual:
	//		return weapons.GenerationUnusualBeginnerCrystal(), true
	//	case equipments.Rare:
	//		return weapons.GenerationRareBeginnerCrystal(), true
	//	case equipments.Epic:
	//		return weapons.GenerationEpicBeginnerCrystal(), true
	//	case equipments.Legendary:
	//		return weapons.GenerationLegendaryBeginnerCrystal(), true
	//	}
	//}
	}

	return nil, false
}