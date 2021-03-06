// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetectDominantLanguageInput struct {
	_ struct{} `type:"structure"`

	// A UTF-8 text string. Each string should contain at least 20 characters and
	// must contain fewer that 5,000 bytes of UTF-8 encoded characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s DetectDominantLanguageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectDominantLanguageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectDominantLanguageInput"}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetectDominantLanguageOutput struct {
	_ struct{} `type:"structure" sensitive:"true"`

	// The languages that Amazon Comprehend detected in the input text. For each
	// language, the response returns the RFC 5646 language code and the level of
	// confidence that Amazon Comprehend has in the accuracy of its inference. For
	// more information about RFC 5646, see Tags for Identifying Languages (https://tools.ietf.org/html/rfc5646)
	// on the IETF Tools web site.
	Languages []DominantLanguage `type:"list"`
}

// String returns the string representation
func (s DetectDominantLanguageOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectDominantLanguage = "DetectDominantLanguage"

// DetectDominantLanguageRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Determines the dominant language of the input text. For a list of languages
// that Amazon Comprehend can detect, see Amazon Comprehend Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
//
//    // Example sending a request using DetectDominantLanguageRequest.
//    req := client.DetectDominantLanguageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DetectDominantLanguage
func (c *Client) DetectDominantLanguageRequest(input *DetectDominantLanguageInput) DetectDominantLanguageRequest {
	op := &aws.Operation{
		Name:       opDetectDominantLanguage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectDominantLanguageInput{}
	}

	req := c.newRequest(op, input, &DetectDominantLanguageOutput{})

	return DetectDominantLanguageRequest{Request: req, Input: input, Copy: c.DetectDominantLanguageRequest}
}

// DetectDominantLanguageRequest is the request type for the
// DetectDominantLanguage API operation.
type DetectDominantLanguageRequest struct {
	*aws.Request
	Input *DetectDominantLanguageInput
	Copy  func(*DetectDominantLanguageInput) DetectDominantLanguageRequest
}

// Send marshals and sends the DetectDominantLanguage API request.
func (r DetectDominantLanguageRequest) Send(ctx context.Context) (*DetectDominantLanguageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectDominantLanguageResponse{
		DetectDominantLanguageOutput: r.Request.Data.(*DetectDominantLanguageOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectDominantLanguageResponse is the response type for the
// DetectDominantLanguage API operation.
type DetectDominantLanguageResponse struct {
	*DetectDominantLanguageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectDominantLanguage request.
func (r *DetectDominantLanguageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
