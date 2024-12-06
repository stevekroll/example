package cmd

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// use init function to add this command to our
// rootCmd this keeps initialization to a single file.
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
			v := viper.New()
			v.SetConfigFile(".env")
			if err := v.ReadInConfig(); err != nil {
				return errors.Unwrap(err)
			}
			// unmarshal viper values into our config struct
			err := v.Unmarshal(versionCfg, unmarshalOpts)
			if err != nil {
				return errors.Unwrap(err)
			}
			// validate required fields are set
			// and values are formatted correctly
			return validator.New(validationOpts).
				StructCtx(cmd.Context(), versionCfg)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := fmt.Println("Version |", versionCfg.Version)
			return err
		},
	}
)
