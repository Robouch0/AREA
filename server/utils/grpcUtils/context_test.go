//
// EPITECH PROJECT, 2025
// AREA
// File description:
// context_test
//

package grpcutils

import "testing"

func TestContextFromUser(t *testing.T) {
	userID := 1
	ctx := CreateContextFromUserID(userID)
	newUserID, err := GetUserIdFromContext(ctx, "service")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if userID != int(newUserID) {
		t.Errorf("Get from context did not work")
		return
	}
}
