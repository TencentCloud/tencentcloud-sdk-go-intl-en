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

package v20201230

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccessInfo struct {
	// Address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Protocol.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type AccountInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Account name.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Account attribute.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Perms []*string `json:"Perms,omitnil,omitempty" name:"Perms"`
}

type CBSSpec struct {
	// Disk type.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type CBSSpecInfo struct {
	// Disk type.Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Size.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Number.Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type CNResourceSpec struct {
	// Node type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Model.
	// 
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Number of nodes.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Disk information.
	DiskSpec *CBSSpec `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`
}

type ChargeProperties struct {
	// 1: requires auto-renewal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Order time range.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit. Valid values: h and m.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Billing type: 0 indicates pay-as-you-go and 1 indicates monthly subscription.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// PREPAID and POSTPAID_BY_HOUR
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

type ConfigHistory struct {
	// id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Creation time.
	// 
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Update time.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// dn/cn
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Parameter name.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// New parameter value.
	ParamNewValue *string `json:"ParamNewValue,omitnil,omitempty" name:"ParamNewValue"`

	// Old parameter value.
	ParamOldValue *string `json:"ParamOldValue,omitnil,omitempty" name:"ParamOldValue"`

	// Status. Valid values: doing and success.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConfigParams struct {
	// Name.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// Value.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParameterValue *string `json:"ParameterValue,omitnil,omitempty" name:"ParameterValue"`

	// Value before modification.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParameterOldValue *string `json:"ParameterOldValue,omitnil,omitempty" name:"ParameterOldValue"`
}

// Predefined struct for user
type CreateInstanceByApiRequestParams struct {
	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Virtual Private Cloud (VPC).
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// Subnet.
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Billing method.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance password.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Resource information.
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tag list.Deprecated, use TagItems.
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version.
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	//  TagItems list.
	TagItems []*Tag `json:"TagItems,omitnil,omitempty" name:"TagItems"`
}

type CreateInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Availability zone.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Virtual Private Cloud (VPC).
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// Subnet.
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Billing method.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance password.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Resource information.
	Resources []*ResourceSpecNew `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tag list.Deprecated, use TagItems.
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Version.
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	//  TagItems list.
	TagItems []*Tag `json:"TagItems,omitnil,omitempty" name:"TagItems"`
}

func (r *CreateInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ChargeProperties")
	delete(f, "AdminPassword")
	delete(f, "Resources")
	delete(f, "Tags")
	delete(f, "ProductVersion")
	delete(f, "TagItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceByApiResponseParams struct {
	// Process ID.Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceByApiResponseParams `json:"Response"`
}

func (r *CreateInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Total number of instances.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Account array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Accounts []*AccountInfo `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBConfigHistoryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBConfigHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBConfigHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBConfigHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBConfigHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBConfigHistoryResponseParams struct {
	// Total count.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// DBConfig history.
	ConfigHistory []*ConfigHistory `json:"ConfigHistory,omitnil,omitempty" name:"ConfigHistory"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBConfigHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBConfigHistoryResponseParams `json:"Response"`
}

func (r *DescribeDBConfigHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBConfigHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParamsRequestParams struct {
	// cn/dn
	NodeTypes []*string `json:"NodeTypes,omitnil,omitempty" name:"NodeTypes"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBParamsRequest struct {
	*tchttp.BaseRequest
	
	// cn/dn
	NodeTypes []*string `json:"NodeTypes,omitnil,omitempty" name:"NodeTypes"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeTypes")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParamsResponseParams struct {
	// Total count.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameters information.
	Items []*ParamItem `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParamsResponseParams `json:"Response"`
}

func (r *DescribeDBParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeErrorLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeErrorLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeErrorLogResponseParams struct {
	// Total count of messages returned.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Error log details.
	ErrorLogDetails []*ErrorLogDetail `json:"ErrorLogDetails,omitnil,omitempty" name:"ErrorLogDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeErrorLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeErrorLogResponseParams `json:"Response"`
}

func (r *DescribeErrorLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceInfoResponseParams struct {
	// Instance description information.
	SimpleInstanceInfo *SimpleInstanceInfo `json:"SimpleInstanceInfo,omitnil,omitempty" name:"SimpleInstanceInfo"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// error msg
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Node list.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNodes []*InstanceNode `json:"InstanceNodes,omitnil,omitempty" name:"InstanceNodes"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsResponseParams struct {
	// Total count of operation records.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// operation records.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Operations []*InstanceOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// Instance description information.
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateRequestParams struct {
	//  InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	//  InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// Instance status. Example: serving.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of instance operation.Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Instance operation name.Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Instance operation progress.Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Instance status description. Example: running.Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Instance process error messages. Example: "Creation failed, insufficient resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// The name of the current step. Example: "Purchasing resources."Note: This field may return null, indicating that no valid values can be obtained.
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// Enabling status of the instance backup task.Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Searches by instance ID.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Searches by instance name.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searched tag list.
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Searches by instance ID.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Searches by instance name.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searched tag list.
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total count of instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// Error message.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleInstancesRequestParams struct {
	// Searches by instance ID.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Searches by instance name.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searches by tag list.
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeSimpleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Searches by instance ID.
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// Searches by instance name.
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Searches by tag list.
	SearchTags []*string `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

func (r *DescribeSimpleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSimpleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSimpleInstancesResponseParams struct {
	// Total count of instance lists.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance list details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstancesList []*InstanceSimpleInfoNew `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// Error message.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSimpleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSimpleInstancesResponseParams `json:"Response"`
}

func (r *DescribeSimpleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSimpleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Sorting method.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Ascending or descending order.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Duration.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Database.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Sorting method.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Ascending or descending order.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Duration.
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Database")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogResponseParams struct {
	// Total count of messages returned.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow SQL log details.
	SlowLogDetails *SlowLogDetail `json:"SlowLogDetails,omitnil,omitempty" name:"SlowLogDetails"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogResponseParams `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeListRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUpgradeListRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Offset.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUpgradeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpgradeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeListResponseParams struct {
	// Details of instance upgrade records.Note: This field may return null, indicating that no valid values can be obtained.
	UpgradeItems []*UpgradeItem `json:"UpgradeItems,omitnil,omitempty" name:"UpgradeItems"`

	// Total count of upgrade records.
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpgradeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpgradeListResponseParams `json:"Response"`
}

func (r *DescribeUpgradeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserHbaConfigRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeUserHbaConfigRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeUserHbaConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserHbaConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserHbaConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserHbaConfigResponseParams struct {
	// Total number of instances.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Hba Config array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HbaConfigs []*HbaConfig `json:"HbaConfigs,omitnil,omitempty" name:"HbaConfigs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserHbaConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserHbaConfigResponseParams `json:"Response"`
}

func (r *DescribeUserHbaConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserHbaConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiRequestParams struct {
	// Instance id. Example: "cdwpg-xxxx".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// Instance id. Example: "cdwpg-xxxx".
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiResponseParams struct {
	// Destroy  process ID.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceByApiResponseParams `json:"Response"`
}

func (r *DestroyInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpecPlus struct {
	// Number of disks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Maximum disk capacity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// Minimum disk capacity.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// Disk type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk type details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// Model type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmClass *string `json:"CvmClass,omitnil,omitempty" name:"CvmClass"`
}

type ErrorLogDetail struct {
	// Username.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// The time an error was reported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorTime *string `json:"ErrorTime,omitnil,omitempty" name:"ErrorTime"`

	// Error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type HbaConfig struct {
	// Type.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Database.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// User.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// IP address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Method.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Indicates whether to perform overwriting.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mask *string `json:"Mask,omitnil,omitempty" name:"Mask"`
}

type InstanceInfo struct {
	// Instance ID 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Kernel version type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Cluster name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Cluster status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Cluster status details.Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Cluster status information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil,omitempty" name:"InstanceStateInfo"`

	// Cluster ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// Creation time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Region.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Region details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Region details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Tag.Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Kernel version.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Character set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// CN node list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CNNodes []*InstanceNodeGroup `json:"CNNodes,omitnil,omitempty" name:"CNNodes"`

	// DN node list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DNNodes []*InstanceNodeGroup `json:"DNNodes,omitnil,omitempty" name:"DNNodes"`

	// Region ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Virtual Private Cloud (VPC).
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Expiration time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Billing mode.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Automatic renewal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Cluster ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Access information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessDetails []*AccessInfo `json:"AccessDetails,omitnil,omitempty" name:"AccessDetails"`
}

type InstanceNode struct {
	// id
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// cn
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// ip
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`
}

type InstanceNodeGroup struct {
	// Model.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Disk information.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDisk *DiskSpecPlus `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// Number of machines.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CvmCount *int64 `json:"CvmCount,omitnil,omitempty" name:"CvmCount"`
}

type InstanceOperation struct {
	// Operation name, such as create_instance, and scaleout_instance
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Cluster ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Operation name description, such as creating, and modifying the cluster name.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Status.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Operation start time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Operation end time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Operation context.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Operation update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Operation UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type InstanceSimpleInfoNew struct {
	// ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Cluster ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Kernel version.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Region ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Region details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Virtual Private Cloud (VPC).
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Start time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Expiration time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Access address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Billing mode.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Automatic renewal.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceStateInfo struct {
	// Instance status. Example: serving.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of instance operation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Instance operation name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Instance operation progress.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *int64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Instance status description. Example: running.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Instance process error messages. Example: "Creation failed, insufficient resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// The name of the current step. Example: "Purchasing resources."
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// Indicates whether there is a backup task in the instance. 1 indicates yes and 0 indicates no.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Request ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// Indicates whether there is a backup task in the cluster. 1 indicates yes and 0 indicates no.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupOpenStatus *int64 `json:"BackupOpenStatus,omitnil,omitempty" name:"BackupOpenStatus"`
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node parameter.
	NodeConfigParams []*NodeConfigParams `json:"NodeConfigParams,omitnil,omitempty" name:"NodeConfigParams"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node parameter.
	NodeConfigParams []*NodeConfigParams `json:"NodeConfigParams,omitnil,omitempty" name:"NodeConfigParams"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeConfigParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// Asynchronous process ID.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the newly modified instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the newly modified instance.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserHbaRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Hba array.
	HbaConfigs []*HbaConfig `json:"HbaConfigs,omitnil,omitempty" name:"HbaConfigs"`
}

type ModifyUserHbaRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Hba array.
	HbaConfigs []*HbaConfig `json:"HbaConfigs,omitnil,omitempty" name:"HbaConfigs"`
}

func (r *ModifyUserHbaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserHbaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "HbaConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserHbaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserHbaResponseParams struct {
	// Task ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Error message.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserHbaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserHbaResponseParams `json:"Response"`
}

func (r *ModifyUserHbaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserHbaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeConfigParams struct {
	// Node type.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Parameter.
	ConfigParams []*ConfigParams `json:"ConfigParams,omitnil,omitempty" name:"ConfigParams"`
}

type NormQueryItem struct {
	// Number of calls.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CallTimes *int64 `json:"CallTimes,omitnil,omitempty" name:"CallTimes"`

	// Number of read-only shared memory blocks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SharedReadBlocks *int64 `json:"SharedReadBlocks,omitnil,omitempty" name:"SharedReadBlocks"`

	// Number of write-only shared memory blocks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SharedWriteBlocks *int64 `json:"SharedWriteBlocks,omitnil,omitempty" name:"SharedWriteBlocks"`

	// Database.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Statement after masking.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NormalQuery *string `json:"NormalQuery,omitnil,omitempty" name:"NormalQuery"`

	// The statement with the longest execution time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxElapsedQuery *string `json:"MaxElapsedQuery,omitnil,omitempty" name:"MaxElapsedQuery"`

	// Total consumption time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CostTime *float64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Client IP address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// Username.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Proportion of total count.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCallTimesPercent *float64 `json:"TotalCallTimesPercent,omitnil,omitempty" name:"TotalCallTimesPercent"`

	// Proportion of total consumption time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCostTimePercent *float64 `json:"TotalCostTimePercent,omitnil,omitempty" name:"TotalCostTimePercent"`

	// Minimum consumption time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MinCostTime *float64 `json:"MinCostTime,omitnil,omitempty" name:"MinCostTime"`

	// Maximum consumption time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxCostTime *float64 `json:"MaxCostTime,omitnil,omitempty" name:"MaxCostTime"`

	// Time of the earliest item.Note: This field may return null, indicating that no valid values can be obtained.
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Time of the latest item.Note: This field may return null, indicating that no valid values can be obtained.
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// Total consumption time of I/O reading.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadCostTime *float64 `json:"ReadCostTime,omitnil,omitempty" name:"ReadCostTime"`

	// Total consumption time I/O writing.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WriteCostTime *float64 `json:"WriteCostTime,omitnil,omitempty" name:"WriteCostTime"`
}

type ParamDetail struct {
	// Parameter name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Default value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Indicates whether the restart is required.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NeedRestart *bool `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Current value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RunningValue *string `json:"RunningValue,omitnil,omitempty" name:"RunningValue"`

	// Value range.
	ValueRange *ValueRange `json:"ValueRange,omitnil,omitempty" name:"ValueRange"`

	// Unit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Introduction in English.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShortDesc *string `json:"ShortDesc,omitnil,omitempty" name:"ShortDesc"`

	// Parameter name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`
}

type ParamItem struct {
	// Node type. Valid values: cn and dn.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Node name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Number of parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Details []*ParamDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type Range struct {
	// Minimum value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// Maximum value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Instanceid.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The username to be modified.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New password.
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instanceid.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The username to be modified.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New password.
	NewPassword *string `json:"NewPassword,omitnil,omitempty" name:"NewPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "NewPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// Error message.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// Resource name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Resource count.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Disk information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSpec *CBSSpecInfo `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// Node type. Valid values: cn and dn.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ResourceSpecNew struct {
	// Resource name.
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Resource count.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Disk information.
	DiskSpec *CBSSpec `json:"DiskSpec,omitnil,omitempty" name:"DiskSpec"`

	// Resource type, DATA.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// Instance name. Example: cdwpg-xxxx.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Types of node that need to restart. Valid values: gtm, cn, dn and fn.
	NodeTypes []*string `json:"NodeTypes,omitnil,omitempty" name:"NodeTypes"`

	// Specifies th ID of nodes that need to restart.
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance name. Example: cdwpg-xxxx.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Types of node that need to restart. Valid values: gtm, cn, dn and fn.
	NodeTypes []*string `json:"NodeTypes,omitnil,omitempty" name:"NodeTypes"`

	// Specifies th ID of nodes that need to restart.
	NodeIds []*string `json:"NodeIds,omitnil,omitempty" name:"NodeIds"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeTypes")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// FlowId.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartInstanceResponseParams `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node type.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Number of scale-out nodes.
	ScaleOutCount *int64 `json:"ScaleOutCount,omitnil,omitempty" name:"ScaleOutCount"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node type.
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Number of scale-out nodes.
	ScaleOutCount *int64 `json:"ScaleOutCount,omitnil,omitempty" name:"ScaleOutCount"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeType")
	delete(f, "ScaleOutCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// FlowId.
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modifies the resource type.
	Case *string `json:"Case,omitnil,omitempty" name:"Case"`

	// Modified parameters.
	ModifySpec *CNResourceSpec `json:"ModifySpec,omitnil,omitempty" name:"ModifySpec"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modifies the resource type.
	Case *string `json:"Case,omitnil,omitempty" name:"Case"`

	// Modified parameters.
	ModifySpec *CNResourceSpec `json:"ModifySpec,omitnil,omitempty" name:"ModifySpec"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Case")
	delete(f, "ModifySpec")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// FlowId.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Specific error.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTags struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1 means only the Tag key is entered without a value, and 0 means both the key and the value are entered.
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SimpleInstanceInfo struct {
	// ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Cluster ID.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Kernel version.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Virtual Private Cloud (VPC).
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserVPCID *string `json:"UserVPCID,omitnil,omitempty" name:"UserVPCID"`

	// Subnet.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserSubnetID *string `json:"UserSubnetID,omitnil,omitempty" name:"UserSubnetID"`

	// Start time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Expiration time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Access address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Automatic renewal switch. 0 indicates automatic renewal is not enabled, and 1 indicates automatic renewal is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Billing mode.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Resource collection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Resources []*ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// Tag list.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Cluster status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type SlowLogDetail struct {
	// Total consumption time.
	TotalTime *float64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// Total number of calls.
	TotalCallTimes *int64 `json:"TotalCallTimes,omitnil,omitempty" name:"TotalCallTimes"`

	// Slow SQL.
	NormalQuerys []*NormQueryItem `json:"NormalQuerys,omitnil,omitempty" name:"NormalQuerys"`
}

type Tag struct {
	// Tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Installation package version.
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// InstanceId.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Installation package version.
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PackageVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// FlowId.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeItem struct {
	// Task name.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Original kernel version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceVersion *string `json:"SourceVersion,omitnil,omitempty" name:"SourceVersion"`

	// Target kernel version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetVersion *string `json:"TargetVersion,omitnil,omitempty" name:"TargetVersion"`

	// Task creation time.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task end time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Task completion status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Operator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

type ValueRange struct {
	// Parameter types. Valid values: enum, string, and section. Enum indicates enumeration, namely utf8, latin1, gbk. String indicates that the returned parameter value is a string. Section indicates that the returned parameter value is a value range, for example, 4-8.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Response parameter when the type is a section.Note: This field may return null, indicating that no valid values can be obtained.
	Range *Range `json:"Range,omitnil,omitempty" name:"Range"`

	// Response parameter when the type is an enum.Note: This field may return null, indicating that no valid values can be obtained.
	Enum []*string `json:"Enum,omitnil,omitempty" name:"Enum"`

	// Response parameter when the type is a string.Note: This field may return null, indicating that no valid values can be obtained.
	String *string `json:"String,omitnil,omitempty" name:"String"`
}