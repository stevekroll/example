package cmd

import "github.com/spf13/cobra"

func init() { rootCmd.AddCommand(webserverCmd) }

var webserverCmd = &cobra.Command{
	Use:        "webserver",
	Aliases:    []string{},
	SuggestFor: []string{},
	Short:      "",
	GroupID:    "",
	Long:       "",
	Example:    "",
	ValidArgs:  []string{},
	// ValidArgsFunction:          func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {},
	// Args:                       func(cmd *cobra.Command, args []string) error {},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            map[string]string{},
	Version:                "",
	// PersistentPreRun:           func(cmd *cobra.Command, args []string) {},
	// PersistentPreRunE:          func(cmd *cobra.Command, args []string) error {},
	// PreRun:                     func(cmd *cobra.Command, args []string) {},
	// PreRunE:                    func(cmd *cobra.Command, args []string) error {},
	// Run:                        func(cmd *cobra.Command, args []string) {},
	// RunE:                       func(cmd *cobra.Command, args []string) error {},
	// PostRun:                    func(cmd *cobra.Command, args []string) {},
	// PostRunE:                   func(cmd *cobra.Command, args []string) error {},
	// PersistentPostRun:          func(cmd *cobra.Command, args []string) {},
	// PersistentPostRunE:         func(cmd *cobra.Command, args []string) error {},
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}
