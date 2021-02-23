// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/konstellation-io/kdl-server/app/api/entity"
)

type AddMembersInput struct {
	ProjectID string   `json:"projectId"`
	MemberIds []string `json:"memberIds"`
}

type AddUserInput struct {
	Email       string             `json:"email"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type APITokenInput struct {
	UserID string  `json:"userId"`
	Name   *string `json:"name"`
}

type CreateProjectInput struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Repository  *RepositoryInput `json:"repository"`
}

type QualityProjectDesc struct {
	Quality *float64 `json:"quality"`
}

type RemoveAPITokenInput struct {
	APITokenID string `json:"apiTokenId"`
}

type RemoveMemberInput struct {
	ProjectID string `json:"projectId"`
	MemberID  string `json:"memberId"`
}

type RemoveUsersInput struct {
	UserIds []string `json:"userIds"`
}

type RepositoryInput struct {
	Type             entity.RepositoryType `json:"type"`
	InternalRepoName *string               `json:"internalRepoName"`
	ExternalRepoURL  *string               `json:"externalRepoUrl"`
}

type SetActiveUserToolsInput struct {
	Active bool `json:"active"`
}

type SetBoolFieldInput struct {
	ID    string `json:"id"`
	Value bool   `json:"value"`
}

type UpdateAccessLevelInput struct {
	UserIds     []string           `json:"userIds"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type UpdateMemberInput struct {
	ProjectID   string             `json:"projectId"`
	MemberID    string             `json:"memberId"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type UpdateProjectInput struct {
	ID         string                        `json:"id"`
	Name       *string                       `json:"name"`
	Repository *UpdateProjectRepositoryInput `json:"repository"`
}

type UpdateProjectRepositoryInput struct {
	URL string `json:"url"`
}
