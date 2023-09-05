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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ActionSummaryOverviewItem struct {
	// Transaction type code
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// Transaction type name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

// Predefined struct for user
type AllocateCustomerCreditRequestParams struct {
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil" name:"ClientUin"`
}

type AllocateCustomerCreditRequest struct {
	*tchttp.BaseRequest
	
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil" name:"ClientUin"`
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
	TotalCredit *float64 `json:"TotalCredit,omitnil" name:"TotalCredit"`

	// The updated available credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type BillDetailData struct {
	// Reseller account
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayerAccountId *int64 `json:"PayerAccountId,omitnil" name:"PayerAccountId"`

	// Customer account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerAccountId *int64 `json:"OwnerAccountId,omitnil" name:"OwnerAccountId"`

	// Operator account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorAccountId *int64 `json:"OperatorAccountId,omitnil" name:"OperatorAccountId"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Billing mode
	// `Monthly subscription` (Monthly subscription)
	// `Pay-As-You-Go resources` (Pay-as-you-go)
	// `Standard RI` (Reserved instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillingMode *string `json:"BillingMode,omitnil" name:"BillingMode"`

	// Project name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// Resource region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Resource AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailabilityZone *string `json:"AvailabilityZone,omitnil" name:"AvailabilityZone"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Subproduct name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductName *string `json:"SubProductName,omitnil" name:"SubProductName"`

	// Settlement type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionType *string `json:"TransactionType,omitnil" name:"TransactionType"`

	// Transaction ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionId *string `json:"TransactionId,omitnil" name:"TransactionId"`

	// Settlement time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionTime *string `json:"TransactionTime,omitnil" name:"TransactionTime"`

	// Start time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageStartTime *string `json:"UsageStartTime,omitnil" name:"UsageStartTime"`

	// End time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageEndTime *string `json:"UsageEndTime,omitnil" name:"UsageEndTime"`

	// Component
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// Component list price
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentListPrice *string `json:"ComponentListPrice,omitnil" name:"ComponentListPrice"`

	// Price unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentPriceMeasurementUnit *string `json:"ComponentPriceMeasurementUnit,omitnil" name:"ComponentPriceMeasurementUnit"`

	// Component usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsage *string `json:"ComponentUsage,omitnil" name:"ComponentUsage"`

	// Component usage unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsageUnit *string `json:"ComponentUsageUnit,omitnil" name:"ComponentUsageUnit"`

	// Resource usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageDuration *string `json:"UsageDuration,omitnil" name:"UsageDuration"`

	// Duration unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationUnit *string `json:"DurationUnit,omitnil" name:"DurationUnit"`

	// Original cost
	// Original cost = component list price * component usage * usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Discount, which defaults to `1`, indicating there is no discount.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountRate *string `json:"DiscountRate,omitnil" name:"DiscountRate"`

	// Currency
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Discounted total
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalAmountAfterDiscount *string `json:"TotalAmountAfterDiscount,omitnil" name:"TotalAmountAfterDiscount"`

	// Voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherDeduction *string `json:"VoucherDeduction,omitnil" name:"VoucherDeduction"`

	// Total cost = discounted total - voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type BusinessSummaryOverviewItem struct {
	// Product code
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// List price accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// Consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type CountryCodeItem struct {
	// Country/region name in English
	EnName *string `json:"EnName,omitnil" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitnil" name:"Name"`


	IOS2 *string `json:"IOS2,omitnil" name:"IOS2"`


	IOS3 *string `json:"IOS3,omitnil" name:"IOS3"`

	// International dialing code
	Code *string `json:"Code,omitnil" name:"Code"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitnil" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitnil" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil" name:"Area"`

	// Extension field, which is left empty by default.
	Extended *string `json:"Extended,omitnil" name:"Extended"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitnil" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitnil" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil" name:"Area"`

	// Extension field, which is left empty by default.
	Extended *string `json:"Extended,omitnil" name:"Extended"`
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
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type CustomerBillDetailData struct {
	// Reseller account
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayerAccountId *int64 `json:"PayerAccountId,omitnil" name:"PayerAccountId"`

	// Customer account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerAccountId *int64 `json:"OwnerAccountId,omitnil" name:"OwnerAccountId"`

	// Operator account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorAccountId *int64 `json:"OperatorAccountId,omitnil" name:"OperatorAccountId"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// Billing mode
	// `Monthly subscription` (Monthly subscription)
	// `Pay-As-You-Go resources` (Pay-as-you-go)
	// `Standard RI` (Reserved instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillingMode *string `json:"BillingMode,omitnil" name:"BillingMode"`

	// Project name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// Resource region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Resource AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailabilityZone *string `json:"AvailabilityZone,omitnil" name:"AvailabilityZone"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Subproduct name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductName *string `json:"SubProductName,omitnil" name:"SubProductName"`

	// Settlement type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionType *string `json:"TransactionType,omitnil" name:"TransactionType"`

	// Transaction ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionId *string `json:"TransactionId,omitnil" name:"TransactionId"`

	// Settlement time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionTime *string `json:"TransactionTime,omitnil" name:"TransactionTime"`

	// Start time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageStartTime *string `json:"UsageStartTime,omitnil" name:"UsageStartTime"`

	// End time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageEndTime *string `json:"UsageEndTime,omitnil" name:"UsageEndTime"`

	// Component
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentType *string `json:"ComponentType,omitnil" name:"ComponentType"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil" name:"ComponentName"`

	// Component list price
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentListPrice *string `json:"ComponentListPrice,omitnil" name:"ComponentListPrice"`

	// Price unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentPriceMeasurementUnit *string `json:"ComponentPriceMeasurementUnit,omitnil" name:"ComponentPriceMeasurementUnit"`

	// Component usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsage *string `json:"ComponentUsage,omitnil" name:"ComponentUsage"`

	// Component usage unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsageUnit *string `json:"ComponentUsageUnit,omitnil" name:"ComponentUsageUnit"`

	// Resource usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageDuration *string `json:"UsageDuration,omitnil" name:"UsageDuration"`

	// Duration unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationUnit *string `json:"DurationUnit,omitnil" name:"DurationUnit"`

	// Original cost
	// Original cost = component list price * component usage * usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Currency
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// Total cost = discounted total - voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

// Predefined struct for user
type DescribeBillDetailRequestParams struct {
	// The queried month in u200dthe format of “YYYY-MM”, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// A pagination parameter that specifies the number of entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A pagination parameter that specifies the current page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Billing mode. Valid values: `prePay` (Monthly subscription), postPay` (Pay-As-You-Go resources).
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values: `prepay_purchase` (Purchase), `prepay_renew` (Renewal), `prepay_modify` (Upgrade/Downgrade), `prepay_return` ( Monthly subscription refund), `postpay_deduct` (Pay-as-you-go), `postpay_deduct_h` (Hourly settlement), `postpay_deduct_d` (Daily settlement), `postpay_deduct_m` (Monthly settlement), `offline_deduct` (Offline project deduction), `online_deduct` (Offline product deduction), `recon_deduct` (Adjustment - deduction), `recon_increase` (Adjustment - compensation), `ripay_purchase` (One-off RI Fee), `postpay_deduct_s` (Spot), `ri_hour_pay` (Hourly RI fee), `prePurchase` (New monthly subscription), `preRenew` (Monthly subscription renewal), `preUpgrade` (Upgrade/Downgrade), `preDowngrade` (Upgrade/Downgrade), `svp_hour_pay` (Hourly Savings Plan fee), `recon_guarantee` (Minimum spend deduction), `pre_purchase` (New monthly subscription), `pre_renew` (Monthly subscription renewal), `pre_upgrade` (Upgrade/Downgrade), `pre_downgrade` (Upgrade/Downgrade).
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// The queried month in u200dthe format of “YYYY-MM”, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// A pagination parameter that specifies the number of entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A pagination parameter that specifies the current page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Billing mode. Valid values: `prePay` (Monthly subscription), postPay` (Pay-As-You-Go resources).
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values: `prepay_purchase` (Purchase), `prepay_renew` (Renewal), `prepay_modify` (Upgrade/Downgrade), `prepay_return` ( Monthly subscription refund), `postpay_deduct` (Pay-as-you-go), `postpay_deduct_h` (Hourly settlement), `postpay_deduct_d` (Daily settlement), `postpay_deduct_m` (Monthly settlement), `offline_deduct` (Offline project deduction), `online_deduct` (Offline product deduction), `recon_deduct` (Adjustment - deduction), `recon_increase` (Adjustment - compensation), `ripay_purchase` (One-off RI Fee), `postpay_deduct_s` (Spot), `ri_hour_pay` (Hourly RI fee), `prePurchase` (New monthly subscription), `preRenew` (Monthly subscription renewal), `preUpgrade` (Upgrade/Downgrade), `preDowngrade` (Upgrade/Downgrade), `svp_hour_pay` (Hourly Savings Plan fee), `recon_guarantee` (Minimum spend deduction), `pre_purchase` (New monthly subscription), `pre_renew` (Monthly subscription renewal), `pre_upgrade` (Upgrade/Downgrade), `pre_downgrade` (Upgrade/Downgrade).
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`
}

func (r *DescribeBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "PageSize")
	delete(f, "Page")
	delete(f, "PayMode")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailResponseParams struct {
	// Data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailSet []*CustomerBillDetailData `json:"DetailSet,omitnil" name:"DetailSet"`

	// Total number of data entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDetailResponseParams `json:"Response"`
}

func (r *DescribeBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeResponseParams struct {
	// Payment mode details in the customer bill data totaled by payment mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByPayModeResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductResponseParams struct {
	// Bill details from the product dimension
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByProductResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionResponseParams struct {
	// Region details in the customer bill data totaled by region
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByRegionResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillDetailRequestParams struct {
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// A pagination parameter that specifies the number of entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A pagination parameter that specifies the current page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` ( Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil" name:"IsConfirmed"`
}

type DescribeCustomerBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// A pagination parameter that specifies the number of entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A pagination parameter that specifies the current page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` ( Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil" name:"IsConfirmed"`
}

func (r *DescribeCustomerBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	delete(f, "Month")
	delete(f, "PageSize")
	delete(f, "Page")
	delete(f, "PayMode")
	delete(f, "ActionType")
	delete(f, "IsConfirmed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillDetailResponseParams struct {
	// Total number of data entries
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailSet []*BillDetailData `json:"DetailSet,omitnil" name:"DetailSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerBillDetailResponseParams `json:"Response"`
}

func (r *DescribeCustomerBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillSummaryRequestParams struct {
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` (Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil" name:"IsConfirmed"`
}

type DescribeCustomerBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil" name:"Month"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` (Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil" name:"IsConfirmed"`
}

func (r *DescribeCustomerBillSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	delete(f, "Month")
	delete(f, "PayMode")
	delete(f, "ActionType")
	delete(f, "IsConfirmed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillSummaryResponseParams struct {
	// Total amount
	TotalCost *float64 `json:"TotalCost,omitnil" name:"TotalCost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerBillSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerBillSummaryResponseParams `json:"Response"`
}

func (r *DescribeCustomerBillSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerInfoData struct {
	// Customer UIN Note: This field may return null, indicating that no valid values can be obtained.
	CustomerUin *string `json:"CustomerUin,omitnil" name:"CustomerUin"`

	// Email Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil" name:"Email"`

	// Mobile number Note: This field may return null, indicating that no valid values can be obtained.
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// Remarks Note: This field may return null, indicating that no valid values can be obtained.
	Mark *string `json:"Mark,omitnil" name:"Mark"`

	// Displayed name Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Binding time Note: This field may return null, indicating that no valid values can be obtained.
	BindTime *string `json:"BindTime,omitnil" name:"BindTime"`

	// Account status Valid values: `0` (Not frozen),  `1` (Frozen).  Note: This field may return null, indicating that no valid values can be obtained.
	AccountStatus *string `json:"AccountStatus,omitnil" name:"AccountStatus"`

	// Identity verification status Note: This field may return null, indicating that no valid values can be obtained.
	AuthStatus *string `json:"AuthStatus,omitnil" name:"AuthStatus"`
}

// Predefined struct for user
type DescribeCustomerInfoRequestParams struct {
	// List of customer UINs
	CustomerUin []*int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

type DescribeCustomerInfoRequest struct {
	*tchttp.BaseRequest
	
	// List of customer UINs
	CustomerUin []*int64 `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

func (r *DescribeCustomerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerInfoResponseParams struct {
	// Customer information Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeCustomerInfoData `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerInfoResponseParams `json:"Response"`
}

func (r *DescribeCustomerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerUinData struct {
	// Customer UIN Note: This field may return null, indicating that no valid values can be obtained.
	CustomerUin *string `json:"CustomerUin,omitnil" name:"CustomerUin"`
}

// Predefined struct for user
type DescribeCustomerUinRequestParams struct {
	// Page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Number of data entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeCustomerUinRequest struct {
	*tchttp.BaseRequest
	
	// Page number
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// Number of data entries per page
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeCustomerUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerUinResponseParams struct {
	// List of customer UINs Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeCustomerUinData `json:"Data,omitnil" name:"Data"`

	// The number of customers
	Total *string `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerUinResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerUinResponseParams `json:"Response"`
}

func (r *DescribeCustomerUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerUinResponse) FromJsonString(s string) error {
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
	Data []*CountryCodeItem `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type PayModeSummaryOverviewItem struct {
	// Billing mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Billing mode name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// Bill details in each payment mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitnil" name:"Detail"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

// Predefined struct for user
type QueryAccountVerificationStatusRequestParams struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil" name:"ClientUin"`
}

type QueryAccountVerificationStatusRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil" name:"ClientUin"`
}

func (r *QueryAccountVerificationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAccountVerificationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAccountVerificationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAccountVerificationStatusResponseParams struct {
	// Account verification status
	AccountStatus *bool `json:"AccountStatus,omitnil" name:"AccountStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryAccountVerificationStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryAccountVerificationStatusResponseParams `json:"Response"`
}

func (r *QueryAccountVerificationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAccountVerificationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreditAllocationHistoryData struct {
	// Allocation time
	AllocatedTime *string `json:"AllocatedTime,omitnil" name:"AllocatedTime"`

	// Operator
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// Allocated credit value
	Credit *float64 `json:"Credit,omitnil" name:"Credit"`

	// The allocated total credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil" name:"AllocatedCredit"`
}

// Predefined struct for user
type QueryCreditAllocationHistoryRequestParams struct {
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type QueryCreditAllocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of record details
	// Note: This field may return null, indicating that no valid values can be obtained.
	History []*QueryCreditAllocationHistoryData `json:"History,omitnil" name:"History"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	UinList []*uint64 `json:"UinList,omitnil" name:"UinList"`
}

type QueryCreditByUinListRequest struct {
	*tchttp.BaseRequest
	
	// User list
	UinList []*uint64 `json:"UinList,omitnil" name:"UinList"`
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
	Data []*QueryDirectCustomersCreditData `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type QueryCreditQuotaRequestParams struct {

}

type QueryCreditQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryCreditQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditQuotaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryCreditQuotaResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditQuotaResponseParams `json:"Response"`
}

func (r *QueryCreditQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomersCreditData struct {
	// Name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Mobile number
	Mobile *string `json:"Mobile,omitnil" name:"Mobile"`

	// Email
	Email *string `json:"Email,omitnil" name:"Email"`

	// Overdue payment flag
	Arrears *string `json:"Arrears,omitnil" name:"Arrears"`

	// Binding time
	AssociationTime *string `json:"AssociationTime,omitnil" name:"AssociationTime"`

	// Expiration time
	RecentExpiry *string `json:"RecentExpiry,omitnil" name:"RecentExpiry"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil" name:"ClientUin"`

	// Credit allocated to a customer
	Credit *float64 `json:"Credit,omitnil" name:"Credit"`

	// The remaining credit of a customer
	RemainingCredit *float64 `json:"RemainingCredit,omitnil" name:"RemainingCredit"`

	// `0`: Identity not verified; `1`: Individual identity verified; `2`: Enterprise identity verified.
	IdentifyType *uint64 `json:"IdentifyType,omitnil" name:"IdentifyType"`

	// Customer remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Forced status
	Force *int64 `json:"Force,omitnil" name:"Force"`
}

// Predefined struct for user
type QueryCustomersCreditRequestParams struct {
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil" name:"Order"`
}

type QueryCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil" name:"Order"`
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
	Data []*QueryCustomersCreditData `json:"Data,omitnil" name:"Data"`

	// Number of customers
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Uin *uint64 `json:"Uin,omitnil" name:"Uin"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil" name:"RemainingCredit"`
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
	Data []*QueryDirectCustomersCreditData `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil" name:"AllocatedCredit"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type QueryVoucherAmountByUinItem struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil" name:"ClientUin"`

	// Voucher quota
	TotalAmount *float64 `json:"TotalAmount,omitnil" name:"TotalAmount"`

	// Voucher amount
	RemainAmount *float64 `json:"RemainAmount,omitnil" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherAmountByUinRequestParams struct {
	// Customer UIN list
	ClientUins []*uint64 `json:"ClientUins,omitnil" name:"ClientUins"`
}

type QueryVoucherAmountByUinRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN list
	ClientUins []*uint64 `json:"ClientUins,omitnil" name:"ClientUins"`
}

func (r *QueryVoucherAmountByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherAmountByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherAmountByUinResponseParams struct {
	// Customer voucher quota information
	Data []*QueryVoucherAmountByUinItem `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryVoucherAmountByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherAmountByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherAmountByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinItem struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil" name:"ClientUin"`

	// The total number of vouchers
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Voucher details
	Data []*QueryVoucherListByUinVoucherItem `json:"Data,omitnil" name:"Data"`
}

// Predefined struct for user
type QueryVoucherListByUinRequestParams struct {
	// Customer UIN list
	ClientUins []*uint64 `json:"ClientUins,omitnil" name:"ClientUins"`

	// Voucher status. If this parameter is not passed in, all status will be queried by default. Valid values: `Unused`, `Used`, `Expired`.
	Status *string `json:"Status,omitnil" name:"Status"`
}

type QueryVoucherListByUinRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN list
	ClientUins []*uint64 `json:"ClientUins,omitnil" name:"ClientUins"`

	// Voucher status. If this parameter is not passed in, all status will be queried by default. Valid values: `Unused`, `Used`, `Expired`.
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *QueryVoucherListByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherListByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherListByUinResponseParams struct {
	// Customer voucher information
	Data []*QueryVoucherListByUinItem `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryVoucherListByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherListByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherListByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinVoucherItem struct {
	// Voucher ID
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// Voucher status
	VoucherStatus *string `json:"VoucherStatus,omitnil" name:"VoucherStatus"`

	// Voucher value
	TotalAmount *float64 `json:"TotalAmount,omitnil" name:"TotalAmount"`

	// Balance
	RemainAmount *float64 `json:"RemainAmount,omitnil" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherPoolRequestParams struct {

}

type QueryVoucherPoolRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryVoucherPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherPoolResponseParams struct {
	// Reseller name
	AgentName *string `json:"AgentName,omitnil" name:"AgentName"`

	// Reseller role type (1: Reseller; 2: Distributor; 3: Second-level reseller)
	AccountType *int64 `json:"AccountType,omitnil" name:"AccountType"`

	// Total quota
	TotalQuota *float64 `json:"TotalQuota,omitnil" name:"TotalQuota"`

	// Remaining quota
	RemainingQuota *float64 `json:"RemainingQuota,omitnil" name:"RemainingQuota"`

	// The number of issued vouchers
	IssuedNum *int64 `json:"IssuedNum,omitnil" name:"IssuedNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryVoucherPoolResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherPoolResponseParams `json:"Response"`
}

func (r *QueryVoucherPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSummaryOverviewItem struct {
	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// Region name
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}