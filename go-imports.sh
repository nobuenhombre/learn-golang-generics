#!/bin/bash
goimports -v -w $(go list -f {{.Dir}} ./...)
