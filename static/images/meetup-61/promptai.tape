# silt without an idea, for three AI ideas in a row: a saturated space, a
# crowded one, an open one. Each new search starts with the last idea in the
# prompt; Ctrl+U clears it.

Output promptai.gif
Source settings.tape

# A live search, not the cache's answer, so that the progress screen shows.
Env SILT_CACHE_TTL "0"

Type "silt"
Sleep 400ms
Enter
Sleep 1.5s

Type "local llm that writes git commit messages from the staged diff"
Sleep 600ms
Enter
Wait+Screen@60s /Matches \(/
Sleep 4s

Type "n"
Sleep 800ms
Ctrl+U
Sleep 400ms
Type "mcp server that exposes a postgres database to an llm"
Sleep 600ms
Enter
Wait+Screen@60s /Matches \(/
Sleep 4s

Type "n"
Sleep 800ms
Ctrl+U
Sleep 400ms
Type "tui that diffs the answers of two local llms side by side"
Sleep 600ms
Enter
Wait+Screen@60s /Matches \(/
Sleep 5s

Type "q"
Sleep 1s
