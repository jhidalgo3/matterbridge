// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageAssignmentRequestObjectRequestBuilder is request builder for AccessPackageAssignmentRequestObject
type AccessPackageAssignmentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageAssignmentRequestObjectRequest
func (b *AccessPackageAssignmentRequestObjectRequestBuilder) Request() *AccessPackageAssignmentRequestObjectRequest {
	return &AccessPackageAssignmentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageAssignmentRequestObjectRequest is request for AccessPackageAssignmentRequestObject
type AccessPackageAssignmentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageAssignmentRequestObject
func (r *AccessPackageAssignmentRequestObjectRequest) Get(ctx context.Context) (resObj *AccessPackageAssignmentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageAssignmentRequestObject
func (r *AccessPackageAssignmentRequestObjectRequest) Update(ctx context.Context, reqObj *AccessPackageAssignmentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageAssignmentRequestObject
func (r *AccessPackageAssignmentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackage is navigation property
func (b *AccessPackageAssignmentRequestObjectRequestBuilder) AccessPackage() *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackage"
	return bb
}

// AccessPackageAssignment is navigation property
func (b *AccessPackageAssignmentRequestObjectRequestBuilder) AccessPackageAssignment() *AccessPackageAssignmentRequestBuilder {
	bb := &AccessPackageAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignment"
	return bb
}

// Requestor is navigation property
func (b *AccessPackageAssignmentRequestObjectRequestBuilder) Requestor() *AccessPackageSubjectRequestBuilder {
	bb := &AccessPackageSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/requestor"
	return bb
}
