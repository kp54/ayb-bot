# \HashtagsApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregationHashtags**](HashtagsApi.md#AggregationHashtags) | **Post** /aggregation/hashtags | aggregation/hashtags
[**ChartsHashtag**](HashtagsApi.md#ChartsHashtag) | **Post** /charts/hashtag | charts/hashtag
[**HashtagsList**](HashtagsApi.md#HashtagsList) | **Post** /hashtags/list | hashtags/list
[**HashtagsSearch**](HashtagsApi.md#HashtagsSearch) | **Post** /hashtags/search | hashtags/search
[**HashtagsTrend**](HashtagsApi.md#HashtagsTrend) | **Post** /hashtags/trend | hashtags/trend
[**HashtagsUsers**](HashtagsApi.md#HashtagsUsers) | **Post** /hashtags/users | hashtags/users
[**NotesSearchByTag**](HashtagsApi.md#NotesSearchByTag) | **Post** /notes/search_by_tag | notes/search_by_tag
[**NotesSearchByTag_0**](HashtagsApi.md#NotesSearchByTag_0) | **Post** /notes/search-by-tag | notes/search-by-tag



## AggregationHashtags

> AggregationHashtags(ctx, body)

aggregation/hashtags

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

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


## ChartsHashtag

> []map[string]interface{} ChartsHashtag(ctx, inlineObject37)

charts/hashtag

ハッシュタグごとのチャートを取得します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject37** | [**InlineObject37**](InlineObject37.md)|  | 

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


## HashtagsList

> []Hashtag HashtagsList(ctx, inlineObject76)

hashtags/list

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject76** | [**InlineObject76**](InlineObject76.md)|  | 

### Return type

[**[]Hashtag**](Hashtag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HashtagsSearch

> []string HashtagsSearch(ctx, inlineObject77)

hashtags/search

ハッシュタグを検索します。  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject77** | [**InlineObject77**](InlineObject77.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HashtagsTrend

> HashtagsTrend(ctx, body)

hashtags/trend

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

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

