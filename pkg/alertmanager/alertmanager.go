package alertmanager

import (
	"fmt"
	"notifier/pkg/sender"
	"time"
)

type AlertManagerRequestBody struct {
	Version          string     `json:"version"`
	GroupKey         string     `json:"-"`
	TruncatedAlerts  int        `json:"-"`
	Status           string     `json:"status"`
	Receiver         string     `json:"receiver"`
	GroupLabels      labels     `json:"-"`
	CommonLabels     labels     `json:"-"`
	CommonAnnotation annotation `json:"-"`
	ExternalURL      string     `json:"-"`
	Alerts           []alert    `json:"alerts"`
}

type alert struct {
	Status          string     `json:"status"`
	AlertLabels     labels     `json:"labels"`
	AlertAnnotation annotation `json:"annotations"`
	StartsAt        string     `json:"-"`
	EndsAt          string     `json:"-"`
	GeneratedURL    string     `json:"-"`
	Fingerprint     string     `json:"-"`
}

type labels struct {
	Receptor string `json:"receptor"`
	Method   string `json:"method"`
	NtfyUrl  string `json:"ntfy_url"`
}

type annotation struct {
	Summary string `json:"summary"`
}

func AlertManagerReq(grb AlertManagerRequestBody) error {
	var err error
	alertTime := time.Now().Format(time.RFC822)
	for _, alert := range grb.Alerts {
		alertStatus := alert.Status
		alertReceptor := alert.AlertLabels.Receptor
		alertSummary := alert.AlertAnnotation.Summary
		alertMessage := alertStatus + "\n" + alertSummary + "\n" + alertTime
		if alert.AlertLabels.NtfyUrl != "" {
			sender.NtfySender(alert.AlertLabels.NtfyUrl, alertMessage)
		}
		if alert.AlertLabels.Method == "sms" {
			err = sender.KavenegarSMSSender(alertReceptor, alertMessage)
			if err != nil {
				fmt.Printf("Kavenegar error is %v\n", err)
			}
		}
		if alert.AlertLabels.Method == "call" {
			err = sender.KavenegarCallSender(alertReceptor, alertMessage)
			if err != nil {
				fmt.Printf("Kavenegar error is %v\n", err)
			}
		}

	}

	if err != nil {
		return err
	}

	return nil
}
