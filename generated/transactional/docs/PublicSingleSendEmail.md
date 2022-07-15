# PublicSingleSendEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | The From header for the email. | [optional] 
**To** | **string** | The recipient of the email. | 
**SendId** | Pointer to **string** | ID for a particular send. No more than one email will be sent per sendId. | [optional] 
**ReplyTo** | Pointer to **[]string** | List of Reply-To header values for the email. | [optional] 
**Cc** | Pointer to **[]string** | List of email addresses to send as Cc. | [optional] 
**Bcc** | Pointer to **[]string** | List of email addresses to send as Bcc. | [optional] 

## Methods

### NewPublicSingleSendEmail

`func NewPublicSingleSendEmail(to string, ) *PublicSingleSendEmail`

NewPublicSingleSendEmail instantiates a new PublicSingleSendEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSingleSendEmailWithDefaults

`func NewPublicSingleSendEmailWithDefaults() *PublicSingleSendEmail`

NewPublicSingleSendEmailWithDefaults instantiates a new PublicSingleSendEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PublicSingleSendEmail) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PublicSingleSendEmail) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PublicSingleSendEmail) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PublicSingleSendEmail) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PublicSingleSendEmail) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PublicSingleSendEmail) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PublicSingleSendEmail) SetTo(v string)`

SetTo sets To field to given value.


### GetSendId

`func (o *PublicSingleSendEmail) GetSendId() string`

GetSendId returns the SendId field if non-nil, zero value otherwise.

### GetSendIdOk

`func (o *PublicSingleSendEmail) GetSendIdOk() (*string, bool)`

GetSendIdOk returns a tuple with the SendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendId

`func (o *PublicSingleSendEmail) SetSendId(v string)`

SetSendId sets SendId field to given value.

### HasSendId

`func (o *PublicSingleSendEmail) HasSendId() bool`

HasSendId returns a boolean if a field has been set.

### GetReplyTo

`func (o *PublicSingleSendEmail) GetReplyTo() []string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *PublicSingleSendEmail) GetReplyToOk() (*[]string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *PublicSingleSendEmail) SetReplyTo(v []string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *PublicSingleSendEmail) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetCc

`func (o *PublicSingleSendEmail) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *PublicSingleSendEmail) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *PublicSingleSendEmail) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *PublicSingleSendEmail) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *PublicSingleSendEmail) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *PublicSingleSendEmail) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *PublicSingleSendEmail) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *PublicSingleSendEmail) HasBcc() bool`

HasBcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


