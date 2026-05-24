package types

import "testing"

func TestSpecTypeCategory(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected SpecCategory
	}{
		{SpecTypeMRD, CategorySource},
		{SpecTypePRD, CategorySource},
		{SpecTypeUXD, CategorySource},
		{SpecTypePress, CategoryGTM},
		{SpecTypeFAQ, CategoryGTM},
		{SpecTypeNarrative1P, CategoryGTM},
		{SpecTypeNarrative6P, CategoryGTM},
		{SpecTypeTRD, CategoryTechnical},
		{SpecTypeIRD, CategoryTechnical},
		{SpecTypeSpec, CategoryOutput},
		{SpecTypeCurrentTruth, CategoryOutput},
		{SpecType("unknown"), ""},
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			got := tt.specType.Category()
			if got != tt.expected {
				t.Errorf("Category() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSpecTypeIsRequired(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected bool
	}{
		{SpecTypeMRD, true},
		{SpecTypePRD, true},
		{SpecTypeUXD, true},
		{SpecTypeTRD, true},
		{SpecTypePress, false},
		{SpecTypeFAQ, false},
		{SpecTypeNarrative1P, false},
		{SpecTypeNarrative6P, false},
		{SpecTypeIRD, false},
		{SpecTypeSpec, false},
		{SpecType("unknown"), false},
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			got := tt.specType.IsRequired()
			if got != tt.expected {
				t.Errorf("IsRequired() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSpecTypeFilename(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected string
	}{
		{SpecTypeMRD, "mrd.md"},
		{SpecTypePRD, "prd.md"},
		{SpecTypeUXD, "uxd.md"},
		{SpecTypePress, "press.md"},
		{SpecTypeTRD, "trd.md"},
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			got := tt.specType.Filename()
			if got != tt.expected {
				t.Errorf("Filename() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSpecTypeEvalFilename(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected string
	}{
		{SpecTypeMRD, "mrd.eval.json"},
		{SpecTypePRD, "prd.eval.json"},
		{SpecTypeTRD, "trd.eval.json"},
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			got := tt.specType.EvalFilename()
			if got != tt.expected {
				t.Errorf("EvalFilename() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSpecTypeDir(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected string
	}{
		{SpecTypeMRD, "source"},
		{SpecTypePRD, "source"},
		{SpecTypeUXD, "source"},
		{SpecTypePress, "gtm"},
		{SpecTypeFAQ, "gtm"},
		{SpecTypeTRD, "technical"},
		{SpecTypeIRD, "technical"},
		{SpecTypeSpec, ""},
		{SpecType("unknown"), ""},
	}

	for _, tt := range tests {
		t.Run(string(tt.specType), func(t *testing.T) {
			got := tt.specType.Dir()
			if got != tt.expected {
				t.Errorf("Dir() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSpecTypeIsValid(t *testing.T) {
	tests := []struct {
		specType SpecType
		expected bool
	}{
		{SpecTypeMRD, true},
		{SpecTypePRD, true},
		{SpecTypeUXD, true},
		{SpecTypePress, true},
		{SpecTypeFAQ, true},
		{SpecTypeTRD, true},
		{SpecTypeSpec, true},
		{SpecType("unknown"), false},
		{SpecType(""), false},
	}

	for _, tt := range tests {
		name := string(tt.specType)
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			got := tt.specType.IsValid()
			if got != tt.expected {
				t.Errorf("IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAllSpecTypes(t *testing.T) {
	all := AllSpecTypes()
	if len(all) != 11 {
		t.Errorf("AllSpecTypes() returned %d types, want 11", len(all))
	}

	// Verify order: source, gtm, technical, output
	categories := make([]SpecCategory, len(all))
	for i, st := range all {
		categories[i] = st.Category()
	}

	// First 3 should be source
	for i := 0; i < 3; i++ {
		if categories[i] != CategorySource {
			t.Errorf("AllSpecTypes()[%d] category = %v, want source", i, categories[i])
		}
	}

	// Next 4 should be gtm
	for i := 3; i < 7; i++ {
		if categories[i] != CategoryGTM {
			t.Errorf("AllSpecTypes()[%d] category = %v, want gtm", i, categories[i])
		}
	}

	// Next 2 should be technical
	for i := 7; i < 9; i++ {
		if categories[i] != CategoryTechnical {
			t.Errorf("AllSpecTypes()[%d] category = %v, want technical", i, categories[i])
		}
	}

	// Last 2 should be output
	for i := 9; i < 11; i++ {
		if categories[i] != CategoryOutput {
			t.Errorf("AllSpecTypes()[%d] category = %v, want output", i, categories[i])
		}
	}
}

func TestSourceSpecTypes(t *testing.T) {
	source := SourceSpecTypes()
	if len(source) != 3 {
		t.Errorf("SourceSpecTypes() returned %d types, want 3", len(source))
	}

	expected := []SpecType{SpecTypeMRD, SpecTypePRD, SpecTypeUXD}
	for i, st := range source {
		if st != expected[i] {
			t.Errorf("SourceSpecTypes()[%d] = %v, want %v", i, st, expected[i])
		}
	}
}

func TestGTMSpecTypes(t *testing.T) {
	gtm := GTMSpecTypes()
	if len(gtm) != 4 {
		t.Errorf("GTMSpecTypes() returned %d types, want 4", len(gtm))
	}

	expected := []SpecType{SpecTypePress, SpecTypeFAQ, SpecTypeNarrative1P, SpecTypeNarrative6P}
	for i, st := range gtm {
		if st != expected[i] {
			t.Errorf("GTMSpecTypes()[%d] = %v, want %v", i, st, expected[i])
		}
	}
}

func TestTechnicalSpecTypes(t *testing.T) {
	tech := TechnicalSpecTypes()
	if len(tech) != 2 {
		t.Errorf("TechnicalSpecTypes() returned %d types, want 2", len(tech))
	}

	expected := []SpecType{SpecTypeTRD, SpecTypeIRD}
	for i, st := range tech {
		if st != expected[i] {
			t.Errorf("TechnicalSpecTypes()[%d] = %v, want %v", i, st, expected[i])
		}
	}
}
