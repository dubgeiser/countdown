from flask_wtf import FlaskForm
from wtforms import IntegerField
from wtforms.validators import InputRequired, NumberRange


class NumbersForm(FlaskForm):
    big_count = IntegerField(
            'big_count', validators=[InputRequired(), NumberRange(0, 4)])
