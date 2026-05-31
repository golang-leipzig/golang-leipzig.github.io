# Synopsis

Run N windows in parallel, record the session.

## Task

Given a config file like synopsis.conf, we want to record with asciinema a tmux
based terminal session, where for each line in the file, we are opening a tmux
window and run the command from one line.

We want record the session and when the commands are all done or after some
timeout, we are stopping the session.

So I imagine the outline to be:

* read @synopsis.conf and determine how many non empty lines we have (N)
* asciinema rec TEMPPATH
* tmux # create N panes, in some sensible layout
* for each command, choose one pane, enter the text, hit enter, move to the next window, enter command, enter, and so on
* wait for a fixed amount of time, e.g. 60s
* close all panes, close tmux
* stop asciinema recording
* move recording from TEMPPATH to synopsis-recording-DATETIME.cast

## Example Output

![](synopsis-recording-20260529-123126-github-light-O3.gif)
