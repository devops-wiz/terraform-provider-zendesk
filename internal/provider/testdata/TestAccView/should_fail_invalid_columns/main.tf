resource "zendesk_view" "test" {
  title = "test_acc_invalid"

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
    columns     = ["status", "blah"]
    group_by    = "status"
    group_order = "asc"
    sort_by     = "assignee"
    sort_order  = "asc"
  }
}
