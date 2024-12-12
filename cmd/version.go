package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// use init function to add our command to the
// rootCmd and keep everything in a single file.
func init() { rootCmd.AddCommand(versionCmd) }

var (
	versionCfg = &struct {
		Version string `config:"VERSION" validate:"required,semver"`
	}{}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the current version of the project",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			// use PreRunE to read in the config and run
			// validation before attempting to execute the command.
			vpr := viper.New()
			vpr.SetConfigFile(".env")
			if err := vpr.ReadInConfig(); err != nil {
				return errors.Unwrap(err)
			}
			// unmarshal viper values into our config struct
			err := vpr.Unmarshal(versionCfg, unmarshalOpts)
			if err != nil {
				return errors.Unwrap(err)
			}
			// validate required fields are set
			// and values are formatted correctly
			return validator.New(validationOpts).
				StructCtx(cmd.Context(), versionCfg)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			str := strings.Join([]string{"Version", versionCfg.Version}, " | ")
			log.Println(str)
			return nil
		},
	}
)
