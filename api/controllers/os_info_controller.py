import platform

from flask import abort

class OSInfoController:

    @staticmethod
    def get_os_info():
        os = platform.system()
        release = platform.release()
        version = platform.version()
        distribution = "N/A"
        machine = platform.machine()
        uname = platform.uname()
        try:
            distribution = platform.linux_distribution()
        except:
            distribution = "N/A"
        result = {
            "os": os,
            "release": release,
            "version": version,
            "distribution": distribution,
            "machine": machine,
            "uname": uname
        }
        return result
