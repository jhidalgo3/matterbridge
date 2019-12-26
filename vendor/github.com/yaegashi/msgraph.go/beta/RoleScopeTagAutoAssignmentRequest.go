// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// RoleScopeTagAutoAssignmentRequestBuilder is request builder for RoleScopeTagAutoAssignment
type RoleScopeTagAutoAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleScopeTagAutoAssignmentRequest
func (b *RoleScopeTagAutoAssignmentRequestBuilder) Request() *RoleScopeTagAutoAssignmentRequest {
	return &RoleScopeTagAutoAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleScopeTagAutoAssignmentRequest is request for RoleScopeTagAutoAssignment
type RoleScopeTagAutoAssignmentRequest struct{ BaseRequest }

// Get performs GET request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Get(ctx context.Context) (resObj *RoleScopeTagAutoAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Update(ctx context.Context, reqObj *RoleScopeTagAutoAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
