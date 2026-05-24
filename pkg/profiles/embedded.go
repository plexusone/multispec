package profiles

import "embed"

// Embed all default profiles
//
//go:embed default/0-1/profile.yaml default/0-1/templates/*.md default/0-1/rubrics/*.rubric.yaml
//go:embed default/startup/profile.yaml default/startup/templates/*.md default/startup/rubrics/*.rubric.yaml
//go:embed default/growth/profile.yaml default/growth/templates/*.md default/growth/rubrics/*.rubric.yaml
//go:embed default/enterprise/profile.yaml default/enterprise/templates/*.md default/enterprise/rubrics/*.rubric.yaml
var defaultProfiles embed.FS

// DefaultLoader returns a loader for built-in default profiles.
func DefaultLoader() Loader {
	return NewResolvingLoader(NewEmbedFSLoader(defaultProfiles, "default"))
}

// DefaultProfileNames returns the names of all default profiles.
var DefaultProfileNames = []string{"0-1", "startup", "growth", "enterprise"}

// IsDefaultProfile returns true if the name is a default profile.
func IsDefaultProfile(name string) bool {
	for _, n := range DefaultProfileNames {
		if n == name {
			return true
		}
	}
	return false
}
