import os
from flask import Flask

app = Flask(__name__)

@app.route('/')
def home():
    return f'Hello, TEST_ENV = {os.environ['TEST_ENV']}'

@app.route('/about')
def about():
    return 'About'
