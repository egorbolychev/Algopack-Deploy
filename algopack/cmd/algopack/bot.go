package algopack

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"algopack/internal/app"
	"algopack/pkg/ctxtool"
)

var (
	workTime int
)

var RootCommand = &cobra.Command{
	Use:   "algopack",
	Short: "Run algopack investment bot",

	Run: runTrade,
}

func init() {
	RootCommand.PersistentFlags().IntVarP(
		&workTime, "time", "t", 0, "the time for which the worker will be launched",
	)
	if err := RootCommand.MarkPersistentFlagRequired("time"); err != nil {
		log.Fatal(err)
	}
}

func runTrade(_ *cobra.Command, _ []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctxWithLogger, err := configureLogger(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctxtool.Logger(ctxWithLogger).Info("Starting Main iteration")

	go mainIteration(ctxWithLogger)

	time.Sleep(time.Duration(workTime) * time.Minute)

	ctxtool.Logger(ctxWithLogger).Info("Bot stopped")
}

func mainIteration(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			ctxtool.Logger(ctx).Info("new iteration started")
			app.TradingIteration(ctx)
		}
		time.Sleep(10 * time.Second)
	}
}

func configureLogger(ctx context.Context) (context.Context, error) {
	logCfg := zap.NewProductionConfig()
	logCfg.OutputPaths = []string{"stdout"}

	logger, err := logCfg.Build()
	if err != nil {
		return nil, err
	}

	logCtx := ctxtool.WithLogger(ctx, logger)
	return logCtx, nil
}
