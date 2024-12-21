//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gmailMessage
//

package gmail

type GmailMessageBody struct {
	AttachmentID string `json:"attachmentId"`
	Size         int    `json:"size"`
	Data         string `json:"data"`
}

type GmailHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GmailMessagePart struct {
	Header []GmailHeader    `json:"header,omitempty"`
	Body   GmailMessageBody `json:"body,omitempty"`
}

type GmailMessage struct {
	ID      string           `json:"id,omitempty"`
	Payload GmailMessagePart `json:"payload,omitempty"`
	Raw     string           `json:"raw,omitempty"`
}

// Give a Gmail message with only the email raw content in it
func NewEmailRequestBody(emailRawContent string) GmailMessage {
	return GmailMessage{
		Raw: emailRawContent,
	}
}
