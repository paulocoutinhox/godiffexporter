# Support with donation
[![Support with donation](http://donation.pcoutinho.com/images/donate-button.png)](http://donation.pcoutinho.com/)

# GoDiffExporter

GoDiffExporter was made to be a simple tool that will parse DIFF files and export to some formats made with Go (Golang).

# Formats

Today this tool support PDF format when export.

# How to use

You can execute on your terminal/console the simple command:

```
git diff > /tmp/diff.txt && godiffexporter -d=/tmp/diff.txt -o=/tmp/diff.pdf
```

Or you can execute command one-by-one:

```
git diff > /tmp/diff.txt
godiffexporter -d=/tmp/diff.txt -o=/tmp/diff.pdf
```

If you are using MAC/OSX you can execute the exporter and open the file directly from terminal:

```
git diff > /tmp/diff.txt && godiffexporter -d=/tmp/diff.txt -o=/tmp/diff.pdf && open /tmp/diff.pdf
```


# Installing

```
go get -u github.com/prsolucoes/godiffexporter
go install github.com/prsolucoes/godiffexporter
```

# Supported By Jetbrains IntelliJ IDEA

![alt text](https://github.com/prsolucoes/goci/raw/master/extras/jetbrains/logo.png "Supported By Jetbrains IntelliJ IDEA")

# Author WebSite

> http://www.pcoutinho.com

# License

MIT