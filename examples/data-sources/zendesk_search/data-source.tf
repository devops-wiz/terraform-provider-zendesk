data "zendesk_search" "test" {
  query = "type:user"
}

output "test_user_email" {
  value = data.zendesk_search.test.results.users[0].email
}