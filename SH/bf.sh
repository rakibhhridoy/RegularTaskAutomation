#!/bin/bash

touch $1.sh
echo "#!/bin/bash" > $1.sh && echo "" >> $1.sh
vim $1.sh
