resource "zendesk_schedule" "test" {
  name      = var.title
  time_zone = "Pacific Time (US & Canada)"

  intervals = {
    sunday = {
      start_time = 5
      end_time   = 13
    }
  }
}

variable "title" {
  type     = string
  nullable = false
}