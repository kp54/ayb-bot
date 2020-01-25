# \UsersApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockingCreate**](UsersApi.md#BlockingCreate) | **Post** /blocking/create | blocking/create
[**BlockingDelete**](UsersApi.md#BlockingDelete) | **Post** /blocking/delete | blocking/delete
[**ChartsActiveUsers**](UsersApi.md#ChartsActiveUsers) | **Post** /charts/active-users | charts/active-users
[**ChartsUserDrive**](UsersApi.md#ChartsUserDrive) | **Post** /charts/user/drive | charts/user/drive
[**ChartsUserFollowing**](UsersApi.md#ChartsUserFollowing) | **Post** /charts/user/following | charts/user/following
[**ChartsUserNotes**](UsersApi.md#ChartsUserNotes) | **Post** /charts/user/notes | charts/user/notes
[**ChartsUserReactions**](UsersApi.md#ChartsUserReactions) | **Post** /charts/user/reactions | charts/user/reactions
[**ChartsUsers**](UsersApi.md#ChartsUsers) | **Post** /charts/users | charts/users
[**FollowingCreate**](UsersApi.md#FollowingCreate) | **Post** /following/create | following/create
[**FollowingDelete**](UsersApi.md#FollowingDelete) | **Post** /following/delete | following/delete
[**HashtagsUsers**](UsersApi.md#HashtagsUsers) | **Post** /hashtags/users | hashtags/users
[**MuteCreate**](UsersApi.md#MuteCreate) | **Post** /mute/create | mute/create
[**MuteDelete**](UsersApi.md#MuteDelete) | **Post** /mute/delete | mute/delete
[**UsernameAvailable**](UsersApi.md#UsernameAvailable) | **Post** /username/available | username/available
[**Users**](UsersApi.md#Users) | **Post** /users | users
[**UsersFollowers**](UsersApi.md#UsersFollowers) | **Post** /users/followers | users/followers
[**UsersFollowing**](UsersApi.md#UsersFollowing) | **Post** /users/following | users/following
[**UsersGetFrequentlyRepliedUsers**](UsersApi.md#UsersGetFrequentlyRepliedUsers) | **Post** /users/get_frequently_replied_users | users/get_frequently_replied_users
[**UsersGetFrequentlyRepliedUsers_0**](UsersApi.md#UsersGetFrequentlyRepliedUsers_0) | **Post** /users/get-frequently-replied-users | users/get-frequently-replied-users
[**UsersListsPull**](UsersApi.md#UsersListsPull) | **Post** /users/lists/pull | users/lists/pull
[**UsersListsPush**](UsersApi.md#UsersListsPush) | **Post** /users/lists/push | users/lists/push
[**UsersNotes**](UsersApi.md#UsersNotes) | **Post** /users/notes | users/notes
[**UsersRecommendation**](UsersApi.md#UsersRecommendation) | **Post** /users/recommendation | users/recommendation
[**UsersRelation**](UsersApi.md#UsersRelation) | **Post** /users/relation | users/relation
[**UsersReportAbuse**](UsersApi.md#UsersReportAbuse) | **Post** /users/report-abuse | users/report-abuse
[**UsersSearch**](UsersApi.md#UsersSearch) | **Post** /users/search | users/search
[**UsersShow**](UsersApi.md#UsersShow) | **Post** /users/show | users/show



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


## ChartsActiveUsers

> []map[string]interface{} ChartsActiveUsers(ctx, inlineObject34)

charts/active-users

アクティブユーザーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject34** | [**InlineObject34**](InlineObject34.md)|  | 

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


## ChartsUserFollowing

> InlineResponse2004 ChartsUserFollowing(ctx, inlineObject42)

charts/user/following

ユーザーごとのフォロー/フォロワーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject42** | [**InlineObject42**](InlineObject42.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChartsUserNotes

> InlineResponse2003Local ChartsUserNotes(ctx, inlineObject43)

charts/user/notes

ユーザーごとの投稿のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject43** | [**InlineObject43**](InlineObject43.md)|  | 

### Return type

[**InlineResponse2003Local**](inline_response_200_3_local.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ChartsUsers

> InlineResponse2005 ChartsUsers(ctx, inlineObject45)

charts/users

ユーザーのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject45** | [**InlineObject45**](InlineObject45.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowingCreate

> FollowingCreate(ctx, inlineObject67)

following/create

指定したユーザーをフォローします。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject67** | [**InlineObject67**](InlineObject67.md)|  | 

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


## FollowingDelete

> FollowingDelete(ctx, inlineObject68)

following/delete

指定したユーザーのフォローを解除します。  **Credential required**: *Yes* / **Permission**: *following-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject68** | [**InlineObject68**](InlineObject68.md)|  | 

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


## HashtagsUsers

> []User HashtagsUsers(ctx, inlineObject78)

hashtags/users

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject78** | [**InlineObject78**](InlineObject78.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## UsernameAvailable

> UsernameAvailable(ctx, inlineObject123)

username/available

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject123** | [**InlineObject123**](InlineObject123.md)|  | 

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


## Users

> []User Users(ctx, inlineObject124)

users

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject124** | [**InlineObject124**](InlineObject124.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersFollowers

> InlineResponse20010 UsersFollowers(ctx, inlineObject125)

users/followers

指定したユーザーのフォロワー一覧を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject125** | [**InlineObject125**](InlineObject125.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersFollowing

> InlineResponse20010 UsersFollowing(ctx, inlineObject126)

users/following

指定したユーザーのフォロー一覧を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject126** | [**InlineObject126**](InlineObject126.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetFrequentlyRepliedUsers

> UsersGetFrequentlyRepliedUsers(ctx, inlineObject127)

users/get_frequently_replied_users

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject127** | [**InlineObject127**](InlineObject127.md)|  | 

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


## UsersGetFrequentlyRepliedUsers_0

> []User UsersGetFrequentlyRepliedUsers_0(ctx, inlineObject128)

users/get-frequently-replied-users

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject128** | [**InlineObject128**](InlineObject128.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

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


## UsersNotes

> []Note UsersNotes(ctx, inlineObject135)

users/notes

指定したユーザーのタイムラインを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject135** | [**InlineObject135**](InlineObject135.md)|  | 

### Return type

[**[]Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRecommendation

> []User UsersRecommendation(ctx, inlineObject136)

users/recommendation

おすすめのユーザー一覧を取得します。  **Credential required**: *Yes* / **Permission**: *account-read*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject136** | [**InlineObject136**](InlineObject136.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRelation

> UsersRelation(ctx, inlineObject137)

users/relation

ユーザー間のリレーションを取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject137** | [**InlineObject137**](InlineObject137.md)|  | 

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


## UsersReportAbuse

> UsersReportAbuse(ctx, inlineObject138)

users/report-abuse

指定したユーザーを迷惑なユーザーであると報告します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject138** | [**InlineObject138**](InlineObject138.md)|  | 

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


## UsersSearch

> []User UsersSearch(ctx, inlineObject139)

users/search

ユーザーを検索します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject139** | [**InlineObject139**](InlineObject139.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersShow

> User UsersShow(ctx, inlineObject140)

users/show

指定したユーザーの情報を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject140** | [**InlineObject140**](InlineObject140.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

