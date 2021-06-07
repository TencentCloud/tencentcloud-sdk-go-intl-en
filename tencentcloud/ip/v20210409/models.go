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

package v20210409

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type CountryCodeItem struct {

	// Country/region name in English
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitempty" name:"Name"`

	// IOS2 standard country/region code
	IOS2 *string `json:"IOS2,omitempty" name:"IOS2"`

	// IOS3 standard country/region code
	IOS3 *string `json:"IOS3,omitempty" name:"IOS3"`

	// Phone code
	Code *string `json:"Code,omitempty" name:"Code"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// Account type of a new customer. Valid value: `business`.
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and special symbols [!@#$%^&*()]. Spaces are not allowed.
	Password *string `json:"Password,omitempty" name:"Password"`

	// Confirm the password. It must be the same as the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// Country code, which can be obtained via the `GetCountryCodes` API, such as `86`.
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// ISO2 standard country code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field.
	Area *string `json:"Area,omitempty" name:"Area"`

	// Expanded field, which is left empty by default.
	Extended *string `json:"Extended,omitempty" name:"Extended"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateAccountRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Account UIN
		Uin *string `json:"Uin,omitempty" name:"Uin"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetCountryCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("GetCountryCodesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetCountryCodesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of country/region codes
		Data []*CountryCodeItem `json:"Data,omitempty" name:"Data"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCountryCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
