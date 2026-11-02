# ObjectSchemaBatchReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAssociationDefinitions** | **bool** | Indicates whether to include association definitions in the response. | 
**IncludeAuditMetadata** | **bool** | Indicates whether to include audit metadata in the response. | 
**IncludePropertyDefinitions** | **bool** | Indicates whether to include property definitions in the response. | 
**Inputs** | **[]string** | An array of strings representing the identifiers of the object schemas to be read. | 

## Methods

### NewObjectSchemaBatchReadRequest

`func NewObjectSchemaBatchReadRequest(includeAssociationDefinitions bool, includeAuditMetadata bool, includePropertyDefinitions bool, inputs []string, ) *ObjectSchemaBatchReadRequest`

NewObjectSchemaBatchReadRequest instantiates a new ObjectSchemaBatchReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSchemaBatchReadRequestWithDefaults

`func NewObjectSchemaBatchReadRequestWithDefaults() *ObjectSchemaBatchReadRequest`

NewObjectSchemaBatchReadRequestWithDefaults instantiates a new ObjectSchemaBatchReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAssociationDefinitions

`func (o *ObjectSchemaBatchReadRequest) GetIncludeAssociationDefinitions() bool`

GetIncludeAssociationDefinitions returns the IncludeAssociationDefinitions field if non-nil, zero value otherwise.

### GetIncludeAssociationDefinitionsOk

`func (o *ObjectSchemaBatchReadRequest) GetIncludeAssociationDefinitionsOk() (*bool, bool)`

GetIncludeAssociationDefinitionsOk returns a tuple with the IncludeAssociationDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAssociationDefinitions

`func (o *ObjectSchemaBatchReadRequest) SetIncludeAssociationDefinitions(v bool)`

SetIncludeAssociationDefinitions sets IncludeAssociationDefinitions field to given value.


### GetIncludeAuditMetadata

`func (o *ObjectSchemaBatchReadRequest) GetIncludeAuditMetadata() bool`

GetIncludeAuditMetadata returns the IncludeAuditMetadata field if non-nil, zero value otherwise.

### GetIncludeAuditMetadataOk

`func (o *ObjectSchemaBatchReadRequest) GetIncludeAuditMetadataOk() (*bool, bool)`

GetIncludeAuditMetadataOk returns a tuple with the IncludeAuditMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuditMetadata

`func (o *ObjectSchemaBatchReadRequest) SetIncludeAuditMetadata(v bool)`

SetIncludeAuditMetadata sets IncludeAuditMetadata field to given value.


### GetIncludePropertyDefinitions

`func (o *ObjectSchemaBatchReadRequest) GetIncludePropertyDefinitions() bool`

GetIncludePropertyDefinitions returns the IncludePropertyDefinitions field if non-nil, zero value otherwise.

### GetIncludePropertyDefinitionsOk

`func (o *ObjectSchemaBatchReadRequest) GetIncludePropertyDefinitionsOk() (*bool, bool)`

GetIncludePropertyDefinitionsOk returns a tuple with the IncludePropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePropertyDefinitions

`func (o *ObjectSchemaBatchReadRequest) SetIncludePropertyDefinitions(v bool)`

SetIncludePropertyDefinitions sets IncludePropertyDefinitions field to given value.


### GetInputs

`func (o *ObjectSchemaBatchReadRequest) GetInputs() []string`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ObjectSchemaBatchReadRequest) GetInputsOk() (*[]string, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ObjectSchemaBatchReadRequest) SetInputs(v []string)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


