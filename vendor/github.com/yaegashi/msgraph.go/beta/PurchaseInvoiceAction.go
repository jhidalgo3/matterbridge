// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PurchaseInvoicePostRequestParameter undocumented
type PurchaseInvoicePostRequestParameter struct {
}

//
type PurchaseInvoicePostRequestBuilder struct{ BaseRequestBuilder }

// Post action undocumented
func (b *PurchaseInvoiceRequestBuilder) Post(reqObj *PurchaseInvoicePostRequestParameter) *PurchaseInvoicePostRequestBuilder {
	bb := &PurchaseInvoicePostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/post"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PurchaseInvoicePostRequest struct{ BaseRequest }

//
func (b *PurchaseInvoicePostRequestBuilder) Request() *PurchaseInvoicePostRequest {
	return &PurchaseInvoicePostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PurchaseInvoicePostRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
