# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import petstore_api
from petstore_api.schemas import NoneClass
from petstore_api.model import shape
from petstore_api.model import shape_or_null
from petstore_api.model.drawing import Drawing


class TestDrawing(unittest.TestCase):
    """Drawing unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test_create_instances(self):
        """
        Validate instance can be created
        """

        inst = shape.Shape(
            shapeType="Triangle",
            triangleType="IsoscelesTriangle"
        )
        from petstore_api.model.isosceles_triangle import IsoscelesTriangle
        assert isinstance(inst, IsoscelesTriangle)

    def test_deserialize_oneof_reference(self):
        """
        Validate the scenario when the type of a OAS property is 'oneOf', and the 'oneOf'
        schema is specified as a reference ($ref), not an inline 'oneOf' schema.
        """
        isosceles_triangle = shape.Shape(
            shapeType="Triangle",
            triangleType="IsoscelesTriangle"
        )
        from petstore_api.model.isosceles_triangle import IsoscelesTriangle
        assert isinstance(isosceles_triangle, IsoscelesTriangle)
        from petstore_api.model.equilateral_triangle import EquilateralTriangle

        inst = Drawing(
            mainShape=isosceles_triangle,
            shapes=[
                shape.Shape(
                    shapeType="Triangle",
                    triangleType="EquilateralTriangle"
                ),
                shape.Shape(
                    shapeType="Triangle",
                    triangleType="IsoscelesTriangle"
                ),
                shape.Shape(
                    shapeType="Triangle",
                    triangleType="EquilateralTriangle"
                ),
                shape.Shape(
                    shapeType="Quadrilateral",
                    quadrilateralType="ComplexQuadrilateral"
                )
            ],
        )
        assert isinstance(inst, Drawing)
        assert isinstance(inst["mainShape"], IsoscelesTriangle)
        self.assertEqual(len(inst["shapes"]), 4)
        from petstore_api.model.complex_quadrilateral import ComplexQuadrilateral
        assert isinstance(inst["shapes"][0], EquilateralTriangle)
        assert isinstance(inst["shapes"][1], IsoscelesTriangle)
        assert isinstance(inst["shapes"][2], EquilateralTriangle)
        assert isinstance(inst["shapes"][3], ComplexQuadrilateral)

        # Validate we cannot assign the None value to mainShape because the 'null' type
        # is not one of the allowed types in the 'Shape' schema.
        err_msg = (r"Invalid inputs given to generate an instance of .+?Shape.+? "
                   r"None of the oneOf schemas matched the input data.")
        with self.assertRaisesRegex(
                petstore_api.ApiValueError,
                err_msg
        ):
            Drawing(
                # 'mainShape' has type 'Shape', which is a oneOf [triangle, quadrilateral]
                # So the None value should not be allowed and an exception should be raised.
                mainShape=None,
            )

        """
        We can pass in a Triangle instance in shapes
        Under the hood it is converted into a dict, and that dict payload
        does validate as a Shape, so this works
        """
        from petstore_api.model.triangle import Triangle
        inst = Drawing(
            mainShape=isosceles_triangle,
            shapes=[
                Triangle(
                    shapeType="Triangle",
                    triangleType="EquilateralTriangle"
                )
            ]
        )
        self.assertEqual(len(inst["shapes"]), 1)
        from petstore_api.model.triangle_interface import TriangleInterface
        shapes = inst["shapes"]
        assert isinstance(shapes[0], shape.Shape)
        assert isinstance(shapes[0], Triangle)
        assert isinstance(shapes[0], EquilateralTriangle)
        assert isinstance(shapes[0], TriangleInterface)

    def test_deserialize_oneof_reference_with_null_type(self):
        """
        Validate the scenario when the type of a OAS property is 'oneOf', and the 'oneOf'
        schema is specified as a reference ($ref), not an inline 'oneOf' schema.
        Further, the 'oneOf' schema has a 'null' type child schema (as introduced in
        OpenAPI 3.1).
        """

        # Validate we can assign the None value to shape_or_null, because the 'null' type
        # is one of the allowed types in the 'ShapeOrNull' schema.
        inst = Drawing(
            # 'shapeOrNull' has type 'ShapeOrNull', which is a oneOf [null, triangle, quadrilateral]
            shapeOrNull=None,
        )
        assert isinstance(inst, Drawing)
        self.assertFalse('mainShape' in inst)
        self.assertTrue('shapeOrNull' in inst)
        self.assertTrue(isinstance(inst["shapeOrNull"], NoneClass))

    def test_deserialize_oneof_reference_with_nullable_type(self):
        """
        Validate the scenario when the type of a OAS property is 'oneOf', and the 'oneOf'
        schema is specified as a reference ($ref), not an inline 'oneOf' schema.
        Further, the 'oneOf' schema has the 'nullable' attribute (as introduced in
        OpenAPI 3.0 and deprecated in 3.1).
        """

        # Validate we can assign the None value to nullableShape, because the NullableShape
        # has the 'nullable: true' attribute.
        inst = Drawing(
            # 'nullableShape' has type 'NullableShape', which is a oneOf [triangle, quadrilateral]
            # and the 'nullable: true' attribute.
            nullableShape=None,
        )
        assert isinstance(inst, Drawing)
        self.assertFalse('mainShape' in inst)
        self.assertTrue('nullableShape' in inst)
        self.assertTrue(isinstance(inst["nullableShape"], NoneClass))


if __name__ == '__main__':
    unittest.main()