#!/bin/bash

set -e

for i in pkg/scap/models/linux_sc/models.go pkg/scap/models/unix_sc/models.go; do
	dir=$(dirname $i)
	name=$(basename $(dirname $i))
	if [[ $name == oval_sc ]]; then
		continue
	fi
	mv $i pkg/scap/models/oval_sc/$name.go
	sed -i -r 's/^(package) [a-z]+_sc$/\1 oval_sc/' pkg/scap/models/oval_sc/$name.go
	sed -i 's/oval_sc\.//' pkg/scap/models/oval_sc/$name.go
	sed -i '/github.com\/0intro\/scap\/pkg\/scap\/models\/oval_sc/d' pkg/scap/models/oval_sc/$name.go
	rmdir $dir
done
