package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

// Execute func.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
