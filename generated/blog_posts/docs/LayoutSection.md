# LayoutSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int32** |  | 
**W** | **int32** |  | 
**Name** | **string** |  | 
**Label** | **string** |  | 
**Type** | **string** |  | 
**Params** | **map[string]map[string]interface{}** | null | 
**Rows** | [**[]map[string]LayoutSection**](map[string]LayoutSection.md) |  | 
**RowMetaData** | [**[]RowMetaData**](RowMetaData.md) |  | 
**Cells** | [**[]LayoutSection**](LayoutSection.md) |  | 
**CssClass** | **string** |  | 
**CssStyle** | **string** |  | 
**CssId** | **string** |  | 
**Styles** | [**Styles**](Styles.md) |  | 

## Methods

### NewLayoutSection

`func NewLayoutSection(x int32, w int32, name string, label string, type_ string, params map[string]map[string]interface{}, rows []map[string]LayoutSection, rowMetaData []RowMetaData, cells []LayoutSection, cssClass string, cssStyle string, cssId string, styles Styles, ) *LayoutSection`

NewLayoutSection instantiates a new LayoutSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutSectionWithDefaults

`func NewLayoutSectionWithDefaults() *LayoutSection`

NewLayoutSectionWithDefaults instantiates a new LayoutSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *LayoutSection) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *LayoutSection) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *LayoutSection) SetX(v int32)`

SetX sets X field to given value.


### GetW

`func (o *LayoutSection) GetW() int32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *LayoutSection) GetWOk() (*int32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *LayoutSection) SetW(v int32)`

SetW sets W field to given value.


### GetName

`func (o *LayoutSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LayoutSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LayoutSection) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *LayoutSection) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LayoutSection) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LayoutSection) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *LayoutSection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LayoutSection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LayoutSection) SetType(v string)`

SetType sets Type field to given value.


### GetParams

`func (o *LayoutSection) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *LayoutSection) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *LayoutSection) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.


### GetRows

`func (o *LayoutSection) GetRows() []map[string]LayoutSection`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *LayoutSection) GetRowsOk() (*[]map[string]LayoutSection, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *LayoutSection) SetRows(v []map[string]LayoutSection)`

SetRows sets Rows field to given value.


### GetRowMetaData

`func (o *LayoutSection) GetRowMetaData() []RowMetaData`

GetRowMetaData returns the RowMetaData field if non-nil, zero value otherwise.

### GetRowMetaDataOk

`func (o *LayoutSection) GetRowMetaDataOk() (*[]RowMetaData, bool)`

GetRowMetaDataOk returns a tuple with the RowMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowMetaData

`func (o *LayoutSection) SetRowMetaData(v []RowMetaData)`

SetRowMetaData sets RowMetaData field to given value.


### GetCells

`func (o *LayoutSection) GetCells() []LayoutSection`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *LayoutSection) GetCellsOk() (*[]LayoutSection, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *LayoutSection) SetCells(v []LayoutSection)`

SetCells sets Cells field to given value.


### GetCssClass

`func (o *LayoutSection) GetCssClass() string`

GetCssClass returns the CssClass field if non-nil, zero value otherwise.

### GetCssClassOk

`func (o *LayoutSection) GetCssClassOk() (*string, bool)`

GetCssClassOk returns a tuple with the CssClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssClass

`func (o *LayoutSection) SetCssClass(v string)`

SetCssClass sets CssClass field to given value.


### GetCssStyle

`func (o *LayoutSection) GetCssStyle() string`

GetCssStyle returns the CssStyle field if non-nil, zero value otherwise.

### GetCssStyleOk

`func (o *LayoutSection) GetCssStyleOk() (*string, bool)`

GetCssStyleOk returns a tuple with the CssStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssStyle

`func (o *LayoutSection) SetCssStyle(v string)`

SetCssStyle sets CssStyle field to given value.


### GetCssId

`func (o *LayoutSection) GetCssId() string`

GetCssId returns the CssId field if non-nil, zero value otherwise.

### GetCssIdOk

`func (o *LayoutSection) GetCssIdOk() (*string, bool)`

GetCssIdOk returns a tuple with the CssId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssId

`func (o *LayoutSection) SetCssId(v string)`

SetCssId sets CssId field to given value.


### GetStyles

`func (o *LayoutSection) GetStyles() Styles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *LayoutSection) GetStylesOk() (*Styles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *LayoutSection) SetStyles(v Styles)`

SetStyles sets Styles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


