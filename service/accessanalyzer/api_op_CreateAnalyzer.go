// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Creates an analyzer.
type CreateAnalyzerInput struct {
	_ struct{} `type:"structure"`

	// The name of the analyzer to create.
	//
	// AnalyzerName is a required field
	AnalyzerName *string `locationName:"analyzerName" min:"1" type:"string" required:"true"`

	// Specifies the archive rules to add for the analyzer. Archive rules automatically
	// archive findings that meet the criteria you define for the rule.
	ArchiveRules []InlineArchiveRule `locationName:"archiveRules" type:"list"`

	// A client token.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The tags to apply to the analyzer.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The type of analyzer to create. Only ACCOUNT analyzers are supported. You
	// can create only one analyzer per account per Region.
	//
	// Type is a required field
	Type Type `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateAnalyzerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAnalyzerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAnalyzerInput"}

	if s.AnalyzerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalyzerName"))
	}
	if s.AnalyzerName != nil && len(*s.AnalyzerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AnalyzerName", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}
	if s.ArchiveRules != nil {
		for i, v := range s.ArchiveRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ArchiveRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAnalyzerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AnalyzerName != nil {
		v := *s.AnalyzerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "analyzerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ArchiveRules != nil {
		v := s.ArchiveRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "archiveRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The response to the request to create an analyzer.
type CreateAnalyzerOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the analyzer that was created by the request.
	Arn *string `locationName:"arn" type:"string"`
}

// String returns the string representation
func (s CreateAnalyzerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAnalyzerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateAnalyzer = "CreateAnalyzer"

// CreateAnalyzerRequest returns a request value for making API operation for
// Access Analyzer.
//
// Creates an analyzer for your account.
//
//    // Example sending a request using CreateAnalyzerRequest.
//    req := client.CreateAnalyzerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/accessanalyzer-2019-11-01/CreateAnalyzer
func (c *Client) CreateAnalyzerRequest(input *CreateAnalyzerInput) CreateAnalyzerRequest {
	op := &aws.Operation{
		Name:       opCreateAnalyzer,
		HTTPMethod: "PUT",
		HTTPPath:   "/analyzer",
	}

	if input == nil {
		input = &CreateAnalyzerInput{}
	}

	req := c.newRequest(op, input, &CreateAnalyzerOutput{})

	return CreateAnalyzerRequest{Request: req, Input: input, Copy: c.CreateAnalyzerRequest}
}

// CreateAnalyzerRequest is the request type for the
// CreateAnalyzer API operation.
type CreateAnalyzerRequest struct {
	*aws.Request
	Input *CreateAnalyzerInput
	Copy  func(*CreateAnalyzerInput) CreateAnalyzerRequest
}

// Send marshals and sends the CreateAnalyzer API request.
func (r CreateAnalyzerRequest) Send(ctx context.Context) (*CreateAnalyzerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAnalyzerResponse{
		CreateAnalyzerOutput: r.Request.Data.(*CreateAnalyzerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAnalyzerResponse is the response type for the
// CreateAnalyzer API operation.
type CreateAnalyzerResponse struct {
	*CreateAnalyzerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAnalyzer request.
func (r *CreateAnalyzerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}