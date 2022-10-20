Operator	            Name	            Description	                                Example

+	                    Addition	        Adds together two values	                x + y	
-	                    Subtraction	        Subtracts one value from another	        x - y	
*	                    Multiplication	    Multiplies two values	                    x * y	
/	                    Division        	Divides one value by another	            x / y	
%	                    Modulus	            Returns the division remainder	            x % y	
++	                    Increment	        Increases the value of a variable by 1	    x++	
--	                    Decrement	        Decreases the value of a variable by 1	    x--








int data type:

1. Signed Integers
Signed integers, declared with one of the int keywords, can store both positive and negative values:

Type        	Size        	Range

int	Depends on platform:
- 32 bits in 32 bit systems -> -2147483648 to 2147483647 in 32 bit systems
- 64 bit in 64 bit systems  -> -9223372036854775808 to 9223372036854775807 in 64 bit systems

int8	8 bits/1 byte	-128 to 127

int16	16 bits/2 byte	-32768 to 32767

int32	32 bits/4 byte	-2147483648 to 2147483647

int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807






2. Unsigned Integers
Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

Go has five keywords/types of unsigned integers:

Type        	Size	        Range

uint Depends on platform:
- 32 bits in 32 bit systems -> 0 to 4294967295 in 32 bit systems
- 64 bit in 64 bit systems  -> 0 to 18446744073709551615 in 64 bit systems

uint8	8 bits/1 byte	0 to 255

uint16	16 bits/2 byte	0 to 65535

uint32	32 bits/4 byte	0 to 4294967295

uint64	64 bits/8 byte	0 to 18446744073709551615











Bitwise operators are used on (binary) numbers:

Operator	Name	                    Description	                                           Example	

&        	AND	                Sets each bit to 1 if both bits are 1                          	x & y	
|	        OR	                Sets each bit to 1 if one of two bits is 1	                    x | y	
^	        XOR	                Sets each bit to 1 if only one of two bits is 1	                x ^ b	
<<	    Zero fill left shift	Shift left by pushing zeros in from the right	                x << 2	
>>	    Signed right shift  	Shift right by pushing copies of the leftmost bit               x >> 2
                                in from the left, and let the rightmost bits 
                                fall off





Floating point number:
The float data types are used to store positive and negative numbers with a decimal point, 
like 35.3, -2.34, or 3597.34987.

The float data type has two keywords:

Type	    Size	            Range
float32	    32 bits	           -3.4e+38 to 3.4e+38.
float64	    64 bits	           -1.7e+308 to +1.7e+308.





Complex Numbers: 

- The complex numbers are divided into two parts are complex64 & complex128
- float32 and float64 are also part of these complex numbers.

complex64  - Complex numbers which contain float32 as a real and imaginary component.
complex128 - Complex numbers which contain float64 as a real and imaginary component.