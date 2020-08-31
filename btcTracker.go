package main

import (
	"context"
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
	var (
		display  string
		duration time.Duration
	)

	for {
		for _, p := range pairs {
			d, err := fetchTicker(context.TODO(), lc, p)
			if err != nil {
				// Probably no internet connection. Sleep and try again later.
				duration = time.Minute * 5
				display = "offline"
			} else {
				duration = time.Minute
				display = strconv.FormatInt(int64(d.Float64()), 10)
			}

			menuet.App().SetMenuState(&menuet.MenuState{
				Title: "₿/ZAR:" + display,
			})

			time.Sleep(duration)
		}
	}
}

func main() {
	lc := luno.NewClient()

	menuet.App().Label = "BTC price checker"

	go displayTickersForever(lc, []string{"XBTZAR"})

	menuet.App().RunApplication()
}
