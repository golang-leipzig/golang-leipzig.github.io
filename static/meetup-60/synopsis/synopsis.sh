#!/usr/bin/env bash
# synopsis: record a tmux session running N commands in parallel panes.
#
# Usage: ./synopsis.sh [config] [timeout-seconds] [layout] [wait]
#   config:  file with one command per line (default: synopsis.conf)
#   timeout: max recording duration in seconds (default: 60)
#   layout:  tmux layout — tiled, even-horizontal, even-vertical,
#            main-horizontal, main-vertical (default: tiled)
#   wait:    "1" to run the full timeout instead of exiting when all
#            commands finish (default: 0)

set -euo pipefail

CONFIG="${1:-synopsis.conf}"
TIMEOUT="${2:-60}"
LAYOUT="${3:-tiled}"
WAIT="${4:-0}"
START_DELAY=2

[[ -f "$CONFIG" ]] || { echo "missing config: $CONFIG" >&2; exit 1; }

mapfile -t CMDS < <(grep -vE '^\s*(#|$)' "$CONFIG")
N="${#CMDS[@]}"
(( N > 0 )) || { echo "no commands in $CONFIG" >&2; exit 1; }

SESSION="synopsis-$$"
TMPCAST=".${SESSION}.cast"
FINAL="synopsis-recording-$(date +%Y%m%d-%H%M%S).cast"

# allow tmux attach to work even from inside an existing tmux
unset TMUX

echo "synopsis: $N commands, timeout ${TIMEOUT}s"

cleanup() { tmux kill-session -t "$SESSION" 2>/dev/null || true; }
trap cleanup EXIT

# --norc avoids ~/.bashrc clobbering PS1
PANE_SHELL="PS1='\$ ' exec bash --norc"

tmux new-session -d -s "$SESSION" "$PANE_SHELL"
for (( i=1; i<N; i++ )); do
    tmux split-window -t "$SESSION" "$PANE_SHELL"
    tmux select-layout -t "$SESSION" "$LAYOUT" >/dev/null
done

PREFIX="exec "
[[ "$WAIT" == "1" ]] && PREFIX=""

(
    sleep "$START_DELAY"
    for (( i=0; i<N; i++ )); do
        # without wait, exec replaces the pane's shell so the pane closes
        # when the command exits; with wait, the shell stays alive until the
        # timeout fires below
        tmux send-keys -t "${SESSION}.${i}" "${PREFIX}${CMDS[$i]}" Enter
    done
    sleep "$TIMEOUT"
    tmux kill-session -t "$SESSION" 2>/dev/null || true
) &
WATCHER=$!

asciinema rec --overwrite "$TMPCAST" -c "tmux attach -t $SESSION" || true
kill "$WATCHER" 2>/dev/null || true

mv "$TMPCAST" "$FINAL"
echo "saved: $FINAL"
