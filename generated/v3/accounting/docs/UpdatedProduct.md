# UpdatedProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncAction** | **string** | The operation to be performed. | 
**UpdatedAt** | **time.Time** | The timestamp (ISO8601 format) when the product was updated in the external accounting system. | 
**Price** | **float32** | The price of the product. | 
**CurrencyCode** | Pointer to **string** | The ISO 4217 currency code that represents the currency of the product price. | [optional] 
**Id** | **string** | The ID of the product in the external accounting system. | 
**Properties** | **map[string]string** | A map of key-value product properties to be imported. | 

## Methods

### NewUpdatedProduct

`func NewUpdatedProduct(syncAction string, updatedAt time.Time, price float32, id string, properties map[string]string, ) *UpdatedProduct`

NewUpdatedProduct instantiates a new UpdatedProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedProductWithDefaults

`func NewUpdatedProductWithDefaults() *UpdatedProduct`

NewUpdatedProductWithDefaults instantiates a new UpdatedProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncAction

`func (o *UpdatedProduct) GetSyncAction() string`

GetSyncAction returns the SyncAction field if non-nil, zero value otherwise.

### GetSyncActionOk

`func (o *UpdatedProduct) GetSyncActionOk() (*string, bool)`

GetSyncActionOk returns a tuple with the SyncAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAction

`func (o *UpdatedProduct) SetSyncAction(v string)`

SetSyncAction sets SyncAction field to given value.


### GetUpdatedAt

`func (o *UpdatedProduct) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatedProduct) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatedProduct) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPrice

`func (o *UpdatedProduct) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatedProduct) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatedProduct) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCurrencyCode

`func (o *UpdatedProduct) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *UpdatedProduct) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *UpdatedProduct) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *UpdatedProduct) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetId

`func (o *UpdatedProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedProduct) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *UpdatedProduct) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpdatedProduct) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpdatedProduct) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


