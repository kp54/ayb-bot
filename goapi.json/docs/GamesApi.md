# \GamesApi

All URIs are relative to *https://misskey.kurume-nct.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GamesReversiGames**](GamesApi.md#GamesReversiGames) | **Post** /games/reversi/games | games/reversi/games
[**GamesReversiGamesShow**](GamesApi.md#GamesReversiGamesShow) | **Post** /games/reversi/games/show | games/reversi/games/show
[**GamesReversiGamesSurrender**](GamesApi.md#GamesReversiGamesSurrender) | **Post** /games/reversi/games/surrender | games/reversi/games/surrender
[**GamesReversiInvitations**](GamesApi.md#GamesReversiInvitations) | **Post** /games/reversi/invitations | games/reversi/invitations
[**GamesReversiMatch**](GamesApi.md#GamesReversiMatch) | **Post** /games/reversi/match | games/reversi/match
[**GamesReversiMatchCancel**](GamesApi.md#GamesReversiMatchCancel) | **Post** /games/reversi/match/cancel | games/reversi/match/cancel



## GamesReversiGames

> GamesReversiGames(ctx, inlineObject72)

games/reversi/games

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject72** | [**InlineObject72**](InlineObject72.md)|  | 

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


## GamesReversiGamesShow

> GamesReversiGamesShow(ctx, inlineObject73)

games/reversi/games/show

No description provided.  **Credential required**: *No*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject73** | [**InlineObject73**](InlineObject73.md)|  | 

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


## GamesReversiGamesSurrender

> GamesReversiGamesSurrender(ctx, inlineObject74)

games/reversi/games/surrender

指定したリバーシの対局で投了します。  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject74** | [**InlineObject74**](InlineObject74.md)|  | 

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


## GamesReversiInvitations

> GamesReversiInvitations(ctx, body)

games/reversi/invitations

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


## GamesReversiMatch

> GamesReversiMatch(ctx, inlineObject75)

games/reversi/match

No description provided.  **Credential required**: *Yes*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject75** | [**InlineObject75**](InlineObject75.md)|  | 

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


## GamesReversiMatchCancel

> GamesReversiMatchCancel(ctx, body)

games/reversi/match/cancel

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

