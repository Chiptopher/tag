#!/bin/bash
#
# MIT License
# 
# Copyright (c) 2018 Christopher M. Boyer
# 
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
# 
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
# 
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
# 
# Build tag executable for all supported files.

###############################################################################
# Build a tag executable for the given operating system and architecture.
# Globals:
#   None
# Arguments:
#   OS
#   ARCH
#   BIN
# Returns:
#   None
###############################################################################
build_executable()
{
    OS=$1
    ARCH=$2
    BUILD_DIR="${3}/${1}/$ARCH"
    BIN="tag"
    if [ $1 = "windows" ]; then
        BIN="$BIN.exe"
    fi
    echo "Building $1:$2 -- $BIN"
    GOOS=$1 GOARCH=$2 go build -o $BUILD_DIR/$BIN ../main.go
    start=$PWD
    cd $BUILD_DIR
    zip "tag.zip" $BIN
    rm $BIN
    cd $start
}

VERSION=$1

if [ ! -d "../bin/stable/" ]; then
    echo ""
    echo "Building stable directory."
    mkdir ../bin/stable
fi

if [ -z $VERSION ]; then
    echo ""
    echo "Version number required.";
else

    DIRECTORY="../bin/$1"
    if [ -d "$DIRECTORY" ]; then
        echo ""
        echo "Build version $VERSION already exists."
    else

        echo ""

        mkdir $DIRECTORY

        build_executable darwin amd64 $DIRECTORY   # macOS 64-bit
        build_executable darwin 386 $DIRECTORY     # macOS 32-bit
        build_executable linux amd64 $DIRECTORY    # linux 64-bit
        build_executable linux 386 $DIRECTORY      # linux 32-bit
        build_executable windows amd64 $DIRECTORY  # windows 64-bit
        build_executable windows 386 $DIRECTORY    # windows 32-bit

        echo ""
        echo "Build Complete."
        echo ""

        LATEST_FLAG=$2
        LATEST_VERSION=$3
        if  [ ! -z $LATEST_FLAG ] && [ $LATEST_FLAG == "--stable" ] &&
            [ ! -z $LATEST_VERSION ] && [ -d "../bin/$LATEST_VERSION" ]; then
            echo "Changing the latest directory to version $LATEST_VERSION."
            echo ""
            rm -rf ../bin/stable/*
            cp -r ../bin/$LATEST_VERSION/* ../bin/stable
        fi
    fi
fi
