# StoreApi

All URIs are relative to *http://petstore.swagger.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteOrder**](StoreApi.md#deleteOrder) | **DELETE** /store/order/{orderId} | Delete purchase order by ID
[**getInventory**](StoreApi.md#getInventory) | **GET** /store/inventory | Returns pet inventories by status
[**getOrderById**](StoreApi.md#getOrderById) | **GET** /store/order/{orderId} | Find purchase order by ID
[**placeOrder**](StoreApi.md#placeOrder) | **POST** /store/order | Place an order for a pet



## deleteOrder

> deleteOrder(orderId)

Delete purchase order by ID

For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors

### Example

```java
// Import classes:
//import org.openapitools.client.api.StoreApi;

StoreApi apiInstance = new StoreApi();
String orderId = null; // String | ID of the order that needs to be deleted
try {
    apiInstance.deleteOrder(orderId);
} catch (ApiException e) {
    System.err.println("Exception when calling StoreApi#deleteOrder");
    e.printStackTrace();
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **String**| ID of the order that needs to be deleted | [default to null]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## getInventory

> Map&lt;String, Integer&gt; getInventory()

Returns pet inventories by status

Returns a map of status codes to quantities

### Example

```java
// Import classes:
//import org.openapitools.client.api.StoreApi;

StoreApi apiInstance = new StoreApi();
try {
    Map<String, Integer> result = apiInstance.getInventory();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling StoreApi#getInventory");
    e.printStackTrace();
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Map&lt;String, Integer&gt;**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getOrderById

> Order getOrderById(orderId)

Find purchase order by ID

For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generate exceptions

### Example

```java
// Import classes:
//import org.openapitools.client.api.StoreApi;

StoreApi apiInstance = new StoreApi();
Long orderId = null; // Long | ID of pet that needs to be fetched
try {
    Order result = apiInstance.getOrderById(orderId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling StoreApi#getOrderById");
    e.printStackTrace();
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **Long**| ID of pet that needs to be fetched | [default to null]

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json


## placeOrder

> Order placeOrder(body)

Place an order for a pet

### Example

```java
// Import classes:
//import org.openapitools.client.api.StoreApi;

StoreApi apiInstance = new StoreApi();
Order body = new Order(); // Order | order placed for purchasing the pet
try {
    Order result = apiInstance.placeOrder(body);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling StoreApi#placeOrder");
    e.printStackTrace();
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Order**](Order.md)| order placed for purchasing the pet |

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

