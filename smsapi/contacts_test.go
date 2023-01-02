// +build unit

package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetContacts(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/list_contacts.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Contacts.GetContacts(ctx, nil)

	expected := createExpectedContactsCollection()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestContactsPageIterator(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/list_contacts.json"))

		assertRequestMethod(t, r, "GET")
	})

	iterator := client.Contacts.GetContactsPageIterator(ctx, nil)

	page, _ := iterator.Next()

	expectedPage := createExpectedContactsCollection()

	if !reflect.DeepEqual(page, expectedPage) {
		t.Errorf("Given: %+v Expected: %+v", page, expectedPage)
	}
}

func TestGetContact(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/contact.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Contacts.GetContact(ctx, "1")

	expected := createContact()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateContact(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	contact := &Contact{
		FirstName:   "test",
		LastName:    "demo",
		PhoneNumber: "111222333",
	}

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/contact.json"))

		assertRequestMethod(t, r, "POST")
	})

	result, _ := client.Contacts.CreateContact(ctx, contact)

	expectedResponse := createContact()

	if !reflect.DeepEqual(result, expectedResponse) {
		t.Errorf("Given: %+v Expected: %+v", result, contact)
	}
}

func TestUpdateContact(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	contact := &Contact{
		FirstName: "test",
		LastName:  "demo",
	}

	mux.HandleFunc("/contacts/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/contact.json"))

		assertRequestMethod(t, r, "PUT")
		assertRequestBody(t, r, new(Contact), contact)
	})

	result, _ := client.Contacts.UpdateContact(ctx, "1", contact)

	expectedResponse := createContact()

	if !reflect.DeepEqual(expectedResponse, result) {
		t.Errorf("Given: %+v Expected: %+v", result, expectedResponse)
	}
}

func TestDeleteContact(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Contacts.DeleteContact(ctx, "1")

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAllContacts(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Contacts.DeleteAllContacts(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestGetContactGroups(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/groups.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Contacts.GetContactGroups(ctx, "1")

	expected := &ContactGroupsCollectionResponse{
		Size: 1,
		Collection: []*ContactGroup{
			createGroup(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetContactGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1/groups/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/group.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Contacts.GetContactGroup(ctx, "1", "1")

	expected := createGroup()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestBindContactToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1/groups/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/groups.json"))

		assertRequestMethod(t, r, "PUT")
	})

	result, _ := client.Contacts.BindContactToGroup(ctx, "1", "1")

	expected := &ContactGroupsCollectionResponse{
		Size: 1,
		Collection: []*ContactGroup{
			createGroup(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestUnbindContactFromGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/1/groups/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Contacts.UnbindContactFromGroup(ctx, "1", "1")

	if err != nil {
		t.Error(err)
	}
}

func TestGetAllGroups(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/groups.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Contacts.GetGroups(ctx)

	expected := &ContactGroupsCollectionResponse{
		Size: 1,
		Collection: []*ContactGroup{
			createGroup(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	group := &ContactGroup{
		Name: "test",
	}

	mux.HandleFunc("/contacts/groups", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/group.json"))

		assertRequestMethod(t, r, "POST")

		if r.Header.Get("Content-Type") != string(ContentTypeXFormUrlencoded) {
			t.Errorf("Content-Type expected to equal: %s", string(ContentTypeXFormUrlencoded))
		}
	})

	result, _ := client.Contacts.CreateGroup(ctx, group)

	expectedResponse := createGroup()

	if !reflect.DeepEqual(expectedResponse, result) {
		t.Errorf("Given: %+v Expected: %+v", result, expectedResponse)
	}
}

func TestUpdateGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/group.json"))
		assertRequestMethod(t, r, "PUT")
	})

	group := createGroup()

	result, _ := client.Contacts.UpdateGroup(ctx, "1", group)

	if !reflect.DeepEqual(result, group) {
		t.Errorf("Given: %+v Expected: %+v", result, group)
	}
}

func TestDeleteGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Contacts.DeleteGroup(ctx, "1")

	if err != nil {
		t.Error(err)
	}
}

func TestMoveContactsToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestQueryParam(t, r, "q", "some-query")
	})

	filters := &ContactListFilters{
		Query: "some-query",
	}

	err := client.Contacts.MoveContactsToGroup(ctx, "1", filters)

	if err != nil {
		t.Error(err)
	}
}

func TestAddContactsToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "gender", "male")
	})

	filters := &ContactListFilters{
		Gender: "male",
	}

	err := client.Contacts.AddContactsToGroup(ctx, "1", filters)

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveContactsFromGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "first_name", "name1")
	})

	filters := &ContactListFilters{
		FirstName: []string{"name1"},
	}

	err := client.Contacts.AddContactsToGroup(ctx, "1", filters)

	if err != nil {
		t.Error(err)
	}
}

func TestAddContactToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/contacts/groups/1/members/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("contacts/contact.json"))

		assertRequestMethod(t, r, "PUT")
	})

	result, _ := client.Contacts.AddContactToGroup(ctx, "1", "1")

	expected := createContact()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func createExpectedContactsCollection() *ContactCollectionResponse {
	contact := createContact()

	expected := &ContactCollectionResponse{
		CollectionMeta: CollectionMeta{
			Size: 1,
		},
		Collection: []*Contact{
			contact,
		},
	}
	return expected
}

func createContact() *Contact {
	return &Contact{
		Id:           "1",
		FirstName:    "Jon",
		LastName:     "Doe",
		BirthdayDate: "1970-01-01",
		PhoneNumber:  "100200300",
		Gender:       "male",
		City:         "",
		Email:        "jondoe@somedomain.com",
		Source:       "",
		DateCreated:  "1970-01-01T00:00:00+00:00",
		DateUpdated:  "1970-01-01T00:00:00+00:00",
		Description:  "Jon Doe",
	}
}

func createGroup() *ContactGroup {
	return &ContactGroup{
		Id:            "1",
		Name:          "group",
		ContactsCount: 0,
		DateCreated:   "1970-01-01T00:00:00+00:00",
		DateUpdated:   "1970-01-01T00:00:00+00:00",
		Description:   "group",
		CreatedBy:     "j.doe",
		Idx:           "",
		Permissions: []*ContactGroupPermissions{
			createGroupPermissions(),
		},
	}
}

func createGroupPermissions() *ContactGroupPermissions {
	return &ContactGroupPermissions{
		Username: "j.doe",
		GroupId:  "1",
		Write:    true,
		Read:     true,
		Send:     true,
	}
}
