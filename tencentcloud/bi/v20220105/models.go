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

package v20220105

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ApiDatasourceConfig struct {
	// API data source parsing result
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldsJsonData *string `json:"FieldsJsonData,omitnil,omitempty" name:"FieldsJsonData"`

	// Connection Type 1: Direct Connection 2: Extraction
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectionType *uint64 `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// Extraction frequency configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	FrequencyConfig *FrequencyConfig `json:"FrequencyConfig,omitnil,omitempty" name:"FrequencyConfig"`

	// Request URL
	// Note: This field may return null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 1:GET 2:POST
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestMethod *uint64 `json:"RequestMethod,omitnil,omitempty" name:"RequestMethod"`

	// Request headers
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestHeader *string `json:"RequestHeader,omitnil,omitempty" name:"RequestHeader"`

	// Request parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestParams *string `json:"RequestParams,omitnil,omitempty" name:"RequestParams"`

	// request body
	// Note: This field may return null, indicating that no valid values can be obtained.
	RequestBody *string `json:"RequestBody,omitnil,omitempty" name:"RequestBody"`

	// Username.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// Valid values: 1: no authentication; 2: BASIC_AUTH.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthorizationType *uint64 `json:"AuthorizationType,omitnil,omitempty" name:"AuthorizationType"`

	// Table ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableId *uint64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Mapping path DbName
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonPathDbNameMap *string `json:"JsonPathDbNameMap,omitnil,omitempty" name:"JsonPathDbNameMap"`

	// Authentication API
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthApi *string `json:"AuthApi,omitnil,omitempty" name:"AuthApi"`

	// Application Key
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// application key
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppSecret *string `json:"AppSecret,omitnil,omitempty" name:"AppSecret"`

	// Data Key
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Data key initialization vector
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretIv *string `json:"SecretIv,omitnil,omitempty" name:"SecretIv"`
}

// Predefined struct for user
type ApplyEmbedIntervalRequestParams struct {
	// Shares the project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Shares the page ID. This field is empty (0) for embedding a dashboard and is not passed for embedding ChatBI.
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Token requiring extension.
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// Alternate field.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// embed: page/dashboard embedding.
	// chatBIEmbed: ChatBI embedding.
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// panel: dashboard; page: page.
	// project, during ChatBI embedding.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type ApplyEmbedIntervalRequest struct {
	*tchttp.BaseRequest
	
	// Shares the project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Shares the page ID. This field is empty (0) for embedding a dashboard and is not passed for embedding ChatBI.
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Token requiring extension.
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// Alternate field.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// embed: page/dashboard embedding.
	// chatBIEmbed: ChatBI embedding.
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// panel: dashboard; page: page.
	// project, during ChatBI embedding.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

func (r *ApplyEmbedIntervalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "BIToken")
	delete(f, "ExtraParam")
	delete(f, "Intention")
	delete(f, "Scope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyEmbedIntervalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyEmbedIntervalResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Result data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *ApplyEmbedTokenInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Result description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyEmbedIntervalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyEmbedIntervalResponseParams `json:"Response"`
}

func (r *ApplyEmbedIntervalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyEmbedIntervalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyEmbedTokenInfo struct {
	// Request result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

type BaseStateAction struct {
	// Whether the edit button is visible.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowEdit *bool `json:"ShowEdit,omitnil,omitempty" name:"ShowEdit"`

	// Whether the edit button is clickable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsEdit *bool `json:"IsEdit,omitnil,omitempty" name:"IsEdit"`

	// Edit button text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EditText *string `json:"EditText,omitnil,omitempty" name:"EditText"`

	// Edit-disabled hint.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EditTips *string `json:"EditTips,omitnil,omitempty" name:"EditTips"`

	// Whether the deletion button is visible.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowDel *bool `json:"ShowDel,omitnil,omitempty" name:"ShowDel"`

	// Whether the deletion button is clickable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDel *bool `json:"IsDel,omitnil,omitempty" name:"IsDel"`

	// Delete button text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DelText *string `json:"DelText,omitnil,omitempty" name:"DelText"`

	// Delete-disabled hint.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DelTips *string `json:"DelTips,omitnil,omitempty" name:"DelTips"`

	// Whether the copy button is visible.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowCopy *bool `json:"ShowCopy,omitnil,omitempty" name:"ShowCopy"`

	// Whether it is visible.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowView *bool `json:"ShowView,omitnil,omitempty" name:"ShowView"`

	// Whether renaming is allowed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShowRename *bool `json:"ShowRename,omitnil,omitempty" name:"ShowRename"`
}

// Predefined struct for user
type ClearEmbedTokenRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Host Account ID
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// panel , page
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// page id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type ClearEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Host Account ID
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// panel , page
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// page id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *ClearEmbedTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearEmbedTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserCorpId")
	delete(f, "Scope")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearEmbedTokenResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Prompt message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Result
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearEmbedTokenResponse struct {
	*tchttp.BaseResponse
	Response *ClearEmbedTokenResponseParams `json:"Response"`
}

func (r *ClearEmbedTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearEmbedTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CorpUserListData struct {
	// List.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*UserIdAndUserName `json:"List,omitnil,omitempty" name:"List"`

	// Total number.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Number of pages.
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type CreateDataTableRequestParams struct {
	// The backend provides a dictionary: table type, 1. database table creation, 2. SQL table creation, 3. Excel upload, 4. API connection, 5. Tencent documentation, 6. cloud database, 7. manually enter, 8. join query.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Data table name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// None.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// folder
	FoldId *uint64 `json:"FoldId,omitnil,omitempty" name:"FoldId"`

	// data source Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// physical table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// excel address
	ExcelUrl *string `json:"ExcelUrl,omitnil,omitempty" name:"ExcelUrl"`

	// configure field
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Multi-table join usage: 1: Data source original table, 2: Local table, 3: Excel table, 4: API table
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// Original table information for multi-table join
	Tables []*JoinSourceTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// Multi-table join association information
	Joins []*JoinRelation `json:"Joins,omitnil,omitempty" name:"Joins"`

	// Additional info, such as api data source info and Tencent document access info
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// whether
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// dependent async transaction id
	ParentTranId *string `json:"ParentTranId,omitnil,omitempty" name:"ParentTranId"`

	// API data source configuration
	ApiDatasourceConfig *ApiDatasourceConfig `json:"ApiDatasourceConfig,omitnil,omitempty" name:"ApiDatasourceConfig"`

	// 1
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// dlc advanced parameter
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// Query database required or not
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// Table remark
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Whether to query field remarks
	QueryFieldRemark *int64 `json:"QueryFieldRemark,omitnil,omitempty" name:"QueryFieldRemark"`

	// Field remarks list
	FieldRemarkList []*FieldRemarkDTO `json:"FieldRemarkList,omitnil,omitempty" name:"FieldRemarkList"`
}

type CreateDataTableRequest struct {
	*tchttp.BaseRequest
	
	// The backend provides a dictionary: table type, 1. database table creation, 2. SQL table creation, 3. Excel upload, 4. API connection, 5. Tencent documentation, 6. cloud database, 7. manually enter, 8. join query.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Data table name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// None.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// folder
	FoldId *uint64 `json:"FoldId,omitnil,omitempty" name:"FoldId"`

	// data source Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// physical table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// excel address
	ExcelUrl *string `json:"ExcelUrl,omitnil,omitempty" name:"ExcelUrl"`

	// configure field
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Multi-table join usage: 1: Data source original table, 2: Local table, 3: Excel table, 4: API table
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// Original table information for multi-table join
	Tables []*JoinSourceTable `json:"Tables,omitnil,omitempty" name:"Tables"`

	// Multi-table join association information
	Joins []*JoinRelation `json:"Joins,omitnil,omitempty" name:"Joins"`

	// Additional info, such as api data source info and Tencent document access info
	ExtInfo *string `json:"ExtInfo,omitnil,omitempty" name:"ExtInfo"`

	// whether
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// dependent async transaction id
	ParentTranId *string `json:"ParentTranId,omitnil,omitempty" name:"ParentTranId"`

	// API data source configuration
	ApiDatasourceConfig *ApiDatasourceConfig `json:"ApiDatasourceConfig,omitnil,omitempty" name:"ApiDatasourceConfig"`

	// 1
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// dlc advanced parameter
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// Query database required or not
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// Table remark
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Whether to query field remarks
	QueryFieldRemark *int64 `json:"QueryFieldRemark,omitnil,omitempty" name:"QueryFieldRemark"`

	// Field remarks list
	FieldRemarkList []*FieldRemarkDTO `json:"FieldRemarkList,omitnil,omitempty" name:"FieldRemarkList"`
}

func (r *CreateDataTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "FoldId")
	delete(f, "DatasourceId")
	delete(f, "TableName")
	delete(f, "Sql")
	delete(f, "ExcelUrl")
	delete(f, "Fields")
	delete(f, "TableNodeType")
	delete(f, "Tables")
	delete(f, "Joins")
	delete(f, "ExtInfo")
	delete(f, "AsyncRequest")
	delete(f, "ParentTranId")
	delete(f, "ApiDatasourceConfig")
	delete(f, "ParamList")
	delete(f, "DlcExtInfo")
	delete(f, "QueryDbData")
	delete(f, "TableComment")
	delete(f, "QueryFieldRemark")
	delete(f, "FieldRemarkList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataTableResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Returned data table id on success
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// Additional Information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Error prompt
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataTableResponseParams `json:"Response"`
}

func (r *CreateDataTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceCloudRequestParams struct {
	// The backend provides dictionaries: domain type. 1. Tencent Cloud, 2. local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Private network IP address of the public cloud.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port of the public cloud.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unified VPC identifier.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Region identifier (gz, bj).
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Product name of the data source.
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type CreateDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// The backend provides dictionaries: domain type. 1. Tencent Cloud, 2. local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Private network IP address of the public cloud.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port of the public cloud.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unified VPC identifier.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Region identifier (gz, bj).
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Product name of the data source.
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

func (r *CreateDatasourceCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "UniqVpcId")
	delete(f, "RegionId")
	delete(f, "ExtraParam")
	delete(f, "InstanceId")
	delete(f, "ProdDbName")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ClusterId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceCloudResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Success No.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Prompt.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatasourceCloudResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasourceCloudResponseParams `json:"Response"`
}

func (r *CreateDatasourceCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceRequestParams struct {
	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// Port.
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// The backend provides dictionaries: domain type. 1. Tencent Cloud, 2. local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Catalog value.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Unified identifier of the Tencent Cloud VPC.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC IP address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Tencent Cloud VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Operation permission limitation.
	OperationAuthLimit []*string `json:"OperationAuthLimit,omitnil,omitempty" name:"OperationAuthLimit"`

	// Enables VPC.
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// Region.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type CreateDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// Port.
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// The backend provides dictionaries: domain type. 1. Tencent Cloud, 2. local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Catalog value.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Unified identifier of the Tencent Cloud VPC.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC IP address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Tencent Cloud VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Operation permission limitation.
	OperationAuthLimit []*string `json:"OperationAuthLimit,omitnil,omitempty" name:"OperationAuthLimit"`

	// Enables VPC.
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// Region.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

func (r *CreateDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbHost")
	delete(f, "DbPort")
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Catalog")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ExtraParam")
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "OperationAuthLimit")
	delete(f, "UseVPC")
	delete(f, "RegionId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasourceResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Data source ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *IdDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Prompt.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasourceResponseParams `json:"Response"`
}

func (r *CreateDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenRequestParams struct {
	// Shares the project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Shares the page ID. This field is empty (0) for embedding a dashboard and is not passed for embedding ChatBI.
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// "embed" indicates page dashboard embedding, and "chatBIEmbed" indicates ChatBI embedding.
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// "page" indicates embedding a page, "panel" indicates embedding the entire dashboard, and "project" is used for ChatBI embedding.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Expiration time. Unit: minutes. Maximum value: 240 (namely, 4 hours). Default value: 240.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Alternate field.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// User enterprise ID (only used for multi-user).
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// User ID (only used for multi-user).
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Access count limit (range: 1-99999). Leave empty to disable access restrictions.
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// Global filter parameters: Applies to all report filter criteria. Should be formatted as a JSON string.
	// ** Currently, only character-type page parameters can be bound to global parameters.
	// **
	// [
	//     {
	// "ParamKey": "name", page parameter name.
	// "JoinType": "AND", // connection method. Currently, only AND is supported.
	//         "WhereList": [
	//             {
	// "Operator": "-neq", // operator. Refer to the following instructions.
	// "Value": [ action value. For a single-value array, only one value is passed.
	//                     "zZWJMD",
	//                     "ZzVGHX",
	// "Hunan province".
	// "Hebei province".
	//                 ]
	//             }
	//         ]
	//     },
	//     {
	//         "ParamKey": "genderParam",
	//         "JoinType": "AND",
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",
	//                 "Value": [
	// "Male".
	//                 ]
	//             }
	//         ]
	//     }
	// ]
	// 
	// 
	// 
	// Currently supported operators.
	//  -neq not equal to != operator.
	//  -eq equal to = operator.
	//  -is in operator.
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// 100: no user bound. Create 1 token at a time. UserCorpId and UserId are optional. ChatBI embedding is not supported.
	// 200: single token per user. Create 1 token at a time. UserCorpId and UserId required.
	// 300: multiple tokens per user. Create multiple tokens at a time. UserCorpId and UserId required.
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// The number of tokens created at one time.
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// Embedded display configurations: Currently for ChatBI embedding scenarios; TableFilter represents data table list filtering, SqlView represents SQL view feature.
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
}

type CreateEmbedTokenRequest struct {
	*tchttp.BaseRequest
	
	// Shares the project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Shares the page ID. This field is empty (0) for embedding a dashboard and is not passed for embedding ChatBI.
	PageId *uint64 `json:"PageId,omitnil,omitempty" name:"PageId"`

	// "embed" indicates page dashboard embedding, and "chatBIEmbed" indicates ChatBI embedding.
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// "page" indicates embedding a page, "panel" indicates embedding the entire dashboard, and "project" is used for ChatBI embedding.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Expiration time. Unit: minutes. Maximum value: 240 (namely, 4 hours). Default value: 240.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Alternate field.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// User enterprise ID (only used for multi-user).
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// User ID (only used for multi-user).
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Access count limit (range: 1-99999). Leave empty to disable access restrictions.
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// Global filter parameters: Applies to all report filter criteria. Should be formatted as a JSON string.
	// ** Currently, only character-type page parameters can be bound to global parameters.
	// **
	// [
	//     {
	// "ParamKey": "name", page parameter name.
	// "JoinType": "AND", // connection method. Currently, only AND is supported.
	//         "WhereList": [
	//             {
	// "Operator": "-neq", // operator. Refer to the following instructions.
	// "Value": [ action value. For a single-value array, only one value is passed.
	//                     "zZWJMD",
	//                     "ZzVGHX",
	// "Hunan province".
	// "Hebei province".
	//                 ]
	//             }
	//         ]
	//     },
	//     {
	//         "ParamKey": "genderParam",
	//         "JoinType": "AND",
	//         "WhereList": [
	//             {
	//                 "Operator": "-neq",
	//                 "Value": [
	// "Male".
	//                 ]
	//             }
	//         ]
	//     }
	// ]
	// 
	// 
	// 
	// Currently supported operators.
	//  -neq not equal to != operator.
	//  -eq equal to = operator.
	//  -is in operator.
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// 100: no user bound. Create 1 token at a time. UserCorpId and UserId are optional. ChatBI embedding is not supported.
	// 200: single token per user. Create 1 token at a time. UserCorpId and UserId required.
	// 300: multiple tokens per user. Create multiple tokens at a time. UserCorpId and UserId required.
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// The number of tokens created at one time.
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// Embedded display configurations: Currently for ChatBI embedding scenarios; TableFilter represents data table list filtering, SqlView represents SQL view feature.
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
}

func (r *CreateEmbedTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "Intention")
	delete(f, "Scope")
	delete(f, "ExpireTime")
	delete(f, "ExtraParam")
	delete(f, "UserCorpId")
	delete(f, "UserId")
	delete(f, "TicketNum")
	delete(f, "GlobalParam")
	delete(f, "TokenType")
	delete(f, "TokenNum")
	delete(f, "ConfigParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEmbedTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEmbedTokenResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *EmbedTokenInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Result description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEmbedTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateEmbedTokenResponseParams `json:"Response"`
}

func (r *CreateEmbedTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEmbedTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePermissionRanksRequestParams struct {
	// <p>page number</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>Mode</p><p>Enumeration value:</p><ul><li>ALL: All</li><li>Specify: Specify</li><li>TAG: Tag</li></ul><p>Default value: ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>Role type</p><p>Enumeration value:</p><ul><li>ROLES: By role</li><li>Others: Other</li></ul><p>Default value: Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>All page numbers</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>Rule information</p>
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// <p>Type</p><p>Enumeration value:</p><ul><li>ROW: row permission</li><li>COLUMN: column permission</li></ul><p>Default value: ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Enabled status</p><p>Enumeration value:</p><ul><li>Open: Turn on</li><li>Close: Turn off</li></ul><p>Default value: Close</p>
	OpenStatus *string `json:"OpenStatus,omitnil,omitempty" name:"OpenStatus"`

	// <p>Project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Row/column permission configuration</p>
	RowColumnConfigList []*RowColumnConfig `json:"RowColumnConfigList,omitnil,omitempty" name:"RowColumnConfigList"`
}

type CreatePermissionRanksRequest struct {
	*tchttp.BaseRequest
	
	// <p>page number</p>
	TableId *int64 `json:"TableId,omitnil,omitempty" name:"TableId"`

	// <p>Mode</p><p>Enumeration value:</p><ul><li>ALL: All</li><li>Specify: Specify</li><li>TAG: Tag</li></ul><p>Default value: ALL</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>Role type</p><p>Enumeration value:</p><ul><li>ROLES: By role</li><li>Others: Other</li></ul><p>Default value: Others</p>
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// <p>All page numbers</p>
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// <p>Rule information</p>
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// <p>Type</p><p>Enumeration value:</p><ul><li>ROW: row permission</li><li>COLUMN: column permission</li></ul><p>Default value: ROW</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Enabled status</p><p>Enumeration value:</p><ul><li>Open: Turn on</li><li>Close: Turn off</li></ul><p>Default value: Close</p>
	OpenStatus *string `json:"OpenStatus,omitnil,omitempty" name:"OpenStatus"`

	// <p>Project ID.</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>Row/column permission configuration</p>
	RowColumnConfigList []*RowColumnConfig `json:"RowColumnConfigList,omitnil,omitempty" name:"RowColumnConfigList"`
}

func (r *CreatePermissionRanksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePermissionRanksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "Mode")
	delete(f, "RoleType")
	delete(f, "RoleId")
	delete(f, "RulerInfo")
	delete(f, "Type")
	delete(f, "OpenStatus")
	delete(f, "ProjectId")
	delete(f, "RowColumnConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePermissionRanksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePermissionRanksResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>Message</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// <p>112</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// <p>1</p>
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePermissionRanksResponse struct {
	*tchttp.BaseResponse
	Response *CreatePermissionRanksResponseParams `json:"Response"`
}

func (r *CreatePermissionRanksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePermissionRanksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectRequestParams struct {
	// Project name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Background color of the logo.
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// Project logo.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Whether to allow user requests.
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// Management platform.
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Background color of the logo.
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// Project logo.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Whether to allow user requests.
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// Management platform.
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ColorCode")
	delete(f, "Logo")
	delete(f, "Mark")
	delete(f, "IsApply")
	delete(f, "DefaultPanelType")
	delete(f, "ManagePlatform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Extra data.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	Data *Data `json:"Data,omitnil,omitempty" name:"Data"`

	// Returned information.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateProjectResponseParams `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleProjectRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// User list (deprecated).
	//
	// Deprecated: UserList is deprecated.
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// User list (new).
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`
}

type CreateUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// User list (deprecated).
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// User list (new).
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`
}

func (r *CreateUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RoleIdList")
	delete(f, "UserList")
	delete(f, "UserInfoList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DataId `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserRoleProjectResponseParams `json:"Response"`
}

func (r *CreateUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleRequestParams struct {
	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// User list (deprecated).
	//
	// Deprecated: UserList is deprecated.
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// User list (new).
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`

	// User group ID list.
	UserGroups []*uint64 `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`
}

type CreateUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// User list (deprecated).
	UserList []*UserIdAndUserName `json:"UserList,omitnil,omitempty" name:"UserList"`

	// User list (new).
	UserInfoList []*UserInfo `json:"UserInfoList,omitnil,omitempty" name:"UserInfoList"`

	// User group ID list.
	UserGroups []*uint64 `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`
}

func (r *CreateUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleIdList")
	delete(f, "UserList")
	delete(f, "UserInfoList")
	delete(f, "UserGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRoleResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DataId `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserRoleResponseParams `json:"Response"`
}

func (r *CreateUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Data struct {
	// Project ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// url
	// Note: This field may return null, indicating that no valid values can be obtained.
	EditUrl *string `json:"EditUrl,omitnil,omitempty" name:"EditUrl"`
}

type DataId struct {
	// Data ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DatasourceInfo struct {
	// Database ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Domain type. Valid values: 1: Tencent Cloud; 2: local.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Database driver.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// Port.
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Creation time.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Modification time.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// Catalog value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Connection type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectType *string `json:"ConnectType,omitnil,omitempty" name:"ConnectType"`

	// Project ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Data source description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Data source status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Source platform.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourcePlat *string `json:"SourcePlat,omitnil,omitempty" name:"SourcePlat"`

	// Extension parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AddInfo *string `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`

	// Project name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Engine type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// Data source owner.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Manager *string `json:"Manager,omitnil,omitempty" name:"Manager"`

	// Operation personnel allowlist.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorWhitelist *string `json:"OperatorWhitelist,omitnil,omitempty" name:"OperatorWhitelist"`

	// VPC information of the data source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// uniqVpc information of the data source.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Data source region information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Operation attributes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	StateAction *BaseStateAction `json:"StateAction,omitnil,omitempty" name:"StateAction"`

	// Updater.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// Permission list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PermissionList []*PermissionGroup `json:"PermissionList,omitnil,omitempty" name:"PermissionList"`

	// Permission value list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthList []*string `json:"AuthList,omitnil,omitempty" name:"AuthList"`

	// Third-party data source identifier.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Cluster ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Data source name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbTypeName *string `json:"DbTypeName,omitnil,omitempty" name:"DbTypeName"`

	// Enable VPC.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// Associated person ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// Associated person name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// Database schema.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type DatasourceInfoData struct {
	// Data source details list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*DatasourceInfo `json:"List,omitnil,omitempty" name:"List"`

	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Total number of pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

// Predefined struct for user
type DeleteDatasourceRequestParams struct {
	// Data source ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// Data source ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasourceResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Expansion.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Information.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatasourceResponseParams `json:"Response"`
}

func (r *DeleteDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectRequestParams struct {
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Random number.
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Random number.
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Seed")
	delete(f, "DefaultPanelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// "".
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectResponseParams `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleProjectRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserRoleProjectResponseParams `json:"Response"`
}

func (r *DeleteUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleRequestParams struct {
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRoleResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserRoleResponseParams `json:"Response"`
}

func (r *DeleteUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceListRequestParams struct {
	// None.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Returns all pages and defaults to false.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Database name search.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// None.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// None.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Permission filter (0: all permissions, 1: access permission, 2: edit permission).
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

type DescribeDatasourceListRequest struct {
	*tchttp.BaseRequest
	
	// None.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Returns all pages and defaults to false.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Database name search.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// None.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// None.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Permission filter (0: all permissions, 1: access permission, 2: edit permission).
	PermissionType *int64 `json:"PermissionType,omitnil,omitempty" name:"PermissionType"`
}

func (r *DescribeDatasourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AllPage")
	delete(f, "DbName")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "Keyword")
	delete(f, "PermissionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// List details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *DatasourceInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// Information.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Information.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasourceListResponseParams `json:"Response"`
}

func (r *DescribeDatasourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePageWidgetListRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type DescribePageWidgetListRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

func (r *DescribePageWidgetListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePageWidgetListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePageWidgetListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePageWidgetListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Extension parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Returned data results.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *WidgetListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Response message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePageWidgetListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePageWidgetListResponseParams `json:"Response"`
}

func (r *DescribePageWidgetListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePageWidgetListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoRequestParams struct {
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

type DescribeProjectInfoRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`
}

func (r *DescribeProjectInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "DefaultPanelType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectInfoResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Project details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *Project `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectInfoResponseParams `json:"Response"`
}

func (r *DescribeProjectInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListRequestParams struct {
	// Page capacity. The initial version defaults to 20 and may change dynamically based on screen width in the future.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page marker.
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Fuzzy search field.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Whether to display all. If true, ignore pagination.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Role information.
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// moduleId set.
	ModuleIdList []*string `json:"ModuleIdList,omitnil,omitempty" name:"ModuleIdList"`
}

type DescribeProjectListRequest struct {
	*tchttp.BaseRequest
	
	// Page capacity. The initial version defaults to 20 and may change dynamically based on screen width in the future.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page marker.
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Fuzzy search field.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Whether to display all. If true, ignore pagination.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Role information.
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`

	// moduleId set.
	ModuleIdList []*string `json:"ModuleIdList,omitnil,omitempty" name:"ModuleIdList"`
}

func (r *DescribeProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Keyword")
	delete(f, "AllPage")
	delete(f, "ModuleCollection")
	delete(f, "ModuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// API information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *ProjectListData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectListResponseParams `json:"Response"`
}

func (r *DescribeProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceFieldListRequestParams struct {
	// data source Id
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql content
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// whether
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// async transaction id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 11
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// DLC extension parameter
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// Query database required or not
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// Data table Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// The backend provides a dictionary: table type, 1. database table creation, 2. SQL table creation, 3. Excel upload, 4. API connection, 5. Tencent documentation, 6. cloud database, 7. manually enter, 8. join query.
	TableType *uint64 `json:"TableType,omitnil,omitempty" name:"TableType"`
}

type DescribeSourceFieldListRequest struct {
	*tchttp.BaseRequest
	
	// data source Id
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// sql content
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// whether
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// async transaction id
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 11
	ParamList []*ParamCreateDTO `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// DLC extension parameter
	DlcExtInfo *string `json:"DlcExtInfo,omitnil,omitempty" name:"DlcExtInfo"`

	// Query database required or not
	QueryDbData *string `json:"QueryDbData,omitnil,omitempty" name:"QueryDbData"`

	// Data table Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// The backend provides a dictionary: table type, 1. database table creation, 2. SQL table creation, 3. Excel upload, 4. API connection, 5. Tencent documentation, 6. cloud database, 7. manually enter, 8. join query.
	TableType *uint64 `json:"TableType,omitnil,omitempty" name:"TableType"`
}

func (r *DescribeSourceFieldListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceFieldListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataSourceId")
	delete(f, "TableName")
	delete(f, "Sql")
	delete(f, "ProjectId")
	delete(f, "AsyncRequest")
	delete(f, "TranId")
	delete(f, "ParamList")
	delete(f, "DlcExtInfo")
	delete(f, "QueryDbData")
	delete(f, "TableId")
	delete(f, "TableType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceFieldListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceFieldListResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional Information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// List of fields in the table
	Data *TableColumnListData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceFieldListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceFieldListResponseParams `json:"Response"`
}

func (r *DescribeSourceFieldListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceFieldListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// None.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// None.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// None.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Whether to filter out enterprise administrators.
	IsFilterPerAuthUser *bool `json:"IsFilterPerAuthUser,omitnil,omitempty" name:"IsFilterPerAuthUser"`

	// Whether to filter out the current user.
	IsFilterCurrentUser *bool `json:"IsFilterCurrentUser,omitnil,omitempty" name:"IsFilterCurrentUser"`

	// Keyword.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeUserProjectListRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// None.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// None.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// None.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Whether to filter out enterprise administrators.
	IsFilterPerAuthUser *bool `json:"IsFilterPerAuthUser,omitnil,omitempty" name:"IsFilterPerAuthUser"`

	// Whether to filter out the current user.
	IsFilterCurrentUser *bool `json:"IsFilterCurrentUser,omitnil,omitempty" name:"IsFilterCurrentUser"`

	// Keyword.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeUserProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AllPage")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "IsFilterPerAuthUser")
	delete(f, "IsFilterCurrentUser")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserProjectListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *CorpUserListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserProjectListResponseParams `json:"Response"`
}

func (r *DescribeUserProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleListRequestParams struct {
	// Page number.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Number of pages.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// All page numbers.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 0: enterprise user. 1: visitor. If left blank, it indicates all users.
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// Keyword for fuzzy search.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether to only obtain users bound to the WeCom app.
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`
}

type DescribeUserRoleListRequest struct {
	*tchttp.BaseRequest
	
	// Page number.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Number of pages.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// All page numbers.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// 0: enterprise user. 1: visitor. If left blank, it indicates all users.
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// Keyword for fuzzy search.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether to only obtain users bound to the WeCom app.
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`
}

func (r *DescribeUserRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "AllPage")
	delete(f, "UserType")
	delete(f, "Keyword")
	delete(f, "ProjectId")
	delete(f, "IsOnlyBindAppUser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Extension description information (providing more exception messages for auxiliary judgment).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *UserRoleListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRoleListResponseParams `json:"Response"`
}

func (r *DescribeUserRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleProjectListRequestParams struct {
	// Page number.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Number of pages.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether to only obtain users bound to the WeCom app.
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`

	// Whether to obtain all the data.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Role code.
	RoleCode *string `json:"RoleCode,omitnil,omitempty" name:"RoleCode"`

	// User ID list.
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeUserRoleProjectListRequest struct {
	*tchttp.BaseRequest
	
	// Page number.
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// Number of pages.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Whether to only obtain users bound to the WeCom app.
	IsOnlyBindAppUser *bool `json:"IsOnlyBindAppUser,omitnil,omitempty" name:"IsOnlyBindAppUser"`

	// Whether to obtain all the data.
	AllPage *bool `json:"AllPage,omitnil,omitempty" name:"AllPage"`

	// Role code.
	RoleCode *string `json:"RoleCode,omitnil,omitempty" name:"RoleCode"`

	// User ID list.
	UserIdList []*string `json:"UserIdList,omitnil,omitempty" name:"UserIdList"`

	// Search keywords.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeUserRoleProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleProjectListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "IsOnlyBindAppUser")
	delete(f, "AllPage")
	delete(f, "RoleCode")
	delete(f, "UserIdList")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRoleProjectListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRoleProjectListResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *UserRoleListData `json:"Data,omitnil,omitempty" name:"Data"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserRoleProjectListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRoleProjectListResponseParams `json:"Response"`
}

func (r *DescribeUserRoleProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRoleProjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EmbedTokenInfo struct {
	// Information identifier.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Token.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BIToken *string `json:"BIToken,omitnil,omitempty" name:"BIToken"`

	// Project ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Updater.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Page ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Backup.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Embedding type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// Expiration time (in minutes), with a maximum value of 240.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// User enterprise ID (only used for multi-user).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserCorpId *string `json:"UserCorpId,omitnil,omitempty" name:"UserCorpId"`

	// User ID (only used for multi-user).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Access count limit (range: 1–99999). Leave empty to disable access restrictions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TicketNum *int64 `json:"TicketNum,omitnil,omitempty" name:"TicketNum"`

	// Global parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GlobalParam *string `json:"GlobalParam,omitnil,omitempty" name:"GlobalParam"`

	// "embed" indicates page dashboard embedding, and "chatBIEmbed" indicates ChatBI embedding.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Intention *string `json:"Intention,omitnil,omitempty" name:"Intention"`

	// 100: no bound user.
	// 200: single token per user.
	// 300: multiple tokens per user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TokenType *int64 `json:"TokenType,omitnil,omitempty" name:"TokenType"`

	// Number of tokens.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TokenNum *int64 `json:"TokenNum,omitnil,omitempty" name:"TokenNum"`

	// Whether it is multiple tokens per user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SingleUserMultiToken *bool `json:"SingleUserMultiToken,omitnil,omitempty" name:"SingleUserMultiToken"`

	// Embedded display configurations: Currently for ChatBI embedding scenarios; TableFilter represents data table list filtering, SqlView represents SQL view feature.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigParam *string `json:"ConfigParam,omitnil,omitempty" name:"ConfigParam"`
}

type EmptyValue struct {
	// Empty value display style type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Display style type for null values corresponds to specific display string
	// Note: This field may return null, indicating that no valid values can be obtained.
	Custom *string `json:"Custom,omitnil,omitempty" name:"Custom"`
}

type EmptyValueConfig struct {
	// Numeric value field null style configuration
	// Note: This field may return null, indicating that no valid values can be obtained.
	Number *EmptyValue `json:"Number,omitnil,omitempty" name:"Number"`

	// Style configuration for empty string fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	String *EmptyValue `json:"String,omitnil,omitempty" name:"String"`
}

type ErrorInfo struct {
	// Error description field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorTip *string `json:"ErrorTip,omitnil,omitempty" name:"ErrorTip"`

	// Original exception message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Error level field.
	// ERROR
	// WARN
	// INFO
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorLevel *string `json:"ErrorLevel,omitnil,omitempty" name:"ErrorLevel"`

	// Documentation link.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DocLink *string `json:"DocLink,omitnil,omitempty" name:"DocLink"`

	// Quick start guide.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FAQ *string `json:"FAQ,omitnil,omitempty" name:"FAQ"`

	// Reserved field 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReservedField *string `json:"ReservedField,omitnil,omitempty" name:"ReservedField"`
}

// Predefined struct for user
type ExportScreenPageRequestParams struct {
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Canvas type. Grid canvas: GRID; Free canvas: FREE.
	CanvasType *string `json:"CanvasType,omitnil,omitempty" name:"CanvasType"`

	// Image export type. Valid values: Base64, and URL (valid period: 1 day).
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// Component IDs. When empty, export the entire page.
	WidgetIds []*string `json:"WidgetIds,omitnil,omitempty" name:"WidgetIds"`

	// Whether it is an async request.
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// Transaction ID.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

type ExportScreenPageRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page ID.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Canvas type. Grid canvas: GRID; Free canvas: FREE.
	CanvasType *string `json:"CanvasType,omitnil,omitempty" name:"CanvasType"`

	// Image export type. Valid values: Base64, and URL (valid period: 1 day).
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// Component IDs. When empty, export the entire page.
	WidgetIds []*string `json:"WidgetIds,omitnil,omitempty" name:"WidgetIds"`

	// Whether it is an async request.
	AsyncRequest *bool `json:"AsyncRequest,omitnil,omitempty" name:"AsyncRequest"`

	// Transaction ID.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`
}

func (r *ExportScreenPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScreenPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageId")
	delete(f, "CanvasType")
	delete(f, "PicType")
	delete(f, "WidgetIds")
	delete(f, "AsyncRequest")
	delete(f, "TranId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportScreenPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportScreenPageResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Extension parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Returned data results.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *PageScreenListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Response message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportScreenPageResponse struct {
	*tchttp.BaseResponse
	Response *ExportScreenPageResponseParams `json:"Response"`
}

func (r *ExportScreenPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportScreenPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldRemarkDTO struct {
	// field name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// Field remarks list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment []*string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type FrequencyConfig struct {
	// Cycle
	// Note: This field may return null, indicating that no valid values can be obtained.
	Frequency *string `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Date
	// Note: This field may return null, indicating that no valid values can be obtained.
	Dates []*int64 `json:"Dates,omitnil,omitempty" name:"Dates"`

	// Time
	// Note: This field may return null, indicating that no valid values can be obtained.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Interval
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalTime *uint64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 1:SECOND,2:MINUTE,3:HOUR,4:DAY
	// Note: This field may return null, indicating that no valid values can be obtained.
	IntervalTimeUnit *uint64 `json:"IntervalTimeUnit,omitnil,omitempty" name:"IntervalTimeUnit"`

	// hourly list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Hours []*int64 `json:"Hours,omitnil,omitempty" name:"Hours"`

	// Minute list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Minute []*int64 `json:"Minute,omitnil,omitempty" name:"Minute"`
}

type IdDTO struct {
	// Request ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// key
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// id
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Transaction ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Transaction status.
	// Value range:.
	// Processing.
	// 2: processing is successful.
	// 3: processing failure.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type JoinRelation struct {
	// Association relationship id, used by the frontend
	// Note: This field may return null, indicating that no valid values can be obtained.
	JoinId *string `json:"JoinId,omitnil,omitempty" name:"JoinId"`

	// Original table node id
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceTableNodeId *string `json:"SourceTableNodeId,omitnil,omitempty" name:"SourceTableNodeId"`

	// Target table node id
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetTableNodeId *string `json:"TargetTableNodeId,omitnil,omitempty" name:"TargetTableNodeId"`

	// Association type of multi-table join
	// Note: This field may return null, indicating that no valid values can be obtained.
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// Field list for joined tables
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fields []*JoinRelationField `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type JoinRelationField struct {
	// Field association id, frontend usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldJoinId *string `json:"FieldJoinId,omitnil,omitempty" name:"FieldJoinId"`

	// Original table field
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceField *TableField `json:"SourceField,omitnil,omitempty" name:"SourceField"`

	// Target table field
	// Note: This field may return null, indicating that no valid values can be obtained.
	TargetField *TableField `json:"TargetField,omitnil,omitempty" name:"TargetField"`
}

type JoinSourceTable struct {
	// 1: Data source original table, 2: Local table, 3: Excel table, 4: API table
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableNodeType *uint64 `json:"TableNodeType,omitnil,omitempty" name:"TableNodeType"`

	// Base Table Node Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// Parent node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Optional, the data source has no ID in the original table.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// Required. Use the original table name for the data source. Use the logical table name for other types.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Field list to display in the base table
	// Note: This field may return null, indicating that no valid values can be obtained.
	Fields []*TableField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Data source ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatasourceId *uint64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// Optional, alias of the data source displayed on the front-end, excel table creation is required
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableAlias *string `json:"TableAlias,omitnil,omitempty" name:"TableAlias"`
}

// Predefined struct for user
type ModifyDatasourceCloudRequestParams struct {
	// The backend provides dictionaries: domain type. 1: Tencent Cloud; 2: local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Primary key.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Private network IP address of the public cloud.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port of the public cloud.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unified VPC identifier.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Region identifier (gz, bj).
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Product name of the data source.
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type ModifyDatasourceCloudRequest struct {
	*tchttp.BaseRequest
	
	// The backend provides dictionaries: domain type. 1: Tencent Cloud; 2: local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Project ID.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Primary key.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Private network IP address of the public cloud.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Private network port of the public cloud.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Unified VPC identifier.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// Region identifier (gz, bj).
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Product name of the data source.
	ProdDbName *string `json:"ProdDbName,omitnil,omitempty" name:"ProdDbName"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

func (r *ModifyDatasourceCloudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceCloudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "ProjectId")
	delete(f, "Id")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "UniqVpcId")
	delete(f, "RegionId")
	delete(f, "ExtraParam")
	delete(f, "InstanceId")
	delete(f, "ProdDbName")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ClusterId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceCloudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceCloudResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// None.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Prompt.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatasourceCloudResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatasourceCloudResponseParams `json:"Response"`
}

func (r *ModifyDatasourceCloudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceCloudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceRequestParams struct {
	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// Port.
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// The backend provides dictionaries: domain type. 1: Tencent Cloud; 2: local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Data source ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Catalog value.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Unified identifier of the Tencent Cloud VPC.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC IP address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Tencent Cloud VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Enable VPC.  
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// Region.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

type ModifyDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// HOST
	DbHost *string `json:"DbHost,omitnil,omitempty" name:"DbHost"`

	// Port.
	DbPort *uint64 `json:"DbPort,omitnil,omitempty" name:"DbPort"`

	// The backend provides dictionaries: domain type. 1: Tencent Cloud; 2: local.
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// Drive.
	// Value range:.
	// MYSQL: MYSQL database.
	// PRESTO: presto database.
	// POSTGRE: PostgreSQL database.
	// DLC: dlc database.
	// MSSQL: microsoft SQL Server database.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Database encoding.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username.
	DbUser *string `json:"DbUser,omitnil,omitempty" name:"DbUser"`

	// Password.
	DbPwd *string `json:"DbPwd,omitnil,omitempty" name:"DbPwd"`

	// Database name.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database alias.
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Data source ID.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Catalog value.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Third-party data source identifier.
	DataOrigin *string `json:"DataOrigin,omitnil,omitempty" name:"DataOrigin"`

	// Third-party project ID.
	DataOriginProjectId *string `json:"DataOriginProjectId,omitnil,omitempty" name:"DataOriginProjectId"`

	// Third-party data source ID.
	DataOriginDatasourceId *string `json:"DataOriginDatasourceId,omitnil,omitempty" name:"DataOriginDatasourceId"`

	// Extension parameter.
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// Unified identifier of the Tencent Cloud VPC.
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// VPC IP address.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// VPC port.
	Vport *string `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Tencent Cloud VPC identifier.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Enable VPC.  
	UseVPC *bool `json:"UseVPC,omitnil,omitempty" name:"UseVPC"`

	// Region.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Database schema.
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// Database version.
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`
}

func (r *ModifyDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbHost")
	delete(f, "DbPort")
	delete(f, "ServiceType")
	delete(f, "DbType")
	delete(f, "Charset")
	delete(f, "DbUser")
	delete(f, "DbPwd")
	delete(f, "DbName")
	delete(f, "SourceName")
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "Catalog")
	delete(f, "DataOrigin")
	delete(f, "DataOriginProjectId")
	delete(f, "DataOriginDatasourceId")
	delete(f, "ExtraParam")
	delete(f, "UniqVpcId")
	delete(f, "Vip")
	delete(f, "Vport")
	delete(f, "VpcId")
	delete(f, "UseVPC")
	delete(f, "RegionId")
	delete(f, "Schema")
	delete(f, "DbVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatasourceResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// None.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Prompt.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatasourceResponseParams `json:"Response"`
}

func (r *ModifyDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectRequestParams struct {
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Color value.
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// Icon.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Available upon request.
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// Seed.
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// 2
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// Project management platform.
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Color value.
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// Icon.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Remarks.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Available upon request.
	IsApply *bool `json:"IsApply,omitnil,omitempty" name:"IsApply"`

	// Seed.
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// Default dashboard.
	// Value range:.
	// 1: project dashboard. 
	// 2: my dashboard.
	DefaultPanelType *int64 `json:"DefaultPanelType,omitnil,omitempty" name:"DefaultPanelType"`

	// 2
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// Project management platform.
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ColorCode")
	delete(f, "Logo")
	delete(f, "Mark")
	delete(f, "IsApply")
	delete(f, "Seed")
	delete(f, "DefaultPanelType")
	delete(f, "PanelScope")
	delete(f, "ManagePlatform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Additional information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Returned data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProjectResponseParams `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserGroupResourceRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID
	UserGroupId *int64 `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// resource
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Entity class
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// Resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type ModifyResourceUserGroupResourceRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID
	UserGroupId *int64 `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// resource
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Entity class
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// Resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *ModifyResourceUserGroupResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserGroupId")
	delete(f, "Resource")
	delete(f, "EntityIds")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceUserGroupResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserGroupResourceResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Data.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceUserGroupResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceUserGroupResourceResponseParams `json:"Response"`
}

func (r *ModifyResourceUserGroupResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserGroupResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// resource
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Entity class
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// Resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type ModifyResourceUserRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// resource
	Resource *UserResourceDTO `json:"Resource,omitnil,omitempty" name:"Resource"`

	// Entity class
	EntityIds []*int64 `json:"EntityIds,omitnil,omitempty" name:"EntityIds"`

	// Resource type.
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *ModifyResourceUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	delete(f, "Resource")
	delete(f, "EntityIds")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceUserResponseParams struct {
	// Custom error information object
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Data.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceUserResponseParams `json:"Response"`
}

func (r *ModifyResourceUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleProjectRequestParams struct {
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// Mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// WeCom app user ID.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`
}

type ModifyUserRoleProjectRequest struct {
	*tchttp.BaseRequest
	
	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// Mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// WeCom app user ID.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`
}

func (r *ModifyUserRoleProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserId")
	delete(f, "RoleIdList")
	delete(f, "Email")
	delete(f, "UserName")
	delete(f, "AppUserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleProjectResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserRoleProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserRoleProjectResponseParams `json:"Response"`
}

func (r *ModifyUserRoleProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleRequestParams struct {
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// Mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Mobile number.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Telephone country code.
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// WeCom user ID.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// Whether to enable mobile phone verification code login (0: disabled, 1: enabled).
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// Whether to enable password expiration reminder (0: disabled, 1: enabled).
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// Force password reset (0: disabled, 1: enabled).
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// Password expiration reminder period: 30, 60, 90 (default), or 180 days.
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`
}

type ModifyUserRoleRequest struct {
	*tchttp.BaseRequest
	
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Role ID list.
	RoleIdList []*int64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// Mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Mobile number.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Telephone country code.
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// WeCom user ID.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// Whether to enable mobile phone verification code login (0: disabled, 1: enabled).
	LoginSecurityStatus *int64 `json:"LoginSecurityStatus,omitnil,omitempty" name:"LoginSecurityStatus"`

	// Whether to enable password expiration reminder (0: disabled, 1: enabled).
	ResetPassWordTip *int64 `json:"ResetPassWordTip,omitnil,omitempty" name:"ResetPassWordTip"`

	// Force password reset (0: disabled, 1: enabled).
	ForceResetPassWord *int64 `json:"ForceResetPassWord,omitnil,omitempty" name:"ForceResetPassWord"`

	// Password expiration reminder period: 30, 60, 90 (default), or 180 days.
	PasswordExpired *int64 `json:"PasswordExpired,omitnil,omitempty" name:"PasswordExpired"`
}

func (r *ModifyUserRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "RoleIdList")
	delete(f, "Email")
	delete(f, "UserName")
	delete(f, "PhoneNumber")
	delete(f, "AreaCode")
	delete(f, "AppUserId")
	delete(f, "LoginSecurityStatus")
	delete(f, "ResetPassWordTip")
	delete(f, "ForceResetPassWord")
	delete(f, "PasswordExpired")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRoleResponseParams struct {
	// Custom error information object.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorInfo *ErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Expansion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserRoleResponseParams `json:"Response"`
}

func (r *ModifyUserRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageScreenListVO struct {
	// Image export type. Valid values: Base64, and URL.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PicType *string `json:"PicType,omitnil,omitempty" name:"PicType"`

	// Image list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*PageScreenVO `json:"List,omitnil,omitempty" name:"List"`

	// Async transaction ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Transaction status.
	// 1: processing; 2: processing successful; 3: processing failed (error content in outer Msg).
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type PageScreenVO struct {
	// Screenshot Base64 or URL.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Component ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WidgetId *string `json:"WidgetId,omitnil,omitempty" name:"WidgetId"`
}

type ParamCreateDTO struct {
	// Parameter name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Default value
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Parameter type, string/datetime/number
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// Format type, yyyy-MM-dd HH:mm:ss.SSS (only time required)
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// Complex type, another format expression, such as YYYY-MM
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComplexType *string `json:"ComplexType,omitnil,omitempty" name:"ComplexType"`

	// Application scope
	// Note: This field may return null, indicating that no valid values can be obtained.
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type PermissionComponent struct {
	// Permission value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// Availability.
	// Valid values:.
	// 
	// - usable.
	// - visible.
	// - disabled: unavailable.
	// - hidden: hide.
	// 
	// Default value: disabled.
	// Example value: disabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncludeType *string `json:"IncludeType,omitnil,omitempty" name:"IncludeType"`

	// Target upgrade version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpgradeVersionType *string `json:"UpgradeVersionType,omitnil,omitempty" name:"UpgradeVersionType"`

	// Supplemental information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// Key for supplementary information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TipsKey *string `json:"TipsKey,omitnil,omitempty" name:"TipsKey"`
}

type PermissionGroup struct {
	// Group name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModuleGroup *string `json:"ModuleGroup,omitnil,omitempty" name:"ModuleGroup"`

	// Permission list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Components []*PermissionComponent `json:"Components,omitnil,omitempty" name:"Components"`
}

type Project struct {
	// Project ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Project logo.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// Project name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Background color of the logo.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ColorCode *string `json:"ColorCode,omitnil,omitempty" name:"ColorCode"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Number of members.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// Number of pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// Last modified report and dashboard names.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastModifyName *string `json:"LastModifyName,omitnil,omitempty" name:"LastModifyName"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Apply *bool `json:"Apply,omitnil,omitempty" name:"Apply"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// ""
	// Note: This field may return null, indicating that no valid values can be obtained.
	Seed *string `json:"Seed,omitnil,omitempty" name:"Seed"`

	// Permission list in the project.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthList []*string `json:"AuthList,omitnil,omitempty" name:"AuthList"`

	// Default dashboard.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PanelScope *string `json:"PanelScope,omitnil,omitempty" name:"PanelScope"`

	// Whether it is managed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsExternalManage *bool `json:"IsExternalManage,omitnil,omitempty" name:"IsExternalManage"`

	// Management platform name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagePlatform *string `json:"ManagePlatform,omitnil,omitempty" name:"ManagePlatform"`

	// Customization parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConfigList []*ProjectConfigList `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUserName *string `json:"CreatedUserName,omitnil,omitempty" name:"CreatedUserName"`

	// Associated person ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// Associated person.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// Number of dashboard pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NormalCount *int64 `json:"NormalCount,omitnil,omitempty" name:"NormalCount"`

	// Number of free canvas pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FreeCount *int64 `json:"FreeCount,omitnil,omitempty" name:"FreeCount"`

	// Number of ad-hoc analysis pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdhocCount *int64 `json:"AdhocCount,omitnil,omitempty" name:"AdhocCount"`

	// Number of pages in the briefing
	// Note: This field may return null, indicating that no valid values can be obtained.
	BriefingCount *int64 `json:"BriefingCount,omitnil,omitempty" name:"BriefingCount"`
}

type ProjectConfigList struct {
	// Module group.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModuleGroup *string `json:"ModuleGroup,omitnil,omitempty" name:"ModuleGroup"`

	// Content.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Components []*ProjectConfigResult `json:"Components,omitnil,omitempty" name:"Components"`
}

type ProjectConfigResult struct {
	// Configuration name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// Configuration mode.
	// Valid values:.
	// 
	// - usable.
	// - visible.
	// - disabled: unavailable.
	// - hidden: hide.
	// 
	// Default value: disabled.
	// Example value: disabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IncludeType *string `json:"IncludeType,omitnil,omitempty" name:"IncludeType"`

	// Additional parameters.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type ProjectListData struct {
	// Array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*Project `json:"List,omitnil,omitempty" name:"List"`

	// Total number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPages *uint64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`
}

type ResourceItem struct {
	// Resource name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// resource value
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceValue *bool `json:"ResourceValue,omitnil,omitempty" name:"ResourceValue"`

	// Changeable
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanChange *bool `json:"CanChange,omitnil,omitempty" name:"CanChange"`

	// Prompt message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type RowColumnConfig struct {
	// Row column permission rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	RulerInfo *string `json:"RulerInfo,omitnil,omitempty" name:"RulerInfo"`

	// Tag value list
	// Note: This field may return null, indicating that no valid values can be obtained.
	TagValueList []*RowColumnTagValue `json:"TagValueList,omitnil,omitempty" name:"TagValueList"`
}

type RowColumnTagValue struct {
	// Tag ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Tag name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag value list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type TableColumn struct {
	// Column name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// alias name
	// Note: This field may return null, indicating that no valid values can be obtained.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// Column type
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Segment type
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// Remarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// excel name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExcelName *string `json:"ExcelName,omitnil,omitempty" name:"ExcelName"`

	// Associated dictionary table Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	DictId *int64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// Associated dictionary table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`

	// Join tables and add fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// Table name to which the field belongs
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Complex format of the target set by the user
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldComplexType *string `json:"FieldComplexType,omitnil,omitempty" name:"FieldComplexType"`

	// format rule
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// Whether to filter empty data fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFilter *bool `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// Compute field type
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcType *string `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Formula content of the calculated field
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcFormula *string `json:"CalcFormula,omitnil,omitempty" name:"CalcFormula"`

	// Chinese formula content of the calculated field
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcDesc *string `json:"CalcDesc,omitnil,omitempty" name:"CalcDesc"`

	// Api data source json path name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonPathName *string `json:"JsonPathName,omitnil,omitempty" name:"JsonPathName"`

	// Geographic type identifier
	// Note: This field may return null, indicating that no valid values can be obtained.
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// Custom map Id
	// Note: This field may return null, indicating that no valid values can be obtained.
	GeoJsonId *uint64 `json:"GeoJsonId,omitnil,omitempty" name:"GeoJsonId"`

	// Style configuration for null value display
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmptyValueConfig *EmptyValueConfig `json:"EmptyValueConfig,omitnil,omitempty" name:"EmptyValueConfig"`

	// Original column name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbFieldName *string `json:"DbFieldName,omitnil,omitempty" name:"DbFieldName"`

	// Whether to copy field operation
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCopyOperation *bool `json:"IsCopyOperation,omitnil,omitempty" name:"IsCopyOperation"`

	// Whether to copy from common fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCopyFromNormal *bool `json:"IsCopyFromNormal,omitnil,omitempty" name:"IsCopyFromNormal"`
}

type TableColumnListData struct {
	// Column list in the table
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*TableColumn `json:"List,omitnil,omitempty" name:"List"`

	// async transaction id
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// Async transaction status
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranStatus *int64 `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`
}

type TableField struct {
	// Field name in the db
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// bi display name
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// Field type in the db
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// Abstract field types after BI categorization, such as string, number, time
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// Complex detail type generated after calculation formula of combination of fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	FieldComplexType *string `json:"FieldComplexType,omitnil,omitempty" name:"FieldComplexType"`

	// Field description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Field calculation formula
	// Note: This field may return null, indicating that no valid values can be obtained.
	FormatRule *string `json:"FormatRule,omitnil,omitempty" name:"FormatRule"`

	// Whether to filter empty data fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsFilter *bool `json:"IsFilter,omitnil,omitempty" name:"IsFilter"`

	// Compute field type
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcType *string `json:"CalcType,omitnil,omitempty" name:"CalcType"`

	// Formula content of the calculated field
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcFormula *string `json:"CalcFormula,omitnil,omitempty" name:"CalcFormula"`

	// Chinese formula content of the calculated field, displayed on the front-end
	// Note: This field may return null, indicating that no valid values can be obtained.
	CalcDesc *string `json:"CalcDesc,omitnil,omitempty" name:"CalcDesc"`

	// Associate dictionary table id
	// Note: This field may return null, indicating that no valid values can be obtained.
	DictId *uint64 `json:"DictId,omitnil,omitempty" name:"DictId"`

	// Associate dictionary table name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DictName *string `json:"DictName,omitnil,omitempty" name:"DictName"`

	// Optional, join tables to add field
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableNodeId *string `json:"TableNodeId,omitnil,omitempty" name:"TableNodeId"`

	// excel
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExcelName *string `json:"ExcelName,omitnil,omitempty" name:"ExcelName"`

	// Optional, join tables to add field, name of the table the field belongs to
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// api data source path name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JsonPathName *string `json:"JsonPathName,omitnil,omitempty" name:"JsonPathName"`

	// Geographic field identifier
	// Note: This field may return null, indicating that no valid values can be obtained.
	Granularity *string `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// Map id
	// Note: This field may return null, indicating that no valid values can be obtained.
	GeoJsonId *uint64 `json:"GeoJsonId,omitnil,omitempty" name:"GeoJsonId"`

	// Style configuration for null value display
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmptyValueConfig *EmptyValueConfig `json:"EmptyValueConfig,omitnil,omitempty" name:"EmptyValueConfig"`

	// Original column name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbFieldName *string `json:"DbFieldName,omitnil,omitempty" name:"DbFieldName"`

	// Whether to copy field operation
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCopyOperation *bool `json:"IsCopyOperation,omitnil,omitempty" name:"IsCopyOperation"`

	// Whether to copy from common fields
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsCopyFromNormal *bool `json:"IsCopyFromNormal,omitnil,omitempty" name:"IsCopyFromNormal"`
}

type UserGroupDTO struct {
	// id
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// User group name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// Parent node ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParentId *int64 `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// Whether it is default.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// Administrator user ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AdminUserId *string `json:"AdminUserId,omitnil,omitempty" name:"AdminUserId"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Location.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Location *int64 `json:"Location,omitnil,omitempty" name:"Location"`
}

type UserIdAndUserName struct {
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Enterprise ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// Email.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Last login time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastLogin *string `json:"LastLogin,omitnil,omitempty" name:"LastLogin"`

	// User status.
	// Valid values:.
	// 
	// - 1: Enable.
	// - 0: disabled.
	// 
	// The default value is 1.
	// Example value: 1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to change password on first login.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FirstModify *int64 `json:"FirstModify,omitnil,omitempty" name:"FirstModify"`

	// Mobile number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Telephone country code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Modifier.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// Change time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// System global role.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GlobalUserName *string `json:"GlobalUserName,omitnil,omitempty" name:"GlobalUserName"`

	// System global role code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GlobalUserCode *string `json:"GlobalUserCode,omitnil,omitempty" name:"GlobalUserCode"`

	// Mobile number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserAliasName *string `json:"AppUserAliasName,omitnil,omitempty" name:"AppUserAliasName"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	InValidateAppRange *bool `json:"InValidateAppRange,omitnil,omitempty" name:"InValidateAppRange"`

	//  -1: No activation required. 0: Inactivated. 1: Activated. Null value represents pending binding.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmailActivationStatus *int64 `json:"EmailActivationStatus,omitnil,omitempty" name:"EmailActivationStatus"`

	// 1
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type UserInfo struct {
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Email.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Mobile number.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Telephone country code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// WeCom account ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// WeCom account name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`
}

type UserResourceDTO struct {
	// Enterprise ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// User ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Resource list
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResourceList []*ResourceItem `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`
}

type UserRoleListData struct {
	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Total number of pages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// List.
	// Note: This field may return null, indicating that no valid values can be obtained.
	List []*UserRoleListDataUserRoleInfo `json:"List,omitnil,omitempty" name:"List"`
}

type UserRoleListDataRoleInfo struct {
	// Role Name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// Role ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// Project ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Project name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Whether it is a global role (0: no; 1: yes).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScopeType *int64 `json:"ScopeType,omitnil,omitempty" name:"ScopeType"`

	// Role key.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ModuleCollection *string `json:"ModuleCollection,omitnil,omitempty" name:"ModuleCollection"`
}

type UserRoleListDataUserRoleInfo struct {
	// Business ID.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// This API is used to obtain the role list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleList []*UserRoleListDataRoleInfo `json:"RoleList,omitnil,omitempty" name:"RoleList"`

	// Role ID list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoleIdList []*uint64 `json:"RoleIdList,omitnil,omitempty" name:"RoleIdList"`

	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Enterprise ID.
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// Email.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Creator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedUser *string `json:"CreatedUser,omitnil,omitempty" name:"CreatedUser"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Updater.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedUser *string `json:"UpdatedUser,omitnil,omitempty" name:"UpdatedUser"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Last login time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastLogin *string `json:"LastLogin,omitnil,omitempty" name:"LastLogin"`

	// Account status.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Mobile number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// Telephone country code.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AreaCode *string `json:"AreaCode,omitnil,omitempty" name:"AreaCode"`

	// Whether it is the root account.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootAccount *bool `json:"RootAccount,omitnil,omitempty" name:"RootAccount"`

	// Whether it is an enterprise administrator.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CorpAdmin *bool `json:"CorpAdmin,omitnil,omitempty" name:"CorpAdmin"`

	// WeCom user ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserId *string `json:"AppUserId,omitnil,omitempty" name:"AppUserId"`

	// Nickname.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserAliasName *string `json:"AppUserAliasName,omitnil,omitempty" name:"AppUserAliasName"`

	// Application username.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppUserName *string `json:"AppUserName,omitnil,omitempty" name:"AppUserName"`

	// Whether it is within the visible range.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InValidateAppRange *bool `json:"InValidateAppRange,omitnil,omitempty" name:"InValidateAppRange"`

	// User openID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AppOpenUserId *string `json:"AppOpenUserId,omitnil,omitempty" name:"AppOpenUserId"`

	// Activation status of email.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmailActivationStatus *int64 `json:"EmailActivationStatus,omitnil,omitempty" name:"EmailActivationStatus"`

	// User group information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserGroupList []*UserGroupDTO `json:"UserGroupList,omitnil,omitempty" name:"UserGroupList"`
}

type WidgetListVO struct {
	// uin
	// Note: This field may return null, indicating that no valid values can be obtained.
	CorpId *string `json:"CorpId,omitnil,omitempty" name:"CorpId"`

	// Project ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Page ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// Component array.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WidgetList []*WidgetVO `json:"WidgetList,omitnil,omitempty" name:"WidgetList"`
}

type WidgetVO struct {
	// Component ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WidgetId *string `json:"WidgetId,omitnil,omitempty" name:"WidgetId"`

	// Component name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WidgetName *string `json:"WidgetName,omitnil,omitempty" name:"WidgetName"`
}