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

package v20180709

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ActionSummaryOverviewItem struct {
	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type AdjustInfoDetail struct {
	// Payer UIN, namely the account ID of the payer. The account ID is the user's unique account identifier on Tencent Cloud.
	// Note: This field may return null, indicating that no valid values can be obtained.
	// Example value: 909619400.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Bill month. Format: yyyy-MM.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// Example value: 2024-10.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Adjustment type.
	// Adjustment: manualAdjustment.
	// Supplementary settlement: supplementarySettlement.
	// Re-settlement: reSettlement.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// Example value: manualAdjustment.
	AdjustType *string `json:"AdjustType,omitnil,omitempty" name:"AdjustType"`

	// Adjustment order number.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// Example value: 2220726096135.
	AdjustNum *string `json:"AdjustNum,omitnil,omitempty" name:"AdjustNum"`

	// Completion time of exception adjustment. Format: yyyy-MM-dd HH:mm:ss.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// Example value: 2022-12-02 12:39:04.
	AdjustCompletionTime *string `json:"AdjustCompletionTime,omitnil,omitempty" name:"AdjustCompletionTime"`

	// Adjustment amount.
	// Note: This field may return null, indicating that no valid value can be obtained.
	// Example value: 333.00000000.
	AdjustAmount *float64 `json:"AdjustAmount,omitnil,omitempty" name:"AdjustAmount"`
}

type AllocationRationExpression struct {
	// Cost allocation unit ID that the sharing rule belongs to.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Sharing proportion occupied by allocation unit, pass 0 for allocation by proportion.
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type AllocationRuleExpression struct {
	// RuleKey: cost allocation dimension.
	// Enumeration value.
	// ownerUin - user UIN.
	// Operator UIN.
	// businessCode - product-level code.
	// productCode - 2-tier product code.
	// itemCode - 4-tier product code.
	// projectId - specifies the project ID.
	// regionId.
	// resourceId - specifies the resource ID.
	// tag - tag key-value pair.
	// payMode - billing mode.
	// instanceType - instance type.
	// actionType - transaction type.
	RuleKey *string `json:"RuleKey,omitnil,omitempty" name:"RuleKey"`

	// Specifies the dimension rules for cost allocation.
	// Enumeration value.
	// in.
	// not in.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Cost allocation dimension value. for example, when RuleKey is businessCode, ["p_cbs","p_sqlserver"] indicates the cost of products at the "p_cbs","p_sqlserver" level.
	RuleValue []*string `json:"RuleValue,omitnil,omitempty" name:"RuleValue"`

	// Logical connection for cost allocation, enumeration values are as follows:.
	// Create and bind a policy query an instance reset the access password of an instance.
	// Create and bind a policy query an instance reset the access password of an instance.
	Connectors *string `json:"Connectors,omitnil,omitempty" name:"Connectors"`

	// Nested rule.
	Children []*AllocationRuleExpression `json:"Children,omitnil,omitempty" name:"Children"`
}

type AllocationRuleOverview struct {
	// Sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Sharing rule name.
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Public area policy type.
	// Enumeration value.
	// 1 - custom sharing proportion. 
	// 2 - proportional allocation. 
	// 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Last update time of the sharing rules.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Overview of cost allocation units.
	AllocationNode []*AllocationUnit `json:"AllocationNode,omitnil,omitempty" name:"AllocationNode"`
}

type AllocationRulesSummary struct {
	// Add new sharing rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Specifies the sharing rule policy type. the enumeration values are as follows:.
	// 1 - custom sharing proportion. 
	// 2 - proportional allocation.
	// 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Sharing rule expression.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Sharing proportion expression, allocation by proportion not passed.
	RatioDetail []*AllocationRationExpression `json:"RatioDetail,omitnil,omitempty" name:"RatioDetail"`
}

type AllocationTree struct {
	// ID of a cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Cost allocation unit name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Subtree.
	Children []*AllocationTree `json:"Children,omitnil,omitempty" name:"Children"`
}

type AllocationUnit struct {
	// Cost allocation unit ID.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Specifies the name of a cost allocation rule.
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`
}

type AnalyseActionTypeDetail struct {
	// Transaction type codeNote: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type nameNote: This field may return null, indicating that no valid values can be obtained.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`
}

type AnalyseAmountDetail struct {
	// Fee typeNote: This field may return null, indicating that no valid values can be obtained.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Whether to displayNote: This field may return null, indicating that no valid values can be obtained.
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`
}

type AnalyseBusinessDetail struct {
	// Product codeNote: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product nameNote: This field may return null, indicating that no valid values can be obtained.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type AnalyseConditionDetail struct {
	// ProductNote: This field may return null, indicating that no valid values can be obtained.
	Business []*AnalyseBusinessDetail `json:"Business,omitnil,omitempty" name:"Business"`

	// ItemNote: This field may return null, indicating that no valid values can be obtained.
	Project []*AnalyseProjectDetail `json:"Project,omitnil,omitempty" name:"Project"`

	// RegionNote: This field may return null, indicating that no valid values can be obtained.
	Region []*AnalyseRegionDetail `json:"Region,omitnil,omitempty" name:"Region"`

	// Billing modeNote: This field may return null, indicating that no valid values can be obtained.
	PayMode []*AnalysePayModeDetail `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction typeNote: This field may return null, indicating that no valid values can be obtained.
	ActionType []*AnalyseActionTypeDetail `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Availability zoneNote: This field may return null, indicating that no valid values can be obtained.
	Zone []*AnalyseZoneDetail `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Resource owner UINNote: This field may return null, indicating that no valid values can be obtained.
	OwnerUin []*AnalyseOwnerUinDetail `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Fee typeNote: This field may return null, indicating that no valid values can be obtained.
	Amount []*AnalyseAmountDetail `json:"Amount,omitnil,omitempty" name:"Amount"`
}

type AnalyseConditions struct {
	// Product name codeNote: This field may return null, indicating that no valid values can be obtained.
	BusinessCodes *string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Sub-product name codeNote: This field may return null, indicating that no valid values can be obtained.
	ProductCodes *string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Component type codeNote: This field may return null, indicating that no valid values can be obtained.
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Availability zone ID: The availability zone ID to which the resource belongsNote: This field may return null, indicating that no valid values can be obtained.
	ZoneIds *string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Region ID: The region ID to which the resource belongsNote: This field may return null, indicating that no valid values can be obtained.
	RegionIds *string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Project ID: The project ID to which the resource belongsNote: This field may return null, indicating that no valid values can be obtained.
	ProjectIds *string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Billing mode prePay (indicates monthly subscription)/postPay (indicates pay-as-you-go billing)Note: This field may return null, indicating that no valid values can be obtained.
	PayModes *string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type. Query transaction type. (Use transaction type code input parameter.)Note: This field may return null, indicating that no valid values can be obtained.
	ActionTypes *string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Cost allocation tag keyNote: This field may return null, indicating that no valid values can be obtained.
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Fee type. Query fee type. (Use fee type code input parameter.) The input parameter enumeration is as follows:cashPayAmount: cash incentivePayAmount: free credits voucherPayAmount: coupons tax:taxes costBeforeTax: price before taxNote: This field may return null, indicating that no valid values can be obtained.
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// User UIN for querying cost analysis dataNote: This field may return null, indicating that no valid values can be obtained.
	PayerUins *string `json:"PayerUins,omitnil,omitempty" name:"PayerUins"`

	// User UIN for using resourcesNote: This field may return null, indicating that no valid values can be obtained.
	OwnerUins *string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Consumption type. Query consumption type. (Use consumption type code input parameter.)Note: This field may return null, indicating that no valid values can be obtained.
	ConsumptionTypes *string `json:"ConsumptionTypes,omitnil,omitempty" name:"ConsumptionTypes"`
}

type AnalyseDetail struct {
	// Time
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Date detailed amountNote: This field may return null, indicating that no valid values can be obtained.
	TimeDetail []*AnalyseTimeDetail `json:"TimeDetail,omitnil,omitempty" name:"TimeDetail"`
}

type AnalyseHeaderDetail struct {
	// Header dateNote: This field may return null, indicating that no valid values can be obtained.
	HeadDetail []*AnalyseHeaderTimeDetail `json:"HeadDetail,omitnil,omitempty" name:"HeadDetail"`

	// TimeNote: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// TotalNote: This field may return null, indicating that no valid values can be obtained.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`
}

type AnalyseHeaderTimeDetail struct {
	// DateNote: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AnalyseOwnerUinDetail struct {
	// User UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type AnalysePayModeDetail struct {
	// Billing mode codeNote: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode nameNote: This field may return null, indicating that no valid values can be obtained.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type AnalyseProjectDetail struct {
	// Project ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Default projectNote: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type AnalyseRegionDetail struct {
	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region nameNote: This field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type AnalyseTimeDetail struct {
	// DateNote: This field may return null, indicating that no valid values can be obtained.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// AmountNote: This field may return null, indicating that no valid values can be obtained.
	Money *string `json:"Money,omitnil,omitempty" name:"Money"`
}

type AnalyseZoneDetail struct {
	// Availability zone IDNote: This field may return null, indicating that no valid values can be obtained.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone nameNote: This field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ApplicableProducts struct {
	// Valid values: `all products` or names of the applicable products (string). Multiple names are separated by commas.
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If `GoodsName` contains multiple product names and `PayMode` is `*`, it indicates that the voucher can be used in all billing modes for each of the products.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type BillDetail struct {
	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct name: The subcategory of a Tencent Cloud product purchased by the user, such as CVM Standard S1.
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project name: The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region: The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Availability zone: availability zone of a resource, e.g. Guangzhou Zone 3
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance ID: The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID: The order number for a monthly subscription purchase
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Transaction ID: The bill number for a deducted payment
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// Transaction time: The time at which a payment was deducted
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Usage start time: The time at which product or service usage starts
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: The time at which product or service usage ends
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Component list
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// Payer account ID: The account ID of the payer, which is the unique identifier of a Tencent Cloud user.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Owner account ID: The account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID: The account or role ID of the operator who purchases or activates a resource
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Tag information. Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Product code. Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code. Note: This field may return null, indicating that no valid values can be obtained.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Transaction type code. Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Region ID. Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Price attribute: A set of attributes which will determine the price of a component, apart from unit price and usage duration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Associated transaction document ID: The ID of the document associated with a transaction, such as a write-off order, the original order showing a deduction error during first settlement, a restructured order, or the original purchase order corresponding to a refund order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Calculation formula: The detailed calculation formula for a specific transaction type, such as refund or configuration change.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing rules: Official website links for detailed billing rules of each product.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Billing dayNote: This field may return null, indicating that no valid values can be obtained.
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`

	// Billing monthNote: This field may return null, indicating that no valid values can be obtained.
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Billing record IDNote: This field may return null, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Domestic and international codesNote: This field may return null, indicating that no valid values can be obtained.
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Domestic and International: The region type to which the resource belongs (domestic, international)Note: This field may return null, indicating that no valid values can be obtained.
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Note attributes (instance configuration): Other note information, such as the reserved instance, the reserved instance type, the transaction type, and the region information on both ends of the CCN product.Note: This field may return null, indicating that no valid values can be obtained.
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// the discount object for the current consumption item, such as official website discount, user discount, and event discount.
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// the discount type for the current consumption item, such as discount and contract price.
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// supplementary description of the discount type, such as 0.2.
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`
}

type BillDetailAssociatedOrder struct {
	// Purchase order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrepayPurchase *string `json:"PrepayPurchase,omitnil,omitempty" name:"PrepayPurchase"`

	// Renewal order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrepayRenew *string `json:"PrepayRenew,omitnil,omitempty" name:"PrepayRenew"`

	// Upgrade order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PrepayModifyUp *string `json:"PrepayModifyUp,omitnil,omitempty" name:"PrepayModifyUp"`

	// Write-off order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReverseOrder *string `json:"ReverseOrder,omitnil,omitempty" name:"ReverseOrder"`

	// The order after discount.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NewOrder *string `json:"NewOrder,omitnil,omitempty" name:"NewOrder"`

	// The original order before discount.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Original *string `json:"Original,omitnil,omitempty" name:"Original"`
}

type BillDetailComponent struct {
	// Component type: The component type of a product or service purchased, such as CVM instance components including CPU and memory.
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// Component name: The specific component of a product or service purchased
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// Component list price: The listed unit price of a component. If a customer has applied for a fixed preferential price or contract price, this parameter will not be displayed by default.
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// Specified price of a component. This parameter has been deprecated.
	//
	// Deprecated: SpecifiedPrice is deprecated.
	SpecifiedPrice *string `json:"SpecifiedPrice,omitnil,omitempty" name:"SpecifiedPrice"`

	// Component price measurement unit: The unit of measurement for a component price, which is composed of USD, usage unit, and duration unit.
	PriceUnit *string `json:"PriceUnit,omitnil,omitempty" name:"PriceUnit"`

	// Component usage: The actually settled usage of a component, which is "Raw usage - Deducted usage (including packages)".
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// Component usage unit: The unit of measurement for component usage
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// Raw usage/duration: The raw usage/duration of a component before deduction. Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// Deducted usage/duration (including packages): The usage/duration deducted with a package. Note: This field may return null, indicating that no valid values can be obtained.
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// Usage duration: The resource usage duration
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit: The unit of measurement for usage duration
	TimeUnitName *string `json:"TimeUnitName,omitnil,omitempty" name:"TimeUnitName"`

	// Original cost: The original cost of a resource, which is "List price x Usage x Usage duration". If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// Discount multiplier: The discount multiplier applied to the cost of the resource. If a customer has applied for a fixed preferential price or contract price or is in a refund scenario, this parameter will not be displayed by default.
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// Total amount after discount: Total amount after discount = (Original cost - RI deduction (cost) - SP deduction (cost)) x Discount multiplier
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Cash credit: The amount paid from the user's cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user's free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Commission credit: The amount paid with the user's commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Component type code. Note: This field may return null, indicating that no valid values can be obtained.
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name code. Note: This field may return null, indicating that no valid values can be obtained.
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component contracted price: The contracted unit price of a component, which is "List price x Discount". Note: This field may return null, indicating that no valid values can be obtained.
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// Instance type: The instance type of a product or service purchased, which can be resource package, RI, SP, or spot instance. Other instance types are not displayed by default. Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RI deduction (duration): The usage duration deducted by RI. Note: This field may return null, indicating that no valid values can be obtained.
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// RI deduction (cost): The amount deducted from the original cost by RI. Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// Savings plan deduction rate: The discount multiplier that applies to the component based on the remaining commitment of the savings plan. Note: This field may return null, indicating that no valid values can be obtained.
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// Cost deduction by SP. This parameter has been deprecated. Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// SP deduction (cost): SP deduction (cost) = Cost deduction by SP / SP deduction rate. Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// Blended discount multiplier: The final discount multiplier that is applied after combining multiple discount types, which is "Total amount after discount / Original cost". Note: This field may return null, indicating that no valid values can be obtained.
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// Configuration description: The specification configuration of an instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentConfig []*BillDetailComponentConfig `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// The tax rate.
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// The tax amount.
	TaxAmount *string `json:"TaxAmount,omitnil,omitempty" name:"TaxAmount"`

	// The currency used for the settlement of a component.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`
}

type BillDetailComponentConfig struct {
	// Configuration description name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Configuration description value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BillDistributionResourceSummary struct {
	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct name: The subcategory of a Tencent Cloud product purchased by the user, such as CVM - Standard S1.
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode: The billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project Name: The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region: The region of a resource, e.g. South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Availability zone: The availability zone of a resource, e.g. Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance ID: The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.	
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, pay-as-you-go deduction, etc.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID: The ID of a monthly subscription order.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Deduction time: The settlement cost deduction time.
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Usage start time: The time at which product or service usage starts.
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: The time at which product or service usage ends.
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Configuration description: The billable item names and usage of a resource, which are displayed on the resource bill only.
	ConfigDesc *string `json:"ConfigDesc,omitnil,omitempty" name:"ConfigDesc"`

	// Extended Field 1: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField1 *string `json:"ExtendField1,omitnil,omitempty" name:"ExtendField1"`

	// Extended field 2: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField2 *string `json:"ExtendField2,omitnil,omitempty" name:"ExtendField2"`

	// Original cost. The original cost of a component = Component price x Usage x Usage duration. If a customer has applied for a fixed preferential price or contract price or if a customer is in a refund scenario, this parameter will not be displayed by default.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Discount multiplier: The discount multiplier that applies to the component. If a customer has applied for a fixed preferential price or contract price or if a customer is in a refund scenario, this parameter will not be displayed by default.
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Offer type.
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// Total amount after discount.
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Voucher payment: The voucher deduction amount.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Cash credit payment: The amount paid through the user's cash account.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit payment: The amount paid with the user's free credit.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Commission credit payment: The amount paid with the user's commission credit.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Extended field 3: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// Extended field 4: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// Extended field 5: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// Tag information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Owner account ID: The account ID of the actual resource user.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID: The account or role ID of the operator who purchases or activates a resource.
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Product code.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Region ID.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Instance type: The instance type of a product or service purchased, which can be resource package, RI, SP, or spot instance. Other instance types are not displayed by default.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RI deduction (cost): The amount deducted from the original cost by RI.	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// Savings plan deduction (disused).
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// SP deduction (cost): The amount of cost deducted by a savings plan based on the component's original cost. SP deduction (cost) = Cost deduction by SP / SP deduction rate	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// Billing monthNote: This field may return null, indicating that no valid values can be obtained.
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillResourceSummary struct {
	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct name: The subcategory of a Tencent Cloud product purchased by the user, such as CVM Computing C5t.
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project name: The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region: The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Availability zone: The availability zone to which a resource belongs, such as Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance ID: The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.	
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID: The order number for a monthly subscription purchase
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Transaction time: The time at which a payment was deducted
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Usage start time: The time at which product or service usage starts
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: The time at which product or service usage ends
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Configuration description: The billable item names and usage of a resource, which are displayed on the resource bill only.
	ConfigDesc *string `json:"ConfigDesc,omitnil,omitempty" name:"ConfigDesc"`

	// Extended field 1: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField1 *string `json:"ExtendField1,omitnil,omitempty" name:"ExtendField1"`

	// Extended field 2: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField2 *string `json:"ExtendField2,omitnil,omitempty" name:"ExtendField2"`

	// Original cost: The original cost of a resource, which is "List price x Usage x Usage duration". If a customer has applied for a fixed preferential price or contract price or applied for a refund, this parameter will not be displayed by default.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Discount multiplier: The discount multiplier applied to the cost of the resource. If a customer has applied for a fixed preferential price or contract price or applied for a refund, this parameter will not be displayed by default.
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Offer type
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// Total amount after discount (Including Tax):  = Total Amount After Discount (Excluding Tax) + TaxAmount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Cash credit: The amount paid from the user's cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user's free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Commission credit: The amount paid with the user's commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Extended field 3: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// Extended field 4: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// Extended field 5: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// Tag information. Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Payer account ID: The account ID of the payer, which is the unique identifier of a Tencent Cloud user.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Owner account ID: The account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID: The account or role ID of the operator who purchases or activates a resource.
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Instance type: The instance type of a product or service purchased, which can be resource package, RI, SP, or spot instance. Other instance types are not displayed by default.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RI deduction (cost): The amount deducted from the original cost by RI	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// Cost deduction by SP. This parameter has been deprecated.
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// SP deduction (cost): SP deduction (cost) = Cost deduction by SP / SP deduction rate	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// Billing monthNote: This field may return null, indicating that no valid values can be obtained.
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillTagInfo struct {
	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type BusinessSummaryInfo struct {
	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Original cost in USD. This parameter became valid when Bill 3.0 took effect in May 2021. Before that, `-` was returned for this parameter. If a customer has applied for a contract price different from the prices listed on the official website, `-` will also be returned for this parameter. 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type BusinessSummaryOverviewItem struct {
	// Product code. Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type BusinessSummaryTotal struct {
	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type ConditionBusiness struct {
	// Product name code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type ConditionPayMode struct {
	// Payment mode
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Payment mode name
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type ConditionProject struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type ConditionRegion struct {
	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type Conditions struct {
	// Only supports two values: 6 and 12.
	TimeRange *uint64 `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// Product name code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Payment mode. Options include prePay and postPay.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Resource keyword
	ResourceKeyword *string `json:"ResourceKeyword,omitnil,omitempty" name:"ResourceKeyword"`

	// Product name code
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Subproduct name code
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Project ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Payment mode. Options include prePay and postPay.
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Whether to hide zero-amount transactions
	HideFreeCost *int64 `json:"HideFreeCost,omitnil,omitempty" name:"HideFreeCost"`

	// Sorting rule. Options include desc and asc.
	OrderByCost *string `json:"OrderByCost,omitnil,omitempty" name:"OrderByCost"`

	// Transaction ID
	BillIds []*string `json:"BillIds,omitnil,omitempty" name:"BillIds"`

	// Component code
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// File ID
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// File type
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// Status
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConsumptionBusinessSummaryDataItem struct {
	// Product name code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cost trend
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Cash
	// Note: This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Bonus
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// VoucherNote: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Share revenueNote: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Region name (only shown in regional summary)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type ConsumptionProjectSummaryDataItem struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Trend
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Product consumption details
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil,omitempty" name:"Business"`

	// Cash
	// Note: This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Bonus
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// VoucherNote: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Share revenueNote: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type ConsumptionRegionSummaryDataItem struct {
	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Trend
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Product consumption details
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil,omitempty" name:"Business"`

	// Cash
	// Note: This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// VoucherNote: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Bonus
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Share revenueNote: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type ConsumptionResourceSummaryConditionValue struct {
	// Product list
	Business []*ConditionBusiness `json:"Business,omitnil,omitempty" name:"Business"`

	// Project list
	Project []*ConditionProject `json:"Project,omitnil,omitempty" name:"Project"`

	// Region list
	Region []*ConditionRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// Payment mode list
	PayMode []*ConditionPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type ConsumptionResourceSummaryDataItem struct {
	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource name
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash expenditure
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Payment mode
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Payment mode name
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Product name code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Consumption type
	ConsumptionTypeName *string `json:"ConsumptionTypeName,omitnil,omitempty" name:"ConsumptionTypeName"`

	// Pre-discount priceNote: This field may return null, indicating that no valid values can be obtained.
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// Start time of feesNote: This field may return null, indicating that no valid values can be obtained.
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// End time of feesNote: This field may return null, indicating that no valid values can be obtained.
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Days
	// Note: This field may return null, indicating that no valid values can be obtained.
	DayDiff *string `json:"DayDiff,omitnil,omitempty" name:"DayDiff"`

	// Daily consumptionNote: This field may return null, indicating that no valid values can be obtained.
	DailyTotalCost *string `json:"DailyTotalCost,omitnil,omitempty" name:"DailyTotalCost"`

	// Order numberNote: This field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// VoucherNote: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Bonus
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Share revenueNote: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Payer UIN: the account ID of the payer, which is the unique identifier of a Tencent Cloud userNote: This field may return null, indicating that no valid values can be obtained.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: the account ID of the actual resource userNote: This field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator UIN: the account ID or role ID of the operator who places orders for prepaid resources or activates postpaid resourcesNote: This field may return null, indicating that no valid values can be obtained.
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Subproduct codeNote: This field may return null, indicating that no valid values can be obtained.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: the subcategory of a product purchased by the user, such as CVM – Standard S1Note: This field may return null, indicating that no valid values can be obtained.
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Region typeNote: This field may return null, indicating that no valid values can be obtained.
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Region type nameNote: This field may return null, indicating that no valid values can be obtained.
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Extended field 1Note: This field may return null, indicating that no valid values can be obtained.
	Extend1 *string `json:"Extend1,omitnil,omitempty" name:"Extend1"`

	// Extended field 2Note: This field may return null, indicating that no valid values can be obtained.
	Extend2 *string `json:"Extend2,omitnil,omitempty" name:"Extend2"`

	// Extended field 3Note: This field may return null, indicating that no valid values can be obtained.
	Extend3 *string `json:"Extend3,omitnil,omitempty" name:"Extend3"`

	// Extended field 4Note: This field may return null, indicating that no valid values can be obtained.
	Extend4 *string `json:"Extend4,omitnil,omitempty" name:"Extend4"`

	// Extended field 5Note: This field may return null, indicating that no valid values can be obtained.
	Extend5 *string `json:"Extend5,omitnil,omitempty" name:"Extend5"`

	// Instance typeNote: This field may return null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type nameNote: This field may return null, indicating that no valid values can be obtained.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Deduction time: the time at which a payment is deductedNote: This field may return null, indicating that no valid values can be obtained.
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Availability zone: availability zone of a resource, e.g. Guangzhou Zone 3Note: This field may return null, indicating that no valid values can be obtained.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Configuration descriptionNote: This field may return null, indicating that no valid values can be obtained.
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// Tag information.Note: This field may return null, indicating that no valid values can be obtained.
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ConsumptionSummaryTotal struct {
	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type ConsumptionSummaryTrend struct {
	// Trend type, upward for rising, downward for falling, none for no change
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Trend value. The value is null when Type is none.Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CosDetailSets struct {
	// Bucket name
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// The start time of the usage
	DosageBeginTime *string `json:"DosageBeginTime,omitnil,omitempty" name:"DosageBeginTime"`

	// The end time of the usage
	DosageEndTime *string `json:"DosageEndTime,omitnil,omitempty" name:"DosageEndTime"`

	// Subproduct name
	SubProductCodeName *string `json:"SubProductCodeName,omitnil,omitempty" name:"SubProductCodeName"`

	// Billable item name
	BillingItemCodeName *string `json:"BillingItemCodeName,omitnil,omitempty" name:"BillingItemCodeName"`

	// Usage
	DosageValue *string `json:"DosageValue,omitnil,omitempty" name:"DosageValue"`

	// Unit of the billable item
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type CostComponentSet struct {
	// Component type name
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// Component name
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// List price
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// List price unit
	PriceUnit *string `json:"PriceUnit,omitnil,omitempty" name:"PriceUnit"`

	// Usage
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// Usage unit
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// Original price
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// Discount
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Discounted price
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// Voucher payment amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Cash payment amount
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Bonus payment amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`
}

type CostDetail struct {
	// Payer UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct name
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode name
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project name
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region Name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Zone name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Resource ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource name
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Type nameNote: This field may return null, indicating that no valid values can be obtained.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Transaction ID
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// Start time of fees
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// End time of fees
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Component details
	ComponentSet []*CostComponentSet `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// Subproduct name code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`
}

// Predefined struct for user
type CreateAllocationRuleRequestParams struct {
	// List of sharing rules.
	RuleList *AllocationRulesSummary `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type CreateAllocationRuleRequest struct {
	*tchttp.BaseRequest
	
	// List of sharing rules.
	RuleList *AllocationRulesSummary `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *CreateAllocationRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleList")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAllocationRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationRuleResponseParams struct {
	// Add new sharing rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAllocationRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAllocationRuleResponseParams `json:"Response"`
}

func (r *CreateAllocationRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationTagRequestParams struct {
	// Cost allocation tag key.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type CreateAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// Cost allocation tag key.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *CreateAllocationTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAllocationTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationTagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAllocationTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateAllocationTagResponseParams `json:"Response"`
}

func (r *CreateAllocationTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationUnitRequestParams struct {
	// New cost allocation unit parent node ID.
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Specifies the name of a newly-added cost allocation unit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type CreateAllocationUnitRequest struct {
	*tchttp.BaseRequest
	
	// New cost allocation unit parent node ID.
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Specifies the name of a newly-added cost allocation unit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *CreateAllocationUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentId")
	delete(f, "Name")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAllocationUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationUnitResponseParams struct {
	// Specifies the ID of a newly-added cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAllocationUnitResponse struct {
	*tchttp.BaseResponse
	Response *CreateAllocationUnitResponseParams `json:"Response"`
}

func (r *CreateAllocationUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGatherRuleRequestParams struct {
	// Cost allocation unit ID that the rule belongs to.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Collection rule details.
	RuleList *GatherRuleSummary `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type CreateGatherRuleRequest struct {
	*tchttp.BaseRequest
	
	// Cost allocation unit ID that the rule belongs to.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Collection rule details.
	RuleList *GatherRuleSummary `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *CreateGatherRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatherRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "RuleList")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGatherRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGatherRuleResponseParams struct {
	// Collection  rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGatherRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateGatherRuleResponseParams `json:"Response"`
}

func (r *CreateGatherRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGatherRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Deal struct {
	// Order ID.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// The status of the order. 1: unpaid; 2: paid; 3: shipping; 4: shipped; 5: shipment failed; 6: refunded; 7: closed case; 8: order expired; 9: order invalidated; 10: product invalidated; 11: third-party payment refused; 12: payment in process
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Payer
	Payer *string `json:"Payer,omitnil,omitempty" name:"Payer"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Creator
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Actual payment amount (pent)
	RealTotalCost *int64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Voucher offset amount (pent)
	VoucherDecline *int64 `json:"VoucherDecline,omitnil,omitempty" name:"VoucherDecline"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product category ID
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// Product details
	ProductInfo []*ProductInfo `json:"ProductInfo,omitnil,omitempty" name:"ProductInfo"`

	// Duration
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Time unit
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Currency unit
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Discount rate
	Policy *float64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// Unit price (cents)
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Original price (cents)
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Product code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct code
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Large order number.
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// Refund formula.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Refund involves order information.
	RefReturnDeals *string `json:"RefReturnDeals,omitnil,omitempty" name:"RefReturnDeals"`

	// Billing mode: `prePay` (prepaid), `postPay` (pay-as-you-go), `riPay` (reserved instance)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type
	// 
	// Modify network mode adjust bandwidth mode.
	// Adjust bandwidth size.
	// `Refund`: refund.
	// downgrade.
	// upgrade configuration.
	// renew.
	// purchase.
	// preMoveOut monthly subscription resource migration out.
	// preMoveIn specifies the monthly subscription resources to migrate.
	// preToPost specifies the prepaid to postpaid conversion.
	// postMoveOut specifies the pay-as-you-go resources to be migrated out.
	// postMoveIn specifies the pay-as-you-go resources for inbound migration.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Product code chinese name.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Sub-Product code chinese name.
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// The resource ID corresponding to the order. If the query parameter `Limit` exceeds 200, null will be returned.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceId []*string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type DeleteAllocationRuleRequestParams struct {
	// The deleted sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DeleteAllocationRuleRequest struct {
	*tchttp.BaseRequest
	
	// The deleted sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DeleteAllocationRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllocationRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAllocationRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllocationRuleResponseParams `json:"Response"`
}

func (r *DeleteAllocationRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationTagRequestParams struct {
	// Cost allocation tag key
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DeleteAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// Cost allocation tag key
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *DeleteAllocationTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllocationTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationTagResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAllocationTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllocationTagResponseParams `json:"Response"`
}

func (r *DeleteAllocationTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationUnitRequestParams struct {
	// Specifies the deleted cost allocation unit ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DeleteAllocationUnitRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the deleted cost allocation unit ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DeleteAllocationUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllocationUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationUnitResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAllocationUnitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllocationUnitResponseParams `json:"Response"`
}

func (r *DeleteAllocationUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatherRuleRequestParams struct {
	// The deleted collection rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DeleteGatherRuleRequest struct {
	*tchttp.BaseRequest
	
	// The deleted collection rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DeleteGatherRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatherRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGatherRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGatherRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGatherRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGatherRuleResponseParams `json:"Response"`
}

func (r *DeleteGatherRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGatherRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// The UIN to query.
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Available account balance in cents, which takes the same calculation rules as `Balance`, `CreditBalance`, and `RealCreditBalance`.
	RealBalance *float64 `json:"RealBalance,omitnil,omitempty" name:"RealBalance"`

	// Cash account balance in cents. Currently, this field is not applied.
	CashAccountBalance *float64 `json:"CashAccountBalance,omitnil,omitempty" name:"CashAccountBalance"`

	// Income account balance in cents. Currently, this field is not applied.
	IncomeIntoAccountBalance *float64 `json:"IncomeIntoAccountBalance,omitnil,omitempty" name:"IncomeIntoAccountBalance"`

	// Present account balance in cents. Currently, this field is not applied.
	PresentAccountBalance *float64 `json:"PresentAccountBalance,omitnil,omitempty" name:"PresentAccountBalance"`

	// Frozen amount in cents.
	FreezeAmount *float64 `json:"FreezeAmount,omitnil,omitempty" name:"FreezeAmount"`

	// Overdue amount in cents, which is when the available credit balance is negative.
	OweAmount *float64 `json:"OweAmount,omitnil,omitempty" name:"OweAmount"`

	// Whether overdue payments are allowed. Currently, this field is not applied.
	IsAllowArrears *bool `json:"IsAllowArrears,omitnil,omitempty" name:"IsAllowArrears"`

	// Whether you have a credit limit. Currently, this field is not applied.
	IsCreditLimited *bool `json:"IsCreditLimited,omitnil,omitempty" name:"IsCreditLimited"`

	// Credit limit in cents. Credit limit－available credit balance = consumption amount
	CreditAmount *float64 `json:"CreditAmount,omitnil,omitempty" name:"CreditAmount"`

	// Available credit balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `RealCreditBalance`.
	CreditBalance *float64 `json:"CreditBalance,omitnil,omitempty" name:"CreditBalance"`

	// Available account balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `CreditBalance`.
	RealCreditBalance *float64 `json:"RealCreditBalance,omitnil,omitempty" name:"RealCreditBalance"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAllocationRuleDetailRequestParams struct {
	// The queried sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocationRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// The queried sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocationRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationRuleDetailResponseParams struct {
	// Sharing rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Sharing rule ownership UIN.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Sharing rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Specifies the public area policy type. the enumeration values are as follows:.
	// 1 - custom sharing proportion. 
	// 2 - proportional allocation. 
	// 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Public sharing rule expression.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Sharing proportion expression. returns when Type is 1 or 2.
	RatioDetail []*AllocationRationExpression `json:"RatioDetail,omitnil,omitempty" name:"RatioDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeAllocationRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationRuleSummaryRequestParams struct {
	// Specifies the data quantity per fetch. the maximum value is 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Public area policy type, for filtering.
	// Enumeration values are as follows:. 
	// 1 - custom sharing proportion. 
	// 2 - proportional allocation. 
	// 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Sharing rule name or cost allocation unit name, used for fuzzy filter criteria.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeAllocationRuleSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the data quantity per fetch. the maximum value is 1000.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Public area policy type, for filtering.
	// Enumeration values are as follows:. 
	// 1 - custom sharing proportion. 
	// 2 - proportional allocation. 
	// 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Sharing rule name or cost allocation unit name, used for fuzzy filter criteria.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeAllocationRuleSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationRuleSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "Type")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationRuleSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationRuleSummaryResponseParams struct {
	// Sharing rule expression.
	RuleList []*AllocationRuleOverview `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Total number of rules.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationRuleSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationRuleSummaryResponseParams `json:"Response"`
}

func (r *DescribeAllocationRuleSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationRuleSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationTreeRequestParams struct {
	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocationTreeRequest struct {
	*tchttp.BaseRequest
	
	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocationTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationTreeResponseParams struct {
	// Cost allocation unit ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Specifies the name of a cost allocation unit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Specifies a subtree.
	Children []*AllocationTree `json:"Children,omitnil,omitempty" name:"Children"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationTreeResponseParams `json:"Response"`
}

func (r *DescribeAllocationTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationUnitDetailRequestParams struct {
	// ID of the queried cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocationUnitDetailRequest struct {
	*tchttp.BaseRequest
	
	// ID of the queried cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocationUnitDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationUnitDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationUnitDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationUnitDetailResponseParams struct {
	// ID of a cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Associated UIN of the cost allocation unit.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Specifies the name of a cost allocation unit.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Cost allocation unit parent node ID.
	ParentId *uint64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Source organization name.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Source organization ID.
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// Specifies remark description.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Cost allocation unit identifier.
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// If a cost allocation unit is set with an collection rule, return the collection rule ID. if no collection rule is set, do not return.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationUnitDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationUnitDetailResponseParams `json:"Response"`
}

func (r *DescribeAllocationUnitDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationUnitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillAdjustInfoRequestParams struct {
	// Format: yyyy-MM.
	// Billing month. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If the TimeFrom and TimeTo are passed, the Month field is invalid.
	// Example: 2024-10.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Format: yyyy-MM-dd.
	// Start date. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If TimeFrom and TimeTo are passed, the Month field is invalid. TimeFrom and TimeTo should represent the same month and be passed in together. Cross-month queries are not supported. The result will include the full month's data.
	// Example: 2024-10-01.
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// Format: yyyy-MM-dd.
	// End date. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If TimeFrom and TimeTo are passed, the Month field is invalid. TimeFrom and TimeTo should represent the same month and be passed in together. Cross-month queries are not supported. The result will include the full month's data.
	// Example: 2024-10-02.
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`
}

type DescribeBillAdjustInfoRequest struct {
	*tchttp.BaseRequest
	
	// Format: yyyy-MM.
	// Billing month. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If the TimeFrom and TimeTo are passed, the Month field is invalid.
	// Example: 2024-10.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Format: yyyy-MM-dd.
	// Start date. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If TimeFrom and TimeTo are passed, the Month field is invalid. TimeFrom and TimeTo should represent the same month and be passed in together. Cross-month queries are not supported. The result will include the full month's data.
	// Example: 2024-10-01.
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// Format: yyyy-MM-dd.
	// End date. Either Month or the combination of TimeFrom and TimeTo needs to be passed. If TimeFrom and TimeTo are passed, the Month field is invalid. TimeFrom and TimeTo should represent the same month and be passed in together. Cross-month queries are not supported. The result will include the full month's data.
	// Example: 2024-10-02.
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`
}

func (r *DescribeBillAdjustInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillAdjustInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TimeFrom")
	delete(f, "TimeTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillAdjustInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillAdjustInfoResponseParams struct {
	// Total amount of data.
	// Example value: 10.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Detailed data.
	// Example value: [].
	Data []*AdjustInfoDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillAdjustInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillAdjustInfoResponseParams `json:"Response"`
}

func (r *DescribeBillAdjustInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillAdjustInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailForOrganizationRequestParams struct {
	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Cycle type, which can be `byUsedTime` (by billing cycle) or `byPayTime` (by deduction time). This value must be the same as the billing period type in Billing Center for that particular month. You can check your billing cycle at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicates whether the total number of records is required, used for pagination.
	// Valid values: `1` (required), `0` (not required).
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Billing mode, which can be `prePay` (monthly subscription) or `postPay` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Monthly subscription purchase
	// Monthly subscription renewal
	// Monthly subscription upgrade/downgrade
	// Monthly subscription refund 
	// Pay-as-you-go deduction 
	// Offline project deduction 
	// Offline product deduction 
	// Adjustment deduction 
	// Adjustment compensation 
	// Hourly pay-as-you-go 
	// Daily pay-as-you-go 
	// Monthly pay-as-you-go 
	// Hourly spot instance 
	// Offline project adjustment compensation 
	// Offline product adjustment compensation 
	// Offer deduction 
	// Offer compensation 
	// Pay-as-you-go resource migration in 
	// Pay-as-you-go resource migration out 
	// Monthly subscription resource migration in 
	// Monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Monthly subscription to pay-as-you-go 
	// Minimum spend deduction 
	// Hourly savings plan fee
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: The ID of the project to which the resource belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product code.
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned by the last response. You can view multiple pages when querying for data after May 2023 to speed up the query. We recommend you use this query method if your data volume is above 100 thousand entries, which can improve query speed by 2-10 times.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type DescribeBillDetailForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Cycle type, which can be `byUsedTime` (by billing cycle) or `byPayTime` (by deduction time). This value must be the same as the billing period type in Billing Center for that particular month. You can check your billing cycle at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicates whether the total number of records is required, used for pagination.
	// Valid values: `1` (required), `0` (not required).
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Billing mode, which can be `prePay` (monthly subscription) or `postPay` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Monthly subscription purchase
	// Monthly subscription renewal
	// Monthly subscription upgrade/downgrade
	// Monthly subscription refund 
	// Pay-as-you-go deduction 
	// Offline project deduction 
	// Offline product deduction 
	// Adjustment deduction 
	// Adjustment compensation 
	// Hourly pay-as-you-go 
	// Daily pay-as-you-go 
	// Monthly pay-as-you-go 
	// Hourly spot instance 
	// Offline project adjustment compensation 
	// Offline product adjustment compensation 
	// Offer deduction 
	// Offer compensation 
	// Pay-as-you-go resource migration in 
	// Pay-as-you-go resource migration out 
	// Monthly subscription resource migration in 
	// Monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Monthly subscription to pay-as-you-go 
	// Minimum spend deduction 
	// Hourly savings plan fee
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: The ID of the project to which the resource belongs.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product code.
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned by the last response. You can view multiple pages when querying for data after May 2023 to speed up the query. We recommend you use this query method if your data volume is above 100 thousand entries, which can improve query speed by 2-10 times.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

func (r *DescribeBillDetailForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailForOrganizationRequest) FromJsonString(s string) error {
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
	delete(f, "PayMode")
	delete(f, "ResourceId")
	delete(f, "ActionType")
	delete(f, "ProjectId")
	delete(f, "BusinessCode")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailForOrganizationResponseParams struct {
	// Details list.
	DetailSet []*DistributionBillDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// Total number of records, which is cached every 24 hours and may be less than the actual total number of records.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Context information of the current request, which can be used in the parameters of the next request to speed up the query.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillDetailForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDetailForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillDetailForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailRequestParams struct {
	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity, maximum is 300
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The period type. byUsedTime: By usage period; byPayTime: By payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page. 
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Month; format: yyyy-mm. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the query range, which should be in the format Y-m-d H:i:s . The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use Month to query the billing details of a month.
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the query range, which should be in the format `Y-m-d H:i:s `. The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use `Month` to query the billing details of a month. 
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Spot
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: ID of the project to which the resource belongs
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned by the last request. You can set `Month` to `2023-05` or later to accelerate queries. We recommend users whose data volume is over 100 thousand entries use the paginated query feature, which can help greatly speed up your queries.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// The account ID of the payer, which is the unique identifier of a Tencent Cloud user. This account is allowed to query its own bills by default. If an organization admin account needs to query the self-pay bills of members, this field should be specified as the member account ID.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity, maximum is 300
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The period type. byUsedTime: By usage period; byPayTime: By payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page. 
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Month; format: yyyy-mm. You only have to enter either Month or BeginTime and EndTime. When you enter values for BeginTime and EndTime, Month becomes invalid. This value must be no earlier than the month when Bill 2.0 is activated; last 24 months data are available.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the query range, which should be in the format Y-m-d H:i:s . The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use Month to query the billing details of a month.
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the query range, which should be in the format `Y-m-d H:i:s `. The query range must be in the last 18 months and cannot be earlier than May 2018 (when Bill 2.0 was introduced). The start time and end time must be in the same month.
	// 
	// Example: tccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --BeginTime '2023-04-01 12:05:15' --EndTime '2023-04-18 12:00:10' --ProjectId 1000000731  --version "2018-07-09"
	// 
	// Alternatively, you can use `Month` to query the billing details of a month. 
	// Example:
	// ccli billing DescribeBillDetail --cli-unfold-argument --Offset 1 --Limit 100 --Month 2023-04  --version "2018-07-09" --ResourceId "disk-oj9okstm"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Spot
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: ID of the project to which the resource belongs
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned by the last request. You can set `Month` to `2023-05` or later to accelerate queries. We recommend users whose data volume is over 100 thousand entries use the paginated query feature, which can help greatly speed up your queries.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// The account ID of the payer, which is the unique identifier of a Tencent Cloud user. This account is allowed to query its own bills by default. If an organization admin account needs to query the self-pay bills of members, this field should be specified as the member account ID.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	DetailSet []*BillDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBillDownloadUrlRequestParams struct {
	// Bill type. Valid values:
	// `billOverview` (L0: PDF bills)
	// `billSummary` (L1: Bill summary)	
	// `billResource` (L2: Bill by instance)	
	// `billDetail` (L3: Bill details)	
	// `billPack` (Bill packs)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Bill month.
	// The earliest month that can be queried is January 2021.
	// L0 bills and bill packs cannot be downloaded for the current month. Please download the current month's bills after it is generated at 19:00 on the 1st day of the next month.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// List of account IDs for downloading the bill. By default, it queries the bill for the current account. If you are an admin account and need to download bills for member accounts with their own payment, input the member account's UIN for this parameter.
	ChildUin []*string `json:"ChildUin,omitnil,omitempty" name:"ChildUin"`
}

type DescribeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Bill type. Valid values:
	// `billOverview` (L0: PDF bills)
	// `billSummary` (L1: Bill summary)	
	// `billResource` (L2: Bill by instance)	
	// `billDetail` (L3: Bill details)	
	// `billPack` (Bill packs)
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Bill month.
	// The earliest month that can be queried is January 2021.
	// L0 bills and bill packs cannot be downloaded for the current month. Please download the current month's bills after it is generated at 19:00 on the 1st day of the next month.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// List of account IDs for downloading the bill. By default, it queries the bill for the current account. If you are an admin account and need to download bills for member accounts with their own payment, input the member account's UIN for this parameter.
	ChildUin []*string `json:"ChildUin,omitnil,omitempty" name:"ChildUin"`
}

func (r *DescribeBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "Month")
	delete(f, "ChildUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDownloadUrlResponseParams struct {
	// Indicates whether the bill file is ready. Valid values: `0` (the file is being generated), `1` (the file has been generated).
	Ready *int64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Billing file download link, valid for 1 day. Note: This field may return null, indicating that no valid values can be obtained.
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryForOrganizationRequestParams struct {
	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm". This value must be no earlier than the month when Bill 2.0 is activated.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Cycle type, which can be `byUsedTime` (by billing cycle) or `byPayTime` (by deduction time). This value must be the same as the billing period type in Billing Center for that particular month. You can check your billing cycle at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Indicates whether the total number of records is required, used for pagination.
	// Valid values: `1` (required), `0` (not required).
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Monthly subscription purchase
	// Monthly subscription renewal
	// Monthly subscription upgrade/downgrade
	// Monthly subscription refund 
	// Pay-as-you-go deduction 
	// Offline project deduction 
	// Offline product deduction 
	// Adjustment deduction 
	// Adjustment compensation 
	// Hourly pay-as-you-go 
	// Daily pay-as-you-go 
	// Monthly pay-as-you-go 
	// Hourly spot instance 
	// Offline project adjustment compensation 
	// Offline product adjustment compensation 
	// Offer deduction 
	// Offer compensation 
	// Pay-as-you-go resource migration in 
	// Pay-as-you-go resource migration out 
	// Monthly subscription resource migration in 
	// Monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Monthly subscription to pay-as-you-go 
	// Minimum spend deduction 
	// Hourly savings plan fee
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Billing mode. Valid values: `prePay`, `postPay`.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes (`BusinessCode`) used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Cost allocation tag key, which can be customized. This parameter can be used for querying bills after January 2021.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Resource tag value. If it is left empty, there are no records with tag values set under this tag key.
	// This parameter can be used for querying bills after January 2021.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillResourceSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm". This value must be no earlier than the month when Bill 2.0 is activated.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Cycle type, which can be `byUsedTime` (by billing cycle) or `byPayTime` (by deduction time). This value must be the same as the billing period type in Billing Center for that particular month. You can check your billing cycle at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Indicates whether the total number of records is required, used for pagination.
	// Valid values: `1` (required), `0` (not required).
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Monthly subscription purchase
	// Monthly subscription renewal
	// Monthly subscription upgrade/downgrade
	// Monthly subscription refund 
	// Pay-as-you-go deduction 
	// Offline project deduction 
	// Offline product deduction 
	// Adjustment deduction 
	// Adjustment compensation 
	// Hourly pay-as-you-go 
	// Daily pay-as-you-go 
	// Monthly pay-as-you-go 
	// Hourly spot instance 
	// Offline project adjustment compensation 
	// Offline product adjustment compensation 
	// Offer deduction 
	// Offer compensation 
	// Pay-as-you-go resource migration in 
	// Pay-as-you-go resource migration out 
	// Monthly subscription resource migration in 
	// Monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Monthly subscription to pay-as-you-go 
	// Minimum spend deduction 
	// Hourly savings plan fee
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Billing mode. Valid values: `prePay`, `postPay`.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes (`BusinessCode`) used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Cost allocation tag key, which can be customized. This parameter can be used for querying bills after January 2021.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Resource tag value. If it is left empty, there are no records with tag values set under this tag key.
	// This parameter can be used for querying bills after January 2021.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

func (r *DescribeBillResourceSummaryForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryForOrganizationRequest) FromJsonString(s string) error {
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
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryForOrganizationResponseParams struct {
	// Resource summary list.
	ResourceSummarySet []*BillDistributionResourceSummary `json:"ResourceSummarySet,omitnil,omitempty" name:"ResourceSummarySet"`

	// Total number of resource summary lists. It will not be returned if `NeedRecordNum` is `0`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillResourceSummaryForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillResourceSummaryForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillResourceSummaryForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryRequestParams struct {
	// Pagination offset. If `Offset` is `0`, it indicates the first page. If `Limit` is `100`, "`Offset` = `100`" indicates the second page, then "`Offset` = `200`" indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm". This value must be no earlier than March 2019, when Bill 2.0 was launched.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The period type. byUsedTime
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Spot
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// ID of the instance to be queried
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// The account ID of the payer, which is the unique identifier of a Tencent Cloud user. This account is allowed to query its own bills by default. If an organization admin account needs to query the self-pay bills of members, this field should be specified as the member account ID.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Cost allocation tag key, which can be customized. This parameter can be used for querying bills after January 2021.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Resource tag value. If it is left empty, there are no records with tag values set under this tag key.
	// This parameter can be used for querying bills after January 2021.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset. If `Offset` is `0`, it indicates the first page. If `Limit` is `100`, "`Offset` = `100`" indicates the second page, then "`Offset` = `200`" indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Bill month in the format of "yyyy-mm". This value must be no earlier than March 2019, when Bill 2.0 was launched.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The period type. byUsedTime
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Hourly settlement
	// Daily settlement
	// Monthly settlement
	// Spot
	// New monthly subscription
	// Monthly subscription renewal
	// Monthly subscription specification adjustment
	// Monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// ID of the instance to be queried
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product code
	// Note: To query the product codes used in the current month, call <a href="https://intl.cloud.tencent.com/document/product/555/35761?from_cn_redirect=1">DescribeBillSummaryByProduct</a>.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// The account ID of the payer, which is the unique identifier of a Tencent Cloud user. This account is allowed to query its own bills by default. If an organization admin account needs to query the self-pay bills of members, this field should be specified as the member account ID.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Cost allocation tag key, which can be customized. This parameter can be used for querying bills after January 2021.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Resource tag value. If it is left empty, there are no records with tag values set under this tag key.
	// This parameter can be used for querying bills after January 2021.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryResponseParams struct {
	// Resource summary list
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitnil,omitempty" name:"ResourceSummarySet"`

	// Total number of resource summary lists, which will not be returned when `NeedRecordNum` is `0`. This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Query bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Detailed cost distribution for all billing modes
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// A bill type, which corresponds to a subtotal type of L0 bills.
	// This parameter has become valid since v3.0 bills took effect in May 2021.
	// Valid values:
	// `consume`: consumption
	// `refund`: refund
	// `adjustment`: bill adjustment
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// A bill type, which corresponds to a subtotal type of L0 bills.
	// This parameter has become valid since v3.0 bills took effect in May 2021.
	// Valid values:
	// `consume`: consumption
	// `refund`: refund
	// `adjustment`: bill adjustment
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Total cost details
	// Note: This field may return null, indicating that no valid value was found.
	SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitnil,omitempty" name:"SummaryTotal"`

	// Cost distribution of all products
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Detailed cost distribution for all projects
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Detailed cost distribution for all regions
	// Note: This field may return null, indicating that no valid value was found.
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Cost allocation tag key, which can be customized.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Resource tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. Query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Cost allocation tag key, which can be customized.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Resource tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Details about cost distribution over different tags
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// Total cost
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SummaryTotal *SummaryTotal `json:"SummaryTotal,omitnil,omitempty" name:"SummaryTotal"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBillSummaryForOrganizationRequestParams struct {
	// Bill month in the format of "yyyy-mm".
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Bill dimension. Valid values: `business`, `project`, `region`, `payMode`, and `tag`.
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag key. Pass in it when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeBillSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-mm".
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Bill dimension. Valid values: `business`, `project`, `region`, `payMode`, and `tag`.
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag key. Pass in it when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryForOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "GroupType")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryForOrganizationResponseParams struct {
	// Indicates whether the data is ready. Valid values: `0` (not ready), `1` (ready). If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Bills summarized by multiple dimensions.
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil,omitempty" name:"SummaryDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillSummaryForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryRequestParams struct {
	// Bill month in the format of "yyyy-mm"
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Bill dimension. Valid values: `business`, `project`, `region`, `payMode`, and `tag`
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag key, which is used when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-mm"
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Bill dimension. Valid values: `business`, `project`, `region`, `payMode`, and `tag`
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag key, which is used when `GroupType` is `tag`.
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
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
	// Indicates whether the data is ready. `0`: Not ready. `1`: Ready. If `Ready` is `0`, it indicates that the current UIN is initializing billing for the first time. Wait for 5-10 minutes and try again.
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Detailed summary of costs by multiple dimensions
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil,omitempty" name:"SummaryDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCostDetailRequestParams struct {
	// The number of entries returned at a time. The maximum value is `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cycle start time in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after cost analysis is activated and within the past 24 months.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Cycle end time in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after cost analysis is activated and within the past 24 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether the total number of records in the list is needed, for frontend pagination1: needed, 0: not needed
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Month, in the format of yyyy-mm. Either Month or BeginTime&EndTime must be entered, and if BeginTime&EndTime is entered, Month becomes invalid. It cannot be earlier than the month when cost analysis is activated. Data of up to 24 months can be retrieved.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Used to query information of a specified product (currently not available)
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Payment mode. Options include prePay and postPay.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Used to query information of a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeCostDetailRequest struct {
	*tchttp.BaseRequest
	
	// The number of entries returned at a time. The maximum value is `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cycle start time in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after cost analysis is activated and within the past 24 months.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Cycle end time in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after cost analysis is activated and within the past 24 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether the total number of records in the list is needed, for frontend pagination1: needed, 0: not needed
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Month, in the format of yyyy-mm. Either Month or BeginTime&EndTime must be entered, and if BeginTime&EndTime is entered, Month becomes invalid. It cannot be earlier than the month when cost analysis is activated. Data of up to 24 months can be retrieved.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Used to query information of a specified product (currently not available)
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Payment mode. Options include prePay and postPay.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Used to query information of a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DescribeCostDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "NeedRecordNum")
	delete(f, "Month")
	delete(f, "ProductCode")
	delete(f, "PayMode")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostDetailResponseParams struct {
	// Consumption details
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailSet []*CostDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// Record countNote: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostDetailResponseParams `json:"Response"`
}

func (r *DescribeCostDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostExplorerSummaryRequestParams struct {
	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Bill type: 1-cost bill, 2-consumption bill
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Statistical period: day-day, month-month;
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Classification dimension (data aggregation dimension). Query classification dimension. (Use classification dimension code input parameter.) Input parameter enumeration value:default = Total only
	// feeType = Fee typebillType = Bill typebusiness = Product
	// product = Sub-productregion=Region
	// zone = Availability zoneactionType = Transaction typepayMode = Billing modetags = Tagproject = ProjectpayerUin = Payer accountownerUin = User account
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Fee type: cost-total cost, totalCost-original price cost
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// Quantity. The maximum value per page is 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Starting page, where PageNo=1 indicates the first page, PageNo=2 indicates the second page, and so on.
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Cost allocation tag value
	TagKeyStr *string `json:"TagKeyStr,omitnil,omitempty" name:"TagKeyStr"`

	// Whether the filter box is needed: 1- indicates it is needed, 0- indicates it is not needed. If it is not specified, it is not required by default.
	NeedConditionValue *string `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// Filter parameters
	Conditions *AnalyseConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type DescribeCostExplorerSummaryRequest struct {
	*tchttp.BaseRequest
	
	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Bill type: 1-cost bill, 2-consumption bill
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// Statistical period: day-day, month-month;
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Classification dimension (data aggregation dimension). Query classification dimension. (Use classification dimension code input parameter.) Input parameter enumeration value:default = Total only
	// feeType = Fee typebillType = Bill typebusiness = Product
	// product = Sub-productregion=Region
	// zone = Availability zoneactionType = Transaction typepayMode = Billing modetags = Tagproject = ProjectpayerUin = Payer accountownerUin = User account
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// Fee type: cost-total cost, totalCost-original price cost
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// Quantity. The maximum value per page is 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Starting page, where PageNo=1 indicates the first page, PageNo=2 indicates the second page, and so on.
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Cost allocation tag value
	TagKeyStr *string `json:"TagKeyStr,omitnil,omitempty" name:"TagKeyStr"`

	// Whether the filter box is needed: 1- indicates it is needed, 0- indicates it is not needed. If it is not specified, it is not required by default.
	NeedConditionValue *string `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// Filter parameters
	Conditions *AnalyseConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

func (r *DescribeCostExplorerSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostExplorerSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "BillType")
	delete(f, "PeriodType")
	delete(f, "Dimensions")
	delete(f, "FeeType")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "TagKeyStr")
	delete(f, "NeedConditionValue")
	delete(f, "Conditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostExplorerSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostExplorerSummaryResponseParams struct {
	// Number of data entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Header informationNote: This field may return null, indicating that no valid values can be obtained.
	Header *AnalyseHeaderDetail `json:"Header,omitnil,omitempty" name:"Header"`

	// Data detailsNote: This field may return null, indicating that no valid values can be obtained.
	Detail []*AnalyseDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Data amountNote: This field may return null, indicating that no valid values can be obtained.
	TotalDetail *AnalyseDetail `json:"TotalDetail,omitnil,omitempty" name:"TotalDetail"`

	// Filter boxNote: This field may return null, indicating that no valid values can be obtained.
	ConditionValue *AnalyseConditionDetail `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostExplorerSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostExplorerSummaryResponseParams `json:"Response"`
}

func (r *DescribeCostExplorerSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostExplorerSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProductRequestParams struct {
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProductResponseParams struct {
	// Data readiness, 0 for not ready, 1 for ready
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Consumption details
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Consumption details summarized by productNote: This field may return null, indicating that no valid values can be obtained.
	Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// Record count. The system returns null when NeedRecordNum is 0.Note: This field may return null, indicating that no valid values can be obtained.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByProductResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProjectRequestParams struct {
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProjectResponseParams struct {
	// Data readiness, 0 for not ready, 1 for ready
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Consumption details
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Consumption details summarized by business
	Data []*ConsumptionProjectSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// Record count. The system returns null when NeedRecordNum is 0.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByProjectResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByRegionRequestParams struct {
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as `EndTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as `BeginTime`. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both `BeginTime` and `EndTime` are `2018-09`, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByRegionResponseParams struct {
	// Data readiness, 0 for not ready, 1 for ready
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Consumption details
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Consumption details summarized by region
	Data []*ConsumptionRegionSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// Record count. The system returns null when NeedRecordNum is 0.Note: This field may return null, indicating that no valid values can be obtained.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByRegionResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByResourceRequestParams struct {
	// The value must be of the same month as EndTime. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both BeginTime and EndTime are 2018-09, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as BeginTime. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both BeginTime and EndTime are 2018-09, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Whether to return filter criteria. 0 for no, 1 for yes. Default is no.
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// Filter criteria. It only supports ResourceKeyword (resource keyword, which supports fuzzy query by resource ID or resource name), ProjectIds (project ID), RegionIds (region ID), PayModes (payment mode, prePay or postPay), HideFreeCost (whether to hide zero-amount transactions, 0 or 1), and OrderByCost (sorting rule by fees, desc or asc).
	Conditions *Conditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type DescribeCostSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// The value must be of the same month as EndTime. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both BeginTime and EndTime are 2018-09, the data returned will be for the entire month of September 2018.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The value must be of the same month as BeginTime. The query period must start and end on the same month and the query result returned will be of the entire month. For example, if both BeginTime and EndTime are 2018-09, the data returned will be for the entire month of September 2018.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Data quantity per fetch. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset, which starts from 0 by default
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// UIN of the user querying the bill data
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Whether to return the record count. 0 for no, 1 for yes. Default is no.
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Whether to return filter criteria. 0 for no, 1 for yes. Default is no.
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// Filter criteria. It only supports ResourceKeyword (resource keyword, which supports fuzzy query by resource ID or resource name), ProjectIds (project ID), RegionIds (region ID), PayModes (payment mode, prePay or postPay), HideFreeCost (whether to hide zero-amount transactions, 0 or 1), and OrderByCost (sorting rule by fees, desc or asc).
	Conditions *Conditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

func (r *DescribeCostSummaryByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	delete(f, "NeedConditionValue")
	delete(f, "Conditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByResourceResponseParams struct {
	// Data readiness, 0 for not ready, 1 for ready
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Consumption detailsNote: This field may return null, indicating that no valid values can be obtained.
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Filter criteria
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConditionValue *ConsumptionResourceSummaryConditionValue `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`

	// Record countNote: This field may return null, indicating that no valid values can be obtained.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Resource consumption detailsNote: This field may return null, indicating that no valid values can be obtained.
	Data []*ConsumptionResourceSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByResourceResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDealsByCondRequestParams struct {
	// Start time 
	// Example :2016-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time 
	// Example:2016-02-01 00:00:00. 
	// It is recommended that the span does not exceed 3 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from, starting from 0. The default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Order status, default is 4 (successful order)
	// Status of the order
	// 1: unpaid
	// 2: paid 
	// 3: shipment in progress
	// 4: shipped
	// 5`: Shipment Failed
	// 6`: Refunded
	// 7`: Ticket closed
	// 8`: Order expired
	// 9`: Order invalid
	// 10: product invalidated
	// 11: third-party payment refused
	// 12: payment in process
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sub-order number
	// Example: 202202021234567
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Large order number.
	// Example: 202202021234566
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// Resource ID
	// Example:ins-a2bb34
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeDealsByCondRequest struct {
	*tchttp.BaseRequest
	
	// Start time 
	// Example :2016-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time 
	// Example:2016-02-01 00:00:00. 
	// It is recommended that the span does not exceed 3 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from, starting from 0. The default is 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Order status, default is 4 (successful order)
	// Status of the order
	// 1: unpaid
	// 2: paid 
	// 3: shipment in progress
	// 4: shipped
	// 5`: Shipment Failed
	// 6`: Refunded
	// 7`: Ticket closed
	// 8`: Order expired
	// 9`: Order invalid
	// 10: product invalidated
	// 11: third-party payment refused
	// 12: payment in process
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sub-order number
	// Example: 202202021234567
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Large order number.
	// Example: 202202021234566
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// Resource ID
	// Example:ins-a2bb34
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DescribeDealsByCondRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDealsByCondRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "OrderId")
	delete(f, "BigDealId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDealsByCondRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDealsByCondResponseParams struct {
	// Order list
	Deals []*Deal `json:"Deals,omitnil,omitempty" name:"Deals"`

	// Total number of orders
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDealsByCondResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDealsByCondResponseParams `json:"Response"`
}

func (r *DescribeDealsByCondResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDealsByCondResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageCosDetailByDateRequestParams struct {
	// The start date of the usage query, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// The end date of the usage query (end date must be in the same month as the start date), such as `2020-09-30`.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Bucket name. You can use `Get Service` to query the list of all buckets under a requester account. For details, see [GET Service (List Buckets)](https://www.tencentcloud.com/document/product/436/8291).
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// The start date of the usage query, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// The end date of the usage query (end date must be in the same month as the start date), such as `2020-09-30`.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Bucket name. You can use `Get Service` to query the list of all buckets under a requester account. For details, see [GET Service (List Buckets)](https://www.tencentcloud.com/document/product/436/8291).
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
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
	DetailSets []*CosDetailSets `json:"DetailSets,omitnil,omitempty" name:"DetailSets"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeGatherRuleDetailRequestParams struct {
	// Specifies the ID of the queried collection rule.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeGatherRuleDetailRequest struct {
	*tchttp.BaseRequest
	
	// Specifies the ID of the queried collection rule.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeGatherRuleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherRuleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatherRuleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatherRuleDetailResponseParams struct {
	// Specifies the collection rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Associated UIN of the collection rule.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Collection rule last update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Collection rule details.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatherRuleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatherRuleDetailResponseParams `json:"Response"`
}

func (r *DescribeGatherRuleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherRuleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagListRequestParams struct {
	// The number of entries returned at a time. The maximum value is `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cost allocation tag key, used for fuzzy search.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type, used for tag filtering. Valid values: `0` (general tags), `1` (cost allocation tags). If it is not specified, all tag keys will be queried.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting order. Valid values: `asc` (ascending order), `desc` (descending order).
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeTagListRequest struct {
	*tchttp.BaseRequest
	
	// The number of entries returned at a time. The maximum value is `1000`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If `Offset` is `0`, it indicates the first page. When `Limit` is `100`, if `Offset` is `100`, it indicates the second page; if `Offset` is `200`, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cost allocation tag key, used for fuzzy search.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type, used for tag filtering. Valid values: `0` (general tags), `1` (cost allocation tags). If it is not specified, all tag keys will be queried.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting order. Valid values: `asc` (ascending order), `desc` (descending order).
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *DescribeTagListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TagKey")
	delete(f, "Status")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagListResponseParams struct {
	// Total number of records.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Tag information.
	Data []*TagDataInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagListResponseParams `json:"Response"`
}

func (r *DescribeTagListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoRequestParams struct {
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The voucher status. Valid values: `unUsed`, `used`, `delivered`, `cancel`, `overdue`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// The voucher order ID.
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// The product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// The campaign ID.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The voucher name.
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// The start date of the voucher issuance, such as `2021-01-01`.
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// The end date of the voucher issuance, such as `2021-01-01`.
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// The field used to sort the records. Valid values: BeginTime, EndTime, CreateTime.
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// Whether to sort the records in ascending or descending order. Valid values: desc, asc.
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// The payment mode. Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If this parameter is empty or `*`, `productCode` and `subProductCode` must also be empty.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// The operator. The default is the UIN of the current user.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The primary types of vouchers are has_price and no_price, which represent the cash voucher with a price and the cash voucher without a price respectively.
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// Voucher subtype: Discount is a discount voucher, and deduct is a deduction voucher.
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The voucher status. Valid values: `unUsed`, `used`, `delivered`, `cancel`, `overdue`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// The voucher order ID.
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// The product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// The campaign ID.
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// The voucher name.
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// The start date of the voucher issuance, such as `2021-01-01`.
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// The end date of the voucher issuance, such as `2021-01-01`.
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// The field used to sort the records. Valid values: BeginTime, EndTime, CreateTime.
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// Whether to sort the records in ascending or descending order. Valid values: desc, asc.
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// The payment mode. Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If this parameter is empty or `*`, `productCode` and `subProductCode` must also be empty.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// The operator. The default is the UIN of the current user.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// The primary types of vouchers are has_price and no_price, which represent the cash voucher with a price and the cash voucher without a price respectively.
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// Voucher subtype: Discount is a discount voucher, and deduct is a deduction voucher.
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`
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
	delete(f, "VoucherMainType")
	delete(f, "VoucherSubType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// The total number of vouchers.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The total voucher balance. The value of this parameter is the total balance (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	TotalBalance *int64 `json:"TotalBalance,omitnil,omitempty" name:"TotalBalance"`

	// The voucher information.
	// Note: This field may return `null`, indicating that no valid value was found.
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitnil,omitempty" name:"VoucherInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// The operator. The default is the UIN of the current.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeVoucherUsageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// The number of records per page. The default is 20, and the maximum is 1,000.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The page number the records start from. The default is 1.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// The operator. The default is the UIN of the current.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The total amount used. The value of this parameter is the total amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	TotalUsedAmount *int64 `json:"TotalUsedAmount,omitnil,omitempty" name:"TotalUsedAmount"`

	// The usage details.
	// Note: This field may return `null`, indicating that no valid value was found.
	UsageRecords []*UsageRecords `json:"UsageRecords,omitnil,omitempty" name:"UsageRecords"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DistributionBillDetail struct {
	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct name: The subcategory of a Tencent Cloud product purchased by the user, such as CVM - Standard S1.
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode: The billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project Name: The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region: The region of a resource, e.g. South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Availability zone: availability zone of a resource, e.g. Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance ID: The object ID of a billed resource, such as a CVM instance ID. This object ID may vary due to various forms and contents of resources in different products.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The resource name set by the user in the console. If it is not set, it will be empty by default.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Transaction type, which can be monthly subscription purchase, monthly subscription renewal, pay-as-you-go deduction, etc.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID: The ID of a monthly subscription order.
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Transaction ID: The ID of a settlement bill.
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// Deduction time: The settlement cost deduction time.
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Usage start time: The time at which product or service usage starts.
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: The time at which product or service usage ends.
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// List of components.
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// Owner account ID: The account ID of the actual resource user.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID: The account or role ID of the operator who purchases or activates a resource.
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Tag information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Product code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Transaction type code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Region ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Price attribute: A set of attributes which will determine the price of a component, apart from unit price and usage duration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Associated transaction document ID: The ID of the document associated with a transaction, such as a write-off order, the original order showing a deduction error during first settlement, a restructured order, or the original purchase order corresponding to a refund order.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Calculation formula: The detailed calculation formula for a specific transaction type, such as refund or configuration change.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing rules: Official website links for detailed billing rules of each product.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Billing monthNote: This field may return null, indicating that no valid values can be obtained.
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Billing dayNote: This field may return null, indicating that no valid values can be obtained.
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`
}

type ExcludedProducts struct {
	// The names of non-applicable products.
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type GatherRuleSummary struct {
	// Cost allocation regular expression.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`
}

// Predefined struct for user
type ModifyAllocationRuleRequestParams struct {
	// The edited sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Edited sharing rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Public sharing policy types, enumeration values are as follows: 1 - custom sharing proportion 2 - proportional allocation 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Edited share rules expression.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Edited sharing proportion expression.
	RatioDetail []*AllocationRationExpression `json:"RatioDetail,omitnil,omitempty" name:"RatioDetail"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type ModifyAllocationRuleRequest struct {
	*tchttp.BaseRequest
	
	// The edited sharing rule ID.
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Edited sharing rule name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Public sharing policy types, enumeration values are as follows: 1 - custom sharing proportion 2 - proportional allocation 3 - allocation by proportion.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Edited share rules expression.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Edited sharing proportion expression.
	RatioDetail []*AllocationRationExpression `json:"RatioDetail,omitnil,omitempty" name:"RatioDetail"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *ModifyAllocationRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllocationRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "RuleDetail")
	delete(f, "RatioDetail")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllocationRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllocationRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAllocationRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllocationRuleResponseParams `json:"Response"`
}

func (r *ModifyAllocationRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllocationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllocationUnitRequestParams struct {
	// ID of the modified cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Modified cost allocation unit name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Modified cost allocation unit source organization name.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Modified allocation unit source organization ID.
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// Cost allocation unit remark description.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type ModifyAllocationUnitRequest struct {
	*tchttp.BaseRequest
	
	// ID of the modified cost allocation unit.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Modified cost allocation unit name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Modified cost allocation unit source organization name.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Modified allocation unit source organization ID.
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// Cost allocation unit remark description.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *ModifyAllocationUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllocationUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "SourceName")
	delete(f, "SourceId")
	delete(f, "Remark")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllocationUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllocationUnitResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAllocationUnitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllocationUnitResponseParams `json:"Response"`
}

func (r *ModifyAllocationUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllocationUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatherRuleRequestParams struct {
	// Edit collection rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Rule detail of the edited collection rule.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type ModifyGatherRuleRequest struct {
	*tchttp.BaseRequest
	
	// Edit collection rule ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Rule detail of the edited collection rule.
	RuleDetail *AllocationRuleExpression `json:"RuleDetail,omitnil,omitempty" name:"RuleDetail"`

	// Month, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *ModifyGatherRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatherRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "RuleDetail")
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGatherRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatherRuleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGatherRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGatherRuleResponseParams `json:"Response"`
}

func (r *ModifyGatherRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatherRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryOverviewItem struct {
	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash balance
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Detailed summary of costs by transaction type
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type ProductInfo struct {
	// Product detail name identifier
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Product details
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ProjectSummaryOverviewItem struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The project to which a resource belongs, which is user-designated. If a resource has not been assigned to a project, it will automatically belong to the default project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type RegionSummaryOverviewItem struct {
	// Region ID
	// Note: This field may return null, indicating that no valid value was found.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region: The region to which a resource belongs, such as South China (Guangzhou).
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type SummaryDetail struct {
	// Bill dimension code. Note: This field may return null, indicating that no valid values can be obtained.
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	// Bill dimension value. Note: This field may return null, indicating that no valid values can be obtained.
	GroupValue *string `json:"GroupValue,omitnil,omitempty" name:"GroupValue"`

	// Original cost in USD. This parameter has become valid since Bill 3.0 took effect in May 2021, and before that `-` was returned for this parameter. If a customer has applied for a contract price different from the prices listed on the official website, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Detailed summary of products. Note: This field may return null, indicating that no valid values can be obtained.
	Business []*BusinessSummaryInfo `json:"Business,omitnil,omitempty" name:"Business"`
}

type SummaryTotal struct {
	// Total amount after discount. Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type TagDataInfo struct {
	// Cost allocation tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type. Valid values: `0` (general tags), `1` (cost allocation tags).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Time to set the cost allocation tag. It will not be returned if `Status` is `0`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TagSummaryOverviewItem struct {
	// Tag value
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// Cost percentage rounded to two decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount. Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account. Note: This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit. Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The amount deducted by using vouchers. Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Commission credit: The amount paid with the user’s commission credit. Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type UsageDetails struct {
	// The name of the product.
	// Note: This field may return `null`, indicating that no valid value was found.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`


	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`
}

type UsageRecords struct {
	// The amount used. The value of this parameter is the amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// The time when the voucher was used.
	UsedTime *string `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// The details of the product purchased.
	// Note: This field may return `null`, indicating that no valid value was found.
	UsageDetails []*UsageDetails `json:"UsageDetails,omitnil,omitempty" name:"UsageDetails"`
}

type VoucherInfos struct {
	// The owner of the voucher.
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// The status of the voucher: `unUsed`, `used`, `delivered`, `cancel`, `overdue`
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// The value of the voucher. The value of this parameter is the voucher value (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	NominalValue *int64 `json:"NominalValue,omitnil,omitempty" name:"NominalValue"`

	// The balance left. The value of this parameter is the balance left (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// The voucher ID.
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// If `PayMode` is `postPay`, this parameter may be `spotpay` (spot instance) or `settle account` (regular pay-as-you-go). If `PayMode` is `prePay`, this parameter may be `purchase`, `renew`, or `modify` (downgrade/upgrade). If `PayMode` is `riPay`, this parameter may be `oneOffFee` (prepayment of reserved instance) or `hourlyFee` (hourly billing of reserved instance). `*` means to query vouchers that support all billing scenarios.
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// The start time of the validity period.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the validity period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The products that are applicable.
	// Note: This field may return `null`, indicating that no valid value was found.
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitnil,omitempty" name:"ApplicableProducts"`

	// The products that are not applicable.
	// Note: This field may return `null`, indicating that no valid value was found.
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitnil,omitempty" name:"ExcludedProducts"`
}