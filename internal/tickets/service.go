package tickets

import (
	"context"
	"errors"
)

type TicketService struct{
	repo Repository
}

type serviceController interface {
	GetTotalTickets(ctx context.Context,  destination string) (int, error)
	AverageDestination(ctx context.Context,  destination string) (float64, error)
}

func NewService (repository Repository) TicketService {
	return TicketService{repo: repository}
}

func (ts TicketService) GetTotalTickets(ctx context.Context,  destination string) (total int, er error){
	ticketsList, err := ts.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		er = errors.New("Can't get tickets list")
		return 
	}
	total = len(ticketsList)
	return 
}

func (ts TicketService) AverageDestination(ctx context.Context,  destination string) (average float64, err error) {
	totalTicketsListPerDestination, err := ts.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		err = errors.New("Can't get tickets list per destination")
		return 
	}

	totalTicketsPerDestination := len(totalTicketsListPerDestination)
	
	ticketsList, err := ts.repo.GetAll(ctx)
	if err != nil {
		err = errors.New("Can't get tickets list")
		return 
	}
	totalTickets := len(ticketsList)
	average = float64(totalTicketsPerDestination) / float64(totalTickets)
	return
}
