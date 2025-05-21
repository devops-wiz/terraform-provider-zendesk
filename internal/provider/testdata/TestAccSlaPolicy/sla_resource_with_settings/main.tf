resource "zendesk_sla_policy" "test" {
  title = var.title
  filter = {
    any = [
      {
        field    = "current_tags",
        operator = "includes",
        value    = "tag1 tag2"
      }
    ]
  }
  policy_metrics = [
    {
      priority       = "low"
      metric         = "agent_work_time"
      target         = 60
      business_hours = false
    }
  ]

  metrics_settings = {
    first_reply_time = {
      fulfill_on_agent_internal_note = true
    }
  }

}

variable "title" {
  type     = string
  nullable = false
}