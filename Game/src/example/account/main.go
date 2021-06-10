package main

import (
	"../../../../Game"
	"../../../utils/logger"
)

func main()  {
	g := game.New()

	u := g.User.New("jklgreentea@gmail.com", "JkLEnot")
	err := g.User.Save(*u)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	u1 := g.User.New("jklkot@gmail.com", "JkLKot")
	err = g.User.Save(*u1)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
}