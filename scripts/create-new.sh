#!/bin/bash

# Constants
declare -r LeetcodeDirName="Leetcode";
declare -r SampleDirName="sample";

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd );

while getopts n:t: flag
do
    case "${flag}" in
        n) name=${OPTARG};;
        t) type=${OPTARG};;
    esac
done

if [ -z ${name+x} ]; then
	echo "Name is unset. Set it with -n";
	exit 1;
fi

if [ -z ${type+x} ]; then
	echo "Type is unset. Set it with -t";
	exit 1;
fi

echo $SCRIPT_DIR;
echo "Name: $name";
echo "Type: $type";

originalDir="$SCRIPT_DIR/../$SampleDirName/$type";
newDir="$SCRIPT_DIR/../$LeetcodeDirName/$name";

cp -r $originalDir $newDir
