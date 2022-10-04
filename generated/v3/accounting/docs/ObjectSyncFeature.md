# ObjectSyncFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToHubSpot** | **bool** | Indicates if syncing the object type from the external accounting system into HubSpot is supported. | 

## Methods

### NewObjectSyncFeature

`func NewObjectSyncFeature(toHubSpot bool, ) *ObjectSyncFeature`

NewObjectSyncFeature instantiates a new ObjectSyncFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSyncFeatureWithDefaults

`func NewObjectSyncFeatureWithDefaults() *ObjectSyncFeature`

NewObjectSyncFeatureWithDefaults instantiates a new ObjectSyncFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToHubSpot

`func (o *ObjectSyncFeature) GetToHubSpot() bool`

GetToHubSpot returns the ToHubSpot field if non-nil, zero value otherwise.

### GetToHubSpotOk

`func (o *ObjectSyncFeature) GetToHubSpotOk() (*bool, bool)`

GetToHubSpotOk returns a tuple with the ToHubSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToHubSpot

`func (o *ObjectSyncFeature) SetToHubSpot(v bool)`

SetToHubSpot sets ToHubSpot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


