data "zendesk_search" "test" {
  query = "email:jacob.potter@dynatrace.com"
}


resource "zendesk_trigger" "test" {
  depends_on = [zendesk_trigger_category.test]
  title      = var.title
  position   = 1
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "notification_zis"
      target          = "slack"
      slack_workspace = "T73ASNB9Q"
      slack_channel   = "C0677RLQDS4"
      slack_title     = "Test Notification"
      value           = "Test body"
    }
  ]
  conditions = {
    any = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      }
  ] }
  category_id = zendesk_trigger_category.test.id
}

resource "zendesk_trigger_category" "test" {
  name = "test_acc_${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}
