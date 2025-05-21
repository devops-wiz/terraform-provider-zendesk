resource "zendesk_macro" "test" {
  title = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field  = "side_conversation_slack"
      target = "help-zendesk"
      value  = "Blah"
    },
    {
      field                = "side_conversation_ticket"
      notification_subject = "Test"
      value                = "Test"
      target               = "SupportAssignable:support_assignable/group/${zendesk_group.test.id}"
      content_type         = "text/html"
    }
  ]
}

resource "zendesk_group" "test" {
  description = "Test"
  is_public   = true
  name        = "Test"
}


variable "title" {
  type     = string
  nullable = false
}
