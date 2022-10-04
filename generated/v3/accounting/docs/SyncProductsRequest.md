# SyncProductsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account in the external accounting system. This is the value that will be passed as &#x60;accountId&#x60; for all outbound calls for the user from HubSpot to the external accounting system. | 
**Products** | [**[]UpdatedProduct**](UpdatedProduct.md) | A list of products to be imported. | 

## Methods

### NewSyncProductsRequest

`func NewSyncProductsRequest(accountId string, products []UpdatedProduct, ) *SyncProductsRequest`

NewSyncProductsRequest instantiates a new SyncProductsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncProductsRequestWithDefaults

`func NewSyncProductsRequestWithDefaults() *SyncProductsRequest`

NewSyncProductsRequestWithDefaults instantiates a new SyncProductsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SyncProductsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SyncProductsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SyncProductsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProducts

`func (o *SyncProductsRequest) GetProducts() []UpdatedProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *SyncProductsRequest) GetProductsOk() (*[]UpdatedProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *SyncProductsRequest) SetProducts(v []UpdatedProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


