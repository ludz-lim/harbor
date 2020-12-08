// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package schemasiface provides an interface to enable mocking the Schemas service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package schemasiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/schemas"
)

// SchemasAPI provides an interface to enable mocking the
// schemas.Schemas service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Schemas.
//    func myFunc(svc schemasiface.SchemasAPI) bool {
//        // Make svc.CreateDiscoverer request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := schemas.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSchemasClient struct {
//        schemasiface.SchemasAPI
//    }
//    func (m *mockSchemasClient) CreateDiscoverer(input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSchemasClient{}
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
type SchemasAPI interface {
	CreateDiscoverer(*schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error)
	CreateDiscovererWithContext(aws.Context, *schemas.CreateDiscovererInput, ...request.Option) (*schemas.CreateDiscovererOutput, error)
	CreateDiscovererRequest(*schemas.CreateDiscovererInput) (*request.Request, *schemas.CreateDiscovererOutput)

	CreateRegistry(*schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error)
	CreateRegistryWithContext(aws.Context, *schemas.CreateRegistryInput, ...request.Option) (*schemas.CreateRegistryOutput, error)
	CreateRegistryRequest(*schemas.CreateRegistryInput) (*request.Request, *schemas.CreateRegistryOutput)

	CreateSchema(*schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error)
	CreateSchemaWithContext(aws.Context, *schemas.CreateSchemaInput, ...request.Option) (*schemas.CreateSchemaOutput, error)
	CreateSchemaRequest(*schemas.CreateSchemaInput) (*request.Request, *schemas.CreateSchemaOutput)

	DeleteDiscoverer(*schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error)
	DeleteDiscovererWithContext(aws.Context, *schemas.DeleteDiscovererInput, ...request.Option) (*schemas.DeleteDiscovererOutput, error)
	DeleteDiscovererRequest(*schemas.DeleteDiscovererInput) (*request.Request, *schemas.DeleteDiscovererOutput)

	DeleteRegistry(*schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error)
	DeleteRegistryWithContext(aws.Context, *schemas.DeleteRegistryInput, ...request.Option) (*schemas.DeleteRegistryOutput, error)
	DeleteRegistryRequest(*schemas.DeleteRegistryInput) (*request.Request, *schemas.DeleteRegistryOutput)

	DeleteResourcePolicy(*schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *schemas.DeleteResourcePolicyInput, ...request.Option) (*schemas.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*schemas.DeleteResourcePolicyInput) (*request.Request, *schemas.DeleteResourcePolicyOutput)

	DeleteSchema(*schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error)
	DeleteSchemaWithContext(aws.Context, *schemas.DeleteSchemaInput, ...request.Option) (*schemas.DeleteSchemaOutput, error)
	DeleteSchemaRequest(*schemas.DeleteSchemaInput) (*request.Request, *schemas.DeleteSchemaOutput)

	DeleteSchemaVersion(*schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error)
	DeleteSchemaVersionWithContext(aws.Context, *schemas.DeleteSchemaVersionInput, ...request.Option) (*schemas.DeleteSchemaVersionOutput, error)
	DeleteSchemaVersionRequest(*schemas.DeleteSchemaVersionInput) (*request.Request, *schemas.DeleteSchemaVersionOutput)

	DescribeCodeBinding(*schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error)
	DescribeCodeBindingWithContext(aws.Context, *schemas.DescribeCodeBindingInput, ...request.Option) (*schemas.DescribeCodeBindingOutput, error)
	DescribeCodeBindingRequest(*schemas.DescribeCodeBindingInput) (*request.Request, *schemas.DescribeCodeBindingOutput)

	DescribeDiscoverer(*schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error)
	DescribeDiscovererWithContext(aws.Context, *schemas.DescribeDiscovererInput, ...request.Option) (*schemas.DescribeDiscovererOutput, error)
	DescribeDiscovererRequest(*schemas.DescribeDiscovererInput) (*request.Request, *schemas.DescribeDiscovererOutput)

	DescribeRegistry(*schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error)
	DescribeRegistryWithContext(aws.Context, *schemas.DescribeRegistryInput, ...request.Option) (*schemas.DescribeRegistryOutput, error)
	DescribeRegistryRequest(*schemas.DescribeRegistryInput) (*request.Request, *schemas.DescribeRegistryOutput)

	DescribeSchema(*schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error)
	DescribeSchemaWithContext(aws.Context, *schemas.DescribeSchemaInput, ...request.Option) (*schemas.DescribeSchemaOutput, error)
	DescribeSchemaRequest(*schemas.DescribeSchemaInput) (*request.Request, *schemas.DescribeSchemaOutput)

	ExportSchema(*schemas.ExportSchemaInput) (*schemas.ExportSchemaOutput, error)
	ExportSchemaWithContext(aws.Context, *schemas.ExportSchemaInput, ...request.Option) (*schemas.ExportSchemaOutput, error)
	ExportSchemaRequest(*schemas.ExportSchemaInput) (*request.Request, *schemas.ExportSchemaOutput)

	GetCodeBindingSource(*schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error)
	GetCodeBindingSourceWithContext(aws.Context, *schemas.GetCodeBindingSourceInput, ...request.Option) (*schemas.GetCodeBindingSourceOutput, error)
	GetCodeBindingSourceRequest(*schemas.GetCodeBindingSourceInput) (*request.Request, *schemas.GetCodeBindingSourceOutput)

	GetDiscoveredSchema(*schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error)
	GetDiscoveredSchemaWithContext(aws.Context, *schemas.GetDiscoveredSchemaInput, ...request.Option) (*schemas.GetDiscoveredSchemaOutput, error)
	GetDiscoveredSchemaRequest(*schemas.GetDiscoveredSchemaInput) (*request.Request, *schemas.GetDiscoveredSchemaOutput)

	GetResourcePolicy(*schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error)
	GetResourcePolicyWithContext(aws.Context, *schemas.GetResourcePolicyInput, ...request.Option) (*schemas.GetResourcePolicyOutput, error)
	GetResourcePolicyRequest(*schemas.GetResourcePolicyInput) (*request.Request, *schemas.GetResourcePolicyOutput)

	ListDiscoverers(*schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error)
	ListDiscoverersWithContext(aws.Context, *schemas.ListDiscoverersInput, ...request.Option) (*schemas.ListDiscoverersOutput, error)
	ListDiscoverersRequest(*schemas.ListDiscoverersInput) (*request.Request, *schemas.ListDiscoverersOutput)

	ListDiscoverersPages(*schemas.ListDiscoverersInput, func(*schemas.ListDiscoverersOutput, bool) bool) error
	ListDiscoverersPagesWithContext(aws.Context, *schemas.ListDiscoverersInput, func(*schemas.ListDiscoverersOutput, bool) bool, ...request.Option) error

	ListRegistries(*schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error)
	ListRegistriesWithContext(aws.Context, *schemas.ListRegistriesInput, ...request.Option) (*schemas.ListRegistriesOutput, error)
	ListRegistriesRequest(*schemas.ListRegistriesInput) (*request.Request, *schemas.ListRegistriesOutput)

	ListRegistriesPages(*schemas.ListRegistriesInput, func(*schemas.ListRegistriesOutput, bool) bool) error
	ListRegistriesPagesWithContext(aws.Context, *schemas.ListRegistriesInput, func(*schemas.ListRegistriesOutput, bool) bool, ...request.Option) error

	ListSchemaVersions(*schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error)
	ListSchemaVersionsWithContext(aws.Context, *schemas.ListSchemaVersionsInput, ...request.Option) (*schemas.ListSchemaVersionsOutput, error)
	ListSchemaVersionsRequest(*schemas.ListSchemaVersionsInput) (*request.Request, *schemas.ListSchemaVersionsOutput)

	ListSchemaVersionsPages(*schemas.ListSchemaVersionsInput, func(*schemas.ListSchemaVersionsOutput, bool) bool) error
	ListSchemaVersionsPagesWithContext(aws.Context, *schemas.ListSchemaVersionsInput, func(*schemas.ListSchemaVersionsOutput, bool) bool, ...request.Option) error

	ListSchemas(*schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error)
	ListSchemasWithContext(aws.Context, *schemas.ListSchemasInput, ...request.Option) (*schemas.ListSchemasOutput, error)
	ListSchemasRequest(*schemas.ListSchemasInput) (*request.Request, *schemas.ListSchemasOutput)

	ListSchemasPages(*schemas.ListSchemasInput, func(*schemas.ListSchemasOutput, bool) bool) error
	ListSchemasPagesWithContext(aws.Context, *schemas.ListSchemasInput, func(*schemas.ListSchemasOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *schemas.ListTagsForResourceInput, ...request.Option) (*schemas.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*schemas.ListTagsForResourceInput) (*request.Request, *schemas.ListTagsForResourceOutput)

	PutCodeBinding(*schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error)
	PutCodeBindingWithContext(aws.Context, *schemas.PutCodeBindingInput, ...request.Option) (*schemas.PutCodeBindingOutput, error)
	PutCodeBindingRequest(*schemas.PutCodeBindingInput) (*request.Request, *schemas.PutCodeBindingOutput)

	PutResourcePolicy(*schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *schemas.PutResourcePolicyInput, ...request.Option) (*schemas.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*schemas.PutResourcePolicyInput) (*request.Request, *schemas.PutResourcePolicyOutput)

	SearchSchemas(*schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error)
	SearchSchemasWithContext(aws.Context, *schemas.SearchSchemasInput, ...request.Option) (*schemas.SearchSchemasOutput, error)
	SearchSchemasRequest(*schemas.SearchSchemasInput) (*request.Request, *schemas.SearchSchemasOutput)

	SearchSchemasPages(*schemas.SearchSchemasInput, func(*schemas.SearchSchemasOutput, bool) bool) error
	SearchSchemasPagesWithContext(aws.Context, *schemas.SearchSchemasInput, func(*schemas.SearchSchemasOutput, bool) bool, ...request.Option) error

	StartDiscoverer(*schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error)
	StartDiscovererWithContext(aws.Context, *schemas.StartDiscovererInput, ...request.Option) (*schemas.StartDiscovererOutput, error)
	StartDiscovererRequest(*schemas.StartDiscovererInput) (*request.Request, *schemas.StartDiscovererOutput)

	StopDiscoverer(*schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error)
	StopDiscovererWithContext(aws.Context, *schemas.StopDiscovererInput, ...request.Option) (*schemas.StopDiscovererOutput, error)
	StopDiscovererRequest(*schemas.StopDiscovererInput) (*request.Request, *schemas.StopDiscovererOutput)

	TagResource(*schemas.TagResourceInput) (*schemas.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *schemas.TagResourceInput, ...request.Option) (*schemas.TagResourceOutput, error)
	TagResourceRequest(*schemas.TagResourceInput) (*request.Request, *schemas.TagResourceOutput)

	UntagResource(*schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *schemas.UntagResourceInput, ...request.Option) (*schemas.UntagResourceOutput, error)
	UntagResourceRequest(*schemas.UntagResourceInput) (*request.Request, *schemas.UntagResourceOutput)

	UpdateDiscoverer(*schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error)
	UpdateDiscovererWithContext(aws.Context, *schemas.UpdateDiscovererInput, ...request.Option) (*schemas.UpdateDiscovererOutput, error)
	UpdateDiscovererRequest(*schemas.UpdateDiscovererInput) (*request.Request, *schemas.UpdateDiscovererOutput)

	UpdateRegistry(*schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error)
	UpdateRegistryWithContext(aws.Context, *schemas.UpdateRegistryInput, ...request.Option) (*schemas.UpdateRegistryOutput, error)
	UpdateRegistryRequest(*schemas.UpdateRegistryInput) (*request.Request, *schemas.UpdateRegistryOutput)

	UpdateSchema(*schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error)
	UpdateSchemaWithContext(aws.Context, *schemas.UpdateSchemaInput, ...request.Option) (*schemas.UpdateSchemaOutput, error)
	UpdateSchemaRequest(*schemas.UpdateSchemaInput) (*request.Request, *schemas.UpdateSchemaOutput)

	WaitUntilCodeBindingExists(*schemas.DescribeCodeBindingInput) error
	WaitUntilCodeBindingExistsWithContext(aws.Context, *schemas.DescribeCodeBindingInput, ...request.WaiterOption) error
}

var _ SchemasAPI = (*schemas.Schemas)(nil)
