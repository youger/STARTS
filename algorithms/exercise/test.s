.section __DATA
    mynum:
    .int 8
    mygs:
    .asciz "%x----%x----%x\n"
.section __TEXT
    .globl main
    main:
        leaq mynum,%eax #将mynum地址复制到%eax
        movl (%eax),%ebx#将%eax内地址所指的内容复制到%ebx
        movl mynum,%ecx#将mynum内容复制到%ecx中
        pushq %ecx
        pushq %ebx
        pushq %eax
        pushq $mygs
        callq printf
        pushq $0
        callq exit
