#!/bin/bash

if [ -n "$1" -o -n "$2" -o -n "$3" ]
then
tmysqlparse=$1
filename=$2
logfile=$3
else
echo "params error"
exit
fi

#echo ${tmysqlparse}/tmysqlparse.sh -c utf8 -f ${logfile}.xml  $filename
${tmysqlparse}/tmysqlparse.sh -c utf8 -f ${logfile}.xml < $filename > /dev/null  2>&1 

#echo ${logfile}.xml
if [ -f ${logfile}.xml ]
then
errcount=`cat ${logfile}.xml|grep syntax_fail|wc -l`

if [ $errcount -eq 0 ]
then
echo 0
else
echo 1
fi
else
echo "${logfile}.xml" is not hava
fi
