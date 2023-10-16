from flask import Flask
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from user.views import user


app = Flask(__name__)

app.register_blueprint(user)

engine = create_engine('sqlite:///example.db')
Session = sessionmaker(bind=engine)


if __name__ == "__main__":
    app.run(debug=True)
