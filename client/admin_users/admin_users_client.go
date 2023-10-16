// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientClient.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/client.gotmpl

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdminCreateUser(params *AdminCreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminCreateUserOK, error)

	AdminDeleteUser(params *AdminDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminDeleteUserOK, error)

	AdminDisableUser(params *AdminDisableUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminDisableUserOK, error)

	AdminEnableUser(params *AdminEnableUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminEnableUserOK, error)

	AdminGetUserAuthTokens(params *AdminGetUserAuthTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminGetUserAuthTokensOK, error)

	AdminLogoutUser(params *AdminLogoutUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminLogoutUserOK, error)

	AdminRevokeUserAuthToken(params *AdminRevokeUserAuthTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminRevokeUserAuthTokenOK, error)

	AdminUpdateUserPassword(params *AdminUpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminUpdateUserPasswordOK, error)

	AdminUpdateUserPermissions(params *AdminUpdateUserPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminUpdateUserPermissionsOK, error)

	GetUserQuota(params *GetUserQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserQuotaOK, error)

	UpdateUserQuota(params *UpdateUserQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserQuotaOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AdminCreateUser creates new user

	If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users:create`.

Note that OrgId is an optional parameter that can be used to assign a new user to a different organization when `auto_assign_org` is set to `true`.
*/
func (a *Client) AdminCreateUser(params *AdminCreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminCreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminCreateUser",
		Method:             "POST",
		PathPattern:        "/admin/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminCreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminCreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminCreateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminDeleteUser deletes global user

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users:delete` and scope `global.users:*`.
*/
func (a *Client) AdminDeleteUser(params *AdminDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminDeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminDeleteUser",
		Method:             "DELETE",
		PathPattern:        "/admin/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminDeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminDeleteUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminDeleteUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminDisableUser disables user

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users:disable` and scope `global.users:1` (userIDScope).
*/
func (a *Client) AdminDisableUser(params *AdminDisableUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminDisableUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminDisableUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminDisableUser",
		Method:             "POST",
		PathPattern:        "/admin/users/{user_id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminDisableUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminDisableUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminDisableUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminEnableUser enables user

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users:enable` and scope `global.users:1` (userIDScope).
*/
func (a *Client) AdminEnableUser(params *AdminEnableUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminEnableUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminEnableUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminEnableUser",
		Method:             "POST",
		PathPattern:        "/admin/users/{user_id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminEnableUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminEnableUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminEnableUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminGetUserAuthTokens returns a list of all auth tokens devices that the user currently have logged in from

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.authtoken:list` and scope `global.users:*`.
*/
func (a *Client) AdminGetUserAuthTokens(params *AdminGetUserAuthTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminGetUserAuthTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetUserAuthTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminGetUserAuthTokens",
		Method:             "GET",
		PathPattern:        "/admin/users/{user_id}/auth-tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminGetUserAuthTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminGetUserAuthTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminGetUserAuthTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminLogoutUser logouts user revokes all auth tokens devices for the user user of issued auth tokens devices will no longer be logged in and will be required to authenticate again upon next activity

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.logout` and scope `global.users:*`.
*/
func (a *Client) AdminLogoutUser(params *AdminLogoutUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminLogoutUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminLogoutUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminLogoutUser",
		Method:             "POST",
		PathPattern:        "/admin/users/{user_id}/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminLogoutUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminLogoutUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminLogoutUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AdminRevokeUserAuthToken revokes auth token for user

	Revokes the given auth token (device) for the user. User of issued auth token (device) will no longer be logged in and will be required to authenticate again upon next activity.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.authtoken:update` and scope `global.users:*`.
*/
func (a *Client) AdminRevokeUserAuthToken(params *AdminRevokeUserAuthTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminRevokeUserAuthTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminRevokeUserAuthTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminRevokeUserAuthToken",
		Method:             "POST",
		PathPattern:        "/admin/users/{user_id}/revoke-auth-token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminRevokeUserAuthTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminRevokeUserAuthTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminRevokeUserAuthToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AdminUpdateUserPassword sets password for user

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.password:update` and scope `global.users:*`.
*/
func (a *Client) AdminUpdateUserPassword(params *AdminUpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminUpdateUserPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateUserPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminUpdateUserPassword",
		Method:             "PUT",
		PathPattern:        "/admin/users/{user_id}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminUpdateUserPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminUpdateUserPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminUpdateUserPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AdminUpdateUserPermissions sets permissions for user

	Only works with Basic Authentication (username and password). See introduction for an explanation.

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.permissions:update` and scope `global.users:*`.
*/
func (a *Client) AdminUpdateUserPermissions(params *AdminUpdateUserPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AdminUpdateUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateUserPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminUpdateUserPermissions",
		Method:             "PUT",
		PathPattern:        "/admin/users/{user_id}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminUpdateUserPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminUpdateUserPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for adminUpdateUserPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserQuota fetches user quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.quotas:list` and scope `global.users:1` (userIDScope).
*/
func (a *Client) GetUserQuota(params *GetUserQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserQuota",
		Method:             "GET",
		PathPattern:        "/admin/users/{user_id}/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserQuota updates user quota

If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `users.quotas:update` and scope `global.users:1` (userIDScope).
*/
func (a *Client) UpdateUserQuota(params *UpdateUserQuotaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserQuotaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUserQuota",
		Method:             "PUT",
		PathPattern:        "/admin/users/{user_id}/quotas/{quota_target}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserQuotaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserQuota: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}