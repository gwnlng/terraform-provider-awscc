// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticbeanstalk_application", applicationDataSource)
}

// applicationDataSource returns the Terraform awscc_elasticbeanstalk_application data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElasticBeanstalk::Application resource.
func applicationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
		//	  "type": "string"
		//	}
		"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your description of the application.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your description of the application.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceLifecycleConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.",
		//	  "properties": {
		//	    "ServiceRole": {
		//	      "description": "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.",
		//	      "type": "string"
		//	    },
		//	    "VersionLifecycleConfig": {
		//	      "additionalProperties": false,
		//	      "description": "Defines lifecycle settings for application versions.",
		//	      "properties": {
		//	        "MaxAgeRule": {
		//	          "additionalProperties": false,
		//	          "description": "Specify a max age rule to restrict the length of time that application versions are retained for an application.",
		//	          "properties": {
		//	            "DeleteSourceFromS3": {
		//	              "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
		//	              "type": "boolean"
		//	            },
		//	            "Enabled": {
		//	              "description": "Specify true to apply the rule, or false to disable it.",
		//	              "type": "boolean"
		//	            },
		//	            "MaxAgeInDays": {
		//	              "description": "Specify the number of days to retain an application versions.",
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "MaxCountRule": {
		//	          "additionalProperties": false,
		//	          "description": "Specify a max count rule to restrict the number of application versions that are retained for an application.",
		//	          "properties": {
		//	            "DeleteSourceFromS3": {
		//	              "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
		//	              "type": "boolean"
		//	            },
		//	            "Enabled": {
		//	              "description": "Specify true to apply the rule, or false to disable it.",
		//	              "type": "boolean"
		//	            },
		//	            "MaxCount": {
		//	              "description": "Specify the maximum number of application versions to retain.",
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource_lifecycle_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ServiceRole
				"service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VersionLifecycleConfig
				"version_lifecycle_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MaxAgeRule
						"max_age_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DeleteSourceFromS3
								"delete_source_from_s3": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specify true to apply the rule, or false to disable it.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: MaxAgeInDays
								"max_age_in_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Specify the number of days to retain an application versions.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Specify a max age rule to restrict the length of time that application versions are retained for an application.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MaxCountRule
						"max_count_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DeleteSourceFromS3
								"delete_source_from_s3": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Specify true to apply the rule, or false to disable it.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: MaxCount
								"max_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Specify the maximum number of application versions to retain.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Specify a max count rule to restrict the number of application versions that are retained for an application.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Defines lifecycle settings for application versions.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ElasticBeanstalk::Application",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::Application").WithTerraformTypeName("awscc_elasticbeanstalk_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":          "ApplicationName",
		"delete_source_from_s3":     "DeleteSourceFromS3",
		"description":               "Description",
		"enabled":                   "Enabled",
		"max_age_in_days":           "MaxAgeInDays",
		"max_age_rule":              "MaxAgeRule",
		"max_count":                 "MaxCount",
		"max_count_rule":            "MaxCountRule",
		"resource_lifecycle_config": "ResourceLifecycleConfig",
		"service_role":              "ServiceRole",
		"version_lifecycle_config":  "VersionLifecycleConfig",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
