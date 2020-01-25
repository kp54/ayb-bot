# \NotesApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartsNotes**](NotesApi.md#ChartsNotes) | **Post** /charts/notes | charts/notes
[**ChartsUserNotes**](NotesApi.md#ChartsUserNotes) | **Post** /charts/user/notes | charts/user/notes
[**DriveFilesAttachedNotes**](NotesApi.md#DriveFilesAttachedNotes) | **Post** /drive/files/attached_notes | drive/files/attached_notes
[**DriveFilesAttachedNotes_0**](NotesApi.md#DriveFilesAttachedNotes_0) | **Post** /drive/files/attached-notes | drive/files/attached-notes
[**IFavorites**](NotesApi.md#IFavorites) | **Post** /i/favorites | i/favorites
[**IPin**](NotesApi.md#IPin) | **Post** /i/pin | i/pin
[**IUnpin**](NotesApi.md#IUnpin) | **Post** /i/unpin | i/unpin
[**Notes**](NotesApi.md#Notes) | **Post** /notes | notes
[**NotesChildren**](NotesApi.md#NotesChildren) | **Post** /notes/children | notes/children
[**NotesConversation**](NotesApi.md#NotesConversation) | **Post** /notes/conversation | notes/conversation
[**NotesCreate**](NotesApi.md#NotesCreate) | **Post** /notes/create | notes/create
[**NotesDelete**](NotesApi.md#NotesDelete) | **Post** /notes/delete | notes/delete
[**NotesFeatured**](NotesApi.md#NotesFeatured) | **Post** /notes/featured | notes/featured
[**NotesGlobalTimeline**](NotesApi.md#NotesGlobalTimeline) | **Post** /notes/global-timeline | notes/global-timeline
[**NotesHybridTimeline**](NotesApi.md#NotesHybridTimeline) | **Post** /notes/hybrid-timeline | notes/hybrid-timeline
[**NotesLocalTimeline**](NotesApi.md#NotesLocalTimeline) | **Post** /notes/local-timeline | notes/local-timeline
[**NotesMentions**](NotesApi.md#NotesMentions) | **Post** /notes/mentions | notes/mentions
[**NotesPollsRecommendation**](NotesApi.md#NotesPollsRecommendation) | **Post** /notes/polls/recommendation | notes/polls/recommendation
[**NotesPollsVote**](NotesApi.md#NotesPollsVote) | **Post** /notes/polls/vote | notes/polls/vote
[**NotesReactions**](NotesApi.md#NotesReactions) | **Post** /notes/reactions | notes/reactions
[**NotesReactionsCreate**](NotesApi.md#NotesReactionsCreate) | **Post** /notes/reactions/create | notes/reactions/create
[**NotesReactionsDelete**](NotesApi.md#NotesReactionsDelete) | **Post** /notes/reactions/delete | notes/reactions/delete
[**NotesRenotes**](NotesApi.md#NotesRenotes) | **Post** /notes/renotes | notes/renotes
[**NotesReplies**](NotesApi.md#NotesReplies) | **Post** /notes/replies | notes/replies
[**NotesSearch**](NotesApi.md#NotesSearch) | **Post** /notes/search | notes/search
[**NotesSearchByTag**](NotesApi.md#NotesSearchByTag) | **Post** /notes/search_by_tag | notes/search_by_tag
[**NotesSearchByTag_0**](NotesApi.md#NotesSearchByTag_0) | **Post** /notes/search-by-tag | notes/search-by-tag
[**NotesShow**](NotesApi.md#NotesShow) | **Post** /notes/show | notes/show
[**NotesState**](NotesApi.md#NotesState) | **Post** /notes/state | notes/state
[**NotesTimeline**](NotesApi.md#NotesTimeline) | **Post** /notes/timeline | notes/timeline
[**NotesUserListTimeline**](NotesApi.md#NotesUserListTimeline) | **Post** /notes/user-list-timeline | notes/user-list-timeline
[**NotesWatchingCreate**](NotesApi.md#NotesWatchingCreate) | **Post** /notes/watching/create | notes/watching/create
[**NotesWatchingDelete**](NotesApi.md#NotesWatchingDelete) | **Post** /notes/watching/delete | notes/watching/delete
[**UsersNotes**](NotesApi.md#UsersNotes) | **Post** /users/notes | users/notes



## ChartsNotes

> InlineResponse2003 ChartsNotes(ctx, inlineObject40)

charts/notes

投稿のチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject40** | [**InlineObject40**](InlineObject40.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

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


## IPin

> IPin(ctx, inlineObject81)

i/pin

指定した投稿をピン留めします。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject81** | [**InlineObject81**](InlineObject81.md)|  | 

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


## IUnpin

> IUnpin(ctx, inlineObject82)

i/unpin

指定した投稿のピン留めを解除します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject82** | [**InlineObject82**](InlineObject82.md)|  | 

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


## Notes

> []Note Notes(ctx, inlineObject94)

notes

投稿を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject94** | [**InlineObject94**](InlineObject94.md)|  | 

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


## NotesChildren

> []Note NotesChildren(ctx, inlineObject95)

notes/children

指定した投稿への返信/引用を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject95** | [**InlineObject95**](InlineObject95.md)|  | 

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


## NotesConversation

> []Note NotesConversation(ctx, inlineObject96)

notes/conversation

指定した投稿の文脈を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject96** | [**InlineObject96**](InlineObject96.md)|  | 

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


## NotesCreate

> InlineResponse2008 NotesCreate(ctx, inlineObject97)

notes/create

投稿します。  **Credential required**: *Yes* / **Permission**: *note-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject97** | [**InlineObject97**](InlineObject97.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesDelete

> NotesDelete(ctx, inlineObject98)

notes/delete

指定した投稿を削除します。  **Credential required**: *Yes* / **Permission**: *note-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject98** | [**InlineObject98**](InlineObject98.md)|  | 

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


## NotesFeatured

> []Note NotesFeatured(ctx, inlineObject101)

notes/featured

Featuredな投稿を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject101** | [**InlineObject101**](InlineObject101.md)|  | 

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


## NotesGlobalTimeline

> []Note NotesGlobalTimeline(ctx, inlineObject102)

notes/global-timeline

グローバルタイムラインを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject102** | [**InlineObject102**](InlineObject102.md)|  | 

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


## NotesHybridTimeline

> []Note NotesHybridTimeline(ctx, inlineObject103)

notes/hybrid-timeline

ハイブリッドタイムラインを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject103** | [**InlineObject103**](InlineObject103.md)|  | 

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


## NotesLocalTimeline

> []Note NotesLocalTimeline(ctx, inlineObject104)

notes/local-timeline

ローカルタイムラインを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject104** | [**InlineObject104**](InlineObject104.md)|  | 

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


## NotesMentions

> []Note NotesMentions(ctx, inlineObject105)

notes/mentions

自分に言及している投稿の一覧を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject105** | [**InlineObject105**](InlineObject105.md)|  | 

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


## NotesPollsRecommendation

> NotesPollsRecommendation(ctx, inlineObject106)

notes/polls/recommendation

おすすめのアンケート一覧を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject106** | [**InlineObject106**](InlineObject106.md)|  | 

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


## NotesPollsVote

> NotesPollsVote(ctx, inlineObject107)

notes/polls/vote

指定した投稿のアンケートに投票します。  **Credential required**: *Yes* / **Permission**: *vote-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject107** | [**InlineObject107**](InlineObject107.md)|  | 

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


## NotesRenotes

> []Note NotesRenotes(ctx, inlineObject111)

notes/renotes

指定した投稿のRenote一覧を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject111** | [**InlineObject111**](InlineObject111.md)|  | 

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


## NotesReplies

> []Note NotesReplies(ctx, inlineObject112)

notes/replies

指定した投稿への返信を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject112** | [**InlineObject112**](InlineObject112.md)|  | 

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


## NotesSearch

> []Note NotesSearch(ctx, inlineObject115)

notes/search

投稿を検索します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject115** | [**InlineObject115**](InlineObject115.md)|  | 

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


## NotesSearchByTag

> []Note NotesSearchByTag(ctx, inlineObject113)

notes/search_by_tag

指定されたタグが付けられた投稿を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject113** | [**InlineObject113**](InlineObject113.md)|  | 

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


## NotesSearchByTag_0

> []Note NotesSearchByTag_0(ctx, inlineObject114)

notes/search-by-tag

指定されたタグが付けられた投稿を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject114** | [**InlineObject114**](InlineObject114.md)|  | 

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


## NotesShow

> Note NotesShow(ctx, inlineObject116)

notes/show

指定した投稿を取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject116** | [**InlineObject116**](InlineObject116.md)|  | 

### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesState

> NotesState(ctx, inlineObject117)

notes/state

指定した投稿の状態を取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject117** | [**InlineObject117**](InlineObject117.md)|  | 

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


## NotesTimeline

> []Note NotesTimeline(ctx, inlineObject118)

notes/timeline

タイムラインを取得します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject118** | [**InlineObject118**](InlineObject118.md)|  | 

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


## NotesWatchingCreate

> NotesWatchingCreate(ctx, inlineObject120)

notes/watching/create

指定した投稿をウォッチします。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject120** | [**InlineObject120**](InlineObject120.md)|  | 

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


## NotesWatchingDelete

> NotesWatchingDelete(ctx, inlineObject121)

notes/watching/delete

指定した投稿のウォッチを解除します。  **Credential required**: *Yes* / **Permission**: *account-write*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject121** | [**InlineObject121**](InlineObject121.md)|  | 

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

