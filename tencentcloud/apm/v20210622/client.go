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

package v20210622

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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


func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
    request = &CreateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
    
    
    return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
    response = &CreateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmInstance
// This API is used to create a business purchase in the APM business system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    return c.CreateApmInstanceWithContext(context.Background(), request)
}

// CreateApmInstance
// This API is used to create a business purchase in the APM business system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstanceWithContext(ctx context.Context, request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    if request == nil {
        request = NewCreateApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApmPrometheusRuleRequest() (request *CreateApmPrometheusRuleRequest) {
    request = &CreateApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmPrometheusRule")
    
    
    return
}

func NewCreateApmPrometheusRuleResponse() (response *CreateApmPrometheusRuleResponse) {
    response = &CreateApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmPrometheusRule
// This API is used to create metric match rules between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULENAMECONFLICT = "FailedOperation.PromRuleNameConflict"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) CreateApmPrometheusRule(request *CreateApmPrometheusRuleRequest) (response *CreateApmPrometheusRuleResponse, err error) {
    return c.CreateApmPrometheusRuleWithContext(context.Background(), request)
}

// CreateApmPrometheusRule
// This API is used to create metric match rules between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULENAMECONFLICT = "FailedOperation.PromRuleNameConflict"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) CreateApmPrometheusRuleWithContext(ctx context.Context, request *CreateApmPrometheusRuleRequest) (response *CreateApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewCreateApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApmSampleConfigRequest() (request *CreateApmSampleConfigRequest) {
    request = &CreateApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmSampleConfig")
    
    
    return
}

func NewCreateApmSampleConfigResponse() (response *CreateApmSampleConfigResponse) {
    response = &CreateApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmSampleConfig
// Create sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLENAMECONFLICT = "FailedOperation.SampleNameConflict"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) CreateApmSampleConfig(request *CreateApmSampleConfigRequest) (response *CreateApmSampleConfigResponse, err error) {
    return c.CreateApmSampleConfigWithContext(context.Background(), request)
}

// CreateApmSampleConfig
// Create sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLENAMECONFLICT = "FailedOperation.SampleNameConflict"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) CreateApmSampleConfigWithContext(ctx context.Context, request *CreateApmSampleConfigRequest) (response *CreateApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewCreateApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProfileTaskRequest() (request *CreateProfileTaskRequest) {
    request = &CreateProfileTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateProfileTask")
    
    
    return
}

func NewCreateProfileTaskResponse() (response *CreateProfileTaskResponse) {
    response = &CreateProfileTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProfileTask
// This API is used to create an event task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTNOTONLINEERROR = "FailedOperation.AgentNotOnlineError"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTERROR = "FailedOperation.AgentVersionNotSupportError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
func (c *Client) CreateProfileTask(request *CreateProfileTaskRequest) (response *CreateProfileTaskResponse, err error) {
    return c.CreateProfileTaskWithContext(context.Background(), request)
}

// CreateProfileTask
// This API is used to create an event task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTNOTONLINEERROR = "FailedOperation.AgentNotOnlineError"
//  FAILEDOPERATION_AGENTVERSIONNOTSUPPORTERROR = "FailedOperation.AgentVersionNotSupportError"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
func (c *Client) CreateProfileTaskWithContext(ctx context.Context, request *CreateProfileTaskRequest) (response *CreateProfileTaskResponse, err error) {
    if request == nil {
        request = NewCreateProfileTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "CreateProfileTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProfileTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProfileTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApmSampleConfigRequest() (request *DeleteApmSampleConfigRequest) {
    request = &DeleteApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DeleteApmSampleConfig")
    
    
    return
}

func NewDeleteApmSampleConfigResponse() (response *DeleteApmSampleConfigResponse) {
    response = &DeleteApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApmSampleConfig
// Delete sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DeleteApmSampleConfig(request *DeleteApmSampleConfigRequest) (response *DeleteApmSampleConfigResponse, err error) {
    return c.DeleteApmSampleConfigWithContext(context.Background(), request)
}

// DeleteApmSampleConfig
// Delete sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DeleteApmSampleConfigWithContext(ctx context.Context, request *DeleteApmSampleConfigRequest) (response *DeleteApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewDeleteApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DeleteApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
    request = &DescribeApmAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
    
    
    return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
    response = &DescribeApmAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmAgent
// Obtaining APM Access Point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    return c.DescribeApmAgentWithContext(context.Background(), request)
}

// DescribeApmAgent
// Obtaining APM Access Point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgentWithContext(ctx context.Context, request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    if request == nil {
        request = NewDescribeApmAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmApplicationConfigRequest() (request *DescribeApmApplicationConfigRequest) {
    request = &DescribeApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmApplicationConfig")
    
    
    return
}

func NewDescribeApmApplicationConfigResponse() (response *DescribeApmApplicationConfigResponse) {
    response = &DescribeApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmApplicationConfig
// This API is used to query application configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeApmApplicationConfig(request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
    return c.DescribeApmApplicationConfigWithContext(context.Background(), request)
}

// DescribeApmApplicationConfig
// This API is used to query application configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeApmApplicationConfigWithContext(ctx context.Context, request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAssociationRequest() (request *DescribeApmAssociationRequest) {
    request = &DescribeApmAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAssociation")
    
    
    return
}

func NewDescribeApmAssociationResponse() (response *DescribeApmAssociationResponse) {
    response = &DescribeApmAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmAssociation
// This API is used to query the relationship between apm Business System and other product.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
func (c *Client) DescribeApmAssociation(request *DescribeApmAssociationRequest) (response *DescribeApmAssociationResponse, err error) {
    return c.DescribeApmAssociationWithContext(context.Background(), request)
}

// DescribeApmAssociation
// This API is used to query the relationship between apm Business System and other product.
//
// error code that may be returned:
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
func (c *Client) DescribeApmAssociationWithContext(ctx context.Context, request *DescribeApmAssociationRequest) (response *DescribeApmAssociationResponse, err error) {
    if request == nil {
        request = NewDescribeApmAssociationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmAssociation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmAssociation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
    request = &DescribeApmInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
    
    
    return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
    response = &DescribeApmInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmInstances
// This API is used to obtain the list of APM business systems.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    return c.DescribeApmInstancesWithContext(context.Background(), request)
}

// DescribeApmInstances
// This API is used to obtain the list of APM business systems.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstancesWithContext(ctx context.Context, request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeApmInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmPrometheusRuleRequest() (request *DescribeApmPrometheusRuleRequest) {
    request = &DescribeApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmPrometheusRule")
    
    
    return
}

func NewDescribeApmPrometheusRuleResponse() (response *DescribeApmPrometheusRuleResponse) {
    response = &DescribeApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmPrometheusRule
// This API is used to query the match rule for metrics between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmPrometheusRule(request *DescribeApmPrometheusRuleRequest) (response *DescribeApmPrometheusRuleResponse, err error) {
    return c.DescribeApmPrometheusRuleWithContext(context.Background(), request)
}

// DescribeApmPrometheusRule
// This API is used to query the match rule for metrics between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmPrometheusRuleWithContext(ctx context.Context, request *DescribeApmPrometheusRuleRequest) (response *DescribeApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewDescribeApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmSampleConfigRequest() (request *DescribeApmSampleConfigRequest) {
    request = &DescribeApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmSampleConfig")
    
    
    return
}

func NewDescribeApmSampleConfigResponse() (response *DescribeApmSampleConfigResponse) {
    response = &DescribeApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmSampleConfig
// Query sampling configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_SAMPLENOTFOUND = "FailedOperation.SampleNotFound"
func (c *Client) DescribeApmSampleConfig(request *DescribeApmSampleConfigRequest) (response *DescribeApmSampleConfigResponse, err error) {
    return c.DescribeApmSampleConfigWithContext(context.Background(), request)
}

// DescribeApmSampleConfig
// Query sampling configuration
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_SAMPLENOTFOUND = "FailedOperation.SampleNotFound"
func (c *Client) DescribeApmSampleConfigWithContext(ctx context.Context, request *DescribeApmSampleConfigRequest) (response *DescribeApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewDescribeApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmServiceMetricRequest() (request *DescribeApmServiceMetricRequest) {
    request = &DescribeApmServiceMetricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmServiceMetric")
    
    
    return
}

func NewDescribeApmServiceMetricResponse() (response *DescribeApmServiceMetricResponse) {
    response = &DescribeApmServiceMetricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmServiceMetric
// This API is used to obtain the list of APM application metrics.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  FAILEDOPERATION_SERVICENOTMATCHAPPIDERR = "FailedOperation.ServiceNotMatchAppIdErr"
func (c *Client) DescribeApmServiceMetric(request *DescribeApmServiceMetricRequest) (response *DescribeApmServiceMetricResponse, err error) {
    return c.DescribeApmServiceMetricWithContext(context.Background(), request)
}

// DescribeApmServiceMetric
// This API is used to obtain the list of APM application metrics.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"
//  FAILEDOPERATION_SERVICENOTMATCHAPPIDERR = "FailedOperation.ServiceNotMatchAppIdErr"
func (c *Client) DescribeApmServiceMetricWithContext(ctx context.Context, request *DescribeApmServiceMetricRequest) (response *DescribeApmServiceMetricResponse, err error) {
    if request == nil {
        request = NewDescribeApmServiceMetricRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeApmServiceMetric")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmServiceMetric require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmServiceMetricResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralApmApplicationConfigRequest() (request *DescribeGeneralApmApplicationConfigRequest) {
    request = &DescribeGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralApmApplicationConfig")
    
    
    return
}

func NewDescribeGeneralApmApplicationConfigResponse() (response *DescribeGeneralApmApplicationConfigResponse) {
    response = &DescribeGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralApmApplicationConfig
// This API is used to query the application configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfig(request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    return c.DescribeGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// DescribeGeneralApmApplicationConfig
// This API is used to query the application configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfigWithContext(ctx context.Context, request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralMetricDataRequest() (request *DescribeGeneralMetricDataRequest) {
    request = &DescribeGeneralMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralMetricData")
    
    
    return
}

func NewDescribeGeneralMetricDataResponse() (response *DescribeGeneralMetricDataResponse) {
    response = &DescribeGeneralMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralMetricData
// This API is a general API used to obtain metric data. Users submit request parameters as needed and receive the corresponding metric data.
//
// The API call frequency is limited to 20 requests per second and 1200 requests per minute. The number of data points per request is limited to 1440.
//
// **Usage of the General Interface for Fetching Metric Data**
//
// DescribeGeneralMetricData is a general interface for querying metric data, supporting flexible data retrieval. The query method of this interface is similar to using the following SQL statement:
//
// SELECT {Metrics} FROM {ViewName} WHERE {Filters} GROUP BY {GroupBy}.
//
// 1. View (ViewName)
//
// Determines the data domain you want to query.
//
// Examples: service_metric (Service Monitoring View), db_metric (Database View), etc. For views supported by APM, please refer to [Metric Views](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=)
//
// 2. Metrics (Metrics)
//
// Used to specify one or more metric items to be included in the returned results.
//
// Examples: request_count (Total Requests), duration_avg (Average Latency), error_rate (Error Rate). For metrics supported by APM, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 3. Filters (Filters)
//
// Supports one or more filtering conditions in the form of Key-Value pairs.
//
// Example: Querying only a specific service: service.name = "order-service". Common dimensions and specific dimensions supported by each ViewName can be used as keys in filtering conditions. For more details, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 4. Aggregation (GroupBy)
//
// Supports one or more aggregation dimensions, equivalent to the GROUP BY clause in SQL.
//
// Example: Grouping by the interface name operation to view the performance of each interface. Common dimensions and specific dimensions supported by each ViewName can be used as aggregation dimensions. For more details, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 5. Granularity (Period)
//
// This parameter determines whether the data needs to be aggregated by time slices.
//
//     - Period = 1 (Time Series Mode): The returned results are aggregated by time slices. The multiple values contained in the TimeSerial and DataSerial correspond one-to-one, representing the aggregation results for specific time slices. Time Series Mode is primarily used for displaying time trend charts.
//
//     - Period = 0 (Summary Statistics Mode): In the returned results, the DataSerial contains only a single value, representing the summarized data for the entire time range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricData(request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    return c.DescribeGeneralMetricDataWithContext(context.Background(), request)
}

// DescribeGeneralMetricData
// This API is a general API used to obtain metric data. Users submit request parameters as needed and receive the corresponding metric data.
//
// The API call frequency is limited to 20 requests per second and 1200 requests per minute. The number of data points per request is limited to 1440.
//
// **Usage of the General Interface for Fetching Metric Data**
//
// DescribeGeneralMetricData is a general interface for querying metric data, supporting flexible data retrieval. The query method of this interface is similar to using the following SQL statement:
//
// SELECT {Metrics} FROM {ViewName} WHERE {Filters} GROUP BY {GroupBy}.
//
// 1. View (ViewName)
//
// Determines the data domain you want to query.
//
// Examples: service_metric (Service Monitoring View), db_metric (Database View), etc. For views supported by APM, please refer to [Metric Views](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=)
//
// 2. Metrics (Metrics)
//
// Used to specify one or more metric items to be included in the returned results.
//
// Examples: request_count (Total Requests), duration_avg (Average Latency), error_rate (Error Rate). For metrics supported by APM, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 3. Filters (Filters)
//
// Supports one or more filtering conditions in the form of Key-Value pairs.
//
// Example: Querying only a specific service: service.name = "order-service". Common dimensions and specific dimensions supported by each ViewName can be used as keys in filtering conditions. For more details, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 4. Aggregation (GroupBy)
//
// Supports one or more aggregation dimensions, equivalent to the GROUP BY clause in SQL.
//
// Example: Grouping by the interface name operation to view the performance of each interface. Common dimensions and specific dimensions supported by each ViewName can be used as aggregation dimensions. For more details, please refer to the [APM Metric Protocol Standard](https://www.tencentcloud.com/document/product/248/68462?has_map=1&lang=en&pg=).
//
// 5. Granularity (Period)
//
// This parameter determines whether the data needs to be aggregated by time slices.
//
//     - Period = 1 (Time Series Mode): The returned results are aggregated by time slices. The multiple values contained in the TimeSerial and DataSerial correspond one-to-one, representing the aggregation results for specific time slices. Time Series Mode is primarily used for displaying time trend charts.
//
//     - Period = 0 (Summary Statistics Mode): In the returned results, the DataSerial contains only a single value, representing the summarized data for the entire time range.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricDataWithContext(ctx context.Context, request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralOTSpanListRequest() (request *DescribeGeneralOTSpanListRequest) {
    request = &DescribeGeneralOTSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralOTSpanList")
    
    
    return
}

func NewDescribeGeneralOTSpanListResponse() (response *DescribeGeneralOTSpanListResponse) {
    response = &DescribeGeneralOTSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralOTSpanList
// General Query OpenTelemetry Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanList(request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    return c.DescribeGeneralOTSpanListWithContext(context.Background(), request)
}

// DescribeGeneralOTSpanList
// General Query OpenTelemetry Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanListWithContext(ctx context.Context, request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralOTSpanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralOTSpanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralOTSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralOTSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralSpanListRequest() (request *DescribeGeneralSpanListRequest) {
    request = &DescribeGeneralSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanList")
    
    
    return
}

func NewDescribeGeneralSpanListResponse() (response *DescribeGeneralSpanListResponse) {
    response = &DescribeGeneralSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralSpanList
// General Query Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanList(request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    return c.DescribeGeneralSpanListWithContext(context.Background(), request)
}

// DescribeGeneralSpanList
// General Query Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanListWithContext(ctx context.Context, request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralSpanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeGeneralSpanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricRecordsRequest() (request *DescribeMetricRecordsRequest) {
    request = &DescribeMetricRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricRecords")
    
    
    return
}

func NewDescribeMetricRecordsResponse() (response *DescribeMetricRecordsResponse) {
    response = &DescribeMetricRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricRecords
// This API is used to query metric list. To query metrics, it is recommended to use the DescribeGeneralMetricData API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecords(request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    return c.DescribeMetricRecordsWithContext(context.Background(), request)
}

// DescribeMetricRecords
// This API is used to query metric list. To query metrics, it is recommended to use the DescribeGeneralMetricData API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecordsWithContext(ctx context.Context, request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeMetricRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeMetricRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceOverviewRequest() (request *DescribeServiceOverviewRequest) {
    request = &DescribeServiceOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceOverview")
    
    
    return
}

func NewDescribeServiceOverviewResponse() (response *DescribeServiceOverviewResponse) {
    response = &DescribeServiceOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceOverview
// This API is used to pull application overview data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverview(request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    return c.DescribeServiceOverviewWithContext(context.Background(), request)
}

// DescribeServiceOverview
// This API is used to pull application overview data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverviewWithContext(ctx context.Context, request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeServiceOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeServiceOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
    request = &DescribeTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeTagValues")
    
    
    return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
    response = &DescribeTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagValues
// This API is used to query dimensional data by dimension name and filter condition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    return c.DescribeTagValuesWithContext(context.Background(), request)
}

// DescribeTagValues
// This API is used to query dimensional data by dimension name and filter condition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValuesWithContext(ctx context.Context, request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "DescribeTagValues")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmApplicationConfigRequest() (request *ModifyApmApplicationConfigRequest) {
    request = &ModifyApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmApplicationConfig")
    
    
    return
}

func NewModifyApmApplicationConfigResponse() (response *ModifyApmApplicationConfigResponse) {
    response = &ModifyApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmApplicationConfig
// Modify application configurations
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOPERATIONCONFIGINVALID = "FailedOperation.AgentOperationConfigInvalid"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
func (c *Client) ModifyApmApplicationConfig(request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
    return c.ModifyApmApplicationConfigWithContext(context.Background(), request)
}

// ModifyApmApplicationConfig
// Modify application configurations
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTOPERATIONCONFIGINVALID = "FailedOperation.AgentOperationConfigInvalid"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
func (c *Client) ModifyApmApplicationConfigWithContext(ctx context.Context, request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmAssociationRequest() (request *ModifyApmAssociationRequest) {
    request = &ModifyApmAssociationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmAssociation")
    
    
    return
}

func NewModifyApmAssociationResponse() (response *ModifyApmAssociationResponse) {
    response = &ModifyApmAssociationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmAssociation
// This API is used to modify the relationship between the apm Business System and other products, including deletion.
//
// error code that may be returned:
//  FAILEDOPERATION_ASSOCIATIONMODIFYREQUESTNOTVALIDERROR = "FailedOperation.AssociationModifyRequestNotValidError"
//  FAILEDOPERATION_CKAFKADIFFASSOCIATIONERROR = "FailedOperation.CKafkaDiffAssociationError"
//  FAILEDOPERATION_CKAFKAEMPTYTOPICERROR = "FailedOperation.CKafkaEmptyTopicError"
//  FAILEDOPERATION_CKAFKAGETROUTEIDFAILEDERROR = "FailedOperation.CKafkaGetRouteIDFailedError"
//  FAILEDOPERATION_CKAFKAGETROUTETIMEOUTERROR = "FailedOperation.CKafkaGetRouteTimeoutError"
//  FAILEDOPERATION_CKAFKANOTAVAILABLEERROR = "FailedOperation.CKafkaNotAvailableError"
//  FAILEDOPERATION_PEERIDNOTAVAILABLE = "FailedOperation.PeerIdNotAvailable"
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
//  FAILEDOPERATION_PROMINSTANCENOTAVAILABLEERROR = "FailedOperation.PromInstanceNotAvailableError"
//  FAILEDOPERATION_ROUTENOTAVAILABLEERROR = "FailedOperation.RouteNotAvailableError"
//  FAILEDOPERATION_TOPICNOTAVAILABLEERROR = "FailedOperation.TopicNotAvailableError"
func (c *Client) ModifyApmAssociation(request *ModifyApmAssociationRequest) (response *ModifyApmAssociationResponse, err error) {
    return c.ModifyApmAssociationWithContext(context.Background(), request)
}

// ModifyApmAssociation
// This API is used to modify the relationship between the apm Business System and other products, including deletion.
//
// error code that may be returned:
//  FAILEDOPERATION_ASSOCIATIONMODIFYREQUESTNOTVALIDERROR = "FailedOperation.AssociationModifyRequestNotValidError"
//  FAILEDOPERATION_CKAFKADIFFASSOCIATIONERROR = "FailedOperation.CKafkaDiffAssociationError"
//  FAILEDOPERATION_CKAFKAEMPTYTOPICERROR = "FailedOperation.CKafkaEmptyTopicError"
//  FAILEDOPERATION_CKAFKAGETROUTEIDFAILEDERROR = "FailedOperation.CKafkaGetRouteIDFailedError"
//  FAILEDOPERATION_CKAFKAGETROUTETIMEOUTERROR = "FailedOperation.CKafkaGetRouteTimeoutError"
//  FAILEDOPERATION_CKAFKANOTAVAILABLEERROR = "FailedOperation.CKafkaNotAvailableError"
//  FAILEDOPERATION_PEERIDNOTAVAILABLE = "FailedOperation.PeerIdNotAvailable"
//  FAILEDOPERATION_PRODUCTNAMENOTAVAILABLE = "FailedOperation.ProductNameNotAvailable"
//  FAILEDOPERATION_PROMINSTANCENOTAVAILABLEERROR = "FailedOperation.PromInstanceNotAvailableError"
//  FAILEDOPERATION_ROUTENOTAVAILABLEERROR = "FailedOperation.RouteNotAvailableError"
//  FAILEDOPERATION_TOPICNOTAVAILABLEERROR = "FailedOperation.TopicNotAvailableError"
func (c *Client) ModifyApmAssociationWithContext(ctx context.Context, request *ModifyApmAssociationRequest) (response *ModifyApmAssociationResponse, err error) {
    if request == nil {
        request = NewModifyApmAssociationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmAssociation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmAssociation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmAssociationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmInstanceRequest() (request *ModifyApmInstanceRequest) {
    request = &ModifyApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmInstance")
    
    
    return
}

func NewModifyApmInstanceResponse() (response *ModifyApmInstanceResponse) {
    response = &ModifyApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmInstance
// This API is used to modify the APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstance(request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    return c.ModifyApmInstanceWithContext(context.Background(), request)
}

// ModifyApmInstance
// This API is used to modify the APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstanceWithContext(ctx context.Context, request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    if request == nil {
        request = NewModifyApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmPrometheusRuleRequest() (request *ModifyApmPrometheusRuleRequest) {
    request = &ModifyApmPrometheusRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmPrometheusRule")
    
    
    return
}

func NewModifyApmPrometheusRuleResponse() (response *ModifyApmPrometheusRuleResponse) {
    response = &ModifyApmPrometheusRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmPrometheusRule
// This API is used to modify metric match rules between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) ModifyApmPrometheusRule(request *ModifyApmPrometheusRuleRequest) (response *ModifyApmPrometheusRuleResponse, err error) {
    return c.ModifyApmPrometheusRuleWithContext(context.Background(), request)
}

// ModifyApmPrometheusRule
// This API is used to modify metric match rules between apm Business System and Prometheus Instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROMRULECONFLICT = "FailedOperation.PromRuleConflict"
//  FAILEDOPERATION_PROMRULEISEMPTYERR = "FailedOperation.PromRuleIsEmptyErr"
//  FAILEDOPERATION_PROMRULEREQUESTNOTVALIDERROR = "FailedOperation.PromRuleRequestNotValidError"
func (c *Client) ModifyApmPrometheusRuleWithContext(ctx context.Context, request *ModifyApmPrometheusRuleRequest) (response *ModifyApmPrometheusRuleResponse, err error) {
    if request == nil {
        request = NewModifyApmPrometheusRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmPrometheusRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmPrometheusRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmPrometheusRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmSampleConfigRequest() (request *ModifyApmSampleConfigRequest) {
    request = &ModifyApmSampleConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmSampleConfig")
    
    
    return
}

func NewModifyApmSampleConfigResponse() (response *ModifyApmSampleConfigResponse) {
    response = &ModifyApmSampleConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmSampleConfig
// Modify sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) ModifyApmSampleConfig(request *ModifyApmSampleConfigRequest) (response *ModifyApmSampleConfigResponse, err error) {
    return c.ModifyApmSampleConfigWithContext(context.Background(), request)
}

// ModifyApmSampleConfig
// Modify sampling configurations
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDOPERATIONTYPE = "FailedOperation.InvalidOperationType"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_OPERATIONNAMEISEMPTY = "FailedOperation.OperationNameIsEmpty"
//  FAILEDOPERATION_SAMPLERULECONFLICT = "FailedOperation.SampleRuleConflict"
func (c *Client) ModifyApmSampleConfigWithContext(ctx context.Context, request *ModifyApmSampleConfigRequest) (response *ModifyApmSampleConfigResponse, err error) {
    if request == nil {
        request = NewModifyApmSampleConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyApmSampleConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmSampleConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmSampleConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGeneralApmApplicationConfigRequest() (request *ModifyGeneralApmApplicationConfigRequest) {
    request = &ModifyGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyGeneralApmApplicationConfig")
    
    
    return
}

func NewModifyGeneralApmApplicationConfigResponse() (response *ModifyGeneralApmApplicationConfigResponse) {
    response = &ModifyGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGeneralApmApplicationConfig
// OpenAPI available for external use. Customers can flexibly specify the fields to be modified, and then add the list of services to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfig(request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    return c.ModifyGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// ModifyGeneralApmApplicationConfig
// OpenAPI available for external use. Customers can flexibly specify the fields to be modified, and then add the list of services to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfigWithContext(ctx context.Context, request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyGeneralApmApplicationConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "ModifyGeneralApmApplicationConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateApmInstanceRequest() (request *TerminateApmInstanceRequest) {
    request = &TerminateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "TerminateApmInstance")
    
    
    return
}

func NewTerminateApmInstanceResponse() (response *TerminateApmInstanceResponse) {
    response = &TerminateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateApmInstance
// Termination of APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstance(request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    return c.TerminateApmInstanceWithContext(context.Background(), request)
}

// TerminateApmInstance
// Termination of APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstanceWithContext(ctx context.Context, request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateApmInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apm", APIVersion, "TerminateApmInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateApmInstanceResponse()
    err = c.Send(request, response)
    return
}
