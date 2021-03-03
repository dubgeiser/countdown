from flask import Flask
from flask import flash
from flask import redirect
from flask import render_template
from flask import request
from flask import session

from countdown import letters as l
from countdown import numbers as n
from countdown.forms import LettersAnswerForm, NumbersForm


app = Flask(__name__)
app.secret_key = b'nm5ZzO1SIHg&0G!v49mbDdR7M^vIh&!1'


@app.route('/')
def home():
    return render_template("index.html")


@app.route('/numbers', methods=['GET', 'POST'])
def numbers():
    form = NumbersForm()
    if form.validate_on_submit():
        return redirect(f'/numbers/{form.big_count.data}')
    else:
        return render_template("numbers/form.html", form=form)


@app.route('/numbers/<int:big_count>', methods=['GET'])
def numbers_generate(big_count):
    try:
        target, selection = n.generate(big_count)
    except Exception as error:
        # TODO Do not blindly put exception messages in UI
        flash(str(error))

        return redirect('/numbers')
    return render_template('numbers/index.html', target=target, selection=selection)


@app.route('/letters', methods=['GET'])
def letters():
    ''' Default entry point for the letters part.
    This resets the session for the letters game.
    '''
    session['game'] = l.new_game()
    return render_template('letters/index.html', selection=session['game']['selection'])


@app.route('/letters/pick/<string:vowel_or_consonant>', methods=['GET'])
def letters_pick(vowel_or_consonant=None):
    ''' Choose either a consonant or a vowel '''
    if vowel_or_consonant is None:
        return redirect('/letters')
    try:
        l.pick(vowel_or_consonant, session['game'])
        session.modified = True
    except Exception as error:
        flash(str(error))
        return redirect('/letters')
    if len(session['game']['selection']) == 9:
        return redirect('/letters/play')
    return render_template('letters/index.html', selection=session['game']['selection'])


@app.route('/letters/play', methods=['GET', 'POST'])
def letters_play():
    ''' All letters are chosen, let's play.
    When finished; fill in your answer in the form.
    '''
    form = LettersAnswerForm(session['game']['selection'])
    if form.validate_on_submit():
        flash(f'{form.data["answer"]} for {len(form.data["answer"])} points.')
        return redirect('/letters')
    return render_template('letters/answer.html',
            selection=session['game']['selection'],
            form=form)
