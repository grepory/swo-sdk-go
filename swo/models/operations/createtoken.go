// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/grepory/swo-sdk-go/swo/models/components"
)

// CreateTokenResponseBody - The request has succeeded.
type CreateTokenResponseBody struct {
	Token string `json:"token"`
}

func (o *CreateTokenResponseBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type CreateTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The request has succeeded.
	Object *CreateTokenResponseBody
}

func (o *CreateTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTokenResponse) GetObject() *CreateTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
