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

type AICallConfig struct {
	// Digital human configuration
	DigitalHuman *DigitalHumanConfig `json:"DigitalHuman,omitnil,omitempty" name:"DigitalHuman"`

	// Enable digital human
	EnableDigitalHuman *bool `json:"EnableDigitalHuman,omitnil,omitempty" name:"EnableDigitalHuman"`

	// Enable voice call
	EnableVoiceCall *bool `json:"EnableVoiceCall,omitnil,omitempty" name:"EnableVoiceCall"`

	// Enable voice interaction
	EnableVoiceInteract *bool `json:"EnableVoiceInteract,omitnil,omitempty" name:"EnableVoiceInteract"`

	// Voice type configuration
	Voice *VoiceConfig `json:"Voice,omitnil,omitempty" name:"Voice"`
}

type AIOptimizeModel struct {
	// Model configuration
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type AccessKeyAuthConfig struct {
	// <p>Access Key field configuration</p>
	ParamList []*AccessKeyParamConfig `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>Access Key passthrough configuration</p>
	PassThroughConfig *AccessKeyPassThroughConfig `json:"PassThroughConfig,omitnil,omitempty" name:"PassThroughConfig"`

	// <p>Access Key usage mode</p><p>Enumeration values:</p><ul><li>1: Access Key pass-through</li></ul>
	UsageMode *int64 `json:"UsageMode,omitnil,omitempty" name:"UsageMode"`
}

type AccessKeyParamConfig struct {
	// <p>Access Key field type. 1: AccessKeyId, 2: AccessKeySecret, 3: SessionToken</p>
	FieldType *int64 `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// <p>Required or not</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>header/query field name</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <p>Default value of the AccessKey, which can be empty.</p>
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type AccessKeyPassThroughConfig struct {
	// <p>Unified injection position of the Access Key field. 0: Header, 1: Query</p>
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`
}

type AccountInfo struct {
	// <p>Employee sub-account id</p>
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// <p>Employee nickname</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// <p>Employee avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

type AddOnPackageInfo struct {
	// <p>Total value-added packages</p>
	AddOnTotal *float64 `json:"AddOnTotal,omitnil,omitempty" name:"AddOnTotal"`

	// <p>Value-added package amount</p>
	AddOnUsage *float64 `json:"AddOnUsage,omitnil,omitempty" name:"AddOnUsage"`

	// <p>Total number of exclusive concurrencies</p>
	ExclusiveConcurrency *uint64 `json:"ExclusiveConcurrency,omitnil,omitempty" name:"ExclusiveConcurrency"`

	// <p>Resource pack status</p><p>Enumeration values:</p><ul><li>1: Available</li><li>2: Exhausted</li><li>3: Expired</li></ul>
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// <p>Exclusive concurrency status</p><p>Enumeration values:</p><ul><li>1: Available</li><li>3: Expired</li></ul>
	ConcurrencyStatus *uint64 `json:"ConcurrencyStatus,omitnil,omitempty" name:"ConcurrencyStatus"`

	// <p>Exclusive tpm</p>
	ExclusiveTpm *uint64 `json:"ExclusiveTpm,omitnil,omitempty" name:"ExclusiveTpm"`

	// <p>Exclusive tpm status</p><p>Enumeration values:</p><ul><li>1: Available</li><li>3: Expired</li></ul>
	ExclusiveTpmStatus *uint64 `json:"ExclusiveTpmStatus,omitnil,omitempty" name:"ExclusiveTpmStatus"`

	// <p>Exclusive computing unit</p>
	ExclusiveComputeUnit *uint64 `json:"ExclusiveComputeUnit,omitnil,omitempty" name:"ExclusiveComputeUnit"`

	// <p>Exclusive computing unit status</p><p>Enumeration values:</p><ul><li>1: Available</li><li>3: Expired</li><li>4: Terminated</li><li>5: Isolated</li><li>6: Not effective</li><li>7: Unavailable (when the package expires)</li></ul>
	ExclusiveComputeUnitStatus *uint64 `json:"ExclusiveComputeUnitStatus,omitnil,omitempty" name:"ExclusiveComputeUnitStatus"`
}

type AgentAdvancedConfig struct {
	// <p>Maximum reasoning round</p>
	MaxReasoningRound *uint64 `json:"MaxReasoningRound,omitnil,omitempty" name:"MaxReasoningRound"`
}

type AgentCollaborationConfig struct {
	// Collaboration mode. Enumeration values: 1: Free transfer: Agents can freely transmit tasks with each other. 2: Workflow orchestration: Collaboration based on predefined processes. 3: Plan-and-Execute: A collaboration mode with separate planning and execution.
	AgentCollaborationMode *int64 `json:"AgentCollaborationMode,omitnil,omitempty" name:"AgentCollaborationMode"`

	// Workflow ID.
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Workflow name.
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type AgentDetail struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent Basic Configuration</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>System prompt</p>
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// <p>Model information</p>
	Model *AgentModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Tool details</p>
	ToolList []*AgentTool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>Plug-in configuration</p>
	PluginList []*AgentPlugin `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// <p>Skill details.</p>
	SkillList []*AgentSkill `json:"SkillList,omitnil,omitempty" name:"SkillList"`

	// <p>Advanced configurations.</p>
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// <p>List of Function Tools executed by the caller</p><p>Input parameter limit: only available in client user state Agent scenario. This field is ignored in business configuration state Agent</p>
	ExternalToolList []*AgentExternalToolConfig `json:"ExternalToolList,omitnil,omitempty" name:"ExternalToolList"`
}

type AgentExternalToolConfig struct {
	// <p>Tool type</p><p>Input parameter limit: currently only support "function"</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Tool name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Tool description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Tool input parameter definition</p>
	Parameters []*RequestParam `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type AgentInput struct {
	// <p>Input source type: 0 user input, 3 custom variable (API parameter)</p>
	InputType *int64 `json:"InputType,omitnil,omitempty" name:"InputType"`

	// <p>User handwriting input</p>
	UserInputValue *AgentUserInputValue `json:"UserInputValue,omitnil,omitempty" name:"UserInputValue"`

	// <p>System parameter</p>
	SystemVariable *AgentSystemVariable `json:"SystemVariable,omitnil,omitempty" name:"SystemVariable"`

	// <p>Custom Variable (API parameter)</p>
	CustomVariableId *string `json:"CustomVariableId,omitnil,omitempty" name:"CustomVariableId"`

	// <p>Environment variable parameter</p>
	EnvVariableId *string `json:"EnvVariableId,omitnil,omitempty" name:"EnvVariableId"`

	// <p>Application variable parameters</p>
	AppVariableId *string `json:"AppVariableId,omitnil,omitempty" name:"AppVariableId"`
}

type AgentModelConfig struct {
	// <p>Unique model id</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>Model alias</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>Model context length character limit</p>
	ContextWordsLimit *uint64 `json:"ContextWordsLimit,omitnil,omitempty" name:"ContextWordsLimit"`

	// <p>Instruction length character limit</p>
	InstructionsWordsLimit *uint64 `json:"InstructionsWordsLimit,omitnil,omitempty" name:"InstructionsWordsLimit"`

	// <p>Model parameters</p>
	ModelParameters *ModelParams `json:"ModelParameters,omitnil,omitempty" name:"ModelParameters"`
}

type AgentPlugin struct {
	// <p>Basic plug-in configuration</p>
	Config *AgentPluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Plugin name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Plug-in icon url</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Plugin description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Plug-in product type</p><p>Enumeration values:</p><ul><li>0: ordinary plug-in</li><li>1: connector plug-in</li></ul>
	PluginClass *int64 `json:"PluginClass,omitnil,omitempty" name:"PluginClass"`

	// <p>Plug-in status</p><p>Enumeration values:</p><ul><li>0: unknown</li><li>1: available</li><li>2: unavailable</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Plug-in authentication configuration status</p><p>Enumeration values:</p><ul><li>0: No authorization required</li><li>1: Unconfigured</li><li>2: Configured</li></ul>
	AuthConfigStatus *int64 `json:"AuthConfigStatus,omitnil,omitempty" name:"AuthConfigStatus"`
}

type AgentPluginConfig struct {
	// <p>plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Plug-in Header parameter</p>
	HeaderParameterList []*AgentPluginParameter `json:"HeaderParameterList,omitnil,omitempty" name:"HeaderParameterList"`

	// <p>Plug-in Query parameter</p>
	QueryParameterList []*AgentPluginParameter `json:"QueryParameterList,omitnil,omitempty" name:"QueryParameterList"`

	// <p>Whether to use CAM one-click authorization. This parameter is valid only when auth_type=2</p>
	EnableCamRoleAuth *bool `json:"EnableCamRoleAuth,omitnil,omitempty" name:"EnableCamRoleAuth"`

	// <p>Authorization type</p><p>Enumeration values:</p><ul><li>0: No authentication</li><li>1: API Key</li><li>2: CAM authorization</li><li>3: OAuth2.0 authorization</li></ul>
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// <p>OAuth authorization and consent mode; 0 - developer authorization; 1 - user authorization (valid only when auth_type=3)</p>
	OAuthConsent *int64 `json:"OAuthConsent,omitnil,omitempty" name:"OAuthConsent"`

	// <p>Credential configuration</p>
	CredentialConfig *AgentPluginCredentialConfig `json:"CredentialConfig,omitnil,omitempty" name:"CredentialConfig"`
}

type AgentPluginCredentialConfig struct {
	// <p>Plugin authentication value source</p><p>Enumeration values:</p><ul><li>0: Unspecified</li><li>1: Use the default authentication value of the plug-in. Only APIKey/AccessKey is supported.</li><li>2: Refer to a credential</li><li>3: Refer to a variable</li></ul>
	AuthValueSource *int64 `json:"AuthValueSource,omitnil,omitempty" name:"AuthValueSource"`

	// <p>Credential ID</p><p>Input parameter limit: required when AuthValueSource=2</p>
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// <p>Parameter configuration</p>
	ParamList []*AgentPluginCredentialParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type AgentPluginCredentialParam struct {
	// <p>Parameter location</p><p>Enumeration values:</p><ul><li>0: Header authentication</li><li>1: Query authentication</li></ul>
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`

	// <p>Parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Source of parameter value</p>
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentPluginParameter struct {
	// <p>Parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Required or not</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>Input value</p>
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentProfile struct {
	// <p>Agent Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Icon URL</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Agent role: 0=primary / 1=secondary</p>
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>Agent description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Application name.</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Developer</p>
	Developer *string `json:"Developer,omitnil,omitempty" name:"Developer"`

	// <p>Primary AgentId, read-only, cannot be updated through API modification</p>
	ParentAgentId *string `json:"ParentAgentId,omitnil,omitempty" name:"ParentAgentId"`
}

type AgentRelease struct {
	// <p>Name.</p>
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// <p>Update time, unix timestamp in seconds (s)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Action description.</p>
	ActionDescription *string `json:"ActionDescription,omitnil,omitempty" name:"ActionDescription"`

	// <p>Change to test</p>
	ReleaseMessage *string `json:"ReleaseMessage,omitnil,omitempty" name:"ReleaseMessage"`
}

type AgentReleasePreview struct {
	// <p>AgentID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Update time, unix timestamp in seconds (s)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Status, status value: 1: Add new, 2: Modify, 3: Delete</p>
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>Action description.</p>
	ActionDescription *string `json:"ActionDescription,omitnil,omitempty" name:"ActionDescription"`

	// <p>Publish a message</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>Release detail</p>
	ReleaseList []*AgentRelease `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`
}

type AgentSkill struct {
	// <p>skillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>skill name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Skill description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>skill display name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>Skill show description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill icon url</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Skill source</p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>Skill version</p>
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`
}

type AgentSkillConfig struct {
	// <p>Skill ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

type AgentSpec struct {
	// <p>Agent Basic Configuration</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>System prompt</p>
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// <p>Primary model configuration</p>
	Model *AgentModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Tool information</p>
	ToolList []*AgentToolConfig `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>Plugin information.</p>
	PluginList []*AgentPluginConfig `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// <p>Skill information.</p>
	SkillList []*AgentSkillConfig `json:"SkillList,omitnil,omitempty" name:"SkillList"`

	// <p>Advanced settings</p>
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`

	// <p>List of Function Tools executed by the caller</p><p>Input parameter limit: only available in client user state Agent scenarios. B-end configuration state Agent ignores this field and</p>
	ExternalToolList []*AgentExternalToolConfig `json:"ExternalToolList,omitnil,omitempty" name:"ExternalToolList"`
}

type AgentSummary struct {
	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent Identity Portrait</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Advanced setting; returned when scope=0</p>
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type AgentSystemVariable struct {
	// <p>System parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Configuration of dialogue history rounds; used only when Input is "dialogue history" in system variables;</p>
	DialogHistoryLimit *int64 `json:"DialogHistoryLimit,omitnil,omitempty" name:"DialogHistoryLimit"`
}

type AgentTool struct {
	// <p>Tool configuration field</p>
	Config *AgentToolBasicConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Tool name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Tool status</p><p>Enumeration values:</p><ul><li>1: available</li><li>2: unavailable</li><li>3: invalid</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Calling method</p><p>Enumeration values:</p><ul><li>0: Non-streaming</li><li>1: Streaming</li></ul>
	StreamMode *int64 `json:"StreamMode,omitnil,omitempty" name:"StreamMode"`

	// <p>Tool access mode</p><p>Enumeration values:</p><ul><li>0: Not specified</li><li>1: Read-only</li><li>2: Write/Delete</li></ul>
	ToolAccessMode *int64 `json:"ToolAccessMode,omitnil,omitempty" name:"ToolAccessMode"`
}

type AgentToolBasicConfig struct {
	// <p>plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Tool id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Tool input parameter list</p>
	InputList []*AgentToolInputParameter `json:"InputList,omitnil,omitempty" name:"InputList"`

	// <p>Tool output parameter list</p>
	OutputList []*AgentToolOutputParameter `json:"OutputList,omitnil,omitempty" name:"OutputList"`

	// <p>Tool Header parameter list</p>
	HeaderParameterList []*AgentPluginParameter `json:"HeaderParameterList,omitnil,omitempty" name:"HeaderParameterList"`

	// <p>Tool Query parameter list</p>
	QueryParameterList []*AgentPluginParameter `json:"QueryParameterList,omitnil,omitempty" name:"QueryParameterList"`

	// <p>Tool source: 0-comes from plug-in, 1-comes from workflow</p>
	ToolSource *uint64 `json:"ToolSource,omitnil,omitempty" name:"ToolSource"`

	// <p>Whether to disable</p>
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`
}

type AgentToolConfig struct {
	// <p>Tool configuration</p>
	Config *AgentToolBasicConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type AgentToolInputParameter struct {
	// <p>Tool name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Tool description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Tool parameter type</p><p>Enumeration values:</p><ul><li>0: STRING</li><li>1: INT</li><li>2: FLOAT</li><li>3: BOOL</li><li>4: OBJECT</li><li>5: ARRAY_STRING</li><li>6: ARRAY_INT</li><li>7: ARRAY_FLOAT</li><li>8: ARRAY_BOOL</li><li>9: ARRAY_OBJECT</li><li>20: ARRAY_ARRAY</li><li>99: NULL</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required or not</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>Sub-parameter, used only when the data type is OBJECT or ARRAY&lt;&gt;</p>
	SubParameterList []*AgentToolInputParameter `json:"SubParameterList,omitnil,omitempty" name:"SubParameterList"`

	// <p>Whether to hide from the model in mode</p>
	IsHidden *bool `json:"IsHidden,omitnil,omitempty" name:"IsHidden"`

	// <p>Parameters in types of OneOf</p>
	OneOfList []*AgentToolInputParameter `json:"OneOfList,omitnil,omitempty" name:"OneOfList"`

	// <p>Parameters in types of AnyOf</p>
	AnyOfList []*AgentToolInputParameter `json:"AnyOfList,omitnil,omitempty" name:"AnyOfList"`

	// <p>Source of parameter value</p>
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentToolOutputParameter struct {
	// <p>Parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Variable description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Parameter type</p><p>Enumeration values:</p><ul><li>0: STRING</li><li>1: INT</li><li>2: FLOAT</li><li>3: BOOL</li><li>4: OBJECT</li><li>5: ARRAY_STRING</li><li>6: ARRAY_INT</li><li>7: ARRAY_FLOAT</li><li>8: ARRAY_BOOL</li><li>9: ARRAY_OBJECT</li><li>20: ARRAY_ARRAY</li><li>99: NULL</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Sub-parameter, used only when the data type is OBJECT or ARRAY_OBJECT</p>
	SubParameterList []*AgentToolOutputParameter `json:"SubParameterList,omitnil,omitempty" name:"SubParameterList"`

	// <p>Parsing method</p>
	RenderMode *int64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`
}

type AgentUserInputValue struct {
	// <p>User input parameter value</p>
	ValueList []*string `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type ApiKeyAuthConfig struct {
	// Key position. HEADER/QUERY
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// Header authentication
	// | 1 | Query authentication |
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`

	// Key parameter name
	KeyParamName *string `json:"KeyParamName,omitnil,omitempty" name:"KeyParamName"`

	// Key parameter value
	KeyParamValue *string `json:"KeyParamValue,omitnil,omitempty" name:"KeyParamValue"`
}

type ApiPluginConfig struct {
	// Authorization configuration message
	AuthConfig *AuthConfig `json:"AuthConfig,omitnil,omitempty" name:"AuthConfig"`
}

type ApiToolConfig struct {
	// <p>Request body parameters</p>
	Body []*RequestParam `json:"Body,omitnil,omitempty" name:"Body"`

	// <p>Example</p>
	Example *ToolExample `json:"Example,omitnil,omitempty" name:"Example"`

	// <p>External call address of API plug-in</p>
	ExternalApiUrl *string `json:"ExternalApiUrl,omitnil,omitempty" name:"ExternalApiUrl"`

	// <p>Header</p>
	Header []*RequestParam `json:"Header,omitnil,omitempty" name:"Header"`

	// <p>Request method</p>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>Output</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// <p>Query parameter</p>
	Query []*RequestParam `json:"Query,omitnil,omitempty" name:"Query"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>STREAM_MODE_UNARY</td><td>0</td><td>Non-streaming</td></tr><tr><td>STREAM_MODE_STREAMING</td><td>1</td><td>Streaming</td></tr></tbody></table>
	StreamMode *int64 `json:"StreamMode,omitnil,omitempty" name:"StreamMode"`

	// <p>Address</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type App struct {
	// <p>Auxiliary information (sub-status/approval/appeal/search for resource/special status)</p>
	AuxiliaryInfo *AppAuxiliaryInfo `json:"AuxiliaryInfo,omitnil,omitempty" name:"AuxiliaryInfo"`

	// <p>Configuration.</p>
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Metadata</p>
	Metadata *AppMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>Application key information</p>
	SecretInfo *AppSecretInfo `json:"SecretInfo,omitnil,omitempty" name:"SecretInfo"`

	// <p>Sharing link information (including access control)</p>
	ShareUrlInfo *AppShareURLInfo `json:"ShareUrlInfo,omitnil,omitempty" name:"ShareUrlInfo"`

	// <p>Status.</p>
	Status *AppStatusInfo `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Business shared configuration</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>List of shared knowledge bases referenced by the application</p>
	SharedKbList []*AppSharedKbInfo `json:"SharedKbList,omitnil,omitempty" name:"SharedKbList"`

	// <p>portal shared configuration</p>
	PortalShareConfig *PortalShareConfig `json:"PortalShareConfig,omitnil,omitempty" name:"PortalShareConfig"`

	// <p>Application-mounted cloud disk</p>
	FileVolumeMountList []*FileVolumeMount `json:"FileVolumeMountList,omitnil,omitempty" name:"FileVolumeMountList"`
}

type AppAdvancedConf struct {
	// <p>Whether to enable context rewrite</p>
	EnableContextRewrite *bool `json:"EnableContextRewrite,omitnil,omitempty" name:"EnableContextRewrite"`

	// <p>Whether to enable image and text retrieval.</p>
	EnableImageTextRetrieval *bool `json:"EnableImageTextRetrieval,omitnil,omitempty" name:"EnableImageTextRetrieval"`

	// <p>Reply flexibility</p>
	ReplyFlexibility *uint64 `json:"ReplyFlexibility,omitnil,omitempty" name:"ReplyFlexibility"`

	// <p>Custom configuration on the dialogue side (shared across all modes, with permission to modify configuration dynamically during a dialogue)</p>
	DialogCustomConfig *DialogCustomConfig `json:"DialogCustomConfig,omitnil,omitempty" name:"DialogCustomConfig"`

	// <p>Intent attainment priority</p>
	IntentAchievement []*IntentAchievementInfo `json:"IntentAchievement,omitnil,omitempty" name:"IntentAchievement"`
}

type AppAppeal struct {
	// Configuration items in an appeal
	AppealingStatus *AppealingStatus `json:"AppealingStatus,omitnil,omitempty" name:"AppealingStatus"`
}

type AppAuxiliaryInfo struct {
	// Appeal information.
	Appeal *AppAppeal `json:"Appeal,omitnil,omitempty" name:"Appeal"`

	// Search for resource status
	SearchResourceStatus *SearchResourceStatusInfo `json:"SearchResourceStatus,omitnil,omitempty" name:"SearchResourceStatus"`

	// Special status information
	SpecialStatusInfo *SpecialStatusInfo `json:"SpecialStatusInfo,omitnil,omitempty" name:"SpecialStatusInfo"`

	// Sub-status information.
	SubStatus *AppSubStatusInfo `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`
}

type AppConfig struct {
	// Trial configuration
	Experience *AppExperienceConfig `json:"Experience,omitnil,omitempty" name:"Experience"`

	// Welcome message configuration
	Greeting *AppGreetingConfig `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// Memory configuration
	Memory *AppMemoryConfig `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Mode-related configuration
	Mode *AppModeConfig `json:"Mode,omitnil,omitempty" name:"Mode"`

	// Model configuration
	Model *AppModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// Online search configuration
	WebSearch *AppWebSearchConfig `json:"WebSearch,omitnil,omitempty" name:"WebSearch"`

	// Workflow configuration
	Workflow *AppWorkflowConfig `json:"Workflow,omitnil,omitempty" name:"Workflow"`
}

type AppExperienceConfig struct {
	// Advanced configuration
	Advanced *AppAdvancedConf `json:"Advanced,omitnil,omitempty" name:"Advanced"`

	// Dialogue experience configuration
	Conversation *ConversationExperience `json:"Conversation,omitnil,omitempty" name:"Conversation"`

	// Role configuration
	Role *RoleConfig `json:"Role,omitnil,omitempty" name:"Role"`
}

type AppFileVolumeMount struct {
	// <p>Application cloud disk configuration</p>
	FileVolumeMount *FileVolumeMount `json:"FileVolumeMount,omitnil,omitempty" name:"FileVolumeMount"`

	// <p>1: Add new 2: Modify 3: Delete</p>
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`
}

type AppGreetingConfig struct {
	// <p>Welcome message content</p>
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// <p>Opening problem list</p>
	OpeningQuestionList []*string `json:"OpeningQuestionList,omitnil,omitempty" name:"OpeningQuestionList"`
}

type AppIdentity struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application name.</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Application mode</p><p>Enumeration values:</p><ul><li>1: Standard mode</li><li>2: agent mode</li><li>3: Single workflow mode</li><li>4: ClawAgent mode</li></ul>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`
}

type AppMemoryConfig struct {
	// Whether to enable long-term memory
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Long memory duration
	LongMemoryDay *uint64 `json:"LongMemoryDay,omitnil,omitempty" name:"LongMemoryDay"`

	// Model configuration
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`

	// prompt content
	PromptContent *string `json:"PromptContent,omitnil,omitempty" name:"PromptContent"`

	// Prompt word mode. Enumeration value: 1: custom
	PromptMode *int64 `json:"PromptMode,omitnil,omitempty" name:"PromptMode"`
}

type AppMetadata struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Application mode. Enumeration values: 1: Standard mode, 2: Agent mode, 3: Single workflow mode, 4: ClawAgent mode
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// App avatar
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// Creation time (Unix timestamp in seconds).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Application description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Superapp name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Space ID.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Update time (Unix timestamp in seconds).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AppModeConfig struct {
	// Multi-agent configuration (Agent mode)
	MultiAgentConfig *MultiAgentConfig `json:"MultiAgentConfig,omitnil,omitempty" name:"MultiAgentConfig"`

	// Single workflow configuration (single workflow mode)
	SingleWorkflowConfig *SingleWorkflowConfig `json:"SingleWorkflowConfig,omitnil,omitempty" name:"SingleWorkflowConfig"`

	// ClawAgent Configuration (ClawAgent mode)
	ClawAgentConfig *ClawAgentConfig `json:"ClawAgentConfig,omitnil,omitempty" name:"ClawAgentConfig"`
}

type AppModelConfig struct {
	// AI one-click optimization model
	AiOptimizeModel *AIOptimizeModel `json:"AiOptimizeModel,omitnil,omitempty" name:"AiOptimizeModel"`

	// Real-time file parsing model
	FileParseModel *FileParseModel `json:"FileParseModel,omitnil,omitempty" name:"FileParseModel"`

	// Generate model configuration
	GenerateModel *GenerateModel `json:"GenerateModel,omitnil,omitempty" name:"GenerateModel"`

	// Multimodal QA model
	MultiModalQaModel *MultiModalQAModel `json:"MultiModalQaModel,omitnil,omitempty" name:"MultiModalQaModel"`

	// Multimodal understanding model
	MultiModalUnderstandingModel *MultiModalUnderstandingModel `json:"MultiModalUnderstandingModel,omitnil,omitempty" name:"MultiModalUnderstandingModel"`

	// Prompt rewrite model
	PromptRewriteModel *PromptRewriteModel `json:"PromptRewriteModel,omitnil,omitempty" name:"PromptRewriteModel"`

	// Thinking model configuration
	ThinkModel *ThinkModel `json:"ThinkModel,omitnil,omitempty" name:"ThinkModel"`
}

type AppOperation struct {
	// Creator
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Creator UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// Creator account (for private scenarios)
	CreatorUserAccount *string `json:"CreatorUserAccount,omitnil,omitempty" name:"CreatorUserAccount"`

	// Last modified (Unix timestamp in seconds).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Last modified by
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`

	// Modifier UIN
	UpdaterUin *string `json:"UpdaterUin,omitnil,omitempty" name:"UpdaterUin"`
}

type AppPluginConfig struct {
	// Application ID for creating a plugin based on a published application
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type AppSecretInfo struct {
	// Application key
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AppShareAccessControl struct {
	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>APP_SHARE_ACCESS_TYPE_PUBLIC</td><td>1</td><td>Public access (accessible to all users)</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_INTERNAL</td><td>2</td><td>Internal access (accessible only to organizational users)</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_ACCOUNT_WHITELIST</td><td>3</td><td>Account allowlist (accessible to specified UIN/mobile phone/mailbox/IP)</td></tr></tbody></table>
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>Whether to enable access control</p><p>Enumeration values:</p><ul><li>true: Enable</li><li>false: Disable</li></ul>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Whitelist information</p>
	Whitelist []*AppShareWhitelistItem `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

type AppShareURLInfo struct {
	// Currently effective access control configuration
	AccessControl *AppShareAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// Share URL
	ShareUrl *string `json:"ShareUrl,omitnil,omitempty" name:"ShareUrl"`
}

type AppShareWhitelistItem struct {
	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_UNSPECIFIED</td> <td>0</td> <td></td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_UIN</td> <td>1</td> <td>UIN account</td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_PHONE</td> <td>2</td> <td>Mobile number</td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_EMAIL</td> <td>3</td> <td>Email address</td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_IP</td> <td>4</td> <td>IP address</td> </tr> <tr> <td>APP_SHARE_WHITELIST_TYPE_RTX</td> <td>5</td> <td>RTX account</td> </tr> </tbody> </table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Allowlist array information</p><p>Parameter format: allowlist value</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AppSharedKbInfo struct {
	// Shared knowledge base ID
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// Name of the shared knowledge base
	KbName *string `json:"KbName,omitnil,omitempty" name:"KbName"`
}

type AppStatusInfo struct {
	// <p>Application status</p><p>Enumeration values:</p><ul><li>1: not launched</li><li>2: running</li><li>3: disabled</li><li>4: importing</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description.</p>
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`
}

type AppSubStatusInfo struct {
	// Approval record ID (valid when `sub_status_list` contains `PUBLISH_APPROVING`)
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// Application sub-status list (possibly in multiple sub-statuses at the same time)
	SubStatusList []*int64 `json:"SubStatusList,omitnil,omitempty" name:"SubStatusList"`
}

type AppSummary struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application mode. Enumeration values: 1: standard mode, 2: Agent mode, 3: single workflow mode, 4: ClawAgent mode</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>Application avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>Application name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Operation information</p>
	OperationInfo *AppOperation `json:"OperationInfo,omitnil,omitempty" name:"OperationInfo"`

	// <p>Status information</p>
	Status *AppStatusInfo `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Sub-status information</p>
	SubStatus *AppSubStatusInfo `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// <p>Resource operation permissions</p>
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`

	// <p>Application description information</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AppToolConfig struct {
	// <p>Input parameter</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>Output parameters</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type AppTrigger struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Execution configuration</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td> <td>1</td> <td>Command execution</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td> <td>2</td> <td>Workflow execution</td> </tr> </tbody> </table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>Number of failures</p>
	FailedCount *string `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// <p>Push channel configuration</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_STATUS_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_STATUS_ENABLED</td><td>1</td><td>Enablement</td></tr><tr><td>APP_TRIGGER_STATUS_PAUSED</td><td>2</td><td>Suspension</td></tr><tr><td>APP_TRIGGER_STATUS_DELETED</td><td>3</td><td>Deleted</td></tr></tbody></table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Number of successes</p>
	SuccessCount *string `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>Trigger configuration</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>Trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Trigger name</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <p>Trigger status</p>
	TriggerStatus *TriggerStatus `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>Scheduled trigger</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook trigger</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AppTriggerInstance struct {
	// <p>Application id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Session id</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Trigger creation time</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>End time.</p>
	FinishedAt *string `json:"FinishedAt,omitnil,omitempty" name:"FinishedAt"`

	// <p>Trigger running instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Request ID</p>
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// <p>Result code</p>
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// <p>Result Summary</p>
	ResultSummary *string `json:"ResultSummary,omitnil,omitempty" name:"ResultSummary"`

	// <p>Single dialogue id</p>
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_INSTANCE_SOURCE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_INSTANCE_SOURCE_APP_TRIGGER</td><td>1</td><td>Originates from an application trigger</td></tr></tbody></table>
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Trigger start execution time</p>
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>TIMER_RUN_STATUS_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>TIMER_RUN_STATUS_PENDING</td> <td>1</td> <td>Waiting to execute</td> </tr> <tr> <td>TIMER_RUN_STATUS_RUNNING</td> <td>2</td> <td>Executing</td> </tr> <tr> <td>TIMER_RUN_STATUS_RETRY_WAIT</td> <td>3</td> <td>Wait for retry</td> </tr> <tr> <td>TIMER_RUN_STATUS_SUCCESS</td> <td>4</td> <td>Successful</td> </tr> <tr> <td>TIMER_RUN_STATUS_DEAD</td> <td>5</td> <td>Failure final state (retry exhausted / non-retriable)</td> </tr> <tr> <td>TIMER_RUN_STATUS_CANCELLED</td> <td>6</td> <td>Cancelled by task pause/deletion/modification</td> </tr> </tbody> </table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>TraceId, used for logging</p>
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// <p>Trigger id</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Workflow running ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type AppTriggerParamBinding struct {
	// <p>Parameter name.</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <p>Parameter type</p><p>Enumeration values:</p><ul><li>0: string</li><li>1: integer</li><li>2: floating-point number</li><li>4: object</li><li>5: string array</li><li>6: integer array</li><li>7: floating-point number array</li><li>8: boolean value array</li><li>3: boolean value</li><li>9: object array</li><li>10: file</li><li>11: document</li><li>12: image</li><li>13: audio</li><li>14: video</li><li>15: file array</li><li>16: document array</li><li>17: image array</li><li>18: audio array</li><li>19: video array</li><li>20: nested array</li><li>22: key</li><li>99: null</li><li>100: unspecified type, for OneOf and AnyOf scenarios</li></ul>
	ParamType *int64 `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// <p>Parameter value</p>
	Value *AppTriggerParamBindingValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type AppTriggerParamBindingConfig struct {
	// <p>Bind parameter list</p>
	ParamList []*AppTriggerParamBinding `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type AppTriggerParamBindingValue struct {
	// <p>Parameter value</p>
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`

	// <p>Application variable name</p>
	VariableName *string `json:"VariableName,omitnil,omitempty" name:"VariableName"`
}

type AppTriggerParamSchema struct {
	// <p>Parameter name</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>String</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>Integer</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>Floating-point number</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>Boolean value</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>Object</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>String array</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>Integer array</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>Floating-point number array</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>Boolean value array</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>Object array</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>Nested array</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>Null</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>Unspecified type, used for OneOf and AnyOf scenarios</td></tr></tbody></table>
	ParamType *int64 `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// <p>Required or optional</p>
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// <p>Sub-parameter list</p>
	SubParamList []*AppTriggerParamSchema `json:"SubParamList,omitnil,omitempty" name:"SubParamList"`
}

type AppTriggerPromptExecuteConfig struct {
	// <p>Trigger execution prompt content</p>
	ExecutePrompt *string `json:"ExecutePrompt,omitnil,omitempty" name:"ExecutePrompt"`

	// <p>api parameter binding</p>
	ParamBindingsApi *AppTriggerParamBindingConfig `json:"ParamBindingsApi,omitnil,omitempty" name:"ParamBindingsApi"`
}

type AppTriggerRunLog struct {
	// <p>Session id</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Execution time</p>
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// <p>End time</p><p>Parameter format: YYYY-MM-DD hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_FIRE_TYPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_FIRE_TYPE_SCHEDULED</td> <td>1</td> <td>Scheduled trigger</td> </tr> <tr> <td>APP_TRIGGER_FIRE_TYPE_WEBHOOK</td> <td>2</td> <td>Webhook trigger</td> </tr> <tr> <td>APP_TRIGGER_FIRE_TYPE_MANUAL_RUN</td> <td>3</td> <td>Manual immediate execution</td> </tr> <tr> <td>APP_TRIGGER_FIRE_TYPE_TEST_RUN</td> <td>4</td> <td>Test execution</td> </tr> </tbody> </table>
	FireType *int64 `json:"FireType,omitnil,omitempty" name:"FireType"`

	// <p>Trigger instance id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>TIMER_RUN_PUSH_STATUS_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>TIMER_RUN_PUSH_STATUS_NONE</td> <td>1</td> <td>Push notification not configured</td> </tr> <tr> <td>TIMER_RUN_PUSH_STATUS_WAITING</td> <td>2</td> <td>Pending</td> </tr> <tr> <td>TIMER_RUN_PUSH_STATUS_SUCCESS</td> <td>3</td> <td>Push successful</td> </tr> <tr> <td>TIMER_RUN_PUSH_STATUS_FAILED</td> <td>4</td> <td>Push failure</td> </tr> </tbody> </table>
	PushStatus *int64 `json:"PushStatus,omitnil,omitempty" name:"PushStatus"`

	// <p>Result code</p>
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// <p>Result Overview</p>
	ResultSummary *string `json:"ResultSummary,omitnil,omitempty" name:"ResultSummary"`

	// <p>Single dialogue id</p>
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// <p>Trigger time</p><p>Parameter format: YYYY:MM:DD hh:mm:ss</p>
	ScheduledFireTime *string `json:"ScheduledFireTime,omitnil,omitempty" name:"ScheduledFireTime"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Execution start time</p><p>Parameter format: YYYY:MM:DD hh:mm:ss</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>TIMER_RUN_STATUS_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>TIMER_RUN_STATUS_PENDING</td> <td>1</td> <td>Waiting to execute</td> </tr> <tr> <td>TIMER_RUN_STATUS_RUNNING</td> <td>2</td> <td>Executing</td> </tr> <tr> <td>TIMER_RUN_STATUS_RETRY_WAIT</td> <td>3</td> <td>Wait for retry</td> </tr> <tr> <td>TIMER_RUN_STATUS_SUCCESS</td> <td>4</td> <td>Successful</td> </tr> <tr> <td>TIMER_RUN_STATUS_DEAD</td> <td>5</td> <td>Failed final state (retry exhausted / non-retriable)</td> </tr> <tr> <td>TIMER_RUN_STATUS_CANCELLED</td> <td>6</td> <td>Cancelled due to task pause/deletion/modification</td> </tr> </tbody> </table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Trigger id</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Whether it has been read</p>
	Unread *bool `json:"Unread,omitnil,omitempty" name:"Unread"`

	// <p>Visitor Id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Workflow run id</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type AppTriggerScheduleConfig struct {
	// <p>Trigger scheduled configuration.</p>
	Schedule *TimerScheduleConfig `json:"Schedule,omitnil,omitempty" name:"Schedule"`
}

type AppTriggerScheduleStatus struct {
	// <p>Last trigger time</p><p>Parameter format: YYYY-MM-DD hh:mm:ss</p>
	LastFireTime *string `json:"LastFireTime,omitnil,omitempty" name:"LastFireTime"`

	// <p>Next trigger time</p><p>Parameter format: YYYY-MM-DD hh:mm:ss</p>
	NextFireTime *string `json:"NextFireTime,omitnil,omitempty" name:"NextFireTime"`

	// <p>Trigger mode</p>
	PolicySummary *string `json:"PolicySummary,omitnil,omitempty" name:"PolicySummary"`
}

type AppTriggerSummary struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td><td>1</td><td>Command execution</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td><td>2</td><td>Workflow execution</td></tr></tbody></table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>Number of failures</p>
	FailedCount *string `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// <p>Last session id</p>
	LastSessionId *string `json:"LastSessionId,omitnil,omitempty" name:"LastSessionId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table> <p>Value range: [0, 2]</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_STATUS_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_STATUS_ENABLED</td> <td>1</td> <td>Enablement</td> </tr> <tr> <td>APP_TRIGGER_STATUS_PAUSED</td> <td>2</td> <td>Suspension</td> </tr> <tr> <td>APP_TRIGGER_STATUS_DELETED</td> <td>3</td> <td>Deleted</td> </tr> </tbody> </table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Number of successes</p>
	SuccessCount *string `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>Trigger id.</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Trigger name</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <p>Trigger execution status</p>
	TriggerStatus *TriggerStatus `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_TYPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_TYPE_SCHEDULED</td> <td>1</td> <td>Scheduled trigger</td> </tr> <tr> <td>APP_TRIGGER_TYPE_WEBHOOK</td> <td>2</td> <td>Webhook trigger</td> </tr> </tbody> </table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>The number of unread logs</p>
	UnreadRunLogCount *string `json:"UnreadRunLogCount,omitnil,omitempty" name:"UnreadRunLogCount"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AppTriggerWebhookConfig struct {
	// <p>Trigger webhook parameter configuration</p>
	ParamSchemaConfig *AppTriggerWebhookParamSchemaConfig `json:"ParamSchemaConfig,omitnil,omitempty" name:"ParamSchemaConfig"`

	// <p>key of the webhook</p>
	WebhookKey *string `json:"WebhookKey,omitnil,omitempty" name:"WebhookKey"`

	// <p>webhook key</p>
	WebhookToken *string `json:"WebhookToken,omitnil,omitempty" name:"WebhookToken"`

	// <p>webhook domain name or IP address</p>
	WebhookUrl *string `json:"WebhookUrl,omitnil,omitempty" name:"WebhookUrl"`
}

type AppTriggerWebhookParamSchemaConfig struct {
	// <p>Trigger API parameter list</p>
	SchemaList []*AppTriggerParamSchema `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`
}

type AppTriggerWebhookStatus struct {
	// <p>Push Webbook address</p>
	WebhookUrl *string `json:"WebhookUrl,omitnil,omitempty" name:"WebhookUrl"`
}

type AppTriggerWorkflowExecuteConfig struct {
	// <p>Workflow API parameter binding</p>
	ParamBindingsApi *AppTriggerParamBindingConfig `json:"ParamBindingsApi,omitnil,omitempty" name:"ParamBindingsApi"`

	// <p>Workflow parameter binding</p>
	ParamBindingsWorkflow *AppTriggerParamBindingConfig `json:"ParamBindingsWorkflow,omitnil,omitempty" name:"ParamBindingsWorkflow"`

	// <p>Workflow ID</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>Workflow name</p>
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type AppVariableValue struct {
	// <p>Variable ID</p>
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// <p>Variable value</p>
	VariableValue *string `json:"VariableValue,omitnil,omitempty" name:"VariableValue"`
}

type AppWebSearchConfig struct {
	// API key
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// Whether enabled
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// Service provider
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// Number of returned results.
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type AppWorkflowConfig struct {
	// Whether to use PDL
	EnablePDL *bool `json:"EnablePDL,omitnil,omitempty" name:"EnablePDL"`
}

type AppealingStatus struct {
	// Whether the avatar is under appeal
	AvatarInAppeal *bool `json:"AvatarInAppeal,omitnil,omitempty" name:"AvatarInAppeal"`

	// Is the fallback response included in the appeal?
	FallbackReplyInAppeal *bool `json:"FallbackReplyInAppeal,omitnil,omitempty" name:"FallbackReplyInAppeal"`

	// Whether the welcome message is required in an appeal
	GreetingInAppeal *bool `json:"GreetingInAppeal,omitnil,omitempty" name:"GreetingInAppeal"`

	// Whether the application name is under appeal
	NameInAppeal *bool `json:"NameInAppeal,omitnil,omitempty" name:"NameInAppeal"`

	// Whether the role description is in the appeal
	RoleInAppeal *bool `json:"RoleInAppeal,omitnil,omitempty" name:"RoleInAppeal"`
}

type AuditLog struct {
	// <p>Employee information</p>
	AccountInfo *AccountInfo `json:"AccountInfo,omitnil,omitempty" name:"AccountInfo"`

	// <p>Application business id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application name</p><p>Name when the operation log is triggered</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>Operation time</p><p>Parameter format: timestamp in seconds</p>
	OperateTime *string `json:"OperateTime,omitnil,omitempty" name:"OperateTime"`

	// <p>Operation type.</p>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>Operation object.</p>
	Biz *string `json:"Biz,omitnil,omitempty" name:"Biz"`

	// <p>Operation content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Unique ID for the operation.</p>
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`
}

type AuditLogMetaField struct {
	// <p>Operation log metadata key.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Name of the operation log metadata</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AuthConfig struct {
	// <p>Authorization mode.</p><p>Enumeration values:</p><ul><li>0: No authentication</li><li>1: API Key authentication</li><li>2: CAM authorization</li><li>3: OAuth 2.0 authorization</li></ul>
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// <p>API Key Authorization Configuration</p>
	ApiKeyAuthConfig *ApiKeyAuthConfig `json:"ApiKeyAuthConfig,omitnil,omitempty" name:"ApiKeyAuthConfig"`

	// <p>CAM Authorization Configuration</p>
	CamAuthConfig *CamAuthConfig `json:"CamAuthConfig,omitnil,omitempty" name:"CamAuthConfig"`

	// <p>OAuth 2.0 authorization configuration</p>
	OAuthConfig *OAuthConfig `json:"OAuthConfig,omitnil,omitempty" name:"OAuthConfig"`

	// <p>AccessKey authorization configuration</p>
	AccessKeyAuthConfig *AccessKeyAuthConfig `json:"AccessKeyAuthConfig,omitnil,omitempty" name:"AccessKeyAuthConfig"`
}

type BackgroundImage struct {
	// Brightness level
	Brightness *uint64 `json:"Brightness,omitnil,omitempty" name:"Brightness"`

	// Landscape image (pc)
	LandscapeImageUrl *string `json:"LandscapeImageUrl,omitnil,omitempty" name:"LandscapeImageUrl"`

	// Original image
	OriginalImageUrl *string `json:"OriginalImageUrl,omitnil,omitempty" name:"OriginalImageUrl"`

	// Long image (mobile phone)
	PortraitImageUrl *string `json:"PortraitImageUrl,omitnil,omitempty" name:"PortraitImageUrl"`

	// Theme color
	ThemeColor *string `json:"ThemeColor,omitnil,omitempty" name:"ThemeColor"`
}

type BasicBilling struct {
	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>UNKNOW</td><td>0</td><td></td></tr><tr><td>TOKEN</td><td>1</td><td>By token</td></tr><tr><td>PAGE_COUNT</td><td>2</td><td>By page number</td></tr><tr><td>TIMES</td><td>3</td><td>By times</td></tr><tr><td>TIMES_THOUSAND</td><td>4</td><td>By thousand times</td></tr><tr><td>SECOND</td><td>5</td><td>By duration</td></tr><tr><td>CHARACTER</td><td>6</td><td>By number of characters</td></tr><tr><td>CHARACTER_THOUSAND</td><td>7</td><td>By thousand characters</td></tr><tr><td>SHEET</td><td>8</td><td>By sheet</td></tr><tr><td>NUMBER</td><td>9</td><td>By count</td></tr></tbody></table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>Cash price</p><p>Unit: CNY</p>
	CashPrice *float64 `json:"CashPrice,omitnil,omitempty" name:"CashPrice"`

	// <p>PU price</p><p>Unit: pu</p>
	PuPrice *float64 `json:"PuPrice,omitnil,omitempty" name:"PuPrice"`
}

type BatchRunSummary struct {
	// 	Number of failures.
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 	Number of successes.
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// Total batches
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Failed subscript
	FailedIndexList []*int64 `json:"FailedIndexList,omitnil,omitempty" name:"FailedIndexList"`

	// 	Batches completed in the previous round
	JustCompletedIndexList []*int64 `json:"JustCompletedIndexList,omitnil,omitempty" name:"JustCompletedIndexList"`

	// Subscript in running
	RunningIndexList []*int64 `json:"RunningIndexList,omitnil,omitempty" name:"RunningIndexList"`

	// Subscript of success
	SuccessIndexList []*int64 `json:"SuccessIndexList,omitnil,omitempty" name:"SuccessIndexList"`
}

type BillingAttribute struct {
	// <p>Property name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Attribute value.</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type CallMCPToolRequestParams struct {
	// <p>Capacity boundary: one request can only call the tool on this plugin</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>thread on the agent-exec side</p>
	ThreadId *string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// <p>Tool Name, used for locating the tool_id in the configured tool_list of the plugin</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>JSON: Original tool input parameters, passed through to the MCP Server as-is</p>
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`
}

type CallMCPToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>Capacity boundary: one request can only call the tool on this plugin</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>thread on the agent-exec side</p>
	ThreadId *string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// <p>Tool Name, used for locating the tool_id in the configured tool_list of the plugin</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>JSON: Original tool input parameters, passed through to the MCP Server as-is</p>
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`
}

func (r *CallMCPToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallMCPToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ThreadId")
	delete(f, "ToolName")
	delete(f, "Arguments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CallMCPToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CallMCPToolResponseParams struct {
	// <p>JSON: The complete original CallToolResult, including structuredContent / isError / _meta</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CallMCPToolResponse struct {
	*tchttp.BaseResponse
	Response *CallMCPToolResponseParams `json:"Response"`
}

func (r *CallMCPToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CallMCPToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallSource struct {
	// <p>Caller ID, whose meaning is determined by subject_type (such as app_id, kb_id)</p>
	SubjectId *string `json:"SubjectId,omitnil,omitempty" name:"SubjectId"`

	// <p>Name of the calling entity</p>
	SubjectName *string `json:"SubjectName,omitnil,omitempty" name:"SubjectName"`

	// <p>Caller type: APP/KB/WIDGET/OPEN_CLAW/KB_RECALL_TEST/WORKBENCH/MODEL_API</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>METRIC_SOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>METRIC_SOURCE_TYPE_APP</td><td>1</td><td>Application development</td></tr><tr><td>METRIC_SOURCE_TYPE_KB</td><td>2</td><td>Knowledge base</td></tr><tr><td>METRIC_SOURCE_TYPE_WIDGET</td><td>3</td><td>Widget</td></tr><tr><td>METRIC_SOURCE_TYPE_OPEN_CLAW</td><td>4</td><td>ClawPro</td></tr><tr><td>METRIC_SOURCE_TYPE_KB_RECALL_TEST</td><td>5</td><td>Knowledge base recall test</td></tr><tr><td>METRIC_SOURCE_TYPE_WORKBENCH</td><td>6</td><td>Intelligent workbench</td></tr><tr><td>METRIC_SOURCE_TYPE_MODEL_API</td><td>7</td><td>Model API call</td></tr></tbody></table>
	SubjectType *int64 `json:"SubjectType,omitnil,omitempty" name:"SubjectType"`
}

type CamAuthConfig struct {
	// <p>Role name.</p>
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// <p>Key location HEADER/QUERY</p><p>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 | Header authentication |<br>| 1 | Request information authentication |</p>
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`

	// <p>SecretId field name</p>
	SecretIdName *string `json:"SecretIdName,omitnil,omitempty" name:"SecretIdName"`

	// <p>SecretKey field name</p>
	SecretKeyName *string `json:"SecretKeyName,omitnil,omitempty" name:"SecretKeyName"`

	// <p>CAM Access Key Field Configuration</p>
	ParamList []*AccessKeyParamConfig `json:"ParamList,omitnil,omitempty" name:"ParamList"`

	// <p>Whether CAM role authorization is supported.</p>
	SupportRoleAuth *bool `json:"SupportRoleAuth,omitnil,omitempty" name:"SupportRoleAuth"`
}

type CategoryModifyFields struct {
	// <p>Category name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CategoryPath struct {
	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>List of category IDs in the path from the root node</p>
	CategoryIdPath []*string `json:"CategoryIdPath,omitnil,omitempty" name:"CategoryIdPath"`

	// <p>List of category names in the path from the root node</p>
	CategoryNamePath []*string `json:"CategoryNamePath,omitnil,omitempty" name:"CategoryNamePath"`
}

type CategoryPermission struct {
	// <p>Whether the current user can add subcategories</p>
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// <p>Whether the current user can delete the category</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>Whether the current user can edit this category</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`
}

// Predefined struct for user
type CheckLabelRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of standard words to be checked (quantity: 1–100)</p>
	TermList []*string `json:"TermList,omitnil,omitempty" name:"TermList"`

	// <p>Tag ID (validates the uniqueness of standard terms under the specified tag)</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`
}

type CheckLabelRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of standard words to be checked (quantity: 1–100)</p>
	TermList []*string `json:"TermList,omitnil,omitempty" name:"TermList"`

	// <p>Tag ID (validates the uniqueness of standard terms under the specified tag)</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`
}

func (r *CheckLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "TermList")
	delete(f, "LabelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckLabelResponseParams struct {
	// <p>Validation result list</p>
	CheckList []*LabelTermCheckResult `json:"CheckList,omitnil,omitempty" name:"CheckList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckLabelResponse struct {
	*tchttp.BaseResponse
	Response *CheckLabelResponseParams `json:"Response"`
}

func (r *CheckLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckResult struct {
	// Whether the verification passed.
	Passed *bool `json:"Passed,omitnil,omitempty" name:"Passed"`

	// Failure reason (filled when passed=false)
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type ClawAgentAgentTeamConfig struct {
	// <p>Whether to enable Agent team collaboration</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>prompt content</p>
	PromptContent *string `json:"PromptContent,omitnil,omitempty" name:"PromptContent"`
}

type ClawAgentChannelContextTimeoutConfig struct {
	// <p>Whether to enable channel context timeout clearing</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Channel context timeout duration</p><p>Value range: [1, 99]</p>
	TimeoutHour *int64 `json:"TimeoutHour,omitnil,omitempty" name:"TimeoutHour"`
}

type ClawAgentConfig struct {
	// <p>Agent Team Collaboration Configuration</p>
	AgentTeamConfig *ClawAgentAgentTeamConfig `json:"AgentTeamConfig,omitnil,omitempty" name:"AgentTeamConfig"`

	// <p>Long-term memory configuration</p>
	LongMemoryConfig *ClawAgentLongMemoryConfig `json:"LongMemoryConfig,omitnil,omitempty" name:"LongMemoryConfig"`

	// <p>Channel context timeout configuration clearing</p>
	ChannelContextTimeoutConfig *ClawAgentChannelContextTimeoutConfig `json:"ChannelContextTimeoutConfig,omitnil,omitempty" name:"ChannelContextTimeoutConfig"`
}

type ClawAgentLongMemoryConfig struct {
	// <p>Whether to enable long-term memory</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CodeToolConfig struct {
	// <p>Code</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Example.</p>
	Example *ToolExample `json:"Example,omitnil,omitempty" name:"Example"`

	// <p>Input parameter</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>Output parameters</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type ComplexBilling struct {
	// <p>Composite Billing List</p>
	ComplexList []*ComplexBillingItem `json:"ComplexList,omitnil,omitempty" name:"ComplexList"`
}

type ComplexBillingItem struct {
	// <p>Composite billing dimensions</p>
	BillingAttributeList []*BillingAttribute `json:"BillingAttributeList,omitnil,omitempty" name:"BillingAttributeList"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>UNKNOW</td><td>0</td><td></td></tr><tr><td>TOKEN</td><td>1</td><td>TOKEN-based</td></tr><tr><td>PAGE_COUNT</td><td>2</td><td>By PAGE NUMBER</td></tr><tr><td>TIMES</td><td>3</td><td>TIMES-based</td></tr><tr><td>TIMES_THOUSAND</td><td>4</td><td>By THOUSAND TIMES</td></tr><tr><td>SECOND</td><td>5</td><td>Duration-based</td></tr><tr><td>CHARACTER</td><td>6</td><td>By NUMBER of characters</td></tr><tr><td>CHARACTER_THOUSAND</td><td>7</td><td>By THOUSAND characters</td></tr><tr><td>SHEET</td><td>8</td><td>By SHEET</td></tr><tr><td>NUMBER</td><td>9</td><td>By COUNT</td></tr></tbody></table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>Cash price</p><p>Unit: CNY</p>
	CashPrice *float64 `json:"CashPrice,omitnil,omitempty" name:"CashPrice"`

	// <p>pu price</p><p>Unit: pu</p>
	PuPrice *float64 `json:"PuPrice,omitnil,omitempty" name:"PuPrice"`
}

type ConcurrencyLimitDetail struct {
	// <p>Call source (subject_type determines the meaning of subject_id/subject_name. For example, for APP, subject_id=app_id and subject_name=app_name)</p>
	CallSource *CallSource `json:"CallSource,omitnil,omitempty" name:"CallSource"`

	// <p>Time when the limit was exceeded (Unix seconds)</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>Request content (the original query text of the user request)</p>
	RequestQuery *string `json:"RequestQuery,omitnil,omitempty" name:"RequestQuery"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type ConflictQA struct {
	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// <p>Knowledge application scope: 1=disabled, 2=development domain only, 3=release domain only, 4=global<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Associated document name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Associated document type</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Source type: 1=documentation generation, 2=batch import, 3=manual entry<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>QA_SOURCE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>QA_SOURCE_TYPE_DOC</td><td>1</td><td>Documentation generation</td></tr><tr><td>QA_SOURCE_TYPE_BATCH_IMPORT</td><td>2</td><td>Batch import</td></tr><tr><td>QA_SOURCE_TYPE_MANUAL</td><td>3</td><td>Manual entry</td></tr></tbody></table></p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>Update time (Unix seconds, for sorting old and new)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type ConflictQASummary struct {
	// <p>Conflict group ID</p>
	ConflictGroupId *string `json:"ConflictGroupId,omitnil,omitempty" name:"ConflictGroupId"`
}

type ConsumptionClassification struct {
	// <p>Consumption scenario (such as reasoning/training/evaluation)</p>
	ConsumptionScene *string `json:"ConsumptionScene,omitnil,omitempty" name:"ConsumptionScene"`

	// <p>Consumption target (for example, specific model name/plug-in name/platform feature name)</p>
	ConsumptionTarget *string `json:"ConsumptionTarget,omitnil,omitempty" name:"ConsumptionTarget"`

	// <p>Consumption type. The set of values is defined by the business side (for example, model/plugin/platform)</p>
	ConsumptionType *string `json:"ConsumptionType,omitnil,omitempty" name:"ConsumptionType"`

	// <p>Package name</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ConsumptionDetail struct {
	// <p>Consumption category (Type/Target/Scenario/Package)</p>
	Classification *ConsumptionClassification `json:"Classification,omitnil,omitempty" name:"Classification"`

	// <p>Consumption time, Unix seconds.</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>Amount source type</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>METRIC_SOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>METRIC_SOURCE_TYPE_APP</td><td>1</td><td>Application development</td></tr><tr><td>METRIC_SOURCE_TYPE_KB</td><td>2</td><td>Knowledge base</td></tr><tr><td>METRIC_SOURCE_TYPE_WIDGET</td><td>3</td><td>Widget</td></tr><tr><td>METRIC_SOURCE_TYPE_OPEN_CLAW</td><td>4</td><td>ClawPro</td></tr><tr><td>METRIC_SOURCE_TYPE_KB_RECALL_TEST</td><td>5</td><td>Knowledge base recall test</td></tr><tr><td>METRIC_SOURCE_TYPE_WORKBENCH</td><td>6</td><td>Intelligent workbench</td></tr><tr><td>METRIC_SOURCE_TYPE_MODEL_API</td><td>7</td><td>Model API calls</td></tr></tbody></table>
	MetricSourceType *int64 `json:"MetricSourceType,omitnil,omitempty" name:"MetricSourceType"`

	// <p>Name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Space name</p>
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// <p>Consumed usage (numeric value/unit/PU consumption)</p>
	Usage *ConsumptionUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// <p>Username</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ConsumptionUsage struct {
	// <p>Consume PU</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`

	// <p>Amount value.</p>
	Usage *float64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// <p>Usage unit. Enumeration value: DosageUnit</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token (default)</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count (page number)</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times (frequency)</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number (count)</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second (ms)</td></tr></tbody></table>
	UsageUnit *int64 `json:"UsageUnit,omitnil,omitempty" name:"UsageUnit"`
}

type ContentFilter struct {
	// <p>Minimum image height (pixel). Images with a height less than this value will be filtered; &lt;=0 means not enabled</p>
	ImageMinHeight *int64 `json:"ImageMinHeight,omitnil,omitempty" name:"ImageMinHeight"`

	// <p>Minimum image width (pixel). Images with a smaller width will be filtered; &lt;=0 means not enabled</p>
	ImageMinWidth *int64 `json:"ImageMinWidth,omitnil,omitempty" name:"ImageMinWidth"`

	// <p>Image name filter rule (multiple regular expressions separated by semicolons &#39;;&#39;, for example &#39;icon;notice;warning;info.*&#39;)</p>
	ImageNamePatterns *string `json:"ImageNamePatterns,omitnil,omitempty" name:"ImageNamePatterns"`
}

type Conversation struct {
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Session ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Session title</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>Client-side AgentId used by the session</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type ConversationAgentTask struct {
	// <p>Task content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Task number</p>
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>Task status. pending: to be executed, processing: processing, success: completed, failed: processing failure, stop: canceled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConversationContent struct {
	// <p>Text content</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>Content type. text: text, image: image, file: file, custom_variables: custom input parameter information, widget_action: widget action information</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Custom parameter data</p>
	CustomParamList []*string `json:"CustomParamList,omitnil,omitempty" name:"CustomParamList"`

	// <p>Custom parameter data</p>
	CustomParams []*string `json:"CustomParams,omitnil,omitempty" name:"CustomParams"`

	// <p>Custom variable data</p>
	CustomVariablesData *string `json:"CustomVariablesData,omitnil,omitempty" name:"CustomVariablesData"`

	// <p>Business form</p>
	EnterpriseCharts *string `json:"EnterpriseCharts,omitnil,omitempty" name:"EnterpriseCharts"`

	// <p>Tab List</p>
	OptionCardList []*string `json:"OptionCardList,omitnil,omitempty" name:"OptionCardList"`

	// <p>Tab list.</p>
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// <p>Tab mode. Enumeration values: 0-OPTION_MODE_SINGLE (single selection), 1-OPTION_MODE_MULTI (multiple selection)</p>
	OptionMode *int64 `json:"OptionMode,omitnil,omitempty" name:"OptionMode"`

	// <p>Reference marker information list</p>
	QuoteInfoList []*ConversationQuoteInfo `json:"QuoteInfoList,omitnil,omitempty" name:"QuoteInfoList"`

	// <p>List of reference superscript information</p>
	QuoteInfos []*ConversationQuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// <p>Reference source list</p>
	ReferenceList []*ConversationReference `json:"ReferenceList,omitnil,omitempty" name:"ReferenceList"`

	// <p>Reference list</p>
	References []*ConversationReference `json:"References,omitnil,omitempty" name:"References"`

	// <p>Associated record ID</p>
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// <p>Intelligent Agent Task List</p>
	TaskList []*ConversationAgentTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>Intelligent Agent Task List</p>
	Tasks []*ConversationAgentTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>Workflow input parameter</p>
	WorkflowInput *string `json:"WorkflowInput,omitnil,omitempty" name:"WorkflowInput"`

	// <p>MCP-APP call information</p>
	McpApp *ConversationMcpApp `json:"McpApp,omitnil,omitempty" name:"McpApp"`
}

type ConversationExperience struct {
	// <p>AI Call Configuration</p>
	AiCall *AICallConfig `json:"AiCall,omitnil,omitempty" name:"AiCall"`

	// <p>Background image configuration</p>
	BackgroundImage *BackgroundImage `json:"BackgroundImage,omitnil,omitempty" name:"BackgroundImage"`

	// <p>Fallback response switch</p>
	EnableFallbackReply *bool `json:"EnableFallbackReply,omitnil,omitempty" name:"EnableFallbackReply"`

	// <p>Whether to use recommended questions.</p>
	EnableRecommended *bool `json:"EnableRecommended,omitnil,omitempty" name:"EnableRecommended"`

	// <p>Whether to use online search.</p>
	EnableWebSearch *bool `json:"EnableWebSearch,omitnil,omitempty" name:"EnableWebSearch"`

	// <p>Fallback response</p>
	FallbackReply *string `json:"FallbackReply,omitnil,omitempty" name:"FallbackReply"`

	// <p>Input box configuration</p>
	InputBoxConfig *InputBoxConfig `json:"InputBoxConfig,omitnil,omitempty" name:"InputBoxConfig"`

	// <p>Output method. Enumeration values: 1: streaming, 2: non-streaming</p>
	Method *int64 `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>Recommended question generation prompt mode. Enumeration value: 1: Prompt that outputs recommended questions combined with the knowledge base only</p>
	RecommendPromptMode *int64 `json:"RecommendPromptMode,omitnil,omitempty" name:"RecommendPromptMode"`

	// <p>A new message interrupts the ongoing message</p>
	EnableNewMessageInterrupt *bool `json:"EnableNewMessageInterrupt,omitnil,omitempty" name:"EnableNewMessageInterrupt"`
}

type ConversationFileCollection struct {
	// <p>Maximum number of files</p>
	MaxFileCount *int64 `json:"MaxFileCount,omitnil,omitempty" name:"MaxFileCount"`

	// <p>Supported file type list</p>
	SupportedFileTypeList []*string `json:"SupportedFileTypeList,omitnil,omitempty" name:"SupportedFileTypeList"`

	// <p>Supported file type list</p>
	SupportedFileTypes []*string `json:"SupportedFileTypes,omitnil,omitempty" name:"SupportedFileTypes"`
}

type ConversationFileInfo struct {
	// <p>File creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Document business ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>File name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>File size</p>
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>File type</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>File download link.</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>Original path of the file in the sandbox</p>
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// <p>File moderation status</p><p>Enumeration values:</p><ul><li>1: under review</li><li>2: pass review</li><li>3: fail review</li></ul>
	AuditStatus *uint64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`
}

type ConversationMcpApp struct {
	// <p>Capacity boundary: A single request is read-only for resources belonging to this plugin</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>ui:// resource, based on which the frontend calls ReadMCPResource to pull HTML</p>
	ResourceUri *string `json:"ResourceUri,omitnil,omitempty" name:"ResourceUri"`

	// <p>thread on the agent-exec side</p>
	ThreadId *string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// <p>JSON: The complete original CallToolResult for replay when rebuilding historical sessions</p>
	ToolResult *string `json:"ToolResult,omitnil,omitempty" name:"ToolResult"`
}

type ConversationMessage struct {
	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Message icon</p>
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// <p>Message ID</p>
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// <p>Message name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Record ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>Message role</p>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>Message status. pending: pending execution, processing: in processing, success: completed, failed: processing failure, stop: canceled</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Message title</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>List of message contents</p>
	ContentList []*ConversationContent `json:"ContentList,omitnil,omitempty" name:"ContentList"`

	// <p>Message content list</p>
	Contents []*ConversationContent `json:"Contents,omitnil,omitempty" name:"Contents"`

	// <p>Type.</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ConversationQuoteInfo struct {
	// <p>Index value of the reference source</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>Reference source location</p>
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`
}

type ConversationRecordErrorInfo struct {
	// <p>Failure error code of the dialogue</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Dialogue failure error message</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ConversationRecordSummary struct {
	// <p>Reply record ID, corresponding to the record_id of the reply message in messages</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>ID of the user's question record, corresponding to the record_id of the user message in messages</p>
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// <p>Duration info of a single dialogue</p>
	TimeUsage *ConversationRecordTimeUsage `json:"TimeUsage,omitnil,omitempty" name:"TimeUsage"`

	// <p>token consumption information of a single dialogue</p>
	TokenUsage *ConversationRecordTokenUsage `json:"TokenUsage,omitnil,omitempty" name:"TokenUsage"`

	// <p>Description of a single dialogue failure; empty if successful</p>
	ErrorInfo *ConversationRecordErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>Status of a single employee assistant dialogue.</p><p>Enumeration values: </p><ul><li>pending: Pending, </li><li>processing: Processing, </li><li>success: Success, </li><li>failed: Failure, </li><li>stop: Stopped.</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConversationRecordTimeUsage struct {
	// <p>Total time consumed for a single dialogue, in milliseconds</p>
	Elapsed *string `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// <p>Time to first token, in milliseconds</p>
	FirstTokenCost *string `json:"FirstTokenCost,omitnil,omitempty" name:"FirstTokenCost"`

	// <p>Total time consumed for model inference, in milliseconds</p>
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type ConversationRecordTokenUsage struct {
	// <p>Total input token count</p>
	InputTokens *string `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>Total output token count</p>
	OutputTokens *string `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>Total token consumption</p>
	TotalTokens *string `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// <p>Total number of cache hit tokens</p>
	CachedTokens *string `json:"CachedTokens,omitnil,omitempty" name:"CachedTokens"`

	// <p>Total number of reasoning tokens</p>
	ReasoningTokens *string `json:"ReasoningTokens,omitnil,omitempty" name:"ReasoningTokens"`
}

type ConversationReference struct {
	// <p>Reference index</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>Reference name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>REFERENCE source TYPE. Enumeration values: 0-APP_REFERENCE_TYPE_UNSPECIFIED (not specified), 1-APP_REFERENCE_TYPE_QA (QA), 2-APP_REFERENCE_TYPE_SEGMENT (fragment), 3-APP_REFERENCE_TYPE_DOC (document), 4-APP_REFERENCE_TYPE_WEB_SEARCH (WEB SEARCH), 5-APP_REFERENCE_TYPE_GRAPH_RAG (GraphRAG)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ConversationResetInfo struct {
	// <p>Last reset millisecond-level timestamp</p>
	ResetTime *string `json:"ResetTime,omitnil,omitempty" name:"ResetTime"`

	// <p>Last boundary reset; this record and earlier records are no longer used as dialogue context</p>
	ResetThroughRecordId *string `json:"ResetThroughRecordId,omitnil,omitempty" name:"ResetThroughRecordId"`
}

type ConversationWidget struct {
	// <p>Base64-encoded original widget data</p>
	EncodedWidget *string `json:"EncodedWidget,omitnil,omitempty" name:"EncodedWidget"`

	// <p>payload from the user's last submission</p>
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// <p>Location of the widget in the text</p>
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// <p>widget status JSON data</p>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// <p>widget view data</p>
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// <p>component_ID</p>
	WidgetId *string `json:"WidgetId,omitnil,omitempty" name:"WidgetId"`

	// <p>Component operation ID</p>
	WidgetRunId *string `json:"WidgetRunId,omitnil,omitempty" name:"WidgetRunId"`
}

type ConversationWorkspace struct {
	// <p>Workspace ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>Storage type</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

// Predefined struct for user
type CopyAgentFromAppRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Target application ID. Required when kind=0</p>
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`

	// <p>Agent type, case-sensitive between B-end configuration Agent and C-end client user Agent</p><p>Enumeration values:</p><ul><li>0: Configuration end Agent</li><li>1: User space Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type CopyAgentFromAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Target application ID. Required when kind=0</p>
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`

	// <p>Agent type, case-sensitive between B-end configuration Agent and C-end client user Agent</p><p>Enumeration values:</p><ul><li>0: Configuration end Agent</li><li>1: User space Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

func (r *CopyAgentFromAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAgentFromAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TargetAppId")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAgentFromAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAgentFromAppResponseParams struct {
	// <p>Main Agent Id.</p>
	ParentAgentId *string `json:"ParentAgentId,omitnil,omitempty" name:"ParentAgentId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAgentFromAppResponse struct {
	*tchttp.BaseResponse
	Response *CopyAgentFromAppResponseParams `json:"Response"`
}

func (r *CopyAgentFromAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAgentFromAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAppRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// target_space_id
	TargetSpaceId *string `json:"TargetSpaceId,omitnil,omitempty" name:"TargetSpaceId"`
}

type CopyAppRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// target_space_id
	TargetSpaceId *string `json:"TargetSpaceId,omitnil,omitempty" name:"TargetSpaceId"`
}

func (r *CopyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TargetSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAppResponseParams struct {
	// new_app_id
	NewAppId *string `json:"NewAppId,omitnil,omitempty" name:"NewAppId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAppResponse struct {
	*tchttp.BaseResponse
	Response *CopyAppResponseParams `json:"Response"`
}

func (r *CopyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CorpShareConfig struct {
	// <p>Business shared switch</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Share scope type. 1: all business staff, 2: designated account, 3: designated space</p>
	ShareScope *int64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`

	// <p>Business shared application tag</p>
	TagIdList []*string `json:"TagIdList,omitnil,omitempty" name:"TagIdList"`

	// <p>Shared range information (for users, StrId is uin and Name is username; for spaces, StrId is space ID and Name is space Name)</p>
	ShareScopeList []*Identity `json:"ShareScopeList,omitnil,omitempty" name:"ShareScopeList"`
}

// Predefined struct for user
type CreateAgentRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent configuration</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Agent type, case-sensitive between B-end configuration Agent and client user-mode Agent</p><p>Enumeration values:</p><ul><li>0: Configuration-end Agent</li><li>1: User-mode Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type CreateAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent configuration</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Agent type, case-sensitive between B-end configuration Agent and client user-mode Agent</p><p>Enumeration values:</p><ul><li>0: Configuration-end Agent</li><li>1: User-mode Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

func (r *CreateAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Agent")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentResponseParams struct {
	// <p>Agent Id</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentResponseParams `json:"Response"`
}

func (r *CreateAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Application mode. Enumeration values: 1: Standard mode, 2: Agent mode, 3: Single workflow mode, 4: ClawAgent mode</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>App avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>Application description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Application name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Application mode. Enumeration values: 1: Standard mode, 2: Agent mode, 3: Single workflow mode, 4: ClawAgent mode</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>App avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>Application description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Application name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "AppMode")
	delete(f, "Avatar")
	delete(f, "Description")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application trigger execution configuration</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td> <td>1</td> <td>Command execution</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td> <td>2</td> <td>Workflow execution</td> </tr> </tbody> </table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>Third-party push configuration</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger configuration</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>Trigger name.</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>Scheduled trigger</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook trigger</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application trigger execution configuration</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td> <td>1</td> <td>Command execution</td> </tr> <tr> <td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td> <td>2</td> <td>Workflow execution</td> </tr> </tbody> </table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>Third-party push configuration</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger configuration</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>Trigger name.</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>Scheduled trigger</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook trigger</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ExecuteConfig")
	delete(f, "ExecuteType")
	delete(f, "PushConfig")
	delete(f, "Scope")
	delete(f, "TriggerConfig")
	delete(f, "TriggerName")
	delete(f, "TriggerType")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppTriggerResponseParams struct {
	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppTriggerResponseParams `json:"Response"`
}

func (r *CreateAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCategoryRequestParams struct {
	// <p>Classification type (0 is not allowed. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Category name (length: 1–64 characters)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Parent category ID</p>
	ParentCategoryId *string `json:"ParentCategoryId,omitnil,omitempty" name:"ParentCategoryId"`
}

type CreateCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Classification type (0 is not allowed. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Category name (length: 1–64 characters)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Parent category ID</p>
	ParentCategoryId *string `json:"ParentCategoryId,omitnil,omitempty" name:"ParentCategoryId"`
}

func (r *CreateCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryType")
	delete(f, "KbId")
	delete(f, "Name")
	delete(f, "ParentCategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCategoryResponseParams struct {
	// <p>ID of the successfully created category</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCategoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateCategoryResponseParams `json:"Response"`
}

func (r *CreateCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelRequestParams struct {

}

type CreateChannelRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChannelResponse struct {
	*tchttp.BaseResponse
	Response *CreateChannelResponseParams `json:"Response"`
}

func (r *CreateChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConversationRequestParams struct {
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Client-side AgentId, available when "modify configuration dynamically in conversation" is enabled in Claw mode</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type CreateConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Client-side AgentId, available when "modify configuration dynamically in conversation" is enabled in Claw mode</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *CreateConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConversationResponseParams struct {
	// <p>Conversation ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConversationResponse struct {
	*tchttp.BaseResponse
	Response *CreateConversationResponseParams `json:"Response"`
}

func (r *CreateConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKBRequestParams struct {
	// <p>Knowledge base type (0 not allowed. Parameter Value: 1=default knowledge base, 2=shared knowledge base)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KB_TYPE_DEFAULT</td><td>1</td><td>Default knowledge base</td></tr><tr><td>KB_TYPE_SHARED</td><td>2</td><td>Shared knowledge base</td></tr></tbody></table></p>
	KbType *int64 `json:"KbType,omitnil,omitempty" name:"KbType"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Writable attribute</p>
	Spec *KBSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// <p>Share sub-type: 1=ordinary, 2=wechat official account<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SHARED_KB_SUB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SHARED_KB_SUB_TYPE_NORMAL</td><td>1</td><td>Ordinary</td></tr><tr><td>SHARED_KB_SUB_TYPE_PUBLIC_ACCOUNT</td><td>2</td><td>WeChat Official Account</td></tr></tbody></table></p>
	SharedSubType *int64 `json:"SharedSubType,omitnil,omitempty" name:"SharedSubType"`
}

type CreateKBRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base type (0 not allowed. Parameter Value: 1=default knowledge base, 2=shared knowledge base)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KB_TYPE_DEFAULT</td><td>1</td><td>Default knowledge base</td></tr><tr><td>KB_TYPE_SHARED</td><td>2</td><td>Shared knowledge base</td></tr></tbody></table></p>
	KbType *int64 `json:"KbType,omitnil,omitempty" name:"KbType"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Writable attribute</p>
	Spec *KBSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// <p>Share sub-type: 1=ordinary, 2=wechat official account<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SHARED_KB_SUB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SHARED_KB_SUB_TYPE_NORMAL</td><td>1</td><td>Ordinary</td></tr><tr><td>SHARED_KB_SUB_TYPE_PUBLIC_ACCOUNT</td><td>2</td><td>WeChat Official Account</td></tr></tbody></table></p>
	SharedSubType *int64 `json:"SharedSubType,omitnil,omitempty" name:"SharedSubType"`
}

func (r *CreateKBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbType")
	delete(f, "SpaceId")
	delete(f, "Spec")
	delete(f, "SharedSubType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKBResponseParams struct {
	// <p>Knowledge base ID after creation</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKBResponse struct {
	*tchttp.BaseResponse
	Response *CreateKBResponseParams `json:"Response"`
}

func (r *CreateKBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag name (length not less than 1 character)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Tag value (standard term + synonym list), where term_id is generated by the backend and left empty upon creation</p>
	TermList []*LabelTerm `json:"TermList,omitnil,omitempty" name:"TermList"`
}

type CreateLabelRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag name (length not less than 1 character)</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Tag value (standard term + synonym list), where term_id is generated by the backend and left empty upon creation</p>
	TermList []*LabelTerm `json:"TermList,omitnil,omitempty" name:"TermList"`
}

func (r *CreateLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "Name")
	delete(f, "TermList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLabelResponseParams struct {
	// <p>Tag ID</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Tag value (standard word + synonym list, including the term_id generated by the backend)</p>
	TermList []*LabelTerm `json:"TermList,omitnil,omitempty" name:"TermList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateLabelResponseParams `json:"Response"`
}

func (r *CreateLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMsgRecordCategoryRequestParams struct {
	// <p>Category name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Parent category business ID. 0 represents a first-level category (unclassified)</p>
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type CreateMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Category name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Parent category business ID. 0 represents a first-level category (unclassified)</p>
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *CreateMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AppId")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMsgRecordCategoryResponseParams struct {
	// <p>Business ID of the created category.</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *CreateMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginRequestParams struct {
	// <p>Basic information of the plug-in.</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Plug-in type configuration</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Tool list of the plug-in</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User sub-account for login (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type CreatePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Basic information of the plug-in.</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Plug-in type configuration</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Tool list of the plug-in</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User sub-account for login (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *CreatePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Profile")
	delete(f, "Config")
	delete(f, "SpaceId")
	delete(f, "ToolList")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginResponseParams struct {
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePluginResponse struct {
	*tchttp.BaseResponse
	Response *CreatePluginResponseParams `json:"Response"`
}

func (r *CreatePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAGenerationTaskRequestParams struct {
	// <p>List of document IDs for QA to be generated (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

type CreateQAGenerationTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of document IDs for QA to be generated (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

func (r *CreateQAGenerationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAGenerationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocIdList")
	delete(f, "KbId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQAGenerationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAGenerationTaskResponseParams struct {
	// <p>Task ID list.</p>
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQAGenerationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateQAGenerationTaskResponseParams `json:"Response"`
}

func (r *CreateQAGenerationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAGenerationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAListRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Batch create (quantity: 1-20)</p>
	QaList []*QACreateSpec `json:"QaList,omitnil,omitempty" name:"QaList"`
}

type CreateQAListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Batch create (quantity: 1-20)</p>
	QaList []*QACreateSpec `json:"QaList,omitnil,omitempty" name:"QaList"`
}

func (r *CreateQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "QaList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAListResponseParams struct {
	// <p>Batch creation result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQAListResponse struct {
	*tchttp.BaseResponse
	Response *CreateQAListResponseParams `json:"Response"`
}

func (r *CreateQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>App sharing access control configuration</p>
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>Channel ID list</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>Business shared configuration</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>Release description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Batch change the debug-only knowledge in the default knowledge base to "effective in both debug and release"</p>
	IsDevToRelease *bool `json:"IsDevToRelease,omitnil,omitempty" name:"IsDevToRelease"`

	// <p>Whether to synchronize publishing as an application template</p>
	IsPublishAsTemplate *bool `json:"IsPublishAsTemplate,omitnil,omitempty" name:"IsPublishAsTemplate"`

	// <p>portal configuration</p>
	PortalShareConfig *PortalShareConfig `json:"PortalShareConfig,omitnil,omitempty" name:"PortalShareConfig"`
}

type CreateReleaseRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>App sharing access control configuration</p>
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>Channel ID list</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>Business shared configuration</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>Release description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Batch change the debug-only knowledge in the default knowledge base to "effective in both debug and release"</p>
	IsDevToRelease *bool `json:"IsDevToRelease,omitnil,omitempty" name:"IsDevToRelease"`

	// <p>Whether to synchronize publishing as an application template</p>
	IsPublishAsTemplate *bool `json:"IsPublishAsTemplate,omitnil,omitempty" name:"IsPublishAsTemplate"`

	// <p>portal configuration</p>
	PortalShareConfig *PortalShareConfig `json:"PortalShareConfig,omitnil,omitempty" name:"PortalShareConfig"`
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
	delete(f, "AppId")
	delete(f, "AppShareAccessControl")
	delete(f, "ChannelIdList")
	delete(f, "CorpShareConfig")
	delete(f, "Description")
	delete(f, "IsDevToRelease")
	delete(f, "IsPublishAsTemplate")
	delete(f, "PortalShareConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseResponseParams struct {
	// <p>need_approval</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// <p>release_id</p>
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

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
type CreateSimilarQuestionRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`
}

type CreateSimilarQuestionRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`
}

func (r *CreateSimilarQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSimilarQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "Question")
	delete(f, "Answer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSimilarQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSimilarQuestionResponseParams struct {
	// <p>Generated similar question list</p>
	QuestionList []*string `json:"QuestionList,omitnil,omitempty" name:"QuestionList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSimilarQuestionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSimilarQuestionResponseParams `json:"Response"`
}

func (r *CreateSimilarQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSimilarQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillRequestParams struct {
	// <p>Skill creation method, required; only allow</p><p>Enumeration values:</p><ul><li>1: FILE_UPLOAD (file upload)</li><li>3: AIGC (AIGC generation)</li></ul>
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// <p>skill package file address (zip); required for both FILE_UPLOAD and AIGC</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill show description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill display name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>Icon address</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Unique business identifier name of the skill (unique within the same enterprise); if not passed, it will be parsed from the skill package</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Version number</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>Version change description</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

type CreateSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill creation method, required; only allow</p><p>Enumeration values:</p><ul><li>1: FILE_UPLOAD (file upload)</li><li>3: AIGC (AIGC generation)</li></ul>
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// <p>skill package file address (zip); required for both FILE_UPLOAD and AIGC</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill show description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill display name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>Icon address</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Unique business identifier name of the skill (unique within the same enterprise); if not passed, it will be parsed from the skill package</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Version number</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>Version change description</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

func (r *CreateSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateType")
	delete(f, "FileUrl")
	delete(f, "SpaceId")
	delete(f, "DisplayDescription")
	delete(f, "DisplayName")
	delete(f, "IconUrl")
	delete(f, "Name")
	delete(f, "SkillVersion")
	delete(f, "UpdateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillResponseParams struct {
	// <p>Skill ID after successful creation</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Version ID after successful creation</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillResponseParams `json:"Response"`
}

func (r *CreateSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillShareRequestParams struct {
	// <p>Required. Application remarks (pop-up "Application Remarks")</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>Required. Original skill_id</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID, required</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Required. id of the shared version (must exceed the shared version)</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Shared configuration</p>
	CorpShareConfig *SkillCorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`
}

type CreateSkillShareRequest struct {
	*tchttp.BaseRequest
	
	// <p>Required. Application remarks (pop-up "Application Remarks")</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>Required. Original skill_id</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID, required</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Required. id of the shared version (must exceed the shared version)</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Shared configuration</p>
	CorpShareConfig *SkillCorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`
}

func (r *CreateSkillShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillShareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyRemark")
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	delete(f, "CorpShareConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillShareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillShareResponseParams struct {
	// <p>Whether the approval process is used (false indicates the shared task is created directly without approval)</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillShareResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillShareResponseParams `json:"Response"`
}

func (r *CreateSkillShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpaceRequestParams struct {
	// Workspace name. Maximum length: 30 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Space description. Maximum length: 150 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSpaceRequest struct {
	*tchttp.BaseRequest
	
	// Workspace name. Maximum length: 30 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Space description. Maximum length: 150 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpaceResponseParams struct {
	// Space id.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpaceResponseParams `json:"Response"`
}

func (r *CreateSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Variable information
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

type CreateVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Variable information
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

func (r *CreateVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Variable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVariableResponseParams struct {
	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVariableResponse struct {
	*tchttp.BaseResponse
	Response *CreateVariableResponseParams `json:"Response"`
}

func (r *CreateVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebSocketTokenRequestParams struct {
	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Application ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateWebSocketTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Application ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateWebSocketTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSocketTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebSocketTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebSocketTokenResponseParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>WebSocket Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>uid, which will be used in the subsequent DescribeUserDialogConfig api</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebSocketTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebSocketTokenResponseParams `json:"Response"`
}

func (r *CreateWebSocketTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSocketTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRunRequestParams struct {
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Execution environment. 1: Development domain, 2: Production domain</p>
	Domain *uint64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>User input.</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>API parameter configuration</p>
	CustomVariableList []*CustomVariable `json:"CustomVariableList,omitnil,omitempty" name:"CustomVariableList"`

	// <p>uid</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Execution environment. 1: Development domain, 2: Production domain</p>
	Domain *uint64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>User input.</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>API parameter configuration</p>
	CustomVariableList []*CustomVariable `json:"CustomVariableList,omitnil,omitempty" name:"CustomVariableList"`

	// <p>uid</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Domain")
	delete(f, "Query")
	delete(f, "CustomVariableList")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRunResponseParams struct {
	// <p>ID of the workflow running instance</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>Creation time (timestamp in milliseconds)</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowRunResponseParams `json:"Response"`
}

func (r *CreateWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceCredentialRequestParams struct {
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Workspace ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateWorkspaceCredentialRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Workspace ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateWorkspaceCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "WorkspaceId")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceCredentialResponseParams struct {
	// <p>Storage type</p>
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>Workspace ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkspaceCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceCredentialResponseParams `json:"Response"`
}

func (r *CreateWorkspaceCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronSchedule struct {
	// Cron expression
	Expression *string `json:"Expression,omitnil,omitempty" name:"Expression"`
}

type CustomVariable struct {
	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DBRetrievalConfig struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DailySchedule struct {
	// Time
	TimeOfDay *string `json:"TimeOfDay,omitnil,omitempty" name:"TimeOfDay"`
}

type DeDuplicateStrategy struct {
	// <p>Verification method: 1=determine whether they are identical based on document content<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DUPLICATE_FILE_CHECK_TYPE_UNKNOWN</td><td>0</td><td>Unknown</td></tr><tr><td>DUPLICATE_FILE_CHECK_TYPE_COS_HASH</td><td>1</td><td>Determine whether they are identical based on document content (cos_hash)</td></tr></tbody></table></p>
	CheckType *int64 `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// <p>Processing method: 1=return an error, 2=skip and return the duplicate document ID<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DUPLICATE_FILE_HANDLE_TYPE_UNKNOWN</td><td>0</td><td>Unknown</td></tr><tr><td>DUPLICATE_FILE_HANDLE_TYPE_RETURN_ERR</td><td>1</td><td>Return an error</td></tr><tr><td>DUPLICATE_FILE_HANDLE_TYPE_SKIP</td><td>2</td><td>Skip and return the duplicate document ID</td></tr></tbody></table></p>
	HandleType *int64 `json:"HandleType,omitnil,omitempty" name:"HandleType"`
}

// Predefined struct for user
type DeleteAgentRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>ID of the agent to delete</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Collaboration mode; 0: Claw mode; 1: Multi-Agent mode
	CollaborationMode *int64 `json:"CollaborationMode,omitnil,omitempty" name:"CollaborationMode"`
}

type DeleteAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>ID of the agent to delete</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// Collaboration mode; 0: Claw mode; 1: Multi-Agent mode
	CollaborationMode *int64 `json:"CollaborationMode,omitnil,omitempty" name:"CollaborationMode"`
}

func (r *DeleteAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "AgentId")
	delete(f, "CollaborationMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentResponseParams `json:"Response"`
}

func (r *DeleteAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppRequestParams struct {
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Reason for deletion (optional, shown during approval)</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Reason for deletion (optional, shown during approval)</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
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
	delete(f, "AppId")
	delete(f, "Reason")
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
type DeleteAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger ID.</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger ID.</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppTriggerResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppTriggerResponseParams `json:"Response"`
}

func (r *DeleteAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCategoryRequestParams struct {
	// <p>Category ID to be deleted</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Classification type (cannot be 0. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

type DeleteCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Category ID to be deleted</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Classification type (cannot be 0. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

func (r *DeleteCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryId")
	delete(f, "CategoryType")
	delete(f, "KbId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCategoryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCategoryResponseParams `json:"Response"`
}

func (r *DeleteCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteChannelRequestParams struct {

}

type DeleteChannelRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteChannelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteChannelResponseParams `json:"Response"`
}

func (r *DeleteChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConversationRequestParams struct {
	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`
}

type DeleteConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`
}

func (r *DeleteConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConversationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConversationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConversationResponseParams `json:"Response"`
}

func (r *DeleteConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocListRequestParams struct {
	// <p>List of document IDs to be deleted (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

type DeleteDocListRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of document IDs to be deleted (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

func (r *DeleteDocListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocIdList")
	delete(f, "KbId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocListResponseParams struct {
	// <p>Batch deletion result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocListResponseParams `json:"Response"`
}

func (r *DeleteDocListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKBRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteKBRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteKBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKBResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKBResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKBResponseParams `json:"Response"`
}

func (r *DeleteKBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLabelListRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag ID list to be deleted (quantity: 1-20)</p>
	LabelIdList []*string `json:"LabelIdList,omitnil,omitempty" name:"LabelIdList"`
}

type DeleteLabelListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag ID list to be deleted (quantity: 1-20)</p>
	LabelIdList []*string `json:"LabelIdList,omitnil,omitempty" name:"LabelIdList"`
}

func (r *DeleteLabelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLabelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "LabelIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLabelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLabelListResponseParams struct {
	// <p>Batch deletion result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLabelListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLabelListResponseParams `json:"Response"`
}

func (r *DeleteLabelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLabelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMsgRecordCategoryRequestParams struct {
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Business ID of the category to be deleted</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type DeleteMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Business ID of the category to be deleted</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

func (r *DeleteMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMsgRecordCategoryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *DeleteMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginRequestParams struct {
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DeletePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DeletePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePluginResponse struct {
	*tchttp.BaseResponse
	Response *DeletePluginResponseParams `json:"Response"`
}

func (r *DeletePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAListRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of QA IDs to be deleted (quantity: 1-20)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`
}

type DeleteQAListRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of QA IDs to be deleted (quantity: 1-20)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`
}

func (r *DeleteQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "QaIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAListResponseParams struct {
	// <p>Batch delete result</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQAListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQAListResponseParams `json:"Response"`
}

func (r *DeleteQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillRequestParams struct {
	// <p>Skill ID, required</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID. Required.</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill ID, required</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID. Required.</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSkillResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSkillResponseParams `json:"Response"`
}

func (r *DeleteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillShareRequestParams struct {
	// <p>Application remarks, required (pop-up "Application Remarks")</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>Original Skill ID, required (frontend does not need to detect the _shared suffix)</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID, required</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Original version ID, required (identical to the version_id passed when publishing via CreateSkillShare)</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DeleteSkillShareRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application remarks, required (pop-up "Application Remarks")</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>Original Skill ID, required (frontend does not need to detect the _shared suffix)</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID, required</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Original version ID, required (identical to the version_id passed when publishing via CreateSkillShare)</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DeleteSkillShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillShareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyRemark")
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSkillShareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillShareResponseParams struct {
	// <p>Whether to go through the approval process (false means the removal has been directly executed without approval)</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSkillShareResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSkillShareResponseParams `json:"Response"`
}

func (r *DeleteSkillShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpaceRequestParams struct {
	// Space id.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteSpaceRequest struct {
	*tchttp.BaseRequest
	
	// Space id.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpaceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSpaceResponseParams `json:"Response"`
}

func (r *DeleteSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

type DeleteVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

func (r *DeleteVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "VariableId")
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVariableResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVariableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVariableResponseParams `json:"Response"`
}

func (r *DeleteVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountListRequestParams struct {
	// <p>Page number</p><p>Starts from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Pagination quantity</p><p>Value range: [1, 100]</p><p>Unit: count</p><p>Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Parameter filtering</p><p>Supports filter queries by SpaceId and NIckName</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAccountListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Page number</p><p>Starts from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Pagination quantity</p><p>Value range: [1, 100]</p><p>Unit: count</p><p>Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Parameter filtering</p><p>Supports filter queries by SpaceId and NIckName</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountListResponseParams struct {
	// <p>Total.</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Employee list</p>
	AccountList []*AccountInfo `json:"AccountList,omitnil,omitempty" name:"AccountList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountListResponseParams `json:"Response"`
}

func (r *DescribeAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDetailRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeAgentDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeAgentDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDetailResponseParams struct {
	// <p>Agent information</p>
	Agent *AgentDetail `json:"Agent,omitnil,omitempty" name:"Agent"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDetailResponseParams `json:"Response"`
}

func (r *DescribeAgentDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentReleasePreviewListRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Page number.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page: between 1 and 200</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query keyword, used for fuzzy matching of heading</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Filter criteria</p><p>Input parameter limits: supports StartTime, EndTime, ActionList, ReleaseStatusList</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAgentReleasePreviewListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Page number.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page: between 1 and 200</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query keyword, used for fuzzy matching of heading</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Filter criteria</p><p>Input parameter limits: supports StartTime, EndTime, ActionList, ReleaseStatusList</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAgentReleasePreviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentReleasePreviewListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentReleasePreviewListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentReleasePreviewListResponseParams struct {
	// <p>Publish preview list</p>
	ReleaseList []*AgentReleasePreview `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentReleasePreviewListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentReleasePreviewListResponseParams `json:"Response"`
}

func (r *DescribeAgentReleasePreviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentReleasePreviewListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSummaryListRequestParams struct {
	// <p>Query scope; 0-single application query; 1-across application query</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application Id. When Scope=0, it is the target application Id (required); when Scope=1, it is not required</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Filter criteria (name: "SearchWord", "SpaceId", "AgentSource", "AppId")</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Number of items per page.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Page number.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeAgentSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Query scope; 0-single application query; 1-across application query</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application Id. When Scope=0, it is the target application Id (required); when Scope=1, it is not required</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Filter criteria (name: "SearchWord", "SpaceId", "AgentSource", "AppId")</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Number of items per page.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Page number.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeAgentSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scope")
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSummaryListResponseParams struct {
	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Agent digest information</p>
	AgentList []*AgentSummary `json:"AgentList,omitnil,omitempty" name:"AgentList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAgentSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application domain: ADP_DOMAIN_DEV(1)=development domain, ADP_DOMAIN_PROD(2)=release domain. Enumeration values: 1: development domain, 2: production domain</p>
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Field mask, specify the fields to return (return all fields if Paths is empty). Paths enumeration values: AppConfig (application configuration), SecretInfo (application key info), ShareUrlInfo (sharing link information), SpecialStatusInfo (special status information), SearchResourceStatus (search resource status), SharedKbList (list of shared knowledge bases referenced by the application), CorpShareConfig (corporate shared configuration), FileVolumeMountList (cloud disk mount list)</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>Specific status type (required when FieldMask includes SpecialStatusInfo). Enumeration values: 1: rollback status, 2: first-time import status</p>
	StatusType *int64 `json:"StatusType,omitnil,omitempty" name:"StatusType"`
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application domain: ADP_DOMAIN_DEV(1)=development domain, ADP_DOMAIN_PROD(2)=release domain. Enumeration values: 1: development domain, 2: production domain</p>
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Field mask, specify the fields to return (return all fields if Paths is empty). Paths enumeration values: AppConfig (application configuration), SecretInfo (application key info), ShareUrlInfo (sharing link information), SpecialStatusInfo (special status information), SearchResourceStatus (search resource status), SharedKbList (list of shared knowledge bases referenced by the application), CorpShareConfig (corporate shared configuration), FileVolumeMountList (cloud disk mount list)</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>Specific status type (required when FieldMask includes SpecialStatusInfo). Enumeration values: 1: rollback status, 2: first-time import status</p>
	StatusType *int64 `json:"StatusType,omitnil,omitempty" name:"StatusType"`
}

func (r *DescribeAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Domain")
	delete(f, "FieldMask")
	delete(f, "StatusType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppResponseParams struct {
	// <p>App details</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppResponseParams `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppStatisticsOverviewRequestParams struct {
	// <p>Application type. Available values: APP_TYPE_RAG (RAG), APP_TYPE_WORKFLOW (Workflow), APP_TYPE_CLAW (Claw), APP_TYPE_MULTIAGENT (MultiAgent). In the current period, only APP_TYPE_RAG takes effect. The rest are reserved values. Importing them will return InvalidParameter.</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified (invalid value, do not import in requests)</td></tr><tr><td>APP_TYPE_RAG</td><td>1</td><td>RAG (uniquely supported in the current period)</td></tr><tr><td>APP_TYPE_WORKFLOW</td><td>2</td><td>Workflow (reserved, not currently supported)</td></tr><tr><td>APP_TYPE_CLAW</td><td>3</td><td>Claw (reserved, not currently supported)</td></tr><tr><td>APP_TYPE_MULTIAGENT</td><td>4</td><td>MultiAgent (reserved, not currently supported)</td></tr></tbody></table>
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: application view (VIEW_TYPE_APP). For scope_id, enter the app ID (decimal string of the uint64 snowflake ID)</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Space ID</p><p>Parameter format: <p>Space ID</p></p><p>Default value: default_space</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Application ID</p><p>Parameter format:</p><p>Application ID</p><p></p><p>Reference value: 2099767969573745984</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Extended filtering. Filter combination rule: multiple items use AND, and value_list under the same item uses OR. Supported Name: space_id (Space ID), channel_type (calling method/channel type; for values, see the trpc.adp.common.v2.ChannelType Enumeration Name). If channel_type is not specified, aggregation is performed by all channels.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAppStatisticsOverviewRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application type. Available values: APP_TYPE_RAG (RAG), APP_TYPE_WORKFLOW (Workflow), APP_TYPE_CLAW (Claw), APP_TYPE_MULTIAGENT (MultiAgent). In the current period, only APP_TYPE_RAG takes effect. The rest are reserved values. Importing them will return InvalidParameter.</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>APP_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified (invalid value, do not import in requests)</td></tr><tr><td>APP_TYPE_RAG</td><td>1</td><td>RAG (uniquely supported in the current period)</td></tr><tr><td>APP_TYPE_WORKFLOW</td><td>2</td><td>Workflow (reserved, not currently supported)</td></tr><tr><td>APP_TYPE_CLAW</td><td>3</td><td>Claw (reserved, not currently supported)</td></tr><tr><td>APP_TYPE_MULTIAGENT</td><td>4</td><td>MultiAgent (reserved, not currently supported)</td></tr></tbody></table>
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: application view (VIEW_TYPE_APP). For scope_id, enter the app ID (decimal string of the uint64 snowflake ID)</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Space ID</p><p>Parameter format: <p>Space ID</p></p><p>Default value: default_space</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Application ID</p><p>Parameter format:</p><p>Application ID</p><p></p><p>Reference value: 2099767969573745984</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Extended filtering. Filter combination rule: multiple items use AND, and value_list under the same item uses OR. Supported Name: space_id (Space ID), channel_type (calling method/channel type; for values, see the trpc.adp.common.v2.ChannelType Enumeration Name). If channel_type is not specified, aggregation is performed by all channels.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAppStatisticsOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "SpaceId")
	delete(f, "AppId")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppStatisticsOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppStatisticsOverviewResponseParams struct {
	// <p>Average first tokens duration (ms)</p>
	AvgFirstTokenTime *string `json:"AvgFirstTokenTime,omitnil,omitempty" name:"AvgFirstTokenTime"`

	// <p>Avg duration of total tokens (ms)</p>
	AvgTotalTokenTime *string `json:"AvgTotalTokenTime,omitnil,omitempty" name:"AvgTotalTokenTime"`

	// <p>Application call success rate (percentage, 0-100)</p>
	CallSuccessRate *float64 `json:"CallSuccessRate,omitnil,omitempty" name:"CallSuccessRate"`

	// <p>Reply type distribution list; statistics by app_type, all reply methods are completed and returned in a fixed sequence. For reply methods with no data, call_count is 0</p>
	ReplyTypeDistributionList []*Distribution `json:"ReplyTypeDistributionList,omitnil,omitempty" name:"ReplyTypeDistributionList"`

	// <p>Total number of calls</p>
	TotalCallCount *string `json:"TotalCallCount,omitnil,omitempty" name:"TotalCallCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppStatisticsOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppStatisticsOverviewResponseParams `json:"Response"`
}

func (r *DescribeAppStatisticsOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSummaryListRequestParams struct {
	// Space ID (required)
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Filter criteria (multiple filters are in an AND relationship with each other, AND multiple Values of the same Filter are in an OR relationship): - AppStatus: APP STATUS, enumeration value, exact match (APP_STATUS_OFFLINE=1/APP_STATUS_RUNNING=2/APP_STATUS_DISABLED=3) - AppMode: application MODE, enumeration value, exact match (APP_MODE_STANDARD=1/APP_MODE_AGENT=2/APP_MODE_SINGLE_WORKFLOW=3/APP_MODE_CLAW_AGENT=4)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy query
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeAppSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// Space ID (required)
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Filter criteria (multiple filters are in an AND relationship with each other, AND multiple Values of the same Filter are in an OR relationship): - AppStatus: APP STATUS, enumeration value, exact match (APP_STATUS_OFFLINE=1/APP_STATUS_RUNNING=2/APP_STATUS_DISABLED=3) - AppMode: application MODE, enumeration value, exact match (APP_MODE_STANDARD=1/APP_MODE_AGENT=2/APP_MODE_SINGLE_WORKFLOW=3/APP_MODE_CLAW_AGENT=4)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy query
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeAppSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSummaryListResponseParams struct {
	// Application abstract list
	AppSummaryList []*AppSummary `json:"AppSummaryList,omitnil,omitempty" name:"AppSummaryList"`

	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAppSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerInstanceRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Trigger running instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Trigger running instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "InstanceId")
	delete(f, "Scope")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerInstanceResponseParams struct {
	// <p>Application trigger instance</p>
	Instance *AppTriggerInstance `json:"Instance,omitnil,omitempty" name:"Instance"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerInstanceResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID.</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID.</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID.</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerResponseParams struct {
	// <p>Application trigger information</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRunLogListRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Filter parameters.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number</p><p>Value range: [1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of records per page</p><p>Value range: [1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerRunLogListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Filter parameters.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number</p><p>Value range: [1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of records per page</p><p>Value range: [1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerRunLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRunLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerRunLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRunLogListResponseParams struct {
	// <p>Log list.</p>
	RunLogList []*AppTriggerRunLog `json:"RunLogList,omitnil,omitempty" name:"RunLogList"`

	// <p>Number of log lists</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerRunLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerRunLogListResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerRunLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRunLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerSummaryListRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Parameter filter list</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number</p><p>Value range: [1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Size per page</p><p>Value range: [1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Blurry query string</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Parameter filter list</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number</p><p>Value range: [1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Size per page</p><p>Value range: [1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Blurry query string</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "Scope")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerSummaryListResponseParams struct {
	// <p>Number of application triggers</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Application trigger list</p>
	TriggerList []*AppTriggerSummary `json:"TriggerList,omitnil,omitempty" name:"TriggerList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogListRequestParams struct {
	// <p>Space id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Number of items per page</p><p>Value range: [1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query start position of ES</p><p>The corresponding API returns SearchAfter</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// <p>Parameter filtering</p><p>Supports Action, BizObject, Content<br>Supports SpaceId, AccountUin, AppId (up to 100)<br>Supports startTime, endTime (timestamp in seconds)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAuditLogListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Space id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Number of items per page</p><p>Value range: [1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Query start position of ES</p><p>The corresponding API returns SearchAfter</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// <p>Parameter filtering</p><p>Supports Action, BizObject, Content<br>Supports SpaceId, AccountUin, AppId (up to 100)<br>Supports startTime, endTime (timestamp in seconds)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAuditLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Limit")
	delete(f, "SearchAfter")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogListResponseParams struct {
	// <p>Operation log list.</p>
	AuditLogList []*AuditLog `json:"AuditLogList,omitnil,omitempty" name:"AuditLogList"`

	// <p>es query start position</p><p>Used for querying the next page</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogListResponseParams `json:"Response"`
}

func (r *DescribeAuditLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogMetaRequestParams struct {

}

type DescribeAuditLogMetaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAuditLogMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogMetaResponseParams struct {
	// <p>Operation type list</p>
	Actions []*AuditLogMetaField `json:"Actions,omitnil,omitempty" name:"Actions"`

	// <p>List of operation objects.</p>
	BizObjects []*AuditLogMetaField `json:"BizObjects,omitnil,omitempty" name:"BizObjects"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogMetaResponseParams `json:"Response"`
}

func (r *DescribeAuditLogMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCategoryListRequestParams struct {
	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): CategoryType - classification type, enumeration value, exact match (CATEGORY_TYPE_DOC=1/CATEGORY_TYPE_QA=2); ParentCategoryId - parent category ID, exact match</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Pagination page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 10. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

type DescribeCategoryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): CategoryType - classification type, enumeration value, exact match (CATEGORY_TYPE_DOC=1/CATEGORY_TYPE_QA=2); ParentCategoryId - parent category ID, exact match</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Pagination page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 10. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

func (r *DescribeCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SummaryListSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCategoryListResponseParams struct {
	// <p>Classification list</p>
	CategoryList []*KBCategory `json:"CategoryList,omitnil,omitempty" name:"CategoryList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCategoryListResponseParams `json:"Response"`
}

func (r *DescribeCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelListRequestParams struct {

}

type DescribeChannelListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeChannelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelListResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChannelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelListResponseParams `json:"Response"`
}

func (r *DescribeChannelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelRequestParams struct {

}

type DescribeChannelRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChannelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelResponseParams `json:"Response"`
}

func (r *DescribeChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyLimitDetailListRequestParams struct {
	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: business view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering. Filter combination rule: multiple items use AND, and value_list within the same item uses OR. Supported Names: concurrency_type (qpm_tpm/dedicated, default qpm_tpm), model_name (required), space_id, app_id/resource_id/source_id (application ID, multiple selection), metric_source_type (METRIC_SOURCE_TYPE_* Enumeration Name or integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeConcurrencyLimitDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: business view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering. Filter combination rule: multiple items use AND, and value_list within the same item uses OR. Supported Names: concurrency_type (qpm_tpm/dedicated, default qpm_tpm), model_name (required), space_id, app_id/resource_id/source_id (application ID, multiple selection), metric_source_type (METRIC_SOURCE_TYPE_* Enumeration Name or integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeConcurrencyLimitDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyLimitDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyLimitDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyLimitDetailListResponseParams struct {
	// <p>Concurrency limit exceeded details list</p>
	ConcurrencyLimitDetailList []*ConcurrencyLimitDetail `json:"ConcurrencyLimitDetailList,omitnil,omitempty" name:"ConcurrencyLimitDetailList"`

	// <p>Total record count, for frontend pagination</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyLimitDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyLimitDetailListResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyLimitDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyLimitDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConflictQARequestParams struct {
	// <p>Conflict group ID</p>
	ConflictGroupId *string `json:"ConflictGroupId,omitnil,omitempty" name:"ConflictGroupId"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

type DescribeConflictQARequest struct {
	*tchttp.BaseRequest
	
	// <p>Conflict group ID</p>
	ConflictGroupId *string `json:"ConflictGroupId,omitnil,omitempty" name:"ConflictGroupId"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

func (r *DescribeConflictQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConflictQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConflictGroupId")
	delete(f, "KbId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConflictQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConflictQAResponseParams struct {
	// <p>List of details of a conflict group</p>
	ConflictQaList []*ConflictQA `json:"ConflictQaList,omitnil,omitempty" name:"ConflictQaList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConflictQAResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConflictQAResponseParams `json:"Response"`
}

func (r *DescribeConflictQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConflictQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConflictQASummaryListRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Common filter (supports filter by status: PENDING/RESOLVED)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 10. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeConflictQASummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Common filter (supports filter by status: PENDING/RESOLVED)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 10. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeConflictQASummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConflictQASummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConflictQASummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConflictQASummaryListResponseParams struct {
	// <p>Conflict question list</p>
	ConflictQaList []*ConflictQASummary `json:"ConflictQaList,omitnil,omitempty" name:"ConflictQaList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConflictQASummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConflictQASummaryListResponseParams `json:"Response"`
}

func (r *DescribeConflictQASummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConflictQASummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumptionDetailListRequestParams struct {
	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: business view / space view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering. Filter combination rule: multiple conditions use AND, values within the same condition use OR. Supported Name: metric_source_type (metric_source_type_* OR integer), source_ids (multi-option source ID), resource_ID/source_ID (select one source ID, this parameter is valid only when source_ids is not passed), space_ID, user_ID</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeConsumptionDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Query time range (Unix seconds)</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: business view / space view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering. Filter combination rule: multiple conditions use AND, values within the same condition use OR. Supported Name: metric_source_type (metric_source_type_* OR integer), source_ids (multi-option source ID), resource_ID/source_ID (select one source ID, this parameter is valid only when source_ids is not passed), space_ID, user_ID</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeConsumptionDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumptionDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumptionDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumptionDetailListResponseParams struct {
	// <p>Resource consumption details list</p>
	ConsumptionDetailList []*ConsumptionDetail `json:"ConsumptionDetailList,omitnil,omitempty" name:"ConsumptionDetailList"`

	// <p>Total record count for frontend pagination</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumptionDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumptionDetailListResponseParams `json:"Response"`
}

func (r *DescribeConsumptionDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumptionDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationListRequestParams struct {
	// <p>Session type. Pass CONVERSATION_TYPE_UNSPECIFIED to indicate all. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Number limit (integer), used in conjunction with Offset</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sub-account Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root account Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Offset (integer), used in conjunction with Limit, starts from 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Client-side AgentId, used when querying sessions created based on the client-side AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeConversationListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session type. Pass CONVERSATION_TYPE_UNSPECIFIED to indicate all. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Keyword</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>Number limit (integer), used in conjunction with Offset</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sub-account Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root account Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Offset (integer), used in conjunction with Limit, starts from 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Client-side AgentId, used when querying sessions created based on the client-side AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeConversationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "Keyword")
	delete(f, "Limit")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "Offset")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationListResponseParams struct {
	// <p>Session list</p>
	ConversationList []*Conversation `json:"ConversationList,omitnil,omitempty" name:"ConversationList"`

	// <p>Session list</p>
	Conversations []*Conversation `json:"Conversations,omitnil,omitempty" name:"Conversations"`

	// <p>Total.</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationListResponseParams `json:"Response"`
}

func (r *DescribeConversationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationMessageListRequestParams struct {
	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Total number of returned records. Default: 10. Maximum: 50. When querying forward or backward, messages of the record specified by record_id are excluded. When querying from the center in both directions, the message of the record specified by record_id is included. The number of returned records is limit / 2 for each direction, rounded up.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sub-user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Query anchor record ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>Query direction relative to record_id. Enumeration values: 0-RECORD_QUERY_DIRECTION_UNSPECIFIED (not specified, compatible with legacy logic, query forward by default), 1-RECORD_QUERY_DIRECTION_BACKWARD (query earlier messages forward from record_id), 2-RECORD_QUERY_DIRECTION_FORWARD (query newer messages backward from record_id), 3-RECORD_QUERY_DIRECTION_BIDIRECTIONAL (query both forward and backward from record_id as the center)</p>
	RecordQueryDirection *int64 `json:"RecordQueryDirection,omitnil,omitempty" name:"RecordQueryDirection"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeConversationMessageListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Total number of returned records. Default: 10. Maximum: 50. When querying forward or backward, messages of the record specified by record_id are excluded. When querying from the center in both directions, the message of the record specified by record_id is included. The number of returned records is limit / 2 for each direction, rounded up.</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Sub-user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Root user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Query anchor record ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>Query direction relative to record_id. Enumeration values: 0-RECORD_QUERY_DIRECTION_UNSPECIFIED (not specified, compatible with legacy logic, query forward by default), 1-RECORD_QUERY_DIRECTION_BACKWARD (query earlier messages forward from record_id), 2-RECORD_QUERY_DIRECTION_FORWARD (query newer messages backward from record_id), 3-RECORD_QUERY_DIRECTION_BIDIRECTIONAL (query both forward and backward from record_id as the center)</p>
	RecordQueryDirection *int64 `json:"RecordQueryDirection,omitnil,omitempty" name:"RecordQueryDirection"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeConversationMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "Limit")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "RecordId")
	delete(f, "RecordQueryDirection")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationMessageListResponseParams struct {
	// <p>First record ID</p>
	FirstRecordId *string `json:"FirstRecordId,omitnil,omitempty" name:"FirstRecordId"`

	// <p>Whether there are more update message directions</p>
	HasMoreAfter *bool `json:"HasMoreAfter,omitnil,omitempty" name:"HasMoreAfter"`

	// <p>Earlier message orientation whether there is more</p>
	HasMoreBefore *bool `json:"HasMoreBefore,omitnil,omitempty" name:"HasMoreBefore"`

	// <p>Last record ID</p>
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// <p>Message list</p>
	MessageList []*ConversationMessage `json:"MessageList,omitnil,omitempty" name:"MessageList"`

	// <p>Message list</p>
	//
	// Deprecated: Messages is deprecated.
	Messages []*ConversationMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// <p>Statistics list of records in a single dialogue, associated with message_list via record_id / related_record_id</p>
	RecordSummaryList []*ConversationRecordSummary `json:"RecordSummaryList,omitnil,omitempty" name:"RecordSummaryList"`

	// <p>Last reset information</p>
	ResetInfo *ConversationResetInfo `json:"ResetInfo,omitnil,omitempty" name:"ResetInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationMessageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationMessageListResponseParams `json:"Response"`
}

func (r *DescribeConversationMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationRequestParams struct {
	// <p>Session ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Primary user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Sub-user Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Primary user Uin.</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Sub-user Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationResponseParams struct {
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Session ID.</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Creation time.</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Update time.</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Workspace</p>
	Workspace *ConversationWorkspace `json:"Workspace,omitnil,omitempty" name:"Workspace"`

	// <p>Session title.</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>Client-side AgentId used for the session.</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationResponseParams `json:"Response"`
}

func (r *DescribeConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocRequestParams struct {
	// <p>Document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: Currently supported Path: DocLink.CosUrl. Other unlisted fields are returned by default.</p>
	ReadMask *FieldMask `json:"ReadMask,omitnil,omitempty" name:"ReadMask"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// <p>Document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: Currently supported Path: DocLink.CosUrl. Other unlisted fields are returned by default.</p>
	ReadMask *FieldMask `json:"ReadMask,omitnil,omitempty" name:"ReadMask"`
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
	delete(f, "DocId")
	delete(f, "KbId")
	delete(f, "ReadMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocResponseParams struct {
	// <p>Documentation link (external URL + COS link)</p>
	DocLink *DocLink `json:"DocLink,omitnil,omitempty" name:"DocLink"`

	// <p>Parsing configuration (split rule, content filtering, etc.)</p>
	ParseConfig *DocParseConfig `json:"ParseConfig,omitnil,omitempty" name:"ParseConfig"`

	// <p>Basic document information</p>
	Summary *DocSummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>Switch configuration</p>
	Switch *DocSwitch `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>Update cycle configuration</p>
	UpdatePeriod *DocUpdatePeriod `json:"UpdatePeriod,omitnil,omitempty" name:"UpdatePeriod"`

	// <p>User access configuration</p>
	UserAccessConfig *UserAccessConfig `json:"UserAccessConfig,omitnil,omitempty" name:"UserAccessConfig"`

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
type DescribeDocSummaryListRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status - document status, enumeration value, exact match; CategoryId - category ID, exact match; SourceType - document source type, enumeration value, exact match; EffectiveDomain - application scope, exact match; CreateTime - creation time, Unix seconds, BETWEEN with [start second, end second]; UpdateTime - update time, Unix seconds, BETWEEN with [start second, end second]</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 10. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query condition (keyword + query scope)</p>
	Query *DocQuery `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

type DescribeDocSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status - document status, enumeration value, exact match; CategoryId - category ID, exact match; SourceType - document source type, enumeration value, exact match; EffectiveDomain - application scope, exact match; CreateTime - creation time, Unix seconds, BETWEEN with [start second, end second]; UpdateTime - update time, Unix seconds, BETWEEN with [start second, end second]</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 10. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query condition (keyword + query scope)</p>
	Query *DocQuery `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

func (r *DescribeDocSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SummaryListSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocSummaryListResponseParams struct {
	// <p>Document list</p>
	DocList []*DocSummary `json:"DocList,omitnil,omitempty" name:"DocList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDocSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocSummaryListResponseParams `json:"Response"`
}

func (r *DescribeDocSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeKBRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeKBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBResponseParams struct {
	// <p>Associated application ID list</p>
	AppIdList []*string `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// <p>Capacity information.</p>
	CapacityInfo *KBCapacity `json:"CapacityInfo,omitnil,omitempty" name:"CapacityInfo"`

	// <p>ES configuration</p>
	EsConfig *ESConfig `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// <p>Model configuration</p>
	ModelConfig *KBModelConfig `json:"ModelConfig,omitnil,omitempty" name:"ModelConfig"`

	// <p>Owner information</p>
	Owner *Operator `json:"Owner,omitnil,omitempty" name:"Owner"`

	// <p>Knowledge base digest information</p>
	Summary *KBSummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBResponseParams `json:"Response"`
}

func (r *DescribeKBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBSummaryListRequestParams struct {
	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Common filter</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 10. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeKBSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Common filter</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default value: 10. Maximum value: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeKBSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBSummaryListResponseParams struct {
	// <p>Knowledge base list</p>
	KbList []*KBSummary `json:"KbList,omitnil,omitempty" name:"KbList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBSummaryListResponseParams `json:"Response"`
}

func (r *DescribeKBSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLabelRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag ID.</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Common filter.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Standard term ID of the scroll loading cursor (pass 0 for the first request, and subsequently pass the TermId of the last item on the previous page)</p>
	LastTermId *string `json:"LastTermId,omitnil,omitempty" name:"LastTermId"`

	// <p>Number of items to load at a time. Default: 10. Maximum: 100</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeLabelRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Tag ID.</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Common filter.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Standard term ID of the scroll loading cursor (pass 0 for the first request, and subsequently pass the TermId of the last item on the previous page)</p>
	LastTermId *string `json:"LastTermId,omitnil,omitempty" name:"LastTermId"`

	// <p>Number of items to load at a time. Default: 10. Maximum: 100</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "LabelId")
	delete(f, "FilterList")
	delete(f, "LastTermId")
	delete(f, "Limit")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLabelResponseParams struct {
	// <p>Basic information.</p>
	Summary *LabelSummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLabelResponseParams `json:"Response"`
}

func (r *DescribeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLabelSummaryListRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Common filter.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number for pagination, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default 10, maximum 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

type DescribeLabelSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Common filter.</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number for pagination, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default 10, maximum 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Keyword search</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

func (r *DescribeLabelSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLabelSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SummaryListSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLabelSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLabelSummaryListResponseParams struct {
	// <p>Tag list.</p>
	LabelList []*LabelSummary `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLabelSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLabelSummaryListResponseParams `json:"Response"`
}

func (r *DescribeLabelSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLabelSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestReleaseRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeLatestReleaseRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeLatestReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLatestReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestReleaseResponseParams struct {
	// Whether there is a release change
	IsChanged *bool `json:"IsChanged,omitnil,omitempty" name:"IsChanged"`

	// Publish information
	ReleaseSummary *ReleaseSummary `json:"ReleaseSummary,omitnil,omitempty" name:"ReleaseSummary"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLatestReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLatestReleaseResponseParams `json:"Response"`
}

func (r *DescribeLatestReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricOverviewListRequestParams struct {
	// <p>Dashboard domain. Required. Determines which domain's KPI data to return.</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>Platform feature usage</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering (resource_type=MODEL). Filter combination rule: multiple items use AND, AND value_list within the same item uses OR. Supported Names: MODEL_Name (MODEL Name), user_id (uid), space_id (space id), resource_id/source_id (source id), metric_source_type (metric_source_type_* Enumeration Name OR integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeMetricOverviewListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Dashboard domain. Required. Determines which domain's KPI data to return.</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>Platform feature usage</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended filtering (resource_type=MODEL). Filter combination rule: multiple items use AND, AND value_list within the same item uses OR. Supported Names: MODEL_Name (MODEL Name), user_id (uid), space_id (space id), resource_id/source_id (source id), metric_source_type (metric_source_type_* Enumeration Name OR integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeMetricOverviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricOverviewListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricOverviewListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricOverviewListResponseParams struct {
	// <p>All domain Overview unified output parameters: KPI Card List, key string identification indicator, client parses by resource_type; for key allowlist, see platform.common.v2.MetricOverview annotation</p>
	MetricList []*MetricOverview `json:"MetricList,omitnil,omitempty" name:"MetricList"`

	// <p>Total record count, equal to the length of MetricList, reserved only for list API consistency</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMetricOverviewListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMetricOverviewListResponseParams `json:"Response"`
}

func (r *DescribeMetricOverviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricOverviewListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListRequestParams struct {
	// <p>Model scenario. 0-case-insensitive, 1-standard generation, 2-standard thinking, 3-Agent thinking, 4-multimodal understanding, 5-multimodal QA, 6-rewrite, 7-long-term memory, 8-natural language to SQL, 9-AI optimization, 10-real-time file parsing, 11-file parsing, 12-GraphRAG, 13-OpenClaw, 14-multimodal Embedding, 15-Rerank, 16-text Embedding, 17-Widget, 18-Claw mode, 19-workflow code generation, 20-workflow large model node, 21-dedicated vectorization for workflow node, 22-workflow parameter extraction, 23-workflow large model knowledge q&a, 24-workflow tag extraction, 25-workflow intent recognition, 26-workflow tab, 27-workflow logical judgment, 28-documentation generation QA, 29-knowledge base Schema</p><p>Enumeration values:</p><ul><li>0: case-insensitive</li><li>1: standard generation</li><li>2: standard thinking</li><li>3: Agent thinking</li><li>4: multimodal understanding</li><li>5: multimodal QA</li><li>6: rewrite</li><li>7: long-term memory</li><li>8: natural language to SQL</li><li>9: AI optimization</li><li>10: real-time file parsing</li><li>11: file parsing</li><li>12: GraphRAG</li><li>13: OpenClaw</li><li>14: multimodal Embedding</li><li>15: Rerank</li><li>16: text Embedding</li><li>17: Widget</li><li>18: Claw mode</li><li>19: workflow code generation</li><li>20: workflow large model node</li><li>21: dedicated vectorization for workflow node</li><li>22: workflow parameter extraction</li><li>23: workflow large model knowledge q&a</li><li>24: workflow tag extraction</li><li>25: workflow intent recognition</li><li>26: workflow tab</li><li>27: workflow logical judgment</li><li>28: documentation generation QA</li><li>29: knowledge base Schema</li></ul>
	ModelScene *int64 `json:"ModelScene,omitnil,omitempty" name:"ModelScene"`

	// <p>Filter conditions (multiple Filters are combined with AND, and multiple Values within the same Filter are combined with OR)<br>DeveloperName: model author name<br>ProviderName: model provider name<br>ProviderType: model provider type</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number. Starts from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 20. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Fuzzy search by keyword</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeModelListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Model scenario. 0-case-insensitive, 1-standard generation, 2-standard thinking, 3-Agent thinking, 4-multimodal understanding, 5-multimodal QA, 6-rewrite, 7-long-term memory, 8-natural language to SQL, 9-AI optimization, 10-real-time file parsing, 11-file parsing, 12-GraphRAG, 13-OpenClaw, 14-multimodal Embedding, 15-Rerank, 16-text Embedding, 17-Widget, 18-Claw mode, 19-workflow code generation, 20-workflow large model node, 21-dedicated vectorization for workflow node, 22-workflow parameter extraction, 23-workflow large model knowledge q&a, 24-workflow tag extraction, 25-workflow intent recognition, 26-workflow tab, 27-workflow logical judgment, 28-documentation generation QA, 29-knowledge base Schema</p><p>Enumeration values:</p><ul><li>0: case-insensitive</li><li>1: standard generation</li><li>2: standard thinking</li><li>3: Agent thinking</li><li>4: multimodal understanding</li><li>5: multimodal QA</li><li>6: rewrite</li><li>7: long-term memory</li><li>8: natural language to SQL</li><li>9: AI optimization</li><li>10: real-time file parsing</li><li>11: file parsing</li><li>12: GraphRAG</li><li>13: OpenClaw</li><li>14: multimodal Embedding</li><li>15: Rerank</li><li>16: text Embedding</li><li>17: Widget</li><li>18: Claw mode</li><li>19: workflow code generation</li><li>20: workflow large model node</li><li>21: dedicated vectorization for workflow node</li><li>22: workflow parameter extraction</li><li>23: workflow large model knowledge q&a</li><li>24: workflow tag extraction</li><li>25: workflow intent recognition</li><li>26: workflow tab</li><li>27: workflow logical judgment</li><li>28: documentation generation QA</li><li>29: knowledge base Schema</li></ul>
	ModelScene *int64 `json:"ModelScene,omitnil,omitempty" name:"ModelScene"`

	// <p>Filter conditions (multiple Filters are combined with AND, and multiple Values within the same Filter are combined with OR)<br>DeveloperName: model author name<br>ProviderName: model provider name<br>ProviderType: model provider type</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number. Starts from 0</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default: 20. Maximum: 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Fuzzy search by keyword</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelScene")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListResponseParams struct {
	// <p>Model list</p>
	ModelList []*Model `json:"ModelList,omitnil,omitempty" name:"ModelList"`

	// <p>Total number of models</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelListResponseParams `json:"Response"`
}

func (r *DescribeModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordCategoryListRequestParams struct {
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeMsgRecordCategoryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeMsgRecordCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgRecordCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordCategoryListResponseParams struct {
	// <p>Message record category tree list</p>
	CategoryList []*MsgRecordCategory `json:"CategoryList,omitnil,omitempty" name:"CategoryList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgRecordCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgRecordCategoryListResponseParams `json:"Response"`
}

func (r *DescribeMsgRecordCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordListRequestParams struct {
	// <p>Application ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <ul><li><strong>ChannelType</strong> : 0-all 2-experience page (Tencent Cloud) 3-web debugging page (Tencent Cloud) 4-experience page (phone number) 5-dialogue API connection 6-application evaluation 7-debug API 10000-WeChat service account 10001-WeChat Subscription Account 10002-WeCom app 10004-WeChat Customer Service 10005-WeChat mini program 10006-Tencent Yuanqi 10007-App Treasure 10009-WeCom intelligent bot 10014-WeCom intelligent bot 10011-LINE 10012-Telegram 10013-DingTalk chatbot 10016-Lark chatbot 30000-scheduled task 30001-trigger </li><li><strong>FeedbackType</strong> : feedback type, -1-all 0-not rated 1-like 2-dislike </li><li><strong>QueryType</strong> : retrieval type (take value as agreed by the platform), SessionId-sessionID Question-question Answer-reply Intent-intent User-user </li><li><strong>Query</strong> : filter criteria corresponding to QueryType </li><li><strong>CategoryId</strong> : category ID </li><li><strong>ReplyMethod</strong> : reply type 0-all reply types 1-reply to rejected questions 2-direct QA reply 3-review failure reply 4-knowledge polish reply 6-workflow reply 8-image understanding reply 9-search engine reply 10-large model direct reply 11-fallback response </li><li><strong>StartTime</strong> : start time, unix second-level timestamp </li><li><strong>EndTime</strong> : end time, unix second-level timestamp </li><li><strong>Cursor</strong> : cursor. For the previous page, take PrevCursor from the response; for the next page, take NextCursor from the response. </li><li><strong>Direction</strong> : direction. next: next page, prev: previous page</li><li><strong>CallResult</strong> : call result: 0-all / 1-success / 2-failure / 3-canceled by the user </li><li><strong>FailReason</strong> : failure reason 0-all 1-review failure 2-QPM limit reached 3-TPM limit reached 4-concurrency limit reached 5-system exception </li><li><strong>Intent</strong> : intent</li></ul>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0. If not passed, it is processed as 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum value: 100. If not specified or set to 0, the default pagination size is used.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Sort condition. Supports sorting by CreateTime only</p>
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeMsgRecordListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <ul><li><strong>ChannelType</strong> : 0-all 2-experience page (Tencent Cloud) 3-web debugging page (Tencent Cloud) 4-experience page (phone number) 5-dialogue API connection 6-application evaluation 7-debug API 10000-WeChat service account 10001-WeChat Subscription Account 10002-WeCom app 10004-WeChat Customer Service 10005-WeChat mini program 10006-Tencent Yuanqi 10007-App Treasure 10009-WeCom intelligent bot 10014-WeCom intelligent bot 10011-LINE 10012-Telegram 10013-DingTalk chatbot 10016-Lark chatbot 30000-scheduled task 30001-trigger </li><li><strong>FeedbackType</strong> : feedback type, -1-all 0-not rated 1-like 2-dislike </li><li><strong>QueryType</strong> : retrieval type (take value as agreed by the platform), SessionId-sessionID Question-question Answer-reply Intent-intent User-user </li><li><strong>Query</strong> : filter criteria corresponding to QueryType </li><li><strong>CategoryId</strong> : category ID </li><li><strong>ReplyMethod</strong> : reply type 0-all reply types 1-reply to rejected questions 2-direct QA reply 3-review failure reply 4-knowledge polish reply 6-workflow reply 8-image understanding reply 9-search engine reply 10-large model direct reply 11-fallback response </li><li><strong>StartTime</strong> : start time, unix second-level timestamp </li><li><strong>EndTime</strong> : end time, unix second-level timestamp </li><li><strong>Cursor</strong> : cursor. For the previous page, take PrevCursor from the response; for the next page, take NextCursor from the response. </li><li><strong>Direction</strong> : direction. next: next page, prev: previous page</li><li><strong>CallResult</strong> : call result: 0-all / 1-success / 2-failure / 3-canceled by the user </li><li><strong>FailReason</strong> : failure reason 0-all 1-review failure 2-QPM limit reached 3-TPM limit reached 4-concurrency limit reached 5-system exception </li><li><strong>Intent</strong> : intent</li></ul>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0. If not passed, it is processed as 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum value: 100. If not specified or set to 0, the default pagination size is used.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Sort condition. Supports sorting by CreateTime only</p>
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeMsgRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordListResponseParams struct {
	// <p>Whether there are more pages</p>
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// <p>Message record list</p>
	MsgRecordList []*MsgRecord `json:"MsgRecordList,omitnil,omitempty" name:"MsgRecordList"`

	// <p>Next page cursor information.</p>
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// <p>Previous page cursor info.</p>
	PrevCursor *string `json:"PrevCursor,omitnil,omitempty" name:"PrevCursor"`

	// <p>Total record count that meets the conditions, used for frontend paginated display</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgRecordListResponseParams `json:"Response"`
}

func (r *DescribeMsgRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeBatchRunRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Index Number</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>Workflow Node ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>Associated node of the sub-workflow</p><p>node_id_a.node_id_b[1]</p>
	SubWorkflowNodePath *string `json:"SubWorkflowNodePath,omitnil,omitempty" name:"SubWorkflowNodePath"`

	// <p>Workflow running ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type DescribeNodeBatchRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Index Number</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>Workflow Node ID</p>
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// <p>Associated node of the sub-workflow</p><p>node_id_a.node_id_b[1]</p>
	SubWorkflowNodePath *string `json:"SubWorkflowNodePath,omitnil,omitempty" name:"SubWorkflowNodePath"`

	// <p>Workflow running ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

func (r *DescribeNodeBatchRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeBatchRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Index")
	delete(f, "NodeId")
	delete(f, "SubWorkflowNodePath")
	delete(f, "WorkflowRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeBatchRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeBatchRunResponseParams struct {
	// <p>Error code</p>
	FailCode *string `json:"FailCode,omitnil,omitempty" name:"FailCode"`

	// <p>Error description</p>
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// <p>Input parameter</p>
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// <p>Input parameter: Cos address</p>
	InputRef *string `json:"InputRef,omitnil,omitempty" name:"InputRef"`

	// <p>Output parameters</p>
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// <p>Output parameter: Cos address</p>
	OutputRef *string `json:"OutputRef,omitnil,omitempty" name:"OutputRef"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_UNSPECIFIED</td> <td>0</td> <td>undefined</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_INIT</td> <td>1</td> <td>initial status</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_RUNNING</td> <td>2</td> <td>running</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_SUCCESS</td> <td>3</td> <td>run successfully</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_FAILED</td> <td>4</td> <td>execution failed</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_CANCELED</td> <td>5</td> <td>canceled</td> </tr> <tr> <td>NODE_BATCH_RUN_STATUS_WAITING_INPUT</td> <td>6</td> <td>pending</td> </tr> </tbody> </table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeBatchRunResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeBatchRunResponseParams `json:"Response"`
}

func (r *DescribeNodeBatchRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeBatchRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRunDetailRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Node operation instance ID</p>
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`
}

type DescribeNodeRunDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Node operation instance ID</p>
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`
}

func (r *DescribeNodeRunDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRunDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "NodeRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeRunDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRunDetailResponseParams struct {
	// <p>Node operation instance details</p>
	NodeRunDetail *NodeRunDetail `json:"NodeRunDetail,omitnil,omitempty" name:"NodeRunDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeRunDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeRunDetailResponseParams `json:"Response"`
}

func (r *DescribeNodeRunDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRunDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginRequestParams struct {
	// <p>Plugin id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current space id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Get the specified field</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>Plug-in display scenario. If not passed or set to 0, the scenario is not limited.</p><p>Enumeration values:</p><ul><li>0: Not limited</li><li>1: Agent Mode</li><li>2: Workflow</li><li>3: Intelligent workbench</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`
}

type DescribePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plugin id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current space id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Get the specified field</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>Plug-in display scenario. If not passed or set to 0, the scenario is not limited.</p><p>Enumeration values:</p><ul><li>0: Not limited</li><li>1: Agent Mode</li><li>2: Workflow</li><li>3: Intelligent workbench</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`
}

func (r *DescribePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	delete(f, "FieldMask")
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginResponseParams struct {
	// <p>Plugin details</p>
	Plugin *Plugin `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginResponseParams `json:"Response"`
}

func (r *DescribePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginSummaryListRequestParams struct {
	// <p>Space ID, used when querying the plug-in list in the space</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Filter criteria list, supporting PluginKind, CategoryKey, PluginSource, PluginId, PluginClass, BillingType, AuthType, IsShared, IsCreatedByMe</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Whether to only return favorite plug-ins. If set to true, only plug-ins favorited by the current user are returned. If set to false or not specified, no filtering by favorite status is applied.</p>
	IsFavoriteOnly *bool `json:"IsFavoriteOnly,omitnil,omitempty" name:"IsFavoriteOnly"`

	// <p>Plug-in display scenario. If not passed or set to 0, the scenario is not limited.</p><p>Enumeration values:</p><ul><li>0: Not limited</li><li>1: Agent Mode</li><li>2: Workflow</li><li>3: Intelligent workbench</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>Page number. Starts from 0</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Size per page.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query content fuzzy matching: plugin name/plugin description/tool name/tool description</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Sorting method.</p><p>Enumeration values:</p><ul><li>0: not specified, default sorting</li><li>1: relevance ranking</li><li>2: sort by update time</li><li>3: default sorting</li><li>4: sort by popularity</li></ul>
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>Filter plug-ins shared by the current space/business</p><p>Value range: [0, 2]</p>
	PluginSpaceRelation *int64 `json:"PluginSpaceRelation,omitnil,omitempty" name:"PluginSpaceRelation"`
}

type DescribePluginSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Space ID, used when querying the plug-in list in the space</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Filter criteria list, supporting PluginKind, CategoryKey, PluginSource, PluginId, PluginClass, BillingType, AuthType, IsShared, IsCreatedByMe</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Whether to only return favorite plug-ins. If set to true, only plug-ins favorited by the current user are returned. If set to false or not specified, no filtering by favorite status is applied.</p>
	IsFavoriteOnly *bool `json:"IsFavoriteOnly,omitnil,omitempty" name:"IsFavoriteOnly"`

	// <p>Plug-in display scenario. If not passed or set to 0, the scenario is not limited.</p><p>Enumeration values:</p><ul><li>0: Not limited</li><li>1: Agent Mode</li><li>2: Workflow</li><li>3: Intelligent workbench</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`

	// <p>Page number. Starts from 0</p>
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Size per page.</p>
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query content fuzzy matching: plugin name/plugin description/tool name/tool description</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Sorting method.</p><p>Enumeration values:</p><ul><li>0: not specified, default sorting</li><li>1: relevance ranking</li><li>2: sort by update time</li><li>3: default sorting</li><li>4: sort by popularity</li></ul>
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`

	// <p>Filter plug-ins shared by the current space/business</p><p>Value range: [0, 2]</p>
	PluginSpaceRelation *int64 `json:"PluginSpaceRelation,omitnil,omitempty" name:"PluginSpaceRelation"`
}

func (r *DescribePluginSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FilterList")
	delete(f, "IsFavoriteOnly")
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SortType")
	delete(f, "PluginSpaceRelation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginSummaryListResponseParams struct {
	// <p>plugin_list</p>
	PluginList []*PluginSummary `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// <p>total_count</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePluginSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginSummaryListResponseParams `json:"Response"`
}

func (r *DescribePluginSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQARequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`
}

type DescribeQARequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`
}

func (r *DescribeQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "QaId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQAResponseParams struct {
	// <p>Fragment highlight information.</p>
	HighlightList []*QASegmentHighlight `json:"HighlightList,omitnil,omitempty" name:"HighlightList"`

	// <p>Fragment content</p>
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// <p>Problem description</p>
	QuestionDescription *string `json:"QuestionDescription,omitnil,omitempty" name:"QuestionDescription"`

	// <p>Similar question list</p>
	SimilarQuestionList []*SimilarQuestion `json:"SimilarQuestionList,omitnil,omitempty" name:"SimilarQuestionList"`

	// <p>Basic information.</p>
	Summary *QASummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQAResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQAResponseParams `json:"Response"`
}

func (r *DescribeQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQASummaryListRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status-QA status, enumeration value, exact match; CategoryId-category ID, exact match; SourceType-QA source type, enumeration value, exact match; EffectiveDomain-application scope, exact match; DocId-associated document ID, exact match; CreateTime-creation time, Unix seconds, BETWEEN passes [start second, end second]; UpdateTime-update time, Unix seconds, BETWEEN passes [start second, end second]</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number for pagination, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default 10, maximum 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query condition (keyword + query scope)</p>
	Query *QAQuery `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

type DescribeQASummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter criteria (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status-QA status, enumeration value, exact match; CategoryId-category ID, exact match; SourceType-QA source type, enumeration value, exact match; EffectiveDomain-application scope, exact match; DocId-associated document ID, exact match; CreateTime-creation time, Unix seconds, BETWEEN passes [start second, end second]; UpdateTime-update time, Unix seconds, BETWEEN passes [start second, end second]</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number for pagination, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Default 10, maximum 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Query condition (keyword + query scope)</p>
	Query *QAQuery `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Switch configuration</p>
	SummaryListSwitch *SummaryListSwitch `json:"SummaryListSwitch,omitnil,omitempty" name:"SummaryListSwitch"`
}

func (r *DescribeQASummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQASummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SummaryListSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQASummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQASummaryListResponseParams struct {
	// <p>Not adopted count</p>
	NotAcceptedCount *uint64 `json:"NotAcceptedCount,omitnil,omitempty" name:"NotAcceptedCount"`

	// <p>Quantity to be checked</p>
	PendingVerifyCount *uint64 `json:"PendingVerifyCount,omitnil,omitempty" name:"PendingVerifyCount"`

	// <p>QA list</p>
	QaList []*QASummary `json:"QaList,omitnil,omitempty" name:"QaList"`

	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQASummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQASummaryListResponseParams `json:"Response"`
}

func (r *DescribeQASummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQASummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseListRequestParams struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeReleaseListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeReleaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseListResponseParams struct {
	// release_list
	ReleaseList []*ReleaseRecord `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`

	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseListResponseParams `json:"Response"`
}

func (r *DescribeReleaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseSummaryRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type DescribeReleaseSummaryRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

func (r *DescribeReleaseSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseSummaryResponseParams struct {
	// Release information
	ReleaseSummary *ReleaseSummary `json:"ReleaseSummary,omitnil,omitempty" name:"ReleaseSummary"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseSummaryResponseParams `json:"Response"`
}

func (r *DescribeReleaseSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceSummaryRequestParams struct {

}

type DescribeResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceSummaryResponseParams struct {
	// <p>Usage information of the billing package</p>
	ResourcePackage *ResourcePackageInfo `json:"ResourcePackage,omitnil,omitempty" name:"ResourcePackage"`

	// <p>Billing value-added package usage information</p>
	AddOnPackage *AddOnPackageInfo `json:"AddOnPackage,omitnil,omitempty" name:"AddOnPackage"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceSummaryResponseParams `json:"Response"`
}

func (r *DescribeResourceSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillCategoryListRequestParams struct {

}

type DescribeSkillCategoryListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSkillCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillCategoryListResponseParams struct {
	// Skill category list
	CategoryList []*SkillCategory `json:"CategoryList,omitnil,omitempty" name:"CategoryList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillCategoryListResponseParams `json:"Response"`
}

func (r *DescribeSkillCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillDetailRequestParams struct {
	// skillID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// Space ID.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// version Filter criteria (multiple filters are in an AND relationship with each other, AND multiple Values of the same Filter are in an OR relationship): - Perspective: Perspective enumeration, single string value. The length of Values must be 1. Multiple Values are deemed as illegal. It only acts on the trimming of version_list returned in the details AND does not determine the visibility of the api itself. Default if not passed: USER (USER = USER Perspective, version_list returns only launch versions / EDITOR = EDITOR Perspective, version_list returns ALL liveness versions / ALL = full Perspective, same as EDITOR)
	VersionFilterList []*Filter `json:"VersionFilterList,omitnil,omitempty" name:"VersionFilterList"`
}

type DescribeSkillDetailRequest struct {
	*tchttp.BaseRequest
	
	// skillID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// Space ID.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// version Filter criteria (multiple filters are in an AND relationship with each other, AND multiple Values of the same Filter are in an OR relationship): - Perspective: Perspective enumeration, single string value. The length of Values must be 1. Multiple Values are deemed as illegal. It only acts on the trimming of version_list returned in the details AND does not determine the visibility of the api itself. Default if not passed: USER (USER = USER Perspective, version_list returns only launch versions / EDITOR = EDITOR Perspective, version_list returns ALL liveness versions / ALL = full Perspective, same as EDITOR)
	VersionFilterList []*Filter `json:"VersionFilterList,omitnil,omitempty" name:"VersionFilterList"`
}

func (r *DescribeSkillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionFilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillDetailResponseParams struct {
	// skill details
	SkillDetail *SkillDetail `json:"SkillDetail,omitnil,omitempty" name:"SkillDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillDetailResponseParams `json:"Response"`
}

func (r *DescribeSkillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillReferenceListRequestParams struct {
	// <p>Skill ID, required</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID. Required.</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeSkillReferenceListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill ID, required</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID. Required.</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeSkillReferenceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillReferenceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillReferenceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillReferenceListResponseParams struct {
	// <p>Reference summary grouped by SkillRefType: when total_count of a type is 0, it is not included in a group (no empty placeholder is returned). This period also writes to three paths: OPENCLAW / AGENT / CORP_ASSISTANT</p>
	ReferenceList []*SkillReferenceGroup `json:"ReferenceList,omitnil,omitempty" name:"ReferenceList"`

	// <p>Whether the current user is allowed to forcibly delete a referenced Skill</p>
	AllowForceModify *bool `json:"AllowForceModify,omitnil,omitempty" name:"AllowForceModify"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillReferenceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillReferenceListResponseParams `json:"Response"`
}

func (r *DescribeSkillReferenceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillReferenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillSummaryListRequestParams struct {
	// Space ID. Required.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Query only the Skills favorited by the current user
	FavoriteOnly *bool `json:"FavoriteOnly,omitnil,omitempty" name:"FavoriteOnly"`

	// Filter criteria (multiple Filters are in an AND relationship, while multiple Values within the same Filter are in an OR relationship):
	// -SkillIdList: List of skill IDs, string array, exact match
	// - ProviderType: Skill provider type, enumerated value array, exact match
	//      (SKILL_PROVIDER_TYPE_OFFICIAL=1/SKILL_PROVIDER_TYPE_THIRD_PARTY=2/SKILL_PROVIDER_TYPE_CUSTOM=3/SKILL_PROVIDER_TYPE_CUSTOM_SHARED=4)
	// -CategoryKey: Category identifier, string array, exact match.
	// - AnalysisStatus: Security detection status, enumeration value array, exact match.
	//      (SKILL_ANALYSIS_PENDING=0/SKILL_ANALYSIS_RUNNING=1/SKILL_ANALYSIS_AVAILABLE=2/SKILL_ANALYSIS_UNAVAILABLE=3/SKILL_ANALYSIS_FAILED=4)
	// - RiskLevel: Risk level, enumerated value array, exact match
	//      (SKILL_RISK_NONE=0/SKILL_RISK_LOW=1/SKILL_RISK_MEDIUM=2/SKILL_RISK_HIGH=3)
	// - SkillStatus: Publishing status in the Skill dimension. Enumerated value array. Exact match. OR between multiple values. Only meaningful when Perspective is EDITOR or ALL.
	// (SKILL_STATUS_INITIALIZED=0/SKILL_STATUS_AUDITING=1/SKILL_STATUS_PENDING_RELEASE=2/SKILL_STATUS_RELEASED=3)
	// - ShareStatus: Shared status, enumeration value array, exact match. This parameter is valid only when ProviderType includes SKILL_PROVIDER_TYPE_CUSTOM/SKILL_PROVIDER_TYPE_CUSTOM_SHARED.
	//      (SHARE_STATUS_UNSHARED=0/SHARE_STATUS_SHARED=1/SHARE_STATUS_APPROVING=2)
	// -Perspective: Perspective enum, string single value. The length of Values must be 1. Multi-value is deemed as illegal. This parameter is valid only when ProviderType=SKILL_PROVIDER_TYPE_CUSTOM. Default if not passed: USER
	// (USER=User perspective. Only Skills with usage permissions are returned / EDITOR=Editor perspective. Only Skills with edit permission are returned / ALL=Full perspective. Skills with any permission bit are returned)
	// -Creator: Filter by creator. String single value. The length of Values must be 1. Multi-value is deemed as illegal. This parameter is valid only when ProviderType=SKILL_PROVIDER_TYPE_CUSTOM.
	// Currently, only the placeholder "$self" is supported, which means only Skills created by the current caller are returned.
	// To scale to a specified identity subsequently, add the agreement here if needed.
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Page number, starting from 0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum value: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy search by name/display name
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeSkillSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// Space ID. Required.
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Query only the Skills favorited by the current user
	FavoriteOnly *bool `json:"FavoriteOnly,omitnil,omitempty" name:"FavoriteOnly"`

	// Filter criteria (multiple Filters are in an AND relationship, while multiple Values within the same Filter are in an OR relationship):
	// -SkillIdList: List of skill IDs, string array, exact match
	// - ProviderType: Skill provider type, enumerated value array, exact match
	//      (SKILL_PROVIDER_TYPE_OFFICIAL=1/SKILL_PROVIDER_TYPE_THIRD_PARTY=2/SKILL_PROVIDER_TYPE_CUSTOM=3/SKILL_PROVIDER_TYPE_CUSTOM_SHARED=4)
	// -CategoryKey: Category identifier, string array, exact match.
	// - AnalysisStatus: Security detection status, enumeration value array, exact match.
	//      (SKILL_ANALYSIS_PENDING=0/SKILL_ANALYSIS_RUNNING=1/SKILL_ANALYSIS_AVAILABLE=2/SKILL_ANALYSIS_UNAVAILABLE=3/SKILL_ANALYSIS_FAILED=4)
	// - RiskLevel: Risk level, enumerated value array, exact match
	//      (SKILL_RISK_NONE=0/SKILL_RISK_LOW=1/SKILL_RISK_MEDIUM=2/SKILL_RISK_HIGH=3)
	// - SkillStatus: Publishing status in the Skill dimension. Enumerated value array. Exact match. OR between multiple values. Only meaningful when Perspective is EDITOR or ALL.
	// (SKILL_STATUS_INITIALIZED=0/SKILL_STATUS_AUDITING=1/SKILL_STATUS_PENDING_RELEASE=2/SKILL_STATUS_RELEASED=3)
	// - ShareStatus: Shared status, enumeration value array, exact match. This parameter is valid only when ProviderType includes SKILL_PROVIDER_TYPE_CUSTOM/SKILL_PROVIDER_TYPE_CUSTOM_SHARED.
	//      (SHARE_STATUS_UNSHARED=0/SHARE_STATUS_SHARED=1/SHARE_STATUS_APPROVING=2)
	// -Perspective: Perspective enum, string single value. The length of Values must be 1. Multi-value is deemed as illegal. This parameter is valid only when ProviderType=SKILL_PROVIDER_TYPE_CUSTOM. Default if not passed: USER
	// (USER=User perspective. Only Skills with usage permissions are returned / EDITOR=Editor perspective. Only Skills with edit permission are returned / ALL=Full perspective. Skills with any permission bit are returned)
	// -Creator: Filter by creator. String single value. The length of Values must be 1. Multi-value is deemed as illegal. This parameter is valid only when ProviderType=SKILL_PROVIDER_TYPE_CUSTOM.
	// Currently, only the placeholder "$self" is supported, which means only Skills created by the current caller are returned.
	// To scale to a specified identity subsequently, add the agreement here if needed.
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Page number, starting from 0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum value: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Fuzzy search by name/display name
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeSkillSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FavoriteOnly")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillSummaryListResponseParams struct {
	// Skill summary list
	SkillSummaryList []*SkillSummary `json:"SkillSummaryList,omitnil,omitempty" name:"SkillSummaryList"`

	// Total number.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillSummaryListResponseParams `json:"Response"`
}

func (r *DescribeSkillSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceListRequestParams struct {
	// Support fuzzy search by space name.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeSpaceListRequest struct {
	*tchttp.BaseRequest
	
	// Support fuzzy search by space name.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeSpaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceListResponseParams struct {
	// Total number.
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Space List
	SpaceList []*Space `json:"SpaceList,omitnil,omitempty" name:"SpaceList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceListResponseParams `json:"Response"`
}

func (r *DescribeSpaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVariableListRequestParams struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeSystemVariableListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeSystemVariableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVariableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemVariableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVariableListResponseParams struct {
	// system_variable_list
	SystemVariableList []*SystemVariable `json:"SystemVariableList,omitnil,omitempty" name:"SystemVariableList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemVariableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemVariableListResponseParams `json:"Response"`
}

func (r *DescribeSystemVariableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVariableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageDetailListRequestParams struct {
	// <p>Resource type, limited to RESOURCE_TYPE_MODEL / RESOURCE_TYPE_PLUGIN</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>Platform feature usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details) (reserved, not currently supported)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended Filter (resource_type=MODEL). Filter combination rule: multiple items use AND, AND values within the same item use OR. Supported names: MODEL_Name, user_id, space_id, resource_id/source_id, metric_source_type (metric_source_type_* OR integer), call_type (invocation type)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeUsageDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Resource type, limited to RESOURCE_TYPE_MODEL / RESOURCE_TYPE_PLUGIN</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>Platform feature usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details) (reserved, not currently supported)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended Filter (resource_type=MODEL). Filter combination rule: multiple items use AND, AND values within the same item use OR. Supported names: MODEL_Name, user_id, space_id, resource_id/source_id, metric_source_type (metric_source_type_* OR integer), call_type (invocation type)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeUsageDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageDetailListResponseParams struct {
	// <p>Total record count for frontend pagination</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Resource invocation time series details list</p>
	UsageDetailList []*UsageDetail `json:"UsageDetailList,omitnil,omitempty" name:"UsageDetailList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageDetailListResponseParams `json:"Response"`
}

func (r *DescribeUsageDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryListRequestParams struct {
	// <p>Resource type, limited to MODEL / PLUGIN / PLATFORM</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_PLATFORM (reserved, not currently supported)</td><td>3</td><td>Platform feature usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details) (reserved, not currently supported)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended Filter (resource_type=MODEL). Filter combination rule: multiple items use AND, AND value_list within the same item uses OR. Supported Name: MODEL_Name (MODEL Name), user_id (uid), space_id (space id), resource_id/source_id (source id), metric_source_type (metric_source_type_* Enumeration Name OR integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeUsageSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Resource type, limited to MODEL / PLUGIN / PLATFORM</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>Model usage</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>Plug-in usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_PLATFORM (reserved, not currently supported)</td><td>3</td><td>Platform feature usage (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>Model concurrency limit exceeded (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>Knowledge base capacity (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>Usage summary (reserved, not currently supported)</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>Resource consumption (billing details) (reserved, not currently supported)</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>Query time range (Unix seconds).</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>View scope: corporate view / space view / application view</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>Extended Filter (resource_type=MODEL). Filter combination rule: multiple items use AND, AND value_list within the same item uses OR. Supported Name: MODEL_Name (MODEL Name), user_id (uid), space_id (space id), resource_id/source_id (source id), metric_source_type (metric_source_type_* Enumeration Name OR integer)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Page number, starting from 0.</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Number of items per page. Maximum: 100.</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeUsageSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryListResponseParams struct {
	// <p>Total record count, used for frontend pagination</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Resource usage aggregation details list</p>
	UsageSummaryList []*UsageSummary `json:"UsageSummaryList,omitnil,omitempty" name:"UsageSummaryList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageSummaryListResponseParams `json:"Response"`
}

func (r *DescribeUsageSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableListRequestParams struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Filter criteria (supported: VariableIdList - variable ID list, VariableType - variable type)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Module type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters.
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// Whether internal variables are required
	NeedInternalVariable *bool `json:"NeedInternalVariable,omitnil,omitempty" name:"NeedInternalVariable"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query keyword.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeVariableListRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Filter criteria (supported: VariableIdList - variable ID list, VariableType - variable type)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// Module type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters.
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// Whether internal variables are required
	NeedInternalVariable *bool `json:"NeedInternalVariable,omitnil,omitempty" name:"NeedInternalVariable"`

	// Page number, starting from 0.
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Number of items per page. Maximum: 100.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Query keyword.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeVariableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "ModuleType")
	delete(f, "NeedInternalVariable")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVariableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableListResponseParams struct {
	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// variable_list
	VariableList []*Variable `json:"VariableList,omitnil,omitempty" name:"VariableList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVariableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVariableListResponseParams `json:"Response"`
}

func (r *DescribeVariableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

type DescribeVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type. Enumeration values: 1: environment parameter, 2: application parameter, 3: system parameter, -1: ALL parameters
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

func (r *DescribeVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "VariableId")
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableResponseParams struct {
	// Variable information
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVariableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVariableResponseParams `json:"Response"`
}

func (r *DescribeVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunRequestParams struct {
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow operation instance ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>The NodePath of the specified sub-Workflow. Format: [&lt;node_id&gt;[&lt;index&gt;].]*&lt;node_id&gt;[&lt;index&gt;]  This path is used to locate a specific Workflow instance (Workflow Run), which can be the main Workflow instance or a sub-instance generated by a sub-Workflow node. The path consists of node identifiers separated by dots (.). Each node identifier is in the format of node id[instance index]. node id (node_id): The id of the node to which the sub-Workflow belongs. Instance index (index): The "instance index" is empty for Workflow nodes, and non-empty for loop and batch processing nodes, starting from 1. Example: "" or unset -> Query the main Workflow instance (parent Workflow) "node_id_a" -> Query the sub-Workflow instance corresponding to Workflow node node_id_a in the main Workflow "node_id_a.node_id_b[1]" -> Query the Workflow execution status of the first loop of loop node node_id_b in the sub-Workflow corresponding to Workflow node node_id_a</p>
	SubWorkflowNodePath *string `json:"SubWorkflowNodePath,omitnil,omitempty" name:"SubWorkflowNodePath"`

	// <p>Whether to return the process diagram configuration of the workflow</p>
	IncludeWorkflowGraph *bool `json:"IncludeWorkflowGraph,omitnil,omitempty" name:"IncludeWorkflowGraph"`
}

type DescribeWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow operation instance ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>The NodePath of the specified sub-Workflow. Format: [&lt;node_id&gt;[&lt;index&gt;].]*&lt;node_id&gt;[&lt;index&gt;]  This path is used to locate a specific Workflow instance (Workflow Run), which can be the main Workflow instance or a sub-instance generated by a sub-Workflow node. The path consists of node identifiers separated by dots (.). Each node identifier is in the format of node id[instance index]. node id (node_id): The id of the node to which the sub-Workflow belongs. Instance index (index): The "instance index" is empty for Workflow nodes, and non-empty for loop and batch processing nodes, starting from 1. Example: "" or unset -> Query the main Workflow instance (parent Workflow) "node_id_a" -> Query the sub-Workflow instance corresponding to Workflow node node_id_a in the main Workflow "node_id_a.node_id_b[1]" -> Query the Workflow execution status of the first loop of loop node node_id_b in the sub-Workflow corresponding to Workflow node node_id_a</p>
	SubWorkflowNodePath *string `json:"SubWorkflowNodePath,omitnil,omitempty" name:"SubWorkflowNodePath"`

	// <p>Whether to return the process diagram configuration of the workflow</p>
	IncludeWorkflowGraph *bool `json:"IncludeWorkflowGraph,omitnil,omitempty" name:"IncludeWorkflowGraph"`
}

func (r *DescribeWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "WorkflowRunId")
	delete(f, "SubWorkflowNodePath")
	delete(f, "IncludeWorkflowGraph")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunResponseParams struct {
	// <p>Workflow running instance details.</p>
	WorkflowRun *WorkflowRun `json:"WorkflowRun,omitnil,omitempty" name:"WorkflowRun"`

	// <p>Node list</p>
	NodeRunList []*NodeRun `json:"NodeRunList,omitnil,omitempty" name:"NodeRunList"`

	// <p>NodePath of the sub-workflow</p>
	SubWorkflowNodePath *string `json:"SubWorkflowNodePath,omitnil,omitempty" name:"SubWorkflowNodePath"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowRunResponseParams `json:"Response"`
}

func (r *DescribeWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunSummaryListRequestParams struct {
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Number of items per page (value ranges from 1 to 200)</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Page number (must be greater than 0)</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Execution environment. 1: Development domain, 2: Production domain -</p>
	Domain *uint64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Parameter filtering</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeWorkflowRunSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID. For how to obtain it, see <a href="https://www.tencentcloud.com/document/product/1759/109469?from_cn_redirect=1#4eecb8c1-6ce4-45f5-8fa2-b269449d8efa">AppId</a></p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Number of items per page (value ranges from 1 to 200)</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>Page number (must be greater than 0)</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>Execution environment. 1: Development domain, 2: Production domain -</p>
	Domain *uint64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Parameter filtering</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeWorkflowRunSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Domain")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowRunSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunSummaryListResponseParams struct {
	// <p>Total.</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Workflow operation instance list</p>
	WorkflowRunSummaryList []*WorkflowRunSummary `json:"WorkflowRunSummaryList,omitnil,omitempty" name:"WorkflowRunSummaryList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowRunSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowRunSummaryListResponseParams `json:"Response"`
}

func (r *DescribeWorkflowRunSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DialogCustomConfig struct {
	// <p>Whether to enable dynamic configuration modification on the dialogue side</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DigitalHumanConfig struct {
	// Digital human avatar asset id
	AssetKey *string `json:"AssetKey,omitnil,omitempty" name:"AssetKey"`

	// Digital human image
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// Digital human avatar name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Preview address of the digital human
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`
}

type Distribution struct {
	// <p>Count of calls for this reply method</p>
	CallCount *string `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>Percentage of this reply method in the total number of calls (percentage, 0–100, rounded to two decimal places; all 0 when no data is available, and the sum of all items is 100 when there is data)</p>
	Percentage *float64 `json:"Percentage,omitnil,omitempty" name:"Percentage"`

	// <p>Reply method name (internationalized according to the request language; falls back to the Enumeration Name of reply_method when i18n is missing)</p>
	ReplyName *string `json:"ReplyName,omitnil,omitempty" name:"ReplyName"`
}

type DocExternalLink struct {
	// <p>External URL</p>
	ExternalUrl *string `json:"ExternalUrl,omitnil,omitempty" name:"ExternalUrl"`

	// <p>Whether to replace the original display</p>
	ReplaceOriginEnabled *bool `json:"ReplaceOriginEnabled,omitnil,omitempty" name:"ReplaceOriginEnabled"`
}

type DocImportSpec struct {
	// <p>File ID (file identifier obtained through the file management service, required)</p>
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// <p>Ownership category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>List of rules regarding repeated files</p>
	DeDuplicateStrategyList []*DeDuplicateStrategy `json:"DeDuplicateStrategyList,omitnil,omitempty" name:"DeDuplicateStrategyList"`

	// <p>Knowledge effective scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>External link</p>
	ExternalLink *DocExternalLink `json:"ExternalLink,omitnil,omitempty" name:"ExternalLink"`

	// <p>Applicable scope (tag condition).</p>
	LabelRefList []*LabelRefIdentity `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Parsing configuration (split rules, content filtering, etc.)</p>
	ParseConfig *DocParseConfig `json:"ParseConfig,omitnil,omitempty" name:"ParseConfig"`

	// <p>Switch configuration</p>
	Switch *DocSwitch `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>Update cycle</p>
	UpdatePeriod *DocUpdatePeriod `json:"UpdatePeriod,omitnil,omitempty" name:"UpdatePeriod"`

	// <p>User access configuration</p>
	UserAccessConfig *UserAccessConfig `json:"UserAccessConfig,omitnil,omitempty" name:"UserAccessConfig"`
}

type DocLifecycle struct {
	// <p>Creation time (Unix seconds).</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Document status: 1=parsing, 2=parsing failure, 3=import failed, 4=under review, 5=review failure, 6=Learning, 7=learning failed, 8=import complete, 9=expired, 10=quota invalid, 11=quota recovering, 12=rename review failure, 13=rename appeal failure, 14=manual appeal in progress, 15=manual appeal failure<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DOC_STATUS_UNKNOWN</td><td>0</td><td></td></tr><tr><td>DOC_STATUS_PARSING</td><td>1</td><td>Parsing</td></tr><tr><td>DOC_STATUS_PARSE_FAIL</td><td>2</td><td>Parsing failure</td></tr><tr><td>DOC_STATUS_IMPORT_FAIL</td><td>3</td><td>Import failed</td></tr><tr><td>DOC_STATUS_AUDITING</td><td>4</td><td>Under review</td></tr><tr><td>DOC_STATUS_AUDIT_FAIL</td><td>5</td><td>Review failure</td></tr><tr><td>DOC_STATUS_LEARNING</td><td>6</td><td>Learning</td></tr><tr><td>DOC_STATUS_LEARN_FAIL</td><td>7</td><td>Learning failed</td></tr><tr><td>DOC_STATUS_IMPORTED</td><td>8</td><td>Import complete</td></tr><tr><td>DOC_STATUS_EXPIRED</td><td>9</td><td>Expired</td></tr><tr><td>DOC_STATUS_QUOTA_INVALID</td><td>10</td><td>Quota invalid</td></tr><tr><td>DOC_STATUS_QUOTA_RECOVERING</td><td>11</td><td>Quota recovering</td></tr><tr><td>DOC_STATUS_RENAME_AUDIT_FAIL</td><td>12</td><td>Rename review failure</td></tr><tr><td>DOC_STATUS_RENAME_APPEAL_FAIL</td><td>13</td><td>Rename appeal failure</td></tr><tr><td>DOC_STATUS_MANUAL_APPEALING</td><td>14</td><td>Manual appeal in progress</td></tr><tr><td>DOC_STATUS_MANUAL_APPEAL_FAIL</td><td>15</td><td>Manual appeal failure</td></tr></tbody></table></p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Status additional information</p>
	StatusMessage *string `json:"StatusMessage,omitnil,omitempty" name:"StatusMessage"`

	// <p>Update time (Unix seconds)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DocLink struct {
	// <p>COS link address, which can be used for preview and download</p>
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// <p>External link</p>
	ExternalLink *DocExternalLink `json:"ExternalLink,omitnil,omitempty" name:"ExternalLink"`
}

type DocMetadata struct {
	// <p>Number of characters in the document</p>
	DocCharCount *string `json:"DocCharCount,omitnil,omitempty" name:"DocCharCount"`

	// <p>File name.</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>File size (byte)</p>
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// <p>File type/extension</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>List of metadata reference field names (used to display which categories and properties of a document are set as metadata)</p>
	RefFieldNameList []*string `json:"RefFieldNameList,omitnil,omitempty" name:"RefFieldNameList"`

	// <p>Source description.</p>
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// <p>Document source type: 1=local upload, 2=webpage link, 3=COS integration, 4=import from external sources<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DOC_SOURCE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>DOC_SOURCE_TYPE_LOCAL</td><td>1</td><td>Local upload</td></tr><tr><td>DOC_SOURCE_TYPE_URL</td><td>2</td><td>Webpage link</td></tr><tr><td>DOC_SOURCE_TYPE_COS</td><td>3</td><td>COS integration</td></tr><tr><td>DOC_SOURCE_TYPE_IMPORT</td><td>4</td><td>Import from external sources</td></tr></tbody></table></p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type DocModifyFields struct {
	// <p>Ownership category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Application scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>External link</p>
	ExternalLink *DocExternalLink `json:"ExternalLink,omitnil,omitempty" name:"ExternalLink"`

	// <p>Tag list.</p>
	LabelRefList []*LabelRefIdentity `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Document name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Parsing configuration (split rules, content filtering, etc.)</p>
	ParseConfig *DocParseConfig `json:"ParseConfig,omitnil,omitempty" name:"ParseConfig"`

	// <p>Switch configuration</p>
	Switch *DocSwitch `json:"Switch,omitnil,omitempty" name:"Switch"`

	// <p>Update cycle</p>
	UpdatePeriod *DocUpdatePeriod `json:"UpdatePeriod,omitnil,omitempty" name:"UpdatePeriod"`

	// <p>User access configuration</p>
	UserAccessConfig *UserAccessConfig `json:"UserAccessConfig,omitnil,omitempty" name:"UserAccessConfig"`
}

type DocOperator struct {
	// <p>Modifier</p>
	Modifier *Operator `json:"Modifier,omitnil,omitempty" name:"Modifier"`

	// <p>Operation permission</p>
	Permission *DocPermission `json:"Permission,omitnil,omitempty" name:"Permission"`
}

type DocParseConfig struct {
	// <p>Content filtering configuration</p>
	ContentFilter *ContentFilter `json:"ContentFilter,omitnil,omitempty" name:"ContentFilter"`

	// <p>Split rule</p>
	SplitRule *string `json:"SplitRule,omitnil,omitempty" name:"SplitRule"`
}

type DocPermission struct {
	// <p>Whether it is deletable.</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>Whether it is editable</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// <p>Whether it can be regenerated</p>
	CanRestart *bool `json:"CanRestart,omitnil,omitempty" name:"CanRestart"`

	// <p>Whether retry is allowed</p>
	CanRetry *bool `json:"CanRetry,omitnil,omitempty" name:"CanRetry"`
}

type DocQuery struct {
	// <p>Query keyword (name fuzzy search)</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Query scope (fields that the query applies to): 1=file name, 2=tag or tag value; multiple selections supported, invalid by default</p>
	QueryScopeList []*int64 `json:"QueryScopeList,omitnil,omitempty" name:"QueryScopeList"`
}

type DocRetrievalConfig struct {
	// <p>Confidence threshold</p>
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Return the top N entries</p>
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type DocSummary struct {
	// <p>Category path (including category ID, category ID path from the root node, and category name path)</p>
	CategoryPath *CategoryPath `json:"CategoryPath,omitnil,omitempty" name:"CategoryPath"`

	// <p>Document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Knowledge effective scope (aggregate application scope + tag condition)</p>
	KnowledgeScope *KnowledgeScope `json:"KnowledgeScope,omitnil,omitempty" name:"KnowledgeScope"`

	// <p>Lifecycle information.</p>
	Lifecycle *DocLifecycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// <p>Meta information.</p>
	Metadata *DocMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>Operator information (aggregate modifier + operation permission)</p>
	OperatorInfo *DocOperator `json:"OperatorInfo,omitnil,omitempty" name:"OperatorInfo"`

	// <p>Task status information</p>
	TaskStatus *DocTaskStatus `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type DocSwitch struct {
	// <p>Whether it is downloadable</p>
	DownloadEnabled *bool `json:"DownloadEnabled,omitnil,omitempty" name:"DownloadEnabled"`

	// <p>Whether to show in the reference</p>
	ReferEnabled *bool `json:"ReferEnabled,omitnil,omitempty" name:"ReferEnabled"`
}

type DocTaskStatus struct {
	// <p>List of completed historical task types</p>
	CompletedTaskTypeList []*int64 `json:"CompletedTaskTypeList,omitnil,omitempty" name:"CompletedTaskTypeList"`

	// <p>List of supported task types in progress</p>
	OngoingTaskTypeList []*int64 `json:"OngoingTaskTypeList,omitnil,omitempty" name:"OngoingTaskTypeList"`
}

type DocUpdatePeriod struct {
	// <p>Whether periodic update is enabled</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Update interval (hr)</p>
	PeriodHour *uint64 `json:"PeriodHour,omitnil,omitempty" name:"PeriodHour"`
}

type DuplexBilling struct {
	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>UNKNOW</td> <td>0</td> <td></td> </tr> <tr> <td>TOKEN</td> <td>1</td> <td>By token</td> </tr> <tr> <td>PAGE_COUNT</td> <td>2</td> <td>By page number</td> </tr> <tr> <td>TIMES</td> <td>3</td> <td>By count</td> </tr> <tr> <td>TIMES_THOUSAND</td> <td>4</td> <td>By thousand counts</td> </tr> <tr> <td>SECOND</td> <td>5</td> <td>By duration</td> </tr> <tr> <td>CHARACTER</td> <td>6</td> <td>By number of characters</td> </tr> <tr> <td>CHARACTER_THOUSAND</td> <td>7</td> <td>By thousand characters</td> </tr> <tr> <td>SHEET</td> <td>8</td> <td>By sheet</td> </tr> <tr> <td>NUMBER</td> <td>9</td> <td>By count</td> </tr> </tbody> </table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>Input cash price</p><p>Unit: CNY</p>
	InputCashPrice *float64 `json:"InputCashPrice,omitnil,omitempty" name:"InputCashPrice"`

	// <p>Input PU price</p><p>Unit: pu</p>
	InputPuPrice *float64 `json:"InputPuPrice,omitnil,omitempty" name:"InputPuPrice"`

	// <p>Output cash price</p><p>Unit: CNY</p>
	OutputCashPrice *float64 `json:"OutputCashPrice,omitnil,omitempty" name:"OutputCashPrice"`

	// <p>Output PU price</p><p>Unit: pu</p>
	OutputPuPrice *float64 `json:"OutputPuPrice,omitnil,omitempty" name:"OutputPuPrice"`
}

type ESConfig struct {
	// <p>Whether storage method modification is supported</p>
	CanModify *bool `json:"CanModify,omitnil,omitempty" name:"CanModify"`

	// <p>ES password (encrypted)</p>
	EncryptedPassword *string `json:"EncryptedPassword,omitnil,omitempty" name:"EncryptedPassword"`

	// <p>ES cluster ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>Storage type: 1=default storage, 2=custom storage<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>ES_STORAGE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>ES_STORAGE_TYPE_DEFAULT</td><td>1</td><td>Default storage</td></tr><tr><td>ES_STORAGE_TYPE_CUSTOM</td><td>2</td><td>Custom storage</td></tr></tbody></table></p>
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>ES userName</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type EffectivePeriod struct {
	// <p>Validity end time (Unix seconds. 0 indicates permanent validity).</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Validity start time (Unix seconds)</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type ExecuteConfig struct {
	// <p>Prompt configuration</p>
	PromptConfig *AppTriggerPromptExecuteConfig `json:"PromptConfig,omitnil,omitempty" name:"PromptConfig"`

	// <p>Workflow configuration.</p>
	WorkflowConfig *AppTriggerWorkflowExecuteConfig `json:"WorkflowConfig,omitnil,omitempty" name:"WorkflowConfig"`
}

type ExpirationAwareness struct {
	// <p>Whether to enable timeliness retrieval. Once enabled, retrieval results will be sorted based on the validity time of the knowledge</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ExpirationPolicy struct {
	// <p>Validity time</p>
	EffectivePeriod *EffectivePeriod `json:"EffectivePeriod,omitnil,omitempty" name:"EffectivePeriod"`

	// <p>Behavior after the validity period expires: 1=NOT_RETRIEVABLE cannot be retrieved, 2=RETRIEVABLE can still be retrieved; meaningless for permanent validity<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>EXPIRE_BEHAVIOR_UNKNOWN</td><td>0</td><td>Unspecified (server uses the default process)</td></tr><tr><td>EXPIRE_BEHAVIOR_NOT_RETRIEVABLE</td><td>1</td><td>Cannot be retrieved (removed upon expiration)</td></tr><tr><td>EXPIRE_BEHAVIOR_RETRIEVABLE</td><td>2</td><td>Can still be retrieved (not removed upon expiration, only the validity period is tagged)</td></tr></tbody></table></p>
	ExpireBehavior *int64 `json:"ExpireBehavior,omitnil,omitempty" name:"ExpireBehavior"`
}

// Predefined struct for user
type ExportQARequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter conditions (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status-QA status, enumeration value, exact match; CategoryId-category ID, exact match; SourceType-QA source type, enumeration value, exact match; EffectiveDomain-application scope, exact match; DocId-associated document ID, exact match; CreateTime-creation time, Unix seconds, BETWEEN passes [start second, end second]; UpdateTime-update time, Unix seconds, BETWEEN passes [start second, end second]; QaId-QA ID list, exact match, supports multiple values</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Query condition (keyword + query scope), consistent with DescribeQASummaryList</p>
	Query *QAQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

type ExportQARequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Filter conditions (multiple Filters are in an AND relationship, and multiple Values of the same Filter are in an OR relationship): Status-QA status, enumeration value, exact match; CategoryId-category ID, exact match; SourceType-QA source type, enumeration value, exact match; EffectiveDomain-application scope, exact match; DocId-associated document ID, exact match; CreateTime-creation time, Unix seconds, BETWEEN passes [start second, end second]; UpdateTime-update time, Unix seconds, BETWEEN passes [start second, end second]; QaId-QA ID list, exact match, supports multiple values</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>Query condition (keyword + query scope), consistent with DescribeQASummaryList</p>
	Query *QAQuery `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ExportQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "FilterList")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAResponseParams struct {
	// <p>Exported task ID (query the completion status via DescribeAsyncTaskStatus)</p>
	ExportTaskId *string `json:"ExportTaskId,omitnil,omitempty" name:"ExportTaskId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportQAResponse struct {
	*tchttp.BaseResponse
	Response *ExportQAResponseParams `json:"Response"`
}

func (r *ExportQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoritePluginRequestParams struct {
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type FavoritePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *FavoritePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoritePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FavoritePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoritePluginResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FavoritePluginResponse struct {
	*tchttp.BaseResponse
	Response *FavoritePluginResponseParams `json:"Response"`
}

func (r *FavoritePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoritePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoriteSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type FavoriteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *FavoriteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoriteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FavoriteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoriteSkillResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FavoriteSkillResponse struct {
	*tchttp.BaseResponse
	Response *FavoriteSkillResponseParams `json:"Response"`
}

func (r *FavoriteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoriteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldMask struct {
	// <p>Parameter name</p><p>Parameter format: path of the specified field to get</p>
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type FileParseModel struct {
	// Model alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Model description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Enhanced mode
	EnhancementMode *string `json:"EnhancementMode,omitnil,omitempty" name:"EnhancementMode"`

	// Unique ID of a model
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Model type
	ModelProviderType *string `json:"ModelProviderType,omitnil,omitempty" name:"ModelProviderType"`

	// Whether to enable formula enhancement
	EnableFormulaEnhancement *bool `json:"EnableFormulaEnhancement,omitnil,omitempty" name:"EnableFormulaEnhancement"`

	// Whether to enable LLM enhancement
	EnableLLMEnhancement *bool `json:"EnableLLMEnhancement,omitnil,omitempty" name:"EnableLLMEnhancement"`

	// Whether to output an HTML table
	OutputHtmlTable *bool `json:"OutputHtmlTable,omitnil,omitempty" name:"OutputHtmlTable"`

	// Supported file type list
	SupportedFileList []*SupportedFileType `json:"SupportedFileList,omitnil,omitempty" name:"SupportedFileList"`
}

type FileVolumeMount struct {
	// <p>Access mode: RO=read-only, RW=read-write; required. Only RW can be used when mounting to /workdir.</p>
	AccessMode *string `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// <p>Absolute mount directory in the sandbox; /workdir can be mounted at most once and only RW can be used.</p>
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// <p>Cloud disk identification</p>
	VolumeId *string `json:"VolumeId,omitnil,omitempty" name:"VolumeId"`

	// <p>Server-side generated mount flag. Reserved for modification, optional for new additions.</p>
	MountId *string `json:"MountId,omitnil,omitempty" name:"MountId"`

	// <p>Path relative to the root directory of cloud disk data; an empty string indicates mounting the entire data root directory.</p>
	RelativePath *string `json:"RelativePath,omitnil,omitempty" name:"RelativePath"`

	// <p>Cloud disk name. It is supplemented by the server only when querying the application configuration and ignored during modification. It is empty when it cannot be obtained and does not impact the mount configuration.</p>
	VolumeName *string `json:"VolumeName,omitnil,omitempty" name:"VolumeName"`
}

type Filter struct {
	// Filter field name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Operator, default IN (backward compatibility)<table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>FILTER_OPERATOR_IN</td><td>0</td><td>Belongs to value_list (default value, backward compatibility; value_list cannot be empty)</td></tr><tr><td>FILTER_OPERATOR_NOT_IN</td><td>1</td><td>Does not belong to value_list (value_list cannot be empty)</td></tr><tr><td>FILTER_OPERATOR_BETWEEN</td><td>2</td><td>Between (inclusive interval [start, end]; value_list must contain exactly 2 elements, one of which can be empty to indicate a unilateral open interval)</td></tr></table>
	Operator *int64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Filter value array
	ValueList []*string `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type FinalRerankConfig struct {
	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type GenerateModel struct {
	// Generative model
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type GraphRAG struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type Identity struct {
	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Digital ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// String ID
	StrId *string `json:"StrId,omitnil,omitempty" name:"StrId"`
}

// Predefined struct for user
type ImportDocListRequestParams struct {
	// <p>List of documents to be imported (quantity: 1-20)</p>
	DocList []*DocImportSpec `json:"DocList,omitnil,omitempty" name:"DocList"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

type ImportDocListRequest struct {
	*tchttp.BaseRequest
	
	// <p>List of documents to be imported (quantity: 1-20)</p>
	DocList []*DocImportSpec `json:"DocList,omitnil,omitempty" name:"DocList"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`
}

func (r *ImportDocListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDocListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocList")
	delete(f, "KbId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportDocListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportDocListResponseParams struct {
	// <p>Batch import result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportDocListResponse struct {
	*tchttp.BaseResponse
	Response *ImportDocListResponseParams `json:"Response"`
}

func (r *ImportDocListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportDocListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputBoxConfig struct {
	// Input box buttons: 1: Upload image, 2: Upload documents, 3: Tencent Docs, 4: Online search
	InputBoxButtons []*int64 `json:"InputBoxButtons,omitnil,omitempty" name:"InputBoxButtons"`
}

type IntentAchievementInfo struct {
	// Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type IntervalSchedule struct {
	// Start time.
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// Enumeration values:
	// | Enumeration value | uint |
	// | --- | --- |
	// | INTERVAL_UNIT_UNSPECIFIED | 0 |
	// | INTERVAL_UNIT_HOUR | 1 |
	// | INTERVAL_UNIT_DAY | 2 |
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Value
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type KBCapacity struct {
	// <p>Maximum number of characters</p>
	MaxCharSize *string `json:"MaxCharSize,omitnil,omitempty" name:"MaxCharSize"`

	// <p>Number of excess characters</p>
	OverCharSize *string `json:"OverCharSize,omitnil,omitempty" name:"OverCharSize"`

	// <p>Used number of characters</p>
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`
}

type KBCategory struct {
	// <p>Whether it can be added</p>
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// <p>Whether it is deletable</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>Whether it is editable</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Subcategory list.</p>
	ChildList []*KBCategory `json:"ChildList,omitnil,omitempty" name:"ChildList"`

	// <p>Is leaf node (no subcategory)</p>
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// <p>The number of category objects</p>
	ItemCount *uint64 `json:"ItemCount,omitnil,omitempty" name:"ItemCount"`

	// <p>Metadata configuration (configuration message when this category is set as metadata)</p>
	MetaValue *MetaValue `json:"MetaValue,omitnil,omitempty" name:"MetaValue"`

	// <p>Category name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type KBModelConfig struct {
	// <p>Embedding model</p>
	EmbeddingModel *string `json:"EmbeddingModel,omitnil,omitempty" name:"EmbeddingModel"`

	// <p>QA extraction model</p>
	QaExtractModel *string `json:"QaExtractModel,omitnil,omitempty" name:"QaExtractModel"`

	// <p>Schema generative model</p>
	SchemaModel *string `json:"SchemaModel,omitnil,omitempty" name:"SchemaModel"`
}

type KBModifyExtendFields struct {
	// <p>Scaling operation: 1=trigger recovery from exceeded state<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KB_EXTENDED_ACTION_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KB_EXTENDED_ACTION_RESUME_EXCEEDED</td><td>1</td><td>Trigger recovery from exceeded state (restore the knowledge base from exceeded status to normal status)</td></tr></tbody></table></p>
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`
}

type KBPermission struct {
	// <p>Whether it is deletable.</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>Whether it is editable</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// <p>Whether or not it can be shared</p>
	CanShare *bool `json:"CanShare,omitnil,omitempty" name:"CanShare"`

	// <p>Whether it is viewable</p>
	CanView *bool `json:"CanView,omitnil,omitempty" name:"CanView"`
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

type KBShareStatus struct {
	// <p>Approval Id.</p>
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// <p>Approval/shared state: 1=PENDING Pending Approval, 2=REJECTED rejected, 3=APPROVED approved, 4=CANCELLED canceled</p>
	SubStatusList []*int64 `json:"SubStatusList,omitnil,omitempty" name:"SubStatusList"`
}

type KBSpaceShare struct {
	// <p>Whether it has been referenced by other spaces</p>
	IsSpaceShared *bool `json:"IsSpaceShared,omitnil,omitempty" name:"IsSpaceShared"`

	// <p>Referenced spatial information.</p>
	SharedSpaceList []*Identity `json:"SharedSpaceList,omitnil,omitempty" name:"SharedSpaceList"`

	// <p>Shared state</p>
	SubStatus *KBShareStatus `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// <p>Whether operation is allowed</p>
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type KBSpec struct {
	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>ES configuration</p>
	EsConfig *ESConfig `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// <p>Model configuration</p>
	ModelConfig *KBModelConfig `json:"ModelConfig,omitnil,omitempty" name:"ModelConfig"`

	// <p>Knowledge base name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Owner ID</p>
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`
}

type KBSummary struct {
	// <p>List of associated applications. Returned only for shared knowledge base.</p>
	AppList []*Identity `json:"AppList,omitnil,omitempty" name:"AppList"`

	// <p>Creation time (Unix seconds).</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Creator</p>
	Creator *Operator `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Number of documents</p>
	DocCount *uint64 `json:"DocCount,omitnil,omitempty" name:"DocCount"`

	// <p>Whether it is excessive</p>
	IsExceeded *bool `json:"IsExceeded,omitnil,omitempty" name:"IsExceeded"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Type: 1=default knowledge base, 2=shared knowledge base<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KB_TYPE_DEFAULT</td><td>1</td><td>Default knowledge base</td></tr><tr><td>KB_TYPE_SHARED</td><td>2</td><td>Shared knowledge base</td></tr></tbody></table></p>
	KbType *int64 `json:"KbType,omitnil,omitempty" name:"KbType"`

	// <p>Last operator. Returned only for shared knowledge bases.</p>
	LatestOperator *Operator `json:"LatestOperator,omitnil,omitempty" name:"LatestOperator"`

	// <p>Knowledge base name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Operation permission.</p>
	Permission *KBPermission `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>List of in-processing statuses</p>
	ProcessingFlagList []*int64 `json:"ProcessingFlagList,omitnil,omitempty" name:"ProcessingFlagList"`

	// <p>Share sub-type: 1=ordinary, 2=wechat official account<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SHARED_KB_SUB_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>SHARED_KB_SUB_TYPE_NORMAL</td><td>1</td><td>Ordinary</td></tr><tr><td>SHARED_KB_SUB_TYPE_PUBLIC_ACCOUNT</td><td>2</td><td>WeChat official account</td></tr></tbody></table></p>
	SharedSubType *int64 `json:"SharedSubType,omitnil,omitempty" name:"SharedSubType"`

	// <p>Update time (Unix seconds)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Knowledge base space share information</p>
	SpaceShare *KBSpaceShare `json:"SpaceShare,omitnil,omitempty" name:"SpaceShare"`
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

type KnowledgeScope struct {
	// <p>Effective scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Applicable scope (tag condition).</p>
	LabelRefList []*LabelRef `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`
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

type LabelModifyFields struct {
	// <p>Tag name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>List of incremental modifications to standard words (add/modify/delete)</p>
	TermModifyList []*LabelTermModifyItem `json:"TermModifyList,omitnil,omitempty" name:"TermModifyList"`
}

type LabelRef struct {
	// <p>Tag ID.</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Tag name.</p>
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// <p>Tag standard term ID list.</p>
	LabelTermIdList []*string `json:"LabelTermIdList,omitnil,omitempty" name:"LabelTermIdList"`

	// <p>Tag standard word list</p>
	LabelTermList []*string `json:"LabelTermList,omitnil,omitempty" name:"LabelTermList"`
}

type LabelRefIdentity struct {
	// <p>Tag ID</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Tag standard term ID list</p>
	LabelTermIdList []*string `json:"LabelTermIdList,omitnil,omitempty" name:"LabelTermIdList"`
}

type LabelRefIdentityList struct {
	// <p>Tag references list.</p>
	ItemList []*LabelRefIdentity `json:"ItemList,omitnil,omitempty" name:"ItemList"`
}

type LabelSummary struct {
	// <p>Tag ID.</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Metadata configuration (configuration message when the tag is set to metadata)</p>
	MetaValue *MetaValue `json:"MetaValue,omitnil,omitempty" name:"MetaValue"`

	// <p>Tag name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Number of resources referring to the tag</p>
	RefCount *uint64 `json:"RefCount,omitnil,omitempty" name:"RefCount"`

	// <p>Tag value (standard word + synonym list)</p>
	TermList []*LabelTerm `json:"TermList,omitnil,omitempty" name:"TermList"`

	// <p>Total number of tag values.</p>
	TermTotalCount *uint64 `json:"TermTotalCount,omitnil,omitempty" name:"TermTotalCount"`
}

type LabelTerm struct {
	// <p>Synonym list</p>
	SynonymList []*string `json:"SynonymList,omitnil,omitempty" name:"SynonymList"`

	// <p>Standard words</p>
	Term *string `json:"Term,omitnil,omitempty" name:"Term"`

	// <p>Standard term ID (generated by the backend, not passed in upon creation)</p>
	TermId *string `json:"TermId,omitnil,omitempty" name:"TermId"`
}

type LabelTermCheckResult struct {
	// Check results
	CheckResult *CheckResult `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// Standard words to be checked
	Term *string `json:"Term,omitnil,omitempty" name:"Term"`

	// Return the corresponding standard term ID if exists.
	TermId *string `json:"TermId,omitnil,omitempty" name:"TermId"`
}

type LabelTermModifyItem struct {
	// <p>Operation type (not allowed to be 0. Value: 1=create, 2=modify, 3=delete)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODIFY_ACTION_UNKNOWN</td><td>0</td><td></td></tr><tr><td>MODIFY_ACTION_CREATE</td><td>1</td><td>Create</td></tr><tr><td>MODIFY_ACTION_UPDATE</td><td>2</td><td>Modify</td></tr><tr><td>MODIFY_ACTION_DELETE</td><td>3</td><td>Delete</td></tr></tbody></table></p>
	ModifyAction *int64 `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`

	// <p>Synonym list (pass the complete synonym collection during CREATE and UPDATE, overriding update)</p>
	SynonymList []*string `json:"SynonymList,omitnil,omitempty" name:"SynonymList"`

	// <p>Standard word (required for CREATE and UPDATE, can be left blank for DELETE)</p>
	Term *string `json:"Term,omitnil,omitempty" name:"Term"`

	// <p>Standard term ID (required for UPDATE and DELETE; leave blank for CREATE, generated by the backend)</p>
	TermId *string `json:"TermId,omitnil,omitempty" name:"TermId"`
}

type MCPPluginConfig struct {
	// <p>MCP plug-in public access address</p>
	ExternalMCPServerUrl *string `json:"ExternalMCPServerUrl,omitnil,omitempty" name:"ExternalMCPServerUrl"`

	// <p>MCP server address</p>
	MCPServerUrl *string `json:"MCPServerUrl,omitnil,omitempty" name:"MCPServerUrl"`

	// <p>MCP transmission type: SSE/Streamable<br>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 | SSE + HTTP mode |<br>| 1 | Streamable HTTP mode |</p>
	MCPTransport *int64 `json:"MCPTransport,omitnil,omitempty" name:"MCPTransport"`

	// <p>header parameters of the MCP plug-in</p>
	PluginHeader []*PluginParam `json:"PluginHeader,omitnil,omitempty" name:"PluginHeader"`

	// <p>query parameter of the MCP plug-in</p>
	PluginQuery []*PluginParam `json:"PluginQuery,omitnil,omitempty" name:"PluginQuery"`

	// <p>SSE persistent connection timeout period, in seconds</p>
	SSEReadTimeout *int64 `json:"SSEReadTimeout,omitnil,omitempty" name:"SSEReadTimeout"`

	// <p>Request timeout period, in seconds</p>
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>Authorization information.</p>
	AuthConfig *AuthConfig `json:"AuthConfig,omitnil,omitempty" name:"AuthConfig"`

	// <p>Whether the interactive interface (MCP Apps) and plug-in-level tags are supported. Default false</p>
	SupportsApps *bool `json:"SupportsApps,omitnil,omitempty" name:"SupportsApps"`
}

type MCPResourceContents struct {
	// <p>Binary content. Either this parameter or text must be specified.</p>
	Blob *string `json:"Blob,omitnil,omitempty" name:"Blob"`

	// <p>JSON: Original _meta of a single resource</p>
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// <p>MIME type. The HTML of the MCP App is text/html;profile=mcp-app</p>
	MimeType *string `json:"MimeType,omitnil,omitempty" name:"MimeType"`

	// <p>Text content. Either this field or blob is required.</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>Resource URI</p>
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
}

type MCPToolConfig struct {
	// <p>Input parameter</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>Output parameters</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// <p>Tool meta information</p>
	Meta *MCPToolMeta `json:"Meta,omitnil,omitempty" name:"Meta"`

	// <p>Whether the interactive interface (MCP Apps) and plug-in-level tags are supported. Default value: false</p>
	SupportsApps *bool `json:"SupportsApps,omitnil,omitempty" name:"SupportsApps"`
}

type MCPToolMeta struct {
	// <p>UI extended meta information of the tool, corresponding to _meta.UI in the MCP protocol, declaring the interactive UI resource (ResourceUri) associated with the tool and the caller Visibility. Returned only when the tool supports MCP Apps or declares Visibility; this field is empty for text tools. For details, see the MCPToolUIMeta structure definition.</p>
	Ui *MCPToolUIMeta `json:"Ui,omitnil,omitempty" name:"Ui"`
}

type MCPToolUIMeta struct {
	// <p>Associated UI resource URI in UI:// scheme, format: UI://&lt;plug-in flag&gt;/&lt;resource name&gt;-&lt;version&gt;. This field is the entry to the interactive UI of MCP Apps. If not empty, it means the tool supports Apps (shown as "text + interactive UI"). If empty, it is a plain text tool. It is automatically identified and filled based on the tool sync result and does not support manual editing.</p>
	ResourceUri *string `json:"ResourceUri,omitnil,omitempty" name:"ResourceUri"`

	// <p>Visibility declaration for tool callers. Value range: model (callable by models), app (callable by UI). Multiple selections allowed, for example [&quot;model&quot;,&quot;app&quot;]. Independent of ResourceUri (SEP-1865) and can exist independently. For example, a backend app-only tool is [&quot;app&quot;]. When ResourceUri is non-empty and this field is not specified, it is normalized to [&quot;model&quot;,&quot;app&quot;] by default. Existing non-Apps tools remain empty.</p><p>Enumeration values:</p><ul><li>model: Supports model</li><li>app: Supports app</li></ul>
	Visibility []*string `json:"Visibility,omitnil,omitempty" name:"Visibility"`
}

type ManualOnlySchedule struct {
	// Enable
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type MetaValue struct {
	// <p>Metadata value name (for display only).</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Whether to refer to all values of this type (when true, ref_value_id should be 0)</p>
	RefAll *bool `json:"RefAll,omitnil,omitempty" name:"RefAll"`

	// <p>Business ID referenced by metadata (property ID, category ID, etc.); this field should be 0 when ref_all=true</p>
	RefValueId *string `json:"RefValueId,omitnil,omitempty" name:"RefValueId"`

	// <p>Metadata usage scenario: 1=retrieve only, 2=both retrieve and generate<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>META_SCENE_UNKNOWN</td><td>0</td><td>Unknown</td></tr><tr><td>META_SCENE_SEARCH_ONLY</td><td>1</td><td>Retrieve only</td></tr><tr><td>META_SCENE_ALL</td><td>2</td><td>Both retrieve and generate</td></tr></tbody></table></p>
	Scene *int64 `json:"Scene,omitnil,omitempty" name:"Scene"`

	// <p>Metadata value type: 1=property tag, 2=document classification, 3=QA classification<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>META_VALUE_TYPE_UNKNOWN</td><td>0</td><td>Unknown</td></tr><tr><td>META_VALUE_TYPE_TAG</td><td>1</td><td>Property tag</td></tr><tr><td>META_VALUE_TYPE_DOC_CATEGORY</td><td>2</td><td>Document classification</td></tr><tr><td>META_VALUE_TYPE_QA_CATEGORY</td><td>3</td><td>QA classification</td></tr></tbody></table></p>
	ValueType *int64 `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

type MetricOverview struct {
	// <p>Metric key. For the value, see the key allowlist in the MetricOverview note.</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Month-over-month percentage. Enter 0 if there is no month-over-month comparison</p>
	Mom *float64 `json:"Mom,omitnil,omitempty" name:"Mom"`

	// <p>Metric unit, enumeration value DosageUnit; for the mapping between key and unit, see the MetricOverview comment allowlist</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token (default)</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count (page number)</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times (count)</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number (count)</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second (ms)</td></tr></tbody></table>
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// <p>Metric value.</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Model struct {
	// <p>Model badge list</p>
	BadgeList []*ModelBadge `json:"BadgeList,omitnil,omitempty" name:"BadgeList"`

	// <p>Model author information</p>
	DeveloperInfo *ModelDeveloperBasic `json:"DeveloperInfo,omitnil,omitempty" name:"DeveloperInfo"`

	// <p>Model restrictive information</p>
	LimitInfo *ModelLimit `json:"LimitInfo,omitnil,omitempty" name:"LimitInfo"`

	// <p>Model basic information</p>
	ModelBasic *ModelBasic `json:"ModelBasic,omitnil,omitempty" name:"ModelBasic"`

	// <p>Model hyperparameter configuration</p>
	ParameterList []*ModelParameter `json:"ParameterList,omitnil,omitempty" name:"ParameterList"`

	// <p>Model attribute configuration</p>
	PropertyList []*ModelProperty `json:"PropertyList,omitnil,omitempty" name:"PropertyList"`

	// <p>Model provider info</p>
	ProviderInfo *ModelProviderBasic `json:"ProviderInfo,omitnil,omitempty" name:"ProviderInfo"`

	// <p>Model status information</p>
	StatusInfo *ModelStatus `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// <p>Model tag list</p>
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ModelBadge struct {
	// <p>Display copy (rendered as final copy)</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>Style theme</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODEL_BADGE_THEME_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>MODEL_BADGE_THEME_INFO</td><td>1</td><td>Info (blue)</td></tr><tr><td>MODEL_BADGE_THEME_SUCCESS</td><td>2</td><td>Success (green)</td></tr><tr><td>MODEL_BADGE_THEME_WARNING</td><td>3</td><td>Warning (orange)</td></tr><tr><td>MODEL_BADGE_THEME_DANGER</td><td>4</td><td>Danger (red)</td></tr></table>
	Theme *int64 `json:"Theme,omitnil,omitempty" name:"Theme"`

	// <p>tooltip copywriting (rendered as final copywriting). If it is empty, it is not displayed.</p>
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// <p>Badge type</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODEL_BADGE_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>MODEL_BADGE_TYPE_FREE_TRIAL</td><td>1</td><td>Limited-time free</td></tr><tr><td>MODEL_BADGE_TYPE_DEPRECATING</td><td>2</td><td>About to go offline</td></tr><tr><td>MODEL_BADGE_TYPE_NEW</td><td>3</td><td>New model</td></tr><tr><td>MODEL_BADGE_TYPE_HOT</td><td>4</td><td>Popular</td></tr></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModelBasic struct {
	// <p>Model description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Model icon address</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Model ID</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>Model type</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODEL_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>MODEL_TYPE_LLM</td><td>1</td><td>LLM model</td></tr><tr><td>MODEL_TYPE_RERANK</td><td>2</td><td>Rerank model</td></tr><tr><td>MODEL_TYPE_EMBEDDING</td><td>3</td><td>Embedding model</td></tr><tr><td>MODEL_TYPE_FILE_PARSE</td><td>4</td><td>Document parsing model</td></tr></table>
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// <p>Model name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModelDetailInfo struct {
	// Model alias
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// Limit on number of historical conversations
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// Unique model ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// Model parameters
	ModelParams *ModelParams `json:"ModelParams,omitnil,omitempty" name:"ModelParams"`
}

type ModelDeveloperBasic struct {
	// <p>Author alias.</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>Author flag</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModelLimit struct {
	// <p>Model context length display copy (for example, 128K, 1000K)</p>
	ContextLengthDescription *string `json:"ContextLengthDescription,omitnil,omitempty" name:"ContextLengthDescription"`

	// <p>Character limit for model dialog box input length</p>
	InputLengthLimit *int64 `json:"InputLengthLimit,omitnil,omitempty" name:"InputLengthLimit"`

	// <p>Maximum input Tokens display copy (for example, 960k, formatted on the server, empty string if not configured)</p>
	MaxInputTokensDescription *string `json:"MaxInputTokensDescription,omitnil,omitempty" name:"MaxInputTokensDescription"`

	// <p>Display copy for the maximum output Tokens (for example, 64k, formatted on the server, empty string if not configured)</p>
	MaxOutputTokensDescription *string `json:"MaxOutputTokensDescription,omitnil,omitempty" name:"MaxOutputTokensDescription"`

	// <p>Model prompt content length character limit</p>
	PromptLengthLimit *int64 `json:"PromptLengthLimit,omitnil,omitempty" name:"PromptLengthLimit"`
}

type ModelParameter struct {
	// <p>Default value.</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>Available Option List</p>
	EnumValueList []*string `json:"EnumValueList,omitnil,omitempty" name:"EnumValueList"`

	// <p>Maximum value (valid only for numerical type).</p>
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// <p>Minimum value (valid only for numerical types).</p>
	MinValue *float64 `json:"MinValue,omitnil,omitempty" name:"MinValue"`

	// <p>Hyperparameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Hyperparameter type. 1-floating-point number, 2-integer, 3-string</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModelParams struct {
	// <p>Whether to enable deep thinking</p>
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

type ModelProperty struct {
	// <p>Property name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Attribute value.</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ModelProviderBasic struct {
	// <p>Model provider alias</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>Model provider name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Model provider type</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODEL_PROVIDER_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified</td></tr><tr><td>MODEL_PROVIDER_TYPE_SELF</td><td>1</td><td>Self-owned provider</td></tr><tr><td>MODEL_PROVIDER_TYPE_CUSTOM</td><td>2</td><td>Custom model provider</td></tr><tr><td>MODEL_PROVIDER_TYPE_THIRD</td><td>3</td><td>Third-party model provider</td></tr></table>
	ProviderType *int64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`
}

type ModelStatus struct {
	// <p>Number of exclusive concurrencies</p>
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// <p>Whether it is exclusive concurrency</p>
	IsExclusive *bool `json:"IsExclusive,omitnil,omitempty" name:"IsExclusive"`

	// <p>Whether direct experience in the intelligent workbench is supported (model plaza details page: dynamically determined based on space_id; when false, the frontend disables the "experience now" button)</p>
	IsExperienceSupported *bool `json:"IsExperienceSupported,omitnil,omitempty" name:"IsExperienceSupported"`

	// <p>Resource status</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>RESOURCE_STATUS_UNSPECIFIED</td><td>0</td><td>Not specified</td></tr><tr><td>RESOURCE_STATUS_AVAILABLE</td><td>1</td><td>Resource available</td></tr><tr><td>RESOURCE_STATUS_EXHAUSTED</td><td>2</td><td>Resource exhausted</td></tr></table>
	ResourceStatus *int64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// <p>RPM call cap (current account dimension: use the configured value if available; otherwise, take default. 0 means not obtained)</p>
	RpmLimit *string `json:"RpmLimit,omitnil,omitempty" name:"RpmLimit"`

	// <p>TPM call upper limit (at the current account level: use the configured value if available; otherwise, take default. 0 means failed to obtain)</p>
	TpmLimit *string `json:"TpmLimit,omitnil,omitempty" name:"TpmLimit"`
}

type ModelUsageDetail struct {
	// <p>Invocation type, originating from billing scene_billing (corresponding to filter.call_type)</p>
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// <p>Whether it is the default knowledge base. It is only meaningful for the business where ResourceType=5 (knowledge base capacity)</p>
	IsDefaultKB *bool `json:"IsDefaultKB,omitnil,omitempty" name:"IsDefaultKB"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>List of consumption measurements for a single call in the MODEL field (authoritative field): lists each measurement type by unit + label. When unit=TOKEN, label is case-sensitive for Token subcategories (input/output/avg_*/cache_*), and being empty indicates total_tokens; unit=PAGE_COUNT refers to the number of pages consumed by the model</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>PU consumption of this call</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`
}

type ModelUsageSummary struct {
	// <p>Number of calls (top-level count of API calls)</p>
	CallCount *float64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>Whether it is the default knowledge base</p>
	IsDefaultKB *bool `json:"IsDefaultKB,omitnil,omitempty" name:"IsDefaultKB"`

	// <p>Model name, flag indicating the AI model used</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>MODEL domain consumption measurement list (authoritative field): lists each measurement category by unit + label. When unit=TOKEN, label is case-sensitive for Token subcategories (input/output/avg_*/cache_*), with an empty label indicating total_tokens; unit=PAGE_COUNT refers to the model's consumed page count</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>Total model PU consumption (sum of PU consumption within the aggregate dimension)</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`
}

// Predefined struct for user
type ModifyAgentRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent Id</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Information of the Agent after modification</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Field path that needs to be updated, for example ["Profile.Name", "Profile.IconUrl", "Instructions", "Model", "ToolList", "PluginList", "SkillList", "AdvancedConfig"]</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent Id</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Information of the Agent after modification</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Field path that needs to be updated, for example ["Profile.Name", "Profile.IconUrl", "Instructions", "Model", "ToolList", "PluginList", "SkillList", "AdvancedConfig"]</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

func (r *ModifyAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "AgentId")
	delete(f, "Agent")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentResponseParams `json:"Response"`
}

func (r *ModifyAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application mode. Enumeration values: 1: Standard mode, 2: Agent mode, 3: Single workflow mode, 4: ClawAgent mode</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>App avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>Application configuration</p>
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Application description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Application name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>List of IDs of shared knowledge bases to refer (full coverage)</p>
	SharedKbIdList []*string `json:"SharedKbIdList,omitnil,omitempty" name:"SharedKbIdList"`

	// <p>App mount cloud disk</p>
	FileVolumeMountList []*AppFileVolumeMount `json:"FileVolumeMountList,omitnil,omitempty" name:"FileVolumeMountList"`

	// <p>Field mask, specifying the fields to update (if Paths is empty, no fields will be updated). Paths enumeration values:<br>[Top level] Name, Avatar, Description, AppMode, SharedKbIdList, FileVolumeMountList<br>[Greeting] Config.Greeting, Config.Greeting.Greeting, Config.Greeting.OpeningQuestionList<br>[Model] Config.Model, Config.Model.ThinkModel, Config.Model.GenerateModel, Config.Model.AiOptimizeModel, Config.Model.FileParseModel, Config.Model.PromptRewriteModel, Config.Model.MultiModalQaModel, Config.Model.MultiModalUnderstandingModel<br>[WebSearch] Config.WebSearch<br>[Memory] Config.Memory, Config.Memory.Enabled, Config.Memory.LongMemoryDay, Config.Memory.Model, Config.Memory.PromptMode, Config.Memory.PromptContent<br>[Mode] Config.Mode, Config.Mode.MultiAgentConfig, Config.Mode.SingleWorkflowConfig,Config.Mode.ClawAgentConfig<br> [Mode.ClawAgentConfig] Config.Mode.ClawAgentConfig.LongMemoryConfig, Config.Mode.ClawAgentConfig.AgentTeamConfig, Config.Mode.ClawAgentConfig.ChannelContextTimeoutConfig<br>[Experience] Config.Experience, Config.Experience.Conversation, Config.Experience.Role, Config.Experience.Advanced<br>[Experience.Conversation] Config.Experience.Conversation.AiCall, Config.Experience.Conversation.BackgroundImage, Config.Experience.Conversation.Method, Config.Experience.Conversation.FallbackReply, Config.Experience.Conversation.Recommended, Config.Experience.Conversation.InputBoxConfig, Config.Experience.Conversation.WebSearch,Config.Experience.Conversation.NewMessageInterrupt<br>[Experience.Conversation.AiCall] Config.Experience.Conversation.AiCall.VoiceInteract, Config.Experience.Conversation.AiCall.VoiceCall, Config.Experience.Conversation.AiCall.DigitalHuman<br>[Experience.Advanced] Config.Experience.Advanced.ContextRewrite, Config.Experience.Advanced.ImageTextRetrieval, Config.Experience.Advanced.IntentAchievement, Config.Experience.Advanced.ReplyFlexibility</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application mode. Enumeration values: 1: Standard mode, 2: Agent mode, 3: Single workflow mode, 4: ClawAgent mode</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>App avatar</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>Application configuration</p>
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Application description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Application name.</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>List of IDs of shared knowledge bases to refer (full coverage)</p>
	SharedKbIdList []*string `json:"SharedKbIdList,omitnil,omitempty" name:"SharedKbIdList"`

	// <p>App mount cloud disk</p>
	FileVolumeMountList []*AppFileVolumeMount `json:"FileVolumeMountList,omitnil,omitempty" name:"FileVolumeMountList"`

	// <p>Field mask, specifying the fields to update (if Paths is empty, no fields will be updated). Paths enumeration values:<br>[Top level] Name, Avatar, Description, AppMode, SharedKbIdList, FileVolumeMountList<br>[Greeting] Config.Greeting, Config.Greeting.Greeting, Config.Greeting.OpeningQuestionList<br>[Model] Config.Model, Config.Model.ThinkModel, Config.Model.GenerateModel, Config.Model.AiOptimizeModel, Config.Model.FileParseModel, Config.Model.PromptRewriteModel, Config.Model.MultiModalQaModel, Config.Model.MultiModalUnderstandingModel<br>[WebSearch] Config.WebSearch<br>[Memory] Config.Memory, Config.Memory.Enabled, Config.Memory.LongMemoryDay, Config.Memory.Model, Config.Memory.PromptMode, Config.Memory.PromptContent<br>[Mode] Config.Mode, Config.Mode.MultiAgentConfig, Config.Mode.SingleWorkflowConfig,Config.Mode.ClawAgentConfig<br> [Mode.ClawAgentConfig] Config.Mode.ClawAgentConfig.LongMemoryConfig, Config.Mode.ClawAgentConfig.AgentTeamConfig, Config.Mode.ClawAgentConfig.ChannelContextTimeoutConfig<br>[Experience] Config.Experience, Config.Experience.Conversation, Config.Experience.Role, Config.Experience.Advanced<br>[Experience.Conversation] Config.Experience.Conversation.AiCall, Config.Experience.Conversation.BackgroundImage, Config.Experience.Conversation.Method, Config.Experience.Conversation.FallbackReply, Config.Experience.Conversation.Recommended, Config.Experience.Conversation.InputBoxConfig, Config.Experience.Conversation.WebSearch,Config.Experience.Conversation.NewMessageInterrupt<br>[Experience.Conversation.AiCall] Config.Experience.Conversation.AiCall.VoiceInteract, Config.Experience.Conversation.AiCall.VoiceCall, Config.Experience.Conversation.AiCall.DigitalHuman<br>[Experience.Advanced] Config.Experience.Advanced.ContextRewrite, Config.Experience.Advanced.ImageTextRetrieval, Config.Experience.Advanced.IntentAchievement, Config.Experience.Advanced.ReplyFlexibility</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
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
	delete(f, "AppId")
	delete(f, "AppMode")
	delete(f, "Avatar")
	delete(f, "Config")
	delete(f, "Description")
	delete(f, "Name")
	delete(f, "SharedKbIdList")
	delete(f, "FileVolumeMountList")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Update time (Unix timestamp, in seconds)</p>
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
type ModifyAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger information</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>Unique ID of the trigger</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Modify field</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ModifyAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Trigger information</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>Unique ID of the trigger</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Modify field</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ModifyAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "Trigger")
	delete(f, "TriggerId")
	delete(f, "UpdateMask")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppTriggerResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppTriggerResponseParams `json:"Response"`
}

func (r *ModifyAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCategoryRequestParams struct {
	// <p>Category ID to be modified (must be greater than 0)</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Classification type (not allowed to be 0. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Modify field content (required, used in conjunction with update_mask)</p>
	Fields *CategoryModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: specifies the field to modify (supported Paths: Name)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>Category ID to be modified (must be greater than 0)</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Classification type (not allowed to be 0. Value: 1=document classification, 2=QA classification)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CATEGORY_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>CATEGORY_TYPE_DOC</td><td>1</td><td>Document classification</td></tr><tr><td>CATEGORY_TYPE_QA</td><td>2</td><td>QA classification</td></tr></tbody></table></p>
	CategoryType *int64 `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// <p>Modify field content (required, used in conjunction with update_mask)</p>
	Fields *CategoryModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: specifies the field to modify (supported Paths: Name)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

func (r *ModifyCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CategoryId")
	delete(f, "CategoryType")
	delete(f, "Fields")
	delete(f, "KbId")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCategoryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCategoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCategoryResponseParams `json:"Response"`
}

func (r *ModifyCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChannelRequestParams struct {

}

type ModifyChannelRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyChannelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChannelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChannelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChannelResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyChannelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChannelResponseParams `json:"Response"`
}

func (r *ModifyChannelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConflictQARequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Conflict group ID (optional if ignored globally)</p>
	ConflictGroupId *string `json:"ConflictGroupId,omitnil,omitempty" name:"ConflictGroupId"`

	// <p>Whether to ignore globally (ignore all pending conflict issues under the current KB)</p>
	IsIgnoreAll *bool `json:"IsIgnoreAll,omitnil,omitempty" name:"IsIgnoreAll"`

	// <p>QA ID list involved in decision-making (required for KEEP and DELETE)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`

	// <p>Conflict resolution policy: 1=reserve, 2=ignore, 3=delete (can be omitted when ignored globally)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CONFLICT_RESOLUTION_UNKNOWN</td><td>0</td><td>-</td></tr><tr><td>CONFLICT_RESOLUTION_KEEP</td><td>1</td><td>Reserve</td></tr><tr><td>CONFLICT_RESOLUTION_IGNORE</td><td>2</td><td>Ignore</td></tr><tr><td>CONFLICT_RESOLUTION_DELETE</td><td>3</td><td>Delete</td></tr></tbody></table></p>
	Resolution *int64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

type ModifyConflictQARequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Conflict group ID (optional if ignored globally)</p>
	ConflictGroupId *string `json:"ConflictGroupId,omitnil,omitempty" name:"ConflictGroupId"`

	// <p>Whether to ignore globally (ignore all pending conflict issues under the current KB)</p>
	IsIgnoreAll *bool `json:"IsIgnoreAll,omitnil,omitempty" name:"IsIgnoreAll"`

	// <p>QA ID list involved in decision-making (required for KEEP and DELETE)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`

	// <p>Conflict resolution policy: 1=reserve, 2=ignore, 3=delete (can be omitted when ignored globally)<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>CONFLICT_RESOLUTION_UNKNOWN</td><td>0</td><td>-</td></tr><tr><td>CONFLICT_RESOLUTION_KEEP</td><td>1</td><td>Reserve</td></tr><tr><td>CONFLICT_RESOLUTION_IGNORE</td><td>2</td><td>Ignore</td></tr><tr><td>CONFLICT_RESOLUTION_DELETE</td><td>3</td><td>Delete</td></tr></tbody></table></p>
	Resolution *int64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

func (r *ModifyConflictQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConflictQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "ConflictGroupId")
	delete(f, "IsIgnoreAll")
	delete(f, "QaIdList")
	delete(f, "Resolution")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConflictQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConflictQAResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConflictQAResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConflictQAResponseParams `json:"Response"`
}

func (r *ModifyConflictQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConflictQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConversationRequestParams struct {
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>User sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Session ID.
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// Session heading
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type ModifyConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session type. Enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (not specified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor trial), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>User sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type is CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Session ID.
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// Session heading
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

func (r *ModifyConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "ConversationId")
	delete(f, "Title")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConversationResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConversationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConversationResponseParams `json:"Response"`
}

func (r *ModifyConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocListRequestParams struct {
	// <p>Document ID list to be modified (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Ownership category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Application scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>External link</p>
	ExternalLink *DocExternalLink `json:"ExternalLink,omitnil,omitempty" name:"ExternalLink"`

	// <p>Tag list.</p>
	LabelRefList *LabelRefIdentityList `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Switch configuration</p>
	Switch *DocSwitch `json:"Switch,omitnil,omitempty" name:"Switch"`
}

type ModifyDocListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Document ID list to be modified (quantity: 1-20)</p>
	DocIdList []*string `json:"DocIdList,omitnil,omitempty" name:"DocIdList"`

	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Ownership category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Application scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>External link</p>
	ExternalLink *DocExternalLink `json:"ExternalLink,omitnil,omitempty" name:"ExternalLink"`

	// <p>Tag list.</p>
	LabelRefList *LabelRefIdentityList `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Switch configuration</p>
	Switch *DocSwitch `json:"Switch,omitnil,omitempty" name:"Switch"`
}

func (r *ModifyDocListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocIdList")
	delete(f, "KbId")
	delete(f, "CategoryId")
	delete(f, "EffectiveDomain")
	delete(f, "ExpirationPolicy")
	delete(f, "ExternalLink")
	delete(f, "LabelRefList")
	delete(f, "Switch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocListResponseParams struct {
	// <p>Batch modification result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocListResponseParams `json:"Response"`
}

func (r *ModifyDocListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocRequestParams struct {
	// <p>Document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Modify field contents (required, used together with update_mask)</p>
	Fields *DocModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Name, CategoryId, EffectiveDomain, LabelRefList, ExternalLink, ExpirationPolicy, UpdatePeriod, Switch, ParseConfig, UserAccessConfig)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyDocRequest struct {
	*tchttp.BaseRequest
	
	// <p>Document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Modify field contents (required, used together with update_mask)</p>
	Fields *DocModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Name, CategoryId, EffectiveDomain, LabelRefList, ExternalLink, ExpirationPolicy, UpdatePeriod, Switch, ParseConfig, UserAccessConfig)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
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
	delete(f, "DocId")
	delete(f, "Fields")
	delete(f, "KbId")
	delete(f, "UpdateMask")
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
type ModifyKBRequestParams struct {
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Writable attribute (use together with update_mask)</p>
	Spec *KBSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// <p>Field mask: specify the field to be modified (snake-shaped field name). Unlisted fields are ignored</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Scaling operation (for hosting special operations that cannot be categorized into regular field modification, such as triggering over-limit recovery. It is effective only when 'extend_fields' is also input in update_mask. For the value, refer to KBExtendedAction: 1=trigger over-limit recovery)</p>
	ExtendFields *KBModifyExtendFields `json:"ExtendFields,omitnil,omitempty" name:"ExtendFields"`
}

type ModifyKBRequest struct {
	*tchttp.BaseRequest
	
	// <p>Knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>Writable attribute (use together with update_mask)</p>
	Spec *KBSpec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// <p>Field mask: specify the field to be modified (snake-shaped field name). Unlisted fields are ignored</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Scaling operation (for hosting special operations that cannot be categorized into regular field modification, such as triggering over-limit recovery. It is effective only when 'extend_fields' is also input in update_mask. For the value, refer to KBExtendedAction: 1=trigger over-limit recovery)</p>
	ExtendFields *KBModifyExtendFields `json:"ExtendFields,omitnil,omitempty" name:"ExtendFields"`
}

func (r *ModifyKBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "Spec")
	delete(f, "UpdateMask")
	delete(f, "ExtendFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKBResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKBResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKBResponseParams `json:"Response"`
}

func (r *ModifyKBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLabelRequestParams struct {
	// <p>Modify field contents (required, use together with update_mask)</p>
	Fields *LabelModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>ID of the tag to be modified</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Name, TermModifyList)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyLabelRequest struct {
	*tchttp.BaseRequest
	
	// <p>Modify field contents (required, use together with update_mask)</p>
	Fields *LabelModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>Associated knowledge base ID</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>ID of the tag to be modified</p>
	LabelId *string `json:"LabelId,omitnil,omitempty" name:"LabelId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Name, TermModifyList)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

func (r *ModifyLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Fields")
	delete(f, "KbId")
	delete(f, "LabelId")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLabelResponseParams struct {
	// <p>Tag value after modification (standard word + synonym list)</p>
	TermList []*LabelTerm `json:"TermList,omitnil,omitempty" name:"TermList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLabelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLabelResponseParams `json:"Response"`
}

func (r *ModifyLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMsgRecordCategoryRequestParams struct {
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Business ID of the category to be modified</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Category name after modification</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>App ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Business ID of the category to be modified</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Category name after modification</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ModifyMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "CategoryId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMsgRecordCategoryResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *ModifyMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginRequestParams struct {
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Plugin version number</p>
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// <p>Plug-in basic information</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Plug-in type configuration</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Specify the fields to update to avoid full coverage</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Tool list of the plug-in. This parameter is not required for mcp plug-ins.</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ModifyPluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Plugin version number</p>
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// <p>Plug-in basic information</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Plug-in type configuration</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Specify the fields to update to avoid full coverage</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>Tool list of the plug-in. This parameter is not required for mcp plug-ins.</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>User Main Account logged in (required in integrator mode)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>User Sub-account logged in (required in integrator mode)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ModifyPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "PluginVersion")
	delete(f, "Profile")
	delete(f, "Config")
	delete(f, "UpdateMask")
	delete(f, "ToolList")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPluginResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPluginResponseParams `json:"Response"`
}

func (r *ModifyPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAListRequestParams struct {
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of QA IDs to be modified (quantity: 1-20)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`

	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Effective scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Whether to adopt (verification passed)</p>
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// <p>Applicable scope (tag condition list).</p>
	LabelRefList *LabelRefIdentityList `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`
}

type ModifyQAListRequest struct {
	*tchttp.BaseRequest
	
	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>List of QA IDs to be modified (quantity: 1-20)</p>
	QaIdList []*string `json:"QaIdList,omitnil,omitempty" name:"QaIdList"`

	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Effective scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Whether to adopt (verification passed)</p>
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// <p>Applicable scope (tag condition list).</p>
	LabelRefList *LabelRefIdentityList `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`
}

func (r *ModifyQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KbId")
	delete(f, "QaIdList")
	delete(f, "CategoryId")
	delete(f, "EffectiveDomain")
	delete(f, "ExpirationPolicy")
	delete(f, "IsAccepted")
	delete(f, "LabelRefList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAListResponseParams struct {
	// <p>Batch modification result.</p>
	ResultList []*OperationResult `json:"ResultList,omitnil,omitempty" name:"ResultList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAListResponseParams `json:"Response"`
}

func (r *ModifyQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQARequestParams struct {
	// <p>Modify field content (used in conjunction with update_mask)</p>
	Fields *QAModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Question, Answer, CategoryId, DocId, LabelRefList, QuestionDescription, ExpirationPolicy, SimilarQuestionList, EffectiveDomain, IsAccepted)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// <p>Modify field content (used in conjunction with update_mask)</p>
	Fields *QAModifyFields `json:"Fields,omitnil,omitempty" name:"Fields"`

	// <p>ID of the knowledge base it belongs to</p>
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// <p>Field mask: specifies the fields to be modified (supported Paths: Question, Answer, CategoryId, DocId, LabelRefList, QuestionDescription, ExpirationPolicy, SimilarQuestionList, EffectiveDomain, IsAccepted)</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
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
	delete(f, "Fields")
	delete(f, "KbId")
	delete(f, "QaId")
	delete(f, "UpdateMask")
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
type ModifySkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>skill package file address (zip); importing it triggers new version generation. It must be imported together with SkillVersion and UpdateDescription</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>Icon address</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill version number (import together with FileUrl)</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>Version change description (import together with FileUrl).</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

type ModifySkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>skill package file address (zip); importing it triggers new version generation. It must be imported together with SkillVersion and UpdateDescription</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>Icon address</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill version number (import together with FileUrl)</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>Version change description (import together with FileUrl).</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

func (r *ModifySkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "DisplayDescription")
	delete(f, "DisplayName")
	delete(f, "FileUrl")
	delete(f, "IconUrl")
	delete(f, "SkillVersion")
	delete(f, "UpdateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySkillResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySkillResponse struct {
	*tchttp.BaseResponse
	Response *ModifySkillResponseParams `json:"Response"`
}

func (r *ModifySkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpaceRequestParams struct {
	// Workspace name. Maximum length: 30 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Space description. Maximum length: 150 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Space id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Fields to update. `Name` and `Description` supported.
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`
}

type ModifySpaceRequest struct {
	*tchttp.BaseRequest
	
	// Workspace name. Maximum length: 30 characters.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Space description. Maximum length: 150 characters.
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Space id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Fields to update. `Name` and `Description` supported.
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`
}

func (r *ModifySpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SpaceId")
	delete(f, "FieldMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpaceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpaceResponseParams `json:"Response"`
}

func (r *ModifySpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Variable information
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

type ModifyVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Variable information
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

func (r *ModifyVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Variable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVariableResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVariableResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVariableResponseParams `json:"Response"`
}

func (r *ModifyVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgRecord struct {
	// Answer.
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Category ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Intent.
	Intent *string `json:"Intent,omitnil,omitempty" name:"Intent"`

	// Intent category
	IntentCategory *string `json:"IntentCategory,omitnil,omitempty" name:"IntentCategory"`

	// Whether it is intelligent categorization
	IsSmart *bool `json:"IsSmart,omitnil,omitempty" name:"IsSmart"`

	// Issue
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// Record ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// method. Enumerates ReplyMethod: 0=not specified, 1=direct reply by large model, 2=conservative reply, 3=refuse to answer, 4=sensitive reply, 5=Q&A pair prioritized reply, 6=welcome message, 7=concurrency limit exceeded, 8=global intervention knowledge, 9=task flow process reply, 10=task flow answer, 11=search engine, 12=knowledge polish, 13=image understanding, 14=real-time document, 15=clarification confirmation, 16=workflow reply, 17=workflow completion, 18=Intelligent Agent reply, 19=multiple intents, 20=interruption, 21=Intelligent Agent plan preview, 22=Intelligent Agent plan result, 23=Intelligent Agent structured output.
	ReplyMethod *int64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`

	// Returned result
	Result *MsgRecordResult `json:"Result,omitnil,omitempty" name:"Result"`

	// Score
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// Session ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Source.
	Source *MsgRecordSource `json:"Source,omitnil,omitempty" name:"Source"`

	// trace_id
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

type MsgRecordCategory struct {
	// <p>Business ID of the category</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Subcategory list, tree nesting</p>
	Children []*MsgRecordCategory `json:"Children,omitnil,omitempty" name:"Children"`

	// <p>Category name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Operation permissions of the current user on this category.</p>
	Permission *CategoryPermission `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>The number of message records in this category</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type MsgRecordResult struct {
	// Indicates the call result of the message: 0=CALL_RESULT_UNKNOWN (all/unknown), 1=CALL_RESULT_SUCCESS (success), 2=CALL_RESULT_FAILED (failed). fail_reason (string) describes the failure reason when the call failed.
	CallResult *int64 `json:"CallResult,omitnil,omitempty" name:"CallResult"`

	// Custom variable, json string
	CustomerVariable *string `json:"CustomerVariable,omitnil,omitempty" name:"CustomerVariable"`

	// Failure cause.
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// Time to first token
	FirstTokenLatency *uint64 `json:"FirstTokenLatency,omitnil,omitempty" name:"FirstTokenLatency"`

	// Number of input tokens.
	InputToken *uint64 `json:"InputToken,omitnil,omitempty" name:"InputToken"`

	// Number of output tokens.
	OutputToken *uint64 `json:"OutputToken,omitnil,omitempty" name:"OutputToken"`

	// Total number of tokens.
	TotalToken *uint64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// Total token duration
	TotalTokenLatency *uint64 `json:"TotalTokenLatency,omitnil,omitempty" name:"TotalTokenLatency"`
}

type MsgRecordSource struct {
	// Channel type for dialogue message access: 0=not specified, 1=agent, 2=experience page (Tencent Cloud), 3=evaluation dialogue, 4=experience page (phone number), 5=dialogue API connection, 6=evaluation task dialogue, 10=workflow debugging, 10000=WeChat official account, 10001=WeChat service account, 10002=WeCom app, 10003=web components, 10004=WeChat Customer Service, 10005=WeChat mini program, 10006=Yuanqi, 10007=App Treasure, 10008=Yuanbao, 10009=WeCom intelligent bot, 10010=Yuanqi API, 10011=LINE, 10012=Telegram, 10100=PC manager, 20001=honor Intelligent Agent platform, 20002=Xiaomi app store; user_id (string) is the unique ID of the visitor under this channel.
	ChannelType *uint64 `json:"ChannelType,omitnil,omitempty" name:"ChannelType"`

	// User ID.
	FromId *string `json:"FromId,omitnil,omitempty" name:"FromId"`

	// User source type of the message sender: 1=user (visitor/C-end user), 2=robot (AI reply), 3=agent (human customer service); from_id (string) is the unique user ID under this source type.
	FromType *uint64 `json:"FromType,omitnil,omitempty" name:"FromType"`

	// User profile photo
	UserAvatar *string `json:"UserAvatar,omitnil,omitempty" name:"UserAvatar"`

	// Visitor ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Visitor name
	UserNickname *string `json:"UserNickname,omitnil,omitempty" name:"UserNickname"`
}

type MultiAgentConfig struct {
	// Agent collaboration configuration
	AgentCollaboration *AgentCollaborationConfig `json:"AgentCollaboration,omitnil,omitempty" name:"AgentCollaboration"`
}

type MultiModalQAModel struct {
	// Model configuration
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type MultiModalUnderstandingModel struct {
	// Model configuration
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type NodeRun struct {
	// ID of the node operation
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`

	// Node ID.
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// ID of the workflow running instance
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// Node name.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Node type.
	// 1: Start node
	// 2: Parameter extraction node
	// 3: Large model node
	// 4: Knowledge Q&A node
	// 5: Knowledge retrieval node
	// 6: Tag extraction node
	// 7: Code execution node
	// 8: Tool node
	// 9: Logic judgment node
	// 10: Reply node
	// 11: Tab node
	// 12: Loop node
	// 13: Intent recognition node
	// 14: Workflow node
	// 15: Plug-in node
	// 16: End node
	// 17: Variable aggregate node data
	// 18: Batch processing node
	// 19: Message Queue node
	// 20: Database node
	// 21: Variable assignment node data
	// 22: Variable conversion node data
	// 23: Agent node
	// 24: Annotation node
	// 25: File collection node
	// 26: Text collection node
	// 27: Widget node
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// Running state. 1: initial status; 2: running; 3: operation successful; 4: operation failure; 5: canceled
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// Total time consumed (ms)
	TotalCostMs *int64 `json:"TotalCostMs,omitnil,omitempty" name:"TotalCostMs"`

	// Total number of tokens consumed
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// Error code.
	FailCode *string `json:"FailCode,omitnil,omitempty" name:"FailCode"`

	// List of branch indexes
	BranchIndexList []*int64 `json:"BranchIndexList,omitnil,omitempty" name:"BranchIndexList"`
}

type NodeRunContent struct {
	// Input variable information
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// Output variable information
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Original output information. Only partial nodes have a value, such as tool nodes and code nodes.
	OriginalOutput *string `json:"OriginalOutput,omitnil,omitempty" name:"OriginalOutput"`

	// URLs of the complete node Input content. (This field has a value only when the Input content exceeds the limit.)
	InputRef *string `json:"InputRef,omitnil,omitempty" name:"InputRef"`

	// URLs of the complete node Output. This field has a value only when the Output exceeds the limit.
	OutputRef *string `json:"OutputRef,omitnil,omitempty" name:"OutputRef"`

	// URLs of the complete original output content of the task. (This field has a value only when the TaskOutput content exceeds the limit.)
	OriginalOutputRef *string `json:"OriginalOutputRef,omitnil,omitempty" name:"OriginalOutputRef"`
}

type NodeRunDetail struct {
	// <p>Node running information</p>
	NodeRun *NodeRun `json:"NodeRun,omitnil,omitempty" name:"NodeRun"`

	// <p>Node operation content.</p>
	NodeRunContent *NodeRunContent `json:"NodeRunContent,omitnil,omitempty" name:"NodeRunContent"`

	// <p>Node operation process</p>
	NodeRunExecution *NodeRunExecution `json:"NodeRunExecution,omitnil,omitempty" name:"NodeRunExecution"`

	// <p>Node operation logs</p>
	NodeRunLog *NodeRunLog `json:"NodeRunLog,omitnil,omitempty" name:"NodeRunLog"`

	// <p>Statistical information.</p>
	StatisticList []*Statistic `json:"StatisticList,omitnil,omitempty" name:"StatisticList"`
}

type NodeRunExecution struct {
	// <p>Batch operation summary</p>
	BatchRunSummary *BatchRunSummary `json:"BatchRunSummary,omitnil,omitempty" name:"BatchRunSummary"`

	// <p>End time (timestamp in milliseconds).</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time (timestamp in milliseconds)</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type NodeRunLog struct {
	// <p>Logs of node.</p>
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// <p>URLs of the complete content of logs of node (has a value only when the Log content exceeds the limit)</p>
	LogRef *string `json:"LogRef,omitnil,omitempty" name:"LogRef"`
}

type OAuthConfig struct {
	// url address of the OAuth service provider authorization page
	AuthorizationUrl *string `json:"AuthorizationUrl,omitnil,omitempty" name:"AuthorizationUrl"`

	// Client ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// Client Key
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// Data scope for the authorization request
	ScopeList []*string `json:"ScopeList,omitnil,omitempty" name:"ScopeList"`

	// url address for obtaining the access token
	TokenUrl *string `json:"TokenUrl,omitnil,omitempty" name:"TokenUrl"`
}

type OnceSchedule struct {
	// Trigger time.
	FireTime *string `json:"FireTime,omitnil,omitempty" name:"FireTime"`
}

type OperationResult struct {
	// <p>Resource ID.</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>Failure reason (populated when succeeded=false)</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>Whether it is successful</p>
	Succeeded *bool `json:"Succeeded,omitnil,omitempty" name:"Succeeded"`
}

type Operator struct {
	// User ID.
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

// Predefined struct for user
type PauseAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type PauseAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *PauseAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseAppTriggerResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *PauseAppTriggerResponseParams `json:"Response"`
}

func (r *PauseAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlatformUsageSummary struct {
	// <p>PLATFORM domain consumption measurement list (authoritative field): list each measurement type by unit + label, where label is the PlatformBizType enumeration name string; typical examples include unit=TIMES + label=PLATFORM_BIZ_TYPE_SECURITY_AUDIT/WEB_SEARCH/OPEN_CLAW/APP_INVOKE, unit=ITEM + label=PLATFORM_BIZ_TYPE_LONG_TERM_MEMORY</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`
}

type Plugin struct {
	// Plug-in configuration
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// Creation time, Unix Timestamp
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Plug-in operation management information
	Operation *PluginOperation `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Plugin ID
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// Plugin version number
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// Basic plug-in information.
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Plug-in statistical information
	Statistics *PluginStatistics `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// <p>Plug-in status. 1: available, 2: unavailable </p><p>Enumeration values:</p><ul><li>1: available</li><li>2: unavailable</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Tool list
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// Update time, Unix timestamp
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Plugin status information by user dimension
	UserState *PluginUserState `json:"UserState,omitnil,omitempty" name:"UserState"`
}

type PluginConfig struct {
	// API plug-in configuration
	ApiPluginConfig *ApiPluginConfig `json:"ApiPluginConfig,omitnil,omitempty" name:"ApiPluginConfig"`

	// App plugin configuration
	AppPluginConfig *AppPluginConfig `json:"AppPluginConfig,omitnil,omitempty" name:"AppPluginConfig"`

	// mcp plug-in configuration
	MCPPluginConfig *MCPPluginConfig `json:"MCPPluginConfig,omitnil,omitempty" name:"MCPPluginConfig"`
}

type PluginOperation struct {
	// Whether external calls are allowed
	AllowExternalAccess *bool `json:"AllowExternalAccess,omitnil,omitempty" name:"AllowExternalAccess"`

	// <p>Billing type.</p><p>Valid values:</p><ul><li>0: free</li><li>1: open beta</li><li>2: official fees</li></ul>
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// Plug-in category identifier
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// Plug-in overview
	Introduction *string `json:"Introduction,omitnil,omitempty" name:"Introduction"`

	// Whether it is a selection
	IsRecommended *bool `json:"IsRecommended,omitnil,omitempty" name:"IsRecommended"`
}

type PluginParam struct {
	// Whether to hide parameter configuration
	IsGlobalHidden *bool `json:"IsGlobalHidden,omitnil,omitempty" name:"IsGlobalHidden"`

	// Whether the parameter is required
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// Parameter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Parameter value.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PluginProfile struct {
	// Plug-in author
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// Plugin description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Plug-in icon url
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// Plugin name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Plug-in product type</p><p>Enumeration values:</p><ul><li>0: ordinary plug-in</li><li>1: connector plug-in</li></ul>
	PluginClass *int64 `json:"PluginClass,omitnil,omitempty" name:"PluginClass"`

	// <p>Plugin type</p><p>Enumeration values:</p><ul><li>0: API</li><li>1: Code</li><li>2: MCP</li><li>3: Application</li></ul>
	PluginKind *int64 `json:"PluginKind,omitnil,omitempty" name:"PluginKind"`

	// <p>Plug-in source</p><p>Enumeration values:</p><ul><li>0: custom plugin</li><li>1: official plug-in</li><li>2: third-party plugin</li></ul>
	PluginSource *int64 `json:"PluginSource,omitnil,omitempty" name:"PluginSource"`
}

type PluginStatistics struct {
	// Plug-in calling volume
	CallCount *uint64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// Number of tools
	ToolCount *int64 `json:"ToolCount,omitnil,omitempty" name:"ToolCount"`
}

type PluginSummary struct {
	// <p>Plug-in configuration message.</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>Whether shared is configured</p>
	IsShared *bool `json:"IsShared,omitnil,omitempty" name:"IsShared"`

	// <p>Plugin operation management information.</p>
	Operation *PluginOperation `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <p>plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Basic information of the plug-in.</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>Space ID that the plug-in belongs to; empty for built-in plugins</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Plugin statistical information.</p>
	Statistics *PluginStatistics `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// <p>Plug-in status. 1: available, 2: unavailable </p><p>Enumeration values:</p><ul><li>1: available</li><li>2: unavailable</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Tool information</p>
	ToolList []*ToolSummary `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>Plugin status information by user dimension</p>
	UserState *PluginUserState `json:"UserState,omitnil,omitempty" name:"UserState"`

	// <p>Update time, Unix timestamp</p><p>Unit: seconds</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type PluginUsageDetail struct {
	// <p>Plugin name.</p>
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// <p>Consumption measurement list for a single PLUGIN domain call (authoritative field): lists each measurement category by unit + label. When unit=TOKEN, label is case-sensitive for Token subcategories (input/output/avg_*). Being empty indicates total_tokens</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>Plug-in Tool Name (tool_name)</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`
}

type PluginUsageSummary struct {
	// <p>Number of calls (top-level count of API calls)</p>
	CallCount *float64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>Consumption measurement list for the PLUGIN domain (authoritative field): lists each measurement category by unit + label. When unit=TOKEN, label is case-sensitive for Token subcategories (input/output/avg_*). Being empty indicates total_tokens</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`
}

type PluginUserState struct {
	// Whether the plug-in has been bookmarked
	IsFavorite *bool `json:"IsFavorite,omitnil,omitempty" name:"IsFavorite"`

	// Whether it is in the plug-in allowlist
	IsInWhiteList *bool `json:"IsInWhiteList,omitnil,omitempty" name:"IsInWhiteList"`

	// <p>Allowlist type, used to represent whether the current user can use this plug-in directly.</p><p>Enumeration values:</p><ul><li>0: Not on the allowlist, fully available</li><li>1: Current user in allowlist</li><li>2: Current user not in allowlist, submit an application</li></ul>
	WhiteListType *int64 `json:"WhiteListType,omitnil,omitempty" name:"WhiteListType"`
}

type PortalShareConfig struct {
	// <p>Business shared switch</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type PromptRewriteModel struct {
	// Model configuration
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type QACreateSpec struct {
	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Associated document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Knowledge effective scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Applicable scope (tag condition)</p>
	LabelRefList []*LabelRefIdentity `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Issue description</p>
	QuestionDescription *string `json:"QuestionDescription,omitnil,omitempty" name:"QuestionDescription"`

	// <p>Similar question list</p>
	SimilarQuestionList []*string `json:"SimilarQuestionList,omitnil,omitempty" name:"SimilarQuestionList"`
}

type QALifecycle struct {
	// <p>Creation time (Unix seconds).</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Status: 1=pending verification, 2=not adopted, 3=import failed, 4=under review, 5=review failed, 6=Learning, 7=learn failed, 8=import completed, 9=expired, 10=quota invalid, 11=quota recovering, 12=manual appeal in progress, 13=manual appeal failed<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>QA_STATUS_UNKNOWN</td><td>0</td><td></td></tr><tr><td>QA_STATUS_PENDING_VERIFY</td><td>1</td><td>Pending verification</td></tr><tr><td>QA_STATUS_NOT_ACCEPTED</td><td>2</td><td>Not adopted</td></tr><tr><td>QA_STATUS_IMPORT_FAIL</td><td>3</td><td>Import failed</td></tr><tr><td>QA_STATUS_AUDITING</td><td>4</td><td>Under review</td></tr><tr><td>QA_STATUS_AUDIT_FAIL</td><td>5</td><td>Review failed</td></tr><tr><td>QA_STATUS_LEARNING</td><td>6</td><td>Learning</td></tr><tr><td>QA_STATUS_LEARN_FAIL</td><td>7</td><td>Learn failed</td></tr><tr><td>QA_STATUS_IMPORTED</td><td>8</td><td>Import completed</td></tr><tr><td>QA_STATUS_EXPIRED</td><td>9</td><td>Expired</td></tr><tr><td>QA_STATUS_QUOTA_INVALID</td><td>10</td><td>Quota invalid</td></tr><tr><td>QA_STATUS_QUOTA_RECOVERING</td><td>11</td><td>Quota recovering</td></tr><tr><td>QA_STATUS_MANUAL_APPEALING</td><td>12</td><td>Manual appeal in progress</td></tr><tr><td>QA_STATUS_MANUAL_APPEAL_FAIL</td><td>13</td><td>Manual appeal failed</td></tr></tbody></table></p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description.</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>Status additional information</p>
	StatusMessage *string `json:"StatusMessage,omitnil,omitempty" name:"StatusMessage"`

	// <p>Update time (Unix seconds).</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type QAMetadata struct {
	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// <p>QA characters</p>
	QaCharCount *string `json:"QaCharCount,omitnil,omitempty" name:"QaCharCount"`

	// <p>QA size (byte, including similar questions)</p>
	QaSize *string `json:"QaSize,omitnil,omitempty" name:"QaSize"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Metadata reference field name list (for displaying which QA categories and attributes are set as metadata)</p>
	RefFieldNameList []*string `json:"RefFieldNameList,omitnil,omitempty" name:"RefFieldNameList"`
}

type QAModifyFields struct {
	// <p>Answer</p>
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// <p>Category ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>Associated document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Knowledge application scope<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	EffectiveDomain *int64 `json:"EffectiveDomain,omitnil,omitempty" name:"EffectiveDomain"`

	// <p>Expiration policy (validity time and behavior after the validity period expires)</p>
	ExpirationPolicy *ExpirationPolicy `json:"ExpirationPolicy,omitnil,omitempty" name:"ExpirationPolicy"`

	// <p>Whether to adopt (verification passed)</p>
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// <p>Applicable scope (tag condition).</p>
	LabelRefList []*LabelRefIdentity `json:"LabelRefList,omitnil,omitempty" name:"LabelRefList"`

	// <p>Query</p>
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// <p>Problem description</p>
	QuestionDescription *string `json:"QuestionDescription,omitnil,omitempty" name:"QuestionDescription"`

	// <p>Similar question modification list</p>
	SimilarQuestionList []*SimilarQuestionModifySpec `json:"SimilarQuestionList,omitnil,omitempty" name:"SimilarQuestionList"`
}

type QAOperator struct {
	// <p>Modifier</p>
	Modifier *Operator `json:"Modifier,omitnil,omitempty" name:"Modifier"`

	// <p>Operation permission</p>
	Permission *QAPermission `json:"Permission,omitnil,omitempty" name:"Permission"`
}

type QAPermission struct {
	// <p>Whether verification is supported (adoption/non-adoption).</p>
	CanAccept *bool `json:"CanAccept,omitnil,omitempty" name:"CanAccept"`

	// <p>Whether it is deletable.</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>Whether it is editable</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`
}

type QAQuery struct {
	// <p>Query keyword (fuzzy search)</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>Query scope (fields that query acts on): 1=question, 2=tag or tag value, 3=answer; multiple selections supported, unavailable by default</p>
	QueryScopeList []*int64 `json:"QueryScopeList,omitnil,omitempty" name:"QueryScopeList"`
}

type QARetrievalConfig struct {
	// <p>Confidence threshold</p>
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Return the top N entries</p>
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type QASegmentHighlight struct {
	// <p>Highlight end position</p>
	EndPos *string `json:"EndPos,omitnil,omitempty" name:"EndPos"`

	// <p>Highlight starting position</p>
	StartPos *string `json:"StartPos,omitnil,omitempty" name:"StartPos"`
}

type QASourceInfo struct {
	// <p>Application scope of the associated document<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_UNKNOWN</td><td>0</td><td></td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_NONE</td><td>1</td><td>Disabled</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_DEV</td><td>2</td><td>Development domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_RELEASE</td><td>3</td><td>Release domain only</td></tr><tr><td>KNOWLEDGE_EFFECTIVE_DOMAIN_ALL</td><td>4</td><td>Global</td></tr></tbody></table></p>
	DocEffectiveDomain *int64 `json:"DocEffectiveDomain,omitnil,omitempty" name:"DocEffectiveDomain"`

	// <p>Associated document ID</p>
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// <p>Associated document name</p>
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// <p>Associated document type</p>
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// <p>Source description.</p>
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// <p>Source type: 1=documentation generation, 2=batch import, 3=manual entry<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>QA_SOURCE_TYPE_UNKNOWN</td><td>0</td><td></td></tr><tr><td>QA_SOURCE_TYPE_DOC</td><td>1</td><td>Documentation generation</td></tr><tr><td>QA_SOURCE_TYPE_BATCH_IMPORT</td><td>2</td><td>Batch import</td></tr><tr><td>QA_SOURCE_TYPE_MANUAL</td><td>3</td><td>Manual entry</td></tr></tbody></table></p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type QASummary struct {
	// <p>Category path (including category ID, category ID path from the root node, and category name path)</p>
	CategoryPath *CategoryPath `json:"CategoryPath,omitnil,omitempty" name:"CategoryPath"`

	// <p>Knowledge effective scope</p>
	KnowledgeScope *KnowledgeScope `json:"KnowledgeScope,omitnil,omitempty" name:"KnowledgeScope"`

	// <p>Lifecycle information.</p>
	Lifecycle *QALifecycle `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// <p>Meta information (question/answer/size statistics)</p>
	Metadata *QAMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>Operator information.</p>
	OperatorInfo *QAOperator `json:"OperatorInfo,omitnil,omitempty" name:"OperatorInfo"`

	// <p>QA ID</p>
	QaId *string `json:"QaId,omitnil,omitempty" name:"QaId"`

	// <p>Similar question statistics</p>
	SimilarQuestion *SimilarQuestionStat `json:"SimilarQuestion,omitnil,omitempty" name:"SimilarQuestion"`

	// <p>Source information.</p>
	SourceInfo *QASourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

// Predefined struct for user
type ReadMCPResourceRequestParams struct {
	// <p>Capacity boundary: one request can only read the resources belonging to this plugin</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>thread on the agent-exec side</p>
	ThreadId *string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// <p>Resource URI, source: _meta.ui.resourceUri of the tool</p>
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
}

type ReadMCPResourceRequest struct {
	*tchttp.BaseRequest
	
	// <p>Capacity boundary: one request can only read the resources belonging to this plugin</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>thread on the agent-exec side</p>
	ThreadId *string `json:"ThreadId,omitnil,omitempty" name:"ThreadId"`

	// <p>Resource URI, source: _meta.ui.resourceUri of the tool</p>
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
}

func (r *ReadMCPResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReadMCPResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "ThreadId")
	delete(f, "Uri")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReadMCPResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReadMCPResourceResponseParams struct {
	// <p>Resource content returned by MCP resources/read</p>
	Contents []*MCPResourceContents `json:"Contents,omitnil,omitempty" name:"Contents"`

	// <p>JSON: The original _meta of Resource. In the MCP App scenario, Host enforces the CSP policy based on ui.csp.</p>
	ResourceMeta *string `json:"ResourceMeta,omitnil,omitempty" name:"ResourceMeta"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReadMCPResourceResponse struct {
	*tchttp.BaseResponse
	Response *ReadMCPResourceResponseParams `json:"Response"`
}

func (r *ReadMCPResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReadMCPResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseRecord struct {
	// Whether it can be exported
	CanExport *bool `json:"CanExport,omitnil,omitempty" name:"CanExport"`

	// Whether it can be rolled back
	CanRollback *bool `json:"CanRollback,omitnil,omitempty" name:"CanRollback"`

	// Release description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Number of release failures
	FailCount *uint64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// Failure cause.
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Release ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

	// Release version
	ReleaseVersion *string `json:"ReleaseVersion,omitnil,omitempty" name:"ReleaseVersion"`

	// Publishing status. Enumeration values: 1: To be released, 2: Releasing, 3: Released successfully, 4: Release failure, 5: Under review, 6: Review successful, 7: Review failed, 8: Release success callback processing, 9: Release suspended, 10: Appeal under review, 11: Appeal review passed, 12: Appeal review not approved
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Status description.
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`

	// Number of successful releases
	SuccessCount *uint64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// Update time (Unix timestamp in seconds).
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Releaser
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`
}

// Predefined struct for user
type ReleaseSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Version ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type ReleaseSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Version ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *ReleaseSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseSkillResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseSkillResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseSkillResponseParams `json:"Response"`
}

func (r *ReleaseSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseSummary struct {
	// <p>Creation time (Unix timestamp, in seconds)</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Release description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Release ID</p>
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

	// <p>Released status. Enumeration values: 1: To be released, 2: Releasing, 3: Released successfully, 4: Release failure, 5: Under review, 6: Review successful, 7: Review failed, 8: Callback processing for successful release, 9: Release suspension, 10: Appeal under review, 11: Appeal passed, 12: Appeal not approved</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Status description.</p>
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`

	// <p>Application sharing access control</p>
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>Release channel ID list</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>Business shared configuration</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>portal configuration</p>
	PortalShareConfig *PortalShareConfig `json:"PortalShareConfig,omitnil,omitempty" name:"PortalShareConfig"`
}

type ReplyContent struct {
	// <p>Message content type<br>text: text<br>image: image<br>file: file<br>option_cards: tab<br>custom_params: user-customized business parameter<br>sandbox: cloud desktop<br>custom_variables: custom input parameter<br>web_search: web search content<br>file_collection: file collection info<br>widget: widget info<br>widget_action: client-side widget action information</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Text content</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>Tab information</p>
	OptionCardList []*string `json:"OptionCardList,omitnil,omitempty" name:"OptionCardList"`

	// <p>File collection info</p>
	FileCollection *ConversationFileCollection `json:"FileCollection,omitnil,omitempty" name:"FileCollection"`

	// <p>Widget information.</p>
	Widget *ConversationWidget `json:"Widget,omitnil,omitempty" name:"Widget"`
}

type ReplyMessage struct {
	// Message array containing multiple message objects
	ContentList []*ReplyContent `json:"ContentList,omitnil,omitempty" name:"ContentList"`
}

type RequestParam struct {
	// <p>Parameters in types of AnyOf</p>
	AnyOf []*RequestParam `json:"AnyOf,omitnil,omitempty" name:"AnyOf"`

	// <p>Default value.</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>Parameter description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Globally hidden and invisible (different from agent_hidden in the Agent scenario). true: globally hidden and invisible; false: visible</p>
	IsGlobalHidden *bool `json:"IsGlobalHidden,omitnil,omitempty" name:"IsGlobalHidden"`

	// <p>Required or optional</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>Parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Parameters in types of OneOf</p>
	OneOf []*RequestParam `json:"OneOf,omitnil,omitempty" name:"OneOf"`

	// <p>Sub-parameter, used only when ParamType is OBJECT or ARRAY&lt;&gt;</p>
	SubParams []*RequestParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>String</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>Integer</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>Floating-point number</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>Boolean value</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>Object</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>Array of strings</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>Array of integers</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>Array of floating-point numbers</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>Array of boolean values</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>Object array</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>Nested array</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>Null</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>Unspecified type, for OneOf and AnyOf scenarios</td></tr></tbody></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type RerankConfig struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

// Predefined struct for user
type ResetConversationRequestParams struct {
	// <p>Session ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor experience), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResetConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>Session ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>Session type enumeration values: 0-CONVERSATION_TYPE_UNSPECIFIED (unspecified; means all in list query), 1-CONVERSATION_TYPE_VISITOR (visitor experience), 2-CONVERSATION_TYPE_EVALUATION (evaluation), 5-CONVERSATION_TYPE_API (API integration), 10-CONVERSATION_TYPE_WORKFLOW (workflow debugging), 20-CONVERSATION_TYPE_SHARE (share a link)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Application key</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>Sub-user Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>Primary user Uin.</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Required when Type=CONVERSATION_TYPE_SHARE. Share code</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Required when Type=CONVERSATION_TYPE_API. Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResetConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConversationResponseParams struct {
	// <p>Reset time, unit: ms</p>
	ResetTime *string `json:"ResetTime,omitnil,omitempty" name:"ResetTime"`

	// <p>recordId of the last record before reset</p>
	ResetThroughRecordId *string `json:"ResetThroughRecordId,omitnil,omitempty" name:"ResetThroughRecordId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetConversationResponse struct {
	*tchttp.BaseResponse
	Response *ResetConversationResponseParams `json:"Response"`
}

func (r *ResetConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceConsumption struct {
	// <p>Feature tag. For the PLATFORM scenario, use the PlatformBizType enumeration name; for MODEL/PLUGIN scenarios, leave it empty</p>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// <p>Consumption measurement unit</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token (default)</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count (page number)</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times (count)</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second (second)</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item (item)</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet (sheet)</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character (character)</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number (count)</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second (millisecond)</td></tr></tbody></table>
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// <p>Consumption value</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type ResourcePackageInfo struct {
	// <p>Package type</p><p>Enumeration values:</p><ul><li>1: Free edition</li><li>2: Pro edition</li><li>3: Enterprise edition</li></ul>
	PackageType *uint64 `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// <p>Total resource package quantity</p>
	ResourceTotal *int64 `json:"ResourceTotal,omitnil,omitempty" name:"ResourceTotal"`

	// <p>Resource package usage</p>
	ResourceUsage *float64 `json:"ResourceUsage,omitnil,omitempty" name:"ResourceUsage"`

	// <p>Knowledge base capacity.</p>
	KnowledgeCapacity *float64 `json:"KnowledgeCapacity,omitnil,omitempty" name:"KnowledgeCapacity"`

	// <p>Knowledge base amount</p>
	KnowledgeUsage *float64 `json:"KnowledgeUsage,omitnil,omitempty" name:"KnowledgeUsage"`

	// <p>Resource pack status</p><p>Enumeration values:</p><ul><li>1: Normal</li><li>3: Expired</li><li>4: About to expire</li></ul>
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type ResponseParam struct {
	// <p>Variable description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Parameter name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>OUTPUT_RENDER_REPLACE</td><td>0</td><td>Overwrite (full replacement)</td></tr><tr><td>OUTPUT_RENDER_APPEND</td><td>1</td><td>Incremental append</td></tr></tbody></table>
	RenderMode *int64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// <p>Only useful for OBJECT or ARRAY_OBJECT types</p>
	SubParams []*ResponseParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>String</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>Integer</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>Floating-point number</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>Boolean value</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>Object</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>Array of strings</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>Array of integers</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>Array of floating-point numbers</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>Array of boolean values</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>Object array</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>Nested array</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>Null value</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>Unspecified type, used for OneOf and AnyOf scenarios</td></tr></tbody></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ResumeAppTriggerRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResumeAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResumeAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeAppTriggerResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ResumeAppTriggerResponseParams `json:"Response"`
}

func (r *ResumeAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeWorkflowRunRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow operation instance ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>User input</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>File information</p>
	FileList []*ConversationFileInfo `json:"FileList,omitnil,omitempty" name:"FileList"`
}

type ResumeWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow operation instance ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// <p>User input</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>File information</p>
	FileList []*ConversationFileInfo `json:"FileList,omitnil,omitempty" name:"FileList"`
}

func (r *ResumeWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "WorkflowRunId")
	delete(f, "Query")
	delete(f, "FileList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeWorkflowRunResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *ResumeWorkflowRunResponseParams `json:"Response"`
}

func (r *ResumeWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetrievalOption struct {
	// <p>Timeliness retrieval enhancement configuration</p>
	ExpirationAwareness *ExpirationAwareness `json:"ExpirationAwareness,omitnil,omitempty" name:"ExpirationAwareness"`

	// <p>GraphRAG Configuration</p>
	GraphRag *GraphRAG `json:"GraphRag,omitnil,omitempty" name:"GraphRag"`

	// <p>Spreadsheet enhancement configuration</p>
	TableEnhancement *TableEnhancement `json:"TableEnhancement,omitnil,omitempty" name:"TableEnhancement"`
}

// Predefined struct for user
type RetryReleaseRequestParams struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Release Task ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type RetryReleaseRequest struct {
	*tchttp.BaseRequest
	
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Release Task ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
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
	delete(f, "AppId")
	delete(f, "ReleaseId")
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

type RoleConfig struct {
	// Role description.
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`
}

// Predefined struct for user
type RollbackReleaseRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type RollbackReleaseRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

func (r *RollbackReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackReleaseResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RollbackReleaseResponseParams `json:"Response"`
}

func (r *RollbackReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunAppTriggerNowRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RunAppTriggerNowRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Application trigger ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>Visitor ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *RunAppTriggerNowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunAppTriggerNowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunAppTriggerNowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunAppTriggerNowResponseParams struct {
	// <p>Application trigger instance ID.</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunAppTriggerNowResponse struct {
	*tchttp.BaseResponse
	Response *RunAppTriggerNowResponseParams `json:"Response"`
}

func (r *RunAppTriggerNowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunAppTriggerNowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveWorkflowDebugContextRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application variable list.</p>
	AppVariableValueList []*AppVariableValue `json:"AppVariableValueList,omitnil,omitempty" name:"AppVariableValueList"`

	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Custom user ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type SaveWorkflowDebugContextRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Application variable list.</p>
	AppVariableValueList []*AppVariableValue `json:"AppVariableValueList,omitnil,omitempty" name:"AppVariableValueList"`

	// <p>Session ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Custom user ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *SaveWorkflowDebugContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveWorkflowDebugContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "AppVariableValueList")
	delete(f, "SessionId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveWorkflowDebugContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveWorkflowDebugContextResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveWorkflowDebugContextResponse struct {
	*tchttp.BaseResponse
	Response *SaveWorkflowDebugContextResponseParams `json:"Response"`
}

func (r *SaveWorkflowDebugContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveWorkflowDebugContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveWorkflowToolNodeOutputRequestParams struct {
	// <p>Callback token</p><p>Request Header X-ADP-Callback-Token</p>
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`

	// <p>Callback result</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`
}

type SaveWorkflowToolNodeOutputRequest struct {
	*tchttp.BaseRequest
	
	// <p>Callback token</p><p>Request Header X-ADP-Callback-Token</p>
	CallbackToken *string `json:"CallbackToken,omitnil,omitempty" name:"CallbackToken"`

	// <p>Callback result</p>
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`
}

func (r *SaveWorkflowToolNodeOutputRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveWorkflowToolNodeOutputRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CallbackToken")
	delete(f, "Result")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveWorkflowToolNodeOutputRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveWorkflowToolNodeOutputResponseParams struct {
	// <p>Status code</p><p>Enumeration values:</p><ul><li>0: success</li><li>1000: failure</li></ul>
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>Error description</p>
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveWorkflowToolNodeOutputResponse struct {
	*tchttp.BaseResponse
	Response *SaveWorkflowToolNodeOutputResponseParams `json:"Response"`
}

func (r *SaveWorkflowToolNodeOutputResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveWorkflowToolNodeOutputResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type SearchResourceStatusInfo struct {
	// Search resource status: AVAILABLE(1)=resource available, EXHAUSTED(2)=resource used up. Enumeration values: 1: resource available, 2: resource used up
	ResourceStatus *int64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type SearchResultPayload struct {
	// <p>Additional information of knowledge graph (JSON string)</p>
	GraphData *string `json:"GraphData,omitnil,omitempty" name:"GraphData"`

	// <p>List of hit image URLs</p>
	ImageUrlList []*string `json:"ImageUrlList,omitnil,omitempty" name:"ImageUrlList"`

	// <p>Table additional information (JSON string)</p>
	SheetInfo *string `json:"SheetInfo,omitnil,omitempty" name:"SheetInfo"`
}

type SimilarQuestion struct {
	// <p>Similar question content</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Whether AI-generated</p>
	IsAiGenerated *bool `json:"IsAiGenerated,omitnil,omitempty" name:"IsAiGenerated"`

	// <p>Similar question ID</p>
	SimilarQuestionId *string `json:"SimilarQuestionId,omitnil,omitempty" name:"SimilarQuestionId"`
}

type SimilarQuestionExtra struct {
	// <p>Text content of the similar question</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Similar question ID</p>
	SimilarQuestionId *string `json:"SimilarQuestionId,omitnil,omitempty" name:"SimilarQuestionId"`
}

type SimilarQuestionModifySpec struct {
	// <p>Content of the similar question (required for CREATE and UPDATE)</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>Operation type: 1=create, 2=modify, 3=delete<table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>MODIFY_ACTION_UNKNOWN</td><td>0</td><td></td></tr><tr><td>MODIFY_ACTION_CREATE</td><td>1</td><td>Create</td></tr><tr><td>MODIFY_ACTION_UPDATE</td><td>2</td><td>Modify</td></tr><tr><td>MODIFY_ACTION_DELETE</td><td>3</td><td>Delete</td></tr></tbody></table></p>
	ModifyAction *int64 `json:"ModifyAction,omitnil,omitempty" name:"ModifyAction"`

	// <p>Similar question ID (required for UPDATE and DELETE)</p>
	SimilarQuestionId *string `json:"SimilarQuestionId,omitnil,omitempty" name:"SimilarQuestionId"`
}

type SimilarQuestionStat struct {
	// <p>Number of similar questions</p>
	SimilarQuestionCount *uint64 `json:"SimilarQuestionCount,omitnil,omitempty" name:"SimilarQuestionCount"`

	// <p>Similar question prompt (show a similar question example)</p>
	SimilarQuestionTips *string `json:"SimilarQuestionTips,omitnil,omitempty" name:"SimilarQuestionTips"`
}

type SingleWorkflowConfig struct {
	// <p>Whether to enable async workflow</p>
	AsyncWorkflow *bool `json:"AsyncWorkflow,omitnil,omitempty" name:"AsyncWorkflow"`

	// <p>Status Release status (UNPUBLISHED: to be released PUBLISHING: publishing PUBLISHED: released PUBLISHED_FAIL: release failure; DRAFT: to be debugged)</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Workflow description</p>
	WorkflowDescription *string `json:"WorkflowDescription,omitnil,omitempty" name:"WorkflowDescription"`

	// <p>Workflow Id</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>Workflow name.</p>
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>Whether the workflow is enabled</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SkillAnalysisInfo struct {
	// Safety detection status
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | To be checked |
	// | 1 | detecting |
	// | 2 | Available |
	// | 3 | Unavailable |
	// | 4 | Detection failed |
	AnalysisStatus *int64 `json:"AnalysisStatus,omitnil,omitempty" name:"AnalysisStatus"`

	// Risk description
	RiskDescription *string `json:"RiskDescription,omitnil,omitempty" name:"RiskDescription"`

	// Risk level
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | risk-free |
	// | 1 | Low risk |
	// | 2 | Medium-risk |
	// | 3 | High-risk |
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Security report redirection url
	SecurityReportUrl *string `json:"SecurityReportUrl,omitnil,omitempty" name:"SecurityReportUrl"`
}

type SkillCategory struct {
	// Category identifier
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// Category name
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`
}

type SkillClassification struct {
	// Billing type of a Skill
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Free |
	// | 1 | payment |
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// Built-in Skill source. This parameter is valid only when create_type is Skill_create_type_BUILTIN.
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Placeholder |
	// | 1 | ADP proprietary |
	// | 2 | Tencent proprietary |
	// | 3 | SkillHub |
	// | 99 | Other |
	BuiltinSource *int64 `json:"BuiltinSource,omitnil,omitempty" name:"BuiltinSource"`

	// Skill category
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// Ways to create a Skill
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Placeholder |
	// | 1 | file upload |
	// 2 | Generated by an enterprise-level shared process |
	// | 3 | AIGC generation |
	// | 99 | Built-in Skill |
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// Skill provider type
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Placeholder |
	// | 1 | Official |
	// | 2 | third party |
	// | 3 | Custom |
	// | 4 | Custom enterprise-level sharing |
	ProviderType *int64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// Skill source link
	SourceLink *string `json:"SourceLink,omitnil,omitempty" name:"SourceLink"`
}

type SkillCorpShareConfig struct {
	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>SHARE_SCOPE_TYPE_UNSPECIFIED</td> <td>0</td> <td></td> </tr> <tr> <td>SHARE_SCOPE_TYPE_ALL</td> <td>1</td> <td></td> </tr> <tr> <td>SHARE_SCOPE_TYPE_ACCOUNT</td> <td>2</td> <td></td> </tr> <tr> <td>SHARE_SCOPE_TYPE_SPACE</td> <td>3</td> <td></td> </tr> </tbody> </table> <p>Enumeration values:</p> <ul> <li>0: Unspecified</li> <li>1: Shared across the business</li> <li>3: Shared by space</li> </ul>
	ShareScope *int64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`

	// <p>Share range information. Only spaces are supported. StrId is the space ID, and Name is the space name.</p>
	ShareScopeList []*Identity `json:"ShareScopeList,omitnil,omitempty" name:"ShareScopeList"`
}

type SkillDetail struct {
	// Call condition abstract
	ReferenceSummaryList []*SkillReferenceSummary `json:"ReferenceSummaryList,omitnil,omitempty" name:"ReferenceSummaryList"`

	// Skill abstract
	SkillSummary *SkillSummary `json:"SkillSummary,omitnil,omitempty" name:"SkillSummary"`

	// Version list
	VersionList []*SkillVersion `json:"VersionList,omitnil,omitempty" name:"VersionList"`
}

type SkillNotice struct {
	// Notification level
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Placeholder |
	// | 1 | Success, string literal: "success" |
	// | 2 | Warning, string literal: "warning" |
	// | 3 | Error, string literal: "error" |
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// Copywriting (post-i18n string)
	NoticeContent *string `json:"NoticeContent,omitnil,omitempty" name:"NoticeContent"`

	// Version ID of the Skill that triggers this notice
	TriggerVersionId *string `json:"TriggerVersionId,omitnil,omitempty" name:"TriggerVersionId"`

	// Notification type 
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Placeholder |
	// | 1 | Release failure |
	// | 2 | Shared approval rejected |
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type SkillProfile struct {
	// <p>Creation time (Unix seconds).</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Creator.</p>
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>Skill description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Skill show description</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>Skill display name</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>Skill icon</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Skill name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Update time (Unix seconds).</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>Space</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type SkillReferenceGroup struct {
	// <p>Reference detail list under this type</p>
	ReferenceSummaryList []*SkillReferenceSummary `json:"ReferenceSummaryList,omitnil,omitempty" name:"ReferenceSummaryList"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>SKILL_REF_UNKNOWN</td> <td>0</td> <td>Placeholder</td> </tr> <tr> <td>SKILL_REF_OPENCLAW</td> <td>1</td> <td>openclaw</td> </tr> <tr> <td>SKILL_REF_AGENT</td> <td>2</td> <td>agent</td> </tr> <tr> <td>SKILL_REF_CORP_ASSISTANT</td> <td>3</td> <td>Business assistant</td> </tr> </tbody> </table>
	ReferenceType *int64 `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>Total number of references of this type</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type SkillReferenceSummary struct {
	// <p>Associated ID</p>
	ReferenceId *string `json:"ReferenceId,omitnil,omitempty" name:"ReferenceId"`

	// <p>Association name</p>
	ReferenceName *string `json:"ReferenceName,omitnil,omitempty" name:"ReferenceName"`

	// <p>Association type</p><p>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 | Placeholder |<br>| 1 | ClawPro |<br>| 2 | agent |</p>
	ReferenceType *int64 `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>Space ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Space name</p>
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// <p>Reference instance owner</p>
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// <p>Application identifier</p>
	AppIdentity *AppIdentity `json:"AppIdentity,omitnil,omitempty" name:"AppIdentity"`
}

type SkillShare struct {
	// <p>Approval ID</p>
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// <p>New skill_id associated after sharing</p>
	ShareSkillId *string `json:"ShareSkillId,omitnil,omitempty" name:"ShareSkillId"`

	// <p>Shared version, for example 1.0.0</p>
	ShareVersion *string `json:"ShareVersion,omitnil,omitempty" name:"ShareVersion"`

	// <p>Shared version ID</p>
	ShareVersionId *string `json:"ShareVersionId,omitnil,omitempty" name:"ShareVersionId"`

	// <p>Original skill_id</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Shared state</p><p>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 | Not shared |<br>| 1 | Shared |<br>| 2 | Pending Approval |</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>Business share scope</p>
	CorpShareConfig *SkillCorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`
}

type SkillSummary struct {
	// Classification information
	ClassificationInfo *SkillClassification `json:"ClassificationInfo,omitnil,omitempty" name:"ClassificationInfo"`

	// Current version information
	CurrentVersionInfo *SkillVersion `json:"CurrentVersionInfo,omitnil,omitempty" name:"CurrentVersionInfo"`

	// Whether the current user collects it
	IsFavorite *bool `json:"IsFavorite,omitnil,omitempty" name:"IsFavorite"`

	// Basic information
	Profile *SkillProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Skill ID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// Skill exception notification list
	NoticeList []*SkillNotice `json:"NoticeList,omitnil,omitempty" name:"NoticeList"`

	// List of resource operation permission bits of the current user for this Skill. This parameter is fixed as an empty array for built-in or shared skills.
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`

	// Share information; there may be two entries, one is shared and the other is under review.
	ShareList []*SkillShare `json:"ShareList,omitnil,omitempty" name:"ShareList"`

	// Skill status 
	// 
	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 | Initialization (no published versions, and the latest version is in INITIALIZED/UNRELEASED) |
	// Security detection in progress (no published versions, and the latest version is in AUDITING)
	// | 2 | To be released (no published versions, and the latest version is in PENDING_RELEASE) |
	// | 3 | Released (any RELEASED version exists, absorbent state) |
	SkillStatus *int64 `json:"SkillStatus,omitnil,omitempty" name:"SkillStatus"`
}

type SkillVersion struct {
	// <p>Detection information</p>
	AnalysisInfo *SkillAnalysisInfo `json:"AnalysisInfo,omitnil,omitempty" name:"AnalysisInfo"`

	// <p>Current effective version number.</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>Current effective version ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// <p>Skill version RELEASE process status:</p><pre><code>  - 0 INITIALIZED      initialization (initial version state)  - 1 AUDITING         under review (f_analysis_status ∈ {PENDING, RUNNING})  - 2 PENDING_RELEASE  awaiting RELEASE (low/medium-risk, waiting for user confirmation to publish)  - 3 RELEASED         RELEASED  - 4 UNRELEASED       not RELEASED (HIGH / UNAVAILABLE / FAILED / user give up, including the semantics of historical "not approved")Decoupled from SkillAnalysisStatus: the former is the RELEASE lifecycle from the user perspective, while the latter is the security detection stage.</code></pre>
	VersionStatus *int64 `json:"VersionStatus,omitnil,omitempty" name:"VersionStatus"`

	// <p>md5 information of the Skill package</p>
	SkillMd5 *string `json:"SkillMd5,omitnil,omitempty" name:"SkillMd5"`

	// <p>Version package address</p>
	SkillUrl *string `json:"SkillUrl,omitnil,omitempty" name:"SkillUrl"`

	// <p>Version creation time (Unix seconds).</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>skill md document</p>
	SkillMarkdownUrl *string `json:"SkillMarkdownUrl,omitnil,omitempty" name:"SkillMarkdownUrl"`

	// <p>Version change description</p>
	UpdateDesc *string `json:"UpdateDesc,omitnil,omitempty" name:"UpdateDesc"`

	// <p>Change user.</p>
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`
}

type Sort struct {
	// <p>Sort field name, for example, create_time</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Sorting order. 1: ascending, 2: descending</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>SORT_ORDER_INVALID</td><td>0</td><td>invalid</td></tr><tr><td>SORT_ORDER_ASC</td><td>1</td><td>ascending</td></tr><tr><td>SORT_ORDER_DESC</td><td>2</td><td>descending</td></tr></tbody></table>
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type Space struct {
	// <p>Space id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>Space name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Space description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Space permissions</p>
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`

	// <p>Creation time.</p><p>Parameter format: timestamp in seconds</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Update time</p><p>Parameter format: timestamp</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SpecialStatusInfo struct {
	// Status (0: not in a special status, 1: in a special status). Enumeration value: 1: in a special status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Statistic struct {
	// <p>Time to first token</p>
	FirstTokenCostMs *uint64 `json:"FirstTokenCostMs,omitnil,omitempty" name:"FirstTokenCostMs"`

	// <p>Input token count</p>
	InputTokens *uint64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>Model name</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>Output token count</p>
	OutputTokens *uint64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>Total reasoning time</p>
	TotalCostMs *uint64 `json:"TotalCostMs,omitnil,omitempty" name:"TotalCostMs"`

	// <p>Total input and output tokens</p>
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

// Predefined struct for user
type StopWorkflowRunListRequestParams struct {
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow running instance ID list</p>
	WorkflowRunIdList []*string `json:"WorkflowRunIdList,omitnil,omitempty" name:"WorkflowRunIdList"`
}

type StopWorkflowRunListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Application ID.</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Workflow running instance ID list</p>
	WorkflowRunIdList []*string `json:"WorkflowRunIdList,omitnil,omitempty" name:"WorkflowRunIdList"`
}

func (r *StopWorkflowRunListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkflowRunListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "WorkflowRunIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWorkflowRunListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkflowRunListResponseParams struct {
	// <p>List of instance IDs stopped successfully</p>
	SuccessIdList []*string `json:"SuccessIdList,omitnil,omitempty" name:"SuccessIdList"`

	// <p>List of instance IDs that failed to stop</p>
	FailedIdList []*string `json:"FailedIdList,omitnil,omitempty" name:"FailedIdList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopWorkflowRunListResponse struct {
	*tchttp.BaseResponse
	Response *StopWorkflowRunListResponseParams `json:"Response"`
}

func (r *StopWorkflowRunListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkflowRunListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummaryListSwitch struct {
	// <p>Whether to display metadata</p>
	ShowMetadataEnabled *bool `json:"ShowMetadataEnabled,omitnil,omitempty" name:"ShowMetadataEnabled"`
}

type SupportedFileType struct {
	// <p>Description.</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>File type (for example, txt, pdf, jpg; extension is recommended)</p>
	FileExt *string `json:"FileExt,omitnil,omitempty" name:"FileExt"`

	// <p>File size limit (unit: byte)</p>
	MaxSizeBytes *string `json:"MaxSizeBytes,omitnil,omitempty" name:"MaxSizeBytes"`
}

type SystemVariable struct {
	// Variable Description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Variable name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type TableEnhancement struct {
	// <p>Whether to enable</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ThinkModel struct {
	// Thinking model
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type TimeRange struct {
	// <p>End time in Unix seconds</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>Start time in Unix seconds</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type TimerPushConfig struct {
	// <p>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 |  |<br>| 1 | Do not push |<br>| 2 | WeChat official account |<br>| 3 | WeCom AI chatbot |</p>
	PushChannel *int64 `json:"PushChannel,omitnil,omitempty" name:"PushChannel"`

	// <p>Push session ID</p>
	PushTargetId *string `json:"PushTargetId,omitnil,omitempty" name:"PushTargetId"`

	// <p>Enumeration values:<br>| uint | Description |<br>| --- | --- |<br>| 0 |  |<br>| 1 | User (WeChat official account openid) |<br>| 2 | Group chat (WeCom chatbot chat_id) |</p>
	PushTargetType *int64 `json:"PushTargetType,omitnil,omitempty" name:"PushTargetType"`

	// <p>url for pushing webhook</p>
	PushWebhookUrl *string `json:"PushWebhookUrl,omitnil,omitempty" name:"PushWebhookUrl"`
}

type TimerScheduleConfig struct {
	// Cron configuration.
	Cron *CronSchedule `json:"Cron,omitnil,omitempty" name:"Cron"`

	// Daily trigger
	Daily *DailySchedule `json:"Daily,omitnil,omitempty" name:"Daily"`

	// Fixed interval
	Interval *IntervalSchedule `json:"Interval,omitnil,omitempty" name:"Interval"`

	// Manual only
	ManualOnly *ManualOnlySchedule `json:"ManualOnly,omitnil,omitempty" name:"ManualOnly"`

	// Single time
	Once *OnceSchedule `json:"Once,omitnil,omitempty" name:"Once"`

	// Enumeration values:
	// | uint | Description |
	// | --- | --- |
	// | 0 |  |
	// Manual only
	// | 2 | Daily |
	// | 3 | Weekly |
	// | 4 | By interval |
	// | 5 | One-time |
	// | 6 | Cron |
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// Time zone
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// Triggers at a fixed time weekly.
	Weekly *WeeklySchedule `json:"Weekly,omitnil,omitempty" name:"Weekly"`
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

type Tool struct {
	// <p>Tool billing information</p>
	Billing *ToolBilling `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>Tool call count</p><p>Unit: count</p>
	CallCount *uint64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>Tool description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Tool name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Plug-in ID</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>TOOL_ACCESS_MODE_UNKNOWN</td><td>0</td><td>Not specified</td></tr><tr><td>TOOL_ACCESS_MODE_READ_ONLY</td><td>1</td><td>Read-only</td></tr><tr><td>TOOL_ACCESS_MODE_WRITE_DELETE</td><td>2</td><td>Write/Delete</td></tr></tbody></table>
	ToolAccessMode *int64 `json:"ToolAccessMode,omitnil,omitempty" name:"ToolAccessMode"`

	// <p>Tool configuration info</p>
	ToolConfig *ToolConfig `json:"ToolConfig,omitnil,omitempty" name:"ToolConfig"`

	// <p>Tool ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type ToolBilling struct {
	// <p>Basic Billing Information</p>
	BasicBilling *BasicBilling `json:"BasicBilling,omitnil,omitempty" name:"BasicBilling"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>BILLING_TYPE_FREE</td> <td>0</td> <td>Free</td> </tr> <tr> <td>BILLING_TYPE_LIMITED_FREE</td> <td>1</td> <td>Limited-time free</td> </tr> <tr> <td>BILLING_TYPE_OFFICIAL_PAID</td> <td>2</td> <td>Official charges</td> </tr> <tr> <td>BILLING_TYPE_OFFICIAL_PAID_OLD_FREE</td> <td>3</td> <td>Official charges (fees for new/upgraded users, limited-time free for existing users)</td> </tr> </tbody> </table>
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// <p>Composite type billing information</p>
	ComplexBilling *ComplexBilling `json:"ComplexBilling,omitnil,omitempty" name:"ComplexBilling"`

	// <p>Bidirectional billing information</p>
	DuplexBilling *DuplexBilling `json:"DuplexBilling,omitnil,omitempty" name:"DuplexBilling"`
}

type ToolConfig struct {
	// <p>API tool configuration message.</p>
	ApiToolConfig *ApiToolConfig `json:"ApiToolConfig,omitnil,omitempty" name:"ApiToolConfig"`

	// <p>Application configuration information</p>
	AppToolConfig *AppToolConfig `json:"AppToolConfig,omitnil,omitempty" name:"AppToolConfig"`

	// <p>Code tool configuration message</p>
	CodeToolConfig *CodeToolConfig `json:"CodeToolConfig,omitnil,omitempty" name:"CodeToolConfig"`

	// <p>Configuration information of the MCP tool</p>
	MCPToolConfig *MCPToolConfig `json:"MCPToolConfig,omitnil,omitempty" name:"MCPToolConfig"`
}

type ToolExample struct {
	// <p>Request parameters</p>
	Request *string `json:"Request,omitnil,omitempty" name:"Request"`

	// <p>Response parameter.</p>
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`
}

type ToolSummary struct {
	// <p>Tool Id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type TriggerConfig struct {
	// <p>Timer configuration.</p>
	ScheduledConfig *AppTriggerScheduleConfig `json:"ScheduledConfig,omitnil,omitempty" name:"ScheduledConfig"`

	// <p>Webhook configuration</p>
	WebhookConfig *AppTriggerWebhookConfig `json:"WebhookConfig,omitnil,omitempty" name:"WebhookConfig"`
}

type TriggerStatus struct {
	// <p>Timer status.</p>
	ScheduledStatus *AppTriggerScheduleStatus `json:"ScheduledStatus,omitnil,omitempty" name:"ScheduledStatus"`

	// <table> <tbody> <tr> <td>Enumeration item</td> <td>Enumeration value</td> <td>Description</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_UNSPECIFIED</td> <td>0</td> <td>Unspecified</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_APP</td> <td>1</td> <td>B-end admin</td> </tr> <tr> <td>APP_TRIGGER_SCOPE_USER</td> <td>2</td> <td>C-end visitor</td> </tr> </tbody> </table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>Visitor id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Webhook status</p>
	WebhookStatus *AppTriggerWebhookStatus `json:"WebhookStatus,omitnil,omitempty" name:"WebhookStatus"`
}

// Predefined struct for user
type UnfavoritePluginRequestParams struct {
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type UnfavoritePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>Plugin_id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>Current workspace id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *UnfavoritePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoritePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnfavoritePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoritePluginResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnfavoritePluginResponse struct {
	*tchttp.BaseResponse
	Response *UnfavoritePluginResponseParams `json:"Response"`
}

func (r *UnfavoritePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoritePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoriteSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type UnfavoriteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>Workspace ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *UnfavoriteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoriteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnfavoriteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoriteSkillResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnfavoriteSkillResponse struct {
	*tchttp.BaseResponse
	Response *UnfavoriteSkillResponseParams `json:"Response"`
}

func (r *UnfavoriteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoriteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageDetail struct {
	// <p>Invocation source</p>
	CallSource *CallSource `json:"CallSource,omitnil,omitempty" name:"CallSource"`

	// <p>Measurement ID, used for reconciliation/tracing</p>
	DosageId *string `json:"DosageId,omitnil,omitempty" name:"DosageId"`

	// <p>Call timestamp (Unix seconds).</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>Exclusive to MODEL domain</p>
	Model *ModelUsageDetail `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>Exclusive to the PLUGIN domain</p>
	Plugin *PluginUsageDetail `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// <p>Call chain tracing ID</p>
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// <p>User ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type UsageSummary struct {
	// <p>Exclusive to MODEL domain</p>
	Model *ModelUsageSummary `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>PLATFORM domain exclusive</p>
	Platform *PlatformUsageSummary `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>Exclusive to PLUGIN domain</p>
	Plugin *PluginUsageSummary `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// <p>Source ID; CORP view=space_id (corporate view grouped by space), SPACE view=app_id (uint64 string), APP view=app_id</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>Source name; CORP view=space_name, SPACE view=app_name, APP view=app_name</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>View type, determines the business meaning of SourceId/SourceName</p><table><tbody><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>VIEW_TYPE_UNSPECIFIED</td><td>0</td><td>Unspecified (invalid value, do not pass in requests)</td></tr><tr><td>VIEW_TYPE_CORP</td><td>1</td><td>Corporate view</td></tr><tr><td>VIEW_TYPE_SPACE</td><td>2</td><td>Space view</td></tr><tr><td>VIEW_TYPE_APP</td><td>3</td><td>Application view</td></tr></tbody></table>
	ViewType *int64 `json:"ViewType,omitnil,omitempty" name:"ViewType"`
}

type UserAccessConfig struct {
	// <p>Customer-defined knowledge ID</p>
	CustomerKnowledgeId *string `json:"CustomerKnowledgeId,omitnil,omitempty" name:"CustomerKnowledgeId"`

	// <p>Whether the document is public</p>
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`
}

type Variable struct {
	// <p>Default file name.</p>
	DefaultFileName *string `json:"DefaultFileName,omitnil,omitempty" name:"DefaultFileName"`

	// <p>Default value.</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>Variable description</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Variable module type</p><p>Enumeration values:</p><ul><li>0: API parameter</li><li>1: environment parameter</li><li>2: application parameter</li><li>3: system parameter</li></ul>
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// <p>Variable name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Variable type</p><p>Enumeration values:</p><ul><li>0: string</li><li>1: integer</li><li>2: floating-point number</li><li>3: boolean value</li><li>4: object</li><li>5: array of strings</li><li>6: array of integers</li><li>7: array of floating-point numbers</li><li>8: array of boolean values</li><li>9: object array</li><li>10: file</li><li>11: document</li><li>12: image</li><li>13: audio</li><li>14: video</li><li>15: file array</li><li>16: document array</li><li>17: image array</li><li>18: audio array</li><li>19: video array</li><li>20: array of arrays</li><li>21: key</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Variable ID</p>
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// <p>Whether to enable network policy (only effective for environment variables)</p>
	EnableEndpoints *bool `json:"EnableEndpoints,omitnil,omitempty" name:"EnableEndpoints"`

	// <p>Network policy list (Supported: precise domain name, *.wildcard domain, with optional protocol/port/path prefix)</p>
	EndpointList []*string `json:"EndpointList,omitnil,omitempty" name:"EndpointList"`

	// <p>Whether it is a built-in variable.</p>
	IsBuiltin *bool `json:"IsBuiltin,omitnil,omitempty" name:"IsBuiltin"`

	// <p>Whether it can be injected into the sandbox environment.</p>
	EnableSandbox *bool `json:"EnableSandbox,omitnil,omitempty" name:"EnableSandbox"`
}

type ViewScope struct {
	// <p>View type; enumeration values: VIEW_TYPE_CORP(1) corporate view, VIEW_TYPE_SPACE(2) space view, VIEW_TYPE_APP(3) application view</p><table><tr><td>Enumeration item</td><td>Enumeration value</td><td>Description</td></tr><tr><td>VIEW_TYPE_UNSPECIFIED</td><td>0</td><td>Not specified (invalid value, do not pass in requests)</td></tr><tr><td>VIEW_TYPE_CORP</td><td>1</td><td>Corporate view</td></tr><tr><td>VIEW_TYPE_SPACE</td><td>2</td><td>Space view</td></tr><tr><td>VIEW_TYPE_APP</td><td>3</td><td>Application view</td></tr></table>
	ViewType *int64 `json:"ViewType,omitnil,omitempty" name:"ViewType"`

	// <p>View scope ID; leave empty for VIEW_TYPE_CORP; enter space_id for VIEW_TYPE_SPACE; enter app_id for VIEW_TYPE_APP (decimal string of uint64 snowflake ID)</p>
	ScopeId *string `json:"ScopeId,omitnil,omitempty" name:"ScopeId"`
}

type VoiceConfig struct {
	// Digital human tone key, which needs to align with the public cloud voice type id.
	TimbreKey *string `json:"TimbreKey,omitnil,omitempty" name:"TimbreKey"`

	// Timbre name
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// Public cloud timbre id
	VoiceType *uint64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

type WeeklySchedule struct {
	// Scheduled configuration (week)
	Times []*WeeklyTime `json:"Times,omitnil,omitempty" name:"Times"`
}

type WeeklyTime struct {
	// Time
	TimeOfDay *string `json:"TimeOfDay,omitnil,omitempty" name:"TimeOfDay"`

	// Day of the week
	Weekday *int64 `json:"Weekday,omitnil,omitempty" name:"Weekday"`
}

type WorkflowRun struct {
	// Basic information of the workflow running instance
	WorkflowRunSummary *WorkflowRunSummary `json:"WorkflowRunSummary,omitnil,omitempty" name:"WorkflowRunSummary"`

	// Workflow run instance output information
	WorkflowRunOutput *WorkflowRunOutput `json:"WorkflowRunOutput,omitnil,omitempty" name:"WorkflowRunOutput"`

	// Workflow run instance metadata
	WorkflowRunMetadata *WorkflowRunMetadata `json:"WorkflowRunMetadata,omitnil,omitempty" name:"WorkflowRunMetadata"`
}

type WorkflowRunMetadata struct {
	// User input.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Name of the main model
	MainModelName *string `json:"MainModelName,omitnil,omitempty" name:"MainModelName"`

	// API parameter configuration
	CustomVariableList []*CustomVariable `json:"CustomVariableList,omitnil,omitempty" name:"CustomVariableList"`

	// Flowchart of the workflow
	WorkflowGraph *string `json:"WorkflowGraph,omitnil,omitempty" name:"WorkflowGraph"`
}

type WorkflowRunOutput struct {
	// Workflow output
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// Current reply message
	LatestMessage *ReplyMessage `json:"LatestMessage,omitnil,omitempty" name:"LatestMessage"`
}

type WorkflowRunProfile struct {
	// Application ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// ID of the workflow running instance	
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// Workflow ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type WorkflowRunResult struct {
	// Running state. 1: Queuing; 2: running; 3: run successfully; 4: fail to run; 5: canceled
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// Total number of tokens consumed
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type WorkflowRunSummary struct {
	// <p>Runtime environment. 1: Development domain, 2: Production domain</p>
	Domain *uint64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>Basic information of workflow operation instance</p>
	WorkflowRunProfile *WorkflowRunProfile `json:"WorkflowRunProfile,omitnil,omitempty" name:"WorkflowRunProfile"`

	// <p>Workflow operation instance result</p>
	WorkflowRunResult *WorkflowRunResult `json:"WorkflowRunResult,omitnil,omitempty" name:"WorkflowRunResult"`

	// <p>Time information of workflow running instance</p>
	WorkflowRunTime *WorkflowRunTime `json:"WorkflowRunTime,omitnil,omitempty" name:"WorkflowRunTime"`
}

type WorkflowRunTime struct {
	// Creation time (a millisecond timestamp).
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Start time (timestamp in milliseconds)
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time (timestamp in milliseconds)
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}