// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ConfirmConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the hosted connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConfirmConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState ConnectionState `locationName:"connectionState" type:"string" enum:"true"`
}

// String returns the string representation
func (s ConfirmConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opConfirmConnection = "ConfirmConnection"

// ConfirmConnectionRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Confirms the creation of the specified hosted connection on an interconnect.
//
// Upon creation, the hosted connection is initially in the Ordering state,
// and remains in this state until the owner confirms creation of the hosted
// connection.
//
//    // Example sending a request using ConfirmConnectionRequest.
//    req := client.ConfirmConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/ConfirmConnection
func (c *Client) ConfirmConnectionRequest(input *ConfirmConnectionInput) ConfirmConnectionRequest {
	op := &aws.Operation{
		Name:       opConfirmConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConfirmConnectionInput{}
	}

	req := c.newRequest(op, input, &ConfirmConnectionOutput{})

	return ConfirmConnectionRequest{Request: req, Input: input, Copy: c.ConfirmConnectionRequest}
}

// ConfirmConnectionRequest is the request type for the
// ConfirmConnection API operation.
type ConfirmConnectionRequest struct {
	*aws.Request
	Input *ConfirmConnectionInput
	Copy  func(*ConfirmConnectionInput) ConfirmConnectionRequest
}

// Send marshals and sends the ConfirmConnection API request.
func (r ConfirmConnectionRequest) Send(ctx context.Context) (*ConfirmConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmConnectionResponse{
		ConfirmConnectionOutput: r.Request.Data.(*ConfirmConnectionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmConnectionResponse is the response type for the
// ConfirmConnection API operation.
type ConfirmConnectionResponse struct {
	*ConfirmConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmConnection request.
func (r *ConfirmConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
