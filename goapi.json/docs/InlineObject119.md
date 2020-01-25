# InlineObject119

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** | リストのID | 
**Limit** | **float32** | 最大数 | [optional] [default to 10]
**SinceId** | **string** | 指定すると、この投稿を基点としてより新しい投稿を取得します | [optional] 
**UntilId** | **string** | 指定すると、この投稿を基点としてより古い投稿を取得します | [optional] 
**SinceDate** | **float32** | 指定した時間を基点としてより新しい投稿を取得します。数値は、1970年1月1日 00:00:00 UTC から指定した日時までの経過時間をミリ秒単位で表します。 | [optional] 
**UntilDate** | **float32** | 指定した時間を基点としてより古い投稿を取得します。数値は、1970年1月1日 00:00:00 UTC から指定した日時までの経過時間をミリ秒単位で表します。 | [optional] 
**IncludeMyRenotes** | **bool** | 自分の行ったRenoteを含めるかどうか | [optional] [default to true]
**IncludeRenotedMyNotes** | **bool** | Renoteされた自分の投稿を含めるかどうか | [optional] [default to true]
**IncludeLocalRenotes** | **bool** | Renoteされたローカルの投稿を含めるかどうか | [optional] [default to true]
**WithFiles** | **bool** | true にすると、ファイルが添付された投稿だけ取得します | [optional] 
**MediaOnly** | **bool** | true にすると、ファイルが添付された投稿だけ取得します (このパラメータは廃止予定です。代わりに withFiles を使ってください。) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


