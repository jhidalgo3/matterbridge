// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10VpnAuthenticationMethod undocumented
type Windows10VpnAuthenticationMethod int

const (
	// Windows10VpnAuthenticationMethodVCertificate undocumented
	Windows10VpnAuthenticationMethodVCertificate Windows10VpnAuthenticationMethod = 0
	// Windows10VpnAuthenticationMethodVUsernameAndPassword undocumented
	Windows10VpnAuthenticationMethodVUsernameAndPassword Windows10VpnAuthenticationMethod = 1
	// Windows10VpnAuthenticationMethodVCustomEapXML undocumented
	Windows10VpnAuthenticationMethodVCustomEapXML Windows10VpnAuthenticationMethod = 2
)

// Windows10VpnAuthenticationMethodPCertificate returns a pointer to Windows10VpnAuthenticationMethodVCertificate
func Windows10VpnAuthenticationMethodPCertificate() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVCertificate
	return &v
}

// Windows10VpnAuthenticationMethodPUsernameAndPassword returns a pointer to Windows10VpnAuthenticationMethodVUsernameAndPassword
func Windows10VpnAuthenticationMethodPUsernameAndPassword() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVUsernameAndPassword
	return &v
}

// Windows10VpnAuthenticationMethodPCustomEapXML returns a pointer to Windows10VpnAuthenticationMethodVCustomEapXML
func Windows10VpnAuthenticationMethodPCustomEapXML() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVCustomEapXML
	return &v
}
