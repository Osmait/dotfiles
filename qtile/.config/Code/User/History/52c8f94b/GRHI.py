from flask import Blueprint

from app import db
from django.contrib.auth.models import User


user = Blueprint('user', __name__)


@user.route('/user')
def create_user():
    db.session.add(User(username='XXXX', email='XXXXXXXXXXXXXXXX'))
    db.session.commit()
    return 'User created!'
