//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sendEmail
//

package gmail

import (
	"area/utils"
	"fmt"
)

// Careful about \r
const (
	createEmailFormat = "From: %s\r\n" +
		"To: %s\r\n" +
		"Subject: %s\r\n\r\n" +
		"%s"
)

type MessageBodyReq struct {
	Raw string `json:"raw,omitempty"`
}

func NewEmailRequestBody(emailRawContent string) MessageBodyReq {
	return MessageBodyReq{
		Raw: emailRawContent,
	}
}

func CreateEmailRawContent(from string, to string, subject string, mailContent string) string {
	mail := fmt.Sprintf(createEmailFormat, from, to, subject, mailContent)

	return utils.EncodeToBase64(mail)
}
