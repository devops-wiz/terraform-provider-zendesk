resource "zendesk_automation" "test" {
  title = var.title
  actions = [
    {
      field = "status"
      value = "open"
    }
  ]
  conditions = {
    all = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      },
      {
        field    = "NEW"
        operator = "is"
        value    = "2"
      }
  ] }

  position = 1
}

variable "title" {
  type     = string
  nullable = false
}
