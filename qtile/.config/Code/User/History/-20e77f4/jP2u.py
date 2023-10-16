
from router import blueprint


@blueprint.route("/", methods=['GET'])
def hello_world():
    return "<p> Hello, World!</p"
