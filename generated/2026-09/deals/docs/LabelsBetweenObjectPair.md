# LabelsBetweenObjectPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromObjectId** | **string** | Source unique ID of the object. | 
**FromObjectTypeId** | **string** | Source object type. | 
**Labels** | **[]string** | An array of strings representing the labels that describe the nature of the association between the two objects. | 
**ToObjectId** | **string** | Target unique ID of the object. | 
**ToObjectTypeId** | **string** | Target object type. | 

## Methods

### NewLabelsBetweenObjectPair

`func NewLabelsBetweenObjectPair(fromObjectId string, fromObjectTypeId string, labels []string, toObjectId string, toObjectTypeId string, ) *LabelsBetweenObjectPair`

NewLabelsBetweenObjectPair instantiates a new LabelsBetweenObjectPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsBetweenObjectPairWithDefaults

`func NewLabelsBetweenObjectPairWithDefaults() *LabelsBetweenObjectPair`

NewLabelsBetweenObjectPairWithDefaults instantiates a new LabelsBetweenObjectPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromObjectId

`func (o *LabelsBetweenObjectPair) GetFromObjectId() string`

GetFromObjectId returns the FromObjectId field if non-nil, zero value otherwise.

### GetFromObjectIdOk

`func (o *LabelsBetweenObjectPair) GetFromObjectIdOk() (*string, bool)`

GetFromObjectIdOk returns a tuple with the FromObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectId

`func (o *LabelsBetweenObjectPair) SetFromObjectId(v string)`

SetFromObjectId sets FromObjectId field to given value.


### GetFromObjectTypeId

`func (o *LabelsBetweenObjectPair) GetFromObjectTypeId() string`

GetFromObjectTypeId returns the FromObjectTypeId field if non-nil, zero value otherwise.

### GetFromObjectTypeIdOk

`func (o *LabelsBetweenObjectPair) GetFromObjectTypeIdOk() (*string, bool)`

GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectTypeId

`func (o *LabelsBetweenObjectPair) SetFromObjectTypeId(v string)`

SetFromObjectTypeId sets FromObjectTypeId field to given value.


### GetLabels

`func (o *LabelsBetweenObjectPair) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LabelsBetweenObjectPair) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LabelsBetweenObjectPair) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetToObjectId

`func (o *LabelsBetweenObjectPair) GetToObjectId() string`

GetToObjectId returns the ToObjectId field if non-nil, zero value otherwise.

### GetToObjectIdOk

`func (o *LabelsBetweenObjectPair) GetToObjectIdOk() (*string, bool)`

GetToObjectIdOk returns a tuple with the ToObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectId

`func (o *LabelsBetweenObjectPair) SetToObjectId(v string)`

SetToObjectId sets ToObjectId field to given value.


### GetToObjectTypeId

`func (o *LabelsBetweenObjectPair) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *LabelsBetweenObjectPair) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *LabelsBetweenObjectPair) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


