import os
from src.app import create_app

app = create_app("dev")

if __name__ == '__main__':
    port = os.getenv('PORT')
    app.run( host='0.0.0.0', port=port )