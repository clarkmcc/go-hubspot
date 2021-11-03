# SyncContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account in the external accounting system. This is the value that will be passed as &#x60;accountId&#x60; for all outbound calls for the user from HubSpot to the external accounting system. | 
**Contacts** | [**[]UpdatedContact**](UpdatedContact.md) | A list of contacts to be imported. | 

## Methods

### NewSyncContactsRequest

`func NewSyncContactsRequest(accountId string, contacts []UpdatedContact, ) *SyncContactsRequest`

NewSyncContactsRequest instantiates a new SyncContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncContactsRequestWithDefaults

`func NewSyncContactsRequestWithDefaults() *SyncContactsRequest`

NewSyncContactsRequestWithDefaults instantiates a new SyncContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SyncContactsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SyncContactsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SyncContactsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetContacts

`func (o *SyncContactsRequest) GetContacts() []UpdatedContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *SyncContactsRequest) GetContactsOk() (*[]UpdatedContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *SyncContactsRequest) SetContacts(v []UpdatedContact)`

SetContacts sets Contacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


