from flask import abort
from api.common.command import run_command

class SecurityScoreController:

    @staticmethod
    def get_lynis_score():
        comm = run_command("lynis audit system --nocolors")
        if comm is not None:
            text = int(comm.split("Hardening index : ")[1].split(" ")[0])
            return text
        else: return abort(500, "Lynis audit system error")
