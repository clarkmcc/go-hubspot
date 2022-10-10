# PublicAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **string** | The ID of the object. | 
**UserId** | **string** | The ID of the user who caused the event. | 
**Timestamp** | **time.Time** | The timestamp at which the event occurred. | 
**ObjectName** | **string** | The internal name of the object in HubSpot. | 
**FullName** | **string** | The name of the user who caused the event. | 
**Event** | **string** | The type of event that took place (CREATED, UPDATED, PUBLISHED, DELETED, UNPUBLISHED). | 
**ObjectType** | **string** | The type of the object (BLOG, LANDING_PAGE, DOMAIN, HUBDB_TABLE etc.) | 

## Methods

### NewPublicAuditLog

`func NewPublicAuditLog(objectId string, userId string, timestamp time.Time, objectName string, fullName string, event string, objectType string, ) *PublicAuditLog`

NewPublicAuditLog instantiates a new PublicAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAuditLogWithDefaults

`func NewPublicAuditLogWithDefaults() *PublicAuditLog`

NewPublicAuditLogWithDefaults instantiates a new PublicAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *PublicAuditLog) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PublicAuditLog) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PublicAuditLog) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetUserId

`func (o *PublicAuditLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PublicAuditLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PublicAuditLog) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTimestamp

`func (o *PublicAuditLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicAuditLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicAuditLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetObjectName

`func (o *PublicAuditLog) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *PublicAuditLog) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *PublicAuditLog) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### GetFullName

`func (o *PublicAuditLog) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PublicAuditLog) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PublicAuditLog) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEvent

`func (o *PublicAuditLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PublicAuditLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PublicAuditLog) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetObjectType

`func (o *PublicAuditLog) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicAuditLog) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicAuditLog) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


