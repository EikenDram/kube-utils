package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update users in Keycloak from a text file",
	Long:  `Update users in Keycloak from a text file`,
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	updateCmd.Flags().StringVarP(&textFile, "file", "f", "", "Name of text file to process")
	updateCmd.MarkFlagRequired("file")
	updateCmd.Flags().StringVarP(&textFormat, "format", "t", "user-pass-raion", "Format of text file")
	updateCmd.Flags().StringVarP(&textDelim, "delim", "d", "space", "Delimiter of text file")
}

func update() {
	//
}
