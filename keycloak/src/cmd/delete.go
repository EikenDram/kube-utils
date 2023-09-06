package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete users in Keycloak from a text file",
	Long: `Delete users in Keycloak from a text file
		By default will only disable users,
		to permanently delete users use --permanent flag`,
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

var permanent bool = false

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&textFile, "file", "f", "", "Name of text file to process")
	deleteCmd.MarkFlagRequired("file")
	deleteCmd.Flags().BoolVarP(&permanent, "permanent", "p", false, "Permanently delete users instead of disabling them")

}

func delete() {
	//
}
