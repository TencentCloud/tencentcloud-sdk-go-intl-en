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

package v20201002

import (
    "encoding/json"
    "errors"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type Attachment struct {

	// Attachment name, which cannot exceed 255 characters. Some attachment types are not supported. For details, see [Attachment Types](https://intl.cloud.tencent.com/document/product/1288/51951?from_cn_redirect=1).
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Attachment content after base64 encoding. A single attachment cannot exceed 5 MB. Note: Tencent Cloud APIs require that a request packet should not exceed 10 MB. If you are sending multiple attachments, the total size of these attachments cannot exceed 10 MB.
	Content *string `json:"Content,omitempty" name:"Content"`
}

type BlackEmailAddress struct {

	// Time when the email address is blocklisted.
	BounceTime *string `json:"BounceTime,omitempty" name:"BounceTime"`

	// Blocklisted email address.
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type CreateEmailAddressRequest struct {
	*tchttp.BaseRequest

	// Your sender address. (You can create up to 10 sender addresses for each domain.)
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// Sender name.
	EmailSenderName *string `json:"EmailSenderName,omitempty" name:"EmailSenderName"`
}

func (r *CreateEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	delete(f, "EmailSenderName")
	if len(f) > 0 {
		return errors.New("CreateEmailAddressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// Your sender domain. You are advised to use a third-level domain, for example, mail.qcloud.com.
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
}

func (r *CreateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return errors.New("CreateEmailIdentityRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Verification type. The value is fixed to `DOMAIN`.
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// Verification passed or not.
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// DNS information that needs to be configured.
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`
}

func (r *CreateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	if len(f) > 0 {
		return errors.New("CreateEmailTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DNSAttributes struct {

	// Record types: CNAME, A, TXT, and MX.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Domain name.
	SendDomain *string `json:"SendDomain,omitempty" name:"SendDomain"`

	// Expected value.
	ExpectedValue *string `json:"ExpectedValue,omitempty" name:"ExpectedValue"`

	// Currently configured value.
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// Approved or not. The default value is `false`.
	Status *bool `json:"Status,omitempty" name:"Status"`
}

type DeleteBlackListRequest struct {
	*tchttp.BaseRequest

	// List of email addresses to be unblocklisted. Enter at least one address.
	EmailAddressList []*string `json:"EmailAddressList,omitempty" name:"EmailAddressList" list`
}

func (r *DeleteBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddressList")
	if len(f) > 0 {
		return errors.New("DeleteBlackListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBlackListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailAddressRequest struct {
	*tchttp.BaseRequest

	// Sender address.
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

func (r *DeleteEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	if len(f) > 0 {
		return errors.New("DeleteEmailAddressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
}

func (r *DeleteEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return errors.New("DeleteEmailIdentityRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *DeleteEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return errors.New("DeleteEmailTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmailIdentity struct {

	// Sender domain.
	IdentityName *string `json:"IdentityName,omitempty" name:"IdentityName"`

	// Verification type. The value is fixed to `DOMAIN`.
	IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

	// Verification passed or not.
	SendingEnabled *bool `json:"SendingEnabled,omitempty" name:"SendingEnabled"`
}

type EmailSender struct {

	// Sender address.
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// Sender name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EmailSenderName *string `json:"EmailSenderName,omitempty" name:"EmailSenderName"`

	// Creation time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitempty" name:"CreatedTimestamp"`
}

type GetEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
}

func (r *GetEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return errors.New("GetEmailIdentityRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Verification type. The value is fixed to `DOMAIN`.
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// Verification passed or not.
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// DNS configuration details.
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`
}

func (r *GetEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return errors.New("GetEmailTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Template content.
		TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetSendEmailStatusRequest struct {
	*tchttp.BaseRequest

	// Sent date. This parameter is required. You can only query the sending status for a single date at a time.
	RequestDate *string `json:"RequestDate,omitempty" name:"RequestDate"`

	// Offset. Default value: `0`
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of pulled entries. The maximum value is `100`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// `MessageId` field returned by the `SendMail` API
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// Recipient email address
	ToEmailAddress *string `json:"ToEmailAddress,omitempty" name:"ToEmailAddress"`
}

func (r *GetSendEmailStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSendEmailStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MessageId")
	delete(f, "ToEmailAddress")
	if len(f) > 0 {
		return errors.New("GetSendEmailStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetSendEmailStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Email sending status list
		EmailStatusList []*SendEmailStatus `json:"EmailStatusList,omitempty" name:"EmailStatusList" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSendEmailStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSendEmailStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetStatisticsReportRequest struct {
	*tchttp.BaseRequest

	// Start date.
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// End date.
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Sender domain.
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// Recipient address type, for example, gmail.com.
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitempty" name:"ReceivingMailboxType"`
}

func (r *GetStatisticsReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticsReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Domain")
	delete(f, "ReceivingMailboxType")
	if len(f) > 0 {
		return errors.New("GetStatisticsReportRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetStatisticsReportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Daily email sending statistics.
		DailyVolumes []*Volume `json:"DailyVolumes,omitempty" name:"DailyVolumes" list`

		// Overall email sending statistics.
		OverallVolume *Volume `json:"OverallVolume,omitempty" name:"OverallVolume"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStatisticsReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticsReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListBlackEmailAddressRequest struct {
	*tchttp.BaseRequest

	// Start date in the format of `YYYY-MM-DD`
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// End date in the format of `YYYY-MM-DD`
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// Common parameter. It must be used with `Offset`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Common parameter. It must be used with `Limit`. Maximum value of `Limit`: `100`.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// You can specify an email address to query.
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// You can specify a task ID to query.
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *ListBlackEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListBlackEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EmailAddress")
	delete(f, "TaskID")
	if len(f) > 0 {
		return errors.New("ListBlackEmailAddressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListBlackEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of blocklisted addresses.
		BlackList []*BlackEmailAddress `json:"BlackList,omitempty" name:"BlackList" list`

		// Total number of blocklisted addresses.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListBlackEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListBlackEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailAddressRequest struct {
	*tchttp.BaseRequest
}

func (r *ListEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ListEmailAddressRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Details of sender addresses.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		EmailSenders []*EmailSender `json:"EmailSenders,omitempty" name:"EmailSenders" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailIdentitiesRequest struct {
	*tchttp.BaseRequest
}

func (r *ListEmailIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ListEmailIdentitiesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of sender domains.
		EmailIdentities []*EmailIdentity `json:"EmailIdentities,omitempty" name:"EmailIdentities" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEmailIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailTemplatesRequest struct {
	*tchttp.BaseRequest

	// Number of templates to get. This parameter is used for pagination.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Template offset to get. This parameter is used for pagination.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListEmailTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return errors.New("ListEmailTemplatesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListEmailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of email templates.
		TemplatesMetadata []*TemplatesMetadata `json:"TemplatesMetadata,omitempty" name:"TemplatesMetadata" list`

		// Total number of templates.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEmailTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailRequest struct {
	*tchttp.BaseRequest

	// Sender address. Enter a sender address, for example, noreply@mail.qcloud.com. To display the sender name, enter the address in the following format:  
	// sender &lt;email address&gt;. For example: 
	// Tencent Cloud team &lt;noreply@mail.qcloud.com&gt;
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

	// Recipient email addresses. You can send an email to up to 50 recipients at a time. Note: the email content will display all recipient addresses. To send one-to-one emails to several recipients, please call the API multiple times to send the emails.
	Destination []*string `json:"Destination,omitempty" name:"Destination" list`

	// Email subject.
	Subject *string `json:"Subject,omitempty" name:"Subject"`

	// Reply-to address. You can enter a valid personal email address that can receive emails. If this field is left empty, reply emails will be sent to Tencent Cloud.
	ReplyToAddresses *string `json:"ReplyToAddresses,omitempty" name:"ReplyToAddresses"`

	// Template when sending emails using a template.
	Template *Template `json:"Template,omitempty" name:"Template"`

	// Email content when sending emails by calling the API.
	Simple *Simple `json:"Simple,omitempty" name:"Simple"`

	// Email attachments
	Attachments []*Attachment `json:"Attachments,omitempty" name:"Attachments" list`
}

func (r *SendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromEmailAddress")
	delete(f, "Destination")
	delete(f, "Subject")
	delete(f, "ReplyToAddresses")
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	if len(f) > 0 {
		return errors.New("SendEmailRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique ID generated when receiving the message
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailStatus struct {

	// `MessageId` field returned by the `SendEmail` API
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// Recipient email address
	ToEmailAddress *string `json:"ToEmailAddress,omitempty" name:"ToEmailAddress"`

	// Sender email address
	FromEmailAddress *string `json:"FromEmailAddress,omitempty" name:"FromEmailAddress"`

	// Tencent Cloud processing status:
	// 0: successful.
	// 1001: internal system exception.
	// 1002: internal system exception.
	// 1003: internal system exception.
	// 1003: internal system exception.
	// 1004: email sending timeout.
	// 1005: internal system exception.
	// 1006: you have sent too many emails to the same address in a short period.
	// 1007: the email address is in the blocklist.
	// 1009: internal system exception.
	// 1010: daily email sending limit exceeded.
	// 1011: no permission to send custom content. Use a template.
	// 2001: no results found.
	// 3007: invalid template ID or unavailable template.
	// 3008: template status exception.
	// 3009: no permission to use this template.
	// 3010: the format of the `TemplateData` field is incorrect. 
	// 3014: unable to send the email because the sender domain is not verified.
	// 3020: the recipient email address is in the blocklist.
	// 3024: failed to pre-check the email address format.
	// 3030: email sending is restricted temporarily due to high bounce rate.
	// 3033: the account has insufficient balance or overdue payment.
	SendStatus *int64 `json:"SendStatus,omitempty" name:"SendStatus"`

	// Recipient processing status:
	// 0: Tencent Cloud has accepted the request and added it to the send queue.
	// 1: the email is delivered successfully, `DeliverTime` indicates the time when the email is delivered successfully.
	// 2: the email is discarded. `DeliverMessage` indicates the reason for discarding.
	// 3: the recipient's ESP rejects the email, probably because the email address does not exist or due to other reasons.
	// 8: the email is delayed by the ESP. `DeliverMessage` indicates the reason for delay.
	DeliverStatus *int64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`

	// Description of the recipient processing status
	DeliverMessage *string `json:"DeliverMessage,omitempty" name:"DeliverMessage"`

	// Timestamp when the request arrives at Tencent Cloud
	RequestTime *int64 `json:"RequestTime,omitempty" name:"RequestTime"`

	// Timestamp when Tencent Cloud delivers the email
	DeliverTime *int64 `json:"DeliverTime,omitempty" name:"DeliverTime"`

	// Whether the recipient has opened the email
	UserOpened *bool `json:"UserOpened,omitempty" name:"UserOpened"`

	// Whether the recipient has clicked the links in the email
	UserClicked *bool `json:"UserClicked,omitempty" name:"UserClicked"`

	// Whether the recipient has unsubscribed from emails sent by the sender
	UserUnsubscribed *bool `json:"UserUnsubscribed,omitempty" name:"UserUnsubscribed"`

	// Whether the recipient has reported the sender
	UserComplainted *bool `json:"UserComplainted,omitempty" name:"UserComplainted"`
}

type Simple struct {

	// HTML code after base64 encoding. To ensure correct display, this parameter should include all code information and cannot contain external CSS.
	Html *string `json:"Html,omitempty" name:"Html"`

	// Plain text content after base64 encoding. If HTML is not involved, the plain text will be displayed in the email. Otherwise, this parameter represents the plain text style of the email.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type Template struct {

	// Template ID. If you don’t have any template, please create one.
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// Variable parameters in the template. Please use `json.dump` to format the JSON object into a string type. The object is a set of key-value pairs. Each key denotes a variable, which is represented by {{key}}. The key will be replaced with the corresponding value (represented by {{value}}) when sending the email.
	TemplateData *string `json:"TemplateData,omitempty" name:"TemplateData"`
}

type TemplateContent struct {

	// HTML code after base64 encoding.
	Html *string `json:"Html,omitempty" name:"Html"`

	// Text content after base64 encoding.
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TemplatesMetadata struct {

	// Creation time.
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitempty" name:"CreatedTimestamp"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`

	// Template status. 1: under review; 0: approved; 2: rejected; other values: unavailable.
	TemplateStatus *int64 `json:"TemplateStatus,omitempty" name:"TemplateStatus"`

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// Review reply
	ReviewReason *string `json:"ReviewReason,omitempty" name:"ReviewReason"`
}

type UpdateEmailIdentityRequest struct {
	*tchttp.BaseRequest

	// Domain to be verified.
	EmailIdentity *string `json:"EmailIdentity,omitempty" name:"EmailIdentity"`
}

func (r *UpdateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return errors.New("UpdateEmailIdentityRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Verification type. The value is fixed to `DOMAIN`.
		IdentityType *string `json:"IdentityType,omitempty" name:"IdentityType"`

		// Verification passed or not.
		VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitempty" name:"VerifiedForSendingStatus"`

		// DNS information that needs to be configured.
		Attributes []*DNSAttributes `json:"Attributes,omitempty" name:"Attributes" list`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEmailTemplateRequest struct {
	*tchttp.BaseRequest

	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitempty" name:"TemplateContent"`

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitempty" name:"TemplateID"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
}

func (r *UpdateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateContent")
	delete(f, "TemplateID")
	delete(f, "TemplateName")
	if len(f) > 0 {
		return errors.New("UpdateEmailTemplateRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Volume struct {

	// Date
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SendDate *string `json:"SendDate,omitempty" name:"SendDate"`

	// Number of email requests.
	RequestCount *uint64 `json:"RequestCount,omitempty" name:"RequestCount"`

	// Number of email requests accepted by Tencent Cloud.
	AcceptedCount *uint64 `json:"AcceptedCount,omitempty" name:"AcceptedCount"`

	// Number of delivered emails.
	DeliveredCount *uint64 `json:"DeliveredCount,omitempty" name:"DeliveredCount"`

	// Number of users (deduplicated) who opened emails.
	OpenedCount *uint64 `json:"OpenedCount,omitempty" name:"OpenedCount"`

	// Number of recipients who clicked on links in emails.
	ClickedCount *uint64 `json:"ClickedCount,omitempty" name:"ClickedCount"`

	// Number of bounced emails.
	BounceCount *uint64 `json:"BounceCount,omitempty" name:"BounceCount"`

	// Number of users who canceled subscriptions.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnsubscribeCount *uint64 `json:"UnsubscribeCount,omitempty" name:"UnsubscribeCount"`
}
