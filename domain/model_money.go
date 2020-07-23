package domain

import (
	"github.com/Rhymond/go-money"
)

type Money struct {
	Amount   int64  `bson:"amount" json:"amount"`
	Currency string `bson:"currency" json:"currency"`
}

// NewMoney creates and returns new instance of Money.
func NewMoney(amount int64, code string) *Money {
	return &Money{
		Amount:   amount,
		Currency: code,
	}
}

// SameCurrency check if given Money is equals by currency.
func (m *Money) SameCurrency(om *Money) bool {
	return m.Currency == om.Currency
}

// Equals checks equality between two Money types.
func (m *Money) Equals(om *Money) (bool, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)
	return thisMoney.Equals(thatMoney)
}

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)

	return thisMoney.GreaterThan(thatMoney)
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)
	return thisMoney.GreaterThanOrEqual(thatMoney)
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)

	return thisMoney.LessThan(thatMoney)
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)

	return thisMoney.LessThanOrEqual(thatMoney)
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool {
	thisMoney := money.New(m.Amount, m.Currency)
	return thisMoney.IsZero()
}

// IsPositive returns boolean of whether the value of Money is positive.
func (m *Money) IsPositive() bool {
	thisMoney := money.New(m.Amount, m.Currency)
	return thisMoney.IsPositive()
}

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool {
	thisMoney := money.New(m.Amount, m.Currency)
	return thisMoney.IsNegative()
}

// Absolute returns new Money struct from given Money using absolute monetary value.
func (m *Money) Absolute() *Money {
	thisMoney := money.New(m.Amount, m.Currency)
	absolute := thisMoney.Absolute()
	return &Money{Amount: absolute.Amount(), Currency: absolute.Currency().Code}
}

// Negative returns new Money struct from given Money using negative monetary value.
func (m *Money) Negative() *Money {
	thisMoney := money.New(m.Amount, m.Currency)
	negative := thisMoney.Negative()
	return &Money{Amount: negative.Amount(), Currency: negative.Currency().Code}
}

// Add returns new Money struct with value representing sum of Self and Other Money.
func (m *Money) Add(om *Money) (*Money, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)
	result, err := thisMoney.Add(thatMoney)
	if err != nil {
		return nil, err
	}

	return &Money{Amount: result.Amount(), Currency: result.Currency().Code}, nil
}

// Subtract returns new Money struct with value representing difference of Self and Other Money.
func (m *Money) Subtract(om *Money) (*Money, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	thatMoney := money.New(om.Amount, om.Currency)
	result, err := thisMoney.Subtract(thatMoney)
	if err != nil {
		return nil, err
	}

	return &Money{Amount: result.Amount(), Currency: result.Currency().Code}, nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier.
func (m *Money) Multiply(mul int64) *Money {
	thisMoney := money.New(m.Amount, m.Currency)
	result := thisMoney.Multiply(mul)

	return &Money{Amount: result.Amount(), Currency: result.Currency().Code}
}

// Round returns new Money struct with value rounded to nearest zero.
func (m *Money) Round() *Money {
	thisMoney := money.New(m.Amount, m.Currency)
	result := thisMoney.Round()

	return &Money{Amount: result.Amount(), Currency: result.Currency().Code}
}

// Split returns slice of Money structs with split Self value in given number.
// After division leftover pennies will be distributed round-robin amongst the parties.
// This means that parties listed first will likely receive more pennies than ones that are listed later.
func (m *Money) Split(n int) ([]*Money, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	result, err := thisMoney.Split(n)
	if err != nil {
		return nil, err
	}
	moneyList := make([]*Money, len(result))
	for index, moneyItem := range result {
		moneyList[index] = &Money{Amount: moneyItem.Amount(), Currency: moneyItem.Currency().Code}
	}
	return moneyList, nil
}

// Allocate returns slice of Money structs with split Self value in given ratios.
// It lets split money by given ratios without losing pennies and as Split operations distributes
// leftover pennies amongst the parties with round-robin principle.
func (m *Money) Allocate(rs ...int) ([]*Money, error) {
	thisMoney := money.New(m.Amount, m.Currency)
	result, err := thisMoney.Allocate(rs...)
	if err != nil {
		return nil, err
	}
	moneyList := make([]*Money, len(result))
	for index, moneyItem := range result {
		moneyList[index] = &Money{Amount: moneyItem.Amount(), Currency: moneyItem.Currency().Code}
	}
	return moneyList, nil
}

// Display lets represent Money struct as string in given Currency value.
func (m *Money) Display() string {
	thisMoney := money.New(m.Amount, m.Currency)
	return thisMoney.Display()
}

// AsMajorUnits lets represent Money struct as subunits (float64) in given Currency value
func (m *Money) AsMajorUnits() float64 {
	thisMoney := money.New(m.Amount, m.Currency)
	return thisMoney.AsMajorUnits()
}
