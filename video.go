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

// Query call members with filter query
//
// Required permissions:
// - ReadCall
func (c *VideoClient) QueryCallMembers(ctx context.Context, request *QueryCallMembersRequest) (*StreamResponse[QueryCallMembersResponse], error) {
	var result QueryCallMembersResponse
	res, err := MakeRequest[QueryCallMembersRequest, QueryCallMembersResponse](c.client, ctx, "POST", "/api/v2/video/call/members", nil, request, &result, nil)
	return res, err
}

// Required permissions:
// - ReadCallStats
func (c *VideoClient) QueryCallStats(ctx context.Context, request *QueryCallStatsRequest) (*StreamResponse[QueryCallStatsResponse], error) {
	var result QueryCallStatsResponse
	res, err := MakeRequest[QueryCallStatsRequest, QueryCallStatsResponse](c.client, ctx, "POST", "/api/v2/video/call/stats", nil, request, &result, nil)
	return res, err
}

// Required permissions:
// - ReadCall
func (c *VideoClient) GetCall(ctx context.Context, _type string, id string, request *GetCallRequest) (*StreamResponse[GetCallResponse], error) {
	var result GetCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	params := extractQueryParams(request)
	res, err := MakeRequest[any, GetCallResponse](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}", params, nil, &result, pathParams)
	return res, err
}

// Sends events:
// - call.updated
//
// Required permissions:
// - UpdateCall
func (c *VideoClient) UpdateCall(ctx context.Context, _type string, id string, request *UpdateCallRequest) (*StreamResponse[UpdateCallResponse], error) {
	var result UpdateCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UpdateCallRequest, UpdateCallResponse](c.client, ctx, "PATCH", "/api/v2/video/call/{type}/{id}", nil, request, &result, pathParams)
	return res, err
}

// Gets or creates a new call
//
// Sends events:
// - call.created
// - call.notification
// - call.ring
//
// Required permissions:
// - CreateCall
// - ReadCall
// - UpdateCallSettings
func (c *VideoClient) GetOrCreateCall(ctx context.Context, _type string, id string, request *GetOrCreateCallRequest) (*StreamResponse[GetOrCreateCallResponse], error) {
	var result GetOrCreateCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[GetOrCreateCallRequest, GetOrCreateCallResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}", nil, request, &result, pathParams)
	return res, err
}

// Block a user, preventing them from joining the call until they are unblocked.
//
// Sends events:
// - call.blocked_user
//
// Required permissions:
// - BlockUser
func (c *VideoClient) BlockUser(ctx context.Context, _type string, id string, request *BlockUserRequest) (*StreamResponse[BlockUserResponse], error) {
	var result BlockUserResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[BlockUserRequest, BlockUserResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/block", nil, request, &result, pathParams)
	return res, err
}

// Sends events:
// - call.deleted
//
// Required permissions:
// - DeleteCall
func (c *VideoClient) DeleteCall(ctx context.Context, _type string, id string, request *DeleteCallRequest) (*StreamResponse[DeleteCallResponse], error) {
	var result DeleteCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[DeleteCallRequest, DeleteCallResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/delete", nil, request, &result, pathParams)
	return res, err
}

// Sends custom event to the call
//
// Sends events:
// - custom
//
// Required permissions:
// - SendEvent
func (c *VideoClient) SendCallEvent(ctx context.Context, _type string, id string, request *SendCallEventRequest) (*StreamResponse[SendCallEventResponse], error) {
	var result SendCallEventResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[SendCallEventRequest, SendCallEventResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/event", nil, request, &result, pathParams)
	return res, err
}

// Required permissions:
// - JoinCall
func (c *VideoClient) CollectUserFeedback(ctx context.Context, _type string, id string, session string, request *CollectUserFeedbackRequest) (*StreamResponse[CollectUserFeedbackResponse], error) {
	var result CollectUserFeedbackResponse
	pathParams := map[string]string{
		"type":    _type,
		"id":      id,
		"session": session,
	}
	res, err := MakeRequest[CollectUserFeedbackRequest, CollectUserFeedbackResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/feedback/{session}", nil, request, &result, pathParams)
	return res, err
}

// Sends events:
// - call.live_started
//
// Required permissions:
// - UpdateCall
func (c *VideoClient) GoLive(ctx context.Context, _type string, id string, request *GoLiveRequest) (*StreamResponse[GoLiveResponse], error) {
	var result GoLiveResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[GoLiveRequest, GoLiveResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/go_live", nil, request, &result, pathParams)
	return res, err
}

// Sends events:
// - call.ended
//
// Required permissions:
// - EndCall
func (c *VideoClient) EndCall(ctx context.Context, _type string, id string) (*StreamResponse[EndCallResponse], error) {
	var result EndCallResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, EndCallResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mark_ended", nil, nil, &result, pathParams)
	return res, err
}

// Sends events:
// - call.member_added
// - call.member_removed
// - call.member_updated
//
// Required permissions:
// - RemoveCallMember
// - UpdateCallMember
// - UpdateCallMemberRole
func (c *VideoClient) UpdateCallMembers(ctx context.Context, _type string, id string, request *UpdateCallMembersRequest) (*StreamResponse[UpdateCallMembersResponse], error) {
	var result UpdateCallMembersResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UpdateCallMembersRequest, UpdateCallMembersResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/members", nil, request, &result, pathParams)
	return res, err
}

// Mutes users in a call
//
// Required permissions:
// - MuteUsers
func (c *VideoClient) MuteUsers(ctx context.Context, _type string, id string, request *MuteUsersRequest) (*StreamResponse[MuteUsersResponse], error) {
	var result MuteUsersResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[MuteUsersRequest, MuteUsersResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/mute_users", nil, request, &result, pathParams)
	return res, err
}

// Pins a track for all users in the call.
//
// Required permissions:
// - PinCallTrack
func (c *VideoClient) VideoPin(ctx context.Context, _type string, id string, request *PinRequest) (*StreamResponse[PinResponse], error) {
	var result PinResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[PinRequest, PinResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/pin", nil, request, &result, pathParams)
	return res, err
}

// Lists recordings
//
// Required permissions:
// - ListRecordings
func (c *VideoClient) ListRecordings(ctx context.Context, _type string, id string) (*StreamResponse[ListRecordingsResponse], error) {
	var result ListRecordingsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, ListRecordingsResponse](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/recordings", nil, nil, &result, pathParams)
	return res, err
}

// Starts RTMP broadcasts for the provided RTMP destinations
//
// Required permissions:
// - StartBroadcasting
func (c *VideoClient) StartRTMPBroadcast(ctx context.Context, _type string, id string, request *StartRTMPBroadcastsRequest) (*StreamResponse[StartRTMPBroadcastsResponse], error) {
	var result StartRTMPBroadcastsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[StartRTMPBroadcastsRequest, StartRTMPBroadcastsResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/rtmp_broadcasts", nil, request, &result, pathParams)
	return res, err
}

// Stop all RTMP broadcasts for the provided call
//
// Required permissions:
// - StopBroadcasting
func (c *VideoClient) StopAllRTMPBroadcasts(ctx context.Context, _type string, id string) (*StreamResponse[StopAllRTMPBroadcastsResponse], error) {
	var result StopAllRTMPBroadcastsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StopAllRTMPBroadcastsResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/rtmp_broadcasts/stop", nil, nil, &result, pathParams)
	return res, err
}

// Stop RTMP broadcasts for the provided RTMP destinations
//
// Required permissions:
// - StopBroadcasting
func (c *VideoClient) StopRTMPBroadcast(ctx context.Context, _type string, id string, name string, request *StopRTMPBroadcastsRequest) (*StreamResponse[StopRTMPBroadcastsResponse], error) {
	var result StopRTMPBroadcastsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
		"name": name,
	}
	res, err := MakeRequest[StopRTMPBroadcastsRequest, StopRTMPBroadcastsResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/rtmp_broadcasts/{name}/stop", nil, request, &result, pathParams)
	return res, err
}

// Starts HLS broadcasting
//
// Required permissions:
// - StartBroadcasting
func (c *VideoClient) StartHLSBroadcasting(ctx context.Context, _type string, id string) (*StreamResponse[StartHLSBroadcastingResponse], error) {
	var result StartHLSBroadcastingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StartHLSBroadcastingResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_broadcasting", nil, nil, &result, pathParams)
	return res, err
}

// Starts recording
//
// Sends events:
// - call.recording_started
//
// Required permissions:
// - StartRecording
func (c *VideoClient) StartRecording(ctx context.Context, _type string, id string, request *StartRecordingRequest) (*StreamResponse[StartRecordingResponse], error) {
	var result StartRecordingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[StartRecordingRequest, StartRecordingResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_recording", nil, request, &result, pathParams)
	return res, err
}

// Starts transcription
//
// Required permissions:
// - StartTranscription
func (c *VideoClient) StartTranscription(ctx context.Context, _type string, id string, request *StartTranscriptionRequest) (*StreamResponse[StartTranscriptionResponse], error) {
	var result StartTranscriptionResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[StartTranscriptionRequest, StartTranscriptionResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/start_transcription", nil, request, &result, pathParams)
	return res, err
}

// Required permissions:
// - ReadCallStats
func (c *VideoClient) GetCallStats(ctx context.Context, _type string, id string, session string) (*StreamResponse[GetCallStatsResponse], error) {
	var result GetCallStatsResponse
	pathParams := map[string]string{
		"type":    _type,
		"id":      id,
		"session": session,
	}
	res, err := MakeRequest[any, GetCallStatsResponse](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/stats/{session}", nil, nil, &result, pathParams)
	return res, err
}

// Stops HLS broadcasting
//
// Required permissions:
// - StopBroadcasting
func (c *VideoClient) StopHLSBroadcasting(ctx context.Context, _type string, id string) (*StreamResponse[StopHLSBroadcastingResponse], error) {
	var result StopHLSBroadcastingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StopHLSBroadcastingResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_broadcasting", nil, nil, &result, pathParams)
	return res, err
}

// Sends events:
// - call.updated
//
// Required permissions:
// - UpdateCall
func (c *VideoClient) StopLive(ctx context.Context, _type string, id string) (*StreamResponse[StopLiveResponse], error) {
	var result StopLiveResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StopLiveResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_live", nil, nil, &result, pathParams)
	return res, err
}

// Stops recording
//
// Sends events:
// - call.recording_stopped
//
// Required permissions:
// - StopRecording
func (c *VideoClient) StopRecording(ctx context.Context, _type string, id string) (*StreamResponse[StopRecordingResponse], error) {
	var result StopRecordingResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StopRecordingResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_recording", nil, nil, &result, pathParams)
	return res, err
}

// Stops transcription
//
// Sends events:
// - call.transcription_stopped
//
// Required permissions:
// - StopTranscription
func (c *VideoClient) StopTranscription(ctx context.Context, _type string, id string) (*StreamResponse[StopTranscriptionResponse], error) {
	var result StopTranscriptionResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, StopTranscriptionResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/stop_transcription", nil, nil, &result, pathParams)
	return res, err
}

// Lists transcriptions
//
// Required permissions:
// - ListTranscriptions
func (c *VideoClient) ListTranscriptions(ctx context.Context, _type string, id string) (*StreamResponse[ListTranscriptionsResponse], error) {
	var result ListTranscriptionsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[any, ListTranscriptionsResponse](c.client, ctx, "GET", "/api/v2/video/call/{type}/{id}/transcriptions", nil, nil, &result, pathParams)
	return res, err
}

// Removes the block for a user on a call. The user will be able to join the call again.
//
// Sends events:
// - call.unblocked_user
//
// Required permissions:
// - BlockUser
func (c *VideoClient) UnblockUser(ctx context.Context, _type string, id string, request *UnblockUserRequest) (*StreamResponse[UnblockUserResponse], error) {
	var result UnblockUserResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UnblockUserRequest, UnblockUserResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unblock", nil, request, &result, pathParams)
	return res, err
}

// Unpins a track for all users in the call.
//
// Required permissions:
// - PinCallTrack
func (c *VideoClient) VideoUnpin(ctx context.Context, _type string, id string, request *UnpinRequest) (*StreamResponse[UnpinResponse], error) {
	var result UnpinResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UnpinRequest, UnpinResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/unpin", nil, request, &result, pathParams)
	return res, err
}

// Updates user permissions
//
// Sends events:
// - call.permissions_updated
//
// Required permissions:
// - UpdateCallPermissions
func (c *VideoClient) UpdateUserPermissions(ctx context.Context, _type string, id string, request *UpdateUserPermissionsRequest) (*StreamResponse[UpdateUserPermissionsResponse], error) {
	var result UpdateUserPermissionsResponse
	pathParams := map[string]string{
		"type": _type,
		"id":   id,
	}
	res, err := MakeRequest[UpdateUserPermissionsRequest, UpdateUserPermissionsResponse](c.client, ctx, "POST", "/api/v2/video/call/{type}/{id}/user_permissions", nil, request, &result, pathParams)
	return res, err
}

// Deletes recording
//
// Required permissions:
// - DeleteRecording
func (c *VideoClient) DeleteRecording(ctx context.Context, _type string, id string, session string, filename string) (*StreamResponse[DeleteRecordingResponse], error) {
	var result DeleteRecordingResponse
	pathParams := map[string]string{
		"type":     _type,
		"id":       id,
		"session":  session,
		"filename": filename,
	}
	res, err := MakeRequest[any, DeleteRecordingResponse](c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/recordings/{filename}", nil, nil, &result, pathParams)
	return res, err
}

// Deletes transcription
//
// Required permissions:
// - DeleteTranscription
func (c *VideoClient) DeleteTranscription(ctx context.Context, _type string, id string, session string, filename string) (*StreamResponse[DeleteTranscriptionResponse], error) {
	var result DeleteTranscriptionResponse
	pathParams := map[string]string{
		"type":     _type,
		"id":       id,
		"session":  session,
		"filename": filename,
	}
	res, err := MakeRequest[any, DeleteTranscriptionResponse](c.client, ctx, "DELETE", "/api/v2/video/call/{type}/{id}/{session}/transcriptions/{filename}", nil, nil, &result, pathParams)
	return res, err
}

// Query calls with filter query
//
// Required permissions:
// - ReadCall
func (c *VideoClient) QueryCalls(ctx context.Context, request *QueryCallsRequest) (*StreamResponse[QueryCallsResponse], error) {
	var result QueryCallsResponse
	res, err := MakeRequest[QueryCallsRequest, QueryCallsResponse](c.client, ctx, "POST", "/api/v2/video/calls", nil, request, &result, nil)
	return res, err
}

func (c *VideoClient) ListCallTypes(ctx context.Context) (*StreamResponse[ListCallTypeResponse], error) {
	var result ListCallTypeResponse
	res, err := MakeRequest[any, ListCallTypeResponse](c.client, ctx, "GET", "/api/v2/video/calltypes", nil, nil, &result, nil)
	return res, err
}

func (c *VideoClient) CreateCallType(ctx context.Context, request *CreateCallTypeRequest) (*StreamResponse[CreateCallTypeResponse], error) {
	var result CreateCallTypeResponse
	res, err := MakeRequest[CreateCallTypeRequest, CreateCallTypeResponse](c.client, ctx, "POST", "/api/v2/video/calltypes", nil, request, &result, nil)
	return res, err
}

func (c *VideoClient) DeleteCallType(ctx context.Context, name string) (*StreamResponse[Response], error) {
	var result Response
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, Response](c.client, ctx, "DELETE", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams)
	return res, err
}

func (c *VideoClient) GetCallType(ctx context.Context, name string) (*StreamResponse[GetCallTypeResponse], error) {
	var result GetCallTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[any, GetCallTypeResponse](c.client, ctx, "GET", "/api/v2/video/calltypes/{name}", nil, nil, &result, pathParams)
	return res, err
}

func (c *VideoClient) UpdateCallType(ctx context.Context, name string, request *UpdateCallTypeRequest) (*StreamResponse[UpdateCallTypeResponse], error) {
	var result UpdateCallTypeResponse
	pathParams := map[string]string{
		"name": name,
	}
	res, err := MakeRequest[UpdateCallTypeRequest, UpdateCallTypeResponse](c.client, ctx, "PUT", "/api/v2/video/calltypes/{name}", nil, request, &result, pathParams)
	return res, err
}

// Returns the list of all edges available for video calls.
func (c *VideoClient) GetEdges(ctx context.Context) (*StreamResponse[GetEdgesResponse], error) {
	var result GetEdgesResponse
	res, err := MakeRequest[any, GetEdgesResponse](c.client, ctx, "GET", "/api/v2/video/edges", nil, nil, &result, nil)
	return res, err
}
