section .data
	msg db "Hello ASM World"

section .text
	global _start

_start:
	mov eax, 1
	mov edi, 1
	mov esi, msg
	mov edx, 15
	syscall

	mov eax, 60
	mov edi, 0
	syscall
