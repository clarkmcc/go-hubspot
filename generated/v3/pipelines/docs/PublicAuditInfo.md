# PublicAuditInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**RawObject** | Pointer to **map[string]interface{}** |  | [optional] 
**FromUserId** | Pointer to **int32** |  | [optional] 
**PortalId** | **int32** |  | 
**Action** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublicAuditInfo

`func NewPublicAuditInfo(identifier string, portalId int32, action string, ) *PublicAuditInfo`

NewPublicAuditInfo instantiates a new PublicAuditInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAuditInfoWithDefaults

`func NewPublicAuditInfoWithDefaults() *PublicAuditInfo`

NewPublicAuditInfoWithDefaults instantiates a new PublicAuditInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PublicAuditInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PublicAuditInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PublicAuditInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetRawObject

`func (o *PublicAuditInfo) GetRawObject() map[string]interface{}`

GetRawObject returns the RawObject field if non-nil, zero value otherwise.

### GetRawObjectOk

`func (o *PublicAuditInfo) GetRawObjectOk() (*map[string]interface{}, bool)`

GetRawObjectOk returns a tuple with the RawObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawObject

`func (o *PublicAuditInfo) SetRawObject(v map[string]interface{})`

SetRawObject sets RawObject field to given value.

### HasRawObject

`func (o *PublicAuditInfo) HasRawObject() bool`

HasRawObject returns a boolean if a field has been set.

### GetFromUserId

`func (o *PublicAuditInfo) GetFromUserId() int32`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *PublicAuditInfo) GetFromUserIdOk() (*int32, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *PublicAuditInfo) SetFromUserId(v int32)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *PublicAuditInfo) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetPortalId

`func (o *PublicAuditInfo) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *PublicAuditInfo) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *PublicAuditInfo) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetAction

`func (o *PublicAuditInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PublicAuditInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PublicAuditInfo) SetAction(v string)`

SetAction sets Action field to given value.


### GetMessage

`func (o *PublicAuditInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PublicAuditInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PublicAuditInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PublicAuditInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *PublicAuditInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicAuditInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicAuditInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PublicAuditInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


