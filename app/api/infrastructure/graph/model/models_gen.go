// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/konstellation-io/kdl-server/app/api/entity"
)

type AddMembersInput struct {
	ProjectID string   `json:"projectId"`
	MemberIds []string `json:"memberIds"`
}

type AddUserInput struct {
	Email       string             `json:"email"`
	AccessLevel entity.AccessLevel `json:"accessLevel"`
}

type APIToken struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CreationDate string `json:"creationDate"`
	LastUsedDate string `json:"lastUsedDate"`
	Token        string `json:"token"`
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

type KnowledgeGraph struct {
	Items []*KnowledgeGraphItem `json:"items"`
}

type KnowledgeGraphItem struct {
	ID          string                `json:"id"`
	Category    KnowledgeGraphItemCat `json:"category"`
	Title       string                `json:"title"`
	Abstract    string                `json:"abstract"`
	Authors     []string              `json:"authors"`
	Score       float64               `json:"score"`
	Date        string                `json:"date"`
	URL         string                `json:"url"`
	IsStarred   bool                  `json:"isStarred"`
	IsDiscarded bool                  `json:"isDiscarded"`
	ExternalID  *string               `json:"externalId"`
	Framework   *string               `json:"framework"`
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
	Type entity.RepositoryType `json:"type"`
	URL  string                `json:"url"`
}

type SSHKey struct {
	Public       string  `json:"public"`
	Private      string  `json:"private"`
	CreationDate string  `json:"creationDate"`
	LastActivity *string `json:"lastActivity"`
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

type User struct {
	ID           string             `json:"id"`
	Email        string             `json:"email"`
	CreationDate string             `json:"creationDate"`
	AccessLevel  entity.AccessLevel `json:"accessLevel"`
	LastActivity *string            `json:"lastActivity"`
	APITokens    []*APIToken        `json:"apiTokens"`
}

type KnowledgeGraphItemCat string

const (
	KnowledgeGraphItemCatPaper KnowledgeGraphItemCat = "Paper"
	KnowledgeGraphItemCatCode  KnowledgeGraphItemCat = "Code"
)

var AllKnowledgeGraphItemCat = []KnowledgeGraphItemCat{
	KnowledgeGraphItemCatPaper,
	KnowledgeGraphItemCatCode,
}

func (e KnowledgeGraphItemCat) IsValid() bool {
	switch e {
	case KnowledgeGraphItemCatPaper, KnowledgeGraphItemCatCode:
		return true
	}
	return false
}

func (e KnowledgeGraphItemCat) String() string {
	return string(e)
}

func (e *KnowledgeGraphItemCat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KnowledgeGraphItemCat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid KnowledgeGraphItemCat", str)
	}
	return nil
}

func (e KnowledgeGraphItemCat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
