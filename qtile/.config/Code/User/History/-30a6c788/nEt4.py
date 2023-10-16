from routes import blueprint


@blueprint.route('/', methods=['GET'])
def hello():
    return 'Hello World!'
