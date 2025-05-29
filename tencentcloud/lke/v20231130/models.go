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

package v20231130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AICallConfig struct {

	EnableVoiceInteract *bool `json:"EnableVoiceInteract,omitnil,omitempty" name:"EnableVoiceInteract"`


	EnableVoiceCall *bool `json:"EnableVoiceCall,omitnil,omitempty" name:"EnableVoiceCall"`


	EnableDigitalHuman *bool `json:"EnableDigitalHuman,omitnil,omitempty" name:"EnableDigitalHuman"`


	Voice *VoiceConfig `json:"Voice,omitnil,omitempty" name:"Voice"`


	DigitalHuman *DigitalHumanConfig `json:"DigitalHuman,omitnil,omitempty" name:"DigitalHuman"`
}

type AgentDebugInfo struct {

	Input *string `json:"Input,omitnil,omitempty" name:"Input"`


	Output *string `json:"Output,omitnil,omitempty" name:"Output"`
}

type AgentProcedure struct {
	// Index.
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// English name of the execution process.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Chinese name for display.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Status constant: processing; success; failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Icon.
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// Agent debugging information.
	Debugging *AgentProcedureDebugging `json:"Debugging,omitnil,omitempty" name:"Debugging"`

	// Whether to switch to Agent. The value can be "main" or "workflow". Leave it blank for no switch.
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Current request execution time, in milliseconds.
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// Workflow node name.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Used to display in which reply bubble the thoughts are placed.
	ReplyIndex *uint64 `json:"ReplyIndex,omitnil,omitempty" name:"ReplyIndex"`

	// Main agent.
	SourceAgentName *string `json:"SourceAgentName,omitnil,omitempty" name:"SourceAgentName"`

	// Registration agent.
	TargetAgentName *string `json:"TargetAgentName,omitnil,omitempty" name:"TargetAgentName"`

	// Icon of Agent.
	AgentIcon *string `json:"AgentIcon,omitnil,omitempty" name:"AgentIcon"`
}

type AgentProcedureDebugging struct {
	// Model thinking content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// The specific text content to be displayed.
	DisplayContent *string `json:"DisplayContent,omitnil,omitempty" name:"DisplayContent"`

	// Display type.
	DisplayType *uint64 `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`

	// Index displayed by the search engine.
	QuoteInfos []*QuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// Specific reference source.
	References []*AgentReference `json:"References,omitnil,omitempty" name:"References"`

	// Display the ongoing status.
	DisplayStatus *string `json:"DisplayStatus,omitnil,omitempty" name:"DisplayStatus"`

	// URL of cloud desktop.
	SandboxUrl *string `json:"SandboxUrl,omitnil,omitempty" name:"SandboxUrl"`

	// URL opened through the browser in cloud desktop.
	DisplayUrl *string `json:"DisplayUrl,omitnil,omitempty" name:"DisplayUrl"`
}

type AgentReference struct {
	// Source document ID.
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// URL.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Document business ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Document name.
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`

	// Q&A business ID.
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// Index of search engine.
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// Title.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type AgentThought struct {
	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Request ID
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// It corresponds to a session, session ID, used for message storage in answering. It can be generated in advance and used when saving messages.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Current request execution time, in milliseconds.
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// Whether it is a workflow currently.
	IsWorkflow *bool `json:"IsWorkflow,omitnil,omitempty" name:"IsWorkflow"`

	// If it is a workflow, workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Detailed thinking process.
	Procedures []*AgentProcedure `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// TraceId.
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// File information
	Files []*FileInfo `json:"Files,omitnil,omitempty" name:"Files"`
}

type ApiVarAttrInfo struct {
	// Custom variable ID.
	ApiVarId *string `json:"ApiVarId,omitnil,omitempty" name:"ApiVarId"`

	// Label ID.
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`
}

type AppConfig struct {
	// Knowledge Q&A management application configuration
	KnowledgeQa *KnowledgeQaConfig `json:"KnowledgeQa,omitnil,omitempty" name:"KnowledgeQa"`

	// Knowledge summary application configuration.
	Summary *SummaryConfig `json:"Summary,omitnil,omitempty" name:"Summary"`

	// Label extraction application configuration.
	Classify *ClassifyConfig `json:"Classify,omitnil,omitempty" name:"Classify"`
}

type AppInfo struct {
	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// Application type description.
	AppTypeDesc *string `json:"AppTypeDesc,omitnil,omitempty" name:"AppTypeDesc"`

	// Application ID.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Application name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Application avatar.
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// Application description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Application status. 1: offline; 2: running; 3: disabled.
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// Status description.
	AppStatusDesc *string `json:"AppStatusDesc,omitnil,omitempty" name:"AppStatusDesc"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Last modifier.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Model name.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Alias of the generative model.
	ModelAliasName *string `json:"ModelAliasName,omitnil,omitempty" name:"ModelAliasName"`

	// Application mode: standard; agent; single_workflow.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Alias of the thought model.
	ThoughtModelAliasName *string `json:"ThoughtModelAliasName,omitnil,omitempty" name:"ThoughtModelAliasName"`
}

type AppModel struct {
	// Model name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Model description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// The round referenced by the context.
	ContextLimit *uint64 `json:"ContextLimit,omitnil,omitempty" name:"ContextLimit"`

	// Model alias.
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// Remaining token quota.
	TokenBalance *float64 `json:"TokenBalance,omitnil,omitempty" name:"TokenBalance"`

	// Whether to use the round referenced by the context.
	IsUseContext *bool `json:"IsUseContext,omitnil,omitempty" name:"IsUseContext"`

	// The number of context memory rounds.
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// Usage type.
	UsageType *string `json:"UsageType,omitnil,omitempty" name:"UsageType"`

	// Model temperature.
	Temperature *string `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// Model TopP.
	TopP *string `json:"TopP,omitnil,omitempty" name:"TopP"`

	// Model resource status: 1: available; 2: exhausted.
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type AttrLabel struct {
	// Label source.
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Label ID.
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// Label identification.
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// Label name.
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// Label value.
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type AttrLabelDetail struct {
	// Label ID.
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// Label identification.
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// Label name.
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// Label value name.
	LabelNames []*string `json:"LabelNames,omitnil,omitempty" name:"LabelNames"`

	// Whether the label is being updated.
	IsUpdating *bool `json:"IsUpdating,omitnil,omitempty" name:"IsUpdating"`

	// Status.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Total number of label values.
	LabelTotalCount *string `json:"LabelTotalCount,omitnil,omitempty" name:"LabelTotalCount"`
}

type AttrLabelRefer struct {
	// Label source, 1: label.
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Label ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Label value ID.
	LabelBizIds []*string `json:"LabelBizIds,omitnil,omitempty" name:"LabelBizIds"`
}

type AttributeLabel struct {
	// Standard word ID.
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// Standard word name.
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// Synonym name.
	SimilarLabels []*string `json:"SimilarLabels,omitnil,omitempty" name:"SimilarLabels"`
}

type BaseConfig struct {
	// Application name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Application avatar url, required as an input parameter in "CreateApp" and "ModifyApp". Description of input parameter: 1. The image of input URL must be jpeg or png, with a size no more than 500kb, and the URL must allow head requests. 2. If the user does not have object storage, use the "Obtain temporary file upload key" (DescribeStorageCredential) API to obtain the COS temporary key and upload path. Upload the avatar to COS by yourself and get the access link.
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// Application description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CallDetail struct {
	// Associated ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Call time.
	CallTime *string `json:"CallTime,omitnil,omitempty" name:"CallTime"`

	// Total token consumption.
	TotalTokenUsage *float64 `json:"TotalTokenUsage,omitnil,omitempty" name:"TotalTokenUsage"`

	// Token consumption for input.
	InputTokenUsage *float64 `json:"InputTokenUsage,omitnil,omitempty" name:"InputTokenUsage"`

	// Token consumption for output.
	OutputTokenUsage *float64 `json:"OutputTokenUsage,omitnil,omitempty" name:"OutputTokenUsage"`

	// Number of search service calls.
	SearchUsage *uint64 `json:"SearchUsage,omitnil,omitempty" name:"SearchUsage"`

	// Model name.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Call type.
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Account.
	UinAccount *string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Application name.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Total number of consumed pages.
	PageUsage *uint64 `json:"PageUsage,omitnil,omitempty" name:"PageUsage"`

	// Filter sub-scenario.
	SubScene *string `json:"SubScene,omitnil,omitempty" name:"SubScene"`
}

type ClassifyConfig struct {
	// Model configuration.
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// Label list.
	Labels []*ClassifyLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Welcome words, within 200 characters.
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
}

type ClassifyLabel struct {
	// Label name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Label description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Label value range.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Context struct {
	// Message record ID.
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// Whether it is a user.
	IsVisitor *bool `json:"IsVisitor,omitnil,omitempty" name:"IsVisitor"`

	// Nickname.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Avatar.
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// Message content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Document information.
	FileInfos []*MsgFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// Response method, 15: clarification confirmation response.
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`
}

// Predefined struct for user
type CreateCorpRequestParams struct {
	// Full name of the corporate.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Contact person's name.
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// Contact person's mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Contact person's mobile phone number.
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`
}

type CreateCorpRequest struct {
	*tchttp.BaseRequest
	
	// Full name of the corporate.
	FullName *string `json:"FullName,omitnil,omitempty" name:"FullName"`

	// Contact person's name.
	ContactName *string `json:"ContactName,omitnil,omitempty" name:"ContactName"`

	// Contact person's mailbox.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Contact person's mobile phone number.
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`
}

func (r *CreateCorpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FullName")
	delete(f, "ContactName")
	delete(f, "Email")
	delete(f, "Telephone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCorpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCorpResponseParams struct {
	// Corporate ID.
	CorpBizId *string `json:"CorpBizId,omitnil,omitempty" name:"CorpBizId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCorpResponse struct {
	*tchttp.BaseResponse
	Response *CreateCorpResponseParams `json:"Response"`
}

func (r *CreateCorpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCorpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateReconstructDocumentFlowConfig struct {
	// The returned form of a table in a markdown file: 
	// 0: the table is returned in MD format;
	// 1: the table is returned in HTML form.
	// The default is 1.
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// The format of smart document parsing results:
	// 0: only return full-text MD;
	// 1: only return OCR original JSON of each page;.
	// 2: only return MD of each page;
	// 3: return full-text MD + OCR original JSON of each page;.
	// 4: return full-text MD + MD of each page.
	// The default value is 3 (return full-text MD + OCR original JSON of each page).
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`
}

// Predefined struct for user
type CreateReconstructDocumentFlowRequestParams struct {
	// File type. Supported file types: pdf, doc, docx, ppt, pptx, md, txt, xls, xlsx, csv, png, jpg, jpeg, bmp, gif, webp, heic, eps, icns, im, pcx, ppm, tiff, xbm, heif, jp2.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The base64 value of the file. File size limit: the downloaded file does not exceed 8m after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// <p>File URL. The file download time does not exceed 15 seconds. Supported image pixels: the length of a single side is between 20-10000px. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. The downloaded file shall not exceed the supported file size after Base64 encoding: </p><table> <tbody> <tr> <td>File Type</td> <td>Supported File Size</td> </tr> <tr> <td>PDF</td> <td>200M</td> </tr> <tr> <td>DOC</td> <td>200M</td> </tr> <tr> <td>DOCX</td> <td>200M</td> </tr> <tr> <td>PPT</td> <td>200M</td> </tr> <tr> <td>PPTX</td> <td>200M</td> </tr> <tr> <td>MD</td> <td>10M</td> </tr> <tr> <td>TXT</td> <td>10M</td> </tr> <tr> <td>XLS</td> <td>20M</td> </tr> <tr> <td>XLSX</td> <td>20M</td> </tr> <tr> <td>CSV</td> <td>20M</td> </tr> <tr> <td>PNG</td> <td>20M</td> </tr> <tr> <td>JPG</td> <td>20M</td> </tr> <tr> <td>JPEG</td> <td>20M</td> </tr> <tr> <td>BMP</td> <td>20M</td> </tr> <tr> <td>GIF</td> <td>20M</td> </tr> <tr> <td>WEBP</td> <td>20M</td> </tr> <tr> <td>HEIC</td> <td>20M</td> </tr> <tr> <td>EPS</td> <td>20M</td> </tr> <tr> <td>ICNS</td> <td>20M</td> </tr> <tr> <td>IM</td> <td>20M</td> </tr> <tr> <td>PCX</td> <td>20M</td> </tr> <tr> <td>PPM</td> <td>20M</td> </tr> <tr> <td>TIFF</td> <td>20M</td> </tr> <tr> <td>XBM</td> <td>20M</td> </tr> <tr> <td>HEIF</td> <td>20M</td> </tr> <tr> <td>JP2</td> <td>20M</td> </tr> </tbody> <colgroup> <col> <col> </colgroup></table>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// When type of the uploaded file is pdf, doc, docx, ppt, or pptx, it specifies the starting page number for file recognition, including the current value. The default is 1, indicating recognition starts from the first page of the file.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// When type of the uploaded file is pdf, doc, docx, orppt, pptx, it specifies the end page number for file recognition, including the current value. The default is 100, indicating recognition up to page 100 of the file. a single call supports recognition of up to 1000 pages, i.e., FileEndPageNumber-FileStartPageNumber should be no more than 1000.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Configuration information for creating a document parsing task.
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateReconstructDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// File type. Supported file types: pdf, doc, docx, ppt, pptx, md, txt, xls, xlsx, csv, png, jpg, jpeg, bmp, gif, webp, heic, eps, icns, im, pcx, ppm, tiff, xbm, heif, jp2.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The base64 value of the file. File size limit: the downloaded file does not exceed 8m after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// <p>File URL. The file download time does not exceed 15 seconds. Supported image pixels: the length of a single side is between 20-10000px. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. The downloaded file shall not exceed the supported file size after Base64 encoding: </p><table> <tbody> <tr> <td>File Type</td> <td>Supported File Size</td> </tr> <tr> <td>PDF</td> <td>200M</td> </tr> <tr> <td>DOC</td> <td>200M</td> </tr> <tr> <td>DOCX</td> <td>200M</td> </tr> <tr> <td>PPT</td> <td>200M</td> </tr> <tr> <td>PPTX</td> <td>200M</td> </tr> <tr> <td>MD</td> <td>10M</td> </tr> <tr> <td>TXT</td> <td>10M</td> </tr> <tr> <td>XLS</td> <td>20M</td> </tr> <tr> <td>XLSX</td> <td>20M</td> </tr> <tr> <td>CSV</td> <td>20M</td> </tr> <tr> <td>PNG</td> <td>20M</td> </tr> <tr> <td>JPG</td> <td>20M</td> </tr> <tr> <td>JPEG</td> <td>20M</td> </tr> <tr> <td>BMP</td> <td>20M</td> </tr> <tr> <td>GIF</td> <td>20M</td> </tr> <tr> <td>WEBP</td> <td>20M</td> </tr> <tr> <td>HEIC</td> <td>20M</td> </tr> <tr> <td>EPS</td> <td>20M</td> </tr> <tr> <td>ICNS</td> <td>20M</td> </tr> <tr> <td>IM</td> <td>20M</td> </tr> <tr> <td>PCX</td> <td>20M</td> </tr> <tr> <td>PPM</td> <td>20M</td> </tr> <tr> <td>TIFF</td> <td>20M</td> </tr> <tr> <td>XBM</td> <td>20M</td> </tr> <tr> <td>HEIF</td> <td>20M</td> </tr> <tr> <td>JP2</td> <td>20M</td> </tr> </tbody> <colgroup> <col> <col> </colgroup></table>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// When type of the uploaded file is pdf, doc, docx, ppt, or pptx, it specifies the starting page number for file recognition, including the current value. The default is 1, indicating recognition starts from the first page of the file.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// When type of the uploaded file is pdf, doc, docx, orppt, pptx, it specifies the end page number for file recognition, including the current value. The default is 100, indicating recognition up to page 100 of the file. a single call supports recognition of up to 1000 pages, i.e., FileEndPageNumber-FileStartPageNumber should be no more than 1000.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Configuration information for creating a document parsing task.
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreateReconstructDocumentFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReconstructDocumentFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "FileBase64")
	delete(f, "FileUrl")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReconstructDocumentFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReconstructDocumentFlowResponseParams struct {
	// Unique task ID. The processing result corresponding to TaskId can be queried through the API [GetReconstructDocumentResult](https://cloud.tencent.com/document/product/1759/107505) within 30 days.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReconstructDocumentFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateReconstructDocumentFlowResponseParams `json:"Response"`
}

func (r *CreateReconstructDocumentFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReconstructDocumentFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseRequestParams struct {
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CreateReleaseRequest struct {
	*tchttp.BaseRequest
	
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

func (r *CreateReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseResponseParams struct {
	// Release ID.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseResponseParams `json:"Response"`
}

func (r *CreateReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppRequestParams struct {
	// Application ID.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DeleteAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppResponseParams `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocRequestParams struct {
	// List of document business IDs.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DeleteDocRequest struct {
	*tchttp.BaseRequest
	
	// List of document business IDs.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DeleteDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocBizIds")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocResponseParams `json:"Response"`
}

func (r *DeleteDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type DeleteQACateRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *DeleteQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQACateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQACateResponseParams `json:"Response"`
}

func (r *DeleteQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQARequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

type DeleteQARequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

func (r *DeleteQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQAResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQAResponseParams `json:"Response"`
}

func (r *DeleteQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// The unique ID of the data source for the rejected question.
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

type DeleteRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// The unique ID of the data source for the rejected question.
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

func (r *DeleteRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "RejectedBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRejectedQuestionResponseParams `json:"Response"`
}

func (r *DeleteRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Attribute ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Quantity loaded each time. 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log in to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Log in to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Query a label or similar labels.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// The label ID of the scroll loading cursor.
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`

	// Query scope: 
	// all (or leave it blank): standard words and similar words 
	// standard: standard words 
	// similar: similar words
	QueryScope *string `json:"QueryScope,omitnil,omitempty" name:"QueryScope"`
}

type DescribeAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Attribute ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Quantity loaded each time. 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Log in to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Log in to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Query a label or similar labels.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// The label ID of the scroll loading cursor.
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`

	// Query scope: 
	// all (or leave it blank): standard words and similar words 
	// standard: standard words 
	// similar: similar words
	QueryScope *string `json:"QueryScope,omitnil,omitempty" name:"QueryScope"`
}

func (r *DescribeAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "Limit")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "LastLabelBizId")
	delete(f, "QueryScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelResponseParams struct {
	// Attribute ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Attribute identifier.
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// Attribute name.
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// Quantity of labels.
	LabelNumber *string `json:"LabelNumber,omitnil,omitempty" name:"LabelNumber"`

	// Label name.
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttributeLabelResponseParams `json:"Response"`
}

func (r *DescribeAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallStatsGraphRequestParams struct {
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Log in to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Log in to user's root sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Filter sub-scenarios (used in document parsing scenarios).
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

type DescribeCallStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Log in to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Log in to user's root sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Filter sub-scenarios (used in document parsing scenarios).
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

func (r *DescribeCallStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	delete(f, "SubScenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallStatsGraphResponseParams struct {
	// Statistical information of API calls.
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeCallStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphRequestParams struct {
	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphResponseParams struct {
	// X-axis: time zone; returns two interval ranges of "minute/hour/day" according to the granularity of query conditions.
	X []*string `json:"X,omitnil,omitempty" name:"X"`

	// Available concurrent Y-axis coordinate.
	AvailableY []*int64 `json:"AvailableY,omitnil,omitempty" name:"AvailableY"`

	// Succeeded to call the concurrent Y-axis coordinate.
	SuccessCallY []*int64 `json:"SuccessCallY,omitnil,omitempty" name:"SuccessCallY"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageRequestParams struct {
	// Model identification.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageRequest struct {
	*tchttp.BaseRequest
	
	// Model identification.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageResponseParams struct {
	// The upper limit of available concurrency.
	AvailableConcurrency *uint64 `json:"AvailableConcurrency,omitnil,omitempty" name:"AvailableConcurrency"`

	// Peak concurrent value.
	ConcurrencyPeak *uint64 `json:"ConcurrencyPeak,omitnil,omitempty" name:"ConcurrencyPeak"`

	// The number of times exceeding the capacity limit of available concurrency.
	ExceedUsageTime *uint64 `json:"ExceedUsageTime,omitnil,omitempty" name:"ExceedUsageTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocRequestParams struct {
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *DescribeDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocResponseParams struct {
	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// File name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File type.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// COS path.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Document status : 1: not generated; 2: generating; 3: generation successful; 4: generation failed; 5: deleting; 6: deleted successfully; 7: under review; 8: review failed; 9: review successful; 10: pending release; 11: releasing; 12: released; 13: learning; 14: learning failed; 15: updating; 16: update failed; 17: parsing; 18: parsing failed; 19: import failed; 20: expired; 21: excessive invalid; 22: excessive invalid recovered.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Document status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Reason for generation failure.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Whether to refer in the answer.
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// Number of Q&A pairs.
	QaNum *int64 `json:"QaNum,omitnil,omitempty" name:"QaNum"`

	// Whether to delete.
	IsDeleted *bool `json:"IsDeleted,omitnil,omitempty" name:"IsDeleted"`

	// Document source.
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Document source description.
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// Whether regeneration is allowed.
	IsAllowRestart *bool `json:"IsAllowRestart,omitnil,omitempty" name:"IsAllowRestart"`

	// Whether Q&A has been deleted.
	IsDeletedQa *bool `json:"IsDeletedQa,omitnil,omitempty" name:"IsDeletedQa"`

	// Whether Q&A is being generated.
	IsCreatingQa *bool `json:"IsCreatingQa,omitnil,omitempty" name:"IsCreatingQa"`

	// Whether deletion is allowed.
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// Whether to allow operation reference switch.
	IsAllowRefer *bool `json:"IsAllowRefer,omitnil,omitempty" name:"IsAllowRefer"`

	// Whether Q&A has been generated.
	IsCreatedQa *bool `json:"IsCreatedQa,omitnil,omitempty" name:"IsCreatedQa"`

	// Document character count.
	DocCharSize *string `json:"DocCharSize,omitnil,omitempty" name:"DocCharSize"`

	// Whether editing is allowed.
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// Applicable scope of labels 1: all, 2: by condition range.
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Label.
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDocResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocResponseParams `json:"Response"`
}

func (r *DescribeDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsagePieGraphRequestParams struct {
	// Application id array.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeKnowledgeUsagePieGraphRequest struct {
	*tchttp.BaseRequest
	
	// Application id array.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeKnowledgeUsagePieGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsagePieGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsagePieGraphResponseParams struct {
	// Total number of characters used by all applications.
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// List of application pie chart details.
	List []*KnowledgeCapacityPieGraphDetail `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsagePieGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsagePieGraphResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsagePieGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageRequestParams struct {

}

type DescribeKnowledgeUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeKnowledgeUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageResponseParams struct {
	// The upper limit of available characters.
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// Number of characters exceeding the capacity limit of available characters.
	ExceedCharSize *string `json:"ExceedCharSize,omitnil,omitempty" name:"ExceedCharSize"`

	// Total number of characters used in the knowledge library.
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsageResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release details.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type DescribeReleaseRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release details.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *DescribeReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseResponseParams struct {
	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Publish description.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Release status (1. pending release; 2. releasing; 3. release successful; 4. release failed; 5. releasing (under review); 6. releasing (review completed); 7. release failed; 9. release paused).
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Release status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseResponseParams `json:"Response"`
}

func (r *DescribeReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyRequestParams struct {
	// Application appkey.
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

type DescribeRobotBizIDByAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// Application appkey.
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

func (r *DescribeRobotBizIDByAppKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRobotBizIDByAppKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyResponseParams struct {
	// Application business ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRobotBizIDByAppKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRobotBizIDByAppKeyResponseParams `json:"Response"`
}

func (r *DescribeRobotBizIDByAppKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchStatsGraphRequestParams struct {
	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Uin list.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application id list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeSearchStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Uin list.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Sub-business type.
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application id list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeSearchStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchStatsGraphResponseParams struct {
	// The statistical result.
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeSearchStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageGraphRequestParams struct {
	// Root account of Tencent Cloud.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Sub-business types of Tencent Cloud Agent Development Platform/TCADP: fileparse (document parsing), Embedding, Rewrite (multi-round rewriting), Concurrency, KnowledgeSummary (knowledge summary), KnowledgeQA (knowledge Q&A), KnowledgeCapacity (knowledge base capacity), SearchEngine (search engine).
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeTokenUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// Root account of Tencent Cloud.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Sub-business types of Tencent Cloud Agent Development Platform/TCADP: fileparse (document parsing), Embedding, Rewrite (multi-round rewriting), Concurrency, KnowledgeSummary (knowledge summary), KnowledgeQA (knowledge Q&A), KnowledgeCapacity (knowledge base capacity), SearchEngine (search engine).
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeTokenUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageGraphResponseParams struct {
	// Total amount of token consumption.
	Total []*Stat `json:"Total,omitnil,omitempty" name:"Total"`

	// Input token consumption.
	Input []*Stat `json:"Input,omitnil,omitempty" name:"Input"`

	// Output token consumption.
	Output []*Stat `json:"Output,omitnil,omitempty" name:"Output"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageRequestParams struct {
	// Root account of Tencent Cloud.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Log in to user's root account (required in the integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business types of Tencent Cloud Agent Development Platform/TCADP: FileParse (document parsing), embedding, Rewrite (multi-round rewriting), Concurrency, KnowledgeSummary (knowledge summary), KnowledgeQA (knowledge Q&A), KnowledgeCapacity (knowledge base capacity), SearchEngine (search engine).
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds (default value: 0).
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds (default value: 0, must be greater than the start timestamp).
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Filter sub-scenario (used in document parsing scenario).
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

type DescribeTokenUsageRequest struct {
	*tchttp.BaseRequest
	
	// Root account of Tencent Cloud.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Log in to user's root account (required in the integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Sub-business types of Tencent Cloud Agent Development Platform/TCADP: FileParse (document parsing), embedding, Rewrite (multi-round rewriting), Concurrency, KnowledgeSummary (knowledge summary), KnowledgeQA (knowledge Q&A), KnowledgeCapacity (knowledge base capacity), SearchEngine (search engine).
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start timestamp, in seconds (default value: 0).
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp, in seconds (default value: 0, must be greater than the start timestamp).
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Filter sub-scenario (used in document parsing scenario).
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

func (r *DescribeTokenUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	delete(f, "SubScenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageResponseParams struct {
	// Total token consumption.
	TotalTokenUsage *float64 `json:"TotalTokenUsage,omitnil,omitempty" name:"TotalTokenUsage"`

	// Input token consumption.
	InputTokenUsage *float64 `json:"InputTokenUsage,omitnil,omitempty" name:"InputTokenUsage"`

	// Output token consumption.
	OutputTokenUsage *float64 `json:"OutputTokenUsage,omitnil,omitempty" name:"OutputTokenUsage"`

	// Number of API calls.
	ApiCallStats *uint64 `json:"ApiCallStats,omitnil,omitempty" name:"ApiCallStats"`

	// Number of search service calls.
	SearchUsage *float64 `json:"SearchUsage,omitnil,omitempty" name:"SearchUsage"`

	// Number of pages consumed by document parsing.
	PageUsage *uint64 `json:"PageUsage,omitnil,omitempty" name:"PageUsage"`

	// Token consumption by splitting.
	SplitTokenUsage *float64 `json:"SplitTokenUsage,omitnil,omitempty" name:"SplitTokenUsage"`

	// Number of Rag retrievals.
	RagSearchUsage *float64 `json:"RagSearchUsage,omitnil,omitempty" name:"RagSearchUsage"`

	// Number of online searches.
	InternetSearchUsage *float64 `json:"InternetSearchUsage,omitnil,omitempty" name:"InternetSearchUsage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Response ID.
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// Log in to user's root account (required in the integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DescribeUnsatisfiedReplyContextRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Response ID.
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// Log in to user's root account (required in the integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DescribeUnsatisfiedReplyContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnsatisfiedReplyContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextResponseParams struct {
	// Context of dissatisfied responses.
	List []*Context `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnsatisfiedReplyContextResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnsatisfiedReplyContextResponseParams `json:"Response"`
}

func (r *DescribeUnsatisfiedReplyContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DigitalHumanConfig struct {

	AssetKey *string `json:"AssetKey,omitnil,omitempty" name:"AssetKey"`


	Name *string `json:"Name,omitnil,omitempty" name:"Name"`


	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

// Predefined struct for user
type ExportQAListRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A business ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// Query parameter.
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportQAListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A business ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// Query parameter.
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAListResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportQAListResponse struct {
	*tchttp.BaseResponse
	Response *ExportQAListResponseParams `json:"Response"`
}

func (r *ExportQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Check to export ID list.
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Retrieve filter.
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Check to export ID list.
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Retrieve filter.
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ExportUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ExportUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtraInfo struct {

	EChartsInfo []*string `json:"EChartsInfo,omitnil,omitempty" name:"EChartsInfo"`
}

type FileInfo struct {
	// File name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File size.
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// URL of the file, address of COS.
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// File type.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// DocID returned after parsing.
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// Creation time.
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type Filters struct {
	// Retrieve user question or answer.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Incorrect retrieval type.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type GenerateQARequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

type GenerateQARequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

func (r *GenerateQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateQAResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateQAResponse struct {
	*tchttp.BaseResponse
	Response *GenerateQAResponseParams `json:"Response"`
}

func (r *GenerateQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordRequestParams struct {
	// Type.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Quantity.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Application AppKey. When Type=5 [API Visitor], this field is required. </br> How to obtain it: </br> 1. After the application is released, obtain it at [Application Page]-[Release Management]-[Call Information]-[API Management]. </br> 2. Refer to item 2 in https://cloud.tencent.com/document/product/1759/109469.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Scenario, experience: 1; formal: 2.
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`

	// Last record ID. Retrieve messages in reverse order. You can only select either MidRecordId or LastRecordId.
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// If this value is input, it means pulling a total of count message records before and after the record ID. You can only select either MidRecordId or LastRecordId.
	MidRecordId *string `json:"MidRecordId,omitnil,omitempty" name:"MidRecordId"`
}

type GetMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// Type.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Quantity.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Application AppKey. When Type=5 [API Visitor], this field is required. </br> How to obtain it: </br> 1. After the application is released, obtain it at [Application Page]-[Release Management]-[Call Information]-[API Management]. </br> 2. Refer to item 2 in https://cloud.tencent.com/document/product/1759/109469.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Scenario, experience: 1; formal: 2.
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`

	// Last record ID. Retrieve messages in reverse order. You can only select either MidRecordId or LastRecordId.
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// If this value is input, it means pulling a total of count message records before and after the record ID. You can only select either MidRecordId or LastRecordId.
	MidRecordId *string `json:"MidRecordId,omitnil,omitempty" name:"MidRecordId"`
}

func (r *GetMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Count")
	delete(f, "SessionId")
	delete(f, "BotAppKey")
	delete(f, "Scene")
	delete(f, "LastRecordId")
	delete(f, "MidRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordResponseParams struct {
	// Session record.
	Records []*MsgRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// The time when session cleared associated context, in milliseconds.
	SessionDisassociatedTimestamp *string `json:"SessionDisassociatedTimestamp,omitnil,omitempty" name:"SessionDisassociatedTimestamp"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *GetMsgRecordResponseParams `json:"Response"`
}

func (r *GetMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *GetTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskType")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task type.
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Task status.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Task messages.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Task parameters.
	Params *TaskParams `json:"Params,omitnil,omitempty" name:"Params"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStatusResponseParams `json:"Response"`
}

func (r *GetTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWsTokenReq_Label struct {
	// Label name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Label value.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetWsTokenRequestParams struct {
	// Access type, 5 - API visitor .
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Application AppKey</br>
	// How to Obtain It:</br>
	// 1. After the application is released, obtain it at [Release Management] - [Call Information] - [API Management] on the application page.</br>
	// 2. See the second item in https://cloud.tencent.com/document/product/1759/109469.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Visitor ID (external input, recommended to be unique, identifies the user currently accessing the session). Length Limit: string(64).
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// Knowledge label, used for search filter in the knowledge base. This field is about to be deactivated. Please use the custom_variables field in the dialogue endpoint API to replace this field.
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

type GetWsTokenRequest struct {
	*tchttp.BaseRequest
	
	// Access type, 5 - API visitor .
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Application AppKey</br>
	// How to Obtain It:</br>
	// 1. After the application is released, obtain it at [Release Management] - [Call Information] - [API Management] on the application page.</br>
	// 2. See the second item in https://cloud.tencent.com/document/product/1759/109469.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Visitor ID (external input, recommended to be unique, identifies the user currently accessing the session). Length Limit: string(64).
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// Knowledge label, used for search filter in the knowledge base. This field is about to be deactivated. Please use the custom_variables field in the dialogue endpoint API to replace this field.
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

func (r *GetWsTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "BotAppKey")
	delete(f, "VisitorBizId")
	delete(f, "VisitorLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWsTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWsTokenResponseParams struct {
	// Token value (valid for 60 seconds, valid only once, multiple validations will report an error).
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// Balance. The balance is valid if it is greater than 0.
	Balance *float64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// The character limit for input in the chat window.
	InputLenLimit *int64 `json:"InputLenLimit,omitnil,omitempty" name:"InputLenLimit"`

	// Application mode: standard; agent; single_workflow.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// SingleWorkflow.
	SingleWorkflow *KnowledgeQaSingleWorkflow `json:"SingleWorkflow,omitnil,omitempty" name:"SingleWorkflow"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWsTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetWsTokenResponseParams `json:"Response"`
}

func (r *GetWsTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistorySummary struct {
	// Assistant.
	Assistant *string `json:"Assistant,omitnil,omitempty" name:"Assistant"`

	// User.
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Dissatisfied response ID.
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type IgnoreUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Dissatisfied response ID.
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *IgnoreUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IgnoreUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *IgnoreUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntentAchievement struct {

	Name *string `json:"Name,omitnil,omitempty" name:"Name"`


	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type InvokeAPI struct {
	// Request method, such as get/post.
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// Request address.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Header parameter.
	HeaderValues []*StrValue `json:"HeaderValues,omitnil,omitempty" name:"HeaderValues"`

	// Input parameter Query.
	QueryValues []*StrValue `json:"QueryValues,omitnil,omitempty" name:"QueryValues"`

	// Original data of a Post request.
	RequestPostBody *string `json:"RequestPostBody,omitnil,omitempty" name:"RequestPostBody"`

	// Returned original data.
	ResponseBody *string `json:"ResponseBody,omitnil,omitempty" name:"ResponseBody"`

	// Output parameter.
	ResponseValues []*ValueInfo `json:"ResponseValues,omitnil,omitempty" name:"ResponseValues"`

	// Exception information.
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`
}

type KnowledgeCapacityPieGraphDetail struct {
	// Current application name.
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Number of characters used by the current application.
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// Proportion of the current application in total usage.
	Proportion *float64 `json:"Proportion,omitnil,omitempty" name:"Proportion"`
}

type KnowledgeQaConfig struct {
	// Welcome words, within 200 characters.
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// Character description, within 4,000 characters. By filling out the description, set the #Character Name, #Style Characteristics, and reachable #Intent of the application. It is recommended to fill in according to the following template, with custom intents no more than 5. 
	// #Character Name:
	// #Style Characteristics:
	// #Output Requirements:
	// #Ability Limitations:
	// The following user intents can be reached.
	// ##Intent Name:
	// ##Intent Description:
	// ##Intent Example:
	// ##Intent Implementation:
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`

	// Generative model configuration.
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// Knowledge search configuration.
	Search []*KnowledgeQaSearch `json:"Search,omitnil,omitempty" name:"Search"`

	// Knowledge management output configuration.
	Output *KnowledgeQaOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Workflow configuration.
	Workflow *KnowledgeWorkflow `json:"Workflow,omitnil,omitempty" name:"Workflow"`

	// Retrieval range.
	SearchRange *SearchRange `json:"SearchRange,omitnil,omitempty" name:"SearchRange"`

	// Application modes: standard, agent, single_workflow.
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// Retrieve a policy.
	SearchStrategy *SearchStrategy `json:"SearchStrategy,omitnil,omitempty" name:"SearchStrategy"`

	// Single workflow ID, which is entered when Pattern is single_workflow.
	SingleWorkflow *KnowledgeQaSingleWorkflow `json:"SingleWorkflow,omitnil,omitempty" name:"SingleWorkflow"`

	// Application associated plug-in.
	Plugins []*KnowledgeQaPlugin `json:"Plugins,omitnil,omitempty" name:"Plugins"`

	// Thinking model configuration.
	ThoughtModel *AppModel `json:"ThoughtModel,omitnil,omitempty" name:"ThoughtModel"`

	// Priority of intent achievement methods.
	IntentAchievements []*IntentAchievement `json:"IntentAchievements,omitnil,omitempty" name:"IntentAchievements"`

	// Whether to enable Image-Text Search.
	ImageTextRetrieval *bool `json:"ImageTextRetrieval,omitnil,omitempty" name:"ImageTextRetrieval"`

	// Configure voice call parameters.
	AiCall *AICallConfig `json:"AiCall,omitnil,omitempty" name:"AiCall"`

	// Session content moderation switch. Note: 1. This feature is enabled by default if no value is input. 2. Tencent Cloud provides a content pre-filtering feature to help filter high-risk or illegal content. If you do not use our filtering feature, you can disable it here. We hereby remind you that you are responsible for ensuring that your content complies with platform policies and laws and regulations, and that you should fulfill your content moderation obligations.
	EnableAudit *bool `json:"EnableAudit,omitnil,omitempty" name:"EnableAudit"`
}

type KnowledgeQaOutput struct {
	// Output method, 1: streaming 2: non-streaming.
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// General model response.
	UseGeneralKnowledge *bool `json:"UseGeneralKnowledge,omitnil,omitempty" name:"UseGeneralKnowledge"`

	// Unknown response, within 300 characters.
	BareAnswer *string `json:"BareAnswer,omitnil,omitempty" name:"BareAnswer"`

	// Whether to display the question clarification switch.
	ShowQuestionClarify *bool `json:"ShowQuestionClarify,omitnil,omitempty" name:"ShowQuestionClarify"`

	// Whether to enable question clarification.
	UseQuestionClarify *bool `json:"UseQuestionClarify,omitnil,omitempty" name:"UseQuestionClarify"`

	// List of keywords for question clarification.
	QuestionClarifyKeywords []*string `json:"QuestionClarifyKeywords,omitnil,omitempty" name:"QuestionClarifyKeywords"`

	// Whether to enable recommended questions.
	UseRecommended *bool `json:"UseRecommended,omitnil,omitempty" name:"UseRecommended"`
}

type KnowledgeQaPlugin struct {
	// Plugin ID.
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// Plugin name.
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// Plugin icon.
	PluginIcon *string `json:"PluginIcon,omitnil,omitempty" name:"PluginIcon"`

	// Tool ID.
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// Tool name.
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// Tool description.
	ToolDesc *string `json:"ToolDesc,omitnil,omitempty" name:"ToolDesc"`

	// Tool input parameter.
	Inputs []*PluginToolReqParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// Whether the plugin is bound to the knowledge library.
	IsBindingKnowledge *bool `json:"IsBindingKnowledge,omitnil,omitempty" name:"IsBindingKnowledge"`
}

type KnowledgeQaSearch struct {
	// Knowledge source: doc (document), qa (question & answering), taskflow (business process), search (search enhancement).
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Q&A - reply flexibility, 1: directly reply if the answer has been accepted. 2: reply after the accepted answer has been polished.
	ReplyFlexibility *uint64 `json:"ReplyFlexibility,omitnil,omitempty" name:"ReplyFlexibility"`

	// Search enhancement - search engine status.
	UseSearchEngine *bool `json:"UseSearchEngine,omitnil,omitempty" name:"UseSearchEngine"`

	// Whether to display the search engine retrieval status.
	ShowSearchEngine *bool `json:"ShowSearchEngine,omitnil,omitempty" name:"ShowSearchEngine"`

	// Knowledge source, whether to select.
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// Maximum number of Q&A recalls, defaults to 2, limited to 5.
	QaTopN *uint64 `json:"QaTopN,omitnil,omitempty" name:"QaTopN"`

	// Maximum number of documents recalls, defaults to 3, limited to 5.
	DocTopN *uint64 `json:"DocTopN,omitnil,omitempty" name:"DocTopN"`

	// Retrieval confidence degree, valid for documents and Q&A. Value range: 0.01 - 0.99.
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// Resource status, 1: the resource is available; 2: the resource is exhausted.
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type KnowledgeQaSingleWorkflow struct {
	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Workflow description.
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// Workflow status, publishing status (UNPUBLISHED; PUBLISHING; PUBLISHED; FAIL).
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether to enable workflow.
	IsEnable *bool `json:"IsEnable,omitnil,omitempty" name:"IsEnable"`

	// Whether to enable asynchronous call of the workflow.
	AsyncWorkflow *bool `json:"AsyncWorkflow,omitnil,omitempty" name:"AsyncWorkflow"`
}

type KnowledgeSummary struct {
	// 1: Q&A; 2: document fragment.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Knowledge content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type KnowledgeWorkflow struct {
	// Whether to enable the workflow.
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// Whether to enable PDL.
	UsePdl *bool `json:"UsePdl,omitnil,omitempty" name:"UsePdl"`
}

type Label struct {
	// Label ID.
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// Label name.
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`
}

// Predefined struct for user
type ListAppRequestParams struct {
	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// Number of items per page, integer.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number, integer.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Keywords: application / modifier.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Login to user's sub-account (required in integrator mode).	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListAppRequest struct {
	*tchttp.BaseRequest
	
	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// Number of items per page, integer.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number, integer.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Keywords: application / modifier.
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Login to user's sub-account (required in integrator mode).	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppResponseParams struct {
	// Quantity.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Label list.
	List []*AppInfo `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppResponse struct {
	*tchttp.BaseResponse
	Response *ListAppResponseParams `json:"Response"`
}

func (r *ListAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Quantity per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type ListAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Quantity per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ListAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelResponseParams struct {
	// Total number.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// List.
	List []*AttrLabelDetail `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ListAttributeLabelResponseParams `json:"Response"`
}

func (r *ListAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type ListQACateRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *ListQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateResponseParams struct {
	// List.
	List []*QACate `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQACateResponse struct {
	*tchttp.BaseResponse
	Response *ListQACateResponseParams `json:"Response"`
}

func (r *ListQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type ListRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ListRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionResponseParams struct {
	// Total number.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// List of rejected questions.
	List []*RejectedQuestion `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ListRejectedQuestionResponseParams `json:"Response"`
}

func (r *ListRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewRequestParams struct {
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Release ticket ID.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// Status (1: newly-added; 2: updated; 3: deleted).
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Release status.
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

type ListReleaseConfigPreviewRequest struct {
	*tchttp.BaseRequest
	
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Release ticket ID.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// Status (1: newly-added; 2: updated; 3: deleted).
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Release status.
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

func (r *ListReleaseConfigPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "Actions")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReleaseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseConfigPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewResponseParams struct {
	// Quantity.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// Configuration item list.
	List []*ReleaseConfigs `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseConfigPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseConfigPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseConfigPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Document status: 7: under review; 8: review failed; 10: pending release; 11: releasing; 12: released; 13: learning; 14: learning failed; 20: expired.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListSelectDocRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Document status: 7: under review; 8: review failed; 10: pending release; 11: releasing; 12: released; 13: learning; 14: learning failed; 20: expired.
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListSelectDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSelectDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocResponseParams struct {
	// Dropdown content.
	List []*Option `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSelectDocResponse struct {
	*tchttp.BaseResponse
	Response *ListSelectDocResponseParams `json:"Response"`
}

func (r *ListSelectDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// User request (question or answer).
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Error type retrieval.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type ListUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// User request (question or answer).
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Error type retrieval.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *ListUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyResponseParams struct {
	// Total number.
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// List of dissatisfied responses.
	List []*UnsatisfiedReply `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ListUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ListUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsageCallDetailRequestParams struct {
	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Uin list.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Call type list.
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Filter sub-scenario.
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

type ListUsageCallDetailRequest struct {
	*tchttp.BaseRequest
	
	// Model identifier.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Uin list.
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// Application ID list.
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// Call type list.
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// Filter sub-scenario.
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`
}

func (r *ListUsageCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsageCallDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "UinAccount")
	delete(f, "AppBizIds")
	delete(f, "CallType")
	delete(f, "SubScenes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsageCallDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsageCallDetailResponseParams struct {
	// Total count of lists.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List.
	List []*CallDetail `json:"List,omitnil,omitempty" name:"List"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUsageCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListUsageCallDetailResponseParams `json:"Response"`
}

func (r *ListUsageCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsageCallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// Application ID.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// Basic application configuration.
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// Application configuration.
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// Login to user's sub-account (required in integrator mode).	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Application type; knowledge_qa - knowledge Q&A management; summary - knowledge summary; classifys - knowledge label extraction.
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// Basic application configuration.
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// Application configuration.
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// Login to user's sub-account (required in integrator mode).	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ModifyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	delete(f, "BaseConfig")
	delete(f, "AppConfig")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// Application.
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppResponseParams `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Label ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Label name.
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// Label identifier (abolished).
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Deleted label value.
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// Newly-added or edited label.
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Label ID.
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// Label name.
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// Label identifier (abolished).
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Deleted label value.
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// Newly-added or edited label.
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *ModifyAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "AttrName")
	delete(f, "AttrKey")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "DeleteLabelBizIds")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAttributeLabelResponseParams `json:"Response"`
}

func (r *ModifyAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// Attribute label applicable scope: 1: all, 2: by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Attribute label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyDocAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// Attribute label applicable scope: 1: all, 2: by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Attribute label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyDocAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyDocAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocCateRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyDocCateRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyDocCateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocCateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Name")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocCateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocCateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocCateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocCateResponseParams `json:"Response"`
}

func (r *ModifyDocCateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocCateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Whether to reference a link.
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// Applicable scope of labels: 1: all; 2: by condition.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Associated labels.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// Web page (or custom link) address.
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// External reference link type: 0: system link 1: custom link.
	// When the value is 1, the weburl field cannot be empty; otherwise, it will not take effect.
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp. 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyDocRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Whether to reference a link.
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// Applicable scope of labels: 1: all; 2: by condition.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// Associated labels.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// Web page (or custom link) address.
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// External reference link type: 0: system link 1: custom link.
	// When the value is 1, the weburl field cannot be empty; otherwise, it will not take effect.
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp. 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	delete(f, "IsRefer")
	delete(f, "AttrRange")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "AttrLabels")
	delete(f, "WebUrl")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocResponseParams `json:"Response"`
}

func (r *ModifyDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// Applicable scope of attribute label: 1: all, 2: by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Attribute label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyQAAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// Applicable scope of attribute label: 1: all, 2: by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Attribute label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyQAAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQAAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyQAAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyQACateRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Category name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Name")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQACateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQACateResponseParams `json:"Response"`
}

func (r *ModifyQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQARequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// Question.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Answer.
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// Custom parameter.
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// Applicable scope of labels: 1. all; 2. by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp, 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Similar question modification information (not passed if there is no modification to the similar question).
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`

	// Problem description.
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Q&A ID.
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// Question.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Answer.
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// Custom parameter.
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// Applicable scope of labels: 1. all; 2. by conditions.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp, 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Similar question modification information (not passed if there is no modification to the similar question).
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`

	// Problem description.
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

func (r *ModifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "CustomParam")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	delete(f, "DocBizId")
	delete(f, "CateBizId")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "SimilarQuestionModify")
	delete(f, "QuestionDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAResponseParams `json:"Response"`
}

func (r *ModifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Rejected question.
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Unique id of the data source for the rejected question source.
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

type ModifyRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Rejected question.
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Unique id of the data source for the rejected question source.
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

func (r *ModifyRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "RejectedBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRejectedQuestionResponseParams `json:"Response"`
}

func (r *ModifyRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgFileInfo struct {
	// Document name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Document size.
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// Document URL.
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// Document type.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// Document ID.
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type MsgRecord struct {
	// Content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// The Session ID corresponding to the current record.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Record ID.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Associated record ID.
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// Whether it is from oneself.
	IsFromSelf *bool `json:"IsFromSelf,omitnil,omitempty" name:"IsFromSelf"`

	// Sender name.
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// Avatar of the sender.
	FromAvatar *string `json:"FromAvatar,omitnil,omitempty" name:"FromAvatar"`

	// Timestamp.
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Whether it is read.
	HasRead *bool `json:"HasRead,omitnil,omitempty" name:"HasRead"`

	// Evaluation.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Whether to rate.
	CanRating *bool `json:"CanRating,omitnil,omitempty" name:"CanRating"`

	// Whether to display the feedback button.
	CanFeedback *bool `json:"CanFeedback,omitnil,omitempty" name:"CanFeedback"`

	// Record type.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Reference source.
	References []*MsgRecordReference `json:"References,omitnil,omitempty" name:"References"`

	// Reason for evaluation.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`

	// Whether it is a large model.
	IsLlmGenerated *bool `json:"IsLlmGenerated,omitnil,omitempty" name:"IsLlmGenerated"`

	// Image URL, which can be public read.
	ImageUrls []*string `json:"ImageUrls,omitnil,omitempty" name:"ImageUrls"`

	// Statistical information of the current token.
	TokenStat *TokenStat `json:"TokenStat,omitnil,omitempty" name:"TokenStat"`

	// Response method.
	// 1: Large model directly replies.
	// 2: Conservative reply, reply to unknown questions.
	// 3: Reply to rejected question.
	// 4: Sensitive response.
	// 5: Directly reply to Q&A pairs. Priority will be given to answering the adopted Q&A pairs.
	// 6: Reply with welcome words.
	// 7: Reply for concurrency limit exceeded.
	// 8: Global intervention knowledge.
	// 9: Reply during the task flow process. When task_flow.type = 0 in history, it is a response from the large model.
	// 10: Reply with task flow answer.
	// 11: Reply from the search engine.
	// 12: Reply after knowledge polishing.
	// 13: Reply with image understanding.
	// 14: Reply based on real-time document.
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`

	// Option tab, used for multi-round dialogue.
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// Task information.
	TaskFlow *TaskFlowInfo `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`

	// File information passed in by the user.
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// Location information of reference source .
	QuoteInfos []*QuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// Information on the thinking process of the agent.
	AgentThought *AgentThought `json:"AgentThought,omitnil,omitempty" name:"AgentThought"`

	// Expanded information.
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// Workflow information.
	WorkFlow *WorkflowInfo `json:"WorkFlow,omitnil,omitempty" name:"WorkFlow"`
}

type MsgRecordReference struct {
	// ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// URL.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Type.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Source document ID.
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type Option struct {
	// Text.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Number of characters in a file.
	CharSize *string `json:"CharSize,omitnil,omitempty" name:"CharSize"`

	// File type.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type PluginToolReqParam struct {
	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter description.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Parameter type, 0: string; 1: int; 2: float; 3: bool; 4: object; 5: array_string; 6: array_int; 7: array_float; 8: array_bool; 9: array_object, 99: null, 100: upspecified.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Whether the parameter is required.
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// Parameter default value.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Sub-parameter. "ParamType" is useful when it is "OBJECT " or "ARRAY<>" type.
	SubParams []*PluginToolReqParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// Whether the plugin parameter configuration is hidden and invisible. true - Hidden and invisible; false - Visible.
	GlobalHidden *bool `json:"GlobalHidden,omitnil,omitempty" name:"GlobalHidden"`

	// OneOf type parameter.
	OneOf []*PluginToolReqParam `json:"OneOf,omitnil,omitempty" name:"OneOf"`

	// AnyOf type parameter.
	AnyOf []*PluginToolReqParam `json:"AnyOf,omitnil,omitempty" name:"AnyOf"`
}

type Procedure struct {
	// English name of execution process.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Chinese name for display.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Status: processing, success, failed.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Number of consumed tokens.
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Debugging information.
	Debugging *ProcedureDebugging `json:"Debugging,omitnil,omitempty" name:"Debugging"`

	// Billing resource status, 1: available; 2: unavailable.
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type ProcedureDebugging struct {
	// Retrieve query.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// System prompt.
	System *string `json:"System,omitnil,omitempty" name:"System"`

	// Multiple rounds of historical information.
	Histories []*HistorySummary `json:"Histories,omitnil,omitempty" name:"Histories"`

	// Retrieve knowledge.
	Knowledge []*KnowledgeSummary `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`

	// Task process.
	TaskFlow *TaskFlowSummary `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`

	// Workflow debugging information.
	WorkFlow *WorkFlowSummary `json:"WorkFlow,omitnil,omitempty" name:"WorkFlow"`

	// Agent debugging information.
	Agent *AgentDebugInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// Custom parameter.
	CustomVariables []*string `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`
}

type QACate struct {
	// Q&A category business ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// Category name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Quantity of Q&As under the category.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Whether it is possible to add.
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// Whether it can be edited.
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// Whether it can be deleted.
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// Subcategory.
	Children []*QACate `json:"Children,omitnil,omitempty" name:"Children"`
}

type QAList struct {
	// Q&A ID.
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// Whether to accept.
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// Question.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Answer.
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`
}

type QAQuery struct {
	// Page number.
	// 
	// 
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// Verification status.
	AcceptStatus []*uint64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// Release status.
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Q&A ID.
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// Source.
	// 
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Query an answer.
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// Query type: filename, attribute label.
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`
}

type QuoteInfo struct {
	// Reference source location.
	Position *uint64 `json:"Position,omitnil,omitempty" name:"Position"`

	// Reference source index sequence.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type RateMsgRecordRequestParams struct {
	// Application appkey.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Message ID.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1: like; 2: dislike.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Reason.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type RateMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// Application appkey.
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// Message ID.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1: like; 2: dislike.
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Reason.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *RateMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotAppKey")
	delete(f, "RecordId")
	delete(f, "Score")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RateMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RateMsgRecordResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RateMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *RateMsgRecordResponseParams `json:"Response"`
}

func (r *RateMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedQuestion struct {
	// Reject question ID.
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`

	// The question that has been rejected.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Status.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Whether editing is allowed.
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// Whether deletion is allowed.
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`
}

type ReleaseConfigs struct {
	// Configuration item description.
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Status.
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// Content after modification.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Content before modification.
	LastValue *string `json:"LastValue,omitnil,omitempty" name:"LastValue"`

	// Modified content (display "content" with priority. If "content" is empty, take "value").
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Reason for failure.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type RetryDocAuditRequestParams struct {
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocAuditRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocAuditResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocAuditResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocAuditResponseParams `json:"Response"`
}

func (r *RetryDocAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseRequestParams struct {
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocParseRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocParseResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocParseResponseParams `json:"Response"`
}

func (r *RetryDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseRequestParams struct {
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release business id.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type RetryReleaseRequest struct {
	*tchttp.BaseRequest
	
	// Robot ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Release business id.
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *RetryReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RetryReleaseResponseParams `json:"Response"`
}

func (r *RetryReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunNodeInfo struct {
	// Node type, 0: unspecified; 1: start node; 2: api node; 3: inquiry node; 4: answer node.
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node name.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Requested API.
	InvokeApi *InvokeAPI `json:"InvokeApi,omitnil,omitempty" name:"InvokeApi"`

	// Values of all slots of the current node, key: SlotID. Return an Null even if there is no value.
	SlotValues []*ValueInfo `json:"SlotValues,omitnil,omitempty" name:"SlotValues"`
}

// Predefined struct for user
type SaveDocRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// File name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File type (md|txt|docx|pdf|xlsx).
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The cos path of the platform, consistent with the UploadPath parameter queried by the DescribeStorageCredential api.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag, short for entity tag, is an information tag that identifies the content of an object when it is created and can be used to check whether the content of the object has changed.
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// Verify the consistency of the uploaded file on the cloud and the local file by validating the crc64 encoding in the cos_hash x-cos-hash-crc64ecma header.<br> After the COS is successfully uploaded, obtain it from the response header.
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// File size.
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Applicable scope of labels: 1: all; 2: by conditional range.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Source (0: source file import; 1: web page import).
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Web page (or custom link) address.
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// Label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// External reference link type: 0: system link; 1: custom link.
	// When the value is 1, the weburl field cannot be empty; otherwise, it will not take effect.
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp. 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Whether to reference a link.
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// Document operation type: 1: batch import (import Q&A pairs in batches); 2: document import (normally import a single document). The default value is 1.<br>Please note that when opt = 1, please download the Excel template from the Tencent Cloud Agent Development Platform/TCADP page.
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type SaveDocRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// File name.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// File type (md|txt|docx|pdf|xlsx).
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// The cos path of the platform, consistent with the UploadPath parameter queried by the DescribeStorageCredential api.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag, short for entity tag, is an information tag that identifies the content of an object when it is created and can be used to check whether the content of the object has changed.
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// Verify the consistency of the uploaded file on the cloud and the local file by validating the crc64 encoding in the cos_hash x-cos-hash-crc64ecma header.<br> After the COS is successfully uploaded, obtain it from the response header.
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// File size.
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Applicable scope of labels: 1: all; 2: by conditional range.
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// Source (0: source file import; 1: web page import).
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// Web page (or custom link) address.
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// Label reference.
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// External reference link type: 0: system link; 1: custom link.
	// When the value is 1, the weburl field cannot be empty; otherwise, it will not take effect.
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// Effective start time, unix timestamp.
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// Effective end time, unix timestamp. 0 indicates permanent validity.
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// Whether to reference a link.
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// Document operation type: 1: batch import (import Q&A pairs in batches); 2: document import (normally import a single document). The default value is 1.<br>Please note that when opt = 1, please download the Excel template from the Tencent Cloud Agent Development Platform/TCADP page.
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`

	// Category ID.
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *SaveDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "CosUrl")
	delete(f, "ETag")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "AttrRange")
	delete(f, "Source")
	delete(f, "WebUrl")
	delete(f, "AttrLabels")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "IsRefer")
	delete(f, "Opt")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveDocResponseParams struct {
	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// Import error message.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Error link.
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// Error link text.
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveDocResponse struct {
	*tchttp.BaseResponse
	Response *SaveDocResponseParams `json:"Response"`
}

func (r *SaveDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRange struct {
	// Search criteria and/or.
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// Custom variable and label relational data.
	ApiVarAttrInfos []*ApiVarAttrInfo `json:"ApiVarAttrInfos,omitnil,omitempty" name:"ApiVarAttrInfos"`
}

type SearchStrategy struct {
	// Retrieval strategy type, 0: hybrid retrieval; 1: semantic retrieval.
	StrategyType *uint64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// Excel retrieval enhancement switch.
	TableEnhancement *bool `json:"TableEnhancement,omitnil,omitempty" name:"TableEnhancement"`
}

type SimilarQuestion struct {
	// Similar question ID.
	SimBizId *string `json:"SimBizId,omitnil,omitempty" name:"SimBizId"`

	// Similar question content.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Similar question review status, 1: audit failure.
	AuditStatus *uint64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`
}

type SimilarQuestionModify struct {
	// List of similar questions (content) to be added.
	AddQuestions []*string `json:"AddQuestions,omitnil,omitempty" name:"AddQuestions"`

	// List of similar questions to be updated.
	UpdateQuestions []*SimilarQuestion `json:"UpdateQuestions,omitnil,omitempty" name:"UpdateQuestions"`

	// List of similar questions to be deleted.
	DeleteQuestions []*SimilarQuestion `json:"DeleteQuestions,omitnil,omitempty" name:"DeleteQuestions"`
}

type Stat struct {
	// X-axis: time zone; return three interval ranges of "minute/hour/day" according to the granularity of the query condition.
	X *string `json:"X,omitnil,omitempty" name:"X"`

	// Y-axis: statistical values in this time period, such as token consumption, call count, or usage information.
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type StatisticInfo struct {
	// Model name.
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// Time consumption of the first token.
	FirstTokenCost *uint64 `json:"FirstTokenCost,omitnil,omitempty" name:"FirstTokenCost"`

	// Total time consumption.
	TotalCost *uint64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Number of input tokens.
	InputTokens *uint64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// Number of output tokens.
	OutputTokens *uint64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// Total number of tokens.
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

// Predefined struct for user
type StopDocParseRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type StopDocParseRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Document ID.
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *StopDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDocParseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopDocParseResponse struct {
	*tchttp.BaseResponse
	Response *StopDocParseResponseParams `json:"Response"`
}

func (r *StopDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrValue struct {
	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SummaryConfig struct {
	// Model configuration.
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// Knowledge summary output configuration.
	Output *SummaryOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// Welcome words, within 200 characters.
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
}

type SummaryOutput struct {
	// Output method: 1. streaming; 2. non-streaming.
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// Output requirement 1: text summary. 2: customized requirement.
	Requirement *uint64 `json:"Requirement,omitnil,omitempty" name:"Requirement"`

	// Custom requirement command.
	RequireCommand *string `json:"RequireCommand,omitnil,omitempty" name:"RequireCommand"`
}

type TaskFlowInfo struct {
	// Task flow ID.
	TaskFlowId *string `json:"TaskFlowId,omitnil,omitempty" name:"TaskFlowId"`

	// Task flow name.
	TaskFlowName *string `json:"TaskFlowName,omitnil,omitempty" name:"TaskFlowName"`

	// Rewrite results of query.
	QueryRewrite *string `json:"QueryRewrite,omitnil,omitempty" name:"QueryRewrite"`

	// Hit intent.
	HitIntent *string `json:"HitIntent,omitnil,omitempty" name:"HitIntent"`

	// Task flow response type.
	// 0: Task flow response.
	// 1: Silent task flow.
	// 2: Pull back script of workflow.
	// 3: Custom response of task flow.
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type TaskFlowSummary struct {
	// Task flow name.
	IntentName *string `json:"IntentName,omitnil,omitempty" name:"IntentName"`

	// Entity list.
	UpdatedSlotValues []*ValueInfo `json:"UpdatedSlotValues,omitnil,omitempty" name:"UpdatedSlotValues"`

	// Node list.
	RunNodes []*RunNodeInfo `json:"RunNodes,omitnil,omitempty" name:"RunNodes"`

	// Intent determination.
	Purposes []*string `json:"Purposes,omitnil,omitempty" name:"Purposes"`
}

type TaskParams struct {
	// Download address. Download through the COS bucket temporary key.
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

type TokenStat struct {
	// Session ID.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Request ID.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// It corresponds to a session, session id, used for storing messages for answering. It can be generated in advance, used when saving messages.
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// Number of consumed tokens.
	UsedCount *uint64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// Number of free tokens.
	FreeCount *uint64 `json:"FreeCount,omitnil,omitempty" name:"FreeCount"`

	// Total number of tokens for orders.
	OrderCount *uint64 `json:"OrderCount,omitnil,omitempty" name:"OrderCount"`

	// Current execution status summary. Constant: processing, success., failed.
	StatusSummary *string `json:"StatusSummary,omitnil,omitempty" name:"StatusSummary"`

	// Chinese display after summarizing the current execution status.
	StatusSummaryTitle *string `json:"StatusSummaryTitle,omitnil,omitempty" name:"StatusSummaryTitle"`

	// Current request execution time, in milliseconds.
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// Number of tokens consumed by the current request.
	TokenCount *uint64 `json:"TokenCount,omitnil,omitempty" name:"TokenCount"`

	// Execution information.
	Procedures []*Procedure `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// Execution process information TraceId.
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

type UnsatisfiedReply struct {
	// Unsatisfied response ID.
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// Message record ID.
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// User question.
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Application response.
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// Error type.
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type UploadAttributeLabelRequestParams struct {
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Filename.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Cos path.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// Verify the consistency of files uploaded to the cloud and local files by validating the crc64 encoding in the x-cos-hash-crc64ecma header.
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// File size.
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type UploadAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Filename.
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Cos path.
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// Verify the consistency of files uploaded to the cloud and local files by validating the crc64 encoding in the x-cos-hash-crc64ecma header.
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// File size.
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *UploadAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "CosUrl")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadAttributeLabelResponseParams struct {
	// Import error.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Error link.
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// Error link text.
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *UploadAttributeLabelResponseParams `json:"Response"`
}

func (r *UploadAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueInfo struct {
	// Value ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value type, 0: unknown or empty; 1: string; 2: integer; 3: float; 4: boolean; 5: array (string array); 6: object_array (structure array); 7: object (structure).
	ValueType *int64 `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// String.
	ValueStr *string `json:"ValueStr,omitnil,omitempty" name:"ValueStr"`

	// Int (return as a string to avoid precision loss).
	ValueInt *string `json:"ValueInt,omitnil,omitempty" name:"ValueInt"`

	// Float.
	ValueFloat *float64 `json:"ValueFloat,omitnil,omitempty" name:"ValueFloat"`

	// Bool.
	ValueBool *bool `json:"ValueBool,omitnil,omitempty" name:"ValueBool"`

	// Array.
	ValueStrArray []*string `json:"ValueStrArray,omitnil,omitempty" name:"ValueStrArray"`
}

// Predefined struct for user
type VerifyQARequestParams struct {
	// Q&A list.
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type VerifyQARequest struct {
	*tchttp.BaseRequest
	
	// Q&A list.
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// Application ID.
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// Login to user's root account (required in integrator mode).
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// Login to user's sub-account (required in integrator mode).
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *VerifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyQAResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyQAResponse struct {
	*tchttp.BaseResponse
	Response *VerifyQAResponseParams `json:"Response"`
}

func (r *VerifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoiceConfig struct {

	VoiceType *uint64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`


	TimbreKey *string `json:"TimbreKey,omitnil,omitempty" name:"TimbreKey"`


	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`
}

type WorkFlowSummary struct {
	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// Workflow running ID.
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// Node information.
	RunNodes []*WorkflowRunNodeInfo `json:"RunNodes,omitnil,omitempty" name:"RunNodes"`

	// Tab.
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// Output results of multiple bubbles.
	Outputs []*string `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// Workflow release time, Unix timestamp.
	WorkflowReleaseTime *string `json:"WorkflowReleaseTime,omitnil,omitempty" name:"WorkflowReleaseTime"`
}

type WorkflowInfo struct {

	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`


	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`


	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`


	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`


	Outputs []*string `json:"Outputs,omitnil,omitempty" name:"Outputs"`


	WorkflowReleaseTime *string `json:"WorkflowReleaseTime,omitnil,omitempty" name:"WorkflowReleaseTime"`
}

type WorkflowRunNodeInfo struct {
	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// Node type.
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Node name.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Status.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Input..
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// Output.
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Task output.
	TaskOutput *string `json:"TaskOutput,omitnil,omitempty" name:"TaskOutput"`

	// Error message.
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// Time consumption.
	CostMilliSeconds *uint64 `json:"CostMilliSeconds,omitnil,omitempty" name:"CostMilliSeconds"`

	// Large model output information.
	StatisticInfos []*StatisticInfo `json:"StatisticInfos,omitnil,omitempty" name:"StatisticInfos"`
}