# PublicObjectSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**Limit** | **int32** |  | 
**After** | **string** |  | 
**Sorts** | **[]string** |  | 
**Properties** | **[]string** |  | 
**FilterGroups** | [**[]FilterGroup**](FilterGroup.md) |  | 

## Methods

### NewPublicObjectSearchRequest

`func NewPublicObjectSearchRequest(limit int32, after string, sorts []string, properties []string, filterGroups []FilterGroup, ) *PublicObjectSearchRequest`

NewPublicObjectSearchRequest instantiates a new PublicObjectSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicObjectSearchRequestWithDefaults

`func NewPublicObjectSearchRequestWithDefaults() *PublicObjectSearchRequest`

NewPublicObjectSearchRequestWithDefaults instantiates a new PublicObjectSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *PublicObjectSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PublicObjectSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PublicObjectSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PublicObjectSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetLimit

`func (o *PublicObjectSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PublicObjectSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PublicObjectSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetAfter

`func (o *PublicObjectSearchRequest) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *PublicObjectSearchRequest) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *PublicObjectSearchRequest) SetAfter(v string)`

SetAfter sets After field to given value.


### GetSorts

`func (o *PublicObjectSearchRequest) GetSorts() []string`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *PublicObjectSearchRequest) GetSortsOk() (*[]string, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *PublicObjectSearchRequest) SetSorts(v []string)`

SetSorts sets Sorts field to given value.


### GetProperties

`func (o *PublicObjectSearchRequest) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PublicObjectSearchRequest) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PublicObjectSearchRequest) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetFilterGroups

`func (o *PublicObjectSearchRequest) GetFilterGroups() []FilterGroup`

GetFilterGroups returns the FilterGroups field if non-nil, zero value otherwise.

### GetFilterGroupsOk

`func (o *PublicObjectSearchRequest) GetFilterGroupsOk() (*[]FilterGroup, bool)`

GetFilterGroupsOk returns a tuple with the FilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroups

`func (o *PublicObjectSearchRequest) SetFilterGroups(v []FilterGroup)`

SetFilterGroups sets FilterGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


