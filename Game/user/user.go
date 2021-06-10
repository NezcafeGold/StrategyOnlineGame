package user

import (
	"../player"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// User - пользователь.
type User struct {
	ID           string         `json:"id"             bson:"_id"`            // ID.
	Email        string         `json:"email"          bson:"email"`          // Почта.
	Nickname     string         `json:"nickname"       bson:"nickname"`       // Ник.
	Player       *player.Player `json:"player"         bson:"player"`         // Структура игрока.
	LastSyncDate time.Time      `json:"last_sync_date" bson:"last_sync_date"` // Дата последний синхронизации
}

// New - создание нового аккаунта.
func New(email, nickname string) *User {
	id := bson.NewObjectId().Hex()

	return &User{
		ID:       id,
		Email:    email,
		Nickname: nickname,
		Player:   player.New(id),
	}
}

// Syncing - синзронизация пользователя.
func (user *User) Syncing() {
	user.Player.Inventory.UpdateSize()
	user.Player.Units.UpdateSize()

	user.LastSyncDate = time.Now()
}