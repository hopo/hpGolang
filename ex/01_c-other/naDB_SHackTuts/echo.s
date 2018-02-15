section .text
	global _start

_start:
	xor rax, rax ;sys_read (init)
	mov rbx, rax
	mov rax, rax
	mov rdx, rax

	sub rsp, 64	;allocate 64byte
	mov rdi, 0
	mov rsi, rsp
	mov rdx, 63 ; ~63
	syscall

	mov rax, 1 ;sys_write
	mov rdi, 1
	mov rsi, rsp
	mov rdx, 63
	syscall

	mov rax, 60
	syscall
