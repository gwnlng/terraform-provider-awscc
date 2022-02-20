// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("awscc_iot_scheduled_audit", scheduledAuditResourceType)
}

// scheduledAuditResourceType returns the Terraform awscc_iot_scheduled_audit resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::ScheduledAudit resource type.
func scheduledAuditResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"day_of_month": {
			// Property: DayOfMonth
			// CloudFormation resource type schema:
			// {
			//   "description": "The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.",
			//   "pattern": "^([1-9]|[12][0-9]|3[01])$|^LAST$",
			//   "type": "string"
			// }
			Description: "The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^([1-9]|[12][0-9]|3[01])$|^LAST$"), ""),
			},
		},
		"day_of_week": {
			// Property: DayOfWeek
			// CloudFormation resource type schema:
			// {
			//   "description": "The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.",
			//   "enum": [
			//     "SUN",
			//     "MON",
			//     "TUE",
			//     "WED",
			//     "THU",
			//     "FRI",
			//     "SAT"
			//   ],
			//   "type": "string"
			// }
			Description: "The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SUN",
					"MON",
					"TUE",
					"WED",
					"THU",
					"FRI",
					"SAT",
				}),
			},
		},
		"frequency": {
			// Property: Frequency
			// CloudFormation resource type schema:
			// {
			//   "description": "How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.",
			//   "enum": [
			//     "DAILY",
			//     "WEEKLY",
			//     "BIWEEKLY",
			//     "MONTHLY"
			//   ],
			//   "type": "string"
			// }
			Description: "How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DAILY",
					"WEEKLY",
					"BIWEEKLY",
					"MONTHLY",
				}),
			},
		},
		"scheduled_audit_arn": {
			// Property: ScheduledAuditArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN (Amazon resource name) of the scheduled audit.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN (Amazon resource name) of the scheduled audit.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"scheduled_audit_name": {
			// Property: ScheduledAuditName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to give to the scheduled audit.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9:_-]+",
			//   "type": "string"
			// }
			Description: "The name you want to give to the scheduled audit.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9:_-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
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
		"target_check_names": {
			// Property: TargetCheckNames
			// CloudFormation resource type schema:
			// {
			//   "description": "Which checks are performed during the scheduled audit. Checks must be enabled for your account.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Which checks are performed during the scheduled audit. Checks must be enabled for your account.",
			Type:        types.SetType{ElemType: types.StringType},
			Required:    true,
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
		Description: "Scheduled audits can be used to specify the checks you want to perform during an audit and how often the audit should be run.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ScheduledAudit").WithTerraformTypeName("awscc_iot_scheduled_audit")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"day_of_month":         "DayOfMonth",
		"day_of_week":          "DayOfWeek",
		"frequency":            "Frequency",
		"key":                  "Key",
		"scheduled_audit_arn":  "ScheduledAuditArn",
		"scheduled_audit_name": "ScheduledAuditName",
		"tags":                 "Tags",
		"target_check_names":   "TargetCheckNames",
		"value":                "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
