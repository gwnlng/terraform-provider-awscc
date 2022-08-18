// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package msk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_msk_serverless_cluster", serverlessClusterDataSourceType)
}

// serverlessClusterDataSourceType returns the Terraform awscc_msk_serverless_cluster data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::MSK::ServerlessCluster resource type.
func serverlessClusterDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"client_authentication": {
			// Property: ClientAuthentication
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Sasl": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Iam": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Enabled": {
			//               "type": "boolean"
			//             }
			//           },
			//           "required": [
			//             "Enabled"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "Iam"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Sasl"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sasl": {
						// Property: Sasl
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"iam": {
									// Property: Iam
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"vpc_configs": {
			// Property: VpcConfigs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "SecurityGroups": {
			//         "insertionOrder": false,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "SubnetIds": {
			//         "insertionOrder": false,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "SubnetIds"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"security_groups": {
						// Property: SecurityGroups
						Type:     types.SetType{ElemType: types.StringType},
						Computed: true,
					},
					"subnet_ids": {
						// Property: SubnetIds
						Type:     types.SetType{ElemType: types.StringType},
						Computed: true,
					},
				},
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
		Description: "Data Source schema for AWS::MSK::ServerlessCluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::ServerlessCluster").WithTerraformTypeName("awscc_msk_serverless_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"client_authentication": "ClientAuthentication",
		"cluster_name":          "ClusterName",
		"enabled":               "Enabled",
		"iam":                   "Iam",
		"sasl":                  "Sasl",
		"security_groups":       "SecurityGroups",
		"subnet_ids":            "SubnetIds",
		"tags":                  "Tags",
		"vpc_configs":           "VpcConfigs",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
