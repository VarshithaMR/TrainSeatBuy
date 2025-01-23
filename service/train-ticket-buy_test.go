package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"

	"TrainSeatBuy/cmd/config"
	"TrainSeatBuy/service/proto/generated"
)

var (
	seatConfig = config.SeatConfig{
		A: config.SectionConfig{
			Count: 5,
		},
		B: config.SectionConfig{
			Count: 5,
		},
	}
	server = NewTicketServiceServer(&seatConfig)
	ctx    = context.Background()
)

func TestSubmitTicket(t *testing.T) {
	req := &generated.SubmitTicketRequest{
		From:      "New York",
		To:        "Los Angeles",
		User:      &generated.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 100,
	}

	receipt, err := server.SubmitRequest(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, "John", receipt.User.FirstName)
	assert.Equal(t, "Doe", receipt.User.LastName)
	assert.Equal(t, "john.doe@example.com", receipt.User.Email)
	assert.NotEmpty(t, receipt.AllocatedSeat)
}

func TestGetDetails(t *testing.T) {
	reqSubmit := &generated.SubmitTicketRequest{
		From:      "New York",
		To:        "Los Angeles",
		User:      &generated.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 100,
	}
	_, err := server.SubmitRequest(ctx, reqSubmit)
	assert.Nil(t, err)

	reqGetDetails := &generated.GetDetailsRequest{
		Email: "john.doe@example.com",
	}
	receipt, err := server.GetDetails(ctx, reqGetDetails)

	assert.Nil(t, err)
	assert.Equal(t, "john.doe@example.com", receipt.User.Email)
	assert.Equal(t, "John", receipt.User.FirstName)
	assert.Equal(t, "Doe", receipt.User.LastName)
	assert.NotEmpty(t, receipt.AllocatedSeat)
}

func TestGetUsersInSection(t *testing.T) {
	reqSubmit1 := &generated.SubmitTicketRequest{
		From:      "New York",
		To:        "Los Angeles",
		User:      &generated.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 100,
	}
	_, err := server.SubmitRequest(ctx, reqSubmit1)
	assert.Nil(t, err)

	reqSubmit2 := &generated.SubmitTicketRequest{
		From:      "Chicago",
		To:        "New York",
		User:      &generated.User{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com"},
		PricePaid: 120,
	}
	_, err = server.SubmitRequest(ctx, reqSubmit2)
	assert.Nil(t, err)

	reqGetUsersInSection := &generated.GetUsersInSectionRequest{
		Section: "A",
	}
	usersInSection, err := server.GetUsersInSection(ctx, reqGetUsersInSection)

	assert.Nil(t, err)
	//
	assert.Len(t, usersInSection.Users, 2)
}

func TestRemoveUser(t *testing.T) {
	server.users["test@example.com"] = generated.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@example.com",
	}
	server.seats["A1"] = "test@example.com"

	tests := []struct {
		name         string
		email        string
		expectError  bool
		expectedMsg  string
		expectedCode codes.Code
	}{
		{
			name:         "Successfully remove user",
			email:        "test@example.com",
			expectedMsg:  "User test@example.com removed successfully",
			expectedCode: codes.OK,
		},
		{
			name:         "User not found",
			email:        "nonexistent@example.com",
			expectedMsg:  "User with email nonexistent@example.com not found",
			expectedCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &generated.RemoveUserRequest{
				Email: tt.email,
			}
			resp, err := server.RemoveUser(ctx, req)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedMsg, resp.Message)
			}
		})
	}
}

func TestModifyUserSeat(t *testing.T) {
	reqSubmit := &generated.SubmitTicketRequest{
		From:      "New York",
		To:        "Los Angeles",
		User:      &generated.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 100,
	}
	_, err := server.SubmitRequest(ctx, reqSubmit)
	assert.Nil(t, err)

	reqModify := &generated.ModifyUserSeatRequest{
		Email:   "john.doe@example.com",
		NewSeat: "A1",
	}
	resp, err := server.ModifyUserSeat(ctx, reqModify)

	assert.Nil(t, err)
	assert.True(t, resp.Success)
}

func TestModifyUserSeat_SeatOccupied(t *testing.T) {
	reqSubmit := &generated.SubmitTicketRequest{
		From:      "New York",
		To:        "Los Angeles",
		User:      &generated.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		PricePaid: 100,
	}
	_, err := server.SubmitRequest(ctx, reqSubmit)

	server.seats["A1"] = "other.user@example.com"

	reqModify := &generated.ModifyUserSeatRequest{
		Email:   "john.doe@example.com",
		NewSeat: "A1",
	}
	resp, err := server.ModifyUserSeat(ctx, reqModify)

	assert.Nil(t, err)
	assert.False(t, resp.Success)
}
