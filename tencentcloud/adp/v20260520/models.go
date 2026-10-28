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

package v20260520

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type DBRetrievalConfig struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DocRetrievalConfig struct {
	// <p>Confidence threshold</p>
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Return the top N entries</p>
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type ExpirationAwareness struct {
	// <p>Whether to enable timeliness retrieval. Once enabled, retrieval results will be sorted based on the validity time of the knowledge</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type FinalRerankConfig struct {
	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type GraphRAG struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type KBRetrievalConfig struct {
	// <p>Database search configuration</p>
	DbRetrievalConfig *DBRetrievalConfig `json:"DbRetrievalConfig,omitnil,omitempty" name:"DbRetrievalConfig"`

	// <p>Document retrieval configuration.</p>
	DocRetrievalConfig *DocRetrievalConfig `json:"DocRetrievalConfig,omitnil,omitempty" name:"DocRetrievalConfig"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Retrieval optional configuration</p>
	OptionConfig *RetrievalOption `json:"OptionConfig,omitnil,omitempty" name:"OptionConfig"`

	// <p>QA search configuration</p>
	QaRetrievalConfig *QARetrievalConfig `json:"QaRetrievalConfig,omitnil,omitempty" name:"QaRetrievalConfig"`

	// <p>ReRank configuration</p>
	RerankConfig *RerankConfig `json:"RerankConfig,omitnil,omitempty" name:"RerankConfig"`

	// <p>Search filter configuration</p>
	SearchFilterConfig *SearchFilterConfig `json:"SearchFilterConfig,omitnil,omitempty" name:"SearchFilterConfig"`

	// <p>Retrieval strategy: 1=hybrid, 2=semantic, 3=keyword, 4=none<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SEARCH_STRATEGY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SEARCH_STRATEGY_TYPE_MIXING</td><td>1</td><td>Hybrid retrieval</td></tr><tr><td>SEARCH_STRATEGY_TYPE_SEMANTIC</td><td>2</td><td>Semantic retrieval</td></tr><tr><td>SEARCH_STRATEGY_TYPE_KEYWORD</td><td>3</td><td>Keyword retrieval</td></tr><tr><td>SEARCH_STRATEGY_TYPE_NONE</td><td>4</td><td>No semantic/vector retrieval</td></tr></tbody></table></p>
	StrategyType *int64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// <p>Text2SQL model</p>
	TextToSqlModel *ModelDetailInfo `json:"TextToSqlModel,omitnil,omitempty" name:"TextToSqlModel"`
}

type KVPair struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KnowledgeResult struct {
	// <p>Confidence degree</p>
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>Hit knowledge type: 1=QA, 2=document fragment, 3=database, 4=map<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_SOURCE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_SOURCE_TYPE_QA</td><td>1</td><td>QA</td></tr><tr><td>KNOWLEDGE_SOURCE_TYPE_DOC</td><td>2</td><td>Document fragment</td></tr><tr><td>KNOWLEDGE_SOURCE_TYPE_DB</td><td>3</td><td>Database</td></tr><tr><td>KNOWLEDGE_SOURCE_TYPE_GRAPH</td><td>4</td><td>Map</td></tr></tbody></table></p>
	KnowledgeType *int64 `json:"KnowledgeType,omitnil,omitempty" name:"KnowledgeType"`

	// <p>Recall type list</p>
	RecallTypeList []*int64 `json:"RecallTypeList,omitnil,omitempty" name:"RecallTypeList"`

	// <p>Result payload</p>
	ResultPayload *SearchResultPayload `json:"ResultPayload,omitnil,omitempty" name:"ResultPayload"`

	// <p>Result type<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SEARCH_RESULT_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SEARCH_RESULT_TYPE_RETRIEVAL</td><td>1</td><td>Standard search result</td></tr><tr><td>SEARCH_RESULT_TYPE_TEXT_TO_SQL</td><td>2</td><td>text2sql result</td></tr><tr><td>SEARCH_RESULT_TYPE_IMAGE_SEARCH_IMAGE</td><td>3</td><td></td></tr><tr><td>SEARCH_RESULT_TYPE_TEXT_SEARCH_IMAGE</td><td>4</td><td></td></tr></tbody></table></p>
	ResultType *int64 `json:"ResultType,omitnil,omitempty" name:"ResultType"`

	// <p>Additional information for similar questions.</p>
	SimilarQuestionExtra *SimilarQuestionExtra `json:"SimilarQuestionExtra,omitnil,omitempty" name:"SimilarQuestionExtra"`

	// <p>Basic information of knowledge fragment</p>
	SnippetProfile *KnowledgeSnippetProfile `json:"SnippetProfile,omitnil,omitempty" name:"SnippetProfile"`

	// <p>Knowledge source information</p>
	SourceInfo *KnowledgeSource `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type KnowledgeSnippetProfile struct {
	// <p>Content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Associated document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Document name</p>
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Knowledge ID</p>
	KnowledgeId *string `json:"KnowledgeId,omitnil,omitempty" name:"KnowledgeId"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Document title</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type KnowledgeSource struct {
	// <p>Whether it is big data</p>
	IsBigData *bool `json:"IsBigData,omitnil,omitempty" name:"IsBigData"`

	// <p>Whether it is a shared knowledge base</p>
	IsShared *bool `json:"IsShared,omitnil,omitempty" name:"IsShared"`

	// <p>Knowledge base name</p>
	KbName *string `json:"KbName,omitnil,omitempty" name:"KbName"`
}

type ModelDetailInfo struct {
	// Model alias.
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Limit on number of historical conversations
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// Unique model ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Model parameters
	ModelParams *ModelParams `json:"ModelParams,omitnil,omitempty" name:"ModelParams"`
}

type ModelParams struct {
	// <p>Whether deep thinking is enabled</p>
	DeepThinking *string `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// <p>Frequency penalty</p>
	FrequencyPenalty *float64 `json:"FrequencyPenalty,omitnil,omitempty" name:"FrequencyPenalty"`

	// <p>Maximum output length</p>
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`

	// <p>Existence penalty</p>
	PresencePenalty *float64 `json:"PresencePenalty,omitnil,omitempty" name:"PresencePenalty"`

	// <p>Deep thinking effect</p>
	ReasoningEffort *string `json:"ReasoningEffort,omitnil,omitempty" name:"ReasoningEffort"`

	// <p>Repetition penalty</p>
	RepetitionPenalty *float64 `json:"RepetitionPenalty,omitnil,omitempty" name:"RepetitionPenalty"`

	// <p>Output format (text, json_object)</p>
	ReplyFormat *string `json:"ReplyFormat,omitnil,omitempty" name:"ReplyFormat"`

	// <p>Random seed</p>
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// <p>Stop sequence</p>
	StopSequenceList []*string `json:"StopSequenceList,omitnil,omitempty" name:"StopSequenceList"`

	// <p>Temperature</p>
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// <p>top_p</p>
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// <p>top_k</p>
	TopK *int64 `json:"TopK,omitnil,omitempty" name:"TopK"`
}

type QARetrievalConfig struct {
	// <p>Confidence threshold</p>
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Return the top N entries</p>
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type RerankConfig struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type RetrievalOption struct {
	// <p>Timeliness retrieval enhancement configuration</p>
	ExpirationAwareness *ExpirationAwareness `json:"ExpirationAwareness,omitnil,omitempty" name:"ExpirationAwareness"`

	// <p>GraphRAG Configuration</p>
	GraphRag *GraphRAG `json:"GraphRag,omitnil,omitempty" name:"GraphRag"`

	// <p>Spreadsheet enhancement configuration</p>
	TableEnhancement *TableEnhancement `json:"TableEnhancement,omitnil,omitempty" name:"TableEnhancement"`
}

type SearchAdvancedConfig struct {
	// <p>Final ReRank Configuration</p>
	FinalRerankConfig *FinalRerankConfig `json:"FinalRerankConfig,omitnil,omitempty" name:"FinalRerankConfig"`

	// <p>Search configuration of each knowledge base</p>
	KbRetrievalList []*KBRetrievalConfig `json:"KbRetrievalList,omitnil,omitempty" name:"KbRetrievalList"`

	// <p>Retrieve knowledge type: 1=document and QA, 2=rejected question<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SEARCH_KNOWLEDGE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SEARCH_KNOWLEDGE_TYPE_DOC_QA</td><td>1</td><td>Document and QA</td></tr><tr><td>SEARCH_KNOWLEDGE_TYPE_REJECTED_QUESTION</td><td>2</td><td>Rejected question</td></tr></tbody></table></p>
	KnowledgeType *int64 `json:"KnowledgeType,omitnil,omitempty" name:"KnowledgeType"`

	// <p>Number of final returned results</p>
	RecallCount *uint64 `json:"RecallCount,omitnil,omitempty" name:"RecallCount"`
}

type SearchBilling struct {
	// <p>Billing tag list</p>
	BillingTagList []*KVPair `json:"BillingTagList,omitnil,omitempty" name:"BillingTagList"`

	// <p>Billing sub-business type</p>
	FinanceSubBusinessType *string `json:"FinanceSubBusinessType,omitnil,omitempty" name:"FinanceSubBusinessType"`
}

type SearchContext struct {
	// <p>Request source<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CALL_SOURCE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>CALL_SOURCE_RAG</td><td>1</td><td>Standard mode</td></tr><tr><td>CALL_SOURCE_WORKFLOW</td><td>2</td><td>Workflow</td></tr><tr><td>CALL_SOURCE_PLUGIN</td><td>3</td><td>Plug-in</td></tr><tr><td>CALL_SOURCE_OPENCLAW</td><td>4</td><td>openclaw</td></tr><tr><td>CALL_SOURCE_RECALL_TEST</td><td>5</td><td>Recall test</td></tr><tr><td>CALL_SOURCE_RECALL_TEST_DIFF</td><td>6</td><td>Recall test in comparison scenarios also needs to trigger the retrieval interface. The difference is that in this case, the frontend does not need to update the latest configuration. Because when comparing 1, 2, and 3, the final saved one may be 2.</td></tr></tbody></table></p>
	CallSource *int64 `json:"CallSource,omitnil,omitempty" name:"CallSource"`

	// <p>adp DOMAIN: 1=development DOMAIN, 2=production DOMAIN<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>adp_DOMAIN_UNSPECIFIED</td><td>0</td><td>not specified</td></tr><tr><td>adp_DOMAIN_DEV</td><td>1</td><td>development DOMAIN</td></tr><tr><td>adp_DOMAIN_PROD</td><td>2</td><td>production DOMAIN</td></tr></tbody></table></p>
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Visitor ID</p>
	VisitorId *string `json:"VisitorId,omitnil,omitempty" name:"VisitorId"`
}

type SearchFilter struct {
	// <p>Search filter type<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SEARCH_FILTER_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SEARCH_FILTER_TYPE_CUSTOMER_LABEL_VALUE</td><td>1</td><td>User-customized tag value</td></tr><tr><td>SEARCH_FILTER_TYPE_CUSTOMER_LABEL_VALUE_ID</td><td>2</td><td>User-customized tag value ID</td></tr><tr><td>SEARCH_FILTER_TYPE_DOC_ID</td><td>3</td><td>Retrieve by specified document ID</td></tr><tr><td>SEARCH_FILTER_TYPE_DOC_CATEGORY_ID</td><td>4</td><td>Retrieve by specified document classification ID</td></tr><tr><td>SEARCH_FILTER_TYPE_DB_TABLE_ID</td><td>5</td><td>Retrieve by specified database table ID</td></tr><tr><td>SEARCH_FILTER_TYPE_KB_SCHEMA_ID</td><td>6</td><td>Specified knowledge base schema ID</td></tr></tbody></table></p>
	FilterType *int64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// <p>Filter value list. According to SearchFilterType, take value 1: pass a custom tag value; 2: pass a custom tag value ID; 3: pass a document ID; 4: pass a category ID</p>
	FilterValueList []*string `json:"FilterValueList,omitnil,omitempty" name:"FilterValueList"`

	// <p>User-customized tag ID</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Logical operator: AND or OR<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>LOGIC_OPR_NOOP</td><td>0</td><td></td></tr><tr><td>LOGIC_OPR_AND</td><td>1</td><td></td></tr><tr><td>LOGIC_OPR_OR</td><td>2</td><td></td></tr></tbody></table></p>
	LogicOp *int64 `json:"LogicOp,omitnil,omitempty" name:"LogicOp"`

	// <p>Nested search filter</p>
	SearchFilterList []*SearchFilter `json:"SearchFilterList,omitnil,omitempty" name:"SearchFilterList"`
}

type SearchFilterConfig struct {
	// <p>Whether to retrieve only selected tags. true: retrieve only knowledge with selected tags; false: simultaneous retrieval of knowledge with selected tags and knowledge without any tags</p>
	OnlyRetrievalSelectedLabel *bool `json:"OnlyRetrievalSelectedLabel,omitnil,omitempty" name:"OnlyRetrievalSelectedLabel"`

	// <p>Search filter</p>
	SearchFilter *SearchFilter `json:"SearchFilter,omitnil,omitempty" name:"SearchFilter"`
}

type SearchInput struct {
	// <p>List of image URLs</p>
	ImageUrlList []*string `json:"ImageUrlList,omitnil,omitempty" name:"ImageUrlList"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Deconstructed sub-problem list</p>
	SubQuestionList []*string `json:"SubQuestionList,omitnil,omitempty" name:"SubQuestionList"`
}

// Predefined struct for user
type SearchKnowledgeRequestParams struct {
	// <p>Advanced search configuration</p>
	AdvancedConfig *SearchAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// <p>Retrieval input</p>
	Input *SearchInput `json:"Input,omitnil,omitempty" name:"Input"`

	// <p>Retrieval context</p>
	Context *SearchContext `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>Billing information</p>
	SearchBilling *SearchBilling `json:"SearchBilling,omitnil,omitempty" name:"SearchBilling"`
}

type SearchKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// <p>Advanced search configuration</p>
	AdvancedConfig *SearchAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// <p>Retrieval input</p>
	Input *SearchInput `json:"Input,omitnil,omitempty" name:"Input"`

	// <p>Retrieval context</p>
	Context *SearchContext `json:"Context,omitnil,omitempty" name:"Context"`

	// <p>Billing information</p>
	SearchBilling *SearchBilling `json:"SearchBilling,omitnil,omitempty" name:"SearchBilling"`
}

func (r *SearchKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AdvancedConfig")
	delete(f, "Input")
	delete(f, "Context")
	delete(f, "SearchBilling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKnowledgeResponseParams struct {
	// <p>Search results list</p>
	KnowledgeList []*KnowledgeResult `json:"KnowledgeList,omitnil,omitempty" name:"KnowledgeList"`

	// <p>token consumption statistics</p>
	TokenUsageList []*TokenUsage `json:"TokenUsageList,omitnil,omitempty" name:"TokenUsageList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *SearchKnowledgeResponseParams `json:"Response"`
}

func (r *SearchKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchResultPayload struct {
	// <p>Additional information of knowledge graph (JSON string)</p>
	GraphData *string `json:"GraphData,omitnil,omitempty" name:"GraphData"`

	// <p>List of hit image URLs</p>
	ImageUrlList []*string `json:"ImageUrlList,omitnil,omitempty" name:"ImageUrlList"`

	// <p>Table additional information (JSON string)</p>
	SheetInfo *string `json:"SheetInfo,omitnil,omitempty" name:"SheetInfo"`
}

type SimilarQuestionExtra struct {
	// <p>Text content of the similar question</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Similar question ID</p>
	SimilarQuestionId *string `json:"SimilarQuestionId,omitnil,omitempty" name:"SimilarQuestionId"`
}

type TableEnhancement struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type TokenUsage struct {
	// <p>Number of completion tokens</p>
	CompletionTokens *uint64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>prompt token count</p>
	PromptTokens *uint64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// <p>Total token count</p>
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}