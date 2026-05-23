package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewTRDRubricSet())
}

// NewTRDRubricSet creates the rubric set for Technical Requirements Documents.
func NewTRDRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeTRD,
		Name:         "TRD Evaluation",
		Description:  "Evaluates Technical Requirements Documents for completeness, clarity, and implementability",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "architecture_clarity",
				Name:        "Architecture Clarity",
				Description: "How clearly the system architecture is described",
				Weight:      0.20,
				Rubric: newRubric("architecture_clarity", "Architecture quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear architecture with diagrams, component descriptions, and data flow").
					AddRangeAnchor(7, 8.9, "Good", "Architecture is understandable with reasonable detail").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic architecture present but lacks detail").
					AddRangeAnchor(3, 4.9, "Needs Work", "Architecture is unclear or incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "No clear architecture described"),
			},
			{
				ID:          "prd_traceability",
				Name:        "PRD Traceability",
				Description: "Whether technical design traces back to PRD requirements",
				Weight:      0.15,
				Rubric: newRubric("prd_traceability", "Traceability quality").
					AddRangeAnchor(9, 10, "Excellent", "Every PRD requirement has corresponding technical design").
					AddRangeAnchor(7, 8.9, "Good", "Most requirements are traced to technical approach").
					AddRangeAnchor(5, 6.9, "Adequate", "Some traceability but gaps exist").
					AddRangeAnchor(3, 4.9, "Needs Work", "Weak connection to PRD").
					AddRangeAnchor(0, 2.9, "Poor", "No traceability to requirements"),
			},
			{
				ID:          "api_design",
				Name:        "API Design",
				Description: "Quality of API/interface definitions",
				Weight:      0.15,
				Rubric: newRubric("api_design", "API quality").
					AddRangeAnchor(9, 10, "Excellent", "Complete API spec with request/response schemas and error handling").
					AddRangeAnchor(7, 8.9, "Good", "Clear API design with reasonable detail").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic API definition present").
					AddRangeAnchor(3, 4.9, "Needs Work", "API design is incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "No API design"),
			},
			{
				ID:          "data_design",
				Name:        "Data Design",
				Description: "Quality of data models and storage design",
				Weight:      0.15,
				Rubric: newRubric("data_design", "Data design quality").
					AddRangeAnchor(9, 10, "Excellent", "Complete data models with relationships, storage rationale, and migration plan").
					AddRangeAnchor(7, 8.9, "Good", "Clear data models with storage decisions").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic data design present").
					AddRangeAnchor(3, 4.9, "Needs Work", "Data design is incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "No data design"),
			},
			{
				ID:          "nfr_coverage",
				Name:        "NFR Coverage",
				Description: "Coverage of non-functional requirements (performance, security, etc.)",
				Weight:      0.20,
				Rubric: newRubric("nfr_coverage", "NFR coverage quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive NFRs with specific targets and measurement plans").
					AddRangeAnchor(7, 8.9, "Good", "Key NFRs covered with reasonable targets").
					AddRangeAnchor(5, 6.9, "Adequate", "Some NFRs but gaps in coverage").
					AddRangeAnchor(3, 4.9, "Needs Work", "NFRs are vague or missing key areas").
					AddRangeAnchor(0, 2.9, "Poor", "No NFR coverage"),
			},
			{
				ID:          "implementability",
				Name:        "Implementability",
				Description: "Whether the design is implementable by an engineer",
				Weight:      0.15,
				Rubric: newRubric("implementability", "Implementation readiness").
					AddRangeAnchor(9, 10, "Excellent", "Design is actionable with clear path to implementation").
					AddRangeAnchor(7, 8.9, "Good", "Design can be implemented with minimal clarification").
					AddRangeAnchor(5, 6.9, "Adequate", "Design needs some clarification before implementation").
					AddRangeAnchor(3, 4.9, "Needs Work", "Too many ambiguities to implement").
					AddRangeAnchor(0, 2.9, "Poor", "Cannot be implemented from this document"),
			},
		},
	}
}
