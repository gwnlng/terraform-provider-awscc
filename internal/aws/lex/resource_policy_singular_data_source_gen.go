// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lex_resource_policy", resourcePolicyDataSourceType)
}

// resourcePolicyDataSourceType returns the Terraform awscc_lex_resource_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lex::ResourcePolicy resource type.
func resourcePolicyDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The Physical ID of the resource policy.",
			//   "type": "string"
			// }
			Description: "The Physical ID of the resource policy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
			//   "type": "object"
			// }
			Description: "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
			//   "maxLength": 1011,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
			Type:        types.StringType,
			Computed:    true,
		},
		"revision_id": {
			// Property: RevisionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
			//   "maxLength": 5,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lex::ResourcePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::ResourcePolicy").WithTerraformTypeName("awscc_lex_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":           "Id",
		"policy":       "Policy",
		"resource_arn": "ResourceArn",
		"revision_id":  "RevisionId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
