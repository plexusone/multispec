# MCP Tools

The MultiSpec MCP server provides the following tools.

## Implemented Tools

### list_projects

List all MultiSpec projects in the current repository.

**Arguments:** None

**Returns:**

```json
{
  "projects": ["user-onboarding", "payment-flow"],
  "specs_dir": "docs/specs",
  "count": 2
}
```

### get_project_status

Get the status and readiness of a specific project.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |

**Returns:**

```json
{
  "project": "user-onboarding",
  "path": "docs/specs/user-onboarding",
  "generated_at": "2024-01-15T10:30:00Z",
  "readiness": {
    "ready": false,
    "summary": "Not ready: 2 blockers",
    "gates": [
      {"name": "Required specs present", "passed": true, "message": "..."},
      {"name": "Evaluations passing", "passed": true, "message": "..."},
      {"name": "Approvals obtained", "passed": false, "message": "..."},
      {"name": "Execution spec generated", "passed": false, "message": "..."}
    ]
  },
  "specs": [...],
  "summary": {
    "total_specs": 10,
    "present_specs": 3,
    "evaluated_specs": 0,
    "approved_specs": 0
  }
}
```

## Stub Tools

The following tools are defined but not yet implemented:

### get_spec

Get the content of a specification.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `spec_type` | string | Yes | Spec type (mrd, prd, uxd, etc.) |

### get_eval

Get evaluation results for a specification.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `spec_type` | string | Yes | Spec type |

### run_eval

Run evaluation on a specification.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `spec_type` | string | Yes | Spec type |

### synthesize

Generate a spec from source documents.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `spec_type` | string | Yes | Spec type to generate |

### reconcile

Generate unified execution spec from approved specs.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |

### approve

Approve a specification.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `spec_type` | string | Yes | Spec type |
| `approver` | string | No | Approver identifier |

### export

Export specs to a target execution system.

**Arguments:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `project` | string | Yes | Project name |
| `target` | string | Yes | Target system (speckit, gsd, gastown, gascity) |

## Error Handling

All tools return errors in a consistent format:

```json
{
  "error": "error message here"
}
```

The response will have `isError: true` set in the MCP CallToolResult.
