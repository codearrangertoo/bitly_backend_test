# \BSDsApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBSDs**](BSDsApi.md#GetBSDs) | **Get** /bsds | Get BSDs



## GetBSDs

> BSDsResponse GetBSDs(ctx).Execute()

Get BSDs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BSDsApi.GetBSDs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BSDsApi.GetBSDs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBSDs`: BSDsResponse
    fmt.Fprintf(os.Stdout, "Response from `BSDsApi.GetBSDs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBSDsRequest struct via the builder pattern


### Return type

[**BSDsResponse**](BSDsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

