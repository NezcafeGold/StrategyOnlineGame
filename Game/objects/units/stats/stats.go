package stats

import (
	"../../resources"
)

// Stats - статы юнита.
type Stats struct {
	GeneralStats GeneralStats `json:"general_stats" bson:"general_stats"` // Общие статы.
	BattleStats  BattleStats  `json:"battle_stats"  bson:"battle_stats"`  // Боевые статы.
	WorkingStats WorkingStats `json:"working_stats" bson:"working_stats"` // Рабочие статы.
}

// GeneralStats - общие статы.
type GeneralStats struct {
	SpeedAttack         float32        `json:"speed_attack"         bson:"speed_attack"`         // Скорость атаки.
	SpeedMove           float32        `json:"speed"                bson:"speed"`                // Скорость передвижения.
	AttackRadius        float32        `json:"attack_radius"        bson:"attack_radius"`        // Радиус атаки (от оружия).
	HP                  float32        `json:"hp"                   bson:"hp"`                   // Хп.
	Defense             float32        `json:"defense"              bson:"defense"`              // Защита.
	Weight              float32        `json:"weight"               bson:"weight"`               // Вес.
	RegenerationHP      float32        `json:"regeneration_hp"      bson:"regeneration_hp"`      // Регенерация хп в сек.
	Stamina             float32        `json:"stamina"              bson:"stamina"`              // Стамина.
	RegenerationStamina float32        `json:"regeneration_stamina" bson:"regeneration_stamina"` // Регенерация стамины в сек.
	FoodWaste           resources.Food `json:"food_waste"           bson:"food_waste"`           // Потребляемая еды в минуту.
	Dexterity           float32        `json:"dexterity"            bson:"dexterity"`            // Ловкость.
	Accuracy            float32        `json:"accuracy"             bson:"accuracy"`             // Меткость.
	Agility             float32        `json:"aqility"              bson:"aqility"`              // Укланение.
	Vampirism           float32        `json:"vampirism"            bson:"vampirism"`            // Вампиризм.
}


