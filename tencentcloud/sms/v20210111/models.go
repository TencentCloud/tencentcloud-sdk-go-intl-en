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

package v20210111

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type SmsPackagesStatistics struct {

	// The creation time of the package in seconds in the format of UNIX timestamp.
	PackageCreateTime *uint64 `json:"PackageCreateTime,omitempty" name:"PackageCreateTime"`

	// The effective time of the package in seconds in the format of UNIX timestamp.
	PackageEffectiveTime *uint64 `json:"PackageEffectiveTime,omitempty" name:"PackageEffectiveTime"`

	// The expiration time of the package in seconds in the format of UNIX timestamp.
	PackageExpiredTime *uint64 `json:"PackageExpiredTime,omitempty" name:"PackageExpiredTime"`

	// Number of SMS messages in the package
	PackageAmount *uint64 `json:"PackageAmount,omitempty" name:"PackageAmount"`

	// Package type. 0: gifted package; 1: purchased package.
	PackageType *uint64 `json:"PackageType,omitempty" name:"PackageType"`

	// Package ID.
	PackageId *uint64 `json:"PackageId,omitempty" name:"PackageId"`

	// Current number of used messages in the package.
	CurrentUsage *uint64 `json:"CurrentUsage,omitempty" name:"CurrentUsage"`
}

type SmsPackagesStatisticsRequest struct {
	*tchttp.BaseRequest

	// The SMS `SdkAppId` generated after an application is added in the [SMS console](https://console.cloud.tencent.com/smsv2/app-manage), such as 1400006666.
	SmsSdkAppId *string `json:"SmsSdkAppId,omitempty" name:"SmsSdkAppId"`

	// Upper limit (number of packages to be pulled)
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Start time in the format of `yyyymmddhh` accurate to the hour, such as 2021050113 (13:00 on May 1, 2021).
	// Note: the creation time of a pulled package should not be earlier than the start time.
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// End time in the format of `yyyymmddhh` accurate to the hour, such as 2021050118 (18:00 on May 1, 2021).
	// Note: `EndTime` must be later than `BeginTime`. The creation time of a pulled package should not be later than the end time.
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *SmsPackagesStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmsPackagesStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmsSdkAppId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return errors.New("SmsPackagesStatisticsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SmsPackagesStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Delivery statistics response packet body.
		SmsPackagesStatisticsSet []*SmsPackagesStatistics `json:"SmsPackagesStatisticsSet,omitempty" name:"SmsPackagesStatisticsSet" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SmsPackagesStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SmsPackagesStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
