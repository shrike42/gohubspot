package gohubspot

import (
	"fmt"
	"strconv"
)

type ContactListsService service

type ContactList struct {
	ParentID       int      `json:"parentId"` // Integer, read-only; The ID of the folder that the list belongs to.  Currently folders can only be managed in the HubSpot app
	Dynamic        bool     `json:"dynamic"`
	Name           string   `json:"name"`
	Filters        Filters  `json:"filters"`
	PortalID       int      `json:"portalId"`
	CreatedAt      UnixTime `json:"createdAt"`
	ListID         int      `json:"listId"`
	UpdatedAt      UnixTime `json:"updatedAt"`
	InternalListID int      `json:"internalListId"`
	Deleteable     bool     `json:"deleteable"`
	Metadata       Metadata `json:"metaData"`
}

type Metadata struct {
	Processing ProcessingType `json:"processing"`

	Size int `json:"size"`
	// Integer, read-only; The number of contacts in the list.

	Error string `json:"error"`
	// String, read-only; Any errors that happened the last time the list was processed.

	LastProcessingStateChangeAt UnixTime `json:"lastProcessingStateChangeAt"`

	LastSizeChangeAt UnixTime `json:"lastSizeChangeAt"`
}

type ContactLists struct {
	Lists []ContactList `json:"lists"`
	Page
}

type Contacts struct {
	Lists   []ContactInList `json:"contacts"`
	HasMore bool            `json:"has-more"`
	Offset  int             `json:"vid-offset"`
}

type contactListsOptions struct {
	listCount int
	offset    int
}

func NewContactListOptions(listCount, offset int) *contactListsOptions {
	if listCount > 250 {
		listCount = 250
	}
	return &contactListsOptions{listCount: listCount, offset: offset}
}

func (s *ContactListsService) GetContactLists(cl *contactListsOptions) (*ContactLists, error) {

	url := "/contacts/v1/lists"

	// Add params
	url += "/?count=" + strconv.Itoa(cl.listCount)
	url += "&offset=" + strconv.Itoa(cl.offset)

	req, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}

	list := new(ContactLists)
	err = s.client.Do(req, list)

	return list, err
}

func (s *ContactListsService) CreateContactList(name string) (*ContactList, error) {
	nameBody := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	url := "/contacts/v1/lists"
	req, err := s.client.Post(url, nameBody)

	if err != nil {
		return nil, err
	}
	list := new(ContactList)
	err = s.client.Do(req, list)
	return list, err
}

func (s *ContactListsService) GetContactList(listId int) (*ContactList, error) {
	url := fmt.Sprintf("/contacts/v1/lists/%d", listId)
	req, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}
	list := new(ContactList)
	err = s.client.Do(req, list)
	return list, err
}

func (s *ContactListsService) GetContacts(cl *contactListsOptions, listId int) (*Contacts, error) {
	url := fmt.Sprintf("/contacts/v1/lists/%d/contacts/all", listId)
	// Add params
	url += "/?count=" + strconv.Itoa(cl.listCount)
	url += "&vidOffset=" + strconv.Itoa(cl.offset)
	url += "&property=firstname"

	req, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}
	cs := new(Contacts)
	err = s.client.Do(req, cs)
	return cs, err
}
