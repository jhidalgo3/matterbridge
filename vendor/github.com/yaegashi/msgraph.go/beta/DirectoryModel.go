// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Directory undocumented
type Directory struct {
	// Entity is the base model of Directory
	Entity
	// DeletedItems undocumented
	DeletedItems []DirectoryObject `json:"deletedItems,omitempty"`
	// FeatureRolloutPolicies undocumented
	FeatureRolloutPolicies []FeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`
}
