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

package v20201002

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AddressUnsubscribeConfigData struct {
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Unsubscription link option 0: do not include unsubscription link 1: simplified chinese 2: english 3: traditional chinese 4: spanish 5: french 6: german 7: japanese 8: korean 9: arabic 10: thai.
	UnsubscribeConfig *string `json:"UnsubscribeConfig,omitnil,omitempty" name:"UnsubscribeConfig"`

	// 0: disabled; 1: enabled.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Attachment struct {
	// Attachment name, which cannot exceed 255 characters. Some attachment types are not supported. For details, see [Attachment Types](https://intl.cloud.tencent.com/document/product/1288/51951?from_cn_redirect=1).
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Base64-encoded attachment content. You can send attachments of up to 4 MB in the total size. Note: The TencentCloud API supports a request packet of up to 8 MB in size, and the size of the attachment content will increase by 1.5 times after Base64 encoding. Therefore, you need to keep the total size of all attachments below 4 MB. If the entire request exceeds 8 MB, the API will return an error.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type BatchSendEmailRequestParams struct {
	// Sender address. Enter a sender address such as `noreply@mail.qcloud.com`. To display the sender name, enter the address in the following format:
	// sender &lt;email address&gt;. For example:
	// Tencent Cloud team &lt;noreply@mail.qcloud.com&gt;
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Email subject
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Task type. `1`: immediate; `2`: scheduled; `3`: recurring
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Reply-to address. You can enter a valid personal email address that can receive emails. If this parameter is left empty, reply emails will fail to be sent.
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// Template when emails are sent using a template
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// Disused
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// Attachment parameters to set when you need to send attachments. This parameter is currently unavailable.
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// Parameter required for a recurring sending task
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// Parameter required for a scheduled sending task
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// Unsubscribe link option. `0`: Do not add unsubscribe link; `1`: English `2`: Simplified Chinese; `3`: Traditional Chinese; `4`: Spanish; `5`: French; `6`: German; `7`: Japanese; `8`: Korean; `9`: Arabic; `10`: Thai
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// Whether to add an ad tag. `0`: Add no tag; `1`: Add before the subject; `2`: Add after the subject.
	ADLocation *uint64 `json:"ADLocation,omitnil,omitempty" name:"ADLocation"`
}

type BatchSendEmailRequest struct {
	*tchttp.BaseRequest
	
	// Sender address. Enter a sender address such as `noreply@mail.qcloud.com`. To display the sender name, enter the address in the following format:
	// sender &lt;email address&gt;. For example:
	// Tencent Cloud team &lt;noreply@mail.qcloud.com&gt;
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Email subject
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Task type. `1`: immediate; `2`: scheduled; `3`: recurring
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Reply-to address. You can enter a valid personal email address that can receive emails. If this parameter is left empty, reply emails will fail to be sent.
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	// Template when emails are sent using a template
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// Disused
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// Attachment parameters to set when you need to send attachments. This parameter is currently unavailable.
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// Parameter required for a recurring sending task
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// Parameter required for a scheduled sending task
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// Unsubscribe link option. `0`: Do not add unsubscribe link; `1`: English `2`: Simplified Chinese; `3`: Traditional Chinese; `4`: Spanish; `5`: French; `6`: German; `7`: Japanese; `8`: Korean; `9`: Arabic; `10`: Thai
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// Whether to add an ad tag. `0`: Add no tag; `1`: Add before the subject; `2`: Add after the subject.
	ADLocation *uint64 `json:"ADLocation,omitnil,omitempty" name:"ADLocation"`
}

func (r *BatchSendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSendEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromEmailAddress")
	delete(f, "ReceiverId")
	delete(f, "Subject")
	delete(f, "TaskType")
	delete(f, "ReplyToAddresses")
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	delete(f, "CycleParam")
	delete(f, "TimedParam")
	delete(f, "Unsubscribe")
	delete(f, "ADLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchSendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSendEmailResponseParams struct {
	// Sending task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchSendEmailResponse struct {
	*tchttp.BaseResponse
	Response *BatchSendEmailResponseParams `json:"Response"`
}

func (r *BatchSendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BlackEmailAddress struct {
	// Time when the email address is blocklisted.
	BounceTime *string `json:"BounceTime,omitnil,omitempty" name:"BounceTime"`

	// Blocklisted email address.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

// Predefined struct for user
type CreateAddressUnsubscribeConfigRequestParams struct {
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Unsubscribe link option. 0: Do not add unsubscribe link; 1: English 2: Simplified Chinese; 3: Traditional Chinese; 4: Spanish; 5: French; 6: German; 7: Japanese; 8: Korean; 9: Arabic; 10: Thai
	UnsubscribeConfig *string `json:"UnsubscribeConfig,omitnil,omitempty" name:"UnsubscribeConfig"`

	// 0: disabled; 1: enabled.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type CreateAddressUnsubscribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Unsubscribe link option. 0: Do not add unsubscribe link; 1: English 2: Simplified Chinese; 3: Traditional Chinese; 4: Spanish; 5: French; 6: German; 7: Japanese; 8: Korean; 9: Arabic; 10: Thai
	UnsubscribeConfig *string `json:"UnsubscribeConfig,omitnil,omitempty" name:"UnsubscribeConfig"`

	// 0: disabled; 1: enabled.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *CreateAddressUnsubscribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressUnsubscribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Address")
	delete(f, "UnsubscribeConfig")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressUnsubscribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressUnsubscribeConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAddressUnsubscribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressUnsubscribeConfigResponseParams `json:"Response"`
}

func (r *CreateAddressUnsubscribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressUnsubscribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailAddressRequestParams struct {
	// Your sender address. (You can create up to 10 sender addresses for each domain.)
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// Sender name.
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`
}

type CreateEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// Your sender address. (You can create up to 10 sender addresses for each domain.)
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// Sender name.
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`
}

func (r *CreateEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	delete(f, "EmailSenderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailAddressResponseParams `json:"Response"`
}

func (r *CreateEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailIdentityRequestParams struct {
	// Your sender domain. You are advised to use a third-level domain, for example, mail.qcloud.com.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type CreateEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Your sender domain. You are advised to use a third-level domain, for example, mail.qcloud.com.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *CreateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailIdentityResponseParams struct {
	// Verification type. The value is fixed to `DOMAIN`.
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// Verification passed or not.
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// DNS information that needs to be configured.
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailIdentityResponseParams `json:"Response"`
}

func (r *CreateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailTemplateRequestParams struct {
	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`
}

type CreateEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`
}

func (r *CreateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "TemplateContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmailTemplateResponseParams struct {
	// Template ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmailTemplateResponseParams `json:"Response"`
}

func (r *CreateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailRequestParams struct {
	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Email address
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

type CreateReceiverDetailRequest struct {
	*tchttp.BaseRequest
	
	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Email address
	Emails []*string `json:"Emails,omitnil,omitempty" name:"Emails"`
}

func (r *CreateReceiverDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	delete(f, "Emails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceiverDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverDetailResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceiverDetailResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceiverDetailResponseParams `json:"Response"`
}

func (r *CreateReceiverDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverRequestParams struct {
	// Recipient group name
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// Recipient group description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateReceiverRequest struct {
	*tchttp.BaseRequest
	
	// Recipient group name
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// Recipient group description
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiversName")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReceiverResponseParams struct {
	// Recipient group ID, by which recipient email addresses are uploaded
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReceiverResponse struct {
	*tchttp.BaseResponse
	Response *CreateReceiverResponseParams `json:"Response"`
}

func (r *CreateReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CycleEmailParam struct {
	// Start time of the task
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// Task recurrence in hours
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// Specifies whether to end the cycle. This parameter is used to update the task. Valid values: 0: No; 1: Yes.
	TermCycle *uint64 `json:"TermCycle,omitnil,omitempty" name:"TermCycle"`
}

type DNSAttributes struct {
	// Record types: CNAME, A, TXT, and MX.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Domain name.
	SendDomain *string `json:"SendDomain,omitnil,omitempty" name:"SendDomain"`

	// Expected value.
	ExpectedValue *string `json:"ExpectedValue,omitnil,omitempty" name:"ExpectedValue"`

	// Currently configured value.
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// Approved or not. The default value is `false`.
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteAddressUnsubscribeConfigRequestParams struct {
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type DeleteAddressUnsubscribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

func (r *DeleteAddressUnsubscribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressUnsubscribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Address")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressUnsubscribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressUnsubscribeConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddressUnsubscribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressUnsubscribeConfigResponseParams `json:"Response"`
}

func (r *DeleteAddressUnsubscribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressUnsubscribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlackListRequestParams struct {
	// List of email addresses to be unblocklisted. Enter at least one address.
	EmailAddressList []*string `json:"EmailAddressList,omitnil,omitempty" name:"EmailAddressList"`
}

type DeleteBlackListRequest struct {
	*tchttp.BaseRequest
	
	// List of email addresses to be unblocklisted. Enter at least one address.
	EmailAddressList []*string `json:"EmailAddressList,omitnil,omitempty" name:"EmailAddressList"`
}

func (r *DeleteBlackListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddressList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlackListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlackListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBlackListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlackListResponseParams `json:"Response"`
}

func (r *DeleteBlackListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlackListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailAddressRequestParams struct {
	// Sender address.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

type DeleteEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// Sender address.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

func (r *DeleteEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailAddressResponseParams `json:"Response"`
}

func (r *DeleteEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailIdentityRequestParams struct {
	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type DeleteEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *DeleteEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailIdentityResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailIdentityResponseParams `json:"Response"`
}

func (r *DeleteEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailTemplateRequestParams struct {
	// Template ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

type DeleteEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

func (r *DeleteEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEmailTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEmailTemplateResponseParams `json:"Response"`
}

func (r *DeleteEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReceiverRequestParams struct {
	// Recipient group ID, which is returned when a recipient group is created.
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`
}

type DeleteReceiverRequest struct {
	*tchttp.BaseRequest
	
	// Recipient group ID, which is returned when a recipient group is created.
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`
}

func (r *DeleteReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReceiverId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReceiverResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReceiverResponseParams `json:"Response"`
}

func (r *DeleteReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmailIdentity struct {
	// Sender domain.
	IdentityName *string `json:"IdentityName,omitnil,omitempty" name:"IdentityName"`

	// Verification type. The value is fixed to `DOMAIN`.
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// Verification passed or not.
	SendingEnabled *bool `json:"SendingEnabled,omitnil,omitempty" name:"SendingEnabled"`

	// Current reputation level
	CurrentReputationLevel *uint64 `json:"CurrentReputationLevel,omitnil,omitempty" name:"CurrentReputationLevel"`

	// Maximum number of messages sent per day
	DailyQuota *uint64 `json:"DailyQuota,omitnil,omitempty" name:"DailyQuota"`
}

type EmailSender struct {
	// Sender address.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// Sender name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EmailSenderName *string `json:"EmailSenderName,omitnil,omitempty" name:"EmailSenderName"`

	// Creation time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitnil,omitempty" name:"CreatedTimestamp"`
}

// Predefined struct for user
type GetEmailIdentityRequestParams struct {
	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type GetEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Sender domain.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *GetEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailIdentityResponseParams struct {
	// Verification type. The value is fixed to `DOMAIN`.
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// Verification passed or not.
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// DNS configuration details.
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *GetEmailIdentityResponseParams `json:"Response"`
}

func (r *GetEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailTemplateRequestParams struct {
	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

type GetEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`
}

func (r *GetEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEmailTemplateResponseParams struct {
	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// Template status. Valid values: `0` (approved); `1` (pending approval); `2` (rejected).
	TemplateStatus *uint64 `json:"TemplateStatus,omitnil,omitempty" name:"TemplateStatus"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *GetEmailTemplateResponseParams `json:"Response"`
}

func (r *GetEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSendEmailStatusRequestParams struct {
	// Date sent. This parameter is required. You can only query the sending status for a single date at a time.
	RequestDate *string `json:"RequestDate,omitnil,omitempty" name:"RequestDate"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The `MessageId` field returned by the `SendMail` API.
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// Recipient email address.
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`
}

type GetSendEmailStatusRequest struct {
	*tchttp.BaseRequest
	
	// Date sent. This parameter is required. You can only query the sending status for a single date at a time.
	RequestDate *string `json:"RequestDate,omitnil,omitempty" name:"RequestDate"`

	// Offset. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Maximum number of pulled entries. Maximum value: `100`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The `MessageId` field returned by the `SendMail` API.
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// Recipient email address.
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`
}

func (r *GetSendEmailStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSendEmailStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSendEmailStatusResponseParams struct {
	// Status of sent emails
	EmailStatusList []*SendEmailStatus `json:"EmailStatusList,omitnil,omitempty" name:"EmailStatusList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSendEmailStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetSendEmailStatusResponseParams `json:"Response"`
}

func (r *GetSendEmailStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSendEmailStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticsReportRequestParams struct {
	// Start date.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Sender domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Recipient address type, for example, gmail.com.
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitnil,omitempty" name:"ReceivingMailboxType"`
}

type GetStatisticsReportRequest struct {
	*tchttp.BaseRequest
	
	// Start date.
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date.
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Sender domain.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Recipient address type, for example, gmail.com.
	ReceivingMailboxType *string `json:"ReceivingMailboxType,omitnil,omitempty" name:"ReceivingMailboxType"`
}

func (r *GetStatisticsReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetStatisticsReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetStatisticsReportResponseParams struct {
	// Daily email sending statistics.
	DailyVolumes []*Volume `json:"DailyVolumes,omitnil,omitempty" name:"DailyVolumes"`

	// Overall email sending statistics.
	OverallVolume *Volume `json:"OverallVolume,omitnil,omitempty" name:"OverallVolume"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetStatisticsReportResponse struct {
	*tchttp.BaseResponse
	Response *GetStatisticsReportResponseParams `json:"Response"`
}

func (r *GetStatisticsReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetStatisticsReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAddressUnsubscribeConfigRequestParams struct {
	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// specifies the maximum number of entries to retrieve, with a cap of 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type ListAddressUnsubscribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// specifies the maximum number of entries to retrieve, with a cap of 100.
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *ListAddressUnsubscribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAddressUnsubscribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAddressUnsubscribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAddressUnsubscribeConfigResponseParams struct {
	// Address-Level unsubscribe configuration.
	AddressUnsubscribeConfigList []*AddressUnsubscribeConfigData `json:"AddressUnsubscribeConfigList,omitnil,omitempty" name:"AddressUnsubscribeConfigList"`

	// Total number.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAddressUnsubscribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *ListAddressUnsubscribeConfigResponseParams `json:"Response"`
}

func (r *ListAddressUnsubscribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAddressUnsubscribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListBlackEmailAddressRequestParams struct {
	// Start date in the format of `YYYY-MM-DD`
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date in the format of `YYYY-MM-DD`
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Common parameter. It must be used with `Offset`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Common parameter. It must be used with `Limit`. Maximum value of `Limit`: `100`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// You can specify an email address to query.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// You can specify a task ID to query.
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type ListBlackEmailAddressRequest struct {
	*tchttp.BaseRequest
	
	// Start date in the format of `YYYY-MM-DD`
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// End date in the format of `YYYY-MM-DD`
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// Common parameter. It must be used with `Offset`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Common parameter. It must be used with `Limit`. Maximum value of `Limit`: `100`.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// You can specify an email address to query.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`

	// You can specify a task ID to query.
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *ListBlackEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListBlackEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListBlackEmailAddressResponseParams struct {
	// List of blocklisted addresses.
	BlackList []*BlackEmailAddress `json:"BlackList,omitnil,omitempty" name:"BlackList"`

	// Total number of blocklisted addresses.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListBlackEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *ListBlackEmailAddressResponseParams `json:"Response"`
}

func (r *ListBlackEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListBlackEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailAddressRequestParams struct {

}

type ListEmailAddressRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListEmailAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailAddressResponseParams struct {
	// Details of sender addresses.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EmailSenders []*EmailSender `json:"EmailSenders,omitnil,omitempty" name:"EmailSenders"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailAddressResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailAddressResponseParams `json:"Response"`
}

func (r *ListEmailAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailIdentitiesRequestParams struct {

}

type ListEmailIdentitiesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListEmailIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailIdentitiesResponseParams struct {
	// List of sender domains.
	EmailIdentities []*EmailIdentity `json:"EmailIdentities,omitnil,omitempty" name:"EmailIdentities"`

	// Maximum reputation level
	MaxReputationLevel *uint64 `json:"MaxReputationLevel,omitnil,omitempty" name:"MaxReputationLevel"`

	// Maximum number of emails sent per domain name
	MaxDailyQuota *uint64 `json:"MaxDailyQuota,omitnil,omitempty" name:"MaxDailyQuota"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailIdentitiesResponseParams `json:"Response"`
}

func (r *ListEmailIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailTemplatesRequestParams struct {
	// Number of templates to get. This parameter is used for pagination.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template offset to get. This parameter is used for pagination.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListEmailTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Number of templates to get. This parameter is used for pagination.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Template offset to get. This parameter is used for pagination.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListEmailTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEmailTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEmailTemplatesResponseParams struct {
	// List of email templates.
	TemplatesMetadata []*TemplatesMetadata `json:"TemplatesMetadata,omitnil,omitempty" name:"TemplatesMetadata"`

	// Total number of templates
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEmailTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *ListEmailTemplatesResponseParams `json:"Response"`
}

func (r *ListEmailTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEmailTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiversRequestParams struct {
	// Offset, starting from 0. The value is an integer.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records to query. The value is an integer not exceeding 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Group status (`1`: to be uploaded; `2` uploading; `3` uploaded). To query groups in all states, do not pass in this parameter.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group name keyword for fuzzy query
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

type ListReceiversRequest struct {
	*tchttp.BaseRequest
	
	// Offset, starting from 0. The value is an integer.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records to query. The value is an integer not exceeding 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Group status (`1`: to be uploaded; `2` uploading; `3` uploaded). To query groups in all states, do not pass in this parameter.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Group name keyword for fuzzy query
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`
}

func (r *ListReceiversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "KeyWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReceiversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiversResponseParams struct {
	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Data record
	Data []*ReceiverData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReceiversResponse struct {
	*tchttp.BaseResponse
	Response *ListReceiversResponseParams `json:"Response"`
}

func (r *ListReceiversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSendTasksRequestParams struct {
	// Offset, starting from 0. The value is an integer. `0` means to skip 0 entries.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records to query. The value is an integer not exceeding 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Task status. `1`: to start; `5`: sending; `6`: sending suspended today; `7`: sending error; `10`: sent. To query tasks in all states, do not pass in this parameter.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Task type. `1`: immediate; `2`: scheduled; `3`: recurring. To query tasks of all types, do not pass in this parameter.
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type ListSendTasksRequest struct {
	*tchttp.BaseRequest
	
	// Offset, starting from 0. The value is an integer. `0` means to skip 0 entries.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of records to query. The value is an integer not exceeding 100.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Task status. `1`: to start; `5`: sending; `6`: sending suspended today; `7`: sending error; `10`: sent. To query tasks in all states, do not pass in this parameter.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Task type. `1`: immediate; `2`: scheduled; `3`: recurring. To query tasks of all types, do not pass in this parameter.
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *ListSendTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSendTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "ReceiverId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSendTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSendTasksResponseParams struct {
	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Data record
	Data []*SendTaskData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSendTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListSendTasksResponseParams `json:"Response"`
}

func (r *ListSendTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSendTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiverData struct {
	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Recipient group name
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`

	// Total number of recipient email addresses
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Recipient group description
	// Note: This field may return `null`, indicating that no valid value can be found.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Group status (`1`: to be uploaded; `2` uploading; `3` uploaded)
	// Note: This field may return `null`, indicating that no valid value can be found.
	ReceiversStatus *uint64 `json:"ReceiversStatus,omitnil,omitempty" name:"ReceiversStatus"`

	// Creation time, such as 2021-09-28 16:40:35
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type SendEmailRequestParams struct {
	// Sender address. Enter a sender address, for example, noreply@mail.qcloud.com.
	// To display the sender name, enter the address in the following format: 
	// Sender <email address>
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Recipient email addresses. You can send an email to up to 50 recipients at a time. Note: the email content will display all recipient addresses. To send one-to-one emails to several recipients, please call the API multiple times to send the emails.
	Destination []*string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// Email subject.
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Reply-to address. You can enter a valid personal email address that can receive emails. If this parameter is left empty, reply emails will fail to be sent.
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`


	Cc []*string `json:"Cc,omitnil,omitempty" name:"Cc"`


	Bcc []*string `json:"Bcc,omitnil,omitempty" name:"Bcc"`

	// Template parameters for template-based sending. As `Simple` has been disused, `Template` is required.
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// Disused
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// Parameters for the attachments to be sent. The TencentCloud API supports a request packet of up to 8 MB in size, and the size of the attachment content will increase by 1.5 times after Base64 encoding. Therefore, you need to keep the total size of all attachments below 4 MB. If the entire request exceeds 8 MB, the API will return an error.
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// Unsubscribe link option. `0`: Do not add unsubscribe link; `1`: English `2`: Simplified Chinese; `3`: Traditional Chinese; `4`: Spanish; `5`: French; `6`: German; `7`: Japanese; `8`: Korean; `9`: Arabic; `10`: Thai
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// Email triggering type. `0` (default): non-trigger-based, suitable for marketing emails and non-immediate emails; `1`: trigger-based, suitable for immediate emails such as emails containing verification codes. If the size of an email exceeds a specified value, the system will automatically choose the non-trigger-based type.
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

type SendEmailRequest struct {
	*tchttp.BaseRequest
	
	// Sender address. Enter a sender address, for example, noreply@mail.qcloud.com.
	// To display the sender name, enter the address in the following format: 
	// Sender <email address>
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Recipient email addresses. You can send an email to up to 50 recipients at a time. Note: the email content will display all recipient addresses. To send one-to-one emails to several recipients, please call the API multiple times to send the emails.
	Destination []*string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// Email subject.
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Reply-to address. You can enter a valid personal email address that can receive emails. If this parameter is left empty, reply emails will fail to be sent.
	ReplyToAddresses *string `json:"ReplyToAddresses,omitnil,omitempty" name:"ReplyToAddresses"`

	Cc []*string `json:"Cc,omitnil,omitempty" name:"Cc"`

	Bcc []*string `json:"Bcc,omitnil,omitempty" name:"Bcc"`

	// Template parameters for template-based sending. As `Simple` has been disused, `Template` is required.
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// Disused
	Simple *Simple `json:"Simple,omitnil,omitempty" name:"Simple"`

	// Parameters for the attachments to be sent. The TencentCloud API supports a request packet of up to 8 MB in size, and the size of the attachment content will increase by 1.5 times after Base64 encoding. Therefore, you need to keep the total size of all attachments below 4 MB. If the entire request exceeds 8 MB, the API will return an error.
	Attachments []*Attachment `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// Unsubscribe link option. `0`: Do not add unsubscribe link; `1`: English `2`: Simplified Chinese; `3`: Traditional Chinese; `4`: Spanish; `5`: French; `6`: German; `7`: Japanese; `8`: Korean; `9`: Arabic; `10`: Thai
	Unsubscribe *string `json:"Unsubscribe,omitnil,omitempty" name:"Unsubscribe"`

	// Email triggering type. `0` (default): non-trigger-based, suitable for marketing emails and non-immediate emails; `1`: trigger-based, suitable for immediate emails such as emails containing verification codes. If the size of an email exceeds a specified value, the system will automatically choose the non-trigger-based type.
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`
}

func (r *SendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
	delete(f, "Cc")
	delete(f, "Bcc")
	delete(f, "Template")
	delete(f, "Simple")
	delete(f, "Attachments")
	delete(f, "Unsubscribe")
	delete(f, "TriggerType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendEmailResponseParams struct {
	// Unique ID generated when receiving the message
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *SendEmailResponseParams `json:"Response"`
}

func (r *SendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendEmailStatus struct {
	// The `MessageId` field returned by the `SendEmail` API
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// Recipient email address
	ToEmailAddress *string `json:"ToEmailAddress,omitnil,omitempty" name:"ToEmailAddress"`

	// Sender email address
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Tencent Cloud processing status
	// 0: Successful.
	// 1001: Internal system exception.
	// 1002: Internal system exception.
	// 1003: Internal system exception.
	// 1003: Internal system exception.
	// 1004: Email sending timed out.
	// 1005: Internal system exception.
	// 1006: You have sent too many emails to the same address in a short period.
	// 1007: The email address is in the blocklist.
	// 1008: The sender domain is rejected by the recipient.
	// 1009: Internal system exception.
	// 1010: The daily email sending limit is exceeded.
	// 1011: You have no permission to send custom content. Use a template.
	// 1013: The sender domain is unsubscribed from by the recipient.
	// 2001: No results were found.
	// 3007: The template ID is invalid or the template is unavailable.
	// 3008: The sender domain is temporarily blocked by the recipient domain.
	// 3009: You have no permission to use this template.
	// 3010: The format of the `TemplateData` field is incorrect. 
	// 3014: The email cannot be sent because the sender domain is not verified.
	// 3020: The recipient email address is in the blocklist.
	// 3024: Failed to precheck the email address format.
	// 3030: Email sending is restricted temporarily due to a high bounce rate.
	// 3033: The account has insufficient balance or overdue payment.
	SendStatus *int64 `json:"SendStatus,omitnil,omitempty" name:"SendStatus"`

	// Recipient processing status
	// 0: Tencent Cloud has accepted the request and added it to the send queue.
	// 1: The email is delivered successfully. `DeliverTime` indicates the time when the email is delivered successfully.
	// 2: The email is discarded. `DeliverMessage` indicates the reason for discarding.
	// 3: The recipient's ESP rejects the email, probably because the email address does not exist or due to other reasons.
	// 8: The email is delayed by the ESP. `DeliverMessage` indicates the reason for delay.
	DeliverStatus *int64 `json:"DeliverStatus,omitnil,omitempty" name:"DeliverStatus"`

	// Description of the recipient processing status
	DeliverMessage *string `json:"DeliverMessage,omitnil,omitempty" name:"DeliverMessage"`

	// Timestamp when the request arrives at Tencent Cloud
	RequestTime *int64 `json:"RequestTime,omitnil,omitempty" name:"RequestTime"`

	// Timestamp when Tencent Cloud delivers the email
	DeliverTime *int64 `json:"DeliverTime,omitnil,omitempty" name:"DeliverTime"`

	// Whether the recipient has opened the email
	UserOpened *bool `json:"UserOpened,omitnil,omitempty" name:"UserOpened"`

	// Whether the recipient has clicked the links in the email
	UserClicked *bool `json:"UserClicked,omitnil,omitempty" name:"UserClicked"`

	// Whether the recipient has unsubscribed from the email sent by the sender
	UserUnsubscribed *bool `json:"UserUnsubscribed,omitnil,omitempty" name:"UserUnsubscribed"`

	// Whether the recipient has reported the sender
	UserComplainted *bool `json:"UserComplainted,omitnil,omitempty" name:"UserComplainted"`
}

type SendTaskData struct {
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Sender address
	FromEmailAddress *string `json:"FromEmailAddress,omitnil,omitempty" name:"FromEmailAddress"`

	// Recipient group ID
	ReceiverId *uint64 `json:"ReceiverId,omitnil,omitempty" name:"ReceiverId"`

	// Task status. `1`: to start; `5`: sending; `6`: sending suspended today; `7`: sending error; `10`: sent
	TaskStatus *uint64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// Task type. `1`: immediate; `2`: scheduled; `3`: recurring
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Number of emails requested to be sent
	RequestCount *uint64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// Number of emails sent
	SendCount *uint64 `json:"SendCount,omitnil,omitempty" name:"SendCount"`

	// Number of emails cached
	CacheCount *uint64 `json:"CacheCount,omitnil,omitempty" name:"CacheCount"`

	// Task creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Task update time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Email subject
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Template and template data
	// Note: This field may return `null`, indicating that no valid value can be found.
	Template *Template `json:"Template,omitnil,omitempty" name:"Template"`

	// Parameters of a recurring task
	// Note: This field may return `null`, indicating that no valid value can be found.
	CycleParam *CycleEmailParam `json:"CycleParam,omitnil,omitempty" name:"CycleParam"`

	// Parameters of a scheduled task
	// Note: This field may return `null`, indicating that no valid value can be found.
	TimedParam *TimedEmailParam `json:"TimedParam,omitnil,omitempty" name:"TimedParam"`

	// Task exception information
	// Note: This field may return `null`, indicating that no valid value can be found.
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// Recipient group name
	ReceiversName *string `json:"ReceiversName,omitnil,omitempty" name:"ReceiversName"`
}

type Simple struct {
	// HTML code after base64 encoding. To ensure correct display, this parameter should include all code information and cannot contain external CSS.
	Html *string `json:"Html,omitnil,omitempty" name:"Html"`

	// Plain text content after base64 encoding. If HTML is not involved, the plain text will be displayed in the email. Otherwise, this parameter represents the plain text style of the email.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type Template struct {
	// Template ID. If you don’t have any template, please create one.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// Variable parameters in the template. Please use `json.dump` to format the JSON object into a string type. The object is a set of key-value pairs. Each key denotes a variable, which is represented by {{key}}. The key will be replaced with the corresponding value (represented by {{value}}) when sending the email.
	// Note: The parameter value cannot be data of a complex type such as HTML.
	// Example: {"name":"xxx","age":"xx"}
	TemplateData *string `json:"TemplateData,omitnil,omitempty" name:"TemplateData"`
}

type TemplateContent struct {
	// HTML code after base64 encoding.
	Html *string `json:"Html,omitnil,omitempty" name:"Html"`

	// Text content after base64 encoding.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type TemplatesMetadata struct {
	// Creation time.
	CreatedTimestamp *uint64 `json:"CreatedTimestamp,omitnil,omitempty" name:"CreatedTimestamp"`

	// Template name.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template status. 1: under review; 0: approved; 2: rejected; other values: unavailable.
	TemplateStatus *int64 `json:"TemplateStatus,omitnil,omitempty" name:"TemplateStatus"`

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// Review reply
	ReviewReason *string `json:"ReviewReason,omitnil,omitempty" name:"ReviewReason"`
}

type TimedEmailParam struct {
	// Start time of a scheduled sending task
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`
}

// Predefined struct for user
type UpdateAddressUnsubscribeConfigRequestParams struct {
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Unsubscribe link option. 0: Do not add unsubscribe link; 1: English 2: Simplified Chinese; 3: Traditional Chinese; 4: Spanish; 5: French; 6: German; 7: Japanese; 8: Korean; 9: Arabic; 10: Thai
	UnsubscribeConfig *string `json:"UnsubscribeConfig,omitnil,omitempty" name:"UnsubscribeConfig"`

	// 0: disable; 1: enable.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateAddressUnsubscribeConfigRequest struct {
	*tchttp.BaseRequest
	
	// Sender address.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Unsubscribe link option. 0: Do not add unsubscribe link; 1: English 2: Simplified Chinese; 3: Traditional Chinese; 4: Spanish; 5: French; 6: German; 7: Japanese; 8: Korean; 9: Arabic; 10: Thai
	UnsubscribeConfig *string `json:"UnsubscribeConfig,omitnil,omitempty" name:"UnsubscribeConfig"`

	// 0: disable; 1: enable.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateAddressUnsubscribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAddressUnsubscribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Address")
	delete(f, "UnsubscribeConfig")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAddressUnsubscribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAddressUnsubscribeConfigResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAddressUnsubscribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAddressUnsubscribeConfigResponseParams `json:"Response"`
}

func (r *UpdateAddressUnsubscribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAddressUnsubscribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailIdentityRequestParams struct {
	// Domain to be verified.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

type UpdateEmailIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Domain to be verified.
	EmailIdentity *string `json:"EmailIdentity,omitnil,omitempty" name:"EmailIdentity"`
}

func (r *UpdateEmailIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EmailIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailIdentityResponseParams struct {
	// Verification type. The value is fixed to `DOMAIN`.
	IdentityType *string `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// Verification passed or not.
	VerifiedForSendingStatus *bool `json:"VerifiedForSendingStatus,omitnil,omitempty" name:"VerifiedForSendingStatus"`

	// DNS information that needs to be configured.
	Attributes []*DNSAttributes `json:"Attributes,omitnil,omitempty" name:"Attributes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailIdentityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailIdentityResponseParams `json:"Response"`
}

func (r *UpdateEmailIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailSmtpPassWordRequestParams struct {
	// SMTP password. Length limit: 64.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Email address. Length limit: 128.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

type UpdateEmailSmtpPassWordRequest struct {
	*tchttp.BaseRequest
	
	// SMTP password. Length limit: 64.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Email address. Length limit: 128.
	EmailAddress *string `json:"EmailAddress,omitnil,omitempty" name:"EmailAddress"`
}

func (r *UpdateEmailSmtpPassWordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailSmtpPassWordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Password")
	delete(f, "EmailAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailSmtpPassWordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailSmtpPassWordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailSmtpPassWordResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailSmtpPassWordResponseParams `json:"Response"`
}

func (r *UpdateEmailSmtpPassWordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailSmtpPassWordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailTemplateRequestParams struct {
	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type UpdateEmailTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template content.
	TemplateContent *TemplateContent `json:"TemplateContent,omitnil,omitempty" name:"TemplateContent"`

	// Template ID.
	TemplateID *uint64 `json:"TemplateID,omitnil,omitempty" name:"TemplateID"`

	// Template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

func (r *UpdateEmailTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEmailTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEmailTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEmailTemplateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEmailTemplateResponseParams `json:"Response"`
}

func (r *UpdateEmailTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEmailTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Volume struct {
	// Date
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SendDate *string `json:"SendDate,omitnil,omitempty" name:"SendDate"`

	// Number of email requests.
	RequestCount *uint64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// Number of email requests accepted by Tencent Cloud.
	AcceptedCount *uint64 `json:"AcceptedCount,omitnil,omitempty" name:"AcceptedCount"`

	// Number of delivered emails.
	DeliveredCount *uint64 `json:"DeliveredCount,omitnil,omitempty" name:"DeliveredCount"`

	// Number of users (deduplicated) who opened emails.
	OpenedCount *uint64 `json:"OpenedCount,omitnil,omitempty" name:"OpenedCount"`

	// Number of recipients who clicked on links in emails.
	ClickedCount *uint64 `json:"ClickedCount,omitnil,omitempty" name:"ClickedCount"`

	// Number of bounced emails.
	BounceCount *uint64 `json:"BounceCount,omitnil,omitempty" name:"BounceCount"`

	// Number of users who canceled subscriptions.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	UnsubscribeCount *uint64 `json:"UnsubscribeCount,omitnil,omitempty" name:"UnsubscribeCount"`
}