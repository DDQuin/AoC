//Here's my approach:

// calculate distances between all valve pairs

// write a recursive function that tells me what the maximum value is that can be gained by opening remaining valves, with a certain time left, when I'm at a given valve, and when there's a given set of valves I may not open anymore (those are my method parameters). To calculate this value, I loop over all non-broken valves (15 in my input) that may still be opened. That's the one I'll move to and open next. Recurse on those. The calculations are straightforward (use the calculated distances) - just be careful with all the details.

// For me this ran in a few milliseconds (using C#), which was fast enough to do part 2 with a relatively brute force approach (consider all possible splits of 14 relevant valves between me and the elephant: about 16000...)

// This was by far the hardest puzzle for me, btw...