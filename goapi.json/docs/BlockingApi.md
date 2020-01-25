# \BlockingApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockingCreate**](BlockingApi.md#BlockingCreate) | **Post** /blocking/create | blocking/create
[**BlockingDelete**](BlockingApi.md#BlockingDelete) | **Post** /blocking/delete | blocking/delete
[**BlockingList**](BlockingApi.md#BlockingList) | **Post** /blocking/list | blocking/list



## BlockingCreate

> BlockingCreate(ctx, inlineObject31)

blocking/create

指定したユーザーをブロックします。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject31** | [**InlineObject31**](InlineObject31.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockingDelete

> BlockingDelete(ctx, inlineObject32)

blocking/delete

指定したユーザーのブロックを解除します。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject32** | [**InlineObject32**](InlineObject32.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockingList

> []Blocking BlockingList(ctx, inlineObject33)

blocking/list

ブロックしているユーザー一覧を取得します。  **Credential required**: *Yes* / **Permission**: *following-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject33** | [**InlineObject33**](InlineObject33.md)|  | 

### Return type

[**[]Blocking**](Blocking.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

