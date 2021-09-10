// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

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
	registry.AddDataSourceTypeFactory("awscc_ec2_ec2_fleet", eC2FleetDataSourceType)
}

// eC2FleetDataSourceType returns the Terraform awscc_ec2_ec2_fleet data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::EC2Fleet resource type.
func eC2FleetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"context": {
			// Property: Context
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"excess_capacity_termination_policy": {
			// Property: ExcessCapacityTerminationPolicy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "termination",
			//     "no-termination"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"fleet_id": {
			// Property: FleetId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"launch_template_configs": {
			// Property: LaunchTemplateConfigs
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "LaunchTemplateSpecification": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "LaunchTemplateId": {
			//             "type": "string"
			//           },
			//           "LaunchTemplateName": {
			//             "maxLength": 128,
			//             "minLength": 3,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Version": {
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Overrides": {
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AvailabilityZone": {
			//               "type": "string"
			//             },
			//             "InstanceType": {
			//               "type": "string"
			//             },
			//             "MaxPrice": {
			//               "type": "string"
			//             },
			//             "Placement": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Affinity": {
			//                   "type": "string"
			//                 },
			//                 "AvailabilityZone": {
			//                   "type": "string"
			//                 },
			//                 "GroupName": {
			//                   "type": "string"
			//                 },
			//                 "HostId": {
			//                   "type": "string"
			//                 },
			//                 "HostResourceGroupArn": {
			//                   "type": "string"
			//                 },
			//                 "PartitionNumber": {
			//                   "type": "integer"
			//                 },
			//                 "SpreadDomain": {
			//                   "type": "string"
			//                 },
			//                 "Tenancy": {
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "Priority": {
			//               "type": "number"
			//             },
			//             "SubnetId": {
			//               "type": "string"
			//             },
			//             "WeightedCapacity": {
			//               "type": "number"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"launch_template_specification": {
						// Property: LaunchTemplateSpecification
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"launch_template_id": {
									// Property: LaunchTemplateId
									Type:     types.StringType,
									Computed: true,
								},
								"launch_template_name": {
									// Property: LaunchTemplateName
									Type:     types.StringType,
									Computed: true,
								},
								"version": {
									// Property: Version
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"overrides": {
						// Property: Overrides
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"availability_zone": {
									// Property: AvailabilityZone
									Type:     types.StringType,
									Computed: true,
								},
								"instance_type": {
									// Property: InstanceType
									Type:     types.StringType,
									Computed: true,
								},
								"max_price": {
									// Property: MaxPrice
									Type:     types.StringType,
									Computed: true,
								},
								"placement": {
									// Property: Placement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"affinity": {
												// Property: Affinity
												Type:     types.StringType,
												Computed: true,
											},
											"availability_zone": {
												// Property: AvailabilityZone
												Type:     types.StringType,
												Computed: true,
											},
											"group_name": {
												// Property: GroupName
												Type:     types.StringType,
												Computed: true,
											},
											"host_id": {
												// Property: HostId
												Type:     types.StringType,
												Computed: true,
											},
											"host_resource_group_arn": {
												// Property: HostResourceGroupArn
												Type:     types.StringType,
												Computed: true,
											},
											"partition_number": {
												// Property: PartitionNumber
												Type:     types.NumberType,
												Computed: true,
											},
											"spread_domain": {
												// Property: SpreadDomain
												Type:     types.StringType,
												Computed: true,
											},
											"tenancy": {
												// Property: Tenancy
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"priority": {
									// Property: Priority
									Type:     types.NumberType,
									Computed: true,
								},
								"subnet_id": {
									// Property: SubnetId
									Type:     types.StringType,
									Computed: true,
								},
								"weighted_capacity": {
									// Property: WeightedCapacity
									Type:     types.NumberType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"on_demand_options": {
			// Property: OnDemandOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllocationStrategy": {
			//       "type": "string"
			//     },
			//     "CapacityReservationOptions": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "UsageStrategy": {
			//           "enum": [
			//             "use-capacity-reservations-first"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "MaxTotalPrice": {
			//       "type": "string"
			//     },
			//     "MinTargetCapacity": {
			//       "type": "integer"
			//     },
			//     "SingleAvailabilityZone": {
			//       "type": "boolean"
			//     },
			//     "SingleInstanceType": {
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allocation_strategy": {
						// Property: AllocationStrategy
						Type:     types.StringType,
						Computed: true,
					},
					"capacity_reservation_options": {
						// Property: CapacityReservationOptions
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"usage_strategy": {
									// Property: UsageStrategy
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"max_total_price": {
						// Property: MaxTotalPrice
						Type:     types.StringType,
						Computed: true,
					},
					"min_target_capacity": {
						// Property: MinTargetCapacity
						Type:     types.NumberType,
						Computed: true,
					},
					"single_availability_zone": {
						// Property: SingleAvailabilityZone
						Type:     types.BoolType,
						Computed: true,
					},
					"single_instance_type": {
						// Property: SingleInstanceType
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"replace_unhealthy_instances": {
			// Property: ReplaceUnhealthyInstances
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"spot_options": {
			// Property: SpotOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllocationStrategy": {
			//       "enum": [
			//         "lowestPrice",
			//         "diversified",
			//         "capacityOptimized"
			//       ],
			//       "type": "string"
			//     },
			//     "InstanceInterruptionBehavior": {
			//       "enum": [
			//         "hibernate",
			//         "stop",
			//         "terminate"
			//       ],
			//       "type": "string"
			//     },
			//     "InstancePoolsToUseCount": {
			//       "type": "integer"
			//     },
			//     "MaxTotalPrice": {
			//       "type": "string"
			//     },
			//     "MinTargetCapacity": {
			//       "type": "integer"
			//     },
			//     "SingleAvailabilityZone": {
			//       "type": "boolean"
			//     },
			//     "SingleInstanceType": {
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allocation_strategy": {
						// Property: AllocationStrategy
						Type:     types.StringType,
						Computed: true,
					},
					"instance_interruption_behavior": {
						// Property: InstanceInterruptionBehavior
						Type:     types.StringType,
						Computed: true,
					},
					"instance_pools_to_use_count": {
						// Property: InstancePoolsToUseCount
						Type:     types.NumberType,
						Computed: true,
					},
					"max_total_price": {
						// Property: MaxTotalPrice
						Type:     types.StringType,
						Computed: true,
					},
					"min_target_capacity": {
						// Property: MinTargetCapacity
						Type:     types.NumberType,
						Computed: true,
					},
					"single_availability_zone": {
						// Property: SingleAvailabilityZone
						Type:     types.BoolType,
						Computed: true,
					},
					"single_instance_type": {
						// Property: SingleInstanceType
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tag_specifications": {
			// Property: TagSpecifications
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ResourceType": {
			//         "enum": [
			//           "client-vpn-endpoint",
			//           "customer-gateway",
			//           "dedicated-host",
			//           "dhcp-options",
			//           "egress-only-internet-gateway",
			//           "elastic-gpu",
			//           "elastic-ip",
			//           "export-image-task",
			//           "export-instance-task",
			//           "fleet",
			//           "fpga-image",
			//           "host-reservation",
			//           "image",
			//           "import-image-task",
			//           "import-snapshot-task",
			//           "instance",
			//           "internet-gateway",
			//           "key-pair",
			//           "launch-template",
			//           "local-gateway-route-table-vpc-association",
			//           "natgateway",
			//           "network-acl",
			//           "network-insights-analysis",
			//           "network-insights-path",
			//           "network-interface",
			//           "placement-group",
			//           "reserved-instances",
			//           "route-table",
			//           "security-group",
			//           "snapshot",
			//           "spot-fleet-request",
			//           "spot-instances-request",
			//           "subnet",
			//           "traffic-mirror-filter",
			//           "traffic-mirror-session",
			//           "traffic-mirror-target",
			//           "transit-gateway",
			//           "transit-gateway-attachment",
			//           "transit-gateway-connect-peer",
			//           "transit-gateway-multicast-domain",
			//           "transit-gateway-route-table",
			//           "volume",
			//           "vpc",
			//           "vpc-flow-log",
			//           "vpc-peering-connection",
			//           "vpn-connection",
			//           "vpn-gateway"
			//         ],
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "type": "string"
			//             },
			//             "Value": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Value",
			//             "Key"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"resource_type": {
						// Property: ResourceType
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
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
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"target_capacity_specification": {
			// Property: TargetCapacitySpecification
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DefaultTargetCapacityType": {
			//       "enum": [
			//         "on-demand",
			//         "spot"
			//       ],
			//       "type": "string"
			//     },
			//     "OnDemandTargetCapacity": {
			//       "type": "integer"
			//     },
			//     "SpotTargetCapacity": {
			//       "type": "integer"
			//     },
			//     "TotalTargetCapacity": {
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "TotalTargetCapacity"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"default_target_capacity_type": {
						// Property: DefaultTargetCapacityType
						Type:     types.StringType,
						Computed: true,
					},
					"on_demand_target_capacity": {
						// Property: OnDemandTargetCapacity
						Type:     types.NumberType,
						Computed: true,
					},
					"spot_target_capacity": {
						// Property: SpotTargetCapacity
						Type:     types.NumberType,
						Computed: true,
					},
					"total_target_capacity": {
						// Property: TotalTargetCapacity
						Type:     types.NumberType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"terminate_instances_with_expiration": {
			// Property: TerminateInstancesWithExpiration
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "maintain",
			//     "request",
			//     "instant"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"valid_from": {
			// Property: ValidFrom
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"valid_until": {
			// Property: ValidUntil
			// CloudFormation resource type schema:
			// {
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
		Description: "Data Source schema for AWS::EC2::EC2Fleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EC2Fleet").WithTerraformTypeName("awscc_ec2_ec2_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"affinity":                            "Affinity",
		"allocation_strategy":                 "AllocationStrategy",
		"availability_zone":                   "AvailabilityZone",
		"capacity_reservation_options":        "CapacityReservationOptions",
		"context":                             "Context",
		"default_target_capacity_type":        "DefaultTargetCapacityType",
		"excess_capacity_termination_policy":  "ExcessCapacityTerminationPolicy",
		"fleet_id":                            "FleetId",
		"group_name":                          "GroupName",
		"host_id":                             "HostId",
		"host_resource_group_arn":             "HostResourceGroupArn",
		"instance_interruption_behavior":      "InstanceInterruptionBehavior",
		"instance_pools_to_use_count":         "InstancePoolsToUseCount",
		"instance_type":                       "InstanceType",
		"key":                                 "Key",
		"launch_template_configs":             "LaunchTemplateConfigs",
		"launch_template_id":                  "LaunchTemplateId",
		"launch_template_name":                "LaunchTemplateName",
		"launch_template_specification":       "LaunchTemplateSpecification",
		"max_price":                           "MaxPrice",
		"max_total_price":                     "MaxTotalPrice",
		"min_target_capacity":                 "MinTargetCapacity",
		"on_demand_options":                   "OnDemandOptions",
		"on_demand_target_capacity":           "OnDemandTargetCapacity",
		"overrides":                           "Overrides",
		"partition_number":                    "PartitionNumber",
		"placement":                           "Placement",
		"priority":                            "Priority",
		"replace_unhealthy_instances":         "ReplaceUnhealthyInstances",
		"resource_type":                       "ResourceType",
		"single_availability_zone":            "SingleAvailabilityZone",
		"single_instance_type":                "SingleInstanceType",
		"spot_options":                        "SpotOptions",
		"spot_target_capacity":                "SpotTargetCapacity",
		"spread_domain":                       "SpreadDomain",
		"subnet_id":                           "SubnetId",
		"tag_specifications":                  "TagSpecifications",
		"tags":                                "Tags",
		"target_capacity_specification":       "TargetCapacitySpecification",
		"tenancy":                             "Tenancy",
		"terminate_instances_with_expiration": "TerminateInstancesWithExpiration",
		"total_target_capacity":               "TotalTargetCapacity",
		"type":                                "Type",
		"usage_strategy":                      "UsageStrategy",
		"valid_from":                          "ValidFrom",
		"valid_until":                         "ValidUntil",
		"value":                               "Value",
		"version":                             "Version",
		"weighted_capacity":                   "WeightedCapacity",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_ec2_fleet", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
