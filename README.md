# Pad Application
This application pads any whole numbers found in a string to the size specified by the user.

## Build

simply run the **build.sh** script. It first runs unit test and if successful it will build the application and produces the **pad** binary.

## Usage

*./pad \<string to be padded\> \<length of the pad\>*

## Assumptions

1. Zero is considered a digit and part of the whole number even if it is at the beginning of the number thus *./pad test01 4* would have the following output **test0001**

1. This application only considers natural numbers e.g. 3.14 would be considered two whole numbers 3 and 14 divided by '.'

## Runtime Analysis

The runtime complexity of this application is **O(n + x*l)** where:

1. **n** is the length of the string

1. **x** is the number of whole numbers within the string

1. **l** is the padding length

it is safe to say that if **l** is not a large number and we dont have that many whole numbers in the string the runtime would be **O(n)**
