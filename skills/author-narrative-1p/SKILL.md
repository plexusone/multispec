# Author Narrative 1-Pager

Guide the user through creating an Executive 1-Pager narrative.

## Overview

A 1-pager is a concise executive summary designed to communicate the essence of a proposal in a single page. It's used for quick alignment, elevator pitches, and stakeholder buy-in.

## Prerequisites

Source specs (MRD, PRD, UXD) should be complete or in progress for context.

## Workflow

### 1. Discovery Questions

Before starting, gather the essential information:

1. **Opportunity**: What is the core opportunity or problem in one sentence?
2. **Solution**: What are we proposing to do about it?
3. **Customer**: Who specifically benefits from this?
4. **Ask**: What decision or resources do you need?

### 2. Initialize Draft

Use `start_draft` to create the 1-pager draft:

```
start_draft(project="<project-name>", spec_type="narrative-1p")
```

### 3. Collaborative Authoring

Work with the user to craft each section concisely:

- **The Opportunity**: 2-3 sentences on the market problem or opportunity
- **Our Solution**: 2-3 sentences on the value proposition
- **Target Customer**: Specific segment definition
- **Key Benefits**: 3-5 concrete, outcome-focused benefits
- **Success Metrics**: 2-3 measurable targets
- **What We Need**: Clear ask and next steps

### 4. Save Progress

Save the draft after each revision:

```
update_draft(project="<project-name>", spec_type="narrative-1p", content="<full-content>")
```

### 5. Evaluate Quality

When complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="narrative-1p")
```

### 6. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="narrative-1p")
```

## Evaluation Criteria

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Opportunity Clarity | 25% | Compelling, timely opportunity |
| Solution Value | 25% | Clear value proposition |
| Target Specificity | 15% | Well-defined customer segment |
| Concrete Benefits | 20% | Outcome-focused, not features |
| Clear Ask | 15% | Specific request and next steps |

## Tips

- Every sentence must earn its place - cut ruthlessly
- Lead with impact, not background
- Benefits should be outcomes customers care about
- The ask should be specific and actionable
- Read it aloud - it should flow naturally

## Next Steps

After 1-pager completion:

1. Share for quick alignment
2. Create 6-pager for detailed analysis
3. Use for stakeholder presentations
