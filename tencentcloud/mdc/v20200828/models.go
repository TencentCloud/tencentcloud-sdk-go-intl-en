// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200828

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type CreateInput struct {
	// Enter a name. It can contain uppercase letters, lowercase letters, digits, and underscores. Length: [1, 32].
	InputName *string `json:"InputName,omitnil,omitempty" name:"InputName"`

	// Input protocol. Selectable: [SRT|RTP|RTMP|RTMP_PULL|RTSP_PULL|RIST].
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Enter description. Length: [0, 255].
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Input IP allowlist in CIDR format.
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// SRT configuration information of the input.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// Input RTP configuration message.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// Primary/replica switch input. Selectable values: `OPEN` and `CLOSE`. Default value: `CLOSE`.
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Input RTMP_PULL configuration message.
	RTMPPullSettings *CreateInputRTMPPullSettings `json:"RTMPPullSettings,omitnil,omitempty" name:"RTMPPullSettings"`

	// Input RTSP_PULL configuration message.
	RTSPPullSettings *CreateInputRTSPPullSettings `json:"RTSPPullSettings,omitnil,omitempty" name:"RTSPPullSettings"`

	// Input HLS_PULL configuration message.
	HLSPullSettings *CreateInputHLSPullSettings `json:"HLSPullSettings,omitnil,omitempty" name:"HLSPullSettings"`

	// Configuration message for smooth streaming during delayed broadcasting.
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil,omitempty" name:"ResilientStream"`

	// Bound input security group ID.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone. This parameter is optional. If disaster recovery is enabled, two different availability zones must be input. Otherwise, only one availability zone can be input.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Input RIST configuration message.
	RISTSettings *CreateInputRISTSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// Region of the input node
	InputRegion *string `json:"InputRegion,omitnil,omitempty" name:"InputRegion"`

	// Hot/cold backup
	FailOverOption *FailOverOption `json:"FailOverOption,omitnil,omitempty" name:"FailOverOption"`
}

type CreateInputHLSPullSettings struct {
	// HLS origin server address. Only one is allowed.
	SourceAddresses []*HLSPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type CreateInputRISTSettings struct {
	// RIST mode. Valid value: `LISTENER`. Default value: `LISTENER`.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// RIST configuration solution. Options: [MAIN|SIMPLE]. Default: MAIN.
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// RIST buffer size in milliseconds. Minimum value: 50 ms. Maximum value: 5000 ms. Default value: 120.
	Buffer *int64 `json:"Buffer,omitnil,omitempty" name:"Buffer"`
}

type CreateInputRTMPPullSettings struct {
	// RTMP origin server address. Only one is allowed.
	SourceAddresses []*RTMPPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type CreateInputRTPSettings struct {
	// Default value: "none". Available value: ['none'].
	FEC *string `json:"FEC,omitnil,omitempty" name:"FEC"`

	// Idle timeout. Default value: 5000 ms. Range: [1000, 10000].
	IdleTimeout *int64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

type CreateInputRTSPPullSettings struct {
	// Origin server address of the RTSP origin server. Only one is allowed.
	SourceAddresses []*RTSPPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type CreateInputSRTSettings struct {
	// SRT mode. Options: [LISTENER|CALLER]. Default value: LISTENER.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Stream Id. It can contain upper- and lower-case letters, digits, and special characters (.#!:&,=_-), with a length of 0 to 512. For the specific format, see: https://github.com/Haivision/srt/blob/master/docs/features/access-control.md#standard-keys.	
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// Delay. Default value: 0. Unit: ms. Range: [0, 3000].
	Latency *int64 `json:"Latency,omitnil,omitempty" name:"Latency"`

	// Receive latency. Default value: 120. Unit: ms. Range: [0, 3000].
	RecvLatency *int64 `json:"RecvLatency,omitnil,omitempty" name:"RecvLatency"`

	// Peer delay. Default value: 0. Unit: ms. Range: [0, 3000].
	PeerLatency *int64 `json:"PeerLatency,omitnil,omitempty" name:"PeerLatency"`

	// Peer timeout. Default value: 5000 ms. Range: [1000, 10000].
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil,omitempty" name:"PeerIdleTimeout"`

	// Decryption key. Empty by default, which means not encrypted. Only ascii characters are supported. Length range: [10, 79].
	Passphrase *string `json:"Passphrase,omitnil,omitempty" name:"Passphrase"`

	// Key length. Default value: 0. Options: [0|16|24|32].
	PbKeyLen *int64 `json:"PbKeyLen,omitnil,omitempty" name:"PbKeyLen"`

	// SRT peer address. Required when Mode is CALLER, and can only contain 1 group.
	SourceAddresses []*SRTSourceAddressReq `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`

	// SRT FEC settings
	FEC *SRTFECSimpleOptions `json:"FEC,omitnil,omitempty" name:"FEC"`
}

type CreateOutputInfo struct {
	// Output name.
	OutputName *string `json:"OutputName,omitnil,omitempty" name:"OutputName"`

	// Output description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Output relay protocol. Supports SRT, RTP, RTMP, RTMP_PULL, RTSP_PULL, and RIST.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Output region.
	OutputRegion *string `json:"OutputRegion,omitnil,omitempty" name:"OutputRegion"`

	// Configuration of the output SRT.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// Configuration of the output RTMP.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitnil,omitempty" name:"RTMPSettings"`

	// Configuration of the output RTP.
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// IP whitelist list in CIDR format, for example, 0.0.0.0/0.
	// Valid when Protocol is RTMP_PULL. If empty, it represents no limit on client IP.
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// Maximum number of concurrent stream pulls. Maximum value: 4. Default value: 4.
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil,omitempty" name:"MaxConcurrent"`

	// ID of the bound input security group.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone. output supports a maximum of one user-submitted input availability zone.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Output type: Internet/TencentCSS/StreamLive
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`

	// Output RIST configuration.
	RISTSettings *CreateOutputRistSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// For streams with multiple audio/video tracks, you can specify the track to be used
	//
	// Deprecated: PidSelector is deprecated.
	PidSelector *PidSelector `json:"PidSelector,omitnil,omitempty" name:"PidSelector"`

	// Output module type, including Pinpoint (single-point output, supports up to four concurrent outputs); MultiMesh (multiple streams output, supports more than four concurrent outputs, currently up to 200). The default type is Pinpoint output. For a single Flow, only one MultiMesh output is allowed per zone.
	OutputKind *string `json:"OutputKind,omitnil,omitempty" name:"OutputKind"`

	// For streams with multiple audio/video tracks, you can specify the track to be used
	StreamSelector *StreamSelector `json:"StreamSelector,omitnil,omitempty" name:"StreamSelector"`
}

type CreateOutputInfoRTPSettings struct {
	// Destination address. You can enter 1 to 2 addresses.
	Destinations []*CreateOutputRTPSettingsDestinations `json:"Destinations,omitnil,omitempty" name:"Destinations"`

	// Can only be filled in with none.
	FEC *string `json:"FEC,omitnil,omitempty" name:"FEC"`

	// Idle timeout period, in milliseconds.
	IdleTimeout *int64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

type CreateOutputRTMPSettings struct {
	// Target address for repost. You can enter 1 to 2 addresses.
	Destinations []*CreateOutputRtmpSettingsDestinations `json:"Destinations,omitnil,omitempty" name:"Destinations"`

	// Chunk size of RTMP, ranging from 4096 to 40960.
	ChunkSize *int64 `json:"ChunkSize,omitnil,omitempty" name:"ChunkSize"`
}

type CreateOutputRTPSettingsDestinations struct {
	// Destination IP for repush.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Destination port for push.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type CreateOutputRistSettings struct {
	// RIST mode. Selectable: [LISTENER|CALLER]. Default: LISTENER.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// RIST configuration solution. Options: [MAIN|SIMPLE]. Default: MAIN.
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// RIST buffer size in milliseconds. Minimum value: 50 ms. Maximum value: 5000 ms. Default value: 120
	Buffer *int64 `json:"Buffer,omitnil,omitempty" name:"Buffer"`
}

type CreateOutputRtmpSettingsDestinations struct {
	// URL for push forwarding, in the format of rtmp://domain/live.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// StreamKey for republishing, in the format of: stream?key=value.
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type CreateOutputSrtSettings struct {
	// Target address for push. Required when Mode is CALLER. Only 1 group can be filled in.
	Destinations []*CreateOutputSrtSettingsDestinations `json:"Destinations,omitnil,omitempty" name:"Destinations"`

	// SRT mode. Valid values: `LISTENER` and `CALLER`. Default value: `CALLER`.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Stream Id for SRT retransmission. It can contain upper- and lower-case letters, digits, and special characters (.#!:&,=_-), with a length of 0 to 512.
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// Total delay for SRT retransmission. Default value: 0. Unit: ms. Range: [0, 3000]. This parameter simultaneously sets the sender and recipient delays (recvlatency and peerlatency) to the same value. Recommended configuration: at least 3 times the RTT to ensure effective handling of packet retransmission and confirmation during network congestion.
	Latency *int64 `json:"Latency,omitnil,omitempty" name:"Latency"`

	// Receive latency for SRT relay. Default value: 120. Unit: ms. Range: [0, 3000].
	// This parameter indicates the time length for the recipient to cache data packets.
	RecvLatency *int64 `json:"RecvLatency,omitnil,omitempty" name:"RecvLatency"`

	// Peer delay for SRT retransmission. Default value: 0. Unit: ms. Range: [0, 3000].
	// This parameter is set by the sender to inform the recipient of the expected buffer time for delay.
	PeerLatency *int64 `json:"PeerLatency,omitnil,omitempty" name:"PeerLatency"`

	// Idle timeout for SRT relay. Default value: 5000 ms. Range: [1000, 10000].
	// If there is no activity on a connection within the set timeout period, it will be closed.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil,omitempty" name:"PeerIdleTimeout"`

	// Encryption key for SRT retweet. It is empty by default, indicating not encrypted. Only ascii characters are allowed, with a length of [10, 79].
	Passphrase *string `json:"Passphrase,omitnil,omitempty" name:"Passphrase"`

	// Key length for SRT retransmission. Default value: 0. Options: [0|16|24|32].
	PbKeyLen *int64 `json:"PbKeyLen,omitnil,omitempty" name:"PbKeyLen"`

	// SRT FEC settings
	FEC *SRTFECFullOptions `json:"FEC,omitnil,omitempty" name:"FEC"`
}

type CreateOutputSrtSettingsDestinations struct {
	// Output IP or domain.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Output port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type CreateStreamLinkFlowRequestParams struct {
	// Stream name.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Maximum bandwidth in bps. Options: [10000000, 20000000].
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// Input group of a stream.
	InputGroup []*CreateInput `json:"InputGroup,omitnil,omitempty" name:"InputGroup"`

	// Media transmission Event ID associated with the Flow. Each Flow can only be associated with one Event.
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Output group of the stream.
	OutputGroup []*CreateOutputInfo `json:"OutputGroup,omitnil,omitempty" name:"OutputGroup"`
}

type CreateStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Stream name.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Maximum bandwidth in bps. Options: [10000000, 20000000].
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// Input group of a stream.
	InputGroup []*CreateInput `json:"InputGroup,omitnil,omitempty" name:"InputGroup"`

	// Media transmission Event ID associated with the Flow. Each Flow can only be associated with one Event.
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Output group of the stream.
	OutputGroup []*CreateOutputInfo `json:"OutputGroup,omitnil,omitempty" name:"OutputGroup"`
}

func (r *CreateStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowName")
	delete(f, "MaxBandwidth")
	delete(f, "InputGroup")
	delete(f, "EventId")
	delete(f, "OutputGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkFlowResponseParams struct {
	// Information of the created Flow.
	Info *DescribeFlow `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkFlowResponseParams `json:"Response"`
}

func (r *CreateStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkInputRequestParams struct {
	// <p>Media transmission stream ID.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Input group of a stream.</p>
	InputGroup []*CreateInput `json:"InputGroup,omitnil,omitempty" name:"InputGroup"`
}

type CreateStreamLinkInputRequest struct {
	*tchttp.BaseRequest
	
	// <p>Media transmission stream ID.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Input group of a stream.</p>
	InputGroup []*CreateInput `json:"InputGroup,omitnil,omitempty" name:"InputGroup"`
}

func (r *CreateStreamLinkInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "InputGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkInputResponseParams struct {
	// <p>Information of the created Flow.</p>
	Info *DescribeFlow `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStreamLinkInputResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkInputResponseParams `json:"Response"`
}

func (r *CreateStreamLinkInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkOutputInfoRequestParams struct {
	// Transport stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output configuration of the transport stream.
	Output *CreateOutputInfo `json:"Output,omitnil,omitempty" name:"Output"`
}

type CreateStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// Transport stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output configuration of the transport stream.
	Output *CreateOutputInfo `json:"Output,omitnil,omitempty" name:"Output"`
}

func (r *CreateStreamLinkOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkOutputInfoResponseParams struct {
	// Output information after creation.
	Info *DescribeOutput `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStreamLinkOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkOutputInfoResponseParams `json:"Response"`
}

func (r *CreateStreamLinkOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkSecurityGroupRequestParams struct {
	// <p>Security group name, supporting letters, digits, and underscores, unique within a Region.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Allowlist, maximum number: 10.</p>
	Whitelist []*string `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

type CreateStreamLinkSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>Security group name, supporting letters, digits, and underscores, unique within a Region.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Allowlist, maximum number: 10.</p>
	Whitelist []*string `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

func (r *CreateStreamLinkSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Whitelist")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkSecurityGroupResponseParams struct {
	// <p>Security group ID.</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStreamLinkSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateStreamLinkSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateStreamLinkSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStreamLinkSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkFlowRequestParams struct {
	// <p>Transport stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DeleteStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Transport stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DeleteStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkFlowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLinkFlowResponseParams `json:"Response"`
}

func (r *DeleteStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkOutputRequestParams struct {
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output Id.
	OutputId *string `json:"OutputId,omitnil,omitempty" name:"OutputId"`
}

type DeleteStreamLinkOutputRequest struct {
	*tchttp.BaseRequest
	
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output Id.
	OutputId *string `json:"OutputId,omitnil,omitempty" name:"OutputId"`
}

func (r *DeleteStreamLinkOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "OutputId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStreamLinkOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStreamLinkOutputResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStreamLinkOutputResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStreamLinkOutputResponseParams `json:"Response"`
}

func (r *DeleteStreamLinkOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStreamLinkOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlow struct {
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Stream name.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Stream status. Valid values: IDLE, RUNNING.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Maximum bandwidth.
	MaxBandwidth *int64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// Input group.
	InputGroup []*DescribeInput `json:"InputGroup,omitnil,omitempty" name:"InputGroup"`

	// Output group.
	OutputGroup []*DescribeOutput `json:"OutputGroup,omitnil,omitempty" name:"OutputGroup"`

	// EventId of the media transmission event associated with this Flow.
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// Creation time of the Flow in UTC format.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Input protocols that can be created for this Flow.
	AllowedInputProtocols []*string `json:"AllowedInputProtocols,omitnil,omitempty" name:"AllowedInputProtocols"`

	// Output protocols allowed to be created by this Flow
	AllowedOutputProtocols []*string `json:"AllowedOutputProtocols,omitnil,omitempty" name:"AllowedOutputProtocols"`
}

type DescribeHLSPullSourceAddress struct {
	// Url of the HLS origin server.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DescribeInput struct {
	// Input Id.
	InputId *string `json:"InputId,omitnil,omitempty" name:"InputId"`

	// Input name.
	InputName *string `json:"InputName,omitnil,omitempty" name:"InputName"`

	// Input description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Input protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Input the address list.
	InputAddressList []*InputAddress `json:"InputAddressList,omitnil,omitempty" name:"InputAddressList"`

	// Input the IP whitelist list.
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// SRT configuration information of the input.
	SRTSettings *DescribeInputSRTSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// Input RTP configuration message.
	RTPSettings *DescribeInputRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// Input region.
	InputRegion *string `json:"InputRegion,omitnil,omitempty" name:"InputRegion"`

	// Input RTMP configuration message.
	RTMPSettings *DescribeInputRTMPSettings `json:"RTMPSettings,omitnil,omitempty" name:"RTMPSettings"`

	// Primary/replica switch input.
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Availability zone configuration. In the current situation with disaster recovery enabled, there can be up to two, corresponding to pipeline 0 and 1 in sequence. Otherwise, there is only one availability zone.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Input RTMP_PULL configuration message.
	RTMPPullSettings *DescribeInputRTMPPullSettings `json:"RTMPPullSettings,omitnil,omitempty" name:"RTMPPullSettings"`

	// Input RTSP_PULL configuration message.
	RTSPPullSettings *DescribeInputRTSPPullSettings `json:"RTSPPullSettings,omitnil,omitempty" name:"RTSPPullSettings"`

	// Input HLS_PULL configuration message.
	HLSPullSettings *DescribeInputHLSPullSettings `json:"HLSPullSettings,omitnil,omitempty" name:"HLSPullSettings"`

	// Configuration message for smooth streaming during delayed broadcasting.
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil,omitempty" name:"ResilientStream"`

	// Bound input security group ID.	
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Input RIST configuration message.
	RISTSettings *DescribeInputRISTSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// Input URL information related to module settings, including the provided streaming address or the configured third-party source stream URL.
	StreamUrls []*StreamUrlDetail `json:"StreamUrls,omitnil,omitempty" name:"StreamUrls"`

	// Disaster recovery configuration item
	FailOverOption *FailOverOption `json:"FailOverOption,omitnil,omitempty" name:"FailOverOption"`
}

type DescribeInputHLSPullSettings struct {
	// HLS origin server address information.
	SourceAddresses []*DescribeHLSPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type DescribeInputRISTSettings struct {
	// RIST mode. Options: [LISTENER|CALLER]. Default value: LISTENER.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// RIST configuration solution. Options: [MAIN|SIMPLE]. Default: MAIN.
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// RIST buffer size in milliseconds. Minimum value: 50 ms. Maximum value: 5000 ms. Default value: 120.
	Buffer *int64 `json:"Buffer,omitnil,omitempty" name:"Buffer"`
}

type DescribeInputRTMPPullSettings struct {
	// RTMP origin server address information.
	SourceAddresses []*DescribeRTMPPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type DescribeInputRTMPSettings struct {
	// RTMP push path.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// RTMP push StreamKey.
	// RTMP push URL splicing rule: rtmp://Ip:1935/AppName/StreamKey
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type DescribeInputRTPSettings struct {
	// Whether to enable FEC.
	FEC *string `json:"FEC,omitnil,omitempty" name:"FEC"`

	// Idle timeout period.
	IdleTimeout *int64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

type DescribeInputRTSPPullSettings struct {
	// RTSP origin server address.
	SourceAddresses []*DescribeRTSPPullSourceAddress `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type DescribeInputSRTSettings struct {
	// SRT mode.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Stream Id.
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// Delay.
	Latency *int64 `json:"Latency,omitnil,omitempty" name:"Latency"`

	// Receive delay.
	RecvLatency *int64 `json:"RecvLatency,omitnil,omitempty" name:"RecvLatency"`

	// Peer delay.
	PeerLatency *int64 `json:"PeerLatency,omitnil,omitempty" name:"PeerLatency"`

	// Idle timeout period of the peer.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil,omitempty" name:"PeerIdleTimeout"`

	// Decryption key.
	Passphrase *string `json:"Passphrase,omitnil,omitempty" name:"Passphrase"`

	// Key length.
	PbKeyLen *int64 `json:"PbKeyLen,omitnil,omitempty" name:"PbKeyLen"`

	// SRT peer address.
	SourceAddresses []*SRTSourceAddressResp `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`

	// FEC settings
	FEC *SRTFECSimpleOptions `json:"FEC,omitnil,omitempty" name:"FEC"`
}

type DescribeOutput struct {
	// <p>Output Id.</p>
	OutputId *string `json:"OutputId,omitnil,omitempty" name:"OutputId"`

	// <p>Output name.</p>
	OutputName *string `json:"OutputName,omitnil,omitempty" name:"OutputName"`

	// <p>Output type.</p>
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`

	// <p>Output description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Output protocol.</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>List of output egress address information.</p>
	OutputAddressList []*OutputAddress `json:"OutputAddressList,omitnil,omitempty" name:"OutputAddressList"`

	// <p>Output region.</p>
	OutputRegion *string `json:"OutputRegion,omitnil,omitempty" name:"OutputRegion"`

	// <p>SRT configuration information of the output.</p>
	SRTSettings *DescribeOutputSRTSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// <p>Output RTP configuration message.</p>
	RTPSettings *DescribeOutputRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// <p>Output RTMP configuration message.</p>
	RTMPSettings *DescribeOutputRTMPSettings `json:"RTMPSettings,omitnil,omitempty" name:"RTMPSettings"`

	// <p>Output RTMP pull configuration message.</p>
	RTMPPullSettings *DescribeOutputRTMPPullSettings `json:"RTMPPullSettings,omitnil,omitempty" name:"RTMPPullSettings"`

	// <p>CIDR allowlist.<br>Valid when Protocol is RTMP_PULL. If empty, it represents no limit on client IP.</p>
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// <p>Output RTSP pull stream configuration message.</p>
	RTSPPullSettings *DescribeOutputRTSPPullSettings `json:"RTSPPullSettings,omitnil,omitempty" name:"RTSPPullSettings"`

	// <p>Output HLS pull configuration message.</p>
	HLSPullSettings *DescribeOutputHLSPullSettings `json:"HLSPullSettings,omitnil,omitempty" name:"HLSPullSettings"`

	// <p>Maximum number of concurrent stream pulls. Maximum value: 4. Default value: 4.</p>
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil,omitempty" name:"MaxConcurrent"`

	// <p>ID of the bound security group.</p>
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// <p>Availability zone. output currently supports a maximum of one.</p>
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// <p>Output RIST configuration message.</p>
	RISTSettings *DescribeOutputRISTSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// <p>For streams with multiple audio/video tracks, you can specify the track to be used</p>
	//
	// Deprecated: PidSelector is deprecated.
	PidSelector *PidSelector `json:"PidSelector,omitnil,omitempty" name:"PidSelector"`

	// <p>Output module type, including Pinpoint (single-point output, supports up to four concurrent outputs); MultiMesh (multiple streams output, supports more than four concurrent outputs, currently up to 200). The default type is Pinpoint output. For one Flow, only one MultiMesh output is allowed in one zone.</p>
	OutputKind *string `json:"OutputKind,omitnil,omitempty" name:"OutputKind"`

	// <p>Output module configuration, related URLs, including the provided pull stream address or the configured relay address for output to a third party</p>
	StreamUrls []*StreamUrlDetail `json:"StreamUrls,omitnil,omitempty" name:"StreamUrls"`

	// <p>For streams with multiple audio/video tracks, you can specify the track to be used</p>
	StreamSelector *StreamSelector `json:"StreamSelector,omitnil,omitempty" name:"StreamSelector"`

	// <p>Output status.</p><p>Enumeration values:</p><ul><li>DISABLED</li><li>ENABLED</li></ul>
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type DescribeOutputHLSPullServerUrl struct {
	// Url of the HLS pull address.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DescribeOutputHLSPullSettings struct {
	// List of HLS pull addresses.
	ServerUrls []*DescribeOutputHLSPullServerUrl `json:"ServerUrls,omitnil,omitempty" name:"ServerUrls"`
}

type DescribeOutputRISTSettings struct {
	// RIST mode. Valid values: `LISTENER`, `CALLER`. Default value: `LISTENER`.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// RIST configuration solution. Options: [MAIN|SIMPLE]. Default: MAIN.
	Profile *string `json:"Profile,omitnil,omitempty" name:"Profile"`

	// RIST buffer size in milliseconds. Minimum value: 50 ms. Maximum value: 5000 ms. Default value: 120.
	Buffer *int64 `json:"Buffer,omitnil,omitempty" name:"Buffer"`

	// Server listening address, used when the RIST mode is LISTENER.
	SourceAddresses []*OutputRISTSourceAddressResp `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`
}

type DescribeOutputRTMPPullServerUrl struct {
	// tcUrl of the RTMP pull address.
	TcUrl *string `json:"TcUrl,omitnil,omitempty" name:"TcUrl"`

	// Stream key of the RTMP pull address.
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type DescribeOutputRTMPPullSettings struct {
	// List of pull addresses.
	ServerUrls []*DescribeOutputRTMPPullServerUrl `json:"ServerUrls,omitnil,omitempty" name:"ServerUrls"`
}

type DescribeOutputRTMPSettings struct {
	// Idle timeout period.
	IdleTimeout *int64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// Chunk size.
	ChunkSize *int64 `json:"ChunkSize,omitnil,omitempty" name:"ChunkSize"`

	// List of target address information for RTMP repush.
	Destinations []*RTMPAddressDestination `json:"Destinations,omitnil,omitempty" name:"Destinations"`
}

type DescribeOutputRTPSettings struct {
	// List of target address information for RTP repush.
	Destinations []*RTPAddressDestination `json:"Destinations,omitnil,omitempty" name:"Destinations"`

	// Whether to enable FEC.
	FEC *string `json:"FEC,omitnil,omitempty" name:"FEC"`

	// Idle timeout period.
	IdleTimeout *int64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`
}

type DescribeOutputRTSPPullServerUrl struct {
	// RTSP pull Url.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DescribeOutputRTSPPullSettings struct {
	// List of RTSP pull addresses.
	ServerUrls []*DescribeOutputRTSPPullServerUrl `json:"ServerUrls,omitnil,omitempty" name:"ServerUrls"`
}

type DescribeOutputSRTSettings struct {
	// Destination address list. Used when the SRT mode is CALLER.
	Destinations []*SRTAddressDestination `json:"Destinations,omitnil,omitempty" name:"Destinations"`

	// Stream Id.
	StreamId *string `json:"StreamId,omitnil,omitempty" name:"StreamId"`

	// Delay.
	Latency *int64 `json:"Latency,omitnil,omitempty" name:"Latency"`

	// Receive delay.
	RecvLatency *int64 `json:"RecvLatency,omitnil,omitempty" name:"RecvLatency"`

	// Peer delay.
	PeerLatency *int64 `json:"PeerLatency,omitnil,omitempty" name:"PeerLatency"`

	// Idle timeout period of the peer.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitnil,omitempty" name:"PeerIdleTimeout"`

	// Encryption key.
	Passphrase *string `json:"Passphrase,omitnil,omitempty" name:"Passphrase"`

	// Encryption key length.
	PbKeyLen *int64 `json:"PbKeyLen,omitnil,omitempty" name:"PbKeyLen"`

	// SRT mode.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Server listening address, used when the SRT mode is LISTENER.
	SourceAddresses []*OutputSRTSourceAddressResp `json:"SourceAddresses,omitnil,omitempty" name:"SourceAddresses"`

	// SRT FEC configuration
	FEC *SRTFECFullOptions `json:"FEC,omitnil,omitempty" name:"FEC"`
}

type DescribeRTMPPullSourceAddress struct {
	// TcUrl of the RTMP origin server.
	TcUrl *string `json:"TcUrl,omitnil,omitempty" name:"TcUrl"`

	// StreamKey of the RTMP origin server.
	// RTMP origin server URL splicing rule: $TcUrl/$StreamKey.
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type DescribeRTSPPullSourceAddress struct {
	// Url of the RTSP origin server.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type DescribeStreamLinkFlowLogsRequestParams struct {
	// <p>Transport stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Start time for statistics, defaults to the last hour, supports querying up to the last 7 days.<br>UTC time, for example, '2020-01-01T12:00:00Z'.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time for statistics, defaulting to one hour after StartTime, with support for querying up to 24 hours of data.<br>UTC time, for example, '2020-01-01T12:00:00Z'.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Input or output type. Valid values: `input`, `output`.</p>
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Primary or backup channel. Option [0|1].</p>
	Pipeline []*string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// <p>Size per page. Default value: 100. Value range: [1, 1000].</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Sort by Timestamp in ascending or descending order. The default is descending order. Options: [desc|asc].</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>Page number. Default value: 1. Value range: [1, 1000].</p>
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

type DescribeStreamLinkFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Transport stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Start time for statistics, defaults to the last hour, supports querying up to the last 7 days.<br>UTC time, for example, '2020-01-01T12:00:00Z'.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>End time for statistics, defaulting to one hour after StartTime, with support for querying up to 24 hours of data.<br>UTC time, for example, '2020-01-01T12:00:00Z'.</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Input or output type. Valid values: `input`, `output`.</p>
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Primary or backup channel. Option [0|1].</p>
	Pipeline []*string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// <p>Size per page. Default value: 100. Value range: [1, 1000].</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Sort by Timestamp in ascending or descending order. The default is descending order. Options: [desc|asc].</p>
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>Page number. Default value: 1. Value range: [1, 1000].</p>
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`
}

func (r *DescribeStreamLinkFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Pipeline")
	delete(f, "PageSize")
	delete(f, "SortType")
	delete(f, "PageNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowLogsResponseParams struct {
	// <p>Log information list.</p>
	Infos []*FlowLogInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// <p>Current page.</p>
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// <p>Size per page.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Total quantity.</p>
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// <p>Total number of pages.</p>
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowLogsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowMediaStatisticsRequestParams struct {
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Query interval. Options: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Defaults to one hour after `StartTime`. You can query the data of up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Peer IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

type DescribeStreamLinkFlowMediaStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Query interval. Options: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Defaults to one hour after `StartTime`. You can query the data of up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Peer IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

func (r *DescribeStreamLinkFlowMediaStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowMediaStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RemoteIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowMediaStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowMediaStatisticsResponseParams struct {
	// List of media data in the transport stream.
	Infos []*FlowMediaInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowMediaStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowMediaStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowMediaStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowMediaStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRealtimeStatusRequestParams struct {
	// <p>Stream ID.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Input an id array. If both the input and output arrays are empty, it signifies a full query.</p>
	InputIds []*string `json:"InputIds,omitnil,omitempty" name:"InputIds"`

	// <p>Output id array. If both input and output arrays are empty, it signifies a full query.</p>
	OutputIds []*string `json:"OutputIds,omitnil,omitempty" name:"OutputIds"`
}

type DescribeStreamLinkFlowRealtimeStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>Stream ID.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// <p>Input an id array. If both the input and output arrays are empty, it signifies a full query.</p>
	InputIds []*string `json:"InputIds,omitnil,omitempty" name:"InputIds"`

	// <p>Output id array. If both input and output arrays are empty, it signifies a full query.</p>
	OutputIds []*string `json:"OutputIds,omitnil,omitempty" name:"OutputIds"`
}

func (r *DescribeStreamLinkFlowRealtimeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRealtimeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "InputIds")
	delete(f, "OutputIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowRealtimeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRealtimeStatusResponseParams struct {
	// <p>Inquiry time. Unit: s.</p>
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>Real-time data information list.</p>
	Datas []*FlowRealtimeStatusItem `json:"Datas,omitnil,omitempty" name:"Datas"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowRealtimeStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowRealtimeStatusResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowRealtimeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRealtimeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowRequestParams struct {
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type DescribeStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *DescribeStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowResponseParams struct {
	// <p>Configuration information of the stream.</p>
	Info *DescribeFlow `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowSRTStatisticsRequestParams struct {
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Default value: one hour after `StartTime`. You can query data up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query interval. Options: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Peer IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

type DescribeStreamLinkFlowSRTStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Default value: one hour after `StartTime`. You can query data up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query interval. Options: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Peer IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

func (r *DescribeStreamLinkFlowSRTStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowSRTStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "RemoteIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowSRTStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowSRTStatisticsResponseParams struct {
	// List of SRT quality data in the transport stream.
	Infos []*FlowSRTInfo `json:"Infos,omitnil,omitempty" name:"Infos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowSRTStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowSRTStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowSRTStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowSRTStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowStatisticsRequestParams struct {
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Query interval. Option: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Defaults to one hour after `StartTime`. You can query the data of up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Peer client or server IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

type DescribeStreamLinkFlowStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// Transport stream ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input or output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Primary or backup channel. Valid values: 0 and 1.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Query interval. Option: [5s|1min|5min|15min].
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// Start time for statistics. Default value: last hour. You can query data from the last 7 days at most.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for statistics. Defaults to one hour after `StartTime`. You can query the data of up to 24 hours.
	// UTC time, for example, '2020-01-01T12:00:00Z'.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Peer client or server IP address.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`
}

func (r *DescribeStreamLinkFlowStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Type")
	delete(f, "InputOutputId")
	delete(f, "Pipeline")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RemoteIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowStatisticsResponseParams struct {
	// List of media data in the transport stream.
	Infos []*FlowStatisticsArray `json:"Infos,omitnil,omitempty" name:"Infos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowStatisticsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowsRequestParams struct {
	// Current page number. Default value: 1.
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Size per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeStreamLinkFlowsRequest struct {
	*tchttp.BaseRequest
	
	// Current page number. Default value: 1.
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Size per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeStreamLinkFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowsResponseParams struct {
	// Configuration information list of the stream.
	Infos []*DescribeFlow `json:"Infos,omitnil,omitempty" name:"Infos"`

	// Current page number.
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Size per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total quantity.
	TotalNum *int64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// Total number of pages.
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkFlowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkFlowsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkRegionsRequestParams struct {

}

type DescribeStreamLinkRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamLinkRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkRegionsResponseParams struct {
	// Media transmission region information.
	Info *StreamLinkRegionInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamLinkRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamLinkRegionsResponseParams `json:"Response"`
}

func (r *DescribeStreamLinkRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamLinkRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailOverOption struct {
	// Hot backup
	FailOverType *string `json:"FailOverType,omitnil,omitempty" name:"FailOverType"`
}

type FlowAudio struct {
	// Frame rate.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate in bps.
	Rate *int64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// Audio Pid.
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
}

type FlowLogInfo struct {
	// Timestamp in seconds.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Input/output type. Valid values: `input`, `output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input or output Id.
	InputOutputId *string `json:"InputOutputId,omitnil,omitempty" name:"InputOutputId"`

	// Protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Event code.
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// Event information.
	EventMessage *string `json:"EventMessage,omitnil,omitempty" name:"EventMessage"`

	// Peer IP.
	RemoteIp *string `json:"RemoteIp,omitnil,omitempty" name:"RemoteIp"`

	// Peer port.
	RemotePort *string `json:"RemotePort,omitnil,omitempty" name:"RemotePort"`

	// Primary/secondary tunnel. 0: primary, 1: backup.
	Pipeline *string `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// Input or output name.
	InputOutputName *string `json:"InputOutputName,omitnil,omitempty" name:"InputOutputName"`
}

type FlowMediaAudio struct {
	// Frame rate.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate in bps.
	Rate *int64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// Audio Pid.
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Flag of the same push.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type FlowMediaInfo struct {
	// Timestamp in seconds.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Total bandwidth.
	Network *int64 `json:"Network,omitnil,omitempty" name:"Network"`

	// Video data of the transport stream.
	Video []*FlowMediaVideo `json:"Video,omitnil,omitempty" name:"Video"`

	// Audio data of the transport stream.
	Audio []*FlowMediaAudio `json:"Audio,omitnil,omitempty" name:"Audio"`

	// Flag of the same push.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Client IP.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`
}

type FlowMediaVideo struct {
	// Frame rate.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate in bps.
	Rate *int64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// Video Pid.
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Flag of the same push.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type FlowRealtimeStatusCommon struct {
	// Current connection status. Connected | Waiting | Idle.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Connection mode. Listener|Caller.
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Connection duration, unit: ms.
	ConnectedTime *int64 `json:"ConnectedTime,omitnil,omitempty" name:"ConnectedTime"`

	// Real-time bitrate, unit in bps.
	Bitrate *int64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`

	// Number of retries.
	Reconnections *int64 `json:"Reconnections,omitnil,omitempty" name:"Reconnections"`
}

type FlowRealtimeStatusItem struct {
	// Type. Valid values: `Input`, `Output`.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Input Id. If Type is Input, this field is not empty.
	InputId *string `json:"InputId,omitnil,omitempty" name:"InputId"`

	// Output Id. If Type is Output, this field is not empty.
	OutputId *string `json:"OutputId,omitnil,omitempty" name:"OutputId"`

	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Protocol. SRT | RTMP.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Common status information.
	CommonStatus *FlowRealtimeStatusCommon `json:"CommonStatus,omitnil,omitempty" name:"CommonStatus"`

	// If it is the SRT Protocol, this field will be present.
	SRTStatus *FlowRealtimeStatusSRT `json:"SRTStatus,omitnil,omitempty" name:"SRTStatus"`

	// If it is the RTMP protocol, this field exists.
	RTMPStatus *FlowRealtimeStatusRTMP `json:"RTMPStatus,omitnil,omitempty" name:"RTMPStatus"`

	// Server IP address.
	ConnectServerIP *string `json:"ConnectServerIP,omitnil,omitempty" name:"ConnectServerIP"`

	// If it is the RTP protocol, this field exists.
	RTPStatus *FlowRealtimeStatusRTP `json:"RTPStatus,omitnil,omitempty" name:"RTPStatus"`
}

type FlowRealtimeStatusRTMP struct {
	// Video frame rate.
	VideoFPS *int64 `json:"VideoFPS,omitnil,omitempty" name:"VideoFPS"`

	// Audio frame rate.
	AudioFPS *int64 `json:"AudioFPS,omitnil,omitempty" name:"AudioFPS"`
}

type FlowRealtimeStatusRTP struct {
	// Number of transmitted packets
	Packets *int64 `json:"Packets,omitnil,omitempty" name:"Packets"`
}

type FlowRealtimeStatusSRT struct {
	// Delay in ms.
	Latency *int64 `json:"Latency,omitnil,omitempty" name:"Latency"`

	// RTT in ms.
	RTT *int64 `json:"RTT,omitnil,omitempty" name:"RTT"`

	// Number of packets sent or received in real time.
	Packets *int64 `json:"Packets,omitnil,omitempty" name:"Packets"`

	// Packet loss rate.
	PacketLossRate *float64 `json:"PacketLossRate,omitnil,omitempty" name:"PacketLossRate"`

	// Retransmission rate.
	RetransmitRate *float64 `json:"RetransmitRate,omitnil,omitempty" name:"RetransmitRate"`

	// Real-time packet loss count.
	DroppedPackets *int64 `json:"DroppedPackets,omitnil,omitempty" name:"DroppedPackets"`

	// Whether to encrypt. On | Off.
	Encryption *string `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type FlowSRTInfo struct {
	// <p>Timestamp in seconds.</p>
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// <p>Packet loss rate on transmitting.</p>
	SendPacketLossRate *int64 `json:"SendPacketLossRate,omitnil,omitempty" name:"SendPacketLossRate"`

	// <p>Retransmission rate.</p>
	SendRetransmissionRate *int64 `json:"SendRetransmissionRate,omitnil,omitempty" name:"SendRetransmissionRate"`

	// <p>Packet loss rate on receiving.</p>
	RecvPacketLossRate *int64 `json:"RecvPacketLossRate,omitnil,omitempty" name:"RecvPacketLossRate"`

	// <p>Receiving retransmission rate.</p>
	RecvRetransmissionRate *int64 `json:"RecvRetransmissionRate,omitnil,omitempty" name:"RecvRetransmissionRate"`

	// <p>RTT latency to the peer.</p>
	RTT *int64 `json:"RTT,omitnil,omitempty" name:"RTT"`

	// <p>Flag the same streaming session.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Number of discarded packets on sending.</p>
	SendPacketDropNumber *int64 `json:"SendPacketDropNumber,omitnil,omitempty" name:"SendPacketDropNumber"`

	// <p>Number of received dropped packets.</p>
	RecvPacketDropNumber *int64 `json:"RecvPacketDropNumber,omitnil,omitempty" name:"RecvPacketDropNumber"`

	// <p>Bandwidth.</p>
	Bandwidth *float64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// <p>Transmit bandwidth</p>
	SendBandwidth *float64 `json:"SendBandwidth,omitnil,omitempty" name:"SendBandwidth"`

	// <p>Receiving bandwidth</p>
	RecvBandwidth *float64 `json:"RecvBandwidth,omitnil,omitempty" name:"RecvBandwidth"`

	// <p>Sent packet count</p>
	SendPackets *int64 `json:"SendPackets,omitnil,omitempty" name:"SendPackets"`

	// <p>Number of received packets</p>
	RecvPackets *int64 `json:"RecvPackets,omitnil,omitempty" name:"RecvPackets"`

	// <p>Sending lost packets</p>
	SendLostPackets *int64 `json:"SendLostPackets,omitnil,omitempty" name:"SendLostPackets"`

	// <p>Accept lost packets</p>
	RecvLostPackets *int64 `json:"RecvLostPackets,omitnil,omitempty" name:"RecvLostPackets"`

	// <p>Sent retransmission packet count</p>
	SendRetransmitPackets *int64 `json:"SendRetransmitPackets,omitnil,omitempty" name:"SendRetransmitPackets"`

	// <p>Number of received retransmission packets</p>
	RecvRetransmitPackets *int64 `json:"RecvRetransmitPackets,omitnil,omitempty" name:"RecvRetransmitPackets"`

	// <p>Flight window size</p>
	FlightSize *int64 `json:"FlightSize,omitnil,omitempty" name:"FlightSize"`

	// <p>Congestion window.</p>
	CongestionWindow *int64 `json:"CongestionWindow,omitnil,omitempty" name:"CongestionWindow"`

	// <p>Sending buffer (ms)</p>
	SendBuffer *int64 `json:"SendBuffer,omitnil,omitempty" name:"SendBuffer"`

	// <p>Receive buffer (ms)</p>
	RecvBuffer *int64 `json:"RecvBuffer,omitnil,omitempty" name:"RecvBuffer"`

	// <p>Sending delay</p>
	SendLatency *int64 `json:"SendLatency,omitnil,omitempty" name:"SendLatency"`

	// <p>Receive latency</p>
	RecvLatency *int64 `json:"RecvLatency,omitnil,omitempty" name:"RecvLatency"`

	// <p>Sent NAK packet count</p>
	SendNAKPackets *int64 `json:"SendNAKPackets,omitnil,omitempty" name:"SendNAKPackets"`

	// <p>Number of received NAK packets</p>
	RecvNAKPackets *int64 `json:"RecvNAKPackets,omitnil,omitempty" name:"RecvNAKPackets"`
}

type FlowStatistics struct {
	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Peer IP.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// Total bandwidth.
	Network *int64 `json:"Network,omitnil,omitempty" name:"Network"`

	// Video data.
	Video []*FlowVideo `json:"Video,omitnil,omitempty" name:"Video"`

	// Audio data.
	Audio []*FlowAudio `json:"Audio,omitnil,omitempty" name:"Audio"`
}

type FlowStatisticsArray struct {
	// Timestamp.
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Statistics for each session.
	FlowStatistics []*FlowStatistics `json:"FlowStatistics,omitnil,omitempty" name:"FlowStatistics"`
}

type FlowVideo struct {
	// Frame rate.
	Fps *int64 `json:"Fps,omitnil,omitempty" name:"Fps"`

	// Bitrate in bps.
	Rate *int64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// Audio Pid.
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
}

type HLSPullSourceAddress struct {
	// Url of the HLS origin server.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type InputAddress struct {
	// IP of the input address.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// The port of the input address.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type ModifyInput struct {
	// Input Id.
	InputId *string `json:"InputId,omitnil,omitempty" name:"InputId"`

	// Enter a name.
	InputName *string `json:"InputName,omitnil,omitempty" name:"InputName"`

	// Input description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// IPs allowed for push, in CIDR format.
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// Configuration of SRT.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// Configuration message of RTP.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// Input protocol. Options: SRT, RTP, RTMP, RTMP_PULL, RTSP_PULL, RIST.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Input primary/replica switch. Option: [OPEN|CLOSE].
	FailOver *string `json:"FailOver,omitnil,omitempty" name:"FailOver"`

	// Configuration message of RTMP_PULL.
	RTMPPullSettings *CreateInputRTMPPullSettings `json:"RTMPPullSettings,omitnil,omitempty" name:"RTMPPullSettings"`

	// Configuration message of RTSP_PULL.
	RTSPPullSettings *CreateInputRTSPPullSettings `json:"RTSPPullSettings,omitnil,omitempty" name:"RTSPPullSettings"`

	// Configuration message of HLS_PULL.
	HLSPullSettings *CreateInputHLSPullSettings `json:"HLSPullSettings,omitnil,omitempty" name:"HLSPullSettings"`

	// Configuration message for smooth stream output during delayed playback.
	ResilientStream *ResilientStreamConf `json:"ResilientStream,omitnil,omitempty" name:"ResilientStream"`

	// ID of the bound input security group. Only one security group can be associated.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Availability zone. Optional. input limit: two availability zones. For this API, only the second availability zone participates in resource allocation. It is valid if disaster recovery is enabled or the RTSP_PULL protocol is involved during switching. The address will be reallocated.	
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Configuration information of RIST.
	RISTSettings *CreateInputRISTSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// Region of the input node.
	InputRegion *string `json:"InputRegion,omitnil,omitempty" name:"InputRegion"`

	// Hot/cold backup related
	FailOverOption *FailOverOption `json:"FailOverOption,omitnil,omitempty" name:"FailOverOption"`
}

type ModifyOutputInfo struct {
	// Id of the Output to be modified.
	OutputId *string `json:"OutputId,omitnil,omitempty" name:"OutputId"`

	// Output name.
	OutputName *string `json:"OutputName,omitnil,omitempty" name:"OutputName"`

	// Output description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Output relay protocol. Supports SRT, RTP, RTMP, RTMP_PULL, RTSP_PULL, and RIST.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Output type: Internet/TencentCSS/StreamLive
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`

	// Output module type, including Pinpoint (single-point output, supports up to four concurrent outputs); MultiMesh (multiple outputs, supports more than four concurrent outputs, currently up to 200). The default type is Pinpoint output. For a single Flow, only one MultiMesh output is allowed in one zone.
	OutputKind *string `json:"OutputKind,omitnil,omitempty" name:"OutputKind"`

	// Configuration of the relayed SRT.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitnil,omitempty" name:"SRTSettings"`

	// Configuration of the relayed RTP.
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitnil,omitempty" name:"RTPSettings"`

	// Configuration for relaying RTMP.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitnil,omitempty" name:"RTMPSettings"`

	// IP whitelist list in CIDR format, for example, 0.0.0.0/0.
	// Valid when Protocol is RTMP_PULL. If empty, it represents no limit on client IP.
	AllowIpList []*string `json:"AllowIpList,omitnil,omitempty" name:"AllowIpList"`

	// Maximum number of concurrent stream pulls. Maximum value: 4. Default value: 4.
	MaxConcurrent *uint64 `json:"MaxConcurrent,omitnil,omitempty" name:"MaxConcurrent"`

	// ID of the bound security group. Only one security group can be associated.	
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// AZ.
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Relay RIST configuration.
	RISTSettings *CreateOutputRistSettings `json:"RISTSettings,omitnil,omitempty" name:"RISTSettings"`

	// For streams with multiple audio/video tracks, you can specify the track to be used.
	//
	// Deprecated: PidSelector is deprecated.
	PidSelector *PidSelector `json:"PidSelector,omitnil,omitempty" name:"PidSelector"`

	// For streams with multiple audio/video tracks, you can specify the track to be used.
	StreamSelector *StreamSelector `json:"StreamSelector,omitnil,omitempty" name:"StreamSelector"`
}

// Predefined struct for user
type ModifyStreamLinkFlowRequestParams struct {
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Stream name to be modified.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`
}

type ModifyStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Stream name to be modified.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`
}

func (r *ModifyStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "FlowName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkFlowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkFlowResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkInputRequestParams struct {
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input information that needs to be modified.
	Input *ModifyInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type ModifyStreamLinkInputRequest struct {
	*tchttp.BaseRequest
	
	// Stream Id.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Input information that needs to be modified.
	Input *ModifyInput `json:"Input,omitnil,omitempty" name:"Input"`
}

func (r *ModifyStreamLinkInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkInputResponseParams struct {
	// Input information after modification.
	Info *DescribeInput `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStreamLinkInputResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkInputResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkOutputInfoRequestParams struct {
	// 01937702ecc509dc0f3269ca3420
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output configuration that needs to be modified.
	Output *ModifyOutputInfo `json:"Output,omitnil,omitempty" name:"Output"`
}

type ModifyStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// 01937702ecc509dc0f3269ca3420
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Output configuration that needs to be modified.
	Output *ModifyOutputInfo `json:"Output,omitnil,omitempty" name:"Output"`
}

func (r *ModifyStreamLinkOutputInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkOutputInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStreamLinkOutputInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStreamLinkOutputInfoResponseParams struct {
	// Output configuration after modification.
	Info *DescribeOutput `json:"Info,omitnil,omitempty" name:"Info"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStreamLinkOutputInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStreamLinkOutputInfoResponseParams `json:"Response"`
}

func (r *ModifyStreamLinkOutputInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStreamLinkOutputInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OutputAddress struct {
	// Outbound IP address.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type OutputRISTSourceAddressResp struct {
	// Listening IP.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Listening port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type OutputSRTSourceAddressResp struct {
	// Listening IP.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Listening port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type PidSelector struct {
	// For streams containing multiple audio tracks, specify the audio track to use by entering a PID. The PID can be a positive integer between 1 and 8191.
	AudioPID []*int64 `json:"AudioPID,omitnil,omitempty" name:"AudioPID"`

	// For streams containing multiple video tracks, specify the video track to be used by entering a PID. You can enter a positive integer between 1 and 8191 for the PID.
	VideoPID []*int64 `json:"VideoPID,omitnil,omitempty" name:"VideoPID"`
}

type RTMPAddressDestination struct {
	// Destination Url for RTMP repush, in the format of 'RTMP://domain/live'.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Target StreamKey for RTMP republishing, in the format of 'steamid?key=value'.
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type RTMPPullSourceAddress struct {
	// TcUrl of the RTMP origin server.
	TcUrl *string `json:"TcUrl,omitnil,omitempty" name:"TcUrl"`

	// StreamKey information of the RTMP origin server.
	StreamKey *string `json:"StreamKey,omitnil,omitempty" name:"StreamKey"`
}

type RTPAddressDestination struct {
	// IP of the destination address for repush.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Port of the destination address for push.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type RTSPPullSourceAddress struct {
	// Url of the RTSP origin server.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type RegionInfo struct {
	// Zone name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ResilientStreamConf struct {
	// Whether to enable smooth delivery during delayed streaming. true: enable; false: disable. Not enabled by default.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Delay time in seconds, ranging from 10 to 300 for the time being.
	BufferTime *uint64 `json:"BufferTime,omitnil,omitempty" name:"BufferTime"`
}

type SRTAddressDestination struct {
	// IP of the target address.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// The port of the destination address.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type SRTFECFullOptions struct {
	// Whether to enable FEC.
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Number of FEC data packet Layout columns. Value range: >0
	Cols *int64 `json:"Cols,omitnil,omitempty" name:"Cols"`

	// Number of FEC data packet Layout rows. Value ranges from >=2 or <=-2.
	Rows *int64 `json:"Rows,omitnil,omitempty" name:"Rows"`

	// ARQ policy when FEC is enabled. Value: "always", "onreq", "never".
	ARQ *string `json:"ARQ,omitnil,omitempty" name:"ARQ"`

	// FEC data packet Layout organization. Value: "even", "staircase"
	Layout *string `json:"Layout,omitnil,omitempty" name:"Layout"`
}

type SRTFECSimpleOptions struct {
	// Whether FEC is enabled
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type SRTSourceAddressReq struct {
	// Peer IP or domain
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Peer port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type SRTSourceAddressResp struct {
	// Peer IP.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Peer port.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type StartStreamLinkFlowRequestParams struct {
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type StartStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *StartStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamLinkFlowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *StartStreamLinkFlowResponseParams `json:"Response"`
}

func (r *StartStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLinkFlowRequestParams struct {
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

type StopStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Stream Id.</p>
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`
}

func (r *StopStreamLinkFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLinkFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamLinkFlowResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *StopStreamLinkFlowResponseParams `json:"Response"`
}

func (r *StopStreamLinkFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamLinkFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StreamLinkRegionInfo struct {
	// List of region information for media direct transmission.
	Regions []*RegionInfo `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type StreamSelector struct {
	// Selection type: PID | TRACK
	SelectorType *string `json:"SelectorType,omitnil,omitempty" name:"SelectorType"`

	// Based on the PID configuration selector
	PidSelector *PidSelector `json:"PidSelector,omitnil,omitempty" name:"PidSelector"`

	// Select a selector based on the Track configuration
	TrackSelector *TrackSelector `json:"TrackSelector,omitnil,omitempty" name:"TrackSelector"`
}

type StreamUrlDetail struct {
	// Describes carrier information and other details.
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// URL 
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Playback: stream playback address;
	// RelayDestination: relay destination address; SourceCaptureUrl: source pull stream address; IngestEndpoint: streaming address
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TrackSelector struct {
	// Video track serial number. It starts from 1. Maximum value: 32.
	VideoIndex []*int64 `json:"VideoIndex,omitnil,omitempty" name:"VideoIndex"`

	// Video track serial number. Starts from 1. Maximum value: 32.
	AudioIndex []*int64 `json:"AudioIndex,omitnil,omitempty" name:"AudioIndex"`
}