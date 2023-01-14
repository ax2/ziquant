# ziquant
Go语言写的一些量化相关工具

## install

go install github.com/ax2/ziquant

安装完成后，假设go的bin目录已经在PATH中，运行ziquant命令可以看到帮助信息，例如：

```bash
some usefull quant utilities

Usage:
  ziquant [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fund        ziquant fund utilities
  help        Help about any command

Flags:
  -h, --help   help for ziquant

Use "ziquant [command] --help" for more information about a command.
```

## examples

- ziquant fund fundlist  </br>
  获取基金列表
