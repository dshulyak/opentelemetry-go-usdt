#!/bin/bpftrace

usdt:./usdtdemo:stacks_tracing:enter
{
    printf("enter %lx %lx %lx %lx %s\n", arg0, arg1, arg2, arg3, str(arg4));
}

usdt:./usdtdemo:stacks_tracing:close
{   
    printf("exit %lx\n", arg0);
}
