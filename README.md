# notifier fire alarm

## What does this project do:

the aim of this project is to send alarm notification from grafana alert manager via kavenegar api.

## state:

In testing stage

## How to:

For running this project you can use docker and should set ADMIN_PASS enviroment variable for sending sms.

```bash
docker image build -t notifier .
docker run -d -e ADMIN_PASS=${admin_pass} notifier

```

for test the project you can use this example

```bash
curl -X GET notifier:5000/api/v1/notifier
```
should gave you the below response

```
{"status" : "ok"}
```