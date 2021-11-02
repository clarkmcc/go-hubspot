# ModelError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human readable message describing the error along with remediation steps where appropriate | [default to null]
**CorrelationId** | **string** | A unique identifier for the request. Include this value with any error reports or support tickets | [default to null]
**Category** | **string** | The error category | [default to null]
**SubCategory** | **string** | A specific category that contains more specific detail about the error | [optional] [default to null]
**Errors** | [**[]ErrorDetail**](ErrorDetail.md) | further information about the error | [optional] [default to null]
**Context** | [**map[string][]string**](array.md) | Context about the error condition | [optional] [default to null]
**Links** | **map[string]string** | A map of link names to associated URIs containing documentation about the error or recommended remediation steps | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

