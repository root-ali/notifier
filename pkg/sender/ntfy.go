package sender

import (
	"bytes"
	"fmt"
	"net/http"
)

func NtfySender(receptor string, message string) error {
	postData := []byte(message)
	_, err := http.Post(receptor, "", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return nil
}
