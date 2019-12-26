// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// EmbeddedSIMDeviceStateRequestBuilder is request builder for EmbeddedSIMDeviceState
type EmbeddedSIMDeviceStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMDeviceStateRequest
func (b *EmbeddedSIMDeviceStateRequestBuilder) Request() *EmbeddedSIMDeviceStateRequest {
	return &EmbeddedSIMDeviceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMDeviceStateRequest is request for EmbeddedSIMDeviceState
type EmbeddedSIMDeviceStateRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Get(ctx context.Context) (resObj *EmbeddedSIMDeviceState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Update(ctx context.Context, reqObj *EmbeddedSIMDeviceState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
