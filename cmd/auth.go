package cmd

import (
	"os"

	"github.com/bitrise-core/bitrise-plugins-io/configs"
	"github.com/bitrise-core/bitrise-plugins-io/services"
	"github.com/bitrise-io/go-utils/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	tokenFlag string
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate",
	Run: func(cmd *cobra.Command, args []string) {
		if err := auth(); err != nil {
			log.Errorf(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringVar(&tokenFlag, "token", "", "Authentication token")
}

func auth() error {
	if tokenFlag == "" {
		return errors.New("Failed to set authentication token, error: invalid number of arguments")
	}

	if err := configs.SetAPIToken(tokenFlag); err != nil {
		return errors.Errorf("Failed to set authentication token, error: %s", err)
	}

	err := services.ValidateAuthToken()
	if err != nil {
		return err
	}
	log.Successf("authenticated")
	return nil
}