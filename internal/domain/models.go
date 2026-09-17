package domain

import (
	"errors"
	"time"
)

var (
	ErrSeatAlreadyBooked  = errors.New("seat is already booked or reserved")
	ErrBookingNotFound    = errors.New("booking not found")
	ErrBookingExpired     = errors.New("booking has expired")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type SeatStatus string

const (
	SeatStatusFree     SeatStatus = "free"
	SeatStatusReserved SeatStatus = "reserved"
	SeatStatusSold     SeatStatus = "sold"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Seat struct {
	ID         int64      `json:"id"`
	EventID    int64      `json:"event_id"`
	RowNumber  int        `json:"row_number"`
	SeatNumber int        `json:"seat_number"`
	Price      int64      `json:"price_kopecks"`
	Status     SeatStatus `json:"status"`
}

type Booking struct {
	ID        int64         `json:"id"`
	SeatID    int64         `json:"seat_id"`
	UserID    int64         `json:"user_id"`
	Status    BookingStatus `json:"status"`
	ExpiresAt time.Time     `json:"expires_at"`
	CreatedAt time.Time     `json:"created_at"`
}
