//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BatchQueryRequest.
func (b BatchQueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "body", b.Body)
	populate(objectMap, "headers", b.Headers)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "method", b.Method)
	populate(objectMap, "path", b.Path)
	populate(objectMap, "workspace", b.Workspace)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryRequest.
func (b *BatchQueryRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "body":
				err = unpopulate(val, "Body", &b.Body)
				delete(rawMsg, key)
		case "headers":
				err = unpopulate(val, "Headers", &b.Headers)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &b.ID)
				delete(rawMsg, key)
		case "method":
				err = unpopulate(val, "Method", &b.Method)
				delete(rawMsg, key)
		case "path":
				err = unpopulate(val, "Path", &b.Path)
				delete(rawMsg, key)
		case "workspace":
				err = unpopulate(val, "Workspace", &b.Workspace)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchQueryResponse.
func (b BatchQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "body", b.Body)
	populate(objectMap, "headers", b.Headers)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "status", b.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryResponse.
func (b *BatchQueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "body":
				err = unpopulate(val, "Body", &b.Body)
				delete(rawMsg, key)
		case "headers":
				err = unpopulate(val, "Headers", &b.Headers)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &b.ID)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &b.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchQueryResults.
func (b BatchQueryResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", b.Error)
	populate(objectMap, "render", &b.Render)
	populate(objectMap, "statistics", &b.Statistics)
	populate(objectMap, "tables", b.Tables)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryResults.
func (b *BatchQueryResults) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &b.Error)
				delete(rawMsg, key)
		case "render":
				err = unpopulate(val, "Render", &b.Render)
				delete(rawMsg, key)
		case "statistics":
				err = unpopulate(val, "Statistics", &b.Statistics)
				delete(rawMsg, key)
		case "tables":
				err = unpopulate(val, "Tables", &b.Tables)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchRequest.
func (b BatchRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "requests", b.Requests)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchRequest.
func (b *BatchRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "requests":
				err = unpopulate(val, "Requests", &b.Requests)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchResponse.
func (b BatchResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "responses", b.Responses)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchResponse.
func (b *BatchResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "responses":
				err = unpopulate(val, "Responses", &b.Responses)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Body.
func (b Body) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "query", b.Query)
	populate(objectMap, "timespan", b.Timespan)
	populate(objectMap, "workspaces", b.Workspaces)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Body.
func (b *Body) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "query":
				err = unpopulate(val, "Query", &b.Query)
				delete(rawMsg, key)
		case "timespan":
				err = unpopulate(val, "Timespan", &b.Timespan)
				delete(rawMsg, key)
		case "workspaces":
				err = unpopulate(val, "Workspaces", &b.Workspaces)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Column.
func (c Column) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Column.
func (c *Column) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
				err = unpopulate(val, "Name", &c.Name)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &c.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", &e.AdditionalProperties)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "resources", e.Resources)
	populate(objectMap, "target", e.Target)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
				err = unpopulate(val, "AdditionalProperties", &e.AdditionalProperties)
				delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
				delete(rawMsg, key)
		case "resources":
				err = unpopulate(val, "Resources", &e.Resources)
				delete(rawMsg, key)
		case "target":
				err = unpopulate(val, "Target", &e.Target)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &e.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorInfo.
func (e ErrorInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", &e.AdditionalProperties)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "innererror", e.Innererror)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorInfo.
func (e *ErrorInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
				err = unpopulate(val, "AdditionalProperties", &e.AdditionalProperties)
				delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "details":
				err = unpopulate(val, "Details", &e.Details)
				delete(rawMsg, key)
		case "innererror":
				err = unpopulate(val, "Innererror", &e.Innererror)
				delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &e.Error)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseAutoGenerated.
func (e ErrorResponseAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponseAutoGenerated.
func (e *ErrorResponseAutoGenerated) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocalizableString.
func (l LocalizableString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "localizedValue", l.LocalizedValue)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocalizableString.
func (l *LocalizableString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "localizedValue":
				err = unpopulate(val, "LocalizedValue", &l.LocalizedValue)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &l.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataValue.
func (m MetadataValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", m.Name)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataValue.
func (m *MetadataValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
				err = unpopulate(val, "Name", &m.Name)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &m.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Metric.
func (m Metric) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "errorCode", m.ErrorCode)
	populate(objectMap, "errorMessage", m.ErrorMessage)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "timeseries", m.Timeseries)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Metric.
func (m *Metric) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayDescription":
				err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
				delete(rawMsg, key)
		case "errorCode":
				err = unpopulate(val, "ErrorCode", &m.ErrorCode)
				delete(rawMsg, key)
		case "errorMessage":
				err = unpopulate(val, "ErrorMessage", &m.ErrorMessage)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &m.ID)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &m.Name)
				delete(rawMsg, key)
		case "timeseries":
				err = unpopulate(val, "Timeseries", &m.Timeseries)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &m.Type)
				delete(rawMsg, key)
		case "unit":
				err = unpopulate(val, "Unit", &m.Unit)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricAvailability.
func (m MetricAvailability) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "retention", m.Retention)
	populate(objectMap, "timeGrain", m.TimeGrain)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricAvailability.
func (m *MetricAvailability) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "retention":
				err = unpopulate(val, "Retention", &m.Retention)
				delete(rawMsg, key)
		case "timeGrain":
				err = unpopulate(val, "TimeGrain", &m.TimeGrain)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricDefinition.
func (m MetricDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", m.Category)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "isDimensionRequired", m.IsDimensionRequired)
	populate(objectMap, "metricAvailabilities", m.MetricAvailabilities)
	populate(objectMap, "metricClass", m.MetricClass)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "namespace", m.Namespace)
	populate(objectMap, "primaryAggregationType", m.PrimaryAggregationType)
	populate(objectMap, "resourceId", m.ResourceID)
	populate(objectMap, "supportedAggregationTypes", m.SupportedAggregationTypes)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricDefinition.
func (m *MetricDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
				err = unpopulate(val, "Category", &m.Category)
				delete(rawMsg, key)
		case "dimensions":
				err = unpopulate(val, "Dimensions", &m.Dimensions)
				delete(rawMsg, key)
		case "displayDescription":
				err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &m.ID)
				delete(rawMsg, key)
		case "isDimensionRequired":
				err = unpopulate(val, "IsDimensionRequired", &m.IsDimensionRequired)
				delete(rawMsg, key)
		case "metricAvailabilities":
				err = unpopulate(val, "MetricAvailabilities", &m.MetricAvailabilities)
				delete(rawMsg, key)
		case "metricClass":
				err = unpopulate(val, "MetricClass", &m.MetricClass)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &m.Name)
				delete(rawMsg, key)
		case "namespace":
				err = unpopulate(val, "Namespace", &m.Namespace)
				delete(rawMsg, key)
		case "primaryAggregationType":
				err = unpopulate(val, "PrimaryAggregationType", &m.PrimaryAggregationType)
				delete(rawMsg, key)
		case "resourceId":
				err = unpopulate(val, "ResourceID", &m.ResourceID)
				delete(rawMsg, key)
		case "supportedAggregationTypes":
				err = unpopulate(val, "SupportedAggregationTypes", &m.SupportedAggregationTypes)
				delete(rawMsg, key)
		case "unit":
				err = unpopulate(val, "Unit", &m.Unit)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricDefinitionCollection.
func (m MetricDefinitionCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricDefinitionCollection.
func (m *MetricDefinitionCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
				err = unpopulate(val, "Value", &m.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricNamespace.
func (m MetricNamespace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "classification", m.Classification)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricNamespace.
func (m *MetricNamespace) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "classification":
				err = unpopulate(val, "Classification", &m.Classification)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &m.ID)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &m.Name)
				delete(rawMsg, key)
		case "properties":
				err = unpopulate(val, "Properties", &m.Properties)
				delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &m.Type)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricNamespaceCollection.
func (m MetricNamespaceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricNamespaceCollection.
func (m *MetricNamespaceCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
				err = unpopulate(val, "Value", &m.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricNamespaceName.
func (m MetricNamespaceName) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metricNamespaceName", m.MetricNamespaceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricNamespaceName.
func (m *MetricNamespaceName) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "metricNamespaceName":
				err = unpopulate(val, "MetricNamespaceName", &m.MetricNamespaceName)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricValue.
func (m MetricValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "average", m.Average)
	populate(objectMap, "count", m.Count)
	populate(objectMap, "maximum", m.Maximum)
	populate(objectMap, "minimum", m.Minimum)
	populateTimeRFC3339(objectMap, "timeStamp", m.TimeStamp)
	populate(objectMap, "total", m.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricValue.
func (m *MetricValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "average":
				err = unpopulate(val, "Average", &m.Average)
				delete(rawMsg, key)
		case "count":
				err = unpopulate(val, "Count", &m.Count)
				delete(rawMsg, key)
		case "maximum":
				err = unpopulate(val, "Maximum", &m.Maximum)
				delete(rawMsg, key)
		case "minimum":
				err = unpopulate(val, "Minimum", &m.Minimum)
				delete(rawMsg, key)
		case "timeStamp":
				err = unpopulateTimeRFC3339(val, "TimeStamp", &m.TimeStamp)
				delete(rawMsg, key)
		case "total":
				err = unpopulate(val, "Total", &m.Total)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Response.
func (r Response) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "cost", r.Cost)
	populate(objectMap, "interval", r.Interval)
	populate(objectMap, "namespace", r.Namespace)
	populate(objectMap, "resourceregion", r.Resourceregion)
	populate(objectMap, "timespan", r.Timespan)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Response.
func (r *Response) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cost":
				err = unpopulate(val, "Cost", &r.Cost)
				delete(rawMsg, key)
		case "interval":
				err = unpopulate(val, "Interval", &r.Interval)
				delete(rawMsg, key)
		case "namespace":
				err = unpopulate(val, "Namespace", &r.Namespace)
				delete(rawMsg, key)
		case "resourceregion":
				err = unpopulate(val, "Resourceregion", &r.Resourceregion)
				delete(rawMsg, key)
		case "timespan":
				err = unpopulate(val, "Timespan", &r.Timespan)
				delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &r.Value)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Results.
func (r Results) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", r.Error)
	populate(objectMap, "render", &r.Render)
	populate(objectMap, "statistics", &r.Statistics)
	populate(objectMap, "tables", r.Tables)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Results.
func (r *Results) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &r.Error)
				delete(rawMsg, key)
		case "render":
				err = unpopulate(val, "Render", &r.Render)
				delete(rawMsg, key)
		case "statistics":
				err = unpopulate(val, "Statistics", &r.Statistics)
				delete(rawMsg, key)
		case "tables":
				err = unpopulate(val, "Tables", &r.Tables)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Table.
func (t Table) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "columns", t.Columns)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "rows", t.Rows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Table.
func (t *Table) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "columns":
				err = unpopulate(val, "Columns", &t.Columns)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &t.Name)
				delete(rawMsg, key)
		case "rows":
				err = unpopulate(val, "Rows", &t.Rows)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TimeSeriesElement.
func (t TimeSeriesElement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "data", t.Data)
	populate(objectMap, "metadatavalues", t.Metadatavalues)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TimeSeriesElement.
func (t *TimeSeriesElement) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
				err = unpopulate(val, "Data", &t.Data)
				delete(rawMsg, key)
		case "metadatavalues":
				err = unpopulate(val, "Metadatavalues", &t.Metadatavalues)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

