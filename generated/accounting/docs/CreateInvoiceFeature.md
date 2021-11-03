# CreateInvoiceFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if creating invoices in HubSpot is supported for the integration. | 
**SubFeatures** | [**CreateInvoiceSubFeatures**](CreateInvoiceSubFeatures.md) |  | 

## Methods

### NewCreateInvoiceFeature

`func NewCreateInvoiceFeature(enabled bool, subFeatures CreateInvoiceSubFeatures, ) *CreateInvoiceFeature`

NewCreateInvoiceFeature instantiates a new CreateInvoiceFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvoiceFeatureWithDefaults

`func NewCreateInvoiceFeatureWithDefaults() *CreateInvoiceFeature`

NewCreateInvoiceFeatureWithDefaults instantiates a new CreateInvoiceFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateInvoiceFeature) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateInvoiceFeature) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateInvoiceFeature) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSubFeatures

`func (o *CreateInvoiceFeature) GetSubFeatures() CreateInvoiceSubFeatures`

GetSubFeatures returns the SubFeatures field if non-nil, zero value otherwise.

### GetSubFeaturesOk

`func (o *CreateInvoiceFeature) GetSubFeaturesOk() (*CreateInvoiceSubFeatures, bool)`

GetSubFeaturesOk returns a tuple with the SubFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeatures

`func (o *CreateInvoiceFeature) SetSubFeatures(v CreateInvoiceSubFeatures)`

SetSubFeatures sets SubFeatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


