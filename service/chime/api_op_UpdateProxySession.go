// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateProxySessionInput struct {
	_ struct{} `type:"structure"`

	// The proxy session capabilities.
	//
	// Capabilities is a required field
	Capabilities []Capability `type:"list" required:"true"`

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int64 `min:"1" type:"integer"`

	// The proxy session ID.
	//
	// ProxySessionId is a required field
	ProxySessionId *string `location:"uri" locationName:"proxySessionId" min:"1" type:"string" required:"true"`

	// The Amazon Chime voice connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateProxySessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProxySessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProxySessionInput"}

	if s.Capabilities == nil {
		invalidParams.Add(aws.NewErrParamRequired("Capabilities"))
	}
	if s.ExpiryMinutes != nil && *s.ExpiryMinutes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ExpiryMinutes", 1))
	}

	if s.ProxySessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProxySessionId"))
	}
	if s.ProxySessionId != nil && len(*s.ProxySessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProxySessionId", 1))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.VoiceConnectorId != nil && len(*s.VoiceConnectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VoiceConnectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProxySessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Capabilities != nil {
		v := s.Capabilities

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Capabilities", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ExpiryMinutes != nil {
		v := *s.ExpiryMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExpiryMinutes", protocol.Int64Value(v), metadata)
	}
	if s.ProxySessionId != nil {
		v := *s.ProxySessionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "proxySessionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateProxySessionOutput struct {
	_ struct{} `type:"structure"`

	// The proxy session details.
	ProxySession *ProxySession `type:"structure"`
}

// String returns the string representation
func (s UpdateProxySessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateProxySessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProxySession != nil {
		v := s.ProxySession

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ProxySession", v, metadata)
	}
	return nil
}

const opUpdateProxySession = "UpdateProxySession"

// UpdateProxySessionRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the specified proxy session details, such as voice or SMS capabilities.
//
//    // Example sending a request using UpdateProxySessionRequest.
//    req := client.UpdateProxySessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateProxySession
func (c *Client) UpdateProxySessionRequest(input *UpdateProxySessionInput) UpdateProxySessionRequest {
	op := &aws.Operation{
		Name:       opUpdateProxySession,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/proxy-sessions/{proxySessionId}",
	}

	if input == nil {
		input = &UpdateProxySessionInput{}
	}

	req := c.newRequest(op, input, &UpdateProxySessionOutput{})

	return UpdateProxySessionRequest{Request: req, Input: input, Copy: c.UpdateProxySessionRequest}
}

// UpdateProxySessionRequest is the request type for the
// UpdateProxySession API operation.
type UpdateProxySessionRequest struct {
	*aws.Request
	Input *UpdateProxySessionInput
	Copy  func(*UpdateProxySessionInput) UpdateProxySessionRequest
}

// Send marshals and sends the UpdateProxySession API request.
func (r UpdateProxySessionRequest) Send(ctx context.Context) (*UpdateProxySessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProxySessionResponse{
		UpdateProxySessionOutput: r.Request.Data.(*UpdateProxySessionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProxySessionResponse is the response type for the
// UpdateProxySession API operation.
type UpdateProxySessionResponse struct {
	*UpdateProxySessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProxySession request.
func (r *UpdateProxySessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
