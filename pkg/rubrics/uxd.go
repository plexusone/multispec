package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewUXDRubricSet())
}

// NewUXDRubricSet creates the rubric set for User Experience Design documents.
func NewUXDRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeUXD,
		Name:         "UXD Evaluation",
		Description:  "Evaluates User Experience Design documents for usability, accessibility, and completeness",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "design_goals",
				Name:        "Design Goals",
				Description: "Clarity of UX objectives and design principles",
				Weight:      0.10,
				Rubric: newRubric("design_goals", "Design goals quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear, measurable UX goals aligned with product strategy and guiding principles").
					AddRangeAnchor(7, 8.9, "Good", "Goals are clear with reasonable principles").
					AddRangeAnchor(5, 6.9, "Adequate", "Goals present but could be more specific").
					AddRangeAnchor(3, 4.9, "Needs Work", "Goals are vague or missing principles").
					AddRangeAnchor(0, 2.9, "Poor", "Design goals missing"),
			},
			{
				ID:          "user_personas",
				Name:        "User Personas",
				Description: "Quality and depth of user persona definitions",
				Weight:      0.15,
				Rubric: newRubric("user_personas", "Persona quality").
					AddRangeAnchor(9, 10, "Excellent", "Rich personas with goals, pain points, behaviors, and scenarios").
					AddRangeAnchor(7, 8.9, "Good", "Well-defined personas with key attributes").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic personas present but shallow").
					AddRangeAnchor(3, 4.9, "Needs Work", "Personas are too generic").
					AddRangeAnchor(0, 2.9, "Poor", "Personas missing or inadequate"),
			},
			{
				ID:          "user_journeys",
				Name:        "User Journeys",
				Description: "Completeness of user journey mapping",
				Weight:      0.15,
				Rubric: newRubric("user_journeys", "Journey mapping quality").
					AddRangeAnchor(9, 10, "Excellent", "Complete journeys with entry/exit points, success criteria, and edge cases").
					AddRangeAnchor(7, 8.9, "Good", "Key journeys mapped with clear flows").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic flows present but incomplete").
					AddRangeAnchor(3, 4.9, "Needs Work", "Journeys are fragmentary").
					AddRangeAnchor(0, 2.9, "Poor", "User journeys missing"),
			},
			{
				ID:          "information_architecture",
				Name:        "Information Architecture",
				Description: "Structure and organization of content",
				Weight:      0.15,
				Rubric: newRubric("information_architecture", "IA quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear site map, navigation structure, and content hierarchy").
					AddRangeAnchor(7, 8.9, "Good", "Reasonable IA with clear navigation").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic structure present").
					AddRangeAnchor(3, 4.9, "Needs Work", "IA is unclear or inconsistent").
					AddRangeAnchor(0, 2.9, "Poor", "Information architecture missing"),
			},
			{
				ID:          "wireframes",
				Name:        "Wireframes and Layouts",
				Description: "Quality and coverage of wireframes",
				Weight:      0.15,
				Rubric: newRubric("wireframes", "Wireframe quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive wireframes for all key screens with element descriptions").
					AddRangeAnchor(7, 8.9, "Good", "Key screens covered with reasonable detail").
					AddRangeAnchor(5, 6.9, "Adequate", "Some wireframes but incomplete coverage").
					AddRangeAnchor(3, 4.9, "Needs Work", "Wireframes are sparse or unclear").
					AddRangeAnchor(0, 2.9, "Poor", "Wireframes missing"),
			},
			{
				ID:          "interaction_design",
				Name:        "Interaction Design",
				Description: "Definition of interaction patterns and behaviors",
				Weight:      0.10,
				Rubric: newRubric("interaction_design", "Interaction design quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear interaction patterns, micro-interactions, and state transitions").
					AddRangeAnchor(7, 8.9, "Good", "Key interactions defined").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic interactions present").
					AddRangeAnchor(3, 4.9, "Needs Work", "Interactions poorly defined").
					AddRangeAnchor(0, 2.9, "Poor", "Interaction design missing"),
			},
			{
				ID:          "accessibility",
				Name:        "Accessibility",
				Description: "WCAG compliance and inclusive design considerations",
				Weight:      0.15,
				Rubric: newRubric("accessibility", "Accessibility quality").
					AddRangeAnchor(9, 10, "Excellent", "WCAG compliance specified, comprehensive checklist, keyboard/screen reader considerations").
					AddRangeAnchor(7, 8.9, "Good", "Accessibility requirements defined with checklist").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic accessibility mentioned").
					AddRangeAnchor(3, 4.9, "Needs Work", "Accessibility is incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "Accessibility not addressed"),
			},
			{
				ID:          "error_states",
				Name:        "Error States and Edge Cases",
				Description: "Handling of errors, empty states, and edge cases",
				Weight:      0.05,
				Rubric: newRubric("error_states", "Error handling quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive error messages, empty states, loading states, and recovery actions").
					AddRangeAnchor(7, 8.9, "Good", "Key error and edge cases covered").
					AddRangeAnchor(5, 6.9, "Adequate", "Some error states defined").
					AddRangeAnchor(3, 4.9, "Needs Work", "Error handling incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "Error states not addressed"),
			},
		},
	}
}
