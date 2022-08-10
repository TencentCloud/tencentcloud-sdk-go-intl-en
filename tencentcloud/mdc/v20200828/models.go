// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateInput struct {
	// Input name, which can contain 1 to 32 letters, digits, and underscores.
	InputName *string `json:"InputName,omitempty" name:"InputName"`

	// Input protocol. Valid values: `SRT`, `RTP`, `RTMP`
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Input description. Length: [0, 255].
	Description *string `json:"Description,omitempty" name:"Description"`

	// Allowlist of input IPs in CIDR format.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`

	// SRT configuration information of input.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration information of input.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// Input failover. Valid values: `OPEN`, `CLOSE` (default)
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`
}

type CreateInputRTPSettings struct {
	// Default value: none. Valid values: ['none'].
	FEC *string `json:"FEC,omitempty" name:"FEC"`

	// Idle timeout period in ms. Default value: 5000. Value range: [1000, 10000].
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
}

type CreateInputSRTSettings struct {
	// The SRT mode. Valid values: LISTENER (default), CALLER.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Stream ID, which can contain 0 to 512 letters, digits, and special characters (.#!:&,=_-).
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Latency in ms. Default value: 0. Value range: [0, 3000].
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// Receive latency in ms. Default value: 120. Value range: [0, 3000].
	RecvLatency *int64 `json:"RecvLatency,omitempty" name:"RecvLatency"`

	// Peer latency in ms. Default value: 0. Value range: [0, 3000].
	PeerLatency *int64 `json:"PeerLatency,omitempty" name:"PeerLatency"`

	// Peer timeout period in ms. Default value: 5000. Value range: [1000, 10000].
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitempty" name:"PeerIdleTimeout"`

	// Decryption key, which is empty by default, indicating not to encrypt. Only ASCII codes can be filled. Length: [10, 79].
	Passphrase *string `json:"Passphrase,omitempty" name:"Passphrase"`

	// Key length. Default value: 0. Valid values: 0, 16, 24, 32.
	PbKeyLen *int64 `json:"PbKeyLen,omitempty" name:"PbKeyLen"`

	// The SRT peer address, which is required if `Mode` is `CALLER`. Only one address is allowed.
	SourceAddresses []*SRTSourceAddressReq `json:"SourceAddresses,omitempty" name:"SourceAddresses"`
}

type CreateOutputInfo struct {
	// The output name.
	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`

	// Description of the output.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The output protocol. Valid values: SRT, RTP, RTMP, RTMP_PULL.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The output region.
	OutputRegion *string `json:"OutputRegion,omitempty" name:"OutputRegion"`

	// The SRT configuration.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// The RTMP configuration.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`

	// The RTP configuration.
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// The IP allowlist. The address must be in CIDR format, such as `0.0.0.0/0`.
	// This parameter is valid if `Protocol` is set to `RTMP_PULL`. If it is left empty, there is no restriction on clients’ IP addresses.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`
}

type CreateOutputInfoRTPSettings struct {
	// The relay destination addresses. One or two addresses are allowed.
	Destinations []*CreateOutputRTPSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// This parameter must be set to `none`.
	FEC *string `json:"FEC,omitempty" name:"FEC"`

	// The timeout period (ms).
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
}

type CreateOutputRTMPSettings struct {
	// The relay destination addresses. One or two addresses are allowed.
	Destinations []*CreateOutputRtmpSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// The RTMP chunk size. Value range: [4096, 40960].
	ChunkSize *int64 `json:"ChunkSize,omitempty" name:"ChunkSize"`
}

type CreateOutputRTPSettingsDestinations struct {
	// The relay destination IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The relay destination port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type CreateOutputRtmpSettingsDestinations struct {
	// The relay URL. Format: `rtmp://domain/live`.
	Url *string `json:"Url,omitempty" name:"Url"`

	// The `StreamKey` for relay. Format: `stream?key=value`.
	StreamKey *string `json:"StreamKey,omitempty" name:"StreamKey"`
}

type CreateOutputSrtSettings struct {
	// The relay destination address, which is required if `Mode` is `CALLER`. Only one address is allowed.
	Destinations []*CreateOutputSrtSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// The stream ID for relay, which can contain 0 to 512 letters, digits, and special characters (.#!:&,=_-).
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// The total latency (ms) of SRT relay. Value range: [0, 3000]. Default: 0.
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// The receive latency (ms) of SRT relay. Value range: [0, 3000]. Default: 120.
	RecvLatency *int64 `json:"RecvLatency,omitempty" name:"RecvLatency"`

	// The peer-to-peer latency (ms) of SRT relay. Value range: [0, 3000]. Default: 0.
	PeerLatency *int64 `json:"PeerLatency,omitempty" name:"PeerLatency"`

	// The timeout period (ms) for the SRT relay peer. Value range: [1000, 10000]. Default: 5000.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitempty" name:"PeerIdleTimeout"`

	// The encryption key for SRT relay, which is empty by default, indicating not to encrypt. Only ASCII codes are allowed. Length: [10, 79].
	Passphrase *string `json:"Passphrase,omitempty" name:"Passphrase"`

	// The key length for SRT relay. Valid values: 0 (default), 16, 24, 32.
	PbKeyLen *int64 `json:"PbKeyLen,omitempty" name:"PbKeyLen"`

	// The SRT mode. Valid values: LISTENER, CALLER (default).
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type CreateOutputSrtSettingsDestinations struct {
	// The output IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The output port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type CreateStreamLinkFlowRequestParams struct {
	// Flow name
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// Maximum bandwidth in bps. Valid values: `10000000`, `20000000`, `50000000`
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Flow input group
	InputGroup []*CreateInput `json:"InputGroup,omitempty" name:"InputGroup"`
}

type CreateStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow name
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// Maximum bandwidth in bps. Valid values: `10000000`, `20000000`, `50000000`
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Flow input group
	InputGroup []*CreateInput `json:"InputGroup,omitempty" name:"InputGroup"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStreamLinkFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStreamLinkFlowResponseParams struct {
	// Information of the created flow
	Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateStreamLinkOutputInfoRequestParams struct {
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The output configuration of the flow.
	Output *CreateOutputInfo `json:"Output,omitempty" name:"Output"`
}

type CreateStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The output configuration of the flow.
	Output *CreateOutputInfo `json:"Output,omitempty" name:"Output"`
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
	// The information of the created output.
	Info *DescribeOutput `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteStreamLinkFlowRequestParams struct {
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DeleteStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Output ID
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`
}

type DeleteStreamLinkOutputRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Output ID
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Flow name.
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// Flow status. Valid values: `IDLE`, `RUNNING`
	State *string `json:"State,omitempty" name:"State"`

	// Maximum bandwidth value.
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Input group.
	InputGroup []*DescribeInput `json:"InputGroup,omitempty" name:"InputGroup"`

	// Output group.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OutputGroup []*DescribeOutput `json:"OutputGroup,omitempty" name:"OutputGroup"`
}

type DescribeInput struct {
	// Input ID.
	InputId *string `json:"InputId,omitempty" name:"InputId"`

	// Input name.
	InputName *string `json:"InputName,omitempty" name:"InputName"`

	// Input description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Input protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Input address list.
	InputAddressList []*InputAddress `json:"InputAddressList,omitempty" name:"InputAddressList"`

	// Input IP allowlist.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`

	// SRT configuration information of input.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SRTSettings *DescribeInputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration information of input.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RTPSettings *DescribeInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// Input region.
	InputRegion *string `json:"InputRegion,omitempty" name:"InputRegion"`

	// RTMP configuration information of an input
	RTMPSettings *DescribeInputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`

	// Input failover
	// Note: this field may return `null`, indicating that no valid value was found.
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`
}

type DescribeInputRTMPSettings struct {
	// Path for RTMP stream pushing
	// Note: this field may return `null`, indicating that no valid value was found.
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// StreamKey for RTMP stream pushing
	// Format of an RTMP stream pushing URL: rtmp://IP address:1935/AppName/StreamKey
	StreamKey *string `json:"StreamKey,omitempty" name:"StreamKey"`
}

type DescribeInputRTPSettings struct {
	// Whether it is FEC.
	FEC *string `json:"FEC,omitempty" name:"FEC"`

	// Idle timeout period.
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
}

type DescribeInputSRTSettings struct {
	// The SRT mode.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// Stream ID.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Latency.
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// Receive latency.
	RecvLatency *int64 `json:"RecvLatency,omitempty" name:"RecvLatency"`

	// Peer latency.
	PeerLatency *int64 `json:"PeerLatency,omitempty" name:"PeerLatency"`

	// Peer idle timeout period.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitempty" name:"PeerIdleTimeout"`

	// Decryption key.
	Passphrase *string `json:"Passphrase,omitempty" name:"Passphrase"`

	// Key length.
	PbKeyLen *int64 `json:"PbKeyLen,omitempty" name:"PbKeyLen"`

	// The SRT peer address.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SourceAddresses []*SRTSourceAddressResp `json:"SourceAddresses,omitempty" name:"SourceAddresses"`
}

type DescribeOutput struct {
	// Output ID.
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`

	// Output name.
	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`

	// Output type.
	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`

	// Output description.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Output protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Output destination address information list.
	OutputAddressList []*OutputAddress `json:"OutputAddressList,omitempty" name:"OutputAddressList"`

	// Output region.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OutputRegion *string `json:"OutputRegion,omitempty" name:"OutputRegion"`

	// SRT configuration information of output.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SRTSettings *DescribeOutputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration information of output.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RTPSettings *DescribeOutputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// RTMP configuration information of output.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RTMPSettings *DescribeOutputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`

	// RTMP pull configuration of the output
	// Note: This field may return `null`, indicating that no valid value was found.
	RTMPPullSettings *DescribeOutputRTMPPullSettings `json:"RTMPPullSettings,omitempty" name:"RTMPPullSettings"`

	// CIDR allowlist
	// This parameter is valid if `Protocol` is set to `RTMP_PULL`. If this parameter is left empty, there is no restriction on clients’ IP addresses.
	// Note: This field may return `null`, indicating that no valid value was found.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`
}

type DescribeOutputRTMPPullServerUrl struct {
	// `tcUrl` of the RTMP pull URL
	TcUrl *string `json:"TcUrl,omitempty" name:"TcUrl"`

	// Stream key of the RTMP pull URL
	StreamKey *string `json:"StreamKey,omitempty" name:"StreamKey"`
}

type DescribeOutputRTMPPullSettings struct {
	// List of pull URLs
	// Note: This field may return `null`, indicating that no valid value was found.
	ServerUrls []*DescribeOutputRTMPPullServerUrl `json:"ServerUrls,omitempty" name:"ServerUrls"`
}

type DescribeOutputRTMPSettings struct {
	// Idle timeout period.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`

	// Chunk size.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ChunkSize *int64 `json:"ChunkSize,omitempty" name:"ChunkSize"`

	// Destination address information list of RTMP push.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Destinations []*RTMPAddressDestination `json:"Destinations,omitempty" name:"Destinations"`
}

type DescribeOutputRTPSettings struct {
	// Destination address information list of RTP push.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Destinations []*RTPAddressDestination `json:"Destinations,omitempty" name:"Destinations"`

	// Whether it is FEC.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FEC *string `json:"FEC,omitempty" name:"FEC"`

	// Idle timeout period.
	// Note: this field may return null, indicating that no valid values can be obtained.
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
}

type DescribeOutputSRTSettings struct {
	// A list of the destination addresses for relay. This parameter is valid if `Mode` is `CALLER`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Destinations []*SRTAddressDestination `json:"Destinations,omitempty" name:"Destinations"`

	// Stream ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Latency.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// Receive latency.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RecvLatency *int64 `json:"RecvLatency,omitempty" name:"RecvLatency"`

	// Peer latency.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeerLatency *int64 `json:"PeerLatency,omitempty" name:"PeerLatency"`

	// Peer idle timeout period.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitempty" name:"PeerIdleTimeout"`

	// Encryption key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Passphrase *string `json:"Passphrase,omitempty" name:"Passphrase"`

	// Encryption key length.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PbKeyLen *int64 `json:"PbKeyLen,omitempty" name:"PbKeyLen"`

	// The SRT mode.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The server’s listen address, which is valid if `Mode` is `LISTENER`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SourceAddresses []*OutputSRTSourceAddressResp `json:"SourceAddresses,omitempty" name:"SourceAddresses"`
}

// Predefined struct for user
type DescribeStreamLinkFlowLogsRequestParams struct {
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type []*string `json:"Type,omitempty" name:"Type"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline []*string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The page size. Value range: [1, 1000]. Default: 100.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Whether to sort the records by timestamp in descending or ascending order. Valid values: desc (default), asc.
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// The page number. Value range: [1, 1000]. Default: 1.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
}

type DescribeStreamLinkFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type []*string `json:"Type,omitempty" name:"Type"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline []*string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The page size. Value range: [1, 1000]. Default: 100.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Whether to sort the records by timestamp in descending or ascending order. Valid values: desc (default), asc.
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// The page number. Value range: [1, 1000]. Default: 1.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`
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
	// A list of the logs.
	Infos []*FlowLogInfo `json:"Infos,omitempty" name:"Infos"`

	// The current page number.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// The number of records per page.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// The total number of records.
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// The total number of pages.
	TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeStreamLinkFlowMediaStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowMediaStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowMediaStatisticsResponseParams struct {
	// A list of the media data.
	Infos []*FlowMediaInfo `json:"Infos,omitempty" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The IDs of the inputs to query. If this parameter and `OutputIds` are both empty, all inputs and outputs are queried.
	InputIds []*string `json:"InputIds,omitempty" name:"InputIds"`

	// The IDs of the outputs to query. If this parameter and `OutputIds` are both empty, all inputs and outputs are queried.
	OutputIds []*string `json:"OutputIds,omitempty" name:"OutputIds"`
}

type DescribeStreamLinkFlowRealtimeStatusRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The IDs of the inputs to query. If this parameter and `OutputIds` are both empty, all inputs and outputs are queried.
	InputIds []*string `json:"InputIds,omitempty" name:"InputIds"`

	// The IDs of the outputs to query. If this parameter and `OutputIds` are both empty, all inputs and outputs are queried.
	OutputIds []*string `json:"OutputIds,omitempty" name:"OutputIds"`
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
	// The timestamp (seconds) of the query.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// A list of the real-time data.
	Datas []*FlowRealtimeStatusItem `json:"Datas,omitempty" name:"Datas"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
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
	// Configuration information of a flow
	Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`
}

type DescribeStreamLinkFlowSRTStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowSRTStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowSRTStatisticsResponseParams struct {
	// A list of the SRT streaming performance data.
	Infos []*FlowSRTInfo `json:"Infos,omitempty" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeStreamLinkFlowStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Whether to query the inputs or outputs. Valid values: input, output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// Whether to query the primary or backup pipeline. Valid values: 0, 1.
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The query interval. Valid values: 5s, 1min, 5min, 15min.
	Period *string `json:"Period,omitempty" name:"Period"`

	// The start time for query, which is 1 hour ago by default. You can query statistics in the last 7 days.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// The end time for query, which is 1 hour after the start time by default. The longest time range allowed for query is 24 hours.
	// It must be in UTC format, such as `2020-01-01T12:00:00Z`.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamLinkFlowStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamLinkFlowStatisticsResponseParams struct {
	// A list of the media data.
	Infos []*FlowStatisticsArray `json:"Infos,omitempty" name:"Infos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Number of the current page. Default value: `1`
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Default value: `10`
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeStreamLinkFlowsRequest struct {
	*tchttp.BaseRequest
	
	// Number of the current page. Default value: `1`
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Default value: `10`
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
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
	// List of the configuration information of the flows
	Infos []*DescribeFlow `json:"Infos,omitempty" name:"Infos"`

	// Number of the current page
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Total number of entries
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// Total number of pages
	TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// StreamLink region information
	Info *StreamLinkRegionInfo `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type FlowAudio struct {
	// The frame rate.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// The bitrate (bps).
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// The audio PID.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type FlowLogInfo struct {
	// The timestamp (seconds).
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// Whether it is an input or output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input or output ID.
	InputOutputId *string `json:"InputOutputId,omitempty" name:"InputOutputId"`

	// The protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The event code.
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// The event information.
	EventMessage *string `json:"EventMessage,omitempty" name:"EventMessage"`

	// The peer IP.
	RemoteIp *string `json:"RemoteIp,omitempty" name:"RemoteIp"`

	// The peer port.
	RemotePort *string `json:"RemotePort,omitempty" name:"RemotePort"`

	// Whether it is a primary or backup pipeline. Valid values: 0 (primary), 1 (backup).
	Pipeline *string `json:"Pipeline,omitempty" name:"Pipeline"`

	// The input or output name.
	InputOutputName *string `json:"InputOutputName,omitempty" name:"InputOutputName"`
}

type FlowMediaAudio struct {
	// The frame rate.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// The bitrate (bps).
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// The audio PID.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// The ID of a push session.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type FlowMediaInfo struct {
	// The timestamp (seconds).
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The total bandwidth.
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// The video data of the flow.
	Video []*FlowMediaVideo `json:"Video,omitempty" name:"Video"`

	// The audio data of the flow.
	Audio []*FlowMediaAudio `json:"Audio,omitempty" name:"Audio"`

	// The ID of a push session.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The client IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
}

type FlowMediaVideo struct {
	// The frame rate.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// The bitrate (bps).
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// The video PID.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// The ID of a push session.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type FlowRealtimeStatusCommon struct {
	// The connection status. Valid values: Connected, Waiting, Idle.
	State *string `json:"State,omitempty" name:"State"`

	// The connection mode. Valid values: Listener, Caller.
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// The connected time.
	ConnectedTime *int64 `json:"ConnectedTime,omitempty" name:"ConnectedTime"`

	// The real-time bitrate (bps).
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// The number of retries.
	Reconnections *int64 `json:"Reconnections,omitempty" name:"Reconnections"`
}

type FlowRealtimeStatusItem struct {
	// Whether it is an input or output. Valid values: Input, Output.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The input ID, which is not empty if `Type` is `Input`.
	InputId *string `json:"InputId,omitempty" name:"InputId"`

	// The output ID, which is not empty if `Type` is `Output`.
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`

	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The protocol used. Valid values: SRT, RTP, RTMP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The common status information.
	CommonStatus *FlowRealtimeStatusCommon `json:"CommonStatus,omitempty" name:"CommonStatus"`

	// This parameter is returned if `Protocol` is `SRT`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SRTStatus *FlowRealtimeStatusSRT `json:"SRTStatus,omitempty" name:"SRTStatus"`

	// This parameter is returned if `Protocol` is `RTMP`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RTMPStatus *FlowRealtimeStatusRTMP `json:"RTMPStatus,omitempty" name:"RTMPStatus"`

	// The server IP.
	ConnectServerIP *string `json:"ConnectServerIP,omitempty" name:"ConnectServerIP"`

	// This parameter is returned if the RTP protocol is used.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RTPStatus *FlowRealtimeStatusRTP `json:"RTPStatus,omitempty" name:"RTPStatus"`
}

type FlowRealtimeStatusRTMP struct {
	// The video frame rate.
	VideoFPS *int64 `json:"VideoFPS,omitempty" name:"VideoFPS"`

	// The audio frame rate.
	AudioFPS *int64 `json:"AudioFPS,omitempty" name:"AudioFPS"`
}

type FlowRealtimeStatusRTP struct {
	// The number of packets transmitted.
	Packets *int64 `json:"Packets,omitempty" name:"Packets"`
}

type FlowRealtimeStatusSRT struct {
	// The latency (ms).
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// RTT (ms).
	RTT *int64 `json:"RTT,omitempty" name:"RTT"`

	// The number of packets sent or received.
	Packets *int64 `json:"Packets,omitempty" name:"Packets"`

	// The packet loss rate.
	PacketLossRate *float64 `json:"PacketLossRate,omitempty" name:"PacketLossRate"`

	// The retransmission rate.
	RetransmitRate *float64 `json:"RetransmitRate,omitempty" name:"RetransmitRate"`

	// The number of packets dropped.
	DroppedPackets *int64 `json:"DroppedPackets,omitempty" name:"DroppedPackets"`

	// Whether to encrypt the stream. Valid values: On, Off.
	Encryption *string `json:"Encryption,omitempty" name:"Encryption"`
}

type FlowSRTInfo struct {
	// The timestamp (seconds).
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The packet loss rate for sending.
	SendPacketLossRate *int64 `json:"SendPacketLossRate,omitempty" name:"SendPacketLossRate"`

	// The retry rate for sending.
	SendRetransmissionRate *int64 `json:"SendRetransmissionRate,omitempty" name:"SendRetransmissionRate"`

	// The packet loss rate for receiving.
	RecvPacketLossRate *int64 `json:"RecvPacketLossRate,omitempty" name:"RecvPacketLossRate"`

	// The retry rate for receiving.
	RecvRetransmissionRate *int64 `json:"RecvRetransmissionRate,omitempty" name:"RecvRetransmissionRate"`

	// The peer RTT.
	RTT *int64 `json:"RTT,omitempty" name:"RTT"`

	// The ID of a push session.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The number of dropped packets for sending.
	SendPacketDropNumber *int64 `json:"SendPacketDropNumber,omitempty" name:"SendPacketDropNumber"`

	// The number of dropped packets for receiving.
	RecvPacketDropNumber *int64 `json:"RecvPacketDropNumber,omitempty" name:"RecvPacketDropNumber"`
}

type FlowStatistics struct {
	// The session ID.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The peer IP.
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// The total bandwidth.
	Network *int64 `json:"Network,omitempty" name:"Network"`

	// The video data.
	Video []*FlowVideo `json:"Video,omitempty" name:"Video"`

	// The audio data.
	Audio []*FlowAudio `json:"Audio,omitempty" name:"Audio"`
}

type FlowStatisticsArray struct {
	// The timestamp.
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// The statistics of all the sessions.
	FlowStatistics []*FlowStatistics `json:"FlowStatistics,omitempty" name:"FlowStatistics"`
}

type FlowVideo struct {
	// The frame rate.
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// The bitrate (bps).
	Rate *int64 `json:"Rate,omitempty" name:"Rate"`

	// The audio PID.
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type InputAddress struct {
	// Input address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Input address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type ModifyInput struct {
	// The input ID.
	InputId *string `json:"InputId,omitempty" name:"InputId"`

	// The input name.
	InputName *string `json:"InputName,omitempty" name:"InputName"`

	// The description of the input.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The IP addresses (CIDR) allowed to push streams.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`

	// The SRT configuration information.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// The RTP configuration information.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// The input protocol. Valid values: SRT, RTP, RTMP.
	// If there is an RTP input, the output must be RTP.
	// If there is an RTMP input, the output must be SRT or RTMP.
	// If there is an SRT input, the output must be SRT.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Whether to enable input failover. Valid values: OPEN, CLOSE.
	FailOver *string `json:"FailOver,omitempty" name:"FailOver"`
}

type ModifyOutputInfo struct {
	// The ID of the output to modify.
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`

	// The output name.
	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`

	// The description of the output.
	Description *string `json:"Description,omitempty" name:"Description"`

	// The output protocol. Valid values: SRT, RTP, RTMP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// The SRT relay configuration.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// The RTP relay configuration.
	RTPSettings *CreateOutputInfoRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// The RTMP relay configuration.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`

	// The IP allowlist. The address must be in CIDR format, such as `0.0.0.0/0`.
	// This parameter is valid if `Protocol` is set to `RTMP_PULL`. If it is left empty, there is no restriction on clients’ IP addresses.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`
}

// Predefined struct for user
type ModifyStreamLinkFlowRequestParams struct {
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Name of the flow to modify
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
}

type ModifyStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Name of the flow to modify
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The input information to modify.
	Input *ModifyInput `json:"Input,omitempty" name:"Input"`
}

type ModifyStreamLinkInputRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The input information to modify.
	Input *ModifyInput `json:"Input,omitempty" name:"Input"`
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
	// The input information after modification.
	Info *DescribeInput `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The output configuration to modify.
	Output *ModifyOutputInfo `json:"Output,omitempty" name:"Output"`
}

type ModifyStreamLinkOutputInfoRequest struct {
	*tchttp.BaseRequest
	
	// The flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// The output configuration to modify.
	Output *ModifyOutputInfo `json:"Output,omitempty" name:"Output"`
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
	// The output configuration after modification.
	Info *DescribeOutput `json:"Info,omitempty" name:"Info"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Output destination IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type OutputSRTSourceAddressResp struct {
	// The listen IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The listen port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type RTMPAddressDestination struct {
	// Destination URL of RTMP push in the format of 'rtmp://domain/live'.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Destination `StreamKey` of RTMP push in the format of 'streamid?key=value'.
	StreamKey *string `json:"StreamKey,omitempty" name:"StreamKey"`
}

type RTPAddressDestination struct {
	// Push destination address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Push destination address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type RegionInfo struct {
	// Region name
	Name *string `json:"Name,omitempty" name:"Name"`
}

type SRTAddressDestination struct {
	// Destination address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Destination address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type SRTSourceAddressReq struct {
	// The peer IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The peer port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type SRTSourceAddressResp struct {
	// The peer IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// The peer port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

// Predefined struct for user
type StartStreamLinkFlowRequestParams struct {
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type StartStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type StopStreamLinkFlowRequest struct {
	*tchttp.BaseRequest
	
	// Flow ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
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
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// List of StreamLink regions
	Regions []*RegionInfo `json:"Regions,omitempty" name:"Regions"`
}