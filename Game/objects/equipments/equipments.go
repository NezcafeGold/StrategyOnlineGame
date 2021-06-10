package equipments

type Quality int      // Качество.
type TypeTool int     // Тип инструмента.
type TypeWeapon int   // Тип оружия.
type TypeArmor int    // Тип брони.

// Качество.
const (
	Usual Quality = iota     // Обычный.
	Unusual Quality = iota   // Необычный.
	Rare Quality = iota      // Редкий.
	Epic Quality = iota      // Эпический.
	Legendary Quality = iota // Легендарный.
)

// Тип инструментов.
const (
	Axe TypeTool = iota     // Топор.
	Pickaxe TypeTool = iota // Кирка.
	Sickle TypeTool = iota  // Серп.
	Hammer TypeTool = iota  // Молоток.
)

// Тип брони.
const (
	Helmet TypeArmor = iota     // Шлем.
	Chestplate TypeArmor = iota // Нагрудник.
	Leggings TypeArmor = iota   // Штаны.
	Boots TypeArmor = iota      // Ботинки
)

// Тип оружия.
const (
	Bow TypeWeapon = iota      // Лук.
	ShortBow TypeWeapon = iota // Короткий лук.
)

// Equipment - снаряжение.
type Equipment struct {
	Weapon    *Weapon   `json:"weapon"    bson:"weapon"`    // Оружие
	Tool      *Tool     `json:"tool"      bson:"tool"`      // Инструмент
	Armors    Armors    `json:"armors"    bson:"armors"`    // Броня
}

// New - создать пустое снаряжение.
func New() Equipment {
	return Equipment{
		Armors: NewArmors(),
	}
}
