// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateLunaClient action.
type CreateLunaClientInput struct {
	_ struct{} `type:"structure"`

	// The contents of a Base64-Encoded X.509 v3 certificate to be installed on
	// the HSMs used by this client.
	//
	// Certificate is a required field
	Certificate *string `min:"600" type:"string" required:"true"`

	// The label for the client.
	Label *string `type:"string"`
}

// String returns the string representation
func (s CreateLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLunaClientInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLunaClientInput"}

	if s.Certificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificate"))
	}
	if s.Certificate != nil && len(*s.Certificate) < 600 {
		invalidParams.Add(aws.NewErrParamMinLen("Certificate", 600))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the CreateLunaClient action.
type CreateLunaClientOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the client.
	ClientArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLunaClient = "CreateLunaClient"

// CreateLunaClientRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Creates an HSM client.
//
//    // Example sending a request using CreateLunaClientRequest.
//    req := client.CreateLunaClientRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/CreateLunaClient
func (c *Client) CreateLunaClientRequest(input *CreateLunaClientInput) CreateLunaClientRequest {
	op := &aws.Operation{
		Name:       opCreateLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLunaClientInput{}
	}

	req := c.newRequest(op, input, &CreateLunaClientOutput{})

	return CreateLunaClientRequest{Request: req, Input: input, Copy: c.CreateLunaClientRequest}
}

// CreateLunaClientRequest is the request type for the
// CreateLunaClient API operation.
type CreateLunaClientRequest struct {
	*aws.Request
	Input *CreateLunaClientInput
	Copy  func(*CreateLunaClientInput) CreateLunaClientRequest
}

// Send marshals and sends the CreateLunaClient API request.
func (r CreateLunaClientRequest) Send(ctx context.Context) (*CreateLunaClientResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLunaClientResponse{
		CreateLunaClientOutput: r.Request.Data.(*CreateLunaClientOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLunaClientResponse is the response type for the
// CreateLunaClient API operation.
type CreateLunaClientResponse struct {
	*CreateLunaClientOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLunaClient request.
func (r *CreateLunaClientResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}