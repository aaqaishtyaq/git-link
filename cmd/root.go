/*
Copyright © 2022 Aaqa Ishtyaq <aaqaishtyaq@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/aaqaishtyaq/git-link/pkg/log"
	"github.com/aaqaishtyaq/git-link/pkg/permalink"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	verbosity, gitRemote string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-link",
	Short: "Generate permalink of a code line/snippet for Github",
	Long: `Generate permalink of a code line/snippet for Github

Usage:

$ git link main.go:1..2

	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := setupLogs(cmd)
		permalink.Generate(args, gitRemote, log)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", zapcore.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().StringVarP(&gitRemote, "remote", "r", "origin", "Git remote for which URL is to be generated")
}

func setupLogs(cmd *cobra.Command) *zap.SugaredLogger {
	verbosity := cmd.Root().Flag("verbosity").Value.String()

	return log.SetUpLogs(os.Stdout, verbosity)
}
