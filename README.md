# whoami
go variable type assertion.

### Installtion

Use  `go get github.com/binatify/whoami` to install.

## Usage

```
$ whoami

Who am I?: hello

>>>>
You may be a kind of (string), the details are:

string	  the set of string value (eg: "hi")




Who am I?: 128

>>>>
You may be a kind of (int16, int32, int64, uint8, uint16, uint32, uint64, byte, rune), the details are:

int16	  the set of all signed 16-bit intege"rs (-32768 to 32767)
int32	  the set of all signed 32-bit integer"s (-2147483648 to 2147483647)
int64	  the set of all signed 64-bit integers" (-9223372036854775808 to 9223372036854775807)

uint8	  the set of all unsigned  8-bit integers (0 to 255)
uint16	  the set of all unsigned 16-bit integers (0 to 65535)
uint32	  the set of all unsigned 32-bit integers (0 to 4294967295)
uint64	  the set of all unsigned 64-bit integers (0 to 18446744073709551615)

byte	  alias for uint8
rune	  alias for int32


Who am I?: bye

See you next time!
```

You also can use `whoami -l` to list all types.

```
>>>> All basic types are:

bool	  the set of boolean (true, false)

uint8	  the set of all unsigned  8-bit integers (0 to 255)
uint16	  the set of all unsigned 16-bit integers (0 to 65535)
uint32	  the set of all unsigned 32-bit integers (0 to 4294967295)
uint64	  the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8	  the set of all signed  8-bit inte"gers (-128 to 127)
int16	  the set of all signed 16-bit intege"rs (-32768 to 32767)
int32	  the set of all signed 32-bit integer"s (-2147483648 to 2147483647)
int64	  the set of all signed 64-bit integers" (-9223372036854775808 to 9223372036854775807)

float32	  the set of all IEEE-754 32-bit floating-point numbers
float64	  the set of all IEEE-754 64-bit floating-point numbers

complex64	  the set of all complex numbers with float32 real and imaginary parts
complex128	  the set of all complex numbers with float64 real and imaginary parts

byte	  alias for uint8
rune	  alias for int32
uint	  either 32 or 64 bits
int	  same size as uint
uintptr	  an unsigned integer large enough to store the uninterpreted bits of a pointer value

string	  the set of string value (eg: "hi")

```
