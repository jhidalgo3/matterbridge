// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SafeSearchFilterType undocumented
type SafeSearchFilterType int

const (
	// SafeSearchFilterTypeVUserDefined undocumented
	SafeSearchFilterTypeVUserDefined SafeSearchFilterType = 0
	// SafeSearchFilterTypeVStrict undocumented
	SafeSearchFilterTypeVStrict SafeSearchFilterType = 1
	// SafeSearchFilterTypeVModerate undocumented
	SafeSearchFilterTypeVModerate SafeSearchFilterType = 2
)

// SafeSearchFilterTypePUserDefined returns a pointer to SafeSearchFilterTypeVUserDefined
func SafeSearchFilterTypePUserDefined() *SafeSearchFilterType {
	v := SafeSearchFilterTypeVUserDefined
	return &v
}

// SafeSearchFilterTypePStrict returns a pointer to SafeSearchFilterTypeVStrict
func SafeSearchFilterTypePStrict() *SafeSearchFilterType {
	v := SafeSearchFilterTypeVStrict
	return &v
}

// SafeSearchFilterTypePModerate returns a pointer to SafeSearchFilterTypeVModerate
func SafeSearchFilterTypePModerate() *SafeSearchFilterType {
	v := SafeSearchFilterTypeVModerate
	return &v
}
