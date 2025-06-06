resource "zendesk_ticket_field" "test" {
  title             = var.title
  type              = "tagger"
  agent_description = "test2"

}

variable "title" {
  type     = string
  nullable = false
}
