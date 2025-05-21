import {
  id = 1500002895221
  to = zendesk_ticket_form.import_test
}

resource "zendesk_ticket_form" "import_test" {
  active = true
  agent_conditions = [
    {
      child_fields = [
        {
          id          = 1900002752645
          is_required = false
          required_on_statuses = {
            statuses = null
            type     = "NO_STATUSES"
          }
        },
      ]
      parent_field_id = 1900002750785
      value           = "problem"
    },
    {
      child_fields = [
        {
          id          = 1500011100682
          is_required = false
          required_on_statuses = {
            statuses = null
            type     = "NO_STATUSES"
          }
        },
        {
          id          = 1900002752745
          is_required = false
          required_on_statuses = {
            statuses = null
            type     = "NO_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011100642
      value           = "esc_status_new"
    },
    {
      child_fields = [
        {
          id          = 1500011100682
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
        {
          id          = 1900002752745
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011100642
      value           = "esc_status_active"
    },
    {
      child_fields = [
        {
          id          = 1500011100682
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
        {
          id          = 1900002752745
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011100642
      value           = "esc_status_complete"
    },
    {
      child_fields = [
        {
          id          = 1900002752685
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1900002752725
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 15279369602839
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 22428196008983
          is_required = false
          required_on_statuses = {
            statuses = null
            type     = "NO_STATUSES"
          }
        },
      ]
      parent_field_id = 4612925021719
      value           = "tse"
    },
    {
      child_fields = [
        {
          id          = 17694655303447
          is_required = false
          required_on_statuses = {
            statuses = null
            type     = "NO_STATUSES"
          }
        },
        {
          id          = 1900002752685
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1900002752725
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 4612925021719
      value           = "ps"
    },
    {
      child_fields = [
        {
          id          = 1900002752725
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1900002752685
          is_required = true
          required_on_statuses = {
            statuses = ["pending", "solved"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 4612925021719
      value           = "tse_to_ps_handoff"
    },
    {
      child_fields = [
        {
          id          = 1900002752685
          is_required = true
          required_on_statuses = {
            statuses = ["open", "pending", "hold", "solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1900002752725
          is_required = true
          required_on_statuses = {
            statuses = ["open", "pending", "hold", "solved"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 4612925021719
      value           = "ps_to_tse_handoff"
    },
    {
      child_fields = [
        {
          id          = 16590742000407
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
        {
          id          = 16590625906839
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
        {
          id          = 16590703321367
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
        {
          id          = 16590745054871
          is_required = true
          required_on_statuses = {
            statuses = null
            type     = "ALL_STATUSES"
          }
        },
      ]
      parent_field_id = 21498705301783
      value           = "security_concern_request"
    },
    {
      child_fields = [
        {
          id          = 1500011100742
          is_required = true
          required_on_statuses = {
            statuses = ["solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1500011047301
          is_required = false
          required_on_statuses = {
            statuses = ["new"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011047201
      value           = "critical_-_1"
    },
    {
      child_fields = [
        {
          id          = 1500011100742
          is_required = true
          required_on_statuses = {
            statuses = ["solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1500011047301
          is_required = false
          required_on_statuses = {
            statuses = ["new"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011047201
      value           = "moderate_-_3"
    },
    {
      child_fields = [
        {
          id          = 1500011100742
          is_required = true
          required_on_statuses = {
            statuses = ["solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1500011047301
          is_required = false
          required_on_statuses = {
            statuses = ["new"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011047201
      value           = "severe_-_2"
    },
    {
      child_fields = [
        {
          id          = 1500011100742
          is_required = true
          required_on_statuses = {
            statuses = ["solved"]
            type     = "SOME_STATUSES"
          }
        },
        {
          id          = 1500011047301
          is_required = false
          required_on_statuses = {
            statuses = ["new"]
            type     = "SOME_STATUSES"
          }
        },
      ]
      parent_field_id = 1500011047201
      value           = "low_-_4"
    },
  ]
  default = true
  end_user_conditions = [
    {
      child_fields = [
        {
          id                   = 16590745054871
          is_required          = true
          required_on_statuses = null
        },
        {
          id                   = 16590625906839
          is_required          = true
          required_on_statuses = null
        },
        {
          id                   = 16590742000407
          is_required          = true
          required_on_statuses = null
        },
        {
          id                   = 16590703321367
          is_required          = true
          required_on_statuses = null
        },
      ]
      parent_field_id = 21498705301783
      value           = "security_concern_request"
    },
  ]
  end_user_display_name = "Support Request"
  end_user_visible      = true
  form_name             = "Support Request"
  position              = 0
  ticket_field_ids      = [1900002750765, 1900002750825, 1900002750845, 9193431385367, 15279369602839, 21498705301783, 1900002750725, 1900002750745, 23637257164183, 16590625906839, 16590742000407, 16590703321367, 16590745054871, 4612925021719, 1900002750785, 1900002750805, 17481352675991, 1900002752645, 1500011047201, 1500011047181, 1500011047301, 12670063125655, 1900002752685, 1900002752725, 1500011100722, 1900002752765, 1900006267625, 1500011100622, 1500011100642, 1900002752745, 1500011100682, 1500011100742, 1500011100702, 1500011100782, 5851640039703, 1500011047261, 1500011047241, 1900002752705, 1900002752665, 1500011100762, 1500011047381, 22428196008983, 1900003190465, 1500011887202, 1900004663765, 5454790585879, 5454804403095, 5454818191511, 9627693352599, 1500011100662, 16260018264087, 17694655303447, 22428586816023, 22901133515415, 23129171346839]
}