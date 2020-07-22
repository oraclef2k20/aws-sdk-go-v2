// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StopHumanLoopInput struct {
	_ struct{} `type:"structure"`

	// The name of the human loop that you want to stop.
	//
	// HumanLoopName is a required field
	HumanLoopName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopHumanLoopInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopHumanLoopInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopHumanLoopInput"}

	if s.HumanLoopName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HumanLoopName"))
	}
	if s.HumanLoopName != nil && len(*s.HumanLoopName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HumanLoopName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopHumanLoopInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.HumanLoopName != nil {
		v := *s.HumanLoopName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StopHumanLoopOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopHumanLoopOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopHumanLoopOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStopHumanLoop = "StopHumanLoop"

// StopHumanLoopRequest returns a request value for making API operation for
// Amazon Augmented AI Runtime.
//
// Stops the specified human loop.
//
//    // Example sending a request using StopHumanLoopRequest.
//    req := client.StopHumanLoopRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-a2i-runtime-2019-11-07/StopHumanLoop
func (c *Client) StopHumanLoopRequest(input *StopHumanLoopInput) StopHumanLoopRequest {
	op := &aws.Operation{
		Name:       opStopHumanLoop,
		HTTPMethod: "POST",
		HTTPPath:   "/human-loops/stop",
	}

	if input == nil {
		input = &StopHumanLoopInput{}
	}

	req := c.newRequest(op, input, &StopHumanLoopOutput{})

	return StopHumanLoopRequest{Request: req, Input: input, Copy: c.StopHumanLoopRequest}
}

// StopHumanLoopRequest is the request type for the
// StopHumanLoop API operation.
type StopHumanLoopRequest struct {
	*aws.Request
	Input *StopHumanLoopInput
	Copy  func(*StopHumanLoopInput) StopHumanLoopRequest
}

// Send marshals and sends the StopHumanLoop API request.
func (r StopHumanLoopRequest) Send(ctx context.Context) (*StopHumanLoopResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopHumanLoopResponse{
		StopHumanLoopOutput: r.Request.Data.(*StopHumanLoopOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopHumanLoopResponse is the response type for the
// StopHumanLoop API operation.
type StopHumanLoopResponse struct {
	*StopHumanLoopOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopHumanLoop request.
func (r *StopHumanLoopResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}