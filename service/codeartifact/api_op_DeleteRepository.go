// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that contains the repository to delete.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`

	// The name of the repository to delete.
	//
	// Repository is a required field
	Repository *string `location:"querystring" locationName:"repository" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRepositoryInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}

	if s.Repository == nil {
		invalidParams.Add(aws.NewErrParamRequired("Repository"))
	}
	if s.Repository != nil && len(*s.Repository) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Repository", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRepositoryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Repository != nil {
		v := *s.Repository

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "repository", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted repository after processing the request.
	Repository *RepositoryDescription `locationName:"repository" type:"structure"`
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRepositoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Repository != nil {
		v := s.Repository

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "repository", v, metadata)
	}
	return nil
}

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest returns a request value for making API operation for
// CodeArtifact.
//
// Deletes a repository.
//
//    // Example sending a request using DeleteRepositoryRequest.
//    req := client.DeleteRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/DeleteRepository
func (c *Client) DeleteRepositoryRequest(input *DeleteRepositoryInput) DeleteRepositoryRequest {
	op := &aws.Operation{
		Name:       opDeleteRepository,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/repository",
	}

	if input == nil {
		input = &DeleteRepositoryInput{}
	}

	req := c.newRequest(op, input, &DeleteRepositoryOutput{})

	return DeleteRepositoryRequest{Request: req, Input: input, Copy: c.DeleteRepositoryRequest}
}

// DeleteRepositoryRequest is the request type for the
// DeleteRepository API operation.
type DeleteRepositoryRequest struct {
	*aws.Request
	Input *DeleteRepositoryInput
	Copy  func(*DeleteRepositoryInput) DeleteRepositoryRequest
}

// Send marshals and sends the DeleteRepository API request.
func (r DeleteRepositoryRequest) Send(ctx context.Context) (*DeleteRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRepositoryResponse{
		DeleteRepositoryOutput: r.Request.Data.(*DeleteRepositoryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRepositoryResponse is the response type for the
// DeleteRepository API operation.
type DeleteRepositoryResponse struct {
	*DeleteRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRepository request.
func (r *DeleteRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
