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

package v20200819

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-08-19"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewSendEmailRequest() (request *SendEmailRequest) {
    request = &SendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dms", APIVersion, "SendEmail")
    
    
    return
}

func NewSendEmailResponse() (response *SendEmailResponse) {
    response = &SendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendEmail
// This API is used to send regular emails.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALCOSERROR = "InternalError.InternalCOSError"
//  INTERNALERROR_INTERNALDBERROR = "InternalError.InternalDBError"
//  INTERNALERROR_INTERNALENCRYPTERROR = "InternalError.InternalEncryptError"
//  INVALIDPARAMETER_INVALIDFROMNAMEMALFORMED = "InvalidParameter.InvalidFromNameMalformed"
//  INVALIDPARAMETER_INVALIDHTMLCONTENTMALFORMED = "InvalidParameter.InvalidHtmlContentMalformed"
//  INVALIDPARAMETER_INVALIDMAILADDRESSNAMEMALFORMED = "InvalidParameter.InvalidMailAddressNameMalformed"
//  INVALIDPARAMETER_INVALIDMAILCONTENTMALFORMED = "InvalidParameter.InvalidMailContentMalformed"
//  INVALIDPARAMETER_INVALIDRECEIVERNAMEMALFORMED = "InvalidParameter.InvalidReceiverNameMalformed"
//  INVALIDPARAMETER_INVALIDSUBJECTMALFORMED = "InvalidParameter.InvalidSubjectMalformed"
//  INVALIDPARAMETER_INVALIDTASKNAMEMALFORMED = "InvalidParameter.InvalidTaskNameMalformed"
//  INVALIDPARAMETER_INVALIDTEXTCONTENTMALFORMED = "InvalidParameter.InvalidTextContentMalformed"
//  RESOURCEINUSE_INVALIDTASKNAMEDUPLICATE = "ResourceInUse.InvalidTaskNameDuplicate"
//  RESOURCENOTFOUND_INVALIDMAILADDRESSNOTFOUND = "ResourceNotFound.InvalidMailAddressNotFound"
//  RESOURCENOTFOUND_INVALIDRECEIVERNOTFOUND = "ResourceNotFound.InvalidReceiverNotFound"
//  RESOURCENOTFOUND_INVALIDREPLYNOTFOUND = "ResourceNotFound.InvalidReplyNotFound"
//  RESOURCENOTFOUND_INVALIDTASKNAMENOTFOUND = "ResourceNotFound.InvalidTaskNameNotFound"
func (c *Client) SendEmail(request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

// SendEmail
// This API is used to send regular emails.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALCOSERROR = "InternalError.InternalCOSError"
//  INTERNALERROR_INTERNALDBERROR = "InternalError.InternalDBError"
//  INTERNALERROR_INTERNALENCRYPTERROR = "InternalError.InternalEncryptError"
//  INVALIDPARAMETER_INVALIDFROMNAMEMALFORMED = "InvalidParameter.InvalidFromNameMalformed"
//  INVALIDPARAMETER_INVALIDHTMLCONTENTMALFORMED = "InvalidParameter.InvalidHtmlContentMalformed"
//  INVALIDPARAMETER_INVALIDMAILADDRESSNAMEMALFORMED = "InvalidParameter.InvalidMailAddressNameMalformed"
//  INVALIDPARAMETER_INVALIDMAILCONTENTMALFORMED = "InvalidParameter.InvalidMailContentMalformed"
//  INVALIDPARAMETER_INVALIDRECEIVERNAMEMALFORMED = "InvalidParameter.InvalidReceiverNameMalformed"
//  INVALIDPARAMETER_INVALIDSUBJECTMALFORMED = "InvalidParameter.InvalidSubjectMalformed"
//  INVALIDPARAMETER_INVALIDTASKNAMEMALFORMED = "InvalidParameter.InvalidTaskNameMalformed"
//  INVALIDPARAMETER_INVALIDTEXTCONTENTMALFORMED = "InvalidParameter.InvalidTextContentMalformed"
//  RESOURCEINUSE_INVALIDTASKNAMEDUPLICATE = "ResourceInUse.InvalidTaskNameDuplicate"
//  RESOURCENOTFOUND_INVALIDMAILADDRESSNOTFOUND = "ResourceNotFound.InvalidMailAddressNotFound"
//  RESOURCENOTFOUND_INVALIDRECEIVERNOTFOUND = "ResourceNotFound.InvalidReceiverNotFound"
//  RESOURCENOTFOUND_INVALIDREPLYNOTFOUND = "ResourceNotFound.InvalidReplyNotFound"
//  RESOURCENOTFOUND_INVALIDTASKNAMENOTFOUND = "ResourceNotFound.InvalidTaskNameNotFound"
func (c *Client) SendEmailWithContext(ctx context.Context, request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    request.SetContext(ctx)
    
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewSendTemplatedEmailRequest() (request *SendTemplatedEmailRequest) {
    request = &SendTemplatedEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dms", APIVersion, "SendTemplatedEmail")
    
    
    return
}

func NewSendTemplatedEmailResponse() (response *SendTemplatedEmailResponse) {
    response = &SendTemplatedEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendTemplatedEmail
// This API is used to send template emails.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALCOSERROR = "InternalError.InternalCOSError"
//  INTERNALERROR_INTERNALDBERROR = "InternalError.InternalDBError"
//  INTERNALERROR_INTERNALENCRYPTERROR = "InternalError.InternalEncryptError"
//  INVALIDPARAMETER_INVALIDFROMNAMEMALFORMED = "InvalidParameter.InvalidFromNameMalformed"
//  INVALIDPARAMETER_INVALIDHTMLCONTENTMALFORMED = "InvalidParameter.InvalidHtmlContentMalformed"
//  INVALIDPARAMETER_INVALIDMAILADDRESSNAMEMALFORMED = "InvalidParameter.InvalidMailAddressNameMalformed"
//  INVALIDPARAMETER_INVALIDMAILCONTENTMALFORMED = "InvalidParameter.InvalidMailContentMalformed"
//  INVALIDPARAMETER_INVALIDRECEIVERNAMEMALFORMED = "InvalidParameter.InvalidReceiverNameMalformed"
//  INVALIDPARAMETER_INVALIDSUBJECTMALFORMED = "InvalidParameter.InvalidSubjectMalformed"
//  INVALIDPARAMETER_INVALIDTASKNAMEMALFORMED = "InvalidParameter.InvalidTaskNameMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATECONTENTMALFORMED = "InvalidParameter.InvalidTemplateContentMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATENAMEMALFORMED = "InvalidParameter.InvalidTemplateNameMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATEVALUEMALFORMED = "InvalidParameter.InvalidTemplateValueMalformed"
//  INVALIDPARAMETER_INVALIDTEXTCONTENTMALFORMED = "InvalidParameter.InvalidTextContentMalformed"
//  RESOURCEINUSE_INVALIDTASKNAMEDUPLICATE = "ResourceInUse.InvalidTaskNameDuplicate"
//  RESOURCENOTFOUND_INVALIDMAILADDRESSNOTFOUND = "ResourceNotFound.InvalidMailAddressNotFound"
//  RESOURCENOTFOUND_INVALIDRECEIVERNOTFOUND = "ResourceNotFound.InvalidReceiverNotFound"
//  RESOURCENOTFOUND_INVALIDREPLYNOTFOUND = "ResourceNotFound.InvalidReplyNotFound"
//  RESOURCENOTFOUND_INVALIDTASKNAMENOTFOUND = "ResourceNotFound.InvalidTaskNameNotFound"
//  RESOURCENOTFOUND_INVALIDTEMPLATENOTFOUND = "ResourceNotFound.InvalidTemplateNotFound"
func (c *Client) SendTemplatedEmail(request *SendTemplatedEmailRequest) (response *SendTemplatedEmailResponse, err error) {
    if request == nil {
        request = NewSendTemplatedEmailRequest()
    }
    
    response = NewSendTemplatedEmailResponse()
    err = c.Send(request, response)
    return
}

// SendTemplatedEmail
// This API is used to send template emails.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALCOSERROR = "InternalError.InternalCOSError"
//  INTERNALERROR_INTERNALDBERROR = "InternalError.InternalDBError"
//  INTERNALERROR_INTERNALENCRYPTERROR = "InternalError.InternalEncryptError"
//  INVALIDPARAMETER_INVALIDFROMNAMEMALFORMED = "InvalidParameter.InvalidFromNameMalformed"
//  INVALIDPARAMETER_INVALIDHTMLCONTENTMALFORMED = "InvalidParameter.InvalidHtmlContentMalformed"
//  INVALIDPARAMETER_INVALIDMAILADDRESSNAMEMALFORMED = "InvalidParameter.InvalidMailAddressNameMalformed"
//  INVALIDPARAMETER_INVALIDMAILCONTENTMALFORMED = "InvalidParameter.InvalidMailContentMalformed"
//  INVALIDPARAMETER_INVALIDRECEIVERNAMEMALFORMED = "InvalidParameter.InvalidReceiverNameMalformed"
//  INVALIDPARAMETER_INVALIDSUBJECTMALFORMED = "InvalidParameter.InvalidSubjectMalformed"
//  INVALIDPARAMETER_INVALIDTASKNAMEMALFORMED = "InvalidParameter.InvalidTaskNameMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATECONTENTMALFORMED = "InvalidParameter.InvalidTemplateContentMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATENAMEMALFORMED = "InvalidParameter.InvalidTemplateNameMalformed"
//  INVALIDPARAMETER_INVALIDTEMPLATEVALUEMALFORMED = "InvalidParameter.InvalidTemplateValueMalformed"
//  INVALIDPARAMETER_INVALIDTEXTCONTENTMALFORMED = "InvalidParameter.InvalidTextContentMalformed"
//  RESOURCEINUSE_INVALIDTASKNAMEDUPLICATE = "ResourceInUse.InvalidTaskNameDuplicate"
//  RESOURCENOTFOUND_INVALIDMAILADDRESSNOTFOUND = "ResourceNotFound.InvalidMailAddressNotFound"
//  RESOURCENOTFOUND_INVALIDRECEIVERNOTFOUND = "ResourceNotFound.InvalidReceiverNotFound"
//  RESOURCENOTFOUND_INVALIDREPLYNOTFOUND = "ResourceNotFound.InvalidReplyNotFound"
//  RESOURCENOTFOUND_INVALIDTASKNAMENOTFOUND = "ResourceNotFound.InvalidTaskNameNotFound"
//  RESOURCENOTFOUND_INVALIDTEMPLATENOTFOUND = "ResourceNotFound.InvalidTemplateNotFound"
func (c *Client) SendTemplatedEmailWithContext(ctx context.Context, request *SendTemplatedEmailRequest) (response *SendTemplatedEmailResponse, err error) {
    if request == nil {
        request = NewSendTemplatedEmailRequest()
    }
    request.SetContext(ctx)
    
    response = NewSendTemplatedEmailResponse()
    err = c.Send(request, response)
    return
}
