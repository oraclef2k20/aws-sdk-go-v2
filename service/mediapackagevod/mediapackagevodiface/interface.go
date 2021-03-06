// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediapackagevodiface provides an interface to enable mocking the AWS Elemental MediaPackage VOD service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediapackagevodiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
)

// ClientAPI provides an interface to enable mocking the
// mediapackagevod.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // MediaPackage Vod.
//    func myFunc(svc mediapackagevodiface.ClientAPI) bool {
//        // Make svc.CreateAsset request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mediapackagevod.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mediapackagevodiface.ClientPI
//    }
//    func (m *mockClientClient) CreateAsset(input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateAssetRequest(*mediapackagevod.CreateAssetInput) mediapackagevod.CreateAssetRequest

	CreatePackagingConfigurationRequest(*mediapackagevod.CreatePackagingConfigurationInput) mediapackagevod.CreatePackagingConfigurationRequest

	CreatePackagingGroupRequest(*mediapackagevod.CreatePackagingGroupInput) mediapackagevod.CreatePackagingGroupRequest

	DeleteAssetRequest(*mediapackagevod.DeleteAssetInput) mediapackagevod.DeleteAssetRequest

	DeletePackagingConfigurationRequest(*mediapackagevod.DeletePackagingConfigurationInput) mediapackagevod.DeletePackagingConfigurationRequest

	DeletePackagingGroupRequest(*mediapackagevod.DeletePackagingGroupInput) mediapackagevod.DeletePackagingGroupRequest

	DescribeAssetRequest(*mediapackagevod.DescribeAssetInput) mediapackagevod.DescribeAssetRequest

	DescribePackagingConfigurationRequest(*mediapackagevod.DescribePackagingConfigurationInput) mediapackagevod.DescribePackagingConfigurationRequest

	DescribePackagingGroupRequest(*mediapackagevod.DescribePackagingGroupInput) mediapackagevod.DescribePackagingGroupRequest

	ListAssetsRequest(*mediapackagevod.ListAssetsInput) mediapackagevod.ListAssetsRequest

	ListPackagingConfigurationsRequest(*mediapackagevod.ListPackagingConfigurationsInput) mediapackagevod.ListPackagingConfigurationsRequest

	ListPackagingGroupsRequest(*mediapackagevod.ListPackagingGroupsInput) mediapackagevod.ListPackagingGroupsRequest

	ListTagsForResourceRequest(*mediapackagevod.ListTagsForResourceInput) mediapackagevod.ListTagsForResourceRequest

	TagResourceRequest(*mediapackagevod.TagResourceInput) mediapackagevod.TagResourceRequest

	UntagResourceRequest(*mediapackagevod.UntagResourceInput) mediapackagevod.UntagResourceRequest

	UpdatePackagingGroupRequest(*mediapackagevod.UpdatePackagingGroupInput) mediapackagevod.UpdatePackagingGroupRequest
}

var _ ClientAPI = (*mediapackagevod.Client)(nil)
