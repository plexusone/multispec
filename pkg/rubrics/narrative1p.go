package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewNarrative1PRubricSet())
}

// NewNarrative1PRubricSet creates the rubric set for 1-Pager Executive Narratives.
func NewNarrative1PRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeNarrative1P,
		Name:         "Narrative 1-Pager Evaluation",
		Description:  "Evaluates Executive 1-Pager narratives for clarity, impact, and completeness",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "opportunity_clarity",
				Name:        "Opportunity Clarity",
				Description: "How clearly the market opportunity or problem is articulated",
				Weight:      0.25,
				Rubric: newRubric("opportunity_clarity", "Opportunity statement quality").
					AddRangeAnchor(9, 10, "Excellent", "Compelling opportunity with clear market timing and urgency").
					AddRangeAnchor(7, 8.9, "Good", "Clear opportunity statement with reasonable justification").
					AddRangeAnchor(5, 6.9, "Adequate", "Opportunity stated but lacks compelling urgency").
					AddRangeAnchor(3, 4.9, "Needs Work", "Opportunity is vague or poorly justified").
					AddRangeAnchor(0, 2.9, "Poor", "No clear opportunity articulated"),
			},
			{
				ID:          "solution_value",
				Name:        "Solution Value Proposition",
				Description: "Clarity of the solution and its value proposition",
				Weight:      0.25,
				Rubric: newRubric("solution_value", "Solution clarity").
					AddRangeAnchor(9, 10, "Excellent", "Crystal clear value proposition with direct link to opportunity").
					AddRangeAnchor(7, 8.9, "Good", "Clear solution with understood value").
					AddRangeAnchor(5, 6.9, "Adequate", "Solution described but value unclear").
					AddRangeAnchor(3, 4.9, "Needs Work", "Solution is vague or disconnected from problem").
					AddRangeAnchor(0, 2.9, "Poor", "No clear solution articulated"),
			},
			{
				ID:          "target_specificity",
				Name:        "Target Customer Specificity",
				Description: "How specifically the target customer is defined",
				Weight:      0.15,
				Rubric: newRubric("target_specificity", "Target customer definition").
					AddRangeAnchor(9, 10, "Excellent", "Specific, well-defined customer segment with clear characteristics").
					AddRangeAnchor(7, 8.9, "Good", "Customer segment identified with key attributes").
					AddRangeAnchor(5, 6.9, "Adequate", "Customer mentioned but too broad").
					AddRangeAnchor(3, 4.9, "Needs Work", "Customer definition is vague").
					AddRangeAnchor(0, 2.9, "Poor", "No target customer defined"),
			},
			{
				ID:          "benefits_concrete",
				Name:        "Concrete Benefits",
				Description: "Whether benefits are outcome-focused and concrete",
				Weight:      0.20,
				Rubric: newRubric("benefits_concrete", "Benefit quality").
					AddRangeAnchor(9, 10, "Excellent", "Benefits are specific, measurable outcomes customers care about").
					AddRangeAnchor(7, 8.9, "Good", "Clear benefits focused on outcomes").
					AddRangeAnchor(5, 6.9, "Adequate", "Benefits listed but feature-focused").
					AddRangeAnchor(3, 4.9, "Needs Work", "Benefits are vague or uncompelling").
					AddRangeAnchor(0, 2.9, "Poor", "No clear benefits stated"),
			},
			{
				ID:          "ask_clarity",
				Name:        "Clear Ask",
				Description: "Whether the ask (resources, decisions needed) is clear",
				Weight:      0.15,
				Rubric: newRubric("ask_clarity", "Ask clarity").
					AddRangeAnchor(9, 10, "Excellent", "Specific ask with clear next steps and timeline").
					AddRangeAnchor(7, 8.9, "Good", "Clear ask with reasonable next steps").
					AddRangeAnchor(5, 6.9, "Adequate", "Ask present but vague").
					AddRangeAnchor(3, 4.9, "Needs Work", "Unclear what is being requested").
					AddRangeAnchor(0, 2.9, "Poor", "No ask or call to action"),
			},
		},
	}
}
