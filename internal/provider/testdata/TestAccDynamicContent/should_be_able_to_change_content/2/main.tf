resource "zendesk_dynamic_content" "test" {
  name = var.title
  variants = [
    {
      content   = "test changed"
      locale_id = data.zendesk_locale.en_us.locale.id
      active    = true
      default   = true
    }
  ]
  default_locale_id = data.zendesk_locale.en_us.locale.id
}

data "zendesk_locale" "en_us" {
  code = "en-US"
}

variable "title" {
  nullable = false
  type     = string
}