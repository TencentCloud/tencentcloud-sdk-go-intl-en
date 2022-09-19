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

package v20170312

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesRequestParams struct {
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeAvailableCvmInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAvailableCvmInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableCvmInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableCvmInstanceTypesResponseParams struct {
	// Array of model configurations
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableCvmInstanceTypesResponseParams `json:"Response"`
}

func (r *DescribeAvailableCvmInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableCvmInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosRequestParams struct {
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	// <li> instance-type - String - Required: No - (Filter) Filter by model.</li>
	// <li> instance-charge-type - String - Required: No - (Filter) Filter by instance billing method. ( POSTPAID_BY_HOUR: pay-as-you-go | SPOTPAID: bidding.)  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCvmZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest
	
	// Filter.
	// <li> zone - String - Required: No - (Filter) Filter by availability zone.</li>
	// <li> instance-family - String - Required: No - (Filter) Filter by model family such as S1, I1, and M1.</li>
	// <li> instance-type - String - Required: No - (Filter) Filter by model.</li>
	// <li> instance-charge-type - String - Required: No - (Filter) Filter by instance billing method. ( POSTPAID_BY_HOUR: pay-as-you-go | SPOTPAID: bidding.)  </li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCvmZoneInstanceConfigInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCvmZoneInstanceConfigInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCvmZoneInstanceConfigInfosResponseParams struct {
	// List of model configurations in the availability zone.
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCvmZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCvmZoneInstanceConfigInfosResponseParams `json:"Response"`
}

func (r *DescribeCvmZoneInstanceConfigInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCvmZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesRequestParams struct {

}

type DescribeInstanceCategoriesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceCategoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceCategoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceCategoriesResponseParams struct {
	// List of CVM instance categories
	InstanceCategorySet []*InstanceCategoryItem `json:"InstanceCategorySet,omitempty" name:"InstanceCategorySet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceCategoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceCategoriesResponseParams `json:"Response"`
}

func (r *DescribeInstanceCategoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceCategoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Externals struct {
	// Release address
	// Note: This field may return null, indicating that no valid value is found.
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty" name:"ReleaseAddress"`

	// Not supported network. Value: <br><li>BASIC: classic network<br><li>VPC1.0: VPC1.0
	// Note: This field may return null, indicating that no valid value was found.
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty" name:"UnsupportNetworks"`

	// Attributes of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty" name:"StorageBlockAttr"`
}

type Filter struct {
	// Filters.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Filter values.
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceCategoryItem struct {
	// Instance type name
	InstanceCategory *string `json:"InstanceCategory,omitempty" name:"InstanceCategory"`

	// List of instance families
	InstanceFamilySet []*string `json:"InstanceFamilySet,omitempty" name:"InstanceFamilySet"`
}

type InstanceTypeConfig struct {
	// Memory size in GB.
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// Number of CPU cores.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

type InstanceTypeQuotaItem struct {
	// Availability zone.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Instance model.
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// Instance billing plan. Valid values: <br><li>POSTPAID_BY_HOUR: pay after use. You are billed for your traffic by the hour.<br><li>`CDHPAID`: [`CDH`](https://intl.cloud.tencent.com/document/product/416?from_cn_redirect=1) billing plan. Applicable to `CDH` only, not the instances on the host.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// ENI type. For example, 25 represents an ENI of 25 GB.
	NetworkCard *int64 `json:"NetworkCard,omitempty" name:"NetworkCard"`

	// Additional data.
	// Note: This field may return null, indicating that no valid value is found.
	Externals *Externals `json:"Externals,omitempty" name:"Externals"`

	// Number of CPU cores of an instance model.
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// Instance memory capacity; unit: `GB`.
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// Instance model family.
	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`

	// Model name.
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// List of local disk specifications. If the parameter returns null, it means that local disks cannot be created.
	LocalDiskTypeList []*LocalDiskType `json:"LocalDiskTypeList,omitempty" name:"LocalDiskTypeList"`

	// Whether an instance model is available. Valid values: <br><li>SELL: available <br><li>SOLD_OUT: sold out
	Status *string `json:"Status,omitempty" name:"Status"`

	// Price of an instance model.
	Price *ItemPrice `json:"Price,omitempty" name:"Price"`

	// Details of out-of-stock items
	// Note: this field may return null, indicating that no valid value is obtained.
	SoldOutReason *string `json:"SoldOutReason,omitempty" name:"SoldOutReason"`

	// Private network bandwidth, in Gbps.
	InstanceBandwidth *float64 `json:"InstanceBandwidth,omitempty" name:"InstanceBandwidth"`

	// The max packet sending and receiving capability (in 10k PPS).
	InstancePps *int64 `json:"InstancePps,omitempty" name:"InstancePps"`

	// Number of local storage blocks.
	StorageBlockAmount *int64 `json:"StorageBlockAmount,omitempty" name:"StorageBlockAmount"`

	// CPU type.
	CpuType *string `json:"CpuType,omitempty" name:"CpuType"`

	// Number of GPUs of the instance.
	Gpu *int64 `json:"Gpu,omitempty" name:"Gpu"`

	// Number of FPGAs of the instance.
	Fpga *int64 `json:"Fpga,omitempty" name:"Fpga"`

	// Descriptive information of the instance.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ItemPrice struct {
	// The original unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// Billing unit for pay-as-you-go mode. Valid values: <br><li>HOUR: billed on an hourly basis. It's used for hourly postpaid instances (`POSTPAID_BY_HOUR`). <br><li>GB: bill by traffic in GB. It's used for postpaid products that are billed by the hourly traffic (`TRAFFIC_POSTPAID_BY_HOUR`).
	// Note: this field may return null, indicating that no valid value is obtained.
	ChargeUnit *string `json:"ChargeUnit,omitempty" name:"ChargeUnit"`

	// The original price of a pay-in-advance instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// Discount price of a prepaid instance, in USD.
	// Note: this field may return null, indicating that no valid value is obtained.
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

	// Percentage of the original price. For example, if you enter "20.0", the discounted price will be 20% of the original price.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// The discounted unit price for pay-as-you-go mode in USD. <br><li>When a billing tier is returned, it indicates the price fo the returned billing tier. For example, if `UnitPriceSecondStep` is returned, it refers to the unit price for the usage between 0 to 96 hours. Otherwise, it refers to that the unit price for unlimited usage.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitempty" name:"UnitPriceDiscount"`

	// Original unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceSecondStep *float64 `json:"UnitPriceSecondStep,omitempty" name:"UnitPriceSecondStep"`

	// Discounted unit price for the usage between 96 to 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountSecondStep *float64 `json:"UnitPriceDiscountSecondStep,omitempty" name:"UnitPriceDiscountSecondStep"`

	// Original unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceThirdStep *float64 `json:"UnitPriceThirdStep,omitempty" name:"UnitPriceThirdStep"`

	// Discounted unit price for the usage after 360 hours in USD. It's applicable to pay-as-you-go mode.
	// Note: this field may return null, indicating that no valid value is obtained.
	UnitPriceDiscountThirdStep *float64 `json:"UnitPriceDiscountThirdStep,omitempty" name:"UnitPriceDiscountThirdStep"`

	// Original 3-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceThreeYear *float64 `json:"OriginalPriceThreeYear,omitempty" name:"OriginalPriceThreeYear"`

	// Discounted 3-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceThreeYear *float64 `json:"DiscountPriceThreeYear,omitempty" name:"DiscountPriceThreeYear"`

	// Discount for 3-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountThreeYear *float64 `json:"DiscountThreeYear,omitempty" name:"DiscountThreeYear"`

	// Original 5-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceFiveYear *float64 `json:"OriginalPriceFiveYear,omitempty" name:"OriginalPriceFiveYear"`

	// Discounted 5-year upfront payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceFiveYear *float64 `json:"DiscountPriceFiveYear,omitempty" name:"DiscountPriceFiveYear"`

	// Discount for 5-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountFiveYear *float64 `json:"DiscountFiveYear,omitempty" name:"DiscountFiveYear"`

	// Original 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	OriginalPriceOneYear *float64 `json:"OriginalPriceOneYear,omitempty" name:"OriginalPriceOneYear"`

	// Discounted 1-year payment, in USD. This parameter is only available to upfront payment mode.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountPriceOneYear *float64 `json:"DiscountPriceOneYear,omitempty" name:"DiscountPriceOneYear"`

	// Discount for 1-year upfront payment. For example, 20.0 indicates 80% off.
	// Note: this field may return `null`, indicating that no valid value was found.
	// Note: this field may return `null`, indicating that no valid value was found.
	DiscountOneYear *float64 `json:"DiscountOneYear,omitempty" name:"DiscountOneYear"`
}

type LocalDiskType struct {
	// Type of a local disk.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Attributes of a local disk.
	PartitionType *string `json:"PartitionType,omitempty" name:"PartitionType"`

	// Minimum size of a local disk.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum size of a local disk.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// Whether a local disk is required during purchase. Valid values:<br><li>REQUIRED: required<br><li>OPTIONAL: optional
	Required *string `json:"Required,omitempty" name:"Required"`
}

type StorageBlock struct {
	// Local HDD storage type. Value: LOCAL_PRO.
	// Note: This field may return null, indicating that no valid value is found.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Minimum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// Maximum capacity of local HDD storage
	// Note: This field may return null, indicating that no valid value is found.
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}