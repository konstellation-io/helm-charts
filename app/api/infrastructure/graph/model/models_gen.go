// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/konstellation-io/kdl-server/app/api/entity"
)

type AddMembersInput struct {
	ProjectID string   `json:"projectId"`
	UserIds   []string `json:"userIds"`
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
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Repository  *RepositoryInput `json:"repository"`
}

type ExternalRepositoryInput struct {
	URL        string                      `json:"url"`
	Username   string                      `json:"username"`
	Credential string                      `json:"credential"`
	AuthMethod entity.RepositoryAuthMethod `json:"authMethod"`
}

type QualityProjectDesc struct {
	Quality int `json:"quality"`
}

type RemoveAPITokenInput struct {
	APITokenID string `json:"apiTokenId"`
}

type RemoveMembersInput struct {
	ProjectID string   `json:"projectId"`
	UserIds   []string `json:"userIds"`
}

type RemoveUsersInput struct {
	UserIds []string `json:"userIds"`
}

type RepositoryInput struct {
	Type     entity.RepositoryType    `json:"type"`
	External *ExternalRepositoryInput `json:"external"`
}

type SetActiveUserToolsInput struct {
	Active    bool    `json:"active"`
	RuntimeID *string `json:"runtimeId"`
}

type SetBoolFieldInput struct {
	ID    string `json:"id"`
	Value bool   `json:"value"`
}

type SyncUsersResponse struct {
	Msg string `json:"msg"`
}

type Topic struct {
	Name      string  `json:"name"`
	Relevance float64 `json:"relevance"`
}

type UpdateAccessLevelInput struct {
	UserIds     []string           `json:"userIds"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type UpdateMembersInput struct {
	ProjectID   string             `json:"projectId"`
	UserIds     []string           `json:"userIds"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type UpdateProjectInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Archived    *bool   `json:"archived"`
}
