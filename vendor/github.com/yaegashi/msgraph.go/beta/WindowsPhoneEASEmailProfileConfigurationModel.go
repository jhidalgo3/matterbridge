// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhoneEASEmailProfileConfiguration By providing configurations in this profile you can instruct the native email client on Windows Phone to communicate with an Exchange server and get email, contacts, calendar, and tasks. Furthermore, you can also specify how much email to sync and how often the device should sync.
type WindowsPhoneEASEmailProfileConfiguration struct {
	// EasEmailProfileConfigurationBase is the base model of WindowsPhoneEASEmailProfileConfiguration
	EasEmailProfileConfigurationBase
	// AccountName Account name.
	AccountName *string `json:"accountName,omitempty"`
	// ApplyOnlyToWindowsPhone81 Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindowsPhone81 *bool `json:"applyOnlyToWindowsPhone81,omitempty"`
	// SyncCalendar Whether or not to sync the calendar.
	SyncCalendar *bool `json:"syncCalendar,omitempty"`
	// SyncContacts Whether or not to sync contacts.
	SyncContacts *bool `json:"syncContacts,omitempty"`
	// SyncTasks Whether or not to sync tasks.
	SyncTasks *bool `json:"syncTasks,omitempty"`
	// DurationOfEmailToSync Duration of email to sync.
	DurationOfEmailToSync *EmailSyncDuration `json:"durationOfEmailToSync,omitempty"`
	// EmailAddressSource Email attribute that is picked from AAD and injected into this profile before installing on the device.
	EmailAddressSource *UserEmailSource `json:"emailAddressSource,omitempty"`
	// EmailSyncSchedule Email sync schedule.
	EmailSyncSchedule *EmailSyncSchedule `json:"emailSyncSchedule,omitempty"`
	// HostName Exchange location that (URL) that the native mail app connects to.
	HostName *string `json:"hostName,omitempty"`
	// RequireSsl Indicates whether or not to use SSL.
	RequireSsl *bool `json:"requireSsl,omitempty"`
}
