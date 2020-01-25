# DriveFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this Drive file. | 
**CreatedAt** | [**time.Time**](time.Time.md) | The date that the Drive file was created on Misskey. | 
**Name** | **string** | The file name with extension. | 
**Type** | **string** | The MIME type of this Drive file. | 
**Md5** | **string** | The MD5 hash of this Drive file. | 
**Datasize** | **float32** | The size of this Drive file. (bytes) | 
**FolderId** | Pointer to **string** | The parent folder ID of this Drive file. | [optional] 
**IsSensitive** | **bool** | Whether this Drive file is sensitive. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


