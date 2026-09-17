package service

import (
	"context"
	"fmt"
	"time"

	"github.com/zga/ticket_reservator/internal/domain"
)

const holdDuration = 10 * time.Minute

type BookingService struct {
	bookingRepo BookingRepository
	seatRepo    SeatRepository
}

func NewBookingService(b BookingRepository, s SeatRepository) *BookingService {
	return &BookingService{
		bookingRepo: b,
		seatRepo:    s,
	}
}

func (s *BookingService) GetHallLayout(ctx context.Context, eventID int64) ([]domain.Seat, error) {
	return s.seatRepo.GetByEventID(ctx, eventID)
}

func (s *BookingService) Reserve(ctx context.Context, userID, seatID int64) (domain.Booking, error) {
	expiresAt := time.Now().Add(holdDuration)

	booking, err := s.bookingRepo.ReserveSeat(ctx, seatID, userID, expiresAt)
	if err != nil {
		return domain.Booking{}, fmt.Errorf("service.Reserve: %w", err)
	}

	return booking, nil
}

func (s *BookingService) Confirm(ctx context.Context, bookingID int64) error {
	if err := s.bookingRepo.ConfirmBooking(ctx, bookingID); err != nil {
		return fmt.Errorf("service.Confirm: %w", err)
	}
	return nil
}

func (s *BookingService) ReleaseExpired(ctx context.Context) (int64, error) {
	return s.bookingRepo.CancelExpiredBookings(ctx)
}
