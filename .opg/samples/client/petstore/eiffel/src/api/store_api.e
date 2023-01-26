note
 description:"[
		OpenAPI Petstore
 		This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
  		The version of the OpenAPI document: 1.0.0
 	    

  	NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

 		 Do not edit the class manually.
 	]"
	date: "$Date$"
	revision: "$Revision$"
	EIS:"Eiffel openapi generator", "src=https://openapi-generator.tech", "protocol=uri"

class
	STORE_API

inherit

    API_I


feature -- API Access


	delete_order (order_id: STRING_32)
			-- Delete purchase order by ID
			-- For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors
			-- 
			-- argument: order_id ID of the order that needs to be deleted (required)
			-- 
			-- 
		require
		local
  			l_path: STRING
  			l_request: API_CLIENT_REQUEST
  			l_response: API_CLIENT_RESPONSE
		do
			reset_error
			create l_request
			
			l_path := "/store/order/{order_id}"
			l_path.replace_substring_all ("{"+"order_id"+"}", api_client.url_encode (order_id.out))


			if attached {STRING} api_client.select_header_accept ({ARRAY [STRING]}<<>>)  as l_accept then
				l_request.add_header(l_accept,"Accept");
			end
			l_request.add_header(api_client.select_header_content_type ({ARRAY [STRING]}<<>>),"Content-Type")
			l_request.set_auth_names ({ARRAY [STRING]}<<>>)
			l_response := api_client.call_api (l_path, "Delete", l_request, agent serializer, Void)
			if l_response.has_error then
				last_error := l_response.error
			end
		end	

	inventory : detachable STRING_TABLE [INTEGER_32]
			-- Returns pet inventories by status
			-- Returns a map of status codes to quantities
			-- 
			-- 
			-- Result STRING_TABLE [INTEGER_32]
		require
		local
  			l_path: STRING
  			l_request: API_CLIENT_REQUEST
  			l_response: API_CLIENT_RESPONSE
		do
			reset_error
			create l_request
			
			l_path := "/store/inventory"


			if attached {STRING} api_client.select_header_accept ({ARRAY [STRING]}<<"application/json">>)  as l_accept then
				l_request.add_header(l_accept,"Accept");
			end
			l_request.add_header(api_client.select_header_content_type ({ARRAY [STRING]}<<>>),"Content-Type")
			l_request.set_auth_names ({ARRAY [STRING]}<<"api_key">>)
			l_response := api_client.call_api (l_path, "Get", l_request, Void, agent deserializer)
			if l_response.has_error then
				last_error := l_response.error
			elseif attached { STRING_TABLE [INTEGER_32] } l_response.data ({ STRING_TABLE [INTEGER_32] }) as l_data then
				Result := l_data
			else
				create last_error.make ("Unknown error: Status response [ " + l_response.status.out + "]")
			end
		end	

	order_by_id (order_id: INTEGER_64): detachable ORDER
			-- Find purchase order by ID
			-- For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generate exceptions
			-- 
			-- argument: order_id ID of pet that needs to be fetched (required)
			-- 
			-- 
			-- Result ORDER
		require
			order_id_is_less_or_equal_than: order_id <= 5 
     		order_id_is_greater_or_equal_than: order_id >= 1 
		local
  			l_path: STRING
  			l_request: API_CLIENT_REQUEST
  			l_response: API_CLIENT_RESPONSE
		do
			reset_error
			create l_request
			
			l_path := "/store/order/{order_id}"
			l_path.replace_substring_all ("{"+"order_id"+"}", api_client.url_encode (order_id.out))


			if attached {STRING} api_client.select_header_accept ({ARRAY [STRING]}<<"application/xml", "application/json">>)  as l_accept then
				l_request.add_header(l_accept,"Accept");
			end
			l_request.add_header(api_client.select_header_content_type ({ARRAY [STRING]}<<>>),"Content-Type")
			l_request.set_auth_names ({ARRAY [STRING]}<<>>)
			l_response := api_client.call_api (l_path, "Get", l_request, Void, agent deserializer)
			if l_response.has_error then
				last_error := l_response.error
			elseif attached { ORDER } l_response.data ({ ORDER }) as l_data then
				Result := l_data
			else
				create last_error.make ("Unknown error: Status response [ " + l_response.status.out + "]")
			end
		end	

	place_order (body: ORDER): detachable ORDER
			-- Place an order for a pet
			-- 
			-- 
			-- argument: body order placed for purchasing the pet (required)
			-- 
			-- 
			-- Result ORDER
		require
		local
  			l_path: STRING
  			l_request: API_CLIENT_REQUEST
  			l_response: API_CLIENT_RESPONSE
		do
			reset_error
			create l_request
			l_request.set_body(body)
			l_path := "/store/order"


			if attached {STRING} api_client.select_header_accept ({ARRAY [STRING]}<<"application/xml", "application/json">>)  as l_accept then
				l_request.add_header(l_accept,"Accept");
			end
			l_request.add_header(api_client.select_header_content_type ({ARRAY [STRING]}<<>>),"Content-Type")
			l_request.set_auth_names ({ARRAY [STRING]}<<>>)
			l_response := api_client.call_api (l_path, "Post", l_request, Void, agent deserializer)
			if l_response.has_error then
				last_error := l_response.error
			elseif attached { ORDER } l_response.data ({ ORDER }) as l_data then
				Result := l_data
			else
				create last_error.make ("Unknown error: Status response [ " + l_response.status.out + "]")
			end
		end	


end
