# Gradient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SideOrCorner** | [**SideOrCorner**](SideOrCorner.md) |  | 
**Angle** | [**Angle**](Angle.md) |  | 
**Colors** | [**[]ColorStop**](ColorStop.md) |  | 

## Methods

### NewGradient

`func NewGradient(sideOrCorner SideOrCorner, angle Angle, colors []ColorStop, ) *Gradient`

NewGradient instantiates a new Gradient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradientWithDefaults

`func NewGradientWithDefaults() *Gradient`

NewGradientWithDefaults instantiates a new Gradient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSideOrCorner

`func (o *Gradient) GetSideOrCorner() SideOrCorner`

GetSideOrCorner returns the SideOrCorner field if non-nil, zero value otherwise.

### GetSideOrCornerOk

`func (o *Gradient) GetSideOrCornerOk() (*SideOrCorner, bool)`

GetSideOrCornerOk returns a tuple with the SideOrCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideOrCorner

`func (o *Gradient) SetSideOrCorner(v SideOrCorner)`

SetSideOrCorner sets SideOrCorner field to given value.


### GetAngle

`func (o *Gradient) GetAngle() Angle`

GetAngle returns the Angle field if non-nil, zero value otherwise.

### GetAngleOk

`func (o *Gradient) GetAngleOk() (*Angle, bool)`

GetAngleOk returns a tuple with the Angle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngle

`func (o *Gradient) SetAngle(v Angle)`

SetAngle sets Angle field to given value.


### GetColors

`func (o *Gradient) GetColors() []ColorStop`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *Gradient) GetColorsOk() (*[]ColorStop, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *Gradient) SetColors(v []ColorStop)`

SetColors sets Colors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


