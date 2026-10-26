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

package v20230517

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AddDeviceData struct {
	// <p>Device iD</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device code (for national standard devices, it is the 20-digit National Standard Encoding we generated for the device; for rtmp devices, it is the 10-digit device code)</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Device name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device access protocol. 1: RTMP; 2: GB.</p><p>Enumeration values:</p><ul><li>1: RTMP</li><li>2: GB</li><li>6: ISUP</li></ul>
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// <p>Device type. 1: IPC. 2: NVR</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Device access service node ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Device access service node name</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Device streaming protocol. 1: UDP, 2: TCP</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device status. 0: unregistered, 1: online, 2: offline, 3: disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Affiliated organization ID of the device</p>
	OrganizationId *int64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Device access gateway ID (deprecated, reserved for compatibility, can be ignored)</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>Gateway access protocol type (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Device access IP (deprecated, reserved for compatibility, can be ignored)</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device Port (deprecated, reserved for compatibility, can be ignored)</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username (deprecated, reserved for compatibility, can be ignored)</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>User ID</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type AddOrgData struct {
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Organization name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Organization parent node ID
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Organization hierarchy
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// User ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Organizational structure
	ParentIds *string `json:"ParentIds,omitnil,omitempty" name:"ParentIds"`

	// Total number of devices
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of devices online
	Online *int64 `json:"Online,omitnil,omitempty" name:"Online"`
}

// Predefined struct for user
type AddOrganizationRequestParams struct {
	// Organization name. Only Chinese, English, digits, spaces, Chinese and English parentheses, underscores (_), and hyphens (-) are supported. The length cannot exceed 64 characters, and the organization name must be unique.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Organization parent node ID (obtain from the DescribeOrganization API; enter 0 for the root organization)
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type AddOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Organization name. Only Chinese, English, digits, spaces, Chinese and English parentheses, underscores (_), and hyphens (-) are supported. The length cannot exceed 64 characters, and the organization name must be unique.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Organization parent node ID (obtain from the DescribeOrganization API; enter 0 for the root organization)
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *AddOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationResponseParams struct {
	// Interface return data for adding an organizational unit
	Data *AddOrgData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationResponseParams `json:"Response"`
}

func (r *AddOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordBackupPlanData struct {
	// Recording cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Recording cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Description of the recording cloud migration plan
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Recording cloud migration plan status. `1`: in normal use. `0`: deleting, unavailable.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// Creation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Modification time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

// Predefined struct for user
type AddRecordBackupPlanRequestParams struct {
	// Recording template ID (template ID associated with the recording schedule, obtained from the ListRecordBackupTemplates API for querying the video cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Recording schedule name (supports Chinese, English, digits, _ and -. The length cannot exceed 32 characters. The plan name must be globally unique and cannot be empty or duplicated.)
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video recording plan description (only supports Chinese, English, digits, _, -, length not exceeding 128 characters)
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Lifecycle (video recording file lifecycle settings, the time when cold and hot storage is managed for files)
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Channel and the device to which the channel belongs (channel information of the device to which a video recording is added. Add no more than 5,000 channels at a time, including the number of channels under the organizational unit directory).
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Add all device channels under the organizational unit directory (Json array, can be empty, total number of channels no more than 5000, including the number of Channel fields)
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

type AddRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// Recording template ID (template ID associated with the recording schedule, obtained from the ListRecordBackupTemplates API for querying the video cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Recording schedule name (supports Chinese, English, digits, _ and -. The length cannot exceed 32 characters. The plan name must be globally unique and cannot be empty or duplicated.)
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video recording plan description (only supports Chinese, English, digits, _, -, length not exceeding 128 characters)
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Lifecycle (video recording file lifecycle settings, the time when cold and hot storage is managed for files)
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Channel and the device to which the channel belongs (channel information of the device to which a video recording is added. Add no more than 5,000 channels at a time, including the number of channels under the organizational unit directory).
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Add all device channels under the organizational unit directory (Json array, can be empty, total number of channels no more than 5000, including the number of Channel fields)
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

func (r *AddRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "PlanName")
	delete(f, "Describe")
	delete(f, "LifeCycle")
	delete(f, "Channels")
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordBackupPlanResponseParams struct {
	// Returned data.
	Data *AddRecordBackupPlanData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordBackupPlanResponseParams `json:"Response"`
}

func (r *AddRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordBackupTemplateData struct {
	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Template name</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (set by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (set by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Creation time.</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>Update time.</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// <p>Cycle offset. Default value: 1, which means to pull yesterday's video recording from the device. 0 means to pull today's video recording from the device. 3 means to pull the video recording from the device from the day before yesterday.</p><p>Default value: 1</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

// Predefined struct for user
type AddRecordBackupTemplateRequestParams struct {
	// <p>Template name (supports only Chinese, English, number, _, -, length not exceeding 32 characters. Template name must be globally unique, cannot be empty, and must be unique)</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (settings can be done by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (settings can be done by week, support setting multiple time periods per day, each time period not less than 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Cycle offset. Default: 1. It means to pull yesterday's video recording to the device. 0 means to pull today's video recording to the device. 3 means to pull the video recording to the device from the day before yesterday.</p><p>Value range: [0, 90]</p><p>Default value: 1</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

type AddRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template name (supports only Chinese, English, number, _, -, length not exceeding 32 characters. Template name must be globally unique, cannot be empty, and must be unique)</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (settings can be done by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (settings can be done by week, support setting multiple time periods per day, each time period not less than 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Cycle offset. Default: 1. It means to pull yesterday's video recording to the device. 0 means to pull today's video recording to the device. 3 means to pull the video recording to the device from the day before yesterday.</p><p>Value range: [0, 90]</p><p>Default value: 1</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

func (r *AddRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TimeSections")
	delete(f, "DevTimeSections")
	delete(f, "Scale")
	delete(f, "DayOffset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordBackupTemplateResponseParams struct {
	// <p>Response data.</p>
	Data *AddRecordBackupTemplateData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *AddRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordPlanRequestParams struct {
	// Real-time cloud migration plan name. It supports Chinese, English, digits, _, and -, with a length not exceeding 32 characters. The plan name must be globally unique and cannot be empty or repeated.
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Real-time cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Cloud plan description. Supports only Chinese, English, digits, _, and -. Length not exceeding 128 characters.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Stream type. default: The stream type is not specified, and the default push type of the device is used. main: Main stream. sub: Substream. Other values are customized based on the device capacity. If not specified, the default type is used. The length cannot exceed 32 bytes.
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Channel information of the device to which video recording is added. You can add no more than 5000 channels only once, including the number of channels under the organizational unit directory.
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Add all device channels under the organization. Json array, can be empty. The total number of channels cannot exceed 5,000 (including the number of Channel fields).
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Video recording supplement mode (0: disabled, 1: enabled). If this field is not specified, it is not enabled by default.
	RepairMode *int64 `json:"RepairMode,omitnil,omitempty" name:"RepairMode"`
}

type AddRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// Real-time cloud migration plan name. It supports Chinese, English, digits, _, and -, with a length not exceeding 32 characters. The plan name must be globally unique and cannot be empty or repeated.
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Real-time cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Cloud plan description. Supports only Chinese, English, digits, _, and -. Length not exceeding 128 characters.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Stream type. default: The stream type is not specified, and the default push type of the device is used. main: Main stream. sub: Substream. Other values are customized based on the device capacity. If not specified, the default type is used. The length cannot exceed 32 bytes.
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Channel information of the device to which video recording is added. You can add no more than 5000 channels only once, including the number of channels under the organizational unit directory.
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Add all device channels under the organization. Json array, can be empty. The total number of channels cannot exceed 5,000 (including the number of Channel fields).
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Video recording supplement mode (0: disabled, 1: enabled). If this field is not specified, it is not enabled by default.
	RepairMode *int64 `json:"RepairMode,omitnil,omitempty" name:"RepairMode"`
}

func (r *AddRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanName")
	delete(f, "TemplateId")
	delete(f, "LifeCycle")
	delete(f, "Describe")
	delete(f, "StreamType")
	delete(f, "Channels")
	delete(f, "OrganizationId")
	delete(f, "RepairMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordPlanResponseParams struct {
	// Returned result
	Data *RecordPlanOptData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordPlanResponseParams `json:"Response"`
}

func (r *AddRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddRecordRetrieveTaskData struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time of video recording retrieval
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Retrieval end time of the video recording
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Retrieval mode. 1: ultrafast mode. Others not supported currently.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Replica validity period
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// Task status. 0: retrieved, 1: retrieving, 2: pending retrieval
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Retrieval capacity, in MB.
	Capacity *float64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// Task description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`
}

// Predefined struct for user
type AddRecordRetrieveTaskRequestParams struct {
	// Task name. It supports Chinese, English, digits, _, and -, with a length not exceeding 32 characters. The name must be globally unique and cannot be empty or repeated.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time for retrieving the video recording in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for retrieving the video recording, in UTC seconds, for example: 1662114146. The period between the start and end times can be up to one day and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Retrieval mode. 1: ultrafast mode. Others not supported currently.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Retrieval replica validity period of the video recording. Minimum: 1 day. Maximum: 365 days.
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// Device channel. A task can contain up to 32 device channels.
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Retrieve the task description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`
}

type AddRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task name. It supports Chinese, English, digits, _, and -, with a length not exceeding 32 characters. The name must be globally unique and cannot be empty or repeated.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time for retrieving the video recording in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time for retrieving the video recording, in UTC seconds, for example: 1662114146. The period between the start and end times can be up to one day and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Retrieval mode. 1: ultrafast mode. Others not supported currently.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Retrieval replica validity period of the video recording. Minimum: 1 day. Maximum: 365 days.
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// Device channel. A task can contain up to 32 device channels.
	Channels []*ChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Retrieve the task description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`
}

func (r *AddRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Mode")
	delete(f, "Expiration")
	delete(f, "Channels")
	delete(f, "Describe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordRetrieveTaskResponseParams struct {
	// Returned results
	Data *AddRecordRetrieveTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *AddRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordTemplateRequestParams struct {
	// Template name. It can contain only Chinese characters, letters, digits, underscores (_), and hyphens (-), with a length not exceeding 32 characters. The template name must be globally unique and cannot be empty or duplicated.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cloud upload time period. Settings can be done by week. Multiple time periods can be set for each day, with each time period not less than 10 minutes.
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`
}

type AddRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name. It can contain only Chinese characters, letters, digits, underscores (_), and hyphens (-), with a length not exceeding 32 characters. The template name must be globally unique and cannot be empty or duplicated.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cloud upload time period. Settings can be done by week. Multiple time periods can be set for each day, with each time period not less than 10 minutes.
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`
}

func (r *AddRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TimeSections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddRecordTemplateResponseParams struct {
	// Returned result
	Data *RecordTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *AddRecordTemplateResponseParams `json:"Response"`
}

func (r *AddRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddStreamAuthData struct {
	// Authentication configuration ID (uuid).
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Whether to enable playback authentication. Valid values: `1`: enable, `0`: disable.
	PullState *int64 `json:"PullState,omitnil,omitempty" name:"PullState"`

	// Playback key (alphanumeric only, length: 0-10 characters)
	PullSecret *string `json:"PullSecret,omitnil,omitempty" name:"PullSecret"`

	// Playback expiration time (unit: minute)
	PullExpired *int64 `json:"PullExpired,omitnil,omitempty" name:"PullExpired"`

	// Whether to enable stream authentication (1: enable, 0: disable)
	PushState *int64 `json:"PushState,omitnil,omitempty" name:"PushState"`

	// Stream key (only alphanumeric characters are supported, with a length of 0-10 characters)
	PushSecret *string `json:"PushSecret,omitnil,omitempty" name:"PushSecret"`

	// Push expiration time (unit: minute)
	PushExpired *int64 `json:"PushExpired,omitnil,omitempty" name:"PushExpired"`

	// User ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

// Predefined struct for user
type AddStreamAuthRequestParams struct {
	// Authentication configuration ID (uuid).
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Whether to enable playback authentication (1: enable, 0: disable)
	PullState *int64 `json:"PullState,omitnil,omitempty" name:"PullState"`

	// Playback key (only alphanumeric characters are supported, with a length of 0-10 characters)
	PullSecret *string `json:"PullSecret,omitnil,omitempty" name:"PullSecret"`

	// Playback expiration time (unit: minute)
	PullExpired *int64 `json:"PullExpired,omitnil,omitempty" name:"PullExpired"`

	// Whether to enable stream authentication (1: enable, 0: disable)
	PushState *int64 `json:"PushState,omitnil,omitempty" name:"PushState"`

	// Stream key (alphanumeric only, length 0-10 characters)
	PushSecret *string `json:"PushSecret,omitnil,omitempty" name:"PushSecret"`

	// Push expiration time (unit: minute)
	PushExpired *int64 `json:"PushExpired,omitnil,omitempty" name:"PushExpired"`
}

type AddStreamAuthRequest struct {
	*tchttp.BaseRequest
	
	// Authentication configuration ID (uuid).
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Whether to enable playback authentication (1: enable, 0: disable)
	PullState *int64 `json:"PullState,omitnil,omitempty" name:"PullState"`

	// Playback key (only alphanumeric characters are supported, with a length of 0-10 characters)
	PullSecret *string `json:"PullSecret,omitnil,omitempty" name:"PullSecret"`

	// Playback expiration time (unit: minute)
	PullExpired *int64 `json:"PullExpired,omitnil,omitempty" name:"PullExpired"`

	// Whether to enable stream authentication (1: enable, 0: disable)
	PushState *int64 `json:"PushState,omitnil,omitempty" name:"PushState"`

	// Stream key (alphanumeric only, length 0-10 characters)
	PushSecret *string `json:"PushSecret,omitnil,omitempty" name:"PushSecret"`

	// Push expiration time (unit: minute)
	PushExpired *int64 `json:"PushExpired,omitnil,omitempty" name:"PushExpired"`
}

func (r *AddStreamAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddStreamAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "PullState")
	delete(f, "PullSecret")
	delete(f, "PullExpired")
	delete(f, "PushState")
	delete(f, "PushSecret")
	delete(f, "PushExpired")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddStreamAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddStreamAuthResponseParams struct {
	// data for setting push/pull stream authentication
	Data *AddStreamAuthData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddStreamAuthResponse struct {
	*tchttp.BaseResponse
	Response *AddStreamAuthResponseParams `json:"Response"`
}

func (r *AddStreamAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddStreamAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserDeviceRequestParams struct {
	// <p>Device name, only supports Chinese, English, numbers, spaces, Chinese and English parentheses, _, -, length not exceeding 128 characters; (device name no need to be globally unique and can be duplicated)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device access protocol (1: RTMP, 2: GB, 6: ISUP)</p><p>Enumeration values:</p><ul><li>1: RTMP</li><li>2: GB</li><li>6: ISUP</li></ul><p>Default value: 2</p>
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// <p>Device type. 1: IPC; 2: NVR. (If the device access protocol is RTMP, only select IPC as the device type.)</p><p>Enumeration values:</p><ul><li>1: IPC</li><li>2: NVR</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>ID of the organization to which the device belongs. Obtain from the organization query api DescribeOrganization.</p>
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Device access service node ID (Value field obtained from the DescribeRegionDomain API for querying available service nodes of a device)</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Device streaming protocol. 1: UDP, 2: TCP. (Valid for national standard devices. UDP by default if not specified)</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password (required for national standard devices, length of 1-64 characters)</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description, length not exceeding 128 characters</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device access gateway ID (deprecated, reserved for compatibility, can be ignored)</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>Gateway access protocol type (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Device access IP (deprecated, reserved for compatibility, can be ignored)</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device port (deprecated, reserved for compatibility, can be ignored)</p><p>Value range: [1, 65535]</p><p>Unit: port</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username (deprecated, reserved for compatibility, can be ignored)</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Device SN (deprecated, reserved for compatibility, can be ignored)</p>
	SNCode *string `json:"SNCode,omitnil,omitempty" name:"SNCode"`

	// <p>Custom AppName for RTMP publishing address (required for RTMP only, support English, digits, _, -, ., length not exceeding 64 characters)</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Custom StreamName for RTMP publishing address (required for RTMP only, supports English, digits, _, -, ., length not exceeding 64 characters)</p>
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

type AddUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Device name, only supports Chinese, English, numbers, spaces, Chinese and English parentheses, _, -, length not exceeding 128 characters; (device name no need to be globally unique and can be duplicated)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device access protocol (1: RTMP, 2: GB, 6: ISUP)</p><p>Enumeration values:</p><ul><li>1: RTMP</li><li>2: GB</li><li>6: ISUP</li></ul><p>Default value: 2</p>
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// <p>Device type. 1: IPC; 2: NVR. (If the device access protocol is RTMP, only select IPC as the device type.)</p><p>Enumeration values:</p><ul><li>1: IPC</li><li>2: NVR</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>ID of the organization to which the device belongs. Obtain from the organization query api DescribeOrganization.</p>
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Device access service node ID (Value field obtained from the DescribeRegionDomain API for querying available service nodes of a device)</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Device streaming protocol. 1: UDP, 2: TCP. (Valid for national standard devices. UDP by default if not specified)</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password (required for national standard devices, length of 1-64 characters)</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description, length not exceeding 128 characters</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device access gateway ID (deprecated, reserved for compatibility, can be ignored)</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>Gateway access protocol type (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Device access IP (deprecated, reserved for compatibility, can be ignored)</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device port (deprecated, reserved for compatibility, can be ignored)</p><p>Value range: [1, 65535]</p><p>Unit: port</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username (deprecated, reserved for compatibility, can be ignored)</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Device SN (deprecated, reserved for compatibility, can be ignored)</p>
	SNCode *string `json:"SNCode,omitnil,omitempty" name:"SNCode"`

	// <p>Custom AppName for RTMP publishing address (required for RTMP only, support English, digits, _, -, ., length not exceeding 64 characters)</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Custom StreamName for RTMP publishing address (required for RTMP only, supports English, digits, _, -, ., length not exceeding 64 characters)</p>
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`
}

func (r *AddUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AccessProtocol")
	delete(f, "Type")
	delete(f, "OrganizationId")
	delete(f, "ClusterId")
	delete(f, "TransportProtocol")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "GatewayId")
	delete(f, "ProtocolType")
	delete(f, "Ip")
	delete(f, "Port")
	delete(f, "Username")
	delete(f, "SNCode")
	delete(f, "AppName")
	delete(f, "StreamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserDeviceResponseParams struct {
	// <p>Add device response data.</p>
	Data *AddDeviceData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *AddUserDeviceResponseParams `json:"Response"`
}

func (r *AddUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddVideoSnapshotConfigChannelRequestParams struct {
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Channel Id list</p>
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`
}

type BatchAddVideoSnapshotConfigChannelRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Channel Id list</p>
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`
}

func (r *BatchAddVideoSnapshotConfigChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddVideoSnapshotConfigChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Channels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAddVideoSnapshotConfigChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAddVideoSnapshotConfigChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchAddVideoSnapshotConfigChannelResponse struct {
	*tchttp.BaseResponse
	Response *BatchAddVideoSnapshotConfigChannelResponseParams `json:"Response"`
}

func (r *BatchAddVideoSnapshotConfigChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAddVideoSnapshotConfigChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAssignVLPlanChannelsRequestParams struct {
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>Operation command: 1=associate, 2=unbind</p>
	Cmd *int64 `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// <p>Channel ID list, up to 500; pass at least one of this parameter and OrganizationIds (when Cmd=1)</p>
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`

	// <p>List of organization IDs (valid only when Cmd=1; does not allow input when Cmd=2); unfold all channels under the organization to participate in association, and take the union with ChannelIds and remove duplicates</p>
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

type BatchAssignVLPlanChannelsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>Operation command: 1=associate, 2=unbind</p>
	Cmd *int64 `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// <p>Channel ID list, up to 500; pass at least one of this parameter and OrganizationIds (when Cmd=1)</p>
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`

	// <p>List of organization IDs (valid only when Cmd=1; does not allow input when Cmd=2); unfold all channels under the organization to participate in association, and take the union with ChannelIds and remove duplicates</p>
	OrganizationIds []*string `json:"OrganizationIds,omitnil,omitempty" name:"OrganizationIds"`
}

func (r *BatchAssignVLPlanChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAssignVLPlanChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanId")
	delete(f, "Cmd")
	delete(f, "ChannelIds")
	delete(f, "OrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchAssignVLPlanChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchAssignVLPlanChannelsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchAssignVLPlanChannelsResponse struct {
	*tchttp.BaseResponse
	Response *BatchAssignVLPlanChannelsResponseParams `json:"Response"`
}

func (r *BatchAssignVLPlanChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchAssignVLPlanChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteVideoDownloadTaskRequestParams struct {
	// List of local recording download task IDs
	DownloadTaskIds []*string `json:"DownloadTaskIds,omitnil,omitempty" name:"DownloadTaskIds"`
}

type BatchDeleteVideoDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// List of local recording download task IDs
	DownloadTaskIds []*string `json:"DownloadTaskIds,omitnil,omitempty" name:"DownloadTaskIds"`
}

func (r *BatchDeleteVideoDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteVideoDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownloadTaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteVideoDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteVideoDownloadTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteVideoDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteVideoDownloadTaskResponseParams `json:"Response"`
}

func (r *BatchDeleteVideoDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteVideoDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchGetVLPlanSourceFileURLRequestParams struct {
	// <p>Analysis result ID list. Up to 200.</p>
	VLPlanResultIds []*string `json:"VLPlanResultIds,omitnil,omitempty" name:"VLPlanResultIds"`

	// <p>URL validity period (in seconds). Value range: [60, 86400]. Default: 600</p>
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

type BatchGetVLPlanSourceFileURLRequest struct {
	*tchttp.BaseRequest
	
	// <p>Analysis result ID list. Up to 200.</p>
	VLPlanResultIds []*string `json:"VLPlanResultIds,omitnil,omitempty" name:"VLPlanResultIds"`

	// <p>URL validity period (in seconds). Value range: [60, 86400]. Default: 600</p>
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

func (r *BatchGetVLPlanSourceFileURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchGetVLPlanSourceFileURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanResultIds")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchGetVLPlanSourceFileURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchGetVLPlanSourceFileURLResponseParams struct {
	// <p>URL list</p>
	Data []*VLPlanResultSnapshotImageUrlInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Total quantity.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchGetVLPlanSourceFileURLResponse struct {
	*tchttp.BaseResponse
	Response *BatchGetVLPlanSourceFileURLResponseParams `json:"Response"`
}

func (r *BatchGetVLPlanSourceFileURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchGetVLPlanSourceFileURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchGetVideoSnapshotImageUrlData struct {
	// <p>Image Url information</p>
	Images []*ImageUrlInfo `json:"Images,omitnil,omitempty" name:"Images"`

	// <p>Valid period of the image Url</p><p>Unit: seconds</p><p>Default value: 600</p>
	ImageUrlExpires *int64 `json:"ImageUrlExpires,omitnil,omitempty" name:"ImageUrlExpires"`
}

// Predefined struct for user
type BatchGetVideoSnapshotImageUrlRequestParams struct {
	// <p>Image Id</p>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// <p>Valid period of the image URL, unit: seconds</p><p>Value range: [60, 86400]</p><p>Default value: 600</p>
	ImageUrlExpires *int64 `json:"ImageUrlExpires,omitnil,omitempty" name:"ImageUrlExpires"`
}

type BatchGetVideoSnapshotImageUrlRequest struct {
	*tchttp.BaseRequest
	
	// <p>Image Id</p>
	ImageIds []*string `json:"ImageIds,omitnil,omitempty" name:"ImageIds"`

	// <p>Valid period of the image URL, unit: seconds</p><p>Value range: [60, 86400]</p><p>Default value: 600</p>
	ImageUrlExpires *int64 `json:"ImageUrlExpires,omitnil,omitempty" name:"ImageUrlExpires"`
}

func (r *BatchGetVideoSnapshotImageUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchGetVideoSnapshotImageUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageIds")
	delete(f, "ImageUrlExpires")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchGetVideoSnapshotImageUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchGetVideoSnapshotImageUrlResponseParams struct {
	// <p>List of screenshot image URL results</p>
	Data *BatchGetVideoSnapshotImageUrlData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchGetVideoSnapshotImageUrlResponse struct {
	*tchttp.BaseResponse
	Response *BatchGetVideoSnapshotImageUrlResponseParams `json:"Response"`
}

func (r *BatchGetVideoSnapshotImageUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchGetVideoSnapshotImageUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateDeviceData struct {
	// Task ID (used to query task progress in the ListSubTasks API for querying the subtask list of a job)
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type BatchOperateDeviceRequestParams struct {
	// <p>Device ID array (obtain from the device list API ListDevices)</p><p>For parameter values, see <a href="https://www.tencentcloud.com/document/api/1344/95871?from_cn_redirect=1">ListDevices</a></p>
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// <p>Operation command</p><p>Enumeration values:</p><ul><li>enable: Enable</li><li>disable: Disable</li><li>delete: Delete</li><li>sync: Synchronize device channel</li></ul>
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`
}

type BatchOperateDeviceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Device ID array (obtain from the device list API ListDevices)</p><p>For parameter values, see <a href="https://www.tencentcloud.com/document/api/1344/95871?from_cn_redirect=1">ListDevices</a></p>
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// <p>Operation command</p><p>Enumeration values:</p><ul><li>enable: Enable</li><li>disable: Disable</li><li>delete: Delete</li><li>sync: Synchronize device channel</li></ul>
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`
}

func (r *BatchOperateDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchOperateDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	delete(f, "Cmd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchOperateDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchOperateDeviceResponseParams struct {
	// <p>Returned results.</p>
	Data *BatchOperateDeviceData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchOperateDeviceResponse struct {
	*tchttp.BaseResponse
	Response *BatchOperateDeviceResponseParams `json:"Response"`
}

func (r *BatchOperateDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchOperateDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BitRateInfo struct {
	// Channel Id
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Bitrate, in kbps.
	Bitrate *float64 `json:"Bitrate,omitnil,omitempty" name:"Bitrate"`
}

// Predefined struct for user
type CallISAPIRequestParams struct {
	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// url resource
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Input parameter
	InputData *string `json:"InputData,omitnil,omitempty" name:"InputData"`
}

type CallISAPIRequest struct {
	*tchttp.BaseRequest
	
	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// url resource
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Input parameter
	InputData *string `json:"InputData,omitnil,omitempty" name:"InputData"`
}

func (r *CallISAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallISAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Url")
	delete(f, "InputData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallISAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallISAPIResponseParams struct {
	// Returned data.
	Data *ISAPIOutputData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CallISAPIResponse struct {
	*tchttp.BaseResponse
	Response *CallISAPIResponseParams `json:"Response"`
}

func (r *CallISAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallISAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallbackURL struct {
	// <p>Image callback URL</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`


	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ChannelAttrInfo struct {
	// Device ID to which the device channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name of the device channel
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

type ChannelInfo struct {
	// Device ID to which the channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device channel ID. A device channel can only be added to one cloud connection plan.
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

// Predefined struct for user
type ControlDevicePTZRequestParams struct {
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Command type (up, down, left, right).
	// Upper left: leftup, upper right: rightup, lower left: leftdown, lower right: rightdown
	// Zoom in: zoom in, zoom out: zoom out
	// Focus far: focusfar, focus near: focusnear
	// Iris in: enlarge aperture, iris out: shrink aperture)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Command description (speed range: 1-8)
	Speed *int64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

type ControlDevicePTZRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Command type (up, down, left, right).
	// Upper left: leftup, upper right: rightup, lower left: leftdown, lower right: rightdown
	// Zoom in: zoom in, zoom out: zoom out
	// Focus far: focusfar, focus near: focusnear
	// Iris in: enlarge aperture, iris out: shrink aperture)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Command description (speed range: 1-8)
	Speed *int64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

func (r *ControlDevicePTZRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Type")
	delete(f, "Speed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDevicePTZRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePTZResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDevicePTZResponse struct {
	*tchttp.BaseResponse
	Response *ControlDevicePTZResponseParams `json:"Response"`
}

func (r *ControlDevicePTZResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePTZResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePresetRequestParams struct {
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Command (goto: preset position call;
	// Set: preset position settings
	// del: Delete preset position)
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Preset position index (supports only index positions 1-10; an error is reported if exceeded)
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type ControlDevicePresetRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Command (goto: preset position call;
	// Set: preset position settings
	// del: Delete preset position)
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Preset position index (supports only index positions 1-10; an error is reported if exceeded)
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

func (r *ControlDevicePresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Cmd")
	delete(f, "Index")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDevicePresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDevicePresetResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDevicePresetResponse struct {
	*tchttp.BaseResponse
	Response *ControlDevicePresetResponseParams `json:"Response"`
}

func (r *ControlDevicePresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDevicePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceSnapshotRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Number of burst shots. Range of optional values: 1–10.
	SnapNum *int64 `json:"SnapNum,omitnil,omitempty" name:"SnapNum"`

	// Snapshot interval. Range of optional values: 1–1800
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Image storage period. Default value: 7 days. Only supports 7, 15, 30, 60, 90, 180, or 365 days.
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

type ControlDeviceSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Number of burst shots. Range of optional values: 1–10.
	SnapNum *int64 `json:"SnapNum,omitnil,omitempty" name:"SnapNum"`

	// Snapshot interval. Range of optional values: 1–1800
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Image storage period. Default value: 7 days. Only supports 7, 15, 30, 60, 90, 180, or 365 days.
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`
}

func (r *ControlDeviceSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "SnapNum")
	delete(f, "Interval")
	delete(f, "Expire")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceSnapshotResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDeviceSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceSnapshotResponseParams `json:"Response"`
}

func (r *ControlDeviceSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceStreamData struct {
	// flv stream address.
	Flv *string `json:"Flv,omitnil,omitempty" name:"Flv"`

	// hls stream address.
	Hls *string `json:"Hls,omitnil,omitempty" name:"Hls"`

	// rtmp stream address.
	Rtmp *string `json:"Rtmp,omitnil,omitempty" name:"Rtmp"`
}

// Predefined struct for user
type ControlDeviceStreamRequestParams struct {
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Stream type (1: main stream;
	// 2: Substream (not allowed to be delivered at the same time as Resolution)
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Resolution (1: QCIF;
	// 2:CIF;
	// 3:4CIF;
	// 4:D1;
	// 5:720P;
	// 6:1080P/I;
	// Custom 1920*1080, etc. (device support required) (not allowed to be sent at the same time as StreamType)
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Whether it is a private network
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

type ControlDeviceStreamRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Stream type (1: main stream;
	// 2: Substream (not allowed to be delivered at the same time as Resolution)
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Resolution (1: QCIF;
	// 2:CIF;
	// 3:4CIF;
	// 4:D1;
	// 5:720P;
	// 6:1080P/I;
	// Custom 1920*1080, etc. (device support required) (not allowed to be sent at the same time as StreamType)
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Whether it is a private network
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

func (r *ControlDeviceStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StreamType")
	delete(f, "Resolution")
	delete(f, "IsInternal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlDeviceStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlDeviceStreamResponseParams struct {
	// Returned data.
	Data *ControlDeviceStreamData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlDeviceStreamResponse struct {
	*tchttp.BaseResponse
	Response *ControlDeviceStreamResponseParams `json:"Response"`
}

func (r *ControlDeviceStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlDeviceStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordTimelineRequestParams struct {
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// End time.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type ControlRecordTimelineRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// End time.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

func (r *ControlRecordTimelineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordTimelineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Start")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ControlRecordTimelineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ControlRecordTimelineResponseParams struct {
	// Returned data.
	Data []*Timeline `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ControlRecordTimelineResponse struct {
	*tchttp.BaseResponse
	Response *ControlRecordTimelineResponseParams `json:"Response"`
}

func (r *ControlRecordTimelineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ControlRecordTimelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVLPlanRequestParams struct {
	// <p>Plan name. Supports Chinese, English, digits, _ and -. Up to 32 characters.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>ID of the associated visual understanding template</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>Plan description. Up to 256 characters.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>Number of days to store screenshot files. Value range: [7, 1080]. Required in image mode. Do not pass this parameter in video mode.</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Callback notification configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type CreateVLPlanRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan name. Supports Chinese, English, digits, _ and -. Up to 32 characters.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>ID of the associated visual understanding template</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>Plan description. Up to 256 characters.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>Number of days to store screenshot files. Value range: [7, 1080]. Required in image mode. Do not pass this parameter in video mode.</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Callback notification configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *CreateVLPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVLPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "VLTemplateId")
	delete(f, "Describe")
	delete(f, "FileExpires")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVLPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVLPlanResponseParams struct {
	// <p>Visual understanding plan information</p>
	Data *VLPlanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVLPlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateVLPlanResponseParams `json:"Response"`
}

func (r *CreateVLPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVLPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVLTemplateRequestParams struct {
	// <p>AI model</p><p>Enumeration values:</p><ul><li>ivc-vl-1.0: View large model</li></ul>
	AIModel *string `json:"AIModel,omitnil,omitempty" name:"AIModel"`

	// <p>Template name. Supports only Chinese, English, digits, _, and -, with a maximum length of 32 characters.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Prompt content, maximum 2MB</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>Analysis mode</p><p>Enumeration values:</p><ul><li>image: image mode</li><li>video: video mode</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>List of effective periods, up to 100 entries</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Template description. Maximum length: 256 characters.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>Image scenario configuration (required when AnalysisMode=image)</p>
	ImageConfig *ImageSceneInfo `json:"ImageConfig,omitnil,omitempty" name:"ImageConfig"`
}

type CreateVLTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>AI model</p><p>Enumeration values:</p><ul><li>ivc-vl-1.0: View large model</li></ul>
	AIModel *string `json:"AIModel,omitnil,omitempty" name:"AIModel"`

	// <p>Template name. Supports only Chinese, English, digits, _, and -, with a maximum length of 32 characters.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Prompt content, maximum 2MB</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>Analysis mode</p><p>Enumeration values:</p><ul><li>image: image mode</li><li>video: video mode</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>List of effective periods, up to 100 entries</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Template description. Maximum length: 256 characters.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>Image scenario configuration (required when AnalysisMode=image)</p>
	ImageConfig *ImageSceneInfo `json:"ImageConfig,omitnil,omitempty" name:"ImageConfig"`
}

func (r *CreateVLTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVLTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AIModel")
	delete(f, "Name")
	delete(f, "Prompt")
	delete(f, "AnalysisMode")
	delete(f, "Times")
	delete(f, "Describe")
	delete(f, "ImageConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVLTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVLTemplateResponseParams struct {
	// <p>Visual understanding template information</p>
	Data *VLTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVLTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateVLTemplateResponseParams `json:"Response"`
}

func (r *CreateVLTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVLTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoDownloadTaskRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Default: 1x. Supports 1x, 2x, 4x, and 8x playback speeds.
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// Expiration time of the transcoded mp4 file (7, 15, 30, 60, 90, 180, and 365 are supported)
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`

	// Download file format. Currently only supports (1: mp4)
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Completion policy (0: The task fails and no MP4 is generated if the stream pull fails but the video recording is incomplete. 1: The task succeeds and an MP4 is generated if the stream pull fails but the video recording is incomplete.)
	CompletionPolicy *int64 `json:"CompletionPolicy,omitnil,omitempty" name:"CompletionPolicy"`
}

type CreateVideoDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Default: 1x. Supports 1x, 2x, 4x, and 8x playback speeds.
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// Expiration time of the transcoded mp4 file (7, 15, 30, 60, 90, 180, and 365 are supported)
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`

	// Download file format. Currently only supports (1: mp4)
	FileType *int64 `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Completion policy (0: The task fails and no MP4 is generated if the stream pull fails but the video recording is incomplete. 1: The task succeeds and an MP4 is generated if the stream pull fails but the video recording is incomplete.)
	CompletionPolicy *int64 `json:"CompletionPolicy,omitnil,omitempty" name:"CompletionPolicy"`
}

func (r *CreateVideoDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Scale")
	delete(f, "Expire")
	delete(f, "FileType")
	delete(f, "CompletionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoDownloadTaskResponseParams struct {
	// Returned results of the download task
	Data *VideoDownloadTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVideoDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateVideoDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoSnapshotConfigRequestParams struct {
	// <p>Screenshot configuration name</p><p>Input parameter limits: Chinese, English, digits, "-", "_", length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Hot storage days for image files</p><p>Enumeration values:</p><ul><li>7: 7 days</li><li>15: 15 days</li><li>30: 30 days</li><li>60: 60 days</li><li>90: 90 days</li><li>180: 180 days</li><li>270: 270 days</li><li>360: 360 days</li><li>540: 540 days</li><li>720: 720 days</li><li>900: 900 days</li><li>1080: 1080 days</li></ul>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Screenshot result callback URL</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type CreateVideoSnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot configuration name</p><p>Input parameter limits: Chinese, English, digits, "-", "_", length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Hot storage days for image files</p><p>Enumeration values:</p><ul><li>7: 7 days</li><li>15: 15 days</li><li>30: 30 days</li><li>60: 60 days</li><li>90: 90 days</li><li>180: 180 days</li><li>270: 270 days</li><li>360: 360 days</li><li>540: 540 days</li><li>720: 720 days</li><li>900: 900 days</li><li>1080: 1080 days</li></ul>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Screenshot result callback URL</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *CreateVideoSnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TemplateId")
	delete(f, "FileExpires")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoSnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoSnapshotConfigResponseParams struct {
	// <p>Screenshot configuration message</p>
	Data *VideoSnapshotConfigInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVideoSnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoSnapshotConfigResponseParams `json:"Response"`
}

func (r *CreateVideoSnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoSnapshotTemplateRequestParams struct {
	// <p>Template name: Chinese, English, digits, "-", "_", with a length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type: 1-Real-time frame capture 2-Low frequency frame capture<br>Note: Real-time frame capture pulls the real-time stream throughout the effective period. Low frequency frame capture pulls the real-time stream as needed during the effective period, with a minimum duration of not less than 1 minute each time.</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Frame capturing frequency: When Type = 1, the value range is [5, 180] seconds; when Type = 2, the value range is [30, 36000] seconds</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Effective period. Note: The array size is no more than 100, and each time period is not less than 1 minute.</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Multiple mode</p>
	MultiMode *MultiVideoSnapshotMode `json:"MultiMode,omitnil,omitempty" name:"MultiMode"`
}

type CreateVideoSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template name: Chinese, English, digits, "-", "_", with a length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type: 1-Real-time frame capture 2-Low frequency frame capture<br>Note: Real-time frame capture pulls the real-time stream throughout the effective period. Low frequency frame capture pulls the real-time stream as needed during the effective period, with a minimum duration of not less than 1 minute each time.</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Frame capturing frequency: When Type = 1, the value range is [5, 180] seconds; when Type = 2, the value range is [30, 36000] seconds</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Effective period. Note: The array size is no more than 100, and each time period is not less than 1 minute.</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Multiple mode</p>
	MultiMode *MultiVideoSnapshotMode `json:"MultiMode,omitnil,omitempty" name:"MultiMode"`
}

func (r *CreateVideoSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Interval")
	delete(f, "Times")
	delete(f, "MultiMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVideoSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVideoSnapshotTemplateResponseParams struct {
	// <p>Screenshot template information</p>
	Data *VideoSnapshotTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVideoSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateVideoSnapshotTemplateResponseParams `json:"Response"`
}

func (r *CreateVideoSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVideoSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainRequestParams struct {
	// Domain name ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain name ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainResponseParams `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationRequestParams struct {
	// Organization ID (obtain from the DescribeOrganization API for querying organizations)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Organization ID (obtain from the DescribeOrganization API for querying organizations)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupPlanRequestParams struct {
	// Recording cloud migration plan ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DeleteRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// Recording cloud migration plan ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DeleteRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupPlanResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordBackupPlanResponseParams `json:"Response"`
}

func (r *DeleteRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupTemplateRequestParams struct {
	// Template ID (obtain from the ListRecordBackupTemplates API for querying the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID (obtain from the ListRecordBackupTemplates API for querying the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordBackupTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *DeleteRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordPlanRequestParams struct {
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DeleteRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DeleteRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordPlanResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordPlanResponseParams `json:"Response"`
}

func (r *DeleteRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRetrieveTaskRequestParams struct {
	// Retrieval task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// Retrieval task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordRetrieveTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *DeleteRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordTemplateRequestParams struct {
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordTemplateResponseParams `json:"Response"`
}

func (r *DeleteRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskRequestParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskResponseParams `json:"Response"`
}

func (r *DeleteTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserDeviceRequestParams struct {
	// Device ID (obtain from the ListDevices API for getting the device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DeleteUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Device ID (obtain from the ListDevices API for getting the device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DeleteUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserDeviceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserDeviceResponseParams `json:"Response"`
}

func (r *DeleteUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVLPlanRequestParams struct {
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`
}

type DeleteVLPlanRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`
}

func (r *DeleteVLPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVLPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVLPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVLPlanResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVLPlanResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVLPlanResponseParams `json:"Response"`
}

func (r *DeleteVLPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVLPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVLTemplateRequestParams struct {
	// <p>Template ID</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`
}

type DeleteVLTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template ID</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`
}

func (r *DeleteVLTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVLTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVLTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVLTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVLTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVLTemplateResponseParams `json:"Response"`
}

func (r *DeleteVLTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVLTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoSnapshotConfigRequestParams struct {
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DeleteVideoSnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DeleteVideoSnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoSnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVideoSnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoSnapshotConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVideoSnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVideoSnapshotConfigResponseParams `json:"Response"`
}

func (r *DeleteVideoSnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoSnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoSnapshotTemplateRequestParams struct {
	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteVideoSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteVideoSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVideoSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVideoSnapshotTemplateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVideoSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVideoSnapshotTemplateResponseParams `json:"Response"`
}

func (r *DeleteVideoSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVideoSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCNAMERequestParams struct {
	// Service node ID (obtain from the DescribeDomainRegion API for querying domain name bindable service nodes)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Domain type. 0: pull streaming domain name. 1: streaming domain.
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

type DescribeCNAMERequest struct {
	*tchttp.BaseRequest
	
	// Service node ID (obtain from the DescribeDomainRegion API for querying domain name bindable service nodes)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Domain type. 0: pull streaming domain name. 1: streaming domain.
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

func (r *DescribeCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DomainType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCNAMEResponseParams struct {
	// CNAME record value
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCNAMEResponseParams `json:"Response"`
}

func (r *DescribeCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceAddrList struct {
	// List of device addresses
	RemoteAddrs []*RemoteAddrInfo `json:"RemoteAddrs,omitnil,omitempty" name:"RemoteAddrs"`
}

type DescribeDeviceChannelData struct {
	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel encoding
	ChannelCode *string `json:"ChannelCode,omitnil,omitempty" name:"ChannelCode"`

	// Channel name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Stream status (0: not transmitted, 1: in transit)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether Ptz is controllable (0: no, 1: yes)
	PTZType *int64 `json:"PTZType,omitnil,omitempty" name:"PTZType"`

	// Channel manufacturer
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// Channel supports resolution (resolution list separated by '/', National Standard Protocol example (6/3), custom example (12800960/640480))
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Channel offline status (0: offline, 1: online)
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DescribeDeviceChannelRequestParams struct {
	// Device ID (obtain from the ListDevices API for getting the device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DescribeDeviceChannelRequest struct {
	*tchttp.BaseRequest
	
	// Device ID (obtain from the ListDevices API for getting the device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DescribeDeviceChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceChannelResponseParams struct {
	// Returned result
	Data []*DescribeDeviceChannelData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceChannelResponseParams `json:"Response"`
}

func (r *DescribeDeviceChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceData struct {
	// <p>Device ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device code (for national standard devices, it is the 20-digit National Standard Encoding we generated for the device; for rtmp devices, it is the 10-digit device code)</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Device name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device access protocol. 1: RTMP, 2: GB, 3: GW</p>
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// <p>Device type. 1: IPC. 2: NVR.</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Device access service node id</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Device access service node name</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Device streaming protocol. 1: UDP; 2: TCP</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>sip service ID</p>
	SipId *string `json:"SipId,omitnil,omitempty" name:"SipId"`

	// <p>sip Service Domain</p>
	SipDomain *string `json:"SipDomain,omitnil,omitempty" name:"SipDomain"`

	// <p>sip service IP address</p>
	SipIp *string `json:"SipIp,omitnil,omitempty" name:"SipIp"`

	// <p>sip server port</p>
	SipPort *int64 `json:"SipPort,omitnil,omitempty" name:"SipPort"`

	// <p>Rtmp device streaming address (valid only for Rtmp devices)</p>
	PushStreamUrl *string `json:"PushStreamUrl,omitnil,omitempty" name:"PushStreamUrl"`

	// <p>Device status. 0: unregistered, 1: online, 2: offline, 3: disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Organization ID of the device</p>
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Device gateway ID, obtained from the query gateway list API (required for gateway access only)</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>Gateway name of the device ownership</p>
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// <p>Device gateway protocol name</p>
	ProtocolTypeName *string `json:"ProtocolTypeName,omitnil,omitempty" name:"ProtocolTypeName"`

	// <p>Gateway access protocol type. 1. Hikvision SDK, 2. Dahua SDK, 3. Uniview SDK, 4. Onvif (required for gateway access only)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Device connectivity IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device Port</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Device region</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>Device vendor</p>
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// <p>Audio on/off (0: off; 1: on). Audio is dropped when turned off.</p>
	AudioSwitch *int64 `json:"AudioSwitch,omitnil,omitempty" name:"AudioSwitch"`

	// <p>Subscription switch (0: off; 1: on), enabled by default. When enabled, it subscribes to device channel adjustments. Valid only for national standard NVR devices</p>
	SubscribeSwitch *int64 `json:"SubscribeSwitch,omitnil,omitempty" name:"SubscribeSwitch"`

	// <p>Custom appName for RTMP publishing address</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Custom streamName for RTMP publishing address</p>
	StreamName *string `json:"StreamName,omitnil,omitempty" name:"StreamName"`

	// <p>Whether to enable silent frames (0: disabled; 1: enabled)</p>
	SilentFrameSwitch *int64 `json:"SilentFrameSwitch,omitnil,omitempty" name:"SilentFrameSwitch"`

	// <p>Rtmp device secure authentication streaming address (valid only for Rtmp devices)</p>
	PushStreamSecureUrl *string `json:"PushStreamSecureUrl,omitnil,omitempty" name:"PushStreamSecureUrl"`
}

type DescribeDeviceListData struct {
	// Device detail list
	List []*DescribeDeviceData `json:"List,omitnil,omitempty" name:"List"`
}

type DescribeDevicePresetData struct {
	// Preset position index. Supports only indexes 1-10.
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// Preset position name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type DescribeDevicePresetRequestParams struct {
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

type DescribeDevicePresetRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the channel query API DescribeDeviceChannel)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`
}

func (r *DescribeDevicePresetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePresetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicePresetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicePresetResponseParams struct {
	// Returned data.
	Data []*DescribeDevicePresetData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDevicePresetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicePresetResponseParams `json:"Response"`
}

func (r *DescribeDevicePresetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicePresetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainData struct {
	// Domain ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Playback domain
	PlayDomain *string `json:"PlayDomain,omitnil,omitempty" name:"PlayDomain"`

	// CNAME record value
	InternalDomain *string `json:"InternalDomain,omitnil,omitempty" name:"InternalDomain"`

	// Whether to upload a certificate (0: No, 1: Yes)
	HaveCert *int64 `json:"HaveCert,omitnil,omitempty" name:"HaveCert"`

	// Service node ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Service node name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// User ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Certificate ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Domain type. 0: Pull streaming domain name; 1: Streaming domain
	DomainType *int64 `json:"DomainType,omitnil,omitempty" name:"DomainType"`
}

type DescribeDomainRegionData struct {
	// Service node description
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Service node ID (the ClusterId required in other APIs)
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Region information
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DescribeDomainRegionRequestParams struct {

}

type DescribeDomainRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRegionResponseParams struct {
	// Returned data.
	Data []*DescribeDomainRegionData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainRegionResponseParams `json:"Response"`
}

func (r *DescribeDomainRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRequestParams struct {

}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainResponseParams struct {
	// Returned data.
	Data []*DescribeDomainData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainResponseParams `json:"Response"`
}

func (r *DescribeDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGBDeviceAddrRequestParams struct {
	// Device ID list
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type DescribeGBDeviceAddrRequest struct {
	*tchttp.BaseRequest
	
	// Device ID list
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

func (r *DescribeGBDeviceAddrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGBDeviceAddrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGBDeviceAddrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGBDeviceAddrResponseParams struct {
	// None.
	Data *DescribeDeviceAddrList `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGBDeviceAddrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGBDeviceAddrResponseParams `json:"Response"`
}

func (r *DescribeGBDeviceAddrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGBDeviceAddrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationData struct {
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Organization name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parent node ID of the organizational unit
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Organizational hierarchy
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// User ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Organizational structure
	ParentIds *string `json:"ParentIds,omitnil,omitempty" name:"ParentIds"`

	// Total number of devices
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of devices online
	Online *int64 `json:"Online,omitnil,omitempty" name:"Online"`
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {

}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// Returned data.
	Data []*DescribeOrganizationData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordBackupPlanData struct {
	// Recording cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Recording cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Description of the recording cloud migration plan
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Recording cloud migration plan status. 1: In normal use. 0: Deleting, cannot be used.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// Creation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Modification time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

// Predefined struct for user
type DescribeRecordBackupPlanRequestParams struct {
	// Recording cloud migration plan ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DescribeRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// Recording cloud migration plan ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DescribeRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordBackupPlanResponseParams struct {
	// Returned data.
	Data *DescribeRecordBackupPlanData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordBackupPlanResponseParams `json:"Response"`
}

func (r *DescribeRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordBackupTemplateData struct {
	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Template name</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (settings can be done by week, support setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (settings can be done by week. You can set multiple time periods each day, each not less than 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Creation time.</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>Update time.</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// <p>Cycle offset. Default value: 1. It means to pull yesterday's video recording from the device. 0 means to pull today's video recording from the device. 3 means to pull the video recording from the day before yesterday from the device.</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

// Predefined struct for user
type DescribeRecordBackupTemplateRequestParams struct {
	// Template ID (obtain from the ListRecordBackupTemplates API used to query the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID (obtain from the ListRecordBackupTemplates API used to query the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordBackupTemplateResponseParams struct {
	// Returned data.
	Data *DescribeRecordBackupTemplateData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *DescribeRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordFileData struct {
	// Notification type. 0: No archived video recording within the period. 1: Archived video recording within the period.
	Tips *int64 `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Exists in array format. If it does not exist, the field content is empty.
	List []*RecordTimeLine `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type DescribeRecordFileRequestParams struct {
	// Device ID to which the channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Search start time in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search end time, UTC seconds. Example: 1662114246. The longest period between the start and end times is one day, and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to carry the playback url for each time period
	WithUrl *bool `json:"WithUrl,omitnil,omitempty" name:"WithUrl"`
}

type DescribeRecordFileRequest struct {
	*tchttp.BaseRequest
	
	// Device ID to which the channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Search start time in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Search end time, UTC seconds. Example: 1662114246. The longest period between the start and end times is one day, and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to carry the playback url for each time period
	WithUrl *bool `json:"WithUrl,omitnil,omitempty" name:"WithUrl"`
}

func (r *DescribeRecordFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "WithUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordFileResponseParams struct {
	// Returned result
	Data *DescribeRecordFileData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordFileResponseParams `json:"Response"`
}

func (r *DescribeRecordFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlanRequestParams struct {
	// Real-time cloud plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type DescribeRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// Real-time cloud plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *DescribeRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlanResponseParams struct {
	// Output result
	Data *RecordPlanBaseInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordPlanResponseParams `json:"Response"`
}

func (r *DescribeRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlaybackUrlRequestParams struct {
	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Replay start time in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Replay end time in UTC seconds, for example: 1662114246. The longest period between the start and end times is one day, and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to obtain the intranet address
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// Whether to enable timestamp correction during cloud video playback. This primarily addresses issues such as timestamp reversal or rollback that cause playback failure.
	CorrectTimestamp *bool `json:"CorrectTimestamp,omitnil,omitempty" name:"CorrectTimestamp"`
}

type DescribeRecordPlaybackUrlRequest struct {
	*tchttp.BaseRequest
	
	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Replay start time in UTC seconds, for example: 1662114146. The longest period between the start and end times is one day and cannot span multiple days.
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Replay end time in UTC seconds, for example: 1662114246. The longest period between the start and end times is one day, and cannot span multiple days.
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether to obtain the intranet address
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// Whether to enable timestamp correction during cloud video playback. This primarily addresses issues such as timestamp reversal or rollback that cause playback failure.
	CorrectTimestamp *bool `json:"CorrectTimestamp,omitnil,omitempty" name:"CorrectTimestamp"`
}

func (r *DescribeRecordPlaybackUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlaybackUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "IsInternal")
	delete(f, "CorrectTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordPlaybackUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordPlaybackUrlResponseParams struct {
	// Returned results.
	Data *RecordPlaybackUrl `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordPlaybackUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordPlaybackUrlResponseParams `json:"Response"`
}

func (r *DescribeRecordPlaybackUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordPlaybackUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordRetrieveTaskData struct {
	// Retrieval task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name for retrieval
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time of video recording retrieval
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Retrieval end time of the video recording
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Retrieval mode. 1: speed mode. Others not supported currently.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Replica validity period
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// Task status. 0: Retrieved, 1: Retrieving, 2: To be retrieved
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Retrieval capacity, in MB.
	Capacity *float64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// Task device channel id
	Channels []*RecordRetrieveTaskChannelInfo `json:"Channels,omitnil,omitempty" name:"Channels"`

	// Task description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Number of task channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`
}

// Predefined struct for user
type DescribeRecordRetrieveTaskRequestParams struct {
	// Cloud recording retrieval task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeRecordRetrieveTaskRequest struct {
	*tchttp.BaseRequest
	
	// Cloud recording retrieval task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeRecordRetrieveTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRetrieveTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRetrieveTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordRetrieveTaskResponseParams struct {
	// Returned result
	Data *DescribeRecordRetrieveTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordRetrieveTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordRetrieveTaskResponseParams `json:"Response"`
}

func (r *DescribeRecordRetrieveTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRetrieveTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTemplateRequestParams struct {
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordTemplateResponseParams struct {
	// Returned results
	Data *RecordTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordTemplateResponseParams `json:"Response"`
}

func (r *DescribeRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStreamAuthData struct {
	// Authentication configuration ID (uuid).
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Whether to enable playback authentication. Valid values: `1`: enable, `0`: disable.
	PullState *int64 `json:"PullState,omitnil,omitempty" name:"PullState"`

	// Playback key (alphanumeric only, length: 0-10 characters)
	PullSecret *string `json:"PullSecret,omitnil,omitempty" name:"PullSecret"`

	// Playback expiration time (unit: minute)
	PullExpired *int64 `json:"PullExpired,omitnil,omitempty" name:"PullExpired"`

	// Whether to enable stream authentication (1: enable, 0: disable)
	PushState *int64 `json:"PushState,omitnil,omitempty" name:"PushState"`

	// Stream key (alphanumeric only, length 0-10 characters)
	PushSecret *string `json:"PushSecret,omitnil,omitempty" name:"PushSecret"`

	// Push expiration time (unit: minute)
	PushExpired *int64 `json:"PushExpired,omitnil,omitempty" name:"PushExpired"`

	// User ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

// Predefined struct for user
type DescribeStreamAuthRequestParams struct {

}

type DescribeStreamAuthRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStreamAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamAuthResponseParams struct {
	// Returned result
	Data *DescribeStreamAuthData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStreamAuthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamAuthResponseParams `json:"Response"`
}

func (r *DescribeStreamAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRequestParams struct {
	// ID of a simple or complex task
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskRequest struct {
	*tchttp.BaseRequest
	
	// ID of a simple or complex task
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResponseParams struct {
	// Task details.
	Data *TaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResponseParams `json:"Response"`
}

func (r *DescribeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceListRequestParams struct {
	// Device Id list
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

type DescribeUserDeviceListRequest struct {
	*tchttp.BaseRequest
	
	// Device Id list
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`
}

func (r *DescribeUserDeviceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDeviceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceListResponseParams struct {
	// Returned result
	Data *DescribeDeviceListData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDeviceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDeviceListResponseParams `json:"Response"`
}

func (r *DescribeUserDeviceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceRequestParams struct {
	// Device ID (obtain from the ListDevices api)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type DescribeUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// Device ID (obtain from the ListDevices api)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *DescribeUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDeviceResponseParams struct {
	// Returned result
	Data *DescribeDeviceData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDeviceResponseParams `json:"Response"`
}

func (r *DescribeUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLPlanRequestParams struct {
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`
}

type DescribeVLPlanRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`
}

func (r *DescribeVLPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVLPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLPlanResponseParams struct {
	// <p>Visual understanding plan information</p>
	Data *VLPlanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVLPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVLPlanResponseParams `json:"Response"`
}

func (r *DescribeVLPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLPlanResultRequestParams struct {
	// <p>Analysis result ID</p>
	VLPlanResultId *string `json:"VLPlanResultId,omitnil,omitempty" name:"VLPlanResultId"`
}

type DescribeVLPlanResultRequest struct {
	*tchttp.BaseRequest
	
	// <p>Analysis result ID</p>
	VLPlanResultId *string `json:"VLPlanResultId,omitnil,omitempty" name:"VLPlanResultId"`
}

func (r *DescribeVLPlanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLPlanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanResultId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVLPlanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLPlanResultResponseParams struct {
	// <p>Analysis result details</p>
	Data *VLPlanResultDetailInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVLPlanResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVLPlanResultResponseParams `json:"Response"`
}

func (r *DescribeVLPlanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLPlanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLTemplateRequestParams struct {
	// <p>Template ID</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`
}

type DescribeVLTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template ID</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`
}

func (r *DescribeVLTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVLTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVLTemplateResponseParams struct {
	// <p>Visual understanding template information</p>
	Data *VLTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVLTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVLTemplateResponseParams `json:"Response"`
}

func (r *DescribeVLTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVLTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoBitRateList struct {
	// List of channel bitrates
	BitRates []*BitRateInfo `json:"BitRates,omitnil,omitempty" name:"BitRates"`
}

// Predefined struct for user
type DescribeVideoBitRateRequestParams struct {
	// Channel ID list
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`
}

type DescribeVideoBitRateRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID list
	ChannelIds []*string `json:"ChannelIds,omitnil,omitempty" name:"ChannelIds"`
}

func (r *DescribeVideoBitRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoBitRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoBitRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoBitRateResponseParams struct {
	// None.
	Data *DescribeVideoBitRateList `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoBitRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoBitRateResponseParams `json:"Response"`
}

func (r *DescribeVideoBitRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoBitRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVideoDownloadUrlData struct {
	// Video recording file download URL
	// Note:
	// The URL is valid for 10 minutes. Access will be denied after expiry. Obtain a new one if needed. 
	// Video file download uses chunked transfer encoding, and the response header is Transfer-Encoding: chunked. 
	// The naming format of the downloaded file is {ChannelId}-{BeginTime}-{EndTime}.{FileType}
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Actual download start time of the video recording
	// Note: This field is returned only when the IsRespActualTime parameter is set to true in the request.
	ActualBeginTime *string `json:"ActualBeginTime,omitnil,omitempty" name:"ActualBeginTime"`

	// Actual download end time of the video recording
	// Note: This field is returned only when the IsRespActualTime parameter is set to true in the request.
	ActualEndTime *string `json:"ActualEndTime,omitnil,omitempty" name:"ActualEndTime"`
}

// Predefined struct for user
type DescribeVideoDownloadUrlRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Download start time in UTC seconds. The longest period between the start and end times is 60 minutes and cannot span across days.
	// Note: The actual downloaded file duration may be longer than this period. You can obtain the actual download start time and end time by specifying the IsRespActualTime parameter. This is because downloads are aligned by TS segments to ensure data integrity and include the entire specified period.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Download end time in UTC seconds. The longest period between the start and end times is 60 minutes and cannot span across days.
	// Note: The actual downloaded file duration may be longer than this period. You can obtain the actual download start time and end time by specifying the IsRespActualTime parameter. This is because downloads are aligned by TS segments to ensure data integrity and include the entire specified period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File format. "mp4": MP4 format. "ts": TS file format.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// data carries the actual download video recording start time and end time required or not
	IsRespActualTime *bool `json:"IsRespActualTime,omitnil,omitempty" name:"IsRespActualTime"`

	// Whether to return the private network download URL. Default false: the public network download URL is returned. If true, the private network download URL is returned.
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// Set the valid period of the URL. The minimum value is 1 second and the maximum value is 86400 seconds. If not set, the default is 600 seconds.
	Expires *int64 `json:"Expires,omitnil,omitempty" name:"Expires"`

	// Whether the downloaded MP4 file supports G711 audio encoding. 
	// Note: If the audio encoding in cloud recording is AAC, then the downloaded MP4 supports AAC encoding by default.
	// If the audio encoding in the cloud recording is G711 and IsSupportG711 is set to true, the downloaded MP4 supports G711 audio encoding.
	// If the audio encoding in the cloud recording is G711 and IsSupportG711 is set to false, the downloaded MP4 does not support G711 audio encoding.
	// This parameter is valid only when FileType is mp4. If left unset, it defaults to false.
	IsSupportG711 *bool `json:"IsSupportG711,omitnil,omitempty" name:"IsSupportG711"`
}

type DescribeVideoDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Download start time in UTC seconds. The longest period between the start and end times is 60 minutes and cannot span across days.
	// Note: The actual downloaded file duration may be longer than this period. You can obtain the actual download start time and end time by specifying the IsRespActualTime parameter. This is because downloads are aligned by TS segments to ensure data integrity and include the entire specified period.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Download end time in UTC seconds. The longest period between the start and end times is 60 minutes and cannot span across days.
	// Note: The actual downloaded file duration may be longer than this period. You can obtain the actual download start time and end time by specifying the IsRespActualTime parameter. This is because downloads are aligned by TS segments to ensure data integrity and include the entire specified period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File format. "mp4": MP4 format. "ts": TS file format.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// data carries the actual download video recording start time and end time required or not
	IsRespActualTime *bool `json:"IsRespActualTime,omitnil,omitempty" name:"IsRespActualTime"`

	// Whether to return the private network download URL. Default false: the public network download URL is returned. If true, the private network download URL is returned.
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`

	// Set the valid period of the URL. The minimum value is 1 second and the maximum value is 86400 seconds. If not set, the default is 600 seconds.
	Expires *int64 `json:"Expires,omitnil,omitempty" name:"Expires"`

	// Whether the downloaded MP4 file supports G711 audio encoding. 
	// Note: If the audio encoding in cloud recording is AAC, then the downloaded MP4 supports AAC encoding by default.
	// If the audio encoding in the cloud recording is G711 and IsSupportG711 is set to true, the downloaded MP4 supports G711 audio encoding.
	// If the audio encoding in the cloud recording is G711 and IsSupportG711 is set to false, the downloaded MP4 does not support G711 audio encoding.
	// This parameter is valid only when FileType is mp4. If left unset, it defaults to false.
	IsSupportG711 *bool `json:"IsSupportG711,omitnil,omitempty" name:"IsSupportG711"`
}

func (r *DescribeVideoDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "FileType")
	delete(f, "IsRespActualTime")
	delete(f, "IsInternal")
	delete(f, "Expires")
	delete(f, "IsSupportG711")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoDownloadUrlResponseParams struct {
	// Returned data structure
	Data *DescribeVideoDownloadUrlData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeVideoDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSnapshotConfigRequestParams struct {
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

type DescribeVideoSnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`
}

func (r *DescribeVideoSnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoSnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSnapshotConfigResponseParams struct {
	// <p>Screenshot configuration details</p>
	Data *VideoSnapshotConfigInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoSnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoSnapshotConfigResponseParams `json:"Response"`
}

func (r *DescribeVideoSnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSnapshotTemplateRequestParams struct {
	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeVideoSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeVideoSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVideoSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVideoSnapshotTemplateResponseParams struct {
	// <p>Screenshot template description</p>
	Data *VideoSnapshotTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVideoSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVideoSnapshotTemplateResponseParams `json:"Response"`
}

func (r *DescribeVideoSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVideoSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EffectiveTime struct {
	// <p>Start time</p><p>Parameter format: hh:mm:ss</p>
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// <p>End time</p><p>Parameter format: hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Repetition cycle: 1-Monday 2-Tuesday 3-Wednesday 4-Thursday 5-Friday 6-Saturday 7-Sunday</p><p>Value range: [1, 7]</p>
	WeekDay *int64 `json:"WeekDay,omitnil,omitempty" name:"WeekDay"`
}

type GBDeviceSnapInfo struct {
	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Download link. Empty value means the stored image expires.
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// Image size in Unit B.
	ImageSize *int64 `json:"ImageSize,omitnil,omitempty" name:"ImageSize"`

	// Creation time of the file.
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Reception time of the image
	ReceivedTime *string `json:"ReceivedTime,omitnil,omitempty" name:"ReceivedTime"`

	// Preview link. Empty value means the stored image expires.
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`

	// National standard signaling session ID, corresponding to the request_id returned by the ControlDeviceSnapshot API.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type ISAPIOutputData struct {
	// Output parameter
	OutputData *string `json:"OutputData,omitnil,omitempty" name:"OutputData"`
}

type ImageSceneInfo struct {
	// <p>Screenshot type</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Frame capturing frequency</p><p>When Type=1, the value range is [5, 180] seconds; when Type=2, the value range is [30, 36000] seconds</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Multiple screenshots mode</p>
	MultiMode *MultiSnapshotMode `json:"MultiMode,omitnil,omitempty" name:"MultiMode"`
}

type ImageUrlInfo struct {
	// <p>Image Id</p>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// <p>Image Url</p>
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>Error information</p><p>Reason for failure to get the image url</p>
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type LifeCycleData struct {
	// Hot storage duration of cloud files in days. Minimum value: 1. Maximum value: 3650.
	Transition *int64 `json:"Transition,omitnil,omitempty" name:"Transition"`

	// Cold storage duration of cloud files in days. 0 means leave unset. The minimum is 60 days when set. The sum of the Expiration and Transition fields is no more than 3650 days.
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`
}

type ListDeviceInfo struct {
	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// Device status. 0: unregistered; 1: online; 2: offline; 3: disabled
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Device streaming protocol. 1: UDP, 2: TCP
	TransportProtocol *uint64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// Device name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Device type. 1: IPC. 2: NVR
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Device password
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Node ID of the device access service
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Service node name
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Access protocol. 1: RTMP, 2: GB, 3: GW
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// Organization ID of the device
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Number of channels
	ChannelNum *uint64 `json:"ChannelNum,omitnil,omitempty" name:"ChannelNum"`
}

// Predefined struct for user
type ListDeviceSnapshotsRequestParams struct {
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device ID (this field is not effective yet)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Query start time. Default query: on the day.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// Query end time. Default query: on the day.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size. Default: 200. Maximum: 2000
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDeviceSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device ID (this field is not effective yet)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Query start time. Default query: on the day.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	// Query end time. Default query: on the day.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size. Default: 200. Maximum: 2000
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDeviceSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeviceSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "DeviceId")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDeviceSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeviceSnapshotsResponseParams struct {
	// List of snapshot information
	Data []*GBDeviceSnapInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Total snapshot results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDeviceSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *ListDeviceSnapshotsResponseParams `json:"Response"`
}

func (r *ListDeviceSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeviceSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDevicesRequestParams struct {
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Whether to obtain the device list of the current level and sub-levels. Default value: false.
	IsContainSubLevel *bool `json:"IsContainSubLevel,omitnil,omitempty" name:"IsContainSubLevel"`

	// Whether the devices already associated with the current user are included. Default value: false.
	IsContainUser *bool `json:"IsContainUser,omitnil,omitempty" name:"IsContainUser"`

	// Device access protocol. 1: RTMP, 2: GB, 3: GW, 4: IVCP (private protocol)
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// Device type. 1: IPC, 2: NVR
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Device status. 0: unregistered; 1: online; 2: offline; 3: disabled.	
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Service node ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Keyword for fuzzy search of devices
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Uin of the current user
	CurrentUin *uint64 `json:"CurrentUin,omitnil,omitempty" name:"CurrentUin"`

	// Page number. Default value: 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Default value: 20. Maximum: 10000 per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListDevicesRequest struct {
	*tchttp.BaseRequest
	
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Whether to obtain the device list of the current level and sub-levels. Default value: false.
	IsContainSubLevel *bool `json:"IsContainSubLevel,omitnil,omitempty" name:"IsContainSubLevel"`

	// Whether the devices already associated with the current user are included. Default value: false.
	IsContainUser *bool `json:"IsContainUser,omitnil,omitempty" name:"IsContainUser"`

	// Device access protocol. 1: RTMP, 2: GB, 3: GW, 4: IVCP (private protocol)
	AccessProtocol *uint64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// Device type. 1: IPC, 2: NVR
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Device status. 0: unregistered; 1: online; 2: offline; 3: disabled.	
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Service node ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Keyword for fuzzy search of devices
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Uin of the current user
	CurrentUin *uint64 `json:"CurrentUin,omitnil,omitempty" name:"CurrentUin"`

	// Page number. Default value: 1.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Default value: 20. Maximum: 10000 per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "IsContainSubLevel")
	delete(f, "IsContainUser")
	delete(f, "AccessProtocol")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "ClusterId")
	delete(f, "Keyword")
	delete(f, "CurrentUin")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDevicesResponseParams struct {
	// Device list details
	Data []*ListDeviceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of devices
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListDevicesResponseParams `json:"Response"`
}

func (r *ListDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListForbidplayChannelsData struct {
	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of devices on the current page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of device channels in this query
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Device channel information list
	List []*ChannelAttrInfo `json:"List,omitnil,omitempty" name:"List"`
}

type ListOrganizationChannelsData struct {
	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of devices on the current page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of device channels in this query
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of device channel information
	List []*OrganizationChannelInfo `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type ListOrganizationChannelsRequestParams struct {
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Maximum number of items per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page. 
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query condition. Query by device name.
	// Only one query condition can take effect at a time. Length not exceeding 32 bytes.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query condition. Proceed to query by channel name.
	// Only one query condition can take effect at a time. Length not exceeding 32 bytes.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

type ListOrganizationChannelsRequest struct {
	*tchttp.BaseRequest
	
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Maximum number of items per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page. 
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query condition. Query by device name.
	// Only one query condition can take effect at a time. Length not exceeding 32 bytes.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query condition. Proceed to query by channel name.
	// Only one query condition can take effect at a time. Length not exceeding 32 bytes.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

func (r *ListOrganizationChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationChannelsResponseParams struct {
	// Returned result
	Data *ListOrganizationChannelsData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationChannelsResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationChannelsResponseParams `json:"Response"`
}

func (r *ListOrganizationChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordBackupPlanData struct {
	// Recording cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Recording cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Description of the recording cloud migration plan
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Recording cloud migration plan status. `1`: in normal use. `0`: deleting, unavailable.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// Creation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Modification time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type ListRecordBackupPlanDevicesData struct {
	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of devices on the current page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of device channels in this query
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of device channel information
	List []*RecordPlanChannelInfo `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type ListRecordBackupPlanDevicesRequestParams struct {
	// Recording schedule ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Query by device name (if empty, this parameter is not used for reference)
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query by channel name (if empty, ignore this parameter)
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Query by organization name (if empty, this parameter is not used for reference)
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Maximum number of items per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ListRecordBackupPlanDevicesRequest struct {
	*tchttp.BaseRequest
	
	// Recording schedule ID (obtain from the ListRecordBackupPlans API for querying the recording cloud migration plan list)
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Query by device name (if empty, this parameter is not used for reference)
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query by channel name (if empty, ignore this parameter)
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Query by organization name (if empty, this parameter is not used for reference)
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Maximum number of items per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *ListRecordBackupPlanDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlanDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	delete(f, "OrganizationName")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupPlanDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlanDevicesResponseParams struct {
	// Returned data.
	Data *ListRecordBackupPlanDevicesData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordBackupPlanDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupPlanDevicesResponseParams `json:"Response"`
}

func (r *ListRecordBackupPlanDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlanDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlansRequestParams struct {

}

type ListRecordBackupPlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupPlansResponseParams struct {
	// Returned data.
	Data []*ListRecordBackupPlanData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupPlansResponseParams `json:"Response"`
}

func (r *ListRecordBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordBackupTemplatesData struct {
	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Template name</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (settings can be done by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (settings can be done by week, supports setting multiple time periods per day, each time period greater than or equal to 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Creation time.</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>Update time.</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// <p>Cycle offset. Default: 1, which means pulling yesterday's data. 0 means pulling today's data. 3 means pulling the day before yesterday's data.</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

// Predefined struct for user
type ListRecordBackupTemplatesRequestParams struct {

}

type ListRecordBackupTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordBackupTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordBackupTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordBackupTemplatesResponseParams struct {
	// Returned data.
	Data []*ListRecordBackupTemplatesData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordBackupTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordBackupTemplatesResponseParams `json:"Response"`
}

func (r *ListRecordBackupTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordBackupTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListRecordPlanDevicesData struct {
	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of devices on the current page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total number of device channels in this query
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of device channel information
	List []*RecordPlanChannelInfo `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type ListRecordPlanDevicesRequestParams struct {
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Maximum number of items per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query by device name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query by channel name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Query by organization name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`
}

type ListRecordPlanDevicesRequest struct {
	*tchttp.BaseRequest
	
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Maximum number of items per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Query by device name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Query by channel name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Query by organization name. If it is empty, ignore this parameter.
	// Only one of the channel name, device name, and organization name is valid at the same time. If multiple fields have values at the same time, query by the priority order of channel name, device name, and organization name. If all are empty, perform a full query. The length cannot exceed 32 bytes.
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`
}

func (r *ListRecordPlanDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	delete(f, "OrganizationName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordPlanDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlanDevicesResponseParams struct {
	// Output result
	Data *ListRecordPlanDevicesData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordPlanDevicesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordPlanDevicesResponseParams `json:"Response"`
}

func (r *ListRecordPlanDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlanDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlansRequestParams struct {

}

type ListRecordPlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordPlansResponseParams struct {
	// Returned results. If a plan exists, the results are in a Json array format. If no plan exists, the field data is empty.
	Data []*RecordPlanBaseInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordPlansResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordPlansResponseParams `json:"Response"`
}

func (r *ListRecordPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordRetrieveTasksRequestParams struct {

}

type ListRecordRetrieveTasksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordRetrieveTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordRetrieveTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordRetrieveTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordRetrieveTasksResponseParams struct {
	// Returned results
	Data []*RecordRetrieveTaskDetailsInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordRetrieveTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordRetrieveTasksResponseParams `json:"Response"`
}

func (r *ListRecordRetrieveTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordRetrieveTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordTemplatesRequestParams struct {

}

type ListRecordTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListRecordTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRecordTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRecordTemplatesResponseParams struct {
	// Output result. If a template exists, it is in Json array format. If no template exists, the field data is empty.
	Data []*RecordTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRecordTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListRecordTemplatesResponseParams `json:"Response"`
}

func (r *ListRecordTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRecordTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSubTasksData struct {
	// Subtask List
	List []*SubTaskData `json:"List,omitnil,omitempty" name:"List"`

	// Subtask quantity
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListSubTasksRequestParams struct {
	// Complex task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Default: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Default: no filtering on this field. Otherwise, filter by task status. Status code: 1-NEW, 2-RUNNING, 3-COMPLETED, 4-FAILED.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListSubTasksRequest struct {
	*tchttp.BaseRequest
	
	// Complex task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Default: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Default: no filtering on this field. Otherwise, filter by task status. Status code: 1-NEW, 2-RUNNING, 3-COMPLETED, 4-FAILED.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListSubTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSubTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSubTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSubTasksResponseParams struct {
	// Returned data.
	Data *ListSubTasksData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSubTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListSubTasksResponseParams `json:"Response"`
}

func (r *ListSubTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSubTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTasksData struct {
	// Task list
	List []*TaskData `json:"List,omitnil,omitempty" name:"List"`

	// Number of Tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListTasksRequestParams struct {
	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. The default value is 20.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Default: do not filter by this field. Otherwise, filter by device operation type. Current values: BatchDeleteUserDevice, BatchDisableDevice, BatchEnableDevice.
	// BatchUpgradeDevice,
	// BatchResetDevice,
	// BatchRebootDevice,
	// BatchRefreshDeviceChannel
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// By default, filtering is not based on this field. Otherwise, filter by task status. Status code: 1 - Pending, 2 - Executing, 3 - Completed, 4 - Cancelled.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Start time.
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListTasksRequest struct {
	*tchttp.BaseRequest
	
	// Page number, which is 1 by default.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. The default value is 20.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Default: do not filter by this field. Otherwise, filter by device operation type. Current values: BatchDeleteUserDevice, BatchDisableDevice, BatchEnableDevice.
	// BatchUpgradeDevice,
	// BatchResetDevice,
	// BatchRebootDevice,
	// BatchRefreshDeviceChannel
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// By default, filtering is not based on this field. Otherwise, filter by task status. Status code: 1 - Pending, 2 - Executing, 3 - Completed, 4 - Cancelled.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Start time.
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Operation")
	delete(f, "Status")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksResponseParams struct {
	// Returned data.
	Data *ListTasksData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListTasksResponseParams `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVLPlansRequestParams struct {
	// <p>Plan name, exact match</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Page number. Default value: 1.</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 20. Maximum: 200.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListVLPlansRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan name, exact match</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Page number. Default value: 1.</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 20. Maximum: 200.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListVLPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVLPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListVLPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVLPlansResponseParams struct {
	// <p>Plan list</p>
	Data []*VLPlanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Total quantity.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListVLPlansResponse struct {
	*tchttp.BaseResponse
	Response *ListVLPlansResponseParams `json:"Response"`
}

func (r *ListVLPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVLPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVLTemplatesRequestParams struct {
	// <p>Template name, exact match</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Analysis mode filter, image / video</p>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>Screenshot type filter (this parameter is valid only when AnalysisMode=image): realtime/low</p>
	CaptureType *string `json:"CaptureType,omitnil,omitempty" name:"CaptureType"`

	// <p>Page number, which is 1 by default.</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 20. Maximum value: 200.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListVLTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template name, exact match</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Analysis mode filter, image / video</p>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>Screenshot type filter (this parameter is valid only when AnalysisMode=image): realtime/low</p>
	CaptureType *string `json:"CaptureType,omitnil,omitempty" name:"CaptureType"`

	// <p>Page number, which is 1 by default.</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 20. Maximum value: 200.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListVLTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVLTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AnalysisMode")
	delete(f, "CaptureType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListVLTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVLTemplatesResponseParams struct {
	// <p>Template list</p>
	Data []*VLTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>Total quantity.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListVLTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListVLTemplatesResponseParams `json:"Response"`
}

func (r *ListVLTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVLTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVideoDownloadTaskData struct {
	// Task list
	List []*VideoDownloadTask `json:"List,omitnil,omitempty" name:"List"`

	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ListVideoDownloadTaskRequestParams struct {
	// Device name. Fuzzy search supported.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Channel name, used for fuzzy search
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Task status (0: preparing, 1: executing, 2: completed, 3: failure)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting rule (only StartTime is supported. A hyphen (-) indicates descending order, e.g., -StartTime)
	SortRule *string `json:"SortRule,omitnil,omitempty" name:"SortRule"`

	// Whether the response carries the preview address (0: without; 1: carry)
	WithPreviewUrl *int64 `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`

	// Page number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Download task ID
	DownloadTaskId *string `json:"DownloadTaskId,omitnil,omitempty" name:"DownloadTaskId"`

	// URL expiration time (in seconds), with a maximum value of 86400 seconds (1 day).
	UrlExpires *int64 `json:"UrlExpires,omitnil,omitempty" name:"UrlExpires"`

	// Task date. Defaults to the current day.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type ListVideoDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// Device name. Fuzzy search supported.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Channel name, used for fuzzy search
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Task status (0: preparing, 1: executing, 2: completed, 3: failure)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting rule (only StartTime is supported. A hyphen (-) indicates descending order, e.g., -StartTime)
	SortRule *string `json:"SortRule,omitnil,omitempty" name:"SortRule"`

	// Whether the response carries the preview address (0: without; 1: carry)
	WithPreviewUrl *int64 `json:"WithPreviewUrl,omitnil,omitempty" name:"WithPreviewUrl"`

	// Page number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Pagination size.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Download task ID
	DownloadTaskId *string `json:"DownloadTaskId,omitnil,omitempty" name:"DownloadTaskId"`

	// URL expiration time (in seconds), with a maximum value of 86400 seconds (1 day).
	UrlExpires *int64 `json:"UrlExpires,omitnil,omitempty" name:"UrlExpires"`

	// Task date. Defaults to the current day.
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *ListVideoDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVideoDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceName")
	delete(f, "ChannelName")
	delete(f, "Status")
	delete(f, "SortRule")
	delete(f, "WithPreviewUrl")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "DownloadTaskId")
	delete(f, "UrlExpires")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListVideoDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVideoDownloadTaskResponseParams struct {
	// Local recording download task list
	Data *ListVideoDownloadTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListVideoDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *ListVideoDownloadTaskResponseParams `json:"Response"`
}

func (r *ListVideoDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVideoDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListVideoSnapshotTemplateData struct {
	// <p>Total.</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Sorting parameter</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>Sort order</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Page number.</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Screenshot template list</p>
	List []*VideoSnapshotTemplateInfo `json:"List,omitnil,omitempty" name:"List"`
}

// Predefined struct for user
type ListVideoSnapshotTemplateRequestParams struct {
	// <p>Template name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Sorting parameter</p><p>Default value: update_time</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>Sort order</p><p>Enumeration values:</p><ul><li>desc: in descending order</li><li>asc: in ascending order</li></ul><p>Default value: desc</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Page number</p><p>Default value: 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page</p><p>Default value: 20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListVideoSnapshotTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>Template name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Sorting parameter</p><p>Default value: update_time</p>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// <p>Sort order</p><p>Enumeration values:</p><ul><li>desc: in descending order</li><li>asc: in ascending order</li></ul><p>Default value: desc</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>Page number</p><p>Default value: 1</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page</p><p>Default value: 20</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListVideoSnapshotTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVideoSnapshotTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Sort")
	delete(f, "Order")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListVideoSnapshotTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListVideoSnapshotTemplateResponseParams struct {
	// <p>Screenshot template result list</p>
	Data *ListVideoSnapshotTemplateData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListVideoSnapshotTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ListVideoSnapshotTemplateResponseParams `json:"Response"`
}

func (r *ListVideoSnapshotTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListVideoSnapshotTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVLPlanRequestParams struct {
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>New plan name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>New plan description.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>ID of the newly associated visual understanding template</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>New storage days, range [7, 1080]. 0 means no modification</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>New callback configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type ModifyVLPlanRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>New plan name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>New plan description.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>ID of the newly associated visual understanding template</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>New storage days, range [7, 1080]. 0 means no modification</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>New callback configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *ModifyVLPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVLPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VLPlanId")
	delete(f, "Name")
	delete(f, "Describe")
	delete(f, "VLTemplateId")
	delete(f, "FileExpires")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVLPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVLPlanResponseParams struct {
	// <p>Visual understanding plan information</p>
	Data *VLPlanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVLPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVLPlanResponseParams `json:"Response"`
}

func (r *ModifyVLPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVLPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVideoSnapshotConfigRequestParams struct {
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Screenshot configuration name</p><p>Input parameter limits: Chinese, English, digits, "-", "_", length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Hot storage days of image files</p><p>Enumeration values:</p><ul><li>7: 7 days</li><li>15: 15 days</li><li>30: 30 days</li><li>60: 60 days</li><li>90: 90 days</li><li>180: 180 days</li><li>270: 270 days</li><li>360: 360 days</li><li>540: 540 days</li><li>720: 720 days</li><li>900: 900 days</li><li>1080: 1080 days</li></ul>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Callback configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

type ModifyVideoSnapshotConfigRequest struct {
	*tchttp.BaseRequest
	
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Screenshot configuration name</p><p>Input parameter limits: Chinese, English, digits, "-", "_", length not exceeding 128 characters</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Hot storage days of image files</p><p>Enumeration values:</p><ul><li>7: 7 days</li><li>15: 15 days</li><li>30: 30 days</li><li>60: 60 days</li><li>90: 90 days</li><li>180: 180 days</li><li>270: 270 days</li><li>360: 360 days</li><li>540: 540 days</li><li>720: 720 days</li><li>900: 900 days</li><li>1080: 1080 days</li></ul>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Callback configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`
}

func (r *ModifyVideoSnapshotConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoSnapshotConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigId")
	delete(f, "Name")
	delete(f, "TemplateId")
	delete(f, "FileExpires")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVideoSnapshotConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVideoSnapshotConfigResponseParams struct {
	// <p>Screenshot configuration message</p>
	Data *VideoSnapshotConfigInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVideoSnapshotConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVideoSnapshotConfigResponseParams `json:"Response"`
}

func (r *ModifyVideoSnapshotConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVideoSnapshotConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiSnapshotMode struct {
	// <p>Whether to enable multiple image mode</p><p>When set to false, single image mode is used, and the Interval and Total parameters are ignored</p><p>Default value: false</p>
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>Frame interception interval</p><p>Value range: [10, 600]</p><p>Unit: second</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Number of captured frames</p><p>Value range: [1, 10]</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type MultiVideoSnapshotMode struct {
	// <p>Whether to enable multiple mode. When set to false, single image mode is used, and the Interval and Total parameters are ignored.</p><p>Default value: false</p>
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>Frame interception interval</p><p>Value range: [10, 600]</p><p>Unit: seconds</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Number of captured frames</p><p>Value range: [1, 10]</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type OrganizationChannelInfo struct {
	// Device ID of the device channel
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name of the device channel
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device channel name.
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Whether the channel is in a cloud migration plan. If so, it cannot be added to another cloud migration plan | true: in a cloud migration plan, false: not in a cloud migration plan
	InPlan *bool `json:"InPlan,omitnil,omitempty" name:"InPlan"`
}

type PlayRecordData struct {
	// Video playback address
	Flv *string `json:"Flv,omitnil,omitempty" name:"Flv"`
}

// Predefined struct for user
type PlayRecordRequestParams struct {
	// Channel ID (obtain from the DescribeDeviceChannel API)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	//  End time.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// Stream type (1: main stream; 2: sub-stream (not allowed to be delivered at the same time as Resolution))
	StreamType *int64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Resolution (1: QCIF; 2: CIF; 3: 4CIF; 4: D1; 5: 720P; 6: 1080P/I; custom 19201080, etc. (device support required) (not allowed to be delivered at the same time as StreamType))
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Whether it is a private network
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

type PlayRecordRequest struct {
	*tchttp.BaseRequest
	
	// Channel ID (obtain from the DescribeDeviceChannel API)
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Start time.
	Start *int64 `json:"Start,omitnil,omitempty" name:"Start"`

	//  End time.
	End *int64 `json:"End,omitnil,omitempty" name:"End"`

	// Stream type (1: main stream; 2: sub-stream (not allowed to be delivered at the same time as Resolution))
	StreamType *int64 `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Resolution (1: QCIF; 2: CIF; 3: 4CIF; 4: D1; 5: 720P; 6: 1080P/I; custom 19201080, etc. (device support required) (not allowed to be delivered at the same time as StreamType))
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// Whether it is a private network
	IsInternal *bool `json:"IsInternal,omitnil,omitempty" name:"IsInternal"`
}

func (r *PlayRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlayRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelId")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "StreamType")
	delete(f, "Resolution")
	delete(f, "IsInternal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PlayRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PlayRecordResponseParams struct {
	// Returned result
	Data *PlayRecordData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PlayRecordResponse struct {
	*tchttp.BaseResponse
	Response *PlayRecordResponseParams `json:"Response"`
}

func (r *PlayRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PlayRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryForbidPlayChannelListRequestParams struct {
	// Sub-user uin, which can also be the root user uin.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Maximum number of items per page. Maximum value: 200. If the number of items exceeds 200, 200 items are returned.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type QueryForbidPlayChannelListRequest struct {
	*tchttp.BaseRequest
	
	// Sub-user uin, which can also be the root user uin.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Maximum number of items per page. Maximum value: 200. If the number of items exceeds 200, 200 items are returned.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Which page.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *QueryForbidPlayChannelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryForbidPlayChannelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryForbidPlayChannelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryForbidPlayChannelListResponseParams struct {
	// Returned result
	Data *ListForbidplayChannelsData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryForbidPlayChannelListResponse struct {
	*tchttp.BaseResponse
	Response *QueryForbidPlayChannelListResponseParams `json:"Response"`
}

func (r *QueryForbidPlayChannelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryForbidPlayChannelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordPlanBaseInfo struct {
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Cloud migration plan description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Stream type. default: default stream type of the device. main: main stream. sub: substream. Other stream types can be customized based on the device capacity.
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Recording schedule status. 1: In normal use. 0: Deleting, cannot be used.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Total number of channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// Video recording supplement mode (0: not enabled, 1: enabled)
	RepairMode *int64 `json:"RepairMode,omitnil,omitempty" name:"RepairMode"`
}

type RecordPlanChannelInfo struct {
	// Device ID to which the device channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name of the device channel
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Affiliated organization name
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Access protocol type of the device to which the channel belongs
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`
}

type RecordPlanOptData struct {
	// Cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Cloud migration plan description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Stream type. default: The default stream type of the device. main: Main stream. sub: Substream. Other values are customized based on the device capacity.
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Video recording supplement mode (0: disabled, 1: enabled)
	RepairMode *int64 `json:"RepairMode,omitnil,omitempty" name:"RepairMode"`
}

type RecordPlaybackUrl struct {
	// hls replay url
	Hls *string `json:"Hls,omitnil,omitempty" name:"Hls"`
}

type RecordRetrieveTaskChannelInfo struct {
	// Device ID to which the device channel belongs
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name of the device channel
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Device channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Device channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Task status. 0: Retrieved, 1: Retrieving, 2: Pending retrieval, 3: No archived video recording
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type RecordRetrieveTaskDetailsInfo struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time of video recording retrieval
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Retrieval end time of the video recording
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Retrieval mode. 1: ultrafast mode. Others not supported currently.
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Replica validity period
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// Task status. 0: Retrieved, 1: Retrieving, 2: Pending retrieval
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Retrieval capacity, in MB.
	Capacity *int64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// Task description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Number of task channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`
}

type RecordTemplateInfo struct {
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cloud upload time period. Settings can be done by week. Multiple time periods can be set for each day, with each time period greater than or equal to 10 minutes.
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`
}

type RecordTemplateTimeSections struct {
	// Day of the week. Value range: 1–7 (corresponding to Monday–Sunday)
	DayOfWeek *int64 `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// Start time. Format: HH:MM:SS. Range: [00:00:00–23:59:59]
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. Format: HH:MM:SS. Range: [00:00:00–23:59:59]
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type RecordTimeLine struct {
	// Segment start time in UTC seconds, for example: 1662114146
	Begin *uint64 `json:"Begin,omitnil,omitempty" name:"Begin"`

	// Segment end time in UTC seconds, for example: 1662114146
	End *uint64 `json:"End,omitnil,omitempty" name:"End"`

	// Playback url of the corresponding recording clip
	HlsUrl *string `json:"HlsUrl,omitnil,omitempty" name:"HlsUrl"`
}

// Predefined struct for user
type RefreshDeviceChannelRequestParams struct {
	// Device ID (obtain from the ListDevices API for getting device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

type RefreshDeviceChannelRequest struct {
	*tchttp.BaseRequest
	
	// Device ID (obtain from the ListDevices API for getting device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`
}

func (r *RefreshDeviceChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDeviceChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshDeviceChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshDeviceChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefreshDeviceChannelResponse struct {
	*tchttp.BaseResponse
	Response *RefreshDeviceChannelResponseParams `json:"Response"`
}

func (r *RefreshDeviceChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDeviceChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteAddrInfo struct {
	// Device Id
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// IP address
	Addr *string `json:"Addr,omitnil,omitempty" name:"Addr"`
}

// Predefined struct for user
type SetForbidPlayChannelsRequestParams struct {
	// Channel parameters to be disabled. Up to 200 channels can be set at a time.
	Channels []*SetForbidplayChannelParam `json:"Channels,omitnil,omitempty" name:"Channels"`

	// User uin, which can be the uin of a sub-user or a root user
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type SetForbidPlayChannelsRequest struct {
	*tchttp.BaseRequest
	
	// Channel parameters to be disabled. Up to 200 channels can be set at a time.
	Channels []*SetForbidplayChannelParam `json:"Channels,omitnil,omitempty" name:"Channels"`

	// User uin, which can be the uin of a sub-user or a root user
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *SetForbidPlayChannelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetForbidPlayChannelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Channels")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetForbidPlayChannelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetForbidPlayChannelsResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetForbidPlayChannelsResponse struct {
	*tchttp.BaseResponse
	Response *SetForbidPlayChannelsResponseParams `json:"Response"`
}

func (r *SetForbidPlayChannelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetForbidPlayChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetForbidplayChannelParam struct {
	// Channel Id
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Whether to forbid stream pushing over the channel
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type SubTaskData struct {
	// Subtask ID
	SubTaskId *string `json:"SubTaskId,omitnil,omitempty" name:"SubTaskId"`

	// Task status. 1: NEW, 2: RUNNING, 3: COMPLETED, 4: FAILED
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Cause of the task failure.
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// Task progress
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Operation type.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Operation type Chinese description
	ActionZhDesc *string `json:"ActionZhDesc,omitnil,omitempty" name:"ActionZhDesc"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Task start time
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// Task creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Task update time
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Task running time in milliseconds
	Runtime *int64 `json:"Runtime,omitnil,omitempty" name:"Runtime"`

	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

type TaskData struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status. 1: NEW, 2: RUNNING, 3: COMPLETED, 4: FAILED
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Reason for failure
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// Progress (0-1)
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Task operation type. A batch task type starts with "Batch".
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Operation type Chinese description
	ActionZhDesc *string `json:"ActionZhDesc,omitnil,omitempty" name:"ActionZhDesc"`

	// Task type. 1. Simple 2. Complex 3. Subtask
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task resource ID (invalid for complex tasks)
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Total number of tasks (valid only for complex tasks)
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Successful Tasks (valid only for complex tasks)
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// Number of failed tasks (valid only for complex tasks)
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// Number of running tasks (valid only for complex tasks)
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// Task start time.
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// Task creation time
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Update task time.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Task running time in milliseconds.
	Runtime *int64 `json:"Runtime,omitnil,omitempty" name:"Runtime"`

	// Device ID
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Device name.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

type Timeline struct {
	// Fragment start time
	Begin *int64 `json:"Begin,omitnil,omitempty" name:"Begin"`

	// Shard end time
	End *int64 `json:"End,omitnil,omitempty" name:"End"`
}

type UpdateDeviceData struct {
	// <p>Device ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device code (for national standard devices, it is the 20-digit National Standard Encoding we generated for the device; for rtmp devices, it is the 10-digit device code)</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Device name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device access protocol. 1: RTMP; 2: GB.</p><p>Enumeration values:</p><ul><li>1: RTMP</li><li>2: GB</li></ul>
	AccessProtocol *int64 `json:"AccessProtocol,omitnil,omitempty" name:"AccessProtocol"`

	// <p>Device type. 1: IPC. 2: NVR</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Device access service node ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>Device access service node name</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// <p>Device streaming protocol. 1: UDP, 2: TCP</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password.</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device status. 0: unregistered, 1: online, 2: offline, 3: disabled</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Affiliated organization ID of the device</p>
	OrganizationId *int64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Device access gateway ID (deprecated, reserved for compatibility, can be ignored)</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>Gateway access protocol type (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Device access IP</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device Port</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>User Id</p>
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

// Predefined struct for user
type UpdateDeviceOrganizationRequestParams struct {
	// Device ID array (obtain from the ListDevices API for getting device list)
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// Organization ID (obtain from the DescribeOrganization API)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

type UpdateDeviceOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Device ID array (obtain from the ListDevices API for getting device list)
	DeviceIds []*string `json:"DeviceIds,omitnil,omitempty" name:"DeviceIds"`

	// Organization ID (obtain from the DescribeOrganization API)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

func (r *UpdateDeviceOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIds")
	delete(f, "OrganizationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceOrganizationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDeviceOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceOrganizationResponseParams `json:"Response"`
}

func (r *UpdateDeviceOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceStatusRequestParams struct {
	// Device ID (obtain from the ListDevices API for getting device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Disable enable status code (2: enable, 3: disable)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateDeviceStatusRequest struct {
	*tchttp.BaseRequest
	
	// Device ID (obtain from the ListDevices API for getting device list)
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Disable enable status code (2: enable, 3: disable)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateDeviceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDeviceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDeviceStatusResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDeviceStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDeviceStatusResponseParams `json:"Response"`
}

func (r *UpdateDeviceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDeviceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrgData struct {
	// Organization ID
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Organization name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Organizational parent node ID
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Organization hierarchy
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// User ID.
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Organizational structure
	ParentIds *string `json:"ParentIds,omitnil,omitempty" name:"ParentIds"`

	// Total number of devices
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of devices online
	Online *int64 `json:"Online,omitnil,omitempty" name:"Online"`
}

// Predefined struct for user
type UpdateOrganizationRequestParams struct {
	// Organization ID (obtain from the DescribeOrganization API for querying organizations)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Organization name. Chinese, English, digits, spaces, Chinese and English parentheses, underscores (_), and hyphens (-) are supported. The length cannot exceed 64 characters, and the organization name must be unique.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type UpdateOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Organization ID (obtain from the DescribeOrganization API for querying organizations)
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Organization name. Chinese, English, digits, spaces, Chinese and English parentheses, underscores (_), and hyphens (-) are supported. The length cannot exceed 64 characters, and the organization name must be unique.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *UpdateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationResponseParams struct {
	// Returned result
	Data *UpdateOrgData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationResponseParams `json:"Response"`
}

func (r *UpdateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordBackupPlanData struct {
	// Recording cloud migration plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Recording cloud migration plan name
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Video cloud template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Description of the recording cloud migration plan
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Cloud file lifecycle
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Recording cloud migration plan status. `1`: in normal use. `0`: deleting, unavailable.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of channels
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// Creation time.
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// Modification time.
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`
}

type UpdateRecordBackupPlanModify struct {
	// Recording cloud migration plan name (supports only Chinese, English, digits, _, and -, length not exceeding 32 characters. The plan name must be globally unique and cannot be empty or repeated. This field is not required if the name is not modified.)
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Recording template ID (obtain from the ListRecordBackupTemplates API for querying the video recording upload template list. This field is not required if the template ID is not modified)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Recording cloud migration plan description (only supports Chinese, English, numbers, _, -, length not exceeding 128 characters. This field is not required if the description is not modified)
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Lifecycle (lifecycle settings for video recording files, managing the time when files are in cold or hot storage. This field is not required if you do not modify the lifecycle)
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Device channels to be added (Json array; not required when no channels are added; the total number of channels added at once cannot exceed 5000, including the number of channels under the organizational directory)
	Add []*ChannelInfo `json:"Add,omitnil,omitempty" name:"Add"`

	// Device channel to delete (Json array containing the IDs of device channels to delete. Not required if no device channel is deleted)
	Del []*string `json:"Del,omitnil,omitempty" name:"Del"`

	// Add the channels of all devices under the organizational unit (Json array, can be empty, and the total number of channels no more than 5000, including the number of channels in the Add field)
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`
}

// Predefined struct for user
type UpdateRecordBackupPlanRequestParams struct {
	// Plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Modified Content
	Mod *UpdateRecordBackupPlanModify `json:"Mod,omitnil,omitempty" name:"Mod"`
}

type UpdateRecordBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Modified Content
	Mod *UpdateRecordBackupPlanModify `json:"Mod,omitnil,omitempty" name:"Mod"`
}

func (r *UpdateRecordBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordBackupPlanResponseParams struct {
	// Returned data.
	Data *UpdateRecordBackupPlanData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRecordBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordBackupPlanResponseParams `json:"Response"`
}

func (r *UpdateRecordBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordBackupTemplateData struct {
	// <p>Template ID</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Template name</p>
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// <p>Cloud Upload Time Period (set by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// <p>Recording time period (set by week, supports setting multiple time periods per day, each time period not less than 10 minutes)</p>
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// <p>Speed multiplication for cloud migration (supports 1x, 2x, and 4x)</p>
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// <p>Creation time.</p>
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// <p>Update time.</p>
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// <p>Cycle offset. Default value: 1, which means to pull yesterday's video recording from the device. 0 means to pull today's video recording from the device. 3 means to pull the video recording from the device from the day before yesterday.</p>
	DayOffset *int64 `json:"DayOffset,omitnil,omitempty" name:"DayOffset"`
}

type UpdateRecordBackupTemplateModify struct {
	// Template name (not required if you do not rename)
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cloud upload time period (settings can be done by week. Multiple time periods can be set for each day, with each time period not less than 10 minutes)
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`

	// Recording time period (set by week. Multiple time periods can be set each day, each not less than 10 minutes)
	DevTimeSections []*RecordTemplateTimeSections `json:"DevTimeSections,omitnil,omitempty" name:"DevTimeSections"`

	// Cloud speed multiplication (supports 1x, 2x, and 4x).
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`
}

// Predefined struct for user
type UpdateRecordBackupTemplateRequestParams struct {
	// Template ID (obtain from the ListRecordBackupTemplates API for querying the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modify the template data for video recording upload to cloud
	Mod *UpdateRecordBackupTemplateModify `json:"Mod,omitnil,omitempty" name:"Mod"`
}

type UpdateRecordBackupTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID (obtain from the ListRecordBackupTemplates API for querying the video recording cloud template list)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modify the template data for video recording upload to cloud
	Mod *UpdateRecordBackupTemplateModify `json:"Mod,omitnil,omitempty" name:"Mod"`
}

func (r *UpdateRecordBackupTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordBackupTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordBackupTemplateResponseParams struct {
	// Returned data.
	Data *UpdateRecordBackupTemplateData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRecordBackupTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordBackupTemplateResponseParams `json:"Response"`
}

func (r *UpdateRecordBackupTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordBackupTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordPlanData struct {
	// Cloud migration plan name, supporting only Chinese, English, digits, _, -, with a length not exceeding 32 characters. The plan name must be globally unique and cannot be empty or duplicated. This field is not required if the name is not modified.
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// Cloud template ID. This field is not required if the template ID is not modified.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Usage plan description. It can contain only Chinese characters, English letters, digits, underscores (_), and hyphens (-), with a length not exceeding 128 characters. This field is not required if you do not modify the description.
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Stream type. default: The stream type is not specified, and the default push type of the device is used. main: main stream. sub: substream. Other: custom based on the device capability set. The length cannot exceed 32 bytes.
	StreamType *string `json:"StreamType,omitnil,omitempty" name:"StreamType"`

	// Lifecycle, file lifecycle settings, the time when cold and hot storage of files is managed. This field is not required if lifecycle is not modified.
	LifeCycle *LifeCycleData `json:"LifeCycle,omitnil,omitempty" name:"LifeCycle"`

	// Device channels to be added. Json array. This field is not required when no channels are added. The total number of channels added at once cannot exceed 5,000, including the number of channels under the organization directory.
	Add []*ChannelInfo `json:"Add,omitnil,omitempty" name:"Add"`

	// Device channels to delete. Json array. The content is the IDs of the device channels to delete. This field is not required if no device channel is deleted.
	Del []*string `json:"Del,omitnil,omitempty" name:"Del"`

	// Organizational directory ID. Add all device channels under the organizational directory. Json array. Can be empty. And the total number of channels is no more than 5,000 (including the number of channels added in the Add field).
	OrganizationId []*string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// Video recording supplement mode (0: not enabled, 1: enabled)
	RepairMode *int64 `json:"RepairMode,omitnil,omitempty" name:"RepairMode"`
}

// Predefined struct for user
type UpdateRecordPlanRequestParams struct {
	// Plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Modify the content of the plan
	Mod *UpdateRecordPlanData `json:"Mod,omitnil,omitempty" name:"Mod"`
}

type UpdateRecordPlanRequest struct {
	*tchttp.BaseRequest
	
	// Plan ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// Modify the content of the plan
	Mod *UpdateRecordPlanData `json:"Mod,omitnil,omitempty" name:"Mod"`
}

func (r *UpdateRecordPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlanId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordPlanResponseParams struct {
	// Returned result
	Data *RecordPlanOptData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRecordPlanResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordPlanResponseParams `json:"Response"`
}

func (r *UpdateRecordPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordTemplateData struct {
	// Template name. If you do not modify the name, you do not need to include this field.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Cloud upload time period. If you do not modify the name, you do not need to include this field.
	TimeSections []*RecordTemplateTimeSections `json:"TimeSections,omitnil,omitempty" name:"TimeSections"`
}

// Predefined struct for user
type UpdateRecordTemplateRequestParams struct {
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modify Content
	Mod *UpdateRecordTemplateData `json:"Mod,omitnil,omitempty" name:"Mod"`
}

type UpdateRecordTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Modify Content
	Mod *UpdateRecordTemplateData `json:"Mod,omitnil,omitempty" name:"Mod"`
}

func (r *UpdateRecordTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Mod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRecordTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRecordTemplateResponseParams struct {
	// Returned result
	Data *RecordTemplateInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRecordTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRecordTemplateResponseParams `json:"Response"`
}

func (r *UpdateRecordTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRecordTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDeviceRequestParams struct {
	// <p>Device ID (obtain from the device list API ListDevices)</p><p>For parameter values, see <a href="https://www.tencentcloud.com/document/api/1344/95871?from_cn_redirect=1">ListDevices</a></p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device name (only supports Chinese, English, numbers, spaces, Chinese and English parentheses, _, -, length not exceeding 128 characters)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device streaming protocol. Only valid for national standard devices. Fill in 0 to make no change (1: UDP, 2: TCP)</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password (only supported for national standard devices, length not exceeding 64 characters)</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description (length not exceeding 128 characters)</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device access IP (deprecated, reserved for compatibility, can be ignored)</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device Port (deprecated, reserved for compatibility, can be ignored)</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username (deprecated, reserved for compatibility, can be ignored)</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Gateway device access protocol (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Audio on/off (0: off; 1: on), enabled by default. Audio is dropped when turned off.</p>
	AudioSwitch *int64 `json:"AudioSwitch,omitnil,omitempty" name:"AudioSwitch"`

	// <p>Subscription switch (0: off; 1: on). Default On. When enabled, it subscribes to device channel transitions. Valid only for national standard NVR devices.</p>
	SubscribeSwitch *int64 `json:"SubscribeSwitch,omitnil,omitempty" name:"SubscribeSwitch"`

	// <p>Whether to enable silent frame (0: disabled; 1: enabled)</p>
	SilentFrameSwitch *int64 `json:"SilentFrameSwitch,omitnil,omitempty" name:"SilentFrameSwitch"`

	// <p>Clock sync switch (only effective for national standard devices)</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Default value: 1</p>
	TimeSyncSwitch *int64 `json:"TimeSyncSwitch,omitnil,omitempty" name:"TimeSyncSwitch"`
}

type UpdateUserDeviceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Device ID (obtain from the device list API ListDevices)</p><p>For parameter values, see <a href="https://www.tencentcloud.com/document/api/1344/95871?from_cn_redirect=1">ListDevices</a></p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device name (only supports Chinese, English, numbers, spaces, Chinese and English parentheses, _, -, length not exceeding 128 characters)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Device streaming protocol. Only valid for national standard devices. Fill in 0 to make no change (1: UDP, 2: TCP)</p>
	TransportProtocol *int64 `json:"TransportProtocol,omitnil,omitempty" name:"TransportProtocol"`

	// <p>Device password (only supported for national standard devices, length not exceeding 64 characters)</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// <p>Device description (length not exceeding 128 characters)</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Device access IP (deprecated, reserved for compatibility, can be ignored)</p>
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// <p>Device Port (deprecated, reserved for compatibility, can be ignored)</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>Device username (deprecated, reserved for compatibility, can be ignored)</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>Gateway device access protocol (deprecated, reserved for compatibility, can be ignored)</p>
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// <p>Audio on/off (0: off; 1: on), enabled by default. Audio is dropped when turned off.</p>
	AudioSwitch *int64 `json:"AudioSwitch,omitnil,omitempty" name:"AudioSwitch"`

	// <p>Subscription switch (0: off; 1: on). Default On. When enabled, it subscribes to device channel transitions. Valid only for national standard NVR devices.</p>
	SubscribeSwitch *int64 `json:"SubscribeSwitch,omitnil,omitempty" name:"SubscribeSwitch"`

	// <p>Whether to enable silent frame (0: disabled; 1: enabled)</p>
	SilentFrameSwitch *int64 `json:"SilentFrameSwitch,omitnil,omitempty" name:"SilentFrameSwitch"`

	// <p>Clock sync switch (only effective for national standard devices)</p><p>Enumeration values:</p><ul><li>0: Off</li><li>1: On</li></ul><p>Default value: 1</p>
	TimeSyncSwitch *int64 `json:"TimeSyncSwitch,omitnil,omitempty" name:"TimeSyncSwitch"`
}

func (r *UpdateUserDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Name")
	delete(f, "TransportProtocol")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Ip")
	delete(f, "Port")
	delete(f, "Username")
	delete(f, "ProtocolType")
	delete(f, "AudioSwitch")
	delete(f, "SubscribeSwitch")
	delete(f, "SilentFrameSwitch")
	delete(f, "TimeSyncSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDeviceResponseParams struct {
	// <p>Response data.</p>
	Data *UpdateDeviceData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserDeviceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserDeviceResponseParams `json:"Response"`
}

func (r *UpdateUserDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VLPlanInfo struct {
	// <p>Plan ID</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>Plan name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Plan description.</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>ID of the visual understanding template</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>Callback notification configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`

	// <p>File expiration time (seconds)</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Number of channels included in the plan</p>
	ChannelCount *int64 `json:"ChannelCount,omitnil,omitempty" name:"ChannelCount"`

	// <p>Creation time.</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>Last update time</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type VLPlanResultDetailInfo struct {
	// <p>Analysis result ID</p>
	VLPlanResultId *string `json:"VLPlanResultId,omitnil,omitempty" name:"VLPlanResultId"`

	// <p>Media type. The current fixed value is image</p>
	MediaType *string `json:"MediaType,omitnil,omitempty" name:"MediaType"`

	// <p>ID of the associated analysis plan</p>
	VLPlanId *string `json:"VLPlanId,omitnil,omitempty" name:"VLPlanId"`

	// <p>Plan name</p>
	VLPlanName *string `json:"VLPlanName,omitnil,omitempty" name:"VLPlanName"`

	// <p>Channel ID</p>
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// <p>Channel name</p>
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// <p>Device ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>Device name</p>
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// <p>Affiliated organization ID</p>
	OrganizationId *string `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// <p>Affiliated organization name</p>
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// <p>Analysis timestamp (seconds)</p>
	VLResultTimestamp *int64 `json:"VLResultTimestamp,omitnil,omitempty" name:"VLResultTimestamp"`

	// <p>Image mode screenshot timestamp (seconds)</p>
	SnapshotTimestamp *int64 `json:"SnapshotTimestamp,omitnil,omitempty" name:"SnapshotTimestamp"`

	// <p>Video clip start timestamp (seconds)</p>
	VideoStartTimestamp *int64 `json:"VideoStartTimestamp,omitnil,omitempty" name:"VideoStartTimestamp"`

	// <p>Video segment end timestamp (seconds)</p>
	VideoEndTimestamp *int64 `json:"VideoEndTimestamp,omitnil,omitempty" name:"VideoEndTimestamp"`

	// <p>Billing method for the number of frames in video mode</p>
	BilledFrameCount *int64 `json:"BilledFrameCount,omitnil,omitempty" name:"BilledFrameCount"`

	// <p>AI analysis result content (JSON string)</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Temporary screenshot download URL (valid for 300 seconds)</p>
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// <p>Screenshot URL validity period (seconds)</p>
	ImageUrlExpire *int64 `json:"ImageUrlExpire,omitnil,omitempty" name:"ImageUrlExpire"`

	// <p>Creation time, format: 2006-01-02 15:04:05</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>Update time, format 2006-01-02 15:04:05</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type VLPlanResultSnapshotImageUrlInfo struct {
	// <p>Analysis result ID</p>
	VLPlanResultId *string `json:"VLPlanResultId,omitnil,omitempty" name:"VLPlanResultId"`

	// <p>Source material download URL (empty if unsuccessful)</p>
	SourceFileURL *string `json:"SourceFileURL,omitnil,omitempty" name:"SourceFileURL"`

	// <p>URL validity period (seconds; 0 on failure)</p>
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`

	// <p>Error information (has a value on failure)</p>
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type VLTemplateInfo struct {
	// <p>Template ID</p>
	VLTemplateId *string `json:"VLTemplateId,omitnil,omitempty" name:"VLTemplateId"`

	// <p>AI model</p><p>Enumeration values:</p><ul><li>IVC-G: Default large model for view</li></ul>
	AIModel *string `json:"AIModel,omitnil,omitempty" name:"AIModel"`

	// <p>AI scenario name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>AI scenario description</p>
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// <p>Prompt content</p>
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// <p>Analysis mode</p><p>Enumeration values:</p><ul><li>image: Image mode</li><li>video: Video mode</li></ul>
	AnalysisMode *string `json:"AnalysisMode,omitnil,omitempty" name:"AnalysisMode"`

	// <p>Image scenario configuration</p>
	ImageConfig *ImageSceneInfo `json:"ImageConfig,omitnil,omitempty" name:"ImageConfig"`

	// <p>Effective period list</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Creation time.</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>Last update time</p>
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// <p>Whether associated with a visual understanding plan</p>
	IsUsed *bool `json:"IsUsed,omitnil,omitempty" name:"IsUsed"`

	// <p>Video scene configuration</p>
	VideoConfig *VideoSceneInfo `json:"VideoConfig,omitnil,omitempty" name:"VideoConfig"`
}

type VideoDownloadTask struct {
	// Download task ID
	DownloadTaskId *string `json:"DownloadTaskId,omitnil,omitempty" name:"DownloadTaskId"`

	// Channel ID
	ChannelId *string `json:"ChannelId,omitnil,omitempty" name:"ChannelId"`

	// Channel name
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// Channel encoding
	ChannelCode *string `json:"ChannelCode,omitnil,omitempty" name:"ChannelCode"`

	// Device name.
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// Device code
	DeviceCode *string `json:"DeviceCode,omitnil,omitempty" name:"DeviceCode"`

	// Task status (0: pending; 1: executing; 2: task complete;
	// 3: Task failure)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Download the recording time period
	VideoTimeSection *string `json:"VideoTimeSection,omitnil,omitempty" name:"VideoTimeSection"`

	// Playback speed
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// Download duration.
	DownloadTime *int64 `json:"DownloadTime,omitnil,omitempty" name:"DownloadTime"`

	// Video recording size
	VideoSize *int64 `json:"VideoSize,omitnil,omitempty" name:"VideoSize"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File download address
	FileDownloadUrl *string `json:"FileDownloadUrl,omitnil,omitempty" name:"FileDownloadUrl"`

	// Reason for failure
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// Lifecycle rule, number of days in hot storage
	Expire *int64 `json:"Expire,omitnil,omitempty" name:"Expire"`

	// mp4 preview address
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`
}

type VideoDownloadTaskData struct {
	// Download task ID
	DownloadTaskId *string `json:"DownloadTaskId,omitnil,omitempty" name:"DownloadTaskId"`
}

type VideoSceneInfo struct {
	// <p>Analysis interval (seconds): trigger an analysis every few seconds (sliding step length)</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Analysis segment duration (seconds): length of the video window fed into the model at a time, range [5, 12800]</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>Frame sampling rate. Default value: 1. Value range: [0.1, 1]. One frame is sampled every 1/Fps seconds and sent to the model.</p>
	Fps *float64 `json:"Fps,omitnil,omitempty" name:"Fps"`
}

type VideoSnapshotConfigInfo struct {
	// <p>Screenshot configuration Id</p>
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Screenshot configuration name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Image storage days</p><p>Unit: days</p>
	FileExpires *int64 `json:"FileExpires,omitnil,omitempty" name:"FileExpires"`

	// <p>Callback configuration</p>
	Callback *CallbackURL `json:"Callback,omitnil,omitempty" name:"Callback"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Last update time</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type VideoSnapshotTemplateInfo struct {
	// <p>Screenshot template Id</p>
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>Screenshot template name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Screenshot type</p><p>Enumeration values:</p><ul><li>1: Real-time frame capture</li><li>2: Low frequency frame capture</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Frame capturing frequency</p>
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// <p>Multiple mode</p>
	MultiMode *MultiVideoSnapshotMode `json:"MultiMode,omitnil,omitempty" name:"MultiMode"`

	// <p>Effective period</p>
	Times []*EffectiveTime `json:"Times,omitnil,omitempty" name:"Times"`

	// <p>Number of configurations associated with the template</p>
	RelatedConfigCount *int64 `json:"RelatedConfigCount,omitnil,omitempty" name:"RelatedConfigCount"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Last update time</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}