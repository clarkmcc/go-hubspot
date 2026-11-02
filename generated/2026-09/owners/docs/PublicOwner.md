# PublicOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Indicates whether the owner is archived. | 
**CreatedAt** | **time.Time** | The date and time when the owner was created. | 
**Email** | Pointer to **string** | The email address of the owner. | [optional] 
**FirstName** | Pointer to **string** | The first name of the owner. | [optional] 
**Id** | **string** | The unique identifier of the owner. | 
**LastName** | Pointer to **string** | The last name of the owner. | [optional] 
**Teams** | Pointer to [**[]PublicTeam**](PublicTeam.md) | An array of teams associated with the owner, where each team is represented by a PublicTeam object. | [optional] 
**Type** | **string** | The type of the owner, which can be either PERSON or QUEUE. | 
**UpdatedAt** | **time.Time** | The date and time when the owner was last updated. | 
**UserId** | Pointer to **int32** | The user ID of the owner. | [optional] 
**UserIdIncludingInactive** | Pointer to **int32** | The user ID of the owner, including inactive users. | [optional] 

## Methods

### NewPublicOwner

`func NewPublicOwner(archived bool, createdAt time.Time, id string, type_ string, updatedAt time.Time, ) *PublicOwner`

NewPublicOwner instantiates a new PublicOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicOwnerWithDefaults

`func NewPublicOwnerWithDefaults() *PublicOwner`

NewPublicOwnerWithDefaults instantiates a new PublicOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *PublicOwner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PublicOwner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PublicOwner) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCreatedAt

`func (o *PublicOwner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicOwner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicOwner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *PublicOwner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PublicOwner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PublicOwner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PublicOwner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *PublicOwner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PublicOwner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PublicOwner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PublicOwner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *PublicOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicOwner) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *PublicOwner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PublicOwner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PublicOwner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PublicOwner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTeams

`func (o *PublicOwner) GetTeams() []PublicTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PublicOwner) GetTeamsOk() (*[]PublicTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PublicOwner) SetTeams(v []PublicTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PublicOwner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetType

`func (o *PublicOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicOwner) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *PublicOwner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicOwner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicOwner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *PublicOwner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicOwner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicOwner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PublicOwner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserIdIncludingInactive

`func (o *PublicOwner) GetUserIdIncludingInactive() int32`

GetUserIdIncludingInactive returns the UserIdIncludingInactive field if non-nil, zero value otherwise.

### GetUserIdIncludingInactiveOk

`func (o *PublicOwner) GetUserIdIncludingInactiveOk() (*int32, bool)`

GetUserIdIncludingInactiveOk returns a tuple with the UserIdIncludingInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdIncludingInactive

`func (o *PublicOwner) SetUserIdIncludingInactive(v int32)`

SetUserIdIncludingInactive sets UserIdIncludingInactive field to given value.

### HasUserIdIncludingInactive

`func (o *PublicOwner) HasUserIdIncludingInactive() bool`

HasUserIdIncludingInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


