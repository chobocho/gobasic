# gobasic
GW-BASIC interpreter written by Go

### 1차 목표
1. +, -, *, /, % 연산 지원  
2. () 지원  


### EBNF
~~~
(* EBNF for a calculator that supports +, -, *, /, %, (), and unary minus. *)

expression   ::= term { ("+" | "-") term }
term         ::= factor { ("*" | "/" | "%") factor }
factor       ::= primary | "-" primary
primary      ::= number | "(" expression ")"
number       ::= digit { digit }
digit        ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
~~~