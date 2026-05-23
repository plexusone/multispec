# MultiSpec

Multi-domain specification orchestration for humans and AI agents.

## Overview

MultiSpec bridges the gap between organizational intent (MRD, PRD, UXD) and executable specifications for AI coding agents. It provides:

- **Domain-specific authoring** - Separate specs for PM, UX, Engineering
- **GTM synthesis** - LLM-generated press releases, FAQs, narratives (Working Backwards)
- **Technical synthesis** - LLM-generated TRD, IRD from source specs
- **Structured evaluation** - Per-domain LLM judges with customizable rubrics
- **Reconciliation** - Conflict detection and tradeoff resolution
- **Target adapters** - Export to SpecKit, GSD, GasTown, GasCity, OpenSpec

## Installation

```bash
go install github.com/plexusone/multispec/cmd/multispec@latest
```

## Quick Start

```bash
# Initialize a new project
multispec init user-onboarding

# Validate project structure
multispec lint

# Check project status
multispec status
multispec status --format json
multispec status --format html > status.html
```

## Directory Structure

```
docs/specs/
├── CONSTITUTION.md                    # Repo-level governance
├── ROADMAP.md                         # Cross-project priorities
└── {project}/                         # kebab-case project name
    ├── source/                        # Human-authored specs
    │   ├── mrd.md
    │   ├── prd.md
    │   └── uxd.md
    ├── gtm/                           # LLM-generated GTM docs
    │   ├── press.md
    │   ├── faq.md
    │   └── narrative.md
    ├── technical/                     # LLM-generated technical docs
    │   ├── trd.md
    │   └── ird.md
    ├── eval/                          # All evaluations
    │   ├── mrd.eval.json
    │   ├── prd.eval.json
    │   └── ...
    ├── .graphize/                     # Requirement graph
    ├── spec.md                        # Reconciled execution spec
    ├── current-truth.md               # Post-ship state
    ├── status.html                    # Readiness report
    ├── index.md                       # MkDocs page
    └── multispec.yaml                 # Configuration
```

## Document Lifecycle

```
HUMAN-AUTHORED (Source)
  MRD.md → PRD.md → UXD.md
           ↓
LLM-GENERATED (GTM) ← Working Backwards
  press.md → faq.md → narrative.md
           ↓
LLM-GENERATED (Technical)
  trd.md → ird.md
           ↓
RECONCILIATION
  All approved specs → spec.md
           ↓
TARGET EXPORT
  spec.md → SpecKit | GSD | GasTown | GasCity | OpenSpec
           ↓
POST-SHIP ALIGNMENT
  spec.md + reality → current-truth.md
```

## CLI Commands

### Implemented

| Command | Description |
|---------|-------------|
| `init <project>` | Initialize a new project with standard directory structure |
| `lint [project]` | Validate directory structure and naming conventions |
| `status` | Show project status and readiness gates |
| `targets` | List available export targets |

### Planned

| Command | Description |
|---------|-------------|
| `eval [type]` | Evaluate specs using LLM judges |
| `synthesize <type>` | Generate specs from source docs |
| `reconcile` | Generate unified execution spec |
| `approve <type>` | Approve a spec for reconciliation |
| `export <target>` | Export to target execution system |
| `graph <cmd>` | Manage requirement graphs |
| `serve` | Start MCP server for AI integration |

## Status Command

The `status` command shows project readiness with multiple output formats:

```bash
# Terminal output with readiness gates
multispec status -p myproject

# JSON for programmatic use
multispec status -p myproject --format json

# HTML report with traffic light indicator
multispec status -p myproject --format html > status.html

# Markdown for embedding in docs
multispec status -p myproject --format markdown

# CI mode - exits non-zero if not ready
multispec status -p myproject --ci
```

### Readiness Gates

| Gate | Description |
|------|-------------|
| Required specs present | All required source specs (mrd, prd, uxd, trd) exist |
| Evaluations passing | No blocking evaluation findings |
| Approvals obtained | All required specs have approvals |
| Execution spec generated | `spec.md` has been created |

## MCP Server

MultiSpec includes an MCP (Model Context Protocol) server for integration with AI coding assistants like Claude Code and Kiro CLI.

```bash
# Run MCP server directly
multispec-mcp
```

### MCP Tools

| Tool | Status | Description |
|------|--------|-------------|
| `list_projects` | Implemented | List all multispec projects |
| `get_project_status` | Implemented | Get project readiness status |
| `get_spec` | Stub | Get specification content |
| `get_eval` | Stub | Get evaluation results |
| `run_eval` | Stub | Run evaluation on a spec |
| `synthesize` | Stub | Generate a spec |
| `reconcile` | Stub | Generate execution spec |
| `approve` | Stub | Approve a specification |
| `export` | Stub | Export to target system |

## Export Targets

| Target | Description |
|--------|-------------|
| `speckit` | GitHub Spec-Kit format |
| `gsd` | Get Shit Done (PLAN.md, STATE.md) |
| `gastown` | GasTown formulas and beads |
| `gascity` | GasCity city.toml configuration |
| `openspec` | OpenSpec portable format (future) |

## Dependencies

- [structured-evaluation](https://github.com/plexusone/structured-evaluation) - Rubric and evaluation types
- [graphize](https://github.com/plexusone/graphize) - Requirement graph extraction

## Development

```bash
# Build
make build

# Test
make test

# Lint
make lint

# Install locally
make install
```

## Project Status

See [ROADMAP.md](docs/specs/ROADMAP.md) for detailed implementation status.

**Current Version:** v0.1.0-dev

| Phase | Status |
|-------|--------|
| Phase 0: Foundation | In Progress |
| Phase 1: Directory Structure | In Progress |
| Phase 2: Evaluation Engine | Not Started |
| Phase 3: GTM & Technical Synthesis | Not Started |
| Phase 4: Reconciliation Engine | Not Started |
| Phase 5: Target Adapters | Not Started |
| Phase 6: Claude Code Integration | In Progress |
| Phase 7: Graphize Integration | Not Started |
| Phase 8: Advanced Features | Not Started |

## License

MIT
