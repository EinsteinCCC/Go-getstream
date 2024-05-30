package getstream

import (
	"context"
)

type CommonClient struct {
	client *Client
}

func NewCommonClient(client *Client) *CommonClient {
	return &CommonClient{
		client: client,
	}
}

func (c *CommonClient) GetApp(ctx context.Context) (GetApplicationResponse, error) {
	var result GetApplicationResponse
	err := MakeRequest[any, GetApplicationResponse, any](c.client, ctx, "GET", "/api/v2/app", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) UpdateApp(ctx context.Context, updateAppRequest UpdateAppRequest) (Response, error) {
	var result Response
	err := MakeRequest[UpdateAppRequest, Response, any](c.client, ctx, "PATCH", "/api/v2/app", nil, &updateAppRequest, &result, nil)
	return result, err
}

func (c *CommonClient) ListBlockLists(ctx context.Context) (ListBlockListResponse, error) {
	var result ListBlockListResponse
	err := MakeRequest[any, ListBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) CreateBlockList(ctx context.Context, createBlockListRequest CreateBlockListRequest) (Response, error) {
	var result Response
	err := MakeRequest[CreateBlockListRequest, Response, any](c.client, ctx, "POST", "/api/v2/blocklists", nil, &createBlockListRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeleteBlockList(ctx context.Context, name string) (Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) GetBlockList(ctx context.Context, name string) (GetBlockListResponse, error) {
	var result GetBlockListResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, GetBlockListResponse, any](c.client, ctx, "GET", "/api/v2/blocklists/{name}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) UpdateBlockList(ctx context.Context, name string, updateBlockListRequest UpdateBlockListRequest) (Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateBlockListRequest, Response, any](c.client, ctx, "PUT", "/api/v2/blocklists/{name}", nil, &updateBlockListRequest, &result, pathParams)
	return result, err
}

func (c *CommonClient) CheckPush(ctx context.Context, checkPushRequest CheckPushRequest) (CheckPushResponse, error) {
	var result CheckPushResponse
	err := MakeRequest[CheckPushRequest, CheckPushResponse, any](c.client, ctx, "POST", "/api/v2/check_push", nil, &checkPushRequest, &result, nil)
	return result, err
}

func (c *CommonClient) CheckSNS(ctx context.Context, checkSNSRequest CheckSNSRequest) (CheckSNSResponse, error) {
	var result CheckSNSResponse
	err := MakeRequest[CheckSNSRequest, CheckSNSResponse, any](c.client, ctx, "POST", "/api/v2/check_sns", nil, &checkSNSRequest, &result, nil)
	return result, err
}

func (c *CommonClient) CheckSQS(ctx context.Context, checkSQSRequest CheckSQSRequest) (CheckSQSResponse, error) {
	var result CheckSQSResponse
	err := MakeRequest[CheckSQSRequest, CheckSQSResponse, any](c.client, ctx, "POST", "/api/v2/check_sqs", nil, &checkSQSRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeleteDevice(ctx context.Context, id string, userId *string) (Response, error) {
	var result Response
	queryParams := map[string]interface{}{
		"id":      id,
		"user_id": userId,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/devices", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) ListDevices(ctx context.Context, userId *string) (ListDevicesResponse, error) {
	var result ListDevicesResponse
	queryParams := map[string]interface{}{
		"user_id": userId,
	}
	err := MakeRequest[any, ListDevicesResponse](c.client, ctx, "GET", "/api/v2/devices", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) CreateDevice(ctx context.Context, createDeviceRequest CreateDeviceRequest) (Response, error) {
	var result Response
	err := MakeRequest[CreateDeviceRequest, Response, any](c.client, ctx, "POST", "/api/v2/devices", nil, &createDeviceRequest, &result, nil)
	return result, err
}

func (c *CommonClient) ExportUsers(ctx context.Context, exportUsersRequest ExportUsersRequest) (ExportUsersResponse, error) {
	var result ExportUsersResponse
	err := MakeRequest[ExportUsersRequest, ExportUsersResponse, any](c.client, ctx, "POST", "/api/v2/export/users", nil, &exportUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) ListExternalStorage(ctx context.Context) (ListExternalStorageResponse, error) {
	var result ListExternalStorageResponse
	err := MakeRequest[any, ListExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) CreateExternalStorage(ctx context.Context, createExternalStorageRequest CreateExternalStorageRequest) (CreateExternalStorageResponse, error) {
	var result CreateExternalStorageResponse
	err := MakeRequest[CreateExternalStorageRequest, CreateExternalStorageResponse, any](c.client, ctx, "POST", "/api/v2/external_storage", nil, &createExternalStorageRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeleteExternalStorage(ctx context.Context, name string) (DeleteExternalStorageResponse, error) {
	var result DeleteExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, DeleteExternalStorageResponse, any](c.client, ctx, "DELETE", "/api/v2/external_storage/{name}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) UpdateExternalStorage(ctx context.Context, name string, updateExternalStorageRequest UpdateExternalStorageRequest) (UpdateExternalStorageResponse, error) {
	var result UpdateExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[UpdateExternalStorageRequest, UpdateExternalStorageResponse, any](c.client, ctx, "PUT", "/api/v2/external_storage/{name}", nil, &updateExternalStorageRequest, &result, pathParams)
	return result, err
}

func (c *CommonClient) CheckExternalStorage(ctx context.Context, name string) (CheckExternalStorageResponse, error) {
	var result CheckExternalStorageResponse
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, CheckExternalStorageResponse, any](c.client, ctx, "GET", "/api/v2/external_storage/{name}/check", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) CreateGuest(ctx context.Context, createGuestRequest CreateGuestRequest) (CreateGuestResponse, error) {
	var result CreateGuestResponse
	err := MakeRequest[CreateGuestRequest, CreateGuestResponse, any](c.client, ctx, "POST", "/api/v2/guest", nil, &createGuestRequest, &result, nil)
	return result, err
}

func (c *CommonClient) CreateImportURL(ctx context.Context, createImportURLRequest CreateImportURLRequest) (CreateImportURLResponse, error) {
	var result CreateImportURLResponse
	err := MakeRequest[CreateImportURLRequest, CreateImportURLResponse, any](c.client, ctx, "POST", "/api/v2/import_urls", nil, &createImportURLRequest, &result, nil)
	return result, err
}

func (c *CommonClient) ListImports(ctx context.Context) (ListImportsResponse, error) {
	var result ListImportsResponse
	err := MakeRequest[any, ListImportsResponse, any](c.client, ctx, "GET", "/api/v2/imports", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) CreateImport(ctx context.Context, createImportRequest CreateImportRequest) (CreateImportResponse, error) {
	var result CreateImportResponse
	err := MakeRequest[CreateImportRequest, CreateImportResponse, any](c.client, ctx, "POST", "/api/v2/imports", nil, &createImportRequest, &result, nil)
	return result, err
}

func (c *CommonClient) GetImport(ctx context.Context, id string) (GetImportResponse, error) {
	var result GetImportResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetImportResponse, any](c.client, ctx, "GET", "/api/v2/imports/{id}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) Unban(ctx context.Context, targetUserId string, channelCid *string, createdBy *string) (Response, error) {
	var result Response
	queryParams := map[string]interface{}{
		"target_user_id": targetUserId,
		"channel_cid":    channelCid,
		"created_by":     createdBy,
	}
	err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/moderation/ban", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) Ban(ctx context.Context, banRequest BanRequest) (Response, error) {
	var result Response
	err := MakeRequest[BanRequest, Response, any](c.client, ctx, "POST", "/api/v2/moderation/ban", nil, &banRequest, &result, nil)
	return result, err
}

func (c *CommonClient) Flag(ctx context.Context, flagRequest FlagRequest) (FlagResponse, error) {
	var result FlagResponse
	err := MakeRequest[FlagRequest, FlagResponse, any](c.client, ctx, "POST", "/api/v2/moderation/flag", nil, &flagRequest, &result, nil)
	return result, err
}

func (c *CommonClient) MuteUser(ctx context.Context, muteUserRequest MuteUserRequest) (MuteUserResponse, error) {
	var result MuteUserResponse
	err := MakeRequest[MuteUserRequest, MuteUserResponse, any](c.client, ctx, "POST", "/api/v2/moderation/mute", nil, &muteUserRequest, &result, nil)
	return result, err
}

func (c *CommonClient) UnmuteUser(ctx context.Context, unmuteUserRequest UnmuteUserRequest) (UnmuteResponse, error) {
	var result UnmuteResponse
	err := MakeRequest[UnmuteUserRequest, UnmuteResponse, any](c.client, ctx, "POST", "/api/v2/moderation/unmute", nil, &unmuteUserRequest, &result, nil)
	return result, err
}

func (c *CommonClient) GetOG(ctx context.Context, url string) (GetOGResponse, error) {
	var result GetOGResponse
	queryParams := map[string]interface{}{
		"url": url,
	}
	err := MakeRequest[any, GetOGResponse](c.client, ctx, "GET", "/api/v2/og", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) ListPermissions(ctx context.Context) (ListPermissionsResponse, error) {
	var result ListPermissionsResponse
	err := MakeRequest[any, ListPermissionsResponse, any](c.client, ctx, "GET", "/api/v2/permissions", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) GetPermission(ctx context.Context, id string) (GetCustomPermissionResponse, error) {
	var result GetCustomPermissionResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetCustomPermissionResponse, any](c.client, ctx, "GET", "/api/v2/permissions/{id}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) ListPushProviders(ctx context.Context) (ListPushProvidersResponse, error) {
	var result ListPushProvidersResponse
	err := MakeRequest[any, ListPushProvidersResponse, any](c.client, ctx, "GET", "/api/v2/push_providers", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) UpsertPushProvider(ctx context.Context, upsertPushProviderRequest UpsertPushProviderRequest) (UpsertPushProviderResponse, error) {
	var result UpsertPushProviderResponse
	err := MakeRequest[UpsertPushProviderRequest, UpsertPushProviderResponse, any](c.client, ctx, "POST", "/api/v2/push_providers", nil, &upsertPushProviderRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeletePushProvider(ctx context.Context, _type string, name string) (Response, error) {
	var result Response
	pathParams := map[string]string{
		"type": _type,
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/push_providers/{type}/{name}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) GetRateLimits(ctx context.Context, serverSide *bool, android *bool, ios *bool, web *bool, endpoints *string) (GetRateLimitsResponse, error) {
	var result GetRateLimitsResponse
	queryParams := map[string]interface{}{
		"server_side": serverSide,
		"android":     android,
		"ios":         ios,
		"web":         web,
		"endpoints":   endpoints,
	}
	err := MakeRequest[any, GetRateLimitsResponse](c.client, ctx, "GET", "/api/v2/rate_limits", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) ListRoles(ctx context.Context) (ListRolesResponse, error) {
	var result ListRolesResponse
	err := MakeRequest[any, ListRolesResponse, any](c.client, ctx, "GET", "/api/v2/roles", nil, nil, &result, nil)
	return result, err
}

func (c *CommonClient) CreateRole(ctx context.Context, createRoleRequest CreateRoleRequest) (CreateRoleResponse, error) {
	var result CreateRoleResponse
	err := MakeRequest[CreateRoleRequest, CreateRoleResponse, any](c.client, ctx, "POST", "/api/v2/roles", nil, &createRoleRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeleteRole(ctx context.Context, name string) (Response, error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/roles/{name}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) GetTask(ctx context.Context, id string) (GetTaskResponse, error) {
	var result GetTaskResponse
	pathParams := map[string]string{
		"id": id,
	}
	err := MakeRequest[any, GetTaskResponse, any](c.client, ctx, "GET", "/api/v2/tasks/{id}", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) QueryUsers(ctx context.Context, payload *QueryUsersPayload) (QueryUsersResponse, error) {
	var result QueryUsersResponse
	queryParams := map[string]interface{}{
		"payload": payload,
	}
	err := MakeRequest[any, QueryUsersResponse](c.client, ctx, "GET", "/api/v2/users", queryParams, nil, &result, nil)
	return result, err
}

func (c *CommonClient) UpdateUsersPartial(ctx context.Context, updateUsersPartialRequest UpdateUsersPartialRequest) (UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest[UpdateUsersPartialRequest, UpdateUsersResponse, any](c.client, ctx, "PATCH", "/api/v2/users", nil, &updateUsersPartialRequest, &result, nil)
	return result, err
}

func (c *CommonClient) UpdateUsers(ctx context.Context, updateUsersRequest UpdateUsersRequest) (UpdateUsersResponse, error) {
	var result UpdateUsersResponse
	err := MakeRequest[UpdateUsersRequest, UpdateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users", nil, &updateUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeactivateUsers(ctx context.Context, deactivateUsersRequest DeactivateUsersRequest) (DeactivateUsersResponse, error) {
	var result DeactivateUsersResponse
	err := MakeRequest[DeactivateUsersRequest, DeactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/deactivate", nil, &deactivateUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeleteUsers(ctx context.Context, deleteUsersRequest DeleteUsersRequest) (DeleteUsersResponse, error) {
	var result DeleteUsersResponse
	err := MakeRequest[DeleteUsersRequest, DeleteUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/delete", nil, &deleteUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) ReactivateUsers(ctx context.Context, reactivateUsersRequest ReactivateUsersRequest) (ReactivateUsersResponse, error) {
	var result ReactivateUsersResponse
	err := MakeRequest[ReactivateUsersRequest, ReactivateUsersResponse, any](c.client, ctx, "POST", "/api/v2/users/reactivate", nil, &reactivateUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) RestoreUsers(ctx context.Context, restoreUsersRequest RestoreUsersRequest) (Response, error) {
	var result Response
	err := MakeRequest[RestoreUsersRequest, Response, any](c.client, ctx, "POST", "/api/v2/users/restore", nil, &restoreUsersRequest, &result, nil)
	return result, err
}

func (c *CommonClient) DeactivateUser(ctx context.Context, userId string, deactivateUserRequest DeactivateUserRequest) (DeactivateUserResponse, error) {
	var result DeactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[DeactivateUserRequest, DeactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/deactivate", nil, &deactivateUserRequest, &result, pathParams)
	return result, err
}

func (c *CommonClient) ExportUser(ctx context.Context, userId string) (ExportUserResponse, error) {
	var result ExportUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[any, ExportUserResponse, any](c.client, ctx, "GET", "/api/v2/users/{user_id}/export", nil, nil, &result, pathParams)
	return result, err
}

func (c *CommonClient) ReactivateUser(ctx context.Context, userId string, reactivateUserRequest ReactivateUserRequest) (ReactivateUserResponse, error) {
	var result ReactivateUserResponse
	pathParams := map[string]string{
		"user_id": userId,
	}
	err := MakeRequest[ReactivateUserRequest, ReactivateUserResponse, any](c.client, ctx, "POST", "/api/v2/users/{user_id}/reactivate", nil, &reactivateUserRequest, &result, pathParams)
	return result, err
}
