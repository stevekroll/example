package cmd

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevekroll/example/internal/webserver"
)

// use init function to add our command to the
// rootCmd and keep everything in a single file.
func init() { rootCmd.AddCommand(webserverCmd) }

var (
	webserverCfg = &struct {
		App      string     `config:"APP"       validate:"required,alpha"`
		Host     string     `config:"HOST"      validate:"required,http_url"`
		Port     string     `config:"PORT"      validate:"required,number"`
		LogLevel slog.Level `config:"LOG_LEVEL" validate:"required,numeric"`
		Timeout  int        `config:"TIMEOUT"   validate:"required,numeric"`
	}{}
	webserverCmd = &cobra.Command{
		Use:   "webserver",
		Short: "Run the webserver",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			// use PreRunE to read in the config and run
			// validation before attempting to execute the command.
			vpr := viper.New()
			vpr.SetConfigFile(".env")
			if err := vpr.ReadInConfig(); err != nil {
				return errors.Unwrap(err)
			}
			// unmarshal viper values into our config struct
			err := vpr.Unmarshal(webserverCfg, unmarshalOpts)
			if err != nil {
				return errors.Unwrap(err)
			}
			// validate required fields are set
			// and values are formatted correctly
			return validator.New(validationOpts).
				StructCtx(cmd.Context(), webserverCfg)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			addr := strings.Join([]string{webserverCfg.Host, webserverCfg.Port}, ":")
			srv := http.Server{
				Addr:              netAddr(addr),
				Handler:           webserver.NewMux(),
				ReadHeaderTimeout: time.Duration(webserverCfg.Timeout),
			}
			return srv.ListenAndServe()
		},
	}
)
