package summoning_scrolls

import (
	"../units"
	"gopkg.in/mgo.v2/bson"
)

// SummoningScrolls - свиток призыва.
type SummoningScrolls struct {
	ID        string    `json:"id"          bson:"_id"`         // ID.
	Title     string    `json:"title"       bson:"title"`       // Название.
	IsUse     bool      `json:"is_use"      bson:"is_use"`      // Статус.
	CanBeSold bool      `json:"can_be_sold" bson:"can_be_sold"` // Можно ли продать.
}

// New - создать свиток призыва.
func New() *SummoningScrolls {
	return &SummoningScrolls{
		ID:        bson.NewObjectId().Hex(),
		Title:     "Свиток призыва юнита",
		IsUse:     false,
		CanBeSold: true,
	}
}

// Use - использовать свиток призыва.
func (summoningScrolls *SummoningScrolls) Use() *units.Unit {
	return units.New()
}
