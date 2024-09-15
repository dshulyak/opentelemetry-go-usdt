Demo of integration USDT with golang's opentelemtry library
===

You will need to [install bpftrace](https://github.com/bpftrace/bpftrace/blob/master/INSTALL.md).

Compile a demo:

```sh
go build -o usdtdemo ./demo
```

Launch bpftrace_usdtdemo.sh:

```sh
sudo ./bpftrace_usdtdemo.sh
```

It will attach to entrypoints to  the followinng entrypoints, you can examine them withh `readelf usdtdemo -a`.

```
Displaying notes found in: .note.stapsdt
  Owner                Data size 	Description
  stapsdt              0x00000055	NT_STAPSDT (SystemTap probe descriptors)
    Provider: stacks_tracing
    Name: enter
    Location: 0x000000000057a7b7, Base: 0x0000000000767a40, Semaphore: 0x0000000000000000
    Arguments: -8@%r12 -8@%r15 -8@%r14 -8@%rbx -8@%rax
  stapsdt              0x00000034	NT_STAPSDT (SystemTap probe descriptors)
    Provider: stacks_tracing
    Name: exit
    Location: 0x000000000057a810, Base: 0x0000000000767a40, Semaphore: 0x0000000000000000
    Arguments: -8@%rdi
```

And on launching `./usdtdemo`, you will see captured output such as:

```
Attaching 2 probes...
enter abefdf2d245cc111 0 0 0 parent
enter c8ffb91b146ffcdc abefdf2d245cc111 0 0 child
exit c8ffb91b146ffcdc
exit abefdf2d245cc111
```

The format is made compatible with my [bpf based profiler](https://github.com/dshulyak/stacks).