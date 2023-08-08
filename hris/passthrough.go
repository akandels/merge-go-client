// This file was auto-generated by Fern from our API Definition.

package hris

type DataPassthroughRequest struct {
	Method MethodEnum `json:"method,omitempty"`
	// <span style="white-space: nowrap">`non-empty`</span>
	Path string `json:"path"`
	// <span style="white-space: nowrap">`non-empty`</span>
	BaseUrlOverride *string `json:"base_url_override,omitempty"`
	// <span style="white-space: nowrap">`non-empty`</span>
	Data *string `json:"data,omitempty"`
	// Pass an array of `MultipartFormField` objects in here instead of using the `data` param if `request_format` is set to `MULTIPART`.
	MultipartFormData []*MultipartFormFieldRequest `json:"multipart_form_data,omitempty"`
	// The headers to use for the request (Merge will handle the account's authorization headers). `Content-Type` header is required for passthrough. Choose content type corresponding to expected format of receiving server.
	Headers       map[string]any     `json:"headers,omitempty"`
	RequestFormat *RequestFormatEnum `json:"request_format,omitempty"`
	// Optional. If true, the response will always be an object of the form `{"type": T, "value": ...}` where `T` will be one of `string, boolean, number, null, array, object`.
	NormalizeResponse *bool `json:"normalize_response,omitempty"`
}
