from io import StringIO
import sys

def decode_binary_string(s):
    return ''.join(chr(int(s[i*8:i*8+8],2)) for i in range(len(s)//8))

raw_binary = str(input("Binary: "))

binary_list = raw_binary.split()
print("Decoded String:",decode_binary_string(raw_binary.replace(" ", "")))
binary_formatted_array = [[0]*8 for _ in range(len(binary_list))]

for byteIndex, byte in enumerate(binary_list):
    for bitIndex, bit in enumerate(byte):
        if bit != '0':
            binary_formatted_array[byteIndex][bitIndex] = pow(2, 7-int(bitIndex))

print(binary_formatted_array)