package rubrics

import "github.com/plexusone/multispec/pkg/types"

func init() {
	Register(NewIRDRubricSet())
}

// NewIRDRubricSet creates the rubric set for Infrastructure Requirements Documents.
func NewIRDRubricSet() *RubricSet {
	return &RubricSet{
		SpecType:     types.SpecTypeIRD,
		Name:         "IRD Evaluation",
		Description:  "Evaluates Infrastructure Requirements Documents for completeness, security, and operability",
		PassCriteria: DefaultPassCriteria(),
		Categories: []*Category{
			{
				ID:          "architecture_completeness",
				Name:        "Architecture Completeness",
				Description: "Whether all infrastructure components are documented",
				Weight:      0.20,
				Rubric: newRubric("architecture_completeness", "Architecture coverage").
					AddRangeAnchor(9, 10, "Excellent", "Complete infrastructure diagram with all components and connections").
					AddRangeAnchor(7, 8.9, "Good", "Infrastructure is well documented with reasonable detail").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic infrastructure documented but gaps exist").
					AddRangeAnchor(3, 4.9, "Needs Work", "Infrastructure documentation is incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "No clear infrastructure architecture"),
			},
			{
				ID:          "security_design",
				Name:        "Security Design",
				Description: "Coverage of security controls, IAM, encryption, and compliance",
				Weight:      0.20,
				Rubric: newRubric("security_design", "Security quality").
					AddRangeAnchor(9, 10, "Excellent", "Comprehensive security design with IAM, encryption, secrets, and compliance").
					AddRangeAnchor(7, 8.9, "Good", "Security controls documented with reasonable coverage").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic security present but incomplete").
					AddRangeAnchor(3, 4.9, "Needs Work", "Security design has significant gaps").
					AddRangeAnchor(0, 2.9, "Poor", "No security design"),
			},
			{
				ID:          "availability_dr",
				Name:        "Availability and DR",
				Description: "Coverage of availability targets, failover, and disaster recovery",
				Weight:      0.15,
				Rubric: newRubric("availability_dr", "Availability/DR quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear SLAs, multi-region strategy, and DR procedures").
					AddRangeAnchor(7, 8.9, "Good", "Availability targets with failover plan").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic availability documented").
					AddRangeAnchor(3, 4.9, "Needs Work", "Availability planning is weak").
					AddRangeAnchor(0, 2.9, "Poor", "No availability or DR planning"),
			},
			{
				ID:          "observability",
				Name:        "Observability",
				Description: "Coverage of logging, metrics, tracing, and alerting",
				Weight:      0.15,
				Rubric: newRubric("observability", "Observability quality").
					AddRangeAnchor(9, 10, "Excellent", "Complete observability stack with dashboards and alert definitions").
					AddRangeAnchor(7, 8.9, "Good", "Logging, metrics, and alerting documented").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic observability present").
					AddRangeAnchor(3, 4.9, "Needs Work", "Observability is incomplete").
					AddRangeAnchor(0, 2.9, "Poor", "No observability planning"),
			},
			{
				ID:          "capacity_cost",
				Name:        "Capacity and Cost",
				Description: "Capacity planning and cost estimation",
				Weight:      0.15,
				Rubric: newRubric("capacity_cost", "Capacity/cost planning").
					AddRangeAnchor(9, 10, "Excellent", "Detailed capacity projections with cost breakdown and scaling strategy").
					AddRangeAnchor(7, 8.9, "Good", "Reasonable capacity planning with cost estimates").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic capacity and cost present").
					AddRangeAnchor(3, 4.9, "Needs Work", "Capacity or cost planning weak").
					AddRangeAnchor(0, 2.9, "Poor", "No capacity or cost planning"),
			},
			{
				ID:          "operability",
				Name:        "Operability",
				Description: "Whether infrastructure can be operated and maintained",
				Weight:      0.15,
				Rubric: newRubric("operability", "Operability quality").
					AddRangeAnchor(9, 10, "Excellent", "Clear IaC approach, CI/CD infrastructure, and runbooks").
					AddRangeAnchor(7, 8.9, "Good", "Infrastructure is operable with documented procedures").
					AddRangeAnchor(5, 6.9, "Adequate", "Basic operational procedures").
					AddRangeAnchor(3, 4.9, "Needs Work", "Operational gaps exist").
					AddRangeAnchor(0, 2.9, "Poor", "Cannot be operated from this document"),
			},
		},
	}
}
