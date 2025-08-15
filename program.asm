section .data
    x   dq 0       ; 8-Byte Variable, initial 0

section .text
    global _start

_start:
    mov     qword [x], 5    ; x = 5

    ; Programm beenden (exit(0))
    mov     rax, 60         ; syscall: exit
    xor     rdi, rdi        ; return code 0
    syscall
