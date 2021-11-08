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

package v20210409

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AssignClientCreditRequest struct {
	*tchttp.BaseRequest

	// Specific value of the credit allocated to the customer
	QuotaNum *float64 `json:"QuotaNum,omitempty" name:"QuotaNum"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *AssignClientCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignClientCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QuotaNum")
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignClientCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AssignClientCreditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignClientCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignClientCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountryCodeItem struct {

	// Country/region name in English
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitempty" name:"Name"`

	// IOS2 standard country/region code
	IOS2 *string `json:"IOS2,omitempty" name:"IOS2"`

	// IOS3 standard country/region code
	IOS3 *string `json:"IOS3,omitempty" name:"IOS3"`

	// Phone code
	Code *string `json:"Code,omitempty" name:"Code"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Account type of a new customer. Valid value: `business`.
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and special symbols [!@#$%^&*()]. Spaces are not allowed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Confirm the password. It must be the same as the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country code, which can be obtained via the `GetCountryCodes` API, such as 852.
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// ISO2 standard country code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Expanded field, which is left empty by default.
	Extended *string `json:"Extended,omitempty" name:"Extended"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountType")
	delete(f, "Mail")
	delete(f, "Password")
	delete(f, "ConfirmPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Area")
	delete(f, "Extended")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account UIN
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCountryCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCountryCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of country/region codes
		Data []*CountryCodeItem `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCountryCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentCreditRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryAgentCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAgentCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentCreditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Allocated credit
		AssignedCreditAmount *float64 `json:"AssignedCreditAmount,omitempty" name:"AssignedCreditAmount"`

		// Total credit
		CustomerCreditAmount *float64 `json:"CustomerCreditAmount,omitempty" name:"CustomerCreditAmount"`

		// Remaining credit
		RemainingCreditAmount *float64 `json:"RemainingCreditAmount,omitempty" name:"RemainingCreditAmount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryClientListItem struct {

	// Name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type
	// Note: this field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Phone
	// Note: this field may return null, indicating that no valid values can be obtained.
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// Email
	// Note: this field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Overdue payment flag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Arrears *string `json:"Arrears,omitempty" name:"Arrears"`

	// Binding time
	// Note: this field may return null, indicating that no valid values can be obtained.
	AssociationTime *string `json:"AssociationTime,omitempty" name:"AssociationTime"`

	// Expiration time
	// Note: this field may return null, indicating that no valid values can be obtained.
	RecentExpiry *string `json:"RecentExpiry,omitempty" name:"RecentExpiry"`

	// Customer UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// Credit granted to customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreditAmount *float64 `json:"CreditAmount,omitempty" name:"CreditAmount"`

	// Customer's remaining credit
	// Note: this field may return null, indicating that no valid values can be obtained.
	RestCreditAmount *float64 `json:"RestCreditAmount,omitempty" name:"RestCreditAmount"`
}

type QueryClientListRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryClientListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryClientListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryClientListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryClientListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Queries the list of customers
	// Note: this field may return null, indicating that no valid values can be obtained.
		Data []*QueryClientListItem `json:"Data,omitempty" name:"Data"`

		// Number of customers
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryClientListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryClientListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreditHistoryRequest struct {
	*tchttp.BaseRequest

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of data entries per page
	PageRow *uint64 `json:"PageRow,omitempty" name:"PageRow"`
}

func (r *QueryCreditHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Page")
	delete(f, "PageRow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreditHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of records
	// Note: this field may return null, indicating that no valid values can be obtained.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// List of record details
	// Note: this field may return null, indicating that no valid values can be obtained.
		History []*QueryUinCreditHistoryData `json:"History,omitempty" name:"History"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCreditHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUinCreditHistoryData struct {

	// Credit allocatee UIN
	CreditAssignUin *uint64 `json:"CreditAssignUin,omitempty" name:"CreditAssignUin"`

	// Allocation time
	AssginTime *string `json:"AssginTime,omitempty" name:"AssginTime"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Allocated credit value
	CreditAmount *float64 `json:"CreditAmount,omitempty" name:"CreditAmount"`
}
