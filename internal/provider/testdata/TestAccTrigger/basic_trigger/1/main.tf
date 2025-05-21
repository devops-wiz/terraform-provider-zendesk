resource "zendesk_trigger" "test" {
  depends_on = [zendesk_trigger_category.test]
  title      = var.title
  position   = 1
  actions = [
    {
      field = "status"
      value = "open"
    },
  ]
  conditions = {
    any = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      },
      {
        field    = "via_id"
        operator = "includes"
        values   = ["4", "75"]
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
