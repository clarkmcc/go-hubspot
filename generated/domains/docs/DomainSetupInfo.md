# DomainSetupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsSslExternally** | **bool** |  | 
**WhoIsEmailAddresses** | **[]string** |  | 

## Methods

### NewDomainSetupInfo

`func NewDomainSetupInfo(supportsSslExternally bool, whoIsEmailAddresses []string, ) *DomainSetupInfo`

NewDomainSetupInfo instantiates a new DomainSetupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSetupInfoWithDefaults

`func NewDomainSetupInfoWithDefaults() *DomainSetupInfo`

NewDomainSetupInfoWithDefaults instantiates a new DomainSetupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsSslExternally

`func (o *DomainSetupInfo) GetSupportsSslExternally() bool`

GetSupportsSslExternally returns the SupportsSslExternally field if non-nil, zero value otherwise.

### GetSupportsSslExternallyOk

`func (o *DomainSetupInfo) GetSupportsSslExternallyOk() (*bool, bool)`

GetSupportsSslExternallyOk returns a tuple with the SupportsSslExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSslExternally

`func (o *DomainSetupInfo) SetSupportsSslExternally(v bool)`

SetSupportsSslExternally sets SupportsSslExternally field to given value.


### GetWhoIsEmailAddresses

`func (o *DomainSetupInfo) GetWhoIsEmailAddresses() []string`

GetWhoIsEmailAddresses returns the WhoIsEmailAddresses field if non-nil, zero value otherwise.

### GetWhoIsEmailAddressesOk

`func (o *DomainSetupInfo) GetWhoIsEmailAddressesOk() (*[]string, bool)`

GetWhoIsEmailAddressesOk returns a tuple with the WhoIsEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhoIsEmailAddresses

`func (o *DomainSetupInfo) SetWhoIsEmailAddresses(v []string)`

SetWhoIsEmailAddresses sets WhoIsEmailAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


