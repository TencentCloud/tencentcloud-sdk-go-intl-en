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
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	DvAuthMethod *string `json:"DvAuthMethod,omitempty" name:"DvAuthMethod"`

	// Domain name
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Certificate type. Currently, the only supported value is 2, which indicates TrustAsia TLS RSA CA.
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// Email address
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// Mobile number
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// Validity period. The default value is 12 months, which is the only supported value currently.
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// Encryption algorithm. Only RSA is supported.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitempty" name:"CsrEncryptAlgo"`

	// Key pair parameter. Only the 2048-bit key pair is supported.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitempty" name:"CsrKeyParameter"`

	// CSR encryption password
	CsrKeyPassword *string `json:"CsrKeyPassword,omitempty" name:"CsrKeyPassword"`

	// Alias
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Original certificate ID, which is used to apply for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("ApplyCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelCertificateOrderRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CancelCertificateOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("CancelCertificateOrderRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelCertificateOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the certificate whose order has been successfully cancelled
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelCertificateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertificateExtra struct {

	// Number of domain names which can be associated with the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainNumber *string `json:"DomainNumber,omitempty" name:"DomainNumber"`

	// Original certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginCertificateId *string `json:"OriginCertificateId,omitempty" name:"OriginCertificateId"`

	// Original ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedBy *string `json:"ReplacedBy,omitempty" name:"ReplacedBy"`

	// New ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedFor *string `json:"ReplacedFor,omitempty" name:"ReplacedFor"`

	// Certificate ID of the new order
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewOrder *string `json:"RenewOrder,omitempty" name:"RenewOrder"`
}

type Certificates struct {

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Certificate source
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitempty" name:"From"`

	// Certificate plan type. `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

	// Primary domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Status value. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: OV/EV certificate, information to be submitted; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

	// Vulnerability scanning status. `INACTIVE`: not activated; `ACTIVE`: activated
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageTypeName *string `json:"PackageTypeName,omitempty" name:"PackageTypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

	// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

	// Project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitempty" name:"ProjectInfo"`

	// Associated Tencent Cloud services. Currently, this parameter is unavailable.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BoundResource []*string `json:"BoundResource,omitempty" name:"BoundResource" list`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitempty" name:"Tags" list`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CommitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("CommitCertificateInformationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CommitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TrustAsia order ID
		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

		// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DeleteCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Deletion result
		DeleteResult *bool `json:"DeleteResult,omitempty" name:"DeleteResult"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DescribeCertificateDetailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
		OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

		// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

		// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
		From *string `json:"From,omitempty" name:"From"`

		// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

		// Certificate plan type. `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: this field may return null, indicating that no valid values can be obtained.
		PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

		// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

		// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
		Alias *string `json:"Alias,omitempty" name:"Alias"`

		// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
		StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

		// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
		VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

		// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
		VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

		// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

		// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

		// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
		ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

		// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
		InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

		// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

		// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

		// Private key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

		// Public key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

		// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
		DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitempty" name:"DvAuthDetail"`

		// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
		VulnerabilityReport *string `json:"VulnerabilityReport,omitempty" name:"VulnerabilityReport"`

		// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
		TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

		// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
		StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

		// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

		// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

		// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

		// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

		// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

		// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubmittedData *SubmittedData `json:"SubmittedData,omitempty" name:"SubmittedData"`

		// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
		RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

		// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
		Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

		// List of associated tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Tags []*Tags `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateOperateLogsRequest struct {
	*tchttp.BaseRequest

	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of requested logs. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Start time. The default value is 15 days ago.
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// End time. The default value is the current time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeCertificateOperateLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeCertificateOperateLogsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of logs that meet query conditions
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// Number of logs returned for this request
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Certificate operation log list
	// Note: this field may return null, indicating that no valid values can be obtained.
		OperateLogs []*OperationLog `json:"OperateLogs,omitempty" name:"OperateLogs" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DescribeCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
		OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

		// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

		// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
		From *string `json:"From,omitempty" name:"From"`

		// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

		// Certificate plan type. `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: this field may return null, indicating that no valid values can be obtained.
		PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

		// Name of the certificate issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
		ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

		// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
		Domain *string `json:"Domain,omitempty" name:"Domain"`

		// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
		Alias *string `json:"Alias,omitempty" name:"Alias"`

		// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
		StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

		// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
		VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

		// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
		VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

		// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

		// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

		// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
		ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

		// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
		InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

		// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

		// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

		// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
		DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitempty" name:"DvAuthDetail"`

		// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
		VulnerabilityReport *string `json:"VulnerabilityReport,omitempty" name:"VulnerabilityReport"`

		// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
		PackageTypeName *string `json:"PackageTypeName,omitempty" name:"PackageTypeName"`

		// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
		StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

		// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

		// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

		// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

		// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

		// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
		IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

		// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
		RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

		// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubmittedData *SubmittedData `json:"SubmittedData,omitempty" name:"SubmittedData"`

		// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
		Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

		// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Tags []*Tags `json:"Tags,omitempty" name:"Tags" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest

	// Pagination offset, starting from 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of certificates on each page. The default value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Keyword for search, which can be a certificate ID, alias, or domain name, for example, a8xHcaIs
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Sorting by expiration time. `DESC`: descending; `ASC`: ascending
	ExpirationSort *string `json:"ExpirationSort,omitempty" name:"ExpirationSort"`

	// Certificate status
	CertificateStatus []*uint64 `json:"CertificateStatus,omitempty" name:"CertificateStatus" list`

	// Whether the certificate can be deployed. `1`: yes; `0`: no
	Deployable *uint64 `json:"Deployable,omitempty" name:"Deployable"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("DescribeCertificatesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number
	// Note: this field may return null, indicating that no valid values can be obtained.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List
	// Note: this field may return null, indicating that no valid values can be obtained.
		Certificates []*Certificates `json:"Certificates,omitempty" name:"Certificates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadCertificateRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DownloadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return errors.New("DownloadCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ZIP content encoded by using Base64. After the content is decoded by using Base64, it can be saved as a ZIP file.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Content *string `json:"Content,omitempty" name:"Content"`

		// MIME type. `application/zip`: ZIP file
	// Note: this field may return null, indicating that no valid values can be obtained.
		ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DvAuthDetail struct {

	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitempty" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitempty" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitempty" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitempty" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitempty" name:"DvAuthKeySubDomain"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuths []*DvAuths `json:"DvAuths,omitempty" name:"DvAuths" list`
}

type DvAuths struct {

	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitempty" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitempty" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitempty" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitempty" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitempty" name:"DvAuthSubDomain"`

	// DV authentication type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthVerifyType *string `json:"DvAuthVerifyType,omitempty" name:"DvAuthVerifyType"`
}

type ModifyCertificateAliasRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Alias
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyCertificateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Alias")
	if len(f) > 0 {
		return errors.New("ModifyCertificateAliasRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ID of the successfully modified certificate
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest

	// ID list of certificates whose projects need to be modified. A maximum of 100 certificate IDs are supported.
	CertificateIdList []*string `json:"CertificateIdList,omitempty" name:"CertificateIdList" list`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyCertificateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return errors.New("ModifyCertificateProjectRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of certificates whose projects were modified successfully
	// Note: this field may return null, indicating that no valid values can be obtained.
		SuccessCertificates []*string `json:"SuccessCertificates,omitempty" name:"SuccessCertificates" list`

		// List of certificates whose projects failed to be modified
	// Note: this field may return null, indicating that no valid values can be obtained.
		FailCertificates []*string `json:"FailCertificates,omitempty" name:"FailCertificates" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {

	// Action performed on logs
	Action *string `json:"Action,omitempty" name:"Action"`

	// Time when the action is performed
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type ProjectInfo struct {

	// Project name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// UIN of the project creator
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreatorUin *uint64 `json:"ProjectCreatorUin,omitempty" name:"ProjectCreatorUin"`

	// Project creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreateTime *string `json:"ProjectCreateTime,omitempty" name:"ProjectCreateTime"`

	// Brief project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectResume *string `json:"ProjectResume,omitempty" name:"ProjectResume"`

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ReplaceCertificateRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	ValidType *string `json:"ValidType,omitempty" name:"ValidType"`

	// Type. `original`: original certificate CSR; `upload`: uploaded manually; `online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// CSR content
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// Password of the key
	CsrkeyPassword *string `json:"CsrkeyPassword,omitempty" name:"CsrkeyPassword"`

	// Reissue reason
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *ReplaceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("ReplaceCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest

	// Certificate ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// CSR generation mode. `online`: generated online; `parse`: uploaded manually
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// Uploaded CSR content
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// Domain name bound with the certificate
	CertificateDomain *string `json:"CertificateDomain,omitempty" name:"CertificateDomain"`

	// Uploaded domain name array (can be uploaded for a multi-domain certificate)
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`

	// Password of the private key
	KeyPassword *string `json:"KeyPassword,omitempty" name:"KeyPassword"`

	// Organization name
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// Division name
	OrganizationDivision *string `json:"OrganizationDivision,omitempty" name:"OrganizationDivision"`

	// Detailed address of the organization
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" name:"OrganizationAddress"`

	// Country where the organization is located, for example, CN (China)
	OrganizationCountry *string `json:"OrganizationCountry,omitempty" name:"OrganizationCountry"`

	// City where the organization is located
	OrganizationCity *string `json:"OrganizationCity,omitempty" name:"OrganizationCity"`

	// Province where the organization is located
	OrganizationRegion *string `json:"OrganizationRegion,omitempty" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`

	// Area code of the fixed-line phone number of the organization
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty" name:"PhoneAreaCode"`

	// Fixed-line phone number of the organization
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Certificate validation method
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// Last name of the administrator
	AdminFirstName *string `json:"AdminFirstName,omitempty" name:"AdminFirstName"`

	// First name of the administrator
	AdminLastName *string `json:"AdminLastName,omitempty" name:"AdminLastName"`

	// Mobile number of the administrator
	AdminPhoneNum *string `json:"AdminPhoneNum,omitempty" name:"AdminPhoneNum"`

	// Email of the administrator
	AdminEmail *string `json:"AdminEmail,omitempty" name:"AdminEmail"`

	// Position of the administrator
	AdminPosition *string `json:"AdminPosition,omitempty" name:"AdminPosition"`

	// Last name of the contact
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// First name of the contact
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// Email of the contact
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// Mobile number of the contact
	ContactNumber *string `json:"ContactNumber,omitempty" name:"ContactNumber"`

	// Position of the contact
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`
}

func (r *SubmitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SubmitCertificateInformationRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SubmitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmittedData struct {

	// CSR type. `online`: CSR generated online; `parse`: CSR pasted
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// CSR content
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// Domain name information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateDomain *string `json:"CertificateDomain,omitempty" name:"CertificateDomain"`

	// DNS information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`

	// Password of the private key
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyPassword *string `json:"KeyPassword,omitempty" name:"KeyPassword"`

	// Enterprise or unit name
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// Division
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationDivision *string `json:"OrganizationDivision,omitempty" name:"OrganizationDivision"`

	// Address
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" name:"OrganizationAddress"`

	// Country
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCountry *string `json:"OrganizationCountry,omitempty" name:"OrganizationCountry"`

	// City
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCity *string `json:"OrganizationCity,omitempty" name:"OrganizationCity"`

	// Province
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationRegion *string `json:"OrganizationRegion,omitempty" name:"OrganizationRegion"`

	// Postal code
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`

	// Area code of the fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty" name:"PhoneAreaCode"`

	// Fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// First name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminFirstName *string `json:"AdminFirstName,omitempty" name:"AdminFirstName"`

	// Last name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminLastName *string `json:"AdminLastName,omitempty" name:"AdminLastName"`

	// Phone number of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPhoneNum *string `json:"AdminPhoneNum,omitempty" name:"AdminPhoneNum"`

	// Email of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminEmail *string `json:"AdminEmail,omitempty" name:"AdminEmail"`

	// Position of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPosition *string `json:"AdminPosition,omitempty" name:"AdminPosition"`

	// First name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// Last name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// Phone number of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactNumber *string `json:"ContactNumber,omitempty" name:"ContactNumber"`

	// Email of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// Position of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`

	// Validation type
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

type Tags struct {

	// Tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UploadCertificateRequest struct {
	*tchttp.BaseRequest

	// Public key of the certificate
	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

	// Private key content. This parameter is required when the certificate type is SVR, and not required when the certificate type is CA.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate. The default value is SVR.
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// Alias
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 
	CertificateUse *string `json:"CertificateUse,omitempty" name:"CertificateUse"`
}

func (r *UploadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("UploadCertificateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Certificate ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
