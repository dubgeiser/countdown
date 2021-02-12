from flask import Flask, render_template

from countdown import generate_numbers


app = Flask(__name__)

@app.route("/")
def home():
    return render_template("index.html")


@app.route("/numbers")
def numbers():
    target, selection = generate_numbers()
    return render_template("numbers.html", target=target, selection=selection)


@app.route("/letters")
def letters():
    return render_template("letters.html")
