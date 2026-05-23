package rubrics

import (
	"testing"

	"github.com/plexusone/multispec/pkg/types"
)

func TestGet(t *testing.T) {
	tests := []struct {
		specType types.SpecType
		wantErr  bool
		wantName string
	}{
		{types.SpecTypeMRD, false, "MRD Evaluation"},
		{types.SpecTypePRD, false, "PRD Evaluation"},
		{types.SpecTypeUXD, false, "UXD Evaluation"},
		{types.SpecTypePress, false, "Press Release Evaluation"},
		{types.SpecTypeFAQ, false, "FAQ Evaluation"},
		{types.SpecTypeNarrative1P, false, "Narrative 1-Pager Evaluation"},
		{types.SpecTypeNarrative6P, false, "Narrative 6-Pager Evaluation"},
		{types.SpecTypeTRD, false, "TRD Evaluation"},
		{types.SpecTypeIRD, false, "IRD Evaluation"},
		{"unknown", true, ""}, // No rubric for unknown type
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			rs, err := Get(tt.specType)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Get(%s) expected error, got nil", tt.specType)
				}
				return
			}
			if err != nil {
				t.Errorf("Get(%s) unexpected error: %v", tt.specType, err)
				return
			}
			if rs.Name != tt.wantName {
				t.Errorf("Get(%s).Name = %q, want %q", tt.specType, rs.Name, tt.wantName)
			}
		})
	}
}

func TestMustGet(t *testing.T) {
	// Should not panic for valid types
	rs := MustGet(types.SpecTypeMRD)
	if rs == nil {
		t.Error("MustGet(MRD) returned nil")
	}

	// Should panic for invalid types
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGet(unknown) did not panic")
		}
	}()
	MustGet("unknown")
}

func TestAvailable(t *testing.T) {
	available := Available()
	if len(available) < 3 {
		t.Errorf("Available() returned %d types, expected at least 3", len(available))
	}
}

func TestHasRubric(t *testing.T) {
	hasRubricTypes := []types.SpecType{
		types.SpecTypeMRD,
		types.SpecTypePRD,
		types.SpecTypeUXD,
		types.SpecTypePress,
		types.SpecTypeFAQ,
		types.SpecTypeNarrative1P,
		types.SpecTypeNarrative6P,
		types.SpecTypeTRD,
		types.SpecTypeIRD,
	}
	for _, st := range hasRubricTypes {
		if !HasRubric(st) {
			t.Errorf("HasRubric(%s) returned false", st)
		}
	}
	if HasRubric("unknown") {
		t.Error("HasRubric(unknown) returned true")
	}
}

func TestRubricSetTotalWeight(t *testing.T) {
	specTypes := []types.SpecType{
		types.SpecTypeMRD,
		types.SpecTypePRD,
		types.SpecTypeUXD,
		types.SpecTypePress,
		types.SpecTypeFAQ,
		types.SpecTypeNarrative1P,
		types.SpecTypeNarrative6P,
		types.SpecTypeTRD,
		types.SpecTypeIRD,
	}

	for _, st := range specTypes {
		rs := MustGet(st)
		total := rs.TotalWeight()

		// Total weight should be 1.0 (100%)
		if total < 0.99 || total > 1.01 {
			t.Errorf("%s TotalWeight() = %f, expected ~1.0", st, total)
		}
	}
}

func TestRubricSetCategoryByID(t *testing.T) {
	rs := MustGet(types.SpecTypePRD)

	cat, ok := rs.CategoryByID("user_stories")
	if !ok {
		t.Fatal("CategoryByID('user_stories') not found")
	}
	if cat.Name != "User Stories" {
		t.Errorf("Category name = %q, want 'User Stories'", cat.Name)
	}
	if cat.Weight != 0.20 {
		t.Errorf("Category weight = %f, want 0.20", cat.Weight)
	}

	_, ok = rs.CategoryByID("nonexistent")
	if ok {
		t.Error("CategoryByID('nonexistent') should return false")
	}
}

func TestCategoryRubrics(t *testing.T) {
	rs := MustGet(types.SpecTypeMRD)

	for _, cat := range rs.Categories {
		if cat.Rubric == nil {
			t.Errorf("Category %q has nil rubric", cat.ID)
		}
	}
}

func TestPassCriteria(t *testing.T) {
	defaultCriteria := DefaultPassCriteria()
	if defaultCriteria.MinScore != 7.0 {
		t.Errorf("DefaultPassCriteria().MinScore = %f, want 7.0", defaultCriteria.MinScore)
	}
	if defaultCriteria.MaxCritical != 0 {
		t.Errorf("DefaultPassCriteria().MaxCritical = %d, want 0", defaultCriteria.MaxCritical)
	}

	strictCriteria := StrictPassCriteria()
	if strictCriteria.MinScore != 8.0 {
		t.Errorf("StrictPassCriteria().MinScore = %f, want 8.0", strictCriteria.MinScore)
	}
	if strictCriteria.MaxMedium != 3 {
		t.Errorf("StrictPassCriteria().MaxMedium = %d, want 3", strictCriteria.MaxMedium)
	}
}
