// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteReportGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the report group to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReportGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReportGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReportGroupInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteReportGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteReportGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReportGroup = "DeleteReportGroup"

// DeleteReportGroupRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// DeleteReportGroup: Deletes a report group. Before you delete a report group,
// you must delete its reports. Use ListReportsForReportGroup (https://docs.aws.amazon.com/codebuild/latest/APIReference/API_ListReportsForReportGroup.html)
// to get the reports in a report group. Use DeleteReport (https://docs.aws.amazon.com/codebuild/latest/APIReference/API_DeleteReport.html)
// to delete the reports. If you call DeleteReportGroup for a report group that
// contains one or more reports, an exception is thrown.
//
//    // Example sending a request using DeleteReportGroupRequest.
//    req := client.DeleteReportGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DeleteReportGroup
func (c *Client) DeleteReportGroupRequest(input *DeleteReportGroupInput) DeleteReportGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteReportGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReportGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteReportGroupOutput{})

	return DeleteReportGroupRequest{Request: req, Input: input, Copy: c.DeleteReportGroupRequest}
}

// DeleteReportGroupRequest is the request type for the
// DeleteReportGroup API operation.
type DeleteReportGroupRequest struct {
	*aws.Request
	Input *DeleteReportGroupInput
	Copy  func(*DeleteReportGroupInput) DeleteReportGroupRequest
}

// Send marshals and sends the DeleteReportGroup API request.
func (r DeleteReportGroupRequest) Send(ctx context.Context) (*DeleteReportGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReportGroupResponse{
		DeleteReportGroupOutput: r.Request.Data.(*DeleteReportGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReportGroupResponse is the response type for the
// DeleteReportGroup API operation.
type DeleteReportGroupResponse struct {
	*DeleteReportGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReportGroup request.
func (r *DeleteReportGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}