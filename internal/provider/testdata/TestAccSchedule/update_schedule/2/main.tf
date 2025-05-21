resource "zendesk_schedule" "test" {
  name      = var.title
  time_zone = "Pacific Time (US & Canada)"

  intervals = {
    sunday = {
      start_time = 5
      end_time   = 13
    }
    monday = {
      start_time = 5
      end_time   = 13
    }
    tuesday = {
      start_time = 4
      end_time   = 19
    }
    wednesday = {
      start_time = 5
      end_time   = 13
    }
    thursday = {
      start_time = 5
      end_time   = 13
    }
    friday = {
      start_time = 5
      end_time   = 13
    }
    saturday = {
      start_time = 5
      end_time   = 13
    }
  }
}

variable "title" {
  type     = string
  nullable = false
}