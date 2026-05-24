package types

import "testing"

func TestDefaultSpecConfig(t *testing.T) {
	cfg := DefaultSpecConfig()

	// Check required specs
	requiredSpecs := []string{"mrd", "prd", "uxd", "trd"}
	for _, spec := range requiredSpecs {
		if !cfg.IsRequired(spec) {
			t.Errorf("Expected %s to be required", spec)
		}
	}

	// Check optional specs
	optionalSpecs := []string{"press", "faq", "ird"}
	for _, spec := range optionalSpecs {
		if cfg.IsRequired(spec) {
			t.Errorf("Expected %s to be optional", spec)
		}
	}
}

func TestSpecConfigIsRequired(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("mrd", &SpecRequirement{Required: false})
	cfg.SetRequirement("custom", &SpecRequirement{Required: true})

	tests := []struct {
		specType string
		expected bool
	}{
		{"mrd", false},     // Override built-in
		{"custom", true},   // Custom spec
		{"prd", true},      // Fall back to built-in default
		{"unknown", false}, // Unknown spec, not required
	}

	for _, tt := range tests {
		t.Run(tt.specType, func(t *testing.T) {
			got := cfg.IsRequired(tt.specType)
			if got != tt.expected {
				t.Errorf("IsRequired(%q) = %v, want %v", tt.specType, got, tt.expected)
			}
		})
	}
}

func TestSpecConfigIsRequiredNil(t *testing.T) {
	var cfg *SpecConfig

	// Should fall back to built-in defaults
	if !cfg.IsRequired("mrd") {
		t.Error("Expected mrd to be required with nil config")
	}

	if cfg.IsRequired("press") {
		t.Error("Expected press to be optional with nil config")
	}
}

func TestSpecConfigGetCategory(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("security", &SpecRequirement{
		Required: true,
		Category: CategorySource,
	})

	tests := []struct {
		specType string
		expected SpecCategory
	}{
		{"mrd", CategorySource},      // Built-in
		{"trd", CategoryTechnical},   // Built-in
		{"security", CategorySource}, // Custom
		{"unknown", ""},              // Unknown
	}

	for _, tt := range tests {
		t.Run(tt.specType, func(t *testing.T) {
			got := cfg.GetCategory(tt.specType)
			if got != tt.expected {
				t.Errorf("GetCategory(%q) = %v, want %v", tt.specType, got, tt.expected)
			}
		})
	}
}

func TestSpecConfigGetTemplate(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("prd", &SpecRequirement{
		Template: "company-prd",
	})

	tests := []struct {
		specType string
		expected string
	}{
		{"prd", "company-prd"}, // Custom template
		{"mrd", "mrd"},         // Default to spec type
	}

	for _, tt := range tests {
		t.Run(tt.specType, func(t *testing.T) {
			got := cfg.GetTemplate(tt.specType)
			if got != tt.expected {
				t.Errorf("GetTemplate(%q) = %v, want %v", tt.specType, got, tt.expected)
			}
		})
	}
}

func TestSpecConfigGetRubric(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("prd", &SpecRequirement{
		Rubric: "company-prd-rubric",
	})

	tests := []struct {
		specType string
		expected string
	}{
		{"prd", "company-prd-rubric"}, // Custom rubric
		{"mrd", "mrd"},                // Default to spec type
	}

	for _, tt := range tests {
		t.Run(tt.specType, func(t *testing.T) {
			got := cfg.GetRubric(tt.specType)
			if got != tt.expected {
				t.Errorf("GetRubric(%q) = %v, want %v", tt.specType, got, tt.expected)
			}
		})
	}
}

func TestSpecConfigCustomSpecs(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("security", &SpecRequirement{Required: true})
	cfg.SetRequirement("compliance", &SpecRequirement{Required: true})
	cfg.SetRequirement("mrd", &SpecRequirement{Required: false}) // Built-in, not custom

	custom := cfg.CustomSpecs()

	if len(custom) != 2 {
		t.Errorf("CustomSpecs() returned %d specs, want 2", len(custom))
	}

	// Check that both custom specs are present
	found := make(map[string]bool)
	for _, name := range custom {
		found[name] = true
	}

	if !found["security"] {
		t.Error("Expected 'security' in CustomSpecs()")
	}
	if !found["compliance"] {
		t.Error("Expected 'compliance' in CustomSpecs()")
	}
	if found["mrd"] {
		t.Error("Did not expect 'mrd' in CustomSpecs() (it's built-in)")
	}
}

func TestSpecConfigMerge(t *testing.T) {
	base := DefaultSpecConfig()
	override := NewSpecConfig()
	override.SetRequirement("mrd", &SpecRequirement{Required: false})
	override.SetRequirement("security", &SpecRequirement{Required: true, Category: CategorySource})

	base.Merge(override)

	if base.IsRequired("mrd") {
		t.Error("Expected mrd to be overridden to optional")
	}

	if !base.IsRequired("security") {
		t.Error("Expected security to be added as required")
	}

	// Original values should remain
	if !base.IsRequired("prd") {
		t.Error("Expected prd to remain required")
	}
}

func TestSpecConfigRequiredSpecs(t *testing.T) {
	cfg := NewSpecConfig()
	cfg.SetRequirement("mrd", &SpecRequirement{Required: true})
	cfg.SetRequirement("prd", &SpecRequirement{Required: true})
	cfg.SetRequirement("uxd", &SpecRequirement{Required: false})
	cfg.SetRequirement("security", &SpecRequirement{Required: true})

	required := cfg.RequiredSpecs()

	// Should include mrd, prd, security, and built-in required specs
	found := make(map[string]bool)
	for _, name := range required {
		found[name] = true
	}

	if !found["mrd"] {
		t.Error("Expected 'mrd' in RequiredSpecs()")
	}
	if !found["prd"] {
		t.Error("Expected 'prd' in RequiredSpecs()")
	}
	if !found["security"] {
		t.Error("Expected 'security' in RequiredSpecs()")
	}
	if found["uxd"] {
		t.Error("Did not expect 'uxd' in RequiredSpecs() (set to optional)")
	}
}
