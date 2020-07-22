// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML name, changing the
// wrapper name.
func (c *Client) HttpPayloadWithXmlName(ctx context.Context, params *HttpPayloadWithXmlNameInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNameOutput, error) {
	stack := middleware.NewStack("HttpPayloadWithXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlName(options.Region), middleware.Before)
	addawsRestxml_serdeOpHttpPayloadWithXmlNameMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "HttpPayloadWithXmlName",
			Err:           err,
		}
	}
	out := result.(*HttpPayloadWithXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNameInput struct {
	Nested *types.PayloadWithXmlName
}

type HttpPayloadWithXmlNameOutput struct {
	Nested *types.PayloadWithXmlName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHttpPayloadWithXmlNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlName{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlName{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPayloadWithXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "HttpPayloadWithXmlName",
	}
}