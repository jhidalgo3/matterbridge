// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ThreatAssessmentResult undocumented
type ThreatAssessmentResult struct {
	// Entity is the base model of ThreatAssessmentResult
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ResultType undocumented
	ResultType *ThreatAssessmentResultType `json:"resultType,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}
