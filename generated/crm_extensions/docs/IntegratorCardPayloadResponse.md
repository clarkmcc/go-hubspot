# IntegratorCardPayloadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of card properties that will be sent in this response. | 
**AllItemsLinkUrl** | Pointer to **string** | URL to a page the integrator has built that displays all details for this card. This URL will be displayed to users under a &#x60;See more [x]&#x60; link if there are more than five items in your response, where &#x60;[x]&#x60; is the value of &#x60;itemLabel&#x60;. | [optional] 
**CardLabel** | Pointer to **string** | The label to be used for the &#x60;allItemsLinkUrl&#x60; link (e.g. &#39;See more tickets&#39;). If not provided, this falls back to the card&#39;s title. | [optional] 
**TopLevelActions** | Pointer to [**TopLevelActions**](TopLevelActions.md) |  | [optional] 
**Sections** | Pointer to [**[]IntegratorObjectResult**](IntegratorObjectResult.md) | A list of up to five valid card sub categories. | [optional] 
**ResponseVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegratorCardPayloadResponse

`func NewIntegratorCardPayloadResponse(totalCount int32, ) *IntegratorCardPayloadResponse`

NewIntegratorCardPayloadResponse instantiates a new IntegratorCardPayloadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorCardPayloadResponseWithDefaults

`func NewIntegratorCardPayloadResponseWithDefaults() *IntegratorCardPayloadResponse`

NewIntegratorCardPayloadResponseWithDefaults instantiates a new IntegratorCardPayloadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *IntegratorCardPayloadResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IntegratorCardPayloadResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IntegratorCardPayloadResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetAllItemsLinkUrl

`func (o *IntegratorCardPayloadResponse) GetAllItemsLinkUrl() string`

GetAllItemsLinkUrl returns the AllItemsLinkUrl field if non-nil, zero value otherwise.

### GetAllItemsLinkUrlOk

`func (o *IntegratorCardPayloadResponse) GetAllItemsLinkUrlOk() (*string, bool)`

GetAllItemsLinkUrlOk returns a tuple with the AllItemsLinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllItemsLinkUrl

`func (o *IntegratorCardPayloadResponse) SetAllItemsLinkUrl(v string)`

SetAllItemsLinkUrl sets AllItemsLinkUrl field to given value.

### HasAllItemsLinkUrl

`func (o *IntegratorCardPayloadResponse) HasAllItemsLinkUrl() bool`

HasAllItemsLinkUrl returns a boolean if a field has been set.

### GetCardLabel

`func (o *IntegratorCardPayloadResponse) GetCardLabel() string`

GetCardLabel returns the CardLabel field if non-nil, zero value otherwise.

### GetCardLabelOk

`func (o *IntegratorCardPayloadResponse) GetCardLabelOk() (*string, bool)`

GetCardLabelOk returns a tuple with the CardLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLabel

`func (o *IntegratorCardPayloadResponse) SetCardLabel(v string)`

SetCardLabel sets CardLabel field to given value.

### HasCardLabel

`func (o *IntegratorCardPayloadResponse) HasCardLabel() bool`

HasCardLabel returns a boolean if a field has been set.

### GetTopLevelActions

`func (o *IntegratorCardPayloadResponse) GetTopLevelActions() TopLevelActions`

GetTopLevelActions returns the TopLevelActions field if non-nil, zero value otherwise.

### GetTopLevelActionsOk

`func (o *IntegratorCardPayloadResponse) GetTopLevelActionsOk() (*TopLevelActions, bool)`

GetTopLevelActionsOk returns a tuple with the TopLevelActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelActions

`func (o *IntegratorCardPayloadResponse) SetTopLevelActions(v TopLevelActions)`

SetTopLevelActions sets TopLevelActions field to given value.

### HasTopLevelActions

`func (o *IntegratorCardPayloadResponse) HasTopLevelActions() bool`

HasTopLevelActions returns a boolean if a field has been set.

### GetSections

`func (o *IntegratorCardPayloadResponse) GetSections() []IntegratorObjectResult`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *IntegratorCardPayloadResponse) GetSectionsOk() (*[]IntegratorObjectResult, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *IntegratorCardPayloadResponse) SetSections(v []IntegratorObjectResult)`

SetSections sets Sections field to given value.

### HasSections

`func (o *IntegratorCardPayloadResponse) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetResponseVersion

`func (o *IntegratorCardPayloadResponse) GetResponseVersion() string`

GetResponseVersion returns the ResponseVersion field if non-nil, zero value otherwise.

### GetResponseVersionOk

`func (o *IntegratorCardPayloadResponse) GetResponseVersionOk() (*string, bool)`

GetResponseVersionOk returns a tuple with the ResponseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseVersion

`func (o *IntegratorCardPayloadResponse) SetResponseVersion(v string)`

SetResponseVersion sets ResponseVersion field to given value.

### HasResponseVersion

`func (o *IntegratorCardPayloadResponse) HasResponseVersion() bool`

HasResponseVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


