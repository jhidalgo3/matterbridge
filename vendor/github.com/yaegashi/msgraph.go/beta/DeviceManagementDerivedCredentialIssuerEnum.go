// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementDerivedCredentialIssuer undocumented
type DeviceManagementDerivedCredentialIssuer int

const (
	// DeviceManagementDerivedCredentialIssuerVIntercede undocumented
	DeviceManagementDerivedCredentialIssuerVIntercede DeviceManagementDerivedCredentialIssuer = 0
	// DeviceManagementDerivedCredentialIssuerVEntrustDatacard undocumented
	DeviceManagementDerivedCredentialIssuerVEntrustDatacard DeviceManagementDerivedCredentialIssuer = 1
	// DeviceManagementDerivedCredentialIssuerVPurebred undocumented
	DeviceManagementDerivedCredentialIssuerVPurebred DeviceManagementDerivedCredentialIssuer = 2
)

// DeviceManagementDerivedCredentialIssuerPIntercede returns a pointer to DeviceManagementDerivedCredentialIssuerVIntercede
func DeviceManagementDerivedCredentialIssuerPIntercede() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVIntercede
	return &v
}

// DeviceManagementDerivedCredentialIssuerPEntrustDatacard returns a pointer to DeviceManagementDerivedCredentialIssuerVEntrustDatacard
func DeviceManagementDerivedCredentialIssuerPEntrustDatacard() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVEntrustDatacard
	return &v
}

// DeviceManagementDerivedCredentialIssuerPPurebred returns a pointer to DeviceManagementDerivedCredentialIssuerVPurebred
func DeviceManagementDerivedCredentialIssuerPPurebred() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVPurebred
	return &v
}
