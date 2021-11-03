# AccountingFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateInvoice** | [**CreateInvoiceFeature**](CreateInvoiceFeature.md) |  | 
**ImportInvoice** | [**ImportInvoiceFeature**](ImportInvoiceFeature.md) |  | 
**Sync** | [**map[string]ObjectSyncFeature**](ObjectSyncFeature.md) | Indicates if syncing objects from the external account system into HubSpot is supported for the integration. This is a map, where the key is one of &#x60;CONTACT&#x60; or &#x60;PRODUCT&#x60;, to indicate which type of object you do or don&#39;t support syncing. For example: &#x60;&#x60;&#x60;   \&quot;sync\&quot;: {     \&quot;CONTACT\&quot;: {       \&quot;toHubSpot\&quot;: true     },     \&quot;PRODUCT\&quot;: {       \&quot;toHubSpot\&quot;: true     }   } &#x60;&#x60;&#x60;  | 

## Methods

### NewAccountingFeatures

`func NewAccountingFeatures(createInvoice CreateInvoiceFeature, importInvoice ImportInvoiceFeature, sync map[string]ObjectSyncFeature, ) *AccountingFeatures`

NewAccountingFeatures instantiates a new AccountingFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingFeaturesWithDefaults

`func NewAccountingFeaturesWithDefaults() *AccountingFeatures`

NewAccountingFeaturesWithDefaults instantiates a new AccountingFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateInvoice

`func (o *AccountingFeatures) GetCreateInvoice() CreateInvoiceFeature`

GetCreateInvoice returns the CreateInvoice field if non-nil, zero value otherwise.

### GetCreateInvoiceOk

`func (o *AccountingFeatures) GetCreateInvoiceOk() (*CreateInvoiceFeature, bool)`

GetCreateInvoiceOk returns a tuple with the CreateInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInvoice

`func (o *AccountingFeatures) SetCreateInvoice(v CreateInvoiceFeature)`

SetCreateInvoice sets CreateInvoice field to given value.


### GetImportInvoice

`func (o *AccountingFeatures) GetImportInvoice() ImportInvoiceFeature`

GetImportInvoice returns the ImportInvoice field if non-nil, zero value otherwise.

### GetImportInvoiceOk

`func (o *AccountingFeatures) GetImportInvoiceOk() (*ImportInvoiceFeature, bool)`

GetImportInvoiceOk returns a tuple with the ImportInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportInvoice

`func (o *AccountingFeatures) SetImportInvoice(v ImportInvoiceFeature)`

SetImportInvoice sets ImportInvoice field to given value.


### GetSync

`func (o *AccountingFeatures) GetSync() map[string]ObjectSyncFeature`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *AccountingFeatures) GetSyncOk() (*map[string]ObjectSyncFeature, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *AccountingFeatures) SetSync(v map[string]ObjectSyncFeature)`

SetSync sets Sync field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


