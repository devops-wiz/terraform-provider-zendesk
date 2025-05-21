resource "zendesk_ticket_field" "subject" {
  title           = "Subject"
  title_in_portal = "Subject"
  type            = "subject"
}