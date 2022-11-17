	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 1
	.globl	_main                           ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$16, %rsp
	movl	$0, -4(%rbp)
LBB0_1:                                 ## =>This Inner Loop Header: Depth=1
	leaq	L_.str(%rip), %rdi
	movb	$0, %al
	callq	_printf
	leaq	-5(%rbp), %rsi
	leaq	L_.str.1(%rip), %rdi
	movb	$0, %al
	callq	_scanf
	leaq	-5(%rbp), %rdi
	callq	_longestPalindrome
	movq	%rax, -16(%rbp)
	movq	-16(%rbp), %rsi
	leaq	L_.str.1(%rip), %rdi
	movb	$0, %al
	callq	_printf
	jmp	LBB0_1
	.cfi_endproc
                                        ## -- End function
	.globl	_longestPalindrome              ## -- Begin function longestPalindrome
	.p2align	4, 0x90
_longestPalindrome:                     ## @longestPalindrome
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$80, %rsp
	movq	%rdi, -16(%rbp)
	movq	-16(%rbp), %rdi
	callq	_strlen
                                        ## kill: def $eax killed $eax killed $rax
	movl	%eax, -20(%rbp)
	cmpl	$1, -20(%rbp)
	jg	LBB1_2
## %bb.1:
	movq	-16(%rbp), %rax
	movq	%rax, -8(%rbp)
	jmp	LBB1_19
LBB1_2:
	movq	$0, -48(%rbp)
	movl	$1, -52(%rbp)
	movl	$0, -56(%rbp)
LBB1_3:                                 ## =>This Loop Header: Depth=1
                                        ##     Child Loop BB1_7 Depth 2
	movl	-56(%rbp), %eax
	movl	-20(%rbp), %ecx
	subl	$2, %ecx
	cmpl	%ecx, %eax
	jge	LBB1_18
## %bb.4:                               ##   in Loop: Header=BB1_3 Depth=1
	movq	-16(%rbp), %rax
	movslq	-56(%rbp), %rcx
	addq	%rcx, %rax
	movq	%rax, -32(%rbp)
	movl	$0, -60(%rbp)
	movl	-56(%rbp), %eax
	movl	%eax, -64(%rbp)
	movq	-16(%rbp), %rax
	movslq	-56(%rbp), %rcx
	addq	%rcx, %rax
	addq	$1, %rax
	movq	%rax, -40(%rbp)
	movq	-16(%rbp), %rax
	movslq	-56(%rbp), %rcx
	movsbl	(%rax,%rcx), %eax
	movq	-16(%rbp), %rcx
	movl	-56(%rbp), %edx
	addl	$2, %edx
	movslq	%edx, %rdx
	movsbl	(%rcx,%rdx), %ecx
	cmpl	%ecx, %eax
	jne	LBB1_6
## %bb.5:                               ##   in Loop: Header=BB1_3 Depth=1
	movq	-40(%rbp), %rax
	addq	$1, %rax
	movq	%rax, -40(%rbp)
	movl	-56(%rbp), %eax
	addl	$1, %eax
	movl	%eax, -64(%rbp)
	movl	$1, -60(%rbp)
LBB1_6:                                 ##   in Loop: Header=BB1_3 Depth=1
	movl	$0, -68(%rbp)
LBB1_7:                                 ##   Parent Loop BB1_3 Depth=1
                                        ## =>  This Inner Loop Header: Depth=2
	xorl	%eax, %eax
                                        ## kill: def $al killed $al killed $eax
	movl	-68(%rbp), %ecx
	cmpl	-64(%rbp), %ecx
	movb	%al, -69(%rbp)                  ## 1-byte Spill
	jge	LBB1_9
## %bb.8:                               ##   in Loop: Header=BB1_7 Depth=2
	movl	-64(%rbp), %eax
	addl	-68(%rbp), %eax
	cmpl	-20(%rbp), %eax
	setl	%al
	movb	%al, -69(%rbp)                  ## 1-byte Spill
LBB1_9:                                 ##   in Loop: Header=BB1_7 Depth=2
	movb	-69(%rbp), %al                  ## 1-byte Reload
	testb	$1, %al
	jne	LBB1_10
	jmp	LBB1_14
LBB1_10:                                ##   in Loop: Header=BB1_7 Depth=2
	movq	-32(%rbp), %rax
	movsbl	(%rax), %eax
	movq	-40(%rbp), %rcx
	movsbl	(%rcx), %ecx
	cmpl	%ecx, %eax
	jne	LBB1_12
## %bb.11:                              ##   in Loop: Header=BB1_7 Depth=2
	xorl	%eax, %eax
	movl	%eax, %ecx
	movl	-68(%rbp), %eax
	addl	$1, %eax
	movl	%eax, -68(%rbp)
	movq	-32(%rbp), %rax
	movslq	-68(%rbp), %rdx
	subq	%rdx, %rcx
	addq	%rcx, %rax
	movq	%rax, -32(%rbp)
	movq	-32(%rbp), %rax
	movslq	-68(%rbp), %rcx
	addq	%rcx, %rax
	movq	%rax, -40(%rbp)
	jmp	LBB1_13
LBB1_12:                                ##   in Loop: Header=BB1_3 Depth=1
	jmp	LBB1_14
LBB1_13:                                ##   in Loop: Header=BB1_7 Depth=2
	jmp	LBB1_7
LBB1_14:                                ##   in Loop: Header=BB1_3 Depth=1
	movl	-68(%rbp), %eax
	shll	$1, %eax
	addl	-60(%rbp), %eax
	movl	%eax, -60(%rbp)
	movl	-60(%rbp), %eax
	cmpl	-52(%rbp), %eax
	jle	LBB1_16
## %bb.15:                              ##   in Loop: Header=BB1_3 Depth=1
	movq	-32(%rbp), %rax
	movq	%rax, -48(%rbp)
	movl	-60(%rbp), %eax
	movl	%eax, -52(%rbp)
LBB1_16:                                ##   in Loop: Header=BB1_3 Depth=1
	jmp	LBB1_17
LBB1_17:                                ##   in Loop: Header=BB1_3 Depth=1
	movl	-56(%rbp), %eax
	addl	$1, %eax
	movl	%eax, -56(%rbp)
	jmp	LBB1_3
LBB1_18:
	movl	-52(%rbp), %esi
	leaq	L_.str.2(%rip), %rdi
	movb	$0, %al
	callq	_printf
	movq	-48(%rbp), %rax
	movq	%rax, -8(%rbp)
LBB1_19:
	movq	-8(%rbp), %rax
	addq	$80, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"input string:\n"

L_.str.1:                               ## @.str.1
	.asciz	"%s"

L_.str.2:                               ## @.str.2
	.asciz	"length %d\n"

.subsections_via_symbols
