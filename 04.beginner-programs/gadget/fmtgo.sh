#!/usr/bin/bash
find . -name "*.go"|xargs sed -i "/^$/d"
find . -name "*.go"|xargs sed -i "s/^func/\nfunc/"
find . -name "*.go"|xargs sed -i "s/^type/\ntype/"
gofmt -s -w .
