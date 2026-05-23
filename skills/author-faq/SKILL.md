# Author FAQ

Guide the user through creating a Frequently Asked Questions document that anticipates and addresses customer concerns.

## Overview

The FAQ document anticipates questions customers will have after reading the Press Release. It covers practical concerns about pricing, availability, getting started, and addresses potential objections. It answers "What will customers want to know?"

## Prerequisites

A Press Release should exist to provide context on what is being announced and the core value proposition.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **Target Audience**: Who are the primary readers of this FAQ?
2. **Key Concerns**: What are the top 3 objections or concerns customers might have?
3. **Pricing Model**: How is the product priced? What factors affect cost?
4. **Getting Started**: What's the path from interest to adoption?
5. **Competition**: How does this compare to alternatives?

### 2. Initialize Draft

Use `start_draft` to create the FAQ draft:

```
start_draft(project="<project-name>", spec_type="faq")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **General Questions**: What is it? Who is it for? Why should I care?
- **Getting Started**: How do I start? What's required? How long does it take?
- **Pricing**: How much does it cost? What's included? Are there free tiers?
- **Technical Questions**: How does it work? What integrations exist?
- **Support**: How do I get help? What SLAs exist?
- **Comparison**: How does this compare to X? Why choose this over Y?

### 4. Save Progress

After each category of questions, save the draft:

```
update_draft(project="<project-name>", spec_type="faq", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="faq")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure questions are phrased as customers would ask them
2. Check that answers are direct and don't dodge the question
3. Verify pricing information is clear with examples
4. Add getting started steps with clear path to success
5. Address likely objections honestly and confidently
6. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="faq")
```

This promotes the draft to the final `gtm/faq.md` location.

## Evaluation Criteria

The FAQ is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Question Coverage | 25% | Comprehensive coverage of likely questions |
| Answer Clarity | 25% | Direct, clear answers with no ambiguity |
| Customer Language | 15% | Questions and answers in customer terms |
| Pricing Transparency | 15% | Clear pricing with examples and guidance |
| Getting Started | 10% | Clear step-by-step adoption path |
| Objection Handling | 10% | Honest, confident responses to concerns |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Write questions as customers would actually phrase them
- Answer questions directly - don't dodge or deflect
- Use customer language, not internal jargon
- Include specific examples for pricing scenarios
- Don't be defensive about objections - address them honestly
- Group questions logically by topic
- Include both obvious and edge-case questions

## Question Categories to Cover

1. **What & Why**: What is this? Why should I use it? What problem does it solve?
2. **Who**: Who is this for? Is it right for my use case?
3. **How**: How does it work? How do I get started?
4. **Pricing**: How much? What's included? Free tier?
5. **Support**: How do I get help? What's the SLA?
6. **Migration**: How do I switch from my current solution?
7. **Security**: Is my data safe? What compliance certifications?
8. **Comparison**: How does this compare to alternatives?

## FAQ Format

```
## [Category Name]

### Q: [Question as customer would ask]

A: [Direct, clear answer]

[Optional: Example, link, or additional context]
```

## Next Steps

After FAQ completion:

1. **Narrative 6-Pager**: Detailed business case for leadership review
2. **TRD**: Technical design for implementation
