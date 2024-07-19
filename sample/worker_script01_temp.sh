#!/bin/bash

# Change to the directory of this script
scriptDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
echo "Script workerServer01 current directory is: $scriptDir"

parentDir="$( cd "$( dirname "$scriptDir" )" &> /dev/null && pwd )"
echo "Script workerServer01 parent directory is: $parentDir"

binPath="$parentDir/bin"
configPath="$parentDir/sample"
workerConfig="$configPath/workerServer01.toml"

# get os kernel name
KERNEL=$(uname -s)

case $KERNEL in
Linux) 	echo "Detecting operating system: Linux";;
Darwin) echo "Detecting operating system:  Darwin (macOS)";;
  *)    echo "Unsupported operating system: $KERNEL"; exit 1;;
esac

if [ "$KERNEL" == "Linux" ]; then
  export LD_LIBRARY_PATH=LIBRARY_DIR_VAR
  echo "Setting  workerServer01 script LD_LIBRARY_PATH to: $LIBRARY_PATH_VAR"
elif [ "$KERNEL" == "Darwin" ]; then
  export DYLD_LIBRARY_PATH=LIBRARY_DIR_VAR
  echo "Setting  workerServer01 script DYLD_LIBRARY_PATH to: $LIBRARY_PATH_VAR"
fi

echo "Starting workerServer01 script Command [$binPath/dbms-worker --config $workerConfig]..."
"$binPath"/dbms-worker --config "$workerConfig"