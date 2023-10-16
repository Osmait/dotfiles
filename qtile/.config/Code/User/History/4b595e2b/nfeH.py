from flask import Flask
from user import blueprint


app = Flask(__name__)

app.register_blueprint(blueprint, url_prefix="/api")


if __name__ == "__main__":
    app.run(debug=True)
