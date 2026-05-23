# MCP Server

MultiSpec includes a Model Context Protocol (MCP) server for integration with AI coding assistants.

## Overview

The MCP server provides a standardized interface for AI assistants like Claude Code and Kiro CLI to interact with MultiSpec projects.

## Installation

The MCP server is built as a separate binary:

```bash
# Build
make build

# The binary is at:
bin/multispec-mcp
```

Or install directly:

```bash
go install github.com/plexusone/multispec/cmd/mcp-server@latest
```

## Configuration

### Claude Code

Add to your Claude Code MCP configuration:

```json
{
  "mcpServers": {
    "multispec": {
      "command": "multispec-mcp",
      "args": []
    }
  }
}
```

### Kiro CLI

Add to your Kiro steering configuration:

```yaml
mcp:
  servers:
    - name: multispec
      command: multispec-mcp
```

## Transport

The MCP server uses stdio transport by default, communicating via stdin/stdout with JSON-RPC 2.0 messages.

## Working Directory

The MCP server operates relative to the current working directory. It looks for:

1. `docs/specs/` directory containing MultiSpec projects
2. Individual project directories with `multispec.yaml`

## Available Tools

See [MCP Tools](tools.md) for the complete list of available tools.

| Tool | Status | Description |
|------|--------|-------------|
| `list_projects` | Implemented | List all projects |
| `get_project_status` | Implemented | Get project readiness |
| `get_spec` | Stub | Get spec content |
| `get_eval` | Stub | Get evaluation results |
| `run_eval` | Stub | Run evaluation |
| `synthesize` | Stub | Generate specs |
| `reconcile` | Stub | Generate execution spec |
| `approve` | Stub | Approve a spec |
| `export` | Stub | Export to target |

## Example Usage

Once configured, you can interact with MultiSpec through your AI assistant:

```
User: What multispec projects are available?
AI: [calls list_projects tool]
    Found 2 projects: user-onboarding, payment-flow

User: What's the status of user-onboarding?
AI: [calls get_project_status tool]
    Project user-onboarding is NOT READY.
    - Required specs present: PASS
    - Evaluations passing: PASS
    - Approvals obtained: FAIL (pending: prd, trd)
    - Execution spec generated: FAIL
```
