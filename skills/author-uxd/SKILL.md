# Author UXD

Guide the user through creating a User Experience Design (UXD) document.

## Overview

A UXD defines how users will interact with the product. It covers personas, user journeys, information architecture, wireframes, and interaction patterns. It answers "How will users experience this?"

## Prerequisites

A PRD should exist for context on user stories and requirements.

## Workflow

### 1. Discovery Questions

Before starting the draft, gather essential information:

1. **User Personas**: Who are the primary and secondary users?
2. **Key Tasks**: What are the most important user tasks?
3. **Entry Points**: How do users discover this feature?
4. **Design Constraints**: Any platform, brand, or accessibility requirements?
5. **Existing Patterns**: What design system or patterns to follow?

### 2. Initialize Draft

Use `start_draft` to create the UXD draft:

```
start_draft(project="<project-name>", spec_type="uxd")
```

This returns the template and instructions.

### 3. Collaborative Authoring

Work with the user to fill in each section:

- **Design Goals**: UX objectives and guiding principles
- **User Personas**: Rich profiles with goals and pain points
- **User Journeys**: End-to-end flows with success criteria
- **Information Architecture**: Site map and navigation structure
- **Wireframes**: Layout and element descriptions
- **Interaction Design**: Patterns, transitions, micro-interactions
- **Accessibility**: WCAG compliance and considerations

### 4. Save Progress

After each major section, save the draft:

```
update_draft(project="<project-name>", spec_type="uxd", content="<full-content>")
```

### 5. Evaluate Quality

When the draft feels complete, run evaluation:

```
eval_draft(project="<project-name>", spec_type="uxd")
```

Review the findings and score. Address any critical or high severity issues.

### 6. Iterate

Based on evaluation feedback:

1. Ensure personas have specific goals and pain points
2. Verify journeys have clear entry/exit points and success criteria
3. Check wireframes cover key screens with element descriptions
4. Add accessibility considerations (color contrast, keyboard nav, screen readers)
5. Define error states and edge cases
6. Re-evaluate until passing

### 7. Finalize

Once evaluation passes:

```
finalize_draft(project="<project-name>", spec_type="uxd")
```

This promotes the draft to the final `source/uxd.md` location.

## Evaluation Criteria

The UXD is evaluated on:

| Category | Weight | What We Look For |
|----------|--------|------------------|
| Design Goals | 10% | Clear, measurable UX objectives |
| User Personas | 15% | Rich profiles with behaviors |
| User Journeys | 15% | Complete flows with success criteria |
| Information Architecture | 15% | Clear navigation and hierarchy |
| Wireframes | 15% | Key screens with element descriptions |
| Interaction Design | 10% | Patterns and transitions defined |
| Accessibility | 15% | WCAG compliance, inclusive design |
| Error States | 5% | Error handling and edge cases |

A passing score is 7.0+ with no critical or high severity findings.

## Tips

- Personas should feel like real people, not demographics
- User journeys should map to PRD user stories
- Use ASCII wireframes in markdown - they're sufficient for specs
- Accessibility is not optional - design for it from the start
- Include both happy path and error states

## Wireframe Format (ASCII)

```
┌──────────────────────────────────────┐
│ Header / Navigation                  │
├──────────────────────────────────────┤
│                                      │
│ Main Content Area                    │
│ - Element A: Description             │
│ - Element B: Description             │
│                                      │
├──────────────────────────────────────┤
│ Footer / Actions                     │
└──────────────────────────────────────┘
```

## Accessibility Checklist

- [ ] Color contrast meets WCAG AA (4.5:1 for text)
- [ ] All interactive elements keyboard accessible
- [ ] Focus indicators visible
- [ ] Alt text for images
- [ ] Form labels properly associated
- [ ] Skip links available
- [ ] Screen reader tested

## Next Steps

After UXD completion:

1. **TRD**: Technical design (synthesized by LLM)
2. **Reconciliation**: Generate unified spec.md
