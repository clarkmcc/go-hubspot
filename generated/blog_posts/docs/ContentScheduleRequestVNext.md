# ContentScheduleRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the object to be scheduled. | 
**PublishDate** | **time.Time** | The date the object should transition from scheduled to published. | 

## Methods

### NewContentScheduleRequestVNext

`func NewContentScheduleRequestVNext(id string, publishDate time.Time, ) *ContentScheduleRequestVNext`

NewContentScheduleRequestVNext instantiates a new ContentScheduleRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentScheduleRequestVNextWithDefaults

`func NewContentScheduleRequestVNextWithDefaults() *ContentScheduleRequestVNext`

NewContentScheduleRequestVNextWithDefaults instantiates a new ContentScheduleRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentScheduleRequestVNext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentScheduleRequestVNext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentScheduleRequestVNext) SetId(v string)`

SetId sets Id field to given value.


### GetPublishDate

`func (o *ContentScheduleRequestVNext) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *ContentScheduleRequestVNext) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *ContentScheduleRequestVNext) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


