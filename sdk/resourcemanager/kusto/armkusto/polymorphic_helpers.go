//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import "encoding/json"

func unmarshalDataConnectionClassification(rawMsg json.RawMessage) (DataConnectionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataConnectionClassification
	switch m["kind"] {
	case string(DataConnectionKindCosmosDb):
		b = &CosmosDbDataConnection{}
	case string(DataConnectionKindEventGrid):
		b = &EventGridDataConnection{}
	case string(DataConnectionKindEventHub):
		b = &EventHubDataConnection{}
	case string(DataConnectionKindIotHub):
		b = &IotHubDataConnection{}
	default:
		b = &DataConnection{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataConnectionClassificationArray(rawMsg json.RawMessage) ([]DataConnectionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataConnectionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataConnectionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDatabaseClassification(rawMsg json.RawMessage) (DatabaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatabaseClassification
	switch m["kind"] {
	case string(KindReadOnlyFollowing):
		b = &ReadOnlyFollowingDatabase{}
	case string(KindReadWrite):
		b = &ReadWriteDatabase{}
	default:
		b = &Database{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatabaseClassificationArray(rawMsg json.RawMessage) ([]DatabaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DatabaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDatabaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
