# OCR too ROB
Problem created by Porter-Gaud alumni Tillson Galloway from Georgia Tech's Greyhat team.

## Problem
I figure the cyber DMV uses this portal to validate my plates. Is there a way to sneak? I tried a test image to see if it parsed correctly:

[test image](test.png)

`ip:port`

## Solution
The website in which 'plates' are uploaded is vulnerable to server-side command injection. Navigating to the `robots.txt` reveals a test image in which the plate is correctly read (in static/uploads/test.png).

If you form a command injection on the plate (the one I used in the [sample solution](solution.png) was `; cat flag; #` ) you can get the flag. Note that to find the location of flag you need to run ls: `; ls; #`.
