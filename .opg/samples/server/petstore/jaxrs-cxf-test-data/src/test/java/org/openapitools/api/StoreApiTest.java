/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package org.openapitools.api;

import java.util.List;
import java.util.Map;
import org.openapitools.model.Order;
import org.junit.Test;
import org.junit.Before;
import static org.junit.Assert.*;

import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import org.apache.cxf.jaxrs.client.JAXRSClientFactory;
import org.apache.cxf.jaxrs.client.ClientConfiguration;
import org.apache.cxf.jaxrs.client.WebClient;
import org.apache.cxf.jaxrs.ext.multipart.Attachment;


import com.fasterxml.jackson.jaxrs.json.JacksonJsonProvider;
import com.fasterxml.jackson.jaxrs.xml.JacksonXMLProvider;
import org.apache.cxf.jaxrs.provider.MultipartProvider;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.junit.runner.RunWith;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.web.server.LocalServerPort;
import java.io.File;
import java.util.Set;
import javax.validation.ConstraintViolation;
import javax.validation.Validation;
import javax.validation.Validator;
import javax.validation.ValidatorFactory;
import com.fasterxml.jackson.databind.node.JsonNodeFactory;
import com.fasterxml.jackson.databind.node.ObjectNode;
import org.junit.BeforeClass;
import org.openapitools.codegen.utils.JsonCache;

/**
 * OpenAPI Petstore
 *
 * <p>This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API tests for StoreApi.
 */
@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
public class StoreApiTest {
    private static Validator validator;
    private static JsonCache cache;

    @BeforeClass
    public static void beforeClass() throws Exception {
        File cacheFile = new File(System.getProperty("jaxrs.test.client.json",
                "C:\\source\\GitHub\\demonfiddler\\openapi-generator\\samples\\server\\petstore\\jaxrs-cxf-test-data\\src\\main\\resources\\test-data.json"));
        cache = JsonCache.Factory.instance.get("test-data").load(cacheFile).child("/org.openapitools.api/StoreApi");

        validator = Validation.buildDefaultValidatorFactory().getValidator();
    }

    @LocalServerPort
    private int serverPort;

    private StoreApi api;

    @Before
    public void setup() {
        List<?> providers = Arrays.asList(new JacksonJsonProvider(), new JacksonXMLProvider(), new MultipartProvider());

        api = JAXRSClientFactory.create("http://localhost:" + serverPort + "/services", StoreApi.class, providers);
        org.apache.cxf.jaxrs.client.Client client = WebClient.client(api);

        ClientConfiguration config = WebClient.getConfig(client); 
    }

    private void validate(Object o) {
        assertNotNull(o);
        Set<ConstraintViolation<Object>> violations = validator.validate(o);
        if (!violations.isEmpty()) {
            StringBuilder message = new StringBuilder("Validation failed");
            for (ConstraintViolation<Object> violation : violations)
                message.append("; ").append(violation.getPropertyPath()).append(": ").append(violation.getMessage());
            fail(message.toString());
        }
    }

    
    /**
     * Delete purchase order by ID
     *
     * For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors
     *
     * @throws ApiException if the API call fails
     */
    @Test
    public void deleteOrderTest() throws Exception {
        String orderId = cache.getString("/deleteOrder/orderId");
        api.deleteOrder(orderId);
    }
    
    /**
     * Returns pet inventories by status
     *
     * Returns a map of status codes to quantities
     *
     * @throws ApiException if the API call fails
     */
    @Test
    public void getInventoryTest() throws Exception {

        Map<String, Integer> response = api.getInventory();
        // TODO: complete test assertions
        validate(response);
    }
    
    /**
     * Find purchase order by ID
     *
     * For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generate exceptions
     *
     * @throws ApiException if the API call fails
     */
    @Test
    public void getOrderByIdTest() throws Exception {
        Long orderId = cache.getLong("/getOrderById/orderId");
        Order response = api.getOrderById(orderId);
        // TODO: complete test assertions
        validate(response);
    }
    
    /**
     * Place an order for a pet
     *
     * @throws ApiException if the API call fails
     */
    @Test
    public void placeOrderTest() throws Exception {
        Order body = cache.getObject("/placeOrder/body", Order.class);
        Order response = api.placeOrder(body);
        // TODO: complete test assertions
        validate(response);
    }
    
}
