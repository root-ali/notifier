package sender

import (
	"fmt"
	"strings"

	"github.com/kavenegar/kavenegar-go"
	"github.com/spf13/viper"
)

func KavenegarSMSSender(receiver string, message string) error {
	viper.BindEnv("KAVENEGAR_API_KEY")
	apiKeyEnv := viper.Get("KAVENEGAR_API_KEY")
	apiKey := fmt.Sprintf("%v", apiKeyEnv)
	api := kavenegar.New(apiKey)
	sender := ""
	receptor := strings.Split(receiver, ",")
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
			return err
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
			return err
		default:
			fmt.Println(err.Error())
			return err
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
		}
	}

	return nil
}

func KavenegarCallSender(receiver string, message string) error {
	viper.BindEnv("KAVENEGAR_API_KEY")
	apiKeyEnv := viper.Get("KAVENEGAR_API_KEY")
	apiKey := fmt.Sprintf("%v", apiKeyEnv)
	api := kavenegar.New(apiKey)
	if res, err := api.Call.MakeTTS(receiver, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
			return err
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
			return err
		default:
			fmt.Println(err.Error())
			return err
		}
	} else {
		fmt.Println(res)
	}

	return nil
}
