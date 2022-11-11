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

package v20210323

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CreateDomainRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Whether the domain is starred. Valid values: yes, no.
	IsMark *string `json:"IsMark,omitempty" name:"IsMark"`
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Whether the domain is starred. Valid values: yes, no.
	IsMark *string `json:"IsMark,omitempty" name:"IsMark"`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "GroupId")
	delete(f, "IsMark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainResponseParams struct {
	// Domain information
	DomainInfo *DomainCreateInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainResponseParams `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record type, which is obtained through the record type API. The value contains uppercase letters, such as `A`.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record split zone, which is obtained through the record split zone API.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// Record value, such as `IP : 200.200.200.200`, `CNAME : cname.dnspod.com`, and `MX : mail.dnspod.com`.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Host record such as `www`. If it is not passed in, it will be `@` by default.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Split zone ID, which is obtained through the record split zone API. The value is a string such as `10=1`. The `RecordLineId` parameter has a higher priority than `RecordLine`. If both of them are passed in, `RecordLineId` will be used first.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX priority, which is required for an MX record and will take effect if the record type is MX. Value range: 1–20.
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL. Value range: 1–604800. The minimum value varies by domain level.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Weight information, which is an integer between 0 and 100. It is supported only for enterprise VIP domains. `0` indicates not to pass in this parameter, i.e., not to set the weight.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Initial status of the record. Valid values: ENABLE, DISABLE. Default value: ENABLE. If `DISABLE` is passed in, the DNS record won't take effect, and the limit on round-robin DNS won't be verified.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type CreateRecordRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record type, which is obtained through the record type API. The value contains uppercase letters, such as `A`.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record split zone, which is obtained through the record split zone API.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// Record value, such as `IP : 200.200.200.200`, `CNAME : cname.dnspod.com`, and `MX : mail.dnspod.com`.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Host record such as `www`. If it is not passed in, it will be `@` by default.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Split zone ID, which is obtained through the record split zone API. The value is a string such as `10=1`. The `RecordLineId` parameter has a higher priority than `RecordLine`. If both of them are passed in, `RecordLineId` will be used first.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX priority, which is required for an MX record and will take effect if the record type is MX. Value range: 1–20.
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL. Value range: 1–604800. The minimum value varies by domain level.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Weight information, which is an integer between 0 and 100. It is supported only for enterprise VIP domains. `0` indicates not to pass in this parameter, i.e., not to set the weight.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Initial status of the record. Valid values: ENABLE, DISABLE. Default value: ENABLE. If `DISABLE` is passed in, the DNS record won't take effect, and the limit on round-robin DNS won't be verified.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *CreateRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "Value")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordLineId")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Weight")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecordResponseParams struct {
	// Record ID
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecordResponseParams `json:"Response"`
}

func (r *CreateRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
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
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteRecordRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DeleteRecordRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRecordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRecordResponseParams `json:"Response"`
}

func (r *DeleteRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeDomainRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
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
	delete(f, "Domain")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainResponseParams struct {
	// Domain information
	DomainInfo *DomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeRecordListRequestParams struct {
	// The domain for which DNS records are to be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The ID of the domain for which DNS records are to be obtained. If `DomainId` is passed in, the system will omit the parameter `Domain`.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// The split zone name.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// The split zone ID. If `RecordLineId` is passed in, the system will omit the parameter `RecordLine`.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// The keyword for searching for DNS records. Host headers and record values are supported.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// The sorting field. Available values: `name`, `line`, `type`, `value`, `weight`, `mx`, and `ttl,updated_on`.
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// The sorting type. Valid values: `ASC` (ascending, default), `DESC` (descending).
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// The offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The limit. It defaults to 100 and can be up to 3,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRecordListRequest struct {
	*tchttp.BaseRequest
	
	// The domain for which DNS records are to be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The ID of the domain for which DNS records are to be obtained. If `DomainId` is passed in, the system will omit the parameter `Domain`.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
	Subdomain *string `json:"Subdomain,omitempty" name:"Subdomain"`

	// The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// The split zone name.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// The split zone ID. If `RecordLineId` is passed in, the system will omit the parameter `RecordLine`.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// The group ID.
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// The keyword for searching for DNS records. Host headers and record values are supported.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// The sorting field. Available values: `name`, `line`, `type`, `value`, `weight`, `mx`, and `ttl,updated_on`.
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// The sorting type. Valid values: `ASC` (ascending, default), `DESC` (descending).
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// The offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The limit. It defaults to 100 and can be up to 3,000.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Subdomain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "RecordLineId")
	delete(f, "GroupId")
	delete(f, "Keyword")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordListResponseParams struct {
	// The record count info.
	RecordCountInfo *RecordCountInfo `json:"RecordCountInfo,omitempty" name:"RecordCountInfo"`

	// The record list result.
	RecordList []*RecordListItem `json:"RecordList,omitempty" name:"RecordList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordListResponseParams `json:"Response"`
}

func (r *DescribeRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeRecordRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordId")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordResponseParams struct {
	// Record information
	RecordInfo *RecordInfo `json:"RecordInfo,omitempty" name:"RecordInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordResponseParams `json:"Response"`
}

func (r *DescribeRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCreateInfo struct {
	// Domain ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain Punycode
	Punycode *string `json:"Punycode,omitempty" name:"Punycode"`

	// NS list of the domain
	GradeNsList []*string `json:"GradeNsList,omitempty" name:"GradeNsList"`
}

type DomainInfo struct {
	// Domain ID
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Domain status
	Status *string `json:"Status,omitempty" name:"Status"`

	// DNS plan level
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// Domain group ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// Whether the domain is starred
	IsMark *string `json:"IsMark,omitempty" name:"IsMark"`

	// TTL (DNS record cache time)
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Whether CNAME flattening is enabled
	CnameSpeedup *string `json:"CnameSpeedup,omitempty" name:"CnameSpeedup"`

	// Domain remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Domain Punycode
	Punycode *string `json:"Punycode,omitempty" name:"Punycode"`

	// DNS status of the domain
	DnsStatus *string `json:"DnsStatus,omitempty" name:"DnsStatus"`

	// NS list of the domain
	DnspodNsList []*string `json:"DnspodNsList,omitempty" name:"DnspodNsList"`

	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain level ID
	GradeLevel *uint64 `json:"GradeLevel,omitempty" name:"GradeLevel"`

	// Domain user ID
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`

	// Whether the domain is a VIP domain
	IsVip *string `json:"IsVip,omitempty" name:"IsVip"`

	// Domain owner account
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// Domain level description
	GradeTitle *string `json:"GradeTitle,omitempty" name:"GradeTitle"`

	// Domain creation time
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// Last update time
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// Tencent Cloud account `Uin`
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// NS list actually used by the domain
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActualNsList []*string `json:"ActualNsList,omitempty" name:"ActualNsList"`

	// Number of domain records
	RecordCount *uint64 `json:"RecordCount,omitempty" name:"RecordCount"`

	// Alias of the domain account owner
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerNick *string `json:"OwnerNick,omitempty" name:"OwnerNick"`
}

// Predefined struct for user
type ModifyDomainRemarkRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Domain remarks. To delete the remarks, submit empty content.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyDomainRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Domain remarks. To delete the remarks, submit empty content.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyDomainRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainRemarkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainRemarkResponseParams `json:"Response"`
}

func (r *ModifyDomainRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainStatusRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain status. Valid values: enable; disable.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type ModifyDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Domain status. Valid values: enable; disable.
	Status *string `json:"Status,omitempty" name:"Status"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *ModifyDomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainStatusResponseParams `json:"Response"`
}

func (r *ModifyDomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordRequestParams struct {
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record type, which is obtained through the record type API. The value contains uppercase letters, such as `A`.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record split zone, which is obtained through the record split zone API.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// Record value, such as `IP : 200.200.200.200`, `CNAME : cname.dnspod.com`, and `MX : mail.dnspod.com`.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Host record such as `www`. If it is not passed in, it will be `@` by default.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Split zone ID, which is obtained through the record split zone API. The value is a string such as `10=1`. The `RecordLineId` parameter has a higher priority than `RecordLine`. If both of them are passed in, `RecordLineId` will be used first.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX priority, which is required for an MX record and will take effect if the record type is MX. Value range: 1–20.
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL. Value range: 1–604800. The minimum value varies by domain level.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Weight information, which is an integer between 0 and 100. It is supported only for enterprise VIP domains. `0` indicates not to pass in this parameter, i.e., not to set the weight.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Initial status of the record. Valid values: ENABLE, DISABLE. Default value: ENABLE. If `DISABLE` is passed in, the DNS record won't take effect, and the limit on round-robin DNS won't be verified.
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyRecordRequest struct {
	*tchttp.BaseRequest
	
	// Domain
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Record type, which is obtained through the record type API. The value contains uppercase letters, such as `A`.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Record split zone, which is obtained through the record split zone API.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// Record value, such as `IP : 200.200.200.200`, `CNAME : cname.dnspod.com`, and `MX : mail.dnspod.com`.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// Domain ID. The `DomainId` parameter has a higher priority than `Domain`. If `DomainId` is passed in, `Domain` will be ignored.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`

	// Host record such as `www`. If it is not passed in, it will be `@` by default.
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Split zone ID, which is obtained through the record split zone API. The value is a string such as `10=1`. The `RecordLineId` parameter has a higher priority than `RecordLine`. If both of them are passed in, `RecordLineId` will be used first.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// MX priority, which is required for an MX record and will take effect if the record type is MX. Value range: 1–20.
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL. Value range: 1–604800. The minimum value varies by domain level.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Weight information, which is an integer between 0 and 100. It is supported only for enterprise VIP domains. `0` indicates not to pass in this parameter, i.e., not to set the weight.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// Initial status of the record. Valid values: ENABLE, DISABLE. Default value: ENABLE. If `DISABLE` is passed in, the DNS record won't take effect, and the limit on round-robin DNS won't be verified.
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RecordType")
	delete(f, "RecordLine")
	delete(f, "Value")
	delete(f, "RecordId")
	delete(f, "DomainId")
	delete(f, "SubDomain")
	delete(f, "RecordLineId")
	delete(f, "MX")
	delete(f, "TTL")
	delete(f, "Weight")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRecordResponseParams struct {
	// Record ID
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRecordResponseParams `json:"Response"`
}

func (r *ModifyRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordCountInfo struct {
	// The subdomain count.
	SubdomainCount *uint64 `json:"SubdomainCount,omitempty" name:"SubdomainCount"`

	// The count of records returned in the list.
	ListCount *uint64 `json:"ListCount,omitempty" name:"ListCount"`

	// The total record count.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type RecordInfo struct {
	// Record ID.
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Subdomain (host record).
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// Record type. For more information, see the `DescribeRecordType` API.
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// Split zone of the DNS record. For more information, see the `DescribeRecordLineList` API.
	RecordLine *string `json:"RecordLine,omitempty" name:"RecordLine"`

	// Split zone ID of the DNS record. For more information, see the `DescribeRecordLineList` API.
	RecordLineId *string `json:"RecordLineId,omitempty" name:"RecordLineId"`

	// Record value.
	Value *string `json:"Value,omitempty" name:"Value"`

	// Record weight.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// MX record value. It is 0 by default for non-MX records.
	MX *uint64 `json:"MX,omitempty" name:"MX"`

	// TTL value of the record.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// Record status. Valid values: 0 (disabled); 1 (enabled).
	Enabled *uint64 `json:"Enabled,omitempty" name:"Enabled"`

	// D-Monitor status of the record.
	// "Ok" : The server is normal.
	// "Warn" : There is an alarm on this record, and the server returns 4XX.
	// "Down" : The server is down.
	// "" : D-Monitor is disabled for this record.
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// Record remarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Last update time of the record.
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// Domain ID.
	DomainId *uint64 `json:"DomainId,omitempty" name:"DomainId"`
}

type RecordListItem struct {
	// The record ID.
	RecordId *uint64 `json:"RecordId,omitempty" name:"RecordId"`

	// The record value.
	Value *string `json:"Value,omitempty" name:"Value"`

	// The record status. Valid values: `ENABLE` (enabled), `DISABLE` (disabled).
	Status *string `json:"Status,omitempty" name:"Status"`

	// The update time.
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// The host name.
	Name *string `json:"Name,omitempty" name:"Name"`

	// The record split zone.
	Line *string `json:"Line,omitempty" name:"Line"`

	// The split zone ID.
	LineId *string `json:"LineId,omitempty" name:"LineId"`

	// The record type.
	Type *string `json:"Type,omitempty" name:"Type"`

	// The record weight, which is required for round-robin DNS records.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// The monitoring status of the record. Valid values: `OK` (normal), `WARN` (warning), and `DOWN` (downtime). It is empty if no monitoring is set or the monitoring is suspended.
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// The record remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// The record cache time.
	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`

	// The MX value, applicable to the MX record only.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MX *uint64 `json:"MX,omitempty" name:"MX"`
}