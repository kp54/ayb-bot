# \AdminApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAbuseUserReports**](AdminApi.md#AdminAbuseUserReports) | **Post** /admin/abuse-user-reports | admin/abuse-user-reports
[**AdminDriveFiles**](AdminApi.md#AdminDriveFiles) | **Post** /admin/drive/files | admin/drive/files
[**AdminDriveShowFile**](AdminApi.md#AdminDriveShowFile) | **Post** /admin/drive/show-file | admin/drive/show-file
[**AdminEmojiAdd**](AdminApi.md#AdminEmojiAdd) | **Post** /admin/emoji/add | admin/emoji/add
[**AdminEmojiList**](AdminApi.md#AdminEmojiList) | **Post** /admin/emoji/list | admin/emoji/list
[**AdminEmojiRemove**](AdminApi.md#AdminEmojiRemove) | **Post** /admin/emoji/remove | admin/emoji/remove
[**AdminEmojiUpdate**](AdminApi.md#AdminEmojiUpdate) | **Post** /admin/emoji/update | admin/emoji/update
[**AdminFederationRemoveAllFollowing**](AdminApi.md#AdminFederationRemoveAllFollowing) | **Post** /admin/federation/remove-all-following | admin/federation/remove-all-following
[**AdminFederationUpdateInstance**](AdminApi.md#AdminFederationUpdateInstance) | **Post** /admin/federation/update-instance | admin/federation/update-instance
[**AdminInvite**](AdminApi.md#AdminInvite) | **Post** /admin/invite | admin/invite
[**AdminLogs**](AdminApi.md#AdminLogs) | **Post** /admin/logs | admin/logs
[**AdminModeratorsAdd**](AdminApi.md#AdminModeratorsAdd) | **Post** /admin/moderators/add | admin/moderators/add
[**AdminModeratorsRemove**](AdminApi.md#AdminModeratorsRemove) | **Post** /admin/moderators/remove | admin/moderators/remove
[**AdminQueueClear**](AdminApi.md#AdminQueueClear) | **Post** /admin/queue/clear | admin/queue/clear
[**AdminQueueJobs**](AdminApi.md#AdminQueueJobs) | **Post** /admin/queue/jobs | admin/queue/jobs
[**AdminQueueStats**](AdminApi.md#AdminQueueStats) | **Post** /admin/queue/stats | admin/queue/stats
[**AdminRemoveAbuseUserReport**](AdminApi.md#AdminRemoveAbuseUserReport) | **Post** /admin/remove-abuse-user-report | admin/remove-abuse-user-report
[**AdminResetPassword**](AdminApi.md#AdminResetPassword) | **Post** /admin/reset-password | admin/reset-password
[**AdminShowUser**](AdminApi.md#AdminShowUser) | **Post** /admin/show-user | admin/show-user
[**AdminShowUsers**](AdminApi.md#AdminShowUsers) | **Post** /admin/show-users | admin/show-users
[**AdminSilenceUser**](AdminApi.md#AdminSilenceUser) | **Post** /admin/silence-user | admin/silence-user
[**AdminSuspendUser**](AdminApi.md#AdminSuspendUser) | **Post** /admin/suspend-user | admin/suspend-user
[**AdminUnsilenceUser**](AdminApi.md#AdminUnsilenceUser) | **Post** /admin/unsilence-user | admin/unsilence-user
[**AdminUnsuspendUser**](AdminApi.md#AdminUnsuspendUser) | **Post** /admin/unsuspend-user | admin/unsuspend-user
[**AdminUnverifyUser**](AdminApi.md#AdminUnverifyUser) | **Post** /admin/unverify-user | admin/unverify-user
[**AdminUpdateMeta**](AdminApi.md#AdminUpdateMeta) | **Post** /admin/update-meta | admin/update-meta
[**AdminUpdateRemoteUser**](AdminApi.md#AdminUpdateRemoteUser) | **Post** /admin/update-remote-user | admin/update-remote-user
[**AdminVerifyUser**](AdminApi.md#AdminVerifyUser) | **Post** /admin/verify-user | admin/verify-user



## AdminAbuseUserReports

> AdminAbuseUserReports(ctx, inlineObject)

admin/abuse-user-reports

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

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


## AdminDriveFiles

> AdminDriveFiles(ctx, inlineObject1)

admin/drive/files

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

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


## AdminDriveShowFile

> AdminDriveShowFile(ctx, inlineObject2)

admin/drive/show-file

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## AdminEmojiAdd

> AdminEmojiAdd(ctx, inlineObject3)

admin/emoji/add

カスタム絵文字を追加します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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


## AdminEmojiList

> AdminEmojiList(ctx, inlineObject4)

admin/emoji/list

カスタム絵文字を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | 

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


## AdminEmojiRemove

> AdminEmojiRemove(ctx, inlineObject5)

admin/emoji/remove

カスタム絵文字を削除します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

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


## AdminEmojiUpdate

> AdminEmojiUpdate(ctx, inlineObject6)

admin/emoji/update

カスタム絵文字を更新します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | 

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


## AdminFederationRemoveAllFollowing

> AdminFederationRemoveAllFollowing(ctx, inlineObject7)

admin/federation/remove-all-following

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | 

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


## AdminFederationUpdateInstance

> AdminFederationUpdateInstance(ctx, inlineObject8)

admin/federation/update-instance

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject8** | [**InlineObject8**](InlineObject8.md)|  | 

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


## AdminInvite

> AdminInvite(ctx, body)

admin/invite

招待コードを発行します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

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


## AdminLogs

> AdminLogs(ctx, inlineObject9)

admin/logs

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject9** | [**InlineObject9**](InlineObject9.md)|  | 

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


## AdminModeratorsAdd

> AdminModeratorsAdd(ctx, inlineObject10)

admin/moderators/add

指定したユーザーをモデレーターにします。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject10** | [**InlineObject10**](InlineObject10.md)|  | 

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


## AdminModeratorsRemove

> AdminModeratorsRemove(ctx, inlineObject11)

admin/moderators/remove

指定したユーザーをモデレーター解除します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject11** | [**InlineObject11**](InlineObject11.md)|  | 

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


## AdminQueueClear

> AdminQueueClear(ctx, body)

admin/queue/clear

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

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


## AdminQueueJobs

> AdminQueueJobs(ctx, inlineObject12)

admin/queue/jobs

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject12** | [**InlineObject12**](InlineObject12.md)|  | 

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


## AdminQueueStats

> AdminQueueStats(ctx, body)

admin/queue/stats

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

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


## AdminRemoveAbuseUserReport

> AdminRemoveAbuseUserReport(ctx, inlineObject13)

admin/remove-abuse-user-report

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject13** | [**InlineObject13**](InlineObject13.md)|  | 

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


## AdminResetPassword

> AdminResetPassword(ctx, inlineObject14)

admin/reset-password

指定したユーザーのパスワードをリセットします。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject14** | [**InlineObject14**](InlineObject14.md)|  | 

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


## AdminShowUser

> AdminShowUser(ctx, inlineObject15)

admin/show-user

指定したユーザーの情報を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject15** | [**InlineObject15**](InlineObject15.md)|  | 

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


## AdminShowUsers

> AdminShowUsers(ctx, inlineObject16)

admin/show-users

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject16** | [**InlineObject16**](InlineObject16.md)|  | 

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


## AdminSilenceUser

> AdminSilenceUser(ctx, inlineObject17)

admin/silence-user

指定したユーザーをサイレンスにします。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject17** | [**InlineObject17**](InlineObject17.md)|  | 

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


## AdminSuspendUser

> AdminSuspendUser(ctx, inlineObject18)

admin/suspend-user

指定したユーザーを凍結します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject18** | [**InlineObject18**](InlineObject18.md)|  | 

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


## AdminUnsilenceUser

> AdminUnsilenceUser(ctx, inlineObject19)

admin/unsilence-user

指定したユーザーのサイレンスを解除します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject19** | [**InlineObject19**](InlineObject19.md)|  | 

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


## AdminUnsuspendUser

> AdminUnsuspendUser(ctx, inlineObject20)

admin/unsuspend-user

指定したユーザーの凍結を解除します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject20** | [**InlineObject20**](InlineObject20.md)|  | 

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


## AdminUnverifyUser

> AdminUnverifyUser(ctx, inlineObject21)

admin/unverify-user

指定したユーザーの公式アカウントを解除します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject21** | [**InlineObject21**](InlineObject21.md)|  | 

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


## AdminUpdateMeta

> AdminUpdateMeta(ctx, inlineObject22)

admin/update-meta

インスタンスの設定を更新します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject22** | [**InlineObject22**](InlineObject22.md)|  | 

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


## AdminUpdateRemoteUser

> AdminUpdateRemoteUser(ctx, inlineObject23)

admin/update-remote-user

指定されたリモートユーザーの情報を更新します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject23** | [**InlineObject23**](InlineObject23.md)|  | 

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


## AdminVerifyUser

> AdminVerifyUser(ctx, inlineObject24)

admin/verify-user

指定したユーザーを公式アカウントにします。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject24** | [**InlineObject24**](InlineObject24.md)|  | 

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

