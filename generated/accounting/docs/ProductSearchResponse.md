# ProductSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Products** | [**[]Product**](Product.md) | The list of products that matched the search criteria. | 

## Methods

### NewProductSearchResponse

`func NewProductSearchResponse(products []Product, ) *ProductSearchResponse`

NewProductSearchResponse instantiates a new ProductSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchResponseWithDefaults

`func NewProductSearchResponseWithDefaults() *ProductSearchResponse`

NewProductSearchResponseWithDefaults instantiates a new ProductSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ProductSearchResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProductSearchResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProductSearchResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ProductSearchResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetProducts

`func (o *ProductSearchResponse) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductSearchResponse) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductSearchResponse) SetProducts(v []Product)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


