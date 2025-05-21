resource "zendesk_automation" "test" {
  title = "Test Automation"
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
}
