// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package backup

// BeginFullBackupOptions contains the optional parameters for the Client.BeginFullBackup method.
type BeginFullBackupOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BeginFullRestoreOptions contains the optional parameters for the Client.BeginFullRestore method.
type BeginFullRestoreOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BeginPreFullBackupOptions contains the optional parameters for the Client.BeginPreFullBackup method.
type BeginPreFullBackupOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BeginPreFullRestoreOptions contains the optional parameters for the Client.BeginPreFullRestore method.
type BeginPreFullRestoreOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BeginSelectiveKeyRestoreOptions contains the optional parameters for the Client.BeginSelectiveKeyRestore method.
type BeginSelectiveKeyRestoreOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}
