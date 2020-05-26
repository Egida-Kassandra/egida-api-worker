from flask_cors import cross_origin
from flask import jsonify
from flask import current_app as app
from api.controllers.os_info_controller import OSInfoController


@app.route('/api/info', methods=['GET'])
@cross_origin()
def get_os_info():
    return jsonify(OSInfoController.get_os_info())
