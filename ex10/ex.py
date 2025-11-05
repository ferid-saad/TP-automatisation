import subprocess

user = "username1"
count = int(subprocess.check_output(f"ps -u {user} | wc -l", shell=True))
print(f"Processes: {count-1}")