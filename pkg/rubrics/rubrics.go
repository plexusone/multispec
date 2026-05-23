// Package rubrics provides evaluation rubrics for spec types.
package rubrics

import (
	"fmt"

	"github.com/plexusone/multispec/pkg/types"
	"github.com/plexusone/structured-evaluation/evaluation"
)

// RubricSet contains all criteria for evaluating a spec type.
type RubricSet struct {
	SpecType    types.SpecType
	Name        string
	Description string
	Categories  []*Category
	PassCriteria
}

// Category represents a scoring category with its rubric.
type Category struct {
	ID          string
	Name        string
	Description string
	Weight      float64
	Rubric      *evaluation.Rubric
}

// PassCriteria defines what constitutes a passing evaluation.
type PassCriteria struct {
	MinScore    float64
	MaxCritical int
	MaxHigh     int
	MaxMedium   int // -1 = unlimited
}

// DefaultPassCriteria returns the default pass criteria.
func DefaultPassCriteria() PassCriteria {
	return PassCriteria{
		MinScore:    7.0,
		MaxCritical: 0,
		MaxHigh:     0,
		MaxMedium:   -1, // unlimited
	}
}

// StrictPassCriteria returns stricter pass criteria.
func StrictPassCriteria() PassCriteria {
	return PassCriteria{
		MinScore:    8.0,
		MaxCritical: 0,
		MaxHigh:     0,
		MaxMedium:   3,
	}
}

// registry maps spec types to their rubric sets.
var registry = make(map[types.SpecType]*RubricSet)

// Register adds a rubric set to the registry.
func Register(rs *RubricSet) {
	registry[rs.SpecType] = rs
}

// Get returns the rubric set for a spec type.
func Get(specType types.SpecType) (*RubricSet, error) {
	rs, ok := registry[specType]
	if !ok {
		return nil, fmt.Errorf("rubric not found for spec type %q", specType)
	}
	return rs, nil
}

// MustGet returns the rubric set for a spec type, panicking on error.
func MustGet(specType types.SpecType) *RubricSet {
	rs, err := Get(specType)
	if err != nil {
		panic(err)
	}
	return rs
}

// Available returns all spec types with registered rubrics.
func Available() []types.SpecType {
	var result []types.SpecType
	for st := range registry {
		result = append(result, st)
	}
	return result
}

// HasRubric returns true if a rubric exists for the spec type.
func HasRubric(specType types.SpecType) bool {
	_, ok := registry[specType]
	return ok
}

// TotalWeight returns the sum of all category weights.
func (rs *RubricSet) TotalWeight() float64 {
	var total float64
	for _, cat := range rs.Categories {
		total += cat.Weight
	}
	return total
}

// CategoryByID returns a category by its ID.
func (rs *RubricSet) CategoryByID(id string) (*Category, bool) {
	for _, cat := range rs.Categories {
		if cat.ID == id {
			return cat, true
		}
	}
	return nil, false
}

// newRubric creates a standard rubric with common anchors.
func newRubric(name, description string) *evaluation.Rubric {
	return evaluation.NewRubric(name, description).
		AddRangeAnchor(9.0, 10.0, "Excellent", "Exceeds expectations in all aspects").
		AddRangeAnchor(7.0, 8.9, "Good", "Meets expectations with minor improvements possible").
		AddRangeAnchor(5.0, 6.9, "Adequate", "Meets minimum requirements but needs improvement").
		AddRangeAnchor(3.0, 4.9, "Needs Work", "Below expectations, requires significant revision").
		AddRangeAnchor(0.0, 2.9, "Poor", "Does not meet basic requirements")
}
