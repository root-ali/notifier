package grafana

import (
	"fmt"
	"notifier/pkg/sender"
	"time"
)

type Service interface {
	GrafanaAlert(GrafanaResponseBody)
}

type GrafanaResponseBody struct {
	Receiver         string     `json:"receiver"`
	Status           string     `json:"status"`
	OrgID            int8       `json:"-"`
	Alerts           []alert    `json:"alerts"`
	GroupLabels      labels     `json:"-"`
	CommonLabels     labels     `json:"-"`
	CommonAnnotation annotation `json:"-"`
	ExternalURL      string     `json:"-"`
	Version          string     `json:"-"`
	GroupKey         string     `json:"-"`
	TruncatedAlerts  int8       `json:"-"`
	Title            string     `json:"-"`
	State            string     `json:"-"`
	Message          string     `json:"-"`
}

type alert struct {
	Status          string     `json:"status"`
	AlertLabels     labels     `json:"labels"`
	AlertAnnotation annotation `json:"annotations"`
	StartsAt        string     `json:"-"`
	EndsAt          string     `json:"-"`
	GeneratedURL    string     `json:"-"`
	Fingerprint     string     `json:"-"`
	SilenceURL      string     `json:"-"`
	DashboardURL    string     `json:"-"`
	PanelURL        string     `json:"-"`
	ValueString     string     `json:"-"`
}

type labels struct {
	AlertName string `json:"alertname"`
	Team      string `json:"-"`
	Zone      string `json:"-"`
	Receptor  string `json:"receptor"`
	Method    string `json:"method"`
	NtfyUrl   string `json:"ntfy_url"`
}

type annotation struct {
	Description string `json:"description"`
	RunbookURL  string `json:"-"`
	Summary     string `json:"summary"`
}

func GrafanaAlert(grb GrafanaResponseBody) error {
	alertTime := time.Now().Format(time.RFC822)
	var err error
	for _, alert := range grb.Alerts {
		alertStatus := alert.Status
		alertReceptor := alert.AlertLabels.Receptor
		alertDescription := alert.AlertAnnotation.Description
		alertMessage := alertStatus + "\n" + alertDescription + "\n" + alertTime
		if alert.AlertLabels.NtfyUrl != "" {
			err = sender.NtfySender(alert.AlertLabels.NtfyUrl, alertMessage)
			if err != nil {
				fmt.Printf("ntfy error is : %v \n ", err)
			}
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
