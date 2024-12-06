package cmd

import (
	"context"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
)

var (
	// root command needs to be executed in your main function
	// and allows us to optionally provide additional commands.
	rootCmd        = &cobra.Command{}
	ExecuteContext = func(ctx context.Context) error {
		return rootCmd.ExecuteContext(ctx)
	}
)

// sets the unmarshal struct tag name to 'config'.
var unmarshalOpts = func(dc *mapstructure.DecoderConfig) {
	dc.TagName = "config"
}

// enables required tag on non-pointer structs to be applied instead of ignored.
var validationOpts = validator.WithRequiredStructEnabled()

// builds the network address from the base url.
var netAddr = func(baseURL string) string {
	u, _ := url.Parse(baseURL)
	return u.Host
}
