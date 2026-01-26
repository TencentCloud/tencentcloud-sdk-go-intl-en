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

package v20180709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewCreateAllocationRuleRequest() (request *CreateAllocationRuleRequest) {
    request = &CreateAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationRule")
    
    
    return
}

func NewCreateAllocationRuleResponse() (response *CreateAllocationRuleResponse) {
    response = &CreateAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationRule
// Create a sharing rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationRule(request *CreateAllocationRuleRequest) (response *CreateAllocationRuleResponse, err error) {
    return c.CreateAllocationRuleWithContext(context.Background(), request)
}

// CreateAllocationRule
// Create a sharing rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationRuleWithContext(ctx context.Context, request *CreateAllocationRuleRequest) (response *CreateAllocationRuleResponse, err error) {
    if request == nil {
        request = NewCreateAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllocationTagRequest() (request *CreateAllocationTagRequest) {
    request = &CreateAllocationTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationTag")
    
    
    return
}

func NewCreateAllocationTagResponse() (response *CreateAllocationTagResponse) {
    response = &CreateAllocationTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationTag
// This API is used to batch set cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAllocationTag(request *CreateAllocationTagRequest) (response *CreateAllocationTagResponse, err error) {
    return c.CreateAllocationTagWithContext(context.Background(), request)
}

// CreateAllocationTag
// This API is used to batch set cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateAllocationTagWithContext(ctx context.Context, request *CreateAllocationTagRequest) (response *CreateAllocationTagResponse, err error) {
    if request == nil {
        request = NewCreateAllocationTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAllocationUnitRequest() (request *CreateAllocationUnitRequest) {
    request = &CreateAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateAllocationUnit")
    
    
    return
}

func NewCreateAllocationUnitResponse() (response *CreateAllocationUnitResponse) {
    response = &CreateAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAllocationUnit
// This API is used to create allocation units.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationUnit(request *CreateAllocationUnitRequest) (response *CreateAllocationUnitResponse, err error) {
    return c.CreateAllocationUnitWithContext(context.Background(), request)
}

// CreateAllocationUnit
// This API is used to create allocation units.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateAllocationUnitWithContext(ctx context.Context, request *CreateAllocationUnitRequest) (response *CreateAllocationUnitResponse, err error) {
    if request == nil {
        request = NewCreateAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGatherRuleRequest() (request *CreateGatherRuleRequest) {
    request = &CreateGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateGatherRule")
    
    
    return
}

func NewCreateGatherRuleResponse() (response *CreateGatherRuleResponse) {
    response = &CreateGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGatherRule
// Create a collection rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateGatherRule(request *CreateGatherRuleRequest) (response *CreateGatherRuleResponse, err error) {
    return c.CreateGatherRuleWithContext(context.Background(), request)
}

// CreateGatherRule
// Create a collection rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateGatherRuleWithContext(ctx context.Context, request *CreateGatherRuleRequest) (response *CreateGatherRuleResponse, err error) {
    if request == nil {
        request = NewCreateGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstance
// Creating an instance resource will generate an order for the newly purchased instance resource and automatically complete the payment using the balance of the Tencent Cloud account. The account calling this API must be granted the finace:trade permission; otherwise, the payment will fail.
//
// Currently, the integrated and supported product for purchase includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_DISTRIBUTEERROR = "FailedOperation.DistributeError"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDGOODSCATEGORYID = "FailedOperation.InvalidGoodsCategoryId"
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  FAILEDOPERATION_DEALCREATEWHITELISTERROR = "FailedOperation.dealCreateWhitelistError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// Creating an instance resource will generate an order for the newly purchased instance resource and automatically complete the payment using the balance of the Tencent Cloud account. The account calling this API must be granted the finace:trade permission; otherwise, the payment will fail.
//
// Currently, the integrated and supported product for purchase includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_DISTRIBUTEERROR = "FailedOperation.DistributeError"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDGOODSCATEGORYID = "FailedOperation.InvalidGoodsCategoryId"
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  FAILEDOPERATION_DEALCREATEWHITELISTERROR = "FailedOperation.dealCreateWhitelistError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "CreateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationRuleRequest() (request *DeleteAllocationRuleRequest) {
    request = &DeleteAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationRule")
    
    
    return
}

func NewDeleteAllocationRuleResponse() (response *DeleteAllocationRuleResponse) {
    response = &DeleteAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationRule
// Delete sharing rule interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteAllocationRule(request *DeleteAllocationRuleRequest) (response *DeleteAllocationRuleResponse, err error) {
    return c.DeleteAllocationRuleWithContext(context.Background(), request)
}

// DeleteAllocationRule
// Delete sharing rule interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteAllocationRuleWithContext(ctx context.Context, request *DeleteAllocationRuleRequest) (response *DeleteAllocationRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationTagRequest() (request *DeleteAllocationTagRequest) {
    request = &DeleteAllocationTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationTag")
    
    
    return
}

func NewDeleteAllocationTagResponse() (response *DeleteAllocationTagResponse) {
    response = &DeleteAllocationTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationTag
// u200cThis API is used to batch cancel cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAllocationTag(request *DeleteAllocationTagRequest) (response *DeleteAllocationTagResponse, err error) {
    return c.DeleteAllocationTagWithContext(context.Background(), request)
}

// DeleteAllocationTag
// u200cThis API is used to batch cancel cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAllocationTagWithContext(ctx context.Context, request *DeleteAllocationTagRequest) (response *DeleteAllocationTagResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAllocationUnitRequest() (request *DeleteAllocationUnitRequest) {
    request = &DeleteAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteAllocationUnit")
    
    
    return
}

func NewDeleteAllocationUnitResponse() (response *DeleteAllocationUnitResponse) {
    response = &DeleteAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAllocationUnit
// Delete a cost allocation unit.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAllocationUnit(request *DeleteAllocationUnitRequest) (response *DeleteAllocationUnitResponse, err error) {
    return c.DeleteAllocationUnitWithContext(context.Background(), request)
}

// DeleteAllocationUnit
// Delete a cost allocation unit.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAllocationUnitWithContext(ctx context.Context, request *DeleteAllocationUnitRequest) (response *DeleteAllocationUnitResponse, err error) {
    if request == nil {
        request = NewDeleteAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGatherRuleRequest() (request *DeleteGatherRuleRequest) {
    request = &DeleteGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DeleteGatherRule")
    
    
    return
}

func NewDeleteGatherRuleResponse() (response *DeleteGatherRuleResponse) {
    response = &DeleteGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGatherRule
// Delete a collection rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteGatherRule(request *DeleteGatherRuleRequest) (response *DeleteGatherRuleResponse, err error) {
    return c.DeleteGatherRuleWithContext(context.Background(), request)
}

// DeleteGatherRule
// Delete a collection rule.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DeleteGatherRuleWithContext(ctx context.Context, request *DeleteGatherRuleRequest) (response *DeleteGatherRuleResponse, err error) {
    if request == nil {
        request = NewDeleteGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DeleteGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountBalanceRequest() (request *DescribeAccountBalanceRequest) {
    request = &DescribeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAccountBalance")
    
    
    return
}

func NewDescribeAccountBalanceResponse() (response *DescribeAccountBalanceResponse) {
    response = &DescribeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountBalance
// This API is used to check the Tencent Cloud account balance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    return c.DescribeAccountBalanceWithContext(context.Background(), request)
}

// DescribeAccountBalance
// This API is used to check the Tencent Cloud account balance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalanceWithContext(ctx context.Context, request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAccountBalance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationRuleDetailRequest() (request *DescribeAllocationRuleDetailRequest) {
    request = &DescribeAllocationRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationRuleDetail")
    
    
    return
}

func NewDescribeAllocationRuleDetailResponse() (response *DescribeAllocationRuleDetailResponse) {
    response = &DescribeAllocationRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationRuleDetail
// This API is used to query sharing rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleDetail(request *DescribeAllocationRuleDetailRequest) (response *DescribeAllocationRuleDetailResponse, err error) {
    return c.DescribeAllocationRuleDetailWithContext(context.Background(), request)
}

// DescribeAllocationRuleDetail
// This API is used to query sharing rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleDetailWithContext(ctx context.Context, request *DescribeAllocationRuleDetailRequest) (response *DescribeAllocationRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationRuleSummaryRequest() (request *DescribeAllocationRuleSummaryRequest) {
    request = &DescribeAllocationRuleSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationRuleSummary")
    
    
    return
}

func NewDescribeAllocationRuleSummaryResponse() (response *DescribeAllocationRuleSummaryResponse) {
    response = &DescribeAllocationRuleSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationRuleSummary
// This API is used to query all sharing rule overviews.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleSummary(request *DescribeAllocationRuleSummaryRequest) (response *DescribeAllocationRuleSummaryResponse, err error) {
    return c.DescribeAllocationRuleSummaryWithContext(context.Background(), request)
}

// DescribeAllocationRuleSummary
// This API is used to query all sharing rule overviews.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationRuleSummaryWithContext(ctx context.Context, request *DescribeAllocationRuleSummaryRequest) (response *DescribeAllocationRuleSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationRuleSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationRuleSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationRuleSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationRuleSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationTreeRequest() (request *DescribeAllocationTreeRequest) {
    request = &DescribeAllocationTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationTree")
    
    
    return
}

func NewDescribeAllocationTreeResponse() (response *DescribeAllocationTreeResponse) {
    response = &DescribeAllocationTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationTree
// This API is used to query the cost tree.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTree(request *DescribeAllocationTreeRequest) (response *DescribeAllocationTreeResponse, err error) {
    return c.DescribeAllocationTreeWithContext(context.Background(), request)
}

// DescribeAllocationTree
// This API is used to query the cost tree.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationTreeWithContext(ctx context.Context, request *DescribeAllocationTreeRequest) (response *DescribeAllocationTreeResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationTreeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationTree")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllocationUnitDetailRequest() (request *DescribeAllocationUnitDetailRequest) {
    request = &DescribeAllocationUnitDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAllocationUnitDetail")
    
    
    return
}

func NewDescribeAllocationUnitDetailResponse() (response *DescribeAllocationUnitDetailResponse) {
    response = &DescribeAllocationUnitDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllocationUnitDetail
// Query the details of a cost allocation unit.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationUnitDetail(request *DescribeAllocationUnitDetailRequest) (response *DescribeAllocationUnitDetailResponse, err error) {
    return c.DescribeAllocationUnitDetailWithContext(context.Background(), request)
}

// DescribeAllocationUnitDetail
// Query the details of a cost allocation unit.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAllocationUnitDetailWithContext(ctx context.Context, request *DescribeAllocationUnitDetailRequest) (response *DescribeAllocationUnitDetailResponse, err error) {
    if request == nil {
        request = NewDescribeAllocationUnitDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeAllocationUnitDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllocationUnitDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllocationUnitDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillAdjustInfoRequest() (request *DescribeBillAdjustInfoRequest) {
    request = &DescribeBillAdjustInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillAdjustInfo")
    
    
    return
}

func NewDescribeBillAdjustInfoResponse() (response *DescribeBillAdjustInfoResponse) {
    response = &DescribeBillAdjustInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillAdjustInfo
// This API is used to check whether the current UIN has any adjustment, enabling customers to proactively obtain the adjustment status faster.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeBillAdjustInfo(request *DescribeBillAdjustInfoRequest) (response *DescribeBillAdjustInfoResponse, err error) {
    return c.DescribeBillAdjustInfoWithContext(context.Background(), request)
}

// DescribeBillAdjustInfo
// This API is used to check whether the current UIN has any adjustment, enabling customers to proactively obtain the adjustment status faster.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeBillAdjustInfoWithContext(ctx context.Context, request *DescribeBillAdjustInfoRequest) (response *DescribeBillAdjustInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBillAdjustInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillAdjustInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillAdjustInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillAdjustInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetail
// u200cThis API is used to get bill details.
//
// Note:
//
// 1. The API request may fail due to network instability or other exceptions. In this case, we recommend you manually retry the request when the API request fails.
//
// 2.If the volume of your bill data is high (for example, if over 200 thousand bill entries are generated for a month), bill data query via APIs may be slow. We recommend you enable bill storage so that you can obtain bill files from COS buckets for analysis. For details, see [Saving Bills to COS](https://intl.cloud.tencent.com/document/product/555/61275?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// u200cThis API is used to get bill details.
//
// Note:
//
// 1. The API request may fail due to network instability or other exceptions. In this case, we recommend you manually retry the request when the API request fails.
//
// 2.If the volume of your bill data is high (for example, if over 200 thousand bill entries are generated for a month), bill data query via APIs may be slow. We recommend you enable bill storage so that you can obtain bill files from COS buckets for analysis. For details, see [Saving Bills to COS](https://intl.cloud.tencent.com/document/product/555/61275?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailForOrganizationRequest() (request *DescribeBillDetailForOrganizationRequest) {
    request = &DescribeBillDetailForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetailForOrganization")
    
    
    return
}

func NewDescribeBillDetailForOrganizationResponse() (response *DescribeBillDetailForOrganizationResponse) {
    response = &DescribeBillDetailForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetailForOrganization
// This API is used to get pay-on-behalf bills of the admin account (bill details).
//
// Note: The API request may fail due to network instability or other exceptions. In this case, we recommend you manually retry the request when the API request fails.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailForOrganization(request *DescribeBillDetailForOrganizationRequest) (response *DescribeBillDetailForOrganizationResponse, err error) {
    return c.DescribeBillDetailForOrganizationWithContext(context.Background(), request)
}

// DescribeBillDetailForOrganization
// This API is used to get pay-on-behalf bills of the admin account (bill details).
//
// Note: The API request may fail due to network instability or other exceptions. In this case, we recommend you manually retry the request when the API request fails.
//
// error code that may be returned:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailForOrganizationWithContext(ctx context.Context, request *DescribeBillDetailForOrganizationRequest) (response *DescribeBillDetailForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDetailForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetailForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDownloadUrlRequest() (request *DescribeBillDownloadUrlRequest) {
    request = &DescribeBillDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDownloadUrl")
    
    
    return
}

func NewDescribeBillDownloadUrlResponse() (response *DescribeBillDownloadUrlResponse) {
    response = &DescribeBillDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDownloadUrl
// This API is used to get bill download URLs for L0, L1, L2, and L3 bills and bill packs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrl(request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    return c.DescribeBillDownloadUrlWithContext(context.Background(), request)
}

// DescribeBillDownloadUrl
// This API is used to get bill download URLs for L0, L1, L2, and L3 bills and bill packs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillDownloadUrlWithContext(ctx context.Context, request *DescribeBillDownloadUrlRequest) (response *DescribeBillDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBillDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryRequest() (request *DescribeBillResourceSummaryRequest) {
    request = &DescribeBillResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummary")
    
    
    return
}

func NewDescribeBillResourceSummaryResponse() (response *DescribeBillResourceSummaryResponse) {
    response = &DescribeBillResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillResourceSummary
// This API is used to get the bill summarized by instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummary(request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    return c.DescribeBillResourceSummaryWithContext(context.Background(), request)
}

// DescribeBillResourceSummary
// This API is used to get the bill summarized by instance.
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryWithContext(ctx context.Context, request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillResourceSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryForOrganizationRequest() (request *DescribeBillResourceSummaryForOrganizationRequest) {
    request = &DescribeBillResourceSummaryForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummaryForOrganization")
    
    
    return
}

func NewDescribeBillResourceSummaryForOrganizationResponse() (response *DescribeBillResourceSummaryForOrganizationResponse) {
    response = &DescribeBillResourceSummaryForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillResourceSummaryForOrganization
// This API is used to get pay-on-behalf bills of the admin account (bills by instance).
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryForOrganization(request *DescribeBillResourceSummaryForOrganizationRequest) (response *DescribeBillResourceSummaryForOrganizationResponse, err error) {
    return c.DescribeBillResourceSummaryForOrganizationWithContext(context.Background(), request)
}

// DescribeBillResourceSummaryForOrganization
// This API is used to get pay-on-behalf bills of the admin account (bills by instance).
//
// error code that may be returned:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBillResourceSummaryForOrganizationWithContext(ctx context.Context, request *DescribeBillResourceSummaryForOrganizationRequest) (response *DescribeBillResourceSummaryForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillResourceSummaryForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummaryForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryRequest() (request *DescribeBillSummaryRequest) {
    request = &DescribeBillSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummary")
    
    
    return
}

func NewDescribeBillSummaryResponse() (response *DescribeBillSummaryResponse) {
    response = &DescribeBillSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummary
// This API is used to get bill details by product, project, region, billing mode, and tag by passing in parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummary(request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    return c.DescribeBillSummaryWithContext(context.Background(), request)
}

// DescribeBillSummary
// This API is used to get bill details by product, project, region, billing mode, and tag by passing in parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryWithContext(ctx context.Context, request *DescribeBillSummaryRequest) (response *DescribeBillSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByPayMode
// This API is used to get the bill summarized by billing mode.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// This API is used to get the bill summarized by billing mode.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProduct
// Gets the bill summarized according to product
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// Gets the bill summarized according to product
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProject")
    
    
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProject
// Gets the bill summarized according to project
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

// DescribeBillSummaryByProject
// Gets the bill summarized according to project
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByRegion
// Gets the bill summarized according to region
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// Gets the bill summarized according to region
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
    request = &DescribeBillSummaryByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByTag")
    
    
    return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
    response = &DescribeBillSummaryByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByTag
// This API is used to get the cost distribution over different tags.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    return c.DescribeBillSummaryByTagWithContext(context.Background(), request)
}

// DescribeBillSummaryByTag
// This API is used to get the cost distribution over different tags.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTagWithContext(ctx context.Context, request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByTagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryByTag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryForOrganizationRequest() (request *DescribeBillSummaryForOrganizationRequest) {
    request = &DescribeBillSummaryForOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryForOrganization")
    
    
    return
}

func NewDescribeBillSummaryForOrganizationResponse() (response *DescribeBillSummaryForOrganizationResponse) {
    response = &DescribeBillSummaryForOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryForOrganization
// This API is used to get bills summarized by product, project, region, billing mode, and tag by passing in parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryForOrganization(request *DescribeBillSummaryForOrganizationRequest) (response *DescribeBillSummaryForOrganizationResponse, err error) {
    return c.DescribeBillSummaryForOrganizationWithContext(context.Background(), request)
}

// DescribeBillSummaryForOrganization
// This API is used to get bills summarized by product, project, region, billing mode, and tag by passing in parameters.
//
// error code that may be returned:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryForOrganizationWithContext(ctx context.Context, request *DescribeBillSummaryForOrganizationRequest) (response *DescribeBillSummaryForOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryForOrganizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeBillSummaryForOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryForOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryForOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostDetailRequest() (request *DescribeCostDetailRequest) {
    request = &DescribeCostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostDetail")
    
    
    return
}

func NewDescribeCostDetailResponse() (response *DescribeCostDetailResponse) {
    response = &DescribeCostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostDetail
// This API is used to query consumption details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostDetail(request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    return c.DescribeCostDetailWithContext(context.Background(), request)
}

// DescribeCostDetail
// This API is used to query consumption details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostDetailWithContext(ctx context.Context, request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCostDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostExplorerSummaryRequest() (request *DescribeCostExplorerSummaryRequest) {
    request = &DescribeCostExplorerSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostExplorerSummary")
    
    
    return
}

func NewDescribeCostExplorerSummaryResponse() (response *DescribeCostExplorerSummaryResponse) {
    response = &DescribeCostExplorerSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostExplorerSummary
// This API is used to view cost analysis details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostExplorerSummary(request *DescribeCostExplorerSummaryRequest) (response *DescribeCostExplorerSummaryResponse, err error) {
    return c.DescribeCostExplorerSummaryWithContext(context.Background(), request)
}

// DescribeCostExplorerSummary
// This API is used to view cost analysis details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCostExplorerSummaryWithContext(ctx context.Context, request *DescribeCostExplorerSummaryRequest) (response *DescribeCostExplorerSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCostExplorerSummaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostExplorerSummary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostExplorerSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostExplorerSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProductRequest() (request *DescribeCostSummaryByProductRequest) {
    request = &DescribeCostSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProduct")
    
    
    return
}

func NewDescribeCostSummaryByProductResponse() (response *DescribeCostSummaryByProductResponse) {
    response = &DescribeCostSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByProduct
// This API is used to obtain consumption details summarized by product.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProduct(request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    return c.DescribeCostSummaryByProductWithContext(context.Background(), request)
}

// DescribeCostSummaryByProduct
// This API is used to obtain consumption details summarized by product.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProductWithContext(ctx context.Context, request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProductRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByProduct")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProjectRequest() (request *DescribeCostSummaryByProjectRequest) {
    request = &DescribeCostSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProject")
    
    
    return
}

func NewDescribeCostSummaryByProjectResponse() (response *DescribeCostSummaryByProjectResponse) {
    response = &DescribeCostSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByProject
// This API is used to obtain consumption details summarized by project.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProject(request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    return c.DescribeCostSummaryByProjectWithContext(context.Background(), request)
}

// DescribeCostSummaryByProject
// This API is used to obtain consumption details summarized by project.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProjectWithContext(ctx context.Context, request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByRegionRequest() (request *DescribeCostSummaryByRegionRequest) {
    request = &DescribeCostSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByRegion")
    
    
    return
}

func NewDescribeCostSummaryByRegionResponse() (response *DescribeCostSummaryByRegionResponse) {
    response = &DescribeCostSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByRegion
// This API is used to obtain consumption details summarized by region.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegion(request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    return c.DescribeCostSummaryByRegionWithContext(context.Background(), request)
}

// DescribeCostSummaryByRegion
// This API is used to obtain consumption details summarized by region.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegionWithContext(ctx context.Context, request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByRegionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByRegion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByResourceRequest() (request *DescribeCostSummaryByResourceRequest) {
    request = &DescribeCostSummaryByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByResource")
    
    
    return
}

func NewDescribeCostSummaryByResourceResponse() (response *DescribeCostSummaryByResourceResponse) {
    response = &DescribeCostSummaryByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCostSummaryByResource
// This API is used to obtain consumption details summarized by resource.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResource(request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    return c.DescribeCostSummaryByResourceWithContext(context.Background(), request)
}

// DescribeCostSummaryByResource
// This API is used to obtain consumption details summarized by resource.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResourceWithContext(ctx context.Context, request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeCostSummaryByResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealsByCondRequest() (request *DescribeDealsByCondRequest) {
    request = &DescribeDealsByCondRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDealsByCond")
    
    
    return
}

func NewDescribeDealsByCondResponse() (response *DescribeDealsByCondResponse) {
    response = &DescribeDealsByCondResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDealsByCond
// Querying orders
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCond(request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    return c.DescribeDealsByCondWithContext(context.Background(), request)
}

// DescribeDealsByCond
// Querying orders
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCondWithContext(ctx context.Context, request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    if request == nil {
        request = NewDescribeDealsByCondRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDealsByCond")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDealsByCond require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDealsByCondResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageCosDetailByDateRequest() (request *DescribeDosageCosDetailByDateRequest) {
    request = &DescribeDosageCosDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    
    return
}

func NewDescribeDosageCosDetailByDateResponse() (response *DescribeDosageCosDetailByDateResponse) {
    response = &DescribeDosageCosDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDosageCosDetailByDate
// This API is used to query COS usage details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDate(request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    return c.DescribeDosageCosDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageCosDetailByDate
// This API is used to query COS usage details.
//
// error code that may be returned:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDateWithContext(ctx context.Context, request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageCosDetailByDateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageCosDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageCosDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGatherRuleDetailRequest() (request *DescribeGatherRuleDetailRequest) {
    request = &DescribeGatherRuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeGatherRuleDetail")
    
    
    return
}

func NewDescribeGatherRuleDetailResponse() (response *DescribeGatherRuleDetailResponse) {
    response = &DescribeGatherRuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGatherRuleDetail
// This API is used to query the collection rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeGatherRuleDetail(request *DescribeGatherRuleDetailRequest) (response *DescribeGatherRuleDetailResponse, err error) {
    return c.DescribeGatherRuleDetailWithContext(context.Background(), request)
}

// DescribeGatherRuleDetail
// This API is used to query the collection rule details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeGatherRuleDetailWithContext(ctx context.Context, request *DescribeGatherRuleDetailRequest) (response *DescribeGatherRuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGatherRuleDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeGatherRuleDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGatherRuleDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGatherRuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagListRequest() (request *DescribeTagListRequest) {
    request = &DescribeTagListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeTagList")
    
    
    return
}

func NewDescribeTagListResponse() (response *DescribeTagListResponse) {
    response = &DescribeTagListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagList
// This API is used to get cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTagList(request *DescribeTagListRequest) (response *DescribeTagListResponse, err error) {
    return c.DescribeTagListWithContext(context.Background(), request)
}

// DescribeTagList
// This API is used to get cost allocation tags.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTagListWithContext(ctx context.Context, request *DescribeTagListRequest) (response *DescribeTagListResponse, err error) {
    if request == nil {
        request = NewDescribeTagListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeTagList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherInfoRequest() (request *DescribeVoucherInfoRequest) {
    request = &DescribeVoucherInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherInfo")
    
    
    return
}

func NewDescribeVoucherInfoResponse() (response *DescribeVoucherInfoResponse) {
    response = &DescribeVoucherInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoucherInfo
// This API is used to query vouchers.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfo(request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    return c.DescribeVoucherInfoWithContext(context.Background(), request)
}

// DescribeVoucherInfo
// This API is used to query vouchers.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfoWithContext(ctx context.Context, request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeVoucherInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherUsageDetailsRequest() (request *DescribeVoucherUsageDetailsRequest) {
    request = &DescribeVoucherUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherUsageDetails")
    
    
    return
}

func NewDescribeVoucherUsageDetailsResponse() (response *DescribeVoucherUsageDetailsResponse) {
    response = &DescribeVoucherUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoucherUsageDetails
// This API is used to query voucher usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetails(request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    return c.DescribeVoucherUsageDetailsWithContext(context.Background(), request)
}

// DescribeVoucherUsageDetails
// This API is used to query voucher usage details.
//
// error code that may be returned:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetailsWithContext(ctx context.Context, request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherUsageDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "DescribeVoucherUsageDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherUsageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherUsageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllocationRuleRequest() (request *ModifyAllocationRuleRequest) {
    request = &ModifyAllocationRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyAllocationRule")
    
    
    return
}

func NewModifyAllocationRuleResponse() (response *ModifyAllocationRuleResponse) {
    response = &ModifyAllocationRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllocationRule
// Edit sharing rules.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAllocationRule(request *ModifyAllocationRuleRequest) (response *ModifyAllocationRuleResponse, err error) {
    return c.ModifyAllocationRuleWithContext(context.Background(), request)
}

// ModifyAllocationRule
// Edit sharing rules.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAllocationRuleWithContext(ctx context.Context, request *ModifyAllocationRuleRequest) (response *ModifyAllocationRuleResponse, err error) {
    if request == nil {
        request = NewModifyAllocationRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyAllocationRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllocationRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllocationRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAllocationUnitRequest() (request *ModifyAllocationUnitRequest) {
    request = &ModifyAllocationUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyAllocationUnit")
    
    
    return
}

func NewModifyAllocationUnitResponse() (response *ModifyAllocationUnitResponse) {
    response = &ModifyAllocationUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAllocationUnit
// This API is used to modify cost allocation unit information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAllocationUnit(request *ModifyAllocationUnitRequest) (response *ModifyAllocationUnitResponse, err error) {
    return c.ModifyAllocationUnitWithContext(context.Background(), request)
}

// ModifyAllocationUnit
// This API is used to modify cost allocation unit information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAllocationUnitWithContext(ctx context.Context, request *ModifyAllocationUnitRequest) (response *ModifyAllocationUnitResponse, err error) {
    if request == nil {
        request = NewModifyAllocationUnitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyAllocationUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAllocationUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAllocationUnitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGatherRuleRequest() (request *ModifyGatherRuleRequest) {
    request = &ModifyGatherRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "ModifyGatherRule")
    
    
    return
}

func NewModifyGatherRuleResponse() (response *ModifyGatherRuleResponse) {
    response = &ModifyGatherRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGatherRule
// Edit a collection rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyGatherRule(request *ModifyGatherRuleRequest) (response *ModifyGatherRuleResponse, err error) {
    return c.ModifyGatherRuleWithContext(context.Background(), request)
}

// ModifyGatherRule
// Edit a collection rule.
//
// error code that may be returned:
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATERROR = "InternalError.DbOperatError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyGatherRuleWithContext(ctx context.Context, request *ModifyGatherRuleRequest) (response *ModifyGatherRuleResponse, err error) {
    if request == nil {
        request = NewModifyGatherRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "ModifyGatherRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGatherRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGatherRuleResponse()
    err = c.Send(request, response)
    return
}

func NewPayDealsRequest() (request *PayDealsRequest) {
    request = &PayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "PayDeals")
    
    
    return
}

func NewPayDealsResponse() (response *PayDealsResponse) {
    response = &PayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PayDeals
// This API is used to pay for an order.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
    return c.PayDealsWithContext(context.Background(), request)
}

// PayDeals
// This API is used to pay for an order.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDealsWithContext(ctx context.Context, request *PayDealsRequest) (response *PayDealsResponse, err error) {
    if request == nil {
        request = NewPayDealsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "PayDeals")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewPayDealsResponse()
    err = c.Send(request, response)
    return
}

func NewRefundInstanceRequest() (request *RefundInstanceRequest) {
    request = &RefundInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "RefundInstance")
    
    
    return
}

func NewRefundInstanceResponse() (response *RefundInstanceResponse) {
    response = &RefundInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefundInstance
// To unsubscribe from an unneeded instance, only the actual payment amount will be refunded, any used vouchers will not be returned. The refunded amount will be credited to your Tencent Cloud account balance by default.The account calling this API must be granted the finace:RefundInstance permission; otherwise, the refund process will fail.
//
// Currently, the integrated and supported product for this operation includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_RESOURCELOCKED = "InvalidParameter.ResourceLocked"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) RefundInstance(request *RefundInstanceRequest) (response *RefundInstanceResponse, err error) {
    return c.RefundInstanceWithContext(context.Background(), request)
}

// RefundInstance
// To unsubscribe from an unneeded instance, only the actual payment amount will be refunded, any used vouchers will not be returned. The refunded amount will be credited to your Tencent Cloud account balance by default.The account calling this API must be granted the finace:RefundInstance permission; otherwise, the refund process will fail.
//
// Currently, the integrated and supported product for this operation includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_NUMLIMITERROR = "FailedOperation.NumLimitError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_RESOURCELOCKED = "InvalidParameter.ResourceLocked"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) RefundInstanceWithContext(ctx context.Context, request *RefundInstanceRequest) (response *RefundInstanceResponse, err error) {
    if request == nil {
        request = NewRefundInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "RefundInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewInstance
// Renewing an instance: when calling this API to renew a server, ensure that your Tencent Cloud account has sufficient balance; otherwise, the renewal will fail. The account calling this API must be granted the finace:tradepermission; otherwise, the renewal will fail.
//
// Currently, the integrated and supported product for renewal includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_DISTRIBUTEERROR = "FailedOperation.DistributeError"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDGOODSCATEGORYID = "FailedOperation.InvalidGoodsCategoryId"
//  FAILEDOPERATION_DEALCREATEWHITELISTERROR = "FailedOperation.dealCreateWhitelistError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_RESOURCELOCKED = "InvalidParameter.ResourceLocked"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// Renewing an instance: when calling this API to renew a server, ensure that your Tencent Cloud account has sufficient balance; otherwise, the renewal will fail. The account calling this API must be granted the finace:tradepermission; otherwise, the renewal will fail.
//
// Currently, the integrated and supported product for renewal includes: Cloud Firewall.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_BUSINESSCHECKERRCODE = "FailedOperation.BusinessCheckErrCode"
//  FAILEDOPERATION_DISTRIBUTEERROR = "FailedOperation.DistributeError"
//  FAILEDOPERATION_GETPRICEPARAMERROR = "FailedOperation.GetPriceParamError"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDGOODSCATEGORYID = "FailedOperation.InvalidGoodsCategoryId"
//  FAILEDOPERATION_DEALCREATEWHITELISTERROR = "FailedOperation.dealCreateWhitelistError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_RESOURCELOCKED = "InvalidParameter.ResourceLocked"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOT_SUPPORT_THIS_ACTION = "UnsupportedOperation.NOT_SUPPORT_THIS_ACTION"
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "billing", APIVersion, "RenewInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}
