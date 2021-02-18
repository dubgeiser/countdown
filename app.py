from flask import Flask
from flask import redirect
from flask import render_template
from flask import request

from countdown import numbers as n
from countdown.forms import NumbersForm


app = Flask(__name__)
app.secret_key = b'nm5ZzO1SIHg&0G!v49mbDdR7M^vIh&!1'


@app.route("/")
def home():
    return render_template("index.html")


@app.route("/numbers", methods=['GET', 'POST'])
def numbers():
    form = NumbersForm()
    if form.validate_on_submit():
        return redirect(f'/numbers/{form.big_count.data}')
    else:
        return render_template("numbers/form.html", form=form)


@app.route("/numbers/<int:big_count>", methods=['GET'])
def numbers_generate(big_count):
    try:
        target, selection = n.generate(big_count)
    except Exception as error:
        return redirect('/numbers') # TODO should put meaningful message or throw 404, ...?
    return render_template('numbers/index.html', target=target, selection=selection)


@app.route("/letters")
def letters():
    return render_template("letters/index.html")
