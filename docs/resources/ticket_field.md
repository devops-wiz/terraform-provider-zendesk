---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zendesk_ticket_field Resource - zendesk"
subcategory: ""
description: |-
  Manages a Zendesk ticket field. A ticket field provides a field in your Zendesk tickets that can store custom data. The field can be visible and editable by both agents and end-users depending on the configuration.
---

# zendesk_ticket_field (Resource)

Manages a Zendesk ticket field. A ticket field provides a field in your Zendesk tickets that can store custom data. The field can be visible and editable by both agents and end-users depending on the configuration.

## Example Usage

```terraform
resource "zendesk_ticket_field" "test2" {
  title = " Terraform test 3"
  type  = "tagger"

  custom_field_options = [
    {
      name  = "Test"
      value = "test_tag"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `title` (String) The title of the ticket field
- `type` (String) Ticket Field Type, acceptable values include status, description, subject, tickettype, priority, group, assignee, custom_status, text, textarea, checkbox, date, integer, decimal, regexp, partial_credit_card, multiselect, tagger.

### Optional

- `active` (Boolean) Whether this field is available
- `agent_description` (String) A description of the ticket field that only agents can see
- `custom_field_options` (Attributes List) Required and presented for a custom ticket field of type 'multiselect' or 'tagger' (see [below for nested schema](#nestedatt--custom_field_options))
- `editable_in_portal` (Boolean) Whether this field is editable by end users in Help Center
- `portal_description` (String) Describes the purpose of the ticket field to users
- `position` (Number) The relative position of the ticket field on a ticket. Note that for accounts with ticket forms, positions are controlled by the different forms
- `regexp_for_validation` (String) For 'regexp' fields only. The validation pattern for a field value to be deemed valid
- `required` (Boolean) If true, agents must enter a value in the field to change the ticket status to solved
- `required_in_portal` (Boolean) If true, end users must enter a value in the field to create the request
- `tag` (String) For 'checkbox' fields only. A tag added to tickets when the checkbox field is selected
- `title_in_portal` (String) The title of the ticket field for end users in Help Center
- `visible_in_portal` (Boolean) Whether this field is visible to end users in Help Center

### Read-Only

- `created_at` (String) The time the custom ticket field was created
- `id` (Number) Ticket Field ID. Automatically assigned when created
- `system_field_options` (Attributes List) (see [below for nested schema](#nestedatt--system_field_options))
- `updated_at` (String) The time the custom ticket field was last updated
- `url` (String) The URL for this resource

<a id="nestedatt--custom_field_options"></a>
### Nested Schema for `custom_field_options`

Required:

- `name` (String)
- `value` (String)

Read-Only:

- `id` (Number)


<a id="nestedatt--system_field_options"></a>
### Nested Schema for `system_field_options`

Read-Only:

- `name` (String) The name of the system field value
- `value` (String) The value of the system field value
