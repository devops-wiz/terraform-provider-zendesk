package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
)

type CredentialsResourceModel struct {
	HeaderName  types.String `tfsdk:"name"`
	HeaderValue types.String `tfsdk:"value"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	Token       types.String `tfsdk:"token"`
}

func (c CredentialsResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":     types.StringType,
		"value":    types.StringType,
		"username": types.StringType,
		"password": types.StringType,
		"token":    types.StringType,
	}
}

type AuthenticationResourceModel struct {
	Type        types.String `tfsdk:"type"`
	Credentials types.Object `tfsdk:"credentials"`
	AddPosition types.String `tfsdk:"add_position"`
}

func (a AuthenticationResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"type":         types.StringType,
		"credentials":  types.ObjectType{AttrTypes: CredentialsResourceModel{}.AttributeTypes()},
		"add_position": types.StringType,
	}
}

var _ ResourceTransform[zendesk.Webhook] = &WebhookResourceModel{}

// WebhookResourceModel describes the resource data model.
type WebhookResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Authentication types.Object `tfsdk:"authentication"`
	Description    types.String `tfsdk:"description"`
	Endpoint       types.String `tfsdk:"endpoint"`
	HttpMethod     types.String `tfsdk:"http_method"`
	CustomHeaders  types.Map    `tfsdk:"custom_headers"`
	RequestFormat  types.String `tfsdk:"request_format"`
	Status         types.String `tfsdk:"status"`
	Subscriptions  types.List   `tfsdk:"subscriptions"`
	Secret         types.String `tfsdk:"secret"`
	CreatedBy      types.String `tfsdk:"created_by"`
	CreatedAt      types.String `tfsdk:"created_at"`
	UpdatedBy      types.String `tfsdk:"updated_by"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
}

func (w *WebhookResourceModel) GetApiModelFromTfModel(ctx context.Context) (newWebhook zendesk.Webhook, diags diag.Diagnostics) {
	newWebhookAuthentication, diags := GetApiWebhookAuthenticationFromTf(ctx, w.Authentication)

	customHeaders := make(map[string]string, len(w.CustomHeaders.Elements()))

	for key, header := range w.CustomHeaders.Elements() {
		customHeaders[key] = header.(types.String).ValueString()
	}

	subs := make([]string, len(w.Subscriptions.Elements()))

	for i, sub := range w.Subscriptions.Elements() {
		subs[i] = sub.(types.String).ValueString()
	}

	if !reflect.DeepEqual(newWebhookAuthentication, &zendesk.WebhookAuthentication{}) {
		newWebhook = zendesk.Webhook{
			Name:           w.Name.ValueString(),
			Description:    w.Description.ValueString(),
			Authentication: newWebhookAuthentication,
			Endpoint:       w.Endpoint.ValueString(),
			HTTPMethod:     w.HttpMethod.ValueString(),
			CustomHeaders:  customHeaders,
			RequestFormat:  w.RequestFormat.ValueString(),
			Status:         w.Status.ValueString(),
			Subscriptions:  subs,
		}
	} else {
		newWebhook = zendesk.Webhook{
			Name:          w.Name.ValueString(),
			Description:   w.Description.ValueString(),
			Endpoint:      w.Endpoint.ValueString(),
			HTTPMethod:    w.HttpMethod.ValueString(),
			CustomHeaders: customHeaders,
			RequestFormat: w.RequestFormat.ValueString(),
			Status:        w.Status.ValueString(),
			Subscriptions: subs,
		}
	}

	return newWebhook, diags
}

func GetApiWebhookAuthenticationFromTf(ctx context.Context, authenticationObj types.Object) (apiAuthentication *zendesk.WebhookAuthentication, diags diag.Diagnostics) {
	var authentication AuthenticationResourceModel

	diags = authenticationObj.As(ctx, &authentication, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})

	if diags.HasError() {
		return apiAuthentication, diags
	}

	credentialsObj := authentication.Credentials

	var credentials CredentialsResourceModel

	diags.Append(credentialsObj.As(ctx, &credentials, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if diags.HasError() {
		return apiAuthentication, diags
	}

	apiCredentials := zendesk.WebhookCredentials{
		HeaderName:  credentials.HeaderName.ValueString(),
		HeaderValue: credentials.HeaderValue.ValueString(),
		Username:    credentials.Username.ValueString(),
		Password:    credentials.Password.ValueString(),
		Token:       credentials.Token.ValueString(),
	}

	if authenticationObj.IsUnknown() || authenticationObj.IsNull() {
		apiAuthentication = nil
	} else {

		apiAuthentication = &zendesk.WebhookAuthentication{
			Type:        authentication.Type.ValueString(),
			AddPosition: authentication.AddPosition.ValueString(),
			Data:        apiCredentials,
		}
	}

	return apiAuthentication, diags
}

func (w *WebhookResourceModel) GetTfModelFromApiModel(ctx context.Context, apiWebhook zendesk.Webhook) (diags diag.Diagnostics) {

	var newTfWebhookAuthentication types.Object

	if apiWebhook.Authentication != nil {
		newTfWebhookAuthentication, diags = getTfWebhookAuthenticationFromApi(ctx, apiWebhook.Authentication)
		if diags.HasError() {
			return diags
		}
	} else {
		newTfWebhookAuthentication = types.ObjectNull(AuthenticationResourceModel{}.AttributeTypes())
	}

	tfCustomHeaders := make(map[string]attr.Value, len(apiWebhook.CustomHeaders))

	var tfMap types.Map

	if len(apiWebhook.CustomHeaders) > 0 {

		for key, header := range apiWebhook.CustomHeaders {
			tfCustomHeaders[key] = types.StringValue(header)
		}

		tfMap, diags = types.MapValue(types.StringType, tfCustomHeaders)

		if diags.HasError() {
			return diags
		}
	} else {
		tfMap = types.MapNull(types.StringType)
	}

	tfSubscriptions := make([]attr.Value, len(apiWebhook.Subscriptions))

	for i, sub := range apiWebhook.Subscriptions {
		tfSubscriptions[i] = types.StringValue(sub)
	}

	tfList, diags := types.ListValue(types.StringType, tfSubscriptions)

	if diags.HasError() {
		return diags
	}

	*w = WebhookResourceModel{
		ID:             types.StringValue(apiWebhook.ID),
		Name:           types.StringValue(apiWebhook.Name),
		Description:    types.StringValue(apiWebhook.Description),
		Authentication: newTfWebhookAuthentication,
		Endpoint:       types.StringValue(apiWebhook.Endpoint),
		HttpMethod:     types.StringValue(apiWebhook.HTTPMethod),
		CustomHeaders:  tfMap,
		RequestFormat:  types.StringValue(apiWebhook.RequestFormat),
		Status:         types.StringValue(apiWebhook.Status),
		Subscriptions:  tfList,
		Secret:         types.StringValue(apiWebhook.SigningSecret.Secret),
		CreatedBy:      types.StringValue(apiWebhook.CreatedBy),
		CreatedAt:      types.StringValue(apiWebhook.CreatedAt.UTC().String()),
		UpdatedBy:      types.StringValue(apiWebhook.UpdatedBy),
		UpdatedAt:      types.StringValue(apiWebhook.UpdatedAt.UTC().String()),
	}

	return diags
}

func getTfWebhookAuthenticationFromApi(ctx context.Context, authentication *zendesk.WebhookAuthentication) (newTfAuthObject types.Object, diags diag.Diagnostics) {
	newTfCredentials := CredentialsResourceModel{}

	switch authentication.Type {
	case "api_key":
		newTfCredentials = CredentialsResourceModel{
			HeaderName:  types.StringValue(authentication.Data.HeaderName),
			HeaderValue: types.StringValue(authentication.Data.HeaderValue),
			Username:    types.StringNull(),
			Password:    types.StringNull(),
			Token:       types.StringNull(),
		}
	case "basic_auth":
		newTfCredentials = CredentialsResourceModel{
			HeaderName:  types.StringNull(),
			HeaderValue: types.StringNull(),
			Username:    types.StringValue(authentication.Data.Username),
			Password:    types.StringValue(authentication.Data.Password),
			Token:       types.StringNull(),
		}
	case "bearer_token":
		newTfCredentials = CredentialsResourceModel{
			HeaderName:  types.StringNull(),
			HeaderValue: types.StringNull(),
			Username:    types.StringNull(),
			Password:    types.StringNull(),
			Token:       types.StringValue(authentication.Data.Token),
		}
	}

	newTfCredsObject, diags := types.ObjectValueFrom(ctx, newTfCredentials.AttributeTypes(), newTfCredentials)

	if diags.HasError() {
		return newTfAuthObject, diags
	}

	newTfAuthentication := AuthenticationResourceModel{
		Type:        types.StringValue(authentication.Type),
		AddPosition: types.StringValue(authentication.AddPosition),
		Credentials: newTfCredsObject,
	}

	newTfAuthObject, diags = types.ObjectValueFrom(ctx, newTfAuthentication.AttributeTypes(), newTfAuthentication)

	if diags.HasError() {
		return newTfAuthObject, diags
	}

	return newTfAuthObject, diags
}
