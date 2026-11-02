# PropertyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculationFormula** | Pointer to **string** | Represents a formula that is used to compute a calculated property. | [optional] 
**CurrencyPropertyName** | Pointer to **string** | A string specifying the name of the property that holds the currency information, applicable when showCurrencySymbol is true. | [optional] 
**Description** | Pointer to **string** | A description of the property that will be shown as help text in HubSpot. | [optional] 
**DisplayOrder** | Pointer to **int32** | Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the Property to be displayed after any positive values. | [optional] 
**FieldType** | Pointer to **string** | Controls how the property appears in HubSpot. | [optional] 
**FormField** | Pointer to **bool** | Whether or not the property can be used in a HubSpot form. | [optional] 
**GroupName** | Pointer to **string** | The name of the property group the property belongs to. | [optional] 
**Hidden** | Pointer to **bool** | If true, the property won&#39;t be visible and can&#39;t be used in HubSpot. | [optional] 
**Label** | Pointer to **string** | A human-readable property label that will be shown in HubSpot. | [optional] 
**NumberDisplayHint** | Pointer to **string** | A string providing a hint for displaying number values. Valid values include &#39;unformatted&#39;, &#39;formatted&#39;, &#39;currency&#39;, &#39;percentage&#39;, &#39;duration&#39;, and &#39;probability&#39;. | [optional] 
**Options** | Pointer to [**[]OptionInput**](OptionInput.md) | A list of valid options for the property. | [optional] 
**ShowCurrencySymbol** | Pointer to **bool** | A boolean indicating whether to display a currency symbol for number fields formatted as currency. | [optional] 
**TextDisplayHint** | Pointer to **string** | A string providing a hint for displaying text values. Valid values include &#39;unformatted_single_line&#39;, &#39;multi_line&#39;, &#39;email&#39;, &#39;phone_number&#39;, &#39;domain_name&#39;, &#39;ip_address&#39;, &#39;physical_address&#39;, and &#39;postal_code&#39;. | [optional] 
**Type** | Pointer to **string** | The data type of the property. | [optional] 

## Methods

### NewPropertyUpdate

`func NewPropertyUpdate() *PropertyUpdate`

NewPropertyUpdate instantiates a new PropertyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyUpdateWithDefaults

`func NewPropertyUpdateWithDefaults() *PropertyUpdate`

NewPropertyUpdateWithDefaults instantiates a new PropertyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculationFormula

`func (o *PropertyUpdate) GetCalculationFormula() string`

GetCalculationFormula returns the CalculationFormula field if non-nil, zero value otherwise.

### GetCalculationFormulaOk

`func (o *PropertyUpdate) GetCalculationFormulaOk() (*string, bool)`

GetCalculationFormulaOk returns a tuple with the CalculationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationFormula

`func (o *PropertyUpdate) SetCalculationFormula(v string)`

SetCalculationFormula sets CalculationFormula field to given value.

### HasCalculationFormula

`func (o *PropertyUpdate) HasCalculationFormula() bool`

HasCalculationFormula returns a boolean if a field has been set.

### GetCurrencyPropertyName

`func (o *PropertyUpdate) GetCurrencyPropertyName() string`

GetCurrencyPropertyName returns the CurrencyPropertyName field if non-nil, zero value otherwise.

### GetCurrencyPropertyNameOk

`func (o *PropertyUpdate) GetCurrencyPropertyNameOk() (*string, bool)`

GetCurrencyPropertyNameOk returns a tuple with the CurrencyPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPropertyName

`func (o *PropertyUpdate) SetCurrencyPropertyName(v string)`

SetCurrencyPropertyName sets CurrencyPropertyName field to given value.

### HasCurrencyPropertyName

`func (o *PropertyUpdate) HasCurrencyPropertyName() bool`

HasCurrencyPropertyName returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *PropertyUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *PropertyUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *PropertyUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *PropertyUpdate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetFieldType

`func (o *PropertyUpdate) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PropertyUpdate) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PropertyUpdate) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PropertyUpdate) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetFormField

`func (o *PropertyUpdate) GetFormField() bool`

GetFormField returns the FormField field if non-nil, zero value otherwise.

### GetFormFieldOk

`func (o *PropertyUpdate) GetFormFieldOk() (*bool, bool)`

GetFormFieldOk returns a tuple with the FormField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormField

`func (o *PropertyUpdate) SetFormField(v bool)`

SetFormField sets FormField field to given value.

### HasFormField

`func (o *PropertyUpdate) HasFormField() bool`

HasFormField returns a boolean if a field has been set.

### GetGroupName

`func (o *PropertyUpdate) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PropertyUpdate) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PropertyUpdate) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PropertyUpdate) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHidden

`func (o *PropertyUpdate) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PropertyUpdate) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PropertyUpdate) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PropertyUpdate) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetLabel

`func (o *PropertyUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PropertyUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PropertyUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PropertyUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNumberDisplayHint

`func (o *PropertyUpdate) GetNumberDisplayHint() string`

GetNumberDisplayHint returns the NumberDisplayHint field if non-nil, zero value otherwise.

### GetNumberDisplayHintOk

`func (o *PropertyUpdate) GetNumberDisplayHintOk() (*string, bool)`

GetNumberDisplayHintOk returns a tuple with the NumberDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDisplayHint

`func (o *PropertyUpdate) SetNumberDisplayHint(v string)`

SetNumberDisplayHint sets NumberDisplayHint field to given value.

### HasNumberDisplayHint

`func (o *PropertyUpdate) HasNumberDisplayHint() bool`

HasNumberDisplayHint returns a boolean if a field has been set.

### GetOptions

`func (o *PropertyUpdate) GetOptions() []OptionInput`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PropertyUpdate) GetOptionsOk() (*[]OptionInput, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PropertyUpdate) SetOptions(v []OptionInput)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PropertyUpdate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetShowCurrencySymbol

`func (o *PropertyUpdate) GetShowCurrencySymbol() bool`

GetShowCurrencySymbol returns the ShowCurrencySymbol field if non-nil, zero value otherwise.

### GetShowCurrencySymbolOk

`func (o *PropertyUpdate) GetShowCurrencySymbolOk() (*bool, bool)`

GetShowCurrencySymbolOk returns a tuple with the ShowCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencySymbol

`func (o *PropertyUpdate) SetShowCurrencySymbol(v bool)`

SetShowCurrencySymbol sets ShowCurrencySymbol field to given value.

### HasShowCurrencySymbol

`func (o *PropertyUpdate) HasShowCurrencySymbol() bool`

HasShowCurrencySymbol returns a boolean if a field has been set.

### GetTextDisplayHint

`func (o *PropertyUpdate) GetTextDisplayHint() string`

GetTextDisplayHint returns the TextDisplayHint field if non-nil, zero value otherwise.

### GetTextDisplayHintOk

`func (o *PropertyUpdate) GetTextDisplayHintOk() (*string, bool)`

GetTextDisplayHintOk returns a tuple with the TextDisplayHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDisplayHint

`func (o *PropertyUpdate) SetTextDisplayHint(v string)`

SetTextDisplayHint sets TextDisplayHint field to given value.

### HasTextDisplayHint

`func (o *PropertyUpdate) HasTextDisplayHint() bool`

HasTextDisplayHint returns a boolean if a field has been set.

### GetType

`func (o *PropertyUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


