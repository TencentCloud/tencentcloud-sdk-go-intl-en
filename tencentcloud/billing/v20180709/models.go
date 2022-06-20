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

	// Transaction type
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// Transaction type name
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Actual cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

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

	// Product name: major categories of Tencent Cloud services, e.g. CVM and TencentDB for MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Sub-product name: sub-categories of Tencent Cloud services, such as CVM-Standard S1
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// Billing mode
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Project: project of a resource
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Region: region of a resource, e.g. South China (Guangzhou)
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Availability zone: availability zone of a resource, e.g. Guangzhou Zone 3
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Instance ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Instance name
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Transaction type
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Order ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// Transaction ID
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// Payment time
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// Service start time
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// Service end time
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// Component list
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitempty" name:"ComponentSet"`

	// Payer's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// User's UIN
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator's UIN
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// Product code
	// Note: This field may return `null`, indicating that no valid value can be found.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Subproduct code
	// Note: This field may return `null`, indicating that no valid value can be found.
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Transaction type
	// Note: This field may return `null`, indicating that no valid value can be found.
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Project ID: ID of the project to which the resource belongs
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BillDetailComponent struct {

	// Component type: type of a resource component, e.g. memory, disk, etc.
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`

	// Component name: name of a resource component, e.g. TencentDB for MySQL-memory
	ItemCodeName *string `json:"ItemCodeName,omitempty" name:"ItemCodeName"`

	// Component published price: original price of a resource component with the original granularity
	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`

	// Specified price of the component
	SpecifiedPrice *string `json:"SpecifiedPrice,omitempty" name:"SpecifiedPrice"`

	// Price unit
	PriceUnit *string `json:"PriceUnit,omitempty" name:"PriceUnit"`

	// Component usage
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// Component usage unit
	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`

	// Usage period
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`

	// Original price of the component
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// Discount rate
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// Total discounted price
	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`

	// Amount paid in voucher
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Amount paid in cash
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Amount paid in trial credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Component type code
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ItemCode *string `json:"ItemCode,omitempty" name:"ItemCode"`

	// Component code
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// Contract price
	ContractPrice *string `json:"ContractPrice,omitempty" name:"ContractPrice"`

	// The special instance (resource pack, reserved instance, savings plan, or spot instance) that is applied to deduction. Valid values:
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The usage duration deducted by a reserved instance. The unit of measurement for deduction is the same as that for usage duration.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	RiTimeSpan *string `json:"RiTimeSpan,omitempty" name:"RiTimeSpan"`

	// The amount deducted by a reserved instance based on the original component cost.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// The discount multiplier that applies to the component based on the remaining commitment of the savings plan.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SPDeductionRate *string `json:"SPDeductionRate,omitempty" name:"SPDeductionRate"`

	// The savings plan deduction amount.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// The amount deducted by a savings plan based on the original component cost.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`

	// The blended discount multiplier that combines the official website discount, reserved instance discount, and savings plan discount. If no reserved instance and savings plan discounts are available, the blended discount multiplier equals the discount multiplier.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BlendedDiscount *string `json:"BlendedDiscount,omitempty" name:"BlendedDiscount"`
}

type BillResourceSummary struct {

	// Product name: major categories of Tencent Cloud services, e.g. CVM and TencentDB for MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Subproduct name, which is the subcategory of a Tencent Cloud product, such as CVM-Standard S1. If no subproduct name can be obtained, `-` is returned.
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// Billing mode
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Project
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Region
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Availability zone
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// Instance ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// Resource instance namDeduction timee
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// Transaction type
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// Order ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// Payment time
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// Service start time
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// Service end time
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// Configuration description
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// Extension field 1
	ExtendField1 *string `json:"ExtendField1,omitempty" name:"ExtendField1"`

	// Extension field 2
	ExtendField2 *string `json:"ExtendField2,omitempty" name:"ExtendField2"`

	// Cost, in USD
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// Discount
	// If different discounts or contract prices are applied, `-` will be returned for this parameter.
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// Total cost after discount, in USD
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Amount paid in voucher, in USD
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Amount paid in cash, in USD
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Amount paid in trial credit, in USD
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Extension field 3
	ExtendField3 *string `json:"ExtendField3,omitempty" name:"ExtendField3"`

	// Extension field 4
	ExtendField4 *string `json:"ExtendField4,omitempty" name:"ExtendField4"`

	// Extension field 5
	ExtendField5 *string `json:"ExtendField5,omitempty" name:"ExtendField5"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Resource owner UIN; '-' is returned if no value is obtained
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator UIN; '-' is returned if no value is obtained
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// The special instance (resource pack, reserved instance, savings plan, or spot instance) that is applied to deduction. Valid values:
	// 
	// ri=Standard RI
	// 
	// svp=Savings Plan
	// 
	// si=Spot Instances
	// 
	// rp=Resource Pack
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// The amount deducted by a reserved instance based on the original component cost.
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// The savings plan deduction amount.
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// The amount deducted by a savings plan based on the original component cost.
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`
}

type BillTagInfo struct {

	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BusinessSummaryOverviewItem struct {

	// Product code
	// Note: This field may return `null`, indicating that no valid value can be found.
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Product name: major categories of Tencent Cloud services, e.g. CVM and TencentDB for MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Actual cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type BusinessSummaryTotal struct {

	// Total cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
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

	// The start time of the period; format: Y-m-d H:i:s. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. BeginTime and EndTime must be inputted as a pair. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The end time of the period; format: Y-m-d H:i:s. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. BeginTime and EndTime must be inputted as a pair. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details list
		DetailSet []*BillDetail `json:"DetailSet,omitempty" name:"DetailSet"`

		// Total number of records
	// Note: This field may return null, indicating that no valid value was found.
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Month; format: yyyy-mm. This value cannot be earlier than the month when Bill 2.0 is enabled. Last 24 months data are available.
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBillResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Resource summary list
		ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitempty" name:"ResourceSummarySet"`

		// Total number of resource summary lists
	// Note: This field may return null, indicating that no valid value was found.
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all billing modes
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Total cost details
	// Note: This field may return null, indicating that no valid value was found.
		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

		// Cost distribution of all products
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all projects
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all regions
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest

	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Cost allocation tag key
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

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. `0`: not ready; `1`: ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Details about cost distribution over different tags
	// Note: This field may return null, indicating that no valid values can be obtained.
		SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

		// Total cost
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		SummaryTotal *SummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// The start time of the promotional campaign.
	TimeFrom *string `json:"TimeFrom,omitempty" name:"TimeFrom"`

	// The end time of the promotional campaign.
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

type DescribeVoucherInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The total number of vouchers.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The total voucher balance. The value of this parameter is the total balance (USD, rounded to 8 decimal places) multiplied by 100,000,000.
		TotalBalance *int64 `json:"TotalBalance,omitempty" name:"TotalBalance"`

		// The voucher information.
	// Note: This field may return `null`, indicating that no valid value was found.
		VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitempty" name:"VoucherInfos"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeVoucherUsageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The total number of vouchers.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The total amount used. The value of this parameter is the total amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
		TotalUsedAmount *int64 `json:"TotalUsedAmount,omitempty" name:"TotalUsedAmount"`

		// The usage details.
	// Note: This field may return `null`, indicating that no valid value was found.
		UsageRecords []*UsageRecords `json:"UsageRecords,omitempty" name:"UsageRecords"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// Billing mode
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// Billing mode name
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// Actual cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Detailed summary of purchases by transaction type
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitempty" name:"Detail"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type ProjectSummaryOverviewItem struct {

	// Project ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// Actual cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type RegionSummaryOverviewItem struct {

	// Region ID
	// Note: This field may return null, indicating that no valid value was found.
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// Actual cost
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type SummaryTotal struct {

	// Total cost
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type TagSummaryOverviewItem struct {

	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// Actual cost
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// Cost percentage rounded to two decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type UsageDetails struct {

	// The name of the product.
	// Note: This field may return `null`, indicating that no valid value was found.
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 
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
