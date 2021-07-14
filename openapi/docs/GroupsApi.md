# \GroupsApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /groups/{group_guid} | Retrieve a Group
[**GetGroupClicks**](GroupsApi.md#GetGroupClicks) | **Get** /groups/{group_guid}/clicks | Get clicks by group
[**GetGroupMetricsByCities**](GroupsApi.md#GetGroupMetricsByCities) | **Get** /groups/{group_guid}/cities | Get Click Metrics for a Group by City
[**GetGroupMetricsByCountries**](GroupsApi.md#GetGroupMetricsByCountries) | **Get** /groups/{group_guid}/countries | Get Click Metrics for a Group by Country
[**GetGroupMetricsByDevices**](GroupsApi.md#GetGroupMetricsByDevices) | **Get** /groups/{group_guid}/devices | Get Click Metrics for a Group by Device Type
[**GetGroupMetricsByReferringNetworks**](GroupsApi.md#GetGroupMetricsByReferringNetworks) | **Get** /groups/{group_guid}/referring_networks | Get Click Metrics for a Group by Referring Networks
[**GetGroupPreferences**](GroupsApi.md#GetGroupPreferences) | **Get** /groups/{group_guid}/preferences | Retrieve Group Preferences
[**GetGroupShortenCounts**](GroupsApi.md#GetGroupShortenCounts) | **Get** /groups/{group_guid}/shorten_counts | Retrieve Group Shorten Counts
[**GetGroupTags**](GroupsApi.md#GetGroupTags) | **Get** /groups/{group_guid}/tags | Retrieve Tags by Group
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /groups | Retrieve Groups
[**UpdateGroup**](GroupsApi.md#UpdateGroup) | **Patch** /groups/{group_guid} | Update a Group
[**UpdateGroupPreferences**](GroupsApi.md#UpdateGroupPreferences) | **Patch** /groups/{group_guid}/preferences | Update Group Preferences



## GetGroup

> Group GetGroup(ctx, groupGuid).Execute()

Retrieve a Group



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroup(context.Background(), groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupClicks

> GroupClicks GetGroupClicks(ctx, groupGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()

Get clicks by group



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupClicks(context.Background(), groupGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupClicks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupClicks`: GroupClicks
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupClicks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupClicksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**GroupClicks**](GroupClicks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMetricsByCities

> CityMetrics GetGroupMetricsByCities(ctx, groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Click Metrics for a Group by City



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupMetricsByCities(context.Background(), groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMetricsByCities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMetricsByCities`: CityMetrics
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMetricsByCities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMetricsByCitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**CityMetrics**](CityMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMetricsByCountries

> ClickMetrics GetGroupMetricsByCountries(ctx, groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Click Metrics for a Group by Country



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupMetricsByCountries(context.Background(), groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMetricsByCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMetricsByCountries`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMetricsByCountries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMetricsByCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**ClickMetrics**](ClickMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMetricsByDevices

> DeviceMetrics GetGroupMetricsByDevices(ctx, groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Click Metrics for a Group by Device Type



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupMetricsByDevices(context.Background(), groupGuid).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMetricsByDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMetricsByDevices`: DeviceMetrics
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMetricsByDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMetricsByDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**DeviceMetrics**](DeviceMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMetricsByReferringNetworks

> ClickMetrics GetGroupMetricsByReferringNetworks(ctx, groupGuid).Execute()

Get Click Metrics for a Group by Referring Networks



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupMetricsByReferringNetworks(context.Background(), groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMetricsByReferringNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMetricsByReferringNetworks`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMetricsByReferringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMetricsByReferringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetGroupPreferences

> GroupPreferences GetGroupPreferences(ctx, groupGuid).Execute()

Retrieve Group Preferences



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupPreferences(context.Background(), groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupPreferences`: GroupPreferences
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupPreferences**](GroupPreferences.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupShortenCounts

> Metrics GetGroupShortenCounts(ctx, groupGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()

Retrieve Group Shorten Counts



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupShortenCounts(context.Background(), groupGuid).Unit(unit).Units(units).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupShortenCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupShortenCounts`: Metrics
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupShortenCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupShortenCountsRequest struct via the builder pattern


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


## GetGroupTags

> Tags GetGroupTags(ctx, groupGuid).Execute()

Retrieve Tags by Group



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroupTags(context.Background(), groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupTags`: Tags
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tags**](Tags.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> Groups GetGroups(ctx).OrganizationGuid(organizationGuid).Execute()

Retrieve Groups



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
    organizationGuid := "Oa1bcd234eF" // string | A GUID for a Bitly organization (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.GetGroups(context.Background()).OrganizationGuid(organizationGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: Groups
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationGuid** | **string** | A GUID for a Bitly organization | 

### Return type

[**Groups**](Groups.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupGuid).GroupUpdate(groupUpdate).Execute()

Update a Group



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    groupUpdate := *openapiclient.NewGroupUpdate() // GroupUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.UpdateGroup(context.Background(), groupGuid).GroupUpdate(groupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdate** | [**GroupUpdate**](GroupUpdate.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupPreferences

> GroupPreferences UpdateGroupPreferences(ctx, groupGuid).GroupPreferences(groupPreferences).Execute()

Update Group Preferences



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group
    groupPreferences := *openapiclient.NewGroupPreferences() // GroupPreferences | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsApi.UpdateGroupPreferences(context.Background(), groupGuid).GroupPreferences(groupPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroupPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupPreferences`: GroupPreferences
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroupPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupPreferences** | [**GroupPreferences**](GroupPreferences.md) |  | 

### Return type

[**GroupPreferences**](GroupPreferences.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

