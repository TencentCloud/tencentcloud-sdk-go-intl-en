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

package v20241204

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2024-12-04"

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


func NewCreateAlarmRequest() (request *CreateAlarmRequest) {
    request = &CreateAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("quota", APIVersion, "CreateAlarm")
    
    
    return
}

func NewCreateAlarmResponse() (response *CreateAlarmResponse) {
    response = &CreateAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarm
// Add alarm rules
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_USERQUOTANOTEXIST = "ResourceNotFound.UserQuotaNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ALARMISEXIST = "UnsupportedOperation.AlarmIsExist"
func (c *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    return c.CreateAlarmWithContext(context.Background(), request)
}

// CreateAlarm
// Add alarm rules
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_USERQUOTANOTEXIST = "ResourceNotFound.UserQuotaNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ALARMISEXIST = "UnsupportedOperation.AlarmIsExist"
func (c *Client) CreateAlarmWithContext(ctx context.Context, request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
    if request == nil {
        request = NewCreateAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "quota", APIVersion, "CreateAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmRequest() (request *DeleteAlarmRequest) {
    request = &DeleteAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("quota", APIVersion, "DeleteAlarm")
    
    
    return
}

func NewDeleteAlarmResponse() (response *DeleteAlarmResponse) {
    response = &DeleteAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarm
// Deletes alarm rules
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    return c.DeleteAlarmWithContext(context.Background(), request)
}

// DeleteAlarm
// Deletes alarm rules
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) DeleteAlarmWithContext(ctx context.Context, request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "quota", APIVersion, "DeleteAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
    request = &DescribeAlarmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("quota", APIVersion, "DescribeAlarms")
    
    
    return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
    response = &DescribeAlarmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarms
// This API is used to query the alarm rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    return c.DescribeAlarmsWithContext(context.Background(), request)
}

// DescribeAlarms
// This API is used to query the alarm rule list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAlarmsWithContext(ctx context.Context, request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "quota", APIVersion, "DescribeAlarms")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarms require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmsResponse()
    err = c.Send(request, response)
    return
}

func NewEnableAlarmRequest() (request *EnableAlarmRequest) {
    request = &EnableAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("quota", APIVersion, "EnableAlarm")
    
    
    return
}

func NewEnableAlarmResponse() (response *EnableAlarmResponse) {
    response = &EnableAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableAlarm
// This API is used to enable alarm rules.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) EnableAlarm(request *EnableAlarmRequest) (response *EnableAlarmResponse, err error) {
    return c.EnableAlarmWithContext(context.Background(), request)
}

// EnableAlarm
// This API is used to enable alarm rules.
//
// error code that may be returned:
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
func (c *Client) EnableAlarmWithContext(ctx context.Context, request *EnableAlarmRequest) (response *EnableAlarmResponse, err error) {
    if request == nil {
        request = NewEnableAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "quota", APIVersion, "EnableAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableAlarmResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlarmRequest() (request *UpdateAlarmRequest) {
    request = &UpdateAlarmRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("quota", APIVersion, "UpdateAlarm")
    
    
    return
}

func NewUpdateAlarmResponse() (response *UpdateAlarmResponse) {
    response = &UpdateAlarmResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlarm
// Modifies alarm rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_USERQUOTANOTEXIST = "ResourceNotFound.UserQuotaNotExist"
//  UNSUPPORTEDOPERATION_ALARMISEXIST = "UnsupportedOperation.AlarmIsExist"
func (c *Client) UpdateAlarm(request *UpdateAlarmRequest) (response *UpdateAlarmResponse, err error) {
    return c.UpdateAlarmWithContext(context.Background(), request)
}

// UpdateAlarm
// Modifies alarm rules
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  RESOURCENOTFOUND_USERQUOTANOTEXIST = "ResourceNotFound.UserQuotaNotExist"
//  UNSUPPORTEDOPERATION_ALARMISEXIST = "UnsupportedOperation.AlarmIsExist"
func (c *Client) UpdateAlarmWithContext(ctx context.Context, request *UpdateAlarmRequest) (response *UpdateAlarmResponse, err error) {
    if request == nil {
        request = NewUpdateAlarmRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "quota", APIVersion, "UpdateAlarm")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlarm require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlarmResponse()
    err = c.Send(request, response)
    return
}
