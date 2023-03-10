/*
 * test
 * test
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonSubTypes;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import com.fasterxml.jackson.annotation.JsonTypeName;
import com.fasterxml.jackson.annotation.JsonValue;
import org.openapitools.client.JSON;
import org.openapitools.client.model.ChildSchemaAllOf;
import org.openapitools.client.model.Parent;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

/**
 * Model tests for ChildSchema
 */
public class ChildSchemaTest {
    private final ChildSchema model = new ChildSchema();

    /**
     * Model tests for ChildSchema
     */
    @Test
    public void testChildSchema() {
        String objJson = "{ \"objectType\": \"ChildSchema\", \"prop1\":\"some_value\" }";
        try {
            JSON j = new JSON();
            ChildSchema obj = j.getMapper().readValue(objJson, ChildSchema.class);
            Assertions.assertEquals(obj.getObjectType(), "ChildSchema");
            Assertions.assertEquals(obj.getProp1(), "some_value");
        } catch (Exception ex) {
            Assertions.fail("Exception '" + ex.getMessage() + "' should not have been raised");
        }
    }

    /**
     * Test the property 'objectType'
     */
    @Test
    public void objectTypeTest() {
        // TODO: test objectType
    }

    /**
     * Test the property 'prop1'
     */
    @Test
    public void prop1Test() {
        // TODO: test prop1
    }

}
