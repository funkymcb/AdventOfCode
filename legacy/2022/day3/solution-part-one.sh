#!/bin/bash

sum_of_priorities=0

split_rucksack () {
    local rucksack="$1"
    COMPARTMENT_ONE=${rucksack:0:${#rucksack}/2}
    COMPARTMENT_TWO=${rucksack:${#rucksack}/2}
}

sum_of_items () {
    local priority
    char_value=$(printf "%d" "'$1")

    if [[ "$1" =~ [a-z] ]]; then
        priority=$(( char_value - 96 ))
        echo "$1 is duplicate. Priority: $priority"
    elif [[ "$1" =~ [A-Z] ]]; then
        priority=$(( char_value - 38 ))
        echo "$1 is duplicate. Priority: $priority"
    fi

    sum_of_priorities=$(( sum_of_priorities + priority ))
}

check_dulicate_items () {
    for (( i=0; i<${#COMPARTMENT_ONE}; i++ )); do
        local char_one=${COMPARTMENT_ONE:$i:1}

        for (( j=0; j<${#COMPARTMENT_TWO}; j++ )); do
            local char_two=${COMPARTMENT_TWO:$j:1}

            if [ "$char_one" == "$char_two" ]; then
                sum_of_items "$char_one"
                break 2
            fi
        done
    done
}

read_file () {
    while read -r rucksack; do
        split_rucksack "$rucksack"
        check_dulicate_items
    done < rucksacks.list
}

read_file
echo "Sum of priorities $sum_of_priorities"
