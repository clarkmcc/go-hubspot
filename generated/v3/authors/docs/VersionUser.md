# VersionUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the User. | 
**Email** | **string** | The email address of the user. | 
**FullName** | **string** | The first and last name of the User. | 

## Methods

### NewVersionUser

`func NewVersionUser(id string, email string, fullName string, ) *VersionUser`

NewVersionUser instantiates a new VersionUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionUserWithDefaults

`func NewVersionUserWithDefaults() *VersionUser`

NewVersionUserWithDefaults instantiates a new VersionUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionUser) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *VersionUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VersionUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VersionUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *VersionUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VersionUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VersionUser) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


