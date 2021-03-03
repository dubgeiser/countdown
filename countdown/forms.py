from flask_wtf import FlaskForm
from wtforms.fields import IntegerField, StringField
from wtforms.validators import InputRequired, NumberRange, Length
from wtforms.validators import ValidationError


class NumbersForm(FlaskForm):
    big_count = IntegerField(
            'big_count', validators=[InputRequired(), NumberRange(0, 4)])


class LettersAnswerForm(FlaskForm):
    answer = StringField('answer', validators=[InputRequired(), Length(2, 9)])

    def __init__(self, selection):
        super().__init__()
        self.selection = selection.copy()

    def validate_answer(self, answer):
        ''' Validate the given answer:
            - Are all the letters in the selection of vowels and consonants?
            - TODO: And if so; is the answer a word that exists?
        '''
        for letter in answer.data:
            if letter not in self.selection:
                raise ValidationError('Uses letters that were not selected')
            self.selection.remove(letter)
        return True
