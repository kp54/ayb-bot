# \FederationApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApShow**](FederationApi.md#ApShow) | **Post** /ap/show | ap/show
[**FederationInstances**](FederationApi.md#FederationInstances) | **Post** /federation/instances | federation/instances
[**FederationShowInstance**](FederationApi.md#FederationShowInstance) | **Post** /federation/show-instance | federation/show-instance



## ApShow

> ApShow(ctx, inlineObject25)

ap/show

URIを指定してActivityPubオブジェクトを参照します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject25** | [**InlineObject25**](InlineObject25.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FederationInstances

> FederationInstances(ctx, inlineObject65)

federation/instances

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject65** | [**InlineObject65**](InlineObject65.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FederationShowInstance

> FederationShowInstance(ctx, inlineObject66)

federation/show-instance

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject66** | [**InlineObject66**](InlineObject66.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

