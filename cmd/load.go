package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var template string
var outputFile string

var loadCmd = &cobra.Command{
	Use:   "intercambio",
	Short: "Intercambio analógico de Disparafilm",
	Long:  "Intercambio es una utilidad para hacer un sorteo estilo secret santa entre la comunidad de usuarios de Disparafilm",
	Args:  cobra.MinimumNArgs(1),
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func Execute() {
	if err := loadCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	loadCmd.PersistentFlags().StringVar(&inputFile, "input", "", "csv input, default is ./donations.csv")
	loadCmd.PersistentFlags().StringVar(&outputFile, "output", "", "csv output, default is ./output.csv")
	loadCmd.PersistentFlags().StringVar(&template, "template", "", "email template, no default")
}
