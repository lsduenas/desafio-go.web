package tickets

import (
	"context"
	"errors"
)

type ticketService struct{
	repo Repository
}

type serviceController interface {
	GetTotalTickets(ctx context.Context,  destination string) (int, error)
	AverageDestination(ctx context.Context,  destination string) (float64, error)
}

func NewService (repository Repository) ticketService {
	return ticketService{repo: repository}
}

func (ts *ticketService) GetTotalTickets(ctx context.Context,  destination string) (total int, er error){
	ticketsList, err := ts.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		er = errors.New("Can't get tickets list")
	}
	total = len(ticketsList)
	return 
}

func (ts ticketService) AverageDestination(ctx context.Context,  destination string) (average float64, er error) {
	totalTicketsListPerDestination, err := ts.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		er = errors.New("Can't get tickets list per destination")
	}
	totalTicketsPerDestination := len(totalTicketsListPerDestination)

	ticketsList, err := ts.repo.GetAll(ctx)
	if err != nil {
		er = errors.New("Can't get tickets list")
	}
	totalTickets := len(ticketsList)

	average = float64(totalTicketsPerDestination / totalTickets)
	return
}
