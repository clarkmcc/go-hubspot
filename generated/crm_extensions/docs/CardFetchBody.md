# CardFetchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | URL to a service endpoints that will respond with card details. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed. | 
**ObjectTypes** | [**[]CardObjectTypeBody**](CardObjectTypeBody.md) | An array of CRM object types where this card should be displayed. HubSpot will call your data fetch URL whenever a user visits a record page of the types defined here. | 

## Methods

### NewCardFetchBody

`func NewCardFetchBody(targetUrl string, objectTypes []CardObjectTypeBody, ) *CardFetchBody`

NewCardFetchBody instantiates a new CardFetchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardFetchBodyWithDefaults

`func NewCardFetchBodyWithDefaults() *CardFetchBody`

NewCardFetchBodyWithDefaults instantiates a new CardFetchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *CardFetchBody) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CardFetchBody) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CardFetchBody) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetObjectTypes

`func (o *CardFetchBody) GetObjectTypes() []CardObjectTypeBody`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *CardFetchBody) GetObjectTypesOk() (*[]CardObjectTypeBody, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *CardFetchBody) SetObjectTypes(v []CardObjectTypeBody)`

SetObjectTypes sets ObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


