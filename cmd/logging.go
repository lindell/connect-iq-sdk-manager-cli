package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func configureLogging(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()

	flags.StringP("log-level", "L", "info", "The level of logging that should be made. Available values: trace, debug, info, error.")
	_ = cmd.RegisterFlagCompletionFunc("log-level", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"trace", "debug", "info", "error"}, cobra.ShellCompDirectiveDefault
	})

	flags.StringP("log-format", "", "text", `The formatting of the logs. Available values: text, json, json-pretty.`)
	_ = cmd.RegisterFlagCompletionFunc("log-format", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"text", "json", "json-pretty"}, cobra.ShellCompDirectiveDefault
	})

	flags.StringP("log-file", "", "-", `The file where all logs should be printed to. "-" means stdout.`)
}

func logFlagInit(cmd *cobra.Command) error {
	// Parse and set log level
	strLevel, _ := cmd.Flags().GetString("log-level")
	logLevel, err := log.ParseLevel(strLevel)
	if err != nil {
		return fmt.Errorf("invalid log-level: %s", strLevel)
	}
	log.SetLevel(logLevel)

	// Parse and set the log format
	strFormat, _ := cmd.Flags().GetString("log-format")

	var formatter log.Formatter
	switch strFormat {
	case "text":
		formatter = &log.TextFormatter{}
	case "json":
		formatter = &log.JSONFormatter{}
	case "json-pretty":
		formatter = &log.JSONFormatter{
			PrettyPrint: true,
		}
	default:
		return fmt.Errorf(`unknown log-format "%s"`, strFormat)
	}

	log.SetFormatter(formatter)

	// Set the output (file)
	strFile, _ := cmd.Flags().GetString("log-file")
	if strFile == "" {
		log.SetOutput(nopWriter{})
	} else if strFile != "-" {
		file, err := os.Create(strFile)
		if err != nil {
			return errors.Wrapf(err, "could not open log-file %s", strFile)
		}
		log.SetOutput(file)
	}

	return nil
}

// nopWriter is a writer that does nothing
type nopWriter struct{}

func (nw nopWriter) Write(bb []byte) (int, error) {
	return len(bb), nil
}
