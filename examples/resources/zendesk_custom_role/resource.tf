resource "zendesk_custom_role" "test" {
  name      = "test_role"
  role_type = 0
  configuration = {
    manage_automations = true
  }
}

