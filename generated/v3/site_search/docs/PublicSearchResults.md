# PublicSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Offset** | **int32** |  | 
**Limit** | **int32** |  | 
**Results** | [**[]ContentSearchResult**](ContentSearchResult.md) |  | 
**SearchTerm** | Pointer to **string** |  | [optional] 
**Page** | **int32** |  | 

## Methods

### NewPublicSearchResults

`func NewPublicSearchResults(total int32, offset int32, limit int32, results []ContentSearchResult, page int32, ) *PublicSearchResults`

NewPublicSearchResults instantiates a new PublicSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSearchResultsWithDefaults

`func NewPublicSearchResultsWithDefaults() *PublicSearchResults`

NewPublicSearchResultsWithDefaults instantiates a new PublicSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PublicSearchResults) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PublicSearchResults) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PublicSearchResults) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *PublicSearchResults) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PublicSearchResults) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PublicSearchResults) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *PublicSearchResults) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PublicSearchResults) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PublicSearchResults) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *PublicSearchResults) GetResults() []ContentSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PublicSearchResults) GetResultsOk() (*[]ContentSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PublicSearchResults) SetResults(v []ContentSearchResult)`

SetResults sets Results field to given value.


### GetSearchTerm

`func (o *PublicSearchResults) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *PublicSearchResults) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *PublicSearchResults) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *PublicSearchResults) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### GetPage

`func (o *PublicSearchResults) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PublicSearchResults) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PublicSearchResults) SetPage(v int32)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


