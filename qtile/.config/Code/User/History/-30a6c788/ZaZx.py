from routes import blueprint


@blueprint.route('/')
def hello():
    return 'Hello World!'
