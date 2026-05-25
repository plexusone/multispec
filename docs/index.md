# MultiSpec

Multi-domain specification orchestration for humans and AI agents.

## What is MultiSpec?

MultiSpec bridges the gap between organizational intent (MRD, PRD, UXD) and executable specifications for AI coding agents. It provides a structured workflow for:

- **Domain-specific authoring** - Separate specs for PM, UX, Engineering
- **GTM synthesis** - LLM-generated press releases, FAQs, narratives (Working Backwards)
- **Technical synthesis** - LLM-generated TRD, IRD from source specs
- **Structured evaluation** - Per-domain LLM judges with customizable rubrics
- **Reconciliation** - Conflict detection and tradeoff resolution
- **Target adapters** - Export to SpecKit, GSD, GasTown, GasCity, OpenSpec

## Working Backwards Flow

MultiSpec implements Amazon's Working Backwards methodology:

```
1. MARKET PROBLEM (human-authored)
   mrd.md
       ↓
2. WORKING BACKWARDS (synthesized, editable)
   press.md  →  faq.md  →  prd.md
   (vision)     (scope)    (requirements)
       ↓
3. STAKEHOLDER REVIEW (synthesized, editable)
   narrative-1p.md / narrative-6p.md
       ↓
4. USER EXPERIENCE (human-authored)
   uxd.md
       ↓
5. TECHNICAL SPECS (synthesized, editable)
   trd.md  →  ird.md
       ↓
6. RECONCILIATION
   All approved specs → spec.md
       ↓
7. AI EXECUTION
   spec.md → SpecKit | GSD | GasTown | GasCity
       ↓
8. POST-SHIP ALIGNMENT
   spec.md + reality → current-truth.md
```

All synthesized documents are committed to git and can be reviewed, edited, and refined by humans or collaboratively with AI assistants.

See the [Working Backwards Guide](concepts/working-backwards.md) for the full methodology.

## Quick Start

```bash
# Install
go install github.com/plexusone/multispec/cmd/multispec@v0.4.0

# Initialize a new project
multispec init user-onboarding

# Initialize with a profile (startup, growth, enterprise)
multispec init my-feature --profile startup

# Validate project structure
multispec lint

# Check project status
multispec status
```

## Key Features

### Readiness Gates

MultiSpec tracks project readiness through configurable gates:

| Gate | Description |
|------|-------------|
| Required specs present | All required source specs exist |
| Evaluations passing | No blocking evaluation findings |
| Approvals obtained | All required specs have approvals |
| Execution spec generated | `spec.md` has been created |

### Multiple Output Formats

Status reports can be generated in multiple formats for different use cases:

- **Text** - Terminal output with icons
- **JSON** - Programmatic access
- **HTML** - Browser-viewable reports with traffic light indicators
- **Markdown** - For embedding in documentation

### MCP Integration

MultiSpec includes an MCP (Model Context Protocol) server for integration with AI coding assistants like Claude Code and Kiro CLI.

## Configuration Profiles

Profiles define which specs are required for different product lifecycle stages:

| Profile | Required Specs | Use Case |
|---------|---------------|----------|
| `0-1` | hypothesis | Idea validation phase |
| `startup` | prd | Pre-PMF startups |
| `growth` | prd, uxd, faq | 1-N scaling phase |
| `enterprise` | mrd, prd, uxd, trd, press, faq, spec | Post-PMF enterprises |

```bash
# List available profiles
multispec profiles list

# Initialize with a profile
multispec init my-project --profile startup

# Export a profile for customization
multispec profiles export enterprise ./my-profile

# Use a custom profile
multispec init my-project --profile-dir ./my-profile
```

Organizations can create custom profiles with their own templates and rubrics. See the [Custom Profiles Guide](guides/custom-profiles.md) for details.

## Project Status

See the [ROADMAP](specs/ROADMAP.md) for detailed implementation status and [Release Notes](releases/v0.4.0.md) for the latest release.

**Current Version:** v0.4.0

| Component | Status |
|-----------|--------|
| CLI (init, lint, status, eval, synthesize, reconcile) | Complete |
| MCP Server (draft workflow, eval) | Complete |
| Evaluation Engine | Complete |
| GTM & Technical Synthesis | Complete |
| Reconciliation | Complete |
| Export (SpecKit) | Complete |
| Graphize Integration | Complete |
| Profiles & Composability | Complete |
