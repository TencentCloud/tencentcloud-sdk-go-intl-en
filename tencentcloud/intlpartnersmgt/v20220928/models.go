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

package v20220928

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type AllocateCustomerCreditRequestParams struct {
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`
}

type AllocateCustomerCreditRequest struct {
	*tchttp.BaseRequest
	
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`
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
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// The updated available credit
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitempty" name:"Name"`


	IOS2 *string `json:"IOS2,omitempty" name:"IOS2"`


	IOS3 *string `json:"IOS3,omitempty" name:"IOS3"`

	// International dialing code
	Code *string `json:"Code,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Extension field, which is left empty by default.
	Extended *string `json:"Extended,omitempty" name:"Extended"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Extension field, which is left empty by default.
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

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Account UIN
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Data []*CountryCodeItem `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AllocatedTime *string `json:"AllocatedTime,omitempty" name:"AllocatedTime"`

	// Operator
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// Allocated credit value
	Credit *float64 `json:"Credit,omitempty" name:"Credit"`

	// The allocated total credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitempty" name:"AllocatedCredit"`
}

// Predefined struct for user
type QueryCreditAllocationHistoryRequestParams struct {
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type QueryCreditAllocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
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
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// List of record details
	// Note: This field may return null, indicating that no valid values can be obtained.
	History []*QueryCreditAllocationHistoryData `json:"History,omitempty" name:"History"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type QueryCreditByUinListRequestParams struct {
	// User list
	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

type QueryCreditByUinListRequest struct {
	*tchttp.BaseRequest
	
	// User list
	UinList []*uint64 `json:"UinList,omitempty" name:"UinList"`
}

func (r *QueryCreditByUinListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditByUinListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditByUinListResponseParams struct {
	// User information list
	Data []*QueryDirectCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCreditByUinListResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditByUinListResponseParams `json:"Response"`
}

func (r *QueryCreditByUinListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomersCreditData struct {
	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Type
	Type *string `json:"Type,omitempty" name:"Type"`

	// Mobile number
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// Email
	Email *string `json:"Email,omitempty" name:"Email"`

	// Overdue payment flag
	Arrears *string `json:"Arrears,omitempty" name:"Arrears"`

	// Binding time
	AssociationTime *string `json:"AssociationTime,omitempty" name:"AssociationTime"`

	// Expiration time
	RecentExpiry *string `json:"RecentExpiry,omitempty" name:"RecentExpiry"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitempty" name:"ClientUin"`

	// Credit allocated to a customer
	Credit *float64 `json:"Credit,omitempty" name:"Credit"`

	// The remaining credit of a customer
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// `0`: Identity not verified; `1`: Individual identity verified; `2`: Enterprise identity verified.
	IdentifyType *uint64 `json:"IdentifyType,omitempty" name:"IdentifyType"`

	// Customer remarks
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Forced status
	Force *int64 `json:"Force,omitempty" name:"Force"`
}

// Predefined struct for user
type QueryCustomersCreditRequestParams struct {
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitempty" name:"Order"`
}

type QueryCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitempty" name:"Order"`
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
	// The list of queried customers
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*QueryCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// Number of customers
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type QueryDirectCustomersCreditData struct {
	// User UIN
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`
}

// Predefined struct for user
type QueryDirectCustomersCreditRequestParams struct {

}

type QueryDirectCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryDirectCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDirectCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDirectCustomersCreditResponseParams struct {
	// Direct customer information list
	Data []*QueryDirectCustomersCreditData `json:"Data,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryDirectCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryDirectCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryDirectCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditResponse) FromJsonString(s string) error {
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
	AllocatedCredit *float64 `json:"AllocatedCredit,omitempty" name:"AllocatedCredit"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitempty" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitempty" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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