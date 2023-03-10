/* 
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using Org.OpenAPITools.Api;
using Org.OpenAPITools.Model;
using Org.OpenAPITools.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace Org.OpenAPITools.Test
{
    /// <summary>
    ///  Class for testing Pet
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class PetTests
    {
        //private Pet instance;

        private long petId = 11088;

        /// <summary>
        /// Test an instance of Pet
        /// </summary>
        [Fact]
        public void PetInstanceTest()
        {
            // TODO uncomment below to test "IsInstanceOfType" Pet
            //Assert.IsType<Pet> (instance, "variable 'instance' is a Pet");
        }


        /// <summary>
        /// Test the property 'Id'
        /// </summary>
        [Fact]
        public void IdTest()
        {
            // TODO unit test for the property 'Id'
        }
        /// <summary>
        /// Test the property 'Category'
        /// </summary>
        [Fact]
        public void CategoryTest()
        {
            // TODO unit test for the property 'Category'
        }
        /// <summary>
        /// Test the property 'Name'
        /// </summary>
        [Fact]
        public void NameTest()
        {
            // TODO unit test for the property 'Name'
        }
        /// <summary>
        /// Test the property 'PhotoUrls'
        /// </summary>
        [Fact]
        public void PhotoUrlsTest()
        {
            // TODO unit test for the property 'PhotoUrls'
        }
        /// <summary>
        /// Test the property 'Tags'
        /// </summary>
        [Fact]
        public void TagsTest()
        {
            // TODO unit test for the property 'Tags'
        }
        /// <summary>
        /// Test the property 'Status'
        /// </summary>
        [Fact]
        public void StatusTest()
        {
            // TODO unit test for the property 'Status'
        }

        /// <summary>
        /// Test serialization
        /// </summary>
        [Fact]
        public void TestSerialization()
        {
            // create pet
            Pet p1 = new Pet(name: "Csharp test", photoUrls: new List<string> { "http://petstore.com/csharp_test" });
            Assert.Equal("{\"name\":\"Csharp test\",\"photoUrls\":[\"http://petstore.com/csharp_test\"]}", JsonConvert.SerializeObject(p1));

            // test additional properties (serialization)
            Pet p2 = new Pet(name: "Csharp test", photoUrls: new List<string> { "http://petstore.com/csharp_test" });
            p2.AdditionalProperties.Add("hello", "world");
            Assert.Equal("{\"name\":\"Csharp test\",\"photoUrls\":[\"http://petstore.com/csharp_test\"],\"hello\":\"world\"}", JsonConvert.SerializeObject(p2));

            // test additional properties (deserialization)
            Pet p3 = JsonConvert.DeserializeObject<Pet>("{\"name\":\"Csharp test\",\"photoUrls\":[\"http://petstore.com/csharp_test\"],\"hello\":\"world\",\"int\":123}");
            Assert.Equal("Csharp test", p3.Name);
            Assert.Equal("world", p3.AdditionalProperties["hello"]);
            Assert.Equal(123L, p3.AdditionalProperties["int"]);
        }

        /// <summary>
        /// Test Equal
        /// </summary>
        [Fact]
        public void TestEqual()
        {
            // create pet
            Pet p1 = new Pet(name: "Csharp test", photoUrls: new List<string> { "http://petstore.com/csharp_test" });
            p1.Id = petId;
            //p1.Name = "Csharp test";
            p1.Status = Pet.StatusEnum.Available;
            // create Category object
            Category category1 = new Category();
            category1.Id = 56;
            category1.Name = "sample category name2";
            List<String> photoUrls1 = new List<String>(new String[] { "sample photoUrls" });
            // create Tag object
            Tag tag1 = new Tag();
            tag1.Id = petId;
            tag1.Name = "csharp sample tag name1";
            List<Tag> tags1 = new List<Tag>(new Tag[] { tag1 });
            p1.Tags = tags1;
            p1.Category = category1;
            p1.PhotoUrls = photoUrls1;

            // create pet 2
            Pet p2 = new Pet(name: "Csharp test", photoUrls: new List<string> { "http://petstore.com/csharp_test" });
            p2.Id = petId;
            p2.Name = "Csharp test";
            p2.Status = Pet.StatusEnum.Available;
            // create Category object
            Category category2 = new Category();
            category2.Id = 56;
            category2.Name = "sample category name2";
            List<String> photoUrls2 = new List<String>(new String[] { "sample photoUrls" });
            // create Tag object
            Tag tag2 = new Tag();
            tag2.Id = petId;
            tag2.Name = "csharp sample tag name1";
            List<Tag> tags2 = new List<Tag>(new Tag[] { tag2 });
            p2.Tags = tags2;
            p2.Category = category2;
            p2.PhotoUrls = photoUrls2;

            // p1 and p2 should be equal (both object and attribute level)
            Assert.True(category1.Equals(category2));
            Assert.True(tags1.SequenceEqual(tags2));
            Assert.True(p1.PhotoUrls.SequenceEqual(p2.PhotoUrls));

            Assert.True(p1.Equals(p2));

            // update attribute to that p1 and p2 are not equal
            category2.Name = "new category name";
            Assert.False(category1.Equals(category2));

            tags2 = new List<Tag>();
            Assert.False(tags1.SequenceEqual(tags2));

            // photoUrls has not changed so it should be equal
            Assert.True(p1.PhotoUrls.SequenceEqual(p2.PhotoUrls));

            Assert.False(p1.Equals(p2));
        }


        /// <summary>
        /// Test Pet deserialization
        /// </summary>
        [Fact]
        public void TestDeserialization()
        {
            // properly deserialized, no exception
            Pet p2 = JsonConvert.DeserializeObject<Pet>("{\n  \"name\": \"csharp test 2\",\n  \"photoUrls\": [\"http://petstore.com/csharp_test\", \"http://petstore.com/csharp_test2\"]\n}");
            // comment out below as the result (json string) is OS dependent.
            //Assert.Equal("{\n  \"name\": \"csharp test 2\",\n  \"photoUrls\": [\n    \"http://petstore.com/csharp_test\",\n    \"http://petstore.com/csharp_test2\"\n  ]\n}", p2.ToJson());

            // Missing `name` to cause exceptions in deserialization
            Assert.Throws<Newtonsoft.Json.JsonSerializationException>(() => JsonConvert.DeserializeObject<Pet>("{\n  \"Something\": \"csharp test 2\",\n  \"photoUrls\": [\"http://petstore.com/csharp_test\", \"http://petstore.com/csharp_test2\"]\n}"));
        }

    }

}
