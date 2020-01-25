# \ListsApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotesUserListTimeline**](ListsApi.md#NotesUserListTimeline) | **Post** /notes/user-list-timeline | notes/user-list-timeline
[**UsersListsCreate**](ListsApi.md#UsersListsCreate) | **Post** /users/lists/create | users/lists/create
[**UsersListsDelete**](ListsApi.md#UsersListsDelete) | **Post** /users/lists/delete | users/lists/delete
[**UsersListsList**](ListsApi.md#UsersListsList) | **Post** /users/lists/list | users/lists/list
[**UsersListsPull**](ListsApi.md#UsersListsPull) | **Post** /users/lists/pull | users/lists/pull
[**UsersListsPush**](ListsApi.md#UsersListsPush) | **Post** /users/lists/push | users/lists/push
[**UsersListsShow**](ListsApi.md#UsersListsShow) | **Post** /users/lists/show | users/lists/show
[**UsersListsUpdate**](ListsApi.md#UsersListsUpdate) | **Post** /users/lists/update | users/lists/update



## NotesUserListTimeline

> []Note NotesUserListTimeline(ctx, inlineObject119)

notes/user-list-timeline

指定したユーザーリストのタイムラインを取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject119** | [**InlineObject119**](InlineObject119.md)|  | 

### Return type

[**[]Note**](Note.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListsCreate

> UsersListsCreate(ctx, inlineObject129)

users/lists/create

ユーザーリストを作成します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject129** | [**InlineObject129**](InlineObject129.md)|  | 

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


## UsersListsDelete

> UsersListsDelete(ctx, inlineObject130)

users/lists/delete

指定したユーザーリストを削除します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject130** | [**InlineObject130**](InlineObject130.md)|  | 

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


## UsersListsList

> []UserList UsersListsList(ctx, body)

users/lists/list

自分の作成したユーザーリスト一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**[]UserList**](UserList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListsPull

> UsersListsPull(ctx, inlineObject131)

users/lists/pull

指定したユーザーリストから指定したユーザーを削除します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject131** | [**InlineObject131**](InlineObject131.md)|  | 

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


## UsersListsPush

> UsersListsPush(ctx, inlineObject132)

users/lists/push

指定したユーザーリストに指定したユーザーを追加します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject132** | [**InlineObject132**](InlineObject132.md)|  | 

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


## UsersListsShow

> UserList UsersListsShow(ctx, inlineObject133)

users/lists/show

指定したユーザーリストの情報を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject133** | [**InlineObject133**](InlineObject133.md)|  | 

### Return type

[**UserList**](UserList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListsUpdate

> UsersListsUpdate(ctx, inlineObject134)

users/lists/update

指定したユーザーリストを更新します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject134** | [**InlineObject134**](InlineObject134.md)|  | 

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

