# \ContainersApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainer**](ContainersApi.md#CreateContainer) | **Post** /functions/v1alpha2/regions/{region}/containers | Create a new container
[**DeleteContainer**](ContainersApi.md#DeleteContainer) | **Delete** /functions/v1alpha2/regions/{region}/containers/{container_id} | Delete a container
[**GetContainer**](ContainersApi.md#GetContainer) | **Get** /functions/v1alpha2/regions/{region}/containers/{container_id} | Get a container
[**ListContainers**](ContainersApi.md#ListContainers) | **Get** /functions/v1alpha2/regions/{region}/containers | List all your containers
[**UpdateContainer**](ContainersApi.md#UpdateContainer) | **Patch** /functions/v1alpha2/regions/{region}/containers/{container_id} | Update an existing container



## CreateContainer

> ScalewayFunctionsV1alpha2Container CreateContainer(ctx, region, inlineObject)

Create a new container

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

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


## DeleteContainer

> ScalewayFunctionsV1alpha2Container DeleteContainer(ctx, region, containerId)

Delete a container

Delete the container associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**containerId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Container**](scaleway.functions.v1alpha2.Container.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainer

> ScalewayFunctionsV1alpha2Container GetContainer(ctx, region, containerId)

Get a container

Get the container associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**containerId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Container**](scaleway.functions.v1alpha2.Container.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainers

> ScalewayFunctionsV1alpha2ListContainersResponse ListContainers(ctx, region, optional)

List all your containers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***ListContainersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListContainersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Page number | [default to 1]
 **pageSize** | **optional.Float32**| Page size | [default to 20]
 **orderBy** | [**optional.Interface of ScalewayFunctionsV1alpha2ListContainersRequestOrderBy**](.md)|  | 
 **namespaceId** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **organizationId** | **optional.String**|  | 

### Return type

[**ScalewayFunctionsV1alpha2ListContainersResponse**](scaleway.functions.v1alpha2.ListContainersResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainer

> ScalewayFunctionsV1alpha2Container UpdateContainer(ctx, region, containerId, inlineObject1)

Update an existing container

Update the container associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**containerId** | **string**|  | 
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

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

