package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewPressRubricSet())
}

// NewPressRubricSet creates the rubric set for Press Release documents.
func NewPressRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypePress,
		Name:         "Press Release Evaluation",
		Description:  "Evaluates Press Release documents for clarity, customer focus, and completeness",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "headline_impact",
				Name:        "Headline Impact",
				Description: "Whether the headline clearly communicates the primary customer benefit",
				Weight:      0.20,
				Rubric: newRubric("headline_impact", "Headline quality").
					AddRangeAnchor(9, 10, "Excellent", "Compelling headline with clear customer benefit and urgency").
					AddRangeAnchor(7, 8.9, "Good", "Clear headline with identifiable benefit").
					AddRangeAnchor(5, 6.9, "Adequate", "Headline present but benefit unclear").
					AddRangeAnchor(3, 4.9, "Needs Work", "Headline is vague or feature-focused").
					AddRangeAnchor(0, 2.9, "Poor", "No clear headline or irrelevant"),
			},
			{
				ID:          "customer_problem",
				Name:        "Customer Problem",
				Description: "How well the customer problem is articulated with authentic voice",
				Weight:      0.20,
				Rubric: newRubric("customer_problem", "Problem articulation").
					AddRangeAnchor(9, 10, "Excellent", "Vivid problem description with authentic customer quotes").
					AddRangeAnchor(7, 8.9, "Good", "Clear problem with customer perspective").
					AddRangeAnchor(5, 6.9, "Adequate", "Problem stated but lacks customer voice").
					AddRangeAnchor(3, 4.9, "Needs Work", "Problem is technical or internal-focused").
					AddRangeAnchor(0, 2.9, "Poor", "No clear customer problem"),
			},
			{
				ID:          "solution_clarity",
				Name:        "Solution Clarity",
				Description: "Whether the solution and benefits are clearly explained",
				Weight:      0.20,
				Rubric: newRubric("solution_clarity", "Solution description quality").
					AddRangeAnchor(9, 10, "Excellent", "Crystal clear solution with concrete, measurable benefits").
					AddRangeAnchor(7, 8.9, "Good", "Clear solution with understandable benefits").
					AddRangeAnchor(5, 6.9, "Adequate", "Solution described but benefits vague").
					AddRangeAnchor(3, 4.9, "Needs Work", "Solution is too technical or confusing").
					AddRangeAnchor(0, 2.9, "Poor", "No clear solution articulated"),
			},
			{
				ID:          "customer_validation",
				Name:        "Customer Validation",
				Description: "Quality of customer quotes and validation",
				Weight:      0.15,
				Rubric: newRubric("customer_validation", "Validation quality").
					AddRangeAnchor(9, 10, "Excellent", "Authentic customer quotes with specific success stories").
					AddRangeAnchor(7, 8.9, "Good", "Credible customer perspective included").
					AddRangeAnchor(5, 6.9, "Adequate", "Generic quotes without specificity").
					AddRangeAnchor(3, 4.9, "Needs Work", "Quotes feel artificial or forced").
					AddRangeAnchor(0, 2.9, "Poor", "No customer validation"),
			},
			{
				ID:          "call_to_action",
				Name:        "Call to Action",
				Description: "Clarity of availability, pricing, and next steps",
				Weight:      0.15,
				Rubric: newRubric("call_to_action", "CTA quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear availability, pricing, and specific action steps").
					AddRangeAnchor(7, 8.9, "Good", "Actionable next steps provided").
					AddRangeAnchor(5, 6.9, "Adequate", "CTA present but vague").
					AddRangeAnchor(3, 4.9, "Needs Work", "Unclear how to proceed").
					AddRangeAnchor(0, 2.9, "Poor", "No call to action"),
			},
			{
				ID:          "readability",
				Name:        "Readability",
				Description: "Whether the press release is written in plain, accessible language",
				Weight:      0.10,
				Rubric: newRubric("readability", "Writing quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear, jargon-free prose accessible to general audience").
					AddRangeAnchor(7, 8.9, "Good", "Readable with minimal jargon").
					AddRangeAnchor(5, 6.9, "Adequate", "Somewhat technical but understandable").
					AddRangeAnchor(3, 4.9, "Needs Work", "Too technical or full of jargon").
					AddRangeAnchor(0, 2.9, "Poor", "Incomprehensible to target audience"),
			},
		},
	}
}
