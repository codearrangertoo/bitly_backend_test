# \CustomBitlinksApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomBitlink**](CustomBitlinksApi.md#AddCustomBitlink) | **Post** /custom_bitlinks | Add Custom Bitlink
[**GetClicksForCustomBitlink**](CustomBitlinksApi.md#GetClicksForCustomBitlink) | **Get** /custom_bitlinks/{custom_bitlink}/clicks | Get Clicks for a Custom Bitlink&#39;s Entire History
[**GetCustomBitlink**](CustomBitlinksApi.md#GetCustomBitlink) | **Get** /custom_bitlinks/{custom_bitlink} | Retrieve Custom Bitlink
[**GetCustomBitlinkMetricsByDestination**](CustomBitlinksApi.md#GetCustomBitlinkMetricsByDestination) | **Get** /custom_bitlinks/{custom_bitlink}/clicks_by_destination | Get Metrics for a Custom Bitlink by Destination
[**UpdateCustomBitlink**](CustomBitlinksApi.md#UpdateCustomBitlink) | **Patch** /custom_bitlinks/{custom_bitlink} | Update Custom Bitlink



## AddCustomBitlink

> CustomBitlink AddCustomBitlink(ctx).AddCustomBitlink(addCustomBitlink).Execute()

Add Custom Bitlink



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
    addCustomBitlink := *openapiclient.NewAddCustomBitlink() // AddCustomBitlink | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomBitlinksApi.AddCustomBitlink(context.Background()).AddCustomBitlink(addCustomBitlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomBitlinksApi.AddCustomBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomBitlink`: CustomBitlink
    fmt.Fprintf(os.Stdout, "Response from `CustomBitlinksApi.AddCustomBitlink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCustomBitlink** | [**AddCustomBitlink**](AddCustomBitlink.md) |  | 

### Return type

[**CustomBitlink**](CustomBitlink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClicksForCustomBitlink

> Clicks GetClicksForCustomBitlink(ctx, customBitlink).Unit(unit).Units(units).UnitReference(unitReference).Execute()

Get Clicks for a Custom Bitlink's Entire History



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
    customBitlink := "chauncey.ly/chauncey" // string | A Custom Bitlink made of the domain and keyword
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomBitlinksApi.GetClicksForCustomBitlink(context.Background(), customBitlink).Unit(unit).Units(units).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomBitlinksApi.GetClicksForCustomBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClicksForCustomBitlink`: Clicks
    fmt.Fprintf(os.Stdout, "Response from `CustomBitlinksApi.GetClicksForCustomBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customBitlink** | **string** | A Custom Bitlink made of the domain and keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClicksForCustomBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**Clicks**](Clicks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomBitlink

> CustomBitlink GetCustomBitlink(ctx, customBitlink).Execute()

Retrieve Custom Bitlink



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
    customBitlink := "chauncey.ly/chauncey" // string | A Custom Bitlink made of the domain and keyword

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomBitlinksApi.GetCustomBitlink(context.Background(), customBitlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomBitlinksApi.GetCustomBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomBitlink`: CustomBitlink
    fmt.Fprintf(os.Stdout, "Response from `CustomBitlinksApi.GetCustomBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customBitlink** | **string** | A Custom Bitlink made of the domain and keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomBitlink**](CustomBitlink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomBitlinkMetricsByDestination

> ClickMetrics GetCustomBitlinkMetricsByDestination(ctx, customBitlink).Unit(unit).Units(units).UnitReference(unitReference).Execute()

Get Metrics for a Custom Bitlink by Destination



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
    customBitlink := "chauncey.ly/chauncey" // string | A Custom Bitlink made of the domain and keyword
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomBitlinksApi.GetCustomBitlinkMetricsByDestination(context.Background(), customBitlink).Unit(unit).Units(units).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomBitlinksApi.GetCustomBitlinkMetricsByDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomBitlinkMetricsByDestination`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `CustomBitlinksApi.GetCustomBitlinkMetricsByDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customBitlink** | **string** | A Custom Bitlink made of the domain and keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomBitlinkMetricsByDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**ClickMetrics**](ClickMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomBitlink

> CustomBitlink UpdateCustomBitlink(ctx, customBitlink).UpdateCustomBitlink(updateCustomBitlink).Execute()

Update Custom Bitlink



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
    customBitlink := "chauncey.ly/chauncey" // string | A Custom Bitlink made of the domain and keyword
    updateCustomBitlink := *openapiclient.NewUpdateCustomBitlink() // UpdateCustomBitlink | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomBitlinksApi.UpdateCustomBitlink(context.Background(), customBitlink).UpdateCustomBitlink(updateCustomBitlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomBitlinksApi.UpdateCustomBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomBitlink`: CustomBitlink
    fmt.Fprintf(os.Stdout, "Response from `CustomBitlinksApi.UpdateCustomBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customBitlink** | **string** | A Custom Bitlink made of the domain and keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomBitlink** | [**UpdateCustomBitlink**](UpdateCustomBitlink.md) |  | 

### Return type

[**CustomBitlink**](CustomBitlink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

