// This file was auto-generated by Fern from our API Definition.

package scopes

import (
	context "context"
	accounting "github.com/merge-api/merge-go-client/accounting"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn More](https://help.merge.dev/en/articles/8828211-common-model-and-field-scopes).
func (c *Client) DefaultScopesRetrieve(ctx context.Context) (*accounting.CommonModelScopeApi, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/default-scopes"

	var response *accounting.CommonModelScopeApi
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn More](https://help.merge.dev/en/articles/8828211-common-model-and-field-scopes).
func (c *Client) LinkedAccountScopesRetrieve(ctx context.Context) (*accounting.CommonModelScopeApi, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/linked-account-scopes"

	var response *accounting.CommonModelScopeApi
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn More](https://help.merge.dev/en/articles/8828211-common-model-and-field-scopes)
func (c *Client) LinkedAccountScopesCreate(ctx context.Context, request *accounting.LinkedAccountCommonModelScopeDeserializerRequest) (*accounting.CommonModelScopeApi, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/linked-account-scopes"

	var response *accounting.CommonModelScopeApi
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
