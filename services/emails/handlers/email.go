package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/weschang15/mailreceivable"
)

type Email struct {
	Address string `json:"address"`
}

type Receivable struct {
	Email string `json:"email"`
	Valid bool `json:"valid"`
	VerifiedMX bool `json:"verifiedMx"`
	VerifiedHost bool `json:"verifiedHost"`
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email Email

	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		log.Fatal(err)
	}

	if err := mailreceivable.Validate(email.Address); err != nil {
		json.NewEncoder(w).Encode(&Receivable{
			Email: email.Address,
			Valid: false,
		})
		return
	}

	if err := mailreceivable.VerifyHost(email.Address); err != nil {
		verifiedMx := err == mailreceivable.ErrHostNotFound
		verifiedHost := err != mailreceivable.ErrHostNotFound

		json.NewEncoder(w).Encode(&Receivable{
			Email: email.Address,
			Valid: true,
			VerifiedMX: verifiedMx,
			VerifiedHost: verifiedHost,
		})

		return
	}
	
	json.NewEncoder(w).Encode(&Receivable{
		Email: email.Address,
		Valid: true,
		VerifiedMX: true,
		VerifiedHost: true,
	})
}