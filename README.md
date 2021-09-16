# notifier fire alarm

## What does this project do:

the aim of this project is to send alarm notification from grafana alert manager via kavenegar api.

## state:

In testing stage

## How to:

For running this project you can use docker and should set ADMIN_PASS enviroment variable for sending sms.

```bash
docker image build -t notifier .
docker run -d -e ADMIN_PASS=${admin_pass} -e KAVENEGAR_API_KEY=${kavenevgar_api_key} notifier

```

for test the project you can use this example

```bash
curl -X GET notifier:5000/api/v1/notifier
```
should gave you the below response

```
{"status" : "ok"}
```
In the notification channel you should set method in POST and admin user in 'admin' and password with your pass in ADMIN_PASS set in environment.

![notifier](https://user-images.githubusercontent.com/25849537/133682805-d42bf901-02d1-4d96-876c-57e50ea5f645.jpg)


for config alert rule in dashboard in getting sms you need config this two tags:
![tags](https://user-images.githubusercontent.com/25849537/133682356-0d372ac3-3c52-47c0-87d6-cef1d139c64d.jpg)

in receptor tags you can config the mobile number you want to receive the alert. you should seperated number with comma ','.
