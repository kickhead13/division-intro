import os

password = os.environ["SUPER_SECRET_PASSWORD"]

with open("shadows.txt", "r") as f:
    sys_password = f.readline().rstrip()
    if password != sys_password:
        print("ERROR")
    else:
        print("PASSWORDS MATCH")
