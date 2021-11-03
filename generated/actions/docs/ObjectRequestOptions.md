# ObjectRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **[]string** | A list of properties of the CRM object to include with the request to the &#x60;actionUrl&#x60;. | 

## Methods

### NewObjectRequestOptions

`func NewObjectRequestOptions(properties []string, ) *ObjectRequestOptions`

NewObjectRequestOptions instantiates a new ObjectRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRequestOptionsWithDefaults

`func NewObjectRequestOptionsWithDefaults() *ObjectRequestOptions`

NewObjectRequestOptionsWithDefaults instantiates a new ObjectRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ObjectRequestOptions) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectRequestOptions) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectRequestOptions) SetProperties(v []string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


