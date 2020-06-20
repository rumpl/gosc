# \NamespacesApi

All URIs are relative to *https://api.scaleway.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](NamespacesApi.md#CreateNamespace) | **Post** /functions/v1alpha2/regions/{region}/namespaces | Create a new namespace
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /functions/v1alpha2/regions/{region}/namespaces/{namespace_id} | Delete an existing namespace
[**GetNamespace**](NamespacesApi.md#GetNamespace) | **Get** /functions/v1alpha2/regions/{region}/namespaces/{namespace_id} | Get a namespace
[**ListNamespaces**](NamespacesApi.md#ListNamespaces) | **Get** /functions/v1alpha2/regions/{region}/namespaces | List all your namespaces
[**UpdateNamespace**](NamespacesApi.md#UpdateNamespace) | **Patch** /functions/v1alpha2/regions/{region}/namespaces/{namespace_id} | Update an existing namespace



## CreateNamespace

> ScalewayFunctionsV1alpha2Namespace CreateNamespace(ctx, region, inlineObject6)

Create a new namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Namespace**](scaleway.functions.v1alpha2.Namespace.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespace

> ScalewayFunctionsV1alpha2Namespace DeleteNamespace(ctx, region, namespaceId)

Delete an existing namespace

Delete the namespace associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**namespaceId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Namespace**](scaleway.functions.v1alpha2.Namespace.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespace

> ScalewayFunctionsV1alpha2Namespace GetNamespace(ctx, region, namespaceId)

Get a namespace

Get the namespace associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**namespaceId** | **string**|  | 

### Return type

[**ScalewayFunctionsV1alpha2Namespace**](scaleway.functions.v1alpha2.Namespace.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespaces

> ScalewayFunctionsV1alpha2ListNamespacesResponse ListNamespaces(ctx, region, optional)

List all your namespaces

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
 **optional** | ***ListNamespacesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNamespacesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Page number | [default to 1]
 **pageSize** | **optional.Float32**| Page size | [default to 20]
 **orderBy** | [**optional.Interface of ScalewayFunctionsV1alpha2ListNamespacesRequestOrderBy**](.md)|  | 
 **name** | **optional.String**|  | 
 **organizationId** | **optional.String**|  | 

### Return type

[**ScalewayFunctionsV1alpha2ListNamespacesResponse**](scaleway.functions.v1alpha2.ListNamespacesResponse.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNamespace

> ScalewayFunctionsV1alpha2Namespace UpdateNamespace(ctx, region, namespaceId, inlineObject7)

Update an existing namespace

Update the space associated with the given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| The region you want to target | 
**namespaceId** | **string**|  | 
**inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | 

### Return type

[**ScalewayFunctionsV1alpha2Namespace**](scaleway.functions.v1alpha2.Namespace.md)

### Authorization

[scaleway](../README.md#scaleway)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

