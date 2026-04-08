package smsapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const contactsApiPath = "/contacts"

type ContactsApi struct {
	client *Client
}

type ContactListFilters struct {
	PaginationFilters

	Query        string   `url:"q,omitempty"`
	OrderBy      string   `url:"order_by,omitempty"`
	PhoneNumber  []string `url:"phone_number,omitempty"`
	Email        []string `url:"email,omitempty"`
	FirstName    []string `url:"first_name,omitempty"`
	LastName     []string `url:"last_name,omitempty"`
	GroupId      []string `url:"group_id,omitempty"`
	Gender       string   `url:"gender,omitempty"`
	BirthdayDate []string `url:"birthday_date,omitempty"`
}

type Contact struct {
	Id           string `json:"id,omitempty" url:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty" url:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty" url:"last_name,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty" url:"phone_number,omitempty"`
	Email        string `json:"email,omitempty" url:"email,omitempty"`
	Gender       string `json:"gender,omitempty" url:"gender,omitempty"`
	BirthdayDate string `json:"birthday_date,omitempty" url:"birthday_date,omitempty"`
	Description  string `json:"description,omitempty" url:"description,omitempty"`
	City         string `json:"city,omitempty" url:"city,omitempty"`
	Source       string `json:"source,omitempty" url:"source,omitempty"`
	DateCreated  string `json:"date_created,omitempty" url:"date_created,omitempty"`
	DateUpdated  string `json:"date_updated,omitempty" url:"date_updated,omitempty"`
}

type ContactCollectionResponse struct {
	CollectionMeta
	Collection []*Contact `json:"collection"`
}

type ContactsCollectionIterator struct {
	i *PageIterator
}

func (b *ContactsCollectionIterator) Next() (*ContactCollectionResponse, error) {
	c := new(ContactCollectionResponse)

	err := b.i.Next(c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (contactsApi *ContactsApi) GetContacts(ctx context.Context, filters *ContactListFilters) (*ContactCollectionResponse, error) {
	var result = new(ContactCollectionResponse)

	uri, _ := addQueryParams("/contacts", filters)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) GetContactsPageIterator(ctx context.Context, filters *ContactListFilters) *ContactsCollectionIterator {
	i := NewPageIterator(contactsApi.client, ctx, contactsApiPath, filters)
	ci := &ContactsCollectionIterator{i}

	return ci
}

func (contactsApi *ContactsApi) CreateContact(ctx context.Context, contact *Contact) (*Contact, error) {
	var result = new(Contact)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPost, contactsApiPath, result, contact)

	return result, err
}

func (contactsApi *ContactsApi) DeleteAllContacts(ctx context.Context) error {
	return contactsApi.client.Delete(ctx, contactsApiPath)
}

func (contactsApi *ContactsApi) GetContact(ctx context.Context, id string) (*Contact, error) {
	uri := fmt.Sprintf("%s/%s", contactsApiPath, id)

	var result = new(Contact)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) UpdateContact(ctx context.Context, id string, contact *Contact) (*Contact, error) {
	uri := fmt.Sprintf("/contacts/%s", id)

	var result = new(Contact)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPut, uri, result, contact)

	return result, err
}

func (contactsApi *ContactsApi) DeleteContact(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/contacts/%s", id)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

type ContactGroup struct {
	Id            string                     `json:"id,omitempty" url:"id,omitempty"`
	Name          string                     `json:"name,omitempty" url:"name,omitempty"`
	Description   string                     `json:"description,omitempty" url:"description,omitempty"`
	ContactsCount int                        `json:"contacts_count,omitempty" url:"contacts_count,omitempty"`
	DateCreated   string                     `json:"date_created,omitempty" url:"date_created,omitempty"`
	DateUpdated   string                     `json:"date_updated,omitempty" url:"date_updated,omitempty"`
	CreatedBy     string                     `json:"created_by,omitempty" url:"created_by,omitempty"`
	Idx           string                     `json:"idx,omitempty" url:"idx,omitempty"`
	Permissions   []*ContactGroupPermissions `json:"permissions,omitempty" url:"permissions,omitempty"`
}

type ContactGroupPermissions struct {
	GroupId  string `json:"group_id,omitempty"`
	Username string `json:"username,omitempty"`
	Write    bool   `json:"write,omitempty"`
	Read     bool   `json:"read,omitempty"`
	Send     bool   `json:"send,omitempty"`
}

type ContactGroupsCollectionResponse struct {
	Size       int             `json:"size"`
	Collection []*ContactGroup `json:"collection"`
}

func (contactsApi *ContactsApi) GetContactGroups(ctx context.Context, id string) (*ContactGroupsCollectionResponse, error) {
	uri := fmt.Sprintf("/contacts/%s/groups", id)

	var result = new(ContactGroupsCollectionResponse)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) GetContactGroup(ctx context.Context, contactId string, groupId string) (*ContactGroup, error) {
	uri := fmt.Sprintf("/contacts/%s/groups/%s", contactId, groupId)

	var result = new(ContactGroup)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) BindContactToGroup(ctx context.Context, contactId string, groupId string) (*ContactGroupsCollectionResponse, error) {
	uri := fmt.Sprintf("/contacts/%s/groups/%s", contactId, groupId)

	var result = new(ContactGroupsCollectionResponse)

	err := contactsApi.client.Put(ctx, uri, result, nil)

	return result, err
}

func (contactsApi *ContactsApi) UnbindContactFromGroup(ctx context.Context, contactId string, groupId string) error {
	uri := fmt.Sprintf("/contacts/%s/groups/%s", contactId, groupId)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

func (contactsApi *ContactsApi) GetGroups(ctx context.Context) (*ContactGroupsCollectionResponse, error) {
	var result = new(ContactGroupsCollectionResponse)

	err := contactsApi.client.Get(ctx, "contacts/groups", result)

	return result, err
}

func (contactsApi *ContactsApi) CreateGroup(ctx context.Context, group *ContactGroup) (*ContactGroup, error) {
	var result = new(ContactGroup)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPost, "contacts/groups", result, group)

	return result, err
}

func (contactsApi *ContactsApi) DeleteAllGroup(ctx context.Context) error {
	err := contactsApi.client.Delete(ctx, "contacts/groups")

	return err
}

func (contactsApi *ContactsApi) UpdateGroup(ctx context.Context, groupId string, group *ContactGroup) (*ContactGroup, error) {
	uri := fmt.Sprintf("contacts/groups/%s", groupId)

	var result = new(ContactGroup)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPut, uri, result, group)

	return result, err
}

func (contactsApi *ContactsApi) GetGroup(ctx context.Context, groupId string) (*ContactGroup, error) {
	uri := fmt.Sprintf("contacts/groups/%s", groupId)

	var result = new(ContactGroup)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) DeleteGroup(ctx context.Context, groupId string) error {
	uri := fmt.Sprintf("contacts/groups/%s", groupId)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

func (contactsApi *ContactsApi) MoveContactsToGroup(ctx context.Context, groupId string, filters *ContactListFilters) error {
	uri := fmt.Sprintf("contacts/groups/%s/members", groupId)

	return contactsApi.client.Urlencoded(ctx, http.MethodPut, uri, nil, filters)
}

func (contactsApi *ContactsApi) AddContactsToGroup(ctx context.Context, groupId string, filters *ContactListFilters) error {
	uri := fmt.Sprintf("contacts/groups/%s/members", groupId)

	return contactsApi.client.Urlencoded(ctx, http.MethodPost, uri, nil, filters)
}

func (contactsApi *ContactsApi) RemoveContactsFromGroup(ctx context.Context, groupId string, filters *ContactListFilters) error {
	uri := fmt.Sprintf("contacts/groups/%s/members", groupId)

	return contactsApi.client.Urlencoded(ctx, http.MethodDelete, uri, nil, filters)
}

func (contactsApi *ContactsApi) AddContactToGroup(ctx context.Context, groupId, contactId string) (*Contact, error) {
	uri := fmt.Sprintf("contacts/groups/%s/members/%s", groupId, contactId)

	var result = new(Contact)

	err := contactsApi.client.Put(ctx, uri, result, nil)

	return result, err
}

func (contactsApi *ContactsApi) GetContactFromGroup(ctx context.Context, groupId, contactId string) (*Contact, error) {
	uri := fmt.Sprintf("contacts/groups/%s/members/%s", groupId, contactId)

	var result = new(Contact)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) RemoveContactFromGroup(ctx context.Context, groupId, contactId string) error {
	uri := fmt.Sprintf("contacts/groups/%s/members/%s", groupId, contactId)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

type ContactGroupPermissionsCollectionResponse struct {
	Size       int                       `json:"size"`
	Collection []ContactGroupPermissions `json:"collection"`
}

func (contactsApi *ContactsApi) GetGroupPermissions(ctx context.Context, groupId string) (*ContactGroupPermissionsCollectionResponse, error) {
	uri := fmt.Sprintf("contacts/groups/%s/permissions", groupId)

	var result = new(ContactGroupPermissionsCollectionResponse)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) AddGroupPermissions(ctx context.Context, groupId string, permissions *ContactGroupPermissions) (*ContactGroupPermissions, error) {
	uri := fmt.Sprintf("contacts/groups/%s/permissions", groupId)

	var result = new(ContactGroupPermissions)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPost, uri, result, permissions)

	return result, err
}

func (contactsApi *ContactsApi) GetUserGroupPermissions(ctx context.Context, groupId, username string) (*ContactGroupPermissions, error) {
	uri := fmt.Sprintf("contacts/groups/%s/permissions/%s", groupId, username)

	var result = new(ContactGroupPermissions)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

func (contactsApi *ContactsApi) AddUserGroupPermissions(ctx context.Context, groupId, username string, permissions *ContactGroupPermissions) (*ContactGroupPermissions, error) {
	uri := fmt.Sprintf("contacts/groups/%s/permissions/%s", groupId, username)

	var result = new(ContactGroupPermissions)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPut, uri, result, permissions)

	return result, err
}

func (contactsApi *ContactsApi) RemoveUserGroupPermissions(ctx context.Context, groupId, username string) error {
	uri := fmt.Sprintf("contacts/groups/%s/permissions/%s", groupId, username)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

type CustomField struct {
	Id   string `json:"id,omitempty" url:"id,omitempty"`
	Name string `json:"name,omitempty" url:"name,omitempty"`
	Type string `json:"type,omitempty" url:"type,omitempty"`
}

type CustomFieldsCollectionResponse struct {
	Size       int           `json:"size"`
	Collection []CustomField `json:"collection"`
}

func (contactsApi *ContactsApi) GetCustomFields(ctx context.Context) (*CustomFieldsCollectionResponse, error) {
	var result = new(CustomFieldsCollectionResponse)

	err := contactsApi.client.Get(ctx, "contacts/fields", result)

	return result, err
}

func (contactsApi *ContactsApi) CreateCustomField(ctx context.Context, name, type_ string) (*CustomField, error) {
	field := &CustomField{
		Name: name,
		Type: type_,
	}

	var result = new(CustomField)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPost, "contacts/fields", result, field)

	return result, err
}

func (contactsApi *ContactsApi) UpdateCustomField(ctx context.Context, fieldId, name string) (*CustomField, error) {
	field := &CustomField{
		Name: name,
	}

	uri := fmt.Sprintf("contacts/fields/%s", fieldId)

	var result = new(CustomField)

	err := contactsApi.client.Urlencoded(ctx, http.MethodPut, uri, result, field)

	return result, err
}

func (contactsApi *ContactsApi) DeleteCustomField(ctx context.Context, fieldId string) error {
	uri := fmt.Sprintf("contacts/fields/%s", fieldId)

	err := contactsApi.client.Delete(ctx, uri)

	return err
}

// AssignContactToGroups assigns a contact to multiple groups at once.
// The request body is urlencoded with numeric keys mapping to group ids
// (e.g. 0=<groupId1>&1=<groupId2>).
func (contactsApi *ContactsApi) AssignContactToGroups(ctx context.Context, contactId string, groupIds []string) (*ContactGroupsCollectionResponse, error) {
	uri := fmt.Sprintf("/contacts/%s/groups", contactId)

	values := url.Values{}
	for i, groupId := range groupIds {
		values.Set(strconv.Itoa(i), groupId)
	}

	var result = new(ContactGroupsCollectionResponse)

	err := contactsApi.client.PostRaw(ctx, uri, strings.NewReader(values.Encode()), ContentTypeXFormUrlencoded, result)

	return result, err
}

func (contactsApi *ContactsApi) CleanTrash(ctx context.Context) error {
	return contactsApi.client.Delete(ctx, "/contacts/trash")
}

func (contactsApi *ContactsApi) RestoreTrash(ctx context.Context) error {
	return contactsApi.client.Urlencoded(ctx, http.MethodPut, "/contacts/trash/restore", nil, nil)
}

type FieldOption struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type FieldOptionsCollectionResponse struct {
	Size       int            `json:"size"`
	Collection []*FieldOption `json:"collection"`
}

func (contactsApi *ContactsApi) GetCustomFieldOptions(ctx context.Context, fieldId string) (*FieldOptionsCollectionResponse, error) {
	uri := fmt.Sprintf("contacts/fields/%s/options", fieldId)

	var result = new(FieldOptionsCollectionResponse)

	err := contactsApi.client.Get(ctx, uri, result)

	return result, err
}

type AvailableField struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    string   `json:"type,omitempty"`
	BuiltIn bool     `json:"built_in,omitempty"`
	Options []string `json:"options,omitempty"`
}

func (contactsApi *ContactsApi) GetAvailableFields(ctx context.Context) ([]*AvailableField, error) {
	var result []*AvailableField

	err := contactsApi.client.Get(ctx, "contacts/fields/available", &result)

	return result, err
}
