from flask import Flask
from user.view import user


app = Flask(__name__)


app.register_blueprint(user, url_prefix="/api")


if __name__ == "__main__":
    app.run(debug=True)
