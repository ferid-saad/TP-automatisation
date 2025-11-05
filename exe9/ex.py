import os

os.system('(crontab -l ; echo "0 2 * * * /path/to/script.sh") | crontab -')