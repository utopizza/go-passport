package session

import (
	"context"
	"testing"
)

func TestCreateSession(t *testing.T) {
	ctx := context.Background()
	userId := int64(996)

	// create
	sessionData, err := CreateSession(ctx, userId)
	if err != nil {
		t.Fatal(err)
	}

	// read
	sessionDataRead, err := ReadSession(ctx, sessionData.Key)
	if err != nil {
		t.Fatal(err)
	}
	if sessionDataRead.UserId != sessionData.UserId {
		t.Fatal("wrong userId in sessionData reading from redis")
	}

	// delete
	if err := DeleteSession(ctx, sessionData.Key); err != nil {
		t.Fatal(err)
	}
}
