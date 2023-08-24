package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize in Keycloak server",
	Long:  `Creates local authorization config for a Keycloak server`,
	Run: func(cmd *cobra.Command, args []string) {
		authorize()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	authCmd.Flags().StringVarP(&keycloakURL, "url", "u", "", "Keycloak URL address")
	authCmd.MarkFlagRequired("url")
	authCmd.Flags().StringVarP(&clientId, "client", "c", "admin-cli", "Keycloak admin-cli client id")
	authCmd.Flags().StringVarP(&clientToken, "token", "t", "", "Keycloak admin-cli client token")
	authCmd.MarkFlagRequired("token")
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func authorize() {
	//test if credentials work
	fmt.Println("Trying to authorize in Keycloak...")
	token, err := getToken()
	check(err)

	//if work, save credentials in .credentials file
	if len(token) > 0 {
		err = saveIni()
		check(err)
		fmt.Println("Authorization successful, configuration saved to .keycloak file")

	} else {
		check(errors.New("returned empty access token"))
	}
}
