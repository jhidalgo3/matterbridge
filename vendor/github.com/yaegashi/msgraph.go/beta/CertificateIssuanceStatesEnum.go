// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CertificateIssuanceStates undocumented
type CertificateIssuanceStates int

const (
	// CertificateIssuanceStatesVUnknown undocumented
	CertificateIssuanceStatesVUnknown CertificateIssuanceStates = 0
	// CertificateIssuanceStatesVChallengeIssued undocumented
	CertificateIssuanceStatesVChallengeIssued CertificateIssuanceStates = 1
	// CertificateIssuanceStatesVChallengeIssueFailed undocumented
	CertificateIssuanceStatesVChallengeIssueFailed CertificateIssuanceStates = 2
	// CertificateIssuanceStatesVRequestCreationFailed undocumented
	CertificateIssuanceStatesVRequestCreationFailed CertificateIssuanceStates = 3
	// CertificateIssuanceStatesVRequestSubmitFailed undocumented
	CertificateIssuanceStatesVRequestSubmitFailed CertificateIssuanceStates = 4
	// CertificateIssuanceStatesVChallengeValidationSucceeded undocumented
	CertificateIssuanceStatesVChallengeValidationSucceeded CertificateIssuanceStates = 5
	// CertificateIssuanceStatesVChallengeValidationFailed undocumented
	CertificateIssuanceStatesVChallengeValidationFailed CertificateIssuanceStates = 6
	// CertificateIssuanceStatesVIssueFailed undocumented
	CertificateIssuanceStatesVIssueFailed CertificateIssuanceStates = 7
	// CertificateIssuanceStatesVIssuePending undocumented
	CertificateIssuanceStatesVIssuePending CertificateIssuanceStates = 8
	// CertificateIssuanceStatesVIssued undocumented
	CertificateIssuanceStatesVIssued CertificateIssuanceStates = 9
	// CertificateIssuanceStatesVResponseProcessingFailed undocumented
	CertificateIssuanceStatesVResponseProcessingFailed CertificateIssuanceStates = 10
	// CertificateIssuanceStatesVResponsePending undocumented
	CertificateIssuanceStatesVResponsePending CertificateIssuanceStates = 11
	// CertificateIssuanceStatesVEnrollmentSucceeded undocumented
	CertificateIssuanceStatesVEnrollmentSucceeded CertificateIssuanceStates = 12
	// CertificateIssuanceStatesVEnrollmentNotNeeded undocumented
	CertificateIssuanceStatesVEnrollmentNotNeeded CertificateIssuanceStates = 13
	// CertificateIssuanceStatesVRevoked undocumented
	CertificateIssuanceStatesVRevoked CertificateIssuanceStates = 14
	// CertificateIssuanceStatesVRemovedFromCollection undocumented
	CertificateIssuanceStatesVRemovedFromCollection CertificateIssuanceStates = 15
	// CertificateIssuanceStatesVRenewVerified undocumented
	CertificateIssuanceStatesVRenewVerified CertificateIssuanceStates = 16
	// CertificateIssuanceStatesVInstallFailed undocumented
	CertificateIssuanceStatesVInstallFailed CertificateIssuanceStates = 17
	// CertificateIssuanceStatesVInstalled undocumented
	CertificateIssuanceStatesVInstalled CertificateIssuanceStates = 18
	// CertificateIssuanceStatesVDeleteFailed undocumented
	CertificateIssuanceStatesVDeleteFailed CertificateIssuanceStates = 19
	// CertificateIssuanceStatesVDeleted undocumented
	CertificateIssuanceStatesVDeleted CertificateIssuanceStates = 20
	// CertificateIssuanceStatesVRenewalRequested undocumented
	CertificateIssuanceStatesVRenewalRequested CertificateIssuanceStates = 21
	// CertificateIssuanceStatesVRequested undocumented
	CertificateIssuanceStatesVRequested CertificateIssuanceStates = 22
)

// CertificateIssuanceStatesPUnknown returns a pointer to CertificateIssuanceStatesVUnknown
func CertificateIssuanceStatesPUnknown() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVUnknown
	return &v
}

// CertificateIssuanceStatesPChallengeIssued returns a pointer to CertificateIssuanceStatesVChallengeIssued
func CertificateIssuanceStatesPChallengeIssued() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVChallengeIssued
	return &v
}

// CertificateIssuanceStatesPChallengeIssueFailed returns a pointer to CertificateIssuanceStatesVChallengeIssueFailed
func CertificateIssuanceStatesPChallengeIssueFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVChallengeIssueFailed
	return &v
}

// CertificateIssuanceStatesPRequestCreationFailed returns a pointer to CertificateIssuanceStatesVRequestCreationFailed
func CertificateIssuanceStatesPRequestCreationFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRequestCreationFailed
	return &v
}

// CertificateIssuanceStatesPRequestSubmitFailed returns a pointer to CertificateIssuanceStatesVRequestSubmitFailed
func CertificateIssuanceStatesPRequestSubmitFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRequestSubmitFailed
	return &v
}

// CertificateIssuanceStatesPChallengeValidationSucceeded returns a pointer to CertificateIssuanceStatesVChallengeValidationSucceeded
func CertificateIssuanceStatesPChallengeValidationSucceeded() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVChallengeValidationSucceeded
	return &v
}

// CertificateIssuanceStatesPChallengeValidationFailed returns a pointer to CertificateIssuanceStatesVChallengeValidationFailed
func CertificateIssuanceStatesPChallengeValidationFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVChallengeValidationFailed
	return &v
}

// CertificateIssuanceStatesPIssueFailed returns a pointer to CertificateIssuanceStatesVIssueFailed
func CertificateIssuanceStatesPIssueFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVIssueFailed
	return &v
}

// CertificateIssuanceStatesPIssuePending returns a pointer to CertificateIssuanceStatesVIssuePending
func CertificateIssuanceStatesPIssuePending() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVIssuePending
	return &v
}

// CertificateIssuanceStatesPIssued returns a pointer to CertificateIssuanceStatesVIssued
func CertificateIssuanceStatesPIssued() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVIssued
	return &v
}

// CertificateIssuanceStatesPResponseProcessingFailed returns a pointer to CertificateIssuanceStatesVResponseProcessingFailed
func CertificateIssuanceStatesPResponseProcessingFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVResponseProcessingFailed
	return &v
}

// CertificateIssuanceStatesPResponsePending returns a pointer to CertificateIssuanceStatesVResponsePending
func CertificateIssuanceStatesPResponsePending() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVResponsePending
	return &v
}

// CertificateIssuanceStatesPEnrollmentSucceeded returns a pointer to CertificateIssuanceStatesVEnrollmentSucceeded
func CertificateIssuanceStatesPEnrollmentSucceeded() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVEnrollmentSucceeded
	return &v
}

// CertificateIssuanceStatesPEnrollmentNotNeeded returns a pointer to CertificateIssuanceStatesVEnrollmentNotNeeded
func CertificateIssuanceStatesPEnrollmentNotNeeded() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVEnrollmentNotNeeded
	return &v
}

// CertificateIssuanceStatesPRevoked returns a pointer to CertificateIssuanceStatesVRevoked
func CertificateIssuanceStatesPRevoked() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRevoked
	return &v
}

// CertificateIssuanceStatesPRemovedFromCollection returns a pointer to CertificateIssuanceStatesVRemovedFromCollection
func CertificateIssuanceStatesPRemovedFromCollection() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRemovedFromCollection
	return &v
}

// CertificateIssuanceStatesPRenewVerified returns a pointer to CertificateIssuanceStatesVRenewVerified
func CertificateIssuanceStatesPRenewVerified() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRenewVerified
	return &v
}

// CertificateIssuanceStatesPInstallFailed returns a pointer to CertificateIssuanceStatesVInstallFailed
func CertificateIssuanceStatesPInstallFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVInstallFailed
	return &v
}

// CertificateIssuanceStatesPInstalled returns a pointer to CertificateIssuanceStatesVInstalled
func CertificateIssuanceStatesPInstalled() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVInstalled
	return &v
}

// CertificateIssuanceStatesPDeleteFailed returns a pointer to CertificateIssuanceStatesVDeleteFailed
func CertificateIssuanceStatesPDeleteFailed() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVDeleteFailed
	return &v
}

// CertificateIssuanceStatesPDeleted returns a pointer to CertificateIssuanceStatesVDeleted
func CertificateIssuanceStatesPDeleted() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVDeleted
	return &v
}

// CertificateIssuanceStatesPRenewalRequested returns a pointer to CertificateIssuanceStatesVRenewalRequested
func CertificateIssuanceStatesPRenewalRequested() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRenewalRequested
	return &v
}

// CertificateIssuanceStatesPRequested returns a pointer to CertificateIssuanceStatesVRequested
func CertificateIssuanceStatesPRequested() *CertificateIssuanceStates {
	v := CertificateIssuanceStatesVRequested
	return &v
}
