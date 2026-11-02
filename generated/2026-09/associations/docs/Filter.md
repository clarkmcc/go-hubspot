# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighValue** | Pointer to **string** | The upper bound value used when the operator is a range type, such as &#39;BETWEEN&#39;. It is a string. | [optional] 
**Operator** | **string** | The operator used to compare the property value. Valid operators include &#39;EQ&#39;, &#39;NEQ&#39;, &#39;LT&#39;, &#39;LTE&#39;, &#39;GT&#39;, &#39;GTE&#39;, &#39;BETWEEN&#39;, &#39;IN&#39;, &#39;NOT_IN&#39;, &#39;HAS_PROPERTY&#39;, &#39;NOT_HAS_PROPERTY&#39;, &#39;CONTAINS_TOKEN&#39;, and &#39;NOT_CONTAINS_TOKEN&#39;. It is a string. | 
**PropertyName** | **string** | The name of the property on which the filter is applied. It is a string. | 
**Value** | Pointer to **string** | The specific value that the property should match. It is a string. | [optional] 
**Values** | Pointer to **[]string** | An array of values used for operators that require multiple values, such as &#39;IN&#39; or &#39;NOT_IN&#39;. Each value in the array is a string. | [optional] 

## Methods

### NewFilter

`func NewFilter(operator string, propertyName string, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighValue

`func (o *Filter) GetHighValue() string`

GetHighValue returns the HighValue field if non-nil, zero value otherwise.

### GetHighValueOk

`func (o *Filter) GetHighValueOk() (*string, bool)`

GetHighValueOk returns a tuple with the HighValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighValue

`func (o *Filter) SetHighValue(v string)`

SetHighValue sets HighValue field to given value.

### HasHighValue

`func (o *Filter) HasHighValue() bool`

HasHighValue returns a boolean if a field has been set.

### GetOperator

`func (o *Filter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Filter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Filter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPropertyName

`func (o *Filter) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *Filter) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *Filter) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.


### GetValue

`func (o *Filter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Filter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Filter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Filter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *Filter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Filter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Filter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Filter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


