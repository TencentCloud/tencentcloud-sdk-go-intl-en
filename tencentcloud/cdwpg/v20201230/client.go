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

package v20201230

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-30"

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


func NewCreateInstanceByApiRequest() (request *CreateInstanceByApiRequest) {
    request = &CreateInstanceByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "CreateInstanceByApi")
    
    
    return
}

func NewCreateInstanceByApiResponse() (response *CreateInstanceByApiResponse) {
    response = &CreateInstanceByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceByApi
// This API is used to create  instance
func (c *Client) CreateInstanceByApi(request *CreateInstanceByApiRequest) (response *CreateInstanceByApiResponse, err error) {
    return c.CreateInstanceByApiWithContext(context.Background(), request)
}

// CreateInstanceByApi
// This API is used to create  instance
func (c *Client) CreateInstanceByApiWithContext(ctx context.Context, request *CreateInstanceByApiRequest) (response *CreateInstanceByApiResponse, err error) {
    if request == nil {
        request = NewCreateInstanceByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceByApiResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccounts
// This API is used to obtain the account list corresponding to the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// This API is used to obtain the account list corresponding to the instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBConfigHistoryRequest() (request *DescribeDBConfigHistoryRequest) {
    request = &DescribeDBConfigHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeDBConfigHistory")
    
    
    return
}

func NewDescribeDBConfigHistoryResponse() (response *DescribeDBConfigHistoryResponse) {
    response = &DescribeDBConfigHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBConfigHistory
// Describe DBConfig History
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDBConfigHistory(request *DescribeDBConfigHistoryRequest) (response *DescribeDBConfigHistoryResponse, err error) {
    return c.DescribeDBConfigHistoryWithContext(context.Background(), request)
}

// DescribeDBConfigHistory
// Describe DBConfig History
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDBConfigHistoryWithContext(ctx context.Context, request *DescribeDBConfigHistoryRequest) (response *DescribeDBConfigHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDBConfigHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBConfigHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBConfigHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParamsRequest() (request *DescribeDBParamsRequest) {
    request = &DescribeDBParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeDBParams")
    
    
    return
}

func NewDescribeDBParamsResponse() (response *DescribeDBParamsResponse) {
    response = &DescribeDBParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBParams
// This API is used to describe instance configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDBParams(request *DescribeDBParamsRequest) (response *DescribeDBParamsResponse, err error) {
    return c.DescribeDBParamsWithContext(context.Background(), request)
}

// DescribeDBParams
// This API is used to describe instance configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDBParamsWithContext(ctx context.Context, request *DescribeDBParamsRequest) (response *DescribeDBParamsResponse, err error) {
    if request == nil {
        request = NewDescribeDBParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorLogRequest() (request *DescribeErrorLogRequest) {
    request = &DescribeErrorLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeErrorLog")
    
    
    return
}

func NewDescribeErrorLogResponse() (response *DescribeErrorLogResponse) {
    response = &DescribeErrorLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeErrorLog
// This API is used to query error logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeErrorLog(request *DescribeErrorLogRequest) (response *DescribeErrorLogResponse, err error) {
    return c.DescribeErrorLogWithContext(context.Background(), request)
}

// DescribeErrorLog
// This API is used to query error logs.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeErrorLogWithContext(ctx context.Context, request *DescribeErrorLogRequest) (response *DescribeErrorLogResponse, err error) {
    if request == nil {
        request = NewDescribeErrorLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// This API is used to query the instance information by an instance ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// This API is used to query the instance information by an instance ID.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceInfoRequest() (request *DescribeInstanceInfoRequest) {
    request = &DescribeInstanceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstanceInfo")
    
    
    return
}

func NewDescribeInstanceInfoResponse() (response *DescribeInstanceInfoResponse) {
    response = &DescribeInstanceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceInfo
// This API is used to get instance information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceInfo(request *DescribeInstanceInfoRequest) (response *DescribeInstanceInfoResponse, err error) {
    return c.DescribeInstanceInfoWithContext(context.Background(), request)
}

// DescribeInstanceInfo
// This API is used to get instance information.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceInfoWithContext(ctx context.Context, request *DescribeInstanceInfoRequest) (response *DescribeInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
    request = &DescribeInstanceNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstanceNodes")
    
    
    return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
    response = &DescribeInstanceNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodes
// This API is used to list nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

// DescribeInstanceNodes
// This API is used to list nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodesWithContext(ctx context.Context, request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstanceOperations")
    
    
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceOperations
// This API is used to get operations of the instance .
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// This API is used to get operations of the instance .
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceOperationsWithContext(ctx context.Context, request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStateRequest() (request *DescribeInstanceStateRequest) {
    request = &DescribeInstanceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstanceState")
    
    
    return
}

func NewDescribeInstanceStateResponse() (response *DescribeInstanceStateResponse) {
    response = &DescribeInstanceStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceState
// This API is used to display instance status, process progress, etc., in the instance details page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceState(request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    return c.DescribeInstanceStateWithContext(context.Background(), request)
}

// DescribeInstanceState
// This API is used to display instance status, process progress, etc., in the instance details page.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceStateWithContext(ctx context.Context, request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to get a list of  instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to get a list of  instances.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSimpleInstancesRequest() (request *DescribeSimpleInstancesRequest) {
    request = &DescribeSimpleInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeSimpleInstances")
    
    
    return
}

func NewDescribeSimpleInstancesResponse() (response *DescribeSimpleInstancesResponse) {
    response = &DescribeSimpleInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSimpleInstances
// This API is used to get a list of instance
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSimpleInstances(request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
    return c.DescribeSimpleInstancesWithContext(context.Background(), request)
}

// DescribeSimpleInstances
// This API is used to get a list of instance
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSimpleInstancesWithContext(ctx context.Context, request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSimpleInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSimpleInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSimpleInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogRequest() (request *DescribeSlowLogRequest) {
    request = &DescribeSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeSlowLog")
    
    
    return
}

func NewDescribeSlowLogResponse() (response *DescribeSlowLogResponse) {
    response = &DescribeSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLog
// This API is used to query slow SQL logs.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSlowLog(request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    return c.DescribeSlowLogWithContext(context.Background(), request)
}

// DescribeSlowLog
// This API is used to query slow SQL logs.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeSlowLogWithContext(ctx context.Context, request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUpgradeListRequest() (request *DescribeUpgradeListRequest) {
    request = &DescribeUpgradeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeUpgradeList")
    
    
    return
}

func NewDescribeUpgradeListResponse() (response *DescribeUpgradeListResponse) {
    response = &DescribeUpgradeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUpgradeList
// This API is used to obtain instance upgrade records.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeUpgradeList(request *DescribeUpgradeListRequest) (response *DescribeUpgradeListResponse, err error) {
    return c.DescribeUpgradeListWithContext(context.Background(), request)
}

// DescribeUpgradeList
// This API is used to obtain instance upgrade records.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
func (c *Client) DescribeUpgradeListWithContext(ctx context.Context, request *DescribeUpgradeListRequest) (response *DescribeUpgradeListResponse, err error) {
    if request == nil {
        request = NewDescribeUpgradeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUpgradeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUpgradeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserHbaConfigRequest() (request *DescribeUserHbaConfigRequest) {
    request = &DescribeUserHbaConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DescribeUserHbaConfig")
    
    
    return
}

func NewDescribeUserHbaConfigResponse() (response *DescribeUserHbaConfigResponse) {
    response = &DescribeUserHbaConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserHbaConfig
// Describe User HbaConfig.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserHbaConfig(request *DescribeUserHbaConfigRequest) (response *DescribeUserHbaConfigResponse, err error) {
    return c.DescribeUserHbaConfigWithContext(context.Background(), request)
}

// DescribeUserHbaConfig
// Describe User HbaConfig.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserHbaConfigWithContext(ctx context.Context, request *DescribeUserHbaConfigRequest) (response *DescribeUserHbaConfigResponse, err error) {
    if request == nil {
        request = NewDescribeUserHbaConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserHbaConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserHbaConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstanceByApiRequest() (request *DestroyInstanceByApiRequest) {
    request = &DestroyInstanceByApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "DestroyInstanceByApi")
    
    
    return
}

func NewDestroyInstanceByApiResponse() (response *DestroyInstanceByApiResponse) {
    response = &DestroyInstanceByApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyInstanceByApi
// This API is used to destroy instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstanceByApi(request *DestroyInstanceByApiRequest) (response *DestroyInstanceByApiResponse, err error) {
    return c.DestroyInstanceByApiWithContext(context.Background(), request)
}

// DestroyInstanceByApi
// This API is used to destroy instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstanceByApiWithContext(ctx context.Context, request *DestroyInstanceByApiRequest) (response *DestroyInstanceByApiResponse, err error) {
    if request == nil {
        request = NewDestroyInstanceByApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyInstanceByApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyInstanceByApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ModifyDBParameters")
    
    
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBParameters
// This API is used to modify instance configurations
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    return c.ModifyDBParametersWithContext(context.Background(), request)
}

// ModifyDBParameters
// This API is used to modify instance configurations
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyDBParametersWithContext(ctx context.Context, request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// This API is used to modify instance information. Only the name of an instance can be modified currently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to modify instance information. Only the name of an instance can be modified currently.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserHbaRequest() (request *ModifyUserHbaRequest) {
    request = &ModifyUserHbaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ModifyUserHba")
    
    
    return
}

func NewModifyUserHbaResponse() (response *ModifyUserHbaResponse) {
    response = &ModifyUserHbaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserHba
// This API is used to modify user Hba configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserHba(request *ModifyUserHbaRequest) (response *ModifyUserHbaResponse, err error) {
    return c.ModifyUserHbaWithContext(context.Background(), request)
}

// ModifyUserHba
// This API is used to modify user Hba configurations.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserHbaWithContext(ctx context.Context, request *ModifyUserHbaRequest) (response *ModifyUserHbaResponse, err error) {
    if request == nil {
        request = NewModifyUserHbaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserHba require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserHbaResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAccountPassword
// This API is used to change account password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// This API is used to change account password.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartInstance
// This API is used by users to proactively restart instances in the console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// This API is used by users to proactively restart instances in the console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// This API is used to scale out instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// This API is used to scale out instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstanceWithContext(ctx context.Context, request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleUpInstanceRequest() (request *ScaleUpInstanceRequest) {
    request = &ScaleUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "ScaleUpInstance")
    
    
    return
}

func NewScaleUpInstanceResponse() (response *ScaleUpInstanceResponse) {
    response = &ScaleUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleUpInstance
// This API is used to scale up instance in the console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstance(request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    return c.ScaleUpInstanceWithContext(context.Background(), request)
}

// ScaleUpInstance
// This API is used to scale up instance in the console.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstanceWithContext(ctx context.Context, request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    if request == nil {
        request = NewScaleUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwpg", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// This API is used to upgrade Instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// This API is used to upgrade Instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
