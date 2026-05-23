# Author PRD

Guide the user through creating a Product Requirements Document (PRD).

## Overview

A PRD defines what to build from the product perspective. It translates market requirements into user stories, functional requirements, and acceptance criteria. It answers "What are we building and for whom?"

## Prerequisites

An MRD should exist or be in progress for context on market needs.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **User Problem**: What specific user problem does this solve?
2. **Goals**: What are 2-3 measurable goals for this feature?
3. **Non-Goals**: What is explicitly out of scope?
4. **User Stories**: Who are the users and what do they need?
5. **Success Metrics**: How will we measure success?

### 2. Initialize Draft

Use `start_draft` to create the PRD draft:

```
start_draft(project="<project-name>", spec_type="prd")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **Problem Statement**: User-centric problem definition
- **Goals/Non-Goals**: Explicit scope boundaries
- **User Stories**: As a [user], I want [goal] so that [reason]
- **Functional Requirements**: Specific, testable capabilities
- **Non-Functional Requirements**: Performance, security, accessibility
- **Success Metrics**: Measurable outcomes with baselines and targets

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="prd", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="prd")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure all user stories have acceptance criteria
2. Check requirements are testable (avoid "fast", "easy", "intuitive")
3. Verify NFRs cover performance, security, and accessibility
4. Add success metrics with measurement methods
5. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="prd")
```

This promotes the draft to the final `source/prd.md` location.

## Evaluation Criteria

The PRD is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Problem Definition | 15% | Specific, measurable user impact |
| Goals and Scope | 15% | SMART goals, explicit non-goals |
| User Stories | 20% | Complete format with acceptance criteria |
| Functional Requirements | 20% | Testable, traceable to stories |
| Non-Functional Requirements | 15% | Coverage of perf/security/a11y |
| Success Metrics | 10% | Baselines, targets, measurement |
| Dependencies | 5% | Identified with owners |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- User stories follow: "As a [type], I want [goal] so that [reason]"
- Every requirement should be testable - avoid subjective terms
- Include acceptance criteria for each user story
- NFRs need specific numbers (e.g., "< 200ms p95 latency")
- Link requirements back to user stories

## User Story Format

```
As a [type of user],
I want [goal]
so that [reason].

Acceptance Criteria:
- Given [context], when [action], then [result]
- ...
```

## Next Steps

After PRD completion:

1. **UXD**: Define user experience and interaction design
2. **TRD**: Technical design (synthesized by LLM)
