
# [phat] -- pomodoro & habit tracker & time logger

[cobra guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)

* pomodoro with user-defined settings
    - defaults: 60/10 x 3 sessions
    - habit: required field (cannot be left blank)
    
* data stored in local CSV file:
    - name of habit             "string"
    - total habit time          "int"
    - habit status              "bool"
        - set minute goal for habit
        - if it reaches the minute goal, status=true

## Base File

"programming",60

programming.goal = 60

if goal met:
    add row to csv_file and add bool

programming,60,true

## Pomodoro function
```bash
0> phat -h programming
```

- starts a pomodoro in the terminal
    -> add habit name to .csv
- counts down the minutes, and seconds
- per session ending, return the time elapsed 
    -> add the time elapsed to .csv file
- run through break function (don't record time)
- until it's the final session
    -> don't run a break if it's the final session
- after the final session, display the stats of the session
    -> data from the .csv file

## Quality of life

- in the user.conf, list habits that can be picked from as defaults
