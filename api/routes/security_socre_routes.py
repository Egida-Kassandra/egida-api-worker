from flask_cors import cross_origin
from flask import jsonify
from flask import current_app as app
from api.controllers.security_score_controller import SecurityScoreController


@app.route('/api/lynis', methods=['GET'])
@cross_origin()
def get_lynis_score():
    return jsonify(SecurityScoreController.get_lynis_score())
