// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EapFastConfiguration undocumented
type EapFastConfiguration int

const (
	// EapFastConfigurationVNoProtectedAccessCredential undocumented
	EapFastConfigurationVNoProtectedAccessCredential EapFastConfiguration = 0
	// EapFastConfigurationVUseProtectedAccessCredential undocumented
	EapFastConfigurationVUseProtectedAccessCredential EapFastConfiguration = 1
	// EapFastConfigurationVUseProtectedAccessCredentialAndProvision undocumented
	EapFastConfigurationVUseProtectedAccessCredentialAndProvision EapFastConfiguration = 2
	// EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously undocumented
	EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously EapFastConfiguration = 3
)

// EapFastConfigurationPNoProtectedAccessCredential returns a pointer to EapFastConfigurationVNoProtectedAccessCredential
func EapFastConfigurationPNoProtectedAccessCredential() *EapFastConfiguration {
	v := EapFastConfigurationVNoProtectedAccessCredential
	return &v
}

// EapFastConfigurationPUseProtectedAccessCredential returns a pointer to EapFastConfigurationVUseProtectedAccessCredential
func EapFastConfigurationPUseProtectedAccessCredential() *EapFastConfiguration {
	v := EapFastConfigurationVUseProtectedAccessCredential
	return &v
}

// EapFastConfigurationPUseProtectedAccessCredentialAndProvision returns a pointer to EapFastConfigurationVUseProtectedAccessCredentialAndProvision
func EapFastConfigurationPUseProtectedAccessCredentialAndProvision() *EapFastConfiguration {
	v := EapFastConfigurationVUseProtectedAccessCredentialAndProvision
	return &v
}

// EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously returns a pointer to EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously
func EapFastConfigurationPUseProtectedAccessCredentialAndProvisionAnonymously() *EapFastConfiguration {
	v := EapFastConfigurationVUseProtectedAccessCredentialAndProvisionAnonymously
	return &v
}
