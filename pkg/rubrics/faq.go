package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewFAQRubricSet())
}

// NewFAQRubricSet creates the rubric set for FAQ documents.
func NewFAQRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeFAQ,
		Name:         "FAQ Evaluation",
		Description:  "Evaluates FAQ documents for completeness, clarity, and anticipation of customer questions",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "question_coverage",
				Name:        "Question Coverage",
				Description: "Whether the FAQ anticipates the questions customers will actually ask",
				Weight:      0.25,
				Rubric: newRubric("question_coverage", "Coverage quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive coverage of likely questions across all topics").
					AddRangeAnchor(7, 8.9, "Good", "Good coverage of main question areas").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic questions covered but gaps exist").
					AddRangeAnchor(3, 4.9, "Needs Work", "Missing important question categories").
					AddRangeAnchor(0, 2.9, "Poor", "Inadequate question coverage"),
			},
			{
				ID:          "answer_clarity",
				Name:        "Answer Clarity",
				Description: "How clearly and directly questions are answered",
				Weight:      0.25,
				Rubric: newRubric("answer_clarity", "Answer quality").
					AddRangeAnchor(9, 10, "Excellent", "Direct, clear answers with no ambiguity").
					AddRangeAnchor(7, 8.9, "Good", "Clear answers that address the question").
					AddRangeAnchor(5, 6.9, "Adequate", "Answers present but could be clearer").
					AddRangeAnchor(3, 4.9, "Needs Work", "Answers are vague or indirect").
					AddRangeAnchor(0, 2.9, "Poor", "Answers don't address questions"),
			},
			{
				ID:          "customer_language",
				Name:        "Customer Language",
				Description: "Whether questions and answers use customer language, not internal jargon",
				Weight:      0.15,
				Rubric: newRubric("customer_language", "Language appropriateness").
					AddRangeAnchor(9, 10, "Excellent", "Questions phrased as customers would ask, answers in plain language").
					AddRangeAnchor(7, 8.9, "Good", "Customer-friendly language throughout").
					AddRangeAnchor(5, 6.9, "Adequate", "Some jargon but generally accessible").
					AddRangeAnchor(3, 4.9, "Needs Work", "Too much internal or technical language").
					AddRangeAnchor(0, 2.9, "Poor", "Written for internal audience, not customers"),
			},
			{
				ID:          "pricing_transparency",
				Name:        "Pricing Transparency",
				Description: "Clarity of pricing and cost-related information",
				Weight:      0.15,
				Rubric: newRubric("pricing_transparency", "Pricing clarity").
					AddRangeAnchor(9, 10, "Excellent", "Clear pricing with examples and cost estimation guidance").
					AddRangeAnchor(7, 8.9, "Good", "Pricing information is clear and complete").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic pricing present but lacks detail").
					AddRangeAnchor(3, 4.9, "Needs Work", "Pricing unclear or hard to understand").
					AddRangeAnchor(0, 2.9, "Poor", "No pricing information"),
			},
			{
				ID:          "getting_started",
				Name:        "Getting Started Clarity",
				Description: "How well the path to adoption is explained",
				Weight:      0.10,
				Rubric: newRubric("getting_started", "Onboarding clarity").
					AddRangeAnchor(9, 10, "Excellent", "Clear step-by-step path from interest to success").
					AddRangeAnchor(7, 8.9, "Good", "Good explanation of how to get started").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic getting started info present").
					AddRangeAnchor(3, 4.9, "Needs Work", "Unclear how to proceed").
					AddRangeAnchor(0, 2.9, "Poor", "No getting started guidance"),
			},
			{
				ID:          "objection_handling",
				Name:        "Objection Handling",
				Description: "Whether the FAQ addresses likely concerns and objections",
				Weight:      0.10,
				Rubric: newRubric("objection_handling", "Objection coverage").
					AddRangeAnchor(9, 10, "Excellent", "Proactively addresses concerns with honest, confident responses").
					AddRangeAnchor(7, 8.9, "Good", "Key objections addressed reasonably").
					AddRangeAnchor(5, 6.9, "Adequate", "Some objections covered").
					AddRangeAnchor(3, 4.9, "Needs Work", "Avoids difficult questions").
					AddRangeAnchor(0, 2.9, "Poor", "No objection handling"),
			},
		},
	}
}
