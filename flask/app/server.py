from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello World! Flask Gunicorn etc.'

if __name__ == '__main__':
    app.run()