# Author TRD

Guide the user through creating a Technical Requirements Document that defines the system architecture and implementation approach.

## Overview

A TRD (Technical Requirements Document) translates product requirements into a technical design. It defines the architecture, APIs, data models, and non-functional requirements that enable implementation. It answers "How will we build this?"

## Prerequisites

An MRD and PRD should exist to provide context on market needs and product requirements. The TRD should trace back to PRD requirements.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **Architecture Style**: Monolith, microservices, serverless, or hybrid?
2. **Tech Stack**: What languages, frameworks, and infrastructure?
3. **Data Requirements**: What data needs to be stored? What are the access patterns?
4. **Integration Points**: What external systems need to integrate?
5. **Scale Requirements**: What are the expected load and growth patterns?
6. **Security Constraints**: What security and compliance requirements exist?

### 2. Initialize Draft

Use `start_draft` to create the TRD draft:

```
start_draft(project="<project-name>", spec_type="trd")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **System Overview**: High-level architecture and components
- **Architecture Decisions**: Key decisions with rationale
- **Component Design**: Detailed design of each component
- **API Design**: Endpoints, request/response schemas, error handling
- **Data Design**: Data models, storage choices, migration strategy
- **Non-Functional Requirements**: Performance, security, reliability, observability
- **Dependencies**: External services and internal dependencies

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="trd", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="trd")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure architecture is clearly described with diagrams
2. Verify every PRD requirement maps to technical design
3. Check API design includes all endpoints with schemas
4. Add data models with relationships and storage rationale
5. Include specific NFR targets (latency, throughput, availability)
6. Confirm the design is implementable by an engineer
7. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="trd")
```

This promotes the draft to the final `technical/trd.md` location.

## Evaluation Criteria

The TRD is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Architecture Clarity | 20% | Clear architecture with diagrams and data flow |
| PRD Traceability | 15% | Every requirement maps to technical design |
| API Design | 15% | Complete spec with schemas and error handling |
| Data Design | 15% | Models with relationships and storage rationale |
| NFR Coverage | 20% | Specific targets for performance, security, etc. |
| Implementability | 15% | Design is actionable with clear implementation path |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Include architecture diagrams (ASCII art or references to diagrams)
- Every PRD requirement should trace to a technical component
- API design should include request/response examples
- Data models should include relationships and indexes
- NFRs need specific, measurable targets
- Consider failure modes and error handling
- Document assumptions and constraints

## Architecture Diagram Format

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Client    │───▶│   API GW    │───▶│   Service   │
└─────────────┘    └─────────────┘    └──────┬──────┘
                                             │
                                             ▼
                                      ┌─────────────┐
                                      │  Database   │
                                      └─────────────┘
```

## API Design Format

```
### POST /api/v1/resource

Create a new resource.

**Request:**
```json
{
  "name": "string",
  "type": "string"
}
```

**Response (201):**
```json
{
  "id": "string",
  "name": "string",
  "created_at": "timestamp"
}
```

**Errors:**
- 400: Invalid request body
- 409: Resource already exists
```

## NFR Format

| Metric | Target | Measurement |
|--------|--------|-------------|
| P95 Latency | < 200ms | DataDog APM |
| Availability | 99.9% | Uptime monitoring |
| Throughput | 1000 RPS | Load testing |

## Next Steps

After TRD completion:

1. **IRD**: Infrastructure requirements for deployment
2. **Implementation**: Begin development based on TRD
