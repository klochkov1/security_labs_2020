from sys import argv
from base64 import b64decode, b64encode
from Crypto.Cipher import AES
from Crypto.Util import Counter

def main():
    if len(argv) < 6:
        print(f"Usage: {argv[0]} [encode|decode] SOURCE KEY VECTOR RESULT")
        exit(1)

    action, source, key, vector, output = argv[1:6]

    with open(source, "rb") as f:
        text = f.read()

    counter = Counter.new(128, initial_value=int(vector.encode("utf-8").hex(), 16))
    cypher = AES.new(key.encode("utf8"), AES.MODE_CTR, counter=counter)  
     
    if action == 'decode':
        result = cypher.decrypt(b64decode(text))
    elif action == 'encode':
        result = b64encode(cypher.encrypt(text))
    else:
        print("Invalid action!")
        exit(1)

    with open(output, "wb") as text_file:
        text_file.write(result)

if __name__ == '__main__':
    main()
