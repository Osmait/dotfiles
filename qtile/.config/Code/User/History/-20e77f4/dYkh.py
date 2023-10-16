
from router import blueprint


@blueprint.route("/hello", methods=['GET'])
def hello_world():
    return "<p> Hello, World!</p"
