# \CampaignsApi

All URIs are relative to *https://api-ssl.bitly.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsApi.md#CreateCampaign) | **Post** /campaigns | Create Campaign
[**CreateChannel**](CampaignsApi.md#CreateChannel) | **Post** /channels | Create Channel
[**GetCampaign**](CampaignsApi.md#GetCampaign) | **Get** /campaigns/{campaign_guid} | Retrieve a Campaign
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | Retrieve Campaigns
[**GetChannel**](CampaignsApi.md#GetChannel) | **Get** /channels/{channel_guid} | Get a Channel
[**GetChannels**](CampaignsApi.md#GetChannels) | **Get** /channels | Retrieve Channels
[**UpdateCampaign**](CampaignsApi.md#UpdateCampaign) | **Patch** /campaigns/{campaign_guid} | Update Campaign
[**UpdateChannel**](CampaignsApi.md#UpdateChannel) | **Patch** /channels/{channel_guid} | Update a Channel



## CreateCampaign

> Campaign CreateCampaign(ctx).CampaignModify(campaignModify).Execute()

Create Campaign



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
    campaignModify := *openapiclient.NewCampaignModify() // CampaignModify | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.CreateCampaign(context.Background()).CampaignModify(campaignModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignModify** | [**CampaignModify**](CampaignModify.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannel

> Channel CreateChannel(ctx).ChannelModify(channelModify).Execute()

Create Channel



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
    channelModify := *openapiclient.NewChannelModify() // ChannelModify | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.CreateChannel(context.Background()).ChannelModify(channelModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: Channel
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelModify** | [**ChannelModify**](ChannelModify.md) |  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, campaignGuid).Execute()

Retrieve a Campaign



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
    campaignGuid := "Ca1bcd2EFGh" // string | A GUID for a Bitly campaign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.GetCampaign(context.Background(), campaignGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | A GUID for a Bitly campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> Campaigns GetCampaigns(ctx).GroupGuid(groupGuid).Execute()

Retrieve Campaigns



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.GetCampaigns(context.Background()).GroupGuid(groupGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaigns`: Campaigns
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupGuid** | **string** | A GUID for a Bitly group | 

### Return type

[**Campaigns**](Campaigns.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> Channel GetChannel(ctx, channelGuid).Execute()

Get a Channel



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
    channelGuid := "Ha1bc2DefGh" // string | A GUID for a Bitly Channel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.GetChannel(context.Background(), channelGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannel`: Channel
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelGuid** | **string** | A GUID for a Bitly Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Channel**](Channel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> Channels GetChannels(ctx).GroupGuid(groupGuid).CampaignGuid(campaignGuid).Execute()

Retrieve Channels



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
    groupGuid := "Ba1bc23dE4F" // string | A GUID for a Bitly group (optional)
    campaignGuid := "Ca1bcd2EFGh" // string | A GUID for a Bitly campaign (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.GetChannels(context.Background()).GroupGuid(groupGuid).CampaignGuid(campaignGuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannels`: Channels
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupGuid** | **string** | A GUID for a Bitly group | 
 **campaignGuid** | **string** | A GUID for a Bitly campaign | 

### Return type

[**Channels**](Channels.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, campaignGuid).CampaignModify(campaignModify).Execute()

Update Campaign



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
    campaignGuid := "Ca1bcd2EFGh" // string | A GUID for a Bitly campaign
    campaignModify := *openapiclient.NewCampaignModify() // CampaignModify | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.UpdateCampaign(context.Background(), campaignGuid).CampaignModify(campaignModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaign`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGuid** | **string** | A GUID for a Bitly campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignModify** | [**CampaignModify**](CampaignModify.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> Channel UpdateChannel(ctx, channelGuid).ChannelModify(channelModify).Execute()

Update a Channel



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
    channelGuid := "Ha1bc2DefGh" // string | A GUID for a Bitly Channel
    channelModify := *openapiclient.NewChannelModify() // ChannelModify | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignsApi.UpdateChannel(context.Background(), channelGuid).ChannelModify(channelModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.UpdateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannel`: Channel
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelGuid** | **string** | A GUID for a Bitly Channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelModify** | [**ChannelModify**](ChannelModify.md) |  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

