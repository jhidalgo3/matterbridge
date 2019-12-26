// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartFontRequestBuilder is request builder for WorkbookChartFont
type WorkbookChartFontRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartFontRequest
func (b *WorkbookChartFontRequestBuilder) Request() *WorkbookChartFontRequest {
	return &WorkbookChartFontRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartFontRequest is request for WorkbookChartFont
type WorkbookChartFontRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChartFont
func (r *WorkbookChartFontRequest) Get(ctx context.Context) (resObj *WorkbookChartFont, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartFont
func (r *WorkbookChartFontRequest) Update(ctx context.Context, reqObj *WorkbookChartFont) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartFont
func (r *WorkbookChartFontRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
