// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserEmailSource undocumented
type UserEmailSource int

const (
	// UserEmailSourceVUserPrincipalName undocumented
	UserEmailSourceVUserPrincipalName UserEmailSource = 0
	// UserEmailSourceVPrimarySMTPAddress undocumented
	UserEmailSourceVPrimarySMTPAddress UserEmailSource = 1
)

// UserEmailSourcePUserPrincipalName returns a pointer to UserEmailSourceVUserPrincipalName
func UserEmailSourcePUserPrincipalName() *UserEmailSource {
	v := UserEmailSourceVUserPrincipalName
	return &v
}

// UserEmailSourcePPrimarySMTPAddress returns a pointer to UserEmailSourceVPrimarySMTPAddress
func UserEmailSourcePPrimarySMTPAddress() *UserEmailSource {
	v := UserEmailSourceVPrimarySMTPAddress
	return &v
}
