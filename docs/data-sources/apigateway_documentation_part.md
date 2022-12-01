---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_documentation_part Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApiGateway::DocumentationPart
---

# awscc_apigateway_documentation_part (Data Source)

Data Source schema for AWS::ApiGateway::DocumentationPart



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `documentation_part_id` (String) The identifier of the documentation Part.
- `location` (Attributes) The location of the API entity that the documentation applies to. (see [below for nested schema](#nestedatt--location))
- `properties` (String) The documentation content map of the targeted API entity.
- `rest_api_id` (String) Identifier of the targeted API entity

<a id="nestedatt--location"></a>
### Nested Schema for `location`

Read-Only:

- `method` (String) The HTTP verb of a method.
- `name` (String) The name of the targeted API entity.
- `path` (String) The URL path of the target.
- `status_code` (String) The HTTP status code of a response.
- `type` (String) The type of API entity that the documentation content applies to.

