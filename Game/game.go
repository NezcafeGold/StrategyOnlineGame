package game

import (
	"./auction"
)

// Game - движок игры.
type Game struct {
	User    *userPage        // Страница пользователей
	Auction *auction.Auction // Страница аукционов
}

// New - создание движка игры.
func New() *Game {
	game := &Game{
		User: newUserPage(),
	}

	game.Auction = auction.New(game.User.ListOnlineUsers)

	return game
}
