// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RuleGroupId of the RuleGroup that you want to delete. RuleGroupId is
	// returned by CreateRuleGroup and by ListRuleGroups.
	//
	// RuleGroupId is a required field
	RuleGroupId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRuleGroupInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RuleGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleGroupId"))
	}
	if s.RuleGroupId != nil && len(*s.RuleGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleGroupId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteRuleGroup request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRuleGroup = "DeleteRuleGroup"

// DeleteRuleGroupRequest returns a request value for making API operation for
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
// Permanently deletes a RuleGroup. You can't delete a RuleGroup if it's still
// used in any WebACL objects or if it still includes any rules.
//
// If you just want to remove a RuleGroup from a WebACL, use UpdateWebACL.
//
// To permanently delete a RuleGroup from AWS WAF, perform the following steps:
//
// Update the RuleGroup to remove rules, if any. For more information, see UpdateRuleGroup.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteRuleGroup request.
//
// Submit a DeleteRuleGroup request.
//
//    // Example sending a request using DeleteRuleGroupRequest.
//    req := client.DeleteRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRuleGroup
func (c *Client) DeleteRuleGroupRequest(input *DeleteRuleGroupInput) DeleteRuleGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRuleGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteRuleGroupOutput{})

	return DeleteRuleGroupRequest{Request: req, Input: input, Copy: c.DeleteRuleGroupRequest}
}

// DeleteRuleGroupRequest is the request type for the
// DeleteRuleGroup API operation.
type DeleteRuleGroupRequest struct {
	*aws.Request
	Input *DeleteRuleGroupInput
	Copy  func(*DeleteRuleGroupInput) DeleteRuleGroupRequest
}

// Send marshals and sends the DeleteRuleGroup API request.
func (r DeleteRuleGroupRequest) Send(ctx context.Context) (*DeleteRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRuleGroupResponse{
		DeleteRuleGroupOutput: r.Request.Data.(*DeleteRuleGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRuleGroupResponse is the response type for the
// DeleteRuleGroup API operation.
type DeleteRuleGroupResponse struct {
	*DeleteRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRuleGroup request.
func (r *DeleteRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}