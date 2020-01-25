# \ReactionsApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartsUserReactions**](ReactionsApi.md#ChartsUserReactions) | **Post** /charts/user/reactions | charts/user/reactions
[**NotesReactions**](ReactionsApi.md#NotesReactions) | **Post** /notes/reactions | notes/reactions
[**NotesReactionsCreate**](ReactionsApi.md#NotesReactionsCreate) | **Post** /notes/reactions/create | notes/reactions/create
[**NotesReactionsDelete**](ReactionsApi.md#NotesReactionsDelete) | **Post** /notes/reactions/delete | notes/reactions/delete



## ChartsUserReactions

> []map[string]interface{} ChartsUserReactions(ctx, inlineObject44)

charts/user/reactions

ユーザーごとの被リアクション数のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject44** | [**InlineObject44**](InlineObject44.md)|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesReactions

> []Reaction NotesReactions(ctx, inlineObject108)

notes/reactions

指定した投稿のリアクション一覧を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject108** | [**InlineObject108**](InlineObject108.md)|  | 

### Return type

[**[]Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesReactionsCreate

> NotesReactionsCreate(ctx, inlineObject109)

notes/reactions/create

指定した投稿にリアクションします。  **Credential required**: *Yes* / **Permission**: *reaction-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject109** | [**InlineObject109**](InlineObject109.md)|  | 

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


## NotesReactionsDelete

> NotesReactionsDelete(ctx, inlineObject110)

notes/reactions/delete

指定した投稿へのリアクションを取り消します。  **Credential required**: *Yes* / **Permission**: *reaction-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject110** | [**InlineObject110**](InlineObject110.md)|  | 

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

