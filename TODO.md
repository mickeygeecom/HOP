# TODO

## Features
- Log/save how players are doing, but this would need some sort of player identifier
- Scheduele "live" quizzes with a lifetime/period (meaning differentiating live vs. persistant quizzes)
- Access Control, f.eg. "keys" or "tokens" in headers in the backend? when using accessing the management frontend
    -> Splitting up endpoints which require "auth", and which that does not.


## Bugs
[Low]
- Correct answer should not be send directly to the client prior to submitting an answer, this opens the opportunity for cheats
