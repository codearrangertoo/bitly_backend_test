# \OrganizationsApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /organizations/{organization_guid} | Retrieve an Organization
[**GetOrganizationShortenCounts**](OrganizationsApi.md#GetOrganizationShortenCounts) | **Get** /organizations/{organization_guid}/shorten_counts | Get Shorten Counts for an Organization
[**GetOrganizations**](OrganizationsApi.md#GetOrganizations) | **Get** /organizations | Retrieve Organizations
[**GetPlanLimits**](OrganizationsApi.md#GetPlanLimits) | **Get** /organizations/{organization_guid}/plan_limits | Get Plan Limits



## GetOrganization

> Organization GetOrganization(ctx, organizationGuid).Execute()

Retrieve an Organization



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
    organizationGuid := "Oa1bcd234eF" // string | A GUID for a Bitly organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetOrganization(context.Background(), organizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationGuid** | **string** | A GUID for a Bitly organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationShortenCounts

> Metrics GetOrganizationShortenCounts(ctx, organizationGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()

Get Shorten Counts for an Organization



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
    organizationGuid := "Oa1bcd234eF" // string | A GUID for a Bitly organization
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetOrganizationShortenCounts(context.Background(), organizationGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationShortenCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationShortenCounts`: Metrics
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationShortenCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationGuid** | **string** | A GUID for a Bitly organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationShortenCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> Organizations GetOrganizations(ctx).Execute()

Retrieve Organizations



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
    resp, r, err := api_client.OrganizationsApi.GetOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: Organizations
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


### Return type

[**Organizations**](Organizations.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanLimits

> PlanLimits GetPlanLimits(ctx, organizationGuid).Execute()

Get Plan Limits



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
    organizationGuid := "Oa1bcd234eF" // string | A GUID for a Bitly organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetPlanLimits(context.Background(), organizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetPlanLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanLimits`: PlanLimits
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetPlanLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationGuid** | **string** | A GUID for a Bitly organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlanLimits**](PlanLimits.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

