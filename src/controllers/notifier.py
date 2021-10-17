import os
import time
from flask import Flask
from flask import request, g, Blueprint, json, Response, jsonify, logging
from flask_httpauth import HTTPBasicAuth
from werkzeug.security import generate_password_hash, check_password_hash
from flask_validate_json import validate_json
import logging
from . import sms

app = Flask(__name__)
logging.basicConfig(level=logging.DEBUG, format=f'%(asctime)s %(levelname)s %(name)s: %(message)s')
notifier = Blueprint('notifier', __name__)
auth = HTTPBasicAuth()
users = {
    "admin": generate_password_hash(os.environ['ADMIN_PASS'])
}

schema = {
      "type": "object",
      "properties": {
        "message": { "type": "string", "minLength": 2, "maxLength": 100 },
        "ruleName": { "type": "string" },
        "state": { "type": "string" },
        "tags" : { "type" : "object"}
      },
      "required": [ "message", "ruleName", "state", "tags" ]
    }

schemaAlertmanager = {
    "type" : "object",
    "properties" : { 
        "status" : { "type" : "string" },
        "receiver" : { "type" : "string" },
        "alerts" : { "type" : "array" }
    },
    "required" : [ "status" , "receiver" , "alerts" ]
}

@notifier.route('/notifier', methods=['POST'])
@auth.login_required
@validate_json(schema)
def grafana():
    clientip = request.remote_addr
    time_of_notif = time.strftime("%d-%m %H:%M:%S")
    app.logger.info("clientip %s , time %s" , clientip , time_of_notif )
    try:
        send_state = request.json["tags"]
        if 'sms' in send_state :
            body = request.json
            message = body["message"]
            ruleName = body["ruleName"]
            state = body["state"]
            receptor = body["tags"]["receptor"]
            sms = "\n".join([message, ruleName, state , time_of_notif ])
            app.logger.info("receptor %s" , receptor)
            app.logger.info("sms %s" , sms)
            if sms.sendSMS(sms, receptor):
                return '{"status" : "OK"}' , 200
            else:
                return '{"status":"error" , "message" : "sms not send"}' , 400
        else:
            app.logger.info("not for send sms notif")
            return '{"status" : "ok", "message" : "not for send sms"}', 200
    except Exception as e :
        app.logger.error(e)
        return '{"status" : "error" , "message" : "requst body error"}', 400
    return '{"status" : "ok", "message" : "notif sent"}', 200

@notifier.route('/notifier', methods=['GET'])
def status():
    return '{"status" : "ok"}', 200

@notifier.route('/amp' , methods=['POST'])
@auth.login_required
@validate_json(schemaAlertmanager)
def alertManager():
    time_of_notif = time.strftime("%d-%m %H:%M:%S")
    app.logger.info("new notifier request received %s" , time_of_notif)
    try:
        body = request.json
        alerts = body["alerts"]
        for alert in alerts:
            if "sms" in alert["labels"]:
                if "receptor" in alert["labels"]:
                    receptor = alert["labels"]["receptor"]
                else:
                    app.logger.info("receptor not found")
                    return '{"status" : "error" , "message" : "receptor not found" }', 200
                if "annotation" in alert:
                    message = alert["status"] + "\n" + alert["annotation"]["summary"] 
                else:
                    app.logger.info("annotation not found")
                    return '{"status":"error" , "message" : "annotation not found"}', 200
                if sms.sendSMS(message, receptor):
                    return '{"status" : "OK"}' , 200
                else:
                    return '{"status":"error" , "message" : "cannot send kavenegar sms"}' , 400
            else:
                app.logger.info("notif not for send via sms")
                return '{"status" : "not send"}' , 200

        app.logger.info("all notification has take care of")
        return '', 200
    except Exception as e:
        app.logger.error(e)
        return '{"status" : "error"}' , 400

@notifier.route('/amp', methods=['GET'])
def status_alertManager():
    return '{"status" : "OK"}', 200

@auth.verify_password
def verify_password(username, password):
    if username in users and \
            check_password_hash(users.get(username), password):
        return username

@auth.error_handler
def auth_error(status):
    res = '{"status" : "error" , "message" : "unauthorized"}'
    return res, 4013