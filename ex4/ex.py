import shutil

usage = shutil.disk_usage('/')
if usage.free / usage.total < 0.1:
    print("Espace disque faible (<10%)")
else:
    print("Espace disque OK")