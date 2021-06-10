package stats

// BattleStats - боевые статы.
type BattleStats struct {
	PhysicallyStats PhysicallyStats `json:"physically_stats" bson:"physically_stats"` // Физ
}

// PhysicallyStats - физ статы.
type PhysicallyStats struct {
	Level                 uint8   `json:"level"                   bson:"level"`                   // Уровень.
	Experience            uint32  `json:"experience"              bson:"experience"`              // Опыт.
	Power                 float32 `json:"power"                   bson:"power"`                   // Сила.
	CriticalAttackChance  float32 `json:"critical_attack_chance"  bson:"critical_attack_chance"`  // Шанс крит урона.
	CriticalAttackPercent float32 `json:"critical_attack_percent" bson:"critical_attack_percent"` // Крит урон.
}

