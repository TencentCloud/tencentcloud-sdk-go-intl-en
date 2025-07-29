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

package v20180724

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

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


func NewBindPrometheusManagedGrafanaRequest() (request *BindPrometheusManagedGrafanaRequest) {
    request = &BindPrometheusManagedGrafanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "BindPrometheusManagedGrafana")
    
    
    return
}

func NewBindPrometheusManagedGrafanaResponse() (response *BindPrometheusManagedGrafanaResponse) {
    response = &BindPrometheusManagedGrafanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindPrometheusManagedGrafana
// This API is used to bind a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) BindPrometheusManagedGrafana(request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    return c.BindPrometheusManagedGrafanaWithContext(context.Background(), request)
}

// BindPrometheusManagedGrafana
// This API is used to bind a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) BindPrometheusManagedGrafanaWithContext(ctx context.Context, request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    if request == nil {
        request = NewBindPrometheusManagedGrafanaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "BindPrometheusManagedGrafana")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindPrometheusManagedGrafana require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindPrometheusManagedGrafanaResponse()
    err = c.Send(request, response)
    return
}

func NewBindingPolicyObjectRequest() (request *BindingPolicyObjectRequest) {
    request = &BindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "BindingPolicyObject")
    
    
    return
}

func NewBindingPolicyObjectResponse() (response *BindingPolicyObjectResponse) {
    response = &BindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindingPolicyObject
// This API is used to bind an alarm policy to a specific object.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObject(request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    return c.BindingPolicyObjectWithContext(context.Background(), request)
}

// BindingPolicyObject
// This API is used to bind an alarm policy to a specific object.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindingPolicyObjectWithContext(ctx context.Context, request *BindingPolicyObjectRequest) (response *BindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewBindingPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "BindingPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindingPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIsPrometheusNewUserRequest() (request *CheckIsPrometheusNewUserRequest) {
    request = &CheckIsPrometheusNewUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CheckIsPrometheusNewUser")
    
    
    return
}

func NewCheckIsPrometheusNewUserResponse() (response *CheckIsPrometheusNewUserResponse) {
    response = &CheckIsPrometheusNewUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckIsPrometheusNewUser
// This API is used to determine whether the user is new to TMP, that is, whether the user has never created a TMP instance in any region.
//
// error code that may be returned:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CheckIsPrometheusNewUser(request *CheckIsPrometheusNewUserRequest) (response *CheckIsPrometheusNewUserResponse, err error) {
    return c.CheckIsPrometheusNewUserWithContext(context.Background(), request)
}

// CheckIsPrometheusNewUser
// This API is used to determine whether the user is new to TMP, that is, whether the user has never created a TMP instance in any region.
//
// error code that may be returned:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CheckIsPrometheusNewUserWithContext(ctx context.Context, request *CheckIsPrometheusNewUserRequest) (response *CheckIsPrometheusNewUserResponse, err error) {
    if request == nil {
        request = NewCheckIsPrometheusNewUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CheckIsPrometheusNewUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIsPrometheusNewUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIsPrometheusNewUserResponse()
    err = c.Send(request, response)
    return
}

func NewCleanGrafanaInstanceRequest() (request *CleanGrafanaInstanceRequest) {
    request = &CleanGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CleanGrafanaInstance")
    
    
    return
}

func NewCleanGrafanaInstanceResponse() (response *CleanGrafanaInstanceResponse) {
    response = &CleanGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CleanGrafanaInstance
// This API is used to forcibly terminate a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CleanGrafanaInstance(request *CleanGrafanaInstanceRequest) (response *CleanGrafanaInstanceResponse, err error) {
    return c.CleanGrafanaInstanceWithContext(context.Background(), request)
}

// CleanGrafanaInstance
// This API is used to forcibly terminate a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CleanGrafanaInstanceWithContext(ctx context.Context, request *CleanGrafanaInstanceRequest) (response *CleanGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewCleanGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CleanGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CleanGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCleanGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
    request = &CreateAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmNotice")
    
    
    return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
    response = &CreateAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmNotice
// This API is used to create a notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    return c.CreateAlarmNoticeWithContext(context.Background(), request)
}

// CreateAlarmNotice
// This API is used to create a notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
    request = &CreateAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlarmPolicy")
    
    
    return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
    response = &CreateAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlarmPolicy
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    return c.CreateAlarmPolicyWithContext(context.Background(), request)
}

// CreateAlarmPolicy
// This API is used to create an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmPolicyWithContext(ctx context.Context, request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAlertRuleRequest() (request *CreateAlertRuleRequest) {
    request = &CreateAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateAlertRule")
    
    
    return
}

func NewCreateAlertRuleResponse() (response *CreateAlertRuleResponse) {
    response = &CreateAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAlertRule
// This API is used to create a Prometheus alerting rule.
//
// 
//
// Note that alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively. For more information, see [Alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAlertRule(request *CreateAlertRuleRequest) (response *CreateAlertRuleResponse, err error) {
    return c.CreateAlertRuleWithContext(context.Background(), request)
}

// CreateAlertRule
// This API is used to create a Prometheus alerting rule.
//
// 
//
// Note that alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively. For more information, see [Alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAlertRuleWithContext(ctx context.Context, request *CreateAlertRuleRequest) (response *CreateAlertRuleResponse, err error) {
    if request == nil {
        request = NewCreateAlertRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateAlertRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExporterIntegrationRequest() (request *CreateExporterIntegrationRequest) {
    request = &CreateExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateExporterIntegration")
    
    
    return
}

func NewCreateExporterIntegrationResponse() (response *CreateExporterIntegrationResponse) {
    response = &CreateExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExporterIntegration
// This API is used to create an exporter integration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
func (c *Client) CreateExporterIntegration(request *CreateExporterIntegrationRequest) (response *CreateExporterIntegrationResponse, err error) {
    return c.CreateExporterIntegrationWithContext(context.Background(), request)
}

// CreateExporterIntegration
// This API is used to create an exporter integration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
func (c *Client) CreateExporterIntegrationWithContext(ctx context.Context, request *CreateExporterIntegrationRequest) (response *CreateExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewCreateExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaInstanceRequest() (request *CreateGrafanaInstanceRequest) {
    request = &CreateGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaInstance")
    
    
    return
}

func NewCreateGrafanaInstanceResponse() (response *CreateGrafanaInstanceResponse) {
    response = &CreateGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaInstance
// This API is used to create a monthly subscribed TCMG instance of the Basic Edition, with auto-renewal enabled and vouchers not allowed by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  FAILEDOPERATION_ZONEUNAVAILABLE = "FailedOperation.ZoneUnavailable"
func (c *Client) CreateGrafanaInstance(request *CreateGrafanaInstanceRequest) (response *CreateGrafanaInstanceResponse, err error) {
    return c.CreateGrafanaInstanceWithContext(context.Background(), request)
}

// CreateGrafanaInstance
// This API is used to create a monthly subscribed TCMG instance of the Basic Edition, with auto-renewal enabled and vouchers not allowed by default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_REGIONUNAVAILABLE = "FailedOperation.RegionUnavailable"
//  FAILEDOPERATION_ZONEUNAVAILABLE = "FailedOperation.ZoneUnavailable"
func (c *Client) CreateGrafanaInstanceWithContext(ctx context.Context, request *CreateGrafanaInstanceRequest) (response *CreateGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaIntegrationRequest() (request *CreateGrafanaIntegrationRequest) {
    request = &CreateGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaIntegration")
    
    
    return
}

func NewCreateGrafanaIntegrationResponse() (response *CreateGrafanaIntegrationResponse) {
    response = &CreateGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaIntegration
// This API is used to create a Grafana integration configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaIntegration(request *CreateGrafanaIntegrationRequest) (response *CreateGrafanaIntegrationResponse, err error) {
    return c.CreateGrafanaIntegrationWithContext(context.Background(), request)
}

// CreateGrafanaIntegration
// This API is used to create a Grafana integration configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaIntegrationWithContext(ctx context.Context, request *CreateGrafanaIntegrationRequest) (response *CreateGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGrafanaNotificationChannelRequest() (request *CreateGrafanaNotificationChannelRequest) {
    request = &CreateGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateGrafanaNotificationChannel")
    
    
    return
}

func NewCreateGrafanaNotificationChannelResponse() (response *CreateGrafanaNotificationChannelResponse) {
    response = &CreateGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGrafanaNotificationChannel
// This API is used to create a Grafana notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaNotificationChannel(request *CreateGrafanaNotificationChannelRequest) (response *CreateGrafanaNotificationChannelResponse, err error) {
    return c.CreateGrafanaNotificationChannelWithContext(context.Background(), request)
}

// CreateGrafanaNotificationChannel
// This API is used to create a Grafana notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateGrafanaNotificationChannelWithContext(ctx context.Context, request *CreateGrafanaNotificationChannelRequest) (response *CreateGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewCreateGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyGroupRequest() (request *CreatePolicyGroupRequest) {
    request = &CreatePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePolicyGroup")
    
    
    return
}

func NewCreatePolicyGroupResponse() (response *CreatePolicyGroupResponse) {
    response = &CreatePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePolicyGroup
// This API is used to add a policy group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    return c.CreatePolicyGroupWithContext(context.Background(), request)
}

// CreatePolicyGroup
// This API is used to add a policy group.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePolicyGroupWithContext(ctx context.Context, request *CreatePolicyGroupRequest) (response *CreatePolicyGroupResponse, err error) {
    if request == nil {
        request = NewCreatePolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAgentRequest() (request *CreatePrometheusAgentRequest) {
    request = &CreatePrometheusAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusAgent")
    
    
    return
}

func NewCreatePrometheusAgentResponse() (response *CreatePrometheusAgentResponse) {
    response = &CreatePrometheusAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAgent
// This API is used to create a Prometheus CVM agent.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAgent(request *CreatePrometheusAgentRequest) (response *CreatePrometheusAgentResponse, err error) {
    return c.CreatePrometheusAgentWithContext(context.Background(), request)
}

// CreatePrometheusAgent
// This API is used to create a Prometheus CVM agent.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAgentWithContext(ctx context.Context, request *CreatePrometheusAgentRequest) (response *CreatePrometheusAgentResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusAlertPolicyRequest() (request *CreatePrometheusAlertPolicyRequest) {
    request = &CreatePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusAlertPolicy")
    
    
    return
}

func NewCreatePrometheusAlertPolicyResponse() (response *CreatePrometheusAlertPolicyResponse) {
    response = &CreatePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusAlertPolicy
// This API is used to create an alerting rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicy(request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    return c.CreatePrometheusAlertPolicyWithContext(context.Background(), request)
}

// CreatePrometheusAlertPolicy
// This API is used to create an alerting rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusAlertPolicyWithContext(ctx context.Context, request *CreatePrometheusAlertPolicyRequest) (response *CreatePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusClusterAgentRequest() (request *CreatePrometheusClusterAgentRequest) {
    request = &CreatePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusClusterAgent")
    
    
    return
}

func NewCreatePrometheusClusterAgentResponse() (response *CreatePrometheusClusterAgentResponse) {
    response = &CreatePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusClusterAgent
// This API is used to associate a cluster with a Cloud Monitor (CM)-integrated Tencent Managed Service for Prometheus (TMP) 2.0 instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgent(request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    return c.CreatePrometheusClusterAgentWithContext(context.Background(), request)
}

// CreatePrometheusClusterAgent
// This API is used to associate a cluster with a Cloud Monitor (CM)-integrated Tencent Managed Service for Prometheus (TMP) 2.0 instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusClusterAgentWithContext(ctx context.Context, request *CreatePrometheusClusterAgentRequest) (response *CreatePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusClusterAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusClusterAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusConfigRequest() (request *CreatePrometheusConfigRequest) {
    request = &CreatePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusConfig")
    
    
    return
}

func NewCreatePrometheusConfigResponse() (response *CreatePrometheusConfigResponse) {
    response = &CreatePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusConfig
// This API is used to create Prometheus configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfig(request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    return c.CreatePrometheusConfigWithContext(context.Background(), request)
}

// CreatePrometheusConfig
// This API is used to create Prometheus configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  RESOURCEINUSE_RESOURCEEXISTALREADY = "ResourceInUse.ResourceExistAlready"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusConfigWithContext(ctx context.Context, request *CreatePrometheusConfigRequest) (response *CreatePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusGlobalNotificationRequest() (request *CreatePrometheusGlobalNotificationRequest) {
    request = &CreatePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusGlobalNotification")
    
    
    return
}

func NewCreatePrometheusGlobalNotificationResponse() (response *CreatePrometheusGlobalNotificationResponse) {
    response = &CreatePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusGlobalNotification
// This API is used to create a global alert notification channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotification(request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    return c.CreatePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// CreatePrometheusGlobalNotification
// This API is used to create a global alert notification channel.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) CreatePrometheusGlobalNotificationWithContext(ctx context.Context, request *CreatePrometheusGlobalNotificationRequest) (response *CreatePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusMultiTenantInstancePostPayModeRequest() (request *CreatePrometheusMultiTenantInstancePostPayModeRequest) {
    request = &CreatePrometheusMultiTenantInstancePostPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusMultiTenantInstancePostPayMode")
    
    
    return
}

func NewCreatePrometheusMultiTenantInstancePostPayModeResponse() (response *CreatePrometheusMultiTenantInstancePostPayModeResponse) {
    response = &CreatePrometheusMultiTenantInstancePostPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusMultiTenantInstancePostPayMode
// This API is used to create a pay-as-you-go Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCE = "FailedOperation.CreateInstance"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePrometheusMultiTenantInstancePostPayMode(request *CreatePrometheusMultiTenantInstancePostPayModeRequest) (response *CreatePrometheusMultiTenantInstancePostPayModeResponse, err error) {
    return c.CreatePrometheusMultiTenantInstancePostPayModeWithContext(context.Background(), request)
}

// CreatePrometheusMultiTenantInstancePostPayMode
// This API is used to create a pay-as-you-go Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_CREATEINSTANCE = "FailedOperation.CreateInstance"
//  FAILEDOPERATION_CREATEINSTANCELIMITED = "FailedOperation.CreateInstanceLimited"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePrometheusMultiTenantInstancePostPayModeWithContext(ctx context.Context, request *CreatePrometheusMultiTenantInstancePostPayModeRequest) (response *CreatePrometheusMultiTenantInstancePostPayModeResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusMultiTenantInstancePostPayModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusMultiTenantInstancePostPayMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusMultiTenantInstancePostPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusMultiTenantInstancePostPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusRecordRuleYamlRequest() (request *CreatePrometheusRecordRuleYamlRequest) {
    request = &CreatePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusRecordRuleYaml")
    
    
    return
}

func NewCreatePrometheusRecordRuleYamlResponse() (response *CreatePrometheusRecordRuleYamlResponse) {
    response = &CreatePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusRecordRuleYaml
// This API is used to create a recording rule in the YAML way.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYaml(request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    return c.CreatePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// CreatePrometheusRecordRuleYaml
// This API is used to create a recording rule in the YAML way.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePrometheusRecordRuleYamlWithContext(ctx context.Context, request *CreatePrometheusRecordRuleYamlRequest) (response *CreatePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusScrapeJobRequest() (request *CreatePrometheusScrapeJobRequest) {
    request = &CreatePrometheusScrapeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusScrapeJob")
    
    
    return
}

func NewCreatePrometheusScrapeJobResponse() (response *CreatePrometheusScrapeJobResponse) {
    response = &CreatePrometheusScrapeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusScrapeJob
// This API is used to create a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusScrapeJob(request *CreatePrometheusScrapeJobRequest) (response *CreatePrometheusScrapeJobResponse, err error) {
    return c.CreatePrometheusScrapeJobWithContext(context.Background(), request)
}

// CreatePrometheusScrapeJob
// This API is used to create a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePrometheusScrapeJobWithContext(ctx context.Context, request *CreatePrometheusScrapeJobRequest) (response *CreatePrometheusScrapeJobResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusScrapeJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusScrapeJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusScrapeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusScrapeJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrometheusTempRequest() (request *CreatePrometheusTempRequest) {
    request = &CreatePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreatePrometheusTemp")
    
    
    return
}

func NewCreatePrometheusTempResponse() (response *CreatePrometheusTempResponse) {
    response = &CreatePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrometheusTemp
// This API is used to create a TMP template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTemp(request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    return c.CreatePrometheusTempWithContext(context.Background(), request)
}

// CreatePrometheusTemp
// This API is used to create a TMP template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) CreatePrometheusTempWithContext(ctx context.Context, request *CreatePrometheusTempRequest) (response *CreatePrometheusTempResponse, err error) {
    if request == nil {
        request = NewCreatePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreatePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordingRuleRequest() (request *CreateRecordingRuleRequest) {
    request = &CreateRecordingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateRecordingRule")
    
    
    return
}

func NewCreateRecordingRuleResponse() (response *CreateRecordingRuleResponse) {
    response = &CreateRecordingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRecordingRule
// This API is used to create a Prometheus recording rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRecordingRule(request *CreateRecordingRuleRequest) (response *CreateRecordingRuleResponse, err error) {
    return c.CreateRecordingRuleWithContext(context.Background(), request)
}

// CreateRecordingRule
// This API is used to create a Prometheus recording rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateRecordingRuleWithContext(ctx context.Context, request *CreateRecordingRuleRequest) (response *CreateRecordingRuleResponse, err error) {
    if request == nil {
        request = NewCreateRecordingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateRecordingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecordingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecordingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSSOAccountRequest() (request *CreateSSOAccountRequest) {
    request = &CreateSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateSSOAccount")
    
    
    return
}

func NewCreateSSOAccountResponse() (response *CreateSSOAccountResponse) {
    response = &CreateSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSSOAccount
// This API is used to authorize a Grafana instance to another Tencent Cloud user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSSOAccount(request *CreateSSOAccountRequest) (response *CreateSSOAccountResponse, err error) {
    return c.CreateSSOAccountWithContext(context.Background(), request)
}

// CreateSSOAccount
// This API is used to authorize a Grafana instance to another Tencent Cloud user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSSOAccountWithContext(ctx context.Context, request *CreateSSOAccountRequest) (response *CreateSSOAccountResponse, err error) {
    if request == nil {
        request = NewCreateSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceDiscoveryRequest() (request *CreateServiceDiscoveryRequest) {
    request = &CreateServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "CreateServiceDiscovery")
    
    
    return
}

func NewCreateServiceDiscoveryResponse() (response *CreateServiceDiscoveryResponse) {
    response = &CreateServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateServiceDiscovery
// This API is used to create a Prometheus scrape configuration in TKE.
//
// <p>Note: The prerequisite is that the corresponding TKE service has been integrated through the Prometheus console. For more information, see
//
// <a href="https://intl.cloud.tencent.com/document/product/248/48859?from_cn_redirect=1" target="_blank">Agent Management</a>.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKECLIENTAUTHFAIL = "FailedOperation.TKEClientAuthFail"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateServiceDiscovery(request *CreateServiceDiscoveryRequest) (response *CreateServiceDiscoveryResponse, err error) {
    return c.CreateServiceDiscoveryWithContext(context.Background(), request)
}

// CreateServiceDiscovery
// This API is used to create a Prometheus scrape configuration in TKE.
//
// <p>Note: The prerequisite is that the corresponding TKE service has been integrated through the Prometheus console. For more information, see
//
// <a href="https://intl.cloud.tencent.com/document/product/248/48859?from_cn_redirect=1" target="_blank">Agent Management</a>.</p>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTED = "FailedOperation.AgentVersionNotSupported"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKECLIENTAUTHFAIL = "FailedOperation.TKEClientAuthFail"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateServiceDiscoveryWithContext(ctx context.Context, request *CreateServiceDiscoveryRequest) (response *CreateServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewCreateServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "CreateServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmNoticesRequest() (request *DeleteAlarmNoticesRequest) {
    request = &DeleteAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmNotices")
    
    
    return
}

func NewDeleteAlarmNoticesResponse() (response *DeleteAlarmNoticesResponse) {
    response = &DeleteAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmNotices
// This API is used to delete an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNotices(request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    return c.DeleteAlarmNoticesWithContext(context.Background(), request)
}

// DeleteAlarmNotices
// This API is used to delete an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNoticesWithContext(ctx context.Context, request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlarmNotices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
    request = &DeleteAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlarmPolicy")
    
    
    return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
    response = &DeleteAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlarmPolicy
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    return c.DeleteAlarmPolicyWithContext(context.Background(), request)
}

// DeleteAlarmPolicy
// This API is used to delete an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAlertRulesRequest() (request *DeleteAlertRulesRequest) {
    request = &DeleteAlertRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteAlertRules")
    
    
    return
}

func NewDeleteAlertRulesResponse() (response *DeleteAlertRulesResponse) {
    response = &DeleteAlertRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAlertRules
// This API is used to batch delete Prometheus alerting rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAlertRules(request *DeleteAlertRulesRequest) (response *DeleteAlertRulesResponse, err error) {
    return c.DeleteAlertRulesWithContext(context.Background(), request)
}

// DeleteAlertRules
// This API is used to batch delete Prometheus alerting rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAlertRulesWithContext(ctx context.Context, request *DeleteAlertRulesRequest) (response *DeleteAlertRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAlertRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteAlertRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAlertRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAlertRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExporterIntegrationRequest() (request *DeleteExporterIntegrationRequest) {
    request = &DeleteExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteExporterIntegration")
    
    
    return
}

func NewDeleteExporterIntegrationResponse() (response *DeleteExporterIntegrationResponse) {
    response = &DeleteExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExporterIntegration
// This API is used to delete an exporter integration.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteExporterIntegration(request *DeleteExporterIntegrationRequest) (response *DeleteExporterIntegrationResponse, err error) {
    return c.DeleteExporterIntegrationWithContext(context.Background(), request)
}

// DeleteExporterIntegration
// This API is used to delete an exporter integration.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteExporterIntegrationWithContext(ctx context.Context, request *DeleteExporterIntegrationRequest) (response *DeleteExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewDeleteExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaInstanceRequest() (request *DeleteGrafanaInstanceRequest) {
    request = &DeleteGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaInstance")
    
    
    return
}

func NewDeleteGrafanaInstanceResponse() (response *DeleteGrafanaInstanceResponse) {
    response = &DeleteGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaInstance
// This API is used to refund a monthly subscribed TCMG instance. Once it is called, the instance cannot be used and will be automatically terminated seven days later.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstance(request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    return c.DeleteGrafanaInstanceWithContext(context.Background(), request)
}

// DeleteGrafanaInstance
// This API is used to refund a monthly subscribed TCMG instance. Once it is called, the instance cannot be used and will be automatically terminated seven days later.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstanceWithContext(ctx context.Context, request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaIntegrationRequest() (request *DeleteGrafanaIntegrationRequest) {
    request = &DeleteGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaIntegration")
    
    
    return
}

func NewDeleteGrafanaIntegrationResponse() (response *DeleteGrafanaIntegrationResponse) {
    response = &DeleteGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaIntegration
// This API is used to delete a Grafana integration configuration.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegration(request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    return c.DeleteGrafanaIntegrationWithContext(context.Background(), request)
}

// DeleteGrafanaIntegration
// This API is used to delete a Grafana integration configuration.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegrationWithContext(ctx context.Context, request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGrafanaNotificationChannelRequest() (request *DeleteGrafanaNotificationChannelRequest) {
    request = &DeleteGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteGrafanaNotificationChannel")
    
    
    return
}

func NewDeleteGrafanaNotificationChannelResponse() (response *DeleteGrafanaNotificationChannelResponse) {
    response = &DeleteGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGrafanaNotificationChannel
// This API is used to delete a Grafana notification channel.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannel(request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    return c.DeleteGrafanaNotificationChannelWithContext(context.Background(), request)
}

// DeleteGrafanaNotificationChannel
// This API is used to delete a Grafana notification channel.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannelWithContext(ctx context.Context, request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyGroupRequest() (request *DeletePolicyGroupRequest) {
    request = &DeletePolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePolicyGroup")
    
    
    return
}

func NewDeletePolicyGroupResponse() (response *DeletePolicyGroupResponse) {
    response = &DeletePolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePolicyGroup
// This API is used to delete an alarm policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroup(request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    return c.DeletePolicyGroupWithContext(context.Background(), request)
}

// DeletePolicyGroup
// This API is used to delete an alarm policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeletePolicyGroupWithContext(ctx context.Context, request *DeletePolicyGroupRequest) (response *DeletePolicyGroupResponse, err error) {
    if request == nil {
        request = NewDeletePolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusAlertPolicyRequest() (request *DeletePrometheusAlertPolicyRequest) {
    request = &DeletePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusAlertPolicy")
    
    
    return
}

func NewDeletePrometheusAlertPolicyResponse() (response *DeletePrometheusAlertPolicyResponse) {
    response = &DeletePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusAlertPolicy
// This API is used to delete a TMP 2.0 instance alerting rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicy(request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    return c.DeletePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DeletePrometheusAlertPolicy
// This API is used to delete a TMP 2.0 instance alerting rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusAlertPolicyWithContext(ctx context.Context, request *DeletePrometheusAlertPolicyRequest) (response *DeletePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusClusterAgentRequest() (request *DeletePrometheusClusterAgentRequest) {
    request = &DeletePrometheusClusterAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusClusterAgent")
    
    
    return
}

func NewDeletePrometheusClusterAgentResponse() (response *DeletePrometheusClusterAgentResponse) {
    response = &DeletePrometheusClusterAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusClusterAgent
// This API is used to disassociate a TMP instance from a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgent(request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    return c.DeletePrometheusClusterAgentWithContext(context.Background(), request)
}

// DeletePrometheusClusterAgent
// This API is used to disassociate a TMP instance from a cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusClusterAgentWithContext(ctx context.Context, request *DeletePrometheusClusterAgentRequest) (response *DeletePrometheusClusterAgentResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusClusterAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusClusterAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusClusterAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusClusterAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusConfigRequest() (request *DeletePrometheusConfigRequest) {
    request = &DeletePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusConfig")
    
    
    return
}

func NewDeletePrometheusConfigResponse() (response *DeletePrometheusConfigResponse) {
    response = &DeletePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusConfig
// This API is used to delete Prometheus configurations. If the target cluster does not exist, a result indicating success will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfig(request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    return c.DeletePrometheusConfigWithContext(context.Background(), request)
}

// DeletePrometheusConfig
// This API is used to delete Prometheus configurations. If the target cluster does not exist, a result indicating success will be returned.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) DeletePrometheusConfigWithContext(ctx context.Context, request *DeletePrometheusConfigRequest) (response *DeletePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusRecordRuleYamlRequest() (request *DeletePrometheusRecordRuleYamlRequest) {
    request = &DeletePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusRecordRuleYaml")
    
    
    return
}

func NewDeletePrometheusRecordRuleYamlResponse() (response *DeletePrometheusRecordRuleYamlResponse) {
    response = &DeletePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusRecordRuleYaml
// This API is used to delete a recording instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYaml(request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    return c.DeletePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// DeletePrometheusRecordRuleYaml
// This API is used to delete a recording instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DeletePrometheusRecordRuleYamlWithContext(ctx context.Context, request *DeletePrometheusRecordRuleYamlRequest) (response *DeletePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusScrapeJobsRequest() (request *DeletePrometheusScrapeJobsRequest) {
    request = &DeletePrometheusScrapeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusScrapeJobs")
    
    
    return
}

func NewDeletePrometheusScrapeJobsResponse() (response *DeletePrometheusScrapeJobsResponse) {
    response = &DeletePrometheusScrapeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusScrapeJobs
// This API is used to delete a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusScrapeJobs(request *DeletePrometheusScrapeJobsRequest) (response *DeletePrometheusScrapeJobsResponse, err error) {
    return c.DeletePrometheusScrapeJobsWithContext(context.Background(), request)
}

// DeletePrometheusScrapeJobs
// This API is used to delete a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeletePrometheusScrapeJobsWithContext(ctx context.Context, request *DeletePrometheusScrapeJobsRequest) (response *DeletePrometheusScrapeJobsResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusScrapeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusScrapeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusScrapeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempRequest() (request *DeletePrometheusTempRequest) {
    request = &DeletePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusTemp")
    
    
    return
}

func NewDeletePrometheusTempResponse() (response *DeletePrometheusTempResponse) {
    response = &DeletePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusTemp
// This API is used to delete a TMP template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTemp(request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    return c.DeletePrometheusTempWithContext(context.Background(), request)
}

// DeletePrometheusTemp
// This API is used to delete a TMP template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempWithContext(ctx context.Context, request *DeletePrometheusTempRequest) (response *DeletePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrometheusTempSyncRequest() (request *DeletePrometheusTempSyncRequest) {
    request = &DeletePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeletePrometheusTempSync")
    
    
    return
}

func NewDeletePrometheusTempSyncResponse() (response *DeletePrometheusTempSyncResponse) {
    response = &DeletePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePrometheusTempSync
// This API is used to unsync a template, which will delete the configuration generated by the template in the target. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSync(request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    return c.DeletePrometheusTempSyncWithContext(context.Background(), request)
}

// DeletePrometheusTempSync
// This API is used to unsync a template, which will delete the configuration generated by the template in the target. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_DBRECORDNOTFOUND = "FailedOperation.DbRecordNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeletePrometheusTempSyncWithContext(ctx context.Context, request *DeletePrometheusTempSyncRequest) (response *DeletePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDeletePrometheusTempSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeletePrometheusTempSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordingRulesRequest() (request *DeleteRecordingRulesRequest) {
    request = &DeleteRecordingRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteRecordingRules")
    
    
    return
}

func NewDeleteRecordingRulesResponse() (response *DeleteRecordingRulesResponse) {
    response = &DeleteRecordingRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRecordingRules
// This API is used to batch delete Prometheus recording rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRecordingRules(request *DeleteRecordingRulesRequest) (response *DeleteRecordingRulesResponse, err error) {
    return c.DeleteRecordingRulesWithContext(context.Background(), request)
}

// DeleteRecordingRules
// This API is used to batch delete Prometheus recording rules.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRecordingRulesWithContext(ctx context.Context, request *DeleteRecordingRulesRequest) (response *DeleteRecordingRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRecordingRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteRecordingRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRecordingRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRecordingRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSSOAccountRequest() (request *DeleteSSOAccountRequest) {
    request = &DeleteSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DeleteSSOAccount")
    
    
    return
}

func NewDeleteSSOAccountResponse() (response *DeleteSSOAccountResponse) {
    response = &DeleteSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSSOAccount
// This API is used to delete an authorized TCMG user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSSOAccount(request *DeleteSSOAccountRequest) (response *DeleteSSOAccountResponse, err error) {
    return c.DeleteSSOAccountWithContext(context.Background(), request)
}

// DeleteSSOAccount
// This API is used to delete an authorized TCMG user.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteSSOAccountWithContext(ctx context.Context, request *DeleteSSOAccountRequest) (response *DeleteSSOAccountResponse, err error) {
    if request == nil {
        request = NewDeleteSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DeleteSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccidentEventListRequest() (request *DescribeAccidentEventListRequest) {
    request = &DescribeAccidentEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAccidentEventList")
    
    
    return
}

func NewDescribeAccidentEventListResponse() (response *DescribeAccidentEventListResponse) {
    response = &DescribeAccidentEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccidentEventList
// This API is used to get the platform event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventList(request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    return c.DescribeAccidentEventListWithContext(context.Background(), request)
}

// DescribeAccidentEventList
// This API is used to get the platform event list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccidentEventListWithContext(ctx context.Context, request *DescribeAccidentEventListRequest) (response *DescribeAccidentEventListResponse, err error) {
    if request == nil {
        request = NewDescribeAccidentEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAccidentEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccidentEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccidentEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmEventsRequest() (request *DescribeAlarmEventsRequest) {
    request = &DescribeAlarmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmEvents")
    
    
    return
}

func NewDescribeAlarmEventsResponse() (response *DescribeAlarmEventsResponse) {
    response = &DescribeAlarmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmEvents
// This API is used to query the list of alarm events.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEvents(request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    return c.DescribeAlarmEventsWithContext(context.Background(), request)
}

// DescribeAlarmEvents
// This API is used to query the list of alarm events.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmEventsWithContext(ctx context.Context, request *DescribeAlarmEventsRequest) (response *DescribeAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmHistoriesRequest() (request *DescribeAlarmHistoriesRequest) {
    request = &DescribeAlarmHistoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmHistories")
    
    
    return
}

func NewDescribeAlarmHistoriesResponse() (response *DescribeAlarmHistoriesResponse) {
    response = &DescribeAlarmHistoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmHistories
// This API is used to query the alarm records.
//
// 
//
// Note: **If you use a sub-account, you can only query the alarm records of authorized projects** or uncategorized products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistories(request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    return c.DescribeAlarmHistoriesWithContext(context.Background(), request)
}

// DescribeAlarmHistories
// This API is used to query the alarm records.
//
// 
//
// Note: **If you use a sub-account, you can only query the alarm records of authorized projects** or uncategorized products.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmHistoriesWithContext(ctx context.Context, request *DescribeAlarmHistoriesRequest) (response *DescribeAlarmHistoriesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmHistoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmHistories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmHistories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmHistoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmMetricsRequest() (request *DescribeAlarmMetricsRequest) {
    request = &DescribeAlarmMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmMetrics")
    
    
    return
}

func NewDescribeAlarmMetricsResponse() (response *DescribeAlarmMetricsResponse) {
    response = &DescribeAlarmMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmMetrics
// This API is used to query the list of alarm metrics.
//
// error code that may be returned:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetrics(request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    return c.DescribeAlarmMetricsWithContext(context.Background(), request)
}

// DescribeAlarmMetrics
// This API is used to query the list of alarm metrics.
//
// error code that may be returned:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetricsWithContext(ctx context.Context, request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeRequest() (request *DescribeAlarmNoticeRequest) {
    request = &DescribeAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotice")
    
    
    return
}

func NewDescribeAlarmNoticeResponse() (response *DescribeAlarmNoticeResponse) {
    response = &DescribeAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotice
// This API is used to query the details of a single notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotice(request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    return c.DescribeAlarmNoticeWithContext(context.Background(), request)
}

// DescribeAlarmNotice
// This API is used to query the details of a single notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeWithContext(ctx context.Context, request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticeCallbacksRequest() (request *DescribeAlarmNoticeCallbacksRequest) {
    request = &DescribeAlarmNoticeCallbacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNoticeCallbacks")
    
    
    return
}

func NewDescribeAlarmNoticeCallbacksResponse() (response *DescribeAlarmNoticeCallbacksResponse) {
    response = &DescribeAlarmNoticeCallbacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNoticeCallbacks
// This API is used to obtain all the callback URLs of an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacks(request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    return c.DescribeAlarmNoticeCallbacksWithContext(context.Background(), request)
}

// DescribeAlarmNoticeCallbacks
// This API is used to obtain all the callback URLs of an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeCallbacksWithContext(ctx context.Context, request *DescribeAlarmNoticeCallbacksRequest) (response *DescribeAlarmNoticeCallbacksResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeCallbacksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNoticeCallbacks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNoticeCallbacks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticeCallbacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
    request = &DescribeAlarmNoticesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmNotices")
    
    
    return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
    response = &DescribeAlarmNoticesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmNotices
// This API is used to query the list of notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    return c.DescribeAlarmNoticesWithContext(context.Background(), request)
}

// DescribeAlarmNotices
// This API is used to query the list of notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmNotices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmNotices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmNoticesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPoliciesRequest() (request *DescribeAlarmPoliciesRequest) {
    request = &DescribeAlarmPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicies")
    
    
    return
}

func NewDescribeAlarmPoliciesResponse() (response *DescribeAlarmPoliciesResponse) {
    response = &DescribeAlarmPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmPolicies
// This API is used to query the list of alarm policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicies(request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    return c.DescribeAlarmPoliciesWithContext(context.Background(), request)
}

// DescribeAlarmPolicies
// This API is used to query the list of alarm policies.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPoliciesWithContext(ctx context.Context, request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
    request = &DescribeAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmPolicy")
    
    
    return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
    response = &DescribeAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmPolicy
// This API is used to get the details of a single alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    return c.DescribeAlarmPolicyWithContext(context.Background(), request)
}

// DescribeAlarmPolicy
// This API is used to get the details of a single alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicyWithContext(ctx context.Context, request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlertRulesRequest() (request *DescribeAlertRulesRequest) {
    request = &DescribeAlertRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlertRules")
    
    
    return
}

func NewDescribeAlertRulesResponse() (response *DescribeAlertRulesResponse) {
    response = &DescribeAlertRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlertRules
// This API is used to query a Prometheus alerting rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlertRules(request *DescribeAlertRulesRequest) (response *DescribeAlertRulesResponse, err error) {
    return c.DescribeAlertRulesWithContext(context.Background(), request)
}

// DescribeAlertRules
// This API is used to query a Prometheus alerting rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAlertRulesWithContext(ctx context.Context, request *DescribeAlertRulesRequest) (response *DescribeAlertRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAlertRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAlertRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlertRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlertRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllNamespacesRequest() (request *DescribeAllNamespacesRequest) {
    request = &DescribeAllNamespacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeAllNamespaces")
    
    
    return
}

func NewDescribeAllNamespacesResponse() (response *DescribeAllNamespacesResponse) {
    response = &DescribeAllNamespacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllNamespaces
// This API is used to query all namespaces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespaces(request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    return c.DescribeAllNamespacesWithContext(context.Background(), request)
}

// DescribeAllNamespaces
// This API is used to query all namespaces.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespacesWithContext(ctx context.Context, request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAllNamespacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeAllNamespaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllNamespaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllNamespacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
    request = &DescribeBaseMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
    
    
    return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
    response = &DescribeBaseMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBaseMetrics
// This API is used to get the attributes of basic metrics.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    return c.DescribeBaseMetricsWithContext(context.Background(), request)
}

// DescribeBaseMetrics
// This API is used to get the attributes of basic metrics.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBaseMetricsWithContext(ctx context.Context, request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeBaseMetricsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBaseMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBaseMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBaseMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicAlarmListRequest() (request *DescribeBasicAlarmListRequest) {
    request = &DescribeBasicAlarmListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBasicAlarmList")
    
    
    return
}

func NewDescribeBasicAlarmListResponse() (response *DescribeBasicAlarmListResponse) {
    response = &DescribeBasicAlarmListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBasicAlarmList
// This API is used to get the basic alarm list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmList(request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    return c.DescribeBasicAlarmListWithContext(context.Background(), request)
}

// DescribeBasicAlarmList
// This API is used to get the basic alarm list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBasicAlarmListWithContext(ctx context.Context, request *DescribeBasicAlarmListRequest) (response *DescribeBasicAlarmListResponse, err error) {
    if request == nil {
        request = NewDescribeBasicAlarmListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBasicAlarmList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicAlarmList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicAlarmListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindingPolicyObjectListRequest() (request *DescribeBindingPolicyObjectListRequest) {
    request = &DescribeBindingPolicyObjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeBindingPolicyObjectList")
    
    
    return
}

func NewDescribeBindingPolicyObjectListResponse() (response *DescribeBindingPolicyObjectListResponse) {
    response = &DescribeBindingPolicyObjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBindingPolicyObjectList
// This API is used to get the bound object list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectList(request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    return c.DescribeBindingPolicyObjectListWithContext(context.Background(), request)
}

// DescribeBindingPolicyObjectList
// This API is used to get the bound object list.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DRUIDTABLENOTFOUND = "FailedOperation.DruidTableNotFound"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBindingPolicyObjectListWithContext(ctx context.Context, request *DescribeBindingPolicyObjectListRequest) (response *DescribeBindingPolicyObjectListResponse, err error) {
    if request == nil {
        request = NewDescribeBindingPolicyObjectListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeBindingPolicyObjectList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindingPolicyObjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindingPolicyObjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterAgentCreatingProgressRequest() (request *DescribeClusterAgentCreatingProgressRequest) {
    request = &DescribeClusterAgentCreatingProgressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeClusterAgentCreatingProgress")
    
    
    return
}

func NewDescribeClusterAgentCreatingProgressResponse() (response *DescribeClusterAgentCreatingProgressResponse) {
    response = &DescribeClusterAgentCreatingProgressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterAgentCreatingProgress
// This API is used to obtain the binding status between the TencentCloud Managed Service for Prometheus instance and the TKE cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterAgentCreatingProgress(request *DescribeClusterAgentCreatingProgressRequest) (response *DescribeClusterAgentCreatingProgressResponse, err error) {
    return c.DescribeClusterAgentCreatingProgressWithContext(context.Background(), request)
}

// DescribeClusterAgentCreatingProgress
// This API is used to obtain the binding status between the TencentCloud Managed Service for Prometheus instance and the TKE cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribeClusterAgentCreatingProgressWithContext(ctx context.Context, request *DescribeClusterAgentCreatingProgressRequest) (response *DescribeClusterAgentCreatingProgressResponse, err error) {
    if request == nil {
        request = NewDescribeClusterAgentCreatingProgressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeClusterAgentCreatingProgress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterAgentCreatingProgress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterAgentCreatingProgressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConditionsTemplateListRequest() (request *DescribeConditionsTemplateListRequest) {
    request = &DescribeConditionsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeConditionsTemplateList")
    
    
    return
}

func NewDescribeConditionsTemplateListResponse() (response *DescribeConditionsTemplateListResponse) {
    response = &DescribeConditionsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConditionsTemplateList
// This API is used to get the trigger condition template.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConditionsTemplateList(request *DescribeConditionsTemplateListRequest) (response *DescribeConditionsTemplateListResponse, err error) {
    return c.DescribeConditionsTemplateListWithContext(context.Background(), request)
}

// DescribeConditionsTemplateList
// This API is used to get the trigger condition template.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeConditionsTemplateListWithContext(ctx context.Context, request *DescribeConditionsTemplateListRequest) (response *DescribeConditionsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeConditionsTemplateListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeConditionsTemplateList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConditionsTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConditionsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDNSConfigRequest() (request *DescribeDNSConfigRequest) {
    request = &DescribeDNSConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeDNSConfig")
    
    
    return
}

func NewDescribeDNSConfigResponse() (response *DescribeDNSConfigResponse) {
    response = &DescribeDNSConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDNSConfig
// This API is used to list Grafana DNS configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDNSConfig(request *DescribeDNSConfigRequest) (response *DescribeDNSConfigResponse, err error) {
    return c.DescribeDNSConfigWithContext(context.Background(), request)
}

// DescribeDNSConfig
// This API is used to list Grafana DNS configurations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDNSConfigWithContext(ctx context.Context, request *DescribeDNSConfigRequest) (response *DescribeDNSConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDNSConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeDNSConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDNSConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDNSConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExporterIntegrationsRequest() (request *DescribeExporterIntegrationsRequest) {
    request = &DescribeExporterIntegrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeExporterIntegrations")
    
    
    return
}

func NewDescribeExporterIntegrationsResponse() (response *DescribeExporterIntegrationsResponse) {
    response = &DescribeExporterIntegrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExporterIntegrations
// This API is used to query the list of exporter integrations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExporterIntegrations(request *DescribeExporterIntegrationsRequest) (response *DescribeExporterIntegrationsResponse, err error) {
    return c.DescribeExporterIntegrationsWithContext(context.Background(), request)
}

// DescribeExporterIntegrations
// This API is used to query the list of exporter integrations.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeExporterIntegrationsWithContext(ctx context.Context, request *DescribeExporterIntegrationsRequest) (response *DescribeExporterIntegrationsResponse, err error) {
    if request == nil {
        request = NewDescribeExporterIntegrationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeExporterIntegrations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExporterIntegrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExporterIntegrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaChannelsRequest() (request *DescribeGrafanaChannelsRequest) {
    request = &DescribeGrafanaChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaChannels")
    
    
    return
}

func NewDescribeGrafanaChannelsResponse() (response *DescribeGrafanaChannelsResponse) {
    response = &DescribeGrafanaChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaChannels
// This API is used to list all Grafana alert channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaChannels(request *DescribeGrafanaChannelsRequest) (response *DescribeGrafanaChannelsResponse, err error) {
    return c.DescribeGrafanaChannelsWithContext(context.Background(), request)
}

// DescribeGrafanaChannels
// This API is used to list all Grafana alert channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaChannelsWithContext(ctx context.Context, request *DescribeGrafanaChannelsRequest) (response *DescribeGrafanaChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaConfigRequest() (request *DescribeGrafanaConfigRequest) {
    request = &DescribeGrafanaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaConfig")
    
    
    return
}

func NewDescribeGrafanaConfigResponse() (response *DescribeGrafanaConfigResponse) {
    response = &DescribeGrafanaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaConfig
// This API is used to list Grafana settings, i.e., the `grafana.ini` file content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaConfig(request *DescribeGrafanaConfigRequest) (response *DescribeGrafanaConfigResponse, err error) {
    return c.DescribeGrafanaConfigWithContext(context.Background(), request)
}

// DescribeGrafanaConfig
// This API is used to list Grafana settings, i.e., the `grafana.ini` file content.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaConfigWithContext(ctx context.Context, request *DescribeGrafanaConfigRequest) (response *DescribeGrafanaConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaEnvironmentsRequest() (request *DescribeGrafanaEnvironmentsRequest) {
    request = &DescribeGrafanaEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaEnvironments")
    
    
    return
}

func NewDescribeGrafanaEnvironmentsResponse() (response *DescribeGrafanaEnvironmentsResponse) {
    response = &DescribeGrafanaEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaEnvironments
// This API is used to list Grafana environment variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaEnvironments(request *DescribeGrafanaEnvironmentsRequest) (response *DescribeGrafanaEnvironmentsResponse, err error) {
    return c.DescribeGrafanaEnvironmentsWithContext(context.Background(), request)
}

// DescribeGrafanaEnvironments
// This API is used to list Grafana environment variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaEnvironmentsWithContext(ctx context.Context, request *DescribeGrafanaEnvironmentsRequest) (response *DescribeGrafanaEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaInstancesRequest() (request *DescribeGrafanaInstancesRequest) {
    request = &DescribeGrafanaInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaInstances")
    
    
    return
}

func NewDescribeGrafanaInstancesResponse() (response *DescribeGrafanaInstancesResponse) {
    response = &DescribeGrafanaInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaInstances
// This API is used to list all Grafana instances under a user account.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeGrafanaInstances(request *DescribeGrafanaInstancesRequest) (response *DescribeGrafanaInstancesResponse, err error) {
    return c.DescribeGrafanaInstancesWithContext(context.Background(), request)
}

// DescribeGrafanaInstances
// This API is used to list all Grafana instances under a user account.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeGrafanaInstancesWithContext(ctx context.Context, request *DescribeGrafanaInstancesRequest) (response *DescribeGrafanaInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaIntegrationsRequest() (request *DescribeGrafanaIntegrationsRequest) {
    request = &DescribeGrafanaIntegrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaIntegrations")
    
    
    return
}

func NewDescribeGrafanaIntegrationsResponse() (response *DescribeGrafanaIntegrationsResponse) {
    response = &DescribeGrafanaIntegrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaIntegrations
// This API is used to list installed Grafana integrations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaIntegrations(request *DescribeGrafanaIntegrationsRequest) (response *DescribeGrafanaIntegrationsResponse, err error) {
    return c.DescribeGrafanaIntegrationsWithContext(context.Background(), request)
}

// DescribeGrafanaIntegrations
// This API is used to list installed Grafana integrations.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaIntegrationsWithContext(ctx context.Context, request *DescribeGrafanaIntegrationsRequest) (response *DescribeGrafanaIntegrationsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaIntegrationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaIntegrations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaIntegrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaIntegrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaNotificationChannelsRequest() (request *DescribeGrafanaNotificationChannelsRequest) {
    request = &DescribeGrafanaNotificationChannelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaNotificationChannels")
    
    
    return
}

func NewDescribeGrafanaNotificationChannelsResponse() (response *DescribeGrafanaNotificationChannelsResponse) {
    response = &DescribeGrafanaNotificationChannelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaNotificationChannels
// This API is used to list Grafana notification channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaNotificationChannels(request *DescribeGrafanaNotificationChannelsRequest) (response *DescribeGrafanaNotificationChannelsResponse, err error) {
    return c.DescribeGrafanaNotificationChannelsWithContext(context.Background(), request)
}

// DescribeGrafanaNotificationChannels
// This API is used to list Grafana notification channels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaNotificationChannelsWithContext(ctx context.Context, request *DescribeGrafanaNotificationChannelsRequest) (response *DescribeGrafanaNotificationChannelsResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaNotificationChannelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaNotificationChannels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaNotificationChannels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaNotificationChannelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGrafanaWhiteListRequest() (request *DescribeGrafanaWhiteListRequest) {
    request = &DescribeGrafanaWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeGrafanaWhiteList")
    
    
    return
}

func NewDescribeGrafanaWhiteListResponse() (response *DescribeGrafanaWhiteListResponse) {
    response = &DescribeGrafanaWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGrafanaWhiteList
// This API is used to list the Grafana allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaWhiteList(request *DescribeGrafanaWhiteListRequest) (response *DescribeGrafanaWhiteListResponse, err error) {
    return c.DescribeGrafanaWhiteListWithContext(context.Background(), request)
}

// DescribeGrafanaWhiteList
// This API is used to list the Grafana allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeGrafanaWhiteListWithContext(ctx context.Context, request *DescribeGrafanaWhiteListRequest) (response *DescribeGrafanaWhiteListResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeGrafanaWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGrafanaWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGrafanaWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstalledPluginsRequest() (request *DescribeInstalledPluginsRequest) {
    request = &DescribeInstalledPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeInstalledPlugins")
    
    
    return
}

func NewDescribeInstalledPluginsResponse() (response *DescribeInstalledPluginsResponse) {
    response = &DescribeInstalledPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstalledPlugins
// This API is used to list the plugins installed in an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstalledPlugins(request *DescribeInstalledPluginsRequest) (response *DescribeInstalledPluginsResponse, err error) {
    return c.DescribeInstalledPluginsWithContext(context.Background(), request)
}

// DescribeInstalledPlugins
// This API is used to list the plugins installed in an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeInstalledPluginsWithContext(ctx context.Context, request *DescribeInstalledPluginsRequest) (response *DescribeInstalledPluginsResponse, err error) {
    if request == nil {
        request = NewDescribeInstalledPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeInstalledPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstalledPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstalledPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMonitorTypesRequest() (request *DescribeMonitorTypesRequest) {
    request = &DescribeMonitorTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeMonitorTypes")
    
    
    return
}

func NewDescribeMonitorTypesResponse() (response *DescribeMonitorTypesResponse) {
    response = &DescribeMonitorTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMonitorTypes
// This API is used to list all the monitoring types supported by CM.
//
// error code that may be returned:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypes(request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    return c.DescribeMonitorTypesWithContext(context.Background(), request)
}

// DescribeMonitorTypes
// This API is used to list all the monitoring types supported by CM.
//
// error code that may be returned:
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypesWithContext(ctx context.Context, request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorTypesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeMonitorTypes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMonitorTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMonitorTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyConditionListRequest() (request *DescribePolicyConditionListRequest) {
    request = &DescribePolicyConditionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyConditionList")
    
    
    return
}

func NewDescribePolicyConditionListResponse() (response *DescribePolicyConditionListResponse) {
    response = &DescribePolicyConditionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyConditionList
// This API is used to get basic alarm policy conditions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionList(request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    return c.DescribePolicyConditionListWithContext(context.Background(), request)
}

// DescribePolicyConditionList
// This API is used to get basic alarm policy conditions.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyConditionListWithContext(ctx context.Context, request *DescribePolicyConditionListRequest) (response *DescribePolicyConditionListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConditionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyConditionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyConditionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyConditionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupInfoRequest() (request *DescribePolicyGroupInfoRequest) {
    request = &DescribePolicyGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupInfo")
    
    
    return
}

func NewDescribePolicyGroupInfoResponse() (response *DescribePolicyGroupInfoResponse) {
    response = &DescribePolicyGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyGroupInfo
// This API is used to get details of a basic policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfo(request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    return c.DescribePolicyGroupInfoWithContext(context.Background(), request)
}

// DescribePolicyGroupInfo
// This API is used to get details of a basic policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupInfoWithContext(ctx context.Context, request *DescribePolicyGroupInfoRequest) (response *DescribePolicyGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyGroupListRequest() (request *DescribePolicyGroupListRequest) {
    request = &DescribePolicyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePolicyGroupList")
    
    
    return
}

func NewDescribePolicyGroupListResponse() (response *DescribePolicyGroupListResponse) {
    response = &DescribePolicyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyGroupList
// This API is used to get the list of basic policy alarm groups.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupList(request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    return c.DescribePolicyGroupListWithContext(context.Background(), request)
}

// DescribePolicyGroupList
// This API is used to get the list of basic policy alarm groups.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePolicyGroupListWithContext(ctx context.Context, request *DescribePolicyGroupListRequest) (response *DescribePolicyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribePolicyGroupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePolicyGroupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductEventListRequest() (request *DescribeProductEventListRequest) {
    request = &DescribeProductEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeProductEventList")
    
    
    return
}

func NewDescribeProductEventListResponse() (response *DescribeProductEventListResponse) {
    response = &DescribeProductEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductEventList
// This API is used to get the list of product events by page.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventList(request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    return c.DescribeProductEventListWithContext(context.Background(), request)
}

// DescribeProductEventList
// This API is used to get the list of product events by page.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProductEventListWithContext(ctx context.Context, request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
    if request == nil {
        request = NewDescribeProductEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeProductEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentInstancesRequest() (request *DescribePrometheusAgentInstancesRequest) {
    request = &DescribePrometheusAgentInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAgentInstances")
    
    
    return
}

func NewDescribePrometheusAgentInstancesResponse() (response *DescribePrometheusAgentInstancesResponse) {
    response = &DescribePrometheusAgentInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAgentInstances
// This API is used to get the list of instances associated with the target cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstances(request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    return c.DescribePrometheusAgentInstancesWithContext(context.Background(), request)
}

// DescribePrometheusAgentInstances
// This API is used to get the list of instances associated with the target cluster.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusAgentInstancesWithContext(ctx context.Context, request *DescribePrometheusAgentInstancesRequest) (response *DescribePrometheusAgentInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAgentInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgentInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAgentsRequest() (request *DescribePrometheusAgentsRequest) {
    request = &DescribePrometheusAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAgents")
    
    
    return
}

func NewDescribePrometheusAgentsResponse() (response *DescribePrometheusAgentsResponse) {
    response = &DescribePrometheusAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAgents
// This API is used to list Prometheus CVM agents.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAgents(request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    return c.DescribePrometheusAgentsWithContext(context.Background(), request)
}

// DescribePrometheusAgents
// This API is used to list Prometheus CVM agents.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAgentsWithContext(ctx context.Context, request *DescribePrometheusAgentsRequest) (response *DescribePrometheusAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusAlertPolicyRequest() (request *DescribePrometheusAlertPolicyRequest) {
    request = &DescribePrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusAlertPolicy")
    
    
    return
}

func NewDescribePrometheusAlertPolicyResponse() (response *DescribePrometheusAlertPolicyResponse) {
    response = &DescribePrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusAlertPolicy
// This API is used to get the list of v2.0 instance alerting rules.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAlertPolicy(request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    return c.DescribePrometheusAlertPolicyWithContext(context.Background(), request)
}

// DescribePrometheusAlertPolicy
// This API is used to get the list of v2.0 instance alerting rules.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusAlertPolicyWithContext(ctx context.Context, request *DescribePrometheusAlertPolicyRequest) (response *DescribePrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusClusterAgentsRequest() (request *DescribePrometheusClusterAgentsRequest) {
    request = &DescribePrometheusClusterAgentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusClusterAgents")
    
    
    return
}

func NewDescribePrometheusClusterAgentsResponse() (response *DescribePrometheusClusterAgentsResponse) {
    response = &DescribePrometheusClusterAgentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusClusterAgents
// This API is used to get the list of clusters associated with the TMP instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribePrometheusClusterAgents(request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    return c.DescribePrometheusClusterAgentsWithContext(context.Background(), request)
}

// DescribePrometheusClusterAgents
// This API is used to get the list of clusters associated with the TMP instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DescribePrometheusClusterAgentsWithContext(ctx context.Context, request *DescribePrometheusClusterAgentsRequest) (response *DescribePrometheusClusterAgentsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusClusterAgentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusClusterAgents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusClusterAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusClusterAgentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusConfigRequest() (request *DescribePrometheusConfigRequest) {
    request = &DescribePrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusConfig")
    
    
    return
}

func NewDescribePrometheusConfigResponse() (response *DescribePrometheusConfigResponse) {
    response = &DescribePrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusConfig
// This API is used to get the Prometheus configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfig(request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    return c.DescribePrometheusConfigWithContext(context.Background(), request)
}

// DescribePrometheusConfig
// This API is used to get the Prometheus configuration.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DB = "FailedOperation.Db"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusConfigWithContext(ctx context.Context, request *DescribePrometheusConfigRequest) (response *DescribePrometheusConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalConfigRequest() (request *DescribePrometheusGlobalConfigRequest) {
    request = &DescribePrometheusGlobalConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusGlobalConfig")
    
    
    return
}

func NewDescribePrometheusGlobalConfigResponse() (response *DescribePrometheusGlobalConfigResponse) {
    response = &DescribePrometheusGlobalConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusGlobalConfig
// This API is used to get the instance-level scrape configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfig(request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    return c.DescribePrometheusGlobalConfigWithContext(context.Background(), request)
}

// DescribePrometheusGlobalConfig
// This API is used to get the instance-level scrape configurations.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalConfigWithContext(ctx context.Context, request *DescribePrometheusGlobalConfigRequest) (response *DescribePrometheusGlobalConfigResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusGlobalConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusGlobalNotificationRequest() (request *DescribePrometheusGlobalNotificationRequest) {
    request = &DescribePrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusGlobalNotification")
    
    
    return
}

func NewDescribePrometheusGlobalNotificationResponse() (response *DescribePrometheusGlobalNotificationResponse) {
    response = &DescribePrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusGlobalNotification
// This API is used to query the global alert notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotification(request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    return c.DescribePrometheusGlobalNotificationWithContext(context.Background(), request)
}

// DescribePrometheusGlobalNotification
// This API is used to query the global alert notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusGlobalNotificationWithContext(ctx context.Context, request *DescribePrometheusGlobalNotificationRequest) (response *DescribePrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceDetailRequest() (request *DescribePrometheusInstanceDetailRequest) {
    request = &DescribePrometheusInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceDetail")
    
    
    return
}

func NewDescribePrometheusInstanceDetailResponse() (response *DescribePrometheusInstanceDetailResponse) {
    response = &DescribePrometheusInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceDetail
// This API is used to get the details of a TMP instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceDetail(request *DescribePrometheusInstanceDetailRequest) (response *DescribePrometheusInstanceDetailResponse, err error) {
    return c.DescribePrometheusInstanceDetailWithContext(context.Background(), request)
}

// DescribePrometheusInstanceDetail
// This API is used to get the details of a TMP instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceDetailWithContext(ctx context.Context, request *DescribePrometheusInstanceDetailRequest) (response *DescribePrometheusInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceInitStatusRequest() (request *DescribePrometheusInstanceInitStatusRequest) {
    request = &DescribePrometheusInstanceInitStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceInitStatus")
    
    
    return
}

func NewDescribePrometheusInstanceInitStatusResponse() (response *DescribePrometheusInstanceInitStatusResponse) {
    response = &DescribePrometheusInstanceInitStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceInitStatus
// This API is used to get the initialization task status of a v2.0 instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstanceInitStatus(request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    return c.DescribePrometheusInstanceInitStatusWithContext(context.Background(), request)
}

// DescribePrometheusInstanceInitStatus
// This API is used to get the initialization task status of a v2.0 instance.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_TASKNOTFOUND = "InternalError.TaskNotFound"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstanceInitStatusWithContext(ctx context.Context, request *DescribePrometheusInstanceInitStatusRequest) (response *DescribePrometheusInstanceInitStatusResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceInitStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceInitStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceInitStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceInitStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstanceUsageRequest() (request *DescribePrometheusInstanceUsageRequest) {
    request = &DescribePrometheusInstanceUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstanceUsage")
    
    
    return
}

func NewDescribePrometheusInstanceUsageResponse() (response *DescribePrometheusInstanceUsageResponse) {
    response = &DescribePrometheusInstanceUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstanceUsage
//  This API is used to query the usage of a pay-as-you-go Tencent Managed Service for Prometheus (TMP) instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceUsage(request *DescribePrometheusInstanceUsageRequest) (response *DescribePrometheusInstanceUsageResponse, err error) {
    return c.DescribePrometheusInstanceUsageWithContext(context.Background(), request)
}

// DescribePrometheusInstanceUsage
//  This API is used to query the usage of a pay-as-you-go Tencent Managed Service for Prometheus (TMP) instance.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusInstanceUsageWithContext(ctx context.Context, request *DescribePrometheusInstanceUsageRequest) (response *DescribePrometheusInstanceUsageResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstanceUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstanceUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstanceUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstanceUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstancesRequest() (request *DescribePrometheusInstancesRequest) {
    request = &DescribePrometheusInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstances")
    
    
    return
}

func NewDescribePrometheusInstancesResponse() (response *DescribePrometheusInstancesResponse) {
    response = &DescribePrometheusInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstances
// This API is used to query the details of one or multiple instances.
//
// <ul>
//
// <li>You can query the details of an instance by its ID, name, or status.</li>
//
// <li>If this parameter is empty, the information of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.</li>
//
// </ul>
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePrometheusInstances(request *DescribePrometheusInstancesRequest) (response *DescribePrometheusInstancesResponse, err error) {
    return c.DescribePrometheusInstancesWithContext(context.Background(), request)
}

// DescribePrometheusInstances
// This API is used to query the details of one or multiple instances.
//
// <ul>
//
// <li>You can query the details of an instance by its ID, name, or status.</li>
//
// <li>If this parameter is empty, the information of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.</li>
//
// </ul>
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribePrometheusInstancesWithContext(ctx context.Context, request *DescribePrometheusInstancesRequest) (response *DescribePrometheusInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusInstancesOverviewRequest() (request *DescribePrometheusInstancesOverviewRequest) {
    request = &DescribePrometheusInstancesOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusInstancesOverview")
    
    
    return
}

func NewDescribePrometheusInstancesOverviewResponse() (response *DescribePrometheusInstancesOverviewResponse) {
    response = &DescribePrometheusInstancesOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusInstancesOverview
// This API is used to obtain the list of Tencent Managed Service for Prometheus (TMP) instances and the clusters associated with them.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverview(request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    return c.DescribePrometheusInstancesOverviewWithContext(context.Background(), request)
}

// DescribePrometheusInstancesOverview
// This API is used to obtain the list of Tencent Managed Service for Prometheus (TMP) instances and the clusters associated with them.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusInstancesOverviewWithContext(ctx context.Context, request *DescribePrometheusInstancesOverviewRequest) (response *DescribePrometheusInstancesOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusInstancesOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusInstancesOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstancesOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusRecordRuleYamlRequest() (request *DescribePrometheusRecordRuleYamlRequest) {
    request = &DescribePrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusRecordRuleYaml")
    
    
    return
}

func NewDescribePrometheusRecordRuleYamlResponse() (response *DescribePrometheusRecordRuleYamlResponse) {
    response = &DescribePrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusRecordRuleYaml
// This API is used to get the YAML list of Prometheus recording rules.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRuleYaml(request *DescribePrometheusRecordRuleYamlRequest) (response *DescribePrometheusRecordRuleYamlResponse, err error) {
    return c.DescribePrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// DescribePrometheusRecordRuleYaml
// This API is used to get the YAML list of Prometheus recording rules.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRuleYamlWithContext(ctx context.Context, request *DescribePrometheusRecordRuleYamlRequest) (response *DescribePrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusRecordRulesRequest() (request *DescribePrometheusRecordRulesRequest) {
    request = &DescribePrometheusRecordRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusRecordRules")
    
    
    return
}

func NewDescribePrometheusRecordRulesResponse() (response *DescribePrometheusRecordRulesResponse) {
    response = &DescribePrometheusRecordRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusRecordRules
// This API is used to get the list of recording rules, including those created by CRD resources in the associated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRules(request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    return c.DescribePrometheusRecordRulesWithContext(context.Background(), request)
}

// DescribePrometheusRecordRules
// This API is used to get the list of recording rules, including those created by CRD resources in the associated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECLIENTCONF = "FailedOperation.KubeClientConf"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusRecordRulesWithContext(ctx context.Context, request *DescribePrometheusRecordRulesRequest) (response *DescribePrometheusRecordRulesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusRecordRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusRecordRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusRecordRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusRecordRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusScrapeJobsRequest() (request *DescribePrometheusScrapeJobsRequest) {
    request = &DescribePrometheusScrapeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusScrapeJobs")
    
    
    return
}

func NewDescribePrometheusScrapeJobsResponse() (response *DescribePrometheusScrapeJobsResponse) {
    response = &DescribePrometheusScrapeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusScrapeJobs
// This API is used to list Prometheus scrape tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusScrapeJobs(request *DescribePrometheusScrapeJobsRequest) (response *DescribePrometheusScrapeJobsResponse, err error) {
    return c.DescribePrometheusScrapeJobsWithContext(context.Background(), request)
}

// DescribePrometheusScrapeJobs
// This API is used to list Prometheus scrape tasks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusScrapeJobsWithContext(ctx context.Context, request *DescribePrometheusScrapeJobsRequest) (response *DescribePrometheusScrapeJobsResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusScrapeJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusScrapeJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusScrapeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTargetsTMPRequest() (request *DescribePrometheusTargetsTMPRequest) {
    request = &DescribePrometheusTargetsTMPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTargetsTMP")
    
    
    return
}

func NewDescribePrometheusTargetsTMPResponse() (response *DescribePrometheusTargetsTMPResponse) {
    response = &DescribePrometheusTargetsTMPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTargetsTMP
// This API is used to get the targets information.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargetsTMP(request *DescribePrometheusTargetsTMPRequest) (response *DescribePrometheusTargetsTMPResponse, err error) {
    return c.DescribePrometheusTargetsTMPWithContext(context.Background(), request)
}

// DescribePrometheusTargetsTMP
// This API is used to get the targets information.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) DescribePrometheusTargetsTMPWithContext(ctx context.Context, request *DescribePrometheusTargetsTMPRequest) (response *DescribePrometheusTargetsTMPResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTargetsTMPRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTargetsTMP")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTargetsTMP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTargetsTMPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempRequest() (request *DescribePrometheusTempRequest) {
    request = &DescribePrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTemp")
    
    
    return
}

func NewDescribePrometheusTempResponse() (response *DescribePrometheusTempResponse) {
    response = &DescribePrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTemp
// This API is used to get the list of templates, where the default template is always on top.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTemp(request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    return c.DescribePrometheusTempWithContext(context.Background(), request)
}

// DescribePrometheusTemp
// This API is used to get the list of templates, where the default template is always on top.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) DescribePrometheusTempWithContext(ctx context.Context, request *DescribePrometheusTempRequest) (response *DescribePrometheusTempResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusTempSyncRequest() (request *DescribePrometheusTempSyncRequest) {
    request = &DescribePrometheusTempSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusTempSync")
    
    
    return
}

func NewDescribePrometheusTempSyncResponse() (response *DescribePrometheusTempSyncResponse) {
    response = &DescribePrometheusTempSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusTempSync
// This API is used to get the information of instances associated with a template. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSync(request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    return c.DescribePrometheusTempSyncWithContext(context.Background(), request)
}

// DescribePrometheusTempSync
// This API is used to get the information of instances associated with a template. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribePrometheusTempSyncWithContext(ctx context.Context, request *DescribePrometheusTempSyncRequest) (response *DescribePrometheusTempSyncResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusTempSyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusTempSync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusTempSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusTempSyncResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrometheusZonesRequest() (request *DescribePrometheusZonesRequest) {
    request = &DescribePrometheusZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribePrometheusZones")
    
    
    return
}

func NewDescribePrometheusZonesResponse() (response *DescribePrometheusZonesResponse) {
    response = &DescribePrometheusZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrometheusZones
// This API is used to list the AZs of Tencent Managed Service for Prometheus (TMP).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusZones(request *DescribePrometheusZonesRequest) (response *DescribePrometheusZonesResponse, err error) {
    return c.DescribePrometheusZonesWithContext(context.Background(), request)
}

// DescribePrometheusZones
// This API is used to list the AZs of Tencent Managed Service for Prometheus (TMP).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribePrometheusZonesWithContext(ctx context.Context, request *DescribePrometheusZonesRequest) (response *DescribePrometheusZonesResponse, err error) {
    if request == nil {
        request = NewDescribePrometheusZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribePrometheusZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingRulesRequest() (request *DescribeRecordingRulesRequest) {
    request = &DescribeRecordingRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeRecordingRules")
    
    
    return
}

func NewDescribeRecordingRulesResponse() (response *DescribeRecordingRulesResponse) {
    response = &DescribeRecordingRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingRules
// This API is used to query Prometheus recording rules by filter.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecordingRules(request *DescribeRecordingRulesRequest) (response *DescribeRecordingRulesResponse, err error) {
    return c.DescribeRecordingRulesWithContext(context.Background(), request)
}

// DescribeRecordingRules
// This API is used to query Prometheus recording rules by filter.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRecordingRulesWithContext(ctx context.Context, request *DescribeRecordingRulesRequest) (response *DescribeRecordingRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeRecordingRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSSOAccountRequest() (request *DescribeSSOAccountRequest) {
    request = &DescribeSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeSSOAccount")
    
    
    return
}

func NewDescribeSSOAccountResponse() (response *DescribeSSOAccountResponse) {
    response = &DescribeSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSSOAccount
// This API is used to list all authorized accounts of the current Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSSOAccount(request *DescribeSSOAccountRequest) (response *DescribeSSOAccountResponse, err error) {
    return c.DescribeSSOAccountWithContext(context.Background(), request)
}

// DescribeSSOAccount
// This API is used to list all authorized accounts of the current Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSSOAccountWithContext(ctx context.Context, request *DescribeSSOAccountRequest) (response *DescribeSSOAccountResponse, err error) {
    if request == nil {
        request = NewDescribeSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceDiscoveryRequest() (request *DescribeServiceDiscoveryRequest) {
    request = &DescribeServiceDiscoveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeServiceDiscovery")
    
    
    return
}

func NewDescribeServiceDiscoveryResponse() (response *DescribeServiceDiscoveryResponse) {
    response = &DescribeServiceDiscoveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceDiscovery
// This API is used to list Prometheus scrape configurations in TKE.
//
// <p>Note: The prerequisite is that the corresponding TKE service has been integrated through the Prometheus console. For more information, see
//
// <a href="https://intl.cloud.tencent.com/document/product/248/48859?from_cn_redirect=1" target="_blank">Agent Management</a>.</p>
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeServiceDiscovery(request *DescribeServiceDiscoveryRequest) (response *DescribeServiceDiscoveryResponse, err error) {
    return c.DescribeServiceDiscoveryWithContext(context.Background(), request)
}

// DescribeServiceDiscovery
// This API is used to list Prometheus scrape configurations in TKE.
//
// <p>Note: The prerequisite is that the corresponding TKE service has been integrated through the Prometheus console. For more information, see
//
// <a href="https://intl.cloud.tencent.com/document/product/248/48859?from_cn_redirect=1" target="_blank">Agent Management</a>.</p>
//
// error code that may be returned:
//  FAILEDOPERATION_ACCESSSTSFAIL = "FailedOperation.AccessSTSFail"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_TKEENDPOINTSTATUSERROR = "FailedOperation.TKEEndpointStatusError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeServiceDiscoveryWithContext(ctx context.Context, request *DescribeServiceDiscoveryRequest) (response *DescribeServiceDiscoveryResponse, err error) {
    if request == nil {
        request = NewDescribeServiceDiscoveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeServiceDiscovery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceDiscovery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceDiscoveryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStatisticDataRequest() (request *DescribeStatisticDataRequest) {
    request = &DescribeStatisticDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DescribeStatisticData")
    
    
    return
}

func NewDescribeStatisticDataResponse() (response *DescribeStatisticDataResponse) {
    response = &DescribeStatisticDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStatisticData
// This API is used to query monitoring data by dimension conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticData(request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    return c.DescribeStatisticDataWithContext(context.Background(), request)
}

// DescribeStatisticData
// This API is used to query monitoring data by dimension conditions.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATACOLUMNNOTFOUND = "FailedOperation.DataColumnNotFound"
//  FAILEDOPERATION_DATAQUERYFAILED = "FailedOperation.DataQueryFailed"
//  FAILEDOPERATION_DATATABLENOTFOUND = "FailedOperation.DataTableNotFound"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DIVISIONBYZERO = "FailedOperation.DivisionByZero"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLBACKFAIL = "InternalError.CallbackFail"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_DEPENDSMQ = "InternalError.DependsMq"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INTERNALERROR_TASKRESULTFORMAT = "InternalError.TaskResultFormat"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPTASK = "InvalidParameter.DupTask"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_MISSAKSK = "InvalidParameter.MissAKSK"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SECRETIDORSECRETKEYERROR = "InvalidParameter.SecretIdOrSecretKeyError"
//  INVALIDPARAMETER_UNSUPPORTEDPRODUCT = "InvalidParameter.UnsupportedProduct"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DASHBOARDNAMEEXISTS = "InvalidParameterValue.DashboardNameExists"
//  INVALIDPARAMETERVALUE_VERSIONMISMATCH = "InvalidParameterValue.VersionMismatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeStatisticDataWithContext(ctx context.Context, request *DescribeStatisticDataRequest) (response *DescribeStatisticDataResponse, err error) {
    if request == nil {
        request = NewDescribeStatisticDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DescribeStatisticData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStatisticData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStatisticDataResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPrometheusInstanceRequest() (request *DestroyPrometheusInstanceRequest) {
    request = &DestroyPrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "DestroyPrometheusInstance")
    
    
    return
}

func NewDestroyPrometheusInstanceResponse() (response *DestroyPrometheusInstanceResponse) {
    response = &DestroyPrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyPrometheusInstance
// This API is used to delete the data of a Prometheus instance. The specified instance must be terminated first.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyPrometheusInstance(request *DestroyPrometheusInstanceRequest) (response *DestroyPrometheusInstanceResponse, err error) {
    return c.DestroyPrometheusInstanceWithContext(context.Background(), request)
}

// DestroyPrometheusInstance
// This API is used to delete the data of a Prometheus instance. The specified instance must be terminated first.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyPrometheusInstanceWithContext(ctx context.Context, request *DestroyPrometheusInstanceRequest) (response *DestroyPrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPrometheusInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "DestroyPrometheusInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewEnableGrafanaInternetRequest() (request *EnableGrafanaInternetRequest) {
    request = &EnableGrafanaInternetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableGrafanaInternet")
    
    
    return
}

func NewEnableGrafanaInternetResponse() (response *EnableGrafanaInternetResponse) {
    response = &EnableGrafanaInternetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableGrafanaInternet
// This API is used to set the Grafana public network access.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaInternet(request *EnableGrafanaInternetRequest) (response *EnableGrafanaInternetResponse, err error) {
    return c.EnableGrafanaInternetWithContext(context.Background(), request)
}

// EnableGrafanaInternet
// This API is used to set the Grafana public network access.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaInternetWithContext(ctx context.Context, request *EnableGrafanaInternetRequest) (response *EnableGrafanaInternetResponse, err error) {
    if request == nil {
        request = NewEnableGrafanaInternetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableGrafanaInternet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGrafanaInternet require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableGrafanaInternetResponse()
    err = c.Send(request, response)
    return
}

func NewEnableGrafanaSSORequest() (request *EnableGrafanaSSORequest) {
    request = &EnableGrafanaSSORequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableGrafanaSSO")
    
    
    return
}

func NewEnableGrafanaSSOResponse() (response *EnableGrafanaSSOResponse) {
    response = &EnableGrafanaSSOResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableGrafanaSSO
// This API is used to set the Grafana SSO through a Tencent Cloud account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaSSO(request *EnableGrafanaSSORequest) (response *EnableGrafanaSSOResponse, err error) {
    return c.EnableGrafanaSSOWithContext(context.Background(), request)
}

// EnableGrafanaSSO
// This API is used to set the Grafana SSO through a Tencent Cloud account.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) EnableGrafanaSSOWithContext(ctx context.Context, request *EnableGrafanaSSORequest) (response *EnableGrafanaSSOResponse, err error) {
    if request == nil {
        request = NewEnableGrafanaSSORequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableGrafanaSSO")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableGrafanaSSO require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableGrafanaSSOResponse()
    err = c.Send(request, response)
    return
}

func NewEnableSSOCamCheckRequest() (request *EnableSSOCamCheckRequest) {
    request = &EnableSSOCamCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "EnableSSOCamCheck")
    
    
    return
}

func NewEnableSSOCamCheckResponse() (response *EnableSSOCamCheckResponse) {
    response = &EnableSSOCamCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableSSOCamCheck
// This API is used to set whether to enable CAM authentication during SSO.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) EnableSSOCamCheck(request *EnableSSOCamCheckRequest) (response *EnableSSOCamCheckResponse, err error) {
    return c.EnableSSOCamCheckWithContext(context.Background(), request)
}

// EnableSSOCamCheck
// This API is used to set whether to enable CAM authentication during SSO.
//
// error code that may be returned:
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
func (c *Client) EnableSSOCamCheckWithContext(ctx context.Context, request *EnableSSOCamCheckRequest) (response *EnableSSOCamCheckResponse, err error) {
    if request == nil {
        request = NewEnableSSOCamCheckRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "EnableSSOCamCheck")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableSSOCamCheck require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableSSOCamCheckResponse()
    err = c.Send(request, response)
    return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
    request = &GetMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
    
    
    return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
    response = &GetMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMonitorData
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, use the [DescribeStatisticData](https://www.tencentcloud.com/document/product/248/39481) API.
//
// You can get the monitoring data of a Tencent Cloud service by passing in its namespace, object dimension description, and monitoring metrics.
//
// API call rate limit: 20 calls/second (1,200 calls/minute). A single request can get the data of up to 10 instances for up to 1,440 data points.
//
// If you need to call a large number of APIs to pull metrics or objects at a time, some APIs may fail to be called due to the rate limit. We suggest you evenly arrange API calls at a time granularity.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    return c.GetMonitorDataWithContext(context.Background(), request)
}

// GetMonitorData
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, use the [DescribeStatisticData](https://www.tencentcloud.com/document/product/248/39481) API.
//
// You can get the monitoring data of a Tencent Cloud service by passing in its namespace, object dimension description, and monitoring metrics.
//
// API call rate limit: 20 calls/second (1,200 calls/minute). A single request can get the data of up to 10 instances for up to 1,440 data points.
//
// If you need to call a large number of APIs to pull metrics or objects at a time, some APIs may fail to be called due to the rate limit. We suggest you evenly arrange API calls at a time granularity.
//
// error code that may be returned:
//  FAILEDOPERATION_ERRNOTOPEN = "FailedOperation.ErrNotOpen"
//  FAILEDOPERATION_ERROWED = "FailedOperation.ErrOwed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMonitorDataWithContext(ctx context.Context, request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
    if request == nil {
        request = NewGetMonitorDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetMonitorData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMonitorDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetPrometheusAgentManagementCommandRequest() (request *GetPrometheusAgentManagementCommandRequest) {
    request = &GetPrometheusAgentManagementCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "GetPrometheusAgentManagementCommand")
    
    
    return
}

func NewGetPrometheusAgentManagementCommandResponse() (response *GetPrometheusAgentManagementCommandResponse) {
    response = &GetPrometheusAgentManagementCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetPrometheusAgentManagementCommand
// This API is used to get the command line for Prometheus agent management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetPrometheusAgentManagementCommand(request *GetPrometheusAgentManagementCommandRequest) (response *GetPrometheusAgentManagementCommandResponse, err error) {
    return c.GetPrometheusAgentManagementCommandWithContext(context.Background(), request)
}

// GetPrometheusAgentManagementCommand
// This API is used to get the command line for Prometheus agent management.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) GetPrometheusAgentManagementCommandWithContext(ctx context.Context, request *GetPrometheusAgentManagementCommandRequest) (response *GetPrometheusAgentManagementCommandResponse, err error) {
    if request == nil {
        request = NewGetPrometheusAgentManagementCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "GetPrometheusAgentManagementCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPrometheusAgentManagementCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetPrometheusAgentManagementCommandResponse()
    err = c.Send(request, response)
    return
}

func NewInstallPluginsRequest() (request *InstallPluginsRequest) {
    request = &InstallPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "InstallPlugins")
    
    
    return
}

func NewInstallPluginsResponse() (response *InstallPluginsResponse) {
    response = &InstallPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallPlugins
// This API is used to install a Grafana plugin.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) InstallPlugins(request *InstallPluginsRequest) (response *InstallPluginsResponse, err error) {
    return c.InstallPluginsWithContext(context.Background(), request)
}

// InstallPlugins
// This API is used to install a Grafana plugin.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) InstallPluginsWithContext(ctx context.Context, request *InstallPluginsRequest) (response *InstallPluginsResponse, err error) {
    if request == nil {
        request = NewInstallPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "InstallPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
    request = &ModifyAlarmNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmNotice")
    
    
    return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
    response = &ModifyAlarmNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmNotice
// This API is used to edit an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    return c.ModifyAlarmNoticeWithContext(context.Background(), request)
}

// ModifyAlarmNotice
// This API is used to edit an alarm notification template.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyConditionRequest() (request *ModifyAlarmPolicyConditionRequest) {
    request = &ModifyAlarmPolicyConditionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyCondition")
    
    
    return
}

func NewModifyAlarmPolicyConditionResponse() (response *ModifyAlarmPolicyConditionResponse) {
    response = &ModifyAlarmPolicyConditionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyCondition
// This API is used to modify the trigger condition of an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyCondition(request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    return c.ModifyAlarmPolicyConditionWithContext(context.Background(), request)
}

// ModifyAlarmPolicyCondition
// This API is used to modify the trigger condition of an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyConditionWithContext(ctx context.Context, request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyConditionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyCondition")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyCondition require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyConditionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyInfoRequest() (request *ModifyAlarmPolicyInfoRequest) {
    request = &ModifyAlarmPolicyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyInfo")
    
    
    return
}

func NewModifyAlarmPolicyInfoResponse() (response *ModifyAlarmPolicyInfoResponse) {
    response = &ModifyAlarmPolicyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyInfo
// This API is used to edit the basic information of a v2.0 alarm policy, including policy name and remarks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfo(request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    return c.ModifyAlarmPolicyInfoWithContext(context.Background(), request)
}

// ModifyAlarmPolicyInfo
// This API is used to edit the basic information of a v2.0 alarm policy, including policy name and remarks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyInfoWithContext(ctx context.Context, request *ModifyAlarmPolicyInfoRequest) (response *ModifyAlarmPolicyInfoResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyNoticeRequest() (request *ModifyAlarmPolicyNoticeRequest) {
    request = &ModifyAlarmPolicyNoticeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyNotice")
    
    
    return
}

func NewModifyAlarmPolicyNoticeResponse() (response *ModifyAlarmPolicyNoticeResponse) {
    response = &ModifyAlarmPolicyNoticeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyNotice
// This API is used to modify the alarm notification template bound to an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyNotice(request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    return c.ModifyAlarmPolicyNoticeWithContext(context.Background(), request)
}

// ModifyAlarmPolicyNotice
// This API is used to modify the alarm notification template bound to an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyNoticeWithContext(ctx context.Context, request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyNoticeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyNotice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyNotice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyNoticeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyStatusRequest() (request *ModifyAlarmPolicyStatusRequest) {
    request = &ModifyAlarmPolicyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyStatus")
    
    
    return
}

func NewModifyAlarmPolicyStatusResponse() (response *ModifyAlarmPolicyStatusResponse) {
    response = &ModifyAlarmPolicyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyStatus
// This API is used to enable/disable an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatus(request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    return c.ModifyAlarmPolicyStatusWithContext(context.Background(), request)
}

// ModifyAlarmPolicyStatus
// This API is used to enable/disable an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DOHTTPTRANSFERFAILED = "FailedOperation.DoHTTPTransferFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatusWithContext(ctx context.Context, request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyTasksRequest() (request *ModifyAlarmPolicyTasksRequest) {
    request = &ModifyAlarmPolicyTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmPolicyTasks")
    
    
    return
}

func NewModifyAlarmPolicyTasksResponse() (response *ModifyAlarmPolicyTasksResponse) {
    response = &ModifyAlarmPolicyTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicyTasks
// This API is used to modify the tasks triggered by alarm policy, which are listed by the value of the `TriggerTasks` field. If an empty array is passed in for `TriggerTasks`, it means unbinding all the trigger tasks from the policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasks(request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    return c.ModifyAlarmPolicyTasksWithContext(context.Background(), request)
}

// ModifyAlarmPolicyTasks
// This API is used to modify the tasks triggered by alarm policy, which are listed by the value of the `TriggerTasks` field. If an empty array is passed in for `TriggerTasks`, it means unbinding all the trigger tasks from the policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyTasksWithContext(ctx context.Context, request *ModifyAlarmPolicyTasksRequest) (response *ModifyAlarmPolicyTasksResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmPolicyTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicyTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmReceiversRequest() (request *ModifyAlarmReceiversRequest) {
    request = &ModifyAlarmReceiversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyAlarmReceivers")
    
    
    return
}

func NewModifyAlarmReceiversResponse() (response *ModifyAlarmReceiversResponse) {
    response = &ModifyAlarmReceiversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmReceivers
// This API is used to modify alarm recipients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceivers(request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    return c.ModifyAlarmReceiversWithContext(context.Background(), request)
}

// ModifyAlarmReceivers
// This API is used to modify alarm recipients.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmReceiversWithContext(ctx context.Context, request *ModifyAlarmReceiversRequest) (response *ModifyAlarmReceiversResponse, err error) {
    if request == nil {
        request = NewModifyAlarmReceiversRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyAlarmReceivers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmReceivers require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmReceiversResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGrafanaInstanceRequest() (request *ModifyGrafanaInstanceRequest) {
    request = &ModifyGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyGrafanaInstance")
    
    
    return
}

func NewModifyGrafanaInstanceResponse() (response *ModifyGrafanaInstanceResponse) {
    response = &ModifyGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGrafanaInstance
// This API is used to modify the attributes of a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) ModifyGrafanaInstance(request *ModifyGrafanaInstanceRequest) (response *ModifyGrafanaInstanceResponse, err error) {
    return c.ModifyGrafanaInstanceWithContext(context.Background(), request)
}

// ModifyGrafanaInstance
// This API is used to modify the attributes of a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) ModifyGrafanaInstanceWithContext(ctx context.Context, request *ModifyGrafanaInstanceRequest) (response *ModifyGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewModifyGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyGroupRequest() (request *ModifyPolicyGroupRequest) {
    request = &ModifyPolicyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPolicyGroup")
    
    
    return
}

func NewModifyPolicyGroupResponse() (response *ModifyPolicyGroupResponse) {
    response = &ModifyPolicyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPolicyGroup
// This API is used to update policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    return c.ModifyPolicyGroupWithContext(context.Background(), request)
}

// ModifyPolicyGroup
// This API is used to update policy group.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPolicyGroupWithContext(ctx context.Context, request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
    if request == nil {
        request = NewModifyPolicyGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPolicyGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPolicyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAgentExternalLabelsRequest() (request *ModifyPrometheusAgentExternalLabelsRequest) {
    request = &ModifyPrometheusAgentExternalLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusAgentExternalLabels")
    
    
    return
}

func NewModifyPrometheusAgentExternalLabelsResponse() (response *ModifyPrometheusAgentExternalLabelsResponse) {
    response = &ModifyPrometheusAgentExternalLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusAgentExternalLabels
// This API is used to modify the external labels of the associated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabels(request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    return c.ModifyPrometheusAgentExternalLabelsWithContext(context.Background(), request)
}

// ModifyPrometheusAgentExternalLabels
// This API is used to modify the external labels of the associated cluster.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusAgentExternalLabelsWithContext(ctx context.Context, request *ModifyPrometheusAgentExternalLabelsRequest) (response *ModifyPrometheusAgentExternalLabelsResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAgentExternalLabelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusAgentExternalLabels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAgentExternalLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAgentExternalLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusAlertPolicyRequest() (request *ModifyPrometheusAlertPolicyRequest) {
    request = &ModifyPrometheusAlertPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusAlertPolicy")
    
    
    return
}

func NewModifyPrometheusAlertPolicyResponse() (response *ModifyPrometheusAlertPolicyResponse) {
    response = &ModifyPrometheusAlertPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusAlertPolicy
// This API is used to modify a TMP 2.0 instance alerting rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyPrometheusAlertPolicy(request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    return c.ModifyPrometheusAlertPolicyWithContext(context.Background(), request)
}

// ModifyPrometheusAlertPolicy
// This API is used to modify a TMP 2.0 instance alerting rule.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
func (c *Client) ModifyPrometheusAlertPolicyWithContext(ctx context.Context, request *ModifyPrometheusAlertPolicyRequest) (response *ModifyPrometheusAlertPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusAlertPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusAlertPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusAlertPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusAlertPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusConfigRequest() (request *ModifyPrometheusConfigRequest) {
    request = &ModifyPrometheusConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusConfig")
    
    
    return
}

func NewModifyPrometheusConfigResponse() (response *ModifyPrometheusConfigResponse) {
    response = &ModifyPrometheusConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusConfig
// This API is used to modify the Prometheus configuration. If there are no configuration items, one will be added.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusConfig(request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    return c.ModifyPrometheusConfigWithContext(context.Background(), request)
}

// ModifyPrometheusConfig
// This API is used to modify the Prometheus configuration. If there are no configuration items, one will be added.
//
// error code that may be returned:
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLUSTERNOTFOUND = "InvalidParameter.ClusterNotFound"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
func (c *Client) ModifyPrometheusConfigWithContext(ctx context.Context, request *ModifyPrometheusConfigRequest) (response *ModifyPrometheusConfigResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusGlobalNotificationRequest() (request *ModifyPrometheusGlobalNotificationRequest) {
    request = &ModifyPrometheusGlobalNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusGlobalNotification")
    
    
    return
}

func NewModifyPrometheusGlobalNotificationResponse() (response *ModifyPrometheusGlobalNotificationResponse) {
    response = &ModifyPrometheusGlobalNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusGlobalNotification
// This API is used to modify the global alert notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotification(request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    return c.ModifyPrometheusGlobalNotificationWithContext(context.Background(), request)
}

// ModifyPrometheusGlobalNotification
// This API is used to modify the global alert notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusGlobalNotificationWithContext(ctx context.Context, request *ModifyPrometheusGlobalNotificationRequest) (response *ModifyPrometheusGlobalNotificationResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusGlobalNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusGlobalNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusGlobalNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusGlobalNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusInstanceAttributesRequest() (request *ModifyPrometheusInstanceAttributesRequest) {
    request = &ModifyPrometheusInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusInstanceAttributes")
    
    
    return
}

func NewModifyPrometheusInstanceAttributesResponse() (response *ModifyPrometheusInstanceAttributesResponse) {
    response = &ModifyPrometheusInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusInstanceAttributes
// This API is used to modify the attributes of a Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPrometheusInstanceAttributes(request *ModifyPrometheusInstanceAttributesRequest) (response *ModifyPrometheusInstanceAttributesResponse, err error) {
    return c.ModifyPrometheusInstanceAttributesWithContext(context.Background(), request)
}

// ModifyPrometheusInstanceAttributes
// This API is used to modify the attributes of a Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyPrometheusInstanceAttributesWithContext(ctx context.Context, request *ModifyPrometheusInstanceAttributesRequest) (response *ModifyPrometheusInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusInstanceAttributesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusInstanceAttributes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusRecordRuleYamlRequest() (request *ModifyPrometheusRecordRuleYamlRequest) {
    request = &ModifyPrometheusRecordRuleYamlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusRecordRuleYaml")
    
    
    return
}

func NewModifyPrometheusRecordRuleYamlResponse() (response *ModifyPrometheusRecordRuleYamlResponse) {
    response = &ModifyPrometheusRecordRuleYamlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusRecordRuleYaml
// This API is used to modify a Prometheus recording instance through YAML.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYaml(request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    return c.ModifyPrometheusRecordRuleYamlWithContext(context.Background(), request)
}

// ModifyPrometheusRecordRuleYaml
// This API is used to modify a Prometheus recording instance through YAML.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) ModifyPrometheusRecordRuleYamlWithContext(ctx context.Context, request *ModifyPrometheusRecordRuleYamlRequest) (response *ModifyPrometheusRecordRuleYamlResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusRecordRuleYamlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusRecordRuleYaml")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusRecordRuleYaml require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusRecordRuleYamlResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrometheusTempRequest() (request *ModifyPrometheusTempRequest) {
    request = &ModifyPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ModifyPrometheusTemp")
    
    
    return
}

func NewModifyPrometheusTempResponse() (response *ModifyPrometheusTempResponse) {
    response = &ModifyPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPrometheusTemp
// This API is used to modify a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTemp(request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    return c.ModifyPrometheusTempWithContext(context.Background(), request)
}

// ModifyPrometheusTemp
// This API is used to modify a template.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyPrometheusTempWithContext(ctx context.Context, request *ModifyPrometheusTempRequest) (response *ModifyPrometheusTempResponse, err error) {
    if request == nil {
        request = NewModifyPrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ModifyPrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewResumeGrafanaInstanceRequest() (request *ResumeGrafanaInstanceRequest) {
    request = &ResumeGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "ResumeGrafanaInstance")
    
    
    return
}

func NewResumeGrafanaInstanceResponse() (response *ResumeGrafanaInstanceResponse) {
    response = &ResumeGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeGrafanaInstance
// This API is used to renew a monthly subscribed TCMG instance for a month without changing the instance edition. It doesn't apply to running instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeGrafanaInstance(request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    return c.ResumeGrafanaInstanceWithContext(context.Background(), request)
}

// ResumeGrafanaInstance
// This API is used to renew a monthly subscribed TCMG instance for a month without changing the instance edition. It doesn't apply to running instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResumeGrafanaInstanceWithContext(ctx context.Context, request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewResumeGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "ResumeGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunPrometheusInstanceRequest() (request *RunPrometheusInstanceRequest) {
    request = &RunPrometheusInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "RunPrometheusInstance")
    
    
    return
}

func NewRunPrometheusInstanceResponse() (response *RunPrometheusInstanceResponse) {
    response = &RunPrometheusInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunPrometheusInstance
// This API is used to initialize a TMP instance, which is called when the integration center is enabled.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstance(request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    return c.RunPrometheusInstanceWithContext(context.Background(), request)
}

// RunPrometheusInstance
// This API is used to initialize a TMP instance, which is called when the integration center is enabled.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
func (c *Client) RunPrometheusInstanceWithContext(ctx context.Context, request *RunPrometheusInstanceRequest) (response *RunPrometheusInstanceResponse, err error) {
    if request == nil {
        request = NewRunPrometheusInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "RunPrometheusInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunPrometheusInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunPrometheusInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSendCustomAlarmMsgRequest() (request *SendCustomAlarmMsgRequest) {
    request = &SendCustomAlarmMsgRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "SendCustomAlarmMsg")
    
    
    return
}

func NewSendCustomAlarmMsgResponse() (response *SendCustomAlarmMsgResponse) {
    response = &SendCustomAlarmMsgResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendCustomAlarmMsg
// This API is used to send a custom alarm notification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCustomAlarmMsg(request *SendCustomAlarmMsgRequest) (response *SendCustomAlarmMsgResponse, err error) {
    return c.SendCustomAlarmMsgWithContext(context.Background(), request)
}

// SendCustomAlarmMsg
// This API is used to send a custom alarm notification.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SendCustomAlarmMsgWithContext(ctx context.Context, request *SendCustomAlarmMsgRequest) (response *SendCustomAlarmMsgResponse, err error) {
    if request == nil {
        request = NewSendCustomAlarmMsgRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "SendCustomAlarmMsg")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendCustomAlarmMsg require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendCustomAlarmMsgResponse()
    err = c.Send(request, response)
    return
}

func NewSetDefaultAlarmPolicyRequest() (request *SetDefaultAlarmPolicyRequest) {
    request = &SetDefaultAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "SetDefaultAlarmPolicy")
    
    
    return
}

func NewSetDefaultAlarmPolicyResponse() (response *SetDefaultAlarmPolicyResponse) {
    response = &SetDefaultAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetDefaultAlarmPolicy
// This API is used to set an alarm policy as the default policy in the current policy type under the current project.
//
// Alarm policies in the same type under the project will be set as non-default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicy(request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    return c.SetDefaultAlarmPolicyWithContext(context.Background(), request)
}

// SetDefaultAlarmPolicy
// This API is used to set an alarm policy as the default policy in the current policy type under the current project.
//
// Alarm policies in the same type under the project will be set as non-default.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetDefaultAlarmPolicyWithContext(ctx context.Context, request *SetDefaultAlarmPolicyRequest) (response *SetDefaultAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewSetDefaultAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "SetDefaultAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDefaultAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDefaultAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewSyncPrometheusTempRequest() (request *SyncPrometheusTempRequest) {
    request = &SyncPrometheusTempRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "SyncPrometheusTemp")
    
    
    return
}

func NewSyncPrometheusTempResponse() (response *SyncPrometheusTempResponse) {
    response = &SyncPrometheusTempResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncPrometheusTemp
// This API is used to sync a template to an instance or cluster. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTemp(request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    return c.SyncPrometheusTempWithContext(context.Background(), request)
}

// SyncPrometheusTemp
// This API is used to sync a template to an instance or cluster. It takes effect for v2 instances.
//
// error code that may be returned:
//  FAILEDOPERATION_COMPONENTCLIENTCOMMON = "FailedOperation.ComponentClientCommon"
//  FAILEDOPERATION_COMPONENTCLIENTHTTP = "FailedOperation.ComponentClientHttp"
//  FAILEDOPERATION_COMPONENTCLIENTUNPACK = "FailedOperation.ComponentClientUnpack"
//  FAILEDOPERATION_DOTRPCTRANSFERFAILED = "FailedOperation.DoTRPCTransferFailed"
//  FAILEDOPERATION_KUBECOMMON = "FailedOperation.KubeCommon"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PARAM = "InternalError.Param"
//  INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"
//  INVALIDPARAMETER_PARAM = "InvalidParameter.Param"
//  INVALIDPARAMETER_PROMCLUSTERNOTFOUND = "InvalidParameter.PromClusterNotFound"
//  INVALIDPARAMETER_PROMINSTANCENOTFOUND = "InvalidParameter.PromInstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SyncPrometheusTempWithContext(ctx context.Context, request *SyncPrometheusTempRequest) (response *SyncPrometheusTempResponse, err error) {
    if request == nil {
        request = NewSyncPrometheusTempRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "SyncPrometheusTemp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncPrometheusTemp require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncPrometheusTempResponse()
    err = c.Send(request, response)
    return
}

func NewTerminatePrometheusInstancesRequest() (request *TerminatePrometheusInstancesRequest) {
    request = &TerminatePrometheusInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "TerminatePrometheusInstances")
    
    
    return
}

func NewTerminatePrometheusInstancesResponse() (response *TerminatePrometheusInstancesResponse) {
    response = &TerminatePrometheusInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminatePrometheusInstances
// This API is used to terminate a pay-as-you-go Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TerminatePrometheusInstances(request *TerminatePrometheusInstancesRequest) (response *TerminatePrometheusInstancesResponse, err error) {
    return c.TerminatePrometheusInstancesWithContext(context.Background(), request)
}

// TerminatePrometheusInstances
// This API is used to terminate a pay-as-you-go Prometheus instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTSNOTINUNINSTALLSTAGE = "FailedOperation.AgentsNotInUninstallStage"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) TerminatePrometheusInstancesWithContext(ctx context.Context, request *TerminatePrometheusInstancesRequest) (response *TerminatePrometheusInstancesResponse, err error) {
    if request == nil {
        request = NewTerminatePrometheusInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "TerminatePrometheusInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminatePrometheusInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminatePrometheusInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingAllPolicyObjectRequest() (request *UnBindingAllPolicyObjectRequest) {
    request = &UnBindingAllPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingAllPolicyObject")
    
    
    return
}

func NewUnBindingAllPolicyObjectResponse() (response *UnBindingAllPolicyObjectResponse) {
    response = &UnBindingAllPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindingAllPolicyObject
// This API is used to delete all bound objects.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObject(request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    return c.UnBindingAllPolicyObjectWithContext(context.Background(), request)
}

// UnBindingAllPolicyObject
// This API is used to delete all bound objects.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingAllPolicyObjectWithContext(ctx context.Context, request *UnBindingAllPolicyObjectRequest) (response *UnBindingAllPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingAllPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnBindingAllPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindingAllPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindingAllPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindingPolicyObjectRequest() (request *UnBindingPolicyObjectRequest) {
    request = &UnBindingPolicyObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnBindingPolicyObject")
    
    
    return
}

func NewUnBindingPolicyObjectResponse() (response *UnBindingPolicyObjectResponse) {
    response = &UnBindingPolicyObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnBindingPolicyObject
// This API is used to delete an object that is bound to a policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObject(request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    return c.UnBindingPolicyObjectWithContext(context.Background(), request)
}

// UnBindingPolicyObject
// This API is used to delete an object that is bound to a policy.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALERTFILTERRULEDELETEFAILED = "FailedOperation.AlertFilterRuleDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYCREATEFAILED = "FailedOperation.AlertPolicyCreateFailed"
//  FAILEDOPERATION_ALERTPOLICYDELETEFAILED = "FailedOperation.AlertPolicyDeleteFailed"
//  FAILEDOPERATION_ALERTPOLICYDESCRIBEFAILED = "FailedOperation.AlertPolicyDescribeFailed"
//  FAILEDOPERATION_ALERTPOLICYMODIFYFAILED = "FailedOperation.AlertPolicyModifyFailed"
//  FAILEDOPERATION_ALERTTRIGGERRULEDELETEFAILED = "FailedOperation.AlertTriggerRuleDeleteFailed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_DBTRANSACTIONBEGINFAILED = "FailedOperation.DbTransactionBeginFailed"
//  FAILEDOPERATION_DBTRANSACTIONCOMMITFAILED = "FailedOperation.DbTransactionCommitFailed"
//  FAILEDOPERATION_DIMQUERYREQUESTFAILED = "FailedOperation.DimQueryRequestFailed"
//  FAILEDOPERATION_DRUIDQUERYFAILED = "FailedOperation.DruidQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_SERVICENOTENABLED = "FailedOperation.ServiceNotEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXETIMEOUT = "InternalError.ExeTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERPARAM = "InvalidParameter.InvalidParameterParam"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_METRICQUOTAEXCEEDED = "LimitExceeded.MetricQuotaExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOTEXISTTASK = "ResourceNotFound.NotExistTask"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnBindingPolicyObjectWithContext(ctx context.Context, request *UnBindingPolicyObjectRequest) (response *UnBindingPolicyObjectResponse, err error) {
    if request == nil {
        request = NewUnBindingPolicyObjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnBindingPolicyObject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnBindingPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnBindingPolicyObjectResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPrometheusManagedGrafanaRequest() (request *UnbindPrometheusManagedGrafanaRequest) {
    request = &UnbindPrometheusManagedGrafanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UnbindPrometheusManagedGrafana")
    
    
    return
}

func NewUnbindPrometheusManagedGrafanaResponse() (response *UnbindPrometheusManagedGrafanaResponse) {
    response = &UnbindPrometheusManagedGrafanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindPrometheusManagedGrafana
// This API is used to unbind a Grafana instance from an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UnbindPrometheusManagedGrafana(request *UnbindPrometheusManagedGrafanaRequest) (response *UnbindPrometheusManagedGrafanaResponse, err error) {
    return c.UnbindPrometheusManagedGrafanaWithContext(context.Background(), request)
}

// UnbindPrometheusManagedGrafana
// This API is used to unbind a Grafana instance from an instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UnbindPrometheusManagedGrafanaWithContext(ctx context.Context, request *UnbindPrometheusManagedGrafanaRequest) (response *UnbindPrometheusManagedGrafanaResponse, err error) {
    if request == nil {
        request = NewUnbindPrometheusManagedGrafanaRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UnbindPrometheusManagedGrafana")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindPrometheusManagedGrafana require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindPrometheusManagedGrafanaResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallGrafanaDashboardRequest() (request *UninstallGrafanaDashboardRequest) {
    request = &UninstallGrafanaDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UninstallGrafanaDashboard")
    
    
    return
}

func NewUninstallGrafanaDashboardResponse() (response *UninstallGrafanaDashboardResponse) {
    response = &UninstallGrafanaDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallGrafanaDashboard
// This API is used to delete a Grafana dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaDashboard(request *UninstallGrafanaDashboardRequest) (response *UninstallGrafanaDashboardResponse, err error) {
    return c.UninstallGrafanaDashboardWithContext(context.Background(), request)
}

// UninstallGrafanaDashboard
// This API is used to delete a Grafana dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UninstallGrafanaDashboardWithContext(ctx context.Context, request *UninstallGrafanaDashboardRequest) (response *UninstallGrafanaDashboardResponse, err error) {
    if request == nil {
        request = NewUninstallGrafanaDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UninstallGrafanaDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallGrafanaDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallGrafanaDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallGrafanaPluginsRequest() (request *UninstallGrafanaPluginsRequest) {
    request = &UninstallGrafanaPluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UninstallGrafanaPlugins")
    
    
    return
}

func NewUninstallGrafanaPluginsResponse() (response *UninstallGrafanaPluginsResponse) {
    response = &UninstallGrafanaPluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallGrafanaPlugins
// This API is used to delete installed plugins.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UninstallGrafanaPlugins(request *UninstallGrafanaPluginsRequest) (response *UninstallGrafanaPluginsResponse, err error) {
    return c.UninstallGrafanaPluginsWithContext(context.Background(), request)
}

// UninstallGrafanaPlugins
// This API is used to delete installed plugins.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UninstallGrafanaPluginsWithContext(ctx context.Context, request *UninstallGrafanaPluginsRequest) (response *UninstallGrafanaPluginsResponse, err error) {
    if request == nil {
        request = NewUninstallGrafanaPluginsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UninstallGrafanaPlugins")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallGrafanaPlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallGrafanaPluginsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertRuleRequest() (request *UpdateAlertRuleRequest) {
    request = &UpdateAlertRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateAlertRule")
    
    
    return
}

func NewUpdateAlertRuleResponse() (response *UpdateAlertRuleResponse) {
    response = &UpdateAlertRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertRule
// This API is used to update a Prometheus alerting rule.
//
// 
//
// Note that alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively. For more information, see [Alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (response *UpdateAlertRuleResponse, err error) {
    return c.UpdateAlertRuleWithContext(context.Background(), request)
}

// UpdateAlertRule
// This API is used to update a Prometheus alerting rule.
//
// 
//
// Note that alert object and alert message are special fields of Prometheus Rule Annotations, which need to be passed in through `annotations` and correspond to `summary` and `description` keys respectively. For more information, see [Alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/).
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleWithContext(ctx context.Context, request *UpdateAlertRuleRequest) (response *UpdateAlertRuleResponse, err error) {
    if request == nil {
        request = NewUpdateAlertRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateAlertRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAlertRuleStateRequest() (request *UpdateAlertRuleStateRequest) {
    request = &UpdateAlertRuleStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateAlertRuleState")
    
    
    return
}

func NewUpdateAlertRuleStateResponse() (response *UpdateAlertRuleStateResponse) {
    response = &UpdateAlertRuleStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAlertRuleState
// This API is used to update the status of a Prometheus alerting rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleState(request *UpdateAlertRuleStateRequest) (response *UpdateAlertRuleStateResponse, err error) {
    return c.UpdateAlertRuleStateWithContext(context.Background(), request)
}

// UpdateAlertRuleState
// This API is used to update the status of a Prometheus alerting rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateAlertRuleStateWithContext(ctx context.Context, request *UpdateAlertRuleStateRequest) (response *UpdateAlertRuleStateResponse, err error) {
    if request == nil {
        request = NewUpdateAlertRuleStateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateAlertRuleState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlertRuleState require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAlertRuleStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDNSConfigRequest() (request *UpdateDNSConfigRequest) {
    request = &UpdateDNSConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateDNSConfig")
    
    
    return
}

func NewUpdateDNSConfigResponse() (response *UpdateDNSConfigResponse) {
    response = &UpdateDNSConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDNSConfig
// This API is used to update the Grafana DNS configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateDNSConfig(request *UpdateDNSConfigRequest) (response *UpdateDNSConfigResponse, err error) {
    return c.UpdateDNSConfigWithContext(context.Background(), request)
}

// UpdateDNSConfig
// This API is used to update the Grafana DNS configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateDNSConfigWithContext(ctx context.Context, request *UpdateDNSConfigRequest) (response *UpdateDNSConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDNSConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateDNSConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDNSConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDNSConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateExporterIntegrationRequest() (request *UpdateExporterIntegrationRequest) {
    request = &UpdateExporterIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateExporterIntegration")
    
    
    return
}

func NewUpdateExporterIntegrationResponse() (response *UpdateExporterIntegrationResponse) {
    response = &UpdateExporterIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateExporterIntegration
// This API is used to update the exporter integration configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateExporterIntegration(request *UpdateExporterIntegrationRequest) (response *UpdateExporterIntegrationResponse, err error) {
    return c.UpdateExporterIntegrationWithContext(context.Background(), request)
}

// UpdateExporterIntegration
// This API is used to update the exporter integration configuration.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTKEFAIL = "FailedOperation.AccessTKEFail"
//  FAILEDOPERATION_AGENTNOTALLOWED = "FailedOperation.AgentNotAllowed"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INSTANCENOTRUNNING = "FailedOperation.InstanceNotRunning"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  FAILEDOPERATION_RESOURCEOPERATING = "FailedOperation.ResourceOperating"
//  FAILEDOPERATION_TKERESOURCECONFLICT = "FailedOperation.TKEResourceConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateExporterIntegrationWithContext(ctx context.Context, request *UpdateExporterIntegrationRequest) (response *UpdateExporterIntegrationResponse, err error) {
    if request == nil {
        request = NewUpdateExporterIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateExporterIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateExporterIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateExporterIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaConfigRequest() (request *UpdateGrafanaConfigRequest) {
    request = &UpdateGrafanaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaConfig")
    
    
    return
}

func NewUpdateGrafanaConfigResponse() (response *UpdateGrafanaConfigResponse) {
    response = &UpdateGrafanaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaConfig
// This API is used to update the Grafana configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaConfig(request *UpdateGrafanaConfigRequest) (response *UpdateGrafanaConfigResponse, err error) {
    return c.UpdateGrafanaConfigWithContext(context.Background(), request)
}

// UpdateGrafanaConfig
// This API is used to update the Grafana configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaConfigWithContext(ctx context.Context, request *UpdateGrafanaConfigRequest) (response *UpdateGrafanaConfigResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaEnvironmentsRequest() (request *UpdateGrafanaEnvironmentsRequest) {
    request = &UpdateGrafanaEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaEnvironments")
    
    
    return
}

func NewUpdateGrafanaEnvironmentsResponse() (response *UpdateGrafanaEnvironmentsResponse) {
    response = &UpdateGrafanaEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaEnvironments
// This API is used to update Grafana environment variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaEnvironments(request *UpdateGrafanaEnvironmentsRequest) (response *UpdateGrafanaEnvironmentsResponse, err error) {
    return c.UpdateGrafanaEnvironmentsWithContext(context.Background(), request)
}

// UpdateGrafanaEnvironments
// This API is used to update Grafana environment variables.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaEnvironmentsWithContext(ctx context.Context, request *UpdateGrafanaEnvironmentsRequest) (response *UpdateGrafanaEnvironmentsResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaEnvironmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaEnvironments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaIntegrationRequest() (request *UpdateGrafanaIntegrationRequest) {
    request = &UpdateGrafanaIntegrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaIntegration")
    
    
    return
}

func NewUpdateGrafanaIntegrationResponse() (response *UpdateGrafanaIntegrationResponse) {
    response = &UpdateGrafanaIntegrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaIntegration
// This API is used to update the Grafana integration configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaIntegration(request *UpdateGrafanaIntegrationRequest) (response *UpdateGrafanaIntegrationResponse, err error) {
    return c.UpdateGrafanaIntegrationWithContext(context.Background(), request)
}

// UpdateGrafanaIntegration
// This API is used to update the Grafana integration configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaIntegrationWithContext(ctx context.Context, request *UpdateGrafanaIntegrationRequest) (response *UpdateGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaIntegrationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaIntegration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaIntegration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaIntegrationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaNotificationChannelRequest() (request *UpdateGrafanaNotificationChannelRequest) {
    request = &UpdateGrafanaNotificationChannelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaNotificationChannel")
    
    
    return
}

func NewUpdateGrafanaNotificationChannelResponse() (response *UpdateGrafanaNotificationChannelResponse) {
    response = &UpdateGrafanaNotificationChannelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaNotificationChannel
// This API is used to update the Grafana notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaNotificationChannel(request *UpdateGrafanaNotificationChannelRequest) (response *UpdateGrafanaNotificationChannelResponse, err error) {
    return c.UpdateGrafanaNotificationChannelWithContext(context.Background(), request)
}

// UpdateGrafanaNotificationChannel
// This API is used to update the Grafana notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaNotificationChannelWithContext(ctx context.Context, request *UpdateGrafanaNotificationChannelRequest) (response *UpdateGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaNotificationChannelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaNotificationChannel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaNotificationChannel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaNotificationChannelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGrafanaWhiteListRequest() (request *UpdateGrafanaWhiteListRequest) {
    request = &UpdateGrafanaWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateGrafanaWhiteList")
    
    
    return
}

func NewUpdateGrafanaWhiteListResponse() (response *UpdateGrafanaWhiteListResponse) {
    response = &UpdateGrafanaWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGrafanaWhiteList
// This API is used to update the Grafana allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaWhiteList(request *UpdateGrafanaWhiteListRequest) (response *UpdateGrafanaWhiteListResponse, err error) {
    return c.UpdateGrafanaWhiteListWithContext(context.Background(), request)
}

// UpdateGrafanaWhiteList
// This API is used to update the Grafana allowlist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UpdateGrafanaWhiteListWithContext(ctx context.Context, request *UpdateGrafanaWhiteListRequest) (response *UpdateGrafanaWhiteListResponse, err error) {
    if request == nil {
        request = NewUpdateGrafanaWhiteListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateGrafanaWhiteList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGrafanaWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGrafanaWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusAgentStatusRequest() (request *UpdatePrometheusAgentStatusRequest) {
    request = &UpdatePrometheusAgentStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusAgentStatus")
    
    
    return
}

func NewUpdatePrometheusAgentStatusResponse() (response *UpdatePrometheusAgentStatusResponse) {
    response = &UpdatePrometheusAgentStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusAgentStatus
// This API is used to update the status of a Prometheus CVM agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusAgentStatus(request *UpdatePrometheusAgentStatusRequest) (response *UpdatePrometheusAgentStatusResponse, err error) {
    return c.UpdatePrometheusAgentStatusWithContext(context.Background(), request)
}

// UpdatePrometheusAgentStatus
// This API is used to update the status of a Prometheus CVM agent.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusAgentStatusWithContext(ctx context.Context, request *UpdatePrometheusAgentStatusRequest) (response *UpdatePrometheusAgentStatusResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusAgentStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusAgentStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusAgentStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusAgentStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrometheusScrapeJobRequest() (request *UpdatePrometheusScrapeJobRequest) {
    request = &UpdatePrometheusScrapeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdatePrometheusScrapeJob")
    
    
    return
}

func NewUpdatePrometheusScrapeJobResponse() (response *UpdatePrometheusScrapeJobResponse) {
    response = &UpdatePrometheusScrapeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePrometheusScrapeJob
// This API is used to update a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusScrapeJob(request *UpdatePrometheusScrapeJobRequest) (response *UpdatePrometheusScrapeJobResponse, err error) {
    return c.UpdatePrometheusScrapeJobWithContext(context.Background(), request)
}

// UpdatePrometheusScrapeJob
// This API is used to update a Prometheus scrape task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdatePrometheusScrapeJobWithContext(ctx context.Context, request *UpdatePrometheusScrapeJobRequest) (response *UpdatePrometheusScrapeJobResponse, err error) {
    if request == nil {
        request = NewUpdatePrometheusScrapeJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdatePrometheusScrapeJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePrometheusScrapeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePrometheusScrapeJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRecordingRuleRequest() (request *UpdateRecordingRuleRequest) {
    request = &UpdateRecordingRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateRecordingRule")
    
    
    return
}

func NewUpdateRecordingRuleResponse() (response *UpdateRecordingRuleResponse) {
    response = &UpdateRecordingRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRecordingRule
// This API is used to update a Prometheus recording rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRecordingRule(request *UpdateRecordingRuleRequest) (response *UpdateRecordingRuleResponse, err error) {
    return c.UpdateRecordingRuleWithContext(context.Background(), request)
}

// UpdateRecordingRule
// This API is used to update a Prometheus recording rule.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BADYAMLFORMAT = "FailedOperation.BadYamlFormat"
//  FAILEDOPERATION_DBQUERYFAILED = "FailedOperation.DbQueryFailed"
//  FAILEDOPERATION_DBRECORDCREATEFAILED = "FailedOperation.DbRecordCreateFailed"
//  FAILEDOPERATION_DBRECORDDELETEFAILED = "FailedOperation.DbRecordDeleteFailed"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
//  FAILEDOPERATION_GENERATEINSTANCEIDFAILED = "FailedOperation.GenerateInstanceIDFailed"
//  FAILEDOPERATION_RESOURCEEXIST = "FailedOperation.ResourceExist"
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateRecordingRuleWithContext(ctx context.Context, request *UpdateRecordingRuleRequest) (response *UpdateRecordingRuleResponse, err error) {
    if request == nil {
        request = NewUpdateRecordingRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateRecordingRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRecordingRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRecordingRuleResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSSOAccountRequest() (request *UpdateSSOAccountRequest) {
    request = &UpdateSSOAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpdateSSOAccount")
    
    
    return
}

func NewUpdateSSOAccountResponse() (response *UpdateSSOAccountResponse) {
    response = &UpdateSSOAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSSOAccount
// This API is used to update the remarks and permission information of an authorized account in an overwriting manner.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateSSOAccount(request *UpdateSSOAccountRequest) (response *UpdateSSOAccountResponse, err error) {
    return c.UpdateSSOAccountWithContext(context.Background(), request)
}

// UpdateSSOAccount
// This API is used to update the remarks and permission information of an authorized account in an overwriting manner.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpdateSSOAccountWithContext(ctx context.Context, request *UpdateSSOAccountRequest) (response *UpdateSSOAccountResponse, err error) {
    if request == nil {
        request = NewUpdateSSOAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpdateSSOAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSSOAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSSOAccountResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGrafanaDashboardRequest() (request *UpgradeGrafanaDashboardRequest) {
    request = &UpgradeGrafanaDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpgradeGrafanaDashboard")
    
    
    return
}

func NewUpgradeGrafanaDashboardResponse() (response *UpgradeGrafanaDashboardResponse) {
    response = &UpgradeGrafanaDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeGrafanaDashboard
// This API is used to update a Grafana dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpgradeGrafanaDashboard(request *UpgradeGrafanaDashboardRequest) (response *UpgradeGrafanaDashboardResponse, err error) {
    return c.UpgradeGrafanaDashboardWithContext(context.Background(), request)
}

// UpgradeGrafanaDashboard
// This API is used to update a Grafana dashboard.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpgradeGrafanaDashboardWithContext(ctx context.Context, request *UpgradeGrafanaDashboardRequest) (response *UpgradeGrafanaDashboardResponse, err error) {
    if request == nil {
        request = NewUpgradeGrafanaDashboardRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpgradeGrafanaDashboard")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGrafanaDashboard require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGrafanaDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeGrafanaInstanceRequest() (request *UpgradeGrafanaInstanceRequest) {
    request = &UpgradeGrafanaInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "UpgradeGrafanaInstance")
    
    
    return
}

func NewUpgradeGrafanaInstanceResponse() (response *UpgradeGrafanaInstanceResponse) {
    response = &UpgradeGrafanaInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeGrafanaInstance
// This API is used to upgrade a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpgradeGrafanaInstance(request *UpgradeGrafanaInstanceRequest) (response *UpgradeGrafanaInstanceResponse, err error) {
    return c.UpgradeGrafanaInstanceWithContext(context.Background(), request)
}

// UpgradeGrafanaInstance
// This API is used to upgrade a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION_RESOURCENOTFOUND = "FailedOperation.ResourceNotFound"
func (c *Client) UpgradeGrafanaInstanceWithContext(ctx context.Context, request *UpgradeGrafanaInstanceRequest) (response *UpgradeGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeGrafanaInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "monitor", APIVersion, "UpgradeGrafanaInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}
