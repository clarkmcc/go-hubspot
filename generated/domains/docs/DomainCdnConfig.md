# DomainCdnConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortalId** | **int32** |  | 
**Created** | **int64** |  | 
**Updated** | **int64** |  | 
**DomainId** | **int64** |  | 
**CertId** | **int64** |  | 
**CertExpirationDate** | **int64** |  | 
**CdnId** | **string** |  | 
**CdnGroupId** | **string** |  | 
**SslStatus** | **string** |  | 
**ValidationMethod** | **string** |  | 
**HttpBody** | **string** |  | 
**HttpUrlPath** | **string** |  | 
**Cname** | **string** |  | 
**CnameTarget** | **string** |  | 
**MinimumTlsVersion** | **string** |  | 
**AlternateOriginHostname** | **string** |  | 
**Id** | **int64** |  | 

## Methods

### NewDomainCdnConfig

`func NewDomainCdnConfig(portalId int32, created int64, updated int64, domainId int64, certId int64, certExpirationDate int64, cdnId string, cdnGroupId string, sslStatus string, validationMethod string, httpBody string, httpUrlPath string, cname string, cnameTarget string, minimumTlsVersion string, alternateOriginHostname string, id int64, ) *DomainCdnConfig`

NewDomainCdnConfig instantiates a new DomainCdnConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCdnConfigWithDefaults

`func NewDomainCdnConfigWithDefaults() *DomainCdnConfig`

NewDomainCdnConfigWithDefaults instantiates a new DomainCdnConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortalId

`func (o *DomainCdnConfig) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *DomainCdnConfig) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *DomainCdnConfig) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.


### GetCreated

`func (o *DomainCdnConfig) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DomainCdnConfig) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DomainCdnConfig) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *DomainCdnConfig) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DomainCdnConfig) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DomainCdnConfig) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.


### GetDomainId

`func (o *DomainCdnConfig) GetDomainId() int64`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainCdnConfig) GetDomainIdOk() (*int64, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainCdnConfig) SetDomainId(v int64)`

SetDomainId sets DomainId field to given value.


### GetCertId

`func (o *DomainCdnConfig) GetCertId() int64`

GetCertId returns the CertId field if non-nil, zero value otherwise.

### GetCertIdOk

`func (o *DomainCdnConfig) GetCertIdOk() (*int64, bool)`

GetCertIdOk returns a tuple with the CertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertId

`func (o *DomainCdnConfig) SetCertId(v int64)`

SetCertId sets CertId field to given value.


### GetCertExpirationDate

`func (o *DomainCdnConfig) GetCertExpirationDate() int64`

GetCertExpirationDate returns the CertExpirationDate field if non-nil, zero value otherwise.

### GetCertExpirationDateOk

`func (o *DomainCdnConfig) GetCertExpirationDateOk() (*int64, bool)`

GetCertExpirationDateOk returns a tuple with the CertExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpirationDate

`func (o *DomainCdnConfig) SetCertExpirationDate(v int64)`

SetCertExpirationDate sets CertExpirationDate field to given value.


### GetCdnId

`func (o *DomainCdnConfig) GetCdnId() string`

GetCdnId returns the CdnId field if non-nil, zero value otherwise.

### GetCdnIdOk

`func (o *DomainCdnConfig) GetCdnIdOk() (*string, bool)`

GetCdnIdOk returns a tuple with the CdnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnId

`func (o *DomainCdnConfig) SetCdnId(v string)`

SetCdnId sets CdnId field to given value.


### GetCdnGroupId

`func (o *DomainCdnConfig) GetCdnGroupId() string`

GetCdnGroupId returns the CdnGroupId field if non-nil, zero value otherwise.

### GetCdnGroupIdOk

`func (o *DomainCdnConfig) GetCdnGroupIdOk() (*string, bool)`

GetCdnGroupIdOk returns a tuple with the CdnGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnGroupId

`func (o *DomainCdnConfig) SetCdnGroupId(v string)`

SetCdnGroupId sets CdnGroupId field to given value.


### GetSslStatus

`func (o *DomainCdnConfig) GetSslStatus() string`

GetSslStatus returns the SslStatus field if non-nil, zero value otherwise.

### GetSslStatusOk

`func (o *DomainCdnConfig) GetSslStatusOk() (*string, bool)`

GetSslStatusOk returns a tuple with the SslStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslStatus

`func (o *DomainCdnConfig) SetSslStatus(v string)`

SetSslStatus sets SslStatus field to given value.


### GetValidationMethod

`func (o *DomainCdnConfig) GetValidationMethod() string`

GetValidationMethod returns the ValidationMethod field if non-nil, zero value otherwise.

### GetValidationMethodOk

`func (o *DomainCdnConfig) GetValidationMethodOk() (*string, bool)`

GetValidationMethodOk returns a tuple with the ValidationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMethod

`func (o *DomainCdnConfig) SetValidationMethod(v string)`

SetValidationMethod sets ValidationMethod field to given value.


### GetHttpBody

`func (o *DomainCdnConfig) GetHttpBody() string`

GetHttpBody returns the HttpBody field if non-nil, zero value otherwise.

### GetHttpBodyOk

`func (o *DomainCdnConfig) GetHttpBodyOk() (*string, bool)`

GetHttpBodyOk returns a tuple with the HttpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpBody

`func (o *DomainCdnConfig) SetHttpBody(v string)`

SetHttpBody sets HttpBody field to given value.


### GetHttpUrlPath

`func (o *DomainCdnConfig) GetHttpUrlPath() string`

GetHttpUrlPath returns the HttpUrlPath field if non-nil, zero value otherwise.

### GetHttpUrlPathOk

`func (o *DomainCdnConfig) GetHttpUrlPathOk() (*string, bool)`

GetHttpUrlPathOk returns a tuple with the HttpUrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUrlPath

`func (o *DomainCdnConfig) SetHttpUrlPath(v string)`

SetHttpUrlPath sets HttpUrlPath field to given value.


### GetCname

`func (o *DomainCdnConfig) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *DomainCdnConfig) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *DomainCdnConfig) SetCname(v string)`

SetCname sets Cname field to given value.


### GetCnameTarget

`func (o *DomainCdnConfig) GetCnameTarget() string`

GetCnameTarget returns the CnameTarget field if non-nil, zero value otherwise.

### GetCnameTargetOk

`func (o *DomainCdnConfig) GetCnameTargetOk() (*string, bool)`

GetCnameTargetOk returns a tuple with the CnameTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameTarget

`func (o *DomainCdnConfig) SetCnameTarget(v string)`

SetCnameTarget sets CnameTarget field to given value.


### GetMinimumTlsVersion

`func (o *DomainCdnConfig) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *DomainCdnConfig) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *DomainCdnConfig) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.


### GetAlternateOriginHostname

`func (o *DomainCdnConfig) GetAlternateOriginHostname() string`

GetAlternateOriginHostname returns the AlternateOriginHostname field if non-nil, zero value otherwise.

### GetAlternateOriginHostnameOk

`func (o *DomainCdnConfig) GetAlternateOriginHostnameOk() (*string, bool)`

GetAlternateOriginHostnameOk returns a tuple with the AlternateOriginHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateOriginHostname

`func (o *DomainCdnConfig) SetAlternateOriginHostname(v string)`

SetAlternateOriginHostname sets AlternateOriginHostname field to given value.


### GetId

`func (o *DomainCdnConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainCdnConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainCdnConfig) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


