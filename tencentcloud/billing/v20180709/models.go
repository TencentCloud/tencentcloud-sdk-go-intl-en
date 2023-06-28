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

package v20180709

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type ActionSummaryOverviewItem struct {
	// Transaction type code
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Transaction type,  which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type ApplicableProducts struct {
	// Valid values: `all products` or names of the applicable products (string). Multiple names are separated by commas.
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If `GoodsName` contains multiple product names and `PayMode` is `*`, it indicates that the voucher can be used in all billing modes for each of the products.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type BillDetail struct {
	// Product name:  The name of a Tencent Cloud product purchased by the user, such as  CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Subproduct name:  The subcategory of a Tencent Cloud product purchased by the user, such as  CVM – Standard S1.
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// Billing mode,  which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Project name:  The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Region:  The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Availability zone: availability zone of a resource, e.g. Guangzhou Zone 3
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Instance ID:  The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Instance name:  The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Order ID:  The order number for a monthly subscription purchase
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// Transaction ID:  The bill number for a deducted payment
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// Transaction time:  The time at which a payment was deducted
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// Usage start time:  The time at which product or service usage starts
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// Usage end time:  The time at which product or service usage ends
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// Component list
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitempty" name:"ComponentSet"`

	// Payer account ID:  The account ID of the payer, which is the unique identifier of a Tencent Cloud user.
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Owner account ID:  The account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator account ID:  The account or role ID of the operator who purchases or activates a resource
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Tag information. Note:  This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// Product code. Note:  This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Subproduct code. Note:  This field may return null, indicating that no valid values can be obtained.
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Transaction type code. Note:  This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Region ID. Note:  This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Price attribute
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceInfo []*string `json:"PriceInfo,omitempty" name:"PriceInfo"`
}

type BillDetailComponent struct {
	// Component type:  The component type of a product or service purchased, such as  CVM instance components including  CPU and memory.
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`

	// Component name:  The specific component of a product or service purchased
	ItemCodeName *string `json:"ItemCodeName,omitempty" name:"ItemCodeName"`

	// Component list price:  The listed unit price of a component. If a customer has applied for a fixed preferential price or contract price, this parameter will not be displayed by default.
	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`

	// Specified price of a component. This parameter has been deprecated.
	//
	// Deprecated: SpecifiedPrice is deprecated.
	SpecifiedPrice *string `json:"SpecifiedPrice,omitempty" name:"SpecifiedPrice"`

	// Component price measurement unit:  The unit of measurement for a component price, which is composed of  USD, usage unit, and duration unit.
	PriceUnit *string `json:"PriceUnit,omitempty" name:"PriceUnit"`

	// Component usage:  The actually settled usage of a component, which is "Raw usage - Deducted usage (including packages)".
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// Component usage unit:  The unit of measurement for component usage
	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`

	// Usage duration:  The resource usage duration
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Duration unit:  The unit of measurement for usage duration
	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`

	// Original cost:  The original cost of a resource, which is "List price x Usage x Usage duration". If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// Discount multiplier:  The discount multiplier applied to the cost of the resource. If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// Total amount after discount:   Total amount after discount = (Original cost - RI deduction (cost) - SP deduction (cost)) x Discount multiplier
	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Component type code. Note:  This field may return null, indicating that no valid values can be obtained.
	ItemCode *string `json:"ItemCode,omitempty" name:"ItemCode"`

	// Component name code. Note:  This field may return null, indicating that no valid values can be obtained.
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// Component contracted price:  The contracted unit price of a component, which is "List price x Discount". Note:  This field may return null, indicating that no valid values can be obtained.
	ContractPrice *string `json:"ContractPrice,omitempty" name:"ContractPrice"`

	// Instance type:  The instance type of a product or service purchased, which can be resource package, RI, SP, or spot instance. Other instance types are not displayed by default. Note:  This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// RI deduction (duration):  The usage duration deducted by RI. Note:  This field may return null, indicating that no valid values can be obtained.
	RiTimeSpan *string `json:"RiTimeSpan,omitempty" name:"RiTimeSpan"`

	// RI deduction (cost):  The amount deducted from the original cost by RI. Note:  This field may return null, indicating that no valid values can be obtained.
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// Savings plan deduction rate:  The discount multiplier that applies to the component based on the remaining commitment of the savings plan. Note:  This field may return null, indicating that no valid values can be obtained.
	SPDeductionRate *string `json:"SPDeductionRate,omitempty" name:"SPDeductionRate"`

	// Cost deduction by SP. This parameter has been deprecated. Note:  This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// SP deduction (cost):  SP deduction (cost) = Cost deduction by SP / SP deduction rate. Note:  This field may return null, indicating that no valid values can be obtained.
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`

	// Blended discount multiplier:  The final discount multiplier that is applied after combining multiple discount types, which is "Total amount after discount / Original cost". Note:  This field may return null, indicating that no valid values can be obtained.
	BlendedDiscount *string `json:"BlendedDiscount,omitempty" name:"BlendedDiscount"`
}

type BillResourceSummary struct {
	// Product name:  The name of a Tencent Cloud product purchased by the user, such as  CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Subproduct name:  The subcategory of a Tencent Cloud product purchased by the user, such as  CVM – Standard S1.
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// Billing mode,  which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Project name:  The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Region:  The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Availability zone:  The availability zone to which a resource belongs, such as Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Instance ID:  The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.	
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Instance name:  The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Transaction type,  which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Order ID:  The order number for a monthly subscription purchase
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// Transaction time:  The time at which a payment was deducted
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// Usage start time:  The time at which product or service usage starts
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// Usage end time:  The time at which product or service usage ends
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// Configuration description:  The billable item names and usage of a resource, which are displayed on the resource bill only.
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// Extended field 1:  Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField1 *string `json:"ExtendField1,omitempty" name:"ExtendField1"`

	// Extended field 2:  Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField2 *string `json:"ExtendField2,omitempty" name:"ExtendField2"`

	// Original cost:  The original cost of a resource, which is "List price x Usage x Usage duration". If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// Discount multiplier:  The discount multiplier applied to the cost of the resource. If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Extended field 3:  Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField3 *string `json:"ExtendField3,omitempty" name:"ExtendField3"`

	// Extended field 4:  Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField4 *string `json:"ExtendField4,omitempty" name:"ExtendField4"`

	// Extended field 5:  Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField5 *string `json:"ExtendField5,omitempty" name:"ExtendField5"`

	// Tag information. Note:  This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// Payer account ID:  The account ID of the payer, which is the unique identifier of a Tencent Cloud user.
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Owner account ID:  The account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator account ID:  The account or role ID of the operator who purchases or activates a resource.
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// Instance type:  The instance type of a product or service purchased, which can be resource package, RI, SP, or spot instance. Other instance types are not displayed by default.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// RI deduction (cost):  The amount deducted from the original cost by RI	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// Cost deduction by SP. This parameter has been deprecated.
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// SP deduction (cost):  SP deduction (cost) = Cost deduction by SP / SP deduction rate	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`
}

type BillTagInfo struct {
	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BusinessSummaryInfo struct {
	// Product code
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Product name:  The name of a Tencent Cloud product purchased by the user, such as  CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Original cost in USD.  This parameter has become valid since Bill 3.0 took effect in May 2021, and before that `-` was returned for this parameter. If a customer has applied for a contract price different from the prices listed on the official website, `-` will also be returned for this parameter. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type BusinessSummaryOverviewItem struct {
	// Product code. Note:  This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Product name:  The name of a Tencent Cloud product purchased by the user, such as  CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type BusinessSummaryTotal struct {
	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type CosDetailSets struct {
	// Bucket name
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// The start time of the usage
	DosageBeginTime *string `json:"DosageBeginTime,omitempty" name:"DosageBeginTime"`

	// The end time of the usage
	DosageEndTime *string `json:"DosageEndTime,omitempty" name:"DosageEndTime"`

	// Subproduct name
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// Billable item name
	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`

	// Usage
	DosageValue *string `json:"DosageValue,omitempty" name:"DosageValue"`

	// Unit of the billable item
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

// Predefined struct for user
type DescribeAccountBalanceRequestParams struct {

}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceResponseParams struct {
	// Available account balance in cents, which takes the same calculation rules as `RealBalance`, `CreditBalance`, and `RealCreditBalance`.
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// The UIN to query.
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// Available account balance in cents, which takes the same calculation rules as `Balance`, `CreditBalance`, and `RealCreditBalance`.
	RealBalance *float64 `json:"RealBalance,omitempty" name:"RealBalance"`

	// Cash account balance in cents. Currently, this field is not applied.
	CashAccountBalance *float64 `json:"CashAccountBalance,omitempty" name:"CashAccountBalance"`

	// Income account balance in cents. Currently, this field is not applied.
	IncomeIntoAccountBalance *float64 `json:"IncomeIntoAccountBalance,omitempty" name:"IncomeIntoAccountBalance"`

	// Present account balance in cents. Currently, this field is not applied.
	PresentAccountBalance *float64 `json:"PresentAccountBalance,omitempty" name:"PresentAccountBalance"`

	// Frozen amount in cents.
	FreezeAmount *float64 `json:"FreezeAmount,omitempty" name:"FreezeAmount"`

	// Overdue amount in cents, which is when the available credit balance is negative.
	OweAmount *float64 `json:"OweAmount,omitempty" name:"OweAmount"`

	// Whether overdue payments are allowed. Currently, this field is not applied.
	IsAllowArrears *bool `json:"IsAllowArrears,omitempty" name:"IsAllowArrears"`

	// Whether you have a credit limit. Currently, this field is not applied.
	IsCreditLimited *bool `json:"IsCreditLimited,omitempty" name:"IsCreditLimited"`

	// Credit limit in cents. Credit limit－available credit balance = consumption amount
	CreditAmount *float64 `json:"CreditAmount,omitempty" name:"CreditAmount"`

	// Available credit balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `RealCreditBalance`.
	CreditBalance *float64 `json:"CreditBalance,omitempty" name:"CreditBalance"`

	// Available account balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `CreditBalance`.
	RealCreditBalance *float64 `json:"RealCreditBalance,omitempty" name:"RealCreditBalance"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountBalanceResponseParams `json:"Response"`
}

func (r *DescribeAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailRequestParams struct {
	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The period type. byUsedTime: By usage period; byPayTime: By payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page. 
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// Month; format: yyyy-mm. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
	Month *string `json:"Month,omitempty" name:"Month"`

	// The start time of the query range, which should be in the format Y-m-d H:i:s . The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use Month to query the billing details of a month.
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the query range, which should be in the format `Y-m-d H:i:s `. The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use `Month` to query the billing details of a month. 
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Action type to query. Valid values:
	// Purchase
	// Renewal
	// Modify
	// Refund
	// Deduction
	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Offline project deduction
	// Offline deduction
	// adjust-CR
	// adjust-DR
	// One-off RI Fee
	// Spot
	// Hourly RI fee
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Project ID: ID of the project to which the resource belongs
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Context information returned by the last request. You can set `Month` to `2023-05` for query by page to accelerate queries. We recommend users whose data volume is over 100 thousand entries use the paginated query feature, which can help greatly speed up your queries.
	Context *string `json:"Context,omitempty" name:"Context"`


	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The period type. byUsedTime: By usage period; byPayTime: By payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page. 
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// Month; format: yyyy-mm. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
	Month *string `json:"Month,omitempty" name:"Month"`

	// The start time of the query range, which should be in the format Y-m-d H:i:s . The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use Month to query the billing details of a month.
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the query range, which should be in the format `Y-m-d H:i:s `. The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use `Month` to query the billing details of a month. 
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Action type to query. Valid values:
	// Purchase
	// Renewal
	// Modify
	// Refund
	// Deduction
	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Offline project deduction
	// Offline deduction
	// adjust-CR
	// adjust-DR
	// One-off RI Fee
	// Spot
	// Hourly RI fee
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Project ID: ID of the project to which the resource belongs
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Context information returned by the last request. You can set `Month` to `2023-05` for query by page to accelerate queries. We recommend users whose data volume is over 100 thousand entries use the paginated query feature, which can help greatly speed up your queries.
	Context *string `json:"Context,omitempty" name:"Context"`

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PeriodType")
	delete(f, "Month")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "NeedRecordNum")
	delete(f, "ProductCode")
	delete(f, "PayMode")
	delete(f, "ResourceId")
	delete(f, "ActionType")
	delete(f, "ProjectId")
	delete(f, "BusinessCode")
	delete(f, "Context")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailResponseParams struct {
	// Details list
	DetailSet []*BillDetail `json:"DetailSet,omitempty" name:"DetailSet"`

	// 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// Context information returned by this request, and the value can be passed in as the value of parameters in the next request to accelerate queries. Note:  This field may return null, indicating that no valid values can be obtained.
	Context *string `json:"Context,omitempty" name:"Context"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillResourceSummaryRequestParams struct {
	// Pagination offset. If `Offset` is `0`, it indicates the first page. If `Limit` is `100`, "`Offset` = `100`" indicates the second page, then "`Offset` = `200`" indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm".  This value must be no earlier than March 2019, when Bill 2.0 was launched.
	Month *string `json:"Month,omitempty" name:"Month"`

	// The period type. byUsedTime: By usage period; byPayTime: by payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// Action type to query. Valid values:
	// Purchase
	// Renewal
	// Modify
	// Refund
	// Deduction
	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Offline project deduction
	// Offline deduction
	// adjust-CR
	// adjust-DR
	// One-off RI Fee
	// Spot
	// Hourly RI fee
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// ID of the instance to be queried
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Billing mode. Valid values: `prePay` (prepaid), `postPay` (postpaid)
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`


	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset. If `Offset` is `0`, it indicates the first page. If `Limit` is `100`, "`Offset` = `100`" indicates the second page, then "`Offset` = `200`" indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm".  This value must be no earlier than March 2019, when Bill 2.0 was launched.
	Month *string `json:"Month,omitempty" name:"Month"`

	// The period type. byUsedTime: By usage period; byPayTime: by payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// Action type to query. Valid values:
	// Purchase
	// Renewal
	// Modify
	// Refund
	// Deduction
	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Offline project deduction
	// Offline deduction
	// adjust-CR
	// adjust-DR
	// One-off RI Fee
	// Spot
	// Hourly RI fee
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// ID of the instance to be queried
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Billing mode. Valid values: `prePay` (prepaid), `postPay` (postpaid)
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeBillResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "NeedRecordNum")
	delete(f, "ActionType")
	delete(f, "ResourceId")
	delete(f, "PayMode")
	delete(f, "BusinessCode")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryResponseParams struct {
	// Resource summary list
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitempty" name:"ResourceSummarySet"`

	// Total number of resource summary lists, which will not be returned when `NeedRecordNum` is `0`. This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillResourceSummaryResponseParams `json:"Response"`
}

func (r *DescribeBillResourceSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeRequestParams struct {
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Query bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Detailed cost distribution for all billing modes
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// A bill type, which corresponds to a subtotal type of L0 bills.
	// This parameter has become valid since v3.0 bills took effect in May 2021.
	// Valid values:
	// `consume`: consumption
	// `refund`: refund
	// `adjustment`: bill adjustment
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// A bill type, which corresponds to a subtotal type of L0 bills.
	// This parameter has become valid since v3.0 bills took effect in May 2021.
	// Valid values:
	// `consume`: consumption
	// `refund`: refund
	// `adjustment`: bill adjustment
	PayType *string `json:"PayType,omitempty" name:"PayType"`
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
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	delete(f, "PayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Total cost details
	// Note: This field may return null, indicating that no valid value was found.
	SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

	// Cost distribution of all products
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillSummaryByProjectRequestParams struct {
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProjectResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Detailed cost distribution for all projects
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByProjectResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionRequestParams struct {
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Detailed cost distribution for all regions
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillSummaryByTagRequestParams struct {
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Cost allocation tag key, which can be customized.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Resource tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Cost allocation tag key, which can be customized.
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Resource tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

func (r *DescribeBillSummaryByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TagKey")
	delete(f, "PayerUin")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByTagResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Details about cost distribution over different tags
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// Total cost
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SummaryTotal *SummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByTagResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryRequestParams struct {
	// Bill month in the format of "2023-04"
	Month *string `json:"Month,omitempty" name:"Month"`

	// Bill dimension. Valid values:  `business`, `project`, `region`, `payMode`, and `tag`
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// Tag key, which is used when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitempty" name:"TagKey"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "2023-04"
	Month *string `json:"Month,omitempty" name:"Month"`

	// Bill dimension. Valid values:  `business`, `project`, `region`, `payMode`, and `tag`
	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`

	// Tag key, which is used when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "GroupType")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryResponseParams struct {
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready.  If `Ready` is `0`, it indicates the current UIN is initializing for the first billing. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// Detailed summary of costs by multiple dimensions
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitempty" name:"SummaryDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageCosDetailByDateRequestParams struct {
	// The start date of the usage query, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// The end date of the usage query (end date must be in the same month as the start date), such as `2020-09-30`.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Bucket name. You can use `Get Service` to query the list of all buckets under a requester account. For details, see [GET Service (List Buckets)](https://www.tencentcloud.com/document/product/436/8291).
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// The start date of the usage query, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// The end date of the usage query (end date must be in the same month as the start date), such as `2020-09-30`.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Bucket name. You can use `Get Service` to query the list of all buckets under a requester account. For details, see [GET Service (List Buckets)](https://www.tencentcloud.com/document/product/436/8291).
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
}

func (r *DescribeDosageCosDetailByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageCosDetailByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "BucketName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDosageCosDetailByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageCosDetailByDateResponseParams struct {
	// Array of usage
	DetailSets []*CosDetailSets `json:"DetailSets,omitempty" name:"DetailSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDosageCosDetailByDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDosageCosDetailByDateResponseParams `json:"Response"`
}

func (r *DescribeDosageCosDetailByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageCosDetailByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoRequestParams struct {
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The voucher status. Valid values: `unUsed`, `used`, `delivered`, `cancel`, `overdue`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// The voucher order ID.
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// The product code.
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// The campaign ID.
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// The voucher name.
	VoucherName *string `json:"VoucherName,omitempty" name:"VoucherName"`

	// The start date of the voucher issuance, such as `2021-01-01`.
	TimeFrom *string `json:"TimeFrom,omitempty" name:"TimeFrom"`

	// The end date of the voucher issuance, such as `2021-01-01`.
	TimeTo *string `json:"TimeTo,omitempty" name:"TimeTo"`

	// The field used to sort the records. Valid values: BeginTime, EndTime, CreateTime.
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// Whether to sort the records in ascending or descending order. Valid values: desc, asc.
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// The payment mode. Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If this parameter is empty or `*`, `productCode` and `subProductCode` must also be empty.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// The operator. The default is the UIN of the current user.
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The voucher status. Valid values: `unUsed`, `used`, `delivered`, `cancel`, `overdue`.
	Status *string `json:"Status,omitempty" name:"Status"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// The voucher order ID.
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// The product code.
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// The campaign ID.
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// The voucher name.
	VoucherName *string `json:"VoucherName,omitempty" name:"VoucherName"`

	// The start date of the voucher issuance, such as `2021-01-01`.
	TimeFrom *string `json:"TimeFrom,omitempty" name:"TimeFrom"`

	// The end date of the voucher issuance, such as `2021-01-01`.
	TimeTo *string `json:"TimeTo,omitempty" name:"TimeTo"`

	// The field used to sort the records. Valid values: BeginTime, EndTime, CreateTime.
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// Whether to sort the records in ascending or descending order. Valid values: desc, asc.
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// The payment mode. Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If this parameter is empty or `*`, `productCode` and `subProductCode` must also be empty.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// The operator. The default is the UIN of the current user.
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeVoucherInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "VoucherId")
	delete(f, "CodeId")
	delete(f, "ProductCode")
	delete(f, "ActivityId")
	delete(f, "VoucherName")
	delete(f, "TimeFrom")
	delete(f, "TimeTo")
	delete(f, "SortField")
	delete(f, "SortOrder")
	delete(f, "PayMode")
	delete(f, "PayScene")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// The total number of vouchers.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The total voucher balance. The value of this parameter is the total balance (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	TotalBalance *int64 `json:"TotalBalance,omitempty" name:"TotalBalance"`

	// The voucher information.
	// Note: This field may return `null`, indicating that no valid value was found.
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitempty" name:"VoucherInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVoucherInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVoucherInfoResponseParams `json:"Response"`
}

func (r *DescribeVoucherInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherUsageDetailsRequestParams struct {
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// The operator. The default is the UIN of the current.
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeVoucherUsageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// The operator. The default is the UIN of the current.
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

func (r *DescribeVoucherUsageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherUsageDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "VoucherId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherUsageDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherUsageDetailsResponseParams struct {
	// The total number of vouchers.
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// The total amount used. The value of this parameter is the total amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	TotalUsedAmount *int64 `json:"TotalUsedAmount,omitempty" name:"TotalUsedAmount"`

	// The usage details.
	// Note: This field may return `null`, indicating that no valid value was found.
	UsageRecords []*UsageRecords `json:"UsageRecords,omitempty" name:"UsageRecords"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVoucherUsageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVoucherUsageDetailsResponseParams `json:"Response"`
}

func (r *DescribeVoucherUsageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherUsageDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExcludedProducts struct {
	// The names of non-applicable products.
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type PayModeSummaryOverviewItem struct {
	// Billing mode code
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Billing mode,  which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// Detailed summary of costs by transaction type
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitempty" name:"Detail"`
}

type ProjectSummaryOverviewItem struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name:  The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type RegionSummaryOverviewItem struct {
	// Region ID
	// Note: This field may return null, indicating that no valid value was found.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region:  The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type SummaryDetail struct {
	// Bill dimension code. Note:  This field may return null, indicating that no valid values can be obtained.
	GroupKey *string `json:"GroupKey,omitempty" name:"GroupKey"`

	// Bill dimension value. Note:  This field may return null, indicating that no valid values can be obtained.
	GroupValue *string `json:"GroupValue,omitempty" name:"GroupValue"`

	// Original cost in USD. This parameter has become valid since Bill 3.0 took effect in May 2021, and before that `-` was returned for this parameter. If a customer has applied for a contract price different from the prices listed on the official website, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// Detailed summary of products. Note:  This field may return null, indicating that no valid values can be obtained.
	Business []*BusinessSummaryInfo `json:"Business,omitempty" name:"Business"`
}

type SummaryTotal struct {
	// Total amount after discount. Note:  This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type TagSummaryOverviewItem struct {
	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Cost percentage rounded to two decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount. Note:  This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cash credit:  The amount paid from the user’s cash account. Note:  This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Free credit:  The amount paid by the user’s free credit. Note:  This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher payment:  The voucher deduction amount. Note:  This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Commission credit:  The amount paid by the user’s commission credit. Note:  This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type UsageDetails struct {
	// The name of the product.
	// Note: This field may return `null`, indicating that no valid value was found.
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`


	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
}

type UsageRecords struct {
	// The amount used. The value of this parameter is the amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	UsedAmount *int64 `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// The time when the voucher was used.
	UsedTime *string `json:"UsedTime,omitempty" name:"UsedTime"`

	// The details of the product purchased.
	// Note: This field may return `null`, indicating that no valid value was found.
	UsageDetails []*UsageDetails `json:"UsageDetails,omitempty" name:"UsageDetails"`
}

type VoucherInfos struct {
	// The owner of the voucher.
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// The status of the voucher: `unUsed`, `used`, `delivered`, `cancel`, `overdue`
	Status *string `json:"Status,omitempty" name:"Status"`

	// The value of the voucher. The value of this parameter is the voucher value (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	NominalValue *int64 `json:"NominalValue,omitempty" name:"NominalValue"`

	// The balance left. The value of this parameter is the balance left (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all.
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// The start time of the validity period.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the validity period.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// The products that are applicable.
	// Note: This field may return `null`, indicating that no valid value was found.
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitempty" name:"ApplicableProducts"`

	// The products that are not applicable.
	// Note: This field may return `null`, indicating that no valid value was found.
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitempty" name:"ExcludedProducts"`
}