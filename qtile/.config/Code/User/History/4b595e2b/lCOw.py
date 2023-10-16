from flask import Flask
from infraestruture.controller import user
from flask_sqlalchemy import SQLAlchemy


dbUser = "osmait"
dbPassword = "admin123"
dbName = "my_store"
dbHost = "localhost"
dbPort = 5432

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = f"postgres://{dbUser}:{dbPassword}@{dbHost}:{dbPort}/{dbName}"
db = SQLAlchemy(app)


app.register_blueprint(user.user, url_prefix='/api')
with app.app_context():
    db.create_all()

if __name__ == '__main__':
    app.run(debug=True)
