# \DeploymentTargetResourceApi

All URIs are relative to *https://ververica.prod.bird.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentTargetUsingPOST**](DeploymentTargetResourceApi.md#CreateDeploymentTargetUsingPOST) | **Post** /api/v1/namespaces/{namespace}/deployment-targets | Create a deployment target
[**DeleteDeploymentTargetUsingDELETE**](DeploymentTargetResourceApi.md#DeleteDeploymentTargetUsingDELETE) | **Delete** /api/v1/namespaces/{namespace}/deployment-targets/{name} | Delete a deployment target
[**GetDeploymentTargetUsingGET**](DeploymentTargetResourceApi.md#GetDeploymentTargetUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployment-targets/{name} | Get a deployment target by name
[**GetDeploymentTargetsUsingGET**](DeploymentTargetResourceApi.md#GetDeploymentTargetsUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployment-targets | List all deployment targets


# **CreateDeploymentTargetUsingPOST**
> DeploymentTarget CreateDeploymentTargetUsingPOST(ctx, deploymentTarget, namespace)
Create a deployment target

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentTarget** | [**DeploymentTarget**](DeploymentTarget.md)| deploymentTarget | 
  **namespace** | **string**| namespace | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeploymentTargetUsingDELETE**
> DeploymentTarget DeleteDeploymentTargetUsingDELETE(ctx, name, namespace)
Delete a deployment target

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTargetUsingGET**
> DeploymentTarget GetDeploymentTargetUsingGET(ctx, name, namespace)
Get a deployment target by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **namespace** | **string**| namespace | 

### Return type

[**DeploymentTarget**](DeploymentTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentTargetsUsingGET**
> ResourceListDeploymentTarget GetDeploymentTargetsUsingGET(ctx, namespace)
List all deployment targets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 

### Return type

[**ResourceListDeploymentTarget**](ResourceList«DeploymentTarget».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

