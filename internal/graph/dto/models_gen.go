// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package dto

import (
	"fmt"
	"io"
	"model"
	"strconv"
	"time"
)

type TeamRequestExportJSON struct {
	V    string `json:"v"`
	Auth model.Auth `json:"auth"`
	Body model.Body `json:"body"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Params []model.Param `json:"params"`
	Headers []model.Header `json:"headers"`
	Endpoint         string `json:"endpoint"`
	TestScript       string `json:"testScript"`
	PreRequestScript string `json:"preRequestScript"`
}

type TeamCollectionExportJSON struct {
	Name string `json:"name"`
	Folders []TeamCollectionExportJSON  `json:"folders"`
	Requests []TeamRequestExportJSON  `json:"requests"`
	Data *string `json:"data"`
}

type TeamCollectionImportDataJSON struct {
	Headers []model.Header `json:"headers"`
	Auth model.Auth `json:"auth"`
}
type TeamCollectionImportJSON struct {
	Name string `json:"name"`
	Folders []TeamCollectionImportJSON  `json:"folders"`
	Requests []TeamRequestExportJSON  `json:"requests"`
	Headers []model.Header `json:"headers"`
	Auth model.Auth `json:"auth"`
	Data TeamCollectionImportDataJSON `json:"data"`
}

type UserRequestExportJSON struct {
	ID string `json:"id"`
	V  string `json:"v"`
	Auth model.Auth `json:"auth"`
	Body model.Body `json:"body"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Params []model.Param `json:"params"`
	Headers []model.Header `json:"headers"`
	Endpoint         string `json:"endpoint"`
	TestScript       string `json:"testScript"`
	PreRequestScript string `json:"preRequestScript"`
}

type UserCollectionExportJSON struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Folders []UserCollectionExportJSON  `json:"folders"`
	Requests []UserRequestExportJSON  `json:"requests"`
	Data *string `json:"data"`
}

type Admin struct {
	// UID of the user
	UID string `json:"uid"`
	// Name of the user (if fetched)
	DisplayName *string `json:"displayName,omitempty"`
	// Email of the user
	Email *string `json:"email,omitempty"`
	// URL to the profile photo of the user (if fetched)
	PhotoURL *string `json:"photoURL,omitempty"`
	// Date when the user account was created
	CreatedOn time.Time `json:"createdOn"`
	// Returns a list of all admin users in infra
	Admins []*model.User `json:"admins"`
	// Returns a user info by UID
	UserInfo *model.User `json:"userInfo"`
	// Returns a list of all the users in infra
	AllUsers []*model.User `json:"allUsers"`
	// Returns a list of all the invited users
	InvitedUsers []*model.InvitedUser `json:"invitedUsers"`
	// Returns a list of all the teams in the infra
	AllTeams []*model.Team `json:"allTeams"`
	// Returns a team info by ID when requested by Admin
	TeamInfo *model.Team `json:"teamInfo"`
	// Return count of all the members in a team
	MembersCountInTeam int `json:"membersCountInTeam"`
	// Return count of all the stored collections in a team
	CollectionCountInTeam int `json:"collectionCountInTeam"`
	// Return count of all the stored requests in a team
	RequestCountInTeam int `json:"requestCountInTeam"`
	// Return count of all the stored environments in a team
	EnvironmentCountInTeam int `json:"environmentCountInTeam"`
	// Return all the pending invitations in a team
	PendingInvitationCountInTeam []*model.TeamInvitation `json:"pendingInvitationCountInTeam"`
	// Return total number of Users in organization
	UsersCount int `json:"usersCount"`
	// Return total number of Teams in organization
	TeamsCount int `json:"teamsCount"`
	// Return total number of Team Collections in organization
	TeamCollectionsCount int `json:"teamCollectionsCount"`
	// Return total number of Team Requests in organization
	TeamRequestsCount int `json:"teamRequestsCount"`
}

type CollectionReorderData struct {
	// Team Collection being moved
	Collection *model.TeamCollection `json:"collection"`
	// Team Collection succeeding the collection being moved in its new position
	NextCollection *model.TeamCollection `json:"nextCollection,omitempty"`
}

type CreateTeamRequestInput struct {
	// ID of the team the collection belongs to
	TeamID string `json:"teamID"`
	// JSON string representing the request data
	Request string `json:"request"`
	// Displayed title of the request
	Title string `json:"title"`
}

type EnableAndDisableSSOArgs struct {
	// Auth Provider
	Provider AuthProvider `json:"provider"`
	// Auth Provider Status
	Status ServiceStatus `json:"status"`
}

type Infra struct {
	// Admin who executed the action
	ExecutedBy *Admin `json:"executedBy"`
	// Returns a list of all admin users in infra
	Admins []*model.User `json:"admins"`
	// Returns a user info by UID
	UserInfo *model.User `json:"userInfo"`
	// Returns a list of all the users in infra
	AllUsers []*model.User `json:"allUsers"`
	// Returns a list of all the invited users
	InvitedUsers []*model.InvitedUser `json:"invitedUsers"`
	// Returns a list of all the teams in the infra
	AllTeams []*model.Team `json:"allTeams"`
	// Returns a team info by ID when requested by Admin
	TeamInfo *model.Team `json:"teamInfo"`
	// Return count of all the members in a team
	MembersCountInTeam int64 `json:"membersCountInTeam"`
	// Return count of all the stored collections in a team
	CollectionCountInTeam int64 `json:"collectionCountInTeam"`
	// Return count of all the stored requests in a team
	RequestCountInTeam int64 `json:"requestCountInTeam"`
	// Return count of all the stored environments in a team
	EnvironmentCountInTeam int64 `json:"environmentCountInTeam"`
	// Return all the pending invitations in a team
	PendingInvitationCountInTeam []*model.TeamInvitation `json:"pendingInvitationCountInTeam"`
	// Return total number of Users in organization
	UsersCount int64 `json:"usersCount"`
	// Return total number of Teams in organization
	TeamsCount int64 `json:"teamsCount"`
	// Return total number of Team Collections in organization
	TeamCollectionsCount int64 `json:"teamCollectionsCount"`
	// Return total number of Team Requests in organization
	TeamRequestsCount int64 `json:"teamRequestsCount"`
	// Returns a list of all the shortcodes in the infra
	AllShortcodes []*ShortcodeWithUserEmail `json:"allShortcodes"`
}

type InfraConfigArgs struct {
	// Infra Config Name
	Name InfraConfigEnum `json:"name"`
	// Infra Config Value
	Value string `json:"value"`
}

type Mutation struct {
}

type Query struct {
}

type RequestReorderData struct {
	// Team Request being moved
	Request *model.TeamRequest `json:"request"`
	// Team Request succeeding the request being moved in its new position
	NextRequest *model.TeamRequest `json:"nextRequest,omitempty"`
}

type ShortcodeCreator struct {
	// Uid of user who created the shortcode
	UID string `json:"uid"`
	// Email of user who created the shortcode
	Email string `json:"email"`
}

type ShortcodeWithUserEmail struct {
	// The 12 digit alphanumeric code
	ID string `json:"id"`
	// JSON string representing the request data
	Request string `json:"request"`
	// JSON string representing the properties for an embed
	Properties *string `json:"properties,omitempty"`
	// Timestamp of when the Shortcode was created
	CreatedOn time.Time `json:"createdOn"`
	// Details of user who created the shortcode
	Creator *ShortcodeCreator `json:"creator,omitempty"`
}

type Subscription struct {
}

type UpdateTeamRequestInput struct {
	// JSON string representing the request data
	Request *string `json:"request,omitempty"`
	// Displayed title of the request
	Title *string `json:"title,omitempty"`
}

type UserCollectionExportJSONData struct {
	// Stringified contents of the collection
	ExportedCollection string `json:"exportedCollection"`
	// Type of the user collection
	CollectionType model.ReqType `json:"collectionType"`
}

type UserCollectionRemovedData struct {
	// ID of User Collection being removed
	ID string `json:"id"`
	// Type of the user collection
	Type model.ReqType `json:"type"`
}

type UserCollectionReorderData struct {
	// User Collection being moved
	UserCollection *model.UserCollection `json:"userCollection"`
	// User Collection succeeding the collection being moved in its new position
	NextUserCollection *model.UserCollection `json:"nextUserCollection,omitempty"`
}

type UserHistoryDeletedManyData struct {
	// Number of user histories deleted
	Count int `json:"count"`
	// Type of the request in the history
	ReqType model.ReqType `json:"reqType"`
}

type UserRequestReorderData struct {
	// User request being moved
	Request *model.UserRequest `json:"request"`
	// User request succeeding the request being moved in its new position
	NextRequest *model.UserRequest `json:"nextRequest,omitempty"`
}

type AuthProvider string

const (
	AuthProviderGoogle    AuthProvider = "GOOGLE"
	AuthProviderGithub    AuthProvider = "GITHUB"
	AuthProviderMicrosoft AuthProvider = "MICROSOFT"
	AuthProviderEmail     AuthProvider = "EMAIL"
)

var AllAuthProvider = []AuthProvider{
	AuthProviderGoogle,
	AuthProviderGithub,
	AuthProviderMicrosoft,
	AuthProviderEmail,
}

func (e AuthProvider) IsValid() bool {
	switch e {
	case AuthProviderGoogle, AuthProviderGithub, AuthProviderMicrosoft, AuthProviderEmail:
		return true
	}
	return false
}

func (e AuthProvider) String() string {
	return string(e)
}

func (e *AuthProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthProvider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthProvider", str)
	}
	return nil
}

func (e AuthProvider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InfraConfigEnum string

const (
	InfraConfigEnumMailerSMTPURL         InfraConfigEnum = "MAILER_SMTP_URL"
	InfraConfigEnumMailerAddressFrom     InfraConfigEnum = "MAILER_ADDRESS_FROM"
	InfraConfigEnumGoogleClientID        InfraConfigEnum = "GOOGLE_CLIENT_ID"
	InfraConfigEnumGoogleClientSecret    InfraConfigEnum = "GOOGLE_CLIENT_SECRET"
	InfraConfigEnumGithubClientID        InfraConfigEnum = "GITHUB_CLIENT_ID"
	InfraConfigEnumGithubClientSecret    InfraConfigEnum = "GITHUB_CLIENT_SECRET"
	InfraConfigEnumMicrosoftClientID     InfraConfigEnum = "MICROSOFT_CLIENT_ID"
	InfraConfigEnumMicrosoftClientSecret InfraConfigEnum = "MICROSOFT_CLIENT_SECRET"
)

var AllInfraConfigEnum = []InfraConfigEnum{
	InfraConfigEnumMailerSMTPURL,
	InfraConfigEnumMailerAddressFrom,
	InfraConfigEnumGoogleClientID,
	InfraConfigEnumGoogleClientSecret,
	InfraConfigEnumGithubClientID,
	InfraConfigEnumGithubClientSecret,
	InfraConfigEnumMicrosoftClientID,
	InfraConfigEnumMicrosoftClientSecret,
}

func (e InfraConfigEnum) IsValid() bool {
	switch e {
	case InfraConfigEnumMailerSMTPURL, InfraConfigEnumMailerAddressFrom, InfraConfigEnumGoogleClientID, InfraConfigEnumGoogleClientSecret, InfraConfigEnumGithubClientID, InfraConfigEnumGithubClientSecret, InfraConfigEnumMicrosoftClientID, InfraConfigEnumMicrosoftClientSecret:
		return true
	}
	return false
}

func (e InfraConfigEnum) String() string {
	return string(e)
}

func (e *InfraConfigEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InfraConfigEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InfraConfigEnum", str)
	}
	return nil
}

func (e InfraConfigEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServiceStatus string

const (
	ServiceStatusEnable  ServiceStatus = "ENABLE"
	ServiceStatusDisable ServiceStatus = "DISABLE"
)

var AllServiceStatus = []ServiceStatus{
	ServiceStatusEnable,
	ServiceStatusDisable,
}

func (e ServiceStatus) IsValid() bool {
	switch e {
	case ServiceStatusEnable, ServiceStatusDisable:
		return true
	}
	return false
}

func (e ServiceStatus) String() string {
	return string(e)
}

func (e *ServiceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceStatus", str)
	}
	return nil
}

func (e ServiceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
