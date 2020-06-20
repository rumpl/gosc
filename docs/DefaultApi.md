# \DefaultApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployContainer**](DefaultApi.md#DeployContainer) | **Post** /functions/v1alpha2/regions/{region}/containers/{container_id}/deploy | 
[**DeployFunction**](DefaultApi.md#DeployFunction) | **Post** /functions/v1alpha2/regions/{region}/functions/{function_id}/deploy | 
[**GetFunctionDownloadURL**](DefaultApi.md#GetFunctionDownloadURL) | **Get** /functions/v1alpha2/regions/{region}/functions/{function_id}/download-url | 
[**GetFunctionUploadURL**](DefaultApi.md#GetFunctionUploadURL) | **Get** /functions/v1alpha2/regions/{region}/functions/{function_id}/upload-url | 
[**IssueJWT**](DefaultApi.md#IssueJWT) | **Get** /functions/v1alpha2/regions/{region}/jwt/issue | 



## DeployContainer

> ScalewayFunctionsV1alpha2Container DeployContainer(ctx, region, containerId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**containerId** | **string**|  | 
**body** | **map[string]interface{}**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Container**](scaleway.functions.v1alpha2.Container.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployFunction

> ScalewayFunctionsV1alpha2Function DeployFunction(ctx, region, functionId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 
**body** | **map[string]interface{}**|  | 

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


## GetFunctionDownloadURL

> ScalewayFunctionsV1alpha2DownloadUrl GetFunctionDownloadURL(ctx, region, functionId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2DownloadUrl**](scaleway.functions.v1alpha2.DownloadURL.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionUploadURL

> ScalewayFunctionsV1alpha2UploadUrl GetFunctionUploadURL(ctx, region, functionId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**functionId** | **string**|  | 
 **optional** | ***GetFunctionUploadURLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFunctionUploadURLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentLength** | **optional.Float32**|  | 

### Return type

[**ScalewayFunctionsV1alpha2UploadUrl**](scaleway.functions.v1alpha2.UploadURL.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueJWT

> ScalewayFunctionsV1alpha2Token IssueJWT(ctx, region, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***IssueJWTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IssueJWTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionId** | **optional.String**|  | 
 **containerId** | **optional.String**|  | 
 **namespaceId** | **optional.String**|  | 
 **expirationDate** | **optional.Time**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Token**](scaleway.functions.v1alpha2.Token.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

