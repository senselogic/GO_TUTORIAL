#!/bin/sh
set -x
../../TOOL/GENESIS/genesis --input_folder gs/ --output_folder go/ --recursive --style .gs .go
