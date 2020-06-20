# \CronsApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCron**](CronsApi.md#CreateCron) | **Post** /functions/v1alpha2/regions/{region}/crons | Create a new cron
[**DeleteCron**](CronsApi.md#DeleteCron) | **Delete** /functions/v1alpha2/regions/{region}/crons/{cron_id} | Delete an existing cron
[**GetCron**](CronsApi.md#GetCron) | **Get** /functions/v1alpha2/regions/{region}/crons/{cron_id} | Get a cron
[**ListCrons**](CronsApi.md#ListCrons) | **Get** /functions/v1alpha2/regions/{region}/crons | List all your crons
[**UpdateCron**](CronsApi.md#UpdateCron) | **Patch** /functions/v1alpha2/regions/{region}/crons/{cron_id} | Update an existing cron



## CreateCron

> ScalewayFunctionsV1alpha2Cron CreateCron(ctx, region, inlineObject2)

Create a new cron

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Cron**](scaleway.functions.v1alpha2.Cron.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCron

> ScalewayFunctionsV1alpha2Cron DeleteCron(ctx, region, cronId)

Delete an existing cron

Delete the cron associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**cronId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Cron**](scaleway.functions.v1alpha2.Cron.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCron

> ScalewayFunctionsV1alpha2Cron GetCron(ctx, region, cronId)

Get a cron

Get the cron associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**cronId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Cron**](scaleway.functions.v1alpha2.Cron.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCrons

> ScalewayFunctionsV1alpha2ListCronsResponse ListCrons(ctx, region, optional)

List all your crons

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***ListCronsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCronsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Page number | [default to 1]
 **pageSize** | **optional.Float32**| Page size | [default to 20]
 **orderBy** | [**optional.Interface of ScalewayFunctionsV1alpha2ListCronsRequestOrderBy**](.md)|  | 
 **applicationId** | **optional.String**|  | 

### Return type

[**ScalewayFunctionsV1alpha2ListCronsResponse**](scaleway.functions.v1alpha2.ListCronsResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCron

> ScalewayFunctionsV1alpha2Cron UpdateCron(ctx, region, cronId, inlineObject3)

Update an existing cron

Update the cron associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**cronId** | **string**|  | 
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Cron**](scaleway.functions.v1alpha2.Cron.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

