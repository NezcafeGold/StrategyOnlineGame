package units

import (
	"../../utils/generate_items"
	"../equipments"
	"./stats"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
)

type Gender int

const (
	Man Gender = iota   // Мужской
	Woman Gender = iota // Женский
)

// Unit - юнит.
type Unit struct {
	ID         string               `json:"id"          bson:"_id"`         // ID
	Name       string               `json:"name"        bson:"name"`        // Имя
	Gender     Gender               `json:"gender"      bson:"gender"`      // Пол
	Level      uint8                `json:"level"       bson:"level"`       // Основный уровень
	Experience uint32               `json:"experience"  bson:"experience"`  // Опыт
	Stats      stats.Stats          `json:"stats"       bson:"stats"`       // Статы
	Equipment  equipments.Equipment `json:"equipments"  bson:"equipments"`  // Снаряжение
	CanBeSold  bool                 `json:"can_be_sold" bson:"can_be_sold"` // Можно ли продать
}

// New - создание нового юнита расы "Человек".
func New() *Unit {
	var src generate_items.CryptoSource
	crnd := rand.New(src)

	switch crnd.Intn(2) {
	case 0: return NewMan()
	case 1: return NewWoman()
	}

	return &Unit{}
}

// NewMan - создание нового юнита расы "Человек" мужского пола.
func NewMan() *Unit {
	return &Unit{
		ID:         bson.NewObjectId().Hex(),
		Name:       "Man",
		Gender:     Man,
		Level:      1,
		Experience: 0,
		Stats:      newStats(),
		Equipment:  equipments.New(),
	}
}

// NewWoman - создание нового юнита расы "Человек" женского пола.
func NewWoman() *Unit {
	return &Unit{
		ID:         bson.NewObjectId().Hex(),
		Name:       "Woman",
		Gender:     Woman,
		Level:      1,
		Experience: 0,
		Stats:      newStats(),
		Equipment:  equipments.New(),
	}
}

// PutOnWeapon - надеть оружие.
func (unit *Unit) PutOnWeapon(weapon *equipments.Weapon) error {
	if unit.Equipment.Weapon != nil {
		return errors.New("Оружие уже надето. ")
	}

	unit.Equipment.Weapon = weapon

	return nil
}

// PutOutWeapon - снять оружие.
func (unit *Unit) PutOutWeapon() (*equipments.Weapon, error) {
	if unit.Equipment.Weapon != nil {
		weapon := unit.Equipment.Weapon
		unit.Equipment.Weapon = nil

		return weapon, nil
	}

	return nil, errors.New("Оружие нельзя снять. ")
}

// PutOnTool - надеть инструмент.
func (unit *Unit) PutOnTool(tool *equipments.Tool) error {
	if unit.Equipment.Tool != nil {
		return errors.New("Инструмент уже надет. ")
	}

	unit.Equipment.Tool = tool

	return nil
}

// PutOutTool - снять инструмент.
func (unit *Unit) PutOutTool() (*equipments.Tool, error) {
	if unit.Equipment.Tool != nil {
		tool := unit.Equipment.Tool
		unit.Equipment.Tool = nil

		return tool, nil
	}

	return nil, errors.New("Нельзя снять инструмент. ")
}

// PutOnHelmet - надеть шлем.
func (unit *Unit) PutOnHelmet(armor *equipments.Armor) error {
	if unit.Equipment.Armors.Helmet != nil {
		return errors.New("Шлем уже надет. ")
	}

	unit.Equipment.Armors.Helmet = armor

	return nil
}

// PutOutHelmet - снять шлем.
func (unit *Unit) PutOutHelmet() (*equipments.Armor, error) {
	if unit.Equipment.Armors.Helmet != nil {
		armor := unit.Equipment.Armors.Helmet
		unit.Equipment.Armors.Helmet = nil

		return armor, nil
	}

	return nil, errors.New("Нельзя снять шлем. ")
}

// PutOnChestplate - надеть нагрудник.
func (unit *Unit) PutOnChestplate(armor *equipments.Armor) error {
	if unit.Equipment.Armors.Chestplate != nil {
		return errors.New("Нагрудник уже надет. ")
	}

	unit.Equipment.Armors.Chestplate = armor

	return nil
}

// PutOutChestplate - снять нагрудник.
func (unit *Unit) PutOutChestplate() (*equipments.Armor, error) {
	if unit.Equipment.Armors.Chestplate != nil {
		armor := unit.Equipment.Armors.Chestplate
		unit.Equipment.Armors.Chestplate = nil

		return armor, nil
	}

	return nil, errors.New("Нельзя снять нагрудник. ")
}

// PutOnLeggings - надеть штаны.
func (unit *Unit) PutOnLeggings(armor *equipments.Armor) error {
	if unit.Equipment.Armors.Leggings != nil {
		return errors.New("Штаны уже надеты. ")
	}

	unit.Equipment.Armors.Leggings = armor

	return nil
}

// PutOutLeggings - снять штаны.
func (unit *Unit) PutOutLeggings() (*equipments.Armor, error) {
	if unit.Equipment.Armors.Leggings != nil {
		armor := unit.Equipment.Armors.Leggings
		unit.Equipment.Armors.Leggings = nil

		return armor, nil
	}

	return nil, errors.New("Нельзя снять штаны. ")
}

// PutOnBoots - надеть ботинки.
func (unit *Unit) PutOnBoots(armor *equipments.Armor) error {
	if unit.Equipment.Armors.Boots != nil {
		return errors.New("Ботинки уже надеты. ")
	}

	unit.Equipment.Armors.Boots = armor

	return nil
}

// PutOutBoots - снять ботинки.
func (unit *Unit) PutOutBoots() (*equipments.Armor, error) {
	if unit.Equipment.Armors.Boots != nil {
		armor := unit.Equipment.Armors.Boots
		unit.Equipment.Armors.Boots = nil

		return armor, nil
	}

	return nil, errors.New("Нельзя снять ботинки. ")
}
