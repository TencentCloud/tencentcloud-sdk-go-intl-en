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

type AccountVpcInfoOut struct {
	// VpcId: vpc-xadsafsdasd
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Region: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// VPC ID: 123456789
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// VPC name: testname
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

// Predefined struct for user
type CreatePrivateDNSAccountRequestParams struct {
	// Private DNS account
	Account *PrivateDNSAccount `json:"Account,omitempty" name:"Account"`
}

type CreatePrivateDNSAccountRequest struct {
	*tchttp.BaseRequest
	
	// Private DNS account
	Account *PrivateDNSAccount `json:"Account,omitempty" name:"Account"`
}

func (r *CreatePrivateDNSAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateDNSAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateDNSAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrivateDNSAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateDNSAccountResponseParams `json:"Response"`
}

func (r *CreatePrivateDNSAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrivateDNSAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneRecordRequestParams struct {
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

// Predefined struct for user
type CreatePrivateZoneRecordResponseParams struct {
	// Record ID
	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateZoneRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePrivateZoneRequestParams struct {
	// Domain name, which must be in the format of standard TLD
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Tags the private domain when it is created
	TagSet []*TagInfo `json:"TagSet,omitempty" name:"TagSet"`

	// Associates the private domain to a VPC when it is created
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable subdomain recursive DNS. Valid values: `ENABLED` (default) and `DISABLED`.
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Associates the private domain to a VPC when it is created
	Vpcs []*VpcInfo `json:"Vpcs,omitempty" name:"Vpcs"`

	// List of authorized accounts' VPCs to associate with the private domain
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`

	// Whether to enable CNAME flattening. Valid values: `ENABLED` (default) and `DISABLED`.
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
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

	// Whether to enable subdomain recursive DNS. Valid values: `ENABLED` (default) and `DISABLED`.
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Associates the private domain to a VPC when it is created
	Vpcs []*VpcInfo `json:"Vpcs,omitempty" name:"Vpcs"`

	// List of authorized accounts' VPCs to associate with the private domain
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`

	// Whether to enable CNAME flattening. Valid values: `ENABLED` (default) and `DISABLED`.
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
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
	delete(f, "CnameSpeedupStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrivateZoneResponseParams struct {
	// Private domain ID, such as zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Private domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrivateZoneResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAccountVpcListRequestParams struct {
	// UIN of account
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

	// Pagination offset, starting from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: `100`. Default value: `20`
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAccountVpcListRequest struct {
	*tchttp.BaseRequest
	
	// UIN of account
	AccountUin *string `json:"AccountUin,omitempty" name:"AccountUin"`

	// Pagination offset, starting from 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: `100`. Default value: `20`
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAccountVpcListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountVpcListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountVpcListResponseParams struct {
	// Number of VPCs
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// VPC list
	VpcSet []*AccountVpcInfoOut `json:"VpcSet,omitempty" name:"VpcSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountVpcListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountVpcListResponseParams `json:"Response"`
}

func (r *DescribeAccountVpcListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountVpcListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogRequestParams struct {
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

// Predefined struct for user
type DescribeAuditLogResponseParams struct {
	// List of operation logs
	Data []*AuditLog `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeDashboardRequestParams struct {

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

// Predefined struct for user
type DescribeDashboardResponseParams struct {
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
}

type DescribeDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDashboardResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePrivateDNSAccountListRequestParams struct {
	// Pagination offset, starting from `0`
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of entries per page. Maximum value: `100`. Default value: `20`
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Filter parameters
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribePrivateDNSAccountListResponseParams struct {
	// Number of Private DNS accounts
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// List of Private DNS accounts
	AccountSet []*PrivateDNSAccount `json:"AccountSet,omitempty" name:"AccountSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivateDNSAccountListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateDNSAccountListResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePrivateZoneServiceRequestParams struct {

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

// Predefined struct for user
type DescribePrivateZoneServiceResponseParams struct {
	// Private DNS service activation status. Valid values: ENABLED, DISABLED
	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrivateZoneServiceResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeQuotaUsageRequestParams struct {

}

type DescribeQuotaUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeQuotaUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaUsageResponseParams struct {
	// TLD quota usage
	TldQuota *TldQuota `json:"TldQuota,omitempty" name:"TldQuota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQuotaUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaUsageResponseParams `json:"Response"`
}

func (r *DescribeQuotaUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRequestDataRequestParams struct {
	// Request volume statistics start time in the format of 2020-11-22 00:00:00
	TimeRangeBegin *string `json:"TimeRangeBegin,omitempty" name:"TimeRangeBegin"`

	// Filter parameter:
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Request volume statistics end time in the format of 2020-11-22 23:59:59
	TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" name:"TimeRangeEnd"`
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

// Predefined struct for user
type DescribeRequestDataResponseParams struct {
	// Request volume statistics table
	Data []*MetricData `json:"Data,omitempty" name:"Data"`

	// Request volume unit time. Valid values: Day, Hour
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRequestDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRequestDataResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPrivateZoneRecordRequestParams struct {
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

// Predefined struct for user
type ModifyPrivateZoneRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneRecordResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPrivateZoneRequestParams struct {
	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Whether to enable CNAME flattening. Valid values: `ENABLED` and `DISABLED`.
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
}

type ModifyPrivateZoneRequest struct {
	*tchttp.BaseRequest
	
	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// Remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED
	DnsForwardStatus *string `json:"DnsForwardStatus,omitempty" name:"DnsForwardStatus"`

	// Whether to enable CNAME flattening. Valid values: `ENABLED` and `DISABLED`.
	CnameSpeedupStatus *string `json:"CnameSpeedupStatus,omitempty" name:"CnameSpeedupStatus"`
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
	delete(f, "CnameSpeedupStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateZoneResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPrivateZoneVpcRequestParams struct {
	// Private domain ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of all VPCs associated with private domain
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// List of authorized accounts' VPCs to associate with the private domain
	AccountVpcSet []*AccountVpcInfo `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`
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

// Predefined struct for user
type ModifyPrivateZoneVpcResponseParams struct {
	// Private domain ID, such as zone-xxxxxx
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// List of VPCs associated with domain
	VpcSet []*VpcInfo `json:"VpcSet,omitempty" name:"VpcSet"`

	// List of authorized accounts' VPCs associated with the private domain
	AccountVpcSet []*AccountVpcInfoOutput `json:"AccountVpcSet,omitempty" name:"AccountVpcSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPrivateZoneVpcResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateZoneVpcResponseParams `json:"Response"`
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

// Predefined struct for user
type SubscribePrivateZoneServiceRequestParams struct {

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

// Predefined struct for user
type SubscribePrivateZoneServiceResponseParams struct {
	// Private DNS service activation status
	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubscribePrivateZoneServiceResponse struct {
	*tchttp.BaseResponse
	Response *SubscribePrivateZoneServiceResponseParams `json:"Response"`
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

type TldQuota struct {
	// Total quota
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Used quota
	Used *int64 `json:"Used,omitempty" name:"Used"`

	// Available quota
	Stock *int64 `json:"Stock,omitempty" name:"Stock"`

	// User’s quota
	Quota *int64 `json:"Quota,omitempty" name:"Quota"`
}

type VpcInfo struct {
	// VpcId: vpc-xadsafsdasd
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC region: ap-guangzhou, ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}