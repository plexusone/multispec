# Author Press Release

Guide the user through creating a Press Release document following the Amazon Working Backwards methodology.

## Overview

A Press Release is written as if the product has already launched successfully. It forces clarity on the customer benefit and value proposition by describing the announcement from the customer's perspective. It answers "What will we announce and why should customers care?"

## Prerequisites

An MRD and PRD should exist to provide context on market opportunity and product requirements.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **Customer Benefit**: What is the single most important benefit for customers?
2. **Problem Statement**: What specific pain point does this solve?
3. **Target Audience**: Who is the primary customer for this announcement?
4. **Differentiator**: What makes this solution unique?
5. **Availability**: When and how will customers access this?

### 2. Initialize Draft

Use `start_draft` to create the Press Release draft:

```
start_draft(project="<project-name>", spec_type="press")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **Headline**: Clear, benefit-focused headline (not feature-focused)
- **Subheadline**: Expand on the headline with context
- **Opening Paragraph**: City, date, and announcement summary
- **Customer Problem**: Describe the problem in customer terms
- **Solution**: How the product solves the problem
- **Customer Quote**: Authentic voice describing the benefit
- **How It Works**: Simple explanation of the solution
- **Executive Quote**: Vision and commitment from leadership
- **Call to Action**: Clear next steps for interested customers

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="press", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="press")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure headline focuses on customer benefit, not features
2. Verify customer problem is described in customer language
3. Check that solution clearly addresses the stated problem
4. Add authentic customer quotes with specific examples
5. Include clear availability and pricing information
6. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="press")
```

This promotes the draft to the final `gtm/press.md` location.

## Evaluation Criteria

The Press Release is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Headline Impact | 20% | Compelling headline with clear customer benefit |
| Customer Problem | 20% | Vivid problem description with authentic voice |
| Solution Clarity | 20% | Crystal clear solution with concrete benefits |
| Customer Validation | 15% | Authentic quotes with specific success stories |
| Call to Action | 15% | Clear availability, pricing, and next steps |
| Readability | 10% | Jargon-free prose accessible to general audience |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Write as if the product has already launched and succeeded
- Focus on customer benefit in the headline, not features
- Use customer language, not internal or technical jargon
- Customer quotes should sound like real people, not marketing
- Be specific about benefits - quantify where possible
- The press release should be understandable by anyone
- Imagine this being read by a journalist or customer

## Press Release Structure

```
HEADLINE: [Benefit-focused, not feature-focused]
SUBHEADLINE: [Expand on headline with context]

CITY, DATE - [Opening paragraph with announcement]

[Customer problem paragraph]

[Solution paragraph]

"[Customer quote describing benefit]" - [Customer Name, Title]

[How it works paragraph]

"[Executive quote with vision]" - [Executive Name, Title]

[Availability and call to action]

About [Company]
[Boilerplate]
```

## Next Steps

After Press Release completion:

1. **FAQ**: Answer anticipated customer questions
2. **Narrative 1-Pager**: Executive summary for stakeholders
