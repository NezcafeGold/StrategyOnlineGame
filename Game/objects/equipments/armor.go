package equipments

// Armor - броня.
type Armor struct {
	ID            string    `json:"id"             bson:"_id"`            // ID
	IDSet         uint32    `json:"id_set"         bson:"id_set"`         // ID сета
	Title         string    `json:"title"          bson:"title"`          // Название
	TypeArmor     TypeArmor `json:"type_armor"     bson:"type_armor"`     // Тип брони
	Durability    uint8     `json:"durability"     bson:"durability"`     // Прочность
	MaxDurability uint8     `json:"max_durability" bson:"max_durability"` // Макс прочность.
	Quality       Quality   `json:"quality"        bson:"quality"`        // Качество
	Defense       float32   `json:"Defense"        bson:"Defense"`        // Защита
	Weight        float32   `json:"weight"         bson:"weight"`         // Вес
	RequiredLevel uint8     `json:"required_level" bson:"required_level"` // Требуемый уровень
	CanBeSold     bool      `json:"can_be_sell"    bson:"can_be_sell"`    // Можно ли продать
}


// Armors - комплект брони.
type Armors struct {
	Helmet     *Armor `json:"helmet"     bson:"helmet"`     // Шлем
	Chestplate *Armor `json:"chestplate" bson:"chestplate"` // Нагрудник
	Leggings   *Armor `json:"leggings"   bson:"leggings"`   // Штаны
	Boots      *Armor `json:"boots"      bson:"boots"`      // Ботинки
}

// NewArmors - создать пустой сет брони.
func NewArmors() Armors {
	return Armors{}
}

// ReduceDurabilityArmor - уменьшить прочность.
func (armor *Armor) ReduceDurabilityArmor(n uint8)  {
	if armor.Durability > 0 {
		if armor.Durability - n > 0 {
			armor.Durability -= n
		} else {
			armor.Durability = 0
		}
	}
}

// IncreaseDurabilityArmor - увеличить прочность.
func (armor *Armor) IncreaseDurabilityArmor(n uint8) {
	if armor.Durability < armor.MaxDurability {
		if armor.Durability + n < armor.MaxDurability {
			armor.Durability += n
		} else {
			armor.Durability = armor.MaxDurability
		}
	}
}