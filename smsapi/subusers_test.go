package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	active := true
	user := &User{
		Credentials: nil,
		Active:      &active,
		Description: "test user",
		Points: &UserPoints{
			FromAccount: 100,
			PerMonth:    10,
		},
	}

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestBody(t, r, new(User), user)
	})

	result, _ := client.Subusers.CreateUser(ctx, user)

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestUpdateAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	inactive := false
	user := &User{
		Credentials: nil,
		Active:      &inactive,
		Description: "updated description",
		Points:      nil,
	}

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "PUT")
		assertRequestBody(t, r, new(User), user)
	})

	result, _ := client.Subusers.UpdateUser(ctx, "1", user)

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Subusers.GetUser(ctx, "1")

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestDeleteAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	client.Subusers.DeleteUser(ctx, "1")
}

func TestGetUsersList(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/users_list.json"))
		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Subusers.ListUsers(ctx, nil)

	expected := &UserCollectionResponse{
		Size: 1,
		Collection: []*UserResponse{
			createUserResponse(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestListUsersWithQuery(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "q", "john")
		fmt.Fprint(w, `{"size":0,"collection":[]}`)
	})

	_, err := client.Subusers.ListUsers(ctx, &UserCollectionFilters{Query: "john"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSubuserShares(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/1/shares", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"sendernames":{"access":"all"},"blacklist":{"access":"none"},"templates":{"access":"selected","templates":["t1"]}}`)
	})

	result, err := client.Subusers.GetShares(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Sendernames.Access != "all" || result.Templates.Access != "selected" || len(result.Templates.Templates) != 1 {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestGetSubuserSendernamesAccess(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/1/shares/sendernames", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"access":"selected","senders":["Info","Alert"]}`)
	})

	result, err := client.Subusers.GetSendernamesAccess(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Access != "selected" || len(result.Senders) != 2 {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestUpdateSubuserSendernamesAccess(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	access := &SubuserAccess{Access: "all"}

	mux.HandleFunc("/subusers/1/shares/sendernames", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestJsonContains(t, r, "access", "all")
	})

	if err := client.Subusers.UpdateSendernamesAccess(ctx, "1", access); err != nil {
		t.Fatal(err)
	}
}

func TestGetSubuserTemplatesAccess(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/1/shares/templates", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"access":"all"}`)
	})

	result, err := client.Subusers.GetTemplatesAccess(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Access != "all" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestUpdateSubuserTemplatesAccess(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	access := &SubuserAccess{Access: "selected", Templates: []string{"t1"}}

	mux.HandleFunc("/subusers/1/shares/templates", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestJsonContains(t, r, "access", "selected")
		fmt.Fprint(w, `{"access":"selected","templates":["t1"]}`)
	})

	result, err := client.Subusers.UpdateTemplatesAccess(ctx, "1", access)
	if err != nil {
		t.Fatal(err)
	}
	if result.Access != "selected" || len(result.Templates) != 1 {
		t.Errorf("Unexpected: %+v", result)
	}
}

func createUserResponse() *UserResponse {
	return &UserResponse{
		Id:          "1",
		Username:    "test",
		Active:      true,
		Description: "",
		Points: &UserPoints{
			FromAccount: 0,
			PerMonth:    0,
		},
	}
}
