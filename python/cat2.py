#The following command will copy the contents of input_file to output.txt(if not create a new)  :

import argparse
import os

my_parser=argparse.ArgumentParser(prog='cat',description='Copy file1 content to file2')

my_parser.add_argument('Path',metavar='Path',type=str,help='Name of file',nargs='+')
my_parser.add_argument('-a',help="Append to outputfile")
my_parser.add_argument('-o',help='Overwrite the outputfile ')

args=my_parser.parse_args()

print(vars(args))
#print(args.Path)


"""input_file=open(args.Path[0],'r')
out_file=open(args.o,'w')

for text in input_file:
    out_file.write(text)
    print(text)
out_file.close()
input_file.close()
f=open(args.o,'r')
for text in f:
    print('o')
    print(text)
print('done')"""

with open(args.Path[0],'r') as input_file:
    if args.a:
        with open(args.a,'a') as out_file:
            for text in input_file:
                out_file.write(text)
            out_file.write('------------------')
    else :
        with open(args.o,'w') as out_file:
            for text in input_file:
                out_file.write(text)
            out_file.write('--------------------')