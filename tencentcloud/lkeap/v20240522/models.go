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

package v20240522

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type CreateReconstructDocumentFlowConfig struct {
	// The form in which tables in a Markdown file are returned.
	// - 0: returned as MD.
	// - 1: returned as HTML.
	// The default is 0.
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// The format of the returned results of intelligent document parsing.
	// - 0: only returns full-text MD.
	// - 1: only returns the OCR original Json for each page.
	// - 2: only returns the MD of each page.
	// - 3: returns the full-text MD and the original OCR Json of each page.
	// - 4: returns full-text MD and MD of each page.
	// The default value is 3 (returns the full-text MD and the original OCR Json of each page).
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`
}

// Predefined struct for user
type CreateReconstructDocumentFlowRequestParams struct {
	// File type.
	// **Supported file types**: PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX, MD, TXT, PNG, JPG, JPEG, CSV, HTML, EPUB, BMP, GIF, WEBP, HEIC, EPS, ICNS, IM, PCX, PPM, TIFF, XBM, HEIF, JP2.
	// **Supported file sizes**: 
	// - Max 100 MB for PDF.
	// - Max 200 MB for DOC, DOCX, PPT, and PPTX .
	// - Max 10 MB for MD, and TXT.
	// - Max 20 MB for others.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// File URL. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. Refer to: [Tencent Cloud COS Documentation](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// The base64 value of the file. Supported file types: PNG, JPG, JPEG, PDF, BMP, TIFF. File size limit: the downloaded file does not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// The starting page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the starting page number for recognition, including the current value.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// The end page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the end page number for recognition, including the current value.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Creates task configuration information for document parsing.
	Config *CreateReconstructDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateReconstructDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// File type.
	// **Supported file types**: PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX, MD, TXT, PNG, JPG, JPEG, CSV, HTML, EPUB, BMP, GIF, WEBP, HEIC, EPS, ICNS, IM, PCX, PPM, TIFF, XBM, HEIF, JP2.
	// **Supported file sizes**: 
	// - Max 100 MB for PDF.
	// - Max 200 MB for DOC, DOCX, PPT, and PPTX .
	// - Max 10 MB for MD, and TXT.
	// - Max 20 MB for others.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// File URL. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. Refer to: [Tencent Cloud COS Documentation](https://cloud.tencent.com/document/product/436/7749)
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// The base64 value of the file. Supported file types: PNG, JPG, JPEG, PDF, BMP, TIFF. File size limit: the downloaded file does not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// The starting page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the starting page number for recognition, including the current value.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// The end page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the end page number for recognition, including the current value.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Creates task configuration information for document parsing.
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
	delete(f, "FileUrl")
	delete(f, "FileBase64")
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
	// Unique task ID. The processing result corresponding to TaskId can be queried through the API [GetReconstructDocumentResult] within 30 days.
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

type CreateSplitDocumentFlowConfig struct {
	// The form in which tables in a Markdown file are returned.
	// - 0: returned as MD.
	// - 1: returned as HTML.
	//
	// Deprecated: TableResultType is deprecated.
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// The format of the returned results of intelligent document parsing.
	// - 0: only returns full-text MD.
	// - 1: only returns the OCR original Json for each page.
	// - 2: only returns the MD of each page.
	// - 3: returns the full-text MD and the original OCR Json of each page.
	// - 4: returns full-text MD and MD of each page.
	// The default value is 3 (returns the full-text MD and the original OCR Json of each page).
	//
	// Deprecated: ResultType is deprecated.
	ResultType *string `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// Whether to enable mllm.
	EnableMllm *bool `json:"EnableMllm,omitnil,omitempty" name:"EnableMllm"`

	// Max segment length.
	MaxChunkSize *int64 `json:"MaxChunkSize,omitnil,omitempty" name:"MaxChunkSize"`
}

// Predefined struct for user
type CreateSplitDocumentFlowRequestParams struct {
	// File type.
	// **Supported file types**: PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX, MD, TXT, PNG, JPG, JPEG, CSV, HTML, EPUB.
	// **Supported file sizes**: 
	// - Max 500 MB for PDF.
	// - Max 200 MB for DOC, DOCX, PPT, and PPTX .
	// - Max 10 MB for MD, and TXT.
	// - Max 20 MB for others.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// File URL. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. 
	// Refer to: [Tencent Cloud COS Documentation](https://cloud.tencent.com/document/product/436/7749).
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// Filename. optional.
	// **The file type suffix shall be included**. This field is required to be specified when the file name cannot be obtained from the passed-in "FileUrl".
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// The base64 value of the file. File size limit: the downloaded file shall not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	//
	// Deprecated: FileBase64 is deprecated.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// The starting page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the starting page number for recognition, including the current value.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// The end page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the end page number for recognition, including the current value.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Configuration message for document splitting task.
	Config *CreateSplitDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type CreateSplitDocumentFlowRequest struct {
	*tchttp.BaseRequest
	
	// File type.
	// **Supported file types**: PDF, DOC, DOCX, XLS, XLSX, PPT, PPTX, MD, TXT, PNG, JPG, JPEG, CSV, HTML, EPUB.
	// **Supported file sizes**: 
	// - Max 500 MB for PDF.
	// - Max 200 MB for DOC, DOCX, PPT, and PPTX .
	// - Max 10 MB for MD, and TXT.
	// - Max 20 MB for others.
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// File URL. It is recommended to store the file in Tencent Cloud as the URL where the file is stored in Tencent Cloud can ensure higher download speed and stability. External URL may affect the speed and stability. 
	// Refer to: [Tencent Cloud COS Documentation](https://cloud.tencent.com/document/product/436/7749).
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// Filename. optional.
	// **The file type suffix shall be included**. This field is required to be specified when the file name cannot be obtained from the passed-in "FileUrl".
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// The base64 value of the file. File size limit: the downloaded file shall not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// The starting page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the starting page number for recognition, including the current value.
	FileStartPageNumber *int64 `json:"FileStartPageNumber,omitnil,omitempty" name:"FileStartPageNumber"`

	// The end page number of the file. When type of the uploaded file is pdf, doc, ppt, or pptx, it specifies the end page number for recognition, including the current value.
	FileEndPageNumber *int64 `json:"FileEndPageNumber,omitnil,omitempty" name:"FileEndPageNumber"`

	// Configuration message for document splitting task.
	Config *CreateSplitDocumentFlowConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CreateSplitDocumentFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSplitDocumentFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "FileName")
	delete(f, "FileBase64")
	delete(f, "FileStartPageNumber")
	delete(f, "FileEndPageNumber")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSplitDocumentFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSplitDocumentFlowResponseParams struct {
	// The unique ID of the splitting task.
	// The splitting results corresponding to the TaskId can be queried through the [GetSplitDocumentResult] API within 30 days.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSplitDocumentFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateSplitDocumentFlowResponseParams `json:"Response"`
}

func (r *CreateSplitDocumentFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSplitDocumentFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DocumentUsage struct {
	// Page number of the document splitting task.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Total number of tokens consumed by the document splitting task.
	//
	// Deprecated: TotalToken is deprecated.
	TotalToken *int64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// Total number of tokens consumed by the document splitting task.
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

// Predefined struct for user
type GetReconstructDocumentResultRequestParams struct {
	// Parsing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetReconstructDocumentResultRequest struct {
	*tchttp.BaseRequest
	
	// Parsing task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetReconstructDocumentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetReconstructDocumentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetReconstructDocumentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetReconstructDocumentResultResponseParams struct {
	// Task status.
	// -Success: execution completed.
	// -Processing: executing.
	// -Pause: paused.
	// -Failed: execution failed.
	// -WaitExecute: pending execution.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Temporary download URL for the parsing result. The file is a zip compressed package, and the URL is valid for 30 minutes.
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// Page number where document parsing fails.
	FailedPages []*ReconstructDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetReconstructDocumentResultResponse struct {
	*tchttp.BaseResponse
	Response *GetReconstructDocumentResultResponseParams `json:"Response"`
}

func (r *GetReconstructDocumentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetReconstructDocumentResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSplitDocumentResultRequestParams struct {
	// Splitting task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetSplitDocumentResultRequest struct {
	*tchttp.BaseRequest
	
	// Splitting task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetSplitDocumentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSplitDocumentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSplitDocumentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSplitDocumentResultResponseParams struct {
	// Task status:
	// -Success: execution completed.
	// -Processing: executing.
	// -Pause: paused.
	// -Failed: execution failed.
	// -WaitExecute: pending execution.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Temporary download URL for the splitting result. The file is a zip compressed package, and the URL is valid for 30 minutes. The compressed package contains the following folders: \*.md, \*.jsonl, \*mllm.json, images.
	// >**jsonl** structure description:.
	// - page_content: Used to generate an embedding library for retrieval purposes. The images in this field will be replaced with placeholders.
	// - org_data: The minimum semantic integrity block corresponding to page_content, used for Q&A model processing.
	// - big_data: The maximum semantic integrity block corresponding to page_content, also used for Q&A model processing.
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// Page number where document splitting fails.
	//
	// Deprecated: FailedPages is deprecated.
	FailedPages []*SplitDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`

	// Amount of the document split task.
	Usage *DocumentUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSplitDocumentResultResponse struct {
	*tchttp.BaseResponse
	Response *GetSplitDocumentResultResponseParams `json:"Response"`
}

func (r *GetSplitDocumentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSplitDocumentResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Message struct {
	// Role.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Chain of thought content. The ReasoningConent parameter only supports output parameters, and is only returned by the deepseek-r1 model.
	ReasoningContent *string `json:"ReasoningContent,omitnil,omitempty" name:"ReasoningContent"`
}

// Predefined struct for user
type QueryRewriteRequestParams struct {
	// The multi-round historical conversation that needs to be rewritten. Each round of historical conversation should include paired inputs of user (question) and assistant (answer). Due to the character limit of the model, a maximum of 4 rounds of conversations can be provided. The last round of conversation will be rewritten.
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// Model name.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type QueryRewriteRequest struct {
	*tchttp.BaseRequest
	
	// The multi-round historical conversation that needs to be rewritten. Each round of historical conversation should include paired inputs of user (question) and assistant (answer). Due to the character limit of the model, a maximum of 4 rounds of conversations can be provided. The last round of conversation will be rewritten.
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// Model name.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

func (r *QueryRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Messages")
	delete(f, "Model")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryRewriteResponseParams struct {
	// Rewritten result.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// Consumption. The numbers of input tokens, output tokens, and total tokens will be returned.
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryRewriteResponse struct {
	*tchttp.BaseResponse
	Response *QueryRewriteResponseParams `json:"Response"`
}

func (r *QueryRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReconstructDocumentFailedPage struct {
	// Page number that failed to parse.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type ReconstructDocumentSSEConfig struct {
	// The form in which tables in a Markdown file are returned.
	// - 0: returned as MD.
	// - 1: returned as HTML.
	// The default is 0.
	TableResultType *string `json:"TableResultType,omitnil,omitempty" name:"TableResultType"`

	// The form in which images in a Markdown file are returned.
	// - 0: returned as URL.
	// - 1: only return the text content extracted from the image in markdown.
	// The default value is 0.
	MarkdownImageResponseType *string `json:"MarkdownImageResponseType,omitnil,omitempty" name:"MarkdownImageResponseType"`

	// Whether the Markdown file contains page number information.
	ReturnPageFormat *bool `json:"ReturnPageFormat,omitnil,omitempty" name:"ReturnPageFormat"`

	// The custom output page number style. {{p}} is a placeholder for the page number. Enable ReturnPageFormat to take effect. If empty, the default style is: <page_num>page {{p}}</page_num>.
	PageFormat *string `json:"PageFormat,omitnil,omitempty" name:"PageFormat"`
}

// Predefined struct for user
type ReconstructDocumentSSERequestParams struct {
	// The base64 value of the file. File size limit: the downloaded file shall not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// Document parsing configuration information.	
	Config *ReconstructDocumentSSEConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type ReconstructDocumentSSERequest struct {
	*tchttp.BaseRequest
	
	// The base64 value of the file. File size limit: the downloaded file shall not exceed 8MB after base64 encoding. File download time does not exceed 3 seconds. Supported image pixels: the length of a single side is between 20-10000px. Either FileUrl or FileBase64 of the file must be provided. If both are provided, only the FileUrl is used.
	FileBase64 *string `json:"FileBase64,omitnil,omitempty" name:"FileBase64"`

	// Document parsing configuration information.	
	Config *ReconstructDocumentSSEConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *ReconstructDocumentSSERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentSSERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileBase64")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReconstructDocumentSSERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReconstructDocumentSSEResponseParams struct {
	// Task ID. The unique identifier of the current request.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Response type. 1: return progress information; 2: return parsing result.
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// Progress. Value range: 0 to 100.
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Progress information.
	ProgressMessage *string `json:"ProgressMessage,omitnil,omitempty" name:"ProgressMessage"`

	// Temporary download URL for the document parsing result. The file is a zip compressed package, and the URL is valid for 30 minutes. The compressed package contains the following folders: \*.md, \*.jsonl, \*mllm.json, images.
	DocumentRecognizeResultUrl *string `json:"DocumentRecognizeResultUrl,omitnil,omitempty" name:"DocumentRecognizeResultUrl"`

	// Page number where document parsing fails.
	FailedPages []*ReconstructDocumentFailedPage `json:"FailedPages,omitnil,omitempty" name:"FailedPages"`


	FailPageNum *int64 `json:"FailPageNum,omitnil,omitempty" name:"FailPageNum"`


	SuccessPageNum *int64 `json:"SuccessPageNum,omitnil,omitempty" name:"SuccessPageNum"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem. As a streaming response API, when the request is successfully completed, the RequestId will be placed in the Header "X-TC-RequestId" of the HTTP response.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReconstructDocumentSSEResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ReconstructDocumentSSEResponseParams `json:"Response"`
}

func (r *ReconstructDocumentSSEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReconstructDocumentSSEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerankRequestParams struct {
	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Document list, up to 20 documents.
	Docs []*string `json:"Docs,omitnil,omitempty" name:"Docs"`

	// Model name. Default: lke-reranker-base.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type RunRerankRequest struct {
	*tchttp.BaseRequest
	
	// Query content.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Document list, up to 20 documents.
	Docs []*string `json:"Docs,omitnil,omitempty" name:"Docs"`

	// Model name. Default: lke-reranker-base.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

func (r *RunRerankRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerankRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Docs")
	delete(f, "Model")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunRerankRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerankResponseParams struct {
	// Relevance. A higher numeric value indicates greater relevance.
	ScoreList []*float64 `json:"ScoreList,omitnil,omitempty" name:"ScoreList"`

	// Consumption. Only returns TotalToken.
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunRerankResponse struct {
	*tchttp.BaseResponse
	Response *RunRerankResponseParams `json:"Response"`
}

func (r *RunRerankResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerankResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitDocumentFailedPage struct {
	// Page number that failed to parse.
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type Usage struct {
	// Number of document pages.
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// Number of input tokens.
	InputTokens *int64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// Number of output tokens.
	OutputTokens *int64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// Total number of tokens.
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}