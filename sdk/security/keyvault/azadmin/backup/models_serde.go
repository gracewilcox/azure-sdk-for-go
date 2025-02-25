// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package backup

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// MarshalJSON implements the json.Marshaller interface for type FullBackupOperation.
func (f FullBackupOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureStorageBlobContainerUri", f.AzureStorageBlobContainerURI)
	populateTimeUnix(objectMap, "endTime", f.EndTime)
	populate(objectMap, "error", f.Error)
	populate(objectMap, "jobId", f.JobID)
	populateTimeUnix(objectMap, "startTime", f.StartTime)
	populate(objectMap, "status", f.Status)
	populate(objectMap, "statusDetails", f.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FullBackupOperation.
func (f *FullBackupOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureStorageBlobContainerUri":
			err = unpopulate(val, "AzureStorageBlobContainerURI", &f.AzureStorageBlobContainerURI)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &f.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &f.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &f.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &f.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &f.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &f.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreOperation.
func (r RestoreOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeUnix(objectMap, "endTime", r.EndTime)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "jobId", r.JobID)
	populateTimeUnix(objectMap, "startTime", r.StartTime)
	populate(objectMap, "status", r.Status)
	populate(objectMap, "statusDetails", r.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreOperation.
func (r *RestoreOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &r.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &r.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &r.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &r.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &r.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &r.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RestoreOperationParameters.
func (r RestoreOperationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "folderToRestore", r.FolderToRestore)
	populate(objectMap, "sasTokenParameters", r.SASTokenParameters)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreOperationParameters.
func (r *RestoreOperationParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "folderToRestore":
			err = unpopulate(val, "FolderToRestore", &r.FolderToRestore)
			delete(rawMsg, key)
		case "sasTokenParameters":
			err = unpopulate(val, "SASTokenParameters", &r.SASTokenParameters)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SASTokenParameters.
func (s SASTokenParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "storageResourceUri", s.StorageResourceURI)
	populate(objectMap, "token", s.Token)
	populate(objectMap, "useManagedIdentity", s.UseManagedIdentity)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SASTokenParameters.
func (s *SASTokenParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "storageResourceUri":
			err = unpopulate(val, "StorageResourceURI", &s.StorageResourceURI)
			delete(rawMsg, key)
		case "token":
			err = unpopulate(val, "Token", &s.Token)
			delete(rawMsg, key)
		case "useManagedIdentity":
			err = unpopulate(val, "UseManagedIdentity", &s.UseManagedIdentity)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SelectiveKeyRestoreOperation.
func (s SelectiveKeyRestoreOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeUnix(objectMap, "endTime", s.EndTime)
	populate(objectMap, "error", s.Error)
	populate(objectMap, "jobId", s.JobID)
	populateTimeUnix(objectMap, "startTime", s.StartTime)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SelectiveKeyRestoreOperation.
func (s *SelectiveKeyRestoreOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateTimeUnix(val, "EndTime", &s.EndTime)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &s.Error)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &s.JobID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeUnix(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &s.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SelectiveKeyRestoreOperationParameters.
func (s SelectiveKeyRestoreOperationParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "folder", s.Folder)
	populate(objectMap, "sasTokenParameters", s.SASTokenParameters)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SelectiveKeyRestoreOperationParameters.
func (s *SelectiveKeyRestoreOperationParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "folder":
			err = unpopulate(val, "Folder", &s.Folder)
			delete(rawMsg, key)
		case "sasTokenParameters":
			err = unpopulate(val, "SASTokenParameters", &s.SASTokenParameters)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
