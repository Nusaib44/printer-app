#!/bin/bash

# Run the application based on the OS
case "$(uname)" in
  Linux)
    ./build/printer-manager-linux
    ;;
  Darwin)
    ./build/printer-manager-darwin
    ;;
  MINGW*|MSYS*|CYGWIN*)
    ./build/printer-manager-windows.exe
    ;;
  *)
    echo "Unsupported OS"
    exit 1
    ;;
esac
