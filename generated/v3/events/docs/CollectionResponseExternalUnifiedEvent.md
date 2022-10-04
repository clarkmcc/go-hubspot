# CollectionResponseExternalUnifiedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ExternalUnifiedEvent**](ExternalUnifiedEvent.md) |  | 
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 

## Methods

### NewCollectionResponseExternalUnifiedEvent

`func NewCollectionResponseExternalUnifiedEvent(results []ExternalUnifiedEvent, ) *CollectionResponseExternalUnifiedEvent`

NewCollectionResponseExternalUnifiedEvent instantiates a new CollectionResponseExternalUnifiedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseExternalUnifiedEventWithDefaults

`func NewCollectionResponseExternalUnifiedEventWithDefaults() *CollectionResponseExternalUnifiedEvent`

NewCollectionResponseExternalUnifiedEventWithDefaults instantiates a new CollectionResponseExternalUnifiedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CollectionResponseExternalUnifiedEvent) GetResults() []ExternalUnifiedEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseExternalUnifiedEvent) GetResultsOk() (*[]ExternalUnifiedEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseExternalUnifiedEvent) SetResults(v []ExternalUnifiedEvent)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseExternalUnifiedEvent) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseExternalUnifiedEvent) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseExternalUnifiedEvent) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseExternalUnifiedEvent) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


