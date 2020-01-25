# \FavoritesApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IFavorites**](FavoritesApi.md#IFavorites) | **Post** /i/favorites | i/favorites
[**NotesFavoritesCreate**](FavoritesApi.md#NotesFavoritesCreate) | **Post** /notes/favorites/create | notes/favorites/create
[**NotesFavoritesDelete**](FavoritesApi.md#NotesFavoritesDelete) | **Post** /notes/favorites/delete | notes/favorites/delete



## IFavorites

> IFavorites(ctx, inlineObject79)

i/favorites

お気に入りに登録した投稿一覧を取得します。  **Credential required**: *Yes* / **Permission**: *favorites-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject79** | [**InlineObject79**](InlineObject79.md)|  | 

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


## NotesFavoritesCreate

> NotesFavoritesCreate(ctx, inlineObject99)

notes/favorites/create

指定した投稿をお気に入りに登録します。  **Credential required**: *Yes* / **Permission**: *favorite-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject99** | [**InlineObject99**](InlineObject99.md)|  | 

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


## NotesFavoritesDelete

> NotesFavoritesDelete(ctx, inlineObject100)

notes/favorites/delete

指定した投稿のお気に入りを解除します。  **Credential required**: *Yes* / **Permission**: *favorite-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject100** | [**InlineObject100**](InlineObject100.md)|  | 

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

