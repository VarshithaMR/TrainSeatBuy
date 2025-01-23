package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	"TrainSeatBuy/cmd/config"
	"TrainSeatBuy/service/proto/generated"
)

type TicketServiceServer struct {
	users          map[string]generated.User
	seats          map[string]string
	availableSeats map[string]int
	generated.UnimplementedTicketServiceServer
}

func NewTicketServiceServer(seatConfig *config.SeatConfig) *TicketServiceServer {
	availableSeats := map[string]int{
		"A": seatConfig.A.Count,
		"B": seatConfig.B.Count,
	}

	return &TicketServiceServer{
		users:          make(map[string]generated.User),
		seats:          make(map[string]string),
		availableSeats: availableSeats,
	}
}

func generateAllSeats(seatConfig map[string]int) []string {
	var allSeats []string
	for section, count := range seatConfig {
		for i := 1; i <= count; i++ {
			seat := fmt.Sprintf("%s%d", section, i)
			allSeats = append(allSeats, seat)
		}
	}
	return allSeats
}

func (s *TicketServiceServer) SubmitRequest(ctx context.Context, req *generated.SubmitTicketRequest) (*generated.TicketReceipt, error) {
	if ctx == nil {
		return nil, errors.New("request context is empty")
	}

	//random seat allocation
	randomAllocatedSeat, err := s.allocateSeatRandomly()
	if err != nil {
		return nil, err
	}

	// Store the user and their allocated seat
	s.users[req.User.Email] = *req.User
	s.seats[randomAllocatedSeat] = req.User.Email

	// Generate the ticket receipt
	receipt := &generated.TicketReceipt{
		From:          req.From,
		To:            req.To,
		User:          req.User,
		PricePaid:     req.PricePaid,
		AllocatedSeat: randomAllocatedSeat,
	}

	return receipt, nil
}

func (s *TicketServiceServer) GetDetails(ctx context.Context, req *generated.GetDetailsRequest) (*generated.TicketReceipt, error) {
	if ctx == nil {
		return nil, errors.New("request context is empty")
	}

	user, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("user with email %s not found", req.Email)
	}

	var allocatedSeat string
	seat, found := s.seats[user.Email]
	if !found {
		allocatedSeat = "No seat allocated"
	} else {
		allocatedSeat = seat
	}

	receipt := &generated.TicketReceipt{
		From:          "",
		To:            "",
		User:          &user,
		PricePaid:     0,
		AllocatedSeat: allocatedSeat,
	}

	return receipt, nil
}

func (s *TicketServiceServer) GetUsersInSection(ctx context.Context, req *generated.GetUsersInSectionRequest) (*generated.UsersInSection, error) {
	if ctx == nil {
		return nil, errors.New("request context is empty")
	}

	var usersInSection []*generated.UserWithSeat

	for seat, email := range s.seats {
		if len(seat) > 0 && seat[:1] == req.Section {
			user, exists := s.users[email]
			if exists {
				usersInSection = append(usersInSection, &generated.UserWithSeat{
					User:          &user,
					AllocatedSeat: seat,
				})
			}
		}
	}

	usersInSec := &generated.UsersInSection{
		Section: req.Section,
		Users:   usersInSection,
	}
	return usersInSec, nil
}

// RemoveUser removes a user from the system based on their email.
func (s *TicketServiceServer) RemoveUser(ctx context.Context, req *generated.RemoveUserRequest) (*generated.RemoveUserResponse, error) {
	if ctx == nil || req.Email == "" {
		return nil, fmt.Errorf("invalid request")
	}

	_, exists := s.users[req.Email]
	if !exists {
		return &generated.RemoveUserResponse{
			Success: false,
			Message: fmt.Sprintf("User with email %s not found", req.Email),
		}, nil
	}

	assignedSeat := ""
	for seat, email := range s.seats {
		if email == req.Email {
			assignedSeat = seat
			break
		}
	}

	if assignedSeat != "" {
		delete(s.seats, assignedSeat)
	}

	delete(s.users, req.Email)

	return &generated.RemoveUserResponse{
		Success: true,
		Message: fmt.Sprintf("User %s removed successfully", req.Email),
	}, nil
}

func (s *TicketServiceServer) ModifyUserSeat(ctx context.Context, req *generated.ModifyUserSeatRequest) (*generated.ModifyUserSeatResponse, error) {
	if req.Email == "" || req.NewSeat == "" {
		return nil, fmt.Errorf("invalid email or new seat provided")
	}

	_, exists := s.users[req.Email]
	if !exists {
		return &generated.ModifyUserSeatResponse{
			Success: false,
			Message: fmt.Sprintf("User with email %s not found", req.Email),
		}, nil
	}

	// Check if the new seat is available
	_, seatAssigned := s.seats[req.NewSeat]
	if seatAssigned {
		return &generated.ModifyUserSeatResponse{
			Success: false,
			Message: fmt.Sprintf("Seat %s is already occupied", req.NewSeat),
		}, nil
	}

	oldSeat := ""
	for seat, email := range s.seats {
		if email == req.Email {
			oldSeat = seat
			break
		}
	}
	if oldSeat != "" {
		delete(s.seats, oldSeat) // Free up the old seat
	}

	// Assign new seat
	s.seats[req.NewSeat] = req.Email

	return &generated.ModifyUserSeatResponse{
		Success: true,
		Message: fmt.Sprintf("User %s has been moved to seat %s", req.Email, req.NewSeat),
	}, nil
}

func (s *TicketServiceServer) allocateSeatRandomly() (string, error) {
	allSeats := generateAllSeats(s.availableSeats)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(allSeats)
	for i := 0; i < n; i++ {
		j := r.Intn(n - i)
		seat := allSeats[i+j]
		_, exists := s.seats[seat]
		if !exists {
			s.seats[seat] = ""
			return seat, nil
		}
		allSeats[i+j], allSeats[n-1] = allSeats[n-1], allSeats[i+j]
	}

	return "", errors.New("no available seats")
}
