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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
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

type CreateStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of the created flow
		Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteStreamLinkOutputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// Push destination address information list.
	// Note: this field may return null, indicating that no valid values can be obtained.
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

type DescribeStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration information of a flow
		Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeStreamLinkFlowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type InputAddress struct {

	// Input address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Input address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
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

type ModifyStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type OutputAddress struct {

	// Output destination IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`
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

type SRTAddressDestination struct {

	// Destination address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Destination address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
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

type StartStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type StopStreamLinkFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
