from flask import Flask
from flask import render_template
from flask import request

from countdown import generate_numbers


app = Flask(__name__)

@app.route("/")
def home():
    return render_template("index.html")


@app.route("/numbers", methods=['GET', 'POST'])
def numbers():
    if request.method == 'POST':
        big_count = request.form['big_count']
        target, selection = generate_numbers(big_count)
        return render_template("numbers/index.html", target=target, selection=selection)
    else:
        return render_template("numbers/index.html")


@app.route("/letters")
def letters():
    return render_template("letters/index.html")
