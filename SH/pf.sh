#!/bin/bash

touch $1.py
echo "#!/usr/bin/env python" > $1.py && echo "" >> $1.py
vim $1.py
