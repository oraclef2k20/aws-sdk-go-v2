// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateThemePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the theme.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// A list of resource permissions to be granted for the theme.
	GrantPermissions []ResourcePermission `type:"list"`

	// A list of resource permissions to be revoked from the theme.
	RevokePermissions []ResourcePermission `type:"list"`

	// The ID for the theme.
	//
	// ThemeId is a required field
	ThemeId *string `location:"uri" locationName:"ThemeId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateThemePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThemePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateThemePermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.ThemeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThemeId"))
	}
	if s.ThemeId != nil && len(*s.ThemeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThemeId", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThemePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GrantPermissions != nil {
		v := s.GrantPermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GrantPermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RevokePermissions != nil {
		v := s.RevokePermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RevokePermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeId != nil {
		v := *s.ThemeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ThemeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateThemePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The resulting list of resource permissions for the theme.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon Resource Name (ARN) of the theme.
	ThemeArn *string `type:"string"`

	// The ID for the theme.
	ThemeId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateThemePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThemePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeArn != nil {
		v := *s.ThemeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThemeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeId != nil {
		v := *s.ThemeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThemeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateThemePermissions = "UpdateThemePermissions"

// UpdateThemePermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the resource permissions for a theme. Permissions apply to the action
// to grant or revoke permissions on, for example "quicksight:DescribeTheme".
//
// Theme permissions apply in groupings. Valid groupings include the following
// for the three levels of permissions, which are user, owner, or no permissions:
//
//    * User "quicksight:DescribeTheme" "quicksight:DescribeThemeAlias" "quicksight:ListThemeAliases"
//    "quicksight:ListThemeVersions"
//
//    * Owner "quicksight:DescribeTheme" "quicksight:DescribeThemeAlias" "quicksight:ListThemeAliases"
//    "quicksight:ListThemeVersions" "quicksight:DeleteTheme" "quicksight:UpdateTheme"
//    "quicksight:CreateThemeAlias" "quicksight:DeleteThemeAlias" "quicksight:UpdateThemeAlias"
//    "quicksight:UpdateThemePermissions" "quicksight:DescribeThemePermissions"
//
//    * To specify no permissions, omit the permissions list.
//
//    // Example sending a request using UpdateThemePermissionsRequest.
//    req := client.UpdateThemePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateThemePermissions
func (c *Client) UpdateThemePermissionsRequest(input *UpdateThemePermissionsInput) UpdateThemePermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateThemePermissions,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/themes/{ThemeId}/permissions",
	}

	if input == nil {
		input = &UpdateThemePermissionsInput{}
	}

	req := c.newRequest(op, input, &UpdateThemePermissionsOutput{})

	return UpdateThemePermissionsRequest{Request: req, Input: input, Copy: c.UpdateThemePermissionsRequest}
}

// UpdateThemePermissionsRequest is the request type for the
// UpdateThemePermissions API operation.
type UpdateThemePermissionsRequest struct {
	*aws.Request
	Input *UpdateThemePermissionsInput
	Copy  func(*UpdateThemePermissionsInput) UpdateThemePermissionsRequest
}

// Send marshals and sends the UpdateThemePermissions API request.
func (r UpdateThemePermissionsRequest) Send(ctx context.Context) (*UpdateThemePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThemePermissionsResponse{
		UpdateThemePermissionsOutput: r.Request.Data.(*UpdateThemePermissionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThemePermissionsResponse is the response type for the
// UpdateThemePermissions API operation.
type UpdateThemePermissionsResponse struct {
	*UpdateThemePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThemePermissions request.
func (r *UpdateThemePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
