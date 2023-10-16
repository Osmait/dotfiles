from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from user.view import user


app = Flask(__name__)


app.register_blueprint(user, url_prefix="/api")

app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:////tmp/test.db'
db = SQLAlchemy(app)

db.create_all()


if __name__ == "__main__":
    app.run(debug=True)
