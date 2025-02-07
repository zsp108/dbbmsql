#!/usr/bin/sh

#sleep 15
run_min_prop=`grep runMins testoracle_proxy.prop`
run_min_real=${run_min_prop##*=}

total_run_sec=$(( $run_min_real * 60 ))

current_run_sec=0
while [ $current_run_sec -le $total_run_sec ]
do
    if [[ -e run.log ]];then
        is_finished=`grep -oE "Measured tpmC" run.log`
        if [[ -n $is_finished ]];then
            break
        fi
        #current_tpmc=`tail -n 1 run.log`
        current_tpmc=`grep -oE "Term-00, Running Average tpmTOTAL: [0-9]{1,}\.[0-9]{1,}\s*Current tpmTOTAL: [0-9]{1,}\s*Memory Usage: [0-9]*MB / [0-9]*MB" run.log|tail -1`
        echo `date +'%Y-%m-%d %H:%M:%S'`" "$current_tpmc
    fi
    sleep 5
    let current_run_sec=$current_run_sec+5
done


