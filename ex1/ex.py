import os, time

log_dir = "/var/log/myapp"
now = time.time()
for f in os.listdir(log_dir):
    file_path = os.path.join(log_dir, f)
    if os.path.isfile(file_path) and now - os.path.getmtime(file_path) > 7 * 86400:
        os.remove(file_path)
        print(f"Deleted {file_path}")