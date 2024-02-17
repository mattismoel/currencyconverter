package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Base Prompt type. It serves as a container for the users input, if all flags are validated correctly.
type Prompt struct {
	Amount float32
	From   string
	To     string
}

// Returns an instance of a prompt struct.
func NewPrompt(amount float32, from string, to string) *Prompt {
	p := &Prompt{
		Amount: amount,
		From:   strings.ToUpper(from),
		To:     strings.ToUpper(to),
	}
	return p
}

// Returns the rate for the prompt.
func (p *Prompt) Rate() (float32, error) {
	err := p.validate()
	if err != nil {
		return 0.0, err
	}

	url := fmt.Sprintf("https://api.frankfurter.app/latest?from=%s", strings.ToUpper(p.From))

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("something went wrong with the api.")
	}

	var r ApiResponse

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return 0.0, err
	}

	val, ok := r.Rates[p.To]
	if !ok {
		return 0.0, fmt.Errorf("could not find the %q rate %q...", "to", p.To)
	}

	return p.Amount * val, nil
}

// Validates a prompt, and returns an error or nil.
func (p *Prompt) validate() error {
	err := validateAmount(p.Amount)
	if err != nil {
		return err
	}

	return nil
}

// Validates an amount.
func validateAmount(amount float32) error {
	if amount < 0 {
		return errors.New("amount is negative")
	}
	return nil
}

