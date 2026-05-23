package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewPRDRubricSet())
}

// NewPRDRubricSet creates the rubric set for Product Requirements Documents.
func NewPRDRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypePRD,
		Name:         "PRD Evaluation",
		Description:  "Evaluates Product Requirements Documents for completeness, testability, and clarity",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "problem_definition",
				Name:        "Problem Definition",
				Description: "Clarity of the problem statement from user perspective",
				Weight:      0.15,
				Rubric: newRubric("problem_definition", "Problem definition quality").
					AddRangeAnchor(9, 10, "Excellent", "Problem is specific, measurable, and clearly impacts users").
					AddRangeAnchor(7, 8.9, "Good", "Problem is clear with user impact stated").
					AddRangeAnchor(5, 6.9, "Adequate", "Problem exists but lacks specificity").
					AddRangeAnchor(3, 4.9, "Needs Work", "Problem is vague or too broad").
					AddRangeAnchor(0, 2.9, "Poor", "No clear problem statement"),
			},
			{
				ID:          "goals_scope",
				Name:        "Goals and Scope",
				Description: "Clarity of goals and explicit non-goals",
				Weight:      0.15,
				Rubric: newRubric("goals_scope", "Goals and scope clarity").
					AddRangeAnchor(9, 10, "Excellent", "Goals are SMART, non-goals explicitly stated, scope is bounded").
					AddRangeAnchor(7, 8.9, "Good", "Clear goals with reasonable scope boundaries").
					AddRangeAnchor(5, 6.9, "Adequate", "Goals present but scope could be clearer").
					AddRangeAnchor(3, 4.9, "Needs Work", "Goals or scope are unclear").
					AddRangeAnchor(0, 2.9, "Poor", "Goals and scope missing"),
			},
			{
				ID:          "user_stories",
				Name:        "User Stories",
				Description: "Quality and completeness of user stories",
				Weight:      0.20,
				Rubric: newRubric("user_stories", "User story quality").
					AddRangeAnchor(9, 10, "Excellent", "All stories follow standard format, are prioritized, with clear acceptance criteria").
					AddRangeAnchor(7, 8.9, "Good", "Stories are clear with acceptance criteria present").
					AddRangeAnchor(5, 6.9, "Adequate", "Stories exist but some lack acceptance criteria").
					AddRangeAnchor(3, 4.9, "Needs Work", "Stories are incomplete or poorly formatted").
					AddRangeAnchor(0, 2.9, "Poor", "User stories missing or inadequate"),
			},
			{
				ID:          "functional_requirements",
				Name:        "Functional Requirements",
				Description: "Completeness and testability of functional requirements",
				Weight:      0.20,
				Rubric: newRubric("functional_requirements", "Functional requirements quality").
					AddRangeAnchor(9, 10, "Excellent", "Requirements are complete, traceable to stories, and testable").
					AddRangeAnchor(7, 8.9, "Good", "Requirements are clear and mostly testable").
					AddRangeAnchor(5, 6.9, "Adequate", "Requirements exist but some are vague").
					AddRangeAnchor(3, 4.9, "Needs Work", "Requirements are incomplete or unclear").
					AddRangeAnchor(0, 2.9, "Poor", "Functional requirements missing"),
			},
			{
				ID:          "nonfunctional_requirements",
				Name:        "Non-Functional Requirements",
				Description: "Coverage of performance, security, reliability requirements",
				Weight:      0.15,
				Rubric: newRubric("nonfunctional_requirements", "NFR quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive NFRs with measurable targets for performance, security, accessibility").
					AddRangeAnchor(7, 8.9, "Good", "Key NFRs defined with reasonable targets").
					AddRangeAnchor(5, 6.9, "Adequate", "Some NFRs present but incomplete coverage").
					AddRangeAnchor(3, 4.9, "Needs Work", "NFRs are missing key areas").
					AddRangeAnchor(0, 2.9, "Poor", "NFRs missing"),
			},
			{
				ID:          "success_metrics",
				Name:        "Success Metrics",
				Description: "Definition of measurable success criteria",
				Weight:      0.10,
				Rubric: newRubric("success_metrics", "Success metrics quality").
					AddRangeAnchor(9, 10, "Excellent", "SMART metrics with baselines, targets, and measurement methods").
					AddRangeAnchor(7, 8.9, "Good", "Metrics defined with targets").
					AddRangeAnchor(5, 6.9, "Adequate", "Some metrics but lacking specificity").
					AddRangeAnchor(3, 4.9, "Needs Work", "Metrics are vague or unmeasurable").
					AddRangeAnchor(0, 2.9, "Poor", "Success metrics missing"),
			},
			{
				ID:          "dependencies",
				Name:        "Dependencies and Risks",
				Description: "Identification of dependencies and open questions",
				Weight:      0.05,
				Rubric: newRubric("dependencies", "Dependencies tracking quality").
					AddRangeAnchor(9, 10, "Excellent", "Dependencies mapped with owners, status, and risk mitigation").
					AddRangeAnchor(7, 8.9, "Good", "Key dependencies identified with owners").
					AddRangeAnchor(5, 6.9, "Adequate", "Some dependencies noted").
					AddRangeAnchor(3, 4.9, "Needs Work", "Dependencies incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "Dependencies not addressed"),
			},
		},
	}
}
