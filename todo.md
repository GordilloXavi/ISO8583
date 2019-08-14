find/code package to make the following changes of basis:
    hex <-> bin
    bin <-> dec
    dec <-> hex

Make methods Decode, Encode, AddFields part of message object

Make main.go to pass messages by arguments and do testing

do some checks with the length of the message (first field in the message)

Do propper eror handling to cover all possible exception: 
    Message still has characters left when we have processed all messageField's
    Slice out of range (message not long enough)
    Bad hexadecimal/binary inputs (non hex caracters)
    Bad encoding (careful with ASCII/utf-8)


