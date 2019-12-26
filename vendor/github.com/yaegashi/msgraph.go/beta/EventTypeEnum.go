// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EventType undocumented
type EventType int

const (
	// EventTypeVSingleInstance undocumented
	EventTypeVSingleInstance EventType = 0
	// EventTypeVOccurrence undocumented
	EventTypeVOccurrence EventType = 1
	// EventTypeVException undocumented
	EventTypeVException EventType = 2
	// EventTypeVSeriesMaster undocumented
	EventTypeVSeriesMaster EventType = 3
)

// EventTypePSingleInstance returns a pointer to EventTypeVSingleInstance
func EventTypePSingleInstance() *EventType {
	v := EventTypeVSingleInstance
	return &v
}

// EventTypePOccurrence returns a pointer to EventTypeVOccurrence
func EventTypePOccurrence() *EventType {
	v := EventTypeVOccurrence
	return &v
}

// EventTypePException returns a pointer to EventTypeVException
func EventTypePException() *EventType {
	v := EventTypeVException
	return &v
}

// EventTypePSeriesMaster returns a pointer to EventTypeVSeriesMaster
func EventTypePSeriesMaster() *EventType {
	v := EventTypeVSeriesMaster
	return &v
}
