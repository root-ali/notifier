from flask import Flask
import logging
from src.controllers.notifier import notifier

def create_app(env_name):
  """
  Create app
  """
  app = Flask(__name__)
 
  logging.basicConfig(level=logging.DEBUG, format=f'%(asctime)s %(levelname)s %(name)s %(threadName)s: %(message)s')

  app.register_blueprint(notifier, url_prefix='/api/v1/')
  return app