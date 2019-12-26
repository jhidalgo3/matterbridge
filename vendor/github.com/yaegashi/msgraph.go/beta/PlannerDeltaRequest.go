// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PlannerDeltaRequestBuilder is request builder for PlannerDelta
type PlannerDeltaRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerDeltaRequest
func (b *PlannerDeltaRequestBuilder) Request() *PlannerDeltaRequest {
	return &PlannerDeltaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerDeltaRequest is request for PlannerDelta
type PlannerDeltaRequest struct{ BaseRequest }

// Get performs GET request for PlannerDelta
func (r *PlannerDeltaRequest) Get(ctx context.Context) (resObj *PlannerDelta, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerDelta
func (r *PlannerDeltaRequest) Update(ctx context.Context, reqObj *PlannerDelta) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerDelta
func (r *PlannerDeltaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
