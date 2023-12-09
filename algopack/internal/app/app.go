package app

import (
	"algopack/internal/api"
	"context"
	"fmt"
	"sync"

	"algopack/internal/parse"
	"algopack/pkg/ctxtool"
)

const (
	SBER = "SBER"
	MOEX = "MOEX"
	MGNT = "MGNT"
	AQUA = "AQUA"
	FLOT = "FLOT"
	QIWI = "QIWI"
)

var tickets = [...]string{MOEX, MGNT, AQUA, FLOT, QIWI}

func TradingIteration(ctx context.Context) {
	var wg sync.WaitGroup

	for _, ticket := range tickets {
		wg.Add(1)

		go func(ticket string) {
			defer wg.Done()

			ticketData, err := parse.ParseTicketData(ctx, ticket)
			if err != nil {
				ctxtool.Logger(ctx).Error(err.Error())
				return
			}

			result, err := api.GetPredictByTicket(ctx, ticketData)
			if err != nil {
				ctxtool.Logger(ctx).Error(err.Error())
				return
			}

			ctxtool.Logger(ctx).Info(fmt.Sprintf("result for ticket %s is %f", result.Title, result.Predict))
		}(ticket)
	}
	wg.Wait()
}
