# Author Narrative 6-Pager

Guide the user through creating an AWS-style 6-pager narrative.

## Overview

The 6-pager is a comprehensive narrative document following Amazon's "Working Backwards" format. It's designed to be read in full before discussion (not presented as slides) and forces rigorous thinking through written prose.

## Prerequisites

Source specs (MRD, PRD, UXD) should be complete or in progress for context.

## Workflow

### 1. Discovery Questions

Before starting, deeply understand each section:

1. **Tenets**: What principles guide this initiative?
2. **Customer**: Who exactly is the customer? What's their specific problem?
3. **Solution**: What's the customer experience from start to finish?
4. **Timing**: Why is now the right time? What happens if we wait?
5. **Business Case**: What's the expected impact? ROI?
6. **Risks**: What could go wrong? What do we depend on?

### 2. Initialize Draft

Use `start_draft` to create the 6-pager draft:

```
start_draft(project="<project-name>", spec_type="narrative-6p")
```

### 3. Collaborative Authoring

Work through each section systematically:

#### Section 1: Introduction & Tenets
- Set context for the proposal
- Define 3-5 guiding principles
- Tenets should help make tradeoff decisions

#### Section 2: Customer Problem
- Define the specific customer (not "users" or "developers")
- Describe their problem in vivid detail
- Include quotes, anecdotes, or scenarios
- Explain how they solve it today and why that's insufficient

#### Section 3: Solution
- Start with the customer experience, not technology
- Walk through the end-to-end journey
- List key capabilities that enable this experience
- Be explicit about what's NOT in scope

#### Section 4: Why Now?
- Market timing factors (technology, trends, regulations)
- Competitive landscape and window of opportunity
- Internal readiness (skills, assets, strategic fit)

#### Section 5: Business Case
- Customer impact (quantified improvements)
- Business impact (revenue, adoption, retention)
- Investment required (team, timeline, resources)
- ROI analysis with assumptions

#### Section 6: Risks & Dependencies
- Key risks with likelihood, impact, and mitigation
- Dependencies on other teams or external parties
- Open questions that need resolution

### 4. Save Progress

Save frequently:

```
update_draft(project="<project-name>", spec_type="narrative-6p", content="<full-content>")
```

### 5. Evaluate Quality

When complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="narrative-6p")
```

### 6. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="narrative-6p")
```

## Evaluation Criteria

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Tenets Clarity | 10% | Actionable guiding principles |
| Customer Depth | 20% | Vivid problem with specificity |
| Solution Narrative | 20% | Customer-centric experience |
| Timing Justification | 15% | Compelling urgency |
| Business Case | 20% | Quantified impact with assumptions |
| Risks/Dependencies | 15% | Thorough with mitigations |

## Tips

- Write in complete sentences and paragraphs, not bullets
- The customer section should make readers feel the pain
- Avoid jargon and acronyms
- Include specific numbers, not "significant" or "many"
- Read it aloud - good writing sounds natural
- Have someone unfamiliar with the topic read it

## The 6-Pager Philosophy

From Amazon's practice:

1. **Narrative over slides**: Prose forces complete thoughts
2. **Silent reading**: Everyone reads first, then discusses
3. **Start with the customer**: Work backwards from their experience
4. **Data-driven**: Quantify claims with real numbers
5. **Long-term thinking**: Consider multi-year implications

## Next Steps

After 6-pager completion:

1. Schedule reading/discussion meeting
2. Distribute in advance (but expect in-meeting reading)
3. Take notes on questions for FAQ section
4. Iterate based on feedback
