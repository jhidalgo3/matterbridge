// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PurchaseInvoiceLineRequestBuilder is request builder for PurchaseInvoiceLine
type PurchaseInvoiceLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns PurchaseInvoiceLineRequest
func (b *PurchaseInvoiceLineRequestBuilder) Request() *PurchaseInvoiceLineRequest {
	return &PurchaseInvoiceLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PurchaseInvoiceLineRequest is request for PurchaseInvoiceLine
type PurchaseInvoiceLineRequest struct{ BaseRequest }

// Get performs GET request for PurchaseInvoiceLine
func (r *PurchaseInvoiceLineRequest) Get(ctx context.Context) (resObj *PurchaseInvoiceLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PurchaseInvoiceLine
func (r *PurchaseInvoiceLineRequest) Update(ctx context.Context, reqObj *PurchaseInvoiceLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PurchaseInvoiceLine
func (r *PurchaseInvoiceLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Account is navigation property
func (b *PurchaseInvoiceLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *PurchaseInvoiceLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
