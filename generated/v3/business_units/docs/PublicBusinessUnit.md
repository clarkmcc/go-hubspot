# PublicBusinessUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoMetadata** | Pointer to [**PublicBusinessUnitLogoMetadata**](PublicBusinessUnitLogoMetadata.md) |  | [optional] 
**Name** | **string** | The Business Unit&#39;s name | 
**Id** | **string** | The Business Unit&#39;s unique ID | 

## Methods

### NewPublicBusinessUnit

`func NewPublicBusinessUnit(name string, id string, ) *PublicBusinessUnit`

NewPublicBusinessUnit instantiates a new PublicBusinessUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicBusinessUnitWithDefaults

`func NewPublicBusinessUnitWithDefaults() *PublicBusinessUnit`

NewPublicBusinessUnitWithDefaults instantiates a new PublicBusinessUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoMetadata

`func (o *PublicBusinessUnit) GetLogoMetadata() PublicBusinessUnitLogoMetadata`

GetLogoMetadata returns the LogoMetadata field if non-nil, zero value otherwise.

### GetLogoMetadataOk

`func (o *PublicBusinessUnit) GetLogoMetadataOk() (*PublicBusinessUnitLogoMetadata, bool)`

GetLogoMetadataOk returns a tuple with the LogoMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoMetadata

`func (o *PublicBusinessUnit) SetLogoMetadata(v PublicBusinessUnitLogoMetadata)`

SetLogoMetadata sets LogoMetadata field to given value.

### HasLogoMetadata

`func (o *PublicBusinessUnit) HasLogoMetadata() bool`

HasLogoMetadata returns a boolean if a field has been set.

### GetName

`func (o *PublicBusinessUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicBusinessUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicBusinessUnit) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicBusinessUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicBusinessUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicBusinessUnit) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


