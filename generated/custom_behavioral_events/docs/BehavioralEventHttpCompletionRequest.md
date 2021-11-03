# BehavioralEventHttpCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Utk** | Pointer to **string** | User token | [optional] 
**Email** | Pointer to **string** | Email of visitor | [optional] 
**EventName** | **string** | Internal name of the event-type to trigger | 
**Properties** | **map[string]string** | Map of properties for the event in the format property internal name - property value | 
**OccurredAt** | Pointer to **time.Time** | The time when this event occurred (if any). If this isn&#39;t set, the current time will be used | [optional] 
**ObjectId** | Pointer to **string** | The object id that this event occurred on. Could be a contact id or a visitor id. | [optional] 

## Methods

### NewBehavioralEventHttpCompletionRequest

`func NewBehavioralEventHttpCompletionRequest(eventName string, properties map[string]string, ) *BehavioralEventHttpCompletionRequest`

NewBehavioralEventHttpCompletionRequest instantiates a new BehavioralEventHttpCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehavioralEventHttpCompletionRequestWithDefaults

`func NewBehavioralEventHttpCompletionRequestWithDefaults() *BehavioralEventHttpCompletionRequest`

NewBehavioralEventHttpCompletionRequestWithDefaults instantiates a new BehavioralEventHttpCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtk

`func (o *BehavioralEventHttpCompletionRequest) GetUtk() string`

GetUtk returns the Utk field if non-nil, zero value otherwise.

### GetUtkOk

`func (o *BehavioralEventHttpCompletionRequest) GetUtkOk() (*string, bool)`

GetUtkOk returns a tuple with the Utk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtk

`func (o *BehavioralEventHttpCompletionRequest) SetUtk(v string)`

SetUtk sets Utk field to given value.

### HasUtk

`func (o *BehavioralEventHttpCompletionRequest) HasUtk() bool`

HasUtk returns a boolean if a field has been set.

### GetEmail

`func (o *BehavioralEventHttpCompletionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BehavioralEventHttpCompletionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BehavioralEventHttpCompletionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BehavioralEventHttpCompletionRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEventName

`func (o *BehavioralEventHttpCompletionRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *BehavioralEventHttpCompletionRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *BehavioralEventHttpCompletionRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetProperties

`func (o *BehavioralEventHttpCompletionRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BehavioralEventHttpCompletionRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BehavioralEventHttpCompletionRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BehavioralEventHttpCompletionRequest) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *BehavioralEventHttpCompletionRequest) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetObjectId

`func (o *BehavioralEventHttpCompletionRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BehavioralEventHttpCompletionRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BehavioralEventHttpCompletionRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BehavioralEventHttpCompletionRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


