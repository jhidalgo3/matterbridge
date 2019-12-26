// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeDeviceClass undocumented
type DeviceManagementExchangeDeviceClass struct {
	// Object is the base model of DeviceManagementExchangeDeviceClass
	Object
	// Name Name of the device class which will be impacted by this rule.
	Name *string `json:"name,omitempty"`
	// Type Type of device which is impacted by this rule e.g. Model, Family
	Type *DeviceManagementExchangeAccessRuleType `json:"type,omitempty"`
}
