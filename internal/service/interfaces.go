package service

import (
	"context"
	"time"

	"github.com/zga/ticket_reservator/internal/domain"
)

type BookingRepository interface {
	ReserveSeat(ctx context.Context, seatID, userID int64, expiresAt time.Time) (domain.Booking, error)
	ConfirmBooking(ctx context.Context, bookingID int64) error
	CancelExpiredBookings(ctx context.Context) (int64, error)
}

type SeatRepository interface {
	GetByEventID(ctx context.Context, eventID int64) ([]domain.Seat, error)
}
