# \LogsApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](LogsApi.md#ListLogs) | **Get** /functions/v1alpha2/regions/{region}/logs | List your application logs



## ListLogs

> ScalewayFunctionsV1alpha2ListLogsResponse ListLogs(ctx, region, optional)

List your application logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***ListLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationId** | **optional.String**|  | 
 **page** | **optional.Float32**| Page number | [default to 1]
 **pageSize** | **optional.Float32**| Page size | [default to 20]
 **orderBy** | [**optional.Interface of ScalewayFunctionsV1alpha2ListLogsRequestOrderBy**](.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2ListLogsResponse**](scaleway.functions.v1alpha2.ListLogsResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

