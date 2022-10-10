# IdentificationTokenGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the visitor that you wish to identify | 
**FirstName** | Pointer to **string** | The first name of the visitor that you wish to identify. This value will only be set in HubSpot for new contacts and existing contacts where first name is unknown. Optional. | [optional] 
**LastName** | Pointer to **string** | The last name of the visitor that you wish to identify. This value will only be set in HubSpot for new contacts and existing contacts where last name is unknown. Optional. | [optional] 

## Methods

### NewIdentificationTokenGenerationRequest

`func NewIdentificationTokenGenerationRequest(email string, ) *IdentificationTokenGenerationRequest`

NewIdentificationTokenGenerationRequest instantiates a new IdentificationTokenGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationTokenGenerationRequestWithDefaults

`func NewIdentificationTokenGenerationRequestWithDefaults() *IdentificationTokenGenerationRequest`

NewIdentificationTokenGenerationRequestWithDefaults instantiates a new IdentificationTokenGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *IdentificationTokenGenerationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentificationTokenGenerationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentificationTokenGenerationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *IdentificationTokenGenerationRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IdentificationTokenGenerationRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IdentificationTokenGenerationRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IdentificationTokenGenerationRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *IdentificationTokenGenerationRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IdentificationTokenGenerationRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IdentificationTokenGenerationRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IdentificationTokenGenerationRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


