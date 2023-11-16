from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"
    
def run():
    while True:
        try:
            app.run()
        except Exception as ex:
            print(ex)