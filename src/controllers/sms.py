from kavenegar import *
import logging
import os

logging.basicConfig(level=logging.DEBUG, format=f'%(asctime)s %(levelname)s %(name)s: %(message)s')

def sendSMS(message , receptor):
    kavenegar_api_key = os.environ['KAVENEGAR_API_KEY']
    try:
        api = KavenegarAPI(kavenegar_api_key)
        params = {
            'receptor': receptor,
            'message': message,
        }
        response = api.sms_send(params)
        logging.info("message has been sent")
        return True
    except APIException as e:
        logging.error(e)
        return False
    except HTTPException as e:
        logging.error(e)
        return False
