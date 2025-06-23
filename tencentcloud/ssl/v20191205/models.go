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

	// Certificate id.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type ApiGatewayInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// API gateway instance details.	
	InstanceList []*ApiGatewayInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of APIGATEWAY instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Whether to query exceptions.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type ApplyCertificateRequestParams struct {
	// Certificate domain validation methods:
	// 
	// DNS_AUTO: Automatically add domain DNS validation. Requires the user's domain to be hosted on 'Cloud DNS' and associated with the same Tencent Cloud account as the certificate application.
	// 
	// DNS: Manually add domain DNS validation. Requires the user to manually add the validation value at their domain's DNS service provider.
	// 
	// FILE: Manually add domain file validation. Requires the user to manually add a specified path file in the root directory of the domain site for file validation. Either HTTP or HTTPS validation will suffice; the domain site must be accessible by overseas CA institutions. The specific access whitelist is: 64.78.193.238, 216.168.247.9, 216.168.249.9, 54.189.196.217.
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// The domain bound to the certificate.
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// The project ID associated with the certificate. Default is 0 (default project)
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate type, optional, currently only type 83 is supported. 83 = trustasia c1 dv free.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// The email associated with the certificate order, By default, it uses the Tencent Cloud account email. If it does not exist, a fixed email address will be used.
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// The mobile phone number associated with the certificate. If it does not exist, a fixed mobile phone number will be used.
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// Certificate valid period, 3 months by default, currently only 3 months is supported.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Encryption algorithm, values can be ECC or RSA, default is RSA.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// Key pair parameters. RSA supports only 2048. ECC supports only prime256v1. When the encryption algorithm is set to ECC, this parameter is mandatory.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// Private key password, currently only used when generating jks, pfx format certificates; private key certificates of other formats are not encrypted.
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// Certificate alias.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Old certificate id, used for certificate renewal (the certificate valid period is within 30 days and not expired), a renewal relationship will be established, which can be hosted; not providing it means applying for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Benefit package ID, used for free certificate expansion package, the free certificate expansion package has been discontinued.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is fasle by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`

	// Other domains bound to the certificate, to be opened. This parameter is not currently supported.
	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`
}

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate domain validation methods:
	// 
	// DNS_AUTO: Automatically add domain DNS validation. Requires the user's domain to be hosted on 'Cloud DNS' and associated with the same Tencent Cloud account as the certificate application.
	// 
	// DNS: Manually add domain DNS validation. Requires the user to manually add the validation value at their domain's DNS service provider.
	// 
	// FILE: Manually add domain file validation. Requires the user to manually add a specified path file in the root directory of the domain site for file validation. Either HTTP or HTTPS validation will suffice; the domain site must be accessible by overseas CA institutions. The specific access whitelist is: 64.78.193.238, 216.168.247.9, 216.168.249.9, 54.189.196.217.
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// The domain bound to the certificate.
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// The project ID associated with the certificate. Default is 0 (default project)
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate type, optional, currently only type 83 is supported. 83 = trustasia c1 dv free.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// The email associated with the certificate order, By default, it uses the Tencent Cloud account email. If it does not exist, a fixed email address will be used.
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// The mobile phone number associated with the certificate. If it does not exist, a fixed mobile phone number will be used.
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// Certificate valid period, 3 months by default, currently only 3 months is supported.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Encryption algorithm, values can be ECC or RSA, default is RSA.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// Key pair parameters. RSA supports only 2048. ECC supports only prime256v1. When the encryption algorithm is set to ECC, this parameter is mandatory.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// Private key password, currently only used when generating jks, pfx format certificates; private key certificates of other formats are not encrypted.
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// Certificate alias.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Old certificate id, used for certificate renewal (the certificate valid period is within 30 days and not expired), a renewal relationship will be established, which can be hosted; not providing it means applying for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Benefit package ID, used for free certificate expansion package, the free certificate expansion package has been discontinued.
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is fasle by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`

	// Other domains bound to the certificate, to be opened. This parameter is not currently supported.
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
	// The newly applied certificate id.
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

type COSInstanceList struct {
	// Region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance details.
	InstanceList []*CosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Total number under the region.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Error message.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Domain name status: rejected - the domain name failed the review or its registration has expired/been canceled; processing - deploying; online - started; offline - closed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Domain billing status, where on indicates enable and off indicates disable.
	HttpsBillingSwitch *string `json:"HttpsBillingSwitch,omitnil,omitempty" name:"HttpsBillingSwitch"`
}

type CdnInstanceList struct {
	// The total number of CDN domains in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CDN domain name details.	
	InstanceList []*CdnInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Whether to query exceptions.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type CertBasicInfo struct {
	// Issuer.
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// Issued to.
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Certificate fingerprint.
	Fingerprint *string `json:"Fingerprint,omitnil,omitempty" name:"Fingerprint"`

	// Certificate valid period start time.
	ValidFrom *string `json:"ValidFrom,omitnil,omitempty" name:"ValidFrom"`

	// Certificate valid period end time.
	ValidTo *string `json:"ValidTo,omitnil,omitempty" name:"ValidTo"`
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

	// Root certificate id.
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// Certificate authentication mode: unidirectional one-way authentication, mutual mutual authentication.
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`
}

type CertificateExtra struct {
	// Quantity of configurable domain names for the certificate.
	DomainNumber *string `json:"DomainNumber,omitnil,omitempty" name:"DomainNumber"`

	// Renew the original certificate id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginCertificateId *string `json:"OriginCertificateId,omitnil,omitempty" name:"OriginCertificateId"`

	// Original ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedBy *string `json:"ReplacedBy,omitnil,omitempty" name:"ReplacedBy"`

	// Reissue certificate id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedFor *string `json:"ReplacedFor,omitnil,omitempty" name:"ReplacedFor"`

	// Renewal certificate id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewOrder *string `json:"RenewOrder,omitnil,omitempty" name:"RenewOrder"`

	// Whether it is a China SM certificate.
	SMCert *int64 `json:"SMCert,omitnil,omitempty" name:"SMCert"`

	// Company type, valid values: 1 (individual); 2 (company).
	CompanyType *int64 `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`
}

type Certificates struct {
	// User uin.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project id.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate source:.
	// trustasia.
	// upload.
	// wosign.
	// sheca.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Certificate package type:.
	// Null: user uploads a certificate (without package type),.
	// 2: trustasia tls rsa ca,. 
	// 3: securesite enhanced enterprise edition (ev pro),. 
	// 4: securesite enhanced (ev). 
	// 5: securesite enterprise professional edition (ov pro).
	// 6: securesite enterprise (ov). 
	// 7: securesite enterprise (ov) wildcard. 
	// 8: geotrust enhanced (ev). 
	// 9: geotrust enterprise (ov) cert. 
	// 10: geotrust enterprise (ov) wildcard cert. 
	// 11: trustasia domain name-based multiple domain names ssl certificate. 
	// 12: trustasia domain name-based (dv) wildcard cert. 
	// 13: trustasia enterprise wildcard (ov) ssl certificate (d3). 
	// 14: trustasia enterprise (ov) ssl certificate (d3). 
	// 15: trustasia enterprise multiple domain names (ov) ssl certificate (d3). 
	// 16: trustasia enhanced (ev) ssl certificate (d3). 
	// 17: trustasia enhanced multiple domain names (ev) ssl certificate (d3). 
	// 18: globalsign enterprise (ov) ssl certificate. 
	// 19: globalsign enterprise wildcard (ov) ssl certificate. 
	// 20: globalsign enhanced (ev) ssl certificate. 
	// 21: trustasia enterprise wildcard multiple domain names (ov) ssl certificate (d3). 
	// 22: globalsign enterprise multiple domain names (ov) ssl certificate. 
	// 23: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 24: globalsign enhanced multiple domain name (ev) ssl certificate.
	// 25: wotrus domain name cert.
	// 26: wotrus domain name multiple domain name cert.
	// 27: wotrus domain name wildcard cert.
	// 28: wotrus enterprise cert.
	// 29: wotrus enterprise multi - domain name certificate.
	// 30: wotrus enterprise wildcard certificate.
	// 31: wotrus enhanced certificate.
	// 32: wotrus enhanced multi - domain name certificate.
	// 33: wotrus - national cryptography domain - type certificate.
	// 34: wotrus-national cryptography domain certificate (multiple domain names).
	// 35: wotrus-national cryptography domain certificate (wildcard).
	// 37: wotrus-national cryptography enterprise certificate.
	// 38: wotrus-national cryptography enterprise certificate (multiple domain names).
	// 39: wotrus-national cryptography enterprise certificate (wildcard).
	// 40: wotrus - enhanced national cryptography certificate.
	// 41: wotrus - enhanced national cryptography certificate (multiple domain names).
	// 42: trustasia - domain name type certificate (wildcard multiple domain names).
	// 43: DNSPod - enterprise (ov) ssl certificate.
	// 44: DNSPod - enterprise (ov) wildcard ssl certificate.
	// 45: DNSPod - enterprise (ov) multiple domain names ssl certificate.
	// 46: DNSPod - enhanced (ev) ssl certificate.
	// 47: DNSPod - enhanced (ev) multiple domain names ssl certificate.
	// 48: DNSPod - domain name-based (dv) ssl certificate.
	// 49: DNSPod - domain name-based (dv) wildcard ssl certificate.
	// 50: DNSPod - domain name-based (dv) multiple domain names ssl certificate.
	// 51: DNSPod (national cryptography) - enterprise (ov) ssl certificate.
	// 52: DNSPod (national cryptography) - enterprise (ov) wildcard ssl certificate.
	// 53: DNSPod (national cryptography) - enterprise (ov) multiple domain names ssl certificate.
	// 54: DNSPod (national cryptography) - domain name-based (dv) ssl certificate.
	// 55: DNSPod (national cryptography) - domain name-based (dv) wildcard ssl certificate.
	// 56: DNSPod (national cryptography) - domain name-based (dv) multiple domain names ssl certificate.
	// 57: securesite enterprise professional edition multiple domain names (ov pro).
	// 58: securesite enterprise multiple domain names (ov).
	// 59: securesite enhanced professional edition multiple domain names (ev pro).
	// 60: securesite enhanced multiple domain names (ev).
	// 61: geotrust enhanced multiple domain names (ev).
	// 75: securesite enterprise (ov).
	// 76: securesite enterprise (ov) wildcard.
	// 77: securesite enhanced (ev).
	// 78: geotrust enterprise (ov).
	// 79: geotrust enterprise wildcard (ov).
	// 80: geotrust enhanced (ev).
	// 81: globalsign enterprise (ov) ssl certificate.
	// 82: globalsign enterprise wildcard (ov) ssl certificate.
	// 83: trustasia c1 dv free.
	// 85: globalsign enhanced (ev) ssl certificate.
	// 88: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 89: globalsign enterprise multiple domain names (ov) ssl certificate.
	// 90: globalsign enhanced multiple domain names (ev) ssl certificate.
	// 91: geotrust enhanced multiple domain names (ev).
	// 92: securesite enterprise pro multiple domain names (ov pro).
	// 93: securesite enterprise multiple domain names (ov).
	// 94: securesite enhanced pro multiple domain names (ev pro).
	// 95: securesite enhanced multiple domain names (ev).
	// 96: securesite ev pro.
	// 97: securesite enterprise professional edition (ov pro).
	// 98: cfca enterprise (ov) ssl certificate.
	// 99: cfca enterprise ov ssl certificate for multiple domain names.
	// 100: cfca ov wildcard ssl certificate.
	// 101: cfca enhanced (ev) ssl certificate.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Certificate type. ca = client certificate; svr = server certificate.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Certificate product name.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Primary domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Remark name.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Certificate status: 0 = under review, 1 = approved, 2 = review failed, 3 = expired, 4 = dns record added automatically, 5 = enterprise certificate, pending documentation submission, 6 = order cancellation in progress, 7 = canceled, 8 = documents submitted, pending upload of confirmation letter, 9 = certificate revocation in progress, 10 = revoked, 11 = reissue in progress, 12 = pending upload of revocation confirmation letter, 13 = free certificate pending documentation submission, 14 = certificate refunded, 15 = certificate migration in progress.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Certificate extended information.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// Vulnerability scanning status: INACTIVE = not enabled, ACTIVE = enabled.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// Status information.
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// Validation type: DNS_AUTO = automatic dns validation, DNS = manual dns validation, FILE = file verification, DNS_PROXY = dns proxy validation, FILE_PROXY = file proxy verification.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// Certificate validation time.
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// Certificate expiration time.
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// Certificate validity period (month).
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// Creation time.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Certificate id.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Multiple domain names contained in the certificate (including the primary domain name).
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// Certificate type name.
	PackageTypeName *string `json:"PackageTypeName,omitnil,omitempty" name:"PackageTypeName"`

	// Status name.
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// Specifies whether the customer is a vip customer. true indicates yes and false indicates no.
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// Specifies whether it is a dv version certificate. true indicates yes and false indicates no.
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// Specifies whether it is a wildcard domain name certificate. true indicates yes and false indicates no.
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// Whether the vulnerability scanning feature is enabled.
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// Whether it is renewable.
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// Project information.
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitnil,omitempty" name:"ProjectInfo"`

	// Associated cloud resources are temporarily unavailable.
	BoundResource []*string `json:"BoundResource,omitnil,omitempty" name:"BoundResource"`

	// Whether it can be deployed.
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// Tag list.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether expiration notice has been ignored.
	IsIgnore *bool `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`

	// Whether it is a China SM certificate.
	IsSM *bool `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// Certificate algorithm.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// Encryption algorithm for upload ca certificate.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// Expiration time for upload ca certificate.
	CAEndTimes []*string `json:"CAEndTimes,omitnil,omitempty" name:"CAEndTimes"`

	// Common name of the upload ca certificate.
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// Certificate prereview information.
	PreAuditInfo *PreAuditInfo `json:"PreAuditInfo,omitnil,omitempty" name:"PreAuditInfo"`

	// Whether to auto-renew.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Hosting status: 0, hosting; 5, resource replacement; 10, hosting completed; -1, not hosted. 
	HostingStatus *int64 `json:"HostingStatus,omitnil,omitempty" name:"HostingStatus"`

	// Hosting completion time.
	HostingCompleteTime *string `json:"HostingCompleteTime,omitnil,omitempty" name:"HostingCompleteTime"`

	// Manage the new certificate id.
	HostingRenewCertId *string `json:"HostingRenewCertId,omitnil,omitempty" name:"HostingRenewCertId"`

	// Existing renewal certificate id.
	HasRenewOrder *string `json:"HasRenewOrder,omitnil,omitempty" name:"HasRenewOrder"`

	// Indicates whether the original certificate is deleted during reissue.
	ReplaceOriCertIsDelete *bool `json:"ReplaceOriCertIsDelete,omitnil,omitempty" name:"ReplaceOriCertIsDelete"`

	// Indicates whether it is about to expire. a certificate is about to expire if it will expire within 30 days.
	IsExpiring *bool `json:"IsExpiring,omitnil,omitempty" name:"IsExpiring"`

	// Add validation expiration date for DV certificate
	DVAuthDeadline *string `json:"DVAuthDeadline,omitnil,omitempty" name:"DVAuthDeadline"`

	// Domain verification passed time.
	ValidationPassedTime *string `json:"ValidationPassedTime,omitnil,omitempty" name:"ValidationPassedTime"`

	// Multiple domain names associated with the certificate.
	CertSANs []*string `json:"CertSANs,omitnil,omitempty" name:"CertSANs"`

	// Domain verification rejection information.
	AwaitingValidationMsg *string `json:"AwaitingValidationMsg,omitnil,omitempty" name:"AwaitingValidationMsg"`

	// Whether to allow downloading.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// Whether all certificate domain names are managed and resolved by dnspod.
	IsDNSPODResolve *bool `json:"IsDNSPODResolve,omitnil,omitempty" name:"IsDNSPODResolve"`

	// Whether the certificate is purchased with benefit points.
	IsPackage *bool `json:"IsPackage,omitnil,omitempty" name:"IsPackage"`

	// Whether there is a private key password.
	KeyPasswordCustomFlag *bool `json:"KeyPasswordCustomFlag,omitnil,omitempty" name:"KeyPasswordCustomFlag"`

	// Types of web servers supported for download: nginx, apache, iis, tomcat, jks, root, other.
	SupportDownloadType *SupportDownloadType `json:"SupportDownloadType,omitnil,omitempty" name:"SupportDownloadType"`

	// Certificate revocation completion time.
	CertRevokedTime *string `json:"CertRevokedTime,omitnil,omitempty" name:"CertRevokedTime"`

	// Hosted resource type list.
	HostingResourceTypes []*string `json:"HostingResourceTypes,omitnil,omitempty" name:"HostingResourceTypes"`

	// Managed configuration information.
	HostingConfig *HostingConfig `json:"HostingConfig,omitnil,omitempty" name:"HostingConfig"`
}

type ClbInstanceDetail struct {
	// The CLB instance ID.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// The CLB instance name.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB listener list.
	Listeners []*ClbListener `json:"Listeners,omitnil,omitempty" name:"Listeners"`
}

type ClbInstanceList struct {
	// The region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CLB instance details.
	InstanceList []*ClbInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of CLB instances in this region.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Whether to query exceptions.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Data of certificate bound to the listener.
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// List of listener rules.
	Rules []*ClbListenerRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// Domain list not matched.
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

type ClbListenerRule struct {
	// The rule ID.
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// The domains bound.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Whether the rule matches the domains to be associated with a certificate.
	IsMatch *bool `json:"IsMatch,omitnil,omitempty" name:"IsMatch"`

	// Certificate data bound to the rule.
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// Domain list not matched.
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`

	// Rule binding path.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type CommitCertificateInformationRequestParams struct {
	// Paid certificate id of materials to be submitted.	
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Certificate domain name verification method:.
	// DNS_AUTO: automatically adds domain dns verification, requiring user domain name resolution to be hosted on [cloud dns](https://console.cloud.tencent.com/cns) and under the same tencent cloud account as the certificate application.
	// DNS: manually add domain dns verification, which requires users to manually add verification values to the domain resolution service provider.
	// FILE: manual addition of domain name file verification. requires the user to manually add a specified path file in the root directory of the domain site for file verification, and either http or https passing is acceptable; the domain site needs to be accessible by overseas ca institutions, with the specific access allowlist being: 64.78.193.238, 216.168.247.9, 216.168.249.9, 54.189.196.217.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Paid certificate id of materials to be submitted.	
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// Certificate domain name verification method:.
	// DNS_AUTO: automatically adds domain dns verification, requiring user domain name resolution to be hosted on [cloud dns](https://console.cloud.tencent.com/cns) and under the same tencent cloud account as the certificate application.
	// DNS: manually add domain dns verification, which requires users to manually add verification values to the domain resolution service provider.
	// FILE: manual addition of domain name file verification. requires the user to manually add a specified path file in the root directory of the domain site for file verification, and either http or https passing is acceptable; the domain site needs to be accessible by overseas ca institutions, with the specific access allowlist being: 64.78.193.238, 216.168.247.9, 216.168.249.9, 54.189.196.217.
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

type CosInstanceDetail struct {
	// Domain name.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Bound certificate id.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// ENABLED: domain name online status.
	// DISABLED: domain name offline status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// bucket name.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// bucket region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
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


	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
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

	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
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
	// Certificate product id. `3`: securesite ev pro; `4`: securesite ev; `5`: securesite ov pro; `6`: securesite ov; `7`: securesite ov wildcard; `8`: geotrust ev; `9`: geotrust ov; `10`: geotrust ov wildcard; `11`: trustasia dv multi-domain; `12`: trustasia dv wildcard; `13`: trustasia ov wildcard d3; `14`: trustasia ov d3; `15`: trustasia ov multi-domain d3; `16`: trustasia ev d3; `17`: trustasia ev multi-domain d3; `18`: globalsign ov; `19`: globalsign ov wildcard; `20`: globalsign ev; `21`: trustasia ov wildcard multi-domain d3; `22`: globalsign ov multi-domain; `23`: globalsign ov wildcard multi-domain; `24`: globalsign ev multi-domain; `25`: wotrus dv; `26`: wotrus dv multi-domain; `27`: wotrus dv wildcard; `28`: wotrus ov; `29`: wotrus ov multi-domain; `30`: wotrus ov wildcard; `31`: wotrus ev; `32`: wotrus ev multi-domain; `33`: DNSPod sm2 dv; `34`: DNSPod sm2 dv multi-domain; `35`: DNSPod sm2 dv wildcard; `37`: DNSPod sm2 ov; `38`: DNSPod sm2 ov multi-domain; `39`: DNSPod sm2 ov wildcard; `40`: DNSPod sm2 ev; `41`: DNSPod sm2 ev multi-domain; `42`: trustasia dv wildcard multi-domain; `43`: dnspod-ov ssl certificate; `44`: dnspod-ov wildcard ssl certificate; `45`: dnspod-ov multi-domain ssl certificate; `46`: dnspod-ev ssl certificate; `47`: dnspod-ev multi-domain ssl certificate; `48`: dnspod-dv ssl certificate; `49`: dnspod-dv wildcard ssl certificate; `50`: dnspod-dv multi-domain ssl certificate; `51`: DNSPod (sm2)-ov ssl certificate; `52`: DNSPod (sm2)-ov wildcard ssl certificate; `53`: DNSPod (sm2)-ov multi-domain ssl certificate; `54`: DNSPod (sm2)-dv ssl certificate; `55`: DNSPod (sm2)-dv wildcard ssl certificate; `56`: DNSPod (sm2)-dv multi-domain ssl certificate; `57`: securesite ov pro multi-domain; `58`: securesite ov multi-domain; `59`: securesite ev pro multi-domain; `60`: securesite ev multi-domain; `61`: geotrust ev multi-domain.
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// Certificate validity period.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Whether to automatically use vouchers: 1 for yes, 0 for no; the default is 1.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Tag, generate tags for certificates.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate product id. `3`: securesite ev pro; `4`: securesite ev; `5`: securesite ov pro; `6`: securesite ov; `7`: securesite ov wildcard; `8`: geotrust ev; `9`: geotrust ov; `10`: geotrust ov wildcard; `11`: trustasia dv multi-domain; `12`: trustasia dv wildcard; `13`: trustasia ov wildcard d3; `14`: trustasia ov d3; `15`: trustasia ov multi-domain d3; `16`: trustasia ev d3; `17`: trustasia ev multi-domain d3; `18`: globalsign ov; `19`: globalsign ov wildcard; `20`: globalsign ev; `21`: trustasia ov wildcard multi-domain d3; `22`: globalsign ov multi-domain; `23`: globalsign ov wildcard multi-domain; `24`: globalsign ev multi-domain; `25`: wotrus dv; `26`: wotrus dv multi-domain; `27`: wotrus dv wildcard; `28`: wotrus ov; `29`: wotrus ov multi-domain; `30`: wotrus ov wildcard; `31`: wotrus ev; `32`: wotrus ev multi-domain; `33`: DNSPod sm2 dv; `34`: DNSPod sm2 dv multi-domain; `35`: DNSPod sm2 dv wildcard; `37`: DNSPod sm2 ov; `38`: DNSPod sm2 ov multi-domain; `39`: DNSPod sm2 ov wildcard; `40`: DNSPod sm2 ev; `41`: DNSPod sm2 ev multi-domain; `42`: trustasia dv wildcard multi-domain; `43`: dnspod-ov ssl certificate; `44`: dnspod-ov wildcard ssl certificate; `45`: dnspod-ov multi-domain ssl certificate; `46`: dnspod-ev ssl certificate; `47`: dnspod-ev multi-domain ssl certificate; `48`: dnspod-dv ssl certificate; `49`: dnspod-dv wildcard ssl certificate; `50`: dnspod-dv multi-domain ssl certificate; `51`: DNSPod (sm2)-ov ssl certificate; `52`: DNSPod (sm2)-ov wildcard ssl certificate; `53`: DNSPod (sm2)-ov multi-domain ssl certificate; `54`: DNSPod (sm2)-dv ssl certificate; `55`: DNSPod (sm2)-dv wildcard ssl certificate; `56`: DNSPod (sm2)-dv multi-domain ssl certificate; `57`: securesite ov pro multi-domain; `58`: securesite ov multi-domain; `59`: securesite ev pro multi-domain; `60`: securesite ev multi-domain; `61`: geotrust ev multi-domain.
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// Certificate validity period.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Whether to automatically use vouchers: 1 for yes, 0 for no; the default is 1.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Tag, generate tags for certificates.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "AutoVoucher")
	delete(f, "Tags")
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

	// Certificate id.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// The forwarding port.
	VirtualPort *string `json:"VirtualPort,omitnil,omitempty" name:"VirtualPort"`
}

type DdosInstanceList struct {
	// The total number of DDOS instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// DDOS instance details.	
	InstanceList []*DdosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Whether to query exceptions.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type DeleteCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// When deleting, check whether the certificate is associated with cloud resources. By default, no check is performed. if you choose to check (the authorization service role SSL_QCSLinkedRoleInReplaceLoadCertificate is required), the deletion will become asynchronous, and the API will return an asynchronous task id. you need to use the DescribeDeleteCertificatesTaskResult API to check whether the deletion is successful.
	IsCheckResource *bool `json:"IsCheckResource,omitnil,omitempty" name:"IsCheckResource"`
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// When deleting, check whether the certificate is associated with cloud resources. By default, no check is performed. if you choose to check (the authorization service role SSL_QCSLinkedRoleInReplaceLoadCertificate is required), the deletion will become asynchronous, and the API will return an asynchronous task id. you need to use the DescribeDeleteCertificatesTaskResult API to check whether the deletion is successful.
	IsCheckResource *bool `json:"IsCheckResource,omitnil,omitempty" name:"IsCheckResource"`
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
	delete(f, "IsCheckResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateResponseParams struct {
	// Deletion result
	DeleteResult *bool `json:"DeleteResult,omitnil,omitempty" name:"DeleteResult"`

	// Asynchronous deletion task id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

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

type DeployRecord struct {
	// Total number.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total successes
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitnil,omitempty" name:"SuccessTotalCount"`

	// Total failed deployments.
	FailedTotalCount *int64 `json:"FailedTotalCount,omitnil,omitempty" name:"FailedTotalCount"`

	// Deployment in progress total number.
	RunningTotalCount *int64 `json:"RunningTotalCount,omitnil,omitempty" name:"RunningTotalCount"`

	// Deployment record type 0 specifies deployment, 1 indicates rollback.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Deployment record detail list.
	RecordDetailList []*DeployRecordList `json:"RecordDetailList,omitnil,omitempty" name:"RecordDetailList"`

	// Hosted resource deployment state: `0` (awaiting deployment), `1` (deployment successful), `2` (deployment failed), `3` (deploying), `4` (rollback successful), `5` (rollback failure).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Managed resource creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DeployRecordItem struct {
	// Deployment record detail ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Bound certificate ID.
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// Deployment instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Specifies the deployment instance name.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Deploy the listener ID.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// Deployment domain name list.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Deploy listener protocol.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Deployment state.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Deployment error information.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Describes the creation time of the deployment record detail.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Specifies the last update time of the deployment record detail.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Deploy listener name.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Whether SNI is enabled.
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// COS bucket name.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// The secret name.
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Deployed TCB region.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CLB type. valid values: 0 (classic clb); 1 (application clb).
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// Certificate authentication mode: UNIDIRECTIONAL for one-way authentication, MUTUAL for MUTUAL authentication.
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// Deployment resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type DeployRecordList struct {
	// Deployment resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Deployment resource detail list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*DeployRecordItem `json:"List,omitnil,omitempty" name:"List"`

	// Total count of deployment resources.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
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
	// Task id, which can be used to query the result of binding cloud resources according to the task id obtained from createcertificatebindresourcesynctask.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The number of cloud resources displayed on each page. The default value is 10, and the maximum value is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Current offset, default is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Result detail of queried resource type. if not provided, all will be queried. valid values include:.
	// - clb.
	// - cdn.
	// - ddos.
	// - live.
	// - vod.
	// - waf.
	// - apigateway.
	// - teo.
	// - tke.
	// - cos.
	// - tse.
	// - tcb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Data of querying region list. clb, tke, waf, api gateway, tcb, cos, and tse support region query, while other resource types do not support.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeCertificateBindResourceTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Task id, which can be used to query the result of binding cloud resources according to the task id obtained from createcertificatebindresourcesynctask.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The number of cloud resources displayed on each page. The default value is 10, and the maximum value is 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Current offset, default is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Result detail of queried resource type. if not provided, all will be queried. valid values include:.
	// - clb.
	// - cdn.
	// - ddos.
	// - live.
	// - vod.
	// - waf.
	// - apigateway.
	// - teo.
	// - tke.
	// - cos.
	// - tse.
	// - tcb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Data of querying region list. clb, tke, waf, api gateway, tcb, cos, and tse support region query, while other resource types do not support.
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

	// The status of the async task. Valid values: `0` for querying, `1` for successful, and `2` for abnormal. If the status is `1`, check the result of `BindResourceResult` ; if the status is `2`, check the `error` .
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The cache time of the current result.
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// Associated TSE resource details
	// Note: This field may return null, indicating that no valid value can be obtained.
	TSE []*TSEInstanceList `json:"TSE,omitnil,omitempty" name:"TSE"`

	// Information of associated cos resource.
	// Note: this field may return null, indicating that no valid values can be obtained.
	COS []*COSInstanceList `json:"COS,omitnil,omitempty" name:"COS"`

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
	// Certificate belonging to user main account uin.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Certificate source:.
	// trustAsia.
	// upload.
	// wosign.
	// sheca.
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Certificate package type:.
	// null: user uploads a certificate (no package type),.
	// 2: trustasia tls rsa ca,. 
	// 3: securesite enhanced enterprise version (ev pro),. 
	// 4: securesite enhanced (ev),. 
	// 5: securesite enterprise pro (ov pro).
	// 6: securesite enterprise (ov). 
	// 7: securesite enterprise (ov) wildcard. 
	// 8: geotrust enhanced (ev). 
	// 9: geotrust enterprise (ov). 
	// 10: geotrust enterprise (ov) wildcard cert. 
	// 11: trustasia domain name-based multiple domain names ssl certificate. 
	// 12: trustasia domain name-based (dv) wildcard cert. 
	// 13: trustasia enterprise wildcard (ov) ssl certificate (d3). 
	// 14: trustasia enterprise (ov) ssl certificate (d3). 
	// 15: trustasia enterprise multiple domain names (ov) ssl certificate (d3). 
	// 16: trustasia enhanced (ev) ssl certificate (d3). 
	// 17: trustasia enhanced multiple domain names (ev) ssl certificate (d3). 
	// 18: globalsign enterprise (ov) ssl certificate. 
	// 19: globalsign enterprise wildcard (ov) ssl certificate. 
	// 20: globalsign enhanced (ev) ssl certificate. 
	// 21: trustasia enterprise wildcard multiple domain names (ov) ssl certificate (d3). 
	// 22: globalsign enterprise multiple domain names (ov) ssl certificate. 
	// 23: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 24: globalsign enhanced multiple domain names (ev) ssl certificate.
	// 25: wotrus domain cert.
	// 26: wotrus multi-domain cert.
	// 27: wotrus wildcard cert.
	// 28: wotrus enterprise cert.
	// 29: wotrus enterprise multi-domain cert.
	// 30: wotrus enterprise wildcard certificate.
	// 31: wotrus enhanced certificate.
	// 32: wotrus enhanced multi-domain name certificate.
	// 33: wotrus-national cryptography domain name certificate.
	// 34: wotrus-national cryptography domain name certificate (multiple domain names).
	// 35: wotrus-national cryptography wildcard certificate.
	// 37: wotrus-national cryptography enterprise certificate.
	// 38: wotrus-national cryptography enterprise certificate (multiple domain names).
	// 39: wotrus-national cryptography enterprise certificate (wildcard).
	// 40: wotrus-national cryptography enhanced certificate.
	// 41: wotrus - national cryptography enhanced certificate (multiple domain names).
	// 42: trustasia - domain name certificate (wildcard multiple domain names).
	// 43: DNSPod - enterprise (ov) ssl certificate.
	// 44: DNSPod - enterprise (ov) wildcard ssl certificate.
	// 45: DNSPod - enterprise (ov) multiple domain names ssl certificate.
	// 46: dnspod-enhanced (ev) ssl certificate.
	// 47: dnspod-enhanced (ev) multiple domain names ssl certificate.
	// 48: dnspod-domain name-based (dv) ssl certificate.
	// 49: dnspod-domain name-based (dv) wildcard ssl certificate.
	// 50: dnspod-domain name-based (dv) multiple domain names ssl certificate.
	// 51: DNSPod (national cryptography) - enterprise (ov) ssl certificate.
	// 52: DNSPod (national cryptography) - enterprise (ov) wildcard ssl certificate.
	// 53: DNSPod (national cryptography) - enterprise (ov) multiple domain names ssl certificate.
	// 54: DNSPod (national cryptography) - domain name-based (dv) ssl certificate.
	// 55: DNSPod (national cryptography) - domain name-based (dv) wildcard ssl certificate.
	// 56: DNSPod (national cryptography) - domain name-based (dv) multiple domain names ssl certificate.
	// 57: securesite enterprise professional version multiple domain names (ov pro).
	// 58: securesite enterprise multiple domain names (ov).
	// 59: securesite enhanced professional version multiple domain names (ev pro).
	// 60: securesite enhanced multiple domain names (ev).
	// 61: geotrust enhanced multiple domain names (ev).
	// 75: securesite enterprise (ov).
	// 76: securesite enterprise (ov) wildcard.
	// 77: securesite enhanced (ev).
	// 78: geotrust enterprise (ov).
	// 79: geotrust enterprise (ov) wildcard.
	// 80: geotrust enhanced (ev).
	// 81: globalsign enterprise (ov) ssl certificate.
	// 82: globalsign enterprise wildcard (ov) ssl certificate.
	// 83: trustasia c1 dv free.
	// 85: globalsign enhanced (ev) ssl certificate.
	// 88: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 89: globalsign enterprise multiple domain names (ov) ssl certificate.
	// 90: globalsign enhanced multiple domain names (ev) ssl certificate.
	// 91: geotrust enhanced multiple domain names (ev).
	// 92: securesite enterprise ov pro for multiple domain names.
	// 93: securesite enterprise for multiple domain names (ov).
	// 94: securesite ev pro for multiple domain names.
	// 95: securesite ev for multiple domain names.
	// 96: securesite ev pro.
	// 97: securesite enterprise professional edition (ov pro).
	// 98: cfca enterprise (ov) ssl certificate.
	// 99: cfca enterprise multiple domain names (ov) ssl certificate.
	// 100: cfca enterprise wildcard (ov) ssl certificate.
	// 101: cfca enhanced (ev) ssl certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Certificate product name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Certificate binds to a common name domain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Certificate status: 0 = under review, 1 = approved, 2 = review failed, 3 = expired, 4 = automatically added dns records, 5 = enterprise certificate, pending document submission, 6 = order cancellation in progress, 7 = canceled, 8 = documents submitted, pending upload of confirmation letter, 9 = certificate revocation in progress, 10 = revoked, 11 = reissue in progress, 12 = pending upload of revocation confirmation letter, 13 = free certificate pending document submission, 14 = certificate has been refunded, 15 = certificate migration in progress.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status information. valid values:.
	// //Common status information.
	// PRE-REVIEWING: in prereview.
	// LEGAL-REVIEWING: in legal review.
	// CA-REVIEWING: in ca review.
	// PENDING-DCV: in domain verification.
	// WAIT-ISSUE: waiting for issue (domain verification passed).
	// Certificate review failure status information.
	// 1. order review failed.
	// 2. ca review failed, and the domain name did not pass the security review.
	// 3. domain name verification timed out, and the order was automatically closed. please reapply for the certificate.
	// 4. the certificate information did not pass the review of the certificate ca agency. the reviewer will call the contact information reserved for the certificate. please pay attention to the incoming call. subsequently, you can resubmit the information through "modify information".
	// To be continuously improved.
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

	// Certificate application time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// CA order id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// Private key certificate; for Chinese SM certificates, it refers to the private key certificate in the signature certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Public key certificate; for Chinese SM certificate, it refers to the public key certificate in the signature certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Certificate domain name verification information.
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

	// Profile information submitted for paid certificates.
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

	// Chinese SM certificate public key, only has value for national cryptography certificates.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EncryptCert *string `json:"EncryptCert,omitnil,omitempty" name:"EncryptCert"`

	// Chinese SM certificate private key certificate, only has value for national cryptography certificates.
	// Note: this field may return null, indicating that no valid values can be obtained.
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitnil,omitempty" name:"EncryptPrivateKey"`

	// SHA1 fingerprint of the signature certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertFingerprint *string `json:"CertFingerprint,omitnil,omitempty" name:"CertFingerprint"`

	// SHA1 fingerprint of the encryption certificate (for Chinese SM certificates only)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptCertFingerprint *string `json:"EncryptCertFingerprint,omitnil,omitempty" name:"EncryptCertFingerprint"`

	// Certificate encryption algorithm (or Chinese SM certificates only).
	// Note: this field may return null, indicating that no valid values can be obtained.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// The authentication value for DV certificate revocation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitnil,omitempty" name:"DvRevokeAuthDetail"`

	// Certificate chain information.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertChainInfo []*CertBasicInfo `json:"CertChainInfo,omitnil,omitempty" name:"CertChainInfo"`

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

	// Number of requested logs, 20 by default, with a maximum value of 1000. if it exceeds 1000, it will be treated as 1000.
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

	// Number of requested logs, 20 by default, with a maximum value of 1000. if it exceeds 1000, it will be treated as 1000.
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

	// Certificate source:
	// trustAsia.
	// upload.
	// wosign.
	// sheca.
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Certificate package type:.
	// Null: user uploads a certificate (without package type),.
	// 2: trustasia tls rsa ca,. 
	// 3: securesite enhanced enterprise edition (ev pro),. 
	// 4: securesite enhanced (ev),. 
	// 5: securesite enterprise professional edition (ov pro).
	// 6: securesite enterprise edition (ov). 
	// 7: securesite enterprise edition (ov) wildcard. 
	// 8: geotrust enhanced (ev). 
	// 9: geotrust enterprise edition (ov). 
	// 10: geotrust enterprise (ov) wildcard cert. 
	// 11: trustasia domain name-based multiple domain names ssl certificate. 
	// 12: trustasia domain name-based (dv) wildcard cert. 
	// 13: trustasia enterprise wildcard (ov) ssl certificate (d3). 
	// 14: trustasia enterprise (ov) ssl certificate (d3). 
	// 15: trustasia enterprise multiple domain names (ov) ssl certificate (d3). 
	// 16: trustasia enhanced (ev) ssl certificate (d3). 
	// 17: trustasia enhanced multiple domain names (ev) ssl certificate (d3). 
	// 18: globalsign enterprise (ov) ssl certificate. 
	// 19: globalsign enterprise wildcard (ov) ssl certificate. 
	// 20: globalsign enhanced (ev) ssl certificate. 
	// 21: trustasia enterprise wildcard multiple domain names (ov) ssl certificate (d3). 
	// 22: globalsign enterprise multiple domain names (ov) ssl certificate. 
	// 23: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 24: globalsign enhanced multiple domain names (ev) ssl certificate.
	// 25: wotrus domain cert.
	// 26: wotrus multi - domain name cert.
	// 27: wotrus wildcard cert.
	// 28: wotrus enterprise cert.
	// 29: wotrus enterprise multi - domain name cert.
	// 30: wotrus enterprise wildcard certificate.
	// 31: wotrus enhanced certificate.
	// 32: wotrus enhanced multi - domain name certificate.
	// 33: wotrus - national cryptography domain - name certificate.
	// 34: wotrus - national cryptography domain - name certificate (multiple domain names).
	// 35: wotrus-national cryptography wildcard domain certificate.
	// 37: wotrus-national cryptography enterprise certificate.
	// 38: wotrus-national cryptography enterprise certificate (multiple domain names).
	// 39: wotrus-national cryptography enterprise certificate (wildcard).
	// 40: wotrus-national cryptography enhanced certificate.
	// 41: wotrus - national cryptography enhanced certificate (multiple domain names).
	// 42: trustasia - domain name certificate (wildcard multiple domain names).
	// 43: DNSPod - enterprise (ov) ssl certificate.
	// 44: DNSPod - enterprise (ov) wildcard ssl certificate.
	// 45: DNSPod - enterprise (ov) multiple domain names ssl certificate.
	// 46: dnspod-enhanced (ev) ssl certificate.
	// 47: dnspod-enhanced (ev) multiple domain names ssl certificate.
	// 48: dnspod-domain name-based (dv) ssl certificate.
	// 49: dnspod-domain name-based (dv) wildcard ssl certificate.
	// 50: dnspod-domain name-based (dv) multiple domain names ssl certificate.
	// 51: DNSPod (national cryptography) - enterprise (ov) ssl certificate.
	// 52: DNSPod (national cryptography) - enterprise (ov) wildcard ssl certificate.
	// 53: DNSPod (national cryptography) - enterprise (ov) multiple domain names ssl certificate.
	// 54: DNSPod (national cryptography) - domain name-based (dv) ssl certificate.
	// 55: DNSPod (national cryptography) - domain name-based (dv) wildcard ssl certificate.
	// 56: DNSPod (national cryptography) - domain name-based (dv) multiple domain names ssl certificate.
	// 57: securesite enterprise professional version multiple domain names (ov pro).
	// 58: securesite enterprise multiple domain names (ov).
	// 59: securesite enhanced professional version multiple domain names (ev pro).
	// 60: securesite enhanced multiple domain names (ev).
	// 61: geotrust enhanced multiple domain names (ev).
	// 75: securesite enterprise (ov).
	// 76: securesite enterprise (ov) wildcard.
	// 77: securesite enhanced (ev).
	// 78: geotrust enterprise (ov).
	// 79: geotrust enterprise (ov) wildcard.
	// 80: geotrust enhanced (ev).
	// 81: globalsign enterprise (ov) ssl certificate.
	// 82: globalsign enterprise wildcard (ov) ssl certificate.
	// 83: trustasia c1 dv free.
	// 85: globalsign enhanced (ev) ssl certificate.
	// 88: globalsign enterprise wildcard multiple domain names (ov) ssl certificate.
	// 89: globalsign enterprise multiple domain names (ov) ssl certificate.
	// 90: globalsign enhanced multiple domain names (ev) ssl certificate.
	// 91: geotrust enhanced multiple domain names (ev).
	// 92: securesite enterprise ov pro for multiple domain names.
	// 93: securesite enterprise for multiple domain names (ov).
	// 94: securesite ev pro for multiple domain names.
	// 95: securesite ev for multiple domain names.
	// 96: securesite ev pro.
	// 97: securesite enterprise professional version (ov pro).
	// 98: cfca enterprise (ov) ssl certificate.
	// 99: cfca enterprise multiple domain names (ov) ssl certificate.
	// 100: cfca enterprise wildcard (ov) ssl certificate.
	// 101: cfca enhanced (ev) ssl certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// Certificate product name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Certificate status: 0 = under review, 1 = approved, 2 = review failed, 3 = expired, 4 = dns records added automatically, 5 = enterprise certificate, pending documentation submission, 6 = order cancellation in progress, 7 = canceled, 8 = documents submitted, pending upload of confirmation letter, 9 = certificate revocation in progress, 10 = revoked, 11 = reissue in progress, 12 = pending upload of revocation confirmation letter, 13 = free certificate pending document submission, 14 = certificate has been refunded, 15 = certificate migration in progress.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status information. valid values:.
	// //Common status information.
	// 1. pre-reviewing: prereviewing.
	// 2. legal-reviewing: under legal review.
	// 3. ca-reviewing: under ca review.
	// 4. pending-dcv: under domain verification.
	// 5. wait-issue: waiting for issuance (domain verification passed).
	// //Certificate review failure status information.
	// Order review failed.
	// CA review failed; the domain name did not pass the security review.
	// Domain verification timed out, and the order was automatically closed. please reapply for the certificate.
	// The certificate information did not pass the review by the certificate authority. the reviewer will call the contact information reserved for the certificate. please pay attention to the incoming call. subsequently, you can resubmit the information through "modify information".
	// To be continuously improved.
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// Validation type: DNS_AUTO = automatic dns validation, DNS = manual dns validation, FILE = file verification, DNS_PROXY = dns proxy validation, FILE_PROXY = file proxy validation.
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

	// All encryption methods of the ca certificate. only valid when the certificate type CertificateType is ca.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// All common names of the ca certificate. only valid when the certificate type CertificateType is ca.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// All expiration times of the ca certificate. only valid when the certificate type CertificateType is ca.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Pagination offset, starting from 0. default is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default is 10. maximum value is 1000; values exceeding 1000 will be treated as 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search keywords, supporting fuzzy match by certificate id, remark name, and certificate domain name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Default sorting is by certificate application time in descending order. Sort by expiration date if the following values are passed: DESC for descending order of certificate expiration time, ASC for ascending order.
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// Certificate status: 0=under review, 1=approved, 2=review failed, 3=expired, 4=dns record added, 5=enterprise certificate, pending submission, 6=order cancellation in progress, 7=canceled, 8=documents submitted, pending upload of confirmation letter, 9=certificate revocation in progress, 10=revoked, 11=reissue in progress, 12=pending upload of revocation confirmation letter, 13=free certificate pending document submission, 14=refunded, 15=certificate migration in progress.
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

	// Filter certificates with specified tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to filter certificates pending issue: 1 for filtering, 0 and null for no filtering.
	IsPendingIssue *int64 `json:"IsPendingIssue,omitnil,omitempty" name:"IsPendingIssue"`

	// Filter certificates by the specified certificate id, only supports certificate ids with permission.
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset, starting from 0. default is 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default is 10. maximum value is 1000; values exceeding 1000 will be treated as 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search keywords, supporting fuzzy match by certificate id, remark name, and certificate domain name.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Default sorting is by certificate application time in descending order. Sort by expiration date if the following values are passed: DESC for descending order of certificate expiration time, ASC for ascending order.
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// Certificate status: 0=under review, 1=approved, 2=review failed, 3=expired, 4=dns record added, 5=enterprise certificate, pending submission, 6=order cancellation in progress, 7=canceled, 8=documents submitted, pending upload of confirmation letter, 9=certificate revocation in progress, 10=revoked, 11=reissue in progress, 12=pending upload of revocation confirmation letter, 13=free certificate pending document submission, 14=refunded, 15=certificate migration in progress.
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

	// Filter certificates with specified tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to filter certificates pending issue: 1 for filtering, 0 and null for no filtering.
	IsPendingIssue *int64 `json:"IsPendingIssue,omitnil,omitempty" name:"IsPendingIssue"`

	// Filter certificates by the specified certificate id, only supports certificate ids with permission.
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
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
	delete(f, "Tags")
	delete(f, "IsPendingIssue")
	delete(f, "CertIds")
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
	//
	// Deprecated: ResourceType is deprecated.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Whether to query the cached results. Valid values: `1` (yes) and `0` (no). By default, the cached results within 30 minutes are queried.
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// The list of filter parameters. FilterKey: domainMatch (query the list of instances with matching or non-matching domains). FilterValue: `1` (default; query the list of instances with matching domains) or `0` (query the list of instances with non-matching domains).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The ID of the deployed certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Paging offset. default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default: 10. maximum value: 200.	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Asynchronous or not. 1 means yes, 0 means no. default: 0.
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

	// Paging offset. default value: 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default: 10. maximum value: 200.	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Asynchronous or not. 1 means yes, 0 means no. default: 0.
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
	// Teo instance list. if no value is obtained, an empty array is returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Deployment record id, which is the record id returned by calling the UpdateCertificateInstance api, or the record id returned by calling the UpdateCertificateRecordRollback rollback api.
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number of items per page. the default is 10. the maximum value is 200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, starting from 0. default is 0.
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeHostUpdateRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// Deployment record id, which is the record id returned by calling the UpdateCertificateInstance api, or the record id returned by calling the UpdateCertificateRecordRollback rollback api.
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number of items per page. the default is 10. the maximum value is 200.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset, starting from 0. default is 0.
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
	// If the total number cannot be obtained, return 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Certificate deployment record list; returns an empty array if no value is obtained.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RecordDetailList []*UpdateRecordDetails `json:"RecordDetailList,omitnil,omitempty" name:"RecordDetailList"`

	// Total number of successes; returns 0 if unavailable.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitnil,omitempty" name:"SuccessTotalCount"`

	// Total number of failures. if it cannot be obtained, return 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailedTotalCount *int64 `json:"FailedTotalCount,omitnil,omitempty" name:"FailedTotalCount"`

	// Total number of deployments in progress; returns 0 if unavailable.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
type DescribeHostUploadUpdateRecordDetailRequestParams struct {
	// Hosting record ID.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number of items per page. default is 10, maximum is 200. exceeds 200 will be set to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeHostUploadUpdateRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// Hosting record ID.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Number of items per page. default is 10, maximum is 200. exceeds 200 will be set to 200.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeHostUploadUpdateRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUploadUpdateRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUploadUpdateRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUploadUpdateRecordDetailResponseParams struct {
	// Hosting record detail list.
	DeployRecordDetail []*DeployRecord `json:"DeployRecordDetail,omitnil,omitempty" name:"DeployRecordDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostUploadUpdateRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUploadUpdateRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeHostUploadUpdateRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUploadUpdateRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUploadUpdateRecordRequestParams struct {
	// Pagination offset, starting from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default is 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Original certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostUploadUpdateRecordRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset, starting from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of items per page. default is 10.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Original certificate ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostUploadUpdateRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUploadUpdateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUploadUpdateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUploadUpdateRecordResponseParams struct {
	// Total number.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Certificate deployment record list.
	DeployRecordList []*UploadUpdateRecordInfo `json:"DeployRecordList,omitnil,omitempty" name:"DeployRecordList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostUploadUpdateRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUploadUpdateRecordResponseParams `json:"Response"`
}

func (r *DescribeHostUploadUpdateRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUploadUpdateRecordResponse) FromJsonString(s string) error {
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
	// Certificate domain name verification record key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// Certificate domain name verification record value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// Certificate domain name verification domain value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// Certificate domain name verification file path, used only for file and file_proxy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// Certificate domain name verification subdomain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitnil,omitempty" name:"DvAuthKeySubDomain"`

	// Certificate domain verification information; multiple domain verifications use this field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuths []*DvAuths `json:"DvAuths,omitnil,omitempty" name:"DvAuths"`
}

type DvAuths struct {
	// Certificate domain name verification record key.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// Certificate domain name verification record value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// Certificate domain name verification domain value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// Certificate domain name verification file path, used only for file and file_proxy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// Certificate domain name verification subdomain.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitnil,omitempty" name:"DvAuthSubDomain"`

	// Certificate domain verification type, valid values:.
	// TXT: add txt record for dns domain verification.
	// FILE: domain file verification.
	// CNAME: add cname record for dns domain verification.
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

type HostingConfig struct {
	// Hosted resource replacement time, defaults to 30 days before the certificate expiration if there is a renewal certificate, then replace.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplaceTime *int64 `json:"ReplaceTime,omitnil,omitempty" name:"ReplaceTime"`

	// Hosted send message type: 0, reminder message before hosted starts (you will receive this reminder message even if there is no renewal certificate); 1, reminder message when hosted starts (you will receive the message reminder only if there is a renewal certificate); 2, reminder message when hosted resource replacement fails; 3 reminder message when hosted resource replacement succeeds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MessageTypes []*int64 `json:"MessageTypes,omitnil,omitempty" name:"MessageTypes"`

	// Resource replacement start time.
	ReplaceStartTime *string `json:"ReplaceStartTime,omitnil,omitempty" name:"ReplaceStartTime"`

	// Resource replacement end time.
	ReplaceEndTime *string `json:"ReplaceEndTime,omitnil,omitempty" name:"ReplaceEndTime"`
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

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

// Predefined struct for user
type ModifyCertificatesExpiringNotificationSwitchRequestParams struct {
	// Certificate ID list. maximum of 50.
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 0: do not ignore notifications. 1: ignore notifications.
	SwitchStatus *uint64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`
}

type ModifyCertificatesExpiringNotificationSwitchRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID list. maximum of 50.
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 0: do not ignore notifications. 1: ignore notifications.
	SwitchStatus *uint64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`
}

func (r *ModifyCertificatesExpiringNotificationSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificatesExpiringNotificationSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "SwitchStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificatesExpiringNotificationSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificatesExpiringNotificationSwitchResponseParams struct {
	// Certificate ID list.
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCertificatesExpiringNotificationSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificatesExpiringNotificationSwitchResponseParams `json:"Response"`
}

func (r *ModifyCertificatesExpiringNotificationSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificatesExpiringNotificationSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {
	// Action performed on logs
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Time when the action is performed
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// Root account.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Sub-Account.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// Certificate id.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Each operation type corresponds to a specific operation description. the following is a textual explanation of each operation type and its description:.
	// 1. apply: indicates applying for a free cert.
	// 2. delete: indicates a deletion.
	// 3. download - represents the download operation.
	// 4. upload: indicates an upload operation.
	// 5. revoke: indicates revoking a cert.
	// 6. cancelRevoke - indicates canceling the revocation operation.
	// 7. updateAlias - indicates updating the remark information.
	// 8. changeProject - indicates assigning the certificate to a certain project.
	// 9. uploadConfirmLetter - indicates uploading the confirmation letter.
	// 10. cancel - indicates canceling the order operation.
	// 11. replace - specifies reissuing a certificate.
	// 12. downloadConfirmLetter - specifies downloading a certificate revocation confirmation letter.
	// 13. editRevokeLetter - specifies uploading a certificate revocation confirmation letter.
	// 14. renewVIP - specifies renewing a paid certificate.
	// 15. applyVIP - specifies applying for a paid certificate.
	// 16. submitInfo - specifies submitting documentation.
	// 17. downloadConfirmLetter - specifies downloading the confirmation letter template.
	// 18. uploadFromYunAPI - indicates uploading via the cloud api.
	// 19. transferIn - indicates the certificate transfer to operation.
	// 20. transferOut - indicates the certificate transfer operation.
	// 21. refund - indicates applying for a refund.
	// 22. multiYearsRenew - indicates multi-year auto-renewal.
	// 23. modifyDownloadLimit - indicates modifying the download limit switch.
	// 24. issued - indicates certificate issuance.
	// 25. domainValidationPassed - indicates domain verification completed.
	// 26. Resubmit - indicates reapplying for a certificate.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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

	// Type. `Original`: original certificate CSR; `Upload`: uploaded manually; `Online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR content, required when uploading manually.
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

	// Type. `Original`: original certificate CSR; `Upload`: uploaded manually; `Online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR content, required when uploading manually.
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
	// Cloud resource types, which support clb, waf, api gateway, cos, tke, tse, and tcb.
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
	// Paid certificate id of materials to be submitted.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// This field is required. Generation method of CSR, valid values are:
	// online: tencent cloud generates the CSR and private key based on the submitted parameter information and stores them encryptedly.
	// parse: generate the CSR and private key by itself, and apply for a certificate by uploading the CSR.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// The content of the uploaded csr.
	// If CsrType is parse, this field is required.
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// The common name bound to the certificate. if a CSR is uploaded, the domain name must be consistent with the common name resolved from the CSR.
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// Other domain names bound to the certificate. not required for single domain and wildcard domain certificates. required for multiple domain names and multiple wildcard domain names.
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// Private key password, which is currently only used for the password when generating jks and pfx format certificates; other formats of private key certificates are not encrypted.	
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// This field is required. Company name.
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// This field is required.  Department name.
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// This field is required. Company's detailed address.
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// This field is required.Country name such as CN.
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// This field is required, which specifies the city where the company is located.
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// This field is required, specifying the province where the company is located.
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// This field is required, the company's fixed-line phone area code.
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// This field is required, the company's landline number.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Certificate validation method. Validation types: DNS_AUTO = Automatic DNS validation (only supported for domains resolved by Tencent Cloud DNS with a normal resolution status), DNS = Manual DNS validation, FILE = File validation.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// This field is required, manager name.
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// This field is required, the manager's surname.
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// This field is required, the manager's mobile phone number.
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// This field is required, the manager's email address.
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// This field is required, the manager position.
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// This field is required, the contact person name.
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// This field is required, the contact person's surname.
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// This field is required, the contact person's email address.
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// This field is required, the contact person's mobile phone number.
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// This field is required, the contact person position.
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`

	// Indicates whether it is a dv certificate. default value is false.
	IsDV *bool `json:"IsDV,omitnil,omitempty" name:"IsDV"`
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Paid certificate id of materials to be submitted.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// This field is required. Generation method of CSR, valid values are:
	// online: tencent cloud generates the CSR and private key based on the submitted parameter information and stores them encryptedly.
	// parse: generate the CSR and private key by itself, and apply for a certificate by uploading the CSR.
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// The content of the uploaded csr.
	// If CsrType is parse, this field is required.
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// The common name bound to the certificate. if a CSR is uploaded, the domain name must be consistent with the common name resolved from the CSR.
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// Other domain names bound to the certificate. not required for single domain and wildcard domain certificates. required for multiple domain names and multiple wildcard domain names.
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// Private key password, which is currently only used for the password when generating jks and pfx format certificates; other formats of private key certificates are not encrypted.	
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// This field is required. Company name.
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// This field is required.  Department name.
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// This field is required. Company's detailed address.
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// This field is required.Country name such as CN.
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// This field is required, which specifies the city where the company is located.
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// This field is required, specifying the province where the company is located.
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// This field is required, the company's fixed-line phone area code.
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// This field is required, the company's landline number.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Certificate validation method. Validation types: DNS_AUTO = Automatic DNS validation (only supported for domains resolved by Tencent Cloud DNS with a normal resolution status), DNS = Manual DNS validation, FILE = File validation.
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// This field is required, manager name.
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// This field is required, the manager's surname.
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// This field is required, the manager's mobile phone number.
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// This field is required, the manager's email address.
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// This field is required, the manager position.
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// This field is required, the contact person name.
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// This field is required, the contact person's surname.
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// This field is required, the contact person's email address.
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// This field is required, the contact person's mobile phone number.
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// This field is required, the contact person position.
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`

	// Indicates whether it is a dv certificate. default value is false.
	IsDV *bool `json:"IsDV,omitnil,omitempty" name:"IsDV"`
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
	delete(f, "IsDV")
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
	// Whether the available format of nginx can be downloaded.
	NGINX *bool `json:"NGINX,omitnil,omitempty" name:"NGINX"`

	// Whether the available format of apache can be downloaded.
	APACHE *bool `json:"APACHE,omitnil,omitempty" name:"APACHE"`

	// Whether the available format of tomcat can be downloaded.
	TOMCAT *bool `json:"TOMCAT,omitnil,omitempty" name:"TOMCAT"`

	// Whether the available format of iis can be downloaded.
	IIS *bool `json:"IIS,omitnil,omitempty" name:"IIS"`

	// Indicates whether the jks format can be downloaded.
	JKS *bool `json:"JKS,omitnil,omitempty" name:"JKS"`

	// Indicates whether other formats can be downloaded.
	OTHER *bool `json:"OTHER,omitnil,omitempty" name:"OTHER"`

	// Indicates whether the root certificate can be downloaded.
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

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Domain status.
	// `Deployed`: deployed;.
	// `Processing`: deploying;.
	// `Applying`: applying;.
	// `Failed`: application failed;.
	// `Issued`: binding failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TeoInstanceList struct {
	// The list of EDGEONE instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of EDGEONE instances.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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
	// The old certificate id for one-click update. by querying the cloud resources bound to this certificate id, and then updating these cloud resources with the new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Resource types that need to be deployed, with optional parameter values (lowercase): clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb, tse, cos.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// New certificate id for one-click update. if this parameter is not provided, the public key certificate and private key certificate must be provided.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// List of regions that need to be deployed (deprecated)
	//
	// Deprecated: Regions is deprecated.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// List of regions where cloud resources need to be deployed. the cloud resource type of the supported region must be passed. valid values: clb, tke, apigateway, waf, tcb, tse, cos.
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// If a public key certificate is uploaded, the private key certificate must also be uploaded, and the CertificateId does not need to be transmitted.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// If a private key certificate is uploaded, then a public key certificate must be uploaded; CertificateId is not required.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Whether to ignore expiration reminder for old certificate  0: do not ignore the notification. 1: ignore the notification, ignore the expiration reminder of OldCertificateId.
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// It specifies whether the same certificate is allowed to be uploaded repeatedly. If the public key and private key certificates are selected for upload, this parameter can be configured. If there are duplicate certificates, the update task will fail.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// Whether to allow downloading. If you choose to upload a public/private key certificate, this parameter can be configured.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// Tag list. If you choose to upload a public/private key certificate, you can configure this parameter.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Project id. If you choose to upload a public/private key certificate, you can configure this parameter.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UpdateCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The old certificate id for one-click update. by querying the cloud resources bound to this certificate id, and then updating these cloud resources with the new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Resource types that need to be deployed, with optional parameter values (lowercase): clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb, tse, cos.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// New certificate id for one-click update. if this parameter is not provided, the public key certificate and private key certificate must be provided.
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// List of regions that need to be deployed (deprecated)
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// List of regions where cloud resources need to be deployed. the cloud resource type of the supported region must be passed. valid values: clb, tke, apigateway, waf, tcb, tse, cos.
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// If a public key certificate is uploaded, the private key certificate must also be uploaded, and the CertificateId does not need to be transmitted.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// If a private key certificate is uploaded, then a public key certificate must be uploaded; CertificateId is not required.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// Whether to ignore expiration reminder for old certificate  0: do not ignore the notification. 1: ignore the notification, ignore the expiration reminder of OldCertificateId.
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// It specifies whether the same certificate is allowed to be uploaded repeatedly. If the public key and private key certificates are selected for upload, this parameter can be configured. If there are duplicate certificates, the update task will fail.
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// Whether to allow downloading. If you choose to upload a public/private key certificate, this parameter can be configured.
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// Tag list. If you choose to upload a public/private key certificate, you can configure this parameter.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Project id. If you choose to upload a public/private key certificate, you can configure this parameter.
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
	// Task id, DeployRecordId of 0 indicates that the task is in progress. repeatedly requesting this api, when DeployRecordId returned is greater than 0, it indicates that the task is created successfully. if not created successfully, an exception will be thrown.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeployRecordId *uint64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Status of the task; 1 indicates successful creation; 0 indicates that there is a task being updated currently, and no new update task has been created; the returned value DeployRecordId is the task id being updated.
	DeployStatus *int64 `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`

	// Task Progress Details.
	// Note: this field may return null, indicating that no valid values can be obtained.
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
	// Record ID for pending retry deployment, which can be obtained through UpdateCertificateInstance. if this parameter is not provided, DeployRecordDetailId must be provided.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Detail ID for pending retry deployment record, which can be obtained through the DescribeHostUpdateRecordDetail api. if this parameter is not provided, DeployRecordId must be provided.
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

type UpdateCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// Record ID for pending retry deployment, which can be obtained through UpdateCertificateInstance. if this parameter is not provided, DeployRecordDetailId must be provided.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Detail ID for pending retry deployment record, which can be obtained through the DescribeHostUpdateRecordDetail api. if this parameter is not provided, DeployRecordId must be provided.
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
	// Update detail record id.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// New and old certificate update - new certificate id.
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// Old and new certificate update - old certificate id.
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// Deployment domain name list
	// Note: This field may return null, indicating that no valid value can be obtained.
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// Type of cloud resource for updating old and new certs.
	// - clb.
	// - cdn.
	// - ddos.
	// - live.
	// - vod.
	// - waf.
	// - apigateway.
	// - teo.
	// - tke.
	// - cos.
	// - tse.
	// - tcb.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Deployment region
	// Note: This field may return null, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Deployment status. valid values:.
	// 0: To be deployed.
	// 1: Deployment successful.
	// 2: Deployment failed.
	// 3: Deploying.
	// 4: Rollback succeeded.
	// 5: Rollback failure.
	// 6: No resource, no need for deployment.
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

	// Listener url (only for CLB).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type UpdateRecordDetails struct {
	// Type of cloud resource for updating old and new certs.
	// - clb.
	// - cdn.
	// - ddos.
	// - live.
	// - vod.
	// - waf.
	// - apigateway.
	// - teo.
	// - tke.
	// - cos.
	// - tse.
	// - tcb.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// The update details of the cloud resource.
	List []*UpdateRecordDetail `json:"List,omitnil,omitempty" name:"List"`

	// The update of the total number of cloud resources.
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
	// Resource type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// Region result list.
	// Note: this field may return null, indicating that no valid values can be obtained.
	UpdateSyncProgressRegions []*UpdateSyncProgressRegion `json:"UpdateSyncProgressRegions,omitnil,omitempty" name:"UpdateSyncProgressRegions"`

	// Asynchronous update progress status: 0, pending, 1 processed, 3 processing.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateSyncProgressRegion struct {
	// Resource type.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Total number
	// .
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Quantity of executions completed.
	// Note: this field may return null, indicating that no valid values can be obtained.
	OffsetCount *int64 `json:"OffsetCount,omitnil,omitempty" name:"OffsetCount"`

	// Asynchronous update progress status: 0, pending, 1 processed, 3 processing.
	// Note: this field may return null, indicating that no valid values can be obtained.
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

	// Certificate Usage/Source, such as CLB, CDN, WAF, LIVE, DDOS.
	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// The list of tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to allow duplicate upload of the same certificate, true: allow uploading certificates with the same fingerprint; false: do not allow uploading certificates with the same fingerprint. default value: true.
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

	// Certificate Usage/Source, such as CLB, CDN, WAF, LIVE, DDOS.
	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// The list of tags.
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Whether to allow duplicate upload of the same certificate, true: allow uploading certificates with the same fingerprint; false: do not allow uploading certificates with the same fingerprint. default value: true.
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

// Predefined struct for user
type UploadUpdateCertificateInstanceRequestParams struct {
	// Old Certificate ID for One-Click Update. Query the tencent cloud resources bound to this certificate ID and use the new certificate to update these resources.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Resource type that needs to be deployed, parameter value (lowercase): clb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Public key certificate.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key certificate.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// The list of regions where cloud resources need to be deployed. The cloud resource type supported in the region is required. Value: clb.
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`
}

type UploadUpdateCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Old Certificate ID for One-Click Update. Query the tencent cloud resources bound to this certificate ID and use the new certificate to update these resources.
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// Resource type that needs to be deployed, parameter value (lowercase): clb.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Public key certificate.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// Private key certificate.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// The list of regions where cloud resources need to be deployed. The cloud resource type supported in the region is required. Value: clb.
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`
}

func (r *UploadUpdateCertificateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldCertificateId")
	delete(f, "ResourceTypes")
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "ResourceTypesRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadUpdateCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadUpdateCertificateInstanceResponseParams struct {
	// The tencent cloud resource update task ID. DeployRecordId equals 0 indicates the task is in progress. Repeatedly request this API, and when DeployRecordId is greater than 0, it means the task has been successfully created. If not successfully created, an exception will be thrown.
	DeployRecordId *uint64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Update task creation status: 1 indicates successful creation; 0 indicates an existing ongoing update task, and no new update task was created. The return value DeployRecordId represents the ID of the ongoing update task.
	DeployStatus *int64 `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`

	// Create progress details for asynchronous update tasks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateSyncProgress []*UpdateSyncProgress `json:"UpdateSyncProgress,omitnil,omitempty" name:"UpdateSyncProgress"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadUpdateCertificateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UploadUpdateCertificateInstanceResponseParams `json:"Response"`
}

func (r *UploadUpdateCertificateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadUpdateCertificateRecordRetryRequestParams struct {
	// Retry deployment record ID, obtained through UpdateCertificateInstance to get the deployment record ID. If this parameter is not provided, DeployRecordDetailId must be provided.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Retry deployment record detail ID, obtained through the DescribeHostUpdateRecordDetail API. If this parameter is not provided, DeployRecordId must be provided.
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

type UploadUpdateCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// Retry deployment record ID, obtained through UpdateCertificateInstance to get the deployment record ID. If this parameter is not provided, DeployRecordDetailId must be provided.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// Retry deployment record detail ID, obtained through the DescribeHostUpdateRecordDetail API. If this parameter is not provided, DeployRecordId must be provided.
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

func (r *UploadUpdateCertificateRecordRetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateRecordRetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "DeployRecordDetailId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadUpdateCertificateRecordRetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadUpdateCertificateRecordRetryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadUpdateCertificateRecordRetryResponse struct {
	*tchttp.BaseResponse
	Response *UploadUpdateCertificateRecordRetryResponseParams `json:"Response"`
}

func (r *UploadUpdateCertificateRecordRetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateRecordRetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadUpdateCertificateRecordRollbackRequestParams struct {
	// Record ID of the certificate pending rollback, obtained through the UpdateCertificateInstance API.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

type UploadUpdateCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// Record ID of the certificate pending rollback, obtained through the UpdateCertificateInstance API.
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

func (r *UploadUpdateCertificateRecordRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateRecordRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadUpdateCertificateRecordRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadUpdateCertificateRecordRollbackResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadUpdateCertificateRecordRollbackResponse struct {
	*tchttp.BaseResponse
	Response *UploadUpdateCertificateRecordRollbackResponseParams `json:"Response"`
}

func (r *UploadUpdateCertificateRecordRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadUpdateCertificateRecordRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadUpdateRecordInfo struct {
	// Record ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Original certificate ID
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// Deployment resource type list.
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// Deployment state.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Deployment time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type VODInstanceList struct {
	// The list of VOD instances.	
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceList []*VodInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The total number of VOD instances in this region.	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
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

	// Whether to query exceptions.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}