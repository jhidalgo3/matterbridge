// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// BookingStaffMemberRequestBuilder is request builder for BookingStaffMember
type BookingStaffMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingStaffMemberRequest
func (b *BookingStaffMemberRequestBuilder) Request() *BookingStaffMemberRequest {
	return &BookingStaffMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingStaffMemberRequest is request for BookingStaffMember
type BookingStaffMemberRequest struct{ BaseRequest }

// Get performs GET request for BookingStaffMember
func (r *BookingStaffMemberRequest) Get(ctx context.Context) (resObj *BookingStaffMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingStaffMember
func (r *BookingStaffMemberRequest) Update(ctx context.Context, reqObj *BookingStaffMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingStaffMember
func (r *BookingStaffMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
