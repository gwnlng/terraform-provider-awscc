// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssm

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
	registry.AddDataSourceTypeFactory("awscc_ssm_resource_data_sync", resourceDataSyncDataSourceType)
}

// resourceDataSyncDataSourceType returns the Terraform awscc_ssm_resource_data_sync data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SSM::ResourceDataSync resource type.
func resourceDataSyncDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"bucket_prefix": {
			// Property: BucketPrefix
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"bucket_region": {
			// Property: BucketRegion
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"kms_key_arn": {
			// Property: KMSKeyArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"s3_destination": {
			// Property: S3Destination
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BucketName": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketPrefix": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketRegion": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "KMSKeyArn": {
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SyncFormat": {
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "BucketName",
			//     "BucketRegion",
			//     "SyncFormat"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket_name": {
						// Property: BucketName
						Type:     types.StringType,
						Computed: true,
					},
					"bucket_prefix": {
						// Property: BucketPrefix
						Type:     types.StringType,
						Computed: true,
					},
					"bucket_region": {
						// Property: BucketRegion
						Type:     types.StringType,
						Computed: true,
					},
					"kms_key_arn": {
						// Property: KMSKeyArn
						Type:     types.StringType,
						Computed: true,
					},
					"sync_format": {
						// Property: SyncFormat
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"sync_format": {
			// Property: SyncFormat
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sync_name": {
			// Property: SyncName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sync_source": {
			// Property: SyncSource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AwsOrganizationsSource": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "OrganizationSourceType": {
			//           "maxLength": 64,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "OrganizationalUnits": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "OrganizationSourceType"
			//       ],
			//       "type": "object"
			//     },
			//     "IncludeFutureRegions": {
			//       "type": "boolean"
			//     },
			//     "SourceRegions": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SourceType": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SourceType",
			//     "SourceRegions"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"aws_organizations_source": {
						// Property: AwsOrganizationsSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"organization_source_type": {
									// Property: OrganizationSourceType
									Type:     types.StringType,
									Computed: true,
								},
								"organizational_units": {
									// Property: OrganizationalUnits
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"include_future_regions": {
						// Property: IncludeFutureRegions
						Type:     types.BoolType,
						Computed: true,
					},
					"source_regions": {
						// Property: SourceRegions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"source_type": {
						// Property: SourceType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"sync_type": {
			// Property: SyncType
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SSM::ResourceDataSync",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourceDataSync").WithTerraformTypeName("awscc_ssm_resource_data_sync")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_organizations_source": "AwsOrganizationsSource",
		"bucket_name":              "BucketName",
		"bucket_prefix":            "BucketPrefix",
		"bucket_region":            "BucketRegion",
		"include_future_regions":   "IncludeFutureRegions",
		"kms_key_arn":              "KMSKeyArn",
		"organization_source_type": "OrganizationSourceType",
		"organizational_units":     "OrganizationalUnits",
		"s3_destination":           "S3Destination",
		"source_regions":           "SourceRegions",
		"source_type":              "SourceType",
		"sync_format":              "SyncFormat",
		"sync_name":                "SyncName",
		"sync_source":              "SyncSource",
		"sync_type":                "SyncType",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ssm_resource_data_sync", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
