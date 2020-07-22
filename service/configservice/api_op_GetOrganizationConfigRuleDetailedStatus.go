// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetOrganizationConfigRuleDetailedStatusInput struct {
	_ struct{} `type:"structure"`

	// A StatusDetailFilters object.
	Filters *StatusDetailFilters `type:"structure"`

	// The maximum number of OrganizationConfigRuleDetailedStatus returned on each
	// page. If you do not specify a number, AWS Config uses the default. The default
	// is 100.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// The name of organization config rule for which you want status details for
	// member accounts.
	//
	// OrganizationConfigRuleName is a required field
	OrganizationConfigRuleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetOrganizationConfigRuleDetailedStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOrganizationConfigRuleDetailedStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOrganizationConfigRuleDetailedStatusInput"}

	if s.OrganizationConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConfigRuleName"))
	}
	if s.OrganizationConfigRuleName != nil && len(*s.OrganizationConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOrganizationConfigRuleDetailedStatusOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// A list of MemberAccountStatus objects.
	OrganizationConfigRuleDetailedStatus []MemberAccountStatus `type:"list"`
}

// String returns the string representation
func (s GetOrganizationConfigRuleDetailedStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOrganizationConfigRuleDetailedStatus = "GetOrganizationConfigRuleDetailedStatus"

// GetOrganizationConfigRuleDetailedStatusRequest returns a request value for making API operation for
// AWS Config.
//
// Returns detailed status for each member account within an organization for
// a given organization config rule.
//
// Only a master account can call this API.
//
//    // Example sending a request using GetOrganizationConfigRuleDetailedStatusRequest.
//    req := client.GetOrganizationConfigRuleDetailedStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetOrganizationConfigRuleDetailedStatus
func (c *Client) GetOrganizationConfigRuleDetailedStatusRequest(input *GetOrganizationConfigRuleDetailedStatusInput) GetOrganizationConfigRuleDetailedStatusRequest {
	op := &aws.Operation{
		Name:       opGetOrganizationConfigRuleDetailedStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOrganizationConfigRuleDetailedStatusInput{}
	}

	req := c.newRequest(op, input, &GetOrganizationConfigRuleDetailedStatusOutput{})

	return GetOrganizationConfigRuleDetailedStatusRequest{Request: req, Input: input, Copy: c.GetOrganizationConfigRuleDetailedStatusRequest}
}

// GetOrganizationConfigRuleDetailedStatusRequest is the request type for the
// GetOrganizationConfigRuleDetailedStatus API operation.
type GetOrganizationConfigRuleDetailedStatusRequest struct {
	*aws.Request
	Input *GetOrganizationConfigRuleDetailedStatusInput
	Copy  func(*GetOrganizationConfigRuleDetailedStatusInput) GetOrganizationConfigRuleDetailedStatusRequest
}

// Send marshals and sends the GetOrganizationConfigRuleDetailedStatus API request.
func (r GetOrganizationConfigRuleDetailedStatusRequest) Send(ctx context.Context) (*GetOrganizationConfigRuleDetailedStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOrganizationConfigRuleDetailedStatusResponse{
		GetOrganizationConfigRuleDetailedStatusOutput: r.Request.Data.(*GetOrganizationConfigRuleDetailedStatusOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOrganizationConfigRuleDetailedStatusResponse is the response type for the
// GetOrganizationConfigRuleDetailedStatus API operation.
type GetOrganizationConfigRuleDetailedStatusResponse struct {
	*GetOrganizationConfigRuleDetailedStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOrganizationConfigRuleDetailedStatus request.
func (r *GetOrganizationConfigRuleDetailedStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}