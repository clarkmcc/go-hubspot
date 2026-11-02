# Paging1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**NextPage**](NextPage.md) |  | [optional] 
**Prev** | Pointer to [**PreviousPage**](PreviousPage.md) |  | [optional] 

## Methods

### NewPaging1

`func NewPaging1() *Paging1`

NewPaging1 instantiates a new Paging1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaging1WithDefaults

`func NewPaging1WithDefaults() *Paging1`

NewPaging1WithDefaults instantiates a new Paging1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *Paging1) GetNext() NextPage`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Paging1) GetNextOk() (*NextPage, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Paging1) SetNext(v NextPage)`

SetNext sets Next field to given value.

### HasNext

`func (o *Paging1) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *Paging1) GetPrev() PreviousPage`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *Paging1) GetPrevOk() (*PreviousPage, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *Paging1) SetPrev(v PreviousPage)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *Paging1) HasPrev() bool`

HasPrev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


