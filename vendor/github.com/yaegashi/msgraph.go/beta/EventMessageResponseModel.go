// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EventMessageResponse undocumented
type EventMessageResponse struct {
	// EventMessage is the base model of EventMessageResponse
	EventMessage
	// ProposedNewTime undocumented
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`
	// ResponseType undocumented
	ResponseType *ResponseType `json:"responseType,omitempty"`
}
