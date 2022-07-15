# PublicSingleSendRequestEgg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **int32** | The content ID for the transactional email, which can be found in email tool UI. | 
**Message** | [**PublicSingleSendEmail**](PublicSingleSendEmail.md) |  | 
**ContactProperties** | Pointer to **map[string]string** | The contactProperties field is a map of contact property values. Each contact property value contains a name and value property. Each property will get set on the contact record and will be visible in the template under {{ contact.NAME }}. Use these properties when you want to set a contact property while you’re sending the email. For example, when sending a reciept you may want to set a last_paid_date property, as the sending of the receipt will have information about the last payment. | [optional] 
**CustomProperties** | Pointer to **map[string]map[string]interface{}** | The customProperties field is a map of property values. Each property value contains a name and value property. Each property will be visible in the template under {{ custom.NAME }}. Note: Custom properties do not currently support arrays. To provide a listing in an email, one workaround is to build an HTML list (either with tables or ul) and specify it as a custom property. | [optional] 

## Methods

### NewPublicSingleSendRequestEgg

`func NewPublicSingleSendRequestEgg(emailId int32, message PublicSingleSendEmail, ) *PublicSingleSendRequestEgg`

NewPublicSingleSendRequestEgg instantiates a new PublicSingleSendRequestEgg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSingleSendRequestEggWithDefaults

`func NewPublicSingleSendRequestEggWithDefaults() *PublicSingleSendRequestEgg`

NewPublicSingleSendRequestEggWithDefaults instantiates a new PublicSingleSendRequestEgg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailId

`func (o *PublicSingleSendRequestEgg) GetEmailId() int32`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *PublicSingleSendRequestEgg) GetEmailIdOk() (*int32, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *PublicSingleSendRequestEgg) SetEmailId(v int32)`

SetEmailId sets EmailId field to given value.


### GetMessage

`func (o *PublicSingleSendRequestEgg) GetMessage() PublicSingleSendEmail`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PublicSingleSendRequestEgg) GetMessageOk() (*PublicSingleSendEmail, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PublicSingleSendRequestEgg) SetMessage(v PublicSingleSendEmail)`

SetMessage sets Message field to given value.


### GetContactProperties

`func (o *PublicSingleSendRequestEgg) GetContactProperties() map[string]string`

GetContactProperties returns the ContactProperties field if non-nil, zero value otherwise.

### GetContactPropertiesOk

`func (o *PublicSingleSendRequestEgg) GetContactPropertiesOk() (*map[string]string, bool)`

GetContactPropertiesOk returns a tuple with the ContactProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProperties

`func (o *PublicSingleSendRequestEgg) SetContactProperties(v map[string]string)`

SetContactProperties sets ContactProperties field to given value.

### HasContactProperties

`func (o *PublicSingleSendRequestEgg) HasContactProperties() bool`

HasContactProperties returns a boolean if a field has been set.

### GetCustomProperties

`func (o *PublicSingleSendRequestEgg) GetCustomProperties() map[string]map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *PublicSingleSendRequestEgg) GetCustomPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *PublicSingleSendRequestEgg) SetCustomProperties(v map[string]map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *PublicSingleSendRequestEgg) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


