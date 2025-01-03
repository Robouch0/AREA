//
// EPITECH PROJECT, 2024
// AREA
// File description:
// encodeToBase64
//

package utils

import "encoding/base64"

func EncodeToBase64(mess string) string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(mess)))
	base64.StdEncoding.Encode(dst, []byte(mess))
	return string(dst)
}

func EncodeToBase64Bytes(mess string) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(mess)))
	base64.StdEncoding.Encode(dst, []byte(mess))
	return dst
}
