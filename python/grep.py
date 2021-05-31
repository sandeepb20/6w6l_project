import argparse
import os
import sys

my_persor=argparse.ArgumentParser(prog='grep',description='Search given string in path file')

my_persor.add_argument('string',metavar='str',action='store',help='String to be searched')
my_persor.add_argument('-path',help='Path to search for str')
my_persor.add_argument('-case',action='store_true',help='search case sensitive')
my_persor.add_argument('-n',action='store_true',help='Print total no str found')
my_persor.add_argument('-v',action='store_true',help='display lines that do not match')

args=my_persor.parse_args()

print(vars(args))


with open(args.path,'r') as input_file:

    count=0
    if args.case:
        print('yes')
        for text in input_file:
            found=0
            word_list=text.split()
            for word in word_list:
                word.upper()
                args.string.upper()
                if(word.upper()==args.string.upper()):
                    found+=1
            if found:
                if not args.v:
                    print(text)
                    count+=1
            else:
                if args.v:
                    print(text)
                    count+=1
    else :
        for text in input_file:
            found=0
            word_list=text.split()
            for word in word_list:
                if(word==args.string):
                    found+=1
            if found:
                if not args.v:
                    print(text)
                    count+=1
            else:
                if args.v:
                    print(text)
                    count+=1
    if args.n:
        print("Total no search found=",count)
