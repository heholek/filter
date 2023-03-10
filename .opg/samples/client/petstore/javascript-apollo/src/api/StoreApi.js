/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */



import ApiClient from "../ApiClient";
import Order from '../model/Order';

/**
* Store service.
* @module api/StoreApi
* @version 1.0.0
*/
export default class StoreApi extends ApiClient {

    /**
    * Constructs a new StoreApi. 
    * @alias module:api/StoreApi
    * @class
    */
    constructor(baseURL = 'http://petstore.swagger.io:80/v2') {
      super(baseURL);
    }


    /**
     * Delete purchase order by ID
     * For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
     * @param {String} orderId ID of the order that needs to be deleted
     * @param requestInit Dynamic configuration. @see {@link https://github.com/apollographql/apollo-server/pull/1277}
     * @return {Promise}
     */
    async deleteOrder(orderId, requestInit) {
      let postBody = null;
      // verify the required parameter 'orderId' is set
      if (orderId === undefined || orderId === null) {
        throw new Error("Missing the required parameter 'orderId' when calling deleteOrder");
      }

      let pathParams = {
        'order_id': orderId
      };
      let queryParams = {
      };
      let headerParams = {
        'User-Agent': 'OpenAPI-Generator/1.0.0/Javascript',
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = [];
      let returnType = null;

      return this.callApi(
        '/store/order/{order_id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, requestInit
      );
    }

    /**
     * Returns pet inventories by status
     * Returns a map of status codes to quantities
     * @param requestInit Dynamic configuration. @see {@link https://github.com/apollographql/apollo-server/pull/1277}
     * @return {Promise<Object.<String, {String: Number}>>}
     */
    async getInventory(requestInit) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
        'User-Agent': 'OpenAPI-Generator/1.0.0/Javascript',
      };
      let formParams = {
      };

      let authNames = ['api_key'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = {'String': 'Number'};

      return this.callApi(
        '/store/inventory', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, requestInit
      );
    }

    /**
     * Find purchase order by ID
     * For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions
     * @param {Number} orderId ID of pet that needs to be fetched
     * @param requestInit Dynamic configuration. @see {@link https://github.com/apollographql/apollo-server/pull/1277}
     * @return {Promise<module:model/Order>}
     */
    async getOrderById(orderId, requestInit) {
      let postBody = null;
      // verify the required parameter 'orderId' is set
      if (orderId === undefined || orderId === null) {
        throw new Error("Missing the required parameter 'orderId' when calling getOrderById");
      }

      let pathParams = {
        'order_id': orderId
      };
      let queryParams = {
      };
      let headerParams = {
        'User-Agent': 'OpenAPI-Generator/1.0.0/Javascript',
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/xml', 'application/json'];
      let returnType = Order;

      return this.callApi(
        '/store/order/{order_id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, requestInit
      );
    }

    /**
     * Place an order for a pet
     * 
     * @param {module:model/Order} order order placed for purchasing the pet
     * @param requestInit Dynamic configuration. @see {@link https://github.com/apollographql/apollo-server/pull/1277}
     * @return {Promise<module:model/Order>}
     */
    async placeOrder(order, requestInit) {
      let postBody = order;
      // verify the required parameter 'order' is set
      if (order === undefined || order === null) {
        throw new Error("Missing the required parameter 'order' when calling placeOrder");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
        'User-Agent': 'OpenAPI-Generator/1.0.0/Javascript',
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/xml', 'application/json'];
      let returnType = Order;

      return this.callApi(
        '/store/order', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, requestInit
      );
    }


}
