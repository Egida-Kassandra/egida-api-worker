from flask import abort
from api.common.command import run_command

class SecurityScoreController:

    @staticmethod
    def get_lynis_score():
        comm = run_command("lynis audit system")
        return abort(500, "Lynis audit system error") if comm is None else comm
