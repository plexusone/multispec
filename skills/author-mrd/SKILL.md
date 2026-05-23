# Author MRD

Guide the user through creating a Market Requirements Document (MRD).

## Overview

An MRD captures the market opportunity, target audience, competitive landscape, and high-level requirements from a market perspective. It answers "What problem are we solving and why?"

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **Market Problem**: What pain point or opportunity exists in the market?
2. **Target Market**: Who are the primary customers? What segments?
3. **Competition**: Who else solves this problem? How will you differentiate?
4. **Business Goals**: What revenue/growth targets does this support?
5. **Timeline**: When does this need to ship?

### 2. Initialize Draft

Use `start_draft` to create the MRD draft:

```
start_draft(project="<project-name>", spec_type="mrd")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **Problem Statement**: Quantify the pain point and cost of inaction
- **Target Market**: Define TAM/SAM/SOM and segment profiles
- **Competitive Landscape**: Analyze direct/indirect competitors
- **Market Requirements**: Prioritize using MoSCoW (Must/Should/Could/Won't)
- **Business Goals**: Set measurable success metrics
- **Risks**: Identify constraints, assumptions, and mitigations

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="mrd", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="mrd")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Identify gaps in problem clarity, market sizing, or requirements
2. Ask clarifying questions to fill gaps
3. Update the draft
4. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="mrd")
```

This promotes the draft to the final `source/mrd.md` location.

## Evaluation Criteria

The MRD is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Problem Statement | 20% | Clear, quantified, timely |
| Target Market | 20% | Defined segments with sizing |
| Competitive Analysis | 15% | Differentiation strategy |
| Requirements Clarity | 20% | Prioritized, rationalized |
| Business Alignment | 15% | Metrics and strategic fit |
| Risk Assessment | 10% | Identified with mitigations |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Start with the problem statement - it anchors everything else
- Use specific numbers for market sizing, not vague estimates
- Competitive analysis should highlight gaps, not just list competitors
- Requirements should trace to customer needs or business goals
- Include both risks and assumptions

## Next Steps

After MRD completion:

1. **PRD**: Define product requirements (what to build)
2. **UXD**: Define user experience (how users interact)
