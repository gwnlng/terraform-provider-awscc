// Code generated by generators/resource/main.go; DO NOT EDIT.

package memorydb

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_memorydb_cluster", clusterResourceType)
}

// clusterResourceType returns the Terraform awscc_memorydb_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MemoryDB::Cluster resource type.
func clusterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"acl_name": {
			// Property: ACLName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Access Control List to associate with the cluster.",
			//   "pattern": "[a-zA-Z][a-zA-Z0-9\\-]*",
			//   "type": "string"
			// }
			Description: "The name of the Access Control List to associate with the cluster.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[a-zA-Z][a-zA-Z0-9\\-]*"), ""),
			},
		},
		"arn": {
			// Property: ARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the cluster.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"auto_minor_version_upgrade": {
			// Property: AutoMinorVersionUpgrade
			// CloudFormation resource type schema:
			// {
			//   "description": "A flag that enables automatic minor version upgrade when set to true.\n\nYou cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.",
			//   "type": "boolean"
			// }
			Description: "A flag that enables automatic minor version upgrade when set to true.\n\nYou cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"cluster_endpoint": {
			// Property: ClusterEndpoint
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The cluster endpoint.",
			//   "properties": {
			//     "Address": {
			//       "description": "The DNS address of the primary read-write node.",
			//       "type": "string"
			//     },
			//     "Port": {
			//       "description": "The port number that the engine is listening on. ",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The cluster endpoint.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"address": {
						// Property: Address
						Description: "The DNS address of the primary read-write node.",
						Type:        types.StringType,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.UseStateForUnknown(),
						},
					},
					"port": {
						// Property: Port
						Description: "The port number that the engine is listening on. ",
						Type:        types.Int64Type,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							tfsdk.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cluster. This value must be unique as it also serves as the cluster identifier.",
			//   "pattern": "[a-z][a-z0-9\\-]*",
			//   "type": "string"
			// }
			Description: "The name of the cluster. This value must be unique as it also serves as the cluster identifier.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("[a-z][a-z0-9\\-]*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional description of the cluster.",
			//   "type": "string"
			// }
			Description: "An optional description of the cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"engine_version": {
			// Property: EngineVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The Redis engine version used by the cluster.",
			//   "type": "string"
			// }
			Description: "The Redis engine version used by the cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"final_snapshot_name": {
			// Property: FinalSnapshotName
			// CloudFormation resource type schema:
			// {
			//   "description": "The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.",
			//   "type": "string"
			// }
			Description: "The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.",
			Type:        types.StringType,
			Optional:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the KMS key used to encrypt the cluster.",
			//   "type": "string"
			// }
			Description: "The ID of the KMS key used to encrypt the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"maintenance_window": {
			// Property: MaintenanceWindow
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
			//   "type": "string"
			// }
			Description: "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
			Type:        types.StringType,
			Optional:    true,
		},
		"node_type": {
			// Property: NodeType
			// CloudFormation resource type schema:
			// {
			//   "description": "The compute and memory capacity of the nodes in the cluster.",
			//   "type": "string"
			// }
			Description: "The compute and memory capacity of the nodes in the cluster.",
			Type:        types.StringType,
			Required:    true,
		},
		"num_replicas_per_shard": {
			// Property: NumReplicasPerShard
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of replicas to apply to each shard. The limit is 5.",
			//   "type": "integer"
			// }
			Description: "The number of replicas to apply to each shard. The limit is 5.",
			Type:        types.Int64Type,
			Optional:    true,
		},
		"num_shards": {
			// Property: NumShards
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of shards the cluster will contain.",
			//   "type": "integer"
			// }
			Description: "The number of shards the cluster will contain.",
			Type:        types.Int64Type,
			Optional:    true,
		},
		"parameter_group_name": {
			// Property: ParameterGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the parameter group associated with the cluster.",
			//   "type": "string"
			// }
			Description: "The name of the parameter group associated with the cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"parameter_group_status": {
			// Property: ParameterGroupStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the parameter group used by the cluster.",
			//   "type": "string"
			// }
			Description: "The status of the parameter group used by the cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			// {
			//   "description": "The port number on which each member of the cluster accepts connections.",
			//   "type": "integer"
			// }
			Description: "The port number on which each member of the cluster accepts connections.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more Amazon VPC security groups associated with this cluster.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "One or more Amazon VPC security groups associated with this cluster.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"snapshot_arns": {
			// Property: SnapshotArns
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"snapshot_name": {
			// Property: SnapshotName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.",
			//   "type": "string"
			// }
			Description: "The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"snapshot_retention_limit": {
			// Property: SnapshotRetentionLimit
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
			//   "type": "integer"
			// }
			Description: "The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
			Type:        types.Int64Type,
			Optional:    true,
		},
		"snapshot_window": {
			// Property: SnapshotWindow
			// CloudFormation resource type schema:
			// {
			//   "description": "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.",
			//   "type": "string"
			// }
			Description: "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
			Type:        types.StringType,
			Optional:    true,
		},
		"sns_topic_status": {
			// Property: SnsTopicStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.",
			//   "type": "string"
			// }
			Description: "The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.",
			Type:        types.StringType,
			Optional:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the cluster. For example, Available, Updating, Creating.",
			//   "type": "string"
			// }
			Description: "The status of the cluster. For example, Available, Updating, Creating.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"subnet_group_name": {
			// Property: SubnetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the subnet group to be used for the cluster.",
			//   "type": "string"
			// }
			Description: "The name of the subnet group to be used for the cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tls_enabled": {
			// Property: TLSEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
			//   "type": "boolean"
			// }
			Description: "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this cluster.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for the tag. May not be null.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value. May be null.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this cluster.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for the tag. May not be null.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The tag's value. May be null.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::Cluster").WithTerraformTypeName("awscc_memorydb_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acl_name":                   "ACLName",
		"address":                    "Address",
		"arn":                        "ARN",
		"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
		"cluster_endpoint":           "ClusterEndpoint",
		"cluster_name":               "ClusterName",
		"description":                "Description",
		"engine_version":             "EngineVersion",
		"final_snapshot_name":        "FinalSnapshotName",
		"key":                        "Key",
		"kms_key_id":                 "KmsKeyId",
		"maintenance_window":         "MaintenanceWindow",
		"node_type":                  "NodeType",
		"num_replicas_per_shard":     "NumReplicasPerShard",
		"num_shards":                 "NumShards",
		"parameter_group_name":       "ParameterGroupName",
		"parameter_group_status":     "ParameterGroupStatus",
		"port":                       "Port",
		"security_group_ids":         "SecurityGroupIds",
		"snapshot_arns":              "SnapshotArns",
		"snapshot_name":              "SnapshotName",
		"snapshot_retention_limit":   "SnapshotRetentionLimit",
		"snapshot_window":            "SnapshotWindow",
		"sns_topic_arn":              "SnsTopicArn",
		"sns_topic_status":           "SnsTopicStatus",
		"status":                     "Status",
		"subnet_group_name":          "SubnetGroupName",
		"tags":                       "Tags",
		"tls_enabled":                "TLSEnabled",
		"value":                      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
