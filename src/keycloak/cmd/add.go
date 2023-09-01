package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add users into Keycloak from a text file",
	Long:  `Add users into Keycloak from a text file`,
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringVarP(&textFile, "file", "f", "", "Name of text file to process")
	addCmd.MarkFlagRequired("file")
	addCmd.Flags().StringVarP(&textFormat, "format", "t", "user-pass-raion", "Format of text file")
	addCmd.Flags().StringVarP(&textDelim, "delim", "d", "space", "Delimiter of text file")
}

func add() {
	//read ini
	err := loadIni()
	check(err)

	//generate bearer token
	token, err := getToken()
	check(err)
	if len(token) > 0 {
		fmt.Println("Authorization successful")

		//test
		var resp = new([]ResponseUser)
		err = getData(token, "master/users", resp)
		check(err)
		fmt.Println()

	} else {
		check(errors.New("returned empty access token"))
	}
}
