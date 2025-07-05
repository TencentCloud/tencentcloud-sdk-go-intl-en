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

package v20210409

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AllocateCustomerCreditRequestParams struct {
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type AllocateCustomerCreditRequest struct {
	*tchttp.BaseRequest
	
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

func (r *AllocateCustomerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddedCredit")
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateCustomerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateCustomerCreditResponseParams struct {
	// The updated total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil,omitempty" name:"TotalCredit"`

	// The updated available credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateCustomerCreditResponse struct {
	*tchttp.BaseResponse
	Response *AllocateCustomerCreditResponseParams `json:"Response"`
}

func (r *AllocateCustomerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CountryCodeItem struct {
	// Country/region name in English
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// IOS2 standard country/region code
	IOS2 *string `json:"IOS2,omitnil,omitempty" name:"IOS2"`

	// IOS3 standard country/region code
	IOS3 *string `json:"IOS3,omitnil,omitempty" name:"IOS3"`

	// Phone code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and special symbols [!@#$%^&*()]. Spaces are not allowed.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Confirm the password. It must be the same as the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// Customer’s country/region code, which can be obtained via the `GetCountryCodes` API, such as “852”.
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// Customer’s ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Expanded field, which is left empty by default.
	Extended *string `json:"Extended,omitnil,omitempty" name:"Extended"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and special symbols [!@#$%^&*()]. Spaces are not allowed.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Confirm the password. It must be the same as the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// Customer’s country/region code, which can be obtained via the `GetCountryCodes` API, such as “852”.
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// Customer’s ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Expanded field, which is left empty by default.
	Extended *string `json:"Extended,omitnil,omitempty" name:"Extended"`
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

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Account UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type GetCountryCodesRequestParams struct {

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

// Predefined struct for user
type GetCountryCodesResponseParams struct {
	// List of country/region codes
	Data []*CountryCodeItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCountryCodesResponse struct {
	*tchttp.BaseResponse
	Response *GetCountryCodesResponseParams `json:"Response"`
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

type QueryCreditAllocationHistoryData struct {
	// Allocation time
	AllocatedTime *string `json:"AllocatedTime,omitnil,omitempty" name:"AllocatedTime"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Allocated credit value
	Credit *float64 `json:"Credit,omitnil,omitempty" name:"Credit"`

	// The allocated total credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil,omitempty" name:"AllocatedCredit"`
}

// Predefined struct for user
type QueryCreditAllocationHistoryRequestParams struct {
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type QueryCreditAllocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *QueryCreditAllocationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditAllocationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditAllocationHistoryResponseParams struct {
	// Total number of records
	// Note: this field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List of record details
	// Note: this field may return null, indicating that no valid values can be obtained.
	History []*QueryCreditAllocationHistoryData `json:"History,omitnil,omitempty" name:"History"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCreditAllocationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditAllocationHistoryResponseParams `json:"Response"`
}

func (r *QueryCreditAllocationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomersCreditData struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Phone
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// Email
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Overdue payment flag
	Arrears *string `json:"Arrears,omitnil,omitempty" name:"Arrears"`

	// Binding time
	AssociationTime *string `json:"AssociationTime,omitnil,omitempty" name:"AssociationTime"`

	// Expiration time
	RecentExpiry *string `json:"RecentExpiry,omitnil,omitempty" name:"RecentExpiry"`

	// The UIN of reseller’s customer
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Credit granted to reseller’s customer
	Credit *float64 `json:"Credit,omitnil,omitempty" name:"Credit"`

	// The remaining credit of reseller’s customer
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// 0: Identity not verified; 1: Individual identity verified; 2: Enterprise identity verified.
	IdentifyType *uint64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// Customer remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Forced status
	Force *int64 `json:"Force,omitnil,omitempty" name:"Force"`
}

// Predefined struct for user
type QueryCustomersCreditRequestParams struct {
	// Search condition type. You can only search by UIN, name, or remarks.
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type QueryCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
	// Search condition type. You can only search by UIN, name, or remarks.
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *QueryCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterType")
	delete(f, "Filter")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomersCreditResponseParams struct {
	// Queries the list of customers
	// Note: this field may return null, indicating that no valid values can be obtained.
	Data []*QueryCustomersCreditData `json:"Data,omitnil,omitempty" name:"Data"`

	// Number of customers
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditRequestParams struct {

}

type QueryPartnerCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryPartnerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPartnerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditResponseParams struct {
	// Allocated credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil,omitempty" name:"AllocatedCredit"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil,omitempty" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryPartnerCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryPartnerCreditResponseParams `json:"Response"`
}

func (r *QueryPartnerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}