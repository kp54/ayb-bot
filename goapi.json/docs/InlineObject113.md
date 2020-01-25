# InlineObject113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | タグ | [optional] 
**Query** | [**[][]string**](array.md) | クエリ | [optional] 
**Following** | Pointer to **bool** |  | [optional] 
**Mute** | **string** |  | [optional] [default to mute_all]
**Reply** | Pointer to **bool** | 返信に限定するか否か | [optional] 
**Renote** | Pointer to **bool** | Renoteに限定するか否か | [optional] 
**WithFiles** | **bool** | true にすると、ファイルが添付された投稿だけ取得します | [optional] 
**Media** | Pointer to **bool** | ファイルが添付された投稿に限定するか否か (このパラメータは廃止予定です。代わりに withFiles を使ってください。) | [optional] 
**Poll** | Pointer to **bool** | アンケートが添付された投稿に限定するか否か | [optional] 
**UntilId** | **string** | 指定すると、この投稿を基点としてより古い投稿を取得します | [optional] 
**SinceDate** | **float32** |  | [optional] 
**UntilDate** | **float32** |  | [optional] 
**Offset** | **float32** |  | [optional] 
**Limit** | **float32** |  | [optional] [default to 10]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


