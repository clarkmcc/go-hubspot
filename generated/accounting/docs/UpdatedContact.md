# UpdatedContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncAction** | **string** | The operation to be performed. | 
**UpdatedAt** | **time.Time** | The timestamp (ISO8601 format) when the customer was updated in the external accounting system. | 
**EmailAddress** | **string** | The customer&#39;s email address | 
**Id** | **string** | The ID of the customer in the external accounting system. | 
**CustomerType** | Pointer to **string** | Designates the type of the customer object. | [optional] 

## Methods

### NewUpdatedContact

`func NewUpdatedContact(syncAction string, updatedAt time.Time, emailAddress string, id string, ) *UpdatedContact`

NewUpdatedContact instantiates a new UpdatedContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedContactWithDefaults

`func NewUpdatedContactWithDefaults() *UpdatedContact`

NewUpdatedContactWithDefaults instantiates a new UpdatedContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncAction

`func (o *UpdatedContact) GetSyncAction() string`

GetSyncAction returns the SyncAction field if non-nil, zero value otherwise.

### GetSyncActionOk

`func (o *UpdatedContact) GetSyncActionOk() (*string, bool)`

GetSyncActionOk returns a tuple with the SyncAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAction

`func (o *UpdatedContact) SetSyncAction(v string)`

SetSyncAction sets SyncAction field to given value.


### GetUpdatedAt

`func (o *UpdatedContact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatedContact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatedContact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmailAddress

`func (o *UpdatedContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UpdatedContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UpdatedContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetId

`func (o *UpdatedContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedContact) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerType

`func (o *UpdatedContact) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *UpdatedContact) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *UpdatedContact) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *UpdatedContact) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


