from flask import Flask, request
import json

app = Flask(__name__)


@app.route("/")
def hello_world():
    return "<p>Hello, 13sai!</p>"


@app.errorhandler(404)
def not_found(error):
    return "<p>404</p>"


@app.route("/api", methods=["POST"])
def api_():
    data = request.data.decode("utf-8")
    if len(data) < 1:
        return json.dumps(
            {"message": "data is null", "status": 1001}, separators=(",", ":")
        )

    ret = {"status": 0, "message": "ok"}
    return json.dumps(ret, separators=(",", ":"))
