// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PermissionClassificationType undocumented
type PermissionClassificationType int

const (
	// PermissionClassificationTypeVLow undocumented
	PermissionClassificationTypeVLow PermissionClassificationType = 1
	// PermissionClassificationTypeVMedium undocumented
	PermissionClassificationTypeVMedium PermissionClassificationType = 2
	// PermissionClassificationTypeVHigh undocumented
	PermissionClassificationTypeVHigh PermissionClassificationType = 3
	// PermissionClassificationTypeVUnknownFutureValue undocumented
	PermissionClassificationTypeVUnknownFutureValue PermissionClassificationType = 4
)

// PermissionClassificationTypePLow returns a pointer to PermissionClassificationTypeVLow
func PermissionClassificationTypePLow() *PermissionClassificationType {
	v := PermissionClassificationTypeVLow
	return &v
}

// PermissionClassificationTypePMedium returns a pointer to PermissionClassificationTypeVMedium
func PermissionClassificationTypePMedium() *PermissionClassificationType {
	v := PermissionClassificationTypeVMedium
	return &v
}

// PermissionClassificationTypePHigh returns a pointer to PermissionClassificationTypeVHigh
func PermissionClassificationTypePHigh() *PermissionClassificationType {
	v := PermissionClassificationTypeVHigh
	return &v
}

// PermissionClassificationTypePUnknownFutureValue returns a pointer to PermissionClassificationTypeVUnknownFutureValue
func PermissionClassificationTypePUnknownFutureValue() *PermissionClassificationType {
	v := PermissionClassificationTypeVUnknownFutureValue
	return &v
}
