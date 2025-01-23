package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

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
	return &TicketServiceServer{
		users:          make(map[string]generated.User),
		seats:          make(map[string]string),
		availableSeats: seatConfig.Seats,
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

func (s *TicketServiceServer) SubmitTicket(req *generated.SubmitTicketRequest) (*generated.TicketReceipt, error) {

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
