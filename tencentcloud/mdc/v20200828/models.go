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

	// Input protocol. Valid values: SRT, RTP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Input description. Length: [0, 255].
	Description *string `json:"Description,omitempty" name:"Description"`

	// Allowlist of input IPs in CIDR format.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`

	// SRT configuration information of input.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration information of input.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`
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

type CreateMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow name.
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// Maximum bandwidth in bps. Valid values: 10000000, 20000000, 50000000.
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`

	// Flow input group.
	InputGroup []*CreateInput `json:"InputGroup,omitempty" name:"InputGroup"`
}

func (r *CreateMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowName")
	delete(f, "MaxBandwidth")
	delete(f, "InputGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of the created flow.
		Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMediaConnectOutputRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Output configuration of a flow.
	Output *CreateOutput `json:"Output,omitempty" name:"Output"`
}

func (r *CreateMediaConnectOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaConnectOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMediaConnectOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMediaConnectOutputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Information of the created output.
		Info *DescribeOutput `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaConnectOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMediaConnectOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOutput struct {

	// Output name.
	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`

	// Output description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Output protocol.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Output region.
	OutputRegion *string `json:"OutputRegion,omitempty" name:"OutputRegion"`

	// SRT configuration of output.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration of output.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// RTMP configuration of output.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`
}

type CreateOutputRTMPSettings struct {

	// Push destination address. You can enter one or two addresses.
	Destinations []*CreateOutputRtmpSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// RTMP chunk size. Value range: [4096, 40960].
	ChunkSize *int64 `json:"ChunkSize,omitempty" name:"ChunkSize"`
}

type CreateOutputRTPSettings struct {

	// Push destination address. You can enter one or two addresses.
	Destinations *CreateOutputRTPSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// Only `none` can be entered.
	FEC *string `json:"FEC,omitempty" name:"FEC"`

	// Idle timeout period.
	IdleTimeout *int64 `json:"IdleTimeout,omitempty" name:"IdleTimeout"`
}

type CreateOutputRTPSettingsDestinations struct {

	// Push destination IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Push destination port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type CreateOutputRtmpSettingsDestinations struct {

	// Push URL in the format of `rtmp://domain/live`.
	Url *string `json:"Url,omitempty" name:"Url"`

	// Push `StreamKey` in the format of `stream?key=value`.
	StreamKey *string `json:"StreamKey,omitempty" name:"StreamKey"`
}

type CreateOutputSrtSettings struct {

	// Push destination address. Please configure one or two addresses.
	Destinations []*CreateOutputSrtSettingsDestinations `json:"Destinations,omitempty" name:"Destinations"`

	// Stream ID of SRT push.
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// Total latency of SRT push.
	Latency *int64 `json:"Latency,omitempty" name:"Latency"`

	// Receive latency of SRT push.
	RecvLatency *int64 `json:"RecvLatency,omitempty" name:"RecvLatency"`

	// Peer latency of SRT push.
	PeerLatency *int64 `json:"PeerLatency,omitempty" name:"PeerLatency"`

	// Peer idle timeout period of SRT push.
	PeerIdleTimeout *int64 `json:"PeerIdleTimeout,omitempty" name:"PeerIdleTimeout"`

	// Encryption key of SRT push.
	Passphrase *string `json:"Passphrase,omitempty" name:"Passphrase"`

	// Key length of SRT push.
	PbKeyLen *int64 `json:"PbKeyLen,omitempty" name:"PbKeyLen"`
}

type CreateOutputSrtSettingsDestinations struct {

	// Output IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Output port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DeleteMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DeleteMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaConnectOutputRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Output ID.
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`
}

func (r *DeleteMediaConnectOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaConnectOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "OutputId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMediaConnectOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaConnectOutputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaConnectOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMediaConnectOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlow struct {

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Flow name.
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// Flow status.
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

type DescribeMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration information of a flow.
		Info *DescribeFlow `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaConnectFlowsRequest struct {
	*tchttp.BaseRequest

	// Number of current pages. Default value: 1.
	PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

	// Number of entries per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeMediaConnectFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaConnectFlowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNum")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMediaConnectFlowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaConnectFlowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Configuration information list of a flow.
		Infos []*DescribeFlow `json:"Infos,omitempty" name:"Infos"`

		// Number of current pages.
		PageNum *int64 `json:"PageNum,omitempty" name:"PageNum"`

		// Number of entries per page.
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// Total number.
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// Total number of pages.
		TotalPage *int64 `json:"TotalPage,omitempty" name:"TotalPage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaConnectFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMediaConnectFlowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type InputAddress struct {

	// Input address IP.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Input address port.
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type ModifyInput struct {

	// Input ID.
	InputId *string `json:"InputId,omitempty" name:"InputId"`

	// Input name.
	InputName *string `json:"InputName,omitempty" name:"InputName"`

	// Input description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Allowed push IP in CIDR format.
	AllowIpList []*string `json:"AllowIpList,omitempty" name:"AllowIpList"`

	// SRT configuration information.
	SRTSettings *CreateInputSRTSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// RTP configuration information.
	RTPSettings *CreateInputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`
}

type ModifyMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Name of the flow to be modified.
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`
}

func (r *ModifyMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "FlowName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaConnectInputRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Information of the input to be modified.
	Input *ModifyInput `json:"Input,omitempty" name:"Input"`
}

func (r *ModifyMediaConnectInputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectInputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaConnectInputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaConnectInputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Input information after modification.
		Info *DescribeInput `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaConnectInputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectInputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaConnectOutputRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// Configuration of the output to be modified.
	Output *ModifyOutput `json:"Output,omitempty" name:"Output"`
}

func (r *ModifyMediaConnectOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	delete(f, "Output")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMediaConnectOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaConnectOutputResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Output configuration after modification.
		Info *DescribeOutput `json:"Info,omitempty" name:"Info"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaConnectOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMediaConnectOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyOutput struct {

	// ID of the output to be modified.
	OutputId *string `json:"OutputId,omitempty" name:"OutputId"`

	// Output name.
	OutputName *string `json:"OutputName,omitempty" name:"OutputName"`

	// Output description.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Output push protocol. Valid values: SRT, RTMP.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// Configuration of SRT push.
	SRTSettings *CreateOutputSrtSettings `json:"SRTSettings,omitempty" name:"SRTSettings"`

	// Configuration of RTP push.
	RTPSettings *CreateOutputRTPSettings `json:"RTPSettings,omitempty" name:"RTPSettings"`

	// Configuration of RTMP push.
	RTMPSettings *CreateOutputRTMPSettings `json:"RTMPSettings,omitempty" name:"RTMPSettings"`
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

type StartMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *StartMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StopMediaConnectFlowRequest struct {
	*tchttp.BaseRequest

	// Flow ID.
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *StopMediaConnectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaConnectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMediaConnectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopMediaConnectFlowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMediaConnectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMediaConnectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
