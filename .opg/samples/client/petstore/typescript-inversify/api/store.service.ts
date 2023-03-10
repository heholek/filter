/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/* tslint:disable:no-unused-variable member-ordering */

import { Observable } from 'rxjs/Observable';

import { map } from 'rxjs/operators';
import IHttpClient from '../IHttpClient';
import { inject, injectable } from 'inversify';
import { IAPIConfiguration } from '../IAPIConfiguration';
import { Headers } from '../Headers';
import HttpResponse from '../HttpResponse';

import { Order } from '../model/order';

import { COLLECTION_FORMATS }  from '../variables';



@injectable()
export class StoreService {
    private basePath: string = 'http://petstore.swagger.io/v2';

    constructor(@inject('IApiHttpClient') private httpClient: IHttpClient,
        @inject('IAPIConfiguration') private APIConfiguration: IAPIConfiguration ) {
        if(this.APIConfiguration.basePath)
            this.basePath = this.APIConfiguration.basePath;
    }

    /**
     * Delete purchase order by ID
     * For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors
     * @param orderId ID of the order that needs to be deleted
     
     */
    public deleteOrder(orderId: string, observe?: 'body', headers?: Headers): Observable<any>;
    public deleteOrder(orderId: string, observe?: 'response', headers?: Headers): Observable<HttpResponse<any>>;
    public deleteOrder(orderId: string, observe: any = 'body', headers: Headers = {}): Observable<any> {
        if (orderId === null || orderId === undefined){
            throw new Error('Required parameter orderId was null or undefined when calling deleteOrder.');
        }

        headers['Accept'] = 'application/json';

        const response: Observable<HttpResponse<any>> = this.httpClient.delete(`${this.basePath}/store/order/${encodeURIComponent(String(orderId))}`, headers);
        if (observe === 'body') {
               return response.pipe(
                   map((httpResponse: HttpResponse) => <any>(httpResponse.response))
               );
        }
        return response;
    }


    /**
     * Returns pet inventories by status
     * Returns a map of status codes to quantities
     
     */
    public getInventory(observe?: 'body', headers?: Headers): Observable<{ [key: string]: number; }>;
    public getInventory(observe?: 'response', headers?: Headers): Observable<HttpResponse<{ [key: string]: number; }>>;
    public getInventory(observe: any = 'body', headers: Headers = {}): Observable<any> {
        // authentication (api_key) required
        if (this.APIConfiguration.apiKeys && this.APIConfiguration.apiKeys['api_key']) {
            headers['api_key'] = this.APIConfiguration.apiKeys['api_key'];
        }
        headers['Accept'] = 'application/json';

        const response: Observable<HttpResponse<{ [key: string]: number; }>> = this.httpClient.get(`${this.basePath}/store/inventory`, headers);
        if (observe === 'body') {
               return response.pipe(
                   map((httpResponse: HttpResponse) => <{ [key: string]: number; }>(httpResponse.response))
               );
        }
        return response;
    }


    /**
     * Find purchase order by ID
     * For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generate exceptions
     * @param orderId ID of pet that needs to be fetched
     
     */
    public getOrderById(orderId: number, observe?: 'body', headers?: Headers): Observable<Order>;
    public getOrderById(orderId: number, observe?: 'response', headers?: Headers): Observable<HttpResponse<Order>>;
    public getOrderById(orderId: number, observe: any = 'body', headers: Headers = {}): Observable<any> {
        if (orderId === null || orderId === undefined){
            throw new Error('Required parameter orderId was null or undefined when calling getOrderById.');
        }

        headers['Accept'] = 'application/xml, application/json';

        const response: Observable<HttpResponse<Order>> = this.httpClient.get(`${this.basePath}/store/order/${encodeURIComponent(String(orderId))}`, headers);
        if (observe === 'body') {
               return response.pipe(
                   map((httpResponse: HttpResponse) => <Order>(httpResponse.response))
               );
        }
        return response;
    }


    /**
     * Place an order for a pet
     * 
     * @param body order placed for purchasing the pet
     
     */
    public placeOrder(body: Order, observe?: 'body', headers?: Headers): Observable<Order>;
    public placeOrder(body: Order, observe?: 'response', headers?: Headers): Observable<HttpResponse<Order>>;
    public placeOrder(body: Order, observe: any = 'body', headers: Headers = {}): Observable<any> {
        if (body === null || body === undefined){
            throw new Error('Required parameter body was null or undefined when calling placeOrder.');
        }

        headers['Accept'] = 'application/xml, application/json';
        headers['Content-Type'] = 'application/json';

        const response: Observable<HttpResponse<Order>> = this.httpClient.post(`${this.basePath}/store/order`, body , headers);
        if (observe === 'body') {
               return response.pipe(
                   map((httpResponse: HttpResponse) => <Order>(httpResponse.response))
               );
        }
        return response;
    }

}
