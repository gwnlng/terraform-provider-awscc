---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediapackage_asset Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaPackage::Asset
---

# awscc_mediapackage_asset (Data Source)

Data Source schema for AWS::MediaPackage::Asset



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The ARN of the Asset.
- **created_at** (String) The time the Asset was initially submitted for Ingest.
- **egress_endpoints** (Attributes List) The list of egress endpoints available for the Asset. (see [below for nested schema](#nestedatt--egress_endpoints))
- **packaging_group_id** (String) The ID of the PackagingGroup for the Asset.
- **resource_id** (String) The resource ID to include in SPEKE key requests.
- **source_arn** (String) ARN of the source object in S3.
- **source_role_arn** (String) The IAM role_arn used to access the source S3 bucket.
- **tags** (Attributes List) A collection of tags associated with a resource (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--egress_endpoints"></a>
### Nested Schema for `egress_endpoints`

Read-Only:

- **packaging_configuration_id** (String) The ID of the PackagingConfiguration being applied to the Asset.
- **url** (String) The URL of the parent manifest for the repackaged Asset.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

