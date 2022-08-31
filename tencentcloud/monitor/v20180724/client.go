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
func (c *Client) BindPrometheusManagedGrafana(request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    return c.BindPrometheusManagedGrafanaWithContext(context.Background(), request)
}

// BindPrometheusManagedGrafana
// This API is used to bind a Grafana instance.
func (c *Client) BindPrometheusManagedGrafanaWithContext(ctx context.Context, request *BindPrometheusManagedGrafanaRequest) (response *BindPrometheusManagedGrafanaResponse, err error) {
    if request == nil {
        request = NewBindPrometheusManagedGrafanaRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindingPolicyObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindingPolicyObjectResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAlarmNoticeWithContext(ctx context.Context, request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewCreateAlarmNoticeRequest()
    }
    
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
// This API is used to create a Cloud Monitor alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
// This API is used to create a Cloud Monitor alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
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
// This API is used to create a Grafana instance.
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
// This API is used to create a Grafana instance.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusAgentResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusMultiTenantInstancePostPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusMultiTenantInstancePostPayModeResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrometheusScrapeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrometheusScrapeJobResponse()
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
// This API is used to delete alarm notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNotices(request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    return c.DeleteAlarmNoticesWithContext(context.Background(), request)
}

// DeleteAlarmNotices
// This API is used to delete alarm notification templates.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmNoticesWithContext(ctx context.Context, request *DeleteAlarmNoticesRequest) (response *DeleteAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmNoticesRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAlarmPolicyWithContext(ctx context.Context, request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAlarmPolicyRequest()
    }
    
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
// This API is used to delete a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstance(request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    return c.DeleteGrafanaInstanceWithContext(context.Background(), request)
}

// DeleteGrafanaInstance
// This API is used to delete a Grafana instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaInstanceWithContext(ctx context.Context, request *DeleteGrafanaInstanceRequest) (response *DeleteGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaInstanceRequest()
    }
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegration(request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    return c.DeleteGrafanaIntegrationWithContext(context.Background(), request)
}

// DeleteGrafanaIntegration
// This API is used to delete a Grafana integration configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaIntegrationWithContext(ctx context.Context, request *DeleteGrafanaIntegrationRequest) (response *DeleteGrafanaIntegrationResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaIntegrationRequest()
    }
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannel(request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    return c.DeleteGrafanaNotificationChannelWithContext(context.Background(), request)
}

// DeleteGrafanaNotificationChannel
// This API is used to delete a Grafana notification channel.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DBRECORDUPDATEFAILED = "FailedOperation.DbRecordUpdateFailed"
func (c *Client) DeleteGrafanaNotificationChannelWithContext(ctx context.Context, request *DeleteGrafanaNotificationChannelRequest) (response *DeleteGrafanaNotificationChannelResponse, err error) {
    if request == nil {
        request = NewDeleteGrafanaNotificationChannelRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePrometheusScrapeJobsResponse()
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetrics(request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    return c.DescribeAlarmMetricsWithContext(context.Background(), request)
}

// DescribeAlarmMetrics
// This API is used to query the list of alarm metrics.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmMetricsWithContext(ctx context.Context, request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmMetricsRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticeWithContext(ctx context.Context, request *DescribeAlarmNoticeRequest) (response *DescribeAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticeRequest()
    }
    
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
// This API is used to get all the callback URLs of an alarm notification template.
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
// This API is used to get all the callback URLs of an alarm notification template.
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmNoticesWithContext(ctx context.Context, request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmNoticesRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPoliciesWithContext(ctx context.Context, request *DescribeAlarmPoliciesRequest) (response *DescribeAlarmPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPoliciesRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAlarmPolicyWithContext(ctx context.Context, request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmPolicyRequest()
    }
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespaces(request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    return c.DescribeAllNamespacesWithContext(context.Background(), request)
}

// DescribeAllNamespaces
// This API is used to query all namespaces.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllNamespacesWithContext(ctx context.Context, request *DescribeAllNamespacesRequest) (response *DescribeAllNamespacesResponse, err error) {
    if request == nil {
        request = NewDescribeAllNamespacesRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindingPolicyObjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindingPolicyObjectListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExporterIntegrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExporterIntegrationsResponse()
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
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
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
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeGrafanaInstancesWithContext(ctx context.Context, request *DescribeGrafanaInstancesRequest) (response *DescribeGrafanaInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeGrafanaInstancesRequest()
    }
    
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
// This API is used to list all the monitor types supported by CM.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypes(request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    return c.DescribeMonitorTypesWithContext(context.Background(), request)
}

// DescribeMonitorTypes
// This API is used to list all the monitor types supported by CM.
//
// error code that may be returned:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMonitorTypesWithContext(ctx context.Context, request *DescribeMonitorTypesRequest) (response *DescribeMonitorTypesResponse, err error) {
    if request == nil {
        request = NewDescribeMonitorTypesRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductEventListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusAgents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusAgentsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusInstancesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrometheusScrapeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrometheusScrapeJobsResponse()
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
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, please use the API [DescribeStatisticData](https://intl.cloud.tencent.com/document/product/248/51845?from_cn_redirect=1).
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
// This API is used to get the monitoring data of Tencent Cloud services except TKE. To pull TKE’s monitoring data, please use the API [DescribeStatisticData](https://intl.cloud.tencent.com/document/product/248/51845?from_cn_redirect=1).
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmNoticeWithContext(ctx context.Context, request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmNoticeRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyConditionWithContext(ctx context.Context, request *ModifyAlarmPolicyConditionRequest) (response *ModifyAlarmPolicyConditionResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyConditionRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyNotice(request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    return c.ModifyAlarmPolicyNoticeWithContext(context.Background(), request)
}

// ModifyAlarmPolicyNotice
// This API is used to modify the alarm notification template bound to an alarm policy.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAlarmPolicyNoticeWithContext(ctx context.Context, request *ModifyAlarmPolicyNoticeRequest) (response *ModifyAlarmPolicyNoticeResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyNoticeRequest()
    }
    
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAlarmPolicyStatusWithContext(ctx context.Context, request *ModifyAlarmPolicyStatusRequest) (response *ModifyAlarmPolicyStatusResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyStatusRequest()
    }
    
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
// This API is used to modify the task triggered by an alarm policy. The `TriggerTasks` field contains the list of triggered tasks. If an empty array is passed in for `TriggerTasks`, it indicates to unbind all the triggered tasks from this policy.
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
// This API is used to modify the task triggered by an alarm policy. The `TriggerTasks` field contains the list of triggered tasks. If an empty array is passed in for `TriggerTasks`, it indicates to unbind all the triggered tasks from this policy.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPolicyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPolicyGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPrometheusInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPrometheusInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewPutMonitorDataRequest() (request *PutMonitorDataRequest) {
    request = &PutMonitorDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("monitor", APIVersion, "PutMonitorData")
    
    
    return
}

func NewPutMonitorDataResponse() (response *PutMonitorDataResponse) {
    response = &PutMonitorDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutMonitorData
// The default API request rate limit is 50 requests/sec.
//
// The default upper limit on metrics of a single tenant is 100.
//
// A maximum of 30 metric/value pairs can be reported at a time. When an error is returned for a request, no metrics/values in the request will be saved.
//
// 
//
// The reporting timestamp is the timestamp when you want to save the data. We recommend that you construct a timestamp at integer minutes.
//
// The time range of a timestamp is from 300 seconds before the current time to the current time.
//
// The data of the same IP metric/value pair must be reported by minute in chronological order.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMonitorData(request *PutMonitorDataRequest) (response *PutMonitorDataResponse, err error) {
    return c.PutMonitorDataWithContext(context.Background(), request)
}

// PutMonitorData
// The default API request rate limit is 50 requests/sec.
//
// The default upper limit on metrics of a single tenant is 100.
//
// A maximum of 30 metric/value pairs can be reported at a time. When an error is returned for a request, no metrics/values in the request will be saved.
//
// 
//
// The reporting timestamp is the timestamp when you want to save the data. We recommend that you construct a timestamp at integer minutes.
//
// The time range of a timestamp is from 300 seconds before the current time to the current time.
//
// The data of the same IP metric/value pair must be reported by minute in chronological order.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMonitorDataWithContext(ctx context.Context, request *PutMonitorDataRequest) (response *PutMonitorDataResponse, err error) {
    if request == nil {
        request = NewPutMonitorDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutMonitorData require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutMonitorDataResponse()
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
// This API is used to restore a Grafana instance.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeGrafanaInstance(request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    return c.ResumeGrafanaInstanceWithContext(context.Background(), request)
}

// ResumeGrafanaInstance
// This API is used to restore a Grafana instance.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResumeGrafanaInstanceWithContext(ctx context.Context, request *ResumeGrafanaInstanceRequest) (response *ResumeGrafanaInstanceResponse, err error) {
    if request == nil {
        request = NewResumeGrafanaInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeGrafanaInstanceResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDefaultAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDefaultAlarmPolicyResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeGrafanaInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeGrafanaInstanceResponse()
    err = c.Send(request, response)
    return
}
