// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidRequiredPasswordType undocumented
type AndroidRequiredPasswordType int

const (
	// AndroidRequiredPasswordTypeVDeviceDefault undocumented
	AndroidRequiredPasswordTypeVDeviceDefault AndroidRequiredPasswordType = 0
	// AndroidRequiredPasswordTypeVAlphabetic undocumented
	AndroidRequiredPasswordTypeVAlphabetic AndroidRequiredPasswordType = 1
	// AndroidRequiredPasswordTypeVAlphanumeric undocumented
	AndroidRequiredPasswordTypeVAlphanumeric AndroidRequiredPasswordType = 2
	// AndroidRequiredPasswordTypeVAlphanumericWithSymbols undocumented
	AndroidRequiredPasswordTypeVAlphanumericWithSymbols AndroidRequiredPasswordType = 3
	// AndroidRequiredPasswordTypeVLowSecurityBiometric undocumented
	AndroidRequiredPasswordTypeVLowSecurityBiometric AndroidRequiredPasswordType = 4
	// AndroidRequiredPasswordTypeVNumeric undocumented
	AndroidRequiredPasswordTypeVNumeric AndroidRequiredPasswordType = 5
	// AndroidRequiredPasswordTypeVNumericComplex undocumented
	AndroidRequiredPasswordTypeVNumericComplex AndroidRequiredPasswordType = 6
	// AndroidRequiredPasswordTypeVAny undocumented
	AndroidRequiredPasswordTypeVAny AndroidRequiredPasswordType = 7
)

// AndroidRequiredPasswordTypePDeviceDefault returns a pointer to AndroidRequiredPasswordTypeVDeviceDefault
func AndroidRequiredPasswordTypePDeviceDefault() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVDeviceDefault
	return &v
}

// AndroidRequiredPasswordTypePAlphabetic returns a pointer to AndroidRequiredPasswordTypeVAlphabetic
func AndroidRequiredPasswordTypePAlphabetic() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVAlphabetic
	return &v
}

// AndroidRequiredPasswordTypePAlphanumeric returns a pointer to AndroidRequiredPasswordTypeVAlphanumeric
func AndroidRequiredPasswordTypePAlphanumeric() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVAlphanumeric
	return &v
}

// AndroidRequiredPasswordTypePAlphanumericWithSymbols returns a pointer to AndroidRequiredPasswordTypeVAlphanumericWithSymbols
func AndroidRequiredPasswordTypePAlphanumericWithSymbols() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVAlphanumericWithSymbols
	return &v
}

// AndroidRequiredPasswordTypePLowSecurityBiometric returns a pointer to AndroidRequiredPasswordTypeVLowSecurityBiometric
func AndroidRequiredPasswordTypePLowSecurityBiometric() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVLowSecurityBiometric
	return &v
}

// AndroidRequiredPasswordTypePNumeric returns a pointer to AndroidRequiredPasswordTypeVNumeric
func AndroidRequiredPasswordTypePNumeric() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVNumeric
	return &v
}

// AndroidRequiredPasswordTypePNumericComplex returns a pointer to AndroidRequiredPasswordTypeVNumericComplex
func AndroidRequiredPasswordTypePNumericComplex() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVNumericComplex
	return &v
}

// AndroidRequiredPasswordTypePAny returns a pointer to AndroidRequiredPasswordTypeVAny
func AndroidRequiredPasswordTypePAny() *AndroidRequiredPasswordType {
	v := AndroidRequiredPasswordTypeVAny
	return &v
}
