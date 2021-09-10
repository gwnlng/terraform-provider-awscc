// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wafv2

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
	registry.AddDataSourceTypeFactory("awscc_wafv2_regex_pattern_set", regexPatternSetDataSourceType)
}

// regexPatternSetDataSourceType returns the Terraform awscc_wafv2_regex_pattern_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::WAFv2::RegexPatternSet resource type.
func regexPatternSetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the WAF entity.",
			//   "type": "string"
			// }
			Description: "ARN of the WAF entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the entity.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Description of the entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the RegexPatternSet",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Id of the RegexPatternSet",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the RegexPatternSet.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the RegexPatternSet.",
			Type:        types.StringType,
			Computed:    true,
		},
		"regular_expression_list": {
			// Property: RegularExpressionList
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
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			// {
			//   "description": "Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.",
			//   "enum": [
			//     "CLOUDFRONT",
			//     "REGIONAL"
			//   ],
			//   "type": "string"
			// }
			Description: "Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
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
		Description: "Data Source schema for AWS::WAFv2::RegexPatternSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::RegexPatternSet").WithTerraformTypeName("awscc_wafv2_regex_pattern_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"id":                      "Id",
		"key":                     "Key",
		"name":                    "Name",
		"regular_expression_list": "RegularExpressionList",
		"scope":                   "Scope",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_wafv2_regex_pattern_set", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
