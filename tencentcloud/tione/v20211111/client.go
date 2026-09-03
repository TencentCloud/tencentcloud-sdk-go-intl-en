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

package v20211111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-11-11"

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


func NewCreateTrainingTaskRequest() (request *CreateTrainingTaskRequest) {
    request = &CreateTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingTask")
    
    
    return
}

func NewCreateTrainingTaskResponse() (response *CreateTrainingTaskResponse) {
    response = &CreateTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTrainingTask
// This API is used to create a model training task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATENAMETASKISCREATING = "FailedOperation.DuplicateNameTaskIsCreating"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CHECKFSPATHACCESSIBILITYFAILED = "InternalError.CheckFSPathAccessibilityFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYVPCINFOFAILED = "InternalError.QueryVPCInfoFailed"
//  INTERNALERROR_VALIDATECREATETASKFAILED = "InternalError.ValidateCreateTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETOUTPUTCONFIGEMPTY = "InvalidParameterValue.AIMarketOutputConfigEmpty"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITILLEGAL = "InvalidParameterValue.BackOffLimitIllegal"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITNOTSUPPORT = "InvalidParameterValue.BackOffLimitNotSupport"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FSPATHINACCESSIBLE = "InvalidParameterValue.FSPathInaccessible"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_PARAMLENGTHEXCEEDLIMIT = "InvalidParameterValue.ParamLengthExceedLimit"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_QUERYVPCINFOFAILED = "InvalidParameterValue.QueryVPCInfoFailed"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_TAIJIRESOURCECONFIGILLEGAL = "InvalidParameterValue.TAIJIResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDDATACONFIG = "InvalidParameterValue.UnsupportedDataConfig"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_SUBNETILLEGAL = "OperationDenied.SubnetIllegal"
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTask(request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    return c.CreateTrainingTaskWithContext(context.Background(), request)
}

// CreateTrainingTask
// This API is used to create a model training task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATENAMETASKISCREATING = "FailedOperation.DuplicateNameTaskIsCreating"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CHECKFSPATHACCESSIBILITYFAILED = "InternalError.CheckFSPathAccessibilityFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYVPCINFOFAILED = "InternalError.QueryVPCInfoFailed"
//  INTERNALERROR_VALIDATECREATETASKFAILED = "InternalError.ValidateCreateTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETOUTPUTCONFIGEMPTY = "InvalidParameterValue.AIMarketOutputConfigEmpty"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITILLEGAL = "InvalidParameterValue.BackOffLimitIllegal"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITNOTSUPPORT = "InvalidParameterValue.BackOffLimitNotSupport"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FSPATHINACCESSIBLE = "InvalidParameterValue.FSPathInaccessible"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_PARAMLENGTHEXCEEDLIMIT = "InvalidParameterValue.ParamLengthExceedLimit"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_QUERYVPCINFOFAILED = "InvalidParameterValue.QueryVPCInfoFailed"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_TAIJIRESOURCECONFIGILLEGAL = "InvalidParameterValue.TAIJIResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDDATACONFIG = "InvalidParameterValue.UnsupportedDataConfig"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_SUBNETILLEGAL = "OperationDenied.SubnetIllegal"
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTaskWithContext(ctx context.Context, request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    if request == nil {
        request = NewCreateTrainingTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateTrainingTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceGroupsRequest() (request *DescribeModelServiceGroupsRequest) {
    request = &DescribeModelServiceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceGroups")
    
    
    return
}

func NewDescribeModelServiceGroupsResponse() (response *DescribeModelServiceGroupsResponse) {
    response = &DescribeModelServiceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelServiceGroups
// This API is used to list online inference service groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroups(request *DescribeModelServiceGroupsRequest) (response *DescribeModelServiceGroupsResponse, err error) {
    return c.DescribeModelServiceGroupsWithContext(context.Background(), request)
}

// DescribeModelServiceGroups
// This API is used to list online inference service groups.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroupsWithContext(ctx context.Context, request *DescribeModelServiceGroupsRequest) (response *DescribeModelServiceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelServiceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceGroupsResponse()
    err = c.Send(request, response)
    return
}
