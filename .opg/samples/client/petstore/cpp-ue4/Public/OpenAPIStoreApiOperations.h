/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator
 * https://github.com/OpenAPITools/openapi-generator
 * Do not edit the class manually.
 */

#pragma once

#include "OpenAPIBaseModel.h"
#include "OpenAPIStoreApi.h"

#include "OpenAPIOrder.h"

namespace OpenAPI
{

/* Delete purchase order by ID
 *
 * For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors
*/
class OPENAPI_API OpenAPIStoreApi::DeleteOrderRequest : public Request
{
public:
    virtual ~DeleteOrderRequest() {}
	void SetupHttpRequest(const FHttpRequestRef& HttpRequest) const final;
	FString ComputePath() const final;

	/* ID of the order that needs to be deleted */
	FString OrderId;
};

class OPENAPI_API OpenAPIStoreApi::DeleteOrderResponse : public Response
{
public:
    virtual ~DeleteOrderResponse() {}
	void SetHttpResponseCode(EHttpResponseCodes::Type InHttpResponseCode) final;
	bool FromJson(const TSharedPtr<FJsonValue>& JsonValue) final;

    
};

/* Returns pet inventories by status
 *
 * Returns a map of status codes to quantities
*/
class OPENAPI_API OpenAPIStoreApi::GetInventoryRequest : public Request
{
public:
    virtual ~GetInventoryRequest() {}
	void SetupHttpRequest(const FHttpRequestRef& HttpRequest) const final;
	FString ComputePath() const final;

};

class OPENAPI_API OpenAPIStoreApi::GetInventoryResponse : public Response
{
public:
    virtual ~GetInventoryResponse() {}
	void SetHttpResponseCode(EHttpResponseCodes::Type InHttpResponseCode) final;
	bool FromJson(const TSharedPtr<FJsonValue>& JsonValue) final;

    TMap<FString, int32> Content;
};

/* Find purchase order by ID
 *
 * For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generate exceptions
*/
class OPENAPI_API OpenAPIStoreApi::GetOrderByIdRequest : public Request
{
public:
    virtual ~GetOrderByIdRequest() {}
	void SetupHttpRequest(const FHttpRequestRef& HttpRequest) const final;
	FString ComputePath() const final;

	/* ID of pet that needs to be fetched */
	int64 OrderId = 0;
};

class OPENAPI_API OpenAPIStoreApi::GetOrderByIdResponse : public Response
{
public:
    virtual ~GetOrderByIdResponse() {}
	void SetHttpResponseCode(EHttpResponseCodes::Type InHttpResponseCode) final;
	bool FromJson(const TSharedPtr<FJsonValue>& JsonValue) final;

    OpenAPIOrder Content;
};

/* Place an order for a pet

*/
class OPENAPI_API OpenAPIStoreApi::PlaceOrderRequest : public Request
{
public:
    virtual ~PlaceOrderRequest() {}
	void SetupHttpRequest(const FHttpRequestRef& HttpRequest) const final;
	FString ComputePath() const final;

	/* order placed for purchasing the pet */
	OpenAPIOrder Body;
};

class OPENAPI_API OpenAPIStoreApi::PlaceOrderResponse : public Response
{
public:
    virtual ~PlaceOrderResponse() {}
	void SetHttpResponseCode(EHttpResponseCodes::Type InHttpResponseCode) final;
	bool FromJson(const TSharedPtr<FJsonValue>& JsonValue) final;

    OpenAPIOrder Content;
};

}
