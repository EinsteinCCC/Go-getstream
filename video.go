package getstream

import (
	"context"
)

type VideoClient struct {
	client *Client
}

func NewVideoClient(client *Client) *VideoClient {
	return &VideoClient{
		client: client,
	}
}
func (c *VideoClient) QueryCallMembers(ctx context.Context, queryCallMembersRequest QueryCallMembersRequest) (QueryCallMembersResponse, error) {
	var result QueryCallMembersResponse
	err := MakeRequest[QueryCallMembersRequest, QueryCallMembersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/members", nil, &queryCallMembersRequest, &result)
	return result, err
}
func (c *VideoClient) QueryCallStats(ctx context.Context, queryCallStatsRequest QueryCallStatsRequest) (QueryCallStatsResponse, error) {
	var result QueryCallStatsResponse
	err := MakeRequest[QueryCallStatsRequest, QueryCallStatsResponse, any](c.client, ctx, "POST", "/api/v2/video/call/stats", nil, &queryCallStatsRequest, &result)
	return result, err
}
func (c *VideoClient) GetCall(ctx context.Context, _type string, id string, membersLimit *int, ring *bool, notify *bool) (GetCallResponse, error) {
	var result GetCallResponse
	pathParams := []string{_type, id}
	queryParams := map[string]interface{}{
		"members_limit": membersLimit,
		"ring":          ring,
		"notify":        notify,
	}
	err := MakeRequest[any, GetCallResponse](c.client, ctx, "GET", "/api/v2/video/call/{_type}/{id}", queryParams, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) UpdateCall(ctx context.Context, _type string, id string, updateCallRequest UpdateCallRequest) (UpdateCallResponse, error) {
	var result UpdateCallResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UpdateCallRequest, UpdateCallResponse, any](c.client, ctx, "PATCH", "/api/v2/video/call/{_type}/{id}", nil, &updateCallRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) GetOrCreateCall(ctx context.Context, _type string, id string, getOrCreateCallRequest GetOrCreateCallRequest) (GetOrCreateCallResponse, error) {
	var result GetOrCreateCallResponse
	pathParams := []string{_type, id}
	err := MakeRequest[GetOrCreateCallRequest, GetOrCreateCallResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}", nil, &getOrCreateCallRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) BlockUser(ctx context.Context, _type string, id string, blockUserRequest BlockUserRequest) (BlockUserResponse, error) {
	var result BlockUserResponse
	pathParams := []string{_type, id}
	err := MakeRequest[BlockUserRequest, BlockUserResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/block", nil, &blockUserRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) SendCallEvent(ctx context.Context, _type string, id string, sendCallEventRequest SendCallEventRequest) (SendCallEventResponse, error) {
	var result SendCallEventResponse
	pathParams := []string{_type, id}
	err := MakeRequest[SendCallEventRequest, SendCallEventResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/event", nil, &sendCallEventRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) CollectUserFeedback(ctx context.Context, _type string, id string, session string, collectUserFeedbackRequest CollectUserFeedbackRequest) (CollectUserFeedbackResponse, error) {
	var result CollectUserFeedbackResponse
	pathParams := []string{_type, id, session}
	err := MakeRequest[CollectUserFeedbackRequest, CollectUserFeedbackResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/feedback/{session}", nil, &collectUserFeedbackRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) GoLive(ctx context.Context, _type string, id string, goLiveRequest GoLiveRequest) (GoLiveResponse, error) {
	var result GoLiveResponse
	pathParams := []string{_type, id}
	err := MakeRequest[GoLiveRequest, GoLiveResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/go_live", nil, &goLiveRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) EndCall(ctx context.Context, _type string, id string) (EndCallResponse, error) {
	var result EndCallResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, EndCallResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/mark_ended", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) UpdateCallMembers(ctx context.Context, _type string, id string, updateCallMembersRequest UpdateCallMembersRequest) (UpdateCallMembersResponse, error) {
	var result UpdateCallMembersResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UpdateCallMembersRequest, UpdateCallMembersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/members", nil, &updateCallMembersRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) MuteUsers(ctx context.Context, _type string, id string, muteUsersRequest MuteUsersRequest) (MuteUsersResponse, error) {
	var result MuteUsersResponse
	pathParams := []string{_type, id}
	err := MakeRequest[MuteUsersRequest, MuteUsersResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/mute_users", nil, &muteUsersRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) VideoPin(ctx context.Context, _type string, id string, pinRequest PinRequest) (PinResponse, error) {
	var result PinResponse
	pathParams := []string{_type, id}
	err := MakeRequest[PinRequest, PinResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/pin", nil, &pinRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) ListRecordings(ctx context.Context, _type string, id string) (ListRecordingsResponse, error) {
	var result ListRecordingsResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, ListRecordingsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{_type}/{id}/recordings", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StartHLSBroadcasting(ctx context.Context, _type string, id string) (StartHLSBroadcastingResponse, error) {
	var result StartHLSBroadcastingResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, StartHLSBroadcastingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/start_broadcasting", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StartRecording(ctx context.Context, _type string, id string, startRecordingRequest StartRecordingRequest) (StartRecordingResponse, error) {
	var result StartRecordingResponse
	pathParams := []string{_type, id}
	err := MakeRequest[StartRecordingRequest, StartRecordingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/start_recording", nil, &startRecordingRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StartTranscription(ctx context.Context, _type string, id string, startTranscriptionRequest StartTranscriptionRequest) (StartTranscriptionResponse, error) {
	var result StartTranscriptionResponse
	pathParams := []string{_type, id}
	err := MakeRequest[StartTranscriptionRequest, StartTranscriptionResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/start_transcription", nil, &startTranscriptionRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) GetCallStats(ctx context.Context, _type string, id string, session string) (GetCallStatsResponse, error) {
	var result GetCallStatsResponse
	pathParams := []string{_type, id, session}
	err := MakeRequest[any, GetCallStatsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{_type}/{id}/stats/{session}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StopHLSBroadcasting(ctx context.Context, _type string, id string) (StopHLSBroadcastingResponse, error) {
	var result StopHLSBroadcastingResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, StopHLSBroadcastingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/stop_broadcasting", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StopLive(ctx context.Context, _type string, id string) (StopLiveResponse, error) {
	var result StopLiveResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, StopLiveResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/stop_live", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StopRecording(ctx context.Context, _type string, id string) (StopRecordingResponse, error) {
	var result StopRecordingResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, StopRecordingResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/stop_recording", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) StopTranscription(ctx context.Context, _type string, id string) (StopTranscriptionResponse, error) {
	var result StopTranscriptionResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, StopTranscriptionResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/stop_transcription", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) ListTranscriptions(ctx context.Context, _type string, id string) (ListTranscriptionsResponse, error) {
	var result ListTranscriptionsResponse
	pathParams := []string{_type, id}
	err := MakeRequest[any, ListTranscriptionsResponse, any](c.client, ctx, "GET", "/api/v2/video/call/{_type}/{id}/transcriptions", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) UnblockUser(ctx context.Context, _type string, id string, unblockUserRequest UnblockUserRequest) (UnblockUserResponse, error) {
	var result UnblockUserResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UnblockUserRequest, UnblockUserResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/unblock", nil, &unblockUserRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) VideoUnpin(ctx context.Context, _type string, id string, unpinRequest UnpinRequest) (UnpinResponse, error) {
	var result UnpinResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UnpinRequest, UnpinResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/unpin", nil, &unpinRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) UpdateUserPermissions(ctx context.Context, _type string, id string, updateUserPermissionsRequest UpdateUserPermissionsRequest) (UpdateUserPermissionsResponse, error) {
	var result UpdateUserPermissionsResponse
	pathParams := []string{_type, id}
	err := MakeRequest[UpdateUserPermissionsRequest, UpdateUserPermissionsResponse, any](c.client, ctx, "POST", "/api/v2/video/call/{_type}/{id}/user_permissions", nil, &updateUserPermissionsRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) DeleteRecording(ctx context.Context, _type string, id string, session string, filename string) (DeleteRecordingResponse, error) {
	var result DeleteRecordingResponse
	pathParams := []string{_type, id, session, filename}
	err := MakeRequest[any, DeleteRecordingResponse, any](c.client, ctx, "DELETE", "/api/v2/video/call/{_type}/{id}/{session}/recordings/{filename}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) DeleteTranscription(ctx context.Context, _type string, id string, session string, filename string) (DeleteTranscriptionResponse, error) {
	var result DeleteTranscriptionResponse
	pathParams := []string{_type, id, session, filename}
	err := MakeRequest[any, DeleteTranscriptionResponse, any](c.client, ctx, "DELETE", "/api/v2/video/call/{_type}/{id}/{session}/transcriptions/{filename}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) QueryCalls(ctx context.Context, queryCallsRequest QueryCallsRequest) (QueryCallsResponse, error) {
	var result QueryCallsResponse
	err := MakeRequest[QueryCallsRequest, QueryCallsResponse, any](c.client, ctx, "POST", "/api/v2/video/calls", nil, &queryCallsRequest, &result)
	return result, err
}
func (c *VideoClient) ListCallTypes(ctx context.Context) (ListCallTypeResponse, error) {
	var result ListCallTypeResponse
	err := MakeRequest[any, ListCallTypeResponse, any](c.client, ctx, "GET", "/api/v2/video/calltypes", nil, nil, &result)
	return result, err
}
func (c *VideoClient) CreateCallType(ctx context.Context, createCallTypeRequest CreateCallTypeRequest) (CreateCallTypeResponse, error) {
	var result CreateCallTypeResponse
	err := MakeRequest[CreateCallTypeRequest, CreateCallTypeResponse, any](c.client, ctx, "POST", "/api/v2/video/calltypes", nil, &createCallTypeRequest, &result)
	return result, err
}
func (c *VideoClient) DeleteCallType(ctx context.Context, name string) (Response, error) {
	var result Response
	pathParams := []string{name}
	err := MakeRequest[any, Response, any](c.client, ctx, "DELETE", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) GetCallType(ctx context.Context, name string) (GetCallTypeResponse, error) {
	var result GetCallTypeResponse
	pathParams := []string{name}
	err := MakeRequest[any, GetCallTypeResponse, any](c.client, ctx, "GET", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams...)
	return result, err
}
func (c *VideoClient) UpdateCallType(ctx context.Context, name string, updateCallTypeRequest UpdateCallTypeRequest) (UpdateCallTypeResponse, error) {
	var result UpdateCallTypeResponse
	pathParams := []string{name}
	err := MakeRequest[UpdateCallTypeRequest, UpdateCallTypeResponse, any](c.client, ctx, "PUT", "/api/v2/video/calltypes/{name}", nil, &updateCallTypeRequest, &result, pathParams...)
	return result, err
}
func (c *VideoClient) GetEdges(ctx context.Context) (GetEdgesResponse, error) {
	var result GetEdgesResponse
	err := MakeRequest[any, GetEdgesResponse, any](c.client, ctx, "GET", "/api/v2/video/edges", nil, nil, &result)
	return result, err
}
