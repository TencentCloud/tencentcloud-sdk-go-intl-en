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

	// Country English Name
	EnName *string `json:"EnName,omitempty" name:"EnName"`

	// Country Chinese Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// IOS2 standard country/region code
	IOS2 *string `json:"IOS2,omitempty" name:"IOS2"`

	// IOS3 standard country/region code
	IOS3 *string `json:"IOS3,omitempty" name:"IOS3"`

	// Phone Code
	Code *string `json:"Code,omitempty" name:"Code"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest

	// The account type identification of the newly created customer. The value of this interface is: business
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// Registered email address. The caller needs to ensure the validity and correctness of the email address.
	// The email format must be met. For example: account@qq.com
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// Account password.
	// Length limit: [8,20].
	// It must also contain numbers, letters and special symbols (!@#$%^&*() and other non-spaces)
	Password *string `json:"Password,omitempty" name:"Password"`

	// Reconfirm the password. It must be the same as the Password value
	ConfirmPassword *string `json:"ConfirmPassword,omitempty" name:"ConfirmPassword"`

	// Customer's mobile phone number. The caller is required to ensure the validity and correctness of the mobile phone number.
	// Length limit: [1,32]. Global mobile phone numbers are supported. For example, 18888888888
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// The country code of the customer. For the value, please refer to the GetCountryCodes interface GetCountryCodes. Such as 86
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// Customer's IOS2 standard country code. Refer to the GetCountryCodes interface for obtaining country codes. It needs to correspond to the CountryCode value. Such as CN
	Area *string `json:"Area,omitempty" name:"Area"`

	// Extension field, default is empty
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

		// The uin of the account
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

		// List of Country Codes
		Data []*CountryCodeItem `json:"Data,omitempty" name:"Data" list`

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
