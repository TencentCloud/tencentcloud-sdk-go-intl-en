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

package v20181225

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-12-25"

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


func NewModifySendChannelOnMsgTypesRequest() (request *ModifySendChannelOnMsgTypesRequest) {
    request = &ModifySendChannelOnMsgTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("message", APIVersion, "ModifySendChannelOnMsgTypes")
    
    
    return
}

func NewModifySendChannelOnMsgTypesResponse() (response *ModifySendChannelOnMsgTypesResponse) {
    response = &ModifySendChannelOnMsgTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySendChannelOnMsgTypes
// This API is used to batch modify delivery methods.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ = "InvalidParameter."
//  INVALIDPARAMETER_MSGTYPENOTEXIST = "InvalidParameter.MsgTypeNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
func (c *Client) ModifySendChannelOnMsgTypes(request *ModifySendChannelOnMsgTypesRequest) (response *ModifySendChannelOnMsgTypesResponse, err error) {
    return c.ModifySendChannelOnMsgTypesWithContext(context.Background(), request)
}

// ModifySendChannelOnMsgTypes
// This API is used to batch modify delivery methods.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ = "InvalidParameter."
//  INVALIDPARAMETER_MSGTYPENOTEXIST = "InvalidParameter.MsgTypeNotExist"
//  INVALIDPARAMETER_USERNOTEXIST = "InvalidParameter.UserNotExist"
func (c *Client) ModifySendChannelOnMsgTypesWithContext(ctx context.Context, request *ModifySendChannelOnMsgTypesRequest) (response *ModifySendChannelOnMsgTypesResponse, err error) {
    if request == nil {
        request = NewModifySendChannelOnMsgTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "message", APIVersion, "ModifySendChannelOnMsgTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySendChannelOnMsgTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySendChannelOnMsgTypesResponse()
    err = c.Send(request, response)
    return
}
