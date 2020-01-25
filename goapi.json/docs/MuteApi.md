# \MuteApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MuteCreate**](MuteApi.md#MuteCreate) | **Post** /mute/create | mute/create
[**MuteDelete**](MuteApi.md#MuteDelete) | **Post** /mute/delete | mute/delete
[**MuteList**](MuteApi.md#MuteList) | **Post** /mute/list | mute/list



## MuteCreate

> MuteCreate(ctx, inlineObject90)

mute/create

ユーザーをミュートします。  **Credential required**: *Yes* / **Permission**: *account/write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject90** | [**InlineObject90**](InlineObject90.md)|  | 

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


## MuteDelete

> MuteDelete(ctx, inlineObject91)

mute/delete

ユーザーのミュートを解除します。  **Credential required**: *Yes* / **Permission**: *account/write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject91** | [**InlineObject91**](InlineObject91.md)|  | 

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


## MuteList

> []Muting MuteList(ctx, inlineObject92)

mute/list

ミュートしているユーザー一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account/read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject92** | [**InlineObject92**](InlineObject92.md)|  | 

### Return type

[**[]Muting**](Muting.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

