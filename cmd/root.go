package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
        "github.com/MalloZup/pipelingo/jenkins"

)

var rootCmd = &cobra.Command{
  Use:   "pipelingo",
  Short: "pipelingo create jenkins pipelines",
  Long: `very long desc for pipelingo`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("beeing called by cobra pipelingo, creating jenkins pipeline for fun and profit\n")
    jenkins.Login()
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
