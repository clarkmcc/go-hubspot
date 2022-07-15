# CollectionResponseSmtpApiTokenViewForwardPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SmtpApiTokenView**](SmtpApiTokenView.md) |  | 
**Paging** | Pointer to [**ForwardPaging**](ForwardPaging.md) |  | [optional] 

## Methods

### NewCollectionResponseSmtpApiTokenViewForwardPaging

`func NewCollectionResponseSmtpApiTokenViewForwardPaging(results []SmtpApiTokenView, ) *CollectionResponseSmtpApiTokenViewForwardPaging`

NewCollectionResponseSmtpApiTokenViewForwardPaging instantiates a new CollectionResponseSmtpApiTokenViewForwardPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseSmtpApiTokenViewForwardPagingWithDefaults

`func NewCollectionResponseSmtpApiTokenViewForwardPagingWithDefaults() *CollectionResponseSmtpApiTokenViewForwardPaging`

NewCollectionResponseSmtpApiTokenViewForwardPagingWithDefaults instantiates a new CollectionResponseSmtpApiTokenViewForwardPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) GetResults() []SmtpApiTokenView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) GetResultsOk() (*[]SmtpApiTokenView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) SetResults(v []SmtpApiTokenView)`

SetResults sets Results field to given value.


### GetPaging

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) GetPaging() ForwardPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) GetPagingOk() (*ForwardPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) SetPaging(v ForwardPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponseSmtpApiTokenViewForwardPaging) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


