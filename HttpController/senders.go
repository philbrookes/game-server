package HttpController

import (
	"encoding/json"
	"errors"
	"net/http"
)

//Sender is the type that all senders must match
type Sender func(http.ResponseWriter, interface{}) error

//SenderFormat is a string which contains a format for sending output in
type SenderFormat string

//SenderFormatJSON contains the json output format string
const SenderFormatJSON = "json"

func sendJSON(rw http.ResponseWriter, res interface{}) error {
	textRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	headers := rw.Header()
	headers.Add("Content-Type", "Application/json")
	_, err = rw.Write(textRes)
	return err
}

//GetSender a sender for the provided format string
func GetSender(format string) (Sender, error) {
	switch format {
	case SenderFormatJSON:
		return sendJSON, nil
	}

	return nil, errors.New("Could not find sender for requested format: " + format)
}
