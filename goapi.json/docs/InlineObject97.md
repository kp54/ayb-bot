# InlineObject97

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | **string** | 投稿の公開範囲 | [optional] [default to VISIBILITY_PUBLIC]
**VisibleUserIds** | **[]string** | (投稿の公開範囲が specified の場合)投稿を閲覧できるユーザー | [optional] 
**Text** | Pointer to **string** | 投稿内容 | [optional] 
**Cw** | Pointer to **string** | コンテンツの警告。このパラメータを指定すると設定したテキストで投稿のコンテンツを隠す事が出来ます。 | [optional] 
**ViaMobile** | **bool** | モバイルデバイスからの投稿か否か。 | [optional] 
**LocalOnly** | **bool** | ローカルのみに投稿か否か。 | [optional] 
**NoExtractMentions** | **bool** | 本文からメンションを展開しないか否か。 | [optional] 
**NoExtractHashtags** | **bool** | 本文からハッシュタグを展開しないか否か。 | [optional] 
**NoExtractEmojis** | **bool** | 本文からカスタム絵文字を展開しないか否か。 | [optional] 
**Geo** | Pointer to [**NotesCreateGeo**](_notes_create_geo.md) |  | [optional] 
**FileIds** | **[]string** | 添付するファイル | [optional] 
**MediaIds** | **[]string** | 添付するファイル (このパラメータは廃止予定です。代わりに fileIds を使ってください。) | [optional] 
**ReplyId** | **string** | 返信対象 | [optional] 
**RenoteId** | **string** | Renote対象 | [optional] 
**Poll** | [**NotesCreatePoll**](_notes_create_poll.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


