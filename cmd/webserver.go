package cmd

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// use init function to add this command to our
// rootCmd this keeps initialization to a single file.
func init() { rootCmd.AddCommand(webserverCmd) }

var (
	webserverCfg = &struct {
		App      string     `config:"APP" validate:"required,alpha"`
		Host     string     `config:"HOST" validate:"required,http_url"`
		Port     string     `config:"PORT" validate:"required,number"`
		LogLevel slog.Level `config:"LOG_LEVEL" validate:"required,numeric"`
	}{}
	webserverCmd = &cobra.Command{
		Use:   "webserver",
		Short: "Run the webserver",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			// use PreRunE to read in the config and run
			// validation before attempting to execute the command.
			v := viper.New()
			v.SetConfigFile(".env")
			if err := v.ReadInConfig(); err != nil {
				return errors.Unwrap(err)
			}
			// unmarshal viper values into our config struct
			err := v.Unmarshal(webserverCfg, unmarshalOpts)
			if err != nil {
				return errors.Unwrap(err)
			}
			// validate required fields are set
			// and values are formatted correctly
			return validator.New(validationOpts).
				StructCtx(cmd.Context(), webserverCfg)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Println(webserverCfg)
			return nil
		},
	}
)
