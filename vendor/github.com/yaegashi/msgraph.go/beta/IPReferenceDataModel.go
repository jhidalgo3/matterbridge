// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IPReferenceData undocumented
type IPReferenceData struct {
	// Object is the base model of IPReferenceData
	Object
	// Asn undocumented
	Asn *int `json:"asn,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// CountryOrRegionCode undocumented
	CountryOrRegionCode *string `json:"countryOrRegionCode,omitempty"`
	// Organization undocumented
	Organization *string `json:"organization,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}
