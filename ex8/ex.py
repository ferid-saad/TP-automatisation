import os

files = [f for f in os.listdir('.') if os.path.isfile(f)]
files_with_sizes = [(f, os.path.getsize(f)) for f in files]
files_with_sizes.sort(key=lambda x: x[1], reverse=True)
for f, size in files_with_sizes[:5]:
    print(f"{f}: {size/1024:.2f} KB")