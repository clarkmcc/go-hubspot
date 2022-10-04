# CardObjectTypeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A CRM object type where this card should be displayed. | 
**PropertiesToSend** | **[]string** | An array of properties that should be sent to this card&#39;s target URL when the data fetch request is made. Must be valid properties for the corresponding CRM object type. | 

## Methods

### NewCardObjectTypeBody

`func NewCardObjectTypeBody(name string, propertiesToSend []string, ) *CardObjectTypeBody`

NewCardObjectTypeBody instantiates a new CardObjectTypeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardObjectTypeBodyWithDefaults

`func NewCardObjectTypeBodyWithDefaults() *CardObjectTypeBody`

NewCardObjectTypeBodyWithDefaults instantiates a new CardObjectTypeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CardObjectTypeBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardObjectTypeBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardObjectTypeBody) SetName(v string)`

SetName sets Name field to given value.


### GetPropertiesToSend

`func (o *CardObjectTypeBody) GetPropertiesToSend() []string`

GetPropertiesToSend returns the PropertiesToSend field if non-nil, zero value otherwise.

### GetPropertiesToSendOk

`func (o *CardObjectTypeBody) GetPropertiesToSendOk() (*[]string, bool)`

GetPropertiesToSendOk returns a tuple with the PropertiesToSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesToSend

`func (o *CardObjectTypeBody) SetPropertiesToSend(v []string)`

SetPropertiesToSend sets PropertiesToSend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


