package units

import (
	"../../utils/generate_items"
	"../resources"
	"./stats"
	"sync"
)

// newStats - создание новых всех стат.
func newStats() stats.Stats {
	return stats.Stats{
		GeneralStats: generateGeneralStats(),
		BattleStats:  generateBattleStats(),
		WorkingStats: generateWorkingStats(),
	}
}

// generateBattleStats - создание новых боевых стат.
func generateBattleStats() stats.BattleStats {
	return stats.BattleStats{
		PhysicallyStats: generatePhysicallyStats(),
	}
}

// generatePhysicallyStats - генерация физ стат.
func generatePhysicallyStats() stats.PhysicallyStats {
	physicallyStats := stats.PhysicallyStats{}
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Генерация физ силы.
	go func() {
		defer wg.Done()

		physicallyStats.Power = generate_items.GenerateValueObjFloat32(generatePowerAllWeight, generatePower)
	}()

	// Генерация крит шанса.
	go func() {
		defer wg.Done()

		physicallyStats.CriticalAttackChance = generate_items.GenerateValueObjFloat32(generateCriticalChanceAllWeight, generateCriticalChance)
	}()

	// Процента крит атаки.
	go func() {
		defer wg.Done()

		physicallyStats.CriticalAttackPercent = generate_items.GenerateValueObjFloat32(generateCriticalAttackPercentAllWeight, generateCriticalAttackPercent)
	}()

	physicallyStats.Level = 1
	physicallyStats.Experience = 0

	wg.Wait()

	return physicallyStats
}

// generateMagicStats - генерация рабочих стат.
func generateWorkingStats() stats.WorkingStats {
	workingStats := stats.WorkingStats{}
	wg := sync.WaitGroup{}
	wg.Add(6)

	// Генерация скорости строительства.
	go func() {
		defer wg.Done()

		speedBuilding := generate_items.GenerateValueObjFloat32(generateBuildingAllWeight, generateBuilding)
		workingStats.Building = stats.Building{
			Level:         1,
			Experience:    0,
			SpeedBuilding: speedBuilding,
		}
	}()

	// Генерация скорости сбора еды.
	go func() {
		defer wg.Done()

		speedGathering := generate_items.GenerateValueObjFloat32(generateGatheringFoodAllWeight, generateGatheringFood)
		workingStats.GatheringFood = stats.GatheringResources{
			Level:          1,
			Experience:     0,
			SpeedGathering: speedGathering,
		}
	}()

	// Генерация скорости сбора дерева.
	go func() {
		defer wg.Done()

		speedGathering := generate_items.GenerateValueObjFloat32(generateGatheringWoodAllWeight, generateGatheringWood)
		workingStats.GatheringFood = stats.GatheringResources{
			Level:          1,
			Experience:     0,
			SpeedGathering: speedGathering,
		}
	}()

	// Генерация скорости сбора камня.
	go func() {
		defer wg.Done()

		speedGathering := generate_items.GenerateValueObjFloat32(generateGatheringStoneAllWeight, generateGatheringStone)
		workingStats.GatheringFood = stats.GatheringResources{
			Level:          1,
			Experience:     0,
			SpeedGathering: speedGathering,
		}
	}()

	// Генерация скорости сбора Железа.
	go func() {
		defer wg.Done()

		speedGathering := generate_items.GenerateValueObjFloat32(generateGatheringIronAllWeight, generateGatheringIron)
		workingStats.GatheringFood = stats.GatheringResources{
			Level:          1,
			Experience:     0,
			SpeedGathering: speedGathering,
		}
	}()

	// Генерация скорости сбора золота.
	go func() {
		defer wg.Done()

		speedGathering := generate_items.GenerateValueObjFloat32(generateGatheringGoldAllWeight, generateGatheringGold)
		workingStats.GatheringFood = stats.GatheringResources{
			Level:          1,
			Experience:     0,
			SpeedGathering: speedGathering,
		}
	}()

	wg.Wait()

	return workingStats
}

// generateGeneralStats - генерация общих стат.
func generateGeneralStats() stats.GeneralStats {
	generalStats := stats.GeneralStats{}
	wg := sync.WaitGroup{}
	wg.Add(8)

	// Генерация хп.
	go func() {
		defer wg.Done()

		generalStats.HP = generate_items.GenerateValueObjFloat32(generateHPAllWeight, generateHP)
	}()

	// Генерация регенерации хп.
	go func() {
		defer wg.Done()

		generalStats.RegenerationHP = generate_items.GenerateValueObjFloat32(generateRegenerationHPAllWeight, generateRegenerationHP)
	}()

	// Генерация стамины.
	go func() {
		defer wg.Done()

		generalStats.Stamina = generate_items.GenerateValueObjFloat32(generateStaminaAllWeight, generateStamina)
	}()

	// Генерация регенерации стамины.
	go func() {
		defer wg.Done()

		generalStats.RegenerationStamina = generate_items.GenerateValueObjFloat32(generateRegenerationStaminaAllWeight, generateRegenerationStamina)
	}()

	// Генерация потребления еды.
	go func() {
		defer wg.Done()

		generalStats.FoodWaste = resources.Food(generate_items.GenerateValueObjFloat32(generateFoodWasteAllWeight, generateFoodWaste))
	}()

	// Генерация укланения.
	go func() {
		defer wg.Done()

		generalStats.Agility = generate_items.GenerateValueObjFloat32(generateAgilityAllWeight, generateAgility)
	}()

	// Генерация ловкости.
	go func() {
		defer wg.Done()

		generalStats.Dexterity = generate_items.GenerateValueObjFloat32(generateDexterityAllWeight, generateDexterity)
	}()

	// Генерация меткости.
	go func() {
		defer wg.Done()

		generalStats.Accuracy = generate_items.GenerateValueObjFloat32(generateAccuracyAllWeight, generateAccuracy)
	}()

	wg.Wait()

	return generalStats
}
