#!/usr/local/bin/bash

sum_of_priorities=0

sum_of_badges () {
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

check_for_badge () {
    local member_one=$1
    local member_two=$2
    local member_three=$3

    for (( i=0; i<${#member_one}; i++ )); do
        local char_one=${member_one:$i:1}
        for (( j=0; j<${#member_two}; j++ )); do
            local char_two=${member_two:$j:1}

            if [ "$char_one" == "$char_two" ]; then
                for (( k=0; k<${#member_three}; k++ )); do
                    local char_three=${member_three:$k:1}

                    if [ "$char_one" == "$char_three" ]; then
                        sum_of_badges "$char_one"
                        break 3
                    fi
                done
            fi
        done
    done
}

read_file () {
    while mapfile -t -n 3 group && ((${#group[@]})); do
        check_for_badge "${group[0]}" "${group[1]}" "${group[2]}"
    done < rucksacks.list
}

read_file
echo "Sum of all group badges: $sum_of_priorities"
