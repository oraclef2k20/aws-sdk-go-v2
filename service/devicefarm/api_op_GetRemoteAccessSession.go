// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get information about the specified remote access
// session.
type GetRemoteAccessSessionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the remote access session about which you
	// want to get session information.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRemoteAccessSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRemoteAccessSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRemoteAccessSessionInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server that lists detailed information about
// the remote access session.
type GetRemoteAccessSessionOutput struct {
	_ struct{} `type:"structure"`

	// A container that lists detailed information about the remote access session.
	RemoteAccessSession *RemoteAccessSession `locationName:"remoteAccessSession" type:"structure"`
}

// String returns the string representation
func (s GetRemoteAccessSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRemoteAccessSession = "GetRemoteAccessSession"

// GetRemoteAccessSessionRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns a link to a currently running remote access session.
//
//    // Example sending a request using GetRemoteAccessSessionRequest.
//    req := client.GetRemoteAccessSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetRemoteAccessSession
func (c *Client) GetRemoteAccessSessionRequest(input *GetRemoteAccessSessionInput) GetRemoteAccessSessionRequest {
	op := &aws.Operation{
		Name:       opGetRemoteAccessSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRemoteAccessSessionInput{}
	}

	req := c.newRequest(op, input, &GetRemoteAccessSessionOutput{})

	return GetRemoteAccessSessionRequest{Request: req, Input: input, Copy: c.GetRemoteAccessSessionRequest}
}

// GetRemoteAccessSessionRequest is the request type for the
// GetRemoteAccessSession API operation.
type GetRemoteAccessSessionRequest struct {
	*aws.Request
	Input *GetRemoteAccessSessionInput
	Copy  func(*GetRemoteAccessSessionInput) GetRemoteAccessSessionRequest
}

// Send marshals and sends the GetRemoteAccessSession API request.
func (r GetRemoteAccessSessionRequest) Send(ctx context.Context) (*GetRemoteAccessSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRemoteAccessSessionResponse{
		GetRemoteAccessSessionOutput: r.Request.Data.(*GetRemoteAccessSessionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRemoteAccessSessionResponse is the response type for the
// GetRemoteAccessSession API operation.
type GetRemoteAccessSessionResponse struct {
	*GetRemoteAccessSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRemoteAccessSession request.
func (r *GetRemoteAccessSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}