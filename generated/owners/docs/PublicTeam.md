# PublicTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Membership** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicTeam

`func NewPublicTeam(id string, name string, ) *PublicTeam`

NewPublicTeam instantiates a new PublicTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTeamWithDefaults

`func NewPublicTeamWithDefaults() *PublicTeam`

NewPublicTeamWithDefaults instantiates a new PublicTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTeam) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PublicTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicTeam) SetName(v string)`

SetName sets Name field to given value.


### GetMembership

`func (o *PublicTeam) GetMembership() string`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *PublicTeam) GetMembershipOk() (*string, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *PublicTeam) SetMembership(v string)`

SetMembership sets Membership field to given value.

### HasMembership

`func (o *PublicTeam) HasMembership() bool`

HasMembership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


