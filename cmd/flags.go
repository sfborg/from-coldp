package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gnames/gnfmt"
	fcoldp "github.com/sfborg/from-coldp/pkg"
	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/spf13/cobra"
)

type flagFunc func(cmd *cobra.Command)

func debugFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("debug")
	if b {
		lopts := &slog.HandlerOptions{Level: slog.LevelDebug}
		handle := slog.NewJSONHandler(os.Stderr, lopts)
		logger := slog.New(handle)
		slog.SetDefault(logger)
	}
}

func cacheDirFlag(cmd *cobra.Command) {
	cacheDir, _ := cmd.Flags().GetString("cache-dir")
	if cacheDir != "" {
		opts = append(opts, config.OptCacheDir(cacheDir))
	}
}

func zipFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("zip-output")
	if b {
		opts = append(opts, config.OptWithZipOutput(b))
	}
}

func quotesFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("quotes-in-csv")
	if b {
		opts = append(opts, config.OptWithQuotes(b))
	}
}

func jobsNumFlag(cmd *cobra.Command) {
	jobs, _ := cmd.Flags().GetInt("jobs-number")
	if jobs > 0 {
		opts = append(opts, config.OptJobsNum(jobs))
	}
}

func versionFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("version")
	if b {
		version := fcoldp.GetVersion()
		fmt.Printf(
			"\nVersion: %s\nBuild:   %s\n\n",
			version.Version,
			version.Build,
		)
		os.Exit(0)
	}
}

func fieldsNumFlag(cmd *cobra.Command) {
	s, _ := cmd.Flags().GetString("wrong-fields-num")
	switch s {
	case "":
		return
	case "stop":
		opts = append(opts, config.OptBadRow(gnfmt.ErrorBadRow))
	case "ignore":
		opts = append(opts, config.OptBadRow(gnfmt.SkipBadRow))
	case "process":
		opts = append(opts, config.OptBadRow(gnfmt.ProcessBadRow))
	default:
		slog.Warn("Unknown setting for wrong-fields-num, keeping default",
			"setting", s)
		slog.Info("Supported values are: 'stop' (default), 'ignore', 'process'")
	}
}
