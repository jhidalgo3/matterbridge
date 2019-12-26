// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SalesOrderLineRequestBuilder is request builder for SalesOrderLine
type SalesOrderLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesOrderLineRequest
func (b *SalesOrderLineRequestBuilder) Request() *SalesOrderLineRequest {
	return &SalesOrderLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesOrderLineRequest is request for SalesOrderLine
type SalesOrderLineRequest struct{ BaseRequest }

// Get performs GET request for SalesOrderLine
func (r *SalesOrderLineRequest) Get(ctx context.Context) (resObj *SalesOrderLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesOrderLine
func (r *SalesOrderLineRequest) Update(ctx context.Context, reqObj *SalesOrderLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesOrderLine
func (r *SalesOrderLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Account is navigation property
func (b *SalesOrderLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesOrderLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
