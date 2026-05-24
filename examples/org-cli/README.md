# Organization CLI Example

This example demonstrates how organizations can build custom CLI tools that include multispec commands alongside their own commands. Templates and rubrics are **compiled into the binary** using `//go:embed`, allowing distribution as a single executable.

## Features

- **Single Binary Distribution**: Templates and rubrics embedded at compile time
- **Custom Templates**: Override default templates with org-specific versions
- **Custom Rubrics**: Define organization-specific evaluation criteria
- **Additional Commands**: Add org-specific commands alongside multispec

## Project Structure

```
examples/org-cli/
├── main.go                     # CLI entry point with embedded loaders
├── templates/
│   └── prd.md                  # Custom PRD template (compiled in)
├── rubrics/
│   └── prd.rubric.yaml         # Custom PRD rubric (compiled in)
└── README.md
```

## Building

```bash
go build -o org-spec ./examples/org-cli
```

The resulting `org-spec` binary contains all templates and rubrics - no external files required.

## Usage

```bash
# Standard multispec commands (using compiled-in org templates/rubrics)
org-spec init my-project        # Uses org PRD template with security section
org-spec lint
org-spec eval prd               # Uses org PRD rubric with security category

# Organization-specific commands
org-spec policy list
org-spec policy apply
```

## How It Works

### Embedding Templates

```go
//go:embed templates/*.md
var orgTemplates embed.FS

cfg.TemplateLoader = templates.NewChainLoader(
    templates.NewEmbedFSLoader(orgTemplates, "templates"),  // Org templates (compiled in)
    templates.EmbeddedLoader(),                              // Fallback to multispec defaults
)
```

Templates in `templates/` are embedded at compile time. When a user runs `org-spec init`, the custom `prd.md` is used instead of the multispec default.

### Embedding Rubrics

```go
//go:embed rubrics/*.rubric.yaml
var orgRubrics embed.FS

cfg.RubricLoader = rubrics.NewChainLoader(
    rubrics.NewEmbedFSLoader(orgRubrics, "rubrics"),  // Org rubrics (compiled in)
    rubrics.EmbeddedLoader(),                          // Fallback to multispec defaults
)
```

Rubrics in `rubrics/` are embedded at compile time. When a user runs `org-spec eval prd`, the custom rubric (with security requirements) is used.

## Customizing for Your Organization

### 1. Create your templates

```
your-cli/templates/
├── prd.md      # Override with org requirements
├── mrd.md      # Override with org requirements
└── security.md # Add custom spec types
```

### 2. Create your rubrics

```yaml
# your-cli/rubrics/prd.rubric.yaml
spec_type: prd
name: "Your Org PRD Rubric"
categories:
  - id: security
    name: "Security Requirements"
    weight: 3.0
    required: true
    # ... criteria
```

### 3. Build your CLI

```go
package main

import (
    "embed"
    "github.com/plexusone/multispec/pkg/cli"
    "github.com/plexusone/multispec/pkg/templates"
    "github.com/plexusone/multispec/pkg/rubrics"
)

//go:embed templates/*.md
var orgTemplates embed.FS

//go:embed rubrics/*.rubric.yaml
var orgRubrics embed.FS

func main() {
    root := &cobra.Command{Use: "your-spec"}

    cfg := cli.DefaultConfig()
    cfg.TemplateLoader = templates.NewChainLoader(
        templates.NewEmbedFSLoader(orgTemplates, "templates"),
        templates.EmbeddedLoader(),
    )
    cfg.RubricLoader = rubrics.NewChainLoader(
        rubrics.NewEmbedFSLoader(orgRubrics, "rubrics"),
        rubrics.EmbeddedLoader(),
    )

    cli.AddCommandsTo(root, cfg)
    root.Execute()
}
```

## Loader Options

| Loader | Source | Use Case |
|--------|--------|----------|
| `EmbeddedLoader()` | Multispec defaults | Fallback to built-in templates/rubrics |
| `NewEmbedFSLoader(fs, dir)` | `embed.FS` | Compile org files into binary |
| `NewFileLoader(dir)` | Filesystem | Runtime loading (dev/testing) |
| `NewChainLoader(...)` | Multiple | Try loaders in order |

## Selective Command Inclusion

If you only need specific commands:

```go
cmds := cli.Commands(cfg)

root.AddCommand(cmds.Init)
root.AddCommand(cmds.Lint)
root.AddCommand(cmds.Eval)
// Omit synthesize, reconcile, etc.
```
