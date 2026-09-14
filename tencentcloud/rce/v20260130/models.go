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

package v20260130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AddPromotionEvent struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The coupon associated with the promotion</p>
	Coupon *Coupon `json:"Coupon,omitnil,omitempty" name:"Coupon"`

	// <p>The point associated with the promotion</p>
	Point *CreditPoint `json:"Point,omitnil,omitempty" name:"Point"`

	// <p>The result of participating the promotion</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131***85678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Address struct {
	// <p>Country</p><p>Parameter format: Compliant with the ISO 3166 standard</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>Province</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>City</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>Region</p>
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// <p>Detailed address</p>
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// <p>Postal code</p>
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`
}

type Amount struct {
	// <p>Currency code</p><p>Parameter format: Compliant with the ISO 4217 standard</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// <p>Original amount in currency</p>
	OriginalAmount *float64 `json:"OriginalAmount,omitnil,omitempty" name:"OriginalAmount"`

	// <p>Current exchange rate of base currency converted to USD</p>
	ExchangeRateUSD *float64 `json:"ExchangeRateUSD,omitnil,omitempty" name:"ExchangeRateUSD"`

	// <p>Current exchange rate of base currency converted to CNY</p>
	ExchangeRateCNY *float64 `json:"ExchangeRateCNY,omitnil,omitempty" name:"ExchangeRateCNY"`
}

type App struct {
	// <p>The operating system your application is running on</p>
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`

	// <p>The operating system version  your application is running on</p>
	OSVersion *string `json:"OSVersion,omitnil,omitempty" name:"OSVersion"`

	// <p>The manufacturer of  the device your application is running on</p>
	DeviceManufacturer *string `json:"DeviceManufacturer,omitnil,omitempty" name:"DeviceManufacturer"`

	// <p>The model of the device your application is running on</p>
	DeviceModel *string `json:"DeviceModel,omitnil,omitempty" name:"DeviceModel"`

	// <p>The ID of the device your application is running on</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>The name of your application</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>The version of your application</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>The language of your application</p>
	ClientLanguage *string `json:"ClientLanguage,omitnil,omitempty" name:"ClientLanguage"`
}

// Predefined struct for user
type AssessDeviceRiskPremiumProRequestParams struct {
	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskPremiumProRequest struct {
	*tchttp.BaseRequest
	
	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskPremiumProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskPremiumProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskPremiumProResponseParams struct {
	// <p>The results of AssessDeviceRiskPremiumPro</p>
	Data *AssessDeviceRiskPremiumRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskPremiumProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskPremiumProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskPremiumProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskPremiumRsp struct {
	// <p>Decision information</p>
	Decision *Decision `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>The risk score information of the device</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>The basic information of the device</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>Basic IP environment information</p>
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessDeviceRiskProRequestParams struct {
	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskProRequest struct {
	*tchttp.BaseRequest
	
	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskProResponseParams struct {
	// <p>The results of AssessDeviceRiskPro</p>
	Data *AssessDeviceRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskRsp struct {
	// <p>The risk score information of the device</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>The basic information of the device</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>Basic IP environment information</p>
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessEnvironmentRiskRequestParams struct {
	// <p>User client IP address(IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessEnvironmentRiskRequest struct {
	*tchttp.BaseRequest
	
	// <p>User client IP address(IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessEnvironmentRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessEnvironmentRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessEnvironmentRiskResponseParams struct {
	// <p>The results of AssessEnvironmentRisk</p>
	Data *AssessEnvironmentRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessEnvironmentRiskResponse struct {
	*tchttp.BaseResponse
	Response *AssessEnvironmentRiskResponseParams `json:"Response"`
}

func (r *AssessEnvironmentRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessEnvironmentRiskRsp struct {
	// <p>The risk score information of the IP environment</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>The basic information of the IP environment</p>
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessRiskRequestParams struct {
	// <p>Event code. Used to specify the scenario node for business access.</p><p>Standard events under the account protection product include:</p><ul><li> login: Log in<p></p></li> <li>register: Register </li><li>sms: SMS </li><li>logout: Log out </li><li>modify_account: Modify account </li><li>modify_password: Modify password </li><li>security_verification: Security verification</li></ul><p>Standard events under the payment protection product include:</p><ul><li>create_order: Create an order </li><li>transaction: Transaction</li><li>charge_back: Chargeback</li></ul><p>Standard events under the promotion protection product include:</p><ul><li>add_promotion: Participate in promotions</li><li>redeem: Redeem a prize </li><li>withdraw: Withdraw</li><li>cust_event: Custom event, cust_xxx </li><li>scan_code: Scan a code </li><li>lucky_draw: Lucky draw </li><li>task: Complete a task </li><li>invitation: Invitation </li><li>claim_red_packet: Receive a red packet </li><li>browse: Browse</li></ul><p>Custom events can be evaluated for risk based on an agreement with RCE</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>The time when the event occurred</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with the ISO 8601 standard</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>The user's current session ID used to associate with the actions before and after logging in. If UserId is not passed, SessionId is required. If missing, an empty string can be filled.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>Client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Event details. The event information is imported based on the event code you input.</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>The user's account ID in your system</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Email of the user</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>Phone number of the user.</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>The details of the browser. If you've already integrated our device SDK, this field is not required</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>The details of the app, os and device.If you've already integrated our device SDK, this field is not required</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`
}

type AssessRiskRequest struct {
	*tchttp.BaseRequest
	
	// <p>Event code. Used to specify the scenario node for business access.</p><p>Standard events under the account protection product include:</p><ul><li> login: Log in<p></p></li> <li>register: Register </li><li>sms: SMS </li><li>logout: Log out </li><li>modify_account: Modify account </li><li>modify_password: Modify password </li><li>security_verification: Security verification</li></ul><p>Standard events under the payment protection product include:</p><ul><li>create_order: Create an order </li><li>transaction: Transaction</li><li>charge_back: Chargeback</li></ul><p>Standard events under the promotion protection product include:</p><ul><li>add_promotion: Participate in promotions</li><li>redeem: Redeem a prize </li><li>withdraw: Withdraw</li><li>cust_event: Custom event, cust_xxx </li><li>scan_code: Scan a code </li><li>lucky_draw: Lucky draw </li><li>task: Complete a task </li><li>invitation: Invitation </li><li>claim_red_packet: Receive a red packet </li><li>browse: Browse</li></ul><p>Custom events can be evaluated for risk based on an agreement with RCE</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>The time when the event occurred</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with the ISO 8601 standard</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>The user's current session ID used to associate with the actions before and after logging in. If UserId is not passed, SessionId is required. If missing, an empty string can be filled.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Device fingerprint token, obtained after integration of the device fingerprint SDK into your website or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>Client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Event details. The event information is imported based on the event code you input.</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>The user's account ID in your system</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Email of the user</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>Phone number of the user.</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>The details of the browser. If you've already integrated our device SDK, this field is not required</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>The details of the app, os and device.If you've already integrated our device SDK, this field is not required</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`
}

func (r *AssessRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventCode")
	delete(f, "EventTime")
	delete(f, "SessionId")
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	delete(f, "EventDetail")
	delete(f, "UserId")
	delete(f, "UserEmail")
	delete(f, "UserPhone")
	delete(f, "Browser")
	delete(f, "App")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessRiskResponseParams struct {
	// <p>The results of AssessRisk</p>
	Data *AssessRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessRiskResponse struct {
	*tchttp.BaseResponse
	Response *AssessRiskResponseParams `json:"Response"`
}

func (r *AssessRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessRiskRsp struct {
	// <p>Decision information</p>
	Decision *Decision `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>Risk score, a scoring result calculated based on the product services you have enabled</p>
	Score *Score `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>Extended information</p>
	ExtraInfo []*Cust `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`
}

type Billing struct {
	// <p>The billing address associated with this user</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>The phone number associated with the bill</p><p>Parameter format: Complies with the E.164 standard, using the format with "+", region code, and number</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>The email associated with the bill</p>
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// <p>The name of the receiver associated with the bill</p>
	Recipient *string `json:"Recipient,omitnil,omitempty" name:"Recipient"`
}

type BrowseEvent struct {
	// <p>Current page type such as home page, search page</p>
	PageType *string `json:"PageType,omitnil,omitempty" name:"PageType"`

	// <p>Currently page URL</p>
	PageUrl *string `json:"PageUrl,omitnil,omitempty" name:"PageUrl"`

	// <p>Browsing duration</p><p>Measurement unit: ms</p>
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// <p>The type of the content in current page such as ad, video, article</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>The ID of the content in current page</p>
	ContentId *string `json:"ContentId,omitnil,omitempty" name:"ContentId"`

	// <p>Previous page type such as home page, search page</p>
	ReferPageType *string `json:"ReferPageType,omitnil,omitempty" name:"ReferPageType"`

	// <p>Previous page URL</p>
	ReferPageUrl *string `json:"ReferPageUrl,omitnil,omitempty" name:"ReferPageUrl"`

	// <p>The ID of the user as guest</p>
	GuestId *string `json:"GuestId,omitnil,omitempty" name:"GuestId"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Browser struct {
	// <p>The user agent of the browser that interacts with the website</p>
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// <p>The language(s) that the client prefers</p><p>Parameter format: Complies with the ISO 3166 standard</p>
	AcceptLanguage *string `json:"AcceptLanguage,omitnil,omitempty" name:"AcceptLanguage"`

	// <p>The language(s) intended for the audience</p><p>Parameter format: Compliant with ISO 3166 standard</p>
	ContentLanguage *string `json:"ContentLanguage,omitnil,omitempty" name:"ContentLanguage"`
}

type Card struct {
	// <p>Bank identification number.The first six or eight digits of the card number</p><p>Parameter format: Compliant with the ISO 13616-1 standard</p>
	CardBin *string `json:"CardBin,omitnil,omitempty" name:"CardBin"`

	// <p>The last four digits of the card number</p><p>Parameter format: Compliant with ISO 13616-1 standard</p>
	LastFourDigits *string `json:"LastFourDigits,omitnil,omitempty" name:"LastFourDigits"`

	// <p>The country where the card issued</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>The bank that issued card</p>
	Bank *string `json:"Bank,omitnil,omitempty" name:"Bank"`

	// <p>the type of the card</p><p>Enumeration value:</p><ul><li>credit: Credit card</li><li>debit: Debit card</li><li>charge: Charge card</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>The brand of the card</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>The level of the card that the bank defined</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>The full name of the person who hold the card</p>
	HolderName *string `json:"HolderName,omitnil,omitempty" name:"HolderName"`

	// <p>The expiration date of the card</p><p>Parameter format: YYYY-MM-DD.</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type ChargeBackEvent struct {
	// <p>The ID of the transaction</p>
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// <p>The ID(s) of the order associated with the transaction</p>
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>The code of the chargeback defined by the card organization, for example: 10.1, 13.1, 4870, 4871</p>
	ChargeBackCode *string `json:"ChargeBackCode,omitnil,omitempty" name:"ChargeBackCode"`

	// <p>The reason of the chargeback defined by the card organization, for example: non-receipt of goods, fraud</p>
	ChargeBackReason *string `json:"ChargeBackReason,omitnil,omitempty" name:"ChargeBackReason"`

	// <p>The process of the chargeback defined by the card organization</p><p>Enumeration values:</p><ul><li>need_response: Merchant needs to respond</li><li>information_supplied: Merchant has provided information</li><li>chargeback_reversed: Chargeback has been canceled</li><li>chargeback_sustained: Chargeback has been established</li></ul>
	ChargeBackProcess *string `json:"ChargeBackProcess,omitnil,omitempty" name:"ChargeBackProcess"`

	// <p>The amount of the chargeback</p>
	ChargeBackAmount *Amount `json:"ChargeBackAmount,omitnil,omitempty" name:"ChargeBackAmount"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ClaimRedPacketEvent struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The Name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The ID of the red packet</p>
	RedPacketId *string `json:"RedPacketId,omitnil,omitempty" name:"RedPacketId"`

	// <p>The type of red packet, for example, random amount, passcode, standard</p>
	RedPacketType *string `json:"RedPacketType,omitnil,omitempty" name:"RedPacketType"`

	// <p>The amount  in the red packet</p>
	RedPacketAmount *Amount `json:"RedPacketAmount,omitnil,omitempty" name:"RedPacketAmount"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Coupon struct {
	// <p>The unique ID of each coupon</p>
	CouponId *string `json:"CouponId,omitnil,omitempty" name:"CouponId"`

	// <p>The name of the coupon</p>
	CouponName *string `json:"CouponName,omitnil,omitempty" name:"CouponName"`

	// <p>The start time of the coupon</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with ISO 8601.</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>The expiration time of the coupon</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with ISO 8601 standard</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>The percentage rate of the coupon. If discount off is 10%,please send 0.1</p>
	PercentageRate *float64 `json:"PercentageRate,omitnil,omitempty" name:"PercentageRate"`

	// <p>The discount amount of the coupon</p>
	DiscountAmount *Amount `json:"DiscountAmount,omitnil,omitempty" name:"DiscountAmount"`

	// <p>The threshold amount of the coupon</p>
	Threshold *float64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type CreateOrderEvent struct {
	// <p>The ID of the order</p>
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>The amount of the order</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>The detail information of the merchant associated with the order</p>
	Merchant *Merchant `json:"Merchant,omitnil,omitempty" name:"Merchant"`

	// <p>The detail information of the bill associated with the order</p>
	Billing *Billing `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>The detail information of the items in the order</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>The detail information of the delivery associated with the order</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`

	// <p>The promotion(s) associated with the order</p>
	Promotions []*Promotion `json:"Promotions,omitnil,omitempty" name:"Promotions"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type CreditPoint struct {
	// <p>The value of the point</p>
	Point *float64 `json:"Point,omitnil,omitempty" name:"Point"`

	// <p>The type of the point</p>
	PointType *string `json:"PointType,omitnil,omitempty" name:"PointType"`
}

type Cust struct {
	// <p>Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Value</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CustEvent struct {
	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type DataScore struct {
	// <p>Risk level</p>
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>Risk label</p>
	RiskLabels []*RiskLabel `json:"RiskLabels,omitnil,omitempty" name:"RiskLabels"`

	// <p>Comprehensive risk score.</p><p>Value ranges from 1 to 1000.</p><p>The larger the value, the larger the risk.</p>
	RiskScore *int64 `json:"RiskScore,omitnil,omitempty" name:"RiskScore"`
}

type Decision struct {
	// <p>Decision result</p><ul><li>pass: Pass</li><li>review: Review</li><li>reject: Reject</li></ul>
	DecisionResult *string `json:"DecisionResult,omitnil,omitempty" name:"DecisionResult"`

	// <p>Decision action when a strategy is matched. Configurable in the console.</p>
	Disposition *string `json:"Disposition,omitnil,omitempty" name:"Disposition"`
}

type Delivery struct {
	// <p>The method of the delivery</p><ul><li>physical</li><li>electronic</li></ul>
	DeliveryMethod *string `json:"DeliveryMethod,omitnil,omitempty" name:"DeliveryMethod"`

	// <p>The fee of the delivery</p>
	DeliveryAmount *Amount `json:"DeliveryAmount,omitnil,omitempty" name:"DeliveryAmount"`

	// <p>The address of the delivery</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>Phone number of the consignee</p><p>parameter format: format with "+", region code, and number that complies with the E.164 standard</p>
	ConsigneePhone *string `json:"ConsigneePhone,omitnil,omitempty" name:"ConsigneePhone"`

	// <p>Email of the consignee</p>
	ConsigneeEmail *string `json:"ConsigneeEmail,omitnil,omitempty" name:"ConsigneeEmail"`

	// <p>Full name of the consignee</p>
	ConsigneeName *string `json:"ConsigneeName,omitnil,omitempty" name:"ConsigneeName"`

	// <p> Whether is the delivery expedited</p>
	Expedited *bool `json:"Expedited,omitnil,omitempty" name:"Expedited"`

	// <p>The carrier of the delivery, usually a logistics company</p>
	DeliveryCarrier *string `json:"DeliveryCarrier,omitnil,omitempty" name:"DeliveryCarrier"`

	// <p>The number(s) used to track the delivery</p>
	DeliveryTracking *string `json:"DeliveryTracking,omitnil,omitempty" name:"DeliveryTracking"`
}

type Device struct {
	// <p>The unique id of device returned by RCE</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>The version of the application</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>Device brand</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>Client IP address</p>
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// <p>Device model</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Network type of the device</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>The package name of the application</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// <p>Device platform</p><p>Enumeration value:</p><ul><li>2: Android</li><li>3: IOS</li><li>4: H5</li><li>5: WeChat Mini Program</li></ul>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>Device system version</p>
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// <p>The build version of SDK</p>
	SdkBuildVersion *string `json:"SdkBuildVersion,omitnil,omitempty" name:"SdkBuildVersion"`

	// <p>Signature verification token. Please contact us to enable signature verification</p>
	SignToken *string `json:"SignToken,omitnil,omitempty" name:"SignToken"`

	// <p>Token generation timestamp, in milliseconds</p>
	TokenTime *string `json:"TokenTime,omitnil,omitempty" name:"TokenTime"`
}

type DigitalOrder struct {
	// <p>The name of the asset</p>
	DigitalAsset *string `json:"DigitalAsset,omitnil,omitempty" name:"DigitalAsset"`

	// <p>The type of the asset</p><p>Enumeration value:</p><ul><li>coin</li><li>commodity</li><li>crypto</li><li>fiat</li><li>token</li><li>stock</li><li>bond</li></ul>
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// <p>The type of trade being made</p><p>Enumeration value:</p><ul><li>limit: Limit order</li><li>market: Market order</li><li>stop_limit: Stop-limit order</li><li>stop_loss: Stop-loss order</li><li>take_profit: Take-profit order</li><li>take_profit_limit: Take-profit limit order</li></ul>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>The quantity of the digital asset</p>
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type Environment struct {
	// <p>The geographical location of the IP address</p>
	Location *IPLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>The basic IP network information</p>
	Network *IPNetwork `json:"Network,omitnil,omitempty" name:"Network"`
}

type EventDetail struct {
	// <p>Login</p>
	Login *LoginEvent `json:"Login,omitnil,omitempty" name:"Login"`

	// <p>Registration</p>
	Register *RegisterEvent `json:"Register,omitnil,omitempty" name:"Register"`

	// <p>Create an order</p>
	CreateOrder *CreateOrderEvent `json:"CreateOrder,omitnil,omitempty" name:"CreateOrder"`

	// <p>Transaction</p>
	Transaction *TransactionEvent `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// <p>SMS</p>
	Sms *SMSEvent `json:"Sms,omitnil,omitempty" name:"Sms"`

	// <p>Chargeback</p>
	ChargeBack *ChargeBackEvent `json:"ChargeBack,omitnil,omitempty" name:"ChargeBack"`

	// <p>Logout</p>
	Logout *LogoutEvent `json:"Logout,omitnil,omitempty" name:"Logout"`

	// <p>Modify account</p>
	ModifyAccount *ModifyAccountEvent `json:"ModifyAccount,omitnil,omitempty" name:"ModifyAccount"`

	// <p>Modify password</p>
	ModifyPassword *ModifyPasswordEvent `json:"ModifyPassword,omitnil,omitempty" name:"ModifyPassword"`

	// <p>Security verification</p>
	SecurityVerification *SecurityVerificationEvent `json:"SecurityVerification,omitnil,omitempty" name:"SecurityVerification"`

	// <p>Participate in promotion activities</p>
	AddPromotion *AddPromotionEvent `json:"AddPromotion,omitnil,omitempty" name:"AddPromotion"`

	// <p>Redeem a prize</p>
	Redeem *RedeemEvent `json:"Redeem,omitnil,omitempty" name:"Redeem"`

	// <p>Withdrawal</p>
	Withdraw *WithdrawEvent `json:"Withdraw,omitnil,omitempty" name:"Withdraw"`

	// <p>Custom event</p>
	CustEvent *CustEvent `json:"CustEvent,omitnil,omitempty" name:"CustEvent"`

	// <p>Scan the QR code</p>
	ScanCode *ScanCodeEvent `json:"ScanCode,omitnil,omitempty" name:"ScanCode"`

	// <p>Lucky draw</p>
	LuckyDraw *LuckyDrawEvent `json:"LuckyDraw,omitnil,omitempty" name:"LuckyDraw"`

	// <p>Perform a task</p>
	Task *TaskEvent `json:"Task,omitnil,omitempty" name:"Task"`

	// <p>Invitation</p>
	Invitation *InvitationEvent `json:"Invitation,omitnil,omitempty" name:"Invitation"`

	// <p>Receive a red packet</p>
	ClaimRedPacket *ClaimRedPacketEvent `json:"ClaimRedPacket,omitnil,omitempty" name:"ClaimRedPacket"`

	// <p>Browse</p>
	Browse *BrowseEvent `json:"Browse,omitnil,omitempty" name:"Browse"`
}

type IPLocation struct {
	// <p>The country of the IP address</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>The region of the IP address</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>The city of the IP address</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>The district of the IP address</p>
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// <p>The longitude of the IP address</p>
	Longitude *string `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// <p>The latitude of the IP address</p>
	Latitude *string `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// <p>The timezone of the IP address</p>
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// <p>The zip code of the IP address</p>
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`
}

type IPNetwork struct {
	// <p>Internet service provider</p>
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// <p>Autonomous system number</p>
	ASN *string `json:"ASN,omitnil,omitempty" name:"ASN"`

	// <p>IP registration organization name</p>
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// <p>Whether it is a reserved IP address</p>
	IsReserved *bool `json:"IsReserved,omitnil,omitempty" name:"IsReserved"`

	// <p>Whether it is a gateway IP address</p>
	IsGateway *bool `json:"IsGateway,omitnil,omitempty" name:"IsGateway"`

	// <p>Whether it belongs to an anycast network</p>
	IsAnycast *bool `json:"IsAnycast,omitnil,omitempty" name:"IsAnycast"`

	// <p>Whether it is from a mobile network</p>
	IsMobile *bool `json:"IsMobile,omitnil,omitempty" name:"IsMobile"`

	// <p>Whether it is a dynamic IP address</p>
	IsDynamic *bool `json:"IsDynamic,omitnil,omitempty" name:"IsDynamic"`

	// <p>Whether it is a network egress</p>
	IsEgress *bool `json:"IsEgress,omitnil,omitempty" name:"IsEgress"`

	// <p>Whether it is used for domain name resolution</p>
	IsDNS *bool `json:"IsDNS,omitnil,omitempty" name:"IsDNS"`

	// <p>Whether it is an educational institution</p>
	IsEducation *bool `json:"IsEducation,omitnil,omitempty" name:"IsEducation"`

	// <p>Whether it is an organization</p>
	IsInstitution *bool `json:"IsInstitution,omitnil,omitempty" name:"IsInstitution"`

	// <p>Whether it is an enterprise dedicated line</p>
	IsCompany *bool `json:"IsCompany,omitnil,omitempty" name:"IsCompany"`

	// <p>Whether it is a residence broadband connection</p>
	IsResidence *bool `json:"IsResidence,omitnil,omitempty" name:"IsResidence"`

	// <p>Whether it is cloud service</p>
	IsCloudService *bool `json:"IsCloudService,omitnil,omitempty" name:"IsCloudService"`

	// <p>Whether it is infrastructure</p>
	IsInfrastructure *bool `json:"IsInfrastructure,omitnil,omitempty" name:"IsInfrastructure"`

	// <p>Whether it is an mail exchange service</p>
	IsMXServer *bool `json:"IsMXServer,omitnil,omitempty" name:"IsMXServer"`
}

type InvitationEvent struct {
	// <p>The ID of the invitee</p>
	InviteeUserId *string `json:"InviteeUserId,omitnil,omitempty" name:"InviteeUserId"`

	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The phone number of the invitee</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number.</p>
	InviteePhone *string `json:"InviteePhone,omitnil,omitempty" name:"InviteePhone"`

	// <p>The code that the inviter sent to the user</p>
	InvitationCode *string `json:"InvitationCode,omitnil,omitempty" name:"InvitationCode"`

	// <p>The url that the inviter sent to the user</p>
	InvitationUrl *string `json:"InvitationUrl,omitnil,omitempty" name:"InvitationUrl"`

	// <p>The channel that inviter used to invite the user </p>
	InvitationChannel *string `json:"InvitationChannel,omitnil,omitempty" name:"InvitationChannel"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Inviter struct {
	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The phone number of the inviter</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number.</p>
	InviterPhone *string `json:"InviterPhone,omitnil,omitempty" name:"InviterPhone"`

	// <p>The code that the inviter sent to the user</p>
	InviteCode *string `json:"InviteCode,omitnil,omitempty" name:"InviteCode"`

	// <p>The channel that inviter used to invite the user</p>
	InviteChannel *string `json:"InviteChannel,omitnil,omitempty" name:"InviteChannel"`
}

type Item struct {
	// <p>The unique ID of the item</p>
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// <p>The name of the item</p>
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// <p>The category of the item</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>The price of the item</p>
	Price *Amount `json:"Price,omitnil,omitempty" name:"Price"`

	// <p>If the item has a UPC (Universal Product Code), please provide it here.</p>
	UPC *string `json:"UPC,omitnil,omitempty" name:"UPC"`

	// <p>If the item has an EAN (European Article Number), please provide it here.</p>
	EAN *string `json:"EAN,omitnil,omitempty" name:"EAN"`

	// <p>If the item has an SKU (Stock Keeping Unit), please provide it here.</p>
	SKU *string `json:"SKU,omitnil,omitempty" name:"SKU"`

	// <p>If the item has an ISBN (International Standard Book Number), please provide it here.</p>
	ISBN *string `json:"ISBN,omitnil,omitempty" name:"ISBN"`

	// <p>The brand of the item</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>The quantity of the item</p>
	Quantity *int64 `json:"Quantity,omitnil,omitempty" name:"Quantity"`

	// <p>The manufacture of the item</p>
	Manufacturer *string `json:"Manufacturer,omitnil,omitempty" name:"Manufacturer"`

	// <p>The tags of the item in your system</p>
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type LoginEvent struct {
	// <p>Basic user information</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>The user name entered when the user logged in</p>
	UserLoginName *string `json:"UserLoginName,omitnil,omitempty" name:"UserLoginName"`

	// <p>Login result</p>
	LoginResult *Result `json:"LoginResult,omitnil,omitempty" name:"LoginResult"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type LogoutEvent struct {
	// <p>The detail information of the user</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>The user name entered when the user logged in</p>
	UserLoginName *string `json:"UserLoginName,omitnil,omitempty" name:"UserLoginName"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type LuckyDrawEvent struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>Number of lucky draw</p><p>Unit: count</p>
	LuckyDrawCount *int64 `json:"LuckyDrawCount,omitnil,omitempty" name:"LuckyDrawCount"`

	// <p>Type of lucky draw</p>
	LuckyDrawType *string `json:"LuckyDrawType,omitnil,omitempty" name:"LuckyDrawType"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Merchant struct {
	// <p>The ID of the merchant</p>
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// <p>The name of the merchant</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Merchant registration time</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with ISO 8601</p>
	RegisterTime *string `json:"RegisterTime,omitnil,omitempty" name:"RegisterTime"`

	// <p>Merchant category code</p><p>Parameter format: 4-digit No. compliant with ISO 18245</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>The phone number of the merchant</p><p>parameter format: format with "+", region code, and number that complies with the E.164 standard</p>
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// <p>The email of the merchant</p>
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// <p>The url of the merchant shop on the website</p>
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// <p>The address of the merchant</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>The level of the merchant</p>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// <p>The type of the merchant</p><p>Enumeration value:</p><ul><li>person: Person</li><li>company: Company</li></ul>
	BusinessType *string `json:"BusinessType,omitnil,omitempty" name:"BusinessType"`

	// <p>The volume of goods on sale of the merchant</p>
	GoodsQuantity *int64 `json:"GoodsQuantity,omitnil,omitempty" name:"GoodsQuantity"`

	// <p>The historical sales volume of the merchant</p>
	HistoricSalesQuantity *int64 `json:"HistoricSalesQuantity,omitnil,omitempty" name:"HistoricSalesQuantity"`

	// <p>The historical sales amount of the merchant</p>
	HistoricSalesAmount *Amount `json:"HistoricSalesAmount,omitnil,omitempty" name:"HistoricSalesAmount"`
}

type ModifyAccountEvent struct {
	// <p>The detail information of the user</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>The personal information of the account when registered</p>
	Person *Person `json:"Person,omitnil,omitempty" name:"Person"`

	// <p>The billing address the user provided when registered</p>
	BillingAddress *Address `json:"BillingAddress,omitnil,omitempty" name:"BillingAddress"`

	// <p>The delivery address the user provided when registered</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ModifyPasswordEvent struct {
	// <p>The reason why the password was updated</p><p>Enumeration value:</p><ul><li>user_modify: User self-initiated modification</li><li>forgot_password: Forget password</li><li>forced_reset: System forcing reset</li></ul>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Order struct {
	// <p>The ID of the order</p>
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>The amount of the order</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>The detail information of the items in the order</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>The detail information of the delivery associated with the order</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`
}

type PaymentMethod struct {
	// <p>Payment method</p><p>Enumeration value:</p><ul><li>cash</li><li>check</li><li>credit_card</li><li>debit_card</li><li>crypto_currency</li><li>digital_wallet</li><li>gift_card</li><li>points</li><li>in_app_purchase</li><li>electronic_fund_transfer</li><li>financing</li><li>invoice</li><li>prepaid_card</li><li>sepa_credit</li></ul>
	PaymentType *string `json:"PaymentType,omitnil,omitempty" name:"PaymentType"`

	// <p>The channel of the payment</p>
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`

	// <p>The details of the card.Required while PaymentMethod is "credit_card","debit_card"</p>
	Card *Card `json:"Card,omitnil,omitempty" name:"Card"`

	// <p>SEPA direct debit mandate</p><p>Enumeration value:</p><ul><li>true: Yes</li><li>false: No</li></ul>
	SEPADirectDebitMandate *bool `json:"SEPADirectDebitMandate,omitnil,omitempty" name:"SEPADirectDebitMandate"`

	// <p>The details of the digital wallet when involved digital trade</p>
	DigitalWallet *Wallet `json:"DigitalWallet,omitnil,omitempty" name:"DigitalWallet"`
}

type PaymentResult struct {
	// <p>The status of the payment</p><p>Enumeration values: </p><ul><li>success: Success, </li><li>failure: Failure.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>The reason why the payment has been declined. e.g.card_declined</p>
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// <p>Whether the 3DS has been used in the payment,  enumeration value:</p><ul><li>Yes: true</li><li>No: false</li></ul>
	ThreeDomainSecure *bool `json:"ThreeDomainSecure,omitnil,omitempty" name:"ThreeDomainSecure"`

	// <p>The ECI code returned when 3DS used</p>
	ECICode *string `json:"ECICode,omitnil,omitempty" name:"ECICode"`

	// <p>Response code from the AVS used for address verification</p>
	AVSCode *string `json:"AVSCode,omitnil,omitempty" name:"AVSCode"`

	// <p>Response code from the CVC used for payment authenticity</p>
	CVCCode *string `json:"CVCCode,omitnil,omitempty" name:"CVCCode"`
}

type Person struct {
	// <p>The full name of the user if provided</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>The gender of the user if provided</p>
	Gender *string `json:"Gender,omitnil,omitempty" name:"Gender"`

	// <p>The birthday of the user if provided</p><p>Parameter format: YYYY-MM-DD.</p>
	Birthday *string `json:"Birthday,omitnil,omitempty" name:"Birthday"`

	// <p>The degree of the user if provided</p>
	Degree *string `json:"Degree,omitnil,omitempty" name:"Degree"`

	// <p>The occupation of the user if provided</p>
	Occupation *string `json:"Occupation,omitnil,omitempty" name:"Occupation"`
}

type Promotion struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The coupon(s) associated with the promotion</p>
	Coupon *Coupon `json:"Coupon,omitnil,omitempty" name:"Coupon"`

	// <p>The point(s) associated with the promotion</p>
	CreditPoint *CreditPoint `json:"CreditPoint,omitnil,omitempty" name:"CreditPoint"`
}

type PromotionCode struct {
	// <p>The ID of the promotion code</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>The type of the promotion code, for example: qrcode, barcode, miniprogram code</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>The url or hyperlink to the image</p>
	ImageLink *string `json:"ImageLink,omitnil,omitempty" name:"ImageLink"`

	// <p>The address where the promotion code worked</p>
	Address *Address `json:"Address,omitnil,omitempty" name:"Address"`

	// <p>The item(s) associated with the promotion code</p>
	Items []*Item `json:"Items,omitnil,omitempty" name:"Items"`
}

type RedeemEvent struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>Order information associated with the redemption</p>
	Order *Order `json:"Order,omitnil,omitempty" name:"Order"`

	// <p>The result of redemption</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type RegisterEvent struct {
	// <p>The result of the register</p>
	RegisterResult *Result `json:"RegisterResult,omitnil,omitempty" name:"RegisterResult"`

	// <p>The detail information of the user</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>The personal information of the account when registered</p>
	Person *Person `json:"Person,omitnil,omitempty" name:"Person"`

	// <p>The billing address the user provided when registered</p>
	BillingAddress *Address `json:"BillingAddress,omitnil,omitempty" name:"BillingAddress"`

	// <p>The delivery address the user provided when registered</p>
	DeliveryAddress *Address `json:"DeliveryAddress,omitnil,omitempty" name:"DeliveryAddress"`

	// <p>The detail information of the inviter who invited the user to your business</p>
	Inviter *Inviter `json:"Inviter,omitnil,omitempty" name:"Inviter"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

// Predefined struct for user
type ReportEventRequestParams struct {
	// <p>Event code. Used to specify the scenario node for business access.</p><p> Standard events under the account protection product include:</p><ul><li> login: Log in<p></p></li> <li>register: Register </li><li>sms: SMS </li><li>logout: Log out </li><li>modify_account: Modify account </li><li>modify_password: Modify password </li><li>security_verification: Security verification</li></ul><p>Standard events under the payment protection product include:</p><ul><li>create_order: Create an order</li><li>transaction: Transaction</li><li>charge_back: Chargeback</li></ul><p>Standard events under the promotion protection product include:</p><ul><li>add_promotion: Participate in promotions </li><li>redeem: Redeem a prize </li><li>withdraw: Withdraw</li><li>cust_event: Custom event, cust_xxx </li><li>scan_code: Scan a code </li><li>lucky_draw: Lucky draw </li><li>task: Complete a task </li><li>invitation: Invitation </li><li>claim_red_packet: Receive a red packet </li><li>browse: Browse</li></ul><p>Custom events can be evaluated for risk based on an agreement with RCE</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>The time when the event occurred</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with the ISO 8601 standard</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>The user's current session ID used to associate with the actions before and after logging in. If UserId is not passed, SessionId is required. If missing, an empty string can be filled.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>The token provided by the SDK integrated in your web site or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Event details, import corresponding event information based on the event code you input</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>The user's account ID in your system</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Email of the user</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>Phone number of the user</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>The details of the browser. If you've already integrated our device SDK, this field is not required</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>The details of the app, os and device.If you've already integrated our device SDK, this field is not required</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`
}

type ReportEventRequest struct {
	*tchttp.BaseRequest
	
	// <p>Event code. Used to specify the scenario node for business access.</p><p> Standard events under the account protection product include:</p><ul><li> login: Log in<p></p></li> <li>register: Register </li><li>sms: SMS </li><li>logout: Log out </li><li>modify_account: Modify account </li><li>modify_password: Modify password </li><li>security_verification: Security verification</li></ul><p>Standard events under the payment protection product include:</p><ul><li>create_order: Create an order</li><li>transaction: Transaction</li><li>charge_back: Chargeback</li></ul><p>Standard events under the promotion protection product include:</p><ul><li>add_promotion: Participate in promotions </li><li>redeem: Redeem a prize </li><li>withdraw: Withdraw</li><li>cust_event: Custom event, cust_xxx </li><li>scan_code: Scan a code </li><li>lucky_draw: Lucky draw </li><li>task: Complete a task </li><li>invitation: Invitation </li><li>claim_red_packet: Receive a red packet </li><li>browse: Browse</li></ul><p>Custom events can be evaluated for risk based on an agreement with RCE</p>
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// <p>The time when the event occurred</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with the ISO 8601 standard</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>The user's current session ID used to associate with the actions before and after logging in. If UserId is not passed, SessionId is required. If missing, an empty string can be filled.</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>The token provided by the SDK integrated in your web site or application</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>User client IP address (IPv4 or IPv6)</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`

	// <p>Event details, import corresponding event information based on the event code you input</p>
	EventDetail *EventDetail `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`

	// <p>The user's account ID in your system</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Email of the user</p>
	UserEmail *string `json:"UserEmail,omitnil,omitempty" name:"UserEmail"`

	// <p>Phone number of the user</p><p>Parameter format: Complies with the E.164 standard format, which includes "+", region code, and number</p>
	UserPhone *string `json:"UserPhone,omitnil,omitempty" name:"UserPhone"`

	// <p>The details of the browser. If you've already integrated our device SDK, this field is not required</p>
	Browser *Browser `json:"Browser,omitnil,omitempty" name:"Browser"`

	// <p>The details of the app, os and device.If you've already integrated our device SDK, this field is not required</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`
}

func (r *ReportEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventCode")
	delete(f, "EventTime")
	delete(f, "SessionId")
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	delete(f, "EventDetail")
	delete(f, "UserId")
	delete(f, "UserEmail")
	delete(f, "UserPhone")
	delete(f, "Browser")
	delete(f, "App")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportEventResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportEventResponse struct {
	*tchttp.BaseResponse
	Response *ReportEventResponseParams `json:"Response"`
}

func (r *ReportEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Result struct {
	// <p>Actual completion status</p><p>Enumeration values:</p><ul><li>success: Success,</li><li>failure: Failure.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Failure reason</p>
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

type RiskLabel struct {
	// <p>The ID of the label</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>The reason of the label</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type SMSEvent struct {
	// <p>The detail information of the user</p>
	UserInfo *User `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// <p>The unique ID of the sms</p>
	SMSId *string `json:"SMSId,omitnil,omitempty" name:"SMSId"`

	// <p>The time that the user received the sms</p><p>Parameter format: Millisecond-level time with UTC time zone compliant with ISO 8601 standard</p>
	ReceivedTime *string `json:"ReceivedTime,omitnil,omitempty" name:"ReceivedTime"`

	// <p>The action of the user after receiving the sms</p><ul><li>no_action: No action from the user</li><li>safe: User confirmation of the correct person's action</li><li>compromised: Feedback from real users indicates third-party action</li></ul>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>The result of the sms</p>
	SMSResult *Result `json:"SMSResult,omitnil,omitempty" name:"SMSResult"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type ScanCodeEvent struct {
	// <p>Promotion code information</p>
	PromotionCode *PromotionCode `json:"PromotionCode,omitnil,omitempty" name:"PromotionCode"`

	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type Score struct {
	// <p>Risk score. Range: 1–1000. The higher the score indicates the higher risk</p>
	RiskScore *int64 `json:"RiskScore,omitnil,omitempty" name:"RiskScore"`

	// <p>Risk label</p>
	RiskLabels []*RiskLabel `json:"RiskLabels,omitnil,omitempty" name:"RiskLabels"`
}

type SecurityVerificationEvent struct {
	// <p>The event type being verified</p><p>Enumeration values:</p><ul><li>register</li><li>login</li><li>modify_account</li><li>modify_password</li><li>create_order</li><li>transaction</li><li>modify_order</li><li>withdraw</li><li>add_promotion</li><li>redeem</li></ul>
	VerificationEvent *string `json:"VerificationEvent,omitnil,omitempty" name:"VerificationEvent"`

	// <p>The type of security verification: sms, phone call, email, captcha, shared knowledge, human face, fingerprint, etc</p>
	VerificationType *string `json:"VerificationType,omitnil,omitempty" name:"VerificationType"`

	// <p>The content of the security verifcation.This value should be passed when the verification type is set to sms, phone_call, email captcha or shared_knowledge</p>
	VerificationContent *string `json:"VerificationContent,omitnil,omitempty" name:"VerificationContent"`

	// <p>The result of security verification</p>
	VerificationResult *Result `json:"VerificationResult,omitnil,omitempty" name:"VerificationResult"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type TaskEvent struct {
	// <p>The ID of the promotion</p>
	PromotionId *string `json:"PromotionId,omitnil,omitempty" name:"PromotionId"`

	// <p>The name of the promotion</p>
	PromotionName *string `json:"PromotionName,omitnil,omitempty" name:"PromotionName"`

	// <p>The description of the promotion</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>The ID of the inviter</p>
	InviterUserId *string `json:"InviterUserId,omitnil,omitempty" name:"InviterUserId"`

	// <p>The ID of the task</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>The name of the task</p>
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// <p>Task type, such as daily check-in, ad viewing, or step accumulation</p>
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// <p>Task completed duration</p><p>Measurement unit: ms</p>
	TaskCostTime *int64 `json:"TaskCostTime,omitnil,omitempty" name:"TaskCostTime"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type TransactionEvent struct {
	// <p>The unique ID of the transaction</p>
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// <p>The ID(s) of the order associated with the transaction</p>
	OrderId []*string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// <p>The amount of the transaction</p>
	PaymentAmount *Amount `json:"PaymentAmount,omitnil,omitempty" name:"PaymentAmount"`

	// <p>The detail information of the payment method associated with the transaction</p>
	PaymentMethod *PaymentMethod `json:"PaymentMethod,omitnil,omitempty" name:"PaymentMethod"`

	// <p>Transaction type</p><p>Enumeration value:</p><ul><li>sale: One-time authorization and deduction (most common)</li><li>authorize: Authorization only (frozen amount)</li><li>capture: Execute deduction (after authorization)</li><li>void: Cancel pending authorization or deduction</li><li>refund: Refund (part or all)</li><li>deposit: Deposit to account</li><li>withdrawal: Withdrawal from account</li><li>transfer: Fund transfer between accounts</li><li>buy: Purchase asset (for example, crypto currency)</li><li>sell: Sell asset</li><li>send: Send fund/asset (for example, cross-wallet transfer)</li><li>receive: Receive fund/asset</li></ul><p>Default value: sale</p>
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// <p>Bill information</p>
	Billing *Billing `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>Delivery information</p>
	Delivery *Delivery `json:"Delivery,omitnil,omitempty" name:"Delivery"`

	// <p>Merchant information</p>
	Merchant *Merchant `json:"Merchant,omitnil,omitempty" name:"Merchant"`

	// <p>Payment result</p>
	PaymentResult *PaymentResult `json:"PaymentResult,omitnil,omitempty" name:"PaymentResult"`

	// <p>The ID of the recipent in transfer transaction</p>
	TransferRecipientUserId *string `json:"TransferRecipientUserId,omitnil,omitempty" name:"TransferRecipientUserId"`

	// <p>The address of the sender in transfer transaction</p>
	TransferSentAddress *Address `json:"TransferSentAddress,omitnil,omitempty" name:"TransferSentAddress"`

	// <p>Physical address of the recipient, applicable to the transfer transaction type.</p>
	TransferReceivedAddress *Address `json:"TransferReceivedAddress,omitnil,omitempty" name:"TransferReceivedAddress"`

	// <p>The digital order(s) associated with the transaction</p>
	DigitalOrders []*DigitalOrder `json:"DigitalOrders,omitnil,omitempty" name:"DigitalOrders"`

	// <p>Wallet to receive crypto currency</p>
	ReceiverWallet *Wallet `json:"ReceiverWallet,omitnil,omitempty" name:"ReceiverWallet"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}

type User struct {
	// <p>The level of the user in your system</p>
	UserLevel *string `json:"UserLevel,omitnil,omitempty" name:"UserLevel"`

	// <p>The point of the user in your system</p>
	UserPoint *CreditPoint `json:"UserPoint,omitnil,omitempty" name:"UserPoint"`

	// <p>The type of the user in your system</p>
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type Wallet struct {
	// <p>Wallet type</p><p>Enumeration value:</p><ul><li>crypto: Crypto currency</li><li>digital: Digital currency</li><li>fiat: Fiat currency</li></ul>
	WalletType *string `json:"WalletType,omitnil,omitempty" name:"WalletType"`

	// <p>The address of the wallet.Usually it is the ID of the wallet.</p>
	WalletAddress *string `json:"WalletAddress,omitnil,omitempty" name:"WalletAddress"`

	// <p>The full name of the person who holds  the wallet</p>
	WalletHolderName *string `json:"WalletHolderName,omitnil,omitempty" name:"WalletHolderName"`

	// <p>The provider of the wallet, such as wechat, alipay, paypal</p>
	WalletProvider *string `json:"WalletProvider,omitnil,omitempty" name:"WalletProvider"`
}

type WithdrawEvent struct {
	// <p>The amount of the withdraw</p>
	Amount *Amount `json:"Amount,omitnil,omitempty" name:"Amount"`

	// <p>The method of the withdraw</p><p>Enumeration value:</p><ul><li>card: bank card</li><li>wallet: digital wallet</li></ul>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>The detail information of the card withdrawn to.Required while the withdraw method is card</p>
	Card *Card `json:"Card,omitnil,omitempty" name:"Card"`

	// <p>The detail information of the wallet withdrawn to.Required while the withdraw method is wallet</p>
	Wallet *Wallet `json:"Wallet,omitnil,omitempty" name:"Wallet"`

	// <p>Withdraw result</p>
	Result *Result `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>The custom parameters agreed with RCE. An array of objects in K:V format. e.g.[{"Key": "ApproverName", "Value": "bob"},{"Key":"ApproverPhone","Value": "+86131****5678"}]</p>
	Cust []*Cust `json:"Cust,omitnil,omitempty" name:"Cust"`
}