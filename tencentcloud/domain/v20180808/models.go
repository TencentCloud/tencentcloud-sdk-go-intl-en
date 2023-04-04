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

package v20180808

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AdminContact struct {
	// The first name.
	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

	// The last name.
	LastName *string `json:"LastName,omitempty" name:"LastName"`

	// The country or region name, such as `CN`.
	Country *string `json:"Country,omitempty" name:"Country"`

	// The province or state name.
	Province *string `json:"Province,omitempty" name:"Province"`

	// The city name.
	City *string `json:"City,omitempty" name:"City"`

	// The address line 1.
	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`

	// The zip code.
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// The email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// The mobile number, such as `+86.13600000000`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// The company or organization name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// The job title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobTitle *string `json:"JobTitle,omitempty" name:"JobTitle"`

	// The address line 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressLineTwo *string `json:"AddressLineTwo,omitempty" name:"AddressLineTwo"`

	// The fax number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fax *string `json:"Fax,omitempty" name:"Fax"`
}

type BatchDomainBuyDetails struct {
	// The details ID.
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// The bulk operation type. Valid values: `new` (register domains), `batch_transfer_prohibition_on` (enable transfer prohibition), `batch_transfer_prohibition_off` (disable transfer prohibition), `batch_update_prohibition_on` (enable update prohibition), `batch_update_prohibition_off` (disable update prohibition).
	Action *string `json:"Action,omitempty" name:"Action"`

	// The domains.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// The status. Valid values: `SUCCESS`, `FAILURE`
	Status *string `json:"Status,omitempty" name:"Status"`

	// The reason for failure.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// The creation time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The update time.
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// Null: The DNS service is not transferred. `false`: The DNS service failed to be transferred. `true`: The DNS service is transferred successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransferDnsResult *bool `json:"TransferDnsResult,omitempty" name:"TransferDnsResult"`

	// The reason for failure, expressed in Chinese.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReasonZh *string `json:"ReasonZh,omitempty" name:"ReasonZh"`

	// The payment status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayStatus *int64 `json:"PayStatus,omitempty" name:"PayStatus"`
}

type BatchDomainBuyLog struct {
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The bulk operation type. Valid values: `new` (register domains), `batch_transfer_prohibition_on` (enable transfer prohibition), `batch_transfer_prohibition_off` (disable transfer prohibition), `batch_update_prohibition_on` (enable update prohibition), `batch_update_prohibition_off` (disable update prohibition).
	Action *string `json:"Action,omitempty" name:"Action"`

	// The quantity.
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// The execution status. Valid values: `doing`, `done`
	Status *string `json:"Status,omitempty" name:"Status"`

	// The submission time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

// Predefined struct for user
type BatchModifyIntlDomainDNSRequestParams struct {
	// The target domains.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The domain DNS array.
	Dns []*string `json:"Dns,omitempty" name:"Dns"`

	// Valid values: `batch_modify_domain_dns1`, `batch_modify_domain_dns2`, `batch_modify_domain_dns3`
	BatchAction *string `json:"BatchAction,omitempty" name:"BatchAction"`
}

type BatchModifyIntlDomainDNSRequest struct {
	*tchttp.BaseRequest
	
	// The target domains.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The domain DNS array.
	Dns []*string `json:"Dns,omitempty" name:"Dns"`

	// Valid values: `batch_modify_domain_dns1`, `batch_modify_domain_dns2`, `batch_modify_domain_dns3`
	BatchAction *string `json:"BatchAction,omitempty" name:"BatchAction"`
}

func (r *BatchModifyIntlDomainDNSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyIntlDomainDNSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Dns")
	delete(f, "BatchAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyIntlDomainDNSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyIntlDomainDNSResponseParams struct {
	// The log ID.
	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyIntlDomainDNSResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyIntlDomainDNSResponseParams `json:"Response"`
}

func (r *BatchModifyIntlDomainDNSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyIntlDomainDNSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyIntlDomainInfoRequestParams struct {
	// The domains whose information is to be modified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Whether to enable the 60-day inter-registrar transfer lock.
	LockTransfer *bool `json:"LockTransfer,omitempty" name:"LockTransfer"`
}

type BatchModifyIntlDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// The domains whose information is to be modified.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Whether to enable the 60-day inter-registrar transfer lock.
	LockTransfer *bool `json:"LockTransfer,omitempty" name:"LockTransfer"`
}

func (r *BatchModifyIntlDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyIntlDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "TemplateId")
	delete(f, "LockTransfer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyIntlDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyIntlDomainInfoResponseParams struct {
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyIntlDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyIntlDomainInfoResponseParams `json:"Response"`
}

func (r *BatchModifyIntlDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyIntlDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BillingContact struct {
	// The first name.
	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

	// The last name.
	LastName *string `json:"LastName,omitempty" name:"LastName"`

	// The country or region name, such as `CN`.
	Country *string `json:"Country,omitempty" name:"Country"`

	// The province or state name.
	Province *string `json:"Province,omitempty" name:"Province"`

	// The city name.
	City *string `json:"City,omitempty" name:"City"`

	// The address line 1.
	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`

	// The zip code.
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// The email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// The mobile number, such as `+86.13600000000`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// The company or organization name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// The job title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobTitle *string `json:"JobTitle,omitempty" name:"JobTitle"`

	// The address line 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressLineTwo *string `json:"AddressLineTwo,omitempty" name:"AddressLineTwo"`

	// The fax number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fax *string `json:"Fax,omitempty" name:"Fax"`
}

// Predefined struct for user
type CheckIntlDomainNewRequestParams struct {
	// The name of the domain to be checked.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Period, in years. If this parameter is left empty, premium domains cannot be queried.
	Period *string `json:"Period,omitempty" name:"Period"`
}

type CheckIntlDomainNewRequest struct {
	*tchttp.BaseRequest
	
	// The name of the domain to be checked.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Period, in years. If this parameter is left empty, premium domains cannot be queried.
	Period *string `json:"Period,omitempty" name:"Period"`
}

func (r *CheckIntlDomainNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntlDomainNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIntlDomainNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIntlDomainNewResponseParams struct {
	// The name of the domain checked.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// Whether the domain is available for registration.
	Available *bool `json:"Available,omitempty" name:"Available"`

	// The reason why the domain cannot be registered.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// Whether the domain is a premium one.
	Premium *bool `json:"Premium,omitempty" name:"Premium"`

	// The domain price.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// Whether the domain name involves restricted words.
	BlackWord *bool `json:"BlackWord,omitempty" name:"BlackWord"`

	// The premium domain description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// The price for renewing the premium domain.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeeRenew *float64 `json:"FeeRenew,omitempty" name:"FeeRenew"`

	// The real price of the domain. For a premium domain, its price varies depending on the period. For a non-premium domain, the price is the 1-year price.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealPrice *float64 `json:"RealPrice,omitempty" name:"RealPrice"`

	// The price for transferring a premium domain in.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeeTransfer *float64 `json:"FeeTransfer,omitempty" name:"FeeTransfer"`

	// The price for redeeming a premium domain.
	FeeRestore *float64 `json:"FeeRestore,omitempty" name:"FeeRestore"`

	// The period (in years) of the domain.
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// The reason why the domain cannot be registered, expressed in Chinese.
	ReasonZh *string `json:"ReasonZh,omitempty" name:"ReasonZh"`

	// The internal error code.
	ResCode *string `json:"ResCode,omitempty" name:"ResCode"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckIntlDomainNewResponse struct {
	*tchttp.BaseResponse
	Response *CheckIntlDomainNewResponseParams `json:"Response"`
}

func (r *CheckIntlDomainNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntlDomainNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlDomainBatchRequestParams struct {
	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The purchase period (years) of the domain. Value range: [1-10]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// The domains (maximum 4,000) to purchase.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The payment method. Valid values: `0` (online payment), `1` (account balance), `2` (package), `3` (offline settlement).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to enable the transfer prohibition lock.
	TransferProhibition *bool `json:"TransferProhibition,omitempty" name:"TransferProhibition"`

	// Whether to enable the update prohibition lock.
	UpdateProhibition *bool `json:"UpdateProhibition,omitempty" name:"UpdateProhibition"`

	// The custom DNS servers
	CustomDns []*string `json:"CustomDns,omitempty" name:"CustomDns"`
}

type CreateIntlDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The purchase period (years) of the domain. Value range: [1-10]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// The domains (maximum 4,000) to purchase.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The payment method. Valid values: `0` (online payment), `1` (account balance), `2` (package), `3` (offline settlement).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to enable the transfer prohibition lock.
	TransferProhibition *bool `json:"TransferProhibition,omitempty" name:"TransferProhibition"`

	// Whether to enable the update prohibition lock.
	UpdateProhibition *bool `json:"UpdateProhibition,omitempty" name:"UpdateProhibition"`

	// The custom DNS servers
	CustomDns []*string `json:"CustomDns,omitempty" name:"CustomDns"`
}

func (r *CreateIntlDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Period")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "TransferProhibition")
	delete(f, "UpdateProhibition")
	delete(f, "CustomDns")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntlDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlDomainBatchResponseParams struct {
	// The bulk purchase log ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIntlDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntlDomainBatchResponseParams `json:"Response"`
}

func (r *CreateIntlDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlPhoneEmailRequestParams struct {
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`

	// The verification code.
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type CreateIntlPhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`

	// The verification code.
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

func (r *CreateIntlPhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlPhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Code")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntlPhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlPhoneEmailResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIntlPhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntlPhoneEmailResponseParams `json:"Response"`
}

func (r *CreateIntlPhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlPhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlTemplateRequestParams struct {
	// The registrant contact.
	RegistrantContact *RegistrantContact `json:"RegistrantContact,omitempty" name:"RegistrantContact"`

	// The admin contact.
	AdminContact *AdminContact `json:"AdminContact,omitempty" name:"AdminContact"`

	// The technical contact.
	TechnicalContact *TechnicalContact `json:"TechnicalContact,omitempty" name:"TechnicalContact"`

	// The contact person for bills.
	BillingContact *BillingContact `json:"BillingContact,omitempty" name:"BillingContact"`

	// The profile type. Valid values: `I` (individual, default), `E` (organization).
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

type CreateIntlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The registrant contact.
	RegistrantContact *RegistrantContact `json:"RegistrantContact,omitempty" name:"RegistrantContact"`

	// The admin contact.
	AdminContact *AdminContact `json:"AdminContact,omitempty" name:"AdminContact"`

	// The technical contact.
	TechnicalContact *TechnicalContact `json:"TechnicalContact,omitempty" name:"TechnicalContact"`

	// The contact person for bills.
	BillingContact *BillingContact `json:"BillingContact,omitempty" name:"BillingContact"`

	// The profile type. Valid values: `I` (individual, default), `E` (organization).
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

func (r *CreateIntlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegistrantContact")
	delete(f, "AdminContact")
	delete(f, "TechnicalContact")
	delete(f, "BillingContact")
	delete(f, "TemplateType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntlTemplateResponseParams struct {
	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIntlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntlTemplateResponseParams `json:"Response"`
}

func (r *CreateIntlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntlPhoneEmailRequestParams struct {
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`
}

type DeleteIntlPhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *DeleteIntlPhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntlPhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntlPhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntlPhoneEmailResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIntlPhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntlPhoneEmailResponseParams `json:"Response"`
}

func (r *DeleteIntlPhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntlPhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntlTemplateRequestParams struct {
	// The unique ID of the target registrant profile.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DeleteIntlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of the target registrant profile.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteIntlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntlTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIntlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntlTemplateResponseParams `json:"Response"`
}

func (r *DeleteIntlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlBatchDetailStatusRequestParams struct {
	// The IDs of the logs to be queried.
	LogIds []*int64 `json:"LogIds,omitempty" name:"LogIds"`
}

type DescribeIntlBatchDetailStatusRequest struct {
	*tchttp.BaseRequest
	
	// The IDs of the logs to be queried.
	LogIds []*int64 `json:"LogIds,omitempty" name:"LogIds"`
}

func (r *DescribeIntlBatchDetailStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlBatchDetailStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlBatchDetailStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlBatchDetailStatusResponseParams struct {
	// The details.
	Details []*IntlBatchDetails `json:"Details,omitempty" name:"Details"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlBatchDetailStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlBatchDetailStatusResponseParams `json:"Response"`
}

func (r *DescribeIntlBatchDetailStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlBatchDetailStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlBatchOperationLogsRequestParams struct {
	// The offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIntlBatchOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// The offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIntlBatchOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlBatchOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlBatchOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlBatchOperationLogsResponseParams struct {
	// The total count.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The log list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DomainBatchLogSet []*BatchDomainBuyLog `json:"DomainBatchLogSet,omitempty" name:"DomainBatchLogSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlBatchOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlBatchOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeIntlBatchOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlBatchOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainBatchDetailsRequestParams struct {
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The sort key. Valid values: `Domain`, `UpdateOn`, `Status`
	OrderByKey *string `json:"OrderByKey,omitempty" name:"OrderByKey"`

	// Valid values: `0` (ascending), `1` (descending).
	OrderBy *int64 `json:"OrderBy,omitempty" name:"OrderBy"`
}

type DescribeIntlDomainBatchDetailsRequest struct {
	*tchttp.BaseRequest
	
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The offset. Default value: 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of returned results. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The sort key. Valid values: `Domain`, `UpdateOn`, `Status`
	OrderByKey *string `json:"OrderByKey,omitempty" name:"OrderByKey"`

	// Valid values: `0` (ascending), `1` (descending).
	OrderBy *int64 `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeIntlDomainBatchDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainBatchDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByKey")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlDomainBatchDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainBatchDetailsResponseParams struct {
	// The total count.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The list of log details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DomainBatchDetailSet []*BatchDomainBuyDetails `json:"DomainBatchDetailSet,omitempty" name:"DomainBatchDetailSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlDomainBatchDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlDomainBatchDetailsResponseParams `json:"Response"`
}

func (r *DescribeIntlDomainBatchDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainBatchDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainListRequestParams struct {
	// The page number in pagination cases.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of records on each page in pagination cases.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The domain keyword for fuzzy search.
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// The registration time sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByRegTime *int64 `json:"OrderByRegTime,omitempty" name:"OrderByRegTime"`

	// The expiration time sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByExpireTime *int64 `json:"OrderByExpireTime,omitempty" name:"OrderByExpireTime"`

	// The domain status.
	// `all`: All domains.
	// `UrgentNeedRenew`: The domains that are in urgent need of renewal.
	// `RedemptionPending`: The domains that are in urgent need of redemption.
	// `nosubmit`: The domains with unverified identities.
	// `tran`: The domains that are being transferred in.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The DNS type. Valid values: `1` (DNSPod), `2` (others).
	DnsStatus *uint64 `json:"DnsStatus,omitempty" name:"DnsStatus"`

	// The domain sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByDomainName *uint64 `json:"OrderByDomainName,omitempty" name:"OrderByDomainName"`
}

type DescribeIntlDomainListRequest struct {
	*tchttp.BaseRequest
	
	// The page number in pagination cases.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// The number of records on each page in pagination cases.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The domain keyword for fuzzy search.
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// The registration time sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByRegTime *int64 `json:"OrderByRegTime,omitempty" name:"OrderByRegTime"`

	// The expiration time sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByExpireTime *int64 `json:"OrderByExpireTime,omitempty" name:"OrderByExpireTime"`

	// The domain status.
	// `all`: All domains.
	// `UrgentNeedRenew`: The domains that are in urgent need of renewal.
	// `RedemptionPending`: The domains that are in urgent need of redemption.
	// `nosubmit`: The domains with unverified identities.
	// `tran`: The domains that are being transferred in.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The DNS type. Valid values: `1` (DNSPod), `2` (others).
	DnsStatus *uint64 `json:"DnsStatus,omitempty" name:"DnsStatus"`

	// The domain sort order. Valid values: `1` (descending), `2` (ascending).
	OrderByDomainName *uint64 `json:"OrderByDomainName,omitempty" name:"OrderByDomainName"`
}

func (r *DescribeIntlDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "KeyWord")
	delete(f, "OrderByRegTime")
	delete(f, "OrderByExpireTime")
	delete(f, "Status")
	delete(f, "DnsStatus")
	delete(f, "OrderByDomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainListResponseParams struct {
	// The domain information set.
	DomainSet []*IntlDomainInfo `json:"DomainSet,omitempty" name:"DomainSet"`

	// The total number of domains.
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlDomainListResponseParams `json:"Response"`
}

func (r *DescribeIntlDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainPriceNewListRequestParams struct {
	// The list of domain suffixes for which you want to query prices. This parameter defaults to all suffixes.
	TldList []*string `json:"TldList,omitempty" name:"TldList"`

	// The purchase year of the domains for which you want to query prices. This parameter defaults to all years.
	Year []*int64 `json:"Year,omitempty" name:"Year"`

	// The domain purchase type. Valid values: `new`, `renew`, `redem` (redeem), `tran` (transfer in).
	Operation []*string `json:"Operation,omitempty" name:"Operation"`
}

type DescribeIntlDomainPriceNewListRequest struct {
	*tchttp.BaseRequest
	
	// The list of domain suffixes for which you want to query prices. This parameter defaults to all suffixes.
	TldList []*string `json:"TldList,omitempty" name:"TldList"`

	// The purchase year of the domains for which you want to query prices. This parameter defaults to all years.
	Year []*int64 `json:"Year,omitempty" name:"Year"`

	// The domain purchase type. Valid values: `new`, `renew`, `redem` (redeem), `tran` (transfer in).
	Operation []*string `json:"Operation,omitempty" name:"Operation"`
}

func (r *DescribeIntlDomainPriceNewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainPriceNewListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TldList")
	delete(f, "Year")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlDomainPriceNewListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainPriceNewListResponseParams struct {
	// The price list of domains.
	PriceList []*PriceInfoNew `json:"PriceList,omitempty" name:"PriceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlDomainPriceNewListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlDomainPriceNewListResponseParams `json:"Response"`
}

func (r *DescribeIntlDomainPriceNewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainPriceNewListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainRequestParams struct {
	// The domain ID.
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeIntlDomainRequest struct {
	*tchttp.BaseRequest
	
	// The domain ID.
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeIntlDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlDomainResponseParams struct {
	// The domain information.
	DomainInfo *IntlDomainInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlDomainResponseParams `json:"Response"`
}

func (r *DescribeIntlDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlPhoneEmailListRequestParams struct {
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The number of records on each page in pagination cases.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number in pagination cases.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeIntlPhoneEmailListRequest struct {
	*tchttp.BaseRequest
	
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The number of records on each page in pagination cases.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number in pagination cases.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeIntlPhoneEmailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlPhoneEmailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlPhoneEmailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlPhoneEmailListResponseParams struct {
	// The list of verified mobile numbers and email addresses.
	PhoneEmailList []*IntlPhoneEmailLists `json:"PhoneEmailList,omitempty" name:"PhoneEmailList"`

	// The total count.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlPhoneEmailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlPhoneEmailListResponseParams `json:"Response"`
}

func (r *DescribeIntlPhoneEmailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlPhoneEmailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlTemplateListRequestParams struct {
	// The offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The maximum number of entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The domain registrant keyword for exact match.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// The type. Valid values: `all` (all types), `I` (individual), `E` (organization).
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeIntlTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// The offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The maximum number of entries.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The domain registrant keyword for exact match.
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// The type. Valid values: `all` (all types), `I` (individual), `E` (organization).
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeIntlTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Keyword")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlTemplateListResponseParams struct {
	// The registrant profile list information.
	TemplateSet []*IntlTemplate `json:"TemplateSet,omitempty" name:"TemplateSet"`

	// The total count.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlTemplateListResponseParams `json:"Response"`
}

func (r *DescribeIntlTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlTemplateRequestParams struct {
	// The unique ID of a registrant profile.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

type DescribeIntlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a registrant profile.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DescribeIntlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntlTemplateResponseParams struct {
	// The details of the registrant profile.
	Template *IntlTemplateInfo `json:"Template,omitempty" name:"Template"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntlTemplateResponseParams `json:"Response"`
}

func (r *DescribeIntlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntlBatchDetails struct {
	// The ID of the bulk task.
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// The task status.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The reason.
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// The reason, expressed in Chinese.
	ReasonZh *string `json:"ReasonZh,omitempty" name:"ReasonZh"`
}

type IntlContactInfo struct {
	// The city name.
	City *string `json:"City,omitempty" name:"City"`

	// The country or region name.
	Country *string `json:"Country,omitempty" name:"Country"`

	// The email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// The domain registrant.
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// The province or state name.
	Province *string `json:"Province,omitempty" name:"Province"`

	// The name of the registrant.
	RegistrantName *string `json:"RegistrantName,omitempty" name:"RegistrantName"`

	// The registrant type. Valid values: `I` (individual), `E` (organization).
	RegistrantType *string `json:"RegistrantType,omitempty" name:"RegistrantType"`

	// The detailed address.
	Street *string `json:"Street,omitempty" name:"Street"`

	// The mobile number.
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// The zip code.
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// The first name.
	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

	// The last name.
	LastName *string `json:"LastName,omitempty" name:"LastName"`

	// The company name.
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`
}

type IntlDomainInfo struct {
	// The auto-renewal flag. Valid values: `0` (disabled by default), `1` (enabled), `2` (disabled).
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// The registration time.
	CreationDate *string `json:"CreationDate,omitempty" name:"CreationDate"`

	// The domain ID.
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// The DNS status. Valid values: `1` (DNSPod), `2` (others), `3` (unknown).
	DnsStatus *int64 `json:"DnsStatus,omitempty" name:"DnsStatus"`

	// The domains.
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// The domain status.
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus"`

	// The purchase status of the domain. Valid values:
	// `ok`: Normal.
	// `RegisterPending`: Pending registration.
	// `RegisterDoing`: Registration in progress.
	// `RegisterFailed`: Registration failed.
	// `RenewPending`: Renewal grace period.
	// `RenewDoing`: Renewing.
	// `RedemptionPending`: Redemption period.
	// `AboutToExpire`: About to expire.
	// `TransferPending`: Pending transfer.
	// `TransferTransing`: Transfer in progress.
	// `TransferFailed`: Transfer failed.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The expiration date.
	ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// The auto-renewal flag. Valid values: `1` (enabled), `2` (disabled), `3` (disabled by default).
	ExpireMessage *int64 `json:"ExpireMessage,omitempty" name:"ExpireMessage"`

	// Whether the domain is a premium one.
	IsPremium *bool `json:"IsPremium,omitempty" name:"IsPremium"`

	// The DNS server of the domain.
	Dns []*string `json:"Dns,omitempty" name:"Dns"`

	// The contact information.
	ContactInfo *IntlContactInfo `json:"ContactInfo,omitempty" name:"ContactInfo"`

	// The number of years available for renewal. The value `0` indicates that the domain has reached its maximum validity period (10 years) and cannot be renewed.
	CanRenewYears *int64 `json:"CanRenewYears,omitempty" name:"CanRenewYears"`

	// The registrar type. If the value is `epp`, the registration time on the page is followed by (UTC). Otherwise, it is followed by (UTC+8).
	RegistrarType *string `json:"RegistrarType,omitempty" name:"RegistrarType"`

	// The account to which the domain belongs.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Whether DNSSEC is supported.
	SupportDnssec *bool `json:"SupportDnssec,omitempty" name:"SupportDnssec"`

	// WHOIS privacy service status. Valid values: `1` (enabled), `2` (disabled), `3` (enabling), `4` (disabling).
	WhoisPrivacy *int64 `json:"WhoisPrivacy,omitempty" name:"WhoisPrivacy"`

	// Valid values: `NotModify` (not modified), `Modifying`, `Failed` (failed to modify)
	ModifyStatus *string `json:"ModifyStatus,omitempty" name:"ModifyStatus"`

	// Valid values: `NotModify` (not modified), `Modifying`, `Failed` (failed to modify)
	DnsModifyStatus *string `json:"DnsModifyStatus,omitempty" name:"DnsModifyStatus"`
}

type IntlPhoneEmailLists struct {
	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`

	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The verification time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type IntlTemplate struct {
	// The registrant contact.
	RegistrantContact *RegistrantContact `json:"RegistrantContact,omitempty" name:"RegistrantContact"`

	// The admin contact.
	AdminContact *AdminContact `json:"AdminContact,omitempty" name:"AdminContact"`

	// The technical contact.
	TechnicalContact *TechnicalContact `json:"TechnicalContact,omitempty" name:"TechnicalContact"`

	// The contact person for bills.
	BillingContact *BillingContact `json:"BillingContact,omitempty" name:"BillingContact"`

	// The creation time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// Whether the profile is the default one.
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`

	// The last update time.
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
}

type IntlTemplateInfo struct {
	// The registrant contact.
	RegistrantContact *RegistrantContact `json:"RegistrantContact,omitempty" name:"RegistrantContact"`

	// The admin contact.
	AdminContact *AdminContact `json:"AdminContact,omitempty" name:"AdminContact"`

	// The technical contact.
	TechnicalContact *TechnicalContact `json:"TechnicalContact,omitempty" name:"TechnicalContact"`

	// The contact person for bills.
	BillingContact *BillingContact `json:"BillingContact,omitempty" name:"BillingContact"`

	// The creation time.
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The registrant type. Valid values: `I` (individual), `E` (organization).
	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`

	// The last updated time.
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// The account ID.
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// Whether the profile is the default one.
	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
}

// Predefined struct for user
type ModifyOwnerIntlBatchRequestParams struct {
	// The domains.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The user ID.
	ToUin *string `json:"ToUin,omitempty" name:"ToUin"`

	// Whether to transfer the DNS service.
	DnsTransfer *bool `json:"DnsTransfer,omitempty" name:"DnsTransfer"`
}

type ModifyOwnerIntlBatchRequest struct {
	*tchttp.BaseRequest
	
	// The domains.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The user ID.
	ToUin *string `json:"ToUin,omitempty" name:"ToUin"`

	// Whether to transfer the DNS service.
	DnsTransfer *bool `json:"DnsTransfer,omitempty" name:"DnsTransfer"`
}

func (r *ModifyOwnerIntlBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnerIntlBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "ToUin")
	delete(f, "DnsTransfer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOwnerIntlBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOwnerIntlBatchResponseParams struct {
	// The ID of the bulk task.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyOwnerIntlBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOwnerIntlBatchResponseParams `json:"Response"`
}

func (r *ModifyOwnerIntlBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOwnerIntlBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceInfoNew struct {
	// The domain suffix, such as `.com`.
	Tld *string `json:"Tld,omitempty" name:"Tld"`

	// The purchase years. Value range: [1-10]
	Year *uint64 `json:"Year,omitempty" name:"Year"`

	// The original price of the domain.
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// The current price of the domain.
	RealPrice *float64 `json:"RealPrice,omitempty" name:"RealPrice"`

	// The domain purchase type. Valid values: `new`, `renew`, `redem` (redeem), `tran` (transfer in).
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// The title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Title *string `json:"Title,omitempty" name:"Title"`
}

type RegistrantContact struct {
	// The first name.
	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

	// The last name.
	LastName *string `json:"LastName,omitempty" name:"LastName"`

	// The country or region name, such as `CN`.
	Country *string `json:"Country,omitempty" name:"Country"`

	// The province or state name.
	Province *string `json:"Province,omitempty" name:"Province"`

	// The city name.
	City *string `json:"City,omitempty" name:"City"`

	// The address line 1.
	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`

	// The zip code.
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// The email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// The mobile number, such as `+86.1360000000`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// The company or organization name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// The job title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobTitle *string `json:"JobTitle,omitempty" name:"JobTitle"`

	// The address line 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressLineTwo *string `json:"AddressLineTwo,omitempty" name:"AddressLineTwo"`

	// The fax number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fax *string `json:"Fax,omitempty" name:"Fax"`
}

// Predefined struct for user
type RenewIntlDomainBatchRequestParams struct {
	// The domains to check.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The period (1 to 10 years). If this parameter is left empty, premium domains cannot be checked.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Payment method. Valid value: `1` (account balance).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type RenewIntlDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// The domains to check.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The period (1 to 10 years). If this parameter is left empty, premium domains cannot be checked.
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// Payment method. Valid value: `1` (account balance).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *RenewIntlDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewIntlDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Period")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewIntlDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewIntlDomainBatchResponseParams struct {
	// The ID of the bulk task.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewIntlDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *RenewIntlDomainBatchResponseParams `json:"Response"`
}

func (r *RenewIntlDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewIntlDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendIntlPhoneEmailCodeRequestParams struct {
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`
}

type SendIntlPhoneEmailCodeRequest struct {
	*tchttp.BaseRequest
	
	// The type. Valid values: `1` (mobile number), `2` (email address).
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// The mobile number or email address.
	Code *string `json:"Code,omitempty" name:"Code"`
}

func (r *SendIntlPhoneEmailCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendIntlPhoneEmailCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendIntlPhoneEmailCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendIntlPhoneEmailCodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendIntlPhoneEmailCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendIntlPhoneEmailCodeResponseParams `json:"Response"`
}

func (r *SendIntlPhoneEmailCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendIntlPhoneEmailCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetIntlDomainAutoRenewRequestParams struct {
	// The domain ID.
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// Whether to enable auto-renewal. Valid values: `1` (enable), `2` (disable).
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

type SetIntlDomainAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// The domain ID.
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// Whether to enable auto-renewal. Valid values: `1` (enable), `2` (disable).
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *SetIntlDomainAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetIntlDomainAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetIntlDomainAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetIntlDomainAutoRenewResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetIntlDomainAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *SetIntlDomainAutoRenewResponseParams `json:"Response"`
}

func (r *SetIntlDomainAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetIntlDomainAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TechnicalContact struct {
	// The first name.
	FirstName *string `json:"FirstName,omitempty" name:"FirstName"`

	// The last name.
	LastName *string `json:"LastName,omitempty" name:"LastName"`

	// The country or region name, such as `CN`.
	Country *string `json:"Country,omitempty" name:"Country"`

	// The province or state name.
	Province *string `json:"Province,omitempty" name:"Province"`

	// The city name.
	City *string `json:"City,omitempty" name:"City"`

	// The address line 1.
	AddressLine *string `json:"AddressLine,omitempty" name:"AddressLine"`

	// The zip code.
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// The email address.
	Email *string `json:"Email,omitempty" name:"Email"`

	// The mobile number, such as `+86.13600000000`.
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// The company or organization name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// The job title.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobTitle *string `json:"JobTitle,omitempty" name:"JobTitle"`

	// The address line 2.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddressLineTwo *string `json:"AddressLineTwo,omitempty" name:"AddressLineTwo"`

	// The fax number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fax *string `json:"Fax,omitempty" name:"Fax"`
}

// Predefined struct for user
type TransferInIntlDomainBatchRequestParams struct {
	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The transfer passwords for the domains.
	PassWords []*string `json:"PassWords,omitempty" name:"PassWords"`

	// The domains to be bulk transferred in.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The payment method. Valid value: `1` (account balance).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to enable the transfer prohibition lock.
	TransferProhibition *bool `json:"TransferProhibition,omitempty" name:"TransferProhibition"`

	// Whether to enable the update prohibition lock.
	UpdateProhibition *bool `json:"UpdateProhibition,omitempty" name:"UpdateProhibition"`

	// Whether to enable the 60-day inter-registrar transfer lock.
	LockTransfer *bool `json:"LockTransfer,omitempty" name:"LockTransfer"`
}

type TransferInIntlDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// The profile ID.
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// The transfer passwords for the domains.
	PassWords []*string `json:"PassWords,omitempty" name:"PassWords"`

	// The domains to be bulk transferred in.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// The payment method. Valid value: `1` (account balance).
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// Whether to enable auto-renewal.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// Whether to enable the transfer prohibition lock.
	TransferProhibition *bool `json:"TransferProhibition,omitempty" name:"TransferProhibition"`

	// Whether to enable the update prohibition lock.
	UpdateProhibition *bool `json:"UpdateProhibition,omitempty" name:"UpdateProhibition"`

	// Whether to enable the 60-day inter-registrar transfer lock.
	LockTransfer *bool `json:"LockTransfer,omitempty" name:"LockTransfer"`
}

func (r *TransferInIntlDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInIntlDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "PassWords")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "TransferProhibition")
	delete(f, "UpdateProhibition")
	delete(f, "LockTransfer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferInIntlDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferInIntlDomainBatchResponseParams struct {
	// The bulk purchase log ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TransferInIntlDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferInIntlDomainBatchResponseParams `json:"Response"`
}

func (r *TransferInIntlDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInIntlDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionIntlBatchRequestParams struct {
	// The domain array.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to enable transfer prohibition. Valid values: `true` (enable), `false` (disable).
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type TransferProhibitionIntlBatchRequest struct {
	*tchttp.BaseRequest
	
	// The domain array.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to enable transfer prohibition. Valid values: `true` (enable), `false` (disable).
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *TransferProhibitionIntlBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionIntlBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferProhibitionIntlBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionIntlBatchResponseParams struct {
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TransferProhibitionIntlBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferProhibitionIntlBatchResponseParams `json:"Response"`
}

func (r *TransferProhibitionIntlBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionIntlBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionIntlBatchRequestParams struct {
	// The domain array.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to enable update prohibition. Valid values: `true` (enable), `false` (disable).
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type UpdateProhibitionIntlBatchRequest struct {
	*tchttp.BaseRequest
	
	// The domain array.
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// Whether to enable update prohibition. Valid values: `true` (enable), `false` (disable).
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateProhibitionIntlBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionIntlBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProhibitionIntlBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionIntlBatchResponseParams struct {
	// The log ID.
	LogId *int64 `json:"LogId,omitempty" name:"LogId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateProhibitionIntlBatchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProhibitionIntlBatchResponseParams `json:"Response"`
}

func (r *UpdateProhibitionIntlBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionIntlBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}