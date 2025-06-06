package provider

import "fmt"

var (
	tmplPath            = "testdata/templates"
	ticketFieldTmpl     = "ticket_field.tf.tmpl"
	ticketFieldTmplPath = fmt.Sprintf("%s/%s", tmplPath, ticketFieldTmpl)
)
