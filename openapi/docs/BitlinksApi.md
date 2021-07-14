# \BitlinksApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBitlink**](BitlinksApi.md#CreateBitlink) | **Post** /shorten | Shorten a Link
[**CreateFullBitlink**](BitlinksApi.md#CreateFullBitlink) | **Post** /bitlinks | Create a Bitlink
[**ExpandBitlink**](BitlinksApi.md#ExpandBitlink) | **Post** /expand | Expand a Bitlink
[**GetBitlink**](BitlinksApi.md#GetBitlink) | **Get** /bitlinks/{bitlink} | Retrieve a Bitlink
[**GetBitlinkQRCode**](BitlinksApi.md#GetBitlinkQRCode) | **Get** /bitlinks/{bitlink}/qr | Get a QR Code
[**GetBitlinksByGroup**](BitlinksApi.md#GetBitlinksByGroup) | **Get** /groups/{group_guid}/bitlinks | Retrieve Bitlinks by Group
[**GetClicksForBitlink**](BitlinksApi.md#GetClicksForBitlink) | **Get** /bitlinks/{bitlink}/clicks | Get Clicks for a Bitlink
[**GetClicksSummaryForBitlink**](BitlinksApi.md#GetClicksSummaryForBitlink) | **Get** /bitlinks/{bitlink}/clicks/summary | Get a Clicks Summary for a Bitlink
[**GetMetricsForBitlinkByCities**](BitlinksApi.md#GetMetricsForBitlinkByCities) | **Get** /bitlinks/{bitlink}/cities | Get Metrics for a Bitlink by City
[**GetMetricsForBitlinkByCountries**](BitlinksApi.md#GetMetricsForBitlinkByCountries) | **Get** /bitlinks/{bitlink}/countries | Get Metrics for a Bitlink by Country
[**GetMetricsForBitlinkByDevices**](BitlinksApi.md#GetMetricsForBitlinkByDevices) | **Get** /bitlinks/{bitlink}/devices | Get Metrics for a Bitlink by Device Type
[**GetMetricsForBitlinkByReferrers**](BitlinksApi.md#GetMetricsForBitlinkByReferrers) | **Get** /bitlinks/{bitlink}/referrers | Get Metrics for a Bitlink by Referrers
[**GetMetricsForBitlinkByReferrersByDomains**](BitlinksApi.md#GetMetricsForBitlinkByReferrersByDomains) | **Get** /bitlinks/{bitlink}/referrers_by_domains | Get Metrics for a Bitlink by Referrers by Domain
[**GetMetricsForBitlinkByReferringDomains**](BitlinksApi.md#GetMetricsForBitlinkByReferringDomains) | **Get** /bitlinks/{bitlink}/referring_domains | Get Metrics for a Bitlink by Referring Domains
[**GetSortedBitlinks**](BitlinksApi.md#GetSortedBitlinks) | **Get** /groups/{group_guid}/bitlinks/{sort} | Retrieve Sorted Bitlinks for Group
[**UpdateBitlink**](BitlinksApi.md#UpdateBitlink) | **Patch** /bitlinks/{bitlink} | Update a Bitlink



## CreateBitlink

> ShortenBitlinkBody CreateBitlink(ctx).Shorten(shorten).Execute()

Shorten a Link



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
    shorten := *openapiclient.NewShorten("LongUrl_example") // Shorten | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.CreateBitlink(context.Background()).Shorten(shorten).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.CreateBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBitlink`: ShortenBitlinkBody
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.CreateBitlink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shorten** | [**Shorten**](Shorten.md) |  | 

### Return type

[**ShortenBitlinkBody**](ShortenBitlinkBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFullBitlink

> BitlinkBody CreateFullBitlink(ctx).FullShorten(fullShorten).Execute()

Create a Bitlink



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
    fullShorten := *openapiclient.NewFullShorten("LongUrl_example") // FullShorten | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.CreateFullBitlink(context.Background()).FullShorten(fullShorten).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.CreateFullBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFullBitlink`: BitlinkBody
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.CreateFullBitlink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFullBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullShorten** | [**FullShorten**](FullShorten.md) |  | 

### Return type

[**BitlinkBody**](BitlinkBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandBitlink

> ExpandedBitlink ExpandBitlink(ctx).ExpandBitlink(expandBitlink).Execute()

Expand a Bitlink



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
    expandBitlink := *openapiclient.NewExpandBitlink() // ExpandBitlink | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.ExpandBitlink(context.Background()).ExpandBitlink(expandBitlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.ExpandBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpandBitlink`: ExpandedBitlink
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.ExpandBitlink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpandBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expandBitlink** | [**ExpandBitlink**](ExpandBitlink.md) |  | 

### Return type

[**ExpandedBitlink**](ExpandedBitlink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBitlink

> BitlinkBody GetBitlink(ctx, bitlink).Execute()

Retrieve a Bitlink



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetBitlink(context.Background(), bitlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBitlink`: BitlinkBody
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BitlinkBody**](BitlinkBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBitlinkQRCode

> BitlinkQR GetBitlinkQRCode(ctx, bitlink).Color(color).ExcludeBitlyLogo(excludeBitlyLogo).ImageFormat(imageFormat).Execute()

Get a QR Code



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    color := "1133ff" // string | A color denoted in hexidecimal RGB format (optional)
    excludeBitlyLogo := false // bool | Removes the Bitly logo from the center of the generated QR image (optional)
    imageFormat := "svg" // string | Determines the image format of the returned QR code (optional) (default to "png")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetBitlinkQRCode(context.Background(), bitlink).Color(color).ExcludeBitlyLogo(excludeBitlyLogo).ImageFormat(imageFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetBitlinkQRCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBitlinkQRCode`: BitlinkQR
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetBitlinkQRCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBitlinkQRCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **color** | **string** | A color denoted in hexidecimal RGB format | 
 **excludeBitlyLogo** | **bool** | Removes the Bitly logo from the center of the generated QR image | 
 **imageFormat** | **string** | Determines the image format of the returned QR code | [default to &quot;png&quot;]

### Return type

[**BitlinkQR**](BitlinkQR.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBitlinksByGroup

> Bitlinks GetBitlinksByGroup(ctx, groupGuid).Size(size).Page(page).Keyword(keyword).Query(query).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ModifiedAfter(modifiedAfter).Archived(archived).Deeplinks(deeplinks).DomainDeeplinks(domainDeeplinks).CampaignGuid(campaignGuid).ChannelGuid(channelGuid).CustomBitlink(customBitlink).Tags(tags).LaunchpadIds(launchpadIds).EncodingLogin(encodingLogin).Execute()

Retrieve Bitlinks by Group



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
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    page := int32(1) // int32 | Integer specifying the numbered result at which to start (optional) (default to 1)
    keyword := "chauncey" // string | Custom keyword to filter on history entries (optional)
    query := "api" // string | the value that you would like to search (optional)
    createdBefore := int32(1501027200) // int32 | Timestamp as an integer unix epoch (optional)
    createdAfter := int32(1501027200) // int32 | Timestamp as an integer unix epoch (optional)
    modifiedAfter := int32(1501027200) // int32 | Timestamp as an integer unix epoch (optional)
    archived := "both" // string | Whether or not to include archived bitlinks (optional) (default to "off")
    deeplinks := "both" // string | Filter to only Bitlinks that contain deeplinks (optional) (default to "both")
    domainDeeplinks := "both" // string | Filter to only Bitlinks that contain deeplinks configured with a custom domain (optional) (default to "both")
    campaignGuid := "Ca1bcd2EFGh" // string | Filter to return only links for the given campaign GUID, can be provided (optional)
    channelGuid := "Ha1bc2DefGh" // string | Filter to return only links for the given channel GUID, can be provided, overrides all other parameters (optional)
    customBitlink := "both" // string |  (optional) (default to "both")
    tags := []string{"Inner_example"} // []string | filter by given tags (optional)
    launchpadIds := []string{"Inner_example"} // []string | filter by launchpad id (optional)
    encodingLogin := []string{"Inner_example"} // []string | Filter by the login of the authenticated user that created the Bitlink (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetBitlinksByGroup(context.Background(), groupGuid).Size(size).Page(page).Keyword(keyword).Query(query).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ModifiedAfter(modifiedAfter).Archived(archived).Deeplinks(deeplinks).DomainDeeplinks(domainDeeplinks).CampaignGuid(campaignGuid).ChannelGuid(channelGuid).CustomBitlink(customBitlink).Tags(tags).LaunchpadIds(launchpadIds).EncodingLogin(encodingLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetBitlinksByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBitlinksByGroup`: Bitlinks
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetBitlinksByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBitlinksByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **page** | **int32** | Integer specifying the numbered result at which to start | [default to 1]
 **keyword** | **string** | Custom keyword to filter on history entries | 
 **query** | **string** | the value that you would like to search | 
 **createdBefore** | **int32** | Timestamp as an integer unix epoch | 
 **createdAfter** | **int32** | Timestamp as an integer unix epoch | 
 **modifiedAfter** | **int32** | Timestamp as an integer unix epoch | 
 **archived** | **string** | Whether or not to include archived bitlinks | [default to &quot;off&quot;]
 **deeplinks** | **string** | Filter to only Bitlinks that contain deeplinks | [default to &quot;both&quot;]
 **domainDeeplinks** | **string** | Filter to only Bitlinks that contain deeplinks configured with a custom domain | [default to &quot;both&quot;]
 **campaignGuid** | **string** | Filter to return only links for the given campaign GUID, can be provided | 
 **channelGuid** | **string** | Filter to return only links for the given channel GUID, can be provided, overrides all other parameters | 
 **customBitlink** | **string** |  | [default to &quot;both&quot;]
 **tags** | **[]string** | filter by given tags | 
 **launchpadIds** | **[]string** | filter by launchpad id | 
 **encodingLogin** | **[]string** | Filter by the login of the authenticated user that created the Bitlink | 

### Return type

[**Bitlinks**](Bitlinks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClicksForBitlink

> Clicks GetClicksForBitlink(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Clicks for a Bitlink



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetClicksForBitlink(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetClicksForBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClicksForBitlink`: Clicks
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetClicksForBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClicksForBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
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


## GetClicksSummaryForBitlink

> ClicksSummary GetClicksSummaryForBitlink(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get a Clicks Summary for a Bitlink



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetClicksSummaryForBitlink(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetClicksSummaryForBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClicksSummaryForBitlink`: ClicksSummary
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetClicksSummaryForBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClicksSummaryForBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**ClicksSummary**](ClicksSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsForBitlinkByCities

> CityMetrics GetMetricsForBitlinkByCities(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by City



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByCities(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByCities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByCities`: CityMetrics
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByCities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByCitiesRequest struct via the builder pattern


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


## GetMetricsForBitlinkByCountries

> ClickMetrics GetMetricsForBitlinkByCountries(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by Country



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByCountries(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByCountries`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByCountries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByCountriesRequest struct via the builder pattern


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
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsForBitlinkByDevices

> DeviceMetrics GetMetricsForBitlinkByDevices(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by Device Type



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByDevices(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByDevices`: DeviceMetrics
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByDevicesRequest struct via the builder pattern


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


## GetMetricsForBitlinkByReferrers

> ClickMetrics GetMetricsForBitlinkByReferrers(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by Referrers



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByReferrers(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByReferrers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByReferrers`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByReferrers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByReferrersRequest struct via the builder pattern


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
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsForBitlinkByReferrersByDomains

> ReferrersByDomains GetMetricsForBitlinkByReferrersByDomains(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by Referrers by Domain



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByReferrersByDomains(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByReferrersByDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByReferrersByDomains`: ReferrersByDomains
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByReferrersByDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByReferrersByDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **size** | **int32** | The quantity of items to be be returned | [default to 50]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 

### Return type

[**ReferrersByDomains**](ReferrersByDomains.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsForBitlinkByReferringDomains

> ClickMetrics GetMetricsForBitlinkByReferringDomains(ctx, bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()

Get Metrics for a Bitlink by Referring Domains



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    unit := "month" // string | A unit of time (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (default to -1)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetMetricsForBitlinkByReferringDomains(context.Background(), bitlink).Unit(unit).Units(units).Size(size).UnitReference(unitReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetMetricsForBitlinkByReferringDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsForBitlinkByReferringDomains`: ClickMetrics
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetMetricsForBitlinkByReferringDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsForBitlinkByReferringDomainsRequest struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSortedBitlinks

> SortedLinks GetSortedBitlinks(ctx, groupGuid, sort).Unit(unit).Units(units).UnitReference(unitReference).Size(size).Execute()

Retrieve Sorted Bitlinks for Group



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
    sort := "clicks" // string | The type of sorting that you would like to do
    unit := "month" // string | A unit of time (optional) (default to "day")
    units := int32(5) // int32 | An integer representing the time units to query data for. pass -1 to return all units of time. (optional) (default to -1)
    unitReference := "2006-01-02T15:04:05-0700" // string | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. (optional)
    size := int32(10) // int32 | The quantity of items to be be returned (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.GetSortedBitlinks(context.Background(), groupGuid, sort).Unit(unit).Units(units).UnitReference(unitReference).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.GetSortedBitlinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSortedBitlinks`: SortedLinks
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.GetSortedBitlinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupGuid** | **string** | A GUID for a Bitly group | 
**sort** | **string** | The type of sorting that you would like to do | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSortedBitlinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unit** | **string** | A unit of time | [default to &quot;day&quot;]
 **units** | **int32** | An integer representing the time units to query data for. pass -1 to return all units of time. | [default to -1]
 **unitReference** | **string** | An ISO-8601 timestamp, indicating the most recent time for which to pull metrics. Will default to current time. | 
 **size** | **int32** | The quantity of items to be be returned | [default to 50]

### Return type

[**SortedLinks**](SortedLinks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBitlink

> BitlinkBody UpdateBitlink(ctx, bitlink).BitlinkBody(bitlinkBody).Execute()

Update a Bitlink



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
    bitlink := "bit.ly/12a4b6c" // string | A Bitlink made of the domain and hash
    bitlinkBody := *openapiclient.NewBitlinkBody() // BitlinkBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BitlinksApi.UpdateBitlink(context.Background(), bitlink).BitlinkBody(bitlinkBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BitlinksApi.UpdateBitlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBitlink`: BitlinkBody
    fmt.Fprintf(os.Stdout, "Response from `BitlinksApi.UpdateBitlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlink** | **string** | A Bitlink made of the domain and hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBitlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bitlinkBody** | [**BitlinkBody**](BitlinkBody.md) |  | 

### Return type

[**BitlinkBody**](BitlinkBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

