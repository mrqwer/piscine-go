#! bin/bash
export interview=$(head -n 179 ./streets/Buckingham_Place | tail -n 1 | sed 's/.*#//')
echo $interview 
cat ./interviews/interview-$interview
echo $MAIN_SUSPECT