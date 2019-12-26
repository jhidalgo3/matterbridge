// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnPremisesAgentGroup undocumented
type OnPremisesAgentGroup struct {
	// Entity is the base model of OnPremisesAgentGroup
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// PublishingType undocumented
	PublishingType *OnPremisesPublishingType `json:"publishingType,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// Agents undocumented
	Agents []OnPremisesAgent `json:"agents,omitempty"`
	// PublishedResources undocumented
	PublishedResources []PublishedResource `json:"publishedResources,omitempty"`
}
