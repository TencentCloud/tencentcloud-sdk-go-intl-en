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

package v20230110

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccountFactoryItem struct {
	// Specifies the unique identifier for account factory baseline item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Baseline item name. specifies a unique name for the feature item. supports a combination of english letters, numbers, chinese characters, and symbols @, &, _, [, ], -. valid values: 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline item english name. specifies a unique name for the baseline item. supports a combination of english letters, digits, spaces, and symbols @, &, _, [], -. valid values: 1-64 english characters.
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// Baseline item weight. the smaller the value, the higher the weight. value range equal to or greater than 0.
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// Specifies whether the baseline item is required (1: required; 0: optional).
	Required *int64 `json:"Required,omitnil,omitempty" name:"Required"`

	// Baseline item dependency. value range of N depends on the count of other baseline items it relies on.
	DependsOn []*DependsOnItem `json:"DependsOn,omitnil,omitempty" name:"DependsOn"`

	// Baseline description, with a length of 2 to 256 english or chinese characters. it is empty by default.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Baseline item english description, with a length of 2 to 1024 english characters. it is empty by default.
	DescriptionEn *string `json:"DescriptionEn,omitnil,omitempty" name:"DescriptionEn"`

	// Baseline classification. length: 2-32 english or chinese characters. values cannot be empty.
	Classify *string `json:"Classify,omitnil,omitempty" name:"Classify"`

	// Baseline english classification, with a length of 2-64 english characters. cannot be empty.
	ClassifyEn *string `json:"ClassifyEn,omitnil,omitempty" name:"ClassifyEn"`
}

type BaselineConfigItem struct {
	// Specifies the unique identifier for account factory baseline item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Account factory baseline item configuration. different baseline item configuration parameters.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`
}

type BaselineInfoItem struct {
	// Specifies the unique identifier for account factory baseline item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Account factory baseline item configuration. different baseline item configuration parameters.
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// Specifies the number of accounts for baseline applications.
	ApplyCount *int64 `json:"ApplyCount,omitnil,omitempty" name:"ApplyCount"`
}

type BaselineStepTaskInfo struct {
	// Specifies the unique Id of the task, which can only contain english letters and digits, and is a 16-character random string.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Specifies the unique identifier for the baseline feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Specifies the member account uin of the applied baseline item.
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// Baseline item application status. Running means the baseline item is in application. Success means the baseline item application is successful. Failed means the baseline item application failure. Pending means the baseline item is Pending application. Skipped means the baseline item is Skipped.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrCode *string `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// Error message
	ErrMessage *string `json:"ErrMessage,omitnil,omitempty" name:"ErrMessage"`

	// Baseline item deployment output.
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Creation time, represented in ISO8601 standard format as yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the last update time in ISO8601 standard representation with format yyyy-MM-dd hh:MM:ss.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type BatchApplyAccountBaselinesRequestParams struct {
	// Member account UIN, which is also the UIN of the account to which the baseline is applied.
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// List of baseline item configuration information.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type BatchApplyAccountBaselinesRequest struct {
	*tchttp.BaseRequest
	
	// Member account UIN, which is also the UIN of the account to which the baseline is applied.
	MemberUinList []*int64 `json:"MemberUinList,omitnil,omitempty" name:"MemberUinList"`

	// List of baseline item configuration information.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *BatchApplyAccountBaselinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUinList")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchApplyAccountBaselinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchApplyAccountBaselinesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchApplyAccountBaselinesResponse struct {
	*tchttp.BaseResponse
	Response *BatchApplyAccountBaselinesResponseParams `json:"Response"`
}

func (r *BatchApplyAccountBaselinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchApplyAccountBaselinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependsOnItem struct {
	// Dependency type. valid values: LandingZoneSetUp or AccountFactorySetUp. LandingZoneSetUp refers to the dependency of landingZone. AccountFactorySetUp refers to the dependency of account factory.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specifies the unique identifier for the feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`
}

// Predefined struct for user
type GetAccountFactoryBaselineRequestParams struct {

}

type GetAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountFactoryBaselineResponseParams struct {
	// Specifies the uin of the main account to which the resource belongs.
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Specifies the baseline item name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// List of baseline item configurations.
	BaselineConfigItems []*BaselineInfoItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`

	// Creation time, represented in ISO8601 standard format as yyyy-MM-dd hh:MM:ss.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the last update time in ISO8601 standard representation with format yyyy-MM-dd hh:MM:ss.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *GetAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *GetAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsRequestParams struct {
	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListAccountFactoryBaselineItemsRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListAccountFactoryBaselineItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccountFactoryBaselineItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccountFactoryBaselineItemsResponseParams struct {
	// Account factory baseline list.
	BaselineItems []*AccountFactoryItem `json:"BaselineItems,omitnil,omitempty" name:"BaselineItems"`

	// Total quantity.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAccountFactoryBaselineItemsResponse struct {
	*tchttp.BaseResponse
	Response *ListAccountFactoryBaselineItemsResponseParams `json:"Response"`
}

func (r *ListAccountFactoryBaselineItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccountFactoryBaselineItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksRequestParams struct {
	// Specifies the unique identifier for the feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListDeployStepTasksRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the unique identifier for the feature item, can only contain english letters, digits, and @, ,._[]-:()()[]+=., with a length of 2-128 characters.
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// Maximum number of returned records. value ranges from 0 to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. valid values are equal to or greater than 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListDeployStepTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDeployStepTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDeployStepTasksResponseParams struct {
	// Account factory baseline function application information list.
	BaselineDeployStepTaskList []*BaselineStepTaskInfo `json:"BaselineDeployStepTaskList,omitnil,omitempty" name:"BaselineDeployStepTaskList"`

	// Total quantity.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDeployStepTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListDeployStepTasksResponseParams `json:"Response"`
}

func (r *ListDeployStepTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDeployStepTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineRequestParams struct {
	// Specifies the baseline name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline configuration. overwrite update. can be accessed through controlcenter:GetAccountFactoryBaseline to query existing baseline configuration. can be accessed through controlcenter:ListAccountFactoryBaselineItems to query supported baseline list.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

type UpdateAccountFactoryBaselineRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the baseline name, which must be unique and can only contain a combination of english letters, digits, chinese characters, and symbols @, &_[]-, with a length of 1-25 chinese or english characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Baseline configuration. overwrite update. can be accessed through controlcenter:GetAccountFactoryBaseline to query existing baseline configuration. can be accessed through controlcenter:ListAccountFactoryBaselineItems to query supported baseline list.
	BaselineConfigItems []*BaselineConfigItem `json:"BaselineConfigItems,omitnil,omitempty" name:"BaselineConfigItems"`
}

func (r *UpdateAccountFactoryBaselineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BaselineConfigItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccountFactoryBaselineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccountFactoryBaselineResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccountFactoryBaselineResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccountFactoryBaselineResponseParams `json:"Response"`
}

func (r *UpdateAccountFactoryBaselineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccountFactoryBaselineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}