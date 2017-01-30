#!/bin/bash
sudo cp ./bin/fwew /usr/local/bin && \
echo "Fwew installed to /usr/local/bin/fwew" && \
cp -r ./.fwew $HOME && \
echo "Fwew data installed to $HOME/.fwew"
