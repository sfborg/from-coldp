/*
Copyright © 2024 Dmitry Mozzherin <dmozzherin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/sfborg/from-coldp/internal/io/sfgarcio"
	"github.com/sfborg/from-coldp/internal/io/sysio"
	fcoldp "github.com/sfborg/from-coldp/pkg"
	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/sfborg/sflib/io/dbio"
	"github.com/sfborg/sflib/io/schemaio"
	"github.com/spf13/cobra"
)

var opts []config.Option

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "from-coldp",
	Short: "Converts CoLDP archive file to SFGA archive",
	Long:  `Converts CoLDP archive file to SFGA archive`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		versionFlag(cmd)
		flags := []flagFunc{
			debugFlag, cacheDirFlag, jobsNumFlag, binFlag, zipFlag, fieldsNumFlag,
		}
		for _, v := range flags {
			v(cmd)
		}

		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}

		slog.Info("Converting CoLDP to SFGA")
		coldpPath := args[0]
		outputPath := args[1]

		ext := filepath.Ext(outputPath)
		if ext == ".sqlite" {
			opts = append(opts, config.OptWithBinOutput(true))
		}

		cfg := config.New(opts...)
		err = sysio.New(cfg).ResetCache()
		if err != nil {
			slog.Error("Cannot initialize file system", "error", err)
			os.Exit(1)
		}

		sfgaSchema := schemaio.New(cfg.GitRepo, cfg.TempRepoDir)
		sfgaDB := dbio.New(cfg.CacheSfgaDir)

		sfarc := sfgarcio.New(cfg, sfgaSchema, sfgaDB)
		err = sfarc.Connect()
		if err != nil {
			slog.Error("Cannot initialize storage", "error", err)
			os.Exit(1)
		}

		fc := fcoldp.New(cfg, sfarc)
		var arc coldp.Archive

		slog.Info("Importing CoLDP data", "file", coldpPath)
		arc, err = fc.GetCoLDP(coldpPath)
		if err != nil {
			slog.Error("Cannot get CoLDP Archive", "error", err)
			os.Exit(1)
		}
		_ = arc

		slog.Info("Exporting data to SQLite")
		err = fc.ImportCoLDP(arc)
		if err != nil {
			slog.Error("Cannot export data", "error", err)
			os.Exit(1)
		}
		//
		// slog.Info("Making SFGArchive")
		// err = fd.ExportSFGA(outputPath)
		// if err != nil {
		// 	slog.Error("Cannot dump data", "error", err)
		// 	os.Exit(1)
		// }
		//
		// slog.Info("CoLDP data has been imported successfully")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("debug", "d", false, "set debug mode")
	rootCmd.Flags().StringP("cache-dir", "c", "", "cache directory for temporary files")
	rootCmd.Flags().StringP("wrong-fields-num", "w", "",
		`how to process rows with wrong fields number
     choices: 'stop', 'skip', 'process'
     default: 'stop'`)
	rootCmd.Flags().IntP("jobs-number", "j", 0, "number of concurrent jobs")
	rootCmd.Flags().BoolP("binary-output", "b", false, "return binary SQLite database")
	rootCmd.Flags().BoolP("zip-output", "z", false, "compress output with zip")
	rootCmd.Flags().BoolP("version", "V", false, "shows app's version")
}
