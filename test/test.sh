#!/bin/bash

script_dir=$(cd $(dirname $0); pwd)
test_dirs="${script_dir}/test-dirs"
jsons="${script_dir}/jsons"
bin="${script_dir}/dir-struct-check"

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
test1_dir="${test_dirs}/test1"
assert `$bin $test1_dir < "${jsons}/1/true.json"` "true"
assert `$bin $test1_dir < "${jsons}/1/false1.json"` "false"
assert `$bin $test1_dir < "${jsons}/1/false2.json"` "false"
assert `$bin $test1_dir < "${jsons}/1/false3.json"` "false"

# test2
test2_dir="${test_dirs}/test2"
assert `$bin $test2_dir < "${jsons}/2/true.json"` "true"
assert `$bin $test2_dir < "${jsons}/2/false1.json"` "false"
assert `$bin $test2_dir < "${jsons}/2/false2.json"` "false"
assert `$bin $test2_dir < "${jsons}/2/false3.json"` "false"
