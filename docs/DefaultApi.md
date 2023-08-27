# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalApiV1RowPost**](DefaultApi.md#ExternalApiV1RowPost) | **Post** /external/api/v1/row | Add a New Contacts Row
[**ExternalApiV1RowPut**](DefaultApi.md#ExternalApiV1RowPut) | **Put** /external/api/v1/row | Update an Existing Contacts Row

# **ExternalApiV1RowPost**
> ExternalApiV1RowPost(ctx, body)
Add a New Contacts Row

Add a new entry to the list of contacts on the Stapl dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| JSON data representing the contact details. The keys correspond to the dynamic column names created on the dashboard. | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalApiV1RowPut**
> ExternalApiV1RowPut(ctx, body)
Update an Existing Contacts Row

Update an already existing contact row on the Stapl dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1RowBody**](V1RowBody.md)| JSON data containing the key (email or ID field) and column data for update. Dynamic column names can be provided. | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

