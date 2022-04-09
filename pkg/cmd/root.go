package cmd

import (
	"github.com/narumiruna/gofit/pkg/repmax"
	"github.com/narumiruna/gofit/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gofit",

	RunE: Run,
}

func Run(cmd *cobra.Command, args []string) error {
	reps, err := cmd.Flags().GetInt("reps")
	if err != nil {
		return err
	}

	weight, err := cmd.Flags().GetFloat64("weight")
	if err != nil {
		return err
	}

	oneRepMax := repmax.CalculateMayhew1RM(types.KiloGram.MulFloat(weight), reps)

	log.Infof("Weight: %v, repetitions: %v, 1RM: %v", weight, reps, oneRepMax)

	return nil
}

func init() {
	RootCmd.Flags().IntP("reps", "r", 8, "repetitions to fatigue")
	RootCmd.Flags().Float64P("weight", "w", 40.0, "repetition weight in kilogram")
}
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		log.WithError(err).Fatalf("execution failed")
	}
}
