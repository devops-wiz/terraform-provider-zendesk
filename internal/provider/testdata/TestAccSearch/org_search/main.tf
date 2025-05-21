data "zendesk_search" "test" {
  query = "type:organization"
}

output "test_org_id" {
  value = data.zendesk_search.test.results.organizations[0].id
}