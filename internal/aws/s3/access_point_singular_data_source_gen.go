// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_access_point", accessPointDataSource)
}

// accessPointDataSource returns the Terraform awscc_s3_access_point data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::AccessPoint resource.
func accessPointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified accesspoint.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified accesspoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the bucket that you want to associate this Access Point with.",
		//	  "maxLength": 255,
		//	  "minLength": 3,
		//	  "type": "string"
		//	}
		"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the bucket that you want to associate this Access Point with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account ID associated with the S3 bucket associated with this access point.",
		//	  "maxLength": 64,
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"bucket_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account ID associated with the S3 bucket associated with this access point.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
		//	  "maxLength": 50,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkOrigin
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
		//	  "enum": [
		//	    "Internet",
		//	    "VPC"
		//	  ],
		//	  "type": "string"
		//	}
		"network_origin": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Access Point Policy you want to apply to this access point.",
		//	  "type": "object"
		//	}
		"policy": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Access Point Policy you want to apply to this access point.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "IsPublic": {
		//	      "description": "Specifies whether the policy is public or not.",
		//	      "enum": [
		//	        "true",
		//	        "false"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"policy_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IsPublic
				"is_public": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether the policy is public or not.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PublicAccessBlockConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
		//	  "properties": {
		//	    "BlockPublicAcls": {
		//	      "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
		//	      "type": "boolean"
		//	    },
		//	    "BlockPublicPolicy": {
		//	      "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
		//	      "type": "boolean"
		//	    },
		//	    "IgnorePublicAcls": {
		//	      "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
		//	      "type": "boolean"
		//	    },
		//	    "RestrictPublicBuckets": {
		//	      "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"public_access_block_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BlockPublicAcls
				"block_public_acls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: BlockPublicPolicy
				"block_public_policy": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IgnorePublicAcls
				"ignore_public_acls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RestrictPublicBuckets
				"restrict_public_buckets": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).",
		//	  "properties": {
		//	    "VpcId": {
		//	      "description": "If this field is specified, this access point will only allow connections from the specified VPC ID.",
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "If this field is specified, this access point will only allow connections from the specified VPC ID.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3::AccessPoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::AccessPoint").WithTerraformTypeName("awscc_s3_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                             "Alias",
		"arn":                               "Arn",
		"block_public_acls":                 "BlockPublicAcls",
		"block_public_policy":               "BlockPublicPolicy",
		"bucket":                            "Bucket",
		"bucket_account_id":                 "BucketAccountId",
		"ignore_public_acls":                "IgnorePublicAcls",
		"is_public":                         "IsPublic",
		"name":                              "Name",
		"network_origin":                    "NetworkOrigin",
		"policy":                            "Policy",
		"policy_status":                     "PolicyStatus",
		"public_access_block_configuration": "PublicAccessBlockConfiguration",
		"restrict_public_buckets":           "RestrictPublicBuckets",
		"vpc_configuration":                 "VpcConfiguration",
		"vpc_id":                            "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
