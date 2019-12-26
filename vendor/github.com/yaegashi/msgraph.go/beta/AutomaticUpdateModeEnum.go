// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AutomaticUpdateMode undocumented
type AutomaticUpdateMode int

const (
	// AutomaticUpdateModeVUserDefined undocumented
	AutomaticUpdateModeVUserDefined AutomaticUpdateMode = 0
	// AutomaticUpdateModeVNotifyDownload undocumented
	AutomaticUpdateModeVNotifyDownload AutomaticUpdateMode = 1
	// AutomaticUpdateModeVAutoInstallAtMaintenanceTime undocumented
	AutomaticUpdateModeVAutoInstallAtMaintenanceTime AutomaticUpdateMode = 2
	// AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime undocumented
	AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime AutomaticUpdateMode = 3
	// AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime undocumented
	AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime AutomaticUpdateMode = 4
	// AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl undocumented
	AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl AutomaticUpdateMode = 5
	// AutomaticUpdateModeVWindowsDefault undocumented
	AutomaticUpdateModeVWindowsDefault AutomaticUpdateMode = 6
)

// AutomaticUpdateModePUserDefined returns a pointer to AutomaticUpdateModeVUserDefined
func AutomaticUpdateModePUserDefined() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVUserDefined
	return &v
}

// AutomaticUpdateModePNotifyDownload returns a pointer to AutomaticUpdateModeVNotifyDownload
func AutomaticUpdateModePNotifyDownload() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVNotifyDownload
	return &v
}

// AutomaticUpdateModePAutoInstallAtMaintenanceTime returns a pointer to AutomaticUpdateModeVAutoInstallAtMaintenanceTime
func AutomaticUpdateModePAutoInstallAtMaintenanceTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAtMaintenanceTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootAtMaintenanceTime returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime
func AutomaticUpdateModePAutoInstallAndRebootAtMaintenanceTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootAtScheduledTime returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime
func AutomaticUpdateModePAutoInstallAndRebootAtScheduledTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootWithoutEndUserControl returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl
func AutomaticUpdateModePAutoInstallAndRebootWithoutEndUserControl() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl
	return &v
}

// AutomaticUpdateModePWindowsDefault returns a pointer to AutomaticUpdateModeVWindowsDefault
func AutomaticUpdateModePWindowsDefault() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVWindowsDefault
	return &v
}
