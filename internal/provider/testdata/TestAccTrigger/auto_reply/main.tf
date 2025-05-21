resource "zendesk_trigger" "test" {
  depends_on = [zendesk_trigger_category.test]
  title      = var.title
  position   = 1
  actions = [
    {
      field = "reply_public"
      value = "Test body"
    },
    {
      field = "reply_internal"
      value = "Test body internal"
    }
  ]
  conditions = {
    any = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      }
    ]
  }
  category_id = zendesk_trigger_category.test.id
}

resource "zendesk_trigger_category" "test" {
  name = "test_acc_${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}
