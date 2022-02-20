// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

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
	registry.AddResourceTypeFactory("awscc_connect_user", userResourceType)
}

// userResourceType returns the Terraform awscc_connect_user resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Connect::User resource type.
func userResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"directory_user_id": {
			// Property: DirectoryUserId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the user account in the directory used for identity management.",
			//   "type": "string"
			// }
			Description: "The identifier of the user account in the directory used for identity management.",
			Type:        types.StringType,
			Optional:    true,
		},
		"hierarchy_group_arn": {
			// Property: HierarchyGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the hierarchy group for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the hierarchy group for the user.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"identity_info": {
			// Property: IdentityInfo
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The information about the identity of the user.",
			//   "properties": {
			//     "Email": {
			//       "description": "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
			//       "type": "string"
			//     },
			//     "FirstName": {
			//       "description": "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
			//       "type": "string"
			//     },
			//     "LastName": {
			//       "description": "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The information about the identity of the user.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"email": {
						// Property: Email
						Description: "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
						Type:        types.StringType,
						Optional:    true,
					},
					"first_name": {
						// Property: FirstName
						Description: "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
						Type:        types.StringType,
						Optional:    true,
					},
					"last_name": {
						// Property: LastName
						Description: "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the Amazon Connect instance.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the Amazon Connect instance.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"password": {
			// Property: Password
			// CloudFormation resource type schema:
			// {
			//   "description": "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
			Type:        types.StringType,
			Optional:    true,
			// Password is a write-only property.
		},
		"phone_config": {
			// Property: PhoneConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The phone settings for the user.",
			//   "properties": {
			//     "AfterContactWorkTimeLimit": {
			//       "description": "The After Call Work (ACW) timeout setting, in seconds.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "AutoAccept": {
			//       "description": "The Auto accept setting.",
			//       "type": "boolean"
			//     },
			//     "DeskPhoneNumber": {
			//       "description": "The phone number for the user's desk phone.",
			//       "type": "string"
			//     },
			//     "PhoneType": {
			//       "description": "The phone type.",
			//       "enum": [
			//         "SOFT_PHONE",
			//         "DESK_PHONE"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "PhoneType"
			//   ],
			//   "type": "object"
			// }
			Description: "The phone settings for the user.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"after_contact_work_time_limit": {
						// Property: AfterContactWorkTimeLimit
						Description: "The After Call Work (ACW) timeout setting, in seconds.",
						Type:        types.Int64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
					"auto_accept": {
						// Property: AutoAccept
						Description: "The Auto accept setting.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"desk_phone_number": {
						// Property: DeskPhoneNumber
						Description: "The phone number for the user's desk phone.",
						Type:        types.StringType,
						Optional:    true,
					},
					"phone_type": {
						// Property: PhoneType
						Description: "The phone type.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SOFT_PHONE",
								"DESK_PHONE",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"routing_profile_arn": {
			// Property: RoutingProfileArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the routing profile for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/routing-profile/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The identifier of the routing profile for the user.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/routing-profile/[-a-zA-Z0-9]*$"), ""),
			},
		},
		"security_profile_arns": {
			// Property: SecurityProfileArns
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more security profile arns for the user",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "The identifier of the security profile for the user.",
			//     "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/security-profile/[-a-zA-Z0-9]*$",
			//     "type": "string"
			//   },
			//   "maxItems": 10,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "One or more security profile arns for the user",
			Type:        types.SetType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 10),
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/security-profile/[-a-zA-Z0-9]*$"), "")),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "One or more tags.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
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
			Description: "One or more tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
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
		"user_arn": {
			// Property: UserArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the user.",
			//   "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent/[-a-zA-Z0-9]*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the user.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"username": {
			// Property: Username
			// CloudFormation resource type schema:
			// {
			//   "description": "The user name for the account.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9\\_\\-\\.\\@]+",
			//   "type": "string"
			// }
			Description: "The user name for the account.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9\\_\\-\\.\\@]+"), ""),
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
		Description: "Resource Type definition for AWS::Connect::User",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::User").WithTerraformTypeName("awscc_connect_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"after_contact_work_time_limit": "AfterContactWorkTimeLimit",
		"auto_accept":                   "AutoAccept",
		"desk_phone_number":             "DeskPhoneNumber",
		"directory_user_id":             "DirectoryUserId",
		"email":                         "Email",
		"first_name":                    "FirstName",
		"hierarchy_group_arn":           "HierarchyGroupArn",
		"identity_info":                 "IdentityInfo",
		"instance_arn":                  "InstanceArn",
		"key":                           "Key",
		"last_name":                     "LastName",
		"password":                      "Password",
		"phone_config":                  "PhoneConfig",
		"phone_type":                    "PhoneType",
		"routing_profile_arn":           "RoutingProfileArn",
		"security_profile_arns":         "SecurityProfileArns",
		"tags":                          "Tags",
		"user_arn":                      "UserArn",
		"username":                      "Username",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Password",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
