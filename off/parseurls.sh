#!/bin/bash

# Usage:
# 1. Go to sc, copy the 'ul' (or whichever tag they have) that contains all the urls (in 'li's)
# 2. Run that through here
# 3. ???
# 4. Profit!!!
ag "href=\"\/[\w-]+\/[\w-]+\"" -o --nonumbers | uniq | sed 's/href//g' | sed 's/sc_likes_part2.html:="/https:\/\/soundcloud.com/g' | sed 's/"//g'