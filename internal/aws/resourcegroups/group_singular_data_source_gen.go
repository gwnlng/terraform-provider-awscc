// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package resourcegroups

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
	registry.AddDataSourceTypeFactory("awscc_resourcegroups_group", groupDataSourceType)
}

// groupDataSourceType returns the Terraform awscc_resourcegroups_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ResourceGroups::Group resource type.
func groupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Resource Group ARN.",
			//   "type": "string"
			// }
			Description: "The Resource Group ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "properties": {
			//       "Parameters": {
			//         "items": {
			//           "properties": {
			//             "Name": {
			//               "type": "string"
			//             },
			//             "Values": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "Type": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameters": {
						// Property: Parameters
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"values": {
									// Property: Values
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
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
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the resource group",
			//   "maxLength": 512,
			//   "type": "string"
			// }
			Description: "The description of the resource group",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the resource group",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "The name of the resource group",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_query": {
			// Property: ResourceQuery
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "Query": {
			//       "properties": {
			//         "ResourceTypeFilters": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "StackIdentifier": {
			//           "type": "string"
			//         },
			//         "TagFilters": {
			//           "items": {
			//             "properties": {
			//               "Key": {
			//                 "type": "string"
			//               },
			//               "Values": {
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Type": {
			//       "enum": [
			//         "TAG_FILTERS_1_0",
			//         "CLOUDFORMATION_STACK_1_0"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"query": {
						// Property: Query
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_type_filters": {
									// Property: ResourceTypeFilters
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"stack_identifier": {
									// Property: StackIdentifier
									Type:     types.StringType,
									Computed: true,
								},
								"tag_filters": {
									// Property: TagFilters
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Computed: true,
											},
											"values": {
												// Property: Values
												Type:     types.ListType{ElemType: types.StringType},
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
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"resources": {
			// Property: Resources
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
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
		Description: "Data Source schema for AWS::ResourceGroups::Group",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResourceGroups::Group").WithTerraformTypeName("awscc_resourcegroups_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"configuration":         "Configuration",
		"description":           "Description",
		"key":                   "Key",
		"name":                  "Name",
		"parameters":            "Parameters",
		"query":                 "Query",
		"resource_query":        "ResourceQuery",
		"resource_type_filters": "ResourceTypeFilters",
		"resources":             "Resources",
		"stack_identifier":      "StackIdentifier",
		"tag_filters":           "TagFilters",
		"tags":                  "Tags",
		"type":                  "Type",
		"value":                 "Value",
		"values":                "Values",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_resourcegroups_group", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
