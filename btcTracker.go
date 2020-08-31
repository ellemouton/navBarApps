package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/caseymrm/menuet"
	"github.com/luno/luno-go"
	"github.com/luno/luno-go/decimal"
)

func fetchTicker(ctx context.Context, lc *luno.Client, pair string) (decimal.Decimal, error) {
	resp, err := lc.GetTicker(ctx, &luno.GetTickerRequest{
		Pair: pair,
	})
	if err != nil {
		return decimal.Zero(), err
	}

	return resp.LastTrade, nil
}

func displayTickersForever(lc *luno.Client, pairs []string) {
	for {
		for _, p := range pairs {
			d, err := fetchTicker(context.TODO(), lc, p)
			if err != nil {
				log.Fatal(err)
			}

			menuet.App().SetMenuState(&menuet.MenuState{
				Title: "₿/ZAR:" + strconv.FormatInt(int64(d.Float64()), 10),
			})

			time.Sleep(time.Minute)
		}
	}
}

func main() {
	lc := luno.NewClient()

	menuet.App().Label = "BTC price checker"

	go displayTickersForever(lc, []string{"XBTZAR"})

	menuet.App().RunApplication()
}
