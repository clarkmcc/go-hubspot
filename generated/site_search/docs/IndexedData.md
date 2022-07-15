# IndexedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the document in HubSpot. | 
**Type** | **string** | The type of document. Can be &#x60;SITE_PAGE&#x60;, &#x60;LANDING_PAGE&#x60;, &#x60;BLOG_POST&#x60;, &#x60;LISTING_PAGE&#x60;, or &#x60;KNOWLEDGE_ARTICLE&#x60;. | 
**Fields** | [**map[string]IndexedField**](IndexedField.md) | The indexed fields in HubSpot. | 

## Methods

### NewIndexedData

`func NewIndexedData(id string, type_ string, fields map[string]IndexedField, ) *IndexedData`

NewIndexedData instantiates a new IndexedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexedDataWithDefaults

`func NewIndexedDataWithDefaults() *IndexedData`

NewIndexedDataWithDefaults instantiates a new IndexedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexedData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexedData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexedData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IndexedData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndexedData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndexedData) SetType(v string)`

SetType sets Type field to given value.


### GetFields

`func (o *IndexedData) GetFields() map[string]IndexedField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexedData) GetFieldsOk() (*map[string]IndexedField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexedData) SetFields(v map[string]IndexedField)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


