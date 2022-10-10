# CardFetchBodyPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | Pointer to **string** | URL to a service endpoint that will respond with details for this card. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed. | [optional] 
**ObjectTypes** | [**[]CardObjectTypeBody**](CardObjectTypeBody.md) | An array of CRM object types where this card should be displayed. HubSpot will call your target URL whenever a user visits a record page of the types defined here. | 

## Methods

### NewCardFetchBodyPatch

`func NewCardFetchBodyPatch(objectTypes []CardObjectTypeBody, ) *CardFetchBodyPatch`

NewCardFetchBodyPatch instantiates a new CardFetchBodyPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardFetchBodyPatchWithDefaults

`func NewCardFetchBodyPatchWithDefaults() *CardFetchBodyPatch`

NewCardFetchBodyPatchWithDefaults instantiates a new CardFetchBodyPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *CardFetchBodyPatch) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CardFetchBodyPatch) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CardFetchBodyPatch) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *CardFetchBodyPatch) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetObjectTypes

`func (o *CardFetchBodyPatch) GetObjectTypes() []CardObjectTypeBody`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *CardFetchBodyPatch) GetObjectTypesOk() (*[]CardObjectTypeBody, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *CardFetchBodyPatch) SetObjectTypes(v []CardObjectTypeBody)`

SetObjectTypes sets ObjectTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


