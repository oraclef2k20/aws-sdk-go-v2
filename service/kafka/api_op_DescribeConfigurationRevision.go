// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeConfigurationRevisionInput struct {
	_ struct{} `type:"structure"`

	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`

	// Revision is a required field
	Revision *int64 `location:"uri" locationName:"revision" type:"long" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationRevisionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationRevisionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationRevisionInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if s.Revision == nil {
		invalidParams.Add(aws.NewErrParamRequired("Revision"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationRevisionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "revision", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Response body for DescribeConfigurationRevision.
type DescribeConfigurationRevisionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string `locationName:"arn" type:"string"`

	// The time when the configuration was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601"`

	// The description of the configuration.
	Description *string `locationName:"description" type:"string"`

	// The revision number.
	Revision *int64 `locationName:"revision" type:"long"`

	// ServerProperties is automatically base64 encoded/decoded by the SDK.
	ServerProperties []byte `locationName:"serverProperties" type:"blob"`
}

// String returns the string representation
func (s DescribeConfigurationRevisionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationRevisionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.Int64Value(v), metadata)
	}
	if s.ServerProperties != nil {
		v := s.ServerProperties

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serverProperties", protocol.QuotedValue{ValueMarshaler: protocol.BytesValue(v)}, metadata)
	}
	return nil
}

const opDescribeConfigurationRevision = "DescribeConfigurationRevision"

// DescribeConfigurationRevisionRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a description of this revision of the configuration.
//
//    // Example sending a request using DescribeConfigurationRevisionRequest.
//    req := client.DescribeConfigurationRevisionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/DescribeConfigurationRevision
func (c *Client) DescribeConfigurationRevisionRequest(input *DescribeConfigurationRevisionInput) DescribeConfigurationRevisionRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationRevision,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/configurations/{arn}/revisions/{revision}",
	}

	if input == nil {
		input = &DescribeConfigurationRevisionInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationRevisionOutput{})

	return DescribeConfigurationRevisionRequest{Request: req, Input: input, Copy: c.DescribeConfigurationRevisionRequest}
}

// DescribeConfigurationRevisionRequest is the request type for the
// DescribeConfigurationRevision API operation.
type DescribeConfigurationRevisionRequest struct {
	*aws.Request
	Input *DescribeConfigurationRevisionInput
	Copy  func(*DescribeConfigurationRevisionInput) DescribeConfigurationRevisionRequest
}

// Send marshals and sends the DescribeConfigurationRevision API request.
func (r DescribeConfigurationRevisionRequest) Send(ctx context.Context) (*DescribeConfigurationRevisionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationRevisionResponse{
		DescribeConfigurationRevisionOutput: r.Request.Data.(*DescribeConfigurationRevisionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationRevisionResponse is the response type for the
// DescribeConfigurationRevision API operation.
type DescribeConfigurationRevisionResponse struct {
	*DescribeConfigurationRevisionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationRevision request.
func (r *DescribeConfigurationRevisionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}