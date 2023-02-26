# {{classname}}

All URIs are relative to *https://risk.charybdis.januus.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScoreBtcAddressGet**](DefaultApi.md#ScoreBtcAddressGet) | **Get** /score/btc/{address} | 
[**ScoreEthAddressGet**](DefaultApi.md#ScoreEthAddressGet) | **Get** /score/eth/{address} | 

# **ScoreBtcAddressGet**
> RiskReport ScoreBtcAddressGet(ctx, address)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**|  | 

### Return type

[**RiskReport**](RiskReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScoreEthAddressGet**
> RiskReport ScoreEthAddressGet(ctx, address)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**|  | 

### Return type

[**RiskReport**](RiskReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

