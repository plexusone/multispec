package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewMRDRubricSet())
}

// NewMRDRubricSet creates the rubric set for Market Requirements Documents.
func NewMRDRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeMRD,
		Name:         "MRD Evaluation",
		Description:  "Evaluates Market Requirements Documents for completeness, clarity, and market insight",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "problem_statement",
				Name:        "Problem Statement",
				Description: "Clarity and relevance of the market problem being addressed",
				Weight:      0.20,
				Rubric: newRubric("problem_statement", "Problem definition quality").
					AddRangeAnchor(9, 10, "Excellent", "Problem is clearly articulated with quantified impact, timing justification, and cost of inaction").
					AddRangeAnchor(7, 8.9, "Good", "Problem is clear and relevant with some supporting data").
					AddRangeAnchor(5, 6.9, "Adequate", "Problem is stated but lacks depth or supporting evidence").
					AddRangeAnchor(3, 4.9, "Needs Work", "Problem is vague or poorly justified").
					AddRangeAnchor(0, 2.9, "Poor", "Problem statement is missing or unclear"),
			},
			{
				ID:          "target_market",
				Name:        "Target Market",
				Description: "Definition and sizing of target market segments",
				Weight:      0.20,
				Rubric: newRubric("target_market", "Market definition quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear TAM/SAM/SOM with detailed segment profiles and sizing methodology").
					AddRangeAnchor(7, 8.9, "Good", "Market segments defined with reasonable sizing estimates").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic market definition present but incomplete").
					AddRangeAnchor(3, 4.9, "Needs Work", "Market definition is too broad or unsupported").
					AddRangeAnchor(0, 2.9, "Poor", "Market definition missing or completely inadequate"),
			},
			{
				ID:          "competitive_analysis",
				Name:        "Competitive Analysis",
				Description: "Understanding of competitive landscape and differentiation",
				Weight:      0.15,
				Rubric: newRubric("competitive_analysis", "Competitive insight quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive analysis of direct/indirect competitors with clear differentiation strategy").
					AddRangeAnchor(7, 8.9, "Good", "Good competitive overview with reasonable differentiation points").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic competitor list but limited analysis").
					AddRangeAnchor(3, 4.9, "Needs Work", "Incomplete or superficial competitive analysis").
					AddRangeAnchor(0, 2.9, "Poor", "Competitive analysis missing"),
			},
			{
				ID:          "requirements_clarity",
				Name:        "Requirements Clarity",
				Description: "Clarity and prioritization of market requirements",
				Weight:      0.20,
				Rubric: newRubric("requirements_clarity", "Requirements definition quality").
					AddRangeAnchor(9, 10, "Excellent", "Requirements are prioritized (MoSCoW), traceable, and well-rationalized").
					AddRangeAnchor(7, 8.9, "Good", "Requirements are clear with reasonable prioritization").
					AddRangeAnchor(5, 6.9, "Adequate", "Requirements listed but prioritization is weak").
					AddRangeAnchor(3, 4.9, "Needs Work", "Requirements are vague or poorly organized").
					AddRangeAnchor(0, 2.9, "Poor", "Requirements section is inadequate"),
			},
			{
				ID:          "business_alignment",
				Name:        "Business Alignment",
				Description: "Alignment with business goals and measurable success metrics",
				Weight:      0.15,
				Rubric: newRubric("business_alignment", "Business alignment quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear strategic alignment with quantified success metrics and timelines").
					AddRangeAnchor(7, 8.9, "Good", "Business goals are clear with defined metrics").
					AddRangeAnchor(5, 6.9, "Adequate", "Business alignment stated but metrics are weak").
					AddRangeAnchor(3, 4.9, "Needs Work", "Business alignment is unclear or missing metrics").
					AddRangeAnchor(0, 2.9, "Poor", "No clear business alignment"),
			},
			{
				ID:          "risk_assessment",
				Name:        "Risk Assessment",
				Description: "Identification of constraints, assumptions, and risks",
				Weight:      0.10,
				Rubric: newRubric("risk_assessment", "Risk assessment quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive risk matrix with impact, probability, and mitigation strategies").
					AddRangeAnchor(7, 8.9, "Good", "Key risks identified with mitigation approaches").
					AddRangeAnchor(5, 6.9, "Adequate", "Risks listed but mitigation is incomplete").
					AddRangeAnchor(3, 4.9, "Needs Work", "Risk assessment is superficial").
					AddRangeAnchor(0, 2.9, "Poor", "Risk assessment missing"),
			},
		},
	}
}
