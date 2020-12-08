// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package timestreamqueryiface provides an interface to enable mocking the Amazon Timestream Query service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package timestreamqueryiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/timestreamquery"
)

// TimestreamQueryAPI provides an interface to enable mocking the
// timestreamquery.TimestreamQuery service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Timestream Query.
//    func myFunc(svc timestreamqueryiface.TimestreamQueryAPI) bool {
//        // Make svc.CancelQuery request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := timestreamquery.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTimestreamQueryClient struct {
//        timestreamqueryiface.TimestreamQueryAPI
//    }
//    func (m *mockTimestreamQueryClient) CancelQuery(input *timestreamquery.CancelQueryInput) (*timestreamquery.CancelQueryOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTimestreamQueryClient{}
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
type TimestreamQueryAPI interface {
	CancelQuery(*timestreamquery.CancelQueryInput) (*timestreamquery.CancelQueryOutput, error)
	CancelQueryWithContext(aws.Context, *timestreamquery.CancelQueryInput, ...request.Option) (*timestreamquery.CancelQueryOutput, error)
	CancelQueryRequest(*timestreamquery.CancelQueryInput) (*request.Request, *timestreamquery.CancelQueryOutput)

	DescribeEndpoints(*timestreamquery.DescribeEndpointsInput) (*timestreamquery.DescribeEndpointsOutput, error)
	DescribeEndpointsWithContext(aws.Context, *timestreamquery.DescribeEndpointsInput, ...request.Option) (*timestreamquery.DescribeEndpointsOutput, error)
	DescribeEndpointsRequest(*timestreamquery.DescribeEndpointsInput) (*request.Request, *timestreamquery.DescribeEndpointsOutput)

	Query(*timestreamquery.QueryInput) (*timestreamquery.QueryOutput, error)
	QueryWithContext(aws.Context, *timestreamquery.QueryInput, ...request.Option) (*timestreamquery.QueryOutput, error)
	QueryRequest(*timestreamquery.QueryInput) (*request.Request, *timestreamquery.QueryOutput)

	QueryPages(*timestreamquery.QueryInput, func(*timestreamquery.QueryOutput, bool) bool) error
	QueryPagesWithContext(aws.Context, *timestreamquery.QueryInput, func(*timestreamquery.QueryOutput, bool) bool, ...request.Option) error
}

var _ TimestreamQueryAPI = (*timestreamquery.TimestreamQuery)(nil)
