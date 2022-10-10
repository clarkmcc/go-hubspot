# HubDbTableCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | Pointer to **string** | The new name for the cloned table | [optional] 
**NewLabel** | Pointer to **string** | The new label for the cloned table | [optional] 
**CopyRows** | **bool** | Specifies whether to copy the rows during clone | 

## Methods

### NewHubDbTableCloneRequest

`func NewHubDbTableCloneRequest(copyRows bool, ) *HubDbTableCloneRequest`

NewHubDbTableCloneRequest instantiates a new HubDbTableCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubDbTableCloneRequestWithDefaults

`func NewHubDbTableCloneRequestWithDefaults() *HubDbTableCloneRequest`

NewHubDbTableCloneRequestWithDefaults instantiates a new HubDbTableCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *HubDbTableCloneRequest) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *HubDbTableCloneRequest) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *HubDbTableCloneRequest) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *HubDbTableCloneRequest) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewLabel

`func (o *HubDbTableCloneRequest) GetNewLabel() string`

GetNewLabel returns the NewLabel field if non-nil, zero value otherwise.

### GetNewLabelOk

`func (o *HubDbTableCloneRequest) GetNewLabelOk() (*string, bool)`

GetNewLabelOk returns a tuple with the NewLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLabel

`func (o *HubDbTableCloneRequest) SetNewLabel(v string)`

SetNewLabel sets NewLabel field to given value.

### HasNewLabel

`func (o *HubDbTableCloneRequest) HasNewLabel() bool`

HasNewLabel returns a boolean if a field has been set.

### GetCopyRows

`func (o *HubDbTableCloneRequest) GetCopyRows() bool`

GetCopyRows returns the CopyRows field if non-nil, zero value otherwise.

### GetCopyRowsOk

`func (o *HubDbTableCloneRequest) GetCopyRowsOk() (*bool, bool)`

GetCopyRowsOk returns a tuple with the CopyRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyRows

`func (o *HubDbTableCloneRequest) SetCopyRows(v bool)`

SetCopyRows sets CopyRows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


