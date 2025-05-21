resource "zendesk_automation" "test" {
  title = "test_acc_bad_config"
  actions = [
    {
      field = "status"
      value = "open"
    }
  ]
  conditions = {
    // Invalid config, no empty fields
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
      },
      {
        field    = ""
        operator = "is"
        value    = "something"
      }
  ] }
}
