package equipments

type Weapon struct {
	ID            string     `json:"id"             bson:"_id"`            // ID.
	IDSet         uint32     `json:"id_set"         bson:"id_set"`         // ID сета.
	Title         string     `json:"title"          bson:"title"`          // Название.
	Durability    uint8      `json:"durability"     bson:"durability"`     // Прочность.
	MaxDurability uint8      `json:"max_durability" bson:"max_durability"` // Макс прочность.
	Quality       Quality    `json:"quality"        bson:"quality"`        // Качество.
	TypeWeapon    TypeWeapon `json:"type_weapon"    bson:"type_weapon"`    // Тип оружия.
	Damage        float32    `json:"damage"         bson:"damage"`         // Наносимый урон.
	SpeedAttack   float32    `json:"speed_attack"   bson:"speed_attack"`   // Скорость атаки.
	AttackRadius  float32    `json:"attack_radius"  bson:"attack_radius"`  // Радиус атаки.
	Weight        float32    `json:"weight"         bson:"weight"`         // Вес.
	RequiredLevel uint8      `json:"required_level" bson:"required_level"` // Требуемый уровень.
	CanBeSold     bool       `json:"can_be_sell"    bson:"can_be_sell"`    // Можно ли продать.
}

// ReduceDurability - уменьшить прочность.
func (weapon *Weapon) ReduceDurability(n uint8)  {
	if weapon.Durability > 0 {
		if weapon.Durability - n > 0 {
			weapon.Durability -= n
		} else {
			weapon.Durability = 0
		}
	}
}

// IncreaseDurability - увеличить прочность.
func (weapon *Weapon) IncreaseDurability(n uint8) {
	if weapon.Durability < weapon.MaxDurability {
		if weapon.Durability + n < weapon.MaxDurability {
			weapon.Durability += n
		} else {
			weapon.Durability = weapon.MaxDurability
		}
	}
}