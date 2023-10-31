// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    notification, err := UnmarshalNotification(bytes)
//    bytes, err = notification.Marshal()

package core

import "encoding/json"

func UnmarshalNotification(data []byte) (Notification, error) {
	var r Notification
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Notification) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Notification event. Appropriate when an event has no data aside from recipient settings.
// If the event requires data, then it should reference the Recipient object definition in a
// separate schema.
type Notification struct {
	NotificationRecipients *Recipients `json:"notification_recipients,omitempty"`
}

// Notification recipients. Should be in a top-level field named "notification_recipients"
type Recipients struct {
	// List of emails to direct the notification to. This won’t override notification's               
	// administrators settings. Emails list will be merged with other settings.                       
	Emails                                                                                   []string `json:"emails,omitempty"`
	// Setting to true ignores all the user preferences on this Recipient setting (It doesn’t         
	// affect other configuration that an Administrator sets on their Notification settings).         
	// Setting to false honors the user preferences.                                                  
	IgnoreUserPreferences                                                                    *bool    `json:"ignore_user_preferences,omitempty"`
	// Setting to true sends an email to the administrators of the account. Setting to false          
	// sends an email to all users of the account.                                                    
	OnlyAdmins                                                                               *bool    `json:"only_admins,omitempty"`
	// List of users to direct the notification to. This won’t override notification's                
	// administrators settings. Users list will be merged with other settings.                        
	Users                                                                                    []string `json:"users,omitempty"`
}
