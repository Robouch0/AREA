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

const (
	createEmailFormat = "From: %s\r\n" +
		"To: %s\r\n" +
		"Subject: %s\r\n\r\n" +
		"%s"
)

func CreateEmailRawContent(from string, to string, subject string, mailContent string) string {
	mail := fmt.Sprintf(createEmailFormat, from, to, subject, mailContent)

	return utils.EncodeToBase64(mail)
}
