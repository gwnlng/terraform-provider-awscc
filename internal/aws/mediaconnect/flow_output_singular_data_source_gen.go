// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect

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
	registry.AddDataSourceTypeFactory("awscc_mediaconnect_flow_output", flowOutputDataSourceType)
}

// flowOutputDataSourceType returns the Terraform awscc_mediaconnect_flow_output data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::MediaConnect::FlowOutput resource type.
func flowOutputDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cidr_allow_list": {
			// Property: CidrAllowList
			// CloudFormation resource type schema:
			// {
			//   "description": "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the output.",
			//   "type": "string"
			// }
			Description: "A description of the output.",
			Type:        types.StringType,
			Computed:    true,
		},
		"destination": {
			// Property: Destination
			// CloudFormation resource type schema:
			// {
			//   "description": "The address where you want to send the output.",
			//   "type": "string"
			// }
			Description: "The address where you want to send the output.",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption": {
			// Property: Encryption
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the encryption of the flow.",
			//   "properties": {
			//     "Algorithm": {
			//       "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			//       "enum": [
			//         "aes128",
			//         "aes192",
			//         "aes256"
			//       ],
			//       "type": "string"
			//     },
			//     "KeyType": {
			//       "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			//       "enum": [
			//         "static-key"
			//       ],
			//       "type": "string"
			//     },
			//     "RoleArn": {
			//       "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			//       "type": "string"
			//     },
			//     "SecretArn": {
			//       "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Algorithm",
			//     "RoleArn",
			//     "SecretArn"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about the encryption of the flow.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"algorithm": {
						// Property: Algorithm
						Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						Type:        types.StringType,
						Computed:    true,
					},
					"key_type": {
						// Property: KeyType
						Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						Type:        types.StringType,
						Computed:    true,
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						Type:        types.StringType,
						Computed:    true,
					},
					"secret_arn": {
						// Property: SecretArn
						Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"max_latency": {
			// Property: MaxLatency
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			//   "type": "integer"
			// }
			Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the output. This value must be unique within the current flow.",
			//   "type": "string"
			// }
			Description: "The name of the output. This value must be unique within the current flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"output_arn": {
			// Property: OutputArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the output.",
			//   "type": "string"
			// }
			Description: "The ARN of the output.",
			Type:        types.StringType,
			Computed:    true,
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			// {
			//   "description": "The port to use when content is distributed to this output.",
			//   "type": "integer"
			// }
			Description: "The port to use when content is distributed to this output.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "description": "The protocol that is used by the source or output.",
			//   "enum": [
			//     "zixi-push",
			//     "rtp-fec",
			//     "rtp",
			//     "zixi-pull",
			//     "rist"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol that is used by the source or output.",
			Type:        types.StringType,
			Computed:    true,
		},
		"remote_id": {
			// Property: RemoteId
			// CloudFormation resource type schema:
			// {
			//   "description": "The remote ID for the Zixi-pull stream.",
			//   "type": "string"
			// }
			Description: "The remote ID for the Zixi-pull stream.",
			Type:        types.StringType,
			Computed:    true,
		},
		"smoothing_latency": {
			// Property: SmoothingLatency
			// CloudFormation resource type schema:
			// {
			//   "description": "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
			//   "type": "integer"
			// }
			Description: "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"stream_id": {
			// Property: StreamId
			// CloudFormation resource type schema:
			// {
			//   "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			//   "type": "string"
			// }
			Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_interface_attachment": {
			// Property: VpcInterfaceAttachment
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The settings for attaching a VPC interface to an output.",
			//   "properties": {
			//     "VpcInterfaceName": {
			//       "description": "The name of the VPC interface to use for this output.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The settings for attaching a VPC interface to an output.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"vpc_interface_name": {
						// Property: VpcInterfaceName
						Description: "The name of the VPC interface to use for this output.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::MediaConnect::FlowOutput",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowOutput").WithTerraformTypeName("awscc_mediaconnect_flow_output")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                "Algorithm",
		"cidr_allow_list":          "CidrAllowList",
		"description":              "Description",
		"destination":              "Destination",
		"encryption":               "Encryption",
		"flow_arn":                 "FlowArn",
		"key_type":                 "KeyType",
		"max_latency":              "MaxLatency",
		"name":                     "Name",
		"output_arn":               "OutputArn",
		"port":                     "Port",
		"protocol":                 "Protocol",
		"remote_id":                "RemoteId",
		"role_arn":                 "RoleArn",
		"secret_arn":               "SecretArn",
		"smoothing_latency":        "SmoothingLatency",
		"stream_id":                "StreamId",
		"vpc_interface_attachment": "VpcInterfaceAttachment",
		"vpc_interface_name":       "VpcInterfaceName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_mediaconnect_flow_output", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
