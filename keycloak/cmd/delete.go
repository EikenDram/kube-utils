package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete users in Keycloak from a text file",
	Long:  `Delete users in Keycloak from a text file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&textFile, "file", "f", "", "Name of text file to process")
	deleteCmd.MarkFlagRequired("file")
}
