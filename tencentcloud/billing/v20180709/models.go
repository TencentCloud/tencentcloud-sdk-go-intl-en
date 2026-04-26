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

	// Transaction type, which can be yearly/monthly subscription purchase, yearly/monthly subscription renewal, or pay-as-you-go deduction.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Cost ratio, to two decimal points
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user's cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user's free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Billing month, e.g. `2019-08`
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// The original cost in USD. This parameter has become valid since v3.0 bills took effect in May 2021, and before that `-` was returned for this parameter. If a customer uses a contract price different from the published price, `-` will also be returned for this parameter.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type AdjustInfoDetail struct {
	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// Bill month, formatted as yyyy-MM.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Adjustment type
	// Bill adjustment: manualAdjustment
	// Supplementary settlement: supplementarySettlement
	// Re-settlement
	AdjustType *string `json:"AdjustType,omitnil,omitempty" name:"AdjustType"`

	// Adjustment Number
	AdjustNum *string `json:"AdjustNum,omitnil,omitempty" name:"AdjustNum"`

	// Abnormal adjustment completion time. Format: yyyy-MM-dd HH:mm:ss
	AdjustCompletionTime *string `json:"AdjustCompletionTime,omitnil,omitempty" name:"AdjustCompletionTime"`

	// Adjustment Amount
	AdjustAmount *float64 `json:"AdjustAmount,omitnil,omitempty" name:"AdjustAmount"`
}

type AllocationAverageData struct {
	// Start month
	BeginMonth *string `json:"BeginMonth,omitnil,omitempty" name:"BeginMonth"`

	// End month.
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// Average value of total fees (discounted total)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type AllocationBillTrendDetail struct {
	// Bill month
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Displayed name of bill month
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Total fees (discounted total): Total fees of the cost allocation unit, Collected Fees (discounted total) + Allocated Fees (discounted total)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type AllocationDetail struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Date: Settlement date
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID (the resource account ID or role ID opened by prepaid resource ordering or postpaid operation)
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode: Resource billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The Project to which a resource belongs, which is independently allocated by the user for the resource in the console. If a resource has not been allocated to an Project, it will be a default Project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name: The region where the resource is located
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// AZ ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone: The availability zone where the resource is located.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Resource ID: Resources vary by product, and the content is not identical. For example, Cloud Virtual Machine (CVM) corresponds to the instance ID. If the product is split, it shows the split item ID, such as COS bucket ID and CDN domain name.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The name set by the user for the resource in the console, which is empty by default if not set. If the product is split, it shows the split resource alias.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Instance type code
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. It is displayed as "-" by default for regular instances.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Split item ID: The ID of the split item involved in the split product, such as COS bucket ID and CDN domain name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// Split item name: The split item involved in the split product
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type: Detailed transaction type
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Order ID: The order number for purchase in the annual and monthly billing mode
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Transaction ID: The settlement and deduction number
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// Deduction time: Deduction time
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Usage start time: Usage start time
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: Product or service usage end time
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component type: The major component category corresponding to the product or service purchased by the user
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// Component list price: The original unit price of the component on the portal (not displayed if the customer enjoys a fixed price/contract price)
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// Component unit price: Discounted unit price of the component. Component unit price = list price * discount.
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// Component Price Unit: Unit of component price, Unit Composition: CNY/usage unit/duration unit
	SinglePriceUnit *string `json:"SinglePriceUnit,omitnil,omitempty" name:"SinglePriceUnit"`

	// Component usage: The actual settlement usage of the component, Component Usage = Original Component Usage - Deducted Usage (including resource packages)
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// Component usage unit: Unit of measurement corresponding to component usage.
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// Usage duration: The duration of resource usage, Component Usage = Original Component Usage Duration - Deducted Duration (including resource packages)
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit: Unit of resource usage duration.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Remark attribute (instance configuration): Additional remark information, such as reserved instance type and transaction type for reserved instances, regional information of both ends for CCN products.
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// Split item usage/duration ratio: Split item usage (duration) ratio, Split Item Usage (Duration) /Total Usage Before Splitting (Duration)
	SplitRatio *string `json:"SplitRatio,omitnil,omitempty" name:"SplitRatio"`

	// Original price of a component: Original Price = Component List Price * Component Usage * Duration of Use (not displayed if the customer enjoys a fixed price/contract price, and not displayed by default in refund scenarios), specified price mode
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Reserved instance deduction duration: The duration of use deducted by reserved instances for this product or service.
	RITimeSpan *string `json:"RITimeSpan,omitnil,omitempty" name:"RITimeSpan"`

	// Original price deducted by a reserved instance: The original price of a component deducted by reserved instances for this product or service
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// Savings plan deduction from original price: Savings Plan Deduction from Original Price = Monetary Value of Savings Plan Deduction/ Savings Plan Deduction Rate
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// Discount rate: The discount rate enjoyed by this resource (it is not shown by default if the customer enjoys a fixed/contract price, and it is also not shown by default in the refund scenario)
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Mixed discount rate: The final discount rate after integrating various discount deductions. Mixed Discount Rate = Discounted total price/Original Price.
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// Discounted total: discounted total = (Original Price - Original Price Deducted by a Reserved Instance - Savings Plan Deduction from Original Price) * Discount Rate
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash account expenditure (CNY): The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Promo voucher expenditure (CNY): The amount paid using various vouchers (such as promo vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Gift account expenditure (CNY): The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty account expenditure (CNY): The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Allocation tag: The resource-bound tag
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Domestic and international codes
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Domestic and international: Resource region type (domestic, international)
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Component name code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name: The specific component of a product or service purchased by the user
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// Associated document ID: Document ID associated with this transaction, such as the original new purchase order corresponding to a refund order
	AssociatedOrder *string `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Price attribute: Other attributes of the component that affect discount pricing besides unit price and duration
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Calculation rule explanation: A detailed explanation to calculations of billing settlement for special transaction types, such as refund and configuration changes.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing Rules: The detailed billing rules for each product shown in the portal explanation link
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Original usage/duration: The original usage of the component before deduction by resource packages.
	// (Currently only TRTC, TEM, Cloud Call Center, and CDZ products support this information display. Other products are being integrated.)
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// Deduction of usage/duration (including resource packages): The amount of usage deducted by resource packages
	// (Currently only TRTC, TEM, Cloud Call Center, and CDZ products support this information display. Other products are being integrated.)
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// Configuration description: Information on specification of resource configuration
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// Cost collection type: The source types of fees, including allocated, collection and unallocated.
	// 0 - Allocation
	// 1 - Collection
	// 2 - Unallocated
	AllocationType *uint64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// CostBeforeTax
	CostBeforeTax *string `json:"CostBeforeTax,omitnil,omitempty" name:"CostBeforeTax"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// AmountBeforeTax
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`

	// Discount object of the current consumption item, such as official website discount, user discount and activity discount.
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// Discount type of the current consumption item, such as discount and contract price.
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// Supplementary description of the offer type, for example: business discount 20% off, the offer type is "discount" and the discount content is "0.8".
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`

	// SPDeduction
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// SPDeduction
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// Currency
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`
}

type AllocationMonthOverviewDetail struct {
	// Collected fees (cash): Cash directly collected to the cost allocation unit based on the collection rules
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// Collected fees (voucher): Resource vouchers directly collected to the cost allocation unit based on the collection rules
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// Collected fees (free credit): Resource free credit directly collected to the cost allocation unit based on the collection rules
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// Collected fees (royalty amount): Resource royalty amount directly collected to the cost allocation unit based on the collection rules
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// Allocated fees (cash): Resource cash allocated to the cost allocation unit based on the allocation rules
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// Allocated fees (vouchers): Resource vouchers allocated to the cost allocation unit based on the allocation rules
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// Allocated fees (free credit): Resource free credit allocated to the cost allocation unit based on the allocation rules
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// Allocated fees (royalty amount): Resource royalty amount allocated to the cost allocation unit based on the allocation rules
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// Total fees (cash): Total fees of the cost allocation unit, Collected Fees (Cash) + Allocated fees (Cash)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// Total fees (voucher): Total fees of the cost allocation unit, Collected Fees (Voucher) + Allocated fees (Voucher)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// Total fees (free credit): Total fees of the cost allocation unit, Collected Fees (Free Credit) + Allocated fees (Free Credit)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// Total fees (royalty amount): Total cost of the cost allocation unit, Collected Fees (Royalty Amount) + Allocated fees (Royalty Amount)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// Collected fees (discounted total): Total resource amount after discount directly collected to the cost allocation unit based on the collection rules
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// Allocated fees (discounted total): Total resource amount after discount directly allocated to the cost allocation unit based on the allocation rules
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// Total fees (discounted total): Total fees of the cost allocation unit, Collected Fees (discounted total) + Allocated fees (discounted total)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Proportion (discounted total): Total fees (discounted total) of the Cost Allocation Unit/Total Fees (discounted total) * 100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// Month-on-month ratio (discounted total): [Total fees (discounted total) of the cost allocation unit in this month - Total fees (discounted total) of the cost allocation unit in the previous month]/Total fees (discounted total) of the cost allocation unit in the previous month * 100%
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Sequential Comparison Arrow
	// upward - Upward
	// downward - Downward
	// none - Stable
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`

	// AllocateCostBeforeTax
	AllocateCostBeforeTax *string `json:"AllocateCostBeforeTax,omitnil,omitempty" name:"AllocateCostBeforeTax"`

	// GatherCostBeforeTax
	GatherCostBeforeTax *string `json:"GatherCostBeforeTax,omitnil,omitempty" name:"GatherCostBeforeTax"`

	// TotalCostBeforeTax
	TotalCostBeforeTax *string `json:"TotalCostBeforeTax,omitnil,omitempty" name:"TotalCostBeforeTax"`

	// AllocateTax
	AllocateTax *string `json:"AllocateTax,omitnil,omitempty" name:"AllocateTax"`

	// GatherTax
	GatherTax *string `json:"GatherTax,omitnil,omitempty" name:"GatherTax"`

	// TotalTax
	TotalTax *string `json:"TotalTax,omitnil,omitempty" name:"TotalTax"`
}

type AllocationOverviewDetail struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Date: Settlement date
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// Collected fees (cash): Cash directly collected to the cost allocation unit based on the collection rules
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// Collected fees (voucher): Resource vouchers directly collected to the cost allocation unit based on the collection rules
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// Collected fees (free credit): Resource free credit directly collected to the cost allocation unit based on the collection rules
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// Collected fees (royalty amount): Resource royalty amount directly collected to the cost allocation unit based on the collection rules
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// Allocated fees (cash): Resource cash allocated to the cost allocation unit based on the allocation rules
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// Allocated fees (voucher): Resource vouchers allocated to the cost allocation unit based on the allocation rules
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// Allocated fees (free credit): Resource free credit allocated to the cost allocation unit based on the allocation rules
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// Allocated fees (royalty amount): Resource royalty amount allocated to the cost allocation unit based on the allocation rules
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// Total fees (cash): Total fees of the cost allocation unit, Collected Fees (Cash) + Allocated Fees (Cash)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// Total fees (voucher): Total fees of the cost allocation unit, Collected Fees (Voucher) + Allocated Fees (Voucher)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// Total fees (free credit): Total fees of the cost allocation unit, Collected Fees (Free Credit) + Allocated Fees (Free Credit)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// Total fees (royalty amount): Total fees of the cost allocation unit, Collected Fees (Royalty Amount) + Allocated Fees (Royalty Amount)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// Collected fees (discounted total): Total resource amount after discount directly collected to the cost allocation unit based on the collection rules
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// Allocated fees (discounted total): Total resource amount after discount directly allocated to the cost allocation unit based on the allocation rules
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// Total fees (discounted total): Total fees of the cost allocation unit, Collected Fees (discounted total) + Allocated Fees (discounted total)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Proportion (discounted total): Total fees (discounted total) of the Cost Allocation Unit/Total Fees (discounted total) * 100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// Month-on-month ratio (discounted total): [Total fees (discounted total) of the cost allocation unit in this month - Total fees (discounted total) of the cost allocation unit in the previous month]/Total fees (discounted total) of the cost allocation unit in the previous month * 100%
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Sequential Comparison Arrow
	// upward - Upward
	// downward - Downward
	// none - Stable
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`

	// GatherCostBeforeTax
	GatherCostBeforeTax *string `json:"GatherCostBeforeTax,omitnil,omitempty" name:"GatherCostBeforeTax"`

	// AllocateCostBeforeTax
	AllocateCostBeforeTax *string `json:"AllocateCostBeforeTax,omitnil,omitempty" name:"AllocateCostBeforeTax"`

	// TotalCostBeforeTax
	TotalCostBeforeTax *string `json:"TotalCostBeforeTax,omitnil,omitempty" name:"TotalCostBeforeTax"`

	// GatherTax
	GatherTax *string `json:"GatherTax,omitnil,omitempty" name:"GatherTax"`

	// AllocateTax
	AllocateTax *string `json:"AllocateTax,omitnil,omitempty" name:"AllocateTax"`

	// TotalTax
	TotalTax *string `json:"TotalTax,omitnil,omitempty" name:"TotalTax"`
}

type AllocationOverviewNode struct {
	// Cost allocation unit ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Name of a cost allocation unit
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Billing unit including a rule flag
	// 0 - No rule exists.
	// 1 - Both collection rules and allocation rules exist.
	// 2 - Only collection rules exist.
	// 3 - Only allocation rules exist.
	Symbol *uint64 `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// Detailed monthly overview of a sub-unit
	Children []*AllocationOverviewNode `json:"Children,omitnil,omitempty" name:"Children"`

	// Monthly overview amount details of a cost allocation bill
	Detail *AllocationMonthOverviewDetail `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type AllocationOverviewTotal struct {
	// Total fees: Total Fees (Cash) + Total Fees (Royalty Amount) + Total Fees (Free Credit) + Total Fees (Voucher)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash: Total fees of cash
	// Note: This field may return null, indicating that no valid values can be obtained.
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: Total fees of free credit
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher: Total fees of voucher
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty amount: Total fees of royalty amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Pre-tax price after discount
	CostBeforeTax *string `json:"CostBeforeTax,omitnil,omitempty" name:"CostBeforeTax"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`
}

type AllocationRationExpression struct {
	// Cost allocation unit ID that the sharing rule belongs to.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Sharing proportion occupied by allocation unit, pass 0 for allocation by proportion.
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type AllocationRule struct {
	// Allocation rule ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Allocation rule name
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
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

type AllocationStat struct {
	// Average cost information
	Average *AllocationAverageData `json:"Average,omitnil,omitempty" name:"Average"`
}

type AllocationSummaryByBusiness struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Date: Settlement date
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// Collected fees (cash): Cash directly collected to the cost allocation unit based on the collection rules
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// Collected fees (voucher): Resource vouchers directly collected to the cost allocation unit based on the collection rules
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// Collected fees (free credit): Resource free credit directly collected to the cost allocation unit based on the collection rules
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// Collected fees (royalty amount): Resource royalty amount directly collected to the cost allocation unit based on the collection rules
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// Allocated fees (cash): Resource cash allocated to the cost allocation unit based on the allocation rules
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// Allocated fees (voucher): Resource vouchers allocated to the cost allocation unit based on the allocation rules
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// Allocated fees (free credit): Resource free credit allocated to the cost allocation unit based on the allocation rules
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// Allocated fees (royalty amount): Resource royalty amount allocated to the cost allocation unit based on the allocation rules
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// Total fees (cash): Total fees of the cost allocation unit, Collected Fees (Cash) + Allocated fees (Cash)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// Total fees (voucher): Total fees of the cost allocation unit, Collected Fees (Vouchers) + Allocated fees (Vouchers)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// Total fees (free credit): Total fees of the cost allocation unit, Collected Fees (Free Credit) + Allocated fees (Free Credit)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// Total fees (royalty amount): Total fees of the cost allocation unit, Collected Fees (Royalty Amount) + Allocated fees (Royalty Amount)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// Collected fees (discounted total): Total resource amount after discount directly collected to the cost allocation unit based on the collection rules
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// Allocated fees (discounted total): Total resource amount after discount allocated to the cost allocation unit based on the allocation rules
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// Total fees (discounted total): Total fees of the cost allocation unit, Collected Fees (discounted total) + Allocated fees (discounted total)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Proportion (discounted total): Total fees (discounted total) of the Cost Allocation Unit/Total Fees (discounted total) * 100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// Month-on-month ratio (discounted total): [Total fees (discounted total) of the cost allocation unit in this month - Total fees (discounted total) of the cost allocation unit in the previous month]/Total fees (discounted total) of the cost allocation unit in the previous month * 100%
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// Sequential Comparison Arrow
	// upward - Upward
	// downward - Downward
	// none - Stable
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Original price of a component: Original Price = Component List Price * Component Usage * Duration of Use (not displayed if the customer enjoys a fixed price/contract price, and not displayed by default in refund scenarios), specified price mode
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Original price deducted by a reserved instance: The original price of a component deducted by reserved instances for this product or service
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// Savings plan deduction from original price: Savings Plan Deduction from Original Price = Monetary Value of Savings Plan Deduction/ Savings Plan Deduction Rate
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// Cash account expenditure (CNY): The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Promo voucher expenditure (CNY): The amount paid using various vouchers (such as promo vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Gift account expenditure (CNY): The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty account expenditure (CNY): The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Discounted total: discounted total = (Original Price - Original Price Deducted by a Reserved Instance - Savings Plan Deduction from Original Price) * Discount Rate
	AllocationRealTotalCost *string `json:"AllocationRealTotalCost,omitnil,omitempty" name:"AllocationRealTotalCost"`

	// Collected fees (tax): Tax directly collected to the cost allocation unit based on the collection rules
	GatherTax *string `json:"GatherTax,omitnil,omitempty" name:"GatherTax"`

	// Allocated fees (tax): Resource tax allocated to the cost allocation unit based on the allocation rules
	AllocateTax *string `json:"AllocateTax,omitnil,omitempty" name:"AllocateTax"`

	// Total fees (tax): Total fees of the cost allocation unit, Collected Fees (Tax) + Allocated fees (Tax)
	TotalTax *string `json:"TotalTax,omitnil,omitempty" name:"TotalTax"`

	// GatherCostBeforeTax
	GatherCostBeforeTax *string `json:"GatherCostBeforeTax,omitnil,omitempty" name:"GatherCostBeforeTax"`

	// AllocateCostBeforeTax
	AllocateCostBeforeTax *string `json:"AllocateCostBeforeTax,omitnil,omitempty" name:"AllocateCostBeforeTax"`

	// TotalCostBeforeTax
	TotalCostBeforeTax *string `json:"TotalCostBeforeTax,omitnil,omitempty" name:"TotalCostBeforeTax"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// CostBeforeTax
	CostBeforeTax *string `json:"CostBeforeTax,omitnil,omitempty" name:"CostBeforeTax"`
}

type AllocationSummaryByItem struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Date: Settlement date
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID (the resource account ID or role ID opened by prepaid resource ordering or postpaid operation)
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode: Resource billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type: Detailed transaction type
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name: The region where the resource is located
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// AZ ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone: The availability zone where the resource is located.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance type code
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. It is displayed as "-" by default for regular instances.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Resource ID: Resources vary by product, and the content is not identical. For example, Cloud Virtual Machine (CVM) corresponds to the instance ID. If the product is split, it shows the split item ID, such as COS bucket ID and CDN domain name.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The name set by the user for the resource in the console, which is empty by default if not set. If the product is split, it shows the split resource alias.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Allocation tag: The resource-bound tag
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The Project to which a resource belongs, which is independently allocated by the user for the resource in the console. If a resource has not been allocated to an Project, it will be a default Project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Cost collection type: The source types of fees, including allocated, collection and unallocated.
	// 0 - Allocation
	// 1 - Collection
	// -1 - Unallocated
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// Original price of a component: Original Price = Component List Price * Component Usage * Duration of Use (not displayed if the customer enjoys a fixed price/contract price, and not displayed by default in refund scenarios), specified price mode
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Reserved instance deduction duration: The duration of use deducted by reserved instances for this product or service.
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// Original price deducted by a reserved instance: The original price of a component deducted by reserved instances for this product or service
	RiCost *string `json:"RiCost,omitnil,omitempty" name:"RiCost"`

	// Discounted total: discounted total = (Original Price - Original Price Deducted by a Reserved Instance - Savings Plan Deduction from Original Price) * Discount Rate
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash account expenditure (CNY): The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Promo voucher expenditure (CNY): The amount paid using various vouchers (such as promo vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Gift account expenditure (CNY): The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty account expenditure (CNY): The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Component name code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name: The specific component of a product or service purchased by the user
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component type: The major component category corresponding to the product or service purchased by the user
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// Split item ID: The ID of the split item involved in the split product, such as COS bucket ID and CDN domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// Split item name: The split item involved in the split product
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// Usage start time: Usage start time
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: Product or service usage end time
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Savings plan deduction from original price: Savings Plan Deduction from Original Price = Monetary Value of Savings Plan Deduction/ Savings Plan Deduction Rate
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// Domestic and international codes
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Domestic and international: Resource region type (domestic, international)
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Component list price: The original unit price of the component on the portal (not displayed if the customer enjoys a fixed price/contract price)
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// Component unit price: Discounted unit price of the component. Component unit price = list price * discount.
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// Component Price Unit: Unit of component price, Unit Composition: CNY/usage unit/duration unit
	SinglePriceUnit *string `json:"SinglePriceUnit,omitnil,omitempty" name:"SinglePriceUnit"`

	// Component usage: The actual settlement usage of the component, Component Usage = Original Component Usage - Deducted Usage (including resource packages)
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// Component usage unit: Unit of measurement corresponding to component usage.
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// Usage duration: The duration of resource usage, Component Usage = Original Component Usage Duration - Deducted Duration (including resource packages)
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Duration unit: Unit of resource usage duration.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// Remark attribute (instance configuration): Additional remark information, such as reserved instance type and transaction type for reserved instances, regional information of both ends for CCN products.
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// Original usage/duration: The original usage of the component before deduction by resource packages.
	// (Currently only TRTC, TEM, Cloud Call Center, and CDZ products support this information display. Other products are being integrated.)
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// Deduction of usage/duration (including resource packages): The amount of usage deducted by resource packages
	// (Currently only TRTC, TEM, Cloud Call Center, and CDZ products support this information display. Other products are being integrated.)
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// Discount rate: The discount rate enjoyed by this resource (it is not shown by default if the customer enjoys a fixed/contract price, and it is also not shown by default in the refund scenario)
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// Mixed discount rate: The final discount rate after integrating various discount deductions. Mixed Discount Rate = Discounted total price/Original Price.
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// Price attribute: Other attributes of the component that affect discount pricing besides unit price and duration
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Calculation rule explanation: A detailed explanation to calculations of billing settlement for special transaction types, such as refund and configuration changes.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing Rules: The detailed billing rules for each product shown in the portal explanation link
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Configuration description: Information on specification of resource configuration
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// SPDeduction
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// Savings plan deduction rate: The discount rate for this component within the available balance limit of the savings plan
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// Currency.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// CostBeforeTax
	CostBeforeTax *string `json:"CostBeforeTax,omitnil,omitempty" name:"CostBeforeTax"`

	// AmountBeforeTax
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`

	// AssociatedOrder
	AssociatedOrder *string `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Discount object of the current consumption item, such as official website discount, user discount and activity discount.
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// Discount type of the current consumption item, such as discount and contract price.
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// Supplementary description of the offer type, for example: business discount 20% off, the offer type is "discount" and the discount content is "0.8".
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type AllocationSummaryByResource struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Date: Settlement date
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID (the resource account ID or role ID opened by prepaid resource ordering or postpaid operation)
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode: Resource billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type: Detailed transaction type
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name: The region where the resource is located
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// AZ ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone: The availability zone where the resource is located.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Instance type code
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. It is displayed as "-" by default for regular instances.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Resource ID: Resources vary by product, and the content is not identical. For example, Cloud Virtual Machine (CVM) corresponds to the instance ID. If the product is split, it shows the split item ID, such as COS bucket ID and CDN domain name.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The name set by the user for the resource in the console, which is empty by default if not set. If the product is split, it shows the split resource alias.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Allocation tag: The resource-bound tag
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The Project to which a resource belongs, which is independently allocated by the user for the resource in the console. If a resource has not been allocated to an Project, it will be a default Project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Cost collection type: The source types of fees, including allocated, collection and unallocated.
	// 0 - Allocation 
	// 1 - Collection 
	// -1 - Unallocated
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// Original price of a component: Original Price = Component List Price * Component Usage * Duration of Use (not displayed if the customer enjoys a fixed price/contract price, and not displayed by default in refund scenarios), specified price mode
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Reserved instance deduction duration: The duration of use deducted by reserved instances for this product or service.
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// Original price deducted by a reserved instance: The original price of a component deducted by reserved instances for this product or service
	RiCost *string `json:"RiCost,omitnil,omitempty" name:"RiCost"`

	// Discounted total: discounted total = (Original Price - Original Price Deducted by a Reserved Instance - Savings Plan Deduction from Original Price) * Discount Rate
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash account expenditure (CNY): The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Promo voucher expenditure (CNY): The amount paid using various vouchers (such as promo vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Gift account expenditure (CNY): The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty account expenditure (CNY): The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Split item ID: The ID of the split item involved in the split product, such as COS bucket ID and CDN domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// Split item name: The split item involved in the split product
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// Usage start time: Usage start time
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// Usage end time: Product or service usage end time
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Savings plan deduction from original price: Savings Plan Deduction from Original Price = Monetary Value of Savings Plan Deduction/ Savings Plan Deduction Rate
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// Domestic and international codes
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Domestic and international: Resource region type (domestic, international)
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Configuration description: Name and usage of each component under the corresponding resource (the total usage if the component is cumulative usage billing type)
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// SPDeduction
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// Currency.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// CostBeforeTax
	CostBeforeTax *string `json:"CostBeforeTax,omitnil,omitempty" name:"CostBeforeTax"`

	// AmountBeforeTax
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
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

type AllocationTreeNode struct {
	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`
}

type AllocationUnit struct {
	// Cost allocation unit ID.
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Specifies the name of a cost allocation rule.
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`
}

type AnalyseActionTypeDetail struct {
	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type Name
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`
}

type AnalyseAmountDetail struct {
	// Fee type
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Indicates whether to display
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`
}

type AnalyseBusinessDetail struct {
	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type AnalyseConditionDetail struct {
	// product
	Business []*AnalyseBusinessDetail `json:"Business,omitnil,omitempty" name:"Business"`

	// Project
	Project []*AnalyseProjectDetail `json:"Project,omitnil,omitempty" name:"Project"`

	// Region.
	Region []*AnalyseRegionDetail `json:"Region,omitnil,omitempty" name:"Region"`

	// Billing mode.
	PayMode []*AnalysePayModeDetail `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type
	ActionType []*AnalyseActionTypeDetail `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Availability zone
	Zone []*AnalyseZoneDetail `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Resource owner Uin
	OwnerUin []*AnalyseOwnerUinDetail `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Fee type
	Amount []*AnalyseAmountDetail `json:"Amount,omitnil,omitempty" name:"Amount"`
}

type AnalyseConditions struct {
	// Product name code
	BusinessCodes *string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Subproduct name code
	ProductCodes *string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Availability zone ID: The availability zone ID where the resource is located.
	ZoneIds *string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Region ID: Resource region ID
	RegionIds *string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Project ID: Project ID of the resource
	ProjectIds *string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Billing mode prePay (monthly subscription)/postPay (pay-as-you-go billing)
	PayModes *string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type. Query transaction type (please use transaction type code as input parameter).
	ActionTypes *string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Cost allocation tag key
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Fee type. Query fee type (please use fee type code input parameter). The input parameter enumeration is as follows:
	// cashPayAmount: Cash 
	// incentivePayAmount: Bonus 
	// voucherPayAmount: Coupon 
	// tax: tax. 
	// costBeforeTax: pre-tax price
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// User UIN for querying cost analysis data
	PayerUins *string `json:"PayerUins,omitnil,omitempty" name:"PayerUins"`

	// User UIN for using resources
	OwnerUins *string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Consumption type. Query consumption type (please use consumption type code input parameter).
	ConsumptionTypes *string `json:"ConsumptionTypes,omitnil,omitempty" name:"ConsumptionTypes"`
}

type AnalyseDetail struct {
	// Time
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Amount
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Date Detail Amount
	TimeDetail []*AnalyseTimeDetail `json:"TimeDetail,omitnil,omitempty" name:"TimeDetail"`
}

type AnalyseHeaderDetail struct {
	// Header date
	HeadDetail []*AnalyseHeaderTimeDetail `json:"HeadDetail,omitnil,omitempty" name:"HeadDetail"`

	// Time
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// total
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`
}

type AnalyseHeaderTimeDetail struct {
	// Date
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AnalyseOwnerUinDetail struct {
	// User uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type AnalysePayModeDetail struct {
	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode Name
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type AnalyseProjectDetail struct {
	// Project ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// default project
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type AnalyseRegionDetail struct {
	// Region ID.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type AnalyseTimeDetail struct {
	// Date
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Amount
	Money *string `json:"Money,omitnil,omitempty" name:"Money"`
}

type AnalyseZoneDetail struct {
	// AZ ID.
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Available zone Name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ApplicableProducts struct {
	// Valid values: `all products` or names of the applicable products (string). Multiple names are separated by commas.
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// Valid values: `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all. If `GoodsName` contains multiple product names and `PayMode` is `*`, it indicates that the voucher can be used in all billing modes for each of the products.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type BillActionType struct {
	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type: Detailed transaction type
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`
}

type BillBusiness struct {
	// Product code
	// 
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type BillBusinessLink struct {
	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Subproduct
	Children []*BillProductLink `json:"Children,omitnil,omitempty" name:"Children"`
}

type BillComponent struct {
	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component type: The major component category corresponding to the product or service purchased by the user
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`
}

type BillDays struct {
	// Date: Settlement date
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`
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

	// Order ID: The sub-order number corresponding to the monthly subscription mode. In the postpaid billing model, the bill amount does not exist as an order concept, and this parameter can be ignored.
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

	// Tag information.
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Price attribute: Other attributes of the component that affect discount pricing besides unit price and duration
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Associated transaction document ID: Document ID associated with this transaction, such as a write-off order, the original order, a resettlement order, or the original purchase order number recorded in a refund order.
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Calculation explanation: A detailed explanation to calculations of billing settlement for special transaction types, such as refund and configuration changes.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing Rules: The detailed billing rules for each product shown in the portal explanation link
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Billing day
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Billing record ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Domestic and international codes
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Domestic and international: Resource region type (domestic, international)
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Remark attribute (instance configuration): Additional remark information, such as reserved instance type and transaction type for reserved instances, regional information of both ends for CCN products.
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// discount object
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// Offer type
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// discount content
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`
}

type BillDetailAssociatedOrder struct {
	// Newly purchased order
	PrepayPurchase *string `json:"PrepayPurchase,omitnil,omitempty" name:"PrepayPurchase"`

	// Renewal order
	PrepayRenew *string `json:"PrepayRenew,omitnil,omitempty" name:"PrepayRenew"`

	// Configuration upgrade order
	PrepayModifyUp *string `json:"PrepayModifyUp,omitnil,omitempty" name:"PrepayModifyUp"`

	// Write-off order
	ReverseOrder *string `json:"ReverseOrder,omitnil,omitempty" name:"ReverseOrder"`

	// Order after discount adjustment
	NewOrder *string `json:"NewOrder,omitnil,omitempty" name:"NewOrder"`

	// Order before discount adjustment
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

	// Original usage/duration: The original usage of the component before deduction by resource packages.
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// Deduction of usage/duration (including resource packages): The amount of usage/duration deducted by resource packages
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

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Component type code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component unit price: Discounted unit price of the component. Component unit price = list price * discount.
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. Normal instance display is not displayed by default.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// RI deduction duration: The duration of use deducted by reserved instances for this product or service.
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// Reserved Instance Deduction Component Original Price: The original price of a component deducted by reserved instances for this product or service
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// Savings plan deduction rate: The discount rate for this component within the available balance limit of the savings plan
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// Cost deduction by SP. This parameter has been deprecated. Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// Original Price of Savings Plan Deduction Component: Savings Plan Deduction from Original Price = Deduction Amount of Savings Plan Package / Savings Plan Deduction Rate
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// Mixed discount rate: The final discount rate after integrating various discount deductions. Mixed Discount Rate = Discounted total price/Component original price.
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// Configuration description: Information on specification of resource configuration
	ComponentConfig []*BillDetailComponentConfig `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Tax.
	TaxAmount *string `json:"TaxAmount,omitnil,omitempty" name:"TaxAmount"`

	// Currency.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`
}

type BillDetailComponentConfig struct {
	// Configuration description name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Configuration description value
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

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Extended field 3: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// Extended field 4: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// Extended field 5: The extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// Tag information.
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

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillInstanceType struct {
	// Instance type code
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. It is displayed as "-" by default for regular instances.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`
}

type BillItem struct {
	// Component name code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name: The specific component of a product or service purchased by the user
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`
}

type BillOperateUin struct {
	// Operator account ID (the resource account ID or role ID opened by prepaid resource ordering or postpaid operation)
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

type BillOwnerUin struct {
	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type BillPayMode struct {
	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode: Resource billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type BillProduct struct {
	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`
}

type BillProductLink struct {
	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Component name
	Children []*BillItem `json:"Children,omitnil,omitempty" name:"Children"`
}

type BillProject struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The Project to which a resource belongs, which is independently allocated by the user for the resource in the console. If a resource has not been allocated to an Project, it will be a default Project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type BillRegion struct {
	// Region ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name: The region where the resource is located
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
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

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Extended field 3: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// Extended field 4: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// Extended field 5: Extended attribute information of a product, which is displayed on the resource bill only.
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// Tag information.
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

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type BillTagInfo struct {
	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type BillZoneId struct {
	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability zone: The availability zone where the resource is located.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type BusinessSummaryInfo struct {
	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: The name of a Tencent Cloud product purchased by the user, such as CVM.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Original price in CNY. The TotalCost field comes into effect after bill 3.0 (May 2021) and returns "-" before bill 3.0. In the current situation of contract price, the TotalCost field returns "-" if a price difference exists with the official website price.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Total amount after discount
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash credit: The amount paid from the user’s cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Free credit: The amount paid with the user’s free credit
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher payment: The voucher deduction amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type BusinessSummaryOverviewItem struct {
	// Product code
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
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Bonus
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty amount
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Cash payment (pre-tax)
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Region name (only shown in regional summary)
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
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Bonus
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Voucher
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty amount
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Cash payment (pre-tax)
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`
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
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Voucher
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Bonus
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty amount
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// Cash payment (pre-tax)
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`
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

	// Original price
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// Fee start time
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// End time of fees
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// Days
	DayDiff *string `json:"DayDiff,omitnil,omitempty" name:"DayDiff"`

	// Daily consumption
	DailyTotalCost *string `json:"DailyTotalCost,omitnil,omitempty" name:"DailyTotalCost"`

	// Order ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Voucher
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Bonus
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty amount
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`

	// Cash payment (pre-tax)
	AmountBeforeTax *string `json:"AmountBeforeTax,omitnil,omitempty" name:"AmountBeforeTax"`

	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator UIN: Operator account ID (ID of the operator who places orders for prepaid resources or activates postpaid resource account, or role ID).
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user, such as Cloud Virtual Machine (CVM)-Standard Type S1
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Region type
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// Region type name.
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// Extension Field 1
	Extend1 *string `json:"Extend1,omitnil,omitempty" name:"Extend1"`

	// Extension Field 2
	Extend2 *string `json:"Extend2,omitnil,omitempty" name:"Extend2"`

	// Extension Field 3
	Extend3 *string `json:"Extend3,omitnil,omitempty" name:"Extend3"`

	// Extension Field 4
	Extend4 *string `json:"Extend4,omitnil,omitempty" name:"Extend4"`

	// Extension Field 5
	Extend5 *string `json:"Extend5,omitnil,omitempty" name:"Extend5"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance Type Name
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Deduction time: Deduction time
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// Availability zone: The availability zone where the resource is located, such as Guangzhou Zone 3.
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// Describing Configurations
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// Tag information.
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ConsumptionSummaryTotal struct {
	// Discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type ConsumptionSummaryTrend struct {
	// Trend type, upward for rising, downward for falling, none for no change
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Trend value. The value of this field is null when Type is none.
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

	// Tax
	Tax *string `json:"Tax,omitnil,omitempty" name:"Tax"`

	// tax rate
	TaxRate *string `json:"TaxRate,omitnil,omitempty" name:"TaxRate"`
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

	// Type name
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

	// Tag information.	
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type CreateAllocationRuleRequestParams struct {
	// List of sharing rules
	RuleList *AllocationRulesSummary `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// Month, the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type CreateAllocationRuleRequest struct {
	*tchttp.BaseRequest
	
	// List of sharing rules
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

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// ClientToken is an unique, case-sensitive string generated by the client with no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544****.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Availability zone code.
	ZoneCode *string `json:"ZoneCode,omitnil,omitempty" name:"ZoneCode"`

	// Payment mode. Available values: PrePay: upfront charge.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product detailed information.
	Parameter *string `json:"Parameter,omitnil,omitempty" name:"Parameter"`

	// Product quantity, default value is 1.
	Quantity *int64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Project id, default value is 0.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Purchase duration, max number is 36, default value is 1.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Purchase duration unit. valid values: 
	// m: month,
	// y: year. 
	// default value is: m.
	PeriodUnit *string `json:"PeriodUnit,omitnil,omitempty" name:"PeriodUnit"`

	// Auto-renewal flag. valid values: NOTIFY_AND_MANUAL_RENEW: manually renew, NOTIFY_AND_AUTO_RENEW: automatically renew, DISABLE_NOTIFY_AND_MANUAL_RENEW: renewal is disabled. 
	// default value is NOTIFY_AND_MANUAL_RENEW.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ClientToken is an unique, case-sensitive string generated by the client with no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544****.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Availability zone code.
	ZoneCode *string `json:"ZoneCode,omitnil,omitempty" name:"ZoneCode"`

	// Payment mode. Available values: PrePay: upfront charge.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Product detailed information.
	Parameter *string `json:"Parameter,omitnil,omitempty" name:"Parameter"`

	// Product quantity, default value is 1.
	Quantity *int64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// Project id, default value is 0.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Purchase duration, max number is 36, default value is 1.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Purchase duration unit. valid values: 
	// m: month,
	// y: year. 
	// default value is: m.
	PeriodUnit *string `json:"PeriodUnit,omitnil,omitempty" name:"PeriodUnit"`

	// Auto-renewal flag. valid values: NOTIFY_AND_MANUAL_RENEW: manually renew, NOTIFY_AND_AUTO_RENEW: automatically renew, DISABLE_NOTIFY_AND_MANUAL_RENEW: renewal is disabled. 
	// default value is NOTIFY_AND_MANUAL_RENEW.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientToken")
	delete(f, "ProductCode")
	delete(f, "SubProductCode")
	delete(f, "RegionCode")
	delete(f, "ZoneCode")
	delete(f, "PayMode")
	delete(f, "Parameter")
	delete(f, "Quantity")
	delete(f, "ProjectId")
	delete(f, "Period")
	delete(f, "PeriodUnit")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// Order ID
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// Instance id list. Empty list will be returned once product delivery is delayed. 
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
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

	// Availability zone Id corresponding to the order
	ZoneCode *string `json:"ZoneCode,omitnil,omitempty" name:"ZoneCode"`
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
	// Query the temporary limit
	TempCredit *bool `json:"TempCredit,omitnil,omitempty" name:"TempCredit"`
}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
	// Query the temporary limit
	TempCredit *bool `json:"TempCredit,omitnil,omitempty" name:"TempCredit"`
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
	delete(f, "TempCredit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceResponseParams struct {
	// Available account balance in cents, which takes the same calculation rules as `RealBalance`, `CreditBalance`, and `RealCreditBalance`.
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// The user Uin for the query
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
	//
	// Deprecated: IsAllowArrears is deprecated.
	IsAllowArrears *bool `json:"IsAllowArrears,omitnil,omitempty" name:"IsAllowArrears"`

	// Whether you have a credit limit. Currently, this field is not applied.
	//
	// Deprecated: IsCreditLimited is deprecated.
	IsCreditLimited *bool `json:"IsCreditLimited,omitnil,omitempty" name:"IsCreditLimited"`

	// Credit limit in cents. Credit limit－available credit balance = consumption amount
	CreditAmount *float64 `json:"CreditAmount,omitnil,omitempty" name:"CreditAmount"`

	// Available credit balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `RealCreditBalance`.
	CreditBalance *float64 `json:"CreditBalance,omitnil,omitempty" name:"CreditBalance"`

	// Available account balance in cents, which takes the same calculation rules as `Balance`, `RealBalance`, and `CreditBalance`.
	RealCreditBalance *float64 `json:"RealCreditBalance,omitnil,omitempty" name:"RealCreditBalance"`

	// Temporary limit, unit cent
	TempCredit *float64 `json:"TempCredit,omitnil,omitempty" name:"TempCredit"`

	// Temporary limit details
	TempAmountInfoList []*UinTempAmountModel `json:"TempAmountInfoList,omitnil,omitempty" name:"TempAmountInfoList"`

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
type DescribeAllocateConditionsRequestParams struct {
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocateConditionsRequest struct {
	*tchttp.BaseRequest
	
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocateConditionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocateConditionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocateConditionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocateConditionsResponseParams struct {
	// Product filter list
	Business []*BillBusinessLink `json:"Business,omitnil,omitempty" name:"Business"`

	// Subproduct filter list
	Product []*BillProduct `json:"Product,omitnil,omitempty" name:"Product"`

	// Component name filter list
	Item []*BillItem `json:"Item,omitnil,omitempty" name:"Item"`

	// Region filter list
	Region []*BillRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance type filter list
	InstanceType []*BillInstanceType `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode filter list
	PayMode []*BillPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Project filter list
	Project []*BillProject `json:"Project,omitnil,omitempty" name:"Project"`

	// Tag filter list
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// User UIN filter list
	OwnerUin []*BillOwnerUin `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator UIN filter list
	OperateUin []*BillOperateUin `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Transaction type filter list
	ActionType []*BillActionType `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocateConditionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocateConditionsResponseParams `json:"Response"`
}

func (r *DescribeAllocateConditionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocateConditionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillConditionsRequestParams struct {
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Date
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// AZ ID
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search criteria
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationBillConditionsRequest struct {
	*tchttp.BaseRequest
	
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Date
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// AZ ID
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search criteria
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationBillConditionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillConditionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationBillConditionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillConditionsResponseParams struct {
	// Product filter list
	Business []*BillBusiness `json:"Business,omitnil,omitempty" name:"Business"`

	// Subproduct filter list
	Product []*BillProduct `json:"Product,omitnil,omitempty" name:"Product"`

	// Component name filter list
	Item []*BillItem `json:"Item,omitnil,omitempty" name:"Item"`

	// Region filter list
	Region []*BillRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance type filter list
	InstanceType []*BillInstanceType `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Billing mode filter list
	PayMode []*BillPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Project filter list
	Project []*BillProject `json:"Project,omitnil,omitempty" name:"Project"`

	// Tag filter list
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// User UIN filter list
	OwnerUin []*BillOwnerUin `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator UIN filter list
	OperateUin []*BillOperateUin `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Date filter list
	BillDay []*BillDays `json:"BillDay,omitnil,omitempty" name:"BillDay"`

	// Transaction type filter list
	ActionType []*BillActionType `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Component type filter list
	Component []*BillComponent `json:"Component,omitnil,omitempty" name:"Component"`

	// Availability zone filter list
	Zone []*BillZoneId `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Cost allocation unit filter list
	AllocationTreeNode []*AllocationTreeNode `json:"AllocationTreeNode,omitnil,omitempty" name:"AllocationTreeNode"`

	// Cost allocation tag key
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationBillConditionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationBillConditionsResponseParams `json:"Response"`
}

func (r *DescribeAllocationBillConditionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillConditionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillDetailRequestParams struct {
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// AZ ID, used for filtering
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code, used for filtering
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

type DescribeAllocationBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// AZ ID, used for filtering
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code, used for filtering
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

func (r *DescribeAllocationBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillDetailResponseParams struct {
	// Total quantity.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Details of a cost allocation bill
	Detail []*AllocationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationBillDetailResponseParams `json:"Response"`
}

func (r *DescribeAllocationBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationMonthOverviewRequestParams struct {
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocationMonthOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocationMonthOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationMonthOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationMonthOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationMonthOverviewResponseParams struct {
	// Monthly overview of a cost allocation bill
	Detail []*AllocationOverviewNode `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationMonthOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationMonthOverviewResponseParams `json:"Response"`
}

func (r *DescribeAllocationMonthOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationMonthOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationOverviewRequestParams struct {
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If offset is 0, it indicates the first page. If limit is 100, then offset is 100, it indicates the second page; if offset is 200, it indicates the third page, and so on
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows: 
	// GatherCashPayAmount - Collected fees (cash)
	// GatherVoucherPayAmount - Collected fees (voucher)
	// GatherIncentivePayAmount - Collected fees (free credit)
	// GatherTransferPayAmount - Collected fees (royalty amount)
	// AllocateCashPayAmount - Allocated fees (cash)
	// AllocateVoucherPayAmount - Allocated fees (voucher)
	// AllocateIncentivePayAmount - Allocated fees (free credit)
	// AllocateTransferPayAmount - Allocated fees (royalty amount)
	// TotalCashPayAmount - Total fees (cash)
	// TotalVoucherPayAmount - Total fees (voucher)
	// TotalIncentivePayAmount - Total fees (free credit)
	// TotalTransferPayAmount - Total fees (royalty amount)
	// GatherRealCost - Collected fees (discounted total)
	// AllocateRealCost - Allocated fees (discounted total)
	// RealTotalCost - Total fees (discounted total)
	// Ratio - Proportion (discounted total)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`
}

type DescribeAllocationOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If offset is 0, it indicates the first page. If limit is 100, then offset is 100, it indicates the second page; if offset is 200, it indicates the third page, and so on
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows: 
	// GatherCashPayAmount - Collected fees (cash)
	// GatherVoucherPayAmount - Collected fees (voucher)
	// GatherIncentivePayAmount - Collected fees (free credit)
	// GatherTransferPayAmount - Collected fees (royalty amount)
	// AllocateCashPayAmount - Allocated fees (cash)
	// AllocateVoucherPayAmount - Allocated fees (voucher)
	// AllocateIncentivePayAmount - Allocated fees (free credit)
	// AllocateTransferPayAmount - Allocated fees (royalty amount)
	// TotalCashPayAmount - Total fees (cash)
	// TotalVoucherPayAmount - Total fees (voucher)
	// TotalIncentivePayAmount - Total fees (free credit)
	// TotalTransferPayAmount - Total fees (royalty amount)
	// GatherRealCost - Collected fees (discounted total)
	// AllocateRealCost - Allocated fees (discounted total)
	// RealTotalCost - Total fees (discounted total)
	// Ratio - Proportion (discounted total)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`
}

func (r *DescribeAllocationOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationOverviewResponseParams struct {
	// Total quantity.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Details of the cost allocation overview
	Detail []*AllocationOverviewDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationOverviewResponseParams `json:"Response"`
}

func (r *DescribeAllocationOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationOverviewResponse) FromJsonString(s string) error {
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
type DescribeAllocationSummaryByBusinessRequestParams struct {
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	// 
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Sorting field, with the enumerated values as follows:
	// GatherCashPayAmount - Collected fees (cash)
	// GatherVoucherPayAmount - Collected fees (voucher)
	// GatherIncentivePayAmount - Collected fees (free credit)
	// GatherTransferPayAmount - Collected fees (royalty amount)
	// AllocateCashPayAmount - Allocated fees (cash)
	// AllocateVoucherPayAmount - Allocated fees (voucher)
	// AllocateIncentivePayAmount - Allocated fees (free credit)
	// AllocateTransferPayAmount - Allocated fees (royalty amount)
	// TotalCashPayAmount - Total fees (cash)
	// TotalVoucherPayAmount - Total fees (voucher)
	// TotalIncentivePayAmount - Total fees (free credit)
	// TotalTransferPayAmount - Total fees (royalty amount)
	// GatherRealCost - Collected fees (discounted total)
	// AllocateRealCost - Allocated fees (discounted total)
	// RealTotalCost - Total fees (discounted total)
	// BusinessCode - Product code
	// Ratio - Proportion (discounted total)
	// Trend - Month-on-month ratio (discounted total)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Date, used for filtering and provided when PeriodType is set to day
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Fuzzy search criteria
	//
	// Deprecated: SearchKey is deprecated.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeAllocationSummaryByBusinessRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	// 
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Sorting field, with the enumerated values as follows:
	// GatherCashPayAmount - Collected fees (cash)
	// GatherVoucherPayAmount - Collected fees (voucher)
	// GatherIncentivePayAmount - Collected fees (free credit)
	// GatherTransferPayAmount - Collected fees (royalty amount)
	// AllocateCashPayAmount - Allocated fees (cash)
	// AllocateVoucherPayAmount - Allocated fees (voucher)
	// AllocateIncentivePayAmount - Allocated fees (free credit)
	// AllocateTransferPayAmount - Allocated fees (royalty amount)
	// TotalCashPayAmount - Total fees (cash)
	// TotalVoucherPayAmount - Total fees (voucher)
	// TotalIncentivePayAmount - Total fees (free credit)
	// TotalTransferPayAmount - Total fees (royalty amount)
	// GatherRealCost - Collected fees (discounted total)
	// AllocateRealCost - Allocated fees (discounted total)
	// RealTotalCost - Total fees (discounted total)
	// BusinessCode - Product code
	// Ratio - Proportion (discounted total)
	// Trend - Month-on-month ratio (discounted total)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Date, used for filtering and provided when PeriodType is set to day
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Fuzzy search criteria
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeAllocationSummaryByBusinessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByBusinessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "SortType")
	delete(f, "Sort")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByBusinessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByBusinessResponseParams struct {
	// Total quantity.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Detailed summary of the cost allocation bill by business
	Detail []*AllocationSummaryByBusiness `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByBusinessResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByBusinessResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByBusinessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByBusinessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByItemRequestParams struct {
	// Quantity, with the maximum value of 1,000
	// 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	// 
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	// 
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	// 
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	// 
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	// 
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	// 
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	// 
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	// 
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	// 
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Availability Zone (AZ) ID, used for filtering
	// 
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	// 
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	// 
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code, used for filtering
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	// 
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type, with the enumerated values as follows:
	// 0 - Allocation
	// 1 - Collection
	// -1 - Unallocated
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationSummaryByItemRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	// 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	// 
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	// 
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	// 
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	// 
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	// 
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	// 
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	// 
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	// 
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	// 
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Availability Zone (AZ) ID, used for filtering
	// 
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	// 
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	// 
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Component type code, used for filtering
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	// 
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type, with the enumerated values as follows:
	// 0 - Allocation
	// 1 - Collection
	// -1 - Unallocated
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationSummaryByItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByItemResponseParams struct {
	// Total quantity.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Details of a Cost Allocation Bill by item
	Detail []*AllocationSummaryByItem `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByItemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByItemResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByResourceRequestParams struct {
	// Quantity, with the maximum value of 1,000
	// 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Availability zone (AZ) ID, used for filtering
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Fuzzy search: supports tag, resource ID, and resource alias
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type, with the enumerated values as follows:
	// 0 - Allocation 
	// 1 - Collection 
	// -1 - Unallocated
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	// 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	// 
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Statistical period, with the enumerated values as follows:
	// month - Month
	// day - Day
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// Sorting field, with the enumerated values as follows:
	// RiTimeSpan - Deduction duration of a reserved instance
	// ExtendPayAmount1 - Original price for the deduction duration of a reserved instance
	// RealCost - Discounted total
	// CashPayAmount - Cash amount
	// VoucherPayAmount - Amount of promo voucher
	// IncentivePayAmount - Amount of free credit
	// TransferPayAmount - Royalty amount
	// Cost - Original price of a component
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Date, used for filtering
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Region ID, used for filtering
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Availability zone (AZ) ID, used for filtering
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Fuzzy search: supports tag, resource ID, and resource alias
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Project ID, used for filtering
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Cost collection type, with the enumerated values as follows:
	// 0 - Allocation 
	// 1 - Collection 
	// -1 - Unallocated
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationSummaryByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByResourceResponseParams struct {
	// Total quantity.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Total amount of a cost allocation bill
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Detailed summary of the cost allocation bill by resource
	Detail []*AllocationSummaryByResource `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByResourceResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByResourceResponse) FromJsonString(s string) error {
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
type DescribeAllocationTrendByMonthRequestParams struct {
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Product code, used for filtering
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`
}

type DescribeAllocationTrendByMonthRequest struct {
	*tchttp.BaseRequest
	
	// Bill month, in the format of 2024-02, which is the current month by default if not provided
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Product code, used for filtering
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`
}

func (r *DescribeAllocationTrendByMonthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTrendByMonthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TreeNodeUniqKey")
	delete(f, "BusinessCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationTrendByMonthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationTrendByMonthResponseParams struct {
	// Current month's expense information
	Current *AllocationBillTrendDetail `json:"Current,omitnil,omitempty" name:"Current"`

	// Previous months' expense information
	Previous []*AllocationBillTrendDetail `json:"Previous,omitnil,omitempty" name:"Previous"`

	// Expense statistical information
	Stat *AllocationStat `json:"Stat,omitnil,omitempty" name:"Stat"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationTrendByMonthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationTrendByMonthResponseParams `json:"Response"`
}

func (r *DescribeAllocationTrendByMonthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTrendByMonthResponse) FromJsonString(s string) error {
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
	// <p>Format: yyyy-MM<br>Bill month. Either month or timeFrom&amp;timeTo must be specified. If timeFrom&amp;timeTo is specified, the month field is invalid.</p>
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// <p>Format: yyyy-MM-dd<br>Start time. Either month or timeFrom&amp;timeTo must be specified. If timeFrom&amp;timeTo is specified, the month field is invalid. timeFrom and timeTo must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month.</p>
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// <p>Format: yyyy-MM-dd<br>End time. Either month or timeFrom&amp;timeTo must be specified. If this field is specified, the month field is invalid. timeFrom and timeTo must be passed together and be in the same month. Cross-month queries are not supported. The query result is data of the entire month.</p>
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// <p>Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.</p>
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillAdjustInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>Format: yyyy-MM<br>Bill month. Either month or timeFrom&amp;timeTo must be specified. If timeFrom&amp;timeTo is specified, the month field is invalid.</p>
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// <p>Format: yyyy-MM-dd<br>Start time. Either month or timeFrom&amp;timeTo must be specified. If timeFrom&amp;timeTo is specified, the month field is invalid. timeFrom and timeTo must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month.</p>
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// <p>Format: yyyy-MM-dd<br>End time. Either month or timeFrom&amp;timeTo must be specified. If this field is specified, the month field is invalid. timeFrom and timeTo must be passed together and be in the same month. Cross-month queries are not supported. The query result is data of the entire month.</p>
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// <p>Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.</p>
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillAdjustInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillAdjustInfoResponseParams struct {
	// <p>Total data</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>Detailed data</p>
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

	// Billing mode, which can be `prePay` (yearly/monthly subscription) or `postPay` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Yearly/monthly subscription purchase
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription upgrade/downgrade
	// Yearly/monthly subscription refund 
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
	// Yearly/monthly subscription resource migration in 
	// Yearly/monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Yearly/monthly subscription to pay-as-you-go 
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

	// Billing mode, which can be `prePay` (yearly/monthly subscription) or `postPay` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// ID of the instance to be queried.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Yearly/monthly subscription purchase
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription upgrade/downgrade
	// Yearly/monthly subscription refund 
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
	// Yearly/monthly subscription resource migration in 
	// Yearly/monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Yearly/monthly subscription to pay-as-you-go 
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

	// Context information of this request can be used in the request parameter of the next request to accelerate query speed.
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
	// Pagination offset. Offset=0 indicates the first page. If Limit=100, Offset=100 indicates the second page, Offset=200 indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `300`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Period type, byUsedTime by billing period/byPayTime by fee deduction cycle. It should be consistent with the billing cycle for the month in the expense center. You can go to the top of the [bill overview](https://console.cloud.tencent.com/expense/bill/overview) page to view and confirm your billing cycle type.
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Total number of records for access list needed for frontend pagination
	// 1: needed, 0: not needed
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	//
	// Deprecated: ProductCode is deprecated.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Hourly settlement
	// Daily settlement
	// Yearly/monthly subscription
	// Spot
	// New yearly/monthly subscription
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription specification adjustment
	// Yearly/monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: Project ID of the resource
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product name code
	// Remark: If needed to obtain BusinessCode used in current month, invoke API: <a href="https://www.tencentcloud.com/document/product/555/35761?from_cn_redirect=1">Get fee distribution by product</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned from the last request. Paginated query of data for months with Month>=2023-05 can speed up query speed. Recommended for users with data volume at tens of thousands level. Query speed can be improved by 2-10x.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset. Offset=0 indicates the first page. If Limit=100, Offset=100 indicates the second page, Offset=200 indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `300`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Period type, byUsedTime by billing period/byPayTime by fee deduction cycle. It should be consistent with the billing cycle for the month in the expense center. You can go to the top of the [bill overview](https://console.cloud.tencent.com/expense/bill/overview) page to view and confirm your billing cycle type.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// The start time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end time of the period in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month queries are not supported and the query results are data for the entire month. Data within the last 18 months can be pulled at most.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Total number of records for access list needed for frontend pagination
	// 1: needed, 0: not needed
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Queries information on a specified product
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Billing mode: prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Queries information on a specified resource
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Hourly settlement
	// Daily settlement
	// Yearly/monthly subscription
	// Spot
	// New yearly/monthly subscription
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription specification adjustment
	// Yearly/monthly subscription refund
	// Adjustment - deduction
	// Adjustment - refund
	// Hourly RI fee
	// One-off RI Fee
	// Hourly Savings Plan fee
	// Offline project deduction
	// Offline product deduction
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Project ID: Project ID of the resource
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Product name code
	// Remark: If needed to obtain BusinessCode used in current month, invoke API: <a href="https://www.tencentcloud.com/document/product/555/35761?from_cn_redirect=1">Get fee distribution by product</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Context information returned from the last request. Paginated query of data for months with Month>=2023-05 can speed up query speed. Recommended for users with data volume at tens of thousands level. Query speed can be improved by 2-10x.
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.
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
	// Detail list
	DetailSet []*BillDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// Total record count, cached once every 24 hours, may be less than the actual total record count
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Context information of this request can be used in the request parameter of the next request to accelerate query speed.
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
	// Billing mode. Enumeration values
	// billOverview=L0-PDF Bill
	// Bill Summary=L1-Summary Bill	
	// billResource=L2-Resource bill	
	// billDetail=L3-Detailed Bill	
	// billPack
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Billing month
	// Earliest start month supported is 2021-01
	// L0-PDF&bill package does not support download for the current month. Please download the monthly bill after billing on the 1st of next month at 19:00.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Downloaded account ID list. By default, the query returns the account statement of the current account. If the group management account needs to download the self-pay bills of member accounts, enter the member account UIN in this field.
	ChildUin []*string `json:"ChildUin,omitnil,omitempty" name:"ChildUin"`
}

type DescribeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Billing mode. Enumeration values
	// billOverview=L0-PDF Bill
	// Bill Summary=L1-Summary Bill	
	// billResource=L2-Resource bill	
	// billDetail=L3-Detailed Bill	
	// billPack
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Billing month
	// Earliest start month supported is 2021-01
	// L0-PDF&bill package does not support download for the current month. Please download the monthly bill after billing on the 1st of next month at 19:00.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Downloaded account ID list. By default, the query returns the account statement of the current account. If the group management account needs to download the self-pay bills of member accounts, enter the member account UIN in this field.
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
	// Whether the bill file is ready. 0: file generating, 1: file generated
	Ready *int64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// Billing file download link, valid for 1 day
	// Note: This field may return null, indicating that no valid values can be obtained.
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
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// Indicates whether the total number of records is required, used for pagination.
	// Valid values: `1` (required), `0` (not required).
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// Transaction type. This parameter needs to be input using the `ActionTypeName` value. Valid values:
	// Yearly/monthly subscription purchase
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription upgrade/downgrade
	// Yearly/monthly subscription refund 
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
	// Yearly/monthly subscription resource migration in 
	// Yearly/monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Yearly/monthly subscription to pay-as-you-go 
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
	// Yearly/monthly subscription purchase
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription upgrade/downgrade
	// Yearly/monthly subscription refund 
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
	// Yearly/monthly subscription resource migration in 
	// Yearly/monthly subscription resource migration out 
	// Prepaid 
	// Hourly 
	// RI refund 
	// Pay-as-you-go reversal 
	// Yearly/monthly subscription to pay-as-you-go 
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
	// Yearly/monthly subscription
	// Spot
	// New yearly/monthly subscription
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription specification adjustment
	// Yearly/monthly subscription refund
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
	// Yearly/monthly subscription
	// Spot
	// New yearly/monthly subscription
	// Yearly/monthly subscription renewal
	// Yearly/monthly subscription specification adjustment
	// Yearly/monthly subscription refund
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
	// <p>Bill month, formatted as 2023-04</p>
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// <p>Billing dimension type. Enumeration values as follows: business, project, region, payMode, tag</p>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// <p>Tag key. Pass GroupType=tag when obtaining dimensional billing by tag.</p>
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>Operator UIN: Operator account ID (ID of the prepaid resource order or postpaid operation, activate postpaid resource account or role ID)</p>
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// <p>Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.</p>
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Bill month, formatted as 2023-04</p>
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// <p>Billing dimension type. Enumeration values as follows: business, project, region, payMode, tag</p>
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// <p>Tag key. Pass GroupType=tag when obtaining dimensional billing by tag.</p>
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// <p>Operator UIN: Operator account ID (ID of the prepaid resource order or postpaid operation, activate postpaid resource account or role ID)</p>
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// <p>Account ID of the payer (Account ID is the unique account identifier for the user in Tencent Cloud). By default, the query returns the account statement of the current account. If the group management account needs to query the self-pay bills of member accounts, enter the member account UIN in this field.</p>
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	delete(f, "OperateUin")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryResponseParams struct {
	// <p>Data readiness, 0 preparing, 1 ready. (Ready=0 indicates the first time initialization billing is in progress for the present UIN, is expected to take 5-10 minutes. Just retry after 10 minutes.)</p>
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// <p>Multidimensional bill summary of consumption detail</p>
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
type DescribeCPQBillingMappingRequestParams struct {
	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `100`.	
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Quoted subproduct name
	SpuName *string `json:"SpuName,omitnil,omitempty" name:"SpuName"`

	// Quoted product name
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`

	// Product name
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Subproduct name
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Component type name
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Component name
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`
}

type DescribeCPQBillingMappingRequest struct {
	*tchttp.BaseRequest
	
	// Offset
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// The number of entries returned at a time. The maximum value is `100`.	
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Quoted subproduct name
	SpuName *string `json:"SpuName,omitnil,omitempty" name:"SpuName"`

	// Quoted product name
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`

	// Product name
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// Subproduct name
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Component type name
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Component name
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`
}

func (r *DescribeCPQBillingMappingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPQBillingMappingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SpuName")
	delete(f, "CategoryName")
	delete(f, "BusinessName")
	delete(f, "ProductName")
	delete(f, "ComponentName")
	delete(f, "ItemName")
	delete(f, "BusinessCode")
	delete(f, "ProductCode")
	delete(f, "ComponentCode")
	delete(f, "ItemCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCPQBillingMappingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCPQBillingMappingResponseParams struct {
	// Return data details
	ResourceSpuSet []*ResourceSpuSet `json:"ResourceSpuSet,omitnil,omitempty" name:"ResourceSpuSet"`

	// 10
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCPQBillingMappingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCPQBillingMappingResponseParams `json:"Response"`
}

func (r *DescribeCPQBillingMappingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCPQBillingMappingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostDetailRequestParams struct {
	// The number of entries returned at a time. The maximum value is `100`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cycle start time. The query granularity is daily. The hour/minute/second parameter must be input in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after consumption bill is enabled and within the past 18 months.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Cycle end time. The query granularity is daily. The hour-minute-second parameter must be imported in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If this field is present, Month becomes invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after consumption bill is enabled and within the past 18 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether the total number of records in the list is needed, for frontend pagination1: needed, 0: not needed
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. It cannot be earlier than the month when the consumption bill was enabled. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Query information of a specified product
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

	// Cycle start time. The query granularity is daily. The hour/minute/second parameter must be input in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be entered, and if this field is present, Month becomes invalid. BeginTime and EndTime must be entered together, and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after consumption bill is enabled and within the past 18 months.
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Cycle end time. The query granularity is daily. The hour-minute-second parameter must be imported in the format of yyyy-mm-dd hh:ii:ss. Either Month or BeginTime&EndTime must be specified. If this field is present, Month becomes invalid. BeginTime and EndTime must be specified together and must be in the same month. Cross-month retrieval is not currently supported. Data retrievable is the data after consumption bill is enabled and within the past 18 months.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Whether the total number of records in the list is needed, for frontend pagination1: needed, 0: not needed
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// The month is in the format of yyyy-mm. Either Month or BeginTime&EndTime must be specified. If BeginTime&EndTime is specified, the Month field is invalid. It cannot be earlier than the month when the consumption bill was enabled. Data within the last 18 months can be pulled at most.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Query information of a specified product
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

	// Fee type: cost-discounted total cost, totalCost-original price cost
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

	// Fee type: cost-discounted total cost, totalCost-original price cost
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
	// Number of data records
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Header information.
	Header *AnalyseHeaderDetail `json:"Header,omitnil,omitempty" name:"Header"`

	// Data details
	Detail []*AnalyseDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// data total
	TotalDetail *AnalyseDetail `json:"TotalDetail,omitnil,omitempty" name:"TotalDetail"`

	// filtering box
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

	// Consumption details summarized by product
	Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// Record count. If NeedRecordNum is 0, null is returned.
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

	// Record count. If NeedRecordNum is 0, null is returned.
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

	// Record count. If NeedRecordNum is 0, null is returned.
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

	// Consumption details
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// Filter criteria
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConditionValue *ConsumptionResourceSummaryConditionValue `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`

	// Record countNote: This field may return null, indicating that no valid values can be obtained.
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Resource consumption details
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

	// Order status
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`
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

	// Order status
	StatusSet []*int64 `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`
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
	delete(f, "StatusSet")
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
	// The start date of the usage query in the format of yyyy-mm-dd, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// The end date of the usage query in the format of yyyy-mm-dd, such as `2020-09-30`. (The end date must be in the same month as the start date. Cross-month queries are not supported.)
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Bucket name. You can use `Get Service` to query the list of all buckets under a requester account. For details, see [GET Service (List Buckets)](https://www.tencentcloud.com/document/product/436/8291).
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// The start date of the usage query in the format of yyyy-mm-dd, such as `2020-09-01`.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// The end date of the usage query in the format of yyyy-mm-dd, such as `2020-09-30`. (The end date must be in the same month as the start date. Cross-month queries are not supported.)
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
type DescribeGatherResourceRequestParams struct {
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Resource directory category, with the enumerated values as follows:
	// all - All 
	// none - Not collected
	GatherType *string `json:"GatherType,omitnil,omitempty" name:"GatherType"`

	// Sorting field, with the enumerated values as follows:
	// realCost - Discounted total
	// cashPayAmount - Cash amount
	// voucherPayAmount - Amount of promo voucher
	// incentivePayAmount - Amount of free credit
	// transferPayAmount - Royalty amount
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Region ID, used for filtering
	RegionIds []*uint64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Project ID, used for filtering
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`
}

type DescribeGatherResourceRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. If Offset is 0, it indicates the first page. If Limit is 100, then Offset is 100, and it indicates the second page. If Offset is 200, it indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Bill month, in the format of 2024-02, which is the current month by default if not provided.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Unique identifier of a billing unit, used for filtering
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Resource directory category, with the enumerated values as follows:
	// all - All 
	// none - Not collected
	GatherType *string `json:"GatherType,omitnil,omitempty" name:"GatherType"`

	// Sorting field, with the enumerated values as follows:
	// realCost - Discounted total
	// cashPayAmount - Cash amount
	// voucherPayAmount - Amount of promo voucher
	// incentivePayAmount - Amount of free credit
	// transferPayAmount - Royalty amount
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// Sorting type, with the enumerated values as follows:
	// asc - Ascending
	// desc - Descending
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// Product code, used for filtering
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// Subproduct code, used for filtering
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// Component name code, used for filtering
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// Region ID, used for filtering
	RegionIds []*uint64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// Instance type code, used for filtering
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// Billing mode code, used for filtering
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// Operator UIN, used for filtering
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// User UIN, used for filtering
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// Fuzzy search: supports tag, resource ID, and resource alias.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// Tag, used for filtering
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Project ID, used for filtering
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// Transaction type code, used for filtering
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`
}

func (r *DescribeGatherResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "TreeNodeUniqKey")
	delete(f, "GatherType")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BusinessCodes")
	delete(f, "ProductCodes")
	delete(f, "ItemCodes")
	delete(f, "RegionIds")
	delete(f, "InstanceTypes")
	delete(f, "PayModes")
	delete(f, "OperateUins")
	delete(f, "OwnerUins")
	delete(f, "SearchKey")
	delete(f, "Tag")
	delete(f, "ProjectIds")
	delete(f, "ActionTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatherResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatherResourceResponseParams struct {
	// Total quantity.
	RecordNum *int64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// Resource collection summary
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GatherResourceSummary []*GatherResourceSummary `json:"GatherResourceSummary,omitnil,omitempty" name:"GatherResourceSummary"`

	// Data update time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatherResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatherResourceResponseParams `json:"Response"`
}

func (r *DescribeGatherResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherResourceResponse) FromJsonString(s string) error {
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
type DescribeRenewInstancesRequestParams struct {
	// Maximum number of instances per page. Value range: 1-100.
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// Token for querying the next page of returned results. NextToken is not needed when calling the API for the first time.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Get the sorting order of the instance. The enumerated values are as follows:
	// false = Ascending (default)
	// true=Descending
	Reverse *bool `json:"Reverse,omitnil,omitempty" name:"Reverse"`

	// Renewal flag. Multiple values separated by commas. Enumeration value as follows:
	// NOTIFY_AND_MANUAL_RENEW: manual renewal.
	// NOTIFY_AND_AUTO_RENEW: automatic renewal.
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: non-renewal upon expiration.
	RenewFlagList []*string `json:"RenewFlagList,omitnil,omitempty" name:"RenewFlagList"`

	// Instance ID. Multiple IDs separated by commas, with a maximum of 100.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Expiry time start, format yyyy-MM-dd HH:mm:ss.
	ExpireTimeStart *string `json:"ExpireTimeStart,omitnil,omitempty" name:"ExpireTimeStart"`

	// Expiry time in the format of yyyy-MM-dd HH:mm:ss.
	ExpireTimeEnd *string `json:"ExpireTimeEnd,omitnil,omitempty" name:"ExpireTimeEnd"`
}

type DescribeRenewInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of instances per page. Value range: 1-100.
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// Token for querying the next page of returned results. NextToken is not needed when calling the API for the first time.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// Get the sorting order of the instance. The enumerated values are as follows:
	// false = Ascending (default)
	// true=Descending
	Reverse *bool `json:"Reverse,omitnil,omitempty" name:"Reverse"`

	// Renewal flag. Multiple values separated by commas. Enumeration value as follows:
	// NOTIFY_AND_MANUAL_RENEW: manual renewal.
	// NOTIFY_AND_AUTO_RENEW: automatic renewal.
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: non-renewal upon expiration.
	RenewFlagList []*string `json:"RenewFlagList,omitnil,omitempty" name:"RenewFlagList"`

	// Instance ID. Multiple IDs separated by commas, with a maximum of 100.
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// Expiry time start, format yyyy-MM-dd HH:mm:ss.
	ExpireTimeStart *string `json:"ExpireTimeStart,omitnil,omitempty" name:"ExpireTimeStart"`

	// Expiry time in the format of yyyy-MM-dd HH:mm:ss.
	ExpireTimeEnd *string `json:"ExpireTimeEnd,omitnil,omitempty" name:"ExpireTimeEnd"`
}

func (r *DescribeRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRenewInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "Reverse")
	delete(f, "RenewFlagList")
	delete(f, "InstanceIdList")
	delete(f, "ExpireTimeStart")
	delete(f, "ExpireTimeEnd")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRenewInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRenewInstancesResponseParams struct {
	// Instance summary list.
	InstanceList []*RenewInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// Token for querying the next page of returned results.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRenewInstancesResponseParams `json:"Response"`
}

func (r *DescribeRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagListRequestParams struct {
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Offset=0 indicates the first page. If Limit=100, Offset=100 indicates the second page, Offset=200 indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cost allocation tag key, used as fuzzy search
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type, enumeration value: 0 ordinary tag, 1 allocation tag, used for filtering. If not passed, get all tag keys.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting method, enumeration value: asc for ascending order, desc for descending order.
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeTagListRequest struct {
	*tchttp.BaseRequest
	
	// Quantity, with the maximum value of 1,000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Offset=0 indicates the first page. If Limit=100, Offset=100 indicates the second page, Offset=200 indicates the third page, and so on.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Cost allocation tag key, used as fuzzy search
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type, enumeration value: 0 ordinary tag, 1 allocation tag, used for filtering. If not passed, get all tag keys.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Sorting method, enumeration value: asc for ascending order, desc for descending order.
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
	// Total number of records
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
	// <p>How many data entries per page, 20 is selected by default, maximum not exceeding 1000</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page number, starts from 1 by default</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Voucher status: pending use: unUsed, Used: used, delivered: delivered, discarded: cancel, expired: overdue</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Voucher id</p>
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// <p>Voucher order id</p>
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// <p>product code</p>
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// <p>Activity id</p>
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// <p>Voucher name</p>
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// <p>Start time of delivery. Example: 2021-01-01</p>
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// <p>Delivery end time. Example: 2021-01-01</p>
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// <p>Specified sorting field: BeginTime start time, EndTime expiry time, CreateTime creation time</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>Specify ascending/descending order: desc, asc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>Payment mode, postPay (postpaid)/prePay (prepaid)/riPay (reserved instance)/"" or "*" means all modes. If payMode is "" or "*", productCode and subProductCode must be empty.</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Payment scenario PayMode=postPay: spotpay - spot instance, "settle account" - standard post-paid. PayMode=prePay: purchase - monthly subscription new purchase, renew - annual and monthly renewal (auto renewal), modify - monthly subscription configuration change. PayMode=riPay: oneOffFee - prepayment of reserved instance, hourlyFee - hourly charge for reserved instance, * - support all payment scenarios.</p>
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// <p>Operator is used by default as user uin</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>The primary types of vouchers are has_price and no_price, which represent the cash voucher with a price and the cash voucher without a price respectively.</p>
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// <p>Voucher subtype: Discount is a discount voucher, and deduct is a deduction voucher.</p>
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`

	// <p>Voucher validity start time</p>
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// <p>Voucher validity time end time</p>
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// <p>Voucher expiration time start time</p>
	EndTimeFrom *string `json:"EndTimeFrom,omitnil,omitempty" name:"EndTimeFrom"`

	// <p>Voucher expiration time end time</p>
	EndTimeTo *string `json:"EndTimeTo,omitnil,omitempty" name:"EndTimeTo"`

	// <p>Voucher issuance start time</p>
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// <p>Voucher issuance time end time</p>
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// <p>Language parameter</p><p>Default value: zh</p><p>Expect the product name to return in Chinese or other languages. Currently only support Chinese and English. Return in Chinese when "zh" is entered or left blank; return in English in other cases.</p>
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>How many data entries per page, 20 is selected by default, maximum not exceeding 1000</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Page number, starts from 1 by default</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Voucher status: pending use: unUsed, Used: used, delivered: delivered, discarded: cancel, expired: overdue</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Voucher id</p>
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// <p>Voucher order id</p>
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// <p>product code</p>
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// <p>Activity id</p>
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// <p>Voucher name</p>
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// <p>Start time of delivery. Example: 2021-01-01</p>
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// <p>Delivery end time. Example: 2021-01-01</p>
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// <p>Specified sorting field: BeginTime start time, EndTime expiry time, CreateTime creation time</p>
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// <p>Specify ascending/descending order: desc, asc</p>
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// <p>Payment mode, postPay (postpaid)/prePay (prepaid)/riPay (reserved instance)/"" or "*" means all modes. If payMode is "" or "*", productCode and subProductCode must be empty.</p>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>Payment scenario PayMode=postPay: spotpay - spot instance, "settle account" - standard post-paid. PayMode=prePay: purchase - monthly subscription new purchase, renew - annual and monthly renewal (auto renewal), modify - monthly subscription configuration change. PayMode=riPay: oneOffFee - prepayment of reserved instance, hourlyFee - hourly charge for reserved instance, * - support all payment scenarios.</p>
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// <p>Operator is used by default as user uin</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// <p>The primary types of vouchers are has_price and no_price, which represent the cash voucher with a price and the cash voucher without a price respectively.</p>
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// <p>Voucher subtype: Discount is a discount voucher, and deduct is a deduction voucher.</p>
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`

	// <p>Voucher validity start time</p>
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// <p>Voucher validity time end time</p>
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// <p>Voucher expiration time start time</p>
	EndTimeFrom *string `json:"EndTimeFrom,omitnil,omitempty" name:"EndTimeFrom"`

	// <p>Voucher expiration time end time</p>
	EndTimeTo *string `json:"EndTimeTo,omitnil,omitempty" name:"EndTimeTo"`

	// <p>Voucher issuance start time</p>
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// <p>Voucher issuance time end time</p>
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	// <p>Language parameter</p><p>Default value: zh</p><p>Expect the product name to return in Chinese or other languages. Currently only support Chinese and English. Return in Chinese when "zh" is entered or left blank; return in English in other cases.</p>
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
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
	delete(f, "StartTimeFrom")
	delete(f, "StartTimeTo")
	delete(f, "EndTimeFrom")
	delete(f, "EndTimeTo")
	delete(f, "CreateTimeFrom")
	delete(f, "CreateTimeTo")
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// <p>Total count</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Total balance (differential)</p>
	TotalBalance *int64 `json:"TotalBalance,omitnil,omitempty" name:"TotalBalance"`

	// <p>Voucher information related to</p>
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitnil,omitempty" name:"VoucherInfos"`

	// <p>Unit of the amount field in the API response</p><p>Default value: micro</p><p>Currency unit: micro (microcent)<br>Voucher issuance and use are processed with 8-digit high-precision, so the currency unit defaults to micro (microcent). If CNY or USD is needed, convert using the following formula:<br>CNY: 1 micro = 0.00000001  yuan<br>USD: 1 micro = 0.00000001  USD</p>
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

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

	// Voucher usage record details
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
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Region ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Price attribute: Other attributes of the component that affect discount pricing besides unit price and duration
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// Associated transaction document ID: Document ID associated with this transaction, such as a write-off order, the original order, a resettlement order, or the original purchase order number recorded in a refund order.
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// Calculation explanation: A detailed explanation to calculations of billing settlement for special transaction types, such as refund and configuration changes.
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// Billing Rules: The detailed billing rules for each product shown in the portal explanation link
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// Billing month
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Billing day
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`
}

type ExcludedProducts struct {
	// The names of non-applicable products.
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// `postPay`: pay-as-you-go; `prePay`: prepaid; `riPay`: reserved instance; empty or `*`: all.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type GatherResourceSummary struct {
	// Payer UIN: Account ID of the payer, which is the unique account identifier for the user in Tencent Cloud.
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// User UIN: Account ID of the actual resource user
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Operator account ID (the resource account ID or role ID opened by prepaid resource ordering or postpaid operation)
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// Instance type code
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Instance type: The type of an instance corresponding to the product service purchased, including resource packages, RI, SP, and spot instances. It is displayed as "-" by default for regular instances.
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// Resource ID: Resources vary by product, and the content is not identical. For example, Cloud Virtual Machine (CVM) corresponds to the instance ID. If the product is split, it shows the split item ID, such as COS bucket ID and CDN domain name.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Instance name: The name set by the user for the resource in the console, which is empty by default if not set. If the product is split, it shows the split resource alias.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// Unique identifier of a cost allocation unit
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// Name of a cost allocation unit
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// Allocation rule ID hit by the resource
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// Allocation rule name hit by the resource
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name: Various cloud products purchased by users
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Component name code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name: The specific component of a product or service purchased by the user
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// Region ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name: The region where the resource is located
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Allocation tag: The resource-bound tag
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Discounted total: discounted total = (Original Price - Original Price Deducted by a Reserved Instance - Savings Plan Deduction from Original Price) * Discount Rate
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash account expenditure (CNY): The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Promo voucher expenditure (CNY): The amount paid using various vouchers (such as promo vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Gift account expenditure (CNY): The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Royalty account expenditure (CNY): The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Cost collection type: The source types of fees, including allocated, collection and unallocated.
	// 0 - Allocation
	// 1 - Collection
	// -1 - Unallocated
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// Information of the current allocation unit
	BelongTreeNodeUniqKey *AllocationTreeNode `json:"BelongTreeNodeUniqKey,omitnil,omitempty" name:"BelongTreeNodeUniqKey"`

	// Information on allocation rules hit by the current resource
	BelongRule *AllocationRule `json:"BelongRule,omitnil,omitempty" name:"BelongRule"`

	// Information on other allocation units
	OtherTreeNodeUniqKeys []*AllocationTreeNode `json:"OtherTreeNodeUniqKeys,omitnil,omitempty" name:"OtherTreeNodeUniqKeys"`

	// Information on other hit rules
	OtherRules []*AllocationRule `json:"OtherRules,omitnil,omitempty" name:"OtherRules"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name: The Project to which a resource belongs, which is independently allocated by the user for the resource in the console. If a resource has not been allocated to an Project, it will be a default Project.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name: Product subdivision type purchased by the user
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// Billing mode code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode: Resource billing mode, which can be monthly subscription or pay-as-you-go.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// Transaction type code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type: Detailed transaction type
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// Split item ID: The ID of the split item involved in the split product, such as COS bucket ID and CDN domain name
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// Split item name: The split item involved in the split product
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`
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

type OperateRsp struct {
	// Operation failure code at the instance dimension
	// Note: This field may return null, indicating that no valid values can be obtained.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Failure reason for operating related resources
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Instance ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

// Predefined struct for user
type PayDealsRequestParams struct {
	// Specifies one or more Sub-order No. that need to pay. must pass either this parameter or the BigDealIds field, but not both.
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// Whether to automatically use a voucher. valid values: 1 (yes), 0 (no). default: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. currently only supports specifying one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Specifies one or more Order No. that need to pay. must pass either this parameter or the OrderIds field.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 0 self pay, 3 group agent, 4 reseller places a product-level payment order for customers. default 0.
	AgentPay *int64 `json:"AgentPay,omitnil,omitempty" name:"AgentPay"`

	// Disregard it.
	CpsUin *string `json:"CpsUin,omitnil,omitempty" name:"CpsUin"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest
	
	// Specifies one or more Sub-order No. that need to pay. must pass either this parameter or the BigDealIds field, but not both.
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// Whether to automatically use a voucher. valid values: 1 (yes), 0 (no). default: 0.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. currently only supports specifying one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Specifies one or more Order No. that need to pay. must pass either this parameter or the OrderIds field.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 0 self pay, 3 group agent, 4 reseller places a product-level payment order for customers. default 0.
	AgentPay *int64 `json:"AgentPay,omitnil,omitempty" name:"AgentPay"`

	// Disregard it.
	CpsUin *string `json:"CpsUin,omitnil,omitempty" name:"CpsUin"`
}

func (r *PayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIds")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "BigDealIds")
	delete(f, "AgentPay")
	delete(f, "CpsUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PayDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PayDealsResponseParams struct {
	// Specifies the array of Sub-order No. with payment successful.
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// Specifies the Id array of resources with payment successful.
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// Specifies the array of Order No. with payment successful.
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PayDealsResponse struct {
	*tchttp.BaseResponse
	Response *PayDealsResponseParams `json:"Response"`
}

func (r *PayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayDealsResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type RefundInstanceRequestParams struct {
	// ClientToken is a unique, case-sensitive string generated by the client, no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544****.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`
}

type RefundInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ClientToken is a unique, case-sensitive string generated by the client, no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544****.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`
}

func (r *RefundInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientToken")
	delete(f, "ProductCode")
	delete(f, "SubProductCode")
	delete(f, "InstanceId")
	delete(f, "RegionCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundInstanceResponseParams struct {
	// Order ID list
	OrderIdList []*string `json:"OrderIdList,omitnil,omitempty" name:"OrderIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefundInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RefundInstanceResponseParams `json:"Response"`
}

func (r *RefundInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSummaryOverviewItem struct {
	// Region ID
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

type RenewInstance struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Product code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct code
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Region encoding
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Instance status:
	// NORMAL
	// ISOLATED Isolated
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Renewal flag:
	// NOTIFY_AND_MANUAL_RENEW: manual renewal
	// NOTIFY_AND_AUTO_RENEW: auto-renewal.
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: non-renewal upon expiration.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Instance expiration time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Instance alias: The name set by the user for the instance in the console, which is empty by default if not set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Product name: Cloud products purchased by users, such as Cloud Virtual Machine (CVM)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Project name: Instance Ownership of the project. User can autonomously assign project to the instance on the console. Default project if not allocated.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Automatic renewal cycle length
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewPeriod *uint64 `json:"RenewPeriod,omitnil,omitempty" name:"RenewPeriod"`

	// Automatic renewal cycle unit: y year, m month
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewPeriodUnit *string `json:"RenewPeriodUnit,omitnil,omitempty" name:"RenewPeriodUnit"`
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// ClientToken is a unique, case-sensitive string generated by the client, with no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Manual renewal duration, upper limit: 36, default value is 1.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Manual renewal duration unit. available values: 
	// m: renew monthly, 
	// y: renew annually. 
	// default value is m.
	PeriodUnit *string `json:"PeriodUnit,omitnil,omitempty" name:"PeriodUnit"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ClientToken is a unique, case-sensitive string generated by the client, with no more than 64 ASCII characters. for example, ClientToken=123e4567-e89b-12d3-a456-42665544.
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code.
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Manual renewal duration, upper limit: 36, default value is 1.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Manual renewal duration unit. available values: 
	// m: renew monthly, 
	// y: renew annually. 
	// default value is m.
	PeriodUnit *string `json:"PeriodUnit,omitnil,omitempty" name:"PeriodUnit"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientToken")
	delete(f, "ProductCode")
	delete(f, "SubProductCode")
	delete(f, "RegionCode")
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "PeriodUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// Order ID list
	OrderIdList []*string `json:"OrderIdList,omitnil,omitempty" name:"OrderIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSpuSet struct {
	// Quoted subproduct (Chinese)
	SpuNameZh *string `json:"SpuNameZh,omitnil,omitempty" name:"SpuNameZh"`

	// Quoted subproduct (English)
	SpuNameEn *string `json:"SpuNameEn,omitnil,omitempty" name:"SpuNameEn"`

	// Quoted product (Chinese)
	CategoryNameZh *string `json:"CategoryNameZh,omitnil,omitempty" name:"CategoryNameZh"`

	// Quoted product (English)
	CategoryNameEn *string `json:"CategoryNameEn,omitnil,omitempty" name:"CategoryNameEn"`

	// Product code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name (Chinese)
	BusinessNameZh *string `json:"BusinessNameZh,omitnil,omitempty" name:"BusinessNameZh"`

	// Product name (English)
	BusinessNameEn *string `json:"BusinessNameEn,omitnil,omitempty" name:"BusinessNameEn"`

	// Subproduct code
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Subproduct name (Chinese)
	ProductNameZh *string `json:"ProductNameZh,omitnil,omitempty" name:"ProductNameZh"`

	// Subproduct name (English)
	ProductNameEn *string `json:"ProductNameEn,omitnil,omitempty" name:"ProductNameEn"`

	// Component type code
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// Component type name (Chinese)
	ComponentNameZh *string `json:"ComponentNameZh,omitnil,omitempty" name:"ComponentNameZh"`

	// Component type name (English)
	ComponentNameEn *string `json:"ComponentNameEn,omitnil,omitempty" name:"ComponentNameEn"`

	// Component code
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// Component name (Chinese)
	ItemNameZh *string `json:"ItemNameZh,omitnil,omitempty" name:"ItemNameZh"`

	// Component name (English)
	ItemNameEn *string `json:"ItemNameEn,omitnil,omitempty" name:"ItemNameEn"`
}

// Predefined struct for user
type SetRenewalRequestParams struct {
	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Instance ID. Only one can be specified.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Renewal flag. Enumeration values are as follows:
	// NOTIFY_AND_MANUAL_RENEW: manual renewal.
	// NOTIFY_AND_AUTO_RENEW: automatic renewal.
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: non-renewal upon expiration.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Automatic renewal cycle length. If left empty, the default value set by product is used.
	// If it is month, support: 1-11
	// If it is year, support: 1-5.
	// Supported range mainly depends on the product side.
	RenewPeriod *string `json:"RenewPeriod,omitnil,omitempty" name:"RenewPeriod"`

	// Automatic renewal cycle unit. If left empty, the default value set by the product is used.
	// Year y, month m
	// Supported range mainly depends on the product side.
	RenewPeriodUnit *string `json:"RenewPeriodUnit,omitnil,omitempty" name:"RenewPeriodUnit"`
}

type SetRenewalRequest struct {
	*tchttp.BaseRequest
	
	// Product code.
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Region code.
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// Instance ID. Only one can be specified.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Renewal flag. Enumeration values are as follows:
	// NOTIFY_AND_MANUAL_RENEW: manual renewal.
	// NOTIFY_AND_AUTO_RENEW: automatic renewal.
	// DISABLE_NOTIFY_AND_MANUAL_RENEW: non-renewal upon expiration.
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Automatic renewal cycle length. If left empty, the default value set by product is used.
	// If it is month, support: 1-11
	// If it is year, support: 1-5.
	// Supported range mainly depends on the product side.
	RenewPeriod *string `json:"RenewPeriod,omitnil,omitempty" name:"RenewPeriod"`

	// Automatic renewal cycle unit. If left empty, the default value set by the product is used.
	// Year y, month m
	// Supported range mainly depends on the product side.
	RenewPeriodUnit *string `json:"RenewPeriodUnit,omitnil,omitempty" name:"RenewPeriodUnit"`
}

func (r *SetRenewalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductCode")
	delete(f, "RegionCode")
	delete(f, "InstanceId")
	delete(f, "RenewFlag")
	delete(f, "RenewPeriod")
	delete(f, "RenewPeriodUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRenewalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRenewalResponseParams struct {
	// Instance list when the operation fails.
	InstanceList []*OperateRsp `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetRenewalResponse struct {
	*tchttp.BaseResponse
	Response *SetRenewalResponseParams `json:"Response"`
}

func (r *SetRenewalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRenewalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryDetail struct {
	// Bill dimension code
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	// Billing dimension value
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

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Product summary information
	Business []*BusinessSummaryInfo `json:"Business,omitnil,omitempty" name:"Business"`
}

type SummaryTotal struct {
	// discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Original price in CNY. The TotalCost field comes into effect after bill 3.0 (May 2021) and returns "-" before bill 3.0. In the current situation of contract price, the TotalCost field returns "-" if a price difference exists with the official website price.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type TagDataInfo struct {
	// Cost allocation tag key.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag type. Valid values: `0` (general tags), `1` (cost allocation tags).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Set the allocation tag time. Ordinary tags do not return.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TagSummaryOverviewItem struct {
	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// Percentage of the fee, with 2 decimal places.
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// discounted total price
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// Cash account expenditure: The amount paid through the cash account
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// Gift account expenditure: The amount paid using free credits
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// Coupon expenditure: The amount paid using various vouchers (such as vouchers and cash vouchers)
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Royalty account expenditure: The amount paid through the royalty account
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// Original price in CNY. The TotalCost field comes into effect after bill 3.0 (May 2021) and returns "-" before bill 3.0. In the current situation of contract price, the TotalCost field returns "-" if a price difference exists with the official website price.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type UinTempAmountModel struct {
	// User UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// temporary limit
	TempAmount *float64 `json:"TempAmount,omitnil,omitempty" name:"TempAmount"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type UsageDetails struct {
	// Product name
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// product details
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// Product code	
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// Sub-product code	
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// Billing item code.	
	BillingItemCode *string `json:"BillingItemCode,omitnil,omitempty" name:"BillingItemCode"`

	// Billing sub-item code.	
	SubBillingItemCode *string `json:"SubBillingItemCode,omitnil,omitempty" name:"SubBillingItemCode"`

	// Product English Name	
	ProductEnName *string `json:"ProductEnName,omitnil,omitempty" name:"ProductEnName"`

	// English name of the sub-product.	
	SubProductEnName *string `json:"SubProductEnName,omitnil,omitempty" name:"SubProductEnName"`

	// billing cycle	
	CalcUnit *string `json:"CalcUnit,omitnil,omitempty" name:"CalcUnit"`

	// payMode is prepay and payScene is common in the current situation
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type UsageRecords struct {
	// The amount used. The value of this parameter is the amount used (USD, rounded to 8 decimal places) multiplied by 100,000,000.
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// The time when the voucher was used.
	UsedTime *string `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// Usage record details
	UsageDetails []*UsageDetails `json:"UsageDetails,omitnil,omitempty" name:"UsageDetails"`

	// Payment mode
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Queried coupon id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// Transaction scene: (adjust: adjust accounts, common: normal transaction scene)
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// Unique ID, corresponding to transaction: prepaid dealName, bill adjustment/postpaid outTradeNo
	SeqId *string `json:"SeqId,omitnil,omitempty" name:"SeqId"`
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

	// Product information applies to
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitnil,omitempty" name:"ApplicableProducts"`

	// Product information not applicable
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitnil,omitempty" name:"ExcludedProducts"`

	// Instructions/Batch Remarks
	PolicyRemark *string `json:"PolicyRemark,omitnil,omitempty" name:"PolicyRemark"`

	// Coupon issuance time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}