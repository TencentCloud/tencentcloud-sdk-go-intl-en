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

package v20191205

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ApiGatewayInstanceDetail struct {
	// The instance ID.
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// The instance name.
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type ApiGatewayInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of APIGATEWAY instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*ApiGatewayInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of APIGATEWAY instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type ApplyCertificateRequestParams struct {
	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation.
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate type. Currently, the only supported value is 2, which indicates TrustAsia TLS RSA CA.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Email address
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// Mobile number
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// Validity period. The default value is 12 months, which is the only supported value currently.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Encryption algorithm. RSA and ECC are supported.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// Key pair parameters. RSA supports only 2048. ECC supports only prime256v1. When the encryption algorithm is set to ECC, this parameter is mandatory.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// CSR encryption password
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Original certificate ID, which is used to apply for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Benefit package ID, which is used to expand the free certificate package
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is no by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`


	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`
}

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation.
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate type. Currently, the only supported value is 2, which indicates TrustAsia TLS RSA CA.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Email address
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// Mobile number
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// Validity period. The default value is 12 months, which is the only supported value currently.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Encryption algorithm. RSA and ECC are supported.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// Key pair parameters. RSA supports only 2048. ECC supports only prime256v1. When the encryption algorithm is set to ECC, this parameter is mandatory.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// CSR encryption password
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Original certificate ID, which is used to apply for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Benefit package ID, which is used to expand the free certificate package
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is no by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`

	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DvAuthMethod")
	delete(f, "DomainName")
	delete(f, "ProjectId")
	delete(f, "PackageType")
	delete(f, "ContactEmail")
	delete(f, "ContactPhone")
	delete(f, "ValidityPeriod")
	delete(f, "CsrEncryptAlgo")
	delete(f, "CsrKeyParameter")
	delete(f, "CsrKeyPassword")
	delete(f, "Alias")
	delete(f, "OldCertificateId")
	delete(f, "PackageId")
	delete(f, "DeleteDnsAutoRecord")
	delete(f, "DnsNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCertificateResponseParams `json:"Response"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteCSRRequestParams struct {
	// The IDs of the CSRs to be deleted, 100 IDs at most.
	CSRIds []*int64 `json:"CSRIds,omitnil,omitempty" name:"CSRIds"`
}

type BatchDeleteCSRRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of the CSRs to be deleted, 100 IDs at most.
	CSRIds []*int64 `json:"CSRIds,omitnil,omitempty" name:"CSRIds"`
}

func (r *BatchDeleteCSRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteCSRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CSRIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteCSRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteCSRResponseParams struct {
	// The IDs of the CSRs successfully deleted.
	Success []*int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteCSRResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteCSRResponseParams `json:"Response"`
}

func (r *BatchDeleteCSRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteCSRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindResourceRegionResult struct {
	// The region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The total number of associated cloud resources.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type BindResourceResult struct {
	// Supported types: CLB, CDN, WAF, LIVE, VOD, DDOS, TKE, APIGATEWAY, TCB, and TEO (EDGEONE).
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The region of associated cloud resources.
	BindResourceRegionResult []*BindResourceRegionResult `json:"BindResourceRegionResult,omitnil,omitempty" name:"BindResourceRegionResult"`
}

type CSRItem struct {
	// The CSR ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The account.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// The domain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	// Note: This field may return null, indicating that no valid values can be obtained.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The remarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`

	// The status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The creation time.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The encryption algorithm.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The algorithm parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`
}

// Predefined struct for user
type CancelAuditCertificateRequestParams struct {
	// The certificate ID.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CancelAuditCertificateRequest struct {
	*tchttp.BaseRequest
	
	// The certificate ID.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CancelAuditCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuditCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAuditCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAuditCertificateResponseParams struct {
	// Whether the operation succeeded.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAuditCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CancelAuditCertificateResponseParams `json:"Response"`
}

func (r *CancelAuditCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuditCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CancelCertificateOrderRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CancelCertificateOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelCertificateOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderResponseParams struct {
	// ID of the certificate whose order has been successfully cancelled
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelCertificateOrderResponse struct {
	*tchttp.BaseResponse
	Response *CancelCertificateOrderResponseParams `json:"Response"`
}

func (r *CancelCertificateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdnInstanceDetail struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The ID of the deployed certificate.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The status of the domain.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The billing status of the domain.
	HttpsBillingSwitch *string `json:"HttpsBillingSwitch,omitnil,omitempty" name:"HttpsBillingSwitch"`
}

type CdnInstanceList struct {
	// The total number of CDN domains in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of CDN domains.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*CdnInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`
}

type CertTaskId struct {
	// The certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The async task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type Certificate struct {
	// The certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The list of domains bound to the certificate.
	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`

	// The root certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// The authentication type. Valid values: `UNIDIRECTIONAL` (one-way authentication) and `MUTUAL` (two-way authentication).
	// Note: This field may return null, indicating that no valid values can be obtained.
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`
}

type CertificateExtra struct {
	// Number of domain names which can be associated with the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainNumber *string `json:"DomainNumber,omitnil,omitempty" name:"DomainNumber"`

	// Original certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginCertificateId *string `json:"OriginCertificateId,omitnil,omitempty" name:"OriginCertificateId"`

	// Original ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedBy *string `json:"ReplacedBy,omitnil,omitempty" name:"ReplacedBy"`

	// New ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedFor *string `json:"ReplacedFor,omitnil,omitempty" name:"ReplacedFor"`

	// Certificate ID of the new order
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewOrder *string `json:"RenewOrder,omitnil,omitempty" name:"RenewOrder"`

	// Whether the certificate is a Chinese SM certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SMCert *int64 `json:"SMCert,omitnil,omitempty" name:"SMCert"`

	// Company type
	// Note: This field may return null, indicating that no valid value can be obtained.
	CompanyType *int64 `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`
}

type Certificates struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate source
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// The certificate plan type. Valid values:
	// null: Certificates uploaded by users (no plan type)
	// `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: This field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Primary domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added for domain names of the DNS_AUTO verification type; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate; `14`: Refunded.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// Vulnerability scanning status. `INACTIVE`: not activated; `ACTIVE`: activated
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageTypeName *string `json:"PackageTypeName,omitnil,omitempty" name:"PackageTypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// Whether it can be renewed 
	// Note: This field may return null, indicating that no valid value can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// Project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitnil,omitempty" name:"ProjectInfo"`

	// Associated Tencent Cloud services. Currently, this parameter is unavailable.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BoundResource []*string `json:"BoundResource,omitnil,omitempty" name:"BoundResource"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether the expiration notification was ignored
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsIgnore *bool `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`

	// Whether the certificate is a Chinese SM certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSM *bool `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// Certificate algorithm
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// Encryption algorithm of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// Expiration time of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEndTimes []*string `json:"CAEndTimes,omitnil,omitempty" name:"CAEndTimes"`

	// Generic name of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// Prereview information of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	PreAuditInfo *PreAuditInfo `json:"PreAuditInfo,omitnil,omitempty" name:"PreAuditInfo"`

	// Whether auto-renewal is enabled.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// The hosting status. Valid values: `0` (hosting), `5` (replacing resources), `10` (hosting completed), and `-1` (not hosted). 
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostingStatus *int64 `json:"HostingStatus,omitnil,omitempty" name:"HostingStatus"`

	// The hosting completion time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostingCompleteTime *string `json:"HostingCompleteTime,omitnil,omitempty" name:"HostingCompleteTime"`

	// The hosted new certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostingRenewCertId *string `json:"HostingRenewCertId,omitnil,omitempty" name:"HostingRenewCertId"`

	// Existing renewed certificate ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	HasRenewOrder *string `json:"HasRenewOrder,omitnil,omitempty" name:"HasRenewOrder"`

	// Whether the original certificate is deleted when a certificate is reissued.
	// Note: This field may return null, indicating that no valid value can be obtained.
	ReplaceOriCertIsDelete *bool `json:"ReplaceOriCertIsDelete,omitnil,omitempty" name:"ReplaceOriCertIsDelete"`

	// Whether the certificate is expiring soon. A certificate is considered to be expiring soon when there are 30 days or less left.
	// Note: This field may return null, indicating that no valid value can be obtained.
	IsExpiring *bool `json:"IsExpiring,omitnil,omitempty" name:"IsExpiring"`

	// Validation expiration time for the addition of the DV certificate
	// Note: This field may return null, indicating that no valid value can be obtained.
	DVAuthDeadline *string `json:"DVAuthDeadline,omitnil,omitempty" name:"DVAuthDeadline"`

	// Domain name validation pass time
	// Note: This field may return null, indicating that no valid value can be obtained.
	ValidationPassedTime *string `json:"ValidationPassedTime,omitnil,omitempty" name:"ValidationPassedTime"`

	// Multiple domain names with which the certificate is associated
	// Note: This field may return null, indicating that no valid value can be obtained.
	CertSANs []*string `json:"CertSANs,omitnil,omitempty" name:"CertSANs"`

	// Domain name validation rejection information
	// Note: This field may return null, indicating that no valid value can be obtained.
	AwaitingValidationMsg *string `json:"AwaitingValidationMsg,omitnil,omitempty" name:"AwaitingValidationMsg"`

	// Whether downloading is allowed
	// Note: This field may return null, indicating that no valid value can be obtained.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`


	IsDNSPODResolve *bool `json:"IsDNSPODResolve,omitnil,omitempty" name:"IsDNSPODResolve"`


	IsPackage *bool `json:"IsPackage,omitnil,omitempty" name:"IsPackage"`


	KeyPasswordCustomFlag *bool `json:"KeyPasswordCustomFlag,omitnil,omitempty" name:"KeyPasswordCustomFlag"`


	SupportDownloadType *SupportDownloadType `json:"SupportDownloadType,omitnil,omitempty" name:"SupportDownloadType"`
}

type ClbInstanceDetail struct {
	// The CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// The CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// The list of CLB listeners.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Listeners []*ClbListener `json:"Listeners,omitnil,omitempty" name:"Listeners"`
}

type ClbInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of CLB instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*ClbInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of CLB instances in this region.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ClbListener struct {
	// The listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The listener name.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Whether to enable SNI. Valid values: `1` (enable) and `0` (disable).
	SniSwitch *uint64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// The listener protocol type. Valid values: `HTTPS` and `TCP_SSL`.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The information of the certificate bound to the listener.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// The list of the listener rules.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rules []*ClbListenerRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// The list of non-matching domains.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

type ClbListenerRule struct {
	// The rule ID.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// The domains bound.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Whether the rule matches the domains to be associated with a certificate.
	IsMatch *bool `json:"IsMatch,omitnil,omitempty" name:"IsMatch"`

	// The certificates associated with the rule.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// The list of non-matching domains.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

// Predefined struct for user
type CommitCertificateInformationRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Domain validation method
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Domain validation method
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

func (r *CommitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitCertificateInformationResponseParams struct {
	// TrustAsia order ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CommitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *CommitCertificateInformationResponseParams `json:"Response"`
}

func (r *CommitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCSRRequestParams struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region code that complies with ISO 3166, such as CN and US.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The encryption algorithm. RSA and ECC are supported.	
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The key pair parameter. For RSA, only the 2048-bit and 4096-bit key pairs are supported. For ECC, only prime256v1 is supported.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`

	// Whether to generate the CSR content. Once the CSR content is generated, the CSR record cannot be modified.
	Generate *bool `json:"Generate,omitnil,omitempty" name:"Generate"`

	// The password of the private key.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// The remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateCSRRequest struct {
	*tchttp.BaseRequest
	
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region code that complies with ISO 3166, such as CN and US.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The encryption algorithm. RSA and ECC are supported.	
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The key pair parameter. For RSA, only the 2048-bit and 4096-bit key pairs are supported. For ECC, only prime256v1 is supported.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`

	// Whether to generate the CSR content. Once the CSR content is generated, the CSR record cannot be modified.
	Generate *bool `json:"Generate,omitnil,omitempty" name:"Generate"`

	// The password of the private key.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// The remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateCSRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCSRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Organization")
	delete(f, "Department")
	delete(f, "Email")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "Country")
	delete(f, "EncryptAlgo")
	delete(f, "KeyParameter")
	delete(f, "Generate")
	delete(f, "KeyPassword")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCSRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCSRResponseParams struct {
	// The CSR ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCSRResponse struct {
	*tchttp.BaseResponse
	Response *CreateCSRResponseParams `json:"Response"`
}

func (r *CreateCSRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCSRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateBindResourceSyncTaskRequestParams struct {
	// The list of certificate IDs, 100 IDs at most.
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// Whether to use the cached results. Valid values: `1` (default) for yes and `0` for no. If any task completed within last 30 minutes exists under the current certificate ID, and the cache is used, the query result of the last task completed within 30 minutes will be read.
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
}

type CreateCertificateBindResourceSyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// The list of certificate IDs, 100 IDs at most.
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// Whether to use the cached results. Valid values: `1` (default) for yes and `0` for no. If any task completed within last 30 minutes exists under the current certificate ID, and the cache is used, the query result of the last task completed within 30 minutes will be read.
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
}

func (r *CreateCertificateBindResourceSyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateBindResourceSyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "IsCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateBindResourceSyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateBindResourceSyncTaskResponseParams struct {
	// The IDs of async tasks for querying cloud resources associated with a certificate.
	CertTaskIds []*CertTaskId `json:"CertTaskIds,omitnil,omitempty" name:"CertTaskIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCertificateBindResourceSyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateBindResourceSyncTaskResponseParams `json:"Response"`
}

func (r *CreateCertificateBindResourceSyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateBindResourceSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateRequestParams struct {
	// Certificate product ID. `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// Certificate validity period. Currently, you can only purchase 1-year certificates.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate product ID. `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// Certificate validity period. Currently, you can only purchase 1-year certificates.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DomainNum")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateResponseParams struct {
	// List of certificate IDs
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// List of order IDs
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateResponseParams `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DdosInstanceDetail struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The protocol type.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// The certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The forwarding port.
	VirtualPort *string `json:"VirtualPort,omitnil,omitempty" name:"VirtualPort"`
}

type DdosInstanceList struct {
	// The total number of DDOS instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of DDOS instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*DdosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`
}

// Predefined struct for user
type DeleteCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateResponseParams struct {
	// Deletion result
	DeleteResult *bool `json:"DeleteResult,omitnil,omitempty" name:"DeleteResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCertificateResponseParams `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSRRequestParams struct {
	// The CSR ID.
	CSRId *int64 `json:"CSRId,omitnil,omitempty" name:"CSRId"`
}

type DescribeCSRRequest struct {
	*tchttp.BaseRequest
	
	// The CSR ID.
	CSRId *int64 `json:"CSRId,omitnil,omitempty" name:"CSRId"`
}

func (r *DescribeCSRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CSRId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCSRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSRResponseParams struct {
	// The CSR ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The account.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The key algorithm.
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The algorithm parameter.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`

	// The remarks.
	Remarks *string `json:"Remarks,omitnil,omitempty" name:"Remarks"`

	// The status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The password of the private key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// The creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The CSR content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CSR *string `json:"CSR,omitnil,omitempty" name:"CSR"`

	// The content of the private key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCSRResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCSRResponseParams `json:"Response"`
}

func (r *DescribeCSRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSRSetRequestParams struct {
	// The number of CSRs on each page. The default value is 10, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The pagination offset, starting from 0.	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The domain for CSR filtering
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The encryption algorithm for CSR filtering
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`
}

type DescribeCSRSetRequest struct {
	*tchttp.BaseRequest
	
	// The number of CSRs on each page. The default value is 10, and the maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The pagination offset, starting from 0.	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The domain for CSR filtering
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The encryption algorithm for CSR filtering
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`
}

func (r *DescribeCSRSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSRSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Domain")
	delete(f, "EncryptAlgo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCSRSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCSRSetResponseParams struct {
	// The total number of CSRs.	
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The list of CSRs.
	Set []*CSRItem `json:"Set,omitnil,omitempty" name:"Set"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCSRSetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCSRSetResponseParams `json:"Response"`
}

func (r *DescribeCSRSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCSRSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskDetailRequestParams struct {
	// The task ID, which is required to query the result of associated cloud resources.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The number of cloud resources displayed on each page. The default value is 10, and the maximum value is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The current offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The types of the resources to be queried. If no value is passed in, all types of resources will be queried.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// The regions of the resources to be queried. Only CLB, TKE, WAF, APIGATEWAY, and TCB resources support the query by region.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeCertificateBindResourceTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// The task ID, which is required to query the result of associated cloud resources.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The number of cloud resources displayed on each page. The default value is 10, and the maximum value is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The current offset.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The types of the resources to be queried. If no value is passed in, all types of resources will be queried.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// The regions of the resources to be queried. Only CLB, TKE, WAF, APIGATEWAY, and TCB resources support the query by region.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

func (r *DescribeCertificateBindResourceTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ResourceTypes")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateBindResourceTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskDetailResponseParams struct {
	// The details of associated CLB resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	CLB []*ClbInstanceList `json:"CLB,omitnil,omitempty" name:"CLB"`

	// The details of associated CDN resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	CDN []*CdnInstanceList `json:"CDN,omitnil,omitempty" name:"CDN"`

	// The details of associated WAF resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	WAF []*WafInstanceList `json:"WAF,omitnil,omitempty" name:"WAF"`

	// The details of associated Anti-DDS resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	DDOS []*DdosInstanceList `json:"DDOS,omitnil,omitempty" name:"DDOS"`

	// The details of associated LIVE resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	LIVE []*LiveInstanceList `json:"LIVE,omitnil,omitempty" name:"LIVE"`

	// The details of associated VOD resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	VOD []*VODInstanceList `json:"VOD,omitnil,omitempty" name:"VOD"`

	// The details of associated TKE resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	TKE []*TkeInstanceList `json:"TKE,omitnil,omitempty" name:"TKE"`

	// The details of associated APIGATEWAY resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	APIGATEWAY []*ApiGatewayInstanceList `json:"APIGATEWAY,omitnil,omitempty" name:"APIGATEWAY"`

	// The details of associated TCB resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	TCB []*TCBInstanceList `json:"TCB,omitnil,omitempty" name:"TCB"`

	// The details of associated TEO resources.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	TEO []*TeoInstanceList `json:"TEO,omitnil,omitempty" name:"TEO"`

	// The status of the async task. Valid values: `0` for querying, `1` for successful, and `2` for abnormal. If the status is `1`, the result of `BindResourceResult` will be displayed; if the status is `2`, the error causes will be displayed.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The cache time of the current result.
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// Associated TSE resource details
	// Note: This field may return null, indicating that no valid value can be obtained.
	TSE []*TSEInstanceList `json:"TSE,omitnil,omitempty" name:"TSE"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateBindResourceTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateBindResourceTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeCertificateBindResourceTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskResultRequestParams struct {
	// The task IDs, which are used to query the results of associated cloud resources, 100 IDs at most.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeCertificateBindResourceTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// The task IDs, which are used to query the results of associated cloud resources, 100 IDs at most.
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeCertificateBindResourceTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateBindResourceTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskResultResponseParams struct {
	// The results of the async tasks for querying associated cloud resources.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SyncTaskBindResourceResult []*SyncTaskBindResourceResult `json:"SyncTaskBindResourceResult,omitnil,omitempty" name:"SyncTaskBindResourceResult"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateBindResourceTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateBindResourceTaskResultResponseParams `json:"Response"`
}

func (r *DescribeCertificateBindResourceTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailResponseParams struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Certificate plan type. null: User-uploaded certificate (no plan type); `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// Private key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Public key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil,omitempty" name:"DvAuthDetail"`

	// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil,omitempty" name:"VulnerabilityReport"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// Multiple domain names included in the certificate (excluding the primary domain name, which uses the `Domain` field)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// Whether the certificate is a paid one.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil,omitempty" name:"SubmittedData"`

	// Whether the certificate can be renewed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// List of associated tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Root certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootCert *RootCertificates `json:"RootCert,omitnil,omitempty" name:"RootCert"`

	// Chinese SM encryption certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptCert *string `json:"EncryptCert,omitnil,omitempty" name:"EncryptCert"`

	// Private key of Chinese SM encryption
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitnil,omitempty" name:"EncryptPrivateKey"`

	// SHA1 fingerprint of the signature certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertFingerprint *string `json:"CertFingerprint,omitnil,omitempty" name:"CertFingerprint"`

	// SHA1 fingerprint of the encryption certificate (for Chinese SM certificates only)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptCertFingerprint *string `json:"EncryptCertFingerprint,omitnil,omitempty" name:"EncryptCertFingerprint"`

	// Certificate algorithm
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// The authentication value for DV certificate revocation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitnil,omitempty" name:"DvRevokeAuthDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateDetailResponseParams `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsRequestParams struct {
	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of requested logs. The default value is 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start time. The default value is 15 days ago.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. The default value is the current time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeCertificateOperateLogsRequest struct {
	*tchttp.BaseRequest
	
	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of requested logs. The default value is 20.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start time. The default value is 15 days ago.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time. The default value is the current time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeCertificateOperateLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateOperateLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsResponseParams struct {
	// Total number of logs that meet query conditions
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// Number of logs returned for this request
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Certificate operation log list
	// Note: this field may return null, indicating that no valid values can be obtained.
	OperateLogs []*OperationLog `json:"OperateLogs,omitnil,omitempty" name:"OperateLogs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateOperateLogsResponseParams `json:"Response"`
}

func (r *DescribeCertificateOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DescribeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateResponseParams struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Certificate plan type. `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Name of the certificate issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil,omitempty" name:"DvAuthDetail"`

	// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil,omitempty" name:"VulnerabilityReport"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageTypeName *string `json:"PackageTypeName,omitnil,omitempty" name:"PackageTypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil,omitempty" name:"SubmittedData"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// All encryption algorithms of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// All common names of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// All expiration time of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEndTimes []*string `json:"CAEndTimes,omitnil,omitempty" name:"CAEndTimes"`

	// The authentication value for DV certificate revocation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitnil,omitempty" name:"DvRevokeAuthDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateResponseParams `json:"Response"`
}

func (r *DescribeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesRequestParams struct {
	// Pagination offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword for search, which can be a certificate ID, alias, or domain name, for example, a8xHcaIs
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Sorting by expiration time. `DESC`: descending; `ASC`: ascending
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// Certificate status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate.
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil,omitempty" name:"CertificateStatus"`

	// Whether the certificate can be deployed. `1`: yes; `0`: no
	Deployable *uint64 `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// Whether to filter uploaded hosted certificates. `1`: Yes; `0`: No.
	Upload *int64 `json:"Upload,omitnil,omitempty" name:"Upload"`

	// Whether to filter renewable certificates. `1`: Yes; `0`: No.
	Renew *int64 `json:"Renew,omitnil,omitempty" name:"Renew"`

	// Filter by source. `upload`: Uploaded certificate; `buy`: Tencent Cloud certificate. If this parameter is left empty, all certificates will be queried.
	FilterSource *string `json:"FilterSource,omitnil,omitempty" name:"FilterSource"`

	// Whether to filter Chinese SM certificates. `1`: Yes; `0`: No.
	IsSM *int64 `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// Whether to filter expiring certificates. `1`: Yes; `0`: No.
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil,omitempty" name:"FilterExpiring"`

	// Whether the certificate can be hosted. Valid values: `1` for yes and `0` for no.
	Hostable *uint64 `json:"Hostable,omitnil,omitempty" name:"Hostable"`
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Keyword for search, which can be a certificate ID, alias, or domain name, for example, a8xHcaIs
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Sorting by expiration time. `DESC`: descending; `ASC`: ascending
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// Certificate status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate.
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil,omitempty" name:"CertificateStatus"`

	// Whether the certificate can be deployed. `1`: yes; `0`: no
	Deployable *uint64 `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// Whether to filter uploaded hosted certificates. `1`: Yes; `0`: No.
	Upload *int64 `json:"Upload,omitnil,omitempty" name:"Upload"`

	// Whether to filter renewable certificates. `1`: Yes; `0`: No.
	Renew *int64 `json:"Renew,omitnil,omitempty" name:"Renew"`

	// Filter by source. `upload`: Uploaded certificate; `buy`: Tencent Cloud certificate. If this parameter is left empty, all certificates will be queried.
	FilterSource *string `json:"FilterSource,omitnil,omitempty" name:"FilterSource"`

	// Whether to filter Chinese SM certificates. `1`: Yes; `0`: No.
	IsSM *int64 `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// Whether to filter expiring certificates. `1`: Yes; `0`: No.
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil,omitempty" name:"FilterExpiring"`

	// Whether the certificate can be hosted. Valid values: `1` for yes and `0` for no.
	Hostable *uint64 `json:"Hostable,omitnil,omitempty" name:"Hostable"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "CertificateType")
	delete(f, "ProjectId")
	delete(f, "ExpirationSort")
	delete(f, "CertificateStatus")
	delete(f, "Deployable")
	delete(f, "Upload")
	delete(f, "Renew")
	delete(f, "FilterSource")
	delete(f, "IsSM")
	delete(f, "FilterExpiring")
	delete(f, "Hostable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesResponseParams struct {
	// Total number
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificates []*Certificates `json:"Certificates,omitnil,omitempty" name:"Certificates"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificatesResponseParams `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTeoInstanceListRequestParams struct {
	// The ID of the certificate to be deployed.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The type of resource for certificate deployment.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Whether to query the cached results. Valid values: `1` (yes) and `0` (no). By default, the cached results within 30 minutes are queried.
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// The list of filter parameters. FilterKey: domainMatch (query the list of instances with matching or non-matching domains). FilterValue: `1` (default; query the list of instances with matching domains) or `0` (query the list of instances with non-matching domains).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The ID of the deployed certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// The pagination offset, starting from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of instances on each page. Default value: 10.	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Whether the query is asynchronous.
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
}

type DescribeHostTeoInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the certificate to be deployed.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The type of resource for certificate deployment.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Whether to query the cached results. Valid values: `1` (yes) and `0` (no). By default, the cached results within 30 minutes are queried.
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// The list of filter parameters. FilterKey: domainMatch (query the list of instances with matching or non-matching domains). FilterValue: `1` (default; query the list of instances with matching domains) or `0` (query the list of instances with non-matching domains).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The ID of the deployed certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// The pagination offset, starting from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of instances on each page. Default value: 10.	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Whether the query is asynchronous.
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
}

func (r *DescribeHostTeoInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTeoInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AsyncCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostTeoInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTeoInstanceListResponseParams struct {
	// The list of EDGEONE instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total count.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostTeoInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostTeoInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostTeoInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTeoInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordDetailRequestParams struct {
	// One-click update record ID
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number per page, 10 by default.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging offset, starting from 0
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeHostUpdateRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// One-click update record ID
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number per page, 10 by default.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Paging offset, starting from 0
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeHostUpdateRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUpdateRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordDetailResponseParams struct {
	// Total count
	// Note: This field may return null, indicating that no valid value can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Certificate deployment record list
	// Note: This field may return null, indicating that no valid value can be obtained.
	RecordDetailList []*UpdateRecordDetails `json:"RecordDetailList,omitnil,omitempty" name:"RecordDetailList"`

	// Total successful deployments
	// Note: This field may return null, indicating that no valid value can be obtained.
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitnil,omitempty" name:"SuccessTotalCount"`

	// Total failed deployments
	// Note: This field may return null, indicating that no valid value can be obtained.
	FailedTotalCount *int64 `json:"FailedTotalCount,omitnil,omitempty" name:"FailedTotalCount"`

	// Total running deployments
	// Note: This field may return null, indicating that no valid value can be obtained.
	RunningTotalCount *int64 `json:"RunningTotalCount,omitnil,omitempty" name:"RunningTotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostUpdateRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUpdateRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeHostUpdateRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordRequestParams struct {
	// Paging offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number per page, 10 by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// New certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Old certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostUpdateRecordRequest struct {
	*tchttp.BaseRequest
	
	// Paging offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number per page, 10 by default.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// New certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Old certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostUpdateRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CertificateId")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUpdateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordResponseParams struct {
	// Total count
	// Note: This field may return null, indicating that no valid value can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Certificate deployment record list
	// Note: This field may return null, indicating that no valid value can be obtained.
	DeployRecordList []*UpdateRecordInfo `json:"DeployRecordList,omitnil,omitempty" name:"DeployRecordList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostUpdateRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUpdateRecordResponseParams `json:"Response"`
}

func (r *DescribeHostUpdateRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DownloadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *DownloadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateResponseParams struct {
	// ZIP content encoded by using Base64. After the content is decoded by using Base64, it can be saved as a ZIP file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// MIME type. `application/zip`: ZIP file
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DownloadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCertificateResponseParams `json:"Response"`
}

func (r *DownloadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DvAuthDetail struct {
	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitnil,omitempty" name:"DvAuthKeySubDomain"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuths []*DvAuths `json:"DvAuths,omitnil,omitempty" name:"DvAuths"`
}

type DvAuths struct {
	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitnil,omitempty" name:"DvAuthSubDomain"`

	// DV authentication type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthVerifyType *string `json:"DvAuthVerifyType,omitnil,omitempty" name:"DvAuthVerifyType"`
}

type Error struct {
	// The error code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// The error message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type Filter struct {
	// The key of the filter parameter.
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// The value of the filter parameter.
	FilterValue *string `json:"FilterValue,omitnil,omitempty" name:"FilterValue"`
}

type GatewayCertificate struct {
	// Gateway certificate ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Gateway certificate information
	// Note: This field may return null, indicating that no valid value can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Bound domain name
	// Note: This field may return null, indicating that no valid value can be obtained.
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// Certificate source
	// Note: This field may return null, indicating that no valid value can be obtained.
	CertSource *string `json:"CertSource,omitnil,omitempty" name:"CertSource"`

	// SSL certificate ID that is currently bound
	// Note: This field may return null, indicating that no valid value can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type LiveInstanceDetail struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The ID of the associated certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The status. Valid values: `-1`: No certificate is associated with the domain.
	// `1`: HTTPS is enabled for the domain.
	// `0`: HTTPS is disabled for the domain.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type LiveInstanceList struct {
	// The total number of LIVE instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The list of LIVE instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`
}

// Predefined struct for user
type ModifyCSRRequestParams struct {
	// The CSR ID.	
	CSRId *int64 `json:"CSRId,omitnil,omitempty" name:"CSRId"`

	// Whether to generate the CSR content. Once the CSR content is generated, the CSR record cannot be modified.
	Generate *bool `json:"Generate,omitnil,omitempty" name:"Generate"`

	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The encryption algorithm. RSA and ECC are supported.	
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The key pair parameter. For RSA, only the 2048-bit and 4096-bit key pairs are supported. For ECC, only prime256v1 is supported.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`

	// The remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// The password of the private key.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`
}

type ModifyCSRRequest struct {
	*tchttp.BaseRequest
	
	// The CSR ID.	
	CSRId *int64 `json:"CSRId,omitnil,omitempty" name:"CSRId"`

	// Whether to generate the CSR content. Once the CSR content is generated, the CSR record cannot be modified.
	Generate *bool `json:"Generate,omitnil,omitempty" name:"Generate"`

	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The organization name.
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// The department.
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// The email address.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// The province.
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// The city.
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// The country or region.
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// The encryption algorithm. RSA and ECC are supported.	
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`

	// The key pair parameter. For RSA, only the 2048-bit and 4096-bit key pairs are supported. For ECC, only prime256v1 is supported.
	KeyParameter *string `json:"KeyParameter,omitnil,omitempty" name:"KeyParameter"`

	// The remarks.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// The password of the private key.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`
}

func (r *ModifyCSRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCSRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CSRId")
	delete(f, "Generate")
	delete(f, "Domain")
	delete(f, "Organization")
	delete(f, "Department")
	delete(f, "Email")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "Country")
	delete(f, "EncryptAlgo")
	delete(f, "KeyParameter")
	delete(f, "Remark")
	delete(f, "KeyPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCSRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCSRResponseParams struct {
	// The CSR ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCSRResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCSRResponseParams `json:"Response"`
}

func (r *ModifyCSRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCSRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateAliasRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type ModifyCertificateAliasRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

func (r *ModifyCertificateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateAliasResponseParams struct {
	// ID of the successfully modified certificate
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCertificateAliasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateAliasResponseParams `json:"Response"`
}

func (r *ModifyCertificateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectRequestParams struct {
	// ID list of certificates whose projects need to be modified. A maximum of 100 certificate IDs are supported.
	CertificateIdList []*string `json:"CertificateIdList,omitnil,omitempty" name:"CertificateIdList"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest
	
	// ID list of certificates whose projects need to be modified. A maximum of 100 certificate IDs are supported.
	CertificateIdList []*string `json:"CertificateIdList,omitnil,omitempty" name:"CertificateIdList"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyCertificateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectResponseParams struct {
	// List of certificates whose projects were modified successfully
	// Note: this field may return null, indicating that no valid values can be obtained.
	SuccessCertificates []*string `json:"SuccessCertificates,omitnil,omitempty" name:"SuccessCertificates"`

	// List of certificates whose projects failed to be modified
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailCertificates []*string `json:"FailCertificates,omitnil,omitempty" name:"FailCertificates"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCertificateProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateProjectResponseParams `json:"Response"`
}

func (r *ModifyCertificateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateResubmitRequestParams struct {
	// The certificate ID.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type ModifyCertificateResubmitRequest struct {
	*tchttp.BaseRequest
	
	// The certificate ID.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *ModifyCertificateResubmitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResubmitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateResubmitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateResubmitResponseParams struct {
	// The certificate ID.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCertificateResubmitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateResubmitResponseParams `json:"Response"`
}

func (r *ModifyCertificateResubmitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResubmitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {
	// Action performed on logs
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Time when the action is performed
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`
}

type PreAuditInfo struct {
	// Total number of years of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPeriod *int64 `json:"TotalPeriod,omitnil,omitempty" name:"TotalPeriod"`

	// Current year of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	NowPeriod *int64 `json:"NowPeriod,omitnil,omitempty" name:"NowPeriod"`

	// Certificate prereview manager ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
}

type ProjectInfo struct {
	// Project name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// UIN of the project creator
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreatorUin *uint64 `json:"ProjectCreatorUin,omitnil,omitempty" name:"ProjectCreatorUin"`

	// Project creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreateTime *string `json:"ProjectCreateTime,omitnil,omitempty" name:"ProjectCreateTime"`

	// Brief project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectResume *string `json:"ProjectResume,omitnil,omitempty" name:"ProjectResume"`

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

// Predefined struct for user
type ReplaceCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	ValidType *string `json:"ValidType,omitnil,omitempty" name:"ValidType"`

	// Type. `original`: original certificate CSR; `upload`: uploaded manually; `online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR content
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// Password of the key
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil,omitempty" name:"CsrkeyPassword"`

	// Reissue reason
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// The CSR encryption algorithm. Valid values: `RSA` (default), `ECC1`, and `SM2`.
	// This parameter is available for selection only when the value of `CsrType` is `Online`.
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitnil,omitempty" name:"CertCSREncryptAlgo"`

	// The CSR encryption parameters. When `CsrEncryptAlgo` is set to `RSA`, `2048` (default) and `4096` are available for selection; and when`CsrEncryptAlgo` is set to `ECC`, `prime256v1` (default) and `secp384r1` are available for selection. 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitnil,omitempty" name:"CertCSRKeyParameter"`
}

type ReplaceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	ValidType *string `json:"ValidType,omitnil,omitempty" name:"ValidType"`

	// Type. `original`: original certificate CSR; `upload`: uploaded manually; `online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR content
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// Password of the key
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil,omitempty" name:"CsrkeyPassword"`

	// Reissue reason
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// The CSR encryption algorithm. Valid values: `RSA` (default), `ECC1`, and `SM2`.
	// This parameter is available for selection only when the value of `CsrType` is `Online`.
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitnil,omitempty" name:"CertCSREncryptAlgo"`

	// The CSR encryption parameters. When `CsrEncryptAlgo` is set to `RSA`, `2048` (default) and `4096` are available for selection; and when`CsrEncryptAlgo` is set to `ECC`, `prime256v1` (default) and `secp384r1` are available for selection. 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitnil,omitempty" name:"CertCSRKeyParameter"`
}

func (r *ReplaceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ValidType")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CsrkeyPassword")
	delete(f, "Reason")
	delete(f, "CertCSREncryptAlgo")
	delete(f, "CertCSRKeyParameter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertificateResponseParams `json:"Response"`
}

func (r *ReplaceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTypeRegions struct {
	// Cloud resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Region list
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type RootCertificates struct {
	// Chinese SM signature certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sign *string `json:"Sign,omitnil,omitempty" name:"Sign"`

	// Chinese SM encryption certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *string `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// Standard certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`
}

// Predefined struct for user
type SubmitCertificateInformationRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// CSR generation mode. `online`: generated online; `parse`: uploaded manually
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// Uploaded CSR content
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// Domain name bound with the certificate
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// Uploaded domain name array (can be uploaded for a multi-domain certificate)
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// Password of the private key
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// Organization name
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Division name
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// Detailed address of the organization
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// Country where the organization is located, for example, CN (China)
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// City where the organization is located
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// Province where the organization is located
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// Area code of the fixed-line phone number of the organization
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// Fixed-line phone number of the organization
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Certificate validation method
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Last name of the administrator
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// First name of the administrator
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// Mobile number of the administrator
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// Email of the administrator
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// Position of the administrator
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// Last name of the contact
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// First name of the contact
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// Email of the contact
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// Mobile number of the contact
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// Position of the contact
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// CSR generation mode. `online`: generated online; `parse`: uploaded manually
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// Uploaded CSR content
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// Domain name bound with the certificate
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// Uploaded domain name array (can be uploaded for a multi-domain certificate)
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// Password of the private key
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// Organization name
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Division name
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// Detailed address of the organization
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// Country where the organization is located, for example, CN (China)
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// City where the organization is located
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// Province where the organization is located
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// Area code of the fixed-line phone number of the organization
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// Fixed-line phone number of the organization
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Certificate validation method
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Last name of the administrator
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// First name of the administrator
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// Mobile number of the administrator
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// Email of the administrator
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// Position of the administrator
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// Last name of the contact
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// First name of the contact
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// Email of the contact
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// Mobile number of the contact
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// Position of the contact
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`
}

func (r *SubmitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CertificateDomain")
	delete(f, "DomainList")
	delete(f, "KeyPassword")
	delete(f, "OrganizationName")
	delete(f, "OrganizationDivision")
	delete(f, "OrganizationAddress")
	delete(f, "OrganizationCountry")
	delete(f, "OrganizationCity")
	delete(f, "OrganizationRegion")
	delete(f, "PostalCode")
	delete(f, "PhoneAreaCode")
	delete(f, "PhoneNumber")
	delete(f, "VerifyType")
	delete(f, "AdminFirstName")
	delete(f, "AdminLastName")
	delete(f, "AdminPhoneNum")
	delete(f, "AdminEmail")
	delete(f, "AdminPosition")
	delete(f, "ContactFirstName")
	delete(f, "ContactLastName")
	delete(f, "ContactEmail")
	delete(f, "ContactNumber")
	delete(f, "ContactPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCertificateInformationResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *SubmitCertificateInformationResponseParams `json:"Response"`
}

func (r *SubmitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmittedData struct {
	// CSR type. `online`: CSR generated online; `parse`: CSR pasted
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR content
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// Domain name information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// DNS information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// Password of the private key
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// Enterprise or unit name
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// Division
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// Address
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// Country
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// City
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// Province
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// Postal code
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// Area code of the fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// Fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// First name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// Last name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// Phone number of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// Email of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// Position of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// First name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// Last name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// Phone number of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// Email of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// Position of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`

	// Validation type
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type SupportDownloadType struct {

	NGINX *bool `json:"NGINX,omitnil,omitempty" name:"NGINX"`


	APACHE *bool `json:"APACHE,omitnil,omitempty" name:"APACHE"`


	TOMCAT *bool `json:"TOMCAT,omitnil,omitempty" name:"TOMCAT"`


	IIS *bool `json:"IIS,omitnil,omitempty" name:"IIS"`


	JKS *bool `json:"JKS,omitnil,omitempty" name:"JKS"`


	OTHER *bool `json:"OTHER,omitnil,omitempty" name:"OTHER"`


	ROOT *bool `json:"ROOT,omitnil,omitempty" name:"ROOT"`
}

type SyncTaskBindResourceResult struct {
	// The task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The associated cloud resources.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BindResourceResult []*BindResourceResult `json:"BindResourceResult,omitnil,omitempty" name:"BindResourceResult"`

	// The status of the async task. Valid values: `0` for querying, `1` for successful, and `2` for abnormal. If the status is `1`, the result of `BindResourceResult` will be displayed; if the status is `2`, the error causes will be displayed.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The error occurred when querying the associated cloud resources.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Error *Error `json:"Error,omitnil,omitempty" name:"Error"`

	// The cache time of the current result.
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type TCBAccessInstance struct {
	// The domain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unified domain status.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnionStatus *int64 `json:"UnionStatus,omitnil,omitempty" name:"UnionStatus"`

	// Whether the domain is preempted. A preempted domain is one that is already associated with another environment. It must be disassociated or re-associated first.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsPreempted *bool `json:"IsPreempted,omitnil,omitempty" name:"IsPreempted"`

	// Whether the domain is added to the ICP blocklist. Valid values: `0` for no and `1` for yes.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ICPStatus *int64 `json:"ICPStatus,omitnil,omitempty" name:"ICPStatus"`

	// The ID of the associated certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type TCBAccessService struct {
	// The list of instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TCBAccessInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The instance count.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TCBEnvironment struct {
	// The unique ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// The source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// The name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TCBEnvironments struct {
	// The TCB environment.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	Environment *TCBEnvironment `json:"Environment,omitnil,omitempty" name:"Environment"`

	// The access service.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessService *TCBAccessService `json:"AccessService,omitnil,omitempty" name:"AccessService"`

	// Whether static hosting is used.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostService *TCBHostService `json:"HostService,omitnil,omitempty" name:"HostService"`
}

type TCBHostInstance struct {
	// The domain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The resolution status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DNSStatus *string `json:"DNSStatus,omitnil,omitempty" name:"DNSStatus"`

	// The ID of the associated certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type TCBHostService struct {
	// The list of instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TCBHostInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The instance count.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TCBInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of TCB environments.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Environments []*TCBEnvironments `json:"Environments,omitnil,omitempty" name:"Environments"`
}

type TSEInstanceDetail struct {
	// Gateway ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// Gateway name
	// Note: This field may return null, indicating that no valid value can be obtained.
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// Gateway certificate list
	// Note: This field may return null, indicating that no valid value can be obtained.
	CertificateList []*GatewayCertificate `json:"CertificateList,omitnil,omitempty" name:"CertificateList"`
}

type TSEInstanceList struct {
	// TSE instance details
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceList []*TSEInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Total TSE instances in this region	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Region	
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type Tags struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TeoInstanceDetail struct {
	// The domain.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The AZ ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// The status of the domain.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TeoInstanceList struct {
	// The list of EDGEONE instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of EDGEONE instances.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TkeIngressDetail struct {
	// The Ingress name.
	IngressName *string `json:"IngressName,omitnil,omitempty" name:"IngressName"`

	// The list of TLS domains.
	TlsDomains []*string `json:"TlsDomains,omitnil,omitempty" name:"TlsDomains"`

	// The list of Ingress domains.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type TkeInstanceDetail struct {
	// The cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// The cluster name.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// The list of cluster namespaces.
	NamespaceList []*TkeNameSpaceDetail `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// The cluster type.
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// The cluster version.
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`
}

type TkeInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of TKE instances.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TkeInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of TKE instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TkeNameSpaceDetail struct {
	// The namespace name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The secret list.
	SecretList []*TkeSecretDetail `json:"SecretList,omitnil,omitempty" name:"SecretList"`
}

type TkeSecretDetail struct {
	// The secret name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The Ingress list.
	IngressList []*TkeIngressDetail `json:"IngressList,omitnil,omitempty" name:"IngressList"`

	// The list of domains that do not match the new certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

// Predefined struct for user
type UpdateCertificateInstanceRequestParams struct {
	// One-click update old certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Type of the resource that needs to be deployed. The following parameter values are optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, and tcb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// One-click update new certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// List of regions that need to be deployed (deprecated)
	//
	// Deprecated: Regions is deprecated.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// List of regions for which cloud resources need to be deployed
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// Public key of the certificate. If the public key of the certificate is uploaded, CertificateId does not need to be uploaded.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key of the certificate. If the public key of the certificate is uploaded, the private key of the certificate is required.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Whether an expiration reminder is ignored for the old certificate. 0: The notification is not ignored. 1: The notification is ignored.
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// Whether repeated uploading of the same certificate is allowed. If the public key of the certificate is uploaded, this parameter can be configured.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// Whether downloading is allowed. If the public key of the certificate is uploaded, this parameter can be configured.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// Tag list. If the public key of the certificate is uploaded, this parameter can be configured.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Project ID. If the public key of the certificate is uploaded, this parameter can be configured.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UpdateCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// One-click update old certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Type of the resource that needs to be deployed. The following parameter values are optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, and tcb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// One-click update new certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// List of regions that need to be deployed (deprecated)
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// List of regions for which cloud resources need to be deployed
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// Public key of the certificate. If the public key of the certificate is uploaded, CertificateId does not need to be uploaded.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key of the certificate. If the public key of the certificate is uploaded, the private key of the certificate is required.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Whether an expiration reminder is ignored for the old certificate. 0: The notification is not ignored. 1: The notification is ignored.
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// Whether repeated uploading of the same certificate is allowed. If the public key of the certificate is uploaded, this parameter can be configured.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// Whether downloading is allowed. If the public key of the certificate is uploaded, this parameter can be configured.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// Tag list. If the public key of the certificate is uploaded, this parameter can be configured.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Project ID. If the public key of the certificate is uploaded, this parameter can be configured.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *UpdateCertificateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldCertificateId")
	delete(f, "ResourceTypes")
	delete(f, "CertificateId")
	delete(f, "Regions")
	delete(f, "ResourceTypesRegions")
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "ExpiringNotificationSwitch")
	delete(f, "Repeatable")
	delete(f, "AllowDownload")
	delete(f, "Tags")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateInstanceResponseParams struct {
	// Cloud resource deployment task ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	DeployRecordId *uint64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Deployment status. 1 indicates that the deployment succeeded, and 0 indicates that the deployment failed.
	DeployStatus *int64 `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`


	UpdateSyncProgress []*UpdateSyncProgress `json:"UpdateSyncProgress,omitnil,omitempty" name:"UpdateSyncProgress"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCertificateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateInstanceResponseParams `json:"Response"`
}

func (r *UpdateCertificateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRetryRequestParams struct {
	// To-be-redeployed record ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// To-be-redeployed record detail ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

type UpdateCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// To-be-redeployed record ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// To-be-redeployed record detail ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

func (r *UpdateCertificateRecordRetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "DeployRecordDetailId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateRecordRetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRetryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCertificateRecordRetryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateRecordRetryResponseParams `json:"Response"`
}

func (r *UpdateCertificateRecordRetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRollbackRequestParams struct {
	// To-be-redeployed record ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

type UpdateCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// To-be-redeployed record ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

func (r *UpdateCertificateRecordRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateRecordRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRollbackResponseParams struct {
	// Rollback deployment record ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCertificateRecordRollbackResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateRecordRollbackResponseParams `json:"Response"`
}

func (r *UpdateCertificateRecordRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordDetail struct {
	// Detail record ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// New certificate ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Old certificate ID
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// Deployment domain name list
	// Note: This field may return null, indicating that no valid value can be obtained.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Deployment resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Deployment region
	// Note: This field may return null, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Deployment status
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Deployment error message
	// Note: This field may return null, indicating that no valid value can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Deployment time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Deployment instance ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Deployment instance name
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Deployment listener ID (only for CLB)
	// Note: This field may return null, indicating that no valid value can be obtained.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Deployment listener name (only for CLB)
	// Note: This field may return null, indicating that no valid value can be obtained.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Protocol
	// Note: This field may return null, indicating that no valid value can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Whether SNI is enabled (only for CLB)
	// Note: This field may return null, indicating that no valid value can be obtained.
	SniSwitch *uint64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// Bucket name (only for COS)
	// Note: This field may return null, indicating that no valid value can be obtained.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Port
	// Note: This field may return null, indicating that no valid value can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Namespace (only for TKE)
	// Note: This field may return null, indicating that no valid value can be obtained.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Secret name (only for TKE)
	// Note: This field may return null, indicating that no valid value can be obtained.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Environment ID
	// Note: This field may return null, indicating that no valid value can be obtained.
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// TCB deployment type
	// Note: This field may return null, indicating that no valid value can be obtained.
	TCBType *string `json:"TCBType,omitnil,omitempty" name:"TCBType"`
}

type UpdateRecordDetails struct {
	// Deployment resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Deployment resource detail list
	List []*UpdateRecordDetail `json:"List,omitnil,omitempty" name:"List"`

	// Total deployment resource count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type UpdateRecordInfo struct {
	// Record ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// New certificate ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Old certificate ID
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// Deployment resource type list
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Deployment region list
	// Note: This field may return null, indicating that no valid value can be obtained.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// Deployment status
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Deployment time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type UpdateSyncProgress struct {

	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`


	UpdateSyncProgressRegions []*UpdateSyncProgressRegion `json:"UpdateSyncProgressRegions,omitnil,omitempty" name:"UpdateSyncProgressRegions"`


	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateSyncProgressRegion struct {

	Region *string `json:"Region,omitnil,omitempty" name:"Region"`


	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`


	OffsetCount *int64 `json:"OffsetCount,omitnil,omitempty" name:"OffsetCount"`


	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UploadCertificateRequestParams struct {
	// Public key of the certificate
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key content. This parameter is required when the certificate type is SVR, and not required when the certificate type is CA.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Certificate type. Valid values: `CA` (CA certificate) and `SVR` (server certificate). Default value: `SVR`
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`


	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// The list of tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether a certificate can be repeatedly uploaded.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`
}

type UploadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Public key of the certificate
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key content. This parameter is required when the certificate type is SVR, and not required when the certificate type is CA.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Certificate type. Valid values: `CA` (CA certificate) and `SVR` (server certificate). Default value: `SVR`
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// The list of tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether a certificate can be repeatedly uploaded.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`
}

func (r *UploadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "CertificateType")
	delete(f, "Alias")
	delete(f, "ProjectId")
	delete(f, "CertificateUse")
	delete(f, "Tags")
	delete(f, "Repeatable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// The ID of the repeatedly uploaded certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RepeatCertId *string `json:"RepeatCertId,omitnil,omitempty" name:"RepeatCertId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *UploadCertificateResponseParams `json:"Response"`
}

func (r *UploadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Base64-encoded confirmation letter file, which must be a JPG, JPEG, PNG, or PDF file of 1 KB to 1.4 MB
	ConfirmLetter *string `json:"ConfirmLetter,omitnil,omitempty" name:"ConfirmLetter"`
}

type UploadConfirmLetterRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Base64-encoded confirmation letter file, which must be a JPG, JPEG, PNG, or PDF file of 1 KB to 1.4 MB
	ConfirmLetter *string `json:"ConfirmLetter,omitnil,omitempty" name:"ConfirmLetter"`
}

func (r *UploadConfirmLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ConfirmLetter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadConfirmLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Whether the operation is successful
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadConfirmLetterResponse struct {
	*tchttp.BaseResponse
	Response *UploadConfirmLetterResponseParams `json:"Response"`
}

func (r *UploadConfirmLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VODInstanceList struct {
	// The list of VOD instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*VodInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of VOD instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type VodInstanceDetail struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The certificate ID.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type WafInstanceDetail struct {
	// The domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// The certificate ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Whether to keep the persistent connection.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Keepalive *uint64 `json:"Keepalive,omitnil,omitempty" name:"Keepalive"`
}

type WafInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The list of WAF instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*WafInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of WAF instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}