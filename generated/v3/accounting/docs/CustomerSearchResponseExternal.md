# CustomerSearchResponseExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | 
**Customers** | [**[]AccountingExtensionCustomer**](AccountingExtensionCustomer.md) | The list of customers that matched the search criteria. | 

## Methods

### NewCustomerSearchResponseExternal

`func NewCustomerSearchResponseExternal(result string, customers []AccountingExtensionCustomer, ) *CustomerSearchResponseExternal`

NewCustomerSearchResponseExternal instantiates a new CustomerSearchResponseExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerSearchResponseExternalWithDefaults

`func NewCustomerSearchResponseExternalWithDefaults() *CustomerSearchResponseExternal`

NewCustomerSearchResponseExternalWithDefaults instantiates a new CustomerSearchResponseExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CustomerSearchResponseExternal) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CustomerSearchResponseExternal) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CustomerSearchResponseExternal) SetResult(v string)`

SetResult sets Result field to given value.


### GetCustomers

`func (o *CustomerSearchResponseExternal) GetCustomers() []AccountingExtensionCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CustomerSearchResponseExternal) GetCustomersOk() (*[]AccountingExtensionCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CustomerSearchResponseExternal) SetCustomers(v []AccountingExtensionCustomer)`

SetCustomers sets Customers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


