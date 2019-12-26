// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidDeviceOwnerCertificateProfileBaseRequestBuilder is request builder for AndroidDeviceOwnerCertificateProfileBase
type AndroidDeviceOwnerCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidDeviceOwnerCertificateProfileBaseRequest
func (b *AndroidDeviceOwnerCertificateProfileBaseRequestBuilder) Request() *AndroidDeviceOwnerCertificateProfileBaseRequest {
	return &AndroidDeviceOwnerCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidDeviceOwnerCertificateProfileBaseRequest is request for AndroidDeviceOwnerCertificateProfileBase
type AndroidDeviceOwnerCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for AndroidDeviceOwnerCertificateProfileBase
func (r *AndroidDeviceOwnerCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *AndroidDeviceOwnerCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidDeviceOwnerCertificateProfileBase
func (r *AndroidDeviceOwnerCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *AndroidDeviceOwnerCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidDeviceOwnerCertificateProfileBase
func (r *AndroidDeviceOwnerCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidDeviceOwnerCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidDeviceOwnerTrustedRootCertificateRequestBuilder {
	bb := &AndroidDeviceOwnerTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
