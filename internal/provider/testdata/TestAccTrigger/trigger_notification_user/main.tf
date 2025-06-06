data "zendesk_search" "test" {
  query = "email:web.wizards.software.dev@gmail.com"
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
      field                = "notification_user"
      target               = data.zendesk_search.test.results.users[0].id
      notification_subject = "Test email"
      value                = "Test body"
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
  name = "${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}
