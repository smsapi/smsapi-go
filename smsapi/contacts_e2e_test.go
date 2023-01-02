// +build e2e

package smsapi

import (
	"log"
	"testing"
)

var (
	contactId string
	groupId   string
	fieldId   string
)

func TestCreateGroup(t *testing.T) {
	group := &ContactGroup{
		Name:        "demo",
		Description: "go-smsapi e2e",
	}

	ctx, cancel := createCtx()
	defer cancel()

	group, err := client.Contacts.CreateGroup(ctx, group)

	if err != nil {
		log.Fatal(err)
	}

	groupId = group.Id
}

func TestRemoveGroup(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Contacts.DeleteGroup(ctx, groupId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateContact(t *testing.T) {
	contact := &Contact{
		FirstName:   "demo",
		PhoneNumber: phoneNumber,
		Description: "go-smsapi e2e",
	}

	ctx, cancel := createCtx()
	defer cancel()

	contact, err := client.Contacts.CreateContact(ctx, contact)

	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveContact(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Contacts.DeleteContact(ctx, contactId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateCustomField(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	field, err := client.Contacts.CreateCustomField(ctx, "field-a", "text")

	if err != nil {
		log.Fatal(err)
	}

	fieldId = field.Id
}

func TestUpdateCustomField(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Contacts.UpdateCustomField(ctx, fieldId, "field-a-a")

	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveCustomField(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Contacts.DeleteCustomField(ctx, fieldId)

	if err != nil {
		log.Fatal(err)
	}
}
