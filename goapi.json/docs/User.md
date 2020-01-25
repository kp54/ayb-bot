# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for this User. | 
**Username** | **string** | The screen name, handle, or alias that this user identifies themselves with. | 
**Name** | Pointer to **string** | The name of the user, as they’ve defined it. | 
**Host** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | The user-defined UTF-8 string describing their account. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date that the user account was created on Misskey. | 
**FollowersCount** | **float32** | The number of followers this account currently has. | [optional] 
**FollowingCount** | **float32** | The number of users this account is following. | [optional] 
**NotesCount** | **float32** | The number of Notes (including renotes) issued by the user. | [optional] 
**IsBot** | **bool** | Whether this account is a bot. | [optional] 
**IsCat** | **bool** | Whether this account is a cat. | [optional] 
**IsAdmin** | **bool** | Whether this account is the admin. | [optional] 
**IsVerified** | **bool** |  | [optional] 
**IsLocked** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


