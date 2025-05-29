package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type LocalDatasource struct {
	client *zendesk.Client
}

func NewLocaleDatasource() datasource.DataSource {
	return &LocalDatasource{}
}

func (l *LocalDatasource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_locale"
}

func (l *LocalDatasource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*zendesk.Client)

	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)

		return
	}

	l.client = client
}

func (l *LocalDatasource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = LocaleSchema
}

func (l *LocalDatasource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var config models.LocaleDatasourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &config)...)

	if response.Diagnostics.HasError() {
		return
	}

	locales, err := l.client.GetLocales(ctx)

	if err != nil {
		response.Diagnostics.AddError("Failed to read locales", fmt.Sprintf("Error getting locales: %s", err))
		return
	}

	locale := utils.SliceFilter(locales, func(locale zendesk.Locale) bool {
		return locale.Locale == config.Code.ValueString()
	})

	if len(locale) == 0 {
		response.Diagnostics.AddError("Failed to read locales", fmt.Sprintf("Locale %s not found", config.Code))
	}

	if len(locale) > 1 {
		response.Diagnostics.AddError("Failed to read locales", fmt.Sprintf("Multiple locales returned for code %s", config.Code))
	}

	config.Locale = &models.LocaleModel{}

	diags := config.Locale.GetTfModelFromApiModel(ctx, locale[0])

	response.Diagnostics.Append(diags...)

	if diags.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, config)...)

}
