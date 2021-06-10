package stats

// WorkingStats - рабочие статы.
type WorkingStats struct {
	GatheringFood  GatheringResources `json:"gathering_food"  bson:"gathering_food"`  // Добаыча еды.
	GatheringWood  GatheringResources `json:"gathering_wood"  bson:"gathering_wood"`  // Добаыча дерева.
	GatheringStone GatheringResources `json:"gathering_stone" bson:"gathering_stone"` // Добаыча камня.
	GatheringIron  GatheringResources `json:"gathering_iron"  bson:"gathering_iron"`  // Добаыча железа.
	GatheringGold  GatheringResources `json:"gathering_gold"  bson:"gathering_gold"`  // Добаыча золота.
	Building       Building           `json:"building"        bson:"building"`        // Строительство.
}

// GatheringResources - статы на сбор ресурсов.
type GatheringResources struct {
	Level               uint8   `json:"level"                 bson:"level"`                 // Уровень.
	Experience          uint32  `json:"experience"            bson:"experience"`            // Опыт.
	SpeedGathering      float32 `json:"speed_gathering"       bson:"speed_gathering"`       // Скорость добычи.
	ChanceDroppingItems float32 `json:"chance_dropping_items" bson:"chance_dropping_items"` // Шанс дропа предметов.
}

// Building - статы на строительство.
type Building struct {
	Level         uint8   `json:"level"          bson:"level"`          // Уровень.
	Experience    uint32  `json:"experience"     bson:"experience"`     // Опыт.
	SpeedBuilding float32 `json:"speed_building" bson:"speed_building"` // Скорость строительства.
}
