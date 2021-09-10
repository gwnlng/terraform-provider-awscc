---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_prefix_list Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::PrefixList
---

# awscc_ec2_prefix_list (Data Source)

Data Source schema for AWS::EC2::PrefixList



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **address_family** (String) Ip Version of Prefix List.
- **arn** (String) The Amazon Resource Name (ARN) of the Prefix List.
- **entries** (Attributes List) Entries of Prefix List. (see [below for nested schema](#nestedatt--entries))
- **max_entries** (Number) Max Entries of Prefix List.
- **owner_id** (String) Owner Id of Prefix List.
- **prefix_list_id** (String) Id of Prefix List.
- **prefix_list_name** (String) Name of Prefix List.
- **tags** (Attributes List) Tags for Prefix List (see [below for nested schema](#nestedatt--tags))
- **version** (Number) Version of Prefix List.

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Read-Only:

- **cidr** (String)
- **description** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

