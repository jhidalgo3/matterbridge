// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessApplications undocumented
type ConditionalAccessApplications struct {
	// Object is the base model of ConditionalAccessApplications
	Object
	// IncludeApplications undocumented
	IncludeApplications []string `json:"includeApplications,omitempty"`
	// ExcludeApplications undocumented
	ExcludeApplications []string `json:"excludeApplications,omitempty"`
	// IncludeUserActions undocumented
	IncludeUserActions []string `json:"includeUserActions,omitempty"`
}
