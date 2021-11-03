# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | [**UnitPrice**](UnitPrice.md) |  | 
**TaxExempt** | **bool** | Identifies if the product is tax exempt or not. | 
**SalesTaxType** | Pointer to [**TaxType**](TaxType.md) |  | [optional] 
**Name** | **string** | The display name of the product. | 
**Description** | Pointer to **string** | The description of the product. | [optional] 
**Id** | **string** | The ID of the product in the external accounting system. | 

## Methods

### NewProduct

`func NewProduct(unitPrice UnitPrice, taxExempt bool, name string, id string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *Product) GetUnitPrice() UnitPrice`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Product) GetUnitPriceOk() (*UnitPrice, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Product) SetUnitPrice(v UnitPrice)`

SetUnitPrice sets UnitPrice field to given value.


### GetTaxExempt

`func (o *Product) GetTaxExempt() bool`

GetTaxExempt returns the TaxExempt field if non-nil, zero value otherwise.

### GetTaxExemptOk

`func (o *Product) GetTaxExemptOk() (*bool, bool)`

GetTaxExemptOk returns a tuple with the TaxExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExempt

`func (o *Product) SetTaxExempt(v bool)`

SetTaxExempt sets TaxExempt field to given value.


### GetSalesTaxType

`func (o *Product) GetSalesTaxType() TaxType`

GetSalesTaxType returns the SalesTaxType field if non-nil, zero value otherwise.

### GetSalesTaxTypeOk

`func (o *Product) GetSalesTaxTypeOk() (*TaxType, bool)`

GetSalesTaxTypeOk returns a tuple with the SalesTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxType

`func (o *Product) SetSalesTaxType(v TaxType)`

SetSalesTaxType sets SalesTaxType field to given value.

### HasSalesTaxType

`func (o *Product) HasSalesTaxType() bool`

HasSalesTaxType returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


