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
