package gohubspot

import "fmt"

type EmailService service

// Remember that this is a one way operation with HubSpot
// You cannot re-optin a after this optout the contact will
// need to go through the email resubscribe process
func (s *EmailService) UnsubscribeFromAll(email string) error {
	url := fmt.Sprintf("/email/public/v1/subscriptions/%s", email)
	// A map in this format translates into
	// {"unsubscribeFromAll": true} when marshaled
	body := map[string]bool{"unsubscribeFromAll": true}
	res := map[string]string{}

	err := s.client.RunPut(url, body, res)
	return err
}
