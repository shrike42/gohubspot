package gohubspot

import "fmt"

type ContactsService service

type contactBase struct {
	Vid          int    `json:"vid"`
	CanonicalVid int    `json:"canonical-vid"`
	MergedVids   []int  `json:"merged-vids"`
	PortalID     int    `json:"portal-id"`
	IsContact    bool   `json:"is-contact"`
	ProfileToken string `json:"profile-token"`
	ProfileURL   string `json:"profile-url"`
}

// Most times Contact properties are lists of properties in the API
type Contact struct {
	contactBase
	Properties []Property `json:"properties"`
}

// Occasionally properties in relation to contacts are NOT lists but named items
type ContactInList struct {
	contactBase
	Properties interface{} `json:"properties"`
}

func (cl *ContactInList) GetStringValue(key string) string {
	// Get the Properties as a map and search map of maps for value
	m := cl.Properties.(map[string]interface{})

	var value string
	if _, ok := m[key]; ok {
		value = m[key].(map[string]interface{})["value"].(string)
	}
	return value
}

func (s *ContactsService) Create(properties Properties) (*IdentityProfile, error) {
	url := "/contacts/v1/contact"
	res := new(IdentityProfile)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) Update(contactID int, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", contactID)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) UpdateByEmail(email string, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s/profile", email)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) CreateOrUpdateByEmail(email string, properties Properties) (*Vid, error) {
	url := fmt.Sprintf("/contacts/v1/contact/createOrUpdate/email/%s", email)

	res := new(Vid)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) DeleteById(id int) (*ContactDeleteResult, error) {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d", id)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(url, res)
	return res, err
}

func (s *ContactsService) Merge(primaryID, secondaryID int) error {
	url := fmt.Sprintf("/contacts/v1/contact/merge-vids/%d/", primaryID)
	secondary := struct {
		SecondaryID int `json:"vidToMerge"`
	}{
		SecondaryID: secondaryID,
	}

	return s.client.RunPost(url, secondary, nil)
}

func (s *ContactsService) GetByToken(token string) (*Contact, error) {
	url := fmt.Sprintf("/contacts/v1/contact/utk/%s/profile", token)
	res := new(Contact)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *ContactsService) GetByVid(vid int) (*ContactInList, error) {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", vid)
	res := new(ContactInList)
	err := s.client.RunGet(url, res)
	return res, err
}
