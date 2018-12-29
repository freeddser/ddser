#!/bin/sh


baseDirForScriptSelf=$(cd "$(dirname "$0")"; pwd)

export MY_BASEDIR_VERSION=${baseDirForScriptSelf}/  #for ./share/english/errmsg.sys
export LD_LIBRARY_PATH=${baseDirForScriptSelf}/:${LD_LIBRARY_PATH}  #for libtmysqlparse.so.18  

#printf "./tmysqlparse %s \n" "$@"


#####################################
#to run ./tmysqlparse
####################################
${baseDirForScriptSelf}/tmysqlparse $@
