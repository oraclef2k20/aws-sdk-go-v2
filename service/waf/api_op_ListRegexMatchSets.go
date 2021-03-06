// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRegexMatchSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of RegexMatchSet objects that you want AWS WAF to return
	// for this request. If you have more RegexMatchSet objects than the number
	// you specify for Limit, the response includes a NextMarker value that you
	// can use to get another batch of RegexMatchSet objects.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more RegexMatchSet objects
	// than the value of Limit, AWS WAF returns a NextMarker value in the response
	// that allows you to list another group of ByteMatchSets. For the second and
	// subsequent ListRegexMatchSets requests, specify the value of NextMarker from
	// the previous response to get information about another batch of RegexMatchSet
	// objects.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRegexMatchSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRegexMatchSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRegexMatchSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRegexMatchSetsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more RegexMatchSet objects than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more RegexMatchSet objects, submit another ListRegexMatchSets request, and
	// specify the NextMarker value from the response in the NextMarker value in
	// the next request.
	NextMarker *string `min:"1" type:"string"`

	// An array of RegexMatchSetSummary objects.
	RegexMatchSets []RegexMatchSetSummary `type:"list"`
}

// String returns the string representation
func (s ListRegexMatchSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRegexMatchSets = "ListRegexMatchSets"

// ListRegexMatchSetsRequest returns a request value for making API operation for
// AWS WAF.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Returns an array of RegexMatchSetSummary objects.
//
//    // Example sending a request using ListRegexMatchSetsRequest.
//    req := client.ListRegexMatchSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListRegexMatchSets
func (c *Client) ListRegexMatchSetsRequest(input *ListRegexMatchSetsInput) ListRegexMatchSetsRequest {
	op := &aws.Operation{
		Name:       opListRegexMatchSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRegexMatchSetsInput{}
	}

	req := c.newRequest(op, input, &ListRegexMatchSetsOutput{})

	return ListRegexMatchSetsRequest{Request: req, Input: input, Copy: c.ListRegexMatchSetsRequest}
}

// ListRegexMatchSetsRequest is the request type for the
// ListRegexMatchSets API operation.
type ListRegexMatchSetsRequest struct {
	*aws.Request
	Input *ListRegexMatchSetsInput
	Copy  func(*ListRegexMatchSetsInput) ListRegexMatchSetsRequest
}

// Send marshals and sends the ListRegexMatchSets API request.
func (r ListRegexMatchSetsRequest) Send(ctx context.Context) (*ListRegexMatchSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRegexMatchSetsResponse{
		ListRegexMatchSetsOutput: r.Request.Data.(*ListRegexMatchSetsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRegexMatchSetsResponse is the response type for the
// ListRegexMatchSets API operation.
type ListRegexMatchSetsResponse struct {
	*ListRegexMatchSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRegexMatchSets request.
func (r *ListRegexMatchSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
