package equipments

type Tool struct {
	ID                  string   `json:"id"                    bson:"_id"`                   // ID.
	IDSet               uint32   `json:"id_set"                bson:"id_set"`                // ID сета.
	Title               string   `json:"title"                 bson:"title"`                 // Название.
	Durability          uint8    `json:"durability"            bson:"durability"`            // Прочность.
	MaxDurability       uint8    `json:"max_durability"        bson:"max_durability"`        // Макс прочность.
	Quality             Quality  `json:"quality"               bson:"quality"`               // Качество.
	TypeTool            TypeTool `json:"type_tool"             bson:"type_tool"`             // Тип инструмента.
	AttackRadius        float32  `json:"attack_radius"         bson:"attack_radius"`         // Радиус атаки.
	ChanceDroppingItems float32  `json:"chance_dropping_items" bson:"chance_dropping_items"` // Шанс дропа предметов.
	SpeedWork           float32  `json:"speed_work"            bson:"speed_work"`            // Скорость работы.
	Weight              float32  `json:"weight"                bson:"weight"`                // Вес.
	RequiredLevel       uint8    `json:"required_level"        bson:"required_level"`        // Требуемый уровень
	CanBeSold           bool     `json:"can_be_sell"           bson:"can_be_sell"`           // Можно ли продать.
}

// ReduceDurability - уменьшить прочность.
func (tool *Tool) ReduceDurability(n uint8)  {
	if tool.Durability > 0 {
		if tool.Durability - n > 0 {
			tool.Durability -= n
		} else {
			tool.Durability = 0
		}
	}
}

// IncreaseDurability - увеличить прочность.
func (tool *Tool) IncreaseDurability(n uint8) {
	if tool.Durability < tool.MaxDurability {
		if tool.Durability + n < tool.MaxDurability {
			tool.Durability += n
		} else {
			tool.Durability = tool.MaxDurability
		}
	}
}