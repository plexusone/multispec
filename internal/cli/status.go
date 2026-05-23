package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/plexusone/multispec/pkg/config"
	"github.com/plexusone/multispec/pkg/status"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show project status and readiness",
	Long: `Show the status of all specs, evaluations, and approvals for a project.

Displays readiness gates and indicates whether the project is ready
for AI-assisted development.`,
	RunE: runStatus,
}

func init() {
	statusCmd.Flags().String("format", "text", "Output format: text, json, html, markdown")
	statusCmd.Flags().Bool("ci", false, "CI mode: exit non-zero if not ready")
}

func runStatus(cmd *cobra.Command, args []string) error {
	projectPath, err := resolveProjectPath(cmd)
	if err != nil {
		return err
	}

	project, err := config.Load(projectPath)
	if err != nil {
		return fmt.Errorf("failed to load project: %w", err)
	}

	report, err := status.Generate(project)
	if err != nil {
		return fmt.Errorf("failed to generate status: %w", err)
	}

	format, _ := cmd.Flags().GetString("format")

	switch format {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(report)
	case "html":
		return status.RenderHTML(os.Stdout, report)
	case "markdown":
		return status.RenderMarkdown(os.Stdout, report)
	default:
		return status.RenderText(os.Stdout, report)
	}
}

func resolveProjectPath(cmd *cobra.Command) (string, error) {
	projectFlag, _ := cmd.Flags().GetString("project")

	if projectFlag != "" {
		// Check if it's an absolute path
		if _, err := os.Stat(projectFlag); err == nil {
			return projectFlag, nil
		}

		// Try as project name under specs directory
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		specsDir, err := config.FindSpecsDir(cwd)
		if err != nil {
			return "", fmt.Errorf("specs directory not found")
		}

		projectPath := config.ProjectPath(specsDir, projectFlag)
		if _, err := os.Stat(projectPath); err != nil {
			return "", fmt.Errorf("project %q not found", projectFlag)
		}

		return projectPath, nil
	}

	// Try to find project root from current directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	projectPath, err := config.FindProjectRoot(cwd)
	if err != nil {
		return "", fmt.Errorf("no project found (use --project flag or run from project directory)")
	}

	return projectPath, nil
}
