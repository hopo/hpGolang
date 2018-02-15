section .data
	STAR db '*'
	EMPTY db 0x0a	;\n

section .text
	global _start

_start:
	mov rax, 1	;WRITE system call set
	mov rdi, 1	;standart print mode
	mov rdx, 1	;print length
	mov r10, 0	;role index
	mov r9, [rsp + 16]	;

	cmp r9, 0	;noinput
	je _done	;prgm break

	mov cl, [r9]	;store in cl. only 1 byte
	movzx r9, cl	
	sub r9, 0x30	;index

	mov r8, r9
	xor r9, r9
	call _syscall

_small:
	cmp r10, r9
	je _up
	mov rsi, STAR	;print '*'
	syscall

	mov rax, 1	;WRITE system call
	inc r10	;j++
	jmp _small	;again

_up:
	cmp r9, r8	;case i == n
	je _down	;down function
	mov rsi, EMPTY	;'\n'
	syscall	;print function exec

	mov rax, 1
	mov r10, 0
	add r9, 1
	jmp _small

_down:
	cmp r9, 0	;if i == 0
	je _done;
	mov rsi, EMPTY	;'\n'
	syscall
	mov rax, 1	;WRITE system call
	mov r10, 0	;init j
	sub r9, 1	;i--
	jmp _big	;again print

_big:
	cmp r10, r9	;i == j
	je	_down
	mov rsi, STAR	;print '*'
	syscall

	mov rax, 1
	inc r10
	jmp _big

_done:
	mov rax, 60
	mov rdi, 0
	syscall

_syscall:
	syscall
	ret
