import shlex
import subprocess
from subprocess import Popen


def run_command(command):
    process = Popen(shlex.split(command), stdout=subprocess.PIPE, universal_newlines=True)
    text = ""
    while True:
        output = process.stdout.readline()
        if process.poll() is not None:
            break
        if output:
            text = text + output.strip()+"\n"
    rc = process.poll()
    return text
