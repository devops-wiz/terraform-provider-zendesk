package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &SLAResource{}
var _ resource.ResourceWithImportState = &SLAResource{}
var _ resource.ResourceWithUpgradeState = &SLAResource{}

type SLAResource struct {
	client *zendesk.Client
}

func NewSLAResource() resource.Resource { return &SLAResource{} }

func (s *SLAResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sla_policy"
}

func (s *SLAResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = SLASchema
}

func (s *SLAResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*zendesk.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	s.client = client
}

func (s *SLAResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.SLAPolicyResourceModel{}, s.client.CreateSLAPolicy)
}

func (s *SLAResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.SLAPolicyResourceModel{}, s.client.GetSLAPolicy)
}

func (s *SLAResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	models.UpdateResource(ctx, request, response, &models.SLAPolicyResourceModel{}, s.client.UpdateSLAPolicy)
}

func (s *SLAResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.SLAPolicy](ctx, request, response, &models.SLAPolicyResourceModel{}, s.client.DeleteSLAPolicy)
}

func (s *SLAResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.SLAPolicyResourceModel{}, s.client.GetSLAPolicy)
}

func (s *SLAResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &SLASchemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData models.SLAPolicyResourceModelV0

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				priorAll := priorStateData.Filter.All
				priorAny := priorStateData.Filter.Any

				conditions := models.UpgradeConditionsV1(priorAll, priorAny)

				upgradedStateData := models.SLAPolicyResourceModel{
					ID:              priorStateData.ID,
					Title:           priorStateData.Title,
					Description:     priorStateData.Description,
					Position:        priorStateData.Position,
					Filter:          conditions,
					PolicyMetrics:   priorStateData.PolicyMetrics,
					MetricsSettings: priorStateData.MetricsSettings,
					CreatedAt:       priorStateData.CreatedAt,
					UpdatedAt:       priorStateData.UpdatedAt,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
