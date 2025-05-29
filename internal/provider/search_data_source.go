package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &SearchDatasource{}
var _ datasource.DataSourceWithConfigure = &SearchDatasource{}

type SearchDatasource struct {
	client *zendesk.Client
}

func (o *SearchDatasource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

	o.client = client
}

func NewSearchDatasource() datasource.DataSource {
	return &SearchDatasource{}
}

func (o *SearchDatasource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_search"
}

func (o *SearchDatasource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = SearchSchema
}

func (o *SearchDatasource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var config models.SearchDatasourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &config)...)

	if response.Diagnostics.HasError() {
		return
	}

	searchOptions := config.GetApiQueryOptionsFromTf()

	searchResults, _, err := o.client.Search(ctx, &searchOptions)

	if err != nil {
		response.Diagnostics.AddError("Error reading search API", fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	response.Diagnostics.Append(config.GetTfModelFromApiModel(ctx, searchResults)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, config)...)

}
