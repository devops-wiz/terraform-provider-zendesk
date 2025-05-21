resource "zendesk_view" "test" {
  title = var.title

  conditions = {
    all = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      }
    ]

    any = [
      {
        field    = "type",
        operator = "is",
        value    = "incident"
      },
      {
        field    = "type",
        operator = "is",
        value    = "task"
      }
    ]
  }


  output = {
    columns     = ["status", "assignee"]
    group_by    = "status"
    group_order = "asc"
    sort_by     = "assignee"
    sort_order  = "asc"
  }
}

variable "title" {
  type     = string
  nullable = false
}
