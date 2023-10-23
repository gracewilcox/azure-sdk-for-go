//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package rbac

// Permission - Role definition permissions.
type Permission struct {
	// Action permissions that are granted.
	Actions []*string

	// Data action permissions that are granted.
	DataActions []*DataAction

	// Action permissions that are excluded but not denied. They may be granted by other role definitions assigned to a principal.
	NotActions []*string

	// Data action permissions that are excluded but not denied. They may be granted by other role definitions assigned to a principal.
	NotDataActions []*DataAction
}

// RoleAssignment - Role Assignments
type RoleAssignment struct {
	// Role assignment properties.
	Properties *RoleAssignmentPropertiesWithScope

	// READ-ONLY; The role assignment ID.
	ID *string

	// READ-ONLY; The role assignment name.
	Name *string

	// READ-ONLY; The role assignment type.
	Type *string
}

// RoleAssignmentCreateParameters - Role assignment create parameters.
type RoleAssignmentCreateParameters struct {
	// REQUIRED; Role assignment properties.
	Properties *RoleAssignmentProperties
}

// RoleAssignmentListResult - Role assignment list operation result.
type RoleAssignmentListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// Role assignment list.
	Value []*RoleAssignment
}

// RoleAssignmentProperties - Role assignment properties.
type RoleAssignmentProperties struct {
	// REQUIRED; The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user,
	// service principal, or security group.
	PrincipalID *string

	// REQUIRED; The role definition ID used in the role assignment.
	RoleDefinitionID *string
}

// RoleAssignmentPropertiesWithScope - Role assignment properties with scope.
type RoleAssignmentPropertiesWithScope struct {
	// The principal ID.
	PrincipalID *string

	// The role definition ID.
	RoleDefinitionID *string

	// The role scope.
	Scope *RoleScope
}

// RoleDefinition - Role definition.
type RoleDefinition struct {
	// Role definition properties.
	Properties *RoleDefinitionProperties

	// READ-ONLY; The role definition ID.
	ID *string

	// READ-ONLY; The role definition name.
	Name *string

	// READ-ONLY; The role definition type.
	Type *RoleDefinitionType
}

// RoleDefinitionCreateParameters - Role definition create parameters.
type RoleDefinitionCreateParameters struct {
	// REQUIRED; Role definition properties.
	Properties *RoleDefinitionProperties
}

// RoleDefinitionListResult - Role definition list operation result.
type RoleDefinitionListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// Role definition list.
	Value []*RoleDefinition
}

// RoleDefinitionProperties - Role definition properties.
type RoleDefinitionProperties struct {
	// Role definition assignable scopes.
	AssignableScopes []*RoleScope

	// The role definition description.
	Description *string

	// Role definition permissions.
	Permissions []*Permission

	// The role name.
	RoleName *string

	// The role type.
	RoleType *RoleType
}
