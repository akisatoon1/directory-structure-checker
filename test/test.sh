#!/bin/bash

wd=`pwd`
test_dirs="${wd}/test-dirs"
jsons="${wd}/jsons"
bin="${wd}/dir-struct-check"

# assert(output str, expected str)
function assert() {
    output=$1
    expected=$2
    if [ "${output}" = "${expected}" ]; then
        echo PASS
    else
        echo ERROR
        echo "expected is ${expected}"
    fi
}

# test1
cd "${test_dirs}/test1"
assert `$bin < "${jsons}/1/true.json"` "true"
assert `$bin < "${jsons}/1/false1.json"` "false"
assert `$bin < "${jsons}/1/false2.json"` "false"
assert `$bin < "${jsons}/1/false3.json"` "false"

# test2
cd "${test_dirs}/test2"
assert `$bin < "${jsons}/2/true.json"` "true"
assert `$bin < "${jsons}/2/false1.json"` "false"
assert `$bin < "${jsons}/2/false2.json"` "false"
assert `$bin < "${jsons}/2/false3.json"` "false"
