// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDirectoryConfigInput struct {
	_ struct{} `type:"structure"`

	// The fully qualified name of the directory (for example, corp.example.com).
	//
	// DirectoryName is a required field
	DirectoryName *string `type:"string" required:"true"`

	// The distinguished names of the organizational units for computer accounts.
	//
	// OrganizationalUnitDistinguishedNames is a required field
	OrganizationalUnitDistinguishedNames []string `type:"list" required:"true"`

	// The credentials for the service account used by the fleet or image builder
	// to connect to the directory.
	//
	// ServiceAccountCredentials is a required field
	ServiceAccountCredentials *ServiceAccountCredentials `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateDirectoryConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectoryConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDirectoryConfigInput"}

	if s.DirectoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryName"))
	}

	if s.OrganizationalUnitDistinguishedNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationalUnitDistinguishedNames"))
	}

	if s.ServiceAccountCredentials == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceAccountCredentials"))
	}
	if s.ServiceAccountCredentials != nil {
		if err := s.ServiceAccountCredentials.Validate(); err != nil {
			invalidParams.AddNested("ServiceAccountCredentials", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDirectoryConfigOutput struct {
	_ struct{} `type:"structure"`

	// Information about the directory configuration.
	DirectoryConfig *DirectoryConfig `type:"structure"`
}

// String returns the string representation
func (s CreateDirectoryConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDirectoryConfig = "CreateDirectoryConfig"

// CreateDirectoryConfigRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a Directory Config object in AppStream 2.0. This object includes
// the configuration information required to join fleets and image builders
// to Microsoft Active Directory domains.
//
//    // Example sending a request using CreateDirectoryConfigRequest.
//    req := client.CreateDirectoryConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateDirectoryConfig
func (c *Client) CreateDirectoryConfigRequest(input *CreateDirectoryConfigInput) CreateDirectoryConfigRequest {
	op := &aws.Operation{
		Name:       opCreateDirectoryConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectoryConfigInput{}
	}

	req := c.newRequest(op, input, &CreateDirectoryConfigOutput{})

	return CreateDirectoryConfigRequest{Request: req, Input: input, Copy: c.CreateDirectoryConfigRequest}
}

// CreateDirectoryConfigRequest is the request type for the
// CreateDirectoryConfig API operation.
type CreateDirectoryConfigRequest struct {
	*aws.Request
	Input *CreateDirectoryConfigInput
	Copy  func(*CreateDirectoryConfigInput) CreateDirectoryConfigRequest
}

// Send marshals and sends the CreateDirectoryConfig API request.
func (r CreateDirectoryConfigRequest) Send(ctx context.Context) (*CreateDirectoryConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDirectoryConfigResponse{
		CreateDirectoryConfigOutput: r.Request.Data.(*CreateDirectoryConfigOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDirectoryConfigResponse is the response type for the
// CreateDirectoryConfig API operation.
type CreateDirectoryConfigResponse struct {
	*CreateDirectoryConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDirectoryConfig request.
func (r *CreateDirectoryConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}