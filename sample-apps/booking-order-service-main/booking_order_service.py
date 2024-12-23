from flask import Flask

app = Flask(__name__)

@app.route("/")
def order_home():
    return "<h1>Welcome to the Booking Order Service</h1>"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5002)
