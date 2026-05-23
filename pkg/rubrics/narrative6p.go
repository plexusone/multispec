package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewNarrative6PRubricSet())
}

// NewNarrative6PRubricSet creates the rubric set for AWS-Style 6-Pager Narratives.
func NewNarrative6PRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeNarrative6P,
		Name:         "Narrative 6-Pager Evaluation",
		Description:  "Evaluates AWS-style 6-pager narratives for depth, rigor, and completeness",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "tenets_clarity",
				Name:        "Tenets and Introduction",
				Description: "Quality of guiding principles and context setting",
				Weight:      0.10,
				Rubric: newRubric("tenets_clarity", "Tenets quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear, actionable tenets that guide tradeoff decisions").
					AddRangeAnchor(7, 8.9, "Good", "Reasonable tenets with clear intent").
					AddRangeAnchor(5, 6.9, "Adequate", "Tenets present but generic").
					AddRangeAnchor(3, 4.9, "Needs Work", "Tenets are vague or not useful").
					AddRangeAnchor(0, 2.9, "Poor", "No tenets or introduction"),
			},
			{
				ID:          "customer_depth",
				Name:        "Customer Problem Depth",
				Description: "How deeply the customer and their problem are understood",
				Weight:      0.20,
				Rubric: newRubric("customer_depth", "Customer understanding").
					AddRangeAnchor(9, 10, "Excellent", "Specific customer with vivid problem description, quotes, and anecdotes").
					AddRangeAnchor(7, 8.9, "Good", "Clear customer definition with good problem articulation").
					AddRangeAnchor(5, 6.9, "Adequate", "Customer and problem described but lacks depth").
					AddRangeAnchor(3, 4.9, "Needs Work", "Customer is generic, problem is surface-level").
					AddRangeAnchor(0, 2.9, "Poor", "No meaningful customer or problem description"),
			},
			{
				ID:          "solution_narrative",
				Name:        "Solution Narrative",
				Description: "Quality of the customer-centric solution description",
				Weight:      0.20,
				Rubric: newRubric("solution_narrative", "Solution quality").
					AddRangeAnchor(9, 10, "Excellent", "Compelling customer journey with clear capabilities and scope").
					AddRangeAnchor(7, 8.9, "Good", "Clear solution from customer perspective").
					AddRangeAnchor(5, 6.9, "Adequate", "Solution described but too technical or feature-focused").
					AddRangeAnchor(3, 4.9, "Needs Work", "Solution is unclear or disconnected from customer").
					AddRangeAnchor(0, 2.9, "Poor", "No coherent solution narrative"),
			},
			{
				ID:          "timing_justification",
				Name:        "Why Now Justification",
				Description: "Strength of the timing and urgency argument",
				Weight:      0.15,
				Rubric: newRubric("timing_justification", "Timing argument").
					AddRangeAnchor(9, 10, "Excellent", "Compelling market, competitive, and internal readiness arguments").
					AddRangeAnchor(7, 8.9, "Good", "Clear timing justification with supporting evidence").
					AddRangeAnchor(5, 6.9, "Adequate", "Timing argument present but weak").
					AddRangeAnchor(3, 4.9, "Needs Work", "No convincing reason for urgency").
					AddRangeAnchor(0, 2.9, "Poor", "No timing justification"),
			},
			{
				ID:          "business_case",
				Name:        "Business Case Rigor",
				Description: "Quality of customer impact, business metrics, and ROI analysis",
				Weight:      0.20,
				Rubric: newRubric("business_case", "Business case quality").
					AddRangeAnchor(9, 10, "Excellent", "Quantified customer and business impact with clear assumptions").
					AddRangeAnchor(7, 8.9, "Good", "Reasonable metrics with directionally correct estimates").
					AddRangeAnchor(5, 6.9, "Adequate", "Some metrics but lacking rigor or assumptions").
					AddRangeAnchor(3, 4.9, "Needs Work", "Weak or unsupported business case").
					AddRangeAnchor(0, 2.9, "Poor", "No business case"),
			},
			{
				ID:          "risks_dependencies",
				Name:        "Risks and Dependencies",
				Description: "Thoroughness of risk identification and mitigation",
				Weight:      0.15,
				Rubric: newRubric("risks_dependencies", "Risk assessment quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive risks with likelihood, impact, and mitigations").
					AddRangeAnchor(7, 8.9, "Good", "Key risks identified with mitigation strategies").
					AddRangeAnchor(5, 6.9, "Adequate", "Some risks listed but incomplete analysis").
					AddRangeAnchor(3, 4.9, "Needs Work", "Risks are superficial or missing mitigations").
					AddRangeAnchor(0, 2.9, "Poor", "No risk assessment"),
			},
		},
	}
}
