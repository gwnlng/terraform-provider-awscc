// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kendra_index", indexDataSourceType)
}

// indexDataSourceType returns the Terraform awscc_kendra_index data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Kendra::Index resource type.
func indexDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"capacity_units": {
			// Property: CapacityUnits
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "QueryCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "StorageCapacityUnits": {
			//       "minimum": 0,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "StorageCapacityUnits",
			//     "QueryCapacityUnits"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"query_capacity_units": {
						// Property: QueryCapacityUnits
						Type:     types.NumberType,
						Computed: true,
					},
					"storage_capacity_units": {
						// Property: StorageCapacityUnits
						Type:     types.NumberType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"document_metadata_configurations": {
			// Property: DocumentMetadataConfigurations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 30,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Relevance": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Duration": {
			//             "maxLength": 10,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Freshness": {
			//             "type": "boolean"
			//           },
			//           "Importance": {
			//             "maximum": 10,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "RankOrder": {
			//             "enum": [
			//               "ASCENDING",
			//               "DESCENDING"
			//             ],
			//             "type": "string"
			//           },
			//           "ValueImportanceItems": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 50,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "maximum": 10,
			//                   "minimum": 1,
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Search": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Displayable": {
			//             "type": "boolean"
			//           },
			//           "Facetable": {
			//             "type": "boolean"
			//           },
			//           "Searchable": {
			//             "type": "boolean"
			//           },
			//           "Sortable": {
			//             "type": "boolean"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Type": {
			//         "enum": [
			//           "STRING_VALUE",
			//           "STRING_LIST_VALUE",
			//           "LONG_VALUE",
			//           "DATE_VALUE"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 500,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
					},
					"relevance": {
						// Property: Relevance
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"duration": {
									// Property: Duration
									Type:     types.StringType,
									Computed: true,
								},
								"freshness": {
									// Property: Freshness
									Type:     types.BoolType,
									Computed: true,
								},
								"importance": {
									// Property: Importance
									Type:     types.NumberType,
									Computed: true,
								},
								"rank_order": {
									// Property: RankOrder
									Type:     types.StringType,
									Computed: true,
								},
								"value_importance_items": {
									// Property: ValueImportanceItems
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
											"value": {
												// Property: Value
												Type:     types.NumberType,
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"search": {
						// Property: Search
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"displayable": {
									// Property: Displayable
									Type:     types.BoolType,
									Computed: true,
								},
								"facetable": {
									// Property: Facetable
									Type:     types.BoolType,
									Computed: true,
								},
								"searchable": {
									// Property: Searchable
									Type:     types.BoolType,
									Computed: true,
								},
								"sortable": {
									// Property: Sortable
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"edition": {
			// Property: Edition
			// CloudFormation resource type schema:
			// {
			//   "description": "Edition of index",
			//   "enum": [
			//     "DEVELOPER_EDITION",
			//     "ENTERPRISE_EDITION"
			//   ],
			//   "type": "string"
			// }
			Description: "Edition of index",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of index",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "type": "string"
			// }
			Description: "Unique ID of index",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of index",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of index",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 1284,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"server_side_encryption_configuration": {
			// Property: ServerSideEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KmsKeyId": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of tags",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Kendra resources",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "List of tags",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"user_context_policy": {
			// Property: UserContextPolicy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ATTRIBUTE_FILTER",
			//     "USER_TOKEN"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"user_token_configurations": {
			// Property: UserTokenConfigurations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "JsonTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "UserNameAttributeField",
			//           "GroupAttributeField"
			//         ],
			//         "type": "object"
			//       },
			//       "JwtTokenTypeConfiguration": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ClaimRegex": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "GroupAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Issuer": {
			//             "maxLength": 65,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "KeyLocation": {
			//             "enum": [
			//               "URL",
			//               "SECRET_MANAGER"
			//             ],
			//             "type": "string"
			//           },
			//           "SecretManagerArn": {
			//             "description": "Role Arn",
			//             "maxLength": 1284,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "URL": {
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "UserNameAttributeField": {
			//             "maxLength": 100,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "KeyLocation"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"json_token_type_configuration": {
						// Property: JsonTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type:     types.StringType,
									Computed: true,
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"jwt_token_type_configuration": {
						// Property: JwtTokenTypeConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"claim_regex": {
									// Property: ClaimRegex
									Type:     types.StringType,
									Computed: true,
								},
								"group_attribute_field": {
									// Property: GroupAttributeField
									Type:     types.StringType,
									Computed: true,
								},
								"issuer": {
									// Property: Issuer
									Type:     types.StringType,
									Computed: true,
								},
								"key_location": {
									// Property: KeyLocation
									Type:     types.StringType,
									Computed: true,
								},
								"secret_manager_arn": {
									// Property: SecretManagerArn
									Description: "Role Arn",
									Type:        types.StringType,
									Computed:    true,
								},
								"url": {
									// Property: URL
									Type:     types.StringType,
									Computed: true,
								},
								"user_name_attribute_field": {
									// Property: UserNameAttributeField
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Kendra::Index",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Index").WithTerraformTypeName("awscc_kendra_index")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"capacity_units":                       "CapacityUnits",
		"claim_regex":                          "ClaimRegex",
		"description":                          "Description",
		"displayable":                          "Displayable",
		"document_metadata_configurations":     "DocumentMetadataConfigurations",
		"duration":                             "Duration",
		"edition":                              "Edition",
		"facetable":                            "Facetable",
		"freshness":                            "Freshness",
		"group_attribute_field":                "GroupAttributeField",
		"id":                                   "Id",
		"importance":                           "Importance",
		"issuer":                               "Issuer",
		"json_token_type_configuration":        "JsonTokenTypeConfiguration",
		"jwt_token_type_configuration":         "JwtTokenTypeConfiguration",
		"key":                                  "Key",
		"key_location":                         "KeyLocation",
		"kms_key_id":                           "KmsKeyId",
		"name":                                 "Name",
		"query_capacity_units":                 "QueryCapacityUnits",
		"rank_order":                           "RankOrder",
		"relevance":                            "Relevance",
		"role_arn":                             "RoleArn",
		"search":                               "Search",
		"searchable":                           "Searchable",
		"secret_manager_arn":                   "SecretManagerArn",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"sortable":                             "Sortable",
		"storage_capacity_units":               "StorageCapacityUnits",
		"tags":                                 "Tags",
		"type":                                 "Type",
		"url":                                  "URL",
		"user_context_policy":                  "UserContextPolicy",
		"user_name_attribute_field":            "UserNameAttributeField",
		"user_token_configurations":            "UserTokenConfigurations",
		"value":                                "Value",
		"value_importance_items":               "ValueImportanceItems",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_kendra_index", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
