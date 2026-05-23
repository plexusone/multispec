// Package cli implements the multispec command-line interface.
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time.
	Version = "dev"

	// Commit is set at build time.
	Commit = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "multispec",
	Short: "Multi-domain specification orchestration for humans and AI agents",
	Long: `MultiSpec bridges the gap between organizational intent (MRD, PRD, UXD)
and executable specifications for AI coding agents.

It provides:
  - Domain-specific authoring (source specs)
  - GTM synthesis (press releases, FAQs, narratives)
  - Technical synthesis (TRD, IRD)
  - Structured evaluation with LLM judges
  - Reconciliation into unified execution specs
  - Export to SpecKit, GSD, GasTown, GasCity, OpenSpec`,
	Version: fmt.Sprintf("%s (commit: %s)", Version, Commit),
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP("project", "p", "", "Project name or path")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")

	// Add subcommands
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(lintCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(evalCmd)
	rootCmd.AddCommand(synthesizeCmd)
	rootCmd.AddCommand(reconcileCmd)
	rootCmd.AddCommand(approveCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(targetsCmd)
	rootCmd.AddCommand(graphCmd)
	rootCmd.AddCommand(serveCmd)
}
