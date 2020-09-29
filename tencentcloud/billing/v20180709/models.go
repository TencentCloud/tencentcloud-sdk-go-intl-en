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
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitempty" name:"ComponentSet" list`

	// Payer's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// User's UIN
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator's UIN
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Tag information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags" list`

	// Product name/code (optional)
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// Subproduct name/code (optional)
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// Transaction type/code (optional)
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
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

	// Component type/code (optional)
	ItemCode *string `json:"ItemCode,omitempty" name:"ItemCode"`

	// Component name/code (optional)
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// Contract price
	ContractPrice *string `json:"ContractPrice,omitempty" name:"ContractPrice"`
}

type BillResourceSummary struct {

	// Product name: major categories of Tencent Cloud services, e.g. CVM and TencentDB for MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// Sub-product name: sub-categories of Tencent Cloud services, such as CVM-Standard S1; if no subproduct name is obtained, '-' is returned.
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

	// Discount rate
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
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags" list`

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Resource owner UIN; '-' is returned if no value is obtained
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// Operator UIN; '-' is returned if no value is obtained
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type BillTagInfo struct {

	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BusinessSummaryOverviewItem struct {

	// Product code
	// Note: This field may return null, indicating that no valid value was found.
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

	// 
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details list
		DetailSet []*BillDetail `json:"DetailSet,omitempty" name:"DetailSet" list`

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

func (r *DescribeBillDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Quantity, maximum is 1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// The period type. byUsedTime: By usage period; byPayTime: by payment period. Must be the same as the period of the current monthly bill of the Billing Center. You can check your bill statistics period type at the top of the [Bill Overview](https://console.cloud.tencent.com/expense/bill/overview) page.
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// Month; format: yyyy-mm. This value cannot be earlier than the month when Bill 2.0 is enabled. Last 24 months data are available.
	Month *string `json:"Month,omitempty" name:"Month"`

	// Indicates whether or not the total number of records of accessing the list is required, used for frontend pages.
	// 1 = yes, 0 = no
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeBillResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillResourceSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Resource summary list
		ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitempty" name:"ResourceSummarySet" list`

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

func (r *DescribeBillResourceSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest

	// Query bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Only beginning in the current month is supported, and it must be the same month as the EndTime. For example, 2018-09-01 00:00:00.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// Only ending in the current month is supported, and it must be the same month as the BeginTime. For example, 2018-09-30 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all billing modes
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Only beginning in the current month is supported, and it must be the same month as the EndTime. For example, 2018-09-01 00:00:00.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// Only ending in the current month is supported, and it must be the same month as the BeginTime. For example, 2018-09-30 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
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
		SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Only beginning in the current month is supported, and it must be the same month as the EndTime. For example, 2018-09-01 00:00:00.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// Only ending in the current month is supported, and it must be the same month as the BeginTime. For example, 2018-09-30 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all projects
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest

	// Queries bill data user's UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Only beginning in the current month is supported, and it must be the same month as the EndTime. For example, 2018-09-01 00:00:00.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// Only ending in the current month is supported, and it must be the same month as the BeginTime. For example, 2018-09-30 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. 0 = not ready, 1 = ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Detailed cost distribution for all regions
	// Note: This field may return null, indicating that no valid value was found.
		SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest

	// Payer UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// Currently the period to be queried must start from a time point in the current month, and the starting time and the end time must be in the same month. Example: 2018-09-01 00:00:00.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// Currently the period to be queried must end at a time point in the current month, and the starting time and the end time must be in the same month. Example: 2018-09-30 23:59:59.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Cost allocation tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Indicates whether or not the data is ready. `0`: not ready; `1`: ready.
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// Details about cost distribution over different tags
	// Note: This field may return null, indicating that no valid values can be obtained.
		SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitempty" name:"Detail" list`

	// Cash amount
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// Trial credit amount
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// Voucher amount
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
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
}
