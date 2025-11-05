import shutil

try:
    shutil.make_archive("/tmp/backup", 'zip', "/etc")
    print(f"Zip archive created successfully at: /tmp/backup.zip")
except Exception as e:
    print(f"Error creating zip archive: {e}")