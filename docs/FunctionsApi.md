# \FunctionsApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](FunctionsApi.md#CreateFunction) | **Post** /functions/v1alpha2/regions/{region}/functions | Create a new function
[**DeleteFunction**](FunctionsApi.md#DeleteFunction) | **Delete** /functions/v1alpha2/regions/{region}/functions/{function_id} | Delete a function
[**GetFunction**](FunctionsApi.md#GetFunction) | **Get** /functions/v1alpha2/regions/{region}/functions/{function_id} | Get a function
[**ListFunctionRuntimes**](FunctionsApi.md#ListFunctionRuntimes) | **Get** /functions/v1alpha2/regions/{region}/runtimes | List function runtimes
[**ListFunctions**](FunctionsApi.md#ListFunctions) | **Get** /functions/v1alpha2/regions/{region}/functions | List all your functions
[**UpdateFunction**](FunctionsApi.md#UpdateFunction) | **Patch** /functions/v1alpha2/regions/{region}/functions/{function_id} | Update an existing function



## CreateFunction

> ScalewayFunctionsV1alpha2Function CreateFunction(ctx, region, inlineObject4)

Create a new function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Function**](scaleway.functions.v1alpha2.Function.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> ScalewayFunctionsV1alpha2Function DeleteFunction(ctx, region, functionId)

Delete a function

Delete the function associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Function**](scaleway.functions.v1alpha2.Function.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> ScalewayFunctionsV1alpha2Function GetFunction(ctx, region, functionId)

Get a function

Get the function associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Function**](scaleway.functions.v1alpha2.Function.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionRuntimes

> ScalewayFunctionsV1alpha2ListFunctionRuntimesResponse ListFunctionRuntimes(ctx, region)

List function runtimes

List available function runtimes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 

### Return type

[**ScalewayFunctionsV1alpha2ListFunctionRuntimesResponse**](scaleway.functions.v1alpha2.ListFunctionRuntimesResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctions

> ScalewayFunctionsV1alpha2ListFunctionsResponse ListFunctions(ctx, region, optional)

List all your functions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***ListFunctionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFunctionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Page number | [default to 1]
 **pageSize** | **optional.Float32**| Page size | [default to 20]
 **orderBy** | [**optional.Interface of ScalewayFunctionsV1alpha2ListFunctionsRequestOrderBy**](.md)|  | 
 **namespaceId** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **organizationId** | **optional.String**|  | 

### Return type

[**ScalewayFunctionsV1alpha2ListFunctionsResponse**](scaleway.functions.v1alpha2.ListFunctionsResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> ScalewayFunctionsV1alpha2Function UpdateFunction(ctx, region, functionId, inlineObject5)

Update an existing function

Update the function associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Function**](scaleway.functions.v1alpha2.Function.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

