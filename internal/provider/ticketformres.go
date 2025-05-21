package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithImportState = &TicketFormResource{}
var _ resource.ResourceWithConfigure = &TicketFormResource{}
var _ resource.ResourceWithUpgradeState = &TicketFormResource{}

type TicketFormResource struct {
	client *zendesk.Client
}

func NewTicketFormResource() resource.Resource {
	return &TicketFormResource{}
}

func (t *TicketFormResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ticket_form"
}

func (t *TicketFormResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = TicketFormSchema
}

// Configure adds the provider configured client to the resource.
func (t *TicketFormResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*zendesk.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *zendesk.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	t.client = client
}

func (t *TicketFormResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.TicketFormResourceModel{}, t.client.CreateTicketForm)
}

func (t *TicketFormResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.TicketFormResourceModel{}, t.client.GetTicketForm)
}

func (t *TicketFormResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	models.UpdateResource(ctx, request, response, &models.TicketFormResourceModel{}, t.client.UpdateTicketForm)
}

func (t *TicketFormResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.TicketForm](ctx, request, response, &models.TicketFormResourceModel{}, t.client.DeleteTicketForm)
}

func (t *TicketFormResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.TicketFormResourceModel{}, t.client.GetTicketForm)
}

func (t *TicketFormResource) UpgradeState(_ context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &TicketFormSchemaV0,
			StateUpgrader: func(ctx context.Context, request resource.UpgradeStateRequest, response *resource.UpgradeStateResponse) {
				var priorState models.TicketFormResourceModelV0

				response.Diagnostics.Append(request.State.Get(ctx, &priorState)...)

				if response.Diagnostics.HasError() {
					return
				}

				priorAgentSet := make([]models.FormConditionsV0, len(priorState.AgentConditions.Elements()))

				response.Diagnostics.Append(priorState.AgentConditions.ElementsAs(ctx, &priorAgentSet, true)...)

				if response.Diagnostics.HasError() {
					return
				}

				newAgentMap, tempDiag := models.UpgradeFormConditionsV1(ctx, priorAgentSet)

				response.Diagnostics.Append(tempDiag...)

				if response.Diagnostics.HasError() {
					return
				}

				newAgentMapTf, tempDiag := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: models.FormConditionsSet{}.AttributeTypes()}, newAgentMap)

				response.Diagnostics.Append(tempDiag...)

				if response.Diagnostics.HasError() {
					return
				}

				priorEndUserSet := make([]models.FormConditionsV0, len(priorState.EndUserConditions.Elements()))

				response.Diagnostics.Append(priorState.EndUserConditions.ElementsAs(ctx, &priorEndUserSet, true)...)

				if response.Diagnostics.HasError() {
					return
				}

				newEndUserMap, tempDiag := models.UpgradeFormConditionsV1(ctx, priorEndUserSet)

				response.Diagnostics.Append(tempDiag...)

				if response.Diagnostics.HasError() {
					return
				}

				newEndUserMapTf, tempDiag := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: models.FormConditionsSet{}.AttributeTypes()}, newEndUserMap)

				response.Diagnostics.Append(tempDiag...)

				if response.Diagnostics.HasError() {
					return
				}

				newState := models.TicketFormResourceModel{
					ID:                priorState.ID,
					Name:              priorState.Name,
					DisplayName:       priorState.DisplayName,
					TicketFieldIds:    priorState.TicketFieldIds,
					AgentConditions:   newAgentMapTf,
					EndUserConditions: newEndUserMapTf,
					Active:            priorState.Active,
					Position:          priorState.Position,
					Default:           priorState.Default,
					EndUserVisible:    priorState.EndUserVisible,
					CreatedAt:         priorState.CreatedAt,
					UpdatedAt:         priorState.UpdatedAt,
					Url:               priorState.Url,
				}

				response.Diagnostics.Append(response.State.Set(ctx, newState)...)

			},
		},
	}
}
