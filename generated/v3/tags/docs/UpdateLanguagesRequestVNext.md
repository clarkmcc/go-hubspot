# UpdateLanguagesRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryId** | **string** | ID of the primary object in the multi-language group. | 
**Languages** | **map[string]string** | Map of object IDs to associated languages of object in the multi-language group. | 

## Methods

### NewUpdateLanguagesRequestVNext

`func NewUpdateLanguagesRequestVNext(primaryId string, languages map[string]string, ) *UpdateLanguagesRequestVNext`

NewUpdateLanguagesRequestVNext instantiates a new UpdateLanguagesRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLanguagesRequestVNextWithDefaults

`func NewUpdateLanguagesRequestVNextWithDefaults() *UpdateLanguagesRequestVNext`

NewUpdateLanguagesRequestVNextWithDefaults instantiates a new UpdateLanguagesRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryId

`func (o *UpdateLanguagesRequestVNext) GetPrimaryId() string`

GetPrimaryId returns the PrimaryId field if non-nil, zero value otherwise.

### GetPrimaryIdOk

`func (o *UpdateLanguagesRequestVNext) GetPrimaryIdOk() (*string, bool)`

GetPrimaryIdOk returns a tuple with the PrimaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryId

`func (o *UpdateLanguagesRequestVNext) SetPrimaryId(v string)`

SetPrimaryId sets PrimaryId field to given value.


### GetLanguages

`func (o *UpdateLanguagesRequestVNext) GetLanguages() map[string]string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *UpdateLanguagesRequestVNext) GetLanguagesOk() (*map[string]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *UpdateLanguagesRequestVNext) SetLanguages(v map[string]string)`

SetLanguages sets Languages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


