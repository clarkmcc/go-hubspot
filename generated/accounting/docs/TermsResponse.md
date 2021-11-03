# TermsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Designates if the response is a success (&#39;OK&#39;) or failure (&#39;ERR&#39;). | [optional] 
**Terms** | [**[]AccountingExtensionTerm**](AccountingExtensionTerm.md) | The list of payment terms that matched the search criteria. | 

## Methods

### NewTermsResponse

`func NewTermsResponse(terms []AccountingExtensionTerm, ) *TermsResponse`

NewTermsResponse instantiates a new TermsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsResponseWithDefaults

`func NewTermsResponseWithDefaults() *TermsResponse`

NewTermsResponseWithDefaults instantiates a new TermsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TermsResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TermsResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TermsResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TermsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTerms

`func (o *TermsResponse) GetTerms() []AccountingExtensionTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *TermsResponse) GetTermsOk() (*[]AccountingExtensionTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *TermsResponse) SetTerms(v []AccountingExtensionTerm)`

SetTerms sets Terms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


