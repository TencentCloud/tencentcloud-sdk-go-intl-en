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

package v20201028

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AccountVpcInfo struct {

	// VpcId: vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC region: ap-guangzhou, ap-shanghai
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitempty" name:"Region"`

	// VPC account: 123456789
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// VPC name: testname
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
}

type AccountVpcInfoOutput struct {

	// UIN of the VPC account
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// Region
	Region *string `json:"Region,omitempty" name:"Region"`
}

type AuditLog struct {

	// Log type
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Log table name
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Total number of logs
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of logs
	DataSet []*AuditLogInfo `json:"DataSet,omitempty" name:"DataSet"`
}

type AuditLogInfo struct {

	// Time
	Date *string `json:"Date,omitempty" name:"Date"`

	// Operator UIN
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// Log content
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CreatePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record type. Valid values: "A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Subdomain, such as "www", "m", and "@"
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Record value, such as IP: 192.168.10.2, CNAME: cname.qcloud.com., and MX: mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// Record weight. Value range: 1–100
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// MX priority, which is required when the record type is MX. Valid values: 5, 10, 15, 20, 30, 40, 50
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// Record cache time. The smaller the value, the faster the record will take effect. Value range: 1–86400s. Default value: 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *CreatePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Record ID
		RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// Domain name, which must be in the format of standard TLD
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Tags the private domain when it is created
	TagSet []*TagInfo `json:"TagSet,omitempty" name:"TagSet"`

	// Associates the private domain to a VPC when it is created
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Associates the private domain to a VPC when it is created
	Vpcs []*VpcInfo `json:"Vpcs,omitempty" name:"Vpcs"`

	// List of authorized accounts' VPCs to associate with the private domain
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

func (r *CreatePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "TagSet")
	delete(f, "VpcSet")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	delete(f, "Vpcs")
	delete(f, "AccountVpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Private domain ID, such as zone-xxxxxx
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// Private domain
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DatePoint struct {

	// Time
	Date *string `json:"Date,omitempty" name:"Date"`

	// Value
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DeletePrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Array of record IDs. `RecordId` takes precedence.
	RecordIdSet []*string `json:"RecordIdSet,omitempty" name:"RecordIdSet"`
}

func (r *DeletePrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Array of private domain IDs. `ZoneId` takes precedence.
	ZoneIdSet []*string `json:"ZoneIdSet,omitempty" name:"ZoneIdSet"`
}

func (r *DeletePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ZoneIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeletePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogRequest struct {
	*tchttp.BaseRequest

	// Request volume statistics start time
	TimeRangeBegin *string `json:"TimeRangeBegin,omitempty" name:"TimeRangeBegin"`

	// Filter parameter. Valid values: ZoneId (private domain ID), Domain (private domain), OperatorUin (operator account ID)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Request volume statistics end time
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`

	// Pagination offset, starting from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: 100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of operation logs
		Data []*AuditLog `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of private domain DNS records
		ZoneTotal *int64 `json:"ZoneTotal,omitempty" name:"ZoneTotal"`

		// Number of VPCs associated with private domain
		ZoneVpcCount *int64 `json:"ZoneVpcCount,omitempty" name:"ZoneVpcCount"`

		// Total number of historical requests
		RequestTotalCount *int64 `json:"RequestTotalCount,omitempty" name:"RequestTotalCount"`

		// Traffic package usage
		FlowUsage []*FlowUsage `json:"FlowUsage,omitempty" name:"FlowUsage"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDNSAccountListRequest struct {
	*tchttp.BaseRequest

	// Pagination offset, starting from `0`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: `100`. Default value: `20`
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivateDNSAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateDNSAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateDNSAccountListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of Private DNS accounts
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of Private DNS accounts
		AccountSet []*PrivateDNSAccount `json:"AccountSet,omitempty" name:"AccountSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateDNSAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateDNSAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneListRequest struct {
	*tchttp.BaseRequest

	// Pagination offset, starting from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: 100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribePrivateZoneListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of private domains
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of private domains
		PrivateZoneSet []*PrivateZone `json:"PrivateZoneSet,omitempty" name:"PrivateZoneSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRecordListRequest struct {
	*tchttp.BaseRequest

	// Private domain ID: zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Filter parameter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Pagination offset, starting from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: 100. Default value: 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePrivateZoneRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRecordListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of DNS records
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of DNS records
		RecordSet []*PrivateZoneRecord `json:"RecordSet,omitempty" name:"RecordSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneRequest struct {
	*tchttp.BaseRequest

	// Domain name, which must be in the format of standard TLD
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribePrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Private domain details
		PrivateZone *PrivateZone `json:"PrivateZone,omitempty" name:"PrivateZone"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Private DNS service activation status. Valid values: ENABLED, DISABLED
		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRequestDataRequest struct {
	*tchttp.BaseRequest

	// Request volume statistics start time in the format of 2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitempty" name:"TimeRangeBegin"`

	// Filter parameter:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Request volume statistics end time in the format of 2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`
}

func (r *DescribeRequestDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRangeBegin")
	delete(f, "Filters")
	delete(f, "TimeRangeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRequestDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRequestDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Request volume statistics table
		Data []*MetricData `json:"Data,omitempty" name:"Data"`

		// Request volume unit time. Valid values: Day, Hour
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRequestDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRequestDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Array of parameter values
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FlowUsage struct {

	// Traffic package type, Valid values: ZONE (private domain); TRAFFIC (DNS traffic package)
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// Traffic package quota
	TotalQuantity *int64 `json:"TotalQuantity,omitempty" name:"TotalQuantity"`

	// Available quota of traffic package
	AvailableQuantity *int64 `json:"AvailableQuantity,omitempty" name:"AvailableQuantity"`
}

type MetricData struct {

	// Resource description
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// Table name
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// Table data
	DataSet []*DatePoint `json:"DataSet,omitempty" name:"DataSet"`
}

type ModifyPrivateZoneRecordRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Record ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Record type. Valid values: "A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Subdomain, such as "www", "m", and "@"
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Record value, such as IP: 192.168.10.2, CNAME: cname.qcloud.com., and MX: mail.qcloud.com.
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// Record weight. Value range: 1–100
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// MX priority, which is required when the record type is MX. Valid values: 5, 10, 15, 20, 30, 40, 50
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// Record cache time. The smaller the value, the faster the record will take effect. Value range: 1–86400s. Default value: 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`
}

func (r *ModifyPrivateZoneRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordId")
	delete(f, "RecordType")
	delete(f, "SubDomain")
	delete(f, "RecordValue")
	delete(f, "Weight")
	delete(f, "MX")
	delete(f, "TTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`
}

func (r *ModifyPrivateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Remark")
	delete(f, "DnsForwardStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneVpcRequest struct {
	*tchttp.BaseRequest

	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of all VPCs associated with private domain
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// List of authorized accounts' VPCs to associate with the private domain
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

func (r *ModifyPrivateZoneVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "VpcSet")
	delete(f, "AccountVpcSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Private domain ID, such as zone-xxxxxx
		ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

		// List of VPCs associated with domain
		VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

		// List of authorized accounts' VPCs associated with the private domain
		AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPrivateZoneVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateZoneVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PrivateDNSAccount struct {

	// Root account UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Root account name
	Account *string `json:"Account,omitempty" name:"Account"`

	// Account name
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`
}

type PrivateZone struct {

	// Private domain ID: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Domain name owner UIN
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Private domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Modification time
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// Number of results
	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// List of bound VPCs
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// Private domain status. Valid values: ENABLED (DNS enabled); SUSPEND (DNS paused); FROZEN (locked)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Subdomain recursive DNS status. Valid values: ENABLED, DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Set of tag key-value pairs
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// List of authorized accounts' VPCs associated with the private domain
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
}

type PrivateZoneRecord struct {

	// Record ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// Private domain ID: zone-xxxxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Subdomain
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Record type. Valid values: "A", "AAAA", "CNAME", "MX", "TXT", "PTR"
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record value
	RecordValue *string `json:"RecordValue,omitempty" name:"RecordValue"`

	// Record cache time. The smaller the value, the faster the record will take effect. Value range: 1–86400s. Default value: 600
	TTL *int64 `json:"TTL,omitempty" name:"TTL"`

	// MX priority, which is required when the record type is MX. Valid values: 5, 10, 15, 20, 30, 40, 50
	// Note: this field may return null, indicating that no valid values can be obtained.
	MX *int64 `json:"MX,omitempty" name:"MX"`

	// Record status: ENABLED
	Status *string `json:"Status,omitempty" name:"Status"`

	// Record weight. Value range: 1–100
	// Note: this field may return null, indicating that no valid values can be obtained.
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// Record creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Record update time
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// Additional information
	// Note: this field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitempty" name:"Extra"`
}

type SubscribePrivateZoneServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *SubscribePrivateZoneServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubscribePrivateZoneServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SubscribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Private DNS service activation status
		ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribePrivateZoneServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubscribePrivateZoneServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type VpcInfo struct {

	// VpcId: vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC region: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}
