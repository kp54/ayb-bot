# \DriveApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartsDrive**](DriveApi.md#ChartsDrive) | **Post** /charts/drive | charts/drive
[**ChartsUserDrive**](DriveApi.md#ChartsUserDrive) | **Post** /charts/user/drive | charts/user/drive
[**Drive**](DriveApi.md#Drive) | **Post** /drive | drive
[**DriveFiles**](DriveApi.md#DriveFiles) | **Post** /drive/files | drive/files
[**DriveFilesAttachedNotes**](DriveApi.md#DriveFilesAttachedNotes) | **Post** /drive/files/attached_notes | drive/files/attached_notes
[**DriveFilesAttachedNotes_0**](DriveApi.md#DriveFilesAttachedNotes_0) | **Post** /drive/files/attached-notes | drive/files/attached-notes
[**DriveFilesCheckExistence**](DriveApi.md#DriveFilesCheckExistence) | **Post** /drive/files/check_existence | drive/files/check_existence
[**DriveFilesCheckExistence_0**](DriveApi.md#DriveFilesCheckExistence_0) | **Post** /drive/files/check-existence | drive/files/check-existence
[**DriveFilesCreate**](DriveApi.md#DriveFilesCreate) | **Post** /drive/files/create | drive/files/create
[**DriveFilesDelete**](DriveApi.md#DriveFilesDelete) | **Post** /drive/files/delete | drive/files/delete
[**DriveFilesFind**](DriveApi.md#DriveFilesFind) | **Post** /drive/files/find | drive/files/find
[**DriveFilesShow**](DriveApi.md#DriveFilesShow) | **Post** /drive/files/show | drive/files/show
[**DriveFilesUpdate**](DriveApi.md#DriveFilesUpdate) | **Post** /drive/files/update | drive/files/update
[**DriveFilesUploadFromUrl**](DriveApi.md#DriveFilesUploadFromUrl) | **Post** /drive/files/upload_from_url | drive/files/upload_from_url
[**DriveFilesUploadFromUrl_0**](DriveApi.md#DriveFilesUploadFromUrl_0) | **Post** /drive/files/upload-from-url | drive/files/upload-from-url
[**DriveFolders**](DriveApi.md#DriveFolders) | **Post** /drive/folders | drive/folders
[**DriveFoldersCreate**](DriveApi.md#DriveFoldersCreate) | **Post** /drive/folders/create | drive/folders/create
[**DriveFoldersDelete**](DriveApi.md#DriveFoldersDelete) | **Post** /drive/folders/delete | drive/folders/delete
[**DriveFoldersFind**](DriveApi.md#DriveFoldersFind) | **Post** /drive/folders/find | drive/folders/find
[**DriveFoldersShow**](DriveApi.md#DriveFoldersShow) | **Post** /drive/folders/show | drive/folders/show
[**DriveFoldersUpdate**](DriveApi.md#DriveFoldersUpdate) | **Post** /drive/folders/update | drive/folders/update
[**DriveStream**](DriveApi.md#DriveStream) | **Post** /drive/stream | drive/stream



## ChartsDrive

> InlineResponse2002 ChartsDrive(ctx, inlineObject35)

charts/drive

ドライブのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject35** | [**InlineObject35**](InlineObject35.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserDrive

> InlineResponse2002Local ChartsUserDrive(ctx, inlineObject41)

charts/user/drive

ユーザーごとのドライブのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject41** | [**InlineObject41**](InlineObject41.md)|  | 

### Return type

[**InlineResponse2002Local**](inline_response_200_2_local.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Drive

> InlineResponse2006 Drive(ctx, body)

drive

ドライブの情報を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFiles

> []DriveFile DriveFiles(ctx, inlineObject46)

drive/files

ドライブのファイル一覧を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject46** | [**InlineObject46**](InlineObject46.md)|  | 

### Return type

[**[]DriveFile**](DriveFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFilesAttachedNotes

> DriveFilesAttachedNotes(ctx, inlineObject47)

drive/files/attached_notes

指定したドライブのファイルが添付されている投稿一覧を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject47** | [**InlineObject47**](InlineObject47.md)|  | 

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


## DriveFilesAttachedNotes_0

> []Note DriveFilesAttachedNotes_0(ctx, inlineObject48)

drive/files/attached-notes

指定したドライブのファイルが添付されている投稿一覧を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject48** | [**InlineObject48**](InlineObject48.md)|  | 

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


## DriveFilesCheckExistence

> DriveFilesCheckExistence(ctx, inlineObject49)

drive/files/check_existence

与えられたMD5ハッシュ値を持つファイルがドライブに存在するかどうかを返します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject49** | [**InlineObject49**](InlineObject49.md)|  | 

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


## DriveFilesCheckExistence_0

> DriveFile DriveFilesCheckExistence_0(ctx, inlineObject50)

drive/files/check-existence

与えられたMD5ハッシュ値を持つファイルがドライブに存在するかどうかを返します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject50** | [**InlineObject50**](InlineObject50.md)|  | 

### Return type

[**DriveFile**](DriveFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFilesCreate

> DriveFile DriveFilesCreate(ctx, inlineObject51)

drive/files/create

ドライブにファイルをアップロードします。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject51** | [**InlineObject51**](InlineObject51.md)|  | 

### Return type

[**DriveFile**](DriveFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFilesDelete

> DriveFilesDelete(ctx, inlineObject52)

drive/files/delete

ドライブのファイルを削除します。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject52** | [**InlineObject52**](InlineObject52.md)|  | 

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


## DriveFilesFind

> DriveFilesFind(ctx, inlineObject53)

drive/files/find

No description provided.  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject53** | [**InlineObject53**](InlineObject53.md)|  | 

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


## DriveFilesShow

> DriveFile DriveFilesShow(ctx, inlineObject54)

drive/files/show

指定したドライブのファイルの情報を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject54** | [**InlineObject54**](InlineObject54.md)|  | 

### Return type

[**DriveFile**](DriveFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFilesUpdate

> DriveFilesUpdate(ctx, inlineObject55)

drive/files/update

指定したドライブのファイルの情報を更新します。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject55** | [**InlineObject55**](InlineObject55.md)|  | 

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


## DriveFilesUploadFromUrl

> DriveFilesUploadFromUrl(ctx, inlineObject56)

drive/files/upload_from_url

ドライブに指定されたURLに存在するファイルをアップロードします。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject56** | [**InlineObject56**](InlineObject56.md)|  | 

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


## DriveFilesUploadFromUrl_0

> DriveFilesUploadFromUrl_0(ctx, inlineObject57)

drive/files/upload-from-url

ドライブに指定されたURLに存在するファイルをアップロードします。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject57** | [**InlineObject57**](InlineObject57.md)|  | 

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


## DriveFolders

> []DriveFolder DriveFolders(ctx, inlineObject58)

drive/folders

ドライブのフォルダ一覧を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject58** | [**InlineObject58**](InlineObject58.md)|  | 

### Return type

[**[]DriveFolder**](DriveFolder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFoldersCreate

> DriveFoldersCreate(ctx, inlineObject59)

drive/folders/create

ドライブのフォルダを作成します。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject59** | [**InlineObject59**](InlineObject59.md)|  | 

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


## DriveFoldersDelete

> DriveFoldersDelete(ctx, inlineObject60)

drive/folders/delete

指定したドライブのフォルダを削除します。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject60** | [**InlineObject60**](InlineObject60.md)|  | 

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


## DriveFoldersFind

> []DriveFolder DriveFoldersFind(ctx, inlineObject61)

drive/folders/find

No description provided.  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject61** | [**InlineObject61**](InlineObject61.md)|  | 

### Return type

[**[]DriveFolder**](DriveFolder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFoldersShow

> DriveFolder DriveFoldersShow(ctx, inlineObject62)

drive/folders/show

指定したドライブのフォルダの情報を取得します。  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject62** | [**InlineObject62**](InlineObject62.md)|  | 

### Return type

[**DriveFolder**](DriveFolder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveFoldersUpdate

> DriveFoldersUpdate(ctx, inlineObject63)

drive/folders/update

指定したドライブのフォルダの情報を更新します。  **Credential required**: *Yes* / **Permission**: *drive-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject63** | [**InlineObject63**](InlineObject63.md)|  | 

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


## DriveStream

> []DriveFile DriveStream(ctx, inlineObject64)

drive/stream

No description provided.  **Credential required**: *Yes* / **Permission**: *drive-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject64** | [**InlineObject64**](InlineObject64.md)|  | 

### Return type

[**[]DriveFile**](DriveFile.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

