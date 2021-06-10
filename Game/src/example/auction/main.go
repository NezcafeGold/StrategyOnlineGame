package main

import (
	"../../../../Game"
	"../../../utils/logger"
)

func main() {
	g := game.New()

	enot, err := g.User.Get("jklgreentea@gmail.com")
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	kot, err := g.User.Get("jklkot@gmail.com")
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}

	{
		err = g.Auction.Resources.Asks.AddFoodOrder(1.2, 1000, enot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
		err = g.Auction.Resources.Asks.AddStoneOrder(1.3, 1000, enot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
		err = g.Auction.Resources.Asks.AddFoodOrder(1.1, 1000, enot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
	}

	{
		err = g.Auction.Resources.Asks.AddStoneOrder(1, 1000, kot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
		err = g.Auction.Resources.Asks.AddIronOrder(1.4, 1000, kot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
		err = g.Auction.Resources.Asks.AddWoodOrder(1.2, 1000, kot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
	}

	{
		err = g.User.Disconnect(enot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}

		err = g.User.Disconnect(kot)
		if err != nil {
			logger.Logger.Error(err.Error())
		}
	}

	g.Auction.Syncing()
}
