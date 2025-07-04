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

package v20201214

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-12-14"

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


func NewCreateDedicatedClusterRequest() (request *CreateDedicatedClusterRequest) {
    request = &CreateDedicatedClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "CreateDedicatedCluster")
    
    
    return
}

func NewCreateDedicatedClusterResponse() (response *CreateDedicatedClusterResponse) {
    response = &CreateDedicatedClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDedicatedCluster
// This API is used to create a CDC.
//
// error code that may be returned:
//  FAILEDOPERATION_CDCBINDVPCFAIL = "FailedOperation.CdcBindVpcFail"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) CreateDedicatedCluster(request *CreateDedicatedClusterRequest) (response *CreateDedicatedClusterResponse, err error) {
    return c.CreateDedicatedClusterWithContext(context.Background(), request)
}

// CreateDedicatedCluster
// This API is used to create a CDC.
//
// error code that may be returned:
//  FAILEDOPERATION_CDCBINDVPCFAIL = "FailedOperation.CdcBindVpcFail"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) CreateDedicatedClusterWithContext(ctx context.Context, request *CreateDedicatedClusterRequest) (response *CreateDedicatedClusterResponse, err error) {
    if request == nil {
        request = NewCreateDedicatedClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDedicatedCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDedicatedClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDedicatedClusterOrderRequest() (request *CreateDedicatedClusterOrderRequest) {
    request = &CreateDedicatedClusterOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "CreateDedicatedClusterOrder")
    
    
    return
}

func NewCreateDedicatedClusterOrderResponse() (response *CreateDedicatedClusterOrderResponse) {
    response = &CreateDedicatedClusterOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDedicatedClusterOrder
// This API is used to create a CDC order.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERCOSSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterCosSize"
//  INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERDATASTEPSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterDataStepSize"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERORDERID = "ResourceNotFound.InvalidDedicatedClusterOrderId"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERTYPEID = "ResourceNotFound.InvalidDedicatedClusterTypeId"
//  UNSUPPORTEDOPERATION_NONCUSTOMERAPPIDNOTSUPPORT = "UnsupportedOperation.NonCustomerAppIdNotSupport"
func (c *Client) CreateDedicatedClusterOrder(request *CreateDedicatedClusterOrderRequest) (response *CreateDedicatedClusterOrderResponse, err error) {
    return c.CreateDedicatedClusterOrderWithContext(context.Background(), request)
}

// CreateDedicatedClusterOrder
// This API is used to create a CDC order.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERCOSSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterCosSize"
//  INVALIDPARAMETERVALUE_INVALIDVALUEDEDICATEDCLUSTERDATASTEPSIZE = "InvalidParameterValue.InvalidValueDedicatedClusterDataStepSize"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERORDERID = "ResourceNotFound.InvalidDedicatedClusterOrderId"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERTYPEID = "ResourceNotFound.InvalidDedicatedClusterTypeId"
//  UNSUPPORTEDOPERATION_NONCUSTOMERAPPIDNOTSUPPORT = "UnsupportedOperation.NonCustomerAppIdNotSupport"
func (c *Client) CreateDedicatedClusterOrderWithContext(ctx context.Context, request *CreateDedicatedClusterOrderRequest) (response *CreateDedicatedClusterOrderResponse, err error) {
    if request == nil {
        request = NewCreateDedicatedClusterOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDedicatedClusterOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDedicatedClusterOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSiteRequest() (request *CreateSiteRequest) {
    request = &CreateSiteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "CreateSite")
    
    
    return
}

func NewCreateSiteResponse() (response *CreateSiteResponse) {
    response = &CreateSiteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSite
// This API is used to create a site.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) CreateSite(request *CreateSiteRequest) (response *CreateSiteResponse, err error) {
    return c.CreateSiteWithContext(context.Background(), request)
}

// CreateSite
// This API is used to create a site.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
func (c *Client) CreateSiteWithContext(ctx context.Context, request *CreateSiteRequest) (response *CreateSiteResponse, err error) {
    if request == nil {
        request = NewCreateSiteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSite require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSiteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDedicatedClustersRequest() (request *DeleteDedicatedClustersRequest) {
    request = &DeleteDedicatedClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DeleteDedicatedClusters")
    
    
    return
}

func NewDeleteDedicatedClustersResponse() (response *DeleteDedicatedClustersResponse) {
    response = &DeleteDedicatedClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDedicatedClusters
// This API is used to delete a CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DeleteDedicatedClusters(request *DeleteDedicatedClustersRequest) (response *DeleteDedicatedClustersResponse, err error) {
    return c.DeleteDedicatedClustersWithContext(context.Background(), request)
}

// DeleteDedicatedClusters
// This API is used to delete a CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DeleteDedicatedClustersWithContext(ctx context.Context, request *DeleteDedicatedClustersRequest) (response *DeleteDedicatedClustersResponse, err error) {
    if request == nil {
        request = NewDeleteDedicatedClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDedicatedClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDedicatedClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSitesRequest() (request *DeleteSitesRequest) {
    request = &DeleteSitesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DeleteSites")
    
    
    return
}

func NewDeleteSitesResponse() (response *DeleteSitesResponse) {
    response = &DeleteSitesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSites
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILDELETESITE = "FailedOperation.FailDeleteSite"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) DeleteSites(request *DeleteSitesRequest) (response *DeleteSitesResponse, err error) {
    return c.DeleteSitesWithContext(context.Background(), request)
}

// DeleteSites
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILDELETESITE = "FailedOperation.FailDeleteSite"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) DeleteSitesWithContext(ctx context.Context, request *DeleteSitesRequest) (response *DeleteSitesResponse, err error) {
    if request == nil {
        request = NewDeleteSitesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSitesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterCosCapacityRequest() (request *DescribeDedicatedClusterCosCapacityRequest) {
    request = &DescribeDedicatedClusterCosCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterCosCapacity")
    
    
    return
}

func NewDescribeDedicatedClusterCosCapacityResponse() (response *DescribeDedicatedClusterCosCapacityResponse) {
    response = &DescribeDedicatedClusterCosCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterCosCapacity
// This API is used to query the Cloud Object Storage (COS) capacity of the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterCosCapacity(request *DescribeDedicatedClusterCosCapacityRequest) (response *DescribeDedicatedClusterCosCapacityResponse, err error) {
    return c.DescribeDedicatedClusterCosCapacityWithContext(context.Background(), request)
}

// DescribeDedicatedClusterCosCapacity
// This API is used to query the Cloud Object Storage (COS) capacity of the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterCosCapacityWithContext(ctx context.Context, request *DescribeDedicatedClusterCosCapacityRequest) (response *DescribeDedicatedClusterCosCapacityResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterCosCapacityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterCosCapacity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterCosCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterHostStatisticsRequest() (request *DescribeDedicatedClusterHostStatisticsRequest) {
    request = &DescribeDedicatedClusterHostStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterHostStatistics")
    
    
    return
}

func NewDescribeDedicatedClusterHostStatisticsResponse() (response *DescribeDedicatedClusterHostStatisticsResponse) {
    response = &DescribeDedicatedClusterHostStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterHostStatistics
// This API is used to query the statistic information of the host in the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCETYPENOTSUPPORT = "InvalidParameter.InstanceTypeNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterHostStatistics(request *DescribeDedicatedClusterHostStatisticsRequest) (response *DescribeDedicatedClusterHostStatisticsResponse, err error) {
    return c.DescribeDedicatedClusterHostStatisticsWithContext(context.Background(), request)
}

// DescribeDedicatedClusterHostStatistics
// This API is used to query the statistic information of the host in the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCETYPENOTSUPPORT = "InvalidParameter.InstanceTypeNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterHostStatisticsWithContext(ctx context.Context, request *DescribeDedicatedClusterHostStatisticsRequest) (response *DescribeDedicatedClusterHostStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterHostStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterHostStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterHostStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterHostsRequest() (request *DescribeDedicatedClusterHostsRequest) {
    request = &DescribeDedicatedClusterHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterHosts")
    
    
    return
}

func NewDescribeDedicatedClusterHostsResponse() (response *DescribeDedicatedClusterHostsResponse) {
    response = &DescribeDedicatedClusterHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterHosts
// This API is used to query host information of the CDC
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCETYPENOTSUPPORT = "InvalidParameter.InstanceTypeNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterHosts(request *DescribeDedicatedClusterHostsRequest) (response *DescribeDedicatedClusterHostsResponse, err error) {
    return c.DescribeDedicatedClusterHostsWithContext(context.Background(), request)
}

// DescribeDedicatedClusterHosts
// This API is used to query host information of the CDC
//
// error code that may be returned:
//  INVALIDPARAMETER_INSTANCETYPENOTSUPPORT = "InvalidParameter.InstanceTypeNotSupport"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterHostsWithContext(ctx context.Context, request *DescribeDedicatedClusterHostsRequest) (response *DescribeDedicatedClusterHostsResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterInstanceTypesRequest() (request *DescribeDedicatedClusterInstanceTypesRequest) {
    request = &DescribeDedicatedClusterInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterInstanceTypes")
    
    
    return
}

func NewDescribeDedicatedClusterInstanceTypesResponse() (response *DescribeDedicatedClusterInstanceTypesResponse) {
    response = &DescribeDedicatedClusterInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterInstanceTypes
// This API is used to query instance types supported by the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterInstanceTypes(request *DescribeDedicatedClusterInstanceTypesRequest) (response *DescribeDedicatedClusterInstanceTypesResponse, err error) {
    return c.DescribeDedicatedClusterInstanceTypesWithContext(context.Background(), request)
}

// DescribeDedicatedClusterInstanceTypes
// This API is used to query instance types supported by the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
func (c *Client) DescribeDedicatedClusterInstanceTypesWithContext(ctx context.Context, request *DescribeDedicatedClusterInstanceTypesRequest) (response *DescribeDedicatedClusterInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterOrdersRequest() (request *DescribeDedicatedClusterOrdersRequest) {
    request = &DescribeDedicatedClusterOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterOrders")
    
    
    return
}

func NewDescribeDedicatedClusterOrdersResponse() (response *DescribeDedicatedClusterOrdersResponse) {
    response = &DescribeDedicatedClusterOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterOrders
// This API is used to query the list of CDC orders.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterOrders(request *DescribeDedicatedClusterOrdersRequest) (response *DescribeDedicatedClusterOrdersResponse, err error) {
    return c.DescribeDedicatedClusterOrdersWithContext(context.Background(), request)
}

// DescribeDedicatedClusterOrders
// This API is used to query the list of CDC orders.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterOrdersWithContext(ctx context.Context, request *DescribeDedicatedClusterOrdersRequest) (response *DescribeDedicatedClusterOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterOrdersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterOrders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterOverviewRequest() (request *DescribeDedicatedClusterOverviewRequest) {
    request = &DescribeDedicatedClusterOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterOverview")
    
    
    return
}

func NewDescribeDedicatedClusterOverviewResponse() (response *DescribeDedicatedClusterOverviewResponse) {
    response = &DescribeDedicatedClusterOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterOverview
// This API is used to query the overview of the CDC
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterOverview(request *DescribeDedicatedClusterOverviewRequest) (response *DescribeDedicatedClusterOverviewResponse, err error) {
    return c.DescribeDedicatedClusterOverviewWithContext(context.Background(), request)
}

// DescribeDedicatedClusterOverview
// This API is used to query the overview of the CDC
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterOverviewWithContext(ctx context.Context, request *DescribeDedicatedClusterOverviewRequest) (response *DescribeDedicatedClusterOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClusterTypesRequest() (request *DescribeDedicatedClusterTypesRequest) {
    request = &DescribeDedicatedClusterTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusterTypes")
    
    
    return
}

func NewDescribeDedicatedClusterTypesResponse() (response *DescribeDedicatedClusterTypesResponse) {
    response = &DescribeDedicatedClusterTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusterTypes
// This API is used to query the configuration list of the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterTypes(request *DescribeDedicatedClusterTypesRequest) (response *DescribeDedicatedClusterTypesResponse, err error) {
    return c.DescribeDedicatedClusterTypesWithContext(context.Background(), request)
}

// DescribeDedicatedClusterTypes
// This API is used to query the configuration list of the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeDedicatedClusterTypesWithContext(ctx context.Context, request *DescribeDedicatedClusterTypesRequest) (response *DescribeDedicatedClusterTypesResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClusterTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusterTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClusterTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedClustersRequest() (request *DescribeDedicatedClustersRequest) {
    request = &DescribeDedicatedClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedClusters")
    
    
    return
}

func NewDescribeDedicatedClustersResponse() (response *DescribeDedicatedClustersResponse) {
    response = &DescribeDedicatedClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedClusters
// This API is used to query the CDC list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeDedicatedClusters(request *DescribeDedicatedClustersRequest) (response *DescribeDedicatedClustersResponse, err error) {
    return c.DescribeDedicatedClustersWithContext(context.Background(), request)
}

// DescribeDedicatedClusters
// This API is used to query the CDC list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
func (c *Client) DescribeDedicatedClustersWithContext(ctx context.Context, request *DescribeDedicatedClustersRequest) (response *DescribeDedicatedClustersResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDedicatedSupportedZonesRequest() (request *DescribeDedicatedSupportedZonesRequest) {
    request = &DescribeDedicatedSupportedZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeDedicatedSupportedZones")
    
    
    return
}

func NewDescribeDedicatedSupportedZonesResponse() (response *DescribeDedicatedSupportedZonesResponse) {
    response = &DescribeDedicatedSupportedZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDedicatedSupportedZones
// This API is used to query the list of AZs supported by the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVALUEREGION = "InvalidParameterValue.InvalidValueRegion"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
func (c *Client) DescribeDedicatedSupportedZones(request *DescribeDedicatedSupportedZonesRequest) (response *DescribeDedicatedSupportedZonesResponse, err error) {
    return c.DescribeDedicatedSupportedZonesWithContext(context.Background(), request)
}

// DescribeDedicatedSupportedZones
// This API is used to query the list of AZs supported by the CDC.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVALUEREGION = "InvalidParameterValue.InvalidValueRegion"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
func (c *Client) DescribeDedicatedSupportedZonesWithContext(ctx context.Context, request *DescribeDedicatedSupportedZonesRequest) (response *DescribeDedicatedSupportedZonesResponse, err error) {
    if request == nil {
        request = NewDescribeDedicatedSupportedZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDedicatedSupportedZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDedicatedSupportedZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSitesRequest() (request *DescribeSitesRequest) {
    request = &DescribeSitesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeSites")
    
    
    return
}

func NewDescribeSitesResponse() (response *DescribeSitesResponse) {
    response = &DescribeSitesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSites
// This API is used to query the site list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVALUEREGION = "InvalidParameterValue.InvalidValueRegion"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
func (c *Client) DescribeSites(request *DescribeSitesRequest) (response *DescribeSitesResponse, err error) {
    return c.DescribeSitesWithContext(context.Background(), request)
}

// DescribeSites
// This API is used to query the site list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDVALUEREGION = "InvalidParameterValue.InvalidValueRegion"
//  INVALIDPARAMETERVALUE_REGION = "InvalidParameterValue.Region"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
func (c *Client) DescribeSitesWithContext(ctx context.Context, request *DescribeSitesRequest) (response *DescribeSitesResponse, err error) {
    if request == nil {
        request = NewDescribeSitesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSites require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSitesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSitesDetailRequest() (request *DescribeSitesDetailRequest) {
    request = &DescribeSitesDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "DescribeSitesDetail")
    
    
    return
}

func NewDescribeSitesDetailResponse() (response *DescribeSitesDetailResponse) {
    response = &DescribeSitesDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSitesDetail
// This API is used to query site details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeSitesDetail(request *DescribeSitesDetailRequest) (response *DescribeSitesDetailResponse, err error) {
    return c.DescribeSitesDetailWithContext(context.Background(), request)
}

// DescribeSitesDetail
// This API is used to query site details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDAPPIDFORMAT = "InvalidParameterValue.InvalidAppIdFormat"
func (c *Client) DescribeSitesDetailWithContext(ctx context.Context, request *DescribeSitesDetailRequest) (response *DescribeSitesDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSitesDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSitesDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSitesDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDedicatedClusterInfoRequest() (request *ModifyDedicatedClusterInfoRequest) {
    request = &ModifyDedicatedClusterInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "ModifyDedicatedClusterInfo")
    
    
    return
}

func NewModifyDedicatedClusterInfoResponse() (response *ModifyDedicatedClusterInfoResponse) {
    response = &ModifyDedicatedClusterInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDedicatedClusterInfo
// This API is used to modify the CDC information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) ModifyDedicatedClusterInfo(request *ModifyDedicatedClusterInfoRequest) (response *ModifyDedicatedClusterInfoResponse, err error) {
    return c.ModifyDedicatedClusterInfoWithContext(context.Background(), request)
}

// ModifyDedicatedClusterInfo
// This API is used to modify the CDC information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  INVALIDPARAMETERVALUE_ZONEMISMATCHREGION = "InvalidParameterValue.ZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORTED = "InvalidParameterValue.ZoneNotSupported"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
//  RESOURCEINSUFFICIENT_CLOUDDISKUNAVAILABLE = "ResourceInsufficient.CloudDiskUnavailable"
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERID = "ResourceNotFound.InvalidDedicatedClusterId"
//  RESOURCENOTFOUND_INVALIDSITEID = "ResourceNotFound.InvalidSiteId"
func (c *Client) ModifyDedicatedClusterInfoWithContext(ctx context.Context, request *ModifyDedicatedClusterInfoRequest) (response *ModifyDedicatedClusterInfoResponse, err error) {
    if request == nil {
        request = NewModifyDedicatedClusterInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDedicatedClusterInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDedicatedClusterInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOrderStatusRequest() (request *ModifyOrderStatusRequest) {
    request = &ModifyOrderStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "ModifyOrderStatus")
    
    
    return
}

func NewModifyOrderStatusResponse() (response *ModifyOrderStatusResponse) {
    response = &ModifyOrderStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOrderStatus
// This API is used to modify the status of large and small orders.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERORDERID = "ResourceNotFound.InvalidDedicatedClusterOrderId"
func (c *Client) ModifyOrderStatus(request *ModifyOrderStatusRequest) (response *ModifyOrderStatusResponse, err error) {
    return c.ModifyOrderStatusWithContext(context.Background(), request)
}

// ModifyOrderStatus
// This API is used to modify the status of large and small orders.
//
// error code that may be returned:
//  RESOURCENOTFOUND_INVALIDDEDICATEDCLUSTERORDERID = "ResourceNotFound.InvalidDedicatedClusterOrderId"
func (c *Client) ModifyOrderStatusWithContext(ctx context.Context, request *ModifyOrderStatusRequest) (response *ModifyOrderStatusResponse, err error) {
    if request == nil {
        request = NewModifyOrderStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOrderStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOrderStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySiteDeviceInfoRequest() (request *ModifySiteDeviceInfoRequest) {
    request = &ModifySiteDeviceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "ModifySiteDeviceInfo")
    
    
    return
}

func NewModifySiteDeviceInfoResponse() (response *ModifySiteDeviceInfoResponse) {
    response = &ModifySiteDeviceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySiteDeviceInfo
// This API is used to modify the information about devices in the equipment room.
//
// error code that may be returned:
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
func (c *Client) ModifySiteDeviceInfo(request *ModifySiteDeviceInfoRequest) (response *ModifySiteDeviceInfoResponse, err error) {
    return c.ModifySiteDeviceInfoWithContext(context.Background(), request)
}

// ModifySiteDeviceInfo
// This API is used to modify the information about devices in the equipment room.
//
// error code that may be returned:
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
func (c *Client) ModifySiteDeviceInfoWithContext(ctx context.Context, request *ModifySiteDeviceInfoRequest) (response *ModifySiteDeviceInfoResponse, err error) {
    if request == nil {
        request = NewModifySiteDeviceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySiteDeviceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySiteDeviceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifySiteInfoRequest() (request *ModifySiteInfoRequest) {
    request = &ModifySiteInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdc", APIVersion, "ModifySiteInfo")
    
    
    return
}

func NewModifySiteInfoResponse() (response *ModifySiteInfoResponse) {
    response = &ModifySiteInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySiteInfo
// This API is used to modify the equipment room information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
func (c *Client) ModifySiteInfo(request *ModifySiteInfoRequest) (response *ModifySiteInfoResponse, err error) {
    return c.ModifySiteInfoWithContext(context.Background(), request)
}

// ModifySiteInfo
// This API is used to modify the equipment room information.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER_ATLEASTONE = "MissingParameter.AtLeastOne"
func (c *Client) ModifySiteInfoWithContext(ctx context.Context, request *ModifySiteInfoRequest) (response *ModifySiteInfoResponse, err error) {
    if request == nil {
        request = NewModifySiteInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySiteInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySiteInfoResponse()
    err = c.Send(request, response)
    return
}
