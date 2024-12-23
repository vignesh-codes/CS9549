from flask import Flask

app = Flask(__name__)

@app.route("/")
def payment_home():
    return "<h1>Welcome to the Booking Payment Service</h1>"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5003)
