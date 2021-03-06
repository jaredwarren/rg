// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// color client HTTP transport
//
// Command:
// $ goa gen github.com/jaredwarren/rg/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the color service endpoint HTTP clients.
type Client struct {
	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Color Doer is the HTTP client used to make requests to the color endpoint.
	ColorDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the color service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UpdateDoer:          doer,
		ColorDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Update returns an endpoint that makes HTTP requests to the color service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("color", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Color returns an endpoint that makes HTTP requests to the color service
// color server.
func (c *Client) Color() goa.Endpoint {
	var (
		decodeResponse = DecodeColorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildColorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ColorDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("color", "color", err)
		}
		return decodeResponse(resp)
	}
}
