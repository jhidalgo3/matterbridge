// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DefenderSubmitSamplesConsentType undocumented
type DefenderSubmitSamplesConsentType int

const (
	// DefenderSubmitSamplesConsentTypeVSendSafeSamplesAutomatically undocumented
	DefenderSubmitSamplesConsentTypeVSendSafeSamplesAutomatically DefenderSubmitSamplesConsentType = 0
	// DefenderSubmitSamplesConsentTypeVAlwaysPrompt undocumented
	DefenderSubmitSamplesConsentTypeVAlwaysPrompt DefenderSubmitSamplesConsentType = 1
	// DefenderSubmitSamplesConsentTypeVNeverSend undocumented
	DefenderSubmitSamplesConsentTypeVNeverSend DefenderSubmitSamplesConsentType = 2
	// DefenderSubmitSamplesConsentTypeVSendAllSamplesAutomatically undocumented
	DefenderSubmitSamplesConsentTypeVSendAllSamplesAutomatically DefenderSubmitSamplesConsentType = 3
)

// DefenderSubmitSamplesConsentTypePSendSafeSamplesAutomatically returns a pointer to DefenderSubmitSamplesConsentTypeVSendSafeSamplesAutomatically
func DefenderSubmitSamplesConsentTypePSendSafeSamplesAutomatically() *DefenderSubmitSamplesConsentType {
	v := DefenderSubmitSamplesConsentTypeVSendSafeSamplesAutomatically
	return &v
}

// DefenderSubmitSamplesConsentTypePAlwaysPrompt returns a pointer to DefenderSubmitSamplesConsentTypeVAlwaysPrompt
func DefenderSubmitSamplesConsentTypePAlwaysPrompt() *DefenderSubmitSamplesConsentType {
	v := DefenderSubmitSamplesConsentTypeVAlwaysPrompt
	return &v
}

// DefenderSubmitSamplesConsentTypePNeverSend returns a pointer to DefenderSubmitSamplesConsentTypeVNeverSend
func DefenderSubmitSamplesConsentTypePNeverSend() *DefenderSubmitSamplesConsentType {
	v := DefenderSubmitSamplesConsentTypeVNeverSend
	return &v
}

// DefenderSubmitSamplesConsentTypePSendAllSamplesAutomatically returns a pointer to DefenderSubmitSamplesConsentTypeVSendAllSamplesAutomatically
func DefenderSubmitSamplesConsentTypePSendAllSamplesAutomatically() *DefenderSubmitSamplesConsentType {
	v := DefenderSubmitSamplesConsentTypeVSendAllSamplesAutomatically
	return &v
}
