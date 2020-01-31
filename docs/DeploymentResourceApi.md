# \DeploymentResourceApi

All URIs are relative to *https://ververica.prod.bird.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentUsingPOST**](DeploymentResourceApi.md#CreateDeploymentUsingPOST) | **Post** /api/v1/namespaces/{namespace}/deployments | Create a deployment
[**DeleteDeploymentUsingDELETE**](DeploymentResourceApi.md#DeleteDeploymentUsingDELETE) | **Delete** /api/v1/namespaces/{namespace}/deployments/{deploymentId} | Delete deployment
[**GetDeploymentUsingGET**](DeploymentResourceApi.md#GetDeploymentUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployments/{deploymentId} | Get a deployment by id
[**GetDeploymentsUsingGET**](DeploymentResourceApi.md#GetDeploymentsUsingGET) | **Get** /api/v1/namespaces/{namespace}/deployments | List all deployments
[**UpdateDeploymentUsingPATCH**](DeploymentResourceApi.md#UpdateDeploymentUsingPATCH) | **Patch** /api/v1/namespaces/{namespace}/deployments/{deploymentId} | Update a deployment


# **CreateDeploymentUsingPOST**
> Deployment CreateDeploymentUsingPOST(ctx, deploymentChange, namespace)
Create a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentChange** | [**Deployment**](Deployment.md)| deploymentChange | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeploymentUsingDELETE**
> Deployment DeleteDeploymentUsingDELETE(ctx, deploymentId, namespace)
Delete deployment

This operation expects the deployment to be in desired state CANCELLED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | [**string**](.md)| deploymentId | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentUsingGET**
> Deployment GetDeploymentUsingGET(ctx, deploymentId, namespace)
Get a deployment by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deploymentId** | [**string**](.md)| deploymentId | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsUsingGET**
> ResourceListDeployment GetDeploymentsUsingGET(ctx, namespace, optional)
List all deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**| namespace | 
 **optional** | ***GetDeploymentsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDeploymentsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSelector** | **optional.String**| labelSelector | 

### Return type

[**ResourceListDeployment**](ResourceList«Deployment».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeploymentUsingPATCH**
> Deployment UpdateDeploymentUsingPATCH(ctx, body, deploymentId, namespace)
Update a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComDataartisansAppmanagerApiV1EntityDeploymentDeployment**](ComDataartisansAppmanagerApiV1EntityDeploymentDeployment.md)|  | 
  **deploymentId** | [**string**](.md)| deploymentId | 
  **namespace** | **string**| namespace | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

