// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package job

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/2016-11-01/job"
)

const (
	DefaultAdlaJobDNSSuffix = original.DefaultAdlaJobDNSSuffix
)

type CompileMode = original.CompileMode

const (
	Full      CompileMode = original.Full
	Semantic  CompileMode = original.Semantic
	SingleBox CompileMode = original.SingleBox
)

type ResourceType = original.ResourceType

const (
	JobManagerResource             ResourceType = original.JobManagerResource
	JobManagerResourceInUserFolder ResourceType = original.JobManagerResourceInUserFolder
	StatisticsResource             ResourceType = original.StatisticsResource
	StatisticsResourceInUserFolder ResourceType = original.StatisticsResourceInUserFolder
	VertexResource                 ResourceType = original.VertexResource
	VertexResourceInUserFolder     ResourceType = original.VertexResourceInUserFolder
)

type Result = original.Result

const (
	Cancelled Result = original.Cancelled
	Failed    Result = original.Failed
	None      Result = original.None
	Succeeded Result = original.Succeeded
)

type SeverityTypes = original.SeverityTypes

const (
	Deprecated    SeverityTypes = original.Deprecated
	Error         SeverityTypes = original.Error
	Info          SeverityTypes = original.Info
	SevereWarning SeverityTypes = original.SevereWarning
	UserWarning   SeverityTypes = original.UserWarning
	Warning       SeverityTypes = original.Warning
)

type State = original.State

const (
	StateAccepted           State = original.StateAccepted
	StateCompiling          State = original.StateCompiling
	StateEnded              State = original.StateEnded
	StateNew                State = original.StateNew
	StatePaused             State = original.StatePaused
	StateQueued             State = original.StateQueued
	StateRunning            State = original.StateRunning
	StateScheduling         State = original.StateScheduling
	StateStarting           State = original.StateStarting
	StateWaitingForCapacity State = original.StateWaitingForCapacity
)

type Type = original.Type

const (
	TypeHive          Type = original.TypeHive
	TypeJobProperties Type = original.TypeJobProperties
	TypeUSQL          Type = original.TypeUSQL
)

type TypeBasicCreateJobProperties = original.TypeBasicCreateJobProperties

const (
	TypeBasicCreateJobPropertiesTypeCreateJobProperties TypeBasicCreateJobProperties = original.TypeBasicCreateJobPropertiesTypeCreateJobProperties
	TypeBasicCreateJobPropertiesTypeUSQL                TypeBasicCreateJobProperties = original.TypeBasicCreateJobPropertiesTypeUSQL
)

type TypeEnum = original.TypeEnum

const (
	Hive TypeEnum = original.Hive
	USQL TypeEnum = original.USQL
)

type BaseClient = original.BaseClient
type BaseJobParameters = original.BaseJobParameters
type BasicCreateJobProperties = original.BasicCreateJobProperties
type BasicProperties = original.BasicProperties
type BuildJobParameters = original.BuildJobParameters
type Client = original.Client
type CreateJobParameters = original.CreateJobParameters
type CreateJobProperties = original.CreateJobProperties
type CreateUSQLJobProperties = original.CreateUSQLJobProperties
type DataPath = original.DataPath
type Diagnostics = original.Diagnostics
type ErrorDetails = original.ErrorDetails
type HiveJobProperties = original.HiveJobProperties
type InfoListResult = original.InfoListResult
type InfoListResultIterator = original.InfoListResultIterator
type InfoListResultPage = original.InfoListResultPage
type Information = original.Information
type InformationBasic = original.InformationBasic
type InnerError = original.InnerError
type PipelineClient = original.PipelineClient
type PipelineInformation = original.PipelineInformation
type PipelineInformationListResult = original.PipelineInformationListResult
type PipelineInformationListResultIterator = original.PipelineInformationListResultIterator
type PipelineInformationListResultPage = original.PipelineInformationListResultPage
type PipelineRunInformation = original.PipelineRunInformation
type Properties = original.Properties
type RecurrenceClient = original.RecurrenceClient
type RecurrenceInformation = original.RecurrenceInformation
type RecurrenceInformationListResult = original.RecurrenceInformationListResult
type RecurrenceInformationListResultIterator = original.RecurrenceInformationListResultIterator
type RecurrenceInformationListResultPage = original.RecurrenceInformationListResultPage
type RelationshipProperties = original.RelationshipProperties
type Resource = original.Resource
type StateAuditRecord = original.StateAuditRecord
type Statistics = original.Statistics
type StatisticsVertexStage = original.StatisticsVertexStage
type USQLJobProperties = original.USQLJobProperties

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewInfoListResultIterator(page InfoListResultPage) InfoListResultIterator {
	return original.NewInfoListResultIterator(page)
}
func NewInfoListResultPage(cur InfoListResult, getNextPage func(context.Context, InfoListResult) (InfoListResult, error)) InfoListResultPage {
	return original.NewInfoListResultPage(cur, getNextPage)
}
func NewPipelineClient() PipelineClient {
	return original.NewPipelineClient()
}
func NewPipelineInformationListResultIterator(page PipelineInformationListResultPage) PipelineInformationListResultIterator {
	return original.NewPipelineInformationListResultIterator(page)
}
func NewPipelineInformationListResultPage(cur PipelineInformationListResult, getNextPage func(context.Context, PipelineInformationListResult) (PipelineInformationListResult, error)) PipelineInformationListResultPage {
	return original.NewPipelineInformationListResultPage(cur, getNextPage)
}
func NewRecurrenceClient() RecurrenceClient {
	return original.NewRecurrenceClient()
}
func NewRecurrenceInformationListResultIterator(page RecurrenceInformationListResultPage) RecurrenceInformationListResultIterator {
	return original.NewRecurrenceInformationListResultIterator(page)
}
func NewRecurrenceInformationListResultPage(cur RecurrenceInformationListResult, getNextPage func(context.Context, RecurrenceInformationListResult) (RecurrenceInformationListResult, error)) RecurrenceInformationListResultPage {
	return original.NewRecurrenceInformationListResultPage(cur, getNextPage)
}
func NewWithoutDefaults(adlaJobDNSSuffix string) BaseClient {
	return original.NewWithoutDefaults(adlaJobDNSSuffix)
}
func PossibleCompileModeValues() []CompileMode {
	return original.PossibleCompileModeValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleResultValues() []Result {
	return original.PossibleResultValues()
}
func PossibleSeverityTypesValues() []SeverityTypes {
	return original.PossibleSeverityTypesValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTypeBasicCreateJobPropertiesValues() []TypeBasicCreateJobProperties {
	return original.PossibleTypeBasicCreateJobPropertiesValues()
}
func PossibleTypeEnumValues() []TypeEnum {
	return original.PossibleTypeEnumValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
