resource "zendesk_trigger" "test" {
  title = "TF Test"
  actions = [
    {
      field  = "notification_webhook",
      target = "01HWNR72XCXH956BBDVPGQB386",
      value  = jsonencode({ "test" : "value" })
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
  name = "test"
}


